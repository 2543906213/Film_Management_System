import React from 'react';
import { FaRobot, FaGraduationCap, FaCamera, FaBook, FaFilm, FaMusic } from 'react-icons/fa';
import './AboutMe.css';

function AboutMe() {
    const items = [
        { icon: <FaGraduationCap color="black" className="about-me-list-icon" />, text: "在读研究生" },
        { icon: <FaCamera color="black" className="about-me-list-icon" />, text: "摄影爱好者" },
        { icon: <FaBook color="black" className="about-me-list-icon" />, text: "阅读爱好者" },
        { icon: <FaFilm color="black" className="about-me-list-icon" />, text: "电影观光人" },
        { icon: <FaMusic color="black" className="about-me-list-icon" />, text: "音乐鉴赏者" },
    ];

    return (
        <div className="about-me-container">
            <div className="about-me-header">
                <FaRobot color="black" className="about-me-header-icon" />
                <span className="about-me-header-text">About me 个人简介</span>
            </div>
            <div className="about-me-list">
                {items.map((item, index) => (
                    <div key={index} className="about-me-list-item">
                        {item.icon}
                        <span>{item.text}</span>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default AboutMe;