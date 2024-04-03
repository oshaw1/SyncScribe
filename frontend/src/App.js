import React, { useState, useEffect } from 'react';
import axios from 'axios';
import LoginModal from './components/LoginModal';
import Sidebar from './components/Sidebar';
import NoteContent from './components/NoteContent';
import Footer from './components/Footer';
import RightSidebar from './components/RightSidebar';
import logo from './images/SS.png';
import './App.css';

const App = () => {
  const [notes, setNotes] = useState([]);
  const [selectedNote, setSelectedNote] = useState(null);
  const [folders, setFolders] = useState([]);
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  const handleLogin = () => {
    // Perform login logic and update the isLoggedIn state
    setIsLoggedIn(true);
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
      {!isLoggedIn && <LoginModal onLogin={handleLogin} />}
      <Sidebar folders={folders} onFolderClick={handleFolderClick} onNoteClick={handleNoteClick} />
      <div className="main-container">
        <div className="header">
          <img src={logo} alt="SyncScribe Logo" className="logo" />
          <span className="app-name">Sync Scribe</span>
        </div>
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