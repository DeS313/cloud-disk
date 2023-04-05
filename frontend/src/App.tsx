import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';

import './sass/app.scss'

import Navbar from './components/Navbar';
import Registration from './components/Registraion';
import { RouteTo } from './route';

function App() {
  return (
    <BrowserRouter>
      <div className='app'>
        <Navbar />
        <div className="wrap">
          <Routes>
            <Route path={RouteTo.REGISTRATION} Component={Registration} />
          </Routes>
        </div>
      </div>
    </BrowserRouter>
  );
}

export default App;
