import React, { useState } from 'react';
import "./Authorization.css"
import { emailValidator, passwordValidator } from './Validators';

function Login() {
    const [data, setData] = useState({
        email: '',
        password: '',
      });
      
      const [validationErrors, setValidationErrors] = useState([])
      const [serverError, setServerError] = useState('')

      const onUpdateField = e => {
        const nextDataState = {
          ...data,
          [e.target.name]: e.target.value,
        };
        setData(nextDataState);
      };

    function handleLogin(e) {
        e.preventDefault();
        
        var errors = [
          emailValidator(data.email),
          passwordValidator(data.password)
        ]

        const isArrayEmpty = errors.every((error) => error === '')
        setServerError('')
        if (isArrayEmpty) {
          fetch('/api/login', {
            method: 'POST',
            body: JSON.stringify(data),
          })
            .then((response) => response.json())
            .then((data) => {
              if (data.message === 'OK') {
                var oneday = new Date()
                oneday.setHours(oneday.getHours() + 24);
                var profileData = {...data, timestamp: oneday}
                localStorage.setItem("user", JSON.stringify(profileData))
                window.location.href = "profile/" + data.username
              } else {
                return setServerError(data.message), setValidationErrors([])
              }
            })
            .catch((error) => {
              console.log(error)
            });
        } else {
          setValidationErrors(errors)
        }  
    }

    return (
      <div className="mainSignUp">
          {serverError ?  
          <div className='validationError'>
            <h3>!{serverError}</h3>
          </div> 
          : null}
            {validationErrors.map((error) => {
              if (error) {
                return (
                  <div className='validationError'>
                    <h3>!{error}</h3>
                  </div>
                )
              }
            })}
        <form className='authorizationForm'>
          <label>Email</label>
            <input type="text" name="email" onChange={onUpdateField}/>
          <label>Password</label>
            <input type="text" name="password" onChange={onUpdateField}/>
          <br></br>
          <br></br>
          <button type="submit" onClick={handleLogin}>Sign In</button>
        </form>
      </div>
    );
}

export default Login;