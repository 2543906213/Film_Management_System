import React from 'react';
import Header from '../components/Header';
import PhotoDetail from '../components/PhotoDetail';

function PhotoDetailPage() {
    return (
        <>
            <Header />        {/* 实现页眉的固定 */}
            <PhotoDetail />
        </>
    );
}
export default PhotoDetailPage