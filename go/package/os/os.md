# os包
## 概述
os包提供了操作系统函数的不依赖平台的接口。设计为Unix风格的，虽然错误处理是go风格的；失败的调用会返回错误值而非错误码。通常错误值里包含更多信息。例如，如果某个使用一个文件名的调用（如Open、Stat）失败了，打印错误时会包含该文件名，错误类型将为*PathError，其内部可以解包获得更多信息。

## type File
file代表打开的文件对象
````
type File struct {}

````

### File提供的方法如下
#### func Create
Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）。如果成功，返回的文件对象可用于I/O；对应的文件描述符具有O_RDWR模式。如果出错，错误底层类型是*PathError。
````
func Create(name string,err error)

````
#### func (*File) Write
Write向文件中些人len(b)字节数据。它返回写入的字节数和可能遇到的任何错误，如果返回n!=len(b),本方法会返回一个非nil的错误
定义
````
func (f *File) Write(b []type)(n int ,err error)

````

#### func (*File) WriteString
类似Write但是接受一个string

