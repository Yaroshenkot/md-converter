package converter

import (
	"regexp"
	"strings"
)

func Parsing(md string) string {
	//таблица
	md = regexp.MustCompile(`(?m)((?:\|.*\|(?:\n|$))+)`).ReplaceAllStringFunc(md, func(match string) string {
		lines := strings.Split(strings.TrimSpace(match), "\n")
		if len(lines) < 3 {
			return match
		}

		var html strings.Builder
		html.WriteString("<table>\n")

		for i, line := range lines {
			line = strings.Trim(line, "|")
			cells := strings.Split(line, "|")

			if i == 0 {
				html.WriteString("  <thead>\n    <tr>\n")
				for _, cell := range cells {
					html.WriteString("      <th>" + strings.TrimSpace(cell) + "</th>\n")
				}
				html.WriteString("    </tr>\n  </thead>\n  <tbody>\n")
			} else if i == 1 {
				continue // Пропускаем разделительную линию
			} else {
				html.WriteString("    <tr>\n")
				for _, cell := range cells {
					html.WriteString("      <td>" + strings.TrimSpace(cell) + "</td>\n")
				}
				html.WriteString("    </tr>\n")
			}
		}

		html.WriteString("  </tbody>\n</table>")
		return html.String()
	})

	//параграфы
	md = regexp.MustCompile(`(?m)^(.*\[([^\]]+)\]\(([^)]+)\).*)$`).ReplaceAllString(md, "<p>$1</p>")
	// Заголовки
	md = regexp.MustCompile("```markdown").ReplaceAllString(md, "```html")
	md = regexp.MustCompile(`(?m)^### (.*)$`).ReplaceAllString(md, "<h3>$1</h3>")
	md = regexp.MustCompile(`(?m)^## (.*)$`).ReplaceAllString(md, "<h2>$1</h2>")
	md = regexp.MustCompile(`(?m)^# (.*)$`).ReplaceAllString(md, "<h1>$1</h1>")

	// Жирный текст
	md = regexp.MustCompile(`\*\*(.*?)\*\*`).ReplaceAllString(md, "<strong>$1</strong>")

	//горизонтальные линии
	md = regexp.MustCompile(`(?m)^---$`).ReplaceAllString(md, "<hr>")
	md = regexp.MustCompile(`(?m)^\*\*\*$`).ReplaceAllString(md, "<hr>")

	//цитата
	//md = regexp.MustCompile(`(?m)^> (.+)$`).ReplaceAllString(md, "<blockquote>$1</blockquote>")

	md = regexp.MustCompile(`(?m)^> (.+)$`).ReplaceAllString(md, "<blockquote>$1</blockquote>")
	// Объединяем последовательные blockquote
	md = regexp.MustCompile(`</blockquote>\s*<blockquote>`).ReplaceAllString(md, "<br>\n")

	//для ссылок
	md = regexp.MustCompile(`!\[(.*?)\]\((.*?)\)`).ReplaceAllString(md, `<img src="$2" alt="$1">`)
	md = regexp.MustCompile(`\[([^\[\]]+)\]\(([^()\s]+)(?:\s+"([^"]+)")?\)`).ReplaceAllString(md, `<a href="$2$3">$1</a>`)

	//изображения
	//md = regexp.MustCompile(`!\[(.*?)\]\((.*?)\)`).ReplaceAllString(md, `<img src="$2" alt="$1">`)

	//код

	md = regexp.MustCompile("(?s)```go(\\w+)?\\n(.*?)```").ReplaceAllString(md, "<pre><code class=\"language-go$1\">$2</code></pre>")

	//Нумерованный список

	md = regexp.MustCompile(`(?m)^\d+\. (.+)$`).ReplaceAllString(md, "<li>$1</li>")
	//md = regexp.MustCompile(`(?s)(<li>.*?</li>)`).ReplaceAllStringFunc(md, func(match string) string {
	//return "<ol>\n" + match + "\n</ol>"
	//})

	// Списки
	lines := strings.Split(md, "\n")
	inList := false
	for i, line := range lines {

		if strings.HasPrefix(line, "- ") {
			if !inList {
				lines[i] = "<ul>\n<li>" + line[2:] + "</li>"
				inList = true
			} else {
				lines[i] = "<li>" + line[2:] + "</li>"
			}
		} else if inList {
			lines[i-1] += "\n</ul>"
			inList = false
		}
	}
	return strings.Join(lines, "\n")

}
