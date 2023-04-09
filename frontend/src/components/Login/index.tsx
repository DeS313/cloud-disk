import React, { ChangeEvent } from 'react'

import s from './styles.module.scss'

import MyInput from '../MyInput'

import { fetchLogin } from '../../store/slices/user/userSlice'
import { useAppDispatch } from '../../store'



const Login: React.FC = () => {
    const [regValue, setRegValue] = React.useState({
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

    const login = async () => {
        try {
            dispatch(fetchLogin(regValue))
            setRegValue({ email: "", password: "" })
        } catch (err) {
            alert(err)
        }
    }


    return (
        <div className={s.login}>
            <div className={s.header}>Авторизация</div>
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

            <button onClick={() => login()} className={s.btn}>Войти</button>
        </div>
    )
}

export default Login