## 链表
### 概述
链表提供了高效的节点重排能力，以及顺序性的节点访问方式。并且通过增删节点来调整链表的长度
场景：列表键，发布与订阅，慢查询，监视器，Redis服务器本身还使用链表来保存多个客户端的状态信息
以及使用链表构建客户端输出缓冲区（output buffer)
###1.链表和链表节点的实现
链表节点:src/adlist.h/ListNode结构来表示
````
    typedef struct listNode{
        //前置节点
        struct listNode *prev;
        //后置节点
        struct listNode *next;
        //节点的值
        void * value
    }listNode;
````
虽然仅仅使用多个listNode结构就可以组成链表，但使用adlist.h/list来持有链表的话，更加方便
````
    typedef struct list{
        //表头节点
        listNode * head;
        //表尾节点
        listNode * tail;
        //链表所包含的节点数
        unsigned long len;
        //节点复制函数
        void *(*dup)(void *prt);
        //节点释放函数
        void (*free)(void *ptr);
        //节点值对比函数
        int (*match)(void *ptr,void *key);
    }list;
````
###2.链表实现的特点
#####1)双端：链表节点带有prev和next指针，获取某个节点的前置节点和后置节点的复杂度都是O(1)
#####2)无环：表头节点的prev指针和表尾节点的next指针都是指向null,对链表的访问以null结束
#####3）带表头指针和表尾指针：通过list结构的head指针和tail指针，程序获取链表节点的表头节点和表尾节点的复杂度为O(1)
#####4)带链表长度计数器，程序获取链中节点数量的复杂度为O(1)
#####5)多态
###链表和链表节点的api
#####1)listSetDupMethod 作用:将给定的函数设置为链表的节点复制函数 时间复杂度：复制函数可以通过链表的dup属性直接获得,O(1)
````
    #define listSetDupMethod(l,m) ((l)->dup = (m))
````
解析：l是list（链表)(l)->dup=(m)就是将m这个方法赋给l的dup属性
#####2）listGetDupMethod 作用:返回链表当前正在使用的复制函数 O(1)
````
    #define listGetDupMethod(l)((l)->dup)
````
解析:l是list(链表)(l)->dup获取复制函数
下面这几个api类似
listSetFreeMethod/listGetFree/listSetMatchMethod/listGetMatchMethod