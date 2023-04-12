import React from 'react'

import s from './styles.module.scss'

import File from '../File'

import { useSelector } from 'react-redux'
import { selectorFiles } from '../../store/slices/file/selectors'

const FileList: React.FC = () => {

    const files = useSelector(selectorFiles)

    return (
        <div className={s.fileList}>
            <div className={s.header}>
                <div className={s.name}>Название</div>
                <div className={s.date}>Дата</div>
                <div className={s.size}>Размер</div>
            </div>
            {files && files.map(file => <File key={file.ID} {...file} />)}
        </div>
    )
}

export default FileList