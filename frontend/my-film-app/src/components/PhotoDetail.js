import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import './PhotoDetail.css';

function PhotoDetail() {
    const { id } = useParams(); // 从 URL 获取照片 ID
    const navigate = useNavigate();
    const [photo, setPhoto] = useState(null);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState(null);
    const backendURL = 'http://localhost:8080';

    useEffect(() => {
        const fetchPhotoDetail = async () => {
            try {
                // 注意：确保后端API路径匹配 - 根据模型名可能是 photoCards 而不是 photos
                const response = await fetch(`${backendURL}/api/photos/${id}`);
                if (!response.ok) {
                    throw new Error('获取照片详情失败');
                }
                const data = await response.json();

                // 数据转换和验证
                const photoData = data; // 如果后端直接返回照片对象
                // 或者 const photoData = data.data; // 如果后端将照片对象包装在data字段中

                if (!photoData) {
                    throw new Error('返回数据格式无效');
                }

                console.log('照片详情数据:', photoData); // 调试用
                setPhoto(photoData);
            } catch (error) {
                console.error('获取照片详情出错:', error);
                setError(error.message);
            } finally {
                setIsLoading(false);
            }
        };

        fetchPhotoDetail();
    }, [id]);

    const handleBackClick = () => {
        navigate('/');
    };

    if (isLoading) {
        return <div className="photo-detail-container">加载中...</div>;
    }

    if (error) {
        return <div className="photo-detail-container">错误：{error}</div>;
    }

    if (!photo) {
        return <div className="photo-detail-container">未找到照片</div>;
    }

    return (
        <div className="photo-detail-container">
            <div className="photo-detail-content">
                {/* 添加返回箭头 */}
                <div className="photo-detail-back-arrow" onClick={handleBackClick}>◀</div>
                <div className="photo-detail-image">
                    <img src={`${backendURL}${photo.photo_url}`} alt={photo.title} />
                </div>
                <div className="photo-detail-info">
                    <h1>{photo.title}</h1>
                    <p><strong>拍摄日期:</strong> {new Date(photo.shooting_date).toLocaleDateString('zh-CN', {
                        year: 'numeric',
                        month: 'long',
                        day: 'numeric'
                    })}</p>
                    <p><strong>描述:</strong> {photo.description || '无描述'}</p>

                    {/* 胶片类型 - 新增字段 */}
                    {photo.film_type && (
                        <p><strong>胶片类型:</strong> {photo.film_type}</p>
                    )}

                    {/* 修改字段名称 - shooting_location 而不是 location */}
                    {photo.shooting_location && (
                        <p><strong>拍摄地点:</strong> {photo.shooting_location}</p>
                    )}

                    {/* 修改字段名称 - camera 而不是 camera_model */}
                    {photo.camera && (
                        <p><strong>相机型号:</strong> {photo.camera}</p>
                    )}

                    <div className="photo-detail-tags">
                        <strong>标签:</strong>
                        <div className="photo-detail-tags-container">
                            {photo.tags && photo.tags.length > 0 ? (
                                photo.tags.map(tag => (
                                    <span key={tag.id} className="photo-detail-photo-tag">{tag.name}</span>
                                ))
                            ) : (
                                <span>无标签</span>
                            )}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default PhotoDetail;