package main;
 
import (
	"bytes"
	"fmt"
)
 
//bytes中Reader和Buffer两个结构的使用
 
func useReader() {
	data := "abcdefghijk";
	//通过[]byte创建Reader
	re := bytes.NewReader([]byte(data));
	//返回未读取部分的长度
	fmt.Println("re len : ",re.Len());
	//返回底层数据总长度
	fmt.Println("re size : ",re.Size());
 
	buf := make([]byte,2);
	for {
		//读取数据
		n,err := re.Read(buf);
		if err != nil {
			break;
		}
		fmt.Println(string(buf[:n]));
	};
 
	//设置偏移量，因为上面的操作已经修改了读取位置等信息
	re.Seek(0,0);
	for {
		//一个字节一个字节的读
		b,err := re.ReadByte();
		if err != nil {
			break;
		}
		fmt.Println(string(b));
	}
 
	re.Seek(0,0);
	off := int64(0);
	for {
		//指定偏移量读取
		n,err := re.ReadAt(buf,off);
		if err != nil {
			break;
		}
		off += int64(n);
		fmt.Println(off,string(buf[:n]));
	}
}
 
func useBuffer() {
	data := "123456789";
	//通过[]byte创建一个Buffer
	bf := bytes.NewBuffer([]byte(data));
 
	//Len()返回未读取的数据长度
	fmt.Println("bf len : ",bf.Len());
 
	//Cap()缓存容量
	fmt.Println("bf cap : ",bf.Cap());
 
	//Bytes()返回未读取的数据切片
	bys := bf.Bytes();
	for _,v := range bys {
		fmt.Print(string(v) + " ");
	}
	fmt.Println();
 
	//Next()返回未读取部分前n字节数据的切片
	for i := 0; i < 10; i++ {
		tmp := bf.Next(1);
		fmt.Print(string(tmp) + " ");
	}
	fmt.Println();
	//再次Next，返回[]byte，说明没有未读取的
	fmt.Println(bf.Next(1));
 
	//重设缓冲，丢弃全部内容
	bf.Reset();
 
	//通过string创建Buffer
	bf2 := bytes.NewBufferString(data);
	//读取第一个 delim 及其之前的内容，返回遇到的错误
	line,_ := bf2.ReadBytes('3');
	fmt.Println(string(line));
	//效果同上，返回string
	line2,_ := bf2.ReadString('7');
	fmt.Println(line2);
 
	//创建一个空Buffer
	bf3 := bytes.Buffer{};
	//自动增加缓存容量，保证有n字节剩余空间
	bf3.Grow(16);
	//写入rune编码，返回写入的字节数和错误。
	n,_ := bf3.WriteRune(rune('中'));
	fmt.Println("bf3 write ",n);
	n,_ = bf3.WriteString("国人");
	fmt.Println("bf3 write ",n);
	//返回未读取的字符串
	fmt.Println(bf3.String());
	//将数据长度截断到n字节
	bf3.Truncate(6);
	fmt.Println(bf3.String());
}
 
func main() {
	//防止main中代码过多，我新建两个函数单独写
	useReader();
	useBuffer();
}
