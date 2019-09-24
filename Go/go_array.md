### Go Array 数组

---

```go
type Array struct {
        Elem  *Type // 元素类型
        Bound int64 // 元素数量
}
```

#### 定义方式

Go数组长度固定 

```text
var <数组名称> [<数组长度>]<元素类型>
```

```go
var arr [5]int                            //一维数组，自动赋值0
arr1 := [5]int{1,2,3,4,5}                 //指定元素初始化数组
arr2 := [...]int{2,4,6}                   //省略号初始化数组，让编译器计算长度
arr :=[4]string{0:"1",1:"2",2:"3",3:"4"}  //索引化初始数组
var grid [4][5]int                        //二维数组
arr3 := [5]int{3:1}                       //Go可以直接对数组的某一位置赋值
```

#### 数组遍历

```go
for i:=0;i<len(arr);i++ {
    fmt.Println(arr[i])
}
```

```go
for i := range arr {
    fmt.Println(arr[i])
}
```

```go
for _,v := range arr {
    fmt.Println(v)
}
```

#### 数组是值类型

· [10]int和[5]int是两种类型

Go中长度也是数组类型的一部分

```go
func main() {
    a := [3]int{5, 78, 8}
    var b [5]int
    b = a // a不能赋值给b因为b的类型和a不同，编译器会报 cannot use a (type [3]int) as type [5]int in assignment。
}
```

· 调用函数func(arr [5]int) 会**`拷贝`**数组

**Go语言特殊的一点，数组位值类型，其他语言为了节省内存都是引用类型**。在Go中引用类型为切片slice。

```go
// 函数传参为拷贝数组
func printArray(arr [5]int) {
	arr[0] = 100 //由于是值类型，虽然在这边会改动，但是在main函数，arr[0]并没有修改
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// 传入指针，就会共享同一块内存，即会修改值
func printPointArray(arr *[5]int) {
	arr[0] = 100 //由于是值类型，虽然在这边会改动，但是在main函数，arr[0]并没有修改
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr [5]int                //一维数组
	arr1 := [5]int{1, 2, 3, 4, 5} //初始化数组
	arr2 := [...]int{2, 4, 6}     //让编译器计算长度
	var grid [4][5]int            //二维数组

	fmt.Println(arr, arr1, arr2)
	fmt.Println(grid)

	printArray(arr)
	printArray(arr1)

	fmt.Println(arr, arr1)

	printPointArray(&arr)
	printPointArray(&arr1)

	fmt.Println(arr, arr1)
}
```

```bash
#output
[0 0 0 0 0] [1 2 3 4 5] [2 4 6]
[[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]
0 100
1 0
2 0
3 0
4 0
0 100
1 2
2 3
3 4
4 5
[0 0 0 0 0] [1 2 3 4 5]
0 100
1 0
2 0
3 0
4 0
0 100
1 2
2 3
3 4
4 5
[100 0 0 0 0] [100 2 3 4 5]
```

