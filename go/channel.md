## channel 是一个引用的数据类型 channel即指通道类型 也指代可以传递某种类型的通道
通道即某个通道类型的值，是该类型的一个实例
## 类型表示方法
关键字type用来为类型声明字段，用来声明一个类型。type myString string
chan T 
type IntChan chan int
var intChan chan int
var intChan chan<- int
var intChan <-chan int
FIFO的消息队列
var strChan := make(chan string,3)
elem,ok := <-strChan
strChan <- "a"
strChan <- "b"
strChan <- "c"
## 操作的特性
## 初始化通道
## 接收元素的值