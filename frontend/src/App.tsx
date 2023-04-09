import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { useSelector } from 'react-redux';

import './sass/app.scss'

import Navbar from './components/Navbar';
import Registration from './components/Registraion';
import { RouteTo } from './route';
import Login from './components/Login';
import { selectIsAuth } from './store/slices/user/selectors';
import { useAppDispatch } from './store';
import { fetchGetUser } from './store/slices/user/userSlice';


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
            {!isAuth &&
              <>
                <Route path={RouteTo.REGISTRATION} Component={Registration} />
                <Route path={RouteTo.LOGIN} Component={Login} />
              </>
            }
          </Routes>
        </div>
      </div>
    </BrowserRouter>
  );
}

export default App;
