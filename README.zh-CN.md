# <%= projectName %>

<%= description %><% if (author) { %>

作者：<%= author %><% } %>

## 技术栈

- Go 1.22
- [Gin](https://github.com/gin-gonic/gin)：HTTP Web 框架
- [Viper](https://github.com/spf13/viper)：配置加载（YAML + 环境变量覆盖）
- Makefile：常用构建 / 运行 / 测试命令

## 目录结构

```
.
├── cmd/
│   └── server/          # 服务入口（main.go）
├── internal/            # 私有业务代码（不允许被外部模块导入）
│   ├── config/          # 配置加载（Viper）
│   ├── handler/         # HTTP 处理器（控制器层）
│   ├── service/         # 业务逻辑（服务层）
│   ├── model/           # 数据结构与领域模型
│   ├── middleware/      # 中间件（日志、panic 恢复）
│   └── router/          # 路由注册
├── pkg/
│   └── strutil/         # 可复用公共工具包示例
├── config.yaml          # 配置文件示例
├── Makefile
└── go.mod               # module: <%= modulePath %>
```

## 快速开始

### 通过 AgileBuilder 生成项目

```bash
ag create --git-url <模板仓库地址> --target <目标目录>
```

按提示填写变量即可生成完整项目；生成后会自动执行 `git init` 并提交首个 commit。

### 本地运行

```bash
go mod tidy        # 拉取依赖
make run           # 或：go run ./cmd/server -config config.yaml
```

服务默认监听 `<%= port %>` 端口，验证：

```bash
curl http://localhost:<%= port %>/healthz
curl "http://localhost:<%= port %>/api/v1/hello?name=gin"
```

## 模板变量

| 变量名 | 说明 | 默认值 |
| --- | --- | --- |
| `projectName` | 项目名称（用于构建产物名、README 标题等） | `demo-service` |
| `modulePath` | Go module 路径（go.mod 的 module 行，也是内部包的导入前缀） | `github.com/example/demo-service` |
| `description` | 项目描述 | `一个 Go HTTP 服务` |
| `author` | 作者名称（可留空） | （空） |
| `port` | 服务监听端口（写入 config.yaml） | `8080` |

## 常用命令

| 命令 | 说明 |
| --- | --- |
| `make build` | 编译二进制到 `bin/` 目录 |
| `make run` | 直接运行服务 |
| `make test` | 运行全部测试（`go test ./...`） |
| `make vet` | 静态检查（`go vet ./...`） |
| `make tidy` | 整理依赖（`go mod tidy`） |
| `make clean` | 清理构建产物 |

## 配置说明

`config.yaml` 示例：

```yaml
server:
  port: <%= port %>
  mode: debug # debug | release | test
log:
  level: info
```

所有配置项均可通过环境变量覆盖，规则是把 `.` 替换为 `_` 并转大写，例如 `SERVER_PORT=9090` 覆盖 `server.port`。

## 测试

项目内置基于 `net/http/httptest` 的接口测试示例（`internal/handler/hello_test.go`）以及工具包单元测试（`pkg/strutil/strutil_test.go`）：

```bash
go test ./...
```

## License

MIT
