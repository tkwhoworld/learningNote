### 动态内存分配
> Programmers use dynamic memory allocators (such asmalloc) to acquire virtual memory (VM) at run time

>Dynamic memory allocators manage an area of process VM known as the heap. <br>
>动态内存分配器维护着一个进行的虚拟内存区域，称为堆（heap)

>Allocator maintains heap as collection of variable sized blocks, which are either allocated or free<br>
>分配器将堆（heap)视为一组不同大小的酷块（block)的集合来维护，每个块就是一个连续的虚拟内存片（chunk),要么是已分配，要么是空闲的。

### Types of allocators
>Explicit allocator: application allocates and frees space 
▪ E.g., malloc and free in C


>Implicit allocator: application allocates, but does not free space
▪ E.g., new and garbage collection in Java

>▪ Based on words of 8-bytes each
>Pointers returned by malloc are double-word aligned
▪ Double word = 2 words

### 分配器的要求和目标
> Explicit Allocators

>Can’t control number or size of allocated blocks<br/>
>不修改已分配的块。分配器只能操作或改变空闲块

>Must respond immediately to malloc requests
▪ i.e., can’t reorder or buffer requests<br>
>立即响应内存分配请求。

>▪ Must allocate blocks from free memory<br>
>不修改已分配的块。分配器只能操作或改变空闲块

>▪ Must align blocks so they satisfy all alignment requirements<br>
>对齐块（对齐要求）。分配器必须要对齐块，使得他们可以保存任何类型的数据对象。<br>

<p>
延伸学习（内存对齐）<br>
参考资料:https://www.cnblogs.com/-citywall123/p/12722800.html <br>

>内存对齐：CPU在读取内存时是一块一块读取的，块的大小可以是2,4,6,8，16（总之是2的倍数）
64位系统默认是8字节，块的大小也可以自己用宏定义。#pragma pack(nSize);

>内存对齐的规则<br>
>1.数据成员对齐的规则：结构体或联合体（union))的数据成员，第一个数据成员放在offset为0的地方，以后每个数据成员存储的起始位置要从该成员大小或者成员
的子成员大小（只要该成员有子成员，比如说是数组，结构体等）的整数倍开始(比如int在３２位机为４字节,则要从４的整数倍地址开始存储。<br>
>2.结构体作为成员：:如果一个结构里有某些结构体成员,则结构体成员要从其内部最大元素大小的整数倍地址开始存储.(struct a里存有struct b,b里有char,int ,double等元素,那b应该从8的整数倍开始存储.)<br>
>3.收尾工作。结构体的总大小,也就是sizeof的结果。必须是其内部最大成员的整数倍.不足的要补齐。
</p>
>▪ Can manipulate and modify only free memory<br>
>不修改已分配的块。分配器只能操作或改变空闲块

>▪ Can’t move the allocated blocks once they are malloc’d<br>
>不修改已分配的块。分配器只能操作或改变空闲块

>只使用堆
<p>
<h3>延伸学习，堆（heap):堆的定义及用法</h3>
参考：作者：Vamei 出处：http://www.cnblogs.com/vamei 欢迎转载，也请保留这段声明。<br>

>堆(heap)又被为优先队列(priority queue)。尽管名为优先队列，但堆并不是队列。回忆一下，在队列中，我们可以进行的限定操作是dequeue和enqueue。dequeue是按照进入队列的先后顺序来取出元素。而在堆中，我们不是按照元素进入队列的先后顺序取出元素的，而是按照元素的优先级取出元素。

>Linux内核中的调度器(scheduler)会按照各个进程的优先级来安排CPU执行哪一个进程。计算机中通常有多个进程，每个进程有不同的优先级(该优先级的计算会综合多个因素，比如进程所需要耗费的时间，进程已经等待的时间，用户的优先级，用户设定的进程优先程度等等)。内核会找到优先级最高的进程，并执行。如果有优先级更高的进程被提交，那么调度器会转而安排该进程运行。优先级比较低的进程则会等待。“堆”是实现调度器的理想数据结构。

(Linux中可以使用nice命令来影响进程的优先级)
</p>

### Performance Goal: Throughput
#### Goals: maximize throughput and peak memory utilization
▪ These goals are often conflicting

>Throughput:Number of completed requests per unit time
Example:
▪ 5,000 malloc calls and 5,000 free calls in 10 seconds
▪ Throughput is 1,000 operations/second

### Fragmentation 碎片
>Poor memory utilization caused by fragmentation
>>internal Fragmentation<br>
external Fragmentation

>造成堆利用率很低的主要原因是一种称为碎片的现象。
>>内部碎片<br>外部碎片

#### Internal Fragmentation (内部碎片)
> For a given block, internal fragmentation occurs if payload is
smaller than block size

>内部碎片是一个已分配的块比有效载荷大时发生。<br>
比如：需要7个bytes的空间，但是由于对齐原则，分配了8个字节，多出的就是内部碎片。

#### External Fragmentation (外部碎片)
>Occurs when there is enough aggregate heap memory,but no single free block is large enough<br>
>外部碎片是当空闲内存合计起来满足一个分配请求，但是没有一个单独的空闲块足够大可以来处理这个请求时发生。

### 实现问题
>How do we know how much memory to free given just a pointer?<br>
>How do we keep track of the free blocks<br>
>What do we do with the extra space when allocating a structure that is smaller then the free block it is placed in?<br>
>How do we pick a block to use for allocation --many might fit?
