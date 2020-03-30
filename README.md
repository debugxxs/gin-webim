## gin-webim 
基于gin+websocket做的即时通讯的小应用，作为`go` `websocket`的入门应用  

## 实现的功能 

- [x] 群聊 
- [x] 私聊 
- [x] 记录在线用户列表 

## 安装 与 启动

> git clone https://github.com/MarichMarck/gin-webim.git   
> cd gin-webim  
> go get -u -v github.com/gin-gonic/gin   
> go get -u -v github.com/gorilla/websocket  
> go run main   // http://localhost:8008 

## 注意点  

上线使用注意调整路径，`router/init.go` 文件中 `16 - 18` 行

## 预览

**全貌**  
![全貌](https://github.com/MarichMarck/gin-webim/blob/master/priview/webim-group.png)  

**私聊**   
![私聊](https://github.com/MarichMarck/gin-webim/blob/master/priview/webim-private.png)  

 
