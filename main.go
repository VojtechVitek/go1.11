package main

import (
	"fmt"

	"github.com/VojtechVitek/go1.11/addons"
)

func main() {
	_ = addons.TestAddon{}
	_ = addons.CustomJS{}

	fmt.Println("OK")
}
