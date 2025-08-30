## 整数集合
整数集合（intset)是集合键的底层实现之一，当一个集合只包含整数值元素的时候，并且数量不多，Redis就会使用
整数集合作为集合键的底层实现。
#####6.1整数集合的实现
intset.h/intset
````
typedef struct intset {
    uint32_t encoding;//编码方式
    uint32_t length;//
    int8_t contents[]//保存数组的元素
}
````

#####6.2升级
