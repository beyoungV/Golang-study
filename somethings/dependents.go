//恢复初始设置；
go env -u 

//Windows 下开启 GO111MODULE 的命令为：
go env -w GO111MODULE=on 

//MacOS 或者 Linux 下开启 GO111MODULE 的命令为：
export GO111MODULE=on 或者 export GO111MODULE=auto


//Windows 下设置 GOPROXY 的命令为：
go env -w GOPROXY=https://goproxy.cn,direct

//MacOS 或 Linux 下设置 GOPROXY 的命令为：
export GOPROXY=https://goproxy.cn


//go get命令下载指定版本的依赖包
//go get 命令，在下载依赖包的同时还可以指定依赖包的版本。
//go get -u命令会将项目中的包升级到最新的次要版本或者修订版本；
//go get -u=patch命令会将项目中的包升级到最新的修订版本；
//go get [包名]@[版本号]命令会下载对应包的指定版本或者将对应包升级到指定的版本。
//提示：go get [包名]@[版本号]命令中版本号可以是 x.y.z 的形式，例如 go get foo@v1.2.3，也可以是 git 上的分支或 tag，例如 go get foo@master，
//还可以是 git 提交时的哈希值，例如 go get foo@e3702bed2。
