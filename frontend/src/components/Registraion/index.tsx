import React, { ChangeEvent, FormEvent } from 'react'

import s from './styles.module.scss'
import MyInput from '../MyInput'
import { IRegState } from './types'
import AuthService from '../../API/authService'
import { fetchRegistration } from '../../store/slices/user/userSlice'
import { useAppDispatch } from '../../store'

const Registration: React.FC = () => {
    const [regValue, setRegValue] = React.useState<IRegState>({
        email: '',
        password: ''
    })
    const dispatch = useAppDispatch()

    const onChangeEmail = (e: ChangeEvent<HTMLInputElement>) => {
        setRegValue({ ...regValue, "email": e.currentTarget.value })
    }

    const onChangePassword = (e: ChangeEvent<HTMLInputElement>) => {
        setRegValue({ ...regValue, "password": e.currentTarget.value })
    }

    const registration = () => {
        dispatch(fetchRegistration(regValue))
        setRegValue({ email: "", password: "" })
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

            <button
                onClick={registration}
                className={s.btn}
            >
                Зарегистрироваться
            </button>
        </div>
    )
}

export default Registration