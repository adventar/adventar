import { Feed } from "feed";
import { getCalendar } from "~/lib/RestClient";

async function generateCalendarFeed(calendarId: number): Promise<string> {
  const calendar = await getCalendar(calendarId);
  const feed = new Feed({
    id: "Adventar",
    title: `${calendar.title} Advent Calendar ${calendar.year}`,
    link: `https://adventar.org/calendars/${calendar.id}`,
    description: calendar.description,
    copyright: ""
  });

  return feed.rss2();
}

export { generateCalendarFeed };
