# HTTP reader with automatic content decompression 

Usage
---

Supported http content encodings: `gzip`, `deflate`, `br` (aka [Brotli](https://en.wikipedia.org/wiki/Brotli) format)

```go
package main

import (
	"github.com/oripio/respreader"
	"log"
	"net/http"
)

func main() {
	url := "https://google.com/"

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bresp, err := respreader.Decode(resp)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(bresp))
}

```

License
---

Brotli and these bindings are open-sourced under the MIT License - see the LICENSE file.
