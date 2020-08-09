package main

import (
    "bufio"
    "fmt"
    "net"
    "strings"
)

// 连接的会话逻辑
func handleSession(conn net.Conn, exitChan chan int) {

    fmt.Println("Session started:")
    
    // 创建一个网络连接数据的读取器
    reader := bufio.NewReader(conn)
    
    // 接收数据的循环
    for {
        // 读取字符串, 直到碰到回车返回
        str, err := reader.ReadString('\n')
        
        // 数据读取正确
        if err == nil {
            // 去掉字符串尾部的回车
            str = strings.TrimSpace(str)
            // 处理Telnet指令
            if !processTelnetCommand(str, exitChan) {
                conn.Close()
                break
            }
            
            // Echo逻辑, 发什么数据, 原样返回
            conn.Write([]byte(str + "\r\n"))
        } else {
            // 发生错误
            fmt.Println("Session closed")
            conn.Close()
            break
        }
    }
}

代码说明如下：
第 11 行是会话入口，传入连接和退出用的通道。handle Session() 函数被并发执行。
第 16 行，使用 bufio 包的 NewReader() 方法，创建一个网络数据读取器，这个 Reader 将输入数据的读取过程进行封装，方便我们迅速获取到需要的数据。
第 19 行，会话处理开始时，从 Socket 连接，通过 reader 读取器读取封包，处理封包后需要继续读取从网络发送过来的下一个封包，因此需要一个会话处理循环。
第 22 行，使用 reader.ReadString() 方法进行封包读取。内部会自动处理粘包过程，直到下一个回车符到达后返回数据。这里认为封包来自 Telnet，每个指令以回车换行符（“\r\n”）结尾。
第 25 行，数据读取正常时，返回 err 为 nil。如果发生连接断开、接收错误等网络错误时，err 就不是 nil 了。
第 28 行，reader.ReadString 读取返回的字符串尾部带有回车符，使用 strings.TrimSpace() 函数将尾部带的回车和空白符去掉。
第 31 行，将 str 字符串传入 Telnet 指令处理函数 processTelnetCommand() 中，同时传入退出控制通道 exitChan。当这个函数返回 false 时，表示需要关闭当前连接。
第 32 行和第 33 行，关闭当前连接并退出会话接受循环。
第 37 行，将有效数据通过 conn 的 Write() 方法写入，同时在字符串尾部添加回车换行符（“\r\n”），数据将被 Socket 发送给连接方。
第 41～43 行，处理当 reader.ReadString() 函数返回错误时，打印错误信息并关闭连接，退出会话并接收循环。
