# Traefik Yaml Generator

## 简介 Breif

`tygen` （Traefik Yaml Generator） 主要用于生成、修改 `docker-compose.yml` 文件，使之能够顺利地搭配 `Traefik` 进行使用。

- [x] 创建配置文件
- [x] 追加配置文件
- [ ] 从环境变量中获取参数
- [ ] 对 label 进行去重（包含二次确认）

## 获取 Get

你可以直接通过 `wget` 获取已发布的 Linux 下的二进制文件：

```shell
wget https://github.com/FrogDar/traefik-yaml-generator/releases/download/v0.0.1/tygen
```

当然，也能通过 `go install` 进行安装（只不过需要自己重命名一下x）

```shell
go install github.com/FrogDar/traefik-yaml-generator
```

> TODO: 让 `go install` 安装的文件名自动改为 `tygen`

## 命令 Commands

- `append`： 为现有的 `docker-compose.yml` 中的服务添加 `Traefik` 所需内容
- `completion`：生成自动补全的 `shell` 脚本
- `create`： 从零创建一个 `docker-compose.yml`，需要至少指定一个 docker compose 的服务名。
- `help`：查看命令帮助

> append: Append subcommand appends labels which traefik needs to the yaml file.
> completion: Generate the autocompletion script for the specified shell
> create: Create subcommand create a docker-compose file with Traefik.
> help: Help about any command

## 选项 Flags

### 中文 Chinese

| 缩写 | 全写          | 值类型 | 效果                                                    |
| ---- | ------------- | ------ | ------------------------------------------------------- |
| -a   | --address     | string | 指定 Traefik 主机地址（默认为"example.com"）            |
| -e   | --entrypoints | string | 指定 Traefik 访问端口（默认为"websecure"）              |
| -h   | --help        |        | 获取帮助                                                |
| -i   | --image       | string | 指定 docker image （默认为 "example/example:latest"）   |
| -n   | --network     | string | 指定 Traeifk 网络名 （默认为 "traefik-global-proxy"）   |
| -o   | --output      | string | 指定输出文件                                            |
| -p   | --port        | int    | 指定容器内部服务端口（默认为 80）                       |
| -r   | --rule        | string | 指定 Traefik 匹配规则（默认为 "Host（`example.com`）"） |
| -s   | --service     | string | 指定 Traefik 服务名 （默认为 "service"）                |
| -t   | --tls         |        | 开启该 Traefik 服务的 TLS 功能 （默认为 true）          |

注：

- `address` 是对 `rule` 的简略，相当于定义 `rule` 为  `` Host(`${address}`) `` 。若两个选项同时定义，将会以 `rule` 为主。
- 若没有指定 `service` ，将会默认以 docker compose 中服务的名字作为 `service` 的值。

### 英文 English

| Short | Long          | Type   | Description                                          |
| ----- | ------------- | ------ | ---------------------------------------------------- |
| -a    | --address     | string | Traefik host address (default "example.com")         |
| -e    | --entrypoints | string | entrypoints for the Traefik (default "websecure")    |
| -h    | --help        |        | help for tygen                                       |
| -i    | --image       | string | docker image (default "example/example:latest")      |
| -n    | --network     | string | network name (default "traefik-global-proxy")        |
| -o    | --output      | string | output file                                          |
| -p    | --port        | int    | internal port for the service (default 80)           |
| -r    | --rule        | string | traefik host rule (default "Host(`example.com`)")    |
| -s    | --service     | string | service name (default "service")                     |
| -t    | --tls         |        | enable tls for the service in Traefik (default true) |