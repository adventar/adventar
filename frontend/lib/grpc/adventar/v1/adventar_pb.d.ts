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
    Id: number;
    Owner: User;
    Title: string;
    Description: string;
    Year: number;
    EntryCount: number;
  }
}

export class CalendarStat {
  constructor ();
  getYear(): number;
  setYear(a: number): void;
  getCalendarsCount(): number;
  setCalendarsCount(a: number): void;
  getEntriesCount(): number;
  setEntriesCount(a: number): void;
  toObject(): CalendarStat.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => CalendarStat;
}

export namespace CalendarStat {
  export type AsObject = {
    Year: number;
    CalendarsCount: number;
    EntriesCount: number;
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
    Title: string;
    Description: string;
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
    CalendarId: number;
    Day: number;
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
    CalendarId: number;
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
    EntryId: number;
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
    Id: number;
    Owner: User;
    Calendar: Calendar;
    Day: number;
    Comment: string;
    Url: string;
    Title: string;
    ImageUrl: string;
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
    CalendarId: number;
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
    Calendar: Calendar;
    EntriesList: Entry[];
  }
}

export class GetUserRequest {
  constructor ();
  getUserId(): number;
  setUserId(a: number): void;
  toObject(): GetUserRequest.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => GetUserRequest;
}

export namespace GetUserRequest {
  export type AsObject = {
    UserId: number;
  }
}

export class ListCalendarStatsRequest {
  constructor ();
  toObject(): ListCalendarStatsRequest.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => ListCalendarStatsRequest;
}

export namespace ListCalendarStatsRequest {
  export type AsObject = {
  }
}

export class ListCalendarStatsResponse {
  constructor ();
  getCalendarStatsList(): CalendarStat[];
  setCalendarStatsList(a: CalendarStat[]): void;
  toObject(): ListCalendarStatsResponse.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => ListCalendarStatsResponse;
}

export namespace ListCalendarStatsResponse {
  export type AsObject = {
    CalendarStatsList: CalendarStat[];
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
    Year: number;
    UserId: number;
    Query: string;
    PageSize: number;
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
    CalendarsList: Calendar[];
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
    UserId: number;
    Year: number;
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
    EntriesList: Entry[];
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
    Jwt: string;
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
    CalendarId: number;
    Title: string;
    Description: string;
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
    EntryId: number;
    Comment: string;
    Url: string;
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
    Name: string;
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
    Id: number;
    Name: string;
    IconUrl: string;
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

