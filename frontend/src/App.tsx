import React from 'react';
import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom';
import { useSelector } from 'react-redux';

import './sass/app.scss'

import Navbar from './components/Navbar';
import Registration from './components/Registraion';
import { RouteTo } from './route';
import Login from './components/Login';
import { selectIsAuth } from './store/slices/user/selectors';
import { useAppDispatch } from './store';
import { fetchGetUser } from './store/slices/user/userSlice';
import Disk from './components/Disk';


function App() {
  const isAuth = useSelector(selectIsAuth)
  const dispatch = useAppDispatch()

  React.useEffect(() => {
    dispatch(fetchGetUser())
  }, [])

  return (
    <BrowserRouter>
      <div className='app'>
        <Navbar />
        <div className="wrap">
          <Routes>
            {!isAuth ?
              <>
                <Route path={RouteTo.REGISTRATION} Component={Registration} />
                <Route path={RouteTo.LOGIN} Component={Login} />
                <Route path="*" element={<Navigate to="/login" replace />} />
              </>
              :
              <>
                <Route path={RouteTo.DISK} Component={Disk} />
                <Route path="*" element={<Navigate to="/" replace />} />
              </>
            }
          </Routes>
        </div>
      </div>
    </BrowserRouter>
  );
}

export default App;
