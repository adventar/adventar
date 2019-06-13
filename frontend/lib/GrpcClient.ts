import {
  SignInRequest,
  UpdateUserRequest,
  GetCalendarRequest,
  CreateCalendarRequest,
  UpdateCalendarRequest,
  DeleteCalendarRequest,
  ListCalendarsRequest,
  CreateEntryRequest,
  DeleteEntryRequest,
  GetUserRequest
} from "~/lib/grpc/adventar/v1/adventar_pb";
import { AdventarClient } from "~/lib/grpc/adventar/v1/adventar_grpc_web_pb";
import { User, Calendar, Entry } from "~/types/adventar";
const client = new AdventarClient("http://localhost:8000", null, null);

export function signIn(token: string): Promise<User> {
  const request = new SignInRequest();
  request.setJwt(token);

  return new Promise((resolve, reject) => {
    client.signIn(request, {}, (err, res) => {
      if (err) {
        reject(err);
      } else {
        resolve({
          id: res.getId(),
          name: res.getName(),
          iconUrl: res.getIconUrl()
        });
      }
    });
  });
}

export function updateUser(name: string, token: string): Promise<User> {
  const request = new UpdateUserRequest();
  request.setName(name);

  return new Promise((resolve, reject) => {
    client.updateUser(request, { authorization: token }, (err, res) => {
      if (err) {
        reject(err);
      } else {
        resolve({
          id: res.getId(),
          name: res.getName(),
          iconUrl: res.getIconUrl()
        });
      }
    });
  });
}

type createCalendarParams = {
  title: string;
  description: string;
  token: string;
};
export function createCalendar({ title, description, token }: createCalendarParams): Promise<number> {
  const request = new CreateCalendarRequest();
  request.setTitle(title);
  request.setDescription(description);

  return new Promise((resolve, reject) => {
    client.createCalendar(request, { authorization: token }, (err, res) => {
      if (err) {
        reject(err);
      } else {
        resolve(res.getId());
      }
    });
  });
}

type updateCalendarParams = {
  id: number;
  title: string;
  description: string;
  token: string;
};
export function updateCalendar({ id, title, description, token }: updateCalendarParams): Promise<void> {
  const request = new UpdateCalendarRequest();
  request.setCalendarId(id);
  request.setTitle(title);
  request.setDescription(description);

  return new Promise((resolve, reject) => {
    client.updateCalendar(request, { authorization: token }, err => {
      if (err) {
        reject(err);
      } else {
        resolve();
      }
    });
  });
}

type deleteCalendarParams = {
  id: number;
  token: string;
};
export function deleteCalendar({ id, token }: deleteCalendarParams): Promise<void> {
  const request = new DeleteCalendarRequest();
  request.setCalendarId(id);

  return new Promise((resolve, reject) => {
    client.deleteCalendar(request, { authorization: token }, err => {
      if (err) {
        reject(err);
      } else {
        resolve();
      }
    });
  });
}

export function getCalendar(id: number): Promise<Calendar> {
  const request = new GetCalendarRequest();
  request.setCalendarId(id);

  return new Promise((resolve, reject) => {
    client.getCalendar(request, {}, (err, res) => {
      if (err) {
        reject(err);
      } else {
        const calendar = res.getCalendar();
        resolve({
          id: calendar.getId(),
          title: calendar.getTitle(),
          description: calendar.getDescription(),
          year: calendar.getYear(),
          entryCount: calendar.getEntryCount(),
          entries: res.getEntriesList().map(entry => {
            return {
              id: entry.getId(),
              owner: {
                id: entry.getOwner().getId(),
                name: entry.getOwner().getName(),
                iconUrl: entry.getOwner().getIconUrl()
              },
              day: entry.getDay()
            };
          })
        });
      }
    });
  });
}

interface listCalendarsParams {
  readonly year: number;
  readonly pageSize?: number;
  readonly query?: string;
}
export function listCalendar({ year, pageSize, query }: listCalendarsParams): Promise<Calendar[]> {
  const request = new ListCalendarsRequest();
  request.setYear(year);
  request.setPageSize(pageSize || 0);
  request.setQuery(query || "");

  return new Promise((resolve, reject) => {
    client.listCalendars(request, {}, (err, res) => {
      if (err) {
        reject(err);
      } else {
        const calendars = res.getCalendarsList().map(calendar => {
          return {
            id: calendar.getId(),
            title: calendar.getTitle(),
            description: calendar.getDescription(),
            year: calendar.getYear(),
            entryCount: calendar.getEntryCount()
          };
        });
        resolve(calendars);
      }
    });
  });
}

type createEntryParams = {
  calendarId: number;
  day: number;
  token: string;
};
export function createEntry({ calendarId, day, token }: createEntryParams): Promise<Entry> {
  const request = new CreateEntryRequest();
  request.setCalendarId(calendarId);
  request.setDay(day);

  return new Promise((resolve, reject) => {
    client.createEntry(request, { authorization: token }, (err, res) => {
      if (err) {
        reject(err);
      } else {
        resolve({ id: res.getId() });
      }
    });
  });
}

type deleteEntryParams = {
  entryId: number;
  token: string;
};
export function deleteEntry({ entryId, token }: deleteEntryParams): Promise<void> {
  const request = new DeleteEntryRequest();
  request.setEntryId(entryId);

  return new Promise((resolve, reject) => {
    client.deleteEntry(request, { authorization: token }, err => {
      if (err) {
        reject(err);
      } else {
        resolve();
      }
    });
  });
}

export function getUser(id: number): Promise<User> {
  const request = new GetUserRequest();
  request.setUserId(id);

  return new Promise((resolve, reject) => {
    client.getUser(request, {}, (err, res) => {
      if (err) {
        reject(err);
      } else {
        resolve({
          id: res.getId(),
          name: res.getName(),
          iconUrl: res.getIconUrl()
        });
      }
    });
  });
}
