import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';

import { IUserState, IFetchUserArgs } from './types';
import AuthService from '../../../API/authService';

const initialState: IUserState = {
  currentUser: {
    ID: '',
    Email: '',
    Password: '',
    DiskSpace: 0,
    UserSpace: 0,
    Avatar: '',
  },
  isAuth: false,
};

export const fetchUser = createAsyncThunk(
  'users/fetchUser',
  async (params: IFetchUserArgs, thunkApi) => {
    try {
      const { email, password } = params;

      const data = await AuthService.registration(email, password);

      return data;
    } catch (e) {
      alert(e);
    }
  },
);

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {},
  extraReducers: (bulder) => {
    bulder.addCase(fetchUser.pending, (state, action) => {});
    bulder.addCase(fetchUser.fulfilled, (state, action) => {
      state.currentUser = action.payload;
      state.isAuth = true;
    });
    bulder.addCase(fetchUser.rejected, (state, action) => {
      alert(action.payload);
      alert(action.error);
    });
  },
});

export default userSlice.reducer;
