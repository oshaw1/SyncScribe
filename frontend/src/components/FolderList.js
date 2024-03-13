import React from 'react';

const FolderList = ({ folders, onFolderClick, onNoteClick }) => {
  return (
    <ul>
      {folders.map((folder) => (
        <li key={folder.id}>
          <div onClick={() => onFolderClick(folder.id)}>{folder.name}</div>
          <ul>
            {folder.notes.map((note) => (
              <li key={note.id} onClick={() => onNoteClick(note.id)}>
                {note.title}
              </li>
            ))}
          </ul>
        </li>
      ))}
    </ul>
  );
};

export default FolderList;