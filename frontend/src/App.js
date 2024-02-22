import React from 'react';
import Profile from './pages/Profile/Profile';
import Home from './pages/Home/Home';
import NotFound from './pages/NotFound';
import SignUp from './pages/authorization/SignUp'
import Login from './pages/authorization/Login';
import Navbar from './components/Navbar/Navbar';
import { Routes, Route } from "react-router-dom";
import "./App.css"

function App() {
    return (
        <div className='App'>
            <Navbar />
            <Routes>
                <Route path="/" exact element={<Home />} />
                <Route path="/profile/:username" element={<Profile />} />
                <Route path="/signup" element={<SignUp />} />
                <Route path="/login" element={<Login />} />
                <Route path="*" element={<NotFound />} />
            </Routes>
        </div>
    );
}

export default App;
//<Route path="/registration" element={<Registration />} />
//<Route path="/login" element={<Login />} />
//