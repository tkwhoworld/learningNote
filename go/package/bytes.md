# bytes 包
bytes包实现了操作[]byte的常用函数。本包的函数和strings包的函数相当类似。
bytes包定义了一些常量和变量，以及一堆函数，还有2个类型。

## 函数
这些函数的开头的字母都是大写，因为包要向外暴露函数，函数名必须大写
### func Compare
#### 定义
func Compare (a, b []byte) int
Compare函数返回一个正式表示两个[]byte切片按字典序比较的结果。如果a==b则返回0,如果a < b则返回-1,否则返回1

