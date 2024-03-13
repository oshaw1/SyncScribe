import React from 'react';
import FolderList from './FolderList';

const Sidebar = ({ folders, onFolderClick, onNoteClick }) => {
  return (
    <div className="sidebar">
      <h2>Account</h2>
      <FolderList
        folders={folders}
        onFolderClick={onFolderClick}
        onNoteClick={onNoteClick}
      />
    </div>
  );
};

export default Sidebar;