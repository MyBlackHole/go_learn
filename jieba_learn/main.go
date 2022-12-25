package main

import (

  "fmt"
  "github.com/wangbin/jiebago"
)

var seg jiebago.Segmenter

func init() {
  seg.LoadDictionary("dict.txt")
}

func print(ch <-chan string) {
  for word := range ch {
    fmt.Printf(" %s /", word)
  }
  fmt.Println()
}

func Example() {
  fmt.Print("【全模式】：")
  print(seg.CutAll("我来到北京清华大学"))

  fmt.Print("【精确模式】：")
  print(seg.Cut("我来到北京清华大学", false))

  fmt.Print("【新词识别】：")
  print(seg.Cut("一、本地疫情2月25日0—24时，重庆市本地无新增新冠肺炎确诊病例报告。截至2月25日24时，重庆市本地无在院确诊病例，累计治愈出院病例570例，累计<em>死亡</em>病例6例，累计报告确诊病例576例。二、境外输入疫情2月25日0—24时，重庆市无新增境外输入新冠肺炎确诊病例报告。截至2月25日24时，重庆市无境外输入在院确诊病例，累计报告境外输入确诊病例15例，均已治愈出院。三、无症状感染者2月25日0—24时，重庆市无新增无症状感染者报告；解除医学观察1例（为尼泊尔输入）。截至2月25日24时，尚有无症状感染者2例（均为尼泊尔输入），正在接受医学观察。四、密切接触者截至2月25日24时，累计追踪到密切接触者26620人，均已解除医学观察。【来源：重庆市卫生健康委员会】版权归原作者所有，向原创致敬", true))

  fmt.Print("【搜索引擎模式】：")
  print(seg.CutForSearch("小明硕士毕业于中国科学院计算所，后在日本京都大学深造", true))
}

func main() {
  Example()
}
