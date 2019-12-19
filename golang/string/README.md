# 运行方法
go build;
./string -func Test1

# 简单结论
**string和[]byte之间的转化是有copy的，可以借助unsafe实现无copy的转化(1-5)**
**显示定义的字符串会被分配到只读内存区，通过[]byte以及unsafe都无法更改其值(1-5)**
**只读内存区的字符串会被复用，即显示定义两个相同内容的字符串，都指向同一个内容(6)**
**常量、拼接次数少的考虑用+，定义好的[]string用join，拼接次数多的用stringBuilder**