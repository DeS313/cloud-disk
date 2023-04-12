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
    const res = await axios.post(
      URLs.SEVER + URLs.LOGIN,
      { email, password },
      {
        headers: {
          'Content-Type': 'application/json',
        },
      },
    );

    if (!res.status) {
      throw (res.statusText, res.data);
    }
    localStorage.setItem('token', res.data.token);
    const data = res.data.user;
    return data;
  }

  static async getUser() {
    const token = localStorage.getItem('token');

    const res = await axios.get(URLs.SEVER + URLs.AUTH, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (!res.status) {
      throw (res.statusText, res.data);
    }

    const data = res.data;
    return data;
  }
}
