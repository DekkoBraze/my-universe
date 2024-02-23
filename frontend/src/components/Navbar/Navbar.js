import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import "./Navbar.css"

function Navbar() {
  //const [profileData, setProfileData] = useState([]);
  const profileData = JSON.parse(localStorage.getItem("user"));
  if (profileData)
  {
    if (profileData.timestamp < new Date()) {
      localStorage.removeItem('user')
    }
    
  }
  function handleLogout(e) {
    localStorage.removeItem('user');
    window.location.href = "/"
  }
  
  return (
    <nav className="bar">
      <Link to="/" className="title">MyUniverse</Link>
      {
        profileData ? (
            <ul className="bar-profileInfo">
              <li>
                <Link className="username" to={"/profile/" + profileData.username}>{profileData.username}</Link>
              </li>
              <li>
                <Link onClick={handleLogout}>Logout</Link>
              </li>
            </ul>
        ) : (
          <ul className="auth">
            <li>
              <Link to="/signup">SignUp</Link>
            </li>
            <li>
              <Link to="/login">Login</Link>
            </li>
          </ul>
        )
      }
    </nav>
  );
}

export default Navbar;
