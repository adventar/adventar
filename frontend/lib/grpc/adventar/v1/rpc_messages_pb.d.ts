import * as jspb from "google-protobuf"

import * as adventar_v1_resources_pb from '../../adventar/v1/resources_pb';

export class ListCalendarsRequest extends jspb.Message {
  getYear(): number;
  setYear(value: number): void;

  getUserId(): number;
  setUserId(value: number): void;

  getQuery(): string;
  setQuery(value: string): void;

  getPageSize(): number;
  setPageSize(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListCalendarsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListCalendarsRequest): ListCalendarsRequest.AsObject;
  static serializeBinaryToWriter(message: ListCalendarsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListCalendarsRequest;
  static deserializeBinaryFromReader(message: ListCalendarsRequest, reader: jspb.BinaryReader): ListCalendarsRequest;
}

export namespace ListCalendarsRequest {
  export type AsObject = {
    year: number,
    userId: number,
    query: string,
    pageSize: number,
  }
}

export class ListCalendarsResponse extends jspb.Message {
  getCalendarsList(): Array<adventar_v1_resources_pb.Calendar>;
  setCalendarsList(value: Array<adventar_v1_resources_pb.Calendar>): void;
  clearCalendarsList(): void;
  addCalendars(value?: adventar_v1_resources_pb.Calendar, index?: number): adventar_v1_resources_pb.Calendar;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListCalendarsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListCalendarsResponse): ListCalendarsResponse.AsObject;
  static serializeBinaryToWriter(message: ListCalendarsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListCalendarsResponse;
  static deserializeBinaryFromReader(message: ListCalendarsResponse, reader: jspb.BinaryReader): ListCalendarsResponse;
}

export namespace ListCalendarsResponse {
  export type AsObject = {
    calendarsList: Array<adventar_v1_resources_pb.Calendar.AsObject>,
  }
}

export class GetCalendarRequest extends jspb.Message {
  getCalendarId(): number;
  setCalendarId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCalendarRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetCalendarRequest): GetCalendarRequest.AsObject;
  static serializeBinaryToWriter(message: GetCalendarRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCalendarRequest;
  static deserializeBinaryFromReader(message: GetCalendarRequest, reader: jspb.BinaryReader): GetCalendarRequest;
}

export namespace GetCalendarRequest {
  export type AsObject = {
    calendarId: number,
  }
}

export class GetCalendarResponse extends jspb.Message {
  getCalendar(): adventar_v1_resources_pb.Calendar | undefined;
  setCalendar(value?: adventar_v1_resources_pb.Calendar): void;
  hasCalendar(): boolean;
  clearCalendar(): void;

  getEntriesList(): Array<adventar_v1_resources_pb.Entry>;
  setEntriesList(value: Array<adventar_v1_resources_pb.Entry>): void;
  clearEntriesList(): void;
  addEntries(value?: adventar_v1_resources_pb.Entry, index?: number): adventar_v1_resources_pb.Entry;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCalendarResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetCalendarResponse): GetCalendarResponse.AsObject;
  static serializeBinaryToWriter(message: GetCalendarResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCalendarResponse;
  static deserializeBinaryFromReader(message: GetCalendarResponse, reader: jspb.BinaryReader): GetCalendarResponse;
}

export namespace GetCalendarResponse {
  export type AsObject = {
    calendar?: adventar_v1_resources_pb.Calendar.AsObject,
    entriesList: Array<adventar_v1_resources_pb.Entry.AsObject>,
  }
}

export class CreateCalendarRequest extends jspb.Message {
  getTitle(): string;
  setTitle(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateCalendarRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateCalendarRequest): CreateCalendarRequest.AsObject;
  static serializeBinaryToWriter(message: CreateCalendarRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateCalendarRequest;
  static deserializeBinaryFromReader(message: CreateCalendarRequest, reader: jspb.BinaryReader): CreateCalendarRequest;
}

export namespace CreateCalendarRequest {
  export type AsObject = {
    title: string,
    description: string,
  }
}

export class UpdateCalendarRequest extends jspb.Message {
  getCalendarId(): number;
  setCalendarId(value: number): void;

