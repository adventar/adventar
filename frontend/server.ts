import { Nuxt } from "nuxt";
import serverless from "serverless-http";
import express from "express";
import config from "~/nuxt.config";
import { generateCalendarFeed } from "~/server/Feed";
import { generateIcal } from "~/server/Ical";

const app = express();
const nuxt = new Nuxt({
  ...config,
  dev: false,
  mode: "universal",
  buildDir: "../../.nuxt"
});

app.get("/calendars/:id.rss", async (req, res) => {
  const calendarId = Number(req.params.id);
  const feed = await generateCalendarFeed(calendarId);
  res.header["Content-Type"] = "application/rss+xml; charset=utf-8";
  res.send(feed);
});

app.get("/users/:id.ics", async (req, res) => {
  const userId = req.params.id;
  const ical = await generateIcal(userId);
  res.header["Content-Type"] = "text/calendar; charset=utf-8";
  res.send(ical);
});

app.get("/calendars/:id", async (req, res, next) => {
  await nuxt.ready();
  nuxt.render(req, res, next);
});

export const handler = serverless(app);
