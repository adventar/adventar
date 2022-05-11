import { Nuxt } from "nuxt";
import serverless from "serverless-http";
import express from "express";
import asyncHandler from "express-async-handler";
import bugsnag from "@bugsnag/js";
import config from "~/nuxt.config";
import { generateCalendarFeed } from "~/server/Feed";
import { generateIcal } from "~/server/Ical";
import { ApiError } from "~/lib/JsonApiClient";
import url from "url";

const bugsnagClient = bugsnag(process.env.BUGSNAG_API_KEY || "");

const app = express();
const nuxt = new Nuxt({
  ...config,
  dev: false,
  mode: "universal",
  buildDir: ".nuxt-server"
});

app.get(
  "/calendars/:id.rss",
  asyncHandler(async (req, res) => {
    const calendarId = Number(req.params.id);
    const { feed, cacheable } = await generateCalendarFeed(calendarId);
    if (cacheable) {
      res.header("Cache-Control", "max-age=31536000");
    }
    res.header("Content-Type", "application/rss+xml; charset=utf-8");
    res.send(feed);
  })
);

app.get(
  "/oembed",
  asyncHandler(async (req, res) => {
    const u = req.query.url;
    if (!u) return res.status(400).send("url is required");
    const { pathname } = url.parse(u);
    const calendarId = pathname && Number(pathname.replace(/\/calendars\/(\d+)$/, "$1"));
    // 火曜スタートであればその年は4週目まで、そうでなければ5週目まである
    // const rowCount = new Date(this.calendar.year, 12, 1).getDay() <= 2 ? 4 : 5;
    // const cellHeight = 63;
    // const headerHeight = 92;
    // const height = headerHeight + cellHeight * rowCount;
    res.json({
      version: "1.0",
      width: "100%",
      height: 362,
      type: "rich",
      provider_name: "Adventar",
      provider_url: "https://adventar.org",
      url: `https://adventar.org/calendars/${calendarId}/embed`
    });
  })
);

app.get(
  "/users/:id.ics",
  asyncHandler(async (req, res) => {
    const userId = req.params.id;
    const ical = await generateIcal(userId);
    res.header("Content-Type", "text/calendar; charset=utf-8");
    res.send(ical);
  })
);

app.get(
  "/calendars/:id",
  asyncHandler(async (req, res, next) => {
    await nuxt.ready();
    nuxt.render(req, res, next);
  })
);

app.use((err, req, res, next) => {
  if (err instanceof ApiError && err.response.status === 404) {
    res.status(404);
    return next(err);
  }

  const opt = {
    request: {
      headers: req.headers,
      httpMethod: req.method,
      url: req.url
    }
  };
  bugsnagClient.notify(err, opt, () => {
    next(err);
  });
});

export const handler = serverless(app);

if (process.env.RUN_LOCAL) {
  const port = process.env.PORT || 3030;
  app.listen(port, () => {
    console.log(`Listen: http://localhost:${port}`); // eslint-disable-line no-console
  });
}
