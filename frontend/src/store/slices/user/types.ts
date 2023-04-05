export interface IUserState {
  currentUser: TUser;
  isAuth: boolean;
}

export type TUser = {
  ID: string;
  Email: string;
  Password: string;
  DiskSpace: number;
  UserSpace: number;
  Avatar: string;
};

export interface IFetchUserArgs {
  email: string;
  password: string;
}
