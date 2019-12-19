有两种形式的边界检测：索引（a[i]）和切片（a[i:j]）
array, slice, string

go build -gcflags="-d=ssa/check_bce/debug=1" bound.go