import fetch from "node-fetch";
import { Calendar, Entry } from "~/types/adventar";

export async function getCalendar(calendarId: number): Promise<Calendar> {
  const resBody = await request("GetCalendar", { calendar_id: calendarId });
  const calendar = resBody.calendar;
  calendar.entries = resBody.entries;

  return calendar;
}

export async function listEntries(userId: number): Promise<Entry[]> {
  const bodyBody = await request("ListEntries", { user_Id: userId });

  return bodyBody.entries;
}

async function request(rpcName: string, body: Record<string, any>) {
  const response = await fetch(`${process.env.API_BASE_URL}/adventar.v1.Adventar/${rpcName}`, {
    method: "POST",
    body: JSON.stringify(body)
  });
  if (!response.ok) {
    throw new Error(response.status);
  }

  return response.json();
}
