# go-typetalk-token-source

Implementation of [oauth2.TokenSource](https://godoc.org/golang.org/x/oauth2#TokenSource) in Typetalk.

## Requires

Go 1.14+

## Usage

```go
import "github.com/typetalk-gadget/go-typetalk-token-source/source"
```

## Example

```go
package main

import (
	"fmt"
	"os"

	"github.com/typetalk-gadget/go-typetalk-token-source/source"
)

func main() {
	ts := source.TokenSource{
		ClientID:     os.Getenv("TYPETALK_CLIENT_ID"),
		ClientSecret: os.Getenv("TYPETALK_CLIENT_SECRET"),
		Scope:        "my topic.read",
	}
	t, err := ts.Token()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("token: %#v\n", t)
}

```

## Bugs and Feedback

For bugs, questions and discussions please use the Github Issues.

## License

[MIT License](http://www.opensource.org/licenses/mit-license.php)

## Author

[vvatanabe](https://github.com/vvatanabe)