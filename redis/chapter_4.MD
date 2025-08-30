## 字典
字典又称为符号表（symbol table)\关联数组（associative array)\映射(map)
是一种用于保存键值对（key-value pair)的抽象数据结构
字典中每个键都是独一无二的。
####4.1字典的实现
Redis的字典使用哈希表作为底层实现，一个哈希表里面可以有多个哈希节点，每个哈希节点就保存了一个字典中的键值对
哈希表dict.h/dictht的结构定义
````
typedef struct dictht{
    dictEntry **table;//哈希表数组
    unsigned long size;//哈希表大小
    unsigned long sizemask;//哈希表大小掩码,用于计算索引值 size-1
    unsigned used;//哈希表已有节点数量
}
````
哈希表节点 dict.h/dictEntry
````
typedof struct dictEntry{
    void *key;//键
    //值
    union{
        void *val;
        uint64_tu64;
        int64_ts64;
    }
    struct dictEntry *next;
}
````

字典 dict.h/dict
````
typedef struct dict{
    //类型特定函数
    dictType *type;
    //私有数据
    void * privdata;
    
    //哈希表
    dictht ht[2]
    
    //rehash索引
    long trehashidx;
    unsigned long iterators;
}
````