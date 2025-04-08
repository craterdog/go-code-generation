/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package module_test

import (
	fmt "fmt"
	mod "github.com/craterdog/go-class-model/v5"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	tes "testing"
)

var modelFiles = []string{
	"./ast/Package.go",
	"./grammar/Package.go",
	"./example/Package.go",
}

func TestRoundTrips(t *tes.T) {
	fmt.Println("Round Trip Tests:")
	for _, modelFile := range modelFiles {
		fmt.Printf("   %v\n", modelFile)
		var bytes, err = osx.ReadFile(modelFile)
		if err != nil {
			panic(err)
		}
		var source = string(bytes)
		var parser = mod.Parser()
		var model = parser.ParseSource(source)
		var validator = mod.Validator()
		validator.ValidateModel(model)
		var formatter = mod.Formatter()
		var actual = formatter.FormatModel(model)
		ass.Equal(t, source, actual)
	}
	fmt.Println("Done.")
}
