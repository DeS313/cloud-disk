import { configureStore } from '@reduxjs/toolkit';
import { useDispatch } from 'react-redux';

import redusers from './slices';

export const store = configureStore({
  reducer: {
    ...redusers,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
export const useAppDispatch = () => useDispatch<AppDispatch>();
