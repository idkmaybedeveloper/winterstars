# winterstars

simple snowflake implem

example:
```go
package main

import (
	"log/slog"

	winterstars "go.notyandex.cloud/winterstars"
)

func main() {
	var (
		meow   = winterstars.Next()      // get snowflake id itself
		sessid = winterstars.SessionID() // get sessionId
	)

	slog.Info("snowflake id", "id", meow)
	slog.Info("sess id", "id", sessid)
}
```