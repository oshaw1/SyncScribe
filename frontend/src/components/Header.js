import React from 'react';
import logo from '../images/SS.png';

const Header = () => {
    return (
      <div className="header">
        <img src={logo} alt="SyncScribe Logo" className="logo" />
        <span className="app-name">Sync Scribe</span>
        <div className="create-note-button-div">
            <button className="create-note-button">Create Note</button>
        </div>
      </div>
    );
  };
  
  export default Header;