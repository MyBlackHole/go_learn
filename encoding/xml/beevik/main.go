package main

import (
	"fmt"
	"github.com/beevik/etree"
)

func main() {
	// L02()
    parse_oss_xml()
}

func parse_oss_xml() {
    xml := etree.NewDocument()
    if err := xml.ReadFromFile("./oss_list.xml"); err != nil {
        panic(err)
    }

    root := xml.Root()
    fmt.Println(root.Attr)

	// root := doc.SelectElement("Contents")
	childElements := root.ChildElements()
    fmt.Println(len(childElements))
}

func test() {

	doc := etree.NewDocument()
	if err := doc.ReadFromFile("bookstore.xml"); err != nil {
		panic(err)
	}

	root := doc.SelectElement("bookstore")
	fmt.Println("ROOT element:", root.Tag)

	for _, book := range root.SelectElements("book") {
		fmt.Println("CHILD element:", book.Tag)
		if title := book.SelectElement("title"); title != nil {
			lang := title.SelectAttrValue("lang", "unknown")
			fmt.Printf("  TITLE: %s (%s)\n", title.Text(), lang)
		}
		for _, attr := range book.Attr {
			fmt.Printf("  ATTR: %s=%s\n", attr.Key, attr.Value)
		}
	}

	Parse(doc.Root())
}

func Parse(root *etree.Element) {
	childElements := root.ChildElements()
	if len(childElements) <= 0 {
		return
	}

	for _, child := range childElements {
		fmt.Printf("PARSE: %s=%s\n", child.Tag, child.Text())
		// for _, attr := range child.Attr {
		// 	fmt.Printf("PARSEATTR: %s=%s\n", attr.Key, attr.Value)
		// }
		Parse(child)
	}
}

func L02() {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("./L02_20220614-20220628.xml"); err != nil {
		panic(err)
	}

	Parse(doc.Root())
}
