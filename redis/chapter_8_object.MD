## 对象
Redis基于基本的数据结构(SDS,链表，字典，压缩表，整数集合，跳跃表)构成了一个对象系统，对象系统包括：
字符串对象，列表对象，哈希对象，集合对象，有序集合对象
对象系统实现了引用计数的内存回收机制，当程序不再使用某个对象的时候，这个对象的内存就会被自动释放
redis带有访问时间记录信息，该信息可以用于计算数据库键的空转时长，在服务器启用了maxmemory功能的情况下，
空转时长较大的那些键会优先被服务器删除
####8.1对象类型和对象编码
Redis使用对象来表示数据库中的键和值，每次当我们在Redis数据库中创建新的键值对时，我们至少会
创建两个对象，一个用来存储键，一个用来存储值
````
redis-> SET msg "hellow word"
````
redisObject的结构如下
server.h/redisObject
````
typedef struct redisObject {
    unsigned type:4;//类型
    unsigned encoding:4;//编码
    unsigned lru:LRU_BITS;/*LRU time (relative to global lru_clock) or
                           *LFU data  (least significant 8bits frequency
                           *and most significant 16bits access time*/
    int refcount;
    void *ptr;//指向底层实现数据结构的指针
}
````
#####8.1.1类型
键总是字符串对象
值字符串对象/列表对象/哈希对象/集合对象/有序集合对象

| 类型常量|对象名称|TYPE命令输出|
|---|---|---|---|
|REDIS_STRING|字符串对象|string|
|REDIS_LIST|列表对象|list|
|REDIS_HASH|哈希对象|hash|
|REDIS_SET|集合对象|set|
|REDIS_ZSET|有序集合对象|zset|
#####8.1.2编码和底层实现
不用类型和编码对象

| 类型 |编码|对象|
|---|---|---|
|REDIS_STRING|REDIS_ENCODING_INT|使用整数值实现的字符串对象|
|REDIS_STRING|REDIS_ENCODING_EMBSTR|使用embstr编码的简单动态字符串实现的字符串对象|
|REDIS_STRING|REDIS_ENCODING_RAW|使用简单动态字符串实现的字符串对象|
|REDIS_LIST|REDIS_ENCODING_ZIPLIST|使用压缩列表实现的列表对象|
|REDIS_LIST|REDIS_ENCODING_LINKEDLIST|使用双端链表实现的列表对象|
|REDIS_HASH|REDIS_ENCODING_ZIPLIST|使用压缩列表实现的哈希对象|
|REDIS_HASH|REDIS_ENCODING_HT|使用字典实现的哈希对象|
|REDIS_SET|REDIS_ENCODING_INTSET|使用整数集合实现的集合对象|
|REDIS_SET|REDIS_ENCODING_HT|使用字典实现的集合对象|
|REDIS_ZSET|REDIS_ENCODING_ZIPLIST|使用压缩列表实现的有序集合对象|
|REDIS_ZSET|REDIS_ENCODING_SKIPLIST|使用跳跃表和字典实现的有序集合对象|

````
redis->OBJECT ENCODING story
````
通过encodings属性来设定对象所使用的编码，而不是为特定类型的对象关联一种
固定的编码，极大地提高了Redis的灵活性和效率，因为Redis可以根据不同的使用场景
来为一个对象设置不同的编码，从而优化对象在某一场景下的效率。

在列表对象包含的元素比较少时，Redis使用压缩列表来作为列表对象的实现：
因为压缩列表比双端链表更加节约内存，并且数据元素少时，在内存中以连续块方式保存的
压缩列表比起双端链表可以更快被载入到缓存中
随着列表元素的越来越多，使用压缩列表来保存元素的优势逐渐消失，对象就会将底层实现过压缩
列表转向功能更强，也更合适存储大量元素的双端链表上面。

####8.2字符串对象
使用set命令
````
redis->set msg "hello world"
````
字符串编码:int raw embstr
int：整数，可用long类型表示
raw:大于32个字节 2次内存分配 2次内存释放
embstr 小于32个字节 一次内存分配 1次内存释放一块连续的内存块
######8.2.1串编码的转换
int->raw
embstr->raw


####8.3列表对象
编码:ziplist/linkedlist
````
redis->RPUSH number 1 "three" 5
````
ziplist(压缩列表)编码的结构图

| redisObject  |
|---|
| type REDIS_LIST  |
|encoding REDIS_ENCODING_ZIPLIST|
|ptr|->|zlbytes|zltail|zllen|1|three|5|zlend|
|...|

linkedlist(双端链表)编码的结构图

| redisObject  |
|---|
| type REDIS_LIST  |
|encoding REDIS_ENCODING_ZIPLIST|
|ptr|->|stringObject 1|->|stringObject "three"|->|stringObject 5|
|...|

字符串对象是唯一一个会被其他4种类型对象嵌套的对象

| sdshrd  |
|---|---|---|
|  free 0 | len 5 | buf|
######8.3.2编码转换
当列表对象可以同时满足一下2个条件的时候，列表对象使用ziplist编码
1列表对象保存的所有字符串元素的长度都小于64字节
2列表对象保存的元素数量小于512个；不能满足这2个条件的需要使用linkedlist
######8.3.2列表命令的实现


####8.4哈希对象
编码：ziplist hashtable
ziplist编码的哈希对象使用压缩列表作为底层实现，每当有新的键值对要加入到哈希对象的时候，
程序会先将保存了键的压缩列表节点推入到压缩表的表尾，然后将保存了值得节点推入到压缩列表表
尾。
````
redis->HSET profile name "TOM"
(integer)1
redis->HSET profile age 25
(integer)1
redis->HSET profile career "Programmer"
(integer)1
````

| redisObject  |
|---|
| type REDIS_HASH  |
|encoding REDIS_ENCODING_ZIPLIST|
|ptr|->|压缩列表|

profile哈希对象的压缩列表底层实现

| zlbytes|zltail|zllen|"name"|"Tom"|"age"|25|"career"|"programmer"|
|---|

使用字典作为底层实现
哈希对象中的每个键值对都是使用一个字典键值对来保存：
字典的每个键都是一个字符串对象，对象中保存了键值对的键
字典的每个值都是一个字符串对象，对象中保持了键值对的值

| RedisObject|
|---|
|type REDIS_HASH|
|encoding REDIS_ENCODING_HT|
|ptr|->|dict|
| | |StringObject "age"|->|StringObject 25|
| | |StringObject "career"|->|StringObject "Programmer"|
| | |StringObject "name"|->|StringObject "Tom"|

######8.4.1编码转换
当满足一下2个条件的时候，哈希对象使用ziplist编码
哈希对象保存的所有键值对的键和值的字符串长度都小于64字节
哈希对象保存的键值对数量小于512个；不能满足这2个条件的哈希对象需要使用hashtable编码
可以通过修改
hash_max_ziplist_value/hash_max_ziplist_entries选项说明
######8.4.2哈希命令的实现

####8.5集合对象
集合对象的编码:intset/hashtable
######8.5.1编码的转换
当集合对象可以同时满足一下2个条件时，对象使用intset编码
集合对象保存的所有元素都是整数值
集合对象保存的元素数量不超过512个
set-max-intset-entries修改第二个参数
######8.5.2集合命令的实现

####8.6有序集合对象
