## GO Emojify


Turn emoji into code
```go
Unemojify(text)
```

Turn code (ex: :beer:) into emoji
```go
Emojify(text)
```

### Example 

create main.go and paste code below

```go
package main

import (
	"log"
	emoji "github.com/kxchange/go-emojify"
)

const (
	contents = "I want to watch :tv: and drink :beer:"
)

func main() {

	log.Printf("Input: %s", contents)

	emojiMsg := emoji.Emojify(contents)
	log.Printf("Emojify: %s", emojiMsg)

	unemojiMsg := emoji.Unemojify(emojiMsg)
	log.Printf("Unemojify: %s", unemojiMsg)
}
```

run the code
```bash
go run main.go
```