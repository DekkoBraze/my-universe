import React, { useState } from "react";
import { Link } from "react-router-dom";
import "./Navbar.css"

function Navbar() {
    return (
        <nav className="bar">
          <Link to="/" className="title">MyUniverse</Link>
          <ul className="auth">
            <li>
              <Link to="/signup">SignUp</Link>
            </li>
            <li>
              <Link to="/login">Login</Link>
            </li>
          </ul>
        </nav>
        );
}

export default Navbar;