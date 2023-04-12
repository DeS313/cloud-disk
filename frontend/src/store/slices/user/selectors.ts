import { RootState } from '../..';

export const selectIsAuth = (state: RootState) => state.userSlice.isAuth;
export const selectUser = (state: RootState) => state.userSlice.currentUser;
