package converter

import (
	"fmt"
	"strings"
)

func EscapeHTML(md string) string {
	md = strings.ReplaceAll(md, "&", "&amp;")
	md = strings.ReplaceAll(md, "<", "&lt;")
	md = strings.ReplaceAll(md, ">", "&gt;")
	md = strings.ReplaceAll(md, "\"", "&quot;")
	md = strings.ReplaceAll(md, "`", "&#39;")

	fmt.Println("ggggg")
	return md

}
