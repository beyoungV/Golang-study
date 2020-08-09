package main

import (

    "fmt"

    "strings"

)

func processTelnetCommand(str string, exitChan chan int) bool {

    // @close指令表示终止本次会话

    if strings.HasPrefix(str, "@close") {

        fmt.Println("Session closed")

        // 告诉外部需要断开连接

        return false

        // @shutdown指令表示终止服务进程

    } else if strings.HasPrefix(str, "@shutdown") {

        fmt.Println("Server shutdown")

        // 往通道中写入0, 阻塞等待接收方处理

        exitChan <- 0

        // 告诉外部需要断开连接

        return false

    }

    // 打印输入的字符串

    fmt.Println(str)

    return true

}

代码说明如下：

第 8 行，处理 Telnet 命令的函数入口，传入有效字符并退出通道。

第 11～16 行，当输入字符串中包含“@close”前缀时，在第 16 行返回 false，表示需要关闭当前会话。

第 19～27 行，当输入字符串中包含“@shutdown”前缀时，第 24 行将 0 写入 exitChan，表示结束服务器。

第 31 行，没有特殊的控制字符时，打印输入的字符串。
