## 简单动态字符串SDS(Simple Dynamic String)
###概述 
redis并没有使用c语言传统的字符串表示，而是自己构建了一个名为简单动态字符串的抽象类型，并将SDS作为Redis的默认字符串表示
###1.定义:src/sds.h/sdshdr
````
{
    int len;//buf数组中的已使用的字节数量,等于SDS所保存的字符的长度
    int free;//记录buf中未使用的字节数
    char buf[];//字节数组，用来保存一系列二进制数据
}
````

|sdshdr|
|---|
|free 0|
|len 5|
|buf|->|'R'|'e'|'d'|'i'|'s'|'\0'|

###2.SDS与C字符的区别
C语言使用长度为N+1的字符数组来表示长度为N的字符串，并且字符数组最后一个总是空字符'\0'.
#####1)获取字符串长度的的复杂度
c语言获取的复杂度为N
因为c语言的字符串是存在字符数组里的，所以要获取长度，需要循环，循环次数为字符串的长度，复杂度为O(N)
如果通过SDS获取字符串长度，则直接通过SDS->len就可以获取，复杂度为O(1)
#####2).杜绝缓冲区溢出
c字符串不记录自身长度带来的另一个问题是容易造成缓冲区溢出（buffer overflow)
举例说明:</string.h>/strcat函数
````
char *strcat(char *dest,const char *src);
````
因为C字符串不记录自身的长度，所以strcat假设用户在执行这个函数的时，已经为dest分配了足够多的内存，可以容纳src的字符串中所有的内容，
但是这个假设不成立的时候，就会造成缓冲区溢出
SDS的strcat的代码如下:
````
sds sdscat(sds s, const char *t) {
    return sdscatlen(s, t, strlen(t));
}

sds sdscatlen(sds s, const void *t, size_t len) {
    size_t curlen = sdslen(s);

    s = sdsMakeRoomFor(s,len);
    if (s == NULL) return NULL;
    memcpy(s+curlen, t, len);
    sdssetlen(s, curlen+len);
    s[curlen+len] = '\0';
    return s;
}

sds sdsMakeRoomFor(sds s, size_t addlen) {
    void *sh, *newsh;
    size_t avail = sdsavail(s);
    size_t len, newlen;
    char type, oldtype = s[-1] & SDS_TYPE_MASK;
    int hdrlen;

    /* Return ASAP if there is enough space left. */
    if (avail >= addlen) return s;

    len = sdslen(s);
    sh = (char*)s-sdsHdrSize(oldtype);
    newlen = (len+addlen);
    if (newlen < SDS_MAX_PREALLOC)
        newlen *= 2;
    else
        newlen += SDS_MAX_PREALLOC;

    type = sdsReqType(newlen);

    /* Don't use type 5: the user is appending to the string and type 5 is
     * not able to remember empty space, so sdsMakeRoomFor() must be called
     * at every appending operation. */
    if (type == SDS_TYPE_5) type = SDS_TYPE_8;

    hdrlen = sdsHdrSize(type);
    if (oldtype==type) {
        newsh = s_realloc(sh, hdrlen+newlen+1);
        if (newsh == NULL) return NULL;
        s = (char*)newsh+hdrlen;
    } else {
        /* Since the header size changes, need to move the string forward,
         * and can't use realloc */
        newsh = s_malloc(hdrlen+newlen+1);
        if (newsh == NULL) return NULL;
        memcpy((char*)newsh+hdrlen, s, len+1);
        s_free(sh);
        s = (char*)newsh+hdrlen;
        s[-1] = type;
        sdssetlen(s, len);
    }
    sdssetalloc(s, newlen);
    return s;
}
````
sdsMakeRoomFor方法实现了在SDS原有的基础上来实现防止缓冲区溢出
#####3).减少了修改字符串时带来的内存重新分配次数
因为C字符串并不记录自身的长度，所以对于一个包含N个字符的C字符串来说，这个C字符串的底层实现总是一个
N+1个字符长的数组。所以对c字符的每一次操作都会修改c的长度，从而需要内存重新分配
为了避免这种缺陷，SDS通过未使用空间解除了字符串长度和底层数组长度之间的关联，通过未使用空间，SDS实现了空间预分配和惰性空间释放
2种优化策略
1.空间预分配
````
#define SDS_MAX_PREALLOC (1024*1024)//MB
````
分配公式:大于1mb,分配的时候，size=len+len(free)+1
        大于1mb,分配的时候,size=len+free(1MB)+1
2.惰性释放空间
SDS需要通过API缩短SDS字符串的时候，不会立即释放内存，而是修改free的值
#####4).二进制安全
C字符串中的字符必须符合某种编码（ASCII),并且除了字符串的末尾之外，字符串里面不能包含字符，否则最先读入的空字符会被当做结尾处理。
所以不能保存像图片、音频、视频那样的二进制数据
