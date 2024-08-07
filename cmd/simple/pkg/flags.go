package pkg

import (
	gg "github.com/PlayerR9/go-generator/generator"
)

var (
	OutputLocFlag *gg.OutputLocVal
)

func init() {
	OutputLocFlag = gg.NewOutputFlag("<type_name>_iterator.go", true)
}
