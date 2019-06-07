import { SignInRequest, UpdateUserRequest, GetCalendarRequest, CreateCalendarRequest } from "~/lib/grpc/adventar/v1/adventar_pb";
import { AdventarClient } from "~/lib/grpc/adventar/v1/adventar_grpc_web_pb";
const client = new AdventarClient("http://localhost:8000", null, null);

export type User = {
  id: number;
  name: string;
  iconUrl: string;
};

export type Calendar = {
  id: number;
  title: string;
  description: string;
};

export function signIn(token: string): Promise<User> {
  const request = new SignInRequest();
  request.setJwt(token);

  return new Promise((resolve, reject) => {
    client.signIn(request, {}, (err, res) => {
      if (err) {
        reject(err);
      }
      else {
        resolve({
          id: res.getId(),
          name: res.getName(),
          iconUrl: res.getIconUrl(),
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
      }
      else {
        resolve({
          id: res.getId(),
          name: res.getName(),
          iconUrl: res.getIconUrl(),
        });
      }
    });
  });
}
type createCalendarParams = {
  title: string;
  description: string;
  token: string;
}
export function createCalendar({ title, description, token }: createCalendarParams): Promise<number> {
  const request = new CreateCalendarRequest();
  request.setTitle(title);
  request.setDescription(description);

  return new Promise((resolve, reject) => {
    client.createCalendar(request, { authorization: token }, (err, res) => {
      if (err) {
        reject(err);
      }
      else {
        resolve(res.getId());
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
      }
      else {
        const calendar = res.getCalendar();
        resolve({
          id: calendar.getId(),
          title: calendar.getTitle(),
          description: calendar.getDescription(),
        });
      }
    });
  });
}
