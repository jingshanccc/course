package util

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

//TransToHtml: 将 md 格式字符串转为 html 格式
func TransToHtml(md string) string {
	unsafe := blackfriday.Run([]byte(md))
	return bluemonday.UGCPolicy().Sanitize(string(unsafe))
}
