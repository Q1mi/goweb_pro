# 指南


## 请按如下顺序启动项目

1. 请根据你的实际情况修改 conf/config.yaml 文件中 MySQL 和 Redis 部分的配置！！！
2. 连接上你的MySQL数据库，按顺序依次执行项目中SQL文件夹中的sql文件，完成建库、建表和导入初始数据
    1. init.sql
    2. bluebell_user.sql
    3. bluebell_community.sql
    4. bluebell_post.sql
3. 执行 `go build -o ./bin/bluebell`，编译可执行文件至项目的bin目录
4. 执行 `./bin/bluebell conf/config.yaml`，启动程序
5. 打开你的浏览器输入 [http://127.0.0.1:8084](http://127.0.0.1:8084)，默认端口是 8084，你可以在配置文件中修改

## 注意事项
1. 确保你的MySQL配置是正确的
2. 确保你的Redis配置是正确的
3. 可点击右上角自行注册测试账号
    