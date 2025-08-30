## 跳跃表
跳跃表(skiplist)是一种有序数据结构，它通过在每个节点中维持多个指向其他节点的指针，从而达到快速访问节点的目的。
跳跃表支持平均O(logN),最坏O(N)复杂度的节点查找，还可以通过顺序性操作来批量处理节点(平衡树)
如果一个有序集合包含的元素数量比较多，又或者成员是比较长的字符串时，redis就会使用跳跃表来作为有序集合键的底层实现。
集群节点中用作内部数据结构
#####5.1跳跃表的实现
redis.h/zskiplistNode:用于表示跳跃表节点
````
typedef struct zskiplistNode {
    sds ele;
    double score;
    struct zskiplistNode *backward;
    struct zskiplistLevel {
        struct zskiplistNode *forword;
        unsigned int span;
    } level[];
}zskiplistNode
````
redis.h/zskiplist 保存跳跃表节点的相关信息
````
typedef struct zskiplist {
    struct zskiplistNode *head,*tail;
    unsigend long length;
    int level;
} zskiplist
````