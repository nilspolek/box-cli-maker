package main

import (
	"fmt"
	"time"

	"github.com/Delta456/box-cli-maker/v2"
)

func main() {
	Box := box.New(box.Config{Px: 2, Py: 1, Type: "Single", Color: "Cyan", TitlePos: "Top"})
	for i:=1; i <= 10; i++  {
		Box.PrintAlingned("Counter", fmt.Sprintf("Counter: %d", i))
		time.Sleep(time.Second * 1)
	}
}
