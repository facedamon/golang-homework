# 博客 API 项目

这是一个基于 Go 语言开发的 RESTful API 博客系统，提供用户认证、文章管理和评论功能。

## 项目特性

- 用户注册和登录
- JWT 认证
- 文章的增删改查
- 评论功能
- 优雅的服务器关闭
- 日志记录
- 数据库连接池管理

## 技术栈

- **后端框架**: Gin (HTTP Web 框架)
- **数据库 ORM**: GORM
- **数据库**: MySQL
- **认证**: JWT (JSON Web Token)
- **配置管理**: Viper
- **日志**: Logrus
- **Go 版本**: 1.25.0

## 项目结构

```
├── main.go           # 应用程序入口
├── go.mod           # Go 模块文件
├── go.sum           # 依赖锁文件
├── config/          # 配置相关
│   ├── config.go    # 配置初始化
│   ├── config.yml   # 配置文件
│   ├── db.go        # 数据库配置
│   └── logrus.go    # 日志配置
├── global/          # 全局变量
│   └── global.go    # 全局变量定义
├── handler/         # 处理器层
│   ├── auth.go      # 认证相关接口
│   ├── comment.go   # 评论相关接口
│   └── post.go      # 文章相关接口
├── middleware/      # 中间件
│   ├── auth_middleware.go    # 认证中间件
│   └── error_middleware.go  # 错误处理中间件
├── model/           # 数据模型
│   └── model.go     # 数据库模型定义
├── router/          # 路由配置
│   └── router.go    # 路由初始化
└── utils/           # 工具函数
    └── utils.go     # 通用工具函数
```

## 运行环境要求

- Go 1.25.0 或更高版本
- MySQL 5.7 或更高版本
- Git (用于克隆项目)

## 安装步骤

### 1. 克隆项目

```bash
git clone https://github.com/facedamon/golang-homework.git
cd golang-homework/task4
```

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 配置数据库

#### 创建 MySQL 数据库

```sql
CREATE DATABASE golang_test CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

#### 修改配置文件

编辑 `config/config.yml` 文件，根据您的数据库配置修改连接信息：

```yaml
database:
  dsn: root:root@tcp(127.0.0.1:3306)/golang_test?charset=utf8&parseTime=true&loc=Local
  maxIdleConns: 10
  maxOpenConns: 100
app:
  port: :8081
```

参数说明：
- `dsn`: 数据库连接字符串，格式为 `用户名:密码@tcp(host:port)/数据库名?参数`
- `maxIdleConns`: 最大空闲连接数
- `maxOpenConns`: 最大打开连接数
- `port`: 应用服务端口

### 4. 启动项目

```bash
go run main.go
```

或者构建后运行：

```bash
go build -o blog.exe
./blog.exe
```

启动成功后，服务将在 `http://localhost:8081` 运行。

## API 接口文档

### 认证接口

#### 用户注册
- **POST** `/api/auth/register`
- 请求体:
```json
{
  "username": "用户名",
  "password": "密码",
  "email": "邮箱"
}
```

#### 用户登录
- **POST** `/api/auth/login`
- 请求体:
```json
{
  "username": "用户名",
  "password": "密码"
}
```

### 文章接口 (需要认证)

#### 创建文章
- **POST** `/api/bus/createPost`
- 请求头: `Authorization: Bearer <token>`
- 请求体:
```json
{
  "title": "文章标题",
  "content": "文章内容"
}
```

#### 获取所有文章
- **GET** `/api/bus/getAllPostList`
- 请求头: `Authorization: Bearer <token>`

#### 根据ID获取文章
- **GET** `/api/bus/getPostById/:id`
- 请求头: `Authorization: Bearer <token>`

#### 更新文章
- **POST** `/api/bus/updatePost`
- 请求头: `Authorization: Bearer <token>`
- 请求体:
```json
{
  "id": 1,
  "title": "新标题",
  "content": "新内容"
}
```

#### 删除文章
- **GET** `/api/bus/deletePostById/:id`
- 请求头: `Authorization: Bearer <token>`

### 评论接口 (需要认证)

#### 创建评论
- **POST** `/api/bus/createComment`
- 请求头: `Authorization: Bearer <token>`
- 请求体:
```json
{
  "content": "评论内容",
  "postId": 1
}
```

