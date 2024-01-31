import React, { useState } from 'react';
import { Route, useNavigate, Routes } from 'react-router-dom';
import './login.css'

const Login = ({ setLoggedIn }) => {
  // vars
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [email, setEmail] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');

  const [login, setLogin] = useState(true);

  const [errors, setErrors] = useState([]);

  // Methods
  const navigate = useNavigate();

  const createUser = async (e) => {
    e.preventDefault();
    
    if (password !== confirmPassword) {
      setErrors(['Passwords do not match'])
      return;
    }

    const response = await fetch('/api/createUser', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username: username, password: password})
    });
    const data = response.json();
    if (data.error !== undefined) {
      return;
    }
    setUsername(data.username);
    let token = `${data.id}${data.username}`
    localStorage.setItem('token', token);
    localStorage.setItem('userId', data.id);
    setLoggedIn(true);
    console.log(`token set to ${token}`)
    navigate('/')

  }

  const handleSubmit = async (e) => {
    e.preventDefault();

    const response = await fetch('/api/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, password })
    });
    let data = await response.json()
    if (data.error === undefined) {
      setUsername(data.username);
      let token = `${data.id}${data.username}`
      localStorage.setItem('token', token);
      localStorage.setItem('userId', data.id);
      setLoggedIn(true);
      console.log(`token set to ${token}`)
      navigate('/')
    }
  };

  const toggleLogin = (value) => {
      setLogin(value);
  }

  // Template
  return (
    <div className="entry_wrapper">
		<header>
			<h3 id="login" onClick={() => {toggleLogin(true)}}> Login </h3>
			<h3 id="register"onClick={() => {toggleLogin(false)}}> Register </h3>
		</header>
		<div className="entry_body">
			<form id="loginForm" style={{display: `${!login ? 'none' : 'block'}`}}>

				<label> Username </label>
				<input type = "text" id = "username" value={username} onChange={(e) => setUsername(e.target.value)} />
		
				<label> Password </label>
				<input type = "text" id = "password" value={password} onChange={(e) => setPassword(e.target.value)}/>
		
				<button onClick={handleSubmit} >Login</button>

			</form>
			<form id="registerForm" style={{display: `${login ? 'none' : 'block'}`}}>

				<label> Username </label>
				<input type ="text" id ="username" value={username} onChange={(e) => setUsername(e.target.value)} />
				
				<label> email </label>
				<input type = "text" id = "email" value={email} onChange={(e) => setEmail(e.target.value)}/>
				
				<label> Password </label>
				<input type = "text" id = "password" value={password} onChange={(e) => setPassword(e.target.value)}/>
				
				<label> Confirm Password </label>
				<input type="text" name="confirmPassword" value={confirmPassword} onChange={(e) => setConfirmPassword(e.target.value)}/>
				
				<button onClick={createUser}>Register</button>
			</form>
		</div>
	</div>
  );
};

export default Login;
