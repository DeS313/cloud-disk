import axios from 'axios';
import { URLs } from '.';
import { type } from 'os';

export default class FilesService {
  static async getFiles(dirID: string) {
    const token = localStorage.getItem('token');
    const res = await axios.get(URLs.SEVER + URLs.FILES + `${dirID ? `?parant=${dirID}` : ''}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    console.log(res.status);
    if (res.status != 200) {
      throw new Error(res.statusText);
    }
    return res.data;
  }
  static async createDir(dirID: string, name: string) {
    const token = localStorage.getItem('token');

    const res = await axios.post(
      URLs.SEVER + URLs.FILES,
      {
        name,
        parent: dirID,
        type: 'dir',
      },
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      },
    );
    console.log(res.status);
    if (res.status != 200) {
      throw new Error(res.statusText);
    }
    return res.data;
  }

  static async uploadFile(dirID: string, file: any) {
    const token = localStorage.getItem('token');
    console.log(file, 'file');
    const formData = new FormData();
    formData.append('file', file);
    if (dirID) {
      formData.append('parent', dirID);
    }
    const res = await axios.post(URLs.SEVER + '/upload', formData, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
      onUploadProgress: (progressEvent) => {
        const totalLength = progressEvent.loaded
          ? progressEvent.total
          : progressEvent.event.target.getResponseHeader('content-length') ||
            progressEvent.event.target.getResponseHeader('x-decompressed-content-length');
        if (totalLength) {
          const progress = Math.round((progressEvent.loaded * 100) / totalLength);
          console.log(progress, 'total');
        }
      },
    });

    if (res.status != 200) {
      throw new Error(res.statusText);
    }
    return res.data;
  }
}
