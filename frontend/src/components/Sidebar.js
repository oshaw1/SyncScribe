import React, { useState } from 'react';
import FolderList from './FolderList';
import profilepicture from '../images/AccountPicture.jpg';
import AccountSettingsModal from './AccountSettingsModal';

const Sidebar = ({ folders, onFolderClick, onNoteClick, onLogout }) => {
  const [showModal, setShowModal] = useState(false);

  const handleAccountClick = () => {
    setShowModal(true);
  };

  const handleCloseModal = () => {
    setShowModal(false);
  };

  return (
    <div className="sidebar">
      <div className="account">
        <button className="account-button" onClick={handleAccountClick}>
          <img src={profilepicture} alt="Account" className="account-picture" />
          Account
        </button>
      </div>
      {showModal && <AccountSettingsModal onClose={handleCloseModal} onLogout={onLogout} />}
      <div className="padding-div"></div>
      <div className="folder-housing">
        <FolderList folders={folders} onFolderClick={onFolderClick} onNoteClick={onNoteClick} />
      </div>
    </div>
  );
};

export default Sidebar;