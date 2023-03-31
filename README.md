# wechatbot
项目基于[openwechat](https://github.com/eatmoreapple/openwechat)开发
### 目前实现了以下功能
 + 群聊@回复
 + 私聊回复
 + 自动通过回复
 
# 配置
### 获取chatGLM部署后端地址
配置项在 config.json ```glm_backend```

### 限制对话轮数
配置项在 config.json ```max_boxes```
每个用户单独配置上下文历史栈，超出限制轮数将清空历史。

### 自动通过
配置项在 config.json ```auto_pass```
设置为```true```将自动通过微信好友请求

### 配置最大并发用户数
配置项在 config.json ```user_count```

# 安装使用

### 获取项目
```git clone https://github.com/physi-cs/wechatbot```

### 进入项目目录
```cd wechatbot```

### 启动项目
1. ```nohup startup.sh > wxglm.log 2>&1 &```进行项目后台启动
2. 使用vim打开 wxglm.log ，找到最新的日志，即为登录链接(在vim命令模式下输入'G'可快速跳转到最后一行)
3. 点击链接使用浏览器打开，微信扫码进行登陆

### 关闭项目
1. ```ps -aux | grep main.go```找到项目对应的进程id
2. ``` kill -9 进程id ```
