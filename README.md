# [舔狗](https://jkl201906112027.info)

### 我们是狗，舔狗。

网上找的舔狗日记数据，用Golang实现CURD相关接口，用于前端(todo)动态化展示.

## 没卵用的特性
- 人肉对每段数据的内容、标点符号进行纠错与优化， 按本人理解进行了划重点标记。
- 接入[和风天气API](https://dev.heweather.com/)，获取城市、气象、气温。免费版每日只有1000次。
- 使用个性字体：站名「造字工房尚雅」；正文「汉仪跳跳体」。为了性能，使用了 [font-spider](https://github.com/aui/font-spider) 对字体进行压缩。字体源文件：/assets/font/.font-spider/；文字承载页：/assets/font/index.html。
 
 
 ## 本地运行
 注意：项目使用Go Module管理依赖
 
 ##### 1. 执行init/init.sql，创建DB和表
 
 ##### 2. 配置初始化环境配置，项目根目录下.env.example文件复制一份重命名为.env,配置改成自己本地的配置
 ```cassandraql
 MYSQL_DSN="username:password@tcp(dbIp:dbPort)/dbName?charset=utf8"
 REDIS_ADDR="127.0.0.1:6379"
 REDIS_PW=""
 REDIS_DB=""
 SESSION_SECRE=""
 GIN_MODE="debug"
 SERVER_LISTON_PORT=8088
 ```
 ##### 3. 项目根目录下执行
 ```cassandraql
 go get
 go mod vendor
 go run main.go
 ```
 ##### 4. 浏览器访问接口验证
 ```
 http://localhost/api/v1/getRandomLetter
 ```
 
 
## TODO LIST
- 动态化，后端用Go提供接口：获取日记、添加日记
- 前后端分离，前端用vue.js重构???
- 重新设计投稿流程，直接前端提交，审核日记
- 配置webhook, 代码提交后自动编译部署

## 后端接口

### 1.随机获取一篇日记
```
接口地址：http://localhost/api/v1/getRandomLetter
method： get
resp json：
{
  "errCode": 0,
  "errMsg": "",
  "data": {
    "id": 78,
    "sender": "",
    "receiver": "",
    "content": "你没再来找我了，我发了一条仅你可见也没任何回应，**你的朋友圈每一条我都点了赞，也没引起你的注意**，我不知道你有没有再想起我，是不是我等的还不够久。",
    "weather_info": "",
    "create_time": "2020-05-05T04:25:40+08:00"
  }
}
```
### 2.新增日记
TODO
