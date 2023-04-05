import React from 'react'
import s from './styles.module.scss'
import { InputProps } from './types'

const MyInput: React.FC<InputProps> = (props) => {
    return (
        <input className={s.input} {...props} />
    )
}

export default MyInput