package main;
import (
	"bytes"
	"fmt"
	"unicode"
)
//bytes包中实现了大量对[]byte操作的函数和两个最主要的Reader和Buffer两个结构
func main() {
	str := "aBcD";
	//转为小写
	fmt.Println(string(bytes.ToLower([]byte(str))));
	//转为大写
	fmt.Println(string(bytes.ToUpper([]byte(str))));
	//转为标题，
	fmt.Println(string(bytes.ToTitle([]byte(str))));
	//自定义映射表
	mycase := unicode.SpecialCase{
		unicode.CaseRange{
			//1,1表示替换规则只影响1到1之间的字符
			1,1,[unicode.MaxCase]rune{
				//大写转换
				'壹' - 1,//小写转换
				'一' - 1,//标题转换
				'小' - 1,},unicode.CaseRange{
			2,2,[unicode.MaxCase]rune{
				'贰' - 2,'二' - 2,'中' - 2,unicode.CaseRange{
			3,3,[unicode.MaxCase]rune{
				'叁' - 3,'三' - 3,'大' - 3,};
	//使用映射表将[]byte中字符修改为小写
	data := bytes.ToLowerSpecial(
		mycase,[]byte{1,3},)
	fmt.Println(string(data));
	//使用映射表将[]byte中字符修改为大写
	data = bytes.ToUpperSpecial(
		mycase,);
	fmt.Println(string(data));
	//使用映射表将[]byte中字符修改为标题
	data = bytes.ToTitleSpecial(
		mycase,);
	fmt.Println(string(data));
 
	//将[]byte中单词首字符修改为Title并返回
	fmt.Println(string(bytes.Title([]byte("abc def"))));
 
	//比较两个[]byte，
	// a < b 返回 -1
	// a == b 返回 0
	// b > b 返回 1
	fmt.Println(bytes.Compare([]byte("a"),[]byte("b")));
 
	//比较两个[]byte是否相等
	fmt.Println(bytes.Equal([]byte("abc"),[]byte("abc")));
 
	//比较两个[]byte是否相等，忽略大写，小写，标题
	fmt.Println(bytes.EqualFold([]byte("ABC"),[]byte("abc")));
 
	//去掉[]byte两边包含在cutset中的字符
	fmt.Println(string(bytes.Trim([]byte(" abc ")," ")));
 
	//去掉左边包含在cutset中的字符
	fmt.Println(string(bytes.TrimLeft([]byte(" abc ")," ")));
 
	//去掉右边包含在cutset中的字符
	fmt.Println(string(bytes.TrimRight([]byte(" abc ")," ")));
 
	//去掉两边空白字符
	fmt.Println(string(bytes.TrimSpace([]byte(" abc "))));
 
	//去掉前缀
	fmt.Println(string(bytes.TrimPrefix([]byte("tb_user"),[]byte("tb_"))));
 
	//去掉后缀
	fmt.Println(string(bytes.TrimSuffix([]byte("user_idx"),[]byte("_idx"))));
 
	//以sep为分隔符，切分为多个[]byte
	tmp := bytes.Split([]byte("ab cd ef"),[]byte(" "));
	for _,v := range tmp {
		fmt.Println(string(v));
	}
 
	//分割最多n个子切片，超出n的部分将不进行切分
	tmp = bytes.SplitN([]byte("ab cd ef"),[]byte(" "),2);
	for _,v := range tmp {
		fmt.Println(string(v));
	}
 
	//以sep为分隔符，切分为多个[]byte，结果包含分隔符，在子串尾部
	tmp = bytes.SplitAfter([]byte("ab,cd,ef"),[]byte(","));
	for _,v := range tmp {
		fmt.Println(string(v));
	}
 
	//分割最多n个子切片，超出n的部分将不进行切分
	tmp = bytes.SplitAfterN([]byte("ab,"),v := range tmp {
		fmt.Println(string(v));
	}
 
	//以空白字符切分
	tmp = bytes.Fields([]byte("a b c d"));
	for _,v := range tmp {
		fmt.Println(string(v));
	}
 
	//以符合函数的字符作为分隔符来切分
	tmp = bytes.FieldsFunc([]byte("asbscsd"),func(r rune) bool {
		if r == rune('s') {
			return true;
		}
		return false;
	});
	for _,v := range tmp {
		fmt.Println(string(v));
	}
 
	//以sep为连接符，拼接[][]byte
	fmt.Println(string(bytes.Join(
		[][]byte{
			[]byte("aa"),[]byte("bb"),[]byte("cc"),[]byte("-"),)));
 
	//重复[]byte，Count次
	fmt.Println(string(bytes.Repeat([]byte("abc"),3)));
 
	//判断是否有前缀
	fmt.Println(bytes.HasPrefix([]byte("is_true"),[]byte("is_")));
 
	//判断是否有后缀
	fmt.Println(bytes.HasSuffix([]byte("chk_on"),[]byte("_on")));
 
	//判断是否包含某个[]byte
	fmt.Println(bytes.Contains([]byte("i am jack"),[]byte("jack")));
 
	//判断是否包含某个rune
	fmt.Println(bytes.ContainsRune([]byte("i from 中国"),rune('中')));
 
	//查找sep在参数一中第一次出现的位置，找不到返回-1
	fmt.Println(bytes.Index([]byte("abcabc"),[]byte("a")));
	fmt.Println(bytes.IndexByte([]byte("cba"),'a'));
	fmt.Println(bytes.IndexRune([]byte("i from 中国"),rune('中')));
 
	//查找chars中任意一个字符在参数一中出现的位置，找不到返回-1
	fmt.Println(bytes.IndexAny([]byte("hello world"),"xy"));
 
	//功能同上，只不过查找最后一次出现的位置
	fmt.Println(bytes.LastIndex([]byte("abcabc"),[]byte("a")));
	fmt.Println(bytes.LastIndexByte([]byte("cba"),'a'));
	fmt.Println(bytes.LastIndexAny([]byte("hello world"),"xy"));
 
	//获取sep中在参数一中出现的次数
	fmt.Println(bytes.Count([]byte("a|b|c"),[]byte("|")));
 
	//将参数一中前n个old替换成new，n小于0则全部替换。
	fmt.Println(string(
		bytes.Replace(
			[]byte("i am jack"),[]byte("i am"),[]byte("我是"),-1,),));
 
	//将[]byte中的字符替换为函数的返回值，如果返回值为负数，则丢弃访字符。
	fmt.Println(string(
		bytes.Map(
			func(r rune) rune {
				if r == 'a' {
					return 'A';
				} else if r == 'c' {
					return -1;
				}
				return r;
			},[]byte("abcd"),));
 
	//将[]byte转换为[]rune
	fmt.Println(string(bytes.Runes([]byte("我是谁"))));
}
