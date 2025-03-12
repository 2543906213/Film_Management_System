/* Navigation.js */
import React from 'react';
import { FaCompass } from 'react-icons/fa';
import './Navigation.css';

function Navigation() {
    return (
        <div className="navigation-container">
            <div className="navigation-header">
                <FaCompass color="black" className="navigation-header-icon" />
                <span className="navigation-header-text">Navigation 导航</span>
            </div>
        </div>
    );
};

export default Navigation;
