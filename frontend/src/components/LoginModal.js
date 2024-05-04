import React, { useState } from 'react';
import './PopupModal.css';
import loginlogo from '.././images/SS.png';
import axios from 'axios';

const LoginModal = ({ onLogin, onCreateAccount, setSidebarStructure }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const fetchSidebarStructure = async (token, userID) => {
    try {
      const response = await axios.post(
        'http://localhost:8080/api/sidebar/build',
        {
          userID,
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );
      console.log('Sidebar structure:', response.data);
      setSidebarStructure(response.data); // Update the sidebarStructure state in the parent component
    } catch (error) {
      console.error('Error fetching sidebar structure:', error);
    }
  };

  const handleLogin = async () => {
    try {
      const response = await axios.post('http://localhost:8080/users/login', {
        username,
        password,
      });
      if (response.data.message === 'Login successful') {
        // Store the token and user ID in local storage
        localStorage.setItem('token', response.data.token);
        localStorage.setItem('userID', response.data.userID);

        fetchSidebarStructure(response.data.token, response.data.userID);

        onLogin();
      } else {
        alert('Invalid credentials. Please try again.');
      }
    } catch (error) {
      console.error('Error logging in:', error);
      if (error.response) {
        // Falls out of the range of 2xx
        console.error('Response data:', error.response.data);
        console.error('Response status:', error.response.status);
        console.error('Response headers:', error.response.headers);
      } else if (error.request) {
        // No response was received
        console.error('Request:', error.request);
      } else {
        // Something happened in setting up the request
        console.error('Error:', error.message);
      }
      alert('An error occurred while logging in. Please try again.');
    }
  };

  return (
    <div>
      <div className="login-modal">
        <div className="login-form">
          <img src={loginlogo} alt="SyncScribe Logo" className="logo" />
          <input
            type="text"
            placeholder="Please enter your Username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
          <input
            type="password"
            placeholder="Please enter your Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <button onClick={handleLogin}>Login</button>
          <br />
          <p>
            or <a href="#" onClick={onCreateAccount}>create an account</a>
          </p>
        </div>
      </div>
    </div>
  );
};

export default LoginModal;