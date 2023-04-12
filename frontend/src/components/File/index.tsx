import React from 'react'

import s from './styles.module.scss'
import DirLogo from '../../assets/img/dir.svg'
import FileLogo from "../../assets/img/file.svg"

import { TFile } from '../../store/slices/file/types'
import { useAppDispatch } from '../../store'
import { pushToStack, setCurrentDir } from '../../store/slices/file/fileSlice'
import { useSelector } from 'react-redux'
import { selectorCurrentFile } from '../../store/slices/file/selectors'

const File: React.FC<TFile> = (props, { Name, Size }) => {
    const date = new Date(props.Date);
    const dispatch = useAppDispatch()
    const currentDir = useSelector(selectorCurrentFile)

    function openDirHandler() {
        if (props.Type === 'dir') {
            dispatch(pushToStack(props.ID))
            dispatch(setCurrentDir(props.ID))
        }
    }

    return (
        <div className={s.file} onClick={() => openDirHandler()}>
            <img src={props.Type === "dir" ? DirLogo : FileLogo} alt={props.Name} className={s.img} />
            <div className={s.name}>{props.Name}</div>
            <div className={s.date}>{date.toLocaleDateString()}</div>
            <div className={s.size}>{props.Size}</div>
        </div>
    )
}

export default File