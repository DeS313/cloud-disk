import React, { ChangeEvent, FormEvent } from 'react'

import s from './styles.module.scss'
import MyInput from '../MyInput'
import { IRegState } from './types'
import AuthService from '../../API/authService'
import { fetchUser } from '../../store/slices/user/userSlice'

const Registration: React.FC = () => {
    const [regValue, setRegValue] = React.useState<IRegState>({
        email: '',
        password: ''
    })

    const onChangeEmail = (e: ChangeEvent<HTMLInputElement>) => {
        setRegValue({ ...regValue, "email": e.currentTarget.value })
    }

    const onChangePassword = (e: ChangeEvent<HTMLInputElement>) => {
        setRegValue({ ...regValue, "password": e.currentTarget.value })
    }


    return (
        <div className={s.registration}>
            <div className={s.header}>Регистрация</div>
            <MyInput
                type='text'
                onChange={onChangeEmail}
                value={regValue.email}
                placeholder='Введите email..' />
            <MyInput
                type='password'
                onChange={onChangePassword}
                value={regValue.password}
                placeholder='Введите пароль...' />

            <button onClick={() => fetchUser(regValue)} className={s.btn}>Войти</button>
        </div>
    )
}

export default Registration