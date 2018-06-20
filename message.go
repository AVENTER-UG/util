package util

import "github.com/russross/blackfriday"

// a wrapper fo printValue to use strings and not string pointers
func printValueStr(message string, length int) string {
	return printValue(&message, length)
}

// this function will add spaces to a string, until the length of the string is like we need it
// thats useful to make the output more pretty
func printValue(message *string, length int) string {
	if message != nil {
		if len(*message) < length {
			*message = *message + " "
			return printValue(message, length)
		}
	} else {
		newMsg := " "
		return printValue(&newMsg, length)
	}
	return *message
}

func markdownRender(content string) string {
	htmlFlags := 0
	htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_FRACTIONS

	renderer := blackfriday.HtmlRenderer(htmlFlags, "", "")

	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS

	return string(blackfriday.Markdown([]byte(content), renderer, extensions))
}
