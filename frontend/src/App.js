import React, { useState, useEffect } from 'react';
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

  const handleLogin = () => {
    setIsLoggedIn(true);
    setShowLoginModal(false);
  };

  const handleLogout = () => {
    setIsLoggedIn(false);
    setShowLoginModal(true);
  };

  const handleCreateAccountSuccess = (username, password) => {
    // Reset the create account modal state
    setShowCreateAccountModal(false);
    // Show the login modal
    setShowLoginModal(true);
  };

  const handleOpenCreateAccountModal = () => {
    setShowCreateAccountModal(true);
  };

  const handleCloseCreateAccountModal = () => {
    setShowCreateAccountModal(false);
  };

  useEffect(() => {
    fetchNotes();
    fetchFolders();
  }, []);

  const fetchNotes = async () => {
    const response = await axios.get('/api/notes');
    setNotes(response.data);
  };

  const fetchFolders = async () => {
    const response = await axios.get('/api/folders');
    setFolders(response.data);
  };

  const handleNoteClick = (noteId) => {
    const note = notes.find((note) => note.id === noteId);
    setSelectedNote(note);
  };

  const handleFolderClick = (folderId) => {
    // Implement the logic to handle folder click if needed
    console.log('Folder clicked:', folderId);
  };

  return (
    <div className="App">
      {!isLoggedIn && (
        <>
          {showLoginModal && <LoginModal onLogin={handleLogin} onCreateAccount={handleOpenCreateAccountModal} />}
          {showCreateAccountModal && (
            <CreateAccountModal onClose={handleCloseCreateAccountModal} onCreateSuccess={handleCreateAccountSuccess} />
          )}
        </>
      )}
      <Sidebar folders={folders} onFolderClick={handleFolderClick} onNoteClick={handleNoteClick} onLogout={handleLogout} />
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