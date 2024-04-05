import React from 'react';
import FolderList from './FolderList';
import profilepicture from '../images/AccountPicture.jpg';

const Sidebar = ({ folders, onFolderClick, onNoteClick }) => {
  return (
    <div className="sidebar">
      <div className="account">
        <button className="account-button">
         <img src={profilepicture} alt="Account" className="account-picture" />
          Account
        </button>
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