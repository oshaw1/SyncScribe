import React, { useState } from 'react';
import logo from '../images/SS.png';
import CreateNoteModal from './CreateNoteModal';

const Header = ({ onCreateSuccess }) => {
  const [isNoteModalOpen, setIsNoteModalOpen] = useState(false);

  const openNoteModal = () => {
    setIsNoteModalOpen(true);
  };

  const closeNoteModal = () => {
    setIsNoteModalOpen(false);
  };

  const handleCreateSuccess = () => {
    // Handle successful note creation
    closeNoteModal();
    onCreateSuccess(); // Notify the parent component about the successful note creation
  };

  return (
    <div className="header">
      <img src={logo} alt="SyncScribe Logo" className="logo" />
      <span className="app-name">Sync Scribe</span>
      <div className="create-note-button-div">
        <button className="create-note-button" onClick={openNoteModal}>
          Create Note
        </button>
      </div>
      {isNoteModalOpen && (
        <CreateNoteModal
          onClose={closeNoteModal}
          onCreateSuccess={handleCreateSuccess}
        />
      )}
    </div>
  );
};

export default Header;