import { json } from 'stream/consumers';
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
    }
  }
}
