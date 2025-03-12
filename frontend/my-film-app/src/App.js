import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import HomePage from './pages/HomePage';
import ContactPage from './pages/ContactPage';
import MessageBoardPage from './pages/MessageBoardPage';
import PhotoDetailPage from './pages//PhotoDetailPage';

// 这个文件是默认的根组件，在这里编写自己的代码，构建 UI。
function App() { //这是一个函数组件。函数组件是一个返回 JSX 的 JavaScript 函数
  return ( //组件的 return 语句返回一个 JSX 元素
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/ContactPage" element={<ContactPage />} />
        <Route path="/MessageBoardPage" element={<MessageBoardPage />} />
        <Route path="/photos/:id" element={<PhotoDetailPage />} />

        {/* <Route path="/" element={<Home />} /> */}
        {/* <Route path="/blog" element={<Blog />} /> */}
        {/* <Route path="/blog/:id" element={<BlogPost />} /> */}
        {/* <Route path="/about" element={<About />} /> */}
      </Routes>
    </Router>
  );
}

export default App;
