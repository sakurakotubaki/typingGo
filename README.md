# typingGo
Go anytime lesson

try Go!

- [Task](https://taskfile.dev/)
- [LayerX Engineer Blog](https://tech.layerx.co.jp/entry/taskfile-dev)

## 正規表現のやり方

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("a[a-z]+e", "appl10e")
	fmt.Println(match)
}
```