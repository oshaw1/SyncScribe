import React, { useState } from 'react';
import axios from 'axios';
import LoginModal from './components/LoginModal';
import CreateAccountModal from './components/CreateAccountModal';
import Sidebar from './components/Sidebar';
import NoteContent from './components/NoteContent';
import Footer from './components/Footer';
import RightSidebar from './components/RightSidebar';
import Header from './components/Header';
import './App.css';

const App = () => {
  const [notes, setNotes] = useState([]);
  const [selectedNote, setSelectedNote] = useState(null);
  const [folders, setFolders] = useState([]);
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [showCreateAccountModal, setShowCreateAccountModal] = useState(false);
  const [showLoginModal, setShowLoginModal] = useState(true);
  const [sidebarStructure, setSidebarStructure] = useState({});

  const handleLogin = () => {
    setIsLoggedIn(true);
    setShowLoginModal(false);
  };

  const handleLogout = () => {
    setIsLoggedIn(false);
    setShowLoginModal(true);
  };

  const handleCreateAccountSuccess = (username, password) => {
    setShowCreateAccountModal(false);
    setShowLoginModal(true);
  };

  const handleOpenCreateAccountModal = () => {
    setShowCreateAccountModal(true);
  };

  const handleCloseCreateAccountModal = () => {
    setShowCreateAccountModal(false);
  };

  const handleNoteClick = (noteId) => {
    console.log('Note clicked:', noteId);
  };

  const handleFolderClick = (folderId) => {
    // Implement the logic to handle folder click if needed
    console.log('Folder clicked:', folderId);
  };

  return (
    <div className="App">
      {!isLoggedIn && (
        <>
          {showLoginModal && (
            <LoginModal
              onLogin={handleLogin}
              onCreateAccount={handleOpenCreateAccountModal}
              setSidebarStructure={setSidebarStructure}
            />
          )}
          {showCreateAccountModal && (
            <CreateAccountModal
              onClose={handleCloseCreateAccountModal}
              onCreateSuccess={handleCreateAccountSuccess}
            />
          )}
        </>
      )}
      <Sidebar
        folders={folders}
        onFolderClick={handleFolderClick}
        onNoteClick={handleNoteClick}
        onLogout={handleLogout}
        sidebarStructure={sidebarStructure}
        setSidebarStructure={setSidebarStructure}
      />
      <div className="main-container">
        <Header />
        <div className="main-content">
          <NoteContent note={selectedNote} />
        </div>
        <Footer />
      </div>
      <RightSidebar />
    </div>
  );
};

export default App;