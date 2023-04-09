import { RootState } from '../..';

export const selectIsAuth = (state: RootState) => state.userSlice.isAuth;
