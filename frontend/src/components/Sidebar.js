import React from 'react';
import FolderList from './FolderList';

const Sidebar = ({ folders, onFolderClick, onNoteClick }) => {
  return (
    <div className="sidebar">
      <div className="account">
        <h2>Account</h2>
      </div>
      <div className="padding-div"></div>
      <div className="folder-housing">
        <FolderList
        folders={folders}
        onFolderClick={onFolderClick}
        onNoteClick={onNoteClick}
      /></div>
    </div>
  );
};

export default Sidebar;