# GoDataStructure

>　常見但是golang沒有內建的資料結構

- [Stack](#Stack)
- [Queue](#Queue)

## 結構

### Queue

```go
queue.go  //使用node製作
　　node  //結點的物件
　　　　3 properties
　　　　value
　　　　prev
　　　　next
　　　　0 functions
　　Queue //list的物件
　　　　 3 properties
　　　　 head //list的頭(座標)
　　　　 rear //list的尾(座標)
　　　　 length //list長度,用來判斷此list的大小
　　　　 5 functions
　　　　 Len
　　　　 Prt
　　　　 Push
　　　　 Peek
　　　　 Get
```

### Stack

```go
stack.go  //使用array製作
　　Stack
　　　　 1 properties
　　　　 data //存放資料的陣列
　　　　 6 functions
　　　　 Len
　　　　 Prt
　　　　 Push
　　　　 Peek
　　　　 Pop
　　　　 Empty
```

## Version

### [1.0.0] 2020.05.01

#### ADD

- Stack
- Queue
