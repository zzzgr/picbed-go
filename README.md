<!--suppress HtmlDeprecatedAttribute -->
<p align="center">
  <img width="150" src="http://p0.meituan.net/csc/992380ad5b41ca912dfdeaec6cad42202228.png" alt="Img">
  <p>图床</p>
</p>


## 🍦功能介绍
- 支持`windows`/`linux`/`mac`以及`amd`/`arm` cpu架构

## 🍼使用方法
### 1. 配置文件
```yaml
server:
  # 网站标题
  name: picbed-go
  # 端口
  port: 8080
  # 公告
  notice: 欢迎光临picbed-go
  # 生产留空即可
  mode: ""
config:
  # bilibili cookie
  bliCookie: '填写bilibili cookie, 格式: SESSDATA=5034xxxxxx2A21;bili_jct=6axxxxxx89'

```

### 2. Server端 (Docker和手动)
#### 方式一： docker
略，自行打包

#### 方式二： 手动编译
编译：
```shell
## Linux arm
go env -w CGO_ENABLED=0  GOOS=linux  GOARCH=amd64
go build

## Linux amd
go env -w CGO_ENABLED=0  GOOS=linux  GOARCH=amd64
go build 

## windows
go env -w CGO_ENABLED=0 GOOS=windows  GOARCH=amd64
go build
```

启动：
```shell
./picbed-go
```

### 3. 访问
访问: `http://ip:port/`


## 特别声明
内容来自互联网，请您必须在下载后的24小时内从计算机或手机中完全删除以上内容.
![image.png](http://p0.meituan.net/csc/d38677d1bf19649fd6d2c7ebbc52404c130438.png)
