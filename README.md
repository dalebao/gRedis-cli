# gRedis-cli

使用golang编写的redis命令行查询工具。

##出发点：
1. 在工作中，会生成很多规律的redis键，如：test_1,test_2，当需要人肉删除test_*键的时候，使用原生redis-cli，痛苦。
2. 在工作中，redis键太多，常常会让人忘记redis键的类型，需要先type再用对应类型的查询命令查询，太累。

##特点：
1. 使用一个命令，查询string,hash,list,set,zset类型的数据
2. 批量查询redis键的ttl
3. 批量查询redis键的类型
4. 使用通配符匹配redis键，选择或直接删除redis键
5. 使用table直观展示redis操作情况

![e](https://github.com/dalebao/gRedis-cli/raw/master/gRedis-cli.png)

##命令与使用：
```
git clone 
cd gRedis-cli
go run main.go
```
按照流程填写服务器连接信息

### get
   查询string,hash,list,set,zset类型的数据
    `get redisKey`
    
    
### keys
   使用通配符匹配redis键，返回redis键与对应类型
    `keys *`
    
### type
   批量查询redis键类型
    `type redisKey1 redisKey2`
    
### ttl
   批量查询redis ttl信息
    `ttl redisKey1 redisKey2`
    
### expire
   设置redis键过期时间
   `expire redisKey1 100` 单位秒
   
### del
   批量删除redis键
   `del redisKey1 redisKey2`

### rdel
   匹配redis键，直接或选择删除redis键
   `rdel redis*`
   
### 退出
   输入 `quit`
   
##接下来要做
1. 继续完善查询功能
2. 考虑是否要增加修改redis键内容
3. 增加配置保存功能，避免重复输入配置信息
4. 思考大量数据redis键的处理方式
5. 期待在issue中与我交流
   
##鸣谢
[命令行构建工具](https://github.com/AlecAivazis/survey)

[golang表格构建工具](https://github.com/modood/table)
    
    