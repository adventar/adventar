import { Calendar } from "~/types/adventar";
import { API_ENDPOINT } from "~/lib/Configuration";

export async function getCalendar(id: number): Promise<Calendar> {
  const response = await fetch(`${API_ENDPOINT}/v1/calendars?calendar_id=${id}`);
  const body = await response.json();
  const calendar = body.calendar;
  calendar.entries = body.entries;
  return calendar;
}
