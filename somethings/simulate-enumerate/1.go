package main

import "fmt"

// 声明芯片类型
type ChipType int

const (
    None ChipType = iota
    CPU    // 中央处理器
    GPU    // 图形处理器
)

func (c ChipType) String() string {
    switch c {
    case None:
        return "None"
    case CPU:
        return "CPU"
    case GPU:
        return "GPU"
    }
    return "N/A"
}

func main() {
    // 输出CPU的值并以整型格式显示
    fmt.Printf("%s %d", CPU, CPU)
}

//代码说明如下：
//第 6 行，将 int 声明为 ChipType 芯片类型。
//第 9 行，将 const 里定义的常量值设为 ChipType 类型，且从 0 开始，每行值加 1。
//第 14 行，定义 ChipType 类型的方法 String()，返回值为字符串类型。
//第 15～22 行，使用 switch 语句判断当前的 ChitType 类型的值，返回对应的字符串。
//第 28 行，按整型的格式输出 CPU 的值。

//String() 方法的 ChipType 在使用上和普通的常量没有区别。当这个类型需要显示为字符串时，Go语言会自动寻找 String() 方法并进行调用。
