信号相关的package
## 操作系统信号（signal）是IPC中唯一一种异步的通信方法，它的本质是用软件来模拟硬件的中断机制。
## Linux支持62个信号
信号的来源：键盘输入、硬件故障、系统函数调用和软件中的非法运算
Linux对每个标准信号都有默认的操作方式，默认的操作方式有如下几种:终止进程、忽略该信号、
终止进程并保存内存信息、停止进程、恢复进程。
go语言通过os/signal中的api实现，通过os/signal包实现了对输入信号的访问。
signal包提供了2个函数
func Notify
func Notify(c chan<- os.Signal, sig ...os.Signal)
