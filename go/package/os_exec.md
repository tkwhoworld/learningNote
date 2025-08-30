exec 包
## 使用方式 import os/exec
## 概述 exec包执行了外部命令。它包装了os.StartProcess函数以便更容易修正输入和输出，使用管道连接I/O
## 理解 go支持的IPC有管道、信号、socket，管道主要是依赖于os/exec实现，其实就是封装了一些linux 命令的执行过程。IPC描述的是进程之间的通信。
管道有一个特点，它是半双工（其实就是单向），就是linux的命令的|，前一个命令的结果是后一个命令的输入。
## Command方法
### func Command(name string, arg ...string) *Cmd
理解 输入命令返回一个*Cmd对象的指针

## Type Cmd
定义:
type Cmd struct {
    // Path是将要执行的命令的路径。
    //
    // 该字段不能为空，如为相对路径会相对于Dir字段。
    Path string
    // Args保管命令的参数，包括命令名作为第一个参数；如果为空切片或者nil，相当于无参数命令。
    //
    // 典型用法下，Path和Args都应被Command函数设定。
    Args []string
    // Env指定进程的环境，如为nil，则是在当前进程的环境下执行。
    Env []string
    // Dir指定命令的工作目录。如为空字符串，会在调用者的进程当前目录下执行。
    Dir string
    // Stdin指定进程的标准输入，如为nil，进程会从空设备读取（os.DevNull）
    Stdin io.Reader
    // Stdout和Stderr指定进程的标准输出和标准错误输出。
    //
    // 如果任一个为nil，Run方法会将对应的文件描述符关联到空设备（os.DevNull）
    //
    // 如果两个字段相同，同一时间最多有一个线程可以写入。
    Stdout io.Writer
    Stderr io.Writer
    // ExtraFiles指定额外被新进程继承的已打开文件流，不包括标准输入、标准输出、标准错误输出。
    // 如果本字段非nil，entry i会变成文件描述符3+i。
    //
    // BUG: 在OS X 10.6系统中，子进程可能会继承不期望的文件描述符。
    // http://golang.org/issue/2603
    ExtraFiles []*os.File
    // SysProcAttr保管可选的、各操作系统特定的sys执行属性。
    // Run方法会将它作为os.ProcAttr的Sys字段传递给os.StartProcess函数。
    SysProcAttr *syscall.SysProcAttr
    // Process是底层的，只执行一次的进程。
    Process *os.Process
    // ProcessState包含一个已经存在的进程的信息，只有在调用Wait或Run后才可用。
    ProcessState *os.ProcessState
    // 内含隐藏或非导出字段
}
