# 博客后端
## 配置
1. 启动 `mongoDB`（需要提前安装），创建数据库 `Articles`，并在该数据库中创建collection：`ArticleData`和`Users`，然后创建一个具有读写权限的用户，用户名`service_test`，密码`123456`

2. 在cmd运行以下命令
   `go get -v github.com/610Masters/Backend`

   

## 运行
在文件夹内运行以下命令
`go run main.go`