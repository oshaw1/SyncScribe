import React, { useState } from 'react';
import axios from 'axios';
import './PopupModal.css';

const AccountSettingsModal = ({ onClose, onLogout }) => {
  const [error, setError] = useState('');
  const [errorColor, setErrorColor] = useState('red');
  const [confirmDelete, setConfirmDelete] = useState(false);

  const handleDeleteUser = async () => {
    try {
      const token = localStorage.getItem('token');
      await axios.delete('http://localhost:8080/users/delete', {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      // Display success message
      setError('User deleted successfully');
      setErrorColor('green');
      // Clear the token from local storage
      localStorage.removeItem('token');
      // Close the account settings modal and trigger logout after a short delay
      setTimeout(() => {
        onClose();
        onLogout();
      }, 2000);
    } catch (error) {
      console.error('Error deleting user:', error);
      // Display error message
      setError('An error occurred while deleting the user. Please try again.');
      setErrorColor('red');
    }
  };

  return (
    <div className="account-settings-modal">
      <div className="login-form">
        <h2>Account Settings</h2>
        {error && <p className="error-message" style={{ color: errorColor }}>{error}</p>}
        {!confirmDelete ? (
          <>
            <p>Delete Account</p>
            <button className="delete-button" onClick={() => setConfirmDelete(true)}>
              Delete Account
            </button>
            <button onClick={onClose}>Cancel</button>
          </>
        ) : (
          <>
            <p>This action cannot be undone. Please confirm deletion.</p>
            <button className="confirm-delete-button" onClick={handleDeleteUser}>
              Confirm Deletion
            </button>
            <button onClick={() => setConfirmDelete(false)}>Cancel</button>
          </>
        )}
      </div>
    </div>
  );
};

export default AccountSettingsModal;