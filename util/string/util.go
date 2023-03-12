package string

import (
	"strings"
)

func SplitCookie(cookie string) map[string]string {
	res := make(map[string]string, 10)
	for _, item := range strings.Split(cookie, ";") {
		item = strings.TrimSpace(item)
		split := strings.Split(item, "=")
		if len(split) == 2 {
			res[split[0]] = split[1]
		}
	}
	return res
}
