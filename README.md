# 一个简单的 toml 配置读取包

这是一个非常简单的配置读取包，功能非常单一，使用非常简单；

安装
```
github.com/zhuCheer/cfg
```

使用

```
cfg.InitConfFile("./config.toml") //设置配置文件地址
value:=cfg.GetInt("database.connection_max") //读取toml配置节点



```

