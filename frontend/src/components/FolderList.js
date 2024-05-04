import React from 'react';

const FolderList = ({ sidebarStructure, expandedFolders, toggleFolder, onFolderClick, onNoteClick }) => {
  const renderItem = (item, level) => {
    if (item.type === 'folder') {
      const isExpanded = expandedFolders.includes(item.id);

      return (
        <div key={item.id} style={{ marginLeft: `${level * 20}px` }}>
          <div onClick={() => onFolderClick(item.id)}>
            {item.name}
            <span onClick={(e) => {
              e.stopPropagation();
              toggleFolder(item.id);
            }}>
              {isExpanded ? ' -' : ' +'}
            </span>
          </div>
          {isExpanded && item.children.map((child) => renderItem(child, level + 1))}
        </div>
      );
    } else if (item.type === 'note') {
      return (
        <div
          key={item.id}
          onClick={() => onNoteClick(item.id)}
          style={{ marginLeft: `${level * 20}px` }}
        >
          {item.name}
        </div>
      );
    }

    return null;
  };

  return (
    <div>
      {Object.values(sidebarStructure).map((item) => renderItem(item, 0))}
    </div>
  );
};

export default FolderList;