package main

import (
	"fmt"
	"sync"
	"time"
)

var iconsOnce sync.Once
var icons map[string]string

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(Icon("hearts"), len(icons))
		}()
	}

	time.Sleep(time.Second)
}

func Icon(name string) string {
	iconsOnce.Do(loadIcons) // one-time initialization
	return icons[name]
}

func loadIcons() {
	icons = map[string]string{
		"spades":   "spades.png",
		"hearts":   "hearts.png",
		"diamonds": "diamonds.png",
		"clubs":    "clubs.png",
	}
}
