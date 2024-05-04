import React, { useState, useEffect } from 'react';
import FolderList from './FolderList';
import profilepicture from '../images/AccountPicture.jpg';
import AccountSettingsModal from './AccountSettingsModal';

const Sidebar = ({ onFolderClick, onNoteClick, onLogout, sidebarStructure, setSidebarStructure }) => {
  const [showModal, setShowModal] = useState(false);
  const [expandedFolders, setExpandedFolders] = useState([]);

  const handleAccountClick = () => {
    setShowModal(true);
  };

  const handleCloseModal = () => {
    setShowModal(false);
  };

  useEffect(() => {
    console.log('Sidebar structure updated:', sidebarStructure);
  }, [sidebarStructure]);

  const toggleFolder = (folderId) => {
    if (expandedFolders.includes(folderId)) {
      setExpandedFolders(expandedFolders.filter((id) => id !== folderId));
    } else {
      setExpandedFolders([...expandedFolders, folderId]);
    }
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
        <FolderList
          sidebarStructure={sidebarStructure}
          expandedFolders={expandedFolders}
          toggleFolder={toggleFolder}
          onFolderClick={onFolderClick}
          onNoteClick={onNoteClick}
        />
      </div>
    </div>
  );
};

export default Sidebar;