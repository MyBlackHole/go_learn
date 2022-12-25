package main

import (
	"fmt"
	"github.com/deckarep/golang-set"
)

// // NewSet创建并返回空集的引用，结果集上的操作是线程安全的
// func NewSet(s ...interface{}) Set {}
// // NewSetFromSlice从现有切片创建并返回集合的引用，结果集上的操作是线程安全的
// func NewSetFromSlice(s []interface{}) Set {}
// // NewSetWith创建并返回具有给定元素的新集合，结果集上的操作是线程安全的
// func NewSetWith(elts ...interface{}) Set {}
// // NewThreadUnsafeSet创建并返回对空集的引用，结果集上的操作是非线程安全的
// func NewThreadUnsafeSet() Set {}
// // NewThreadUnsafeSetFromSlice创建并返回对现有切片中集合的引用，结果集上的操作是非线程安全的。
// func NewThreadUnsafeSetFromSlice(s []interface{}) Set {}

func main() {

	// type void struct{}
	// var member void

	// set := make(map[string]void) // New empty set
	// set["Foo"] = member          // Add
	// for k := range set {         // Loop
	// 	fmt.Println(k)
	// }
	// delete(set, "Foo")      // Delete
	// size := len(set)        // Size
	// _, exists := set["Foo"] // Membership

	// 默认创建的线程安全的，如果无需线程安全
	// 可以使用 NewThreadUnsafeSet 创建，使用方法都是一样的。
	s1 := mapset.NewSet(1, 2, 3, 4)
	fmt.Println("s1 contains 3: ", s1.Contains(3))
	fmt.Println("s1 contains 5: ", s1.Contains(5))

	// interface 参数，可以传递任意类型
	s1.Add("poloxue")
	fmt.Println("s1 contains poloxue: ", s1.Contains("poloxue"))
	s1.Remove(3)
	fmt.Println("s1 contains 3: ", s1.Contains(3))

	s2 := mapset.NewSet(1, 3, 4, 5)

	// 并集
	fmt.Println(s1.Union(s2))
}
