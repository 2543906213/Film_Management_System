import React from "react";
import Header from '../components/Header';
import HeroImage from '../components/HeroImage';
import PageTitle from "../components/PageTitle";



function MessageBoard() {
    // 背景图片
    const MessageBoardBanner = process.env.PUBLIC_URL + '/images/MessageBoardBanner.jpg'
    // 头像图片
    const ProfileImage = process.env.PUBLIC_URL + '/images/ProfileImage.png';
    // 页面标题
    const title = "MessageBoard";
    return (

        <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100 p-6">
            <Header />        {/* 实现页眉的固定 */}
            <HeroImage imageSrc={MessageBoardBanner} /> {/* 实现背景图片随浏览器大小自动调整 */}
            <PageTitle imageSrc={ProfileImage} title={title} />      {/* 实现头像和标题的自动调整 */}
        </div>
    );
};

export default MessageBoard;
