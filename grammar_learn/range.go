package main

import "fmt"

func main() {
	kvs := map[string]string{}
	kvs["a"] = "python"
	for key, value := range kvs {
		fmt.Println(key, "=>", value)
	}

	var kvs1 map[string]string
	kvs1 = make(map[string]string)
	kvs1["b"] = "Go"
	kvs1["c"] = "Goc"
	// for key, value := range kvs1 {
	// 	fmt.Println(key, "=>", value)
	// 	// goto Loop
	// }

	Boop:
		for key, value := range kvs1 {
			fmt.Println(key, "=>", value)
			break Boop
		}

	// Loop:
	// 	fmt.Print("ok")

}
