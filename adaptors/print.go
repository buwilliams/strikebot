package adaptors

import (
    "fmt"
)

func Println(options map[string]interface{}) {
    message := options["message"].(string)
    fmt.Println(message)
}
