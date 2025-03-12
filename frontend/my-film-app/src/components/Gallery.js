import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import './Gallery.css';

function Gallery() {
    // 定义导航
    const navigate = useNavigate();
    // 首先，通过 useState 钩子定义了四个状态变量：
    // photos 用于存储照片数据，
    // selectedTags 用于存储选中的标签，
    // isLoading 用于指示数据是否正在加载，
    // error 用于存储可能发生的错误信息。
    const [photos, setPhotos] = useState([]);
    const [selectedTags, setSelectedTags] = useState(new Set());
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState(null);
    const backendURL = 'http://localhost:8080';

    // 向后端发送 GET 请求获取数据
    useEffect(() => {
        const fetchPhotos = async () => {
            try {
                const response = await fetch(`${backendURL}/api/photos`); // 替换为实际后端 URL
                if (!response.ok) {
                    throw new Error('Failed to fetch photos');
                }
                const data = await response.json();
                // console.log('Fetched data:', data); // 调试信息

                // 转换后端数据为组件所需格式
                const transformedPhotos = data.map(photo => ({
                    id: photo.id, // 保留原始ID
                    src: `${backendURL}${photo.photo_url}`,
                    width: 300, // 假设固定宽度
                    height: 400, // 假设固定高度
                    tags: photo.tags.map(tag => tag.name), // 提取标签名称
                    title: photo.title,
                    date: new Date(photo.shooting_date).toLocaleDateString('en-US', {
                        month: 'long',
                        day: 'numeric',
                        year: 'numeric'
                    }), // 格式化日期为 "January 31, 2025"
                }));
                // console.log('Transformed photos:', transformedPhotos); // 调试信息
                setPhotos(transformedPhotos);
            } catch (error) {
                console.error('Error fetching photos:', error); // 调试信息
                setError(error.message);
            } finally {
                setIsLoading(false);
            }
        };

        fetchPhotos();
    }, []);

    // 处理照片卡片点击，导航到详情页
    const handlePhotoClick = (photoId) => {
        navigate(`/photos/${photoId}`);
    };

    // 如果数据加载中，显示加载状态
    if (isLoading) {
        return <div className="gallery-container">加载中...</div>;
    }

    // 如果发生错误，显示错误信息
    if (error) {
        return <div className="gallery-container">错误：{error}</div>;
    }

    // 提取唯一标签并排序
    const uniqueTags = Array.from(
        new Set(photos.map(photo => photo.tags).flat())
    ).sort();

    // 处理标签点击，切换选择状态
    const handleTagClick = (tag) => {
        const newSelectedTags = new Set(selectedTags);
        if (newSelectedTags.has(tag)) {
            newSelectedTags.delete(tag);
        } else {
            newSelectedTags.add(tag);
        }
        setSelectedTags(newSelectedTags);
    };

    // 根据选定标签过滤照片
    const filteredPhotos = photos.filter(photo => {
        if (selectedTags.size === 0) {
            return true;
        }
        return photo.tags.some(tag => selectedTags.has(tag));
    });

    return (
        <div className="gallery-container">
            <div className="gallery-header">
                <div className="gallery-tag-filter">
                    {uniqueTags.map(tag => (
                        <button
                            key={tag}
                            className={selectedTags.has(tag) ? 'selected' : ''}
                            onClick={() => handleTagClick(tag)}
                        >
                            {tag}
                        </button>
                    ))}
                </div>
            </div>
            <div className="gallery-grid">
                {filteredPhotos.map((photo, index) => (
                    <div
                        key={index}
                        className="gallery-card"
                        onClick={() => handlePhotoClick(photo.id)}
                    >
                        {photo.src ? (
                            <img src={photo.src} alt={photo.title} className="gallery-card-image" />
                        ) : (
                            <div className="gallery-card-placeholder">{photo.title}</div>
                        )}
                        <div className="gallery-card-content">
                            <h3>{photo.title}</h3>
                            <p>{photo.date}</p>
                            {photo.tags.map(tag => (
                                <span key={tag} className="gallery-tag">
                                    <span className="gallery-tag-dot"></span> {tag}
                                </span>
                            ))}
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default Gallery;