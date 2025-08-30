## 虚拟内存-概念

### Processes 进程

>Definition: A process is an instance of a running program.<br>
>定义：进程是允许中的程序实例

>Processes provides each program with two key abstractions<br>
>进程提供给程序2个关键的抽象

>Logical control flow 逻辑控制流
>> Each program seems to have exclusive use the CPU.<br>
>> 每个程序似乎独占使用CPU<br>
>> Provided by Kernel mechanism called context switching<br>
>> 内核提供上下文切换

>Private address space 私有地址空间
>>Each Program seems to have exclusive use of mani memory<br>
>>Provided by kernel machanism called virtual memory

### 进程状态 processes state
>Runing 运行<br>
>▪ Process is executing (or waiting to, as we’ll see later this week)<br>
>进程要么在CPU上执行，要么等待被执行且最终会被内存调度<br>

>Stopped 停止 <br>
>▪ Process execution is suspended until further notice (covered later)<br>
>进程的执行被挂起（suspended),且不会被调度直到进一步的通知

>Terminated 终止<br>
>▪ Process is stopped permanently<br>
>进程永远地停止了

### Terminating Processes 终止进程
>Process becomes terminated for one of three reasons:<br>
>>▪ Returning from the main routine<br>
>>从主程序返回<br>
>>▪ Calling the exit function<br>
>>调用exit函数

### Creating Processes 创建进程
>Parent process creates a new running child process by calling fork<br>
>父进程通过调用fork方法创建一个运行中的子进程
>int fork(void)<br>
>>▪ Returns 0 to the child process, child’s PID to parent process<br>
>>▪ Child is almost identical(完全相同的) to parent...<br>
>>>▪ Child get an identical (but separate) copy of the parent’s virtual address space.
>>>相同但是独立的地址空间。地址空间是相同的，每个进程有相同的用户栈、相同的本地变量值，相同的堆，相同的全局变量。
以及相同的代码。然而子进程和父进程是独立的进程，他们都有自己的私有的空间，任何改变都是独立的。
>>>▪ Child gets identical copies of the parent’s open file descriptors
>>>共享文件
>>>▪ Child has a different PID than the parent
>>>不同的进程ID


### Address Space 地址空间
>寻址方式<br>
>A System Using Physical Addressing<br>
>>Used in “simple” systems like embedded microcontrollers in
devices like cars, elevators, and digital picture frames<br>
>>这种系统主要使用在汽车、电梯和数码相框等设备。<br>
> A System Using Virtual Addressing<br>
> 使用虚拟内存
>>Used in all modern servers, laptops, and smart phones
>>One of the great ideas in computer science

>Linear address space: Ordered set of contiguous non-negative integer
addresses:
{0, 1, 2, 3 … }<br>
>如果地址空间中的整数是连续的，我们称为线性地址空间（Linear address space)

>Virtual address space: Set of N = 2^n virtual addresses
{0, 1, 2, 3, …, N-1}<br>
>在一个带虚拟内存的系统中，CPU从一个有N=2^n个地址的地址中间中生成虚拟地址。N=2^n.n为32位或64位

>Physical address space: Set of M = 2m physical addresses
{0, 1, 2, 3, …, M-1}<br>
>物理地址空间

### 虚拟内存
>⬛Uses main memory efficiently
>>▪ Use DRAM as a cache for parts of a virtual address space
>⬛Simplifies memory management
>>▪ Each process gets the same uniform linear address space
>⬛Isolates address spaces
>>▪ One process can’t interfere with another’s memory
>>▪ User program cannot access privileged kernel information and code

#### VM as a Tool for Caching 虚拟内存作为缓存的工具
>Conceptually,virtual memory is an array of N contigous bytes stored on disk<br>
>概念上而言，虚拟内存是一个存放在磁盘上的N个连续字节组成的数组

>The contents of the array on disk are cached in physical memory (DRAM cache)
>▪ These cache blocks are called pages (size is P = 2p bytes)

>磁盘上数组的内容被缓存到主存中。和存储器层次


#### VM as a Tool for Memory Management
#### 虚拟内存作为内存管理的工具

>⬛Key idea: each process has its own virtual address space
>虚拟内存（VM）为每个进程提供独立的地址空间<br>
>▪ It can view memory as a simple linear array<br>
>可以将内存视为简单的线性数组<br>
>▪ Mapping function scatters addresses through physical memory<br>
▪ Well-chosen mappings can improve locality<br>

>⬛Simplifying memory allocation
>简化内存分配
>▪ Each virtual page can be mapped to any physical page
>每个虚拟页都能映射到物理页
>▪ A virtual page can be stored in different physical pages at different times

>⬛Sharing code and data among processes
>▪ Map virtual pages to the same physical page
>简化共享
>独立地址空间为操作系统提供了一个管理用户进程和操作系统自身之间的共享的一致机制。

>Simplifying Linking and Loading<br>
>简化链接和加载<br>
>▪ Each program has similar virtual address space<br>
>每个程序都有相似的虚拟地址空间<br>
>Code,data and heap always start at the same addresses.

>Allocate virtual page for .text and .data setions & creates PETs marked as invalid
>The .text and .data sections are copied,page by page,on demand by the virtual memory system.

>虚拟内存还使得容易向内存中加载可执行文件和共享文件对象文件。

#### VM as a Tool for Memory Protection 虚拟内存作为内存保护的工具


