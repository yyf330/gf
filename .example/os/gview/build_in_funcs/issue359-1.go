package main

import (
	"fmt"

	"github.com/yyf330/gf/frame/g"
)

func main() {
	tplContent := `
{{"我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人我是中国人"| strlimit 10  "..."}}
`
	content, err := g.View().ParseContent(tplContent, nil)
	fmt.Println(err)
	fmt.Println(content)
}
