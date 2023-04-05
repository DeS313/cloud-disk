import React from 'react'

import s from './styles.module.scss'
import Logo from '../../assets/img/navbar-logo.svg'
import { NavLink } from 'react-router-dom'
import { RouteTo } from '../../route'

const Navbar: React.FC = () => {
    return (
        <div className={s.navbar}>
            <div className={s.container}>
                <img src={Logo} alt="logo" className={s.logo} />
                <div className={s.header}>GRM CLOUD</div>
                <div className={s.login}><NavLink to={RouteTo.LOGIN}>Войти</NavLink></div>
                <div className={s.registration}><NavLink to={RouteTo.REGISTRATION}>Регистрация</NavLink></div>
            </div>
        </div >
    )
}

export default Navbar