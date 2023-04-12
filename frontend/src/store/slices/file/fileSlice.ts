import { PayloadAction, createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import { IFilesState, TCreateDir, TFile, TUploadFile } from './types';
import FilesService from '../../../API/filesService';

const initialState: IFilesState = {
  files: [],
  currentDir: '',
  dirStack: [],
};

export const fetchGetFiles = createAsyncThunk(
  'files/fetchGetFiles',
  async (dirID: string, thunkApi) => {
    try {
      const data = await FilesService.getFiles(dirID);
      return data.file as TFile[];
    } catch (e) {
      alert(e);
    }
  },
);

export const fetchGetFile = createAsyncThunk(
  'files/fetchGetFile',
  async (obj: TUploadFile, thunkApi) => {
    try {
      const { dirID, file } = obj;
      const data = await FilesService.uploadFile(dirID, file);
      return data;
    } catch (e) {
      alert(e);
    }
  },
);

export const fetchPostFiles = createAsyncThunk(
  'files/fetchPostFiles',
  async (obj: TCreateDir, thunkApi) => {
    try {
      const data = await FilesService.createDir(obj.dirID, obj.name);
      return data;
    } catch (e) {
      alert(e);
    }
  },
);

const fileSlice = createSlice({
  name: 'file',
  initialState,
  reducers: {
    setFiles(state, action) {
      state.files.push(action.payload);
    },
    setCurrentDir(state, action) {
      state.currentDir = action.payload;
    },
    pushToStack(state, action: PayloadAction<string>) {
      state.dirStack.push(action.payload);
    },
    popToStack(state) {
      state.dirStack.pop();
    },
  },
  extraReducers: (bulder) => {
    bulder.addCase(fetchGetFiles.pending, (state, action) => {
      state.files = [];
    });
    bulder.addCase(fetchGetFiles.fulfilled, (state, action) => {
      state.files = action.payload as TFile[];
    });
    bulder.addCase(fetchGetFiles.rejected, (state, action) => {});

    bulder.addCase(fetchPostFiles.pending, (state, action) => {});
    bulder.addCase(fetchPostFiles.fulfilled, (state, action) => {
      state.files = action.payload;
    });

    bulder.addCase(fetchGetFile.pending, (state, action) => {});
    bulder.addCase(fetchGetFile.fulfilled, (state, action) => {
      state.files = [...state.files, action.payload];
    });
  },
});

export const { setCurrentDir, pushToStack, popToStack } = fileSlice.actions;

export default fileSlice.reducer;
