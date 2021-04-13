package url

import (
	"fmt"
	"strings"
)

const fieldColor = "\033[38;5;184m"
const resetColor = "\033[0m"

func (urlObj *URLObject) String() string {
	str := strings.Builder{}

	str.WriteString(resetColor)
	str.WriteString(fieldColor)
	str.WriteString("Short: ")
	str.WriteString(resetColor)
	str.WriteString(urlObj.Short)
	str.WriteString("\n")

	str.WriteString(fieldColor)
	str.WriteString("Long : ")
	str.WriteString(resetColor)
	str.WriteString(urlObj.Long)
	str.WriteString("\n")

	str.WriteString(fieldColor)
	str.WriteString("CreatedAt: ")
	str.WriteString(resetColor)
	str.WriteString(fmt.Sprint(urlObj.CreatedAt))
	str.WriteString(" , ")
	str.WriteString(fieldColor)
	str.WriteString("ExpireAt: ")
	str.WriteString(resetColor)
	str.WriteString(fmt.Sprint(urlObj.ExpireAt))
	str.WriteString("\n")

	str.WriteString(fieldColor)
	str.WriteString("KeepForever: ")
	str.WriteString(resetColor)
	str.WriteString(fmt.Sprint(urlObj.KeepForever))

	return str.String()
}
