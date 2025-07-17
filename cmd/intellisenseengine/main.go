// cmd/intellisenseengine/main.go
package main

import (
"flag"
"log"
"os"

"intellisenseengine/internal/intellisenseengine"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := intellisenseengine.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
