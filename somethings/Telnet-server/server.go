package main

import (
    "fmt"
    "net"
)

// 服务逻辑, 传入地址和退出的通道
func server(address string, exitChan chan int) {
    
    // 根据给定地址进行侦听
    l, err := net.Listen("tcp", address)
    
    // 如果侦听发生错误, 打印错误并退出
    if err != nil {
        fmt.Println(err.Error())
        exitChan <- 1
    }
    
    // 打印侦听地址, 表示侦听成功
    fmt.Println("listen: " + address)
    
    // 延迟关闭侦听器
    defer l.Close()
    
    // 侦听循环
    for {
        // 新连接没有到来时, Accept是阻塞的
        conn, err := l.Accept()
        // 发生任何的侦听错误, 打印错误并退出服务器
        if err != nil {
            fmt.Println(err.Error())
            continue
        }
        
        // 根据连接开启会话, 这个过程需要并行执行
        go handleSession(conn, exitChan)
    }
}

//代码说明如下：
第 9 行，接受连接的入口，address 为传入的地址，退出服务器使用 exitChan 的通道控制。往 exitChan 写入一个整型值时，进程将以整型值作为程序返回值来结束服务器。
第 12 行，使用 net 包的 Listen() 函数进行侦听。这个函数需要提供两个参数，第一个参数为协议类型，本例需要做的是 TCP 连接，因此填入“tcp”；address 为地址，格式为“主机:端口号”。
第 15 行，如果侦听发生错误，通过第 17 行，往 exitChan 中写入非 0 值结束服务器，同时打印侦听错误。
第 24 行，使用 defer，将侦听器的结束延迟调用。
第 27 行，侦听开始后，开始进行连接接受，每次接受连接后需要继续接受新的连接，周而复始。
第 30 行，服务器接受了一个连接。在没有连接时，Accept() 函数调用后会一直阻塞。连接到来时，返回 conn 和错误变量，conn 的类型是 *tcp.Conn。
第 33 行，某些情况下，连接接受会发生错误，不影响服务器逻辑，这时重新进行新连接接受。
第 39 行，每个连接会生成一个会话。这个会话的处理与接受逻辑需要并行执行，彼此不干扰。
