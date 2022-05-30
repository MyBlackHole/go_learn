package main
 
import (
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
)
 
func main() {
	pdfGenerator, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}
	// Set global options
	pdfGenerator.Orientation.Set(wkhtmltopdf.OrientationLandscape)//设置为“风景(Landscape)”或“肖像(Portrait)”模式, 默认是肖像模块(Portrait)
	pdfGenerator.Grayscale.Set(false)   //指定是否以灰度图生成PDF文档。占用的空间更小
 
	url := "https://www.baidu.com/"
	pdfGenerator.AddPage(wkhtmltopdf.NewPage(url))
	err = pdfGenerator.Create()
	if err  != nil {
		log.Fatal(err)
	}
	err = pdfGenerator.WriteFile("./crawler.pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done!")
}
