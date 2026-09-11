# AGENTS.md

本文件为 AI 编码代理提供 <%= projectName %> 项目的导航指南。

## 项目概览

<%= description %>。基于 Go 1.22 + Gin 的 HTTP 服务，module 路径为 `<%= modulePath %>`。

## 目录结构与职责

| 目录 | 职责 |
| --- | --- |
| `cmd/server` | 程序入口，只做配置加载与服务启动，不写业务逻辑 |
| `internal/config` | 配置结构与加载逻辑（Viper），新增配置项在 `Config` 结构体中加字段 |
| `internal/router` | 路由注册中心，新增接口在 `New()` 中挂载 |
| `internal/handler` | HTTP 处理器：解析请求参数、调用 service、返回响应；禁止写业务逻辑 |
| `internal/service` | 业务逻辑层，不依赖 gin.Context |
| `internal/model` | 请求 / 响应结构体与领域模型 |
| `internal/middleware` | 通用中间件（现有日志、panic 恢复） |
| `pkg/strutil` | 可被外部复用的公共工具包示例；`pkg/` 下只放与业务无关的通用代码 |

## 代码约定

- 分层调用方向：`handler -> service -> model`，禁止反向依赖；`internal/` 内代码不允许被模块外导入。
- 每个包都要有包级注释（`// Package xxx ...`），导出标识符要有文档注释。
- 错误处理：底层返回 `error` 并用 `fmt.Errorf("...: %w", err)` 包装；handler 层统一转换为 JSON 错误响应。
- 新增配置：修改 `internal/config/config.go` 的结构体 + `config.yaml` 示例，必要时加 `v.SetDefault`。
- 新增接口：依次在 model（结构体）、service（逻辑）、handler（处理器）、router（注册）四层添加代码，并为 handler 补 httptest 测试。
- 测试命名：与被测文件同目录的 `xxx_test.go`；接口测试使用 `net/http/httptest` + `router.New()`，参考 `internal/handler/hello_test.go`。

## 常用命令

```bash
make tidy    # 首次拉取依赖
make run     # 本地运行（默认端口 <%= port %>）
make test    # 全部测试
make vet     # 静态检查
make build   # 编译到 bin/
```

提交代码前请确保 `go vet ./...` 与 `go test ./...` 全部通过，并使用 `gofmt` 格式化。
