package main
import (
    "image"
    "image/color"
    "image/png"
    "log"
    "math"
    "os"
)
func main() {
    // 图片大小
    const size = 300
    // 根据给定大小创建灰度图
    pic := image.NewGray(image.Rect(0, 0, size, size))
    // 遍历每个像素
    for x := 0; x < size; x++ {
        for y := 0; y < size; y++ {
            // 填充为白色
            pic.SetGray(x, y, color.Gray{255})
        }
    }
    // 从0到最大像素生成x坐标
    for x := 0; x < size; x++ {
        // 让sin的值的范围在0~2Pi之间
        s := float64(x) * 2 * math.Pi / size
        // sin的幅度为一半的像素。向下偏移一半像素并翻转
        y := size/2 - math.Sin(s)*size/2
        // 用黑色绘制sin轨迹
        pic.SetGray(x, int(y), color.Gray{0})
    }
    // 创建文件
    file, err := os.Create("sin.png") 
    if err != nil {
        log.Fatal(err)
    }
    // 使用png格式将数据写入文件
    png.Encode(file, pic) //将image信息写入文件中
    // 关闭文件
    file.Close()
}



//代码说明如下：
//第 12 行，声明一个 size 常量，值为 300。
//第 14 行，使用 image 包的 NewGray() 函数创建一个图片对象，使用区域由 image.Rect 结构提供，image.Rect 描述一个方形的两个定位点 (x1,y1) 和 (x2,y2)，image.Rect(0,0,size,size) 表示使用完整灰度图像素，尺寸为宽 300，长 300。
//第 16 行和第 17 行，遍历灰度图的所有像素。
//第 19 行，将每一个像素的灰度设为 255，也就是白色。
//灰度图是一种常见的图片格式，一般情况下颜色由 8 位组成，灰度范围为 0～255，0 表示黑色，255 表示白色。
//初始化好的灰度图默认的灰度值都是 0，对的是黑色，由于显示效果的效果不是很好，所以这里将所有像素设置为 255，也就是白色

//第 23 行，生成 0 到 size（300）的 x 坐标轴。
//第 25 行，计算 math.Sin 的定义域，这段代码等效为：
//rate := x / size
//s := rate * 2 * math.Pi

//第 32 行，创建 sin.png 的文件。
//第 33 行，如果创建文件失败，返回错误，打印错误并终止。
//第 37 行，使用 PNG 包，将图形对象写入文件中。
//第 39 行，关闭文件。
