import React, { ChangeEvent } from 'react'

import { useSelector } from 'react-redux'

import s from './styles.module.scss'

import FileList from '../FileList'

import { useAppDispatch } from '../../store'
import { selectDirStack, selectorCurrentFile } from '../../store/slices/file/selectors'
import { fetchGetFile, fetchGetFiles, fetchPostFiles, popToStack, setCurrentDir } from '../../store/slices/file/fileSlice'

import Popup from '../Popup'
import MyInput from '../MyInput'



const Disk: React.FC = () => {
    const dispatch = useAppDispatch()
    const currentDir = useSelector(selectorCurrentFile)
    const dirStack = useSelector(selectDirStack)

    const [isPopup, setIsPopup] = React.useState(false)

    React.useEffect(() => {
        dispatch(fetchGetFiles(currentDir))
    }, [currentDir])

    function openPopup() {
        // console.log(dir)
        setIsPopup(true)
    }

    function closePopup() {
        setIsPopup(false)
    }

    function createHandler(name: string) {
        dispatch(fetchPostFiles({ dirID: currentDir, name }))
    }

    function backClickHandler() {
        const backDirID = dirStack[dirStack.length - 2]
        dispatch(popToStack())
        dispatch(setCurrentDir(backDirID))
    }

    function fileUploadHandler(e: ChangeEvent<HTMLInputElement>) {
        const files = e.currentTarget.files as FileList
        for (let i = 0; i < files.length; i++) {
            let file = files[i]
            console.log(file)
            dispatch(fetchGetFile({ dirID: currentDir, file }))

        }
    }

    return (
        <div className={s.disk}>
            <div className={s.btns}>
                <button
                    disabled={dirStack.length > 0 ? false : true}
                    onClick={() => backClickHandler()}
                    className={s.back}
                >
                    Назад
                </button>
                <button
                    className={s.create}
                    onClick={() => openPopup()}
                >
                    Создать папку
                </button>
                <div className={s.upload}>
                    <label htmlFor="disk-upload-input" className={s.label}>Загрузить  файл</label>
                    <input
                        multiple={true}
                        type="file"
                        id='disk-upload-input'
                        className={s.input}

                        onChange={(e) => fileUploadHandler(e)}
                    />
                </div>
                {isPopup && <Popup createHandler={createHandler} closePopup={closePopup} />}
            </div>

            <FileList />
        </div>
    )
}

export default Disk