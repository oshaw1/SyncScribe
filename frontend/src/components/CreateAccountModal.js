import React, { useState } from 'react';
import axios from 'axios';
import './LoginModal.css';
import loginlogo from '.././images/SS.png';

const CreateAccountModal = ({ onClose, onCreateSuccess }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleCreateAccount = async () => {
    try {
      const response = await axios.post('http://localhost:8080/users/create', {
        username,
        password,
      });
      console.log('Account created successfully:', response.data);

      // After successful account creation, notify the parent component
      onCreateSuccess(username, password);
    } catch (error) {
      console.error('Error creating account:', error);
      setError('An error occurred while creating the account. Please try again.');
    }
  };

  return (
    <div className="login-modal">
      <div className="login-form">
        <img src={loginlogo} alt="SyncScribe Logo" className="logo" />
        <h2>Create Account</h2>
        {error && <p className="error-message">{error}</p>}
        <input
          type="text"
          placeholder="Username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <button onClick={handleCreateAccount}>Create Account</button>
        <button onClick={onClose}>Cancel</button>
      </div>
    </div>
  );
};

export default CreateAccountModal;