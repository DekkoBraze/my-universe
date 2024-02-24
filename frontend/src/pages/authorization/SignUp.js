import React, { useState } from 'react';
import "./Authorization.css"
import { emailValidator, passwordValidator, passwordVerificationValidator, usernameValidator } from './Validators';

function SignUp() {
    const [data, setData] = useState({
        email: '',
        username: '',
        password: '',
        passwordVerification: '',
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

    function handleSignUp(e) {
        e.preventDefault();
        
        var errors = [
          emailValidator(data.email),
          usernameValidator(data.username), 
          passwordValidator(data.password), 
          passwordVerificationValidator(data.passwordVerification, data.password)
        ]

        const isArrayEmpty = errors.every((error) => error === '')
        setServerError('')
        if(isArrayEmpty) {
          fetch('/api/signup', {
            method: 'POST',
            body: JSON.stringify(data),
            mode: 'no-cors',
            headers: {
              'Content-Type': 'application/json',
            }
          })
            .then((response) => response.json())
            .then((data) => {
              if (data.message === 'OK') {
                window.location.href = "login"
              } else {
                return setServerError(data.message), setValidationErrors([])
              }
            })
            .catch((error) => {
              console.log(error)
            })
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
            <form className="authorizationForm">
                <label>Email</label>
                  <input type="text" name="email" onChange={onUpdateField}/>
                <label>Username</label>
                  <input type="text" name="username" onChange={onUpdateField}/>
                <label>Password</label>
                  <input type="text" name="password" onChange={onUpdateField}/>
                <label>Verify password</label>
                  <input type="text" name="passwordVerification" onChange={onUpdateField}/>
                  <br></br>
                  <br></br>
                <button type="submit" onClick={handleSignUp}>Create Account</button>
            </form>
        </div>
    );
}

export default SignUp;