  getTitle(): string;
  setTitle(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateCalendarRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateCalendarRequest): UpdateCalendarRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateCalendarRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateCalendarRequest;
  static deserializeBinaryFromReader(message: UpdateCalendarRequest, reader: jspb.BinaryReader): UpdateCalendarRequest;
}

export namespace UpdateCalendarRequest {
  export type AsObject = {
    calendarId: number,
    title: string,
    description: string,
  }
}

export class DeleteCalendarRequest extends jspb.Message {
  getCalendarId(): number;
  setCalendarId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteCalendarRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteCalendarRequest): DeleteCalendarRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteCalendarRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteCalendarRequest;
  static deserializeBinaryFromReader(message: DeleteCalendarRequest, reader: jspb.BinaryReader): DeleteCalendarRequest;
}

export namespace DeleteCalendarRequest {
  export type AsObject = {
    calendarId: number,
  }
}

export class ListEntriesRequest extends jspb.Message {
  getUserId(): number;
  setUserId(value: number): void;

  getYear(): number;
  setYear(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListEntriesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListEntriesRequest): ListEntriesRequest.AsObject;
  static serializeBinaryToWriter(message: ListEntriesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListEntriesRequest;
  static deserializeBinaryFromReader(message: ListEntriesRequest, reader: jspb.BinaryReader): ListEntriesRequest;
}

export namespace ListEntriesRequest {
  export type AsObject = {
    userId: number,
    year: number,
  }
}

export class ListEntriesResponse extends jspb.Message {
  getEntriesList(): Array<adventar_v1_resources_pb.Entry>;
  setEntriesList(value: Array<adventar_v1_resources_pb.Entry>): void;
  clearEntriesList(): void;
  addEntries(value?: adventar_v1_resources_pb.Entry, index?: number): adventar_v1_resources_pb.Entry;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListEntriesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListEntriesResponse): ListEntriesResponse.AsObject;
  static serializeBinaryToWriter(message: ListEntriesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListEntriesResponse;
  static deserializeBinaryFromReader(message: ListEntriesResponse, reader: jspb.BinaryReader): ListEntriesResponse;
}

export namespace ListEntriesResponse {
  export type AsObject = {
    entriesList: Array<adventar_v1_resources_pb.Entry.AsObject>,
  }
}

export class CreateEntryRequest extends jspb.Message {
  getCalendarId(): number;
  setCalendarId(value: number): void;

  getDay(): number;
  setDay(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateEntryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateEntryRequest): CreateEntryRequest.AsObject;
  static serializeBinaryToWriter(message: CreateEntryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateEntryRequest;
  static deserializeBinaryFromReader(message: CreateEntryRequest, reader: jspb.BinaryReader): CreateEntryRequest;
}

export namespace CreateEntryRequest {
  export type AsObject = {
    calendarId: number,
    day: number,
  }
}

export class UpdateEntryRequest extends jspb.Message {
  getEntryId(): number;
  setEntryId(value: number): void;

  getComment(): string;
  setComment(value: string): void;

  getUrl(): string;
  setUrl(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateEntryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateEntryRequest): UpdateEntryRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateEntryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateEntryRequest;
  static deserializeBinaryFromReader(message: UpdateEntryRequest, reader: jspb.BinaryReader): UpdateEntryRequest;
}

export namespace UpdateEntryRequest {
  export type AsObject = {
    entryId: number,
    comment: string,
    url: string,
  }
}

export class DeleteEntryRequest extends jspb.Message {
  getEntryId(): number;
  setEntryId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteEntryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteEntryRequest): DeleteEntryRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteEntryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteEntryRequest;
  static deserializeBinaryFromReader(message: DeleteEntryRequest, reader: jspb.BinaryReader): DeleteEntryRequest;
}

export namespace DeleteEntryRequest {
  export type AsObject = {
    entryId: number,
  }
}

export class SignInRequest extends jspb.Message {
  getJwt(): string;
  setJwt(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SignInRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SignInRequest): SignInRequest.AsObject;
  static serializeBinaryToWriter(message: SignInRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SignInRequest;
  static deserializeBinaryFromReader(message: SignInRequest, reader: jspb.BinaryReader): SignInRequest;
}

export namespace SignInRequest {
  export type AsObject = {
    jwt: string,
  }
}

export class GetUserRequest extends jspb.Message {
  getUserId(): number;
  setUserId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserRequest): GetUserRequest.AsObject;
  static serializeBinaryToWriter(message: GetUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserRequest;
  static deserializeBinaryFromReader(message: GetUserRequest, reader: jspb.BinaryReader): GetUserRequest;
}

export namespace GetUserRequest {
  export type AsObject = {
    userId: number,
  }
}

export class UpdateUserRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateUserRequest): UpdateUserRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateUserRequest;
  static deserializeBinaryFromReader(message: UpdateUserRequest, reader: jspb.BinaryReader): UpdateUserRequest;
}

export namespace UpdateUserRequest {
  export type AsObject = {
    name: string,
  }
}

