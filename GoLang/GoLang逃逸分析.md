# GoLang 逃逸分析

**逃逸分析是编译器用来决定你的程序中值的位置的过程**。**特别地，编译器执行静态代码分析，以确定一个构造体的实例化值是否会逃逸到堆**。在 Go 语言中，你没有可用的关键字或者函数，能够直接让编译器做这个决定。只能够通过你写代码的方式来作出这个决定。

## 逃逸场景

### 指针逃逸

Go 可以返回局部变量指针，这其实是一个典型的变量逃逸案例，示例代码如下：

```go
package main

type Student struct {
    Name string
    Age  int
}

func StudentRegister(name string, age int) *Student {
    s := new(Student) //局部变量s逃逸到堆

    s.Name = name
    s.Age = age

    return s
}

func main() {
    StudentRegister("Jim", 18)
}
```

### 栈空间不足逃逸（空间开辟过大）

```go
package main

func Slice() {
    s := make([]int, 10000, 10000)

    for index, _ := range s {
        s[index] = index
    }
}

func main() {
    Slice()
}
```

当栈空间不足以存放当前对象时或无法判断当前切片长度时会将对象分配到堆中。

### 动态类型逃逸（不确定长度大小）

很多函数参数为 interface 类型，比如 fmt.Println(a …interface{})，编译期间很难确定其参数的具体类型，也能产生逃逸。

```go
package main

import "fmt"

func main() {
    s := "Escape"
    fmt.Println(s)
}
```

### 闭包引用对象逃逸

```go
package main

import "fmt"

func Fibonacci() func() int {
    a, b := 0, 1
    return func() int {
        a, b = b, a+b
        return a
    }
}

func main() {
    f := Fibonacci()

    for i := 0; i < 10; i++ {
        fmt.Printf("Fibonacci: %d\n", f())
    }
}
```

Fibonacci()函数中原本属于局部变量的 a 和 b 由于闭包的引用，不得不将二者放到堆上，以致产生逃逸。

## 逃逸分析的作用

1. 逃逸分析的好处是为了**减少gc的压力**，不逃逸的对象分配在栈上，当函数返回时就回收了资源，不需要gc标记清除。
2. 逃逸分析完后可以确定哪些变量可以分配在栈上，**栈的分配比堆快，性能好**(逃逸的局部变量会在堆上分配 ,而没有发生逃逸的则有编译器在栈上分配)。
3. **同步消除**，如果你定义的对象的方法上有同步锁，但在运行时，却只有一个线程在访问，此时逃逸分析后的机器码，会去掉同步锁运行。

## 逃逸总结

- 栈上分配内存比在堆中分配内存有更高的效率
- 栈上分配的内存不需要GC处理
- 堆上分配的内存使用完毕会交给GC处理
- 逃逸分析目的是决定内分配地址是栈还是堆
- 逃逸分析在编译阶段完成

## 函数传递指针真的比传值效率高吗

我们知道传递指针可以减少底层值的拷贝，可以提高效率，但是如果拷贝的数据量小，由于指针传递会产生逃逸，可能会使用堆，也可能会增加GC的负担，所以传递指针不一定是高效的。
