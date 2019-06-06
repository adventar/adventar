import { SignInRequest, UpdateUserRequest } from "~/lib/grpc/adventar/v1/adventar_pb";
import { AdventarClient } from "~/lib/grpc/adventar/v1/adventar_grpc_web_pb";
const client = new AdventarClient("http://localhost:8000", null, null);

export type User = {
  id: number;
  name: string;
  iconUrl: string;
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
