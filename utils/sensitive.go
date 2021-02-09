package utils

import (
	"fmt"
	"github.com/importcjj/sensitive"
	"sync"
)

var sOnce sync.Once

var Filter = sensitive.New()

func init() {
	fmt.Print("====================init =========Filter===============")
	sOnce.Do(func() {
		Filter.AddWord("垃圾")
		Filter.LoadWordDict("data.txt")
	})
}
