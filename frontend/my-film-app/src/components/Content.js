/* Content.js */
import React from 'react';
import { FaCameraRetro } from 'react-icons/fa';
import Gallery from './Gallery'
import './Content.css';

function Content() {
    return (
        <div className="content-container">
            <div className="content-header">
                <FaCameraRetro color="black" className="content-header-icon" />
                <span className="content-header-text">Gallery 照片墙</span>
            </div>
            <Gallery />
        </div>
    );
}

export default Content;
