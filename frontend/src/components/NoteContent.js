import React, { useEffect, useState } from 'react';
import axios from 'axios';
import TextEditor from './TextEditor';

const NoteContent = ({ noteId }) => {
  const [noteContent, setNoteContent] = useState('');
  const [noteTitle, setNoteTitle] = useState('');

  useEffect(() => {
    if (noteId) {
      axios.get(`http://localhost:8080/notes/getNote?noteId=${noteId}`)
        .then(response => {
          const note = response.data;
          if(note.Content == ""){
            setNoteContent("Type here to start making notes.")
          }
          setNoteContent(note.Content);
          setNoteTitle(note.Title);
        })
        .catch(error => {
          console.error('Error fetching note:', error);
        });
    }
  }, [noteId]);

  return (
    <div>
        <TextEditor
        title={noteTitle}
        setTitle={setNoteTitle}
        content={noteContent}
        setContent={setNoteContent}
      />
    </div>
  );
};

export default NoteContent;