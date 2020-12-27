# OAuth 2.0
有三个角色：存储资源/数据的服务器A、想要获取数据的第三方应用B、资源/数据所有者C

有四种授权模式：授权码(code)、隐藏式、密码式(password)、凭证式
- 授权码(code)：适用于后端交互的web应用，B将通过A的授权接口得到时限较短的授权码，此时需要C授权，A授权后将携带授权码跳转到B指定的回调地址
，B将使用授权码和自己的ID和Secret向A请求时限较长的访问令牌。
- 密码式(password)：适用于系统高度信任的应用，只要应用拥有了用户名和密码，就可以直接访问A的获取令牌接口，通过密码校验后获得令牌。

## 项目中应用
OAuth相关接口设计在gateway服务中。

在登录流程中，用户访问获取授权码接口，登录密码的校验在用户授权环节进行，获取授权码后跳转到
获取令牌接口，利用授权码和应用程序信息交换令牌。

## 获取授权码
- 接口： /oauth/authorize
- 方式：GET
- 参数： 
```
{
  response_type：授权类型，此处的值固定为 code（必选）
  client_id：在您注册应用程序时得到的 Client ID（必选）
  scope：请求权限范围，多个权限使用空格分隔（见下文）（必选）
  redirect_uri：授权完成后重定向至的 URL，即您的应用程序用于接受授权结果的地址（必选）
  state：完成授权后传回的唯一字符串，认证服务器会原封不动地返回这个值，防止伪造攻击（可选）  
}
```
- 响应： 
```
{
  code：表示授权码（必选）
  state：如果步骤 1 客户端请求中包含该参数，步骤 2 重定向地址中也必须一模一样包含该参数（可选）  
}
```
## 跳转redirect_uri
第三方应用的ID和Secret存储在C中，由于A和C为同一服务，因此跳转到的redirect_uri完成的工作是将client_id
client_secret和获取到的code作为参数交换令牌

- 接口：/oauth/redirect
- 方式：get
- 参数：
```
{
    "code": "",
    "state": "" (if have)
}
```
- 响应：获取client_id和client_secret重定向到获取令牌接口

## 获取令牌

- 接口：/oauth/token
- 方式：POST
- 参数：
```
{
  grant_type：表示使用的授权模式，此处的值固定为 authorization_code（必选）
  code：表示上一步获得的授权码（必选）
  client_id：在您注册应用程序时得到的 Client ID（必选）
  client_secret ：在您注册应用程序时得到的 Client Secret（必选）  
}
```
- 响应： 
```
{
  "access_token": "xoxp-23984754863-2348975623103",
  "expires_in": 86400,
  "scope": "store pet",
  "refresh_token": "tGzv3JOkF0XG5Qx2TlKWIA"
}
```

## go-oauth2
模块划分为server、generate、store、manage，generate定义了授权码（authorize code）和令牌（token）的结构和创建方式，
store负责提供存储client和token，manage用于自定义配置，如使用自定义的token生成jwttoken，
使用mysql存储注册到系统的client信息，使用redis存储token；核心模块server，定义了一系列Handler
来处理获取授权码、校验密码、获取令牌等请求

- PasswordAuthorizationHandler: 校验密码返回USERID
- UserAuthorizationHandler: 获取USERID
- AuthorizeScopeHandler: 权限范围校验
- AccessTokenExpHandler: 设置令牌过期时间
- ClientAuthorizedHandler: 校验client_id


其工作过程为：
- 获取授权码

根据路由绑定将请求转发到server提供的handleAuthorizeRequest方法，
校验request相关参数后，进入用户授权环节，调用UserAuthorizationHandler，获取USERID。然后通过
AuthorizeScopeHandler、AccessTokenExpHandler处理权限范围和设置令牌过期时间，
然后调用GetAuthorizeToken方法，在其中通过ClientAuthorizedHandler校验client_id、ClientScopeHandler
校验client是否拥有对应的权限scope，然后调用GenerateAuthToken方法，创建授权码。
- 获取令牌

根据路由绑定将请求转发到server提供的HandleTokenRequest方法，在ValidationTokenRequest方法中，
根据不同的grant_type，处理传入参数。在密码式中，调用PasswordAuthorizationHandler校验密码返回USERID，
校验通过之后，调用GetAccessToken方法，在GenerateAccessToken中，通过getClient获取传入client_id对应的
client信息，如果client实现了ClientPasswordVerifier接口，则调用该接口的VerifyPassword方法校验client_id
和client_secret，否则将简单比对传入的client_secret和查询得到的client信息。校验通过之后，通过AccessGenerate
接口的Token方法生成令牌assess_token

