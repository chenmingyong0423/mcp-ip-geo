<h1 align="center">
  mcp-ip-geo
</h1>

<div align="center">
  <a href="./README.md">English</a>
</div>

---

`mcp-ip-geo` 是一个 MCP 服务器，基于 `ip-api.com` 提供 `IP` 地理位置查询服务（国家、省份、城市等）。

# 从源码构建

## 本地构建

### 使用 Go 命令

```bash
# 在类 Unix 系统（Linux/macOS）上
go build -o mcp-ip-geo ./cmd/mcp-ip-geo

# 在 Windows 上
go build -o mcp-ip-geo.exe .\cmd\mcp-ip-geo
```

### 使用 Docker

1. 构建 Docker 镜像：

    ```bash
    docker build -t mcp-ip-geo-server .
    ```

2. 运行容器：

    ```bash
    docker run -d --name mcp-ip-geo-server -p 8000:8000 mcp-ip-geo-server
    ```

## 安装预编译版本

使用 `Go` 安装最新版本的服务：

```bash
go install github.com/chenmingyong0423/mcp-ip-geo/cmd/mcp-ip-geo@latest
```

# MCP 集成

你可以通过以下两种方式集成 `mcp-ip-geo` 服务：

## 🖥 可执行文件集成（本地运行）

```json
{
  "mcpServers": {
    "mcp-ip-geo": {
      "command": "/path/to/mcp-ip-geo"
    }
  }
}
```

## 🌐 HTTP 接口集成（连接到已运行的服务）

```json
{
  "mcpServers": {
    "mcp-ip-geo": {
      "url": "http://host:port/mcp"
    }
  }
}
```

# License
> ⚠ 注意：本项目使用了 ip-api.com 免费版本，其 `API` 服务仅限非商业用途。若您打算将本项目用于商业目的，请务必遵守其服务条款，或购买其付费版本：https://ip-api.com/
