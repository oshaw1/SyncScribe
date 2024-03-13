import React from 'react';

const NoteContent = ({ note }) => {
  if (!note) {
    return <div>Select a note to view its contents.</div>;
  }

  return (
    <div className="note-content">
      <h2>{note.title}</h2>
      <p>{note.content}</p>
    </div>
  );
};

export default NoteContent;