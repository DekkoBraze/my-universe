import React, { useState } from 'react';

function SignUp() {
    const [data, setData] = useState({
        email: '',
        username: '',
        password: '',
        passwordVerification: '',
      });

    function handleEmailChange(e) {
        setData({...data, email: e.target.value})
    }

    function handleUsernameChange(e) {
        setData({...data, username: e.target.value})
    }

    function handlePasswordChange(e) {
        setData({...data, password: e.target.value})
    }

    function handlePasswordVerificationChange(e) {
        setData({...data, passwordVerification: e.target.value})
        console.log(JSON.stringify(data))
    }

    function handleSignUp(e) {
        e.preventDefault();

        fetch('http://localhost:8000/signup', {
          method: 'POST',
          body: JSON.stringify(data),
          mode: 'no-cors',
          headers: {
            'Content-Type': 'application/json',
          }
        })
          .then((response) => response.json())
          .then((data) => {
            // Handle the response data
          })
          .catch((error) => {
            console.log(error)
          });

          window.location.href = "login"
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
                <label>Username</label>
                <br></br>
                <input type="text" name="username" onChange={handleUsernameChange}/>
                <br></br>
                <br></br>
                <label>Password</label>
                <br></br>
                <input type="text" name="password" onChange={handlePasswordChange}/>
                <br></br>
                <br></br>
                <label>Verify password</label>
                <br></br>
                <input type="text" name="password_verification" onChange={handlePasswordVerificationChange}/>
                <br></br>
                <br></br>
                <button type="submit" onClick={handleSignUp}>Create Account</button>
            </form>
        </div>
    );
}

export default SignUp;