import * as jspb from "google-protobuf"

export class User extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getName(): string;
  setName(value: string): void;

  getIconUrl(): string;
  setIconUrl(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id: number,
    name: string,
    iconUrl: string,
  }
}

export class Calendar extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getOwner(): User | undefined;
  setOwner(value?: User): void;
  hasOwner(): boolean;
  clearOwner(): void;

  getTitle(): string;
  setTitle(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  getYear(): number;
  setYear(value: number): void;

  getEntryCount(): number;
  setEntryCount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Calendar.AsObject;
  static toObject(includeInstance: boolean, msg: Calendar): Calendar.AsObject;
  static serializeBinaryToWriter(message: Calendar, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Calendar;
  static deserializeBinaryFromReader(message: Calendar, reader: jspb.BinaryReader): Calendar;
}

export namespace Calendar {
  export type AsObject = {
    id: number,
    owner?: User.AsObject,
    title: string,
    description: string,
    year: number,
    entryCount: number,
  }
}

export class Entry extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getOwner(): User | undefined;
  setOwner(value?: User): void;
  hasOwner(): boolean;
  clearOwner(): void;

  getCalendar(): Calendar | undefined;
  setCalendar(value?: Calendar): void;
  hasCalendar(): boolean;
  clearCalendar(): void;

  getDay(): number;
  setDay(value: number): void;

  getComment(): string;
  setComment(value: string): void;

  getUrl(): string;
  setUrl(value: string): void;

  getTitle(): string;
  setTitle(value: string): void;

  getImageUrl(): string;
  setImageUrl(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Entry.AsObject;
  static toObject(includeInstance: boolean, msg: Entry): Entry.AsObject;
  static serializeBinaryToWriter(message: Entry, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Entry;
  static deserializeBinaryFromReader(message: Entry, reader: jspb.BinaryReader): Entry;
}

export namespace Entry {
  export type AsObject = {
    id: number,
    owner?: User.AsObject,
    calendar?: Calendar.AsObject,
    day: number,
    comment: string,
    url: string,
    title: string,
    imageUrl: string,
  }
}

