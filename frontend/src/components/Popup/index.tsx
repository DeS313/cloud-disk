import React from 'react'

import s from './styles.module.scss'
import MyInput from '../MyInput'

type PopupProps = {
    closePopup: () => void
    createHandler: (name: string) => void
}

const Popup: React.FC<PopupProps> = ({ closePopup, createHandler }) => {
    const [dirName, setDirName] = React.useState("")

    return (
        <div className={s.popup} onClick={() => closePopup()}>
            <div className={s.content} onClick={(e) => e.stopPropagation()}>
                <div className={s.header}>
                    <div className={s.title}>Создать новую папку</div>
                    <button
                        onClick={() => closePopup()}
                        className={s.close}
                    >
                        X
                    </button>
                </div>
                <MyInput
                    type='text'
                    placeholder='Введите название папки...'
                    value={dirName}
                    onChange={(e) => setDirName(e.currentTarget.value)}
                />
                <button
                    className={s.create}
                    onClick={() => createHandler(dirName)}
                >
                    Создать
                </button>
            </div>
        </div>
    )
}

export default Popup