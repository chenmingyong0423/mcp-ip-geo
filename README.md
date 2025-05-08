<h1 align="center">
  📝 mcp-ip-geo
</h1>

<div align="center">
  English | <a href="./README-zh_CN.md">简体中文</a>
</div>

---

`mcp-ip-geo` is an `MCP` (`Model Context Protocol`) server that provides IP geolocation lookup services (country, region, city, etc.) using the ip-api.com service.

# 🔌 MCP Integration

You can integrate the `mcp-ip-geo` service in two ways:

## Method 1: Using the go run command (Stdio)

Add the following to your `MCP` configuration to run the latest version directly from GitHub:

```json
{
  "mcpServers": {
    "mcp-ip-geo": {
      "command": "go run github.com/chenmingyong0423/mcp-ip-geo/cmd/mcp-ip-geo@latest"
    }
  }
}
```

## Method 2: Using Docker (Streamable HTTP)

### 🐳 Docker Deployment

#### Step 1: Clone the repository

```bash
git clone https://github.com/chenmingyong0423/mcp-ip-geo.git
cd mcp-ip-geo
```

#### Step 2: Build the Docker image

```bash
docker build -t mcp-ip-geo-server .
```

#### Step 3: Run the container

```bash
docker run -d --name mcp-ip-geo-server -p 8000:8000 mcp-ip-geo-server
```

Once running successfully, the service will listen on `0.0.0.0:8000` within the container (listening on all network interfaces), and can be accessed via `http://<server-address>:8000/mcp`, where `<server-address>` can be:
- Local development environment: Use `localhost` or `127.0.0.1`
- LAN environment: Use the server's internal IP address (e.g., `192.168.x.x`)
- Public network environment: Use the server's public IP address or domain name

> Note: The service is configured to listen on the `0.0.0.0` address inside the container, which is standard practice for containerized applications, ensuring the service can be accessed from outside the container.

#### Step 4: Configure MCP

Add the following to your `MCP` configuration:

```json
{
  "mcpServers": {
    "mcp-ip-geo": {
      "url": "http://<server-address>:8000/mcp"
    }
  }
}
```

Replace `<server-address>` with the actual server IP address or domain name of your deployment environment.

# ⚠️ License Notice

> Note: This project uses the free version of ip-api.com, which is **for non-commercial use only**. If you intend to use this project for commercial purposes, please comply with their terms of service or purchase the paid version: https://ip-api.com/
