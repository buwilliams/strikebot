package adaptors

import (
    "fmt"
)

func Println(options map[string]interface{}) {
    message := options["message"].(string)
    fmt.Printf("[%p] %s\n", &message, message)
}
