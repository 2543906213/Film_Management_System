import React from 'react';
import './PageTitle.css';

const PageTitle = ({ imageSrc, title }) => {
    return (
        <div className="pagetitle-container">
            <img
                src={imageSrc}
                alt="头像"
                className="pagetitle-avatar"
            />
            <div className="pagetitle-title">{title}</div>

        </div>
    );
};

export default PageTitle;
