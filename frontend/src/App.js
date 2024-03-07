import React, { useState, useEffect } from 'react';
import axios from 'axios';
import NoteList from './components/NoteList';
import './App.css';

const App = () => {
  const [notes, setNotes] = useState([]);
  const [newNote, setNewNote] = useState('');

  useEffect(() => {
    fetchNotes();
  }, []);

  const fetchNotes = async () => {
    const response = await axios.get('/api/notes');
    setNotes(response.data);
  };

  const handleInputChange = (event) => {
    setNewNote(event.target.value);
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    await axios.post('/api/notes', { content: newNote });
    setNewNote('');
    fetchNotes();
  };

  return (
    <div>
      <h1>Sync Scribe</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Enter a new note"
          value={newNote}
          onChange={handleInputChange}
        />
        <button type="submit">Add Note</button>
      </form>
      <NoteList notes={notes} />
    </div>
  );
};

export default App;