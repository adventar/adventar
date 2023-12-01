import { getCalendar } from "~/lib/JsonApiClient";

type OembedResponse = {
  version: string;
  width: "100%";
  height: number;
  type: "rich";
  provider_name: "Adventar";
  provider_url: "https://adventar.org";
  url: string;
};

const ROW_HEIGHT = 75;
const BASE_HEIGHT = 362;

export async function generateOembed(calendarId: number): Promise<OembedResponse> {
  const calendar = await getCalendar(calendarId);
  const height = isFiveWeeks(calendar.year) ? BASE_HEIGHT + ROW_HEIGHT : BASE_HEIGHT;
  return {
    version: "1.0",
    width: "100%",
    height,
    type: "rich",
    provider_name: "Adventar",
    provider_url: "https://adventar.org",
    url: `https://adventar.org/calendars/${calendarId}/embed`
  };
}

// 12月1日が木曜、金曜、土曜であれば25日までのカレンダーが5行になる
function isFiveWeeks(year: number): boolean {
  const day = new Date(year, 11, 1).getDay();
  return day === 4 || day === 5 || day === 6;
}
