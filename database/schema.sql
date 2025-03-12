-- 创建数据库，并设置字符集为 utf8mb4 以支持多语言
CREATE DATABASE IF NOT EXISTS photo_blog
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci;
USE photo_blog;

-- 表1：photo_cards - 存储照片卡的基本信息
CREATE TABLE IF NOT EXISTS photo_cards (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,          -- 照片标题
    description TEXT,                     -- 照片描述
    photo_url TEXT NOT NULL,              -- 照片地址
    shooting_date DATE,                   -- 拍摄日期（年月日）
    shooting_location VARCHAR(255),       -- 拍摄地点
    film_type VARCHAR(255),               -- 胶片类型
    camera VARCHAR(255),                  -- 相机
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- 创建时间
);

-- 表2：tags - 存储所有标签，保证唯一性
CREATE TABLE IF NOT EXISTS tags (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE     -- 标签名称
);

-- 表3：photo_card_tags - 关联表，建立照片卡与标签之间的多对多关系
CREATE TABLE IF NOT EXISTS photo_card_tags (
    photo_id INT NOT NULL,
    tag_id INT NOT NULL,
    PRIMARY KEY (photo_id, tag_id),
    FOREIGN KEY (photo_id) REFERENCES photo_cards(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);
