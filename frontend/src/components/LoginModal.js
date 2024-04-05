import React, { useState } from 'react';
import './LoginModal.css';
import loginlogo from '.././images/SS.png';

const LoginModal = ({ onLogin, onCreateAccount  }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleLogin = () => {
    // Perform login logic here
    onLogin();
  };

  return (
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
        <br></br>
        <p>or <a href="#" onClick={onCreateAccount}>create an account</a></p>
      </div>
    </div>
  );
};

export default LoginModal;