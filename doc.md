#1.创建go项目

mkdir holy-go
cd holy-go
go mod init github.com/xulei131401/holy-go
goctl api new admin
go mod tidy

cd script && bash start.sh
访问服务：curl -i -X GET http://localhost:8888/from/you

#2.配置config
增加jwt,mysql,redis等配置

#3.修改api文件重新生成api
goctl api go -api admin.api -dir .


#4.生成model
$ cd service/user/model
$ goctl model mysql ddl -src user.sql -dir . -c
$ goctl model mysql ddl -src service/student/model/student.sql -dir service/student/model/ -c
或者
$ goctl model mysql datasource -url="$datasource" -table="user" -c -dir .

#5.删除.idea
git rm --cached -r .idea
然后再提交

#6.初始化项目自己的goctl模板，方便做定制化
#goctl template init --home $HOME/template
goctl template init --home /Users/xulei/jungle/githubproject/goctl-template/template

[//]: # (#7.新增routes包)
[//]: # (拷贝handler目录下的route.go到routes目录下)

#9.api文件需要自己定制化一部分

#10.jwt token搞定
#11.数据库


#############go-zero最佳实践
1.api定义的结构体只允许使用json tag，不要使用自带的tag，方便接入第三方库验证参数
2.经过实验发现，现在的api文件定义支持validate tag，
3.curl模拟请求

curl --location --request POST 'localhost:8888/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
"username": "foo",
"number": 10
}'


#########问题
1.@server能放到service里吗？prefix


#########定制化工作
1.搞清楚goctl api的原理，可以自己定制化
2.对接gorm，gormt生成的文件也需要定制化
3.session组件
4.gorm的封装,增删改查，分页，集合，pdo属性判断等




