import * as grpcWeb from 'grpc-web';
import {
  Calendar,
  CreateCalendarRequest,
  CreateEntryRequest,
  DeleteCalendarRequest,
  DeleteEntryRequest,
  Entry,
  GetCalendarRequest,
  GetCalendarResponse,
  GetUserRequest,
  ListCalendarsRequest,
  ListCalendarsResponse,
  ListEntriesRequest,
  ListEntriesResponse,
  SignInRequest,
  UpdateCalendarRequest,
  UpdateEntryRequest,
  UpdateUserRequest,
  User,
  Empty} from './adventar_pb';

export class AdventarClient {
  constructor (hostname: string,
               credentials: null | { [index: string]: string; },
               options: null | { [index: string]: string; });

  listCalendars(
    request: ListCalendarsRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: ListCalendarsResponse) => void
  ): grpcWeb.ClientReadableStream<ListCalendarsResponse>;

  getCalendar(
    request: GetCalendarRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: GetCalendarResponse) => void
  ): grpcWeb.ClientReadableStream<GetCalendarResponse>;

  createCalendar(
    request: CreateCalendarRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: Calendar) => void
  ): grpcWeb.ClientReadableStream<Calendar>;

  updateCalendar(
    request: UpdateCalendarRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: Calendar) => void
  ): grpcWeb.ClientReadableStream<Calendar>;

  deleteCalendar(
    request: DeleteCalendarRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: Empty) => void
  ): grpcWeb.ClientReadableStream<Empty>;

  listEntries(
    request: ListEntriesRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: ListEntriesResponse) => void
  ): grpcWeb.ClientReadableStream<ListEntriesResponse>;

  createEntry(
    request: CreateEntryRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: Entry) => void
  ): grpcWeb.ClientReadableStream<Entry>;

  updateEntry(
    request: UpdateEntryRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: Entry) => void
  ): grpcWeb.ClientReadableStream<Entry>;

  deleteEntry(
    request: DeleteEntryRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: Empty) => void
  ): grpcWeb.ClientReadableStream<Empty>;

  signIn(
    request: SignInRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: User) => void
  ): grpcWeb.ClientReadableStream<User>;

  getUser(
    request: GetUserRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: User) => void
  ): grpcWeb.ClientReadableStream<User>;

  updateUser(
    request: UpdateUserRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: User) => void
  ): grpcWeb.ClientReadableStream<User>;

}

export class AdventarPromiseClient {
  constructor (hostname: string,
               credentials: null | { [index: string]: string; },
               options: null | { [index: string]: string; });

  listCalendars(
    request: ListCalendarsRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<ListCalendarsResponse>;

  getCalendar(
    request: GetCalendarRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<GetCalendarResponse>;

  createCalendar(
    request: CreateCalendarRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<Calendar>;

  updateCalendar(
    request: UpdateCalendarRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<Calendar>;

  deleteCalendar(
    request: DeleteCalendarRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<Empty>;

  listEntries(
    request: ListEntriesRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<ListEntriesResponse>;

  createEntry(
    request: CreateEntryRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<Entry>;

  updateEntry(
    request: UpdateEntryRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<Entry>;

  deleteEntry(
    request: DeleteEntryRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<Empty>;

  signIn(
    request: SignInRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<User>;

  getUser(
    request: GetUserRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<User>;

  updateUser(
    request: UpdateUserRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<User>;

}

