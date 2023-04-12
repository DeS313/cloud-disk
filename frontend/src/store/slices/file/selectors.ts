import { useSelector } from 'react-redux';
import { RootState } from '../..';
import { Root } from 'react-dom/client';

export const selectorCurrentFile = (state: RootState) => state.fileSlice.currentDir;
export const selectorFiles = (state: RootState) => state.fileSlice.files;
export const selectDirStack = (state: RootState) => state.fileSlice.dirStack;
