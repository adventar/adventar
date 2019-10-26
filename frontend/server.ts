import { Nuxt } from "nuxt";
import serverless from "serverless-http";
import express from "express";
import asyncHandler from "express-async-handler";
import bugsnag from "@bugsnag/js";
import config from "~/nuxt.config";
import { generateCalendarFeed, ExpiredCalendarError } from "~/server/Feed";
import { generateIcal } from "~/server/Ical";

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
    const feed = await generateCalendarFeed(calendarId);
    res.header("Content-Type", "application/rss+xml; charset=utf-8");
    res.send(feed);
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
  if (err instanceof ExpiredCalendarError) {
    res.status(400);
    return next(err);
  }

  if (err.response && err.response.status === 404) {
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
    console.log(`Listen: http://localhost:${port}`);
  });
}
