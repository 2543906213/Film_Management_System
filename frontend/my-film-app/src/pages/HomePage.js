import React from 'react';
import Header from '../components/Header';
import HeroImage from '../components/HeroImage';
import PageTitle from '../components/PageTitle';
import IntroSection from '../components/IntroSection';
import MainLayout from '../components/MainLayout';

function Home() {
    // 背景图片
    const HomeBanner = process.env.PUBLIC_URL + '/images/HomeBanner.jpg';
    // 头像图片
    const ProfileImage = process.env.PUBLIC_URL + '/images/ProfileImage.png';
    // 页面标题
    const title = "我的摄影作品";
    return (
        <>
            <Header />        {/* 实现页眉的固定 */}
            <HeroImage imageSrc={HomeBanner} />     {/* 实现背景图片随浏览器大小自动调整 */}
            <PageTitle imageSrc={ProfileImage} title={title} />       {/* 实现头像和标题的自动调整 */}
            <IntroSection />  {/* 介绍语，能够实现缓慢出现的功能 */}
            <MainLayout />    {/* 主体内容，分为左侧导航栏+右侧内容栏 */}
        </>
    );
}

export default Home;