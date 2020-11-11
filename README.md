[toc]
# 更新日志
```
2020年11月11日 11:04:18 修改VUE引用为生产环境
2020年11月8日 18:32:10 在session 跨域问题上碰壁，重写鉴权部分
2020年10月29日 19:54:48 添加go gin框架搭建的快速路由
2020年10月27日 11:14:38 使用VUE重构前端
更早以前
前端demo
...古老版本
```

# 介绍与展示

This is a personal homepage that contains blogs.
这是一个包含博客的个人主页。
展示:[我的主页](http://yearagain.com)

设计思路:前端后端分离设计,异步完成数据交互

# 快速开始

如果你要使用后端[这里](https://github.com/yearagain/ihome/tree/main/background/bygin)

你可以使用route目录中的go程序快速建立静态路由程序（这和后端并不是一个系统）

你可能需要调整

bgsite.js login.js blog.js 文件开头的接口路径

ayth.js addblog.html blog_show.html 中的接口路径

login.js中的加密盐


# 前端引用库

前端:bootstrap,vue
md5处理：blueimp-md5
markdown处理:[marked](https://github.com/markedjs/marked)
markdown编辑器：[EasyMDE](https://github.com/Ionaru/easy-markdown-editor)

# 后端
目录：background/bygin 是一个基于gin框架的后端处理
详见：[后端bygin](https://github.com/yearagain/ihome/tree/main/background/bygin)

# 路由

route目录提供了一个使用go构建的快速静态路由

# 字体
仓耳周珂正大榜书：http://tsanger.cn/product/194

# 未来计划

不断的美化和更新

更多的内容交给后端

~~使用VUE重构前端~~