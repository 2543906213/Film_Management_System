// MainLayout.js
import React from 'react';
import AboutMe from './AboutMe';
import Navigator from './Navigation';
import Content from './Content';
import './MainLayout.css';

function MainLayout() {
    return (
        <div className="grid-container">
            <div className="left-column">
                <div className="left-top"> < AboutMe /></div>
                <div className="left-bottom"><Navigator /></div>
            </div>
            <div className="right-column"><Content /></div>
        </div>
    );
}

export default MainLayout;
