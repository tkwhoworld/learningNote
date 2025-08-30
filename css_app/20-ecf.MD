## 异常控制流 ecf Exceptional Control Flow:Exceptions and Processes

### Control Flow
>From startup to shutdown,each CPU core simply reads and executes(interprets) a sequence of instructions,one at time *，
This sequence is the CPU’s control flow (or flow of control)
>计算机从开机到关机，每个CPU Core(每个核都有一个程序计数器)在一个时刻都在读取和运行一系列的指令
这个序列就是CPU的控制流。

>控制流就是指令的执行顺序？

### Altering the Control Flow
>two mechanisms for changing control flow: 改变控制流的2种机制
>>▪ Jumps and branches
>>▪ Call and return
>React to changes in program state

### Exceptional Control Flow
>Exists at all levels of a computer system
>存在所有级别的机器中

### Low level mechanisms 低级机制(偏向底层？)
>1. Exceptions 
>> Change in control flow in response to a system event 
>> Implemented using combination of hardware and OS software

### Higher level mechanisms 高级机制（软件层面?)
>▪ 2. Process context switch
>▪ 3. Signals
>▪ 4. Nonlocal jumps: setjmp() and longjmp()

### Exceptions（异常）
> An exception is a transfer of control to the OS kernel in response 
to some event (i.e., change in processor state)
>异常就是控制流中的突变，用来响应处理器状态中的某些变化

>▪ Kernel is the memory-resident part of the OS

>Examples of events: Divide by 0, arithmetic overflow, page fault, I/O 
request completes, typing Ctrl-C

### Exception Tables 异常表

### ECF:Asynchronous(Interrupts) Synchronous(Traps、Faults、Aborts)

#### Asynchronous Exceptions (Interrupts) 异步异常