#### 获取文章评论
- **GET** `/api/bus/getCommentsByPostId/:id`
- 请求头: `Authorization: Bearer <token>`

## 中间件说明

项目中使用了两个重要的中间件来处理认证和错误：

### 1. 认证中间件 (AuthMiddleWare)

**文件位置**: `middleware/auth_middleware.go`

**功能说明**:
- 验证请求头中的 JWT Token
- 解析 Token 获取用户信息
- 将用户信息存储到上下文中供后续处理器使用

**工作流程**:
1. 从请求头 `Authorization` 中获取 Token
2. 验证 Token 是否存在，格式是否正确（Bearer + JWT）
3. 使用 `utils.ParseJWT()` 解析 Token 获取用户名
4. 根据用户名查询数据库验证用户是否存在
5. 将用户信息存储到 Gin 上下文中，供后续处理器使用
6. 如果任何步骤失败，返回 401 未授权错误并中止请求

**使用方式**:
```go
// 应用到需要认证的路由组
api := g.Group("/api/bus")
api.Use(middleware.AuthMiddleWare())
```

**错误处理**:
- Token 缺失：返回"未授权"错误
- Token 无效：返回"未授权"错误并包含具体错误信息
- 用户不存在：返回"未授权"错误

### 2. 错误处理中间件 (ErrorHandler)

**文件位置**: `middleware/error_middleware.go`

**功能说明**:
- 统一处理应用中的错误响应
- 区分自定义错误和系统错误
- 提供统一的错误响应格式

**工作流程**:
1. 在所有处理器执行完成后检查是否有错误
2. 遍历 Gin 上下文中的错误列表
3. 判断错误类型：
   - 如果是自定义错误类型 (`global.R`)，直接返回
   - 如果是系统错误，包装为统一格式返回
4. 返回 JSON 格式的错误响应

**使用方式**:
```go
// 全局应用到所有路由
g := gin.Default()
g.Use(middleware.ErrorHandler())
```

**错误响应格式**:
```json
{
  "code": 500,
  "msg": "服务器异常",
  "data": "具体错误信息"
}
```

**自定义错误**:
项目中定义了 `global.R` 结构体作为统一的响应格式：
```go
type R struct {
    Code int         `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
}
```

可以使用 `global.NewError()` 创建自定义错误：
```go
err := global.NewError(400, "参数错误")
ctx.Error(err)  // 将错误添加到上下文中
```

### JWT Token 处理

**Token 生成** (`utils.GenerateJWT`):
- 使用 HMAC SHA256 算法
- 包含用户名和过期时间（24小时）
- 密钥为 "secret"
- 返回格式：`Bearer <jwt_token>`

**Token 解析** (`utils.ParseJWT`):
- 验证 Bearer 前缀
- 解析 JWT Token
- 验证签名和有效期
- 返回用户名

**密码处理**:
- 使用 bcrypt 算法加密密码（cost=12）
- 提供密码验证功能

## 开发指南

### 数据库迁移

项目启动时会自动执行数据库迁移，创建必要的表结构：
- users (用户表)
- posts (文章表)
- comments (评论表)

### 日志文件

应用运行时会生成日志文件，记录应用的运行状态和错误信息。

### 优雅关闭

应用支持优雅关闭，当收到 `SIGINT` 或 `SIGTERM` 信号时，会等待正在处理的请求完成后再关闭服务。

## 故障排除

### 常见问题

1. **数据库连接失败**
   - 检查 MySQL 服务是否启动
   - 验证配置文件中的数据库连接信息
   - 确保数据库用户有足够的权限

2. **端口被占用**
   - 修改 `config/config.yml` 中的端口配置
   - 或者终止占用端口的进程

3. **依赖下载失败**
   - 配置 Go 代理：`go env -w GOPROXY=https://goproxy.cn,direct`
   - 或使用其他可用的代理

### 调试模式

开发时可以设置环境变量来启用调试模式：

```bash
export GIN_MODE=debug
go run main.go
```

## 贡献

欢迎提交 Issue 和 Pull Request 来改进项目。

## 许可证

本项目采用 MIT 许可证。