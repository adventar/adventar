import ical from "ical-generator";
import { listEntries } from "~/lib/JsonApiClient";

async function generateIcal(userId: number): Promise<string> {
  const entries = await listEntries(userId);
  const events = entries.map(e => {
    const calendar = e.calendar!;
    return {
      summary: `${calendar.title} Advent Calendar ${calendar.year}`,
      start: new Date(calendar.year, 11, e.day),
      end: new Date(calendar.year, 11, e.day)
    };
  });
  const cal = ical({
    name: "Adventar",
    domain: "adventar.org",
    prodId: { company: "adventar", product: "ical-generator", language: "JA" },
    timezone: "Asia/Tokyo",
    events
  });
  return cal.toString();
}

export { generateIcal };
