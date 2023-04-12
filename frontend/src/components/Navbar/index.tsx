import React from 'react'

import s from './styles.module.scss'
import Logo from '../../assets/img/navbar-logo.svg'
import { NavLink } from 'react-router-dom'
import { RouteTo } from '../../route'
import { useSelector } from 'react-redux'
import { selectIsAuth, selectUser } from '../../store/slices/user/selectors'
import { useAppDispatch } from '../../store'
import { logout } from '../../store/slices/user/userSlice'
import { useDispatch } from 'react-redux'

const Navbar: React.FC = () => {
    const isAuth = useSelector(selectIsAuth)
    const user = useSelector(selectUser)
    const dispatch = useDispatch()
    return (
        <div className={s.navbar}>
            <div className={s.container}>
                <img src={Logo} alt="logo" className={s.logo} />
                <div className={s.header}>GRM CLOUD</div>
                {
                    !isAuth ?
                        <>
                            <div className={s.login}>
                                <NavLink to={RouteTo.LOGIN}>Войти</NavLink>
                            </div>
                            <div className={s.registration}>
                                <NavLink to={RouteTo.REGISTRATION}>Регистрация</NavLink>
                            </div>
                        </>
                        :
                        <>
                            <div>{user && user.Email}</div>
                            <div
                                onClick={() => dispatch(logout())}
                                className={s.logout}>
                                Выход
                            </div>
                        </>
                }

            </div>
        </div >
    )
}

export default Navbar