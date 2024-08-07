package main

import (
	gg "github.com/PlayerR9/go-generator/generator"
	pkg "github.com/PlayerR9/iterators/cmd/simple/pkg"
)

func main() {
	gg.ParseFlags()

	res, err := pkg.Generator.Generate(pkg.OutputLocFlag, "<type_name>_iterator.go", &pkg.GenData{})
	if err != nil {
		pkg.Logger.Fatalf("Failed to generate code: %s", err.Error())
	}

	dest, err := res.WriteFile("")
	if err != nil {
		pkg.Logger.Fatalf("Failed to write code: %s", err.Error())
	}

	pkg.Logger.Printf("Successfully generated code: %s", dest)
}
