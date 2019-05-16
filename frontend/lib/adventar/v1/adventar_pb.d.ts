export class Calendar {
  constructor ();
  getId(): number;
  setId(a: number): void;
  getOwner(): User;
  setOwner(a: User): void;
  getTitle(): string;
  setTitle(a: string): void;
  getDescription(): string;
  setDescription(a: string): void;
  getYear(): number;
  setYear(a: number): void;
  getEntryCount(): number;
  setEntryCount(a: number): void;
  toObject(): Calendar.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => Calendar;
}

export namespace Calendar {
  export type AsObject = {
    id: number;
    owner: User;
    title: string;
    description: string;
    year: number;
    entryCount: number;
  }
}

export class CreateCalendarRequest {
  constructor ();
  getTitle(): string;
  setTitle(a: string): void;
  getDescription(): string;
  setDescription(a: string): void;
  toObject(): CreateCalendarRequest.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => CreateCalendarRequest;
}

export namespace CreateCalendarRequest {
  export type AsObject = {
    title: string;
    description: string;
  }
}

export class CreateEntryRequest {
  constructor ();
  getCalendarId(): number;
  setCalendarId(a: number): void;
  getDay(): number;
  setDay(a: number): void;
  toObject(): CreateEntryRequest.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => CreateEntryRequest;
}

export namespace CreateEntryRequest {
  export type AsObject = {
    calendarId: number;
    day: number;
  }
}

export class DeleteCalendarRequest {
  constructor ();
  getCalendarId(): number;
  setCalendarId(a: number): void;
  toObject(): DeleteCalendarRequest.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => DeleteCalendarRequest;
}

export namespace DeleteCalendarRequest {
  export type AsObject = {
    calendarId: number;
  }
}

export class DeleteEntryRequest {
  constructor ();
  getEntryId(): number;
  setEntryId(a: number): void;
  toObject(): DeleteEntryRequest.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => DeleteEntryRequest;
}

export namespace DeleteEntryRequest {
  export type AsObject = {
    entryId: number;
  }
}

export class Entry {
  constructor ();
  getId(): number;
  setId(a: number): void;
  getOwner(): User;
  setOwner(a: User): void;
  getCalendar(): Calendar;
  setCalendar(a: Calendar): void;
  getDay(): number;
  setDay(a: number): void;
  getComment(): string;
  setComment(a: string): void;
  getUrl(): string;
  setUrl(a: string): void;
  getTitle(): string;
  setTitle(a: string): void;
  getImageUrl(): string;
  setImageUrl(a: string): void;
  toObject(): Entry.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => Entry;
}

export namespace Entry {
  export type AsObject = {
    id: number;
    owner: User;
    calendar: Calendar;
    day: number;
    comment: string;
    url: string;
    title: string;
    imageUrl: string;
  }
}

export class GetCalendarRequest {
  constructor ();
  getCalendarId(): number;
  setCalendarId(a: number): void;
  toObject(): GetCalendarRequest.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => GetCalendarRequest;
}

export namespace GetCalendarRequest {
  export type AsObject = {
    calendarId: number;
  }
}

export class GetCalendarResponse {
  constructor ();
  getCalendar(): Calendar;
  setCalendar(a: Calendar): void;
  getEntriesList(): Entry[];
  setEntriesList(a: Entry[]): void;
  toObject(): GetCalendarResponse.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => GetCalendarResponse;
}

export namespace GetCalendarResponse {
  export type AsObject = {
    calendar: Calendar;
    entriesList: Entry[];
  }
}

export class ListCalendarsRequest {
  constructor ();
  getYear(): number;
  setYear(a: number): void;
  getUserId(): number;
  setUserId(a: number): void;
  getQuery(): string;
  setQuery(a: string): void;
  getPageSize(): number;
  setPageSize(a: number): void;
  toObject(): ListCalendarsRequest.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => ListCalendarsRequest;
}

export namespace ListCalendarsRequest {
  export type AsObject = {
    year: number;
    userId: number;
    query: string;
    pageSize: number;
  }
}

export class ListCalendarsResponse {
  constructor ();
  getCalendarsList(): Calendar[];
  setCalendarsList(a: Calendar[]): void;
  toObject(): ListCalendarsResponse.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => ListCalendarsResponse;
}

export namespace ListCalendarsResponse {
  export type AsObject = {
    calendarsList: Calendar[];
  }
}

export class ListEntriesRequest {
  constructor ();
  getUserId(): number;
  setUserId(a: number): void;
  getYear(): number;
  setYear(a: number): void;
  toObject(): ListEntriesRequest.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => ListEntriesRequest;
}

export namespace ListEntriesRequest {
  export type AsObject = {
    userId: number;
    year: number;
  }
}

export class ListEntriesResponse {
  constructor ();
  getEntriesList(): Entry[];
  setEntriesList(a: Entry[]): void;
  toObject(): ListEntriesResponse.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => ListEntriesResponse;
}

export namespace ListEntriesResponse {
  export type AsObject = {
    entriesList: Entry[];
  }
}

export class SignInRequest {
  constructor ();
  getJwt(): string;
  setJwt(a: string): void;
  toObject(): SignInRequest.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => SignInRequest;
}

export namespace SignInRequest {
  export type AsObject = {
    jwt: string;
  }
}

export class UpdateCalendarRequest {
  constructor ();
  getCalendarId(): number;
  setCalendarId(a: number): void;
  getTitle(): string;
  setTitle(a: string): void;
  getDescription(): string;
  setDescription(a: string): void;
  toObject(): UpdateCalendarRequest.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => UpdateCalendarRequest;
}

export namespace UpdateCalendarRequest {
  export type AsObject = {
    calendarId: number;
    title: string;
    description: string;
  }
}

export class UpdateEntryRequest {
  constructor ();
  getEntryId(): number;
  setEntryId(a: number): void;
  getComment(): string;
  setComment(a: string): void;
  getUrl(): string;
  setUrl(a: string): void;
  toObject(): UpdateEntryRequest.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => UpdateEntryRequest;
}

export namespace UpdateEntryRequest {
  export type AsObject = {
    entryId: number;
    comment: string;
    url: string;
  }
}

export class UpdateUserRequest {
  constructor ();
  getName(): string;
  setName(a: string): void;
  toObject(): UpdateUserRequest.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => UpdateUserRequest;
}

export namespace UpdateUserRequest {
  export type AsObject = {
    name: string;
  }
}

export class User {
  constructor ();
  getId(): number;
  setId(a: number): void;
  getName(): string;
  setName(a: string): void;
  getIconUrl(): string;
  setIconUrl(a: string): void;
  toObject(): User.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => User;
}

export namespace User {
  export type AsObject = {
    id: number;
    name: string;
    iconUrl: string;
  }
}

export class Empty {
  constructor ();
  toObject(): Empty.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => Empty;
}

export namespace Empty {
  export type AsObject = {
  }
}

