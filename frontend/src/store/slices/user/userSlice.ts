import { PayloadAction, createAsyncThunk, createSlice } from '@reduxjs/toolkit';

import { IUserState, IFetchUserArgs } from './types';
import AuthService from '../../../API/authService';
import { useAppDispatch } from '../..';

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

export const fetchRegistration = createAsyncThunk(
  'users/fetchRegistration',
  async (params: IFetchUserArgs, thunkApi) => {
    try {
      const { email, password } = params;
      const data = await AuthService.registration(email, password);
      console.log(data);
      return data;
    } catch (e) {
      alert(e);
    }
  },
);

export const fetchLogin = createAsyncThunk(
  'users/fetchLogin',
  async (params: IFetchUserArgs, thunkApi) => {
    const { email, password } = params;
    const user = await AuthService.login(email, password);
    return user;
  },
);

export const fetchGetUser = createAsyncThunk('users/fetchGetUser', async (thunkApi) => {
  const user = await AuthService.getUser();

  return user;
});

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    logout(state) {
      state.currentUser = {
        ID: '',
        Email: '',
        Password: '',
        DiskSpace: 0,
        UserSpace: 0,
        Avatar: '',
      };
      state.isAuth = false;
    },
  },
  extraReducers: (bulder) => {
    bulder.addCase(fetchRegistration.pending, (state, action) => {});
    bulder.addCase(fetchRegistration.fulfilled, (state, action) => {
      alert('ok');
    });
    bulder.addCase(fetchRegistration.rejected, (state, action) => {
      alert(action.payload);
      alert(action.error);
    });

    bulder.addCase(fetchLogin.pending, (state, action) => {});
    bulder.addCase(fetchLogin.fulfilled, (state, action) => {
      state.currentUser = action.payload;
      state.isAuth = true;
    });
    bulder.addCase(fetchLogin.rejected, (state) => {
      state.isAuth = false;
    });
    bulder.addCase(fetchGetUser.pending, () => {});
    bulder.addCase(fetchGetUser.fulfilled, (state, action) => {
      state.currentUser = action.payload;
      state.isAuth = true;
    });
    bulder.addCase(fetchGetUser.rejected, (state, action) => {
      state.isAuth = false;
    });
  },
});

export const { logout } = userSlice.actions;

export default userSlice.reducer;
