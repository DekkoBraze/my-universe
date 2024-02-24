export const emailValidator = email => {
    if (!email) {
      return "Field 'Email' is required";
    } else if (!new RegExp(/\S+@\S+\.\S+/).test(email)) {
      return "Incorrect email format";
    }
    return "";
  };
  
  export const passwordValidator = password => {
    if (!password) {
      return "Field 'Password' is required";
    } else if (password.length < 8) {
      return "Password must have a minimum 8 characters";
    }
    return "";
  };

  export const usernameValidator = username => {
    if (!username) {
      return "Field 'Username' is required"
    }
    return "";
  }
  
  export const passwordVerificationValidator = (confirmPassword, password) => {
    if (!confirmPassword) {
      return "Field 'Verify password' is required";
    } else if (confirmPassword !== password) {
      return "Passwords do not match";
    }
    return "";
  };

  export const dateOfBirthValidator = (dateOfBirth) => {
    if (!dateOfBirth) {
      return "Field 'Date of birth' is required";
    }
    return "";
  };