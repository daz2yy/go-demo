# go-demo
Just Do It！


## 优雅关停
- 不优雅的方式
  - 在 Linux 终端输入 Ctrl + C （SIGINT）
  - 发送 SIGTERM 信号，例如：kill -9 或者 systemctl stop 等
- 问题
  - 有些请求正在处理，服务端直接退出，造成客户端报错：连接中断，请求失败
  - 程序需要做一些清理的工作；例如：等待进程内任务队列的任务执行完成，设置拒绝接受新消息等

- 优雅关停
  - 通过拦截 SIGINT 和 SIGTERM 信号来实现优雅关停

- 测试
  - 运行 shutdown.go
  - 访问 localhost:18080
  - 在 5s 内发送 Ctrl+C
  - 观察结果