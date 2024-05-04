import React, { useState } from 'react';
import axios from 'axios';
import './PopupModal.css';

const CreateNoteModal = ({ onClose, onCreateSuccess }) => {
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [tags, setTags] = useState('');
  const [error, setError] = useState('');

  const handleCreateNote = async () => {
    try {
      const token = localStorage.getItem('token');
      const userID = localStorage.getItem('userID');
  
      const response = await axios.post('http://localhost:8080/notes/create', {
        title,
        content,
        tags: tags.split(',').map((tag) => tag.trim()),
        userID,
      }, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
  
      console.log('Note created successfully:', response.data);
      onCreateSuccess();
    } catch (error) {
      console.error('Error creating note:', error);
      setError('An error occurred while creating the note. Please try again.');
    }
  };

  return (
    <div className="create-note-modal">
      <div className="note-form">
        <h2>Create Note</h2>
        {error && <p className="error-message">{error}</p>}
        <input
          type="text"
          placeholder="Title"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
        <textarea
          placeholder="Content"
          value={content}
          onChange={(e) => setContent(e.target.value)}
        ></textarea>
        <input
          type="text"
          placeholder="Tags (comma-separated)"
          value={tags}
          onChange={(e) => setTags(e.target.value)}
        />
        <button onClick={handleCreateNote}>Create Note</button>
        <button onClick={onClose}>Cancel</button>
      </div>
    </div>
  );
};

export default CreateNoteModal;