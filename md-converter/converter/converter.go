package converter

func ConvertMDToHTML(mdContent string) (string, error) {
	html := Parsing(mdContent)
	escape := EscapeHTML(html)

	return escape, nil

}
