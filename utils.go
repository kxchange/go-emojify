package emojify

import (
	"strconv"
	"strings"
)

// see 
// https://stackoverflow.com/questions/55700149/print-emoji-from-unicode-literal-loaded-from-file
func unquoteCodePoint(s string) (string, error) {
    r, err := strconv.ParseInt(strings.TrimPrefix(s, "\\U"), 16, 32)
    return string(r), err
}