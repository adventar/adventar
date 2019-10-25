import { Feed } from "feed";
import { getCalendar } from "~/lib/JsonApiClient";

async function generateCalendarFeed(calendarId: number): Promise<string> {
  const calendar = await getCalendar(calendarId);
  const feed = new Feed({
    id: "Adventar",
    title: `${calendar.title} Advent Calendar ${calendar.year}`,
    link: `https://adventar.org/calendars/${calendar.id}`,
    description: calendar.description,
    generator: "Adventar",
    updated: new Date(calendar.year, 11, calendar.entries && calendar.entries[0] ? calendar.entries[0].day : 1),
    copyright: ""
  });

  if (calendar && calendar.entries) {
    calendar.entries.reverse().forEach(entry => {
      if (!entry.url) return;
      const description = `${calendar.title} Advent Calendar ${calendar.year} ${entry.day}日目`;
      feed.addItem({
        guid: entry.id.toString(),
        title: (entry.title || entry.comment || description).trim(),
        description,
        link: entry.url,
        date: new Date(calendar.year, 11, entry.day)
      });
    });
  }

  return feed.rss2();
}

export { generateCalendarFeed };
