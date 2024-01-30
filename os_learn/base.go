// //获取主机名
// os.Hostname()
// //获取当前目录
// os.Getwd()
// //获取用户ID
// os.Getuid()
// //获取有效用户ID
// os.Geteuid()
// //获取组ID
// os.Getgid()
// //获取有效组ID
// os.Getegid()
// //获取进程ID
// os.Getpid()
// //获取父进程ID
// os.Getppid()
// //获取环境变量的值
// os.Getenv("GOPATH")
// //设置环境变量的值
// os.Setenv("NAME", "lqz")

// //清除所有环境变量（慎用）
// os.Clearenv()

// //改变当前工作目录
// os.Chdir("/Users/liuqingzheng")

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//获取主机名
	fmt.Println(os.Hostname())
	//获取当前目录
	fmt.Println(os.Getwd())
	//获取用户ID
	fmt.Println(os.Getuid())
	//获取有效用户ID
	fmt.Println(os.Geteuid())
	//获取组ID
	fmt.Println(os.Getgid())
	//获取有效组ID
	fmt.Println(os.Getegid())
	//获取进程ID
	fmt.Println(os.Getpid())
	//获取父进程ID
	fmt.Println(os.Getppid())
	//获取环境变量的值
	fmt.Println(os.Getenv("GOPATH"))
	//设置环境变量的值
	os.Setenv("NAME", "lqz")

	//清除所有环境变量（慎用）
	//os.Clearenv()
	//改变当前工作目录
	os.Chdir("/Users/liuqingzheng")
	fmt.Println(os.Getwd())
}

func main1() {
	//创建目录
	os.Mkdir("abc", os.ModePerm) // 0777
	//创建多级目录
	os.MkdirAll("xxxx/x", os.ModePerm)
	//删除文件或目录
	os.Remove("xxx/x")
	//删除指定目录下所有文件
	os.RemoveAll("xxxxx")
	//重命名文件
	os.Rename("./test1.txt", "./twst1_new.txt")
	// 获取文件信息-->FileInfo接口类型，有很多方法，如文件大小，是否是文件夹等等
	info, _ := os.Stat("./lqz")
	fmt.Println(info.Size())
	//创建文件
	f1, _ := os.Create("./test.txt")
	defer f1.Close()
	//修改文件权限
	os.Chmod("./test.txt", 0777)
	////修改文件所有者
	os.Chown("./test.txt", 0, 0)
	//修改文件的访问时间和修改时间
	os.Chtimes("./test.txt", time.Now().Add(time.Hour), time.Now().Add(time.Hour))
	//打开文件
	//func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
	//以读写方式打开文件，如果不存在，则创建
	file, err := os.OpenFile("./test2.txt", os.O_RDWR|os.O_CREATE, 0766)
}

// //*********flag 取值*********
// const (
// 	O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
// 	O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
// 	O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
// 	O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
// 	O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
// 	O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
// 	O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
// 	O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
// )

// //*********fFileMode取值*********f
// const (
// 	// 单字符是被String方法用于格式化的属性缩写。
// 	ModeDir        FileMode = 1 << (32 - 1 - iota) // d: 目录
// 	ModeAppend                                     // a: 只能写入，且只能写入到末尾
// 	ModeExclusive                                  // l: 用于执行
// 	ModeTemporary                                  // T: 临时文件（非备份文件）
// 	ModeSymlink                                    // L: 符号链接（不是快捷方式文件）
// 	ModeDevice                                     // D: 设备
// 	ModeNamedPipe                                  // p: 命名管道（FIFO）
// 	ModeSocket                                     // S: Unix域socket
// 	ModeSetuid                                     // u: 表示文件具有其创建者用户id权限
// 	ModeSetgid                                     // g: 表示文件具有其创建者组id的权限
// 	ModeCharDevice                                 // c: 字符设备，需已设置ModeDevice
// 	ModeSticky                                     // t: 只有root/创建者能删除/移动文件
// 	// 覆盖所有类型位（用于通过&获取类型位），对普通文件，所有这些位都不应被设置
// 	ModeType          = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
// 	ModePerm FileMode = 0777 // 覆盖所有Unix权限位（用于通过&获取类型位）
// )
