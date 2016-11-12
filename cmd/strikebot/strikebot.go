package main

import (
    "os"
    "fmt"
    bot "github.com/buwilliams/strikebot/bot"
)

func main() {
    strikebotArgs := os.Args
    if len(strikebotArgs) == 2 {
        bot.Strike(strikebotArgs[1])
    } else {
        fmt.Println("Usage: strikebot <map.json>")
    }
}
