# Gin-Scaffold: 工程化 Gin 后端脚手架

![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue)
![Gin Version](https://img.shields.io/badge/Gin-1.1%2B-green)
![License](https://img.shields.io/badge/License-MIT-yellow)

## 项目概述
Gin-Scaffold 是一个基于 Gin 框架构建的工程化后端脚手架，整合了依赖注入、日志、缓存、认证、限流等核心组件，提供标准化的项目结构和最佳实践，旨在帮助开发者快速搭建生产级别的 Go 后端服务。

### 核心特性
- **工程化架构**：严格的分层设计，支持模块化扩展，符合 Go 项目最佳实践。
- **依赖注入**：基于 Google Wire 实现组件解耦，简化测试和维护。
- **配置管理**：基于 Viper 实现高效配置解析和管理。
- **安全防护**：集成 JWT 认证（含黑名单机制）、请求限流中间件。
- **高效缓存**：Redis 连接工厂支持单机多 DB、集群模式，适配不同部署场景。
- **日志系统**：Zap + Lumberjack 实现结构化日志输出与日志轮转。
- **数据访问**：GORM 通用 Repository 封装，简化数据库操作。
- **全局 ID**：基于 SonyflakeX 生成分布式唯一 ID，避免 ID 冲突。
- **丰富中间件**：内置日志、多规则限流、Auth认证、CORS、Gzip、TraceID等常用中间件，支持自定义扩展。

## 技术栈
| 类别         | 技术选型                           |
|--------------|--------------------------------|
| Web 框架     | Gin 1.10+                      |
| 依赖注入     | Google Wire                    |
| 缓存         | Redis (go-redis/v9)            |
| 日志         | Zap + Lumberjack               |
| 认证         | JWT (golang-jwt/jwt)           |
| ORM          | GORM 1.3+                      |
| 分布式 ID    | SonyflakeX                     |
| 配置管理     | Viper (可选扩展)                   |
| 限流         | 基于 Redis/内存的限流算法               |
| 部署         | Docker + Docker Compose (可选扩展) |

## 快速开始

### 环境准备
- Go 1.23.5+
- Redis 7.2+ (用于缓存、限流、JWT 黑名单)
- MySQL 8.0+ (可选，用于数据持久化)

### 安装依赖
```bash
# 克隆项目
git clone https://github.com/0xjasoncao/gin-scaffold
cd gin-scaffold

# 安装依赖
go mod download	

# 项目启动
go build main.go
main start -c 配置文件的文件夹地址

