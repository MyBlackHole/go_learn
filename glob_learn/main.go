package main

import "github.com/gobwas/glob"
import "fmt"

func main() {
	var g glob.Glob

	// create simple glob
	g = glob.MustCompile("*.github.com")
	fmt.Println(g.Match("api.github.com")) // true

	// quote meta characters and then create simple glob
	g = glob.MustCompile(glob.QuoteMeta("*.github.com"))
	fmt.Println(g.Match("*.github.com")) // true

	// create new glob with set of delimiters as ["."]
	g = glob.MustCompile("api.*.com", '.')
	fmt.Println(g.Match("api.github.com")) // true
	fmt.Println(g.Match("api.gi.hub.com")) // false

	// create new glob with set of delimiters as ["."]
	// but now with super wildcard
	g = glob.MustCompile("api.**.com", '.')
	fmt.Println(g.Match("api.github.com")) // true
	fmt.Println(g.Match("api.gi.hub.com")) // true

	// create glob with single symbol wildcard
	g = glob.MustCompile("?at")
	fmt.Println(g.Match("cat")) // true
	fmt.Println(g.Match("fat")) // true
	fmt.Println(g.Match("at"))  // false

	// create glob with single symbol wildcard and delimiters ['f']
	g = glob.MustCompile("?at", 'f')
	fmt.Println(g.Match("cat")) // true
	fmt.Println(g.Match("fat")) // false
	fmt.Println(g.Match("at"))  // false

	// create glob with character-list matchers
	g = glob.MustCompile("[abc]at")
	fmt.Println(g.Match("cat")) // true
	fmt.Println(g.Match("bat")) // true
	fmt.Println(g.Match("fat")) // false
	fmt.Println(g.Match("at"))  // false

	// create glob with character-list matchers
	g = glob.MustCompile("[!abc]at")
	fmt.Println(g.Match("cat")) // false
	fmt.Println(g.Match("bat")) // false
	fmt.Println(g.Match("fat")) // true
	fmt.Println(g.Match("at"))  // false

	// create glob with character-range matchers
	g = glob.MustCompile("[a-c]at")
	fmt.Println(g.Match("cat")) // true
	fmt.Println(g.Match("bat")) // true
	fmt.Println(g.Match("fat")) // false
	fmt.Println(g.Match("at"))  // false

	// create glob with character-range matchers
	g = glob.MustCompile("[!a-c]at")
	fmt.Println(g.Match("cat")) // false
	fmt.Println(g.Match("bat")) // false
	fmt.Println(g.Match("fat")) // true
	fmt.Println(g.Match("at"))  // false

	// create glob with pattern-alternatives list
	g = glob.MustCompile("{cat,bat,[fr]at}")
	fmt.Println(g.Match("cat")) // true
	fmt.Println(g.Match("bat")) // true
	fmt.Println(g.Match("fat")) // true
	fmt.Println(g.Match("rat")) // true
	fmt.Println(g.Match("at"))  // false
	fmt.Println(g.Match("zat")) // false
}
