package emojify

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
	"github.com/stretchr/testify/assert"
)

var (
	msg = "saya suka :beer: :beer: :couplekiss_woman_woman:"
)

func TestEmojiRevCodeMap(t *testing.T) {
	emojiMsg := Emojify(":beer:")
	r, _ := utf8.DecodeRuneInString(emojiMsg)
	hexKey, _ := unquoteCodePoint(fmt.Sprintf(`%X`, r))

	val, _ := emojiRevCodeMap[string(hexKey)]
	assert.Equal(t,":beer:", val[0], "Not valid emoji code")
}

func TestEmojify(t *testing.T) {
	emojiMsg := Emojify(msg)
	assert.NotContains(t, ":beer:", emojiMsg, "Contents should not contains emoji code")
}

func TestEmojifyFail(t *testing.T) {
	emojiMsg := Emojify(":cheek:")
	assert.Equal(t, ":cheek:", emojiMsg, "Contents output should :cheek:")
}

func TestUnEmojify(t *testing.T) {
	emojiMsg := Emojify(msg)
	unEmojiMsg := Unemojify(emojiMsg)
	assert.Contains(t, unEmojiMsg, ":beer:", "Contents should contains emoji code")
}

func TestUnemojiFail(t *testing.T) {
	emojiMsg := Unemoji(":cheek:")
	assert.Equal(t, ":cheek:", emojiMsg, "Contents output should :cheek:")
}

func TestEmojifyFunc(t *testing.T) {
	emojiMsg := EmojifyFunc(msg, func(name string)string{
		name = strings.ReplaceAll(name,":","")
		return fmt.Sprintf(`<span class="emoji emoji-%s"></span>`, name)
	})

	assert.Contains(t, emojiMsg, "emoji-beer", "Should contains class emoji-beer")
}


func BenchmarkEmojiRevCodeMap(t *testing.B) {
	emojiMsg := Emojify(":beer:")
	r, _ := utf8.DecodeRuneInString(emojiMsg)
	hexKey, _ := unquoteCodePoint(fmt.Sprintf(`%X`, r))
	val, _ := emojiRevCodeMap[string(hexKey)]
	fmt.Sprintf("%v", val)
}

func BenchmarkEmojify(t *testing.B) {
	Emojify(msg)
}

func BenchmarkUnemojify(t *testing.B) {
	emojiMsg := Emojify(msg)
	Unemojify(emojiMsg)
}

func BenchmarkEmojifyFunc(t *testing.B) {
	EmojifyFunc(msg, func(name string)string{
		name = strings.ReplaceAll(name,":","")
		return fmt.Sprintf(`<span class="emoji emoji-%s"></span>`, name)
	})
}