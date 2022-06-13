#1.创建go项目

mkdir holy-go
cd holy-go
go mod init holy-go
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










