export interface User {
  id: number;
  name: string;
  iconUrl: string;
}

export interface Calendar {
  id: number;
  title: string;
  owner: User;
  description: string;
  year: number;
  entryCount: number;
  entries?: Entry[];
}

export interface Entry {
  id: number;
  owner: User;
  day: number;
  comment: string;
  url: string;
  title: string;
  imageUrl: string;
}
