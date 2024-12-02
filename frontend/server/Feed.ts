import { Feed } from "feed";
import { getCalendar } from "~/lib/JsonApiClient";
import { getToday } from "~/lib/Configuration";

async function generateCalendarFeed(calendarId: number): Promise<{ feed: string; cacheable: boolean }> {
  const calendar = await getCalendar(calendarId);
  const today = getToday();
  // 前年より前のカレンダーは更新されることがないのでcacheする。slackbot対策。
  const cacheable = calendar.year < today.getFullYear();

  const feed = new Feed({
    id: "Adventar",
    title: `${calendar.title} Advent Calendar ${calendar.year}`,
    // slackbot が link の URL にアクセスしてそうなので一回消す
    // link: `https://adventar.org/calendars/${calendar.id}`,
    description: calendar.description,
    generator: "Adventar",
    updated: new Date(calendar.year, 11, calendar.entries && calendar.entries[0] ? calendar.entries[0].day : 1),
    copyright: ""
  });

  if (calendar && calendar.entries) {
    calendar.entries.reverse().forEach(entry => {
      if (!entry.url) return;
      if (entry.day > today.getDate()) return;
      const description = `${calendar.title} Advent Calendar ${calendar.year} ${entry.day}日目`;
      feed.addItem({
        guid: entry.id.toString(),
        title: (entry.title || entry.comment || description).trim(),
        description,
        link: entry.url,
        date: new Date(calendar.year, 11, entry.day),
        extra: {
          "dc:creator": entry.owner ? entry.owner.name : ""
        }
      });
    });
  }

  // https://github.com/adventar/adventar/issues/52#issuecomment-2509769282
  const rss = feed
    .rss2()
    .replace('<rss version="2.0">', '<rss version="2.0" xmlns:dc="http://purl.org/dc/elements/1.1/">');

  return { feed: rss, cacheable };
}

export { generateCalendarFeed };
