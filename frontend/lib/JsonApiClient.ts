import fetch, { Response } from "node-fetch";
import { Calendar, Entry } from "~/types/adventar";

export async function getCalendar(calendarId: number): Promise<Calendar> {
  const resBody = await request("GetCalendar", { calendar_id: calendarId });
  const calendar = resBody.calendar;
  calendar.entries = resBody.entries;

  return calendar;
}

export async function listEntries(userId: number): Promise<Entry[]> {
  const resBody = await request("ListEntries", { user_id: userId });

  return resBody.entries;
}

async function request(rpcName: string, body: Record<string, any>) {
  const response = await fetch(`${process.env.API_BASE_URL}/adventar.v1.Adventar/${rpcName}`, {
    method: "POST",
    body: JSON.stringify(body)
  });
  if (!response.ok) {
    throw new ApiError(response);
  }

  return response.json();
}

export class ApiError extends Error {
  constructor(public response: Response) {
    super(`API request failed: ${response.status}`);
  }
}