### 使用JWT创建AssessToken
在创建oAuthServer时通过manager指定生成token的方式为JWTAccessGenerate，
其Token方法接收传入参数signKey、signKeyId、signingMethod
```
生成的token结构: {
  Header: {
    "typ": "jwt",
    "alg": signingMethod.Alg(),
    "kid": signKeyId
  },
  Claims: {
    "Audience": client_id,
    "Subject": userID,
    "ExpiresAt": expireTime
  },
  Method: signingMethod
}
```

# Git
## 修改上次提交信息 commit message
```bash
# 进入修改 (vim)
git commit --amend
# 修改后退出
:wq
#推送到仓库
git push -f
```
## 拷贝git仓库（包含提交日志）
```bash
# 从原始仓库拷贝一份不包含工作区的裸版本库
git clone --bare git@原仓库地址
# 进入项目目录并推送到新的仓库地址
cd xx
git push --mirror git@目标仓库地址
```
## 代码关联到多个git仓库
```bash
# 查看当前代码关联的仓库
git remote -v
# 关联新的仓库
git remote add 仓库名 地址
# 查看是否生效
git remote -v
```

# Proto
```bash
# 位于目录: 项目目录的上级目录
# 编译dto下的文件
protoc --proto_path=. --go_out=. course/user-srv/proto/dto/user.proto
# 编译微服务文件
protoc --proto_path=. --go_out=. --micro_out=. course/user-srv/proto/user/user.proto
```

# Docker
## Consul

```bash
# 启动第一个server节点
docker run --name consul1 -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600 consul agent -server -bootstrap-except 2 -ui -bind=0.0.0.0 -client=0.0.0.0
# 获取第一个节点的ip
docker inspect --format '{{ .NetworkSettings.IPAddress }}' consul1
# 开启其他server节点并加入第一个节点
docker run --name consul2 -d -p 8501:8500 consul agent -server -ui -bind=0.0.0.0 -client=0.0.0.0 -join x.x.x.x
docker run --name consul3 -d -p 8502:8500 consul agent -server -ui -bind=0.0.0.0 -client=0.0.0.0 -join x.x.x.x
# 查看集群状态
docker exec -it consul1 consul members
```

## Mysql
```
docker run -itd -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql
```
### 高版本关闭only_full_group_by
```
# 进入/etc/mysql/mysql.conf.d 在mysqld.cnf中添加
sql_mode=STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION
```

## MariaDB

```
docker run --name mariadb -v /var/lib/mysql:/var/lib/mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root@123456 -e MYSQL_ROOT_HOST=% -d mariadb --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
```

## Redis
```
docker run -itd -p 6379:6379 redis
```

## 自动化发布API
- 基础命令
```bash
docker run -it --rm \
  -e ACCESS_TOKEN={访问令牌} \
  -e APIDOC_TEAM={企业域名} \
  -e APIDOC_PROJECT={项目地址名称} \
  -e APIDOC_ID={API 文档资源 ID} \
  -e APIDOC_RELEASE_TYPE=file \
  -v {API 数据文件路径}/data.txt:/opt/data.txt \
  ecoding/apidoc-publisher
```
- 编写dockerfile构建自己的镜像用于发布API
> 主要是在Dockerfile中，定义需要的环境变量，以及API文件路径。
```bash
# 运行 ecoding/apidoc-publisher，从容器中获取其发布API的脚本
docker cp 容器ID:/my_shell/api_doc_release.sh /xx/
# 修改脚本内容，增加下载文件命令（line 52-53）
```
![](https://gitee.com/jingshanccc/image/raw/master/image/image-20201201120717958.png)
```bash
# 编写Dockerfile
vim Dockerfile
FROM ecoding/apidoc-publisher:20200115.1
ADD api_doc_release.sh /my_shell/api_doc_release.sh
ENV FILE_URL=https://gitee.com/jingshanccc/course/raw/main/doc/openapi.yaml
ENV ACCESS_TOKEN=your token
ENV APIDOC_TEAM=your team
ENV APIDOC_PROJECT=your project
ENV APIDOC_ID=1
ENV APIDOC_RELEASE_TYPE=file
# 构建镜像
docker build -t xx/xxx .
```