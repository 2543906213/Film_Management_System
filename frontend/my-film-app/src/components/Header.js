import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import './Header.css';

const Header = () => {
  // 初始化 navigate
  const navigate = useNavigate();
  // 用于存储搜索输入框的内容
  const [searchQuery, setSearchQuery] = useState('');
  // 用于导航到不同的页面
  const [showSearchInput, setShowSearchInput] = useState(false);

  // 切换搜索输入框的显示状态
  const handleSearchClick = () => {
    setShowSearchInput(!showSearchInput);
  };

  // 处理搜索表单的提交事件，阻止默认提交行为，并在控制台输出搜索内容。
  const handleSearchSubmit = (e) => {
    e.preventDefault();
    // 这里可以处理全站搜索逻辑，例如跳转到搜索结果页面
    console.log("搜索内容:", searchQuery);
    // 示例：跳转到 /search 页面并传递查询参数
    // navigate(`/search?q=${encodeURIComponent(searchQuery)}`);
  };

  return (
    <header className="header">
      <div className="header-container">

        {/* 左侧：HOME 链接 */}
        <div className="header-left">
          <Link to="/" className="home-link">HOME</Link>
        </div>

        {/* 右侧：三个按钮 */}
        <div className="header-right">
          {/* 第一个按钮：联系我 */}
          <button className="header-button" onClick={() => navigate('/contactpage')}>
            <span className="icon">📞</span> Contact
          </button>
          {/* 第二个按钮：留言板 */}
          <button className="header-button" onClick={() => navigate('/messageboardpage')}>
            <span className="icon">💬</span> MessageBoard
          </button>
          {/* 第三个按钮：搜索按钮 */}
          <button className="header-button" onClick={handleSearchClick}>
            <span className="icon">🔍</span> Search
          </button>
          {/* 隐藏：搜索框 */}
          {showSearchInput && (
            <form onSubmit={handleSearchSubmit} className="search-form">
              <input
                type="text"
                placeholder="输入搜索内容..."
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                className="search-input"
              />
            </form>
          )}
        </div>
      </div>
    </header>
  );
};

export default Header;
