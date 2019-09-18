### Go 循环语句 For

---

#### 基础使用

```GO
for i := 0 ; i < 100 ; i++ {
    fmt.Println(i)
}
```

· go语言的for，条件不需要括号。

· 可以省略初始条件，结束条件，递增表达式。

· for 条件内可以定义变量。

· 全部省略就是一个死循环。

```go
for {
    ...
}
```

· 当只有if条件的时候就是一个while

```go
var i = 0 
for i < 100 {
    fmt.Println(i)
    i++
}
```

#### 语法糖 for-range

##### 例子：下面代码会不会停止？

```go
func main() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
}
```

##### range用法

```go
for i := range a {
    fmt.Println(i)
}
```

##### range表达式左侧

在range内可用两种方式assign给左边变量即(i)

```wiki
· =
· :=   这种方式,在go中每次循环都会复用左边宣告的变量
```

##### range表达式右侧类型

可以用于range的类型如下

```wiki
array,pointer to an array,slice,string,map,可接受的channel
```

**range 表达式会在开始循环前被 evaluated 一次,循环开始之前会先把 range 表达式复制给一个变量**.

>evaluated我的理解是值赋予一个变量的动作，可用这个来理解，"append(slice1, 1, 2, 3, 4) evaluated but not used"，相当于是算出一个值后没有把对应的值赋给一个变量。

Go中，无论我们对什么赋值，都会被复制。如果赋值了一个指针，那我们就复制了一个指针副本。如果赋值了一个结构体，那我们就复制了一个结构体副本。
##### range对语法糖的还原

· array

```go
// The loop we generate:
//   len_temp := len(range)
//   range_temp := range
//   for index_temp = 0; index_temp < len_temp; index_temp++ {
//           value_temp = range_temp[index_temp]
//           index = index_temp
//           value = value_temp
//           original body
//   }
```

· slice 

```go
//   for_temp := range
//   len_temp := len(for_temp)
//   for index_temp = 0; index_temp < len_temp; index_temp++ {
//           value_temp = for_temp[index_temp]
//           index = index_temp
//           value = value_temp
//           original body
//   }
```

· channel

```go
//for{
//    v,ok=<-ch
//    if!ok{
//        break
//    }
//    original body
//}
```

##### 回到例子

```go
for_temp := v
len_temp := len(for_temp)
for index_temp = 0; index_temp < len_temp; index_temp++ {
        value_temp = for_temp[index_temp]
        index = index_temp
        value = value_temp
        v = append(v, index)
}
```

在循环开始前对这个结构体生成副本然后赋值给 for_temp，后面的循环实际上是在对 for_temp 进行迭代。任何对于原始变量 v本身（而非对其背后指向的数组）的更改都和生成的副本 for_temp没有关系。但其背后指向的数组还是以指针的形式共享给 v和 for_temp，所以 v[i] = 1这样的语句仍然可以工作。

##### 延伸：range map会发生什么？

```wiki
· 在 range 循环里对 maps 做添加或删除元素的操作是安全的。
· 如果在循环中对 map 添加了一个元素，那么这个元素并不一定会出现在后续的迭代中。
```

第一点，对map的复制，是复制map结构体的指针。所以range对指针副本做操作的内存地址和对maps指针指向的内存地址一样，所以操作的内容会一致。

第二点，map本质是hash table，哈希表内部数组里的元素并不是以特定顺序存放。最后一个添加的元素有可能经过哈希后被放到了内部数组里的第一个索引位，我们确实没有办法预测当前添加的元素是否会出现在后续的迭代中，毕竟在添加元素的时候很可能已经遍历过了第一个索引位。