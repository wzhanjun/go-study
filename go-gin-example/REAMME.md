## 

### 使用Docker部署应用

```
## Dockerfile

FROM scratch

WORKDIR $GOPATH/src/github.com/EDDYCJY/go-gin-example
COPY . $GOPATH/src/github.com/EDDYCJY/go-gin-example

EXPOSE 8000
CMD ["./go-gin-example"]

## 编译
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-gin-example .

### windows下编译linux环境
SET CGO_ENABLED=0
SET GOOS=linux
go build -a -installsuffix cgo -o go-gin-example .

## 构建镜像
docker build -t gin-blog-docker-scratch .

## 运行go镜像
docker run --link mysql:mysql -p 8000:8000 gin-blog-docker-scratch

## 挂载数据卷
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=rootroot -v /data/docker-mysql:/var/lib/mysql -d mysql:5.7

```