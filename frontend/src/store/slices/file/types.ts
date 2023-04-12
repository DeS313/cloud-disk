export type TFile = {
  ID: string;
  Name: string;
  Type: string;
  AccessLink: string;
  Path: string;
  Date: string;
  Size: number;
  UserID: string;
  ParrentID: string;
};

export type TCreateDir = {
  dirID: string;
  name: string;
};

export type TUploadFile = {
  dirID: string;
  file: any;
};

export interface IFilesState {
  currentDir: string;
  files: TFile[];
  dirStack: string[];
}
