import { Calendar } from "~/types/adventar";

export async function getCalendar(id: number): Promise<Calendar> {
  const response = await fetch(`${process.env.apiBaseUrl}/v1/calendars?calendar_id=${id}`);
  const body = await response.json();
  const calendar = body.calendar;
  calendar.entries = body.entries;
  return calendar;
}
