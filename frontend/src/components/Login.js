import React, { useState } from 'react';
import { Route, useNavigate, Routes } from 'react-router-dom';
import './../App.css';

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
    const data = await response.json();
    if (data.status !== 'Success') {
      setErrors([data.Error]);
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
    <div className='flex flex-col items-center justify-center min-h-screen'>
  <div className='bg-white p-8 rounded shadow-lg max-w-md'>
    <div className='mb-4'>
      {errors.map((error, index) => (
        <p key={index} className='text-red-500'>{error}</p>
      ))}
    </div>
    <header className='mb-4'>
      <h3
        id='login'
        className={`cursor-pointer ${login ? 'text-blue-500' : 'text-gray-500'}`}
        onClick={() => { toggleLogin(true) }}
      >
        Login
      </h3>
      <h3
        id='register'
        className={`cursor-pointer ${!login ? 'text-blue-500' : 'text-gray-500'}`}
        onClick={() => { toggleLogin(false) }}
      >
        Register
      </h3>
    </header>
    <div className='entry_body'>
      <form
        id='loginForm'
        style={{ display: `${login ? 'block' : 'none'}` }}
        className='mb-4'
      >
        <label className='block mb-2'>Username</label>
        <input
          type='text'
          id='username'
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          className='w-full border p-2 rounded'
        />

        <label className='block mb-2'>Password</label>
        <input
          type='password'
          id='password'
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          className='w-full border p-2 rounded'
        />

        <button onClick={handleSubmit} className='mt-4 bg-blue-500 text-white p-2 rounded'>
          Login
        </button>
      </form>
      <form
        id='registerForm'
        style={{ display: `${!login ? 'block' : 'none'}` }}
        className='mb-4'
      >
        <label className='block mb-2'>Username</label>
        <input
          type='text'
          id='username'
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          className='w-full border p-2 rounded'
        />

        <label className='block mb-2'>Email</label>
        <input
          type='text'
          id='email'
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          className='w-full border p-2 rounded'
        />

        <label className='block mb-2'>Password</label>
        <input
          type='password'
          id='password'
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          className='w-full border p-2 rounded'
        />

        <label className='block mb-2'>Confirm Password</label>
        <input
          type='password'
          name='confirmPassword'
          value={confirmPassword}
          onChange={(e) => setConfirmPassword(e.target.value)}
          className='w-full border p-2 rounded'
        />

        <button onClick={createUser} className='mt-4 bg-blue-500 text-white p-2 rounded'>
          Register
        </button>
      </form>
    </div>
  </div>
</div>
  );
};

export default Login;
