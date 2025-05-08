<h1 align="center">
  📝 mcp-ip-geo
</h1>

<div align="center">
  <a href="./README.md">English</a> | 中文
</div>

---

`mcp-ip-geo` 是一个 `MCP` (`Model Context Protocol`) 服务器，基于 `ip-api.com` 提供 `IP` 地理位置查询服务，可获取 `IP` 地址对应的国家、省份、城市等地理信息。

# 🔌 MCP 集成

您可以通过以下两种方式集成 `mcp-ip-geo` 服务：

## 方式一：使用 go run 命令集成（Stdio 标准输入输出）

通过在 `MCP` 配置中添加以下内容，可直接从 `GitHub` 运行最新版本：

```json
{
  "mcpServers": {
    "mcp-ip-geo": {
      "command": "go run github.com/chenmingyong0423/mcp-ip-geo/cmd/mcp-ip-geo@latest"
    }
  }
}
```

## 方式二：使用 Docker 命令部署服务（Streamable HTTP）

### 🐳 使用 Docker 部署

#### 步骤 1: 克隆仓库

```bash
git clone https://github.com/chenmingyong0423/mcp-ip-geo.git
cd mcp-ip-geo
```

#### 步骤 2: 构建 Docker 镜像

```bash
docker build -t mcp-ip-geo-server .
```

#### 步骤 3: 运行容器

```bash
docker run -d --name mcp-ip-geo-server -p 8000:8000 mcp-ip-geo-server
```

成功运行后，服务将在容器内以 `0.0.0.0:8000` 监听（即监听所有网络接口），可通过 `http://<服务器地址>:8000/mcp` 访问，其中`<服务器地址>`可以是：
- 本地开发环境：使用 `localhost` 或 `127.0.0.1`
- 局域网环境：使用服务器的内网IP地址（如 `192.168.x.x`）
- 公网环境：使用服务器的公网 `IP` 地址或域名

> 注意：服务在容器内配置为监听 `0.0.0.0` 地址，这是容器化应用的标准做法，确保服务可以从容器外部访问。

#### 步骤 4: 配置 MCP

在 `MCP` 配置中添加以下内容：

```json
{
  "mcpServers": {
    "mcp-ip-geo": {
      "url": "http://<服务器地址>:8000/mcp"
    }
  }
}
```

请将 `<服务器地址>` 替换为实际部署环境的服务器 `IP` 地址或域名。

# ⚠️ 许可说明

> 注意：本项目使用了 ip-api.com 免费版本，其 API 服务**仅限非商业用途**。若您打算将本项目用于商业目的，请务必遵守其服务条款，或购买其付费版本：https://ip-api.com/
