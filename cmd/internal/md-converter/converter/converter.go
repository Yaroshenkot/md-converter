package converter

import "fmt"

func ConvertMDToHTML(mdContent string) string {

	var n bool
	html := Parsing(mdContent)

	fmt.Println("Нужно ли выполнить экранирование?")
	fmt.Scan(&n)

	if n {
		html = EscapeHTML(html)
	}

	return html

}
