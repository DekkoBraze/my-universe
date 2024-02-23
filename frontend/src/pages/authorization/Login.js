import React, { useState } from 'react';

function Login() {
    const [data, setData] = useState({
        email: '',
        password: '',
      });

    function handleEmailChange(e) {
        setData({...data, email: e.target.value})
    }

    function handlePasswordChange(e) {
        setData({...data, password: e.target.value})
    }

    function handleLogin(e) {
        e.preventDefault();
        
        fetch('/api/login', {
          method: 'POST',
          body: JSON.stringify(data),
        })
          .then((response) => response.json())
          .then((data) => {
            var oneday = new Date()
            oneday.setHours(oneday.getHours() + 24);
            var profileData = {...data, timestamp: oneday}
            localStorage.setItem("user", JSON.stringify(profileData))
            window.location.href = "profile/" + data.username
          })
          .catch((error) => {
            console.log(error)
          });
    }

    return (
        <div style={{
            textAlign: "center",
        }}>
            <form>
                <label>Email</label>
                <br></br>
                <input type="text" name="email" onChange={handleEmailChange}/>
                <br></br>
                <br></br>
                <label>Password</label>
                <br></br>
                <input type="text" name="password" onChange={handlePasswordChange}/>
                <br></br>
                <br></br>
                <button type="submit" onClick={handleLogin}>Sign In</button>
            </form>
        </div>
    );
}

export default Login;