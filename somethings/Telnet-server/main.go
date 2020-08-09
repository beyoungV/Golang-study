package main

import (
    "os"
)

func main() {

    // 创建一个程序结束码的通道
    exitChan := make(chan int)

    // 将服务器并发运行
    go server("127.0.0.1:7001", exitChan)

    // 通道阻塞, 等待接收返回值
    code := <-exitChan

    // 标记程序返回值并退出
    os.Exit(code)
}

代码说明如下：
第 10 行，创建一个整型的无缓冲通道作为退出信号。
第 13 行，接受连接的过程可以并发操作，使用 go 将 server() 函数开启 goroutine。
第 16 行，从 exitChan 中取出返回值。如果取不到数据就一直阻塞。
第 19 行，将程序返回值传入 os.Exit() 函数中并终止程序。
