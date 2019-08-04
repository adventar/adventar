import { Nuxt } from "nuxt";
import serverless from "serverless-http";
import express from "express";
import config from "~/nuxt.config";
import { generateCalendarFeed } from "~/server/Feed";

const app = express();
const nuxt = new Nuxt({
  ...config,
  dev: false,
  mode: "universal",
  buildDir: "../../.nuxt"
});

app.get("/calendars/:id.rss", async (req, res) => {
  const id = Number(req.params.id);
  const feed = await generateCalendarFeed(id);
  res.header["Content-Type"] = "application/rss+xml; charset=utf-8";
  res.send(feed);
});

app.get("/users/:id.ical", (req, res) => {
  const id = req.params.id;
  console.log(id);
  res.send("ok");
});

app.get("/calendars/:id", async (req, res, next) => {
  await nuxt.ready();
  nuxt.render(req, res, next);
});

export const handler = serverless(app);
