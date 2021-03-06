package main

import "fmt"

// 本函数测试入口参数和返回值情况
func dummy(b int) int {
    
    // 声明一个变量c并赋值
    var c int
    c = b
    
    return c
}

// 空函数, 什么也不做
func void() {
}

func main() {
    
    // 声明a变量并打印
    var a int
    
    // 调用void()函数
    void()
    
    // 打印a变量的值和dummy()函数返回
    fmt.Println(a, dummy(0))
}

//代码说明如下：
//第 6 行，dummy() 函数拥有一个参数，返回一个整型值，用来测试函数参数和返回值分析情况。
//第 9 行，声明变量 c，用于演示函数临时变量通过函数返回值返回后的情况。
//第 16 行，这是一个空函数，测试没有任何参数函数的分析情况。
//第 22 行，在 main() 中声明变量 a，测试 main() 中变量的分析情况。
//第 25 行，调用 void() 函数，没有返回值，测试 void() 调用后的分析情况。
//第 28 行，打印 a 和 dummy(0) 的返回值，测试函数返回值没有变量接收时的分析情

//栈（Stack）是一种拥有特殊规则的线性表数据结构。
//栈只允许从线性表的同一端放入和取出数据，按照后进先出（LIFO，Last InFirst Out）的顺序。

//堆在内存分配中类似于往一个房间里摆放各种家具，家具的尺寸有大有小，分配内存时，需要找一块足够装下家具的空间再摆放家具。
//经过反复摆放和腾空家具后，房间里的空间会变得乱七八糟，此时再往这个空间里摆放家具会发现虽然有足够的空间，但各个空间分布在不同的区域，没有一段连续的空间来摆放家具。
//此时，内存分配器就需要对这些空间进行调整优化。
