import {
  SignInRequest,
  UpdateUserRequest,
  ListCalendarsRequest,
  GetCalendarRequest,
  CreateCalendarRequest,
  UpdateCalendarRequest,
  DeleteCalendarRequest,
  ListEntriesRequest,
  CreateEntryRequest,
  UpdateEntryRequest,
  DeleteEntryRequest,
  GetUserRequest,
  ListCalendarStatsRequest
} from "~/lib/grpc/adventar/v1/adventar_pb";
import { AdventarPromiseClient } from "~/lib/grpc/adventar/v1/adventar_grpc_web_pb";
import { User, Calendar, Entry, CalendarStat } from "~/types/adventar";

const client = new AdventarPromiseClient(process.env.API_BASE_URL || "", null, null);

export async function signIn(token: string, iconUrl: string): Promise<User> {
  const request = new SignInRequest();
  request.setJwt(token);
  request.setIconUrl(iconUrl);

  const user = await client.signIn(request, {});

  return {
    id: user.getId(),
    name: user.getName(),
    iconUrl: user.getIconUrl()
  };
}

export async function updateUser(name: string, token: string): Promise<User> {
  const request = new UpdateUserRequest();
  request.setName(name);

  const user = await client.updateUser(request, { authorization: token });

  return {
    id: user.getId(),
    name: user.getName(),
    iconUrl: user.getIconUrl()
  };
}

type createCalendarParams = {
  title: string;
  description: string;
  token: string;
};
export async function createCalendar({ title, description, token }: createCalendarParams): Promise<number> {
  const request = new CreateCalendarRequest();
  request.setTitle(title);
  request.setDescription(description);

  const calendar = await client.createCalendar(request, { authorization: token });

  return calendar.getId();
}

type updateCalendarParams = {
  id: number;
  title: string;
  description: string;
  token: string;
};
export async function updateCalendar({ id, title, description, token }: updateCalendarParams): Promise<void> {
  const request = new UpdateCalendarRequest();
  request.setCalendarId(id);
  request.setTitle(title);
  request.setDescription(description);

  await client.updateCalendar(request, { authorization: token });
}

type deleteCalendarParams = {
  id: number;
  token: string;
};
export async function deleteCalendar({ id, token }: deleteCalendarParams): Promise<void> {
  const request = new DeleteCalendarRequest();
  request.setCalendarId(id);

  await client.deleteCalendar(request, { authorization: token });
}

export async function getCalendar(id: number): Promise<Calendar> {
  const request = new GetCalendarRequest();
  request.setCalendarId(id);

  const response = await client.getCalendar(request, {});
  const calendar = response.getCalendar();
  const entries = response.getEntriesList();

  return {
    id: calendar.getId(),
    title: calendar.getTitle(),
    description: calendar.getDescription(),
    year: calendar.getYear(),
    owner: {
      id: calendar.getOwner().getId(),
      name: calendar.getOwner().getName(),
      iconUrl: calendar.getOwner().getIconUrl()
    },
    entryCount: calendar.getEntryCount(),
    entries: entries.map(entry => {
      return {
        id: entry.getId(),
        owner: {
          id: entry.getOwner().getId(),
          name: entry.getOwner().getName(),
          iconUrl: entry.getOwner().getIconUrl()
        },
        day: entry.getDay(),
        comment: entry.getComment(),
        url: entry.getUrl(),
        title: entry.getTitle(),
        imageUrl: entry.getImageUrl()
      };
    })
  };
}

interface listCalendarsParams {
  readonly year: number;
  readonly userId?: number;
  readonly pageSize?: number;
  readonly query?: string;
}
export async function listCalendars({ year, userId, pageSize, query }: listCalendarsParams): Promise<Calendar[]> {
  const request = new ListCalendarsRequest();
  request.setYear(year);
  request.setPageSize(pageSize || 0);
  request.setUserId(userId || 0);
  request.setQuery(query || "");

  const response = await client.listCalendars(request, {});

  return response.getCalendarsList().map(calendar => {
    return {
      id: calendar.getId(),
      title: calendar.getTitle(),
      description: calendar.getDescription(),
      year: calendar.getYear(),
      owner: {
        id: calendar.getOwner().getId(),
        name: calendar.getOwner().getName(),
        iconUrl: calendar.getOwner().getIconUrl()
      },
      entryCount: calendar.getEntryCount()
    };
  });
}

type listEntriesParams = {
  year: number;
  userId: number;
};
export async function listEntries({ year, userId }: listEntriesParams): Promise<Entry[]> {
  const request = new ListEntriesRequest();
  request.setYear(year);
  request.setUserId(userId);
  const response = await client.listEntries(request, {});

  return response.getEntriesList().map(entry => {
    return {
      id: entry.getId(),
      owner: {
        id: entry.getOwner().getId(),
        name: entry.getOwner().getName(),
        iconUrl: entry.getOwner().getIconUrl()
      },
      calendar: {
        id: entry.getCalendar().getId(),
        title: entry.getCalendar().getTitle(),
        year: entry.getCalendar().getYear()
      },
      day: entry.getDay(),
      comment: entry.getComment(),
      url: entry.getUrl(),
      title: entry.getTitle(),
      imageUrl: entry.getImageUrl()
    };
  });
}

type createEntryParams = {
  calendarId: number;
  day: number;
  token: string;
};
export async function createEntry({ calendarId, day, token }: createEntryParams): Promise<Entry> {
  const request = new CreateEntryRequest();
  request.setCalendarId(calendarId);
  request.setDay(day);

  const entry = await client.createEntry(request, { authorization: token });

  return {
    id: entry.getId(),
    day,
    comment: entry.getComment(),
    url: entry.getUrl(),
    title: entry.getTitle(),
    imageUrl: entry.getImageUrl()
  };
}

type updateEntryParams = {
  entryId: number;
  comment: string;
  url: string;
  token: string;
};
export async function updateEntry({ entryId, comment, url, token }: updateEntryParams): Promise<number> {
  const request = new UpdateEntryRequest();
  request.setEntryId(entryId);
  request.setComment(comment);
  request.setUrl(url);

  const entry = await client.updateEntry(request, { authorization: token });

  return entry.getId();
}

type deleteEntryParams = {
  entryId: number;
  token: string;
};
export async function deleteEntry({ entryId, token }: deleteEntryParams): Promise<void> {
  const request = new DeleteEntryRequest();
  request.setEntryId(entryId);

  await client.deleteEntry(request, { authorization: token });
}

export async function getUser(id: number): Promise<User> {
  const request = new GetUserRequest();
  request.setUserId(id);

  const user = await client.getUser(request, {});

  return {
    id: user.getId(),
    name: user.getName(),
    iconUrl: user.getIconUrl()
  };
}

export async function listCalendarStats(): Promise<CalendarStat[]> {
  const request = new ListCalendarStatsRequest();
  const res = await client.listCalendarStats(request, {});

  return res.getCalendarStatsList().map(s => {
    return {
      year: s.getYear(),
      calendarsCount: s.getCalendarsCount(),
      entriesCount: s.getEntriesCount()
    };
  });
}
