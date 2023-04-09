import axios from 'axios';
import { URLs } from '.';

export default class AuthService {
  static async registration(email: string, password: string) {
    try {
      const data = JSON.stringify({
        email,
        password,
      });
      const res = await fetch(URLs.SEVER + URLs.REGISTRATION, {
        method: 'POST',
        mode: 'cors',
        headers: {
          'Content-Type': 'application/json',
        },
        body: data,
      });

      if (!res.ok) {
        throw new Error(res.statusText);
      }

      return res.json();
    } catch (error) {
      alert(error);
      throw error;
    }
  }

  static async login(email: string, password: string) {
    try {
      const user = JSON.stringify({
        email,
        password,
      });
      const res = await fetch(URLs.SEVER + URLs.LOGIN, {
        method: 'POST',
        mode: 'cors',
        headers: {
          'Content-Type': 'application/json',
        },
        body: user,
      });

      if (!res.ok) {
        throw new Error(res.statusText);
      }
      const data = await res.json();
      localStorage.setItem('token', data['token']);
      return data.user;
    } catch (error) {
      alert(error);
      localStorage.removeItem('token');
      throw error;
    }
  }

  static async getUser() {
    try {
      const token = localStorage.getItem('token');
      const res = await fetch(URLs.SEVER + URLs.AUTH, {
        mode: 'cors',
        headers: {
          'Content-Type': 'application/json',
          Authorization: 'Bearer ' + token,
        },
      });

      if (!res.ok) {
        throw res.statusText;
      }

      const data = await res.json();
      return data;
    } catch (error) {
      alert(error);
      throw error;
    }
  }
}
