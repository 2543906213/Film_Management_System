# 胶片摄影管理系统

## 项目概述

胶片摄影管理系统是一个全栈应用程序，专为展示和管理胶片摄影作品而设计。该系统由Go语言编写的后端API和React开发的前端页面组成，支持照片作品的展示、标签分类、详情查看等功能。

## 功能特点

### 前端功能
- 响应式照片墙展示
- 作品详情页面查看
- 按标签筛选照片
- 留言板功能
- 联系页面
- 个人简介展示

### 后端功能
- RESTful API 接口
- 照片作品的CRUD操作
- 标签管理系统
- 评论管理
- 文件上传与管理
- MySQL数据存储

## 技术栈

### 后端
- 语言：[Go](https://go.dev/) 1.23.4
- Web框架：[Gin](https://github.com/gin-gonic/gin) 1.10.0
- ORM工具：[GORM](https://gorm.io/) 1.25.12
- 数据库：MySQL
- 其他依赖：
  - CORS支持：gin-contrib/cors
  - 数据库驱动：gorm.io/driver/mysql

### 前端
- 框架：[React](https://reactjs.org/) 19.0.0
- 路由：[React Router](https://reactrouter.com/) 7.2.0
- HTTP客户端：[Axios](https://axios-http.com/) 1.7.9
- UI组件：[React Icons](https://react-icons.github.io/react-icons/) 5.5.0
- 照片展示：react-photo-album 3.0.2
- 构建工具：Create React App

## 系统架构

```
Film_Management_System
├── backend/                # Go后端
│   ├── internal/           # 内部包
│   │   ├── config/         # 配置相关
│   │   ├── controllers/    # 控制器
│   │   ├── database/       # 数据库连接和操作
│   │   ├── models/         # 数据模型
│   │   └── routes/         # 路由定义
│   ├── pkg/                # 公共包
│   │   └── utils/          # 工具函数
│   └── main.go             # 程序入口
└── frontend/               # React前端
    └── my-film-app/        # React应用
        ├── public/         # 静态资源
        └── src/            # 源代码
            ├── components/ # 组件
            └── pages/      # 页面
```

## 快速开始

### 后端安装

1. 确保安装了Go 1.23.4或更高版本
2. 配置MySQL数据库
3. 克隆仓库并进入后端目录

```bash
git clone https://github.com/yourusername/Film_Management_System.git
cd Film_Management_System/backend
```

4. 修改配置文件 config.json，更新数据库连接信息

5. 运行后端服务

```bash
go run main.go
```

服务默认运行在 `http://localhost:8080`

### 前端安装

1. 确保安装了Node.js和npm
2. 进入前端目录

```bash
cd Film_Management_System/frontend/my-film-app
```

3. 安装依赖

```bash
npm install
```

4. 启动开发服务器

```bash
npm start
```

应用默认运行在 `http://localhost:3000`

## API接口说明

### 照片相关接口

| 方法   | 端点              | 描述               | 参数                                                    |
|-------|-------------------|-------------------|--------------------------------------------------------|
| GET   | /api/photos       | 获取所有照片         | 无                                                     |
| GET   | /api/photos/:id   | 获取特定照片         | id: 照片ID                                              |
| POST  | /api/photos       | 创建新照片           | title, description, photoFile, shooting_date, tags等    |
| PUT   | /api/photos/:id   | 更新照片信息         | id: 照片ID, 其他同POST                                  |
| DELETE| /api/photos/:id   | 删除照片            | id: 照片ID                                              |

### 评论相关接口

| 方法   | 端点              | 描述               | 参数                                                    |
|-------|-------------------|-------------------|--------------------------------------------------------|
| GET   | /api/comments     | 获取评论           | photo_id (可选): 指定照片的评论                          |
| POST  | /api/comments     | 添加评论           | photo_id, content, user_name                           |
| DELETE| /api/comments/:id | 删除评论           | id: 评论ID                                              |

## 数据模型

### 照片卡(PhotoCard)
- ID: 唯一标识符
- Title: 标题
- Description: 描述
- PhotoURL: 照片URL
- ShootingDate: 拍摄日期
- ShootingLocation: 拍摄地点
- FilmType: 胶片类型
- Camera: 相机型号
- CreatedAt: 创建时间
- Tags: 关联标签

### 标签(Tag)
- ID: 唯一标识符
- Name: 标签名称
- PhotoCards: 关联的照片卡

### 评论(Comment)
- ID: 唯一标识符
- PhotoID: 关联照片ID
- UserName: 用户名
- Content: 评论内容
- CreatedAt: 创建时间

## 贡献指南

1. Fork本仓库
2. 创建您的特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交您的更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 打开Pull Request

## 许可证

[选择适当的开源许可证，如MIT, Apache 2.0等]

## 联系方式

项目开发者 - [里包恩] - [2543906213@qq.com]

项目链接: [https://github.com/2543906213/Film_Management_System](https://github.com/2543906213/Film_Management_System)

## 致谢

- 感谢所有贡献者和项目依赖的开源社区
- 特别感谢[列出特别感谢的人或组织]