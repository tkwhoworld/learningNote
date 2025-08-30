## 概述
bufio包实现了有缓冲的I/O。它包装一个io.Reader或io.Writer接口对象，创建另一个也实现了该接口，且同时还提供了缓冲和一些文本I/O的帮助函数的对象。

### type Reader 
Reader实现了给一个io.Reader接口对象附加缓冲。可以理解为一个Reader对象，其
````
type Reader struct {
    // 内含隐藏或非导出字段
}
````

### func (*Reader) ReadLine