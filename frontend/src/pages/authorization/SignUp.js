import React, { useState } from 'react';
import "./Authorization.css"
import { emailValidator, passwordValidator, passwordVerificationValidator, usernameValidator, dateOfBirthValidator, genderValidator } from './Validators';
import { CountryDropdown } from 'react-country-region-selector';

function SignUp() {
    const [data, setData] = useState({
        email: '',
        username: '',
        password: '',
        passwordVerification: '',
        avatar: null,
        dateOfBirth: '',
        gender: '',
        country: '',
        status: '',
        description: ''
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

    const onUpdateCountry = e => {
      const nextDataState = {
        ...data,
        country: e,
      };
      setData(nextDataState);
    };

    function onAvatarUpdate(e) {
      const nextDataState = {
        ...data,
        avatar: e.target.files[0],
      };
      setData(nextDataState);
    }

    function removeAvatar() {
      const nextDataState = {
        ...data,
        avatar: null,
      };
      setData(nextDataState);
    }

    function handleSignUp(e) {
        e.preventDefault();

        var errors = [
          emailValidator(data.email),
          usernameValidator(data.username), 
          passwordValidator(data.password), 
          passwordVerificationValidator(data.passwordVerification, data.password),
          dateOfBirthValidator(data.dateOfBirth),
          genderValidator(data.gender)
        ]

        const isArrayEmpty = errors.every((error) => error === '')
        setServerError('')
        if(isArrayEmpty) {
          fetch('/api/signup', {
            method: 'POST',
            body: JSON.stringify(data),
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
              <div className='authoInfo'>
                <label>Email*</label>
                  <input type="email" name="email" onChange={onUpdateField}/>
                <label>Username*</label>
                  <input type="text" name="username" onChange={onUpdateField}/>
                <label>Password*</label>
                  <input type="password" name="password" onChange={onUpdateField}/>
                <label>Verify password*</label>
                  <input type="password" name="passwordVerification" onChange={onUpdateField}/>
              </div>
              <div className='personalInfo'>
              <label>Avatar</label>
              {data.avatar && (
                <div>
                <img
                  alt="not found"
                  width={"125px"}
                  src={URL.createObjectURL(data.avatar)}
                />
                <br></br>
                <button onClick={() => removeAvatar()}>Remove</button>
                </div>
              )}
              <br></br>
              <input
                type="file"
                name="myImage"
                onChange={onAvatarUpdate}
              />
                <label>Date of birth*</label>
                  <input type="date" name="dateOfBirth" onChange={onUpdateField}/>
                <label>Gender*</label>
                <select name="gender" onChange={onUpdateField}> 
                  <option value="-">-</option> 
                  <option value="Male">Male</option> 
                  <option value="Female">Female</option> 
                  <option value="Other">Other</option> 
                </select>
                <label>Country</label>
                  <CountryDropdown
                  name='country'
                  value={data.country}
                  onChange={onUpdateCountry} />
                <label>Status</label>
                  <input type="text" name="status" onChange={onUpdateField} maxLength={30}/>
                <label>Description</label>
                  <input type="text" name="description" onChange={onUpdateField} maxLength={100}/>
              </div>
                  <br></br>
                  <br></br>
                <button type="submit" onClick={handleSignUp}>Create Account</button>
            </form>
        </div>
    );
}

export default SignUp;