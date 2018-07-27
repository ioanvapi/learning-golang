# memberlist

## Run

这里通过一个简单的 http 服务查询和插入数据, 找两台机器, 第一台执行:

```
memberlist
```

会生成 gossip 监听的服务 ip 和端口
使用上面的 ip 和端口在第二台执行

```
memberlist --members=xxx.xxx.xxx.xxx:xxxx
```

那么一个 gossip 的网络就搭建完成了

```
# add
curl "http://localhost:4001/add?key=foo&val=bar"

# get
curl "http://另一台机器:4001/get?key=foo"

# delete
curl "http://localhost:4001/del?key=foo"
```
