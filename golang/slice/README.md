# 执行方式
go build
./slice -func Test1

# benchmark执行方式
go test -bench=. test4_test.go test4.go main.go

# 简单结论
**slice 由三个成员组成  data, len, cap**
**data指向底层的数组**
**len slice的长度**
**cap slice的容量**
**常用操作 len(), cap(), copy(), append()**
**1. 默认len是0， cap是0， 关联到底层数组后，len是关联的元素个数，cap是数组的长度**
**2. append操作可能会触发扩容操作，cap会按照2倍扩容，会产生新数组，并且进行关联。扩容后slice地址不变，底层数组地址改变**
**3. slice作为函数参数使用时，函数内对slice的赋值操作会影响外部slice，append不会。如果需要内部append生效，考虑使用指针**
**4. copy() 不会改变slice的len和cap，只是底层数组值的copy**