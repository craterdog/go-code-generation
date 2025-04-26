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
	not "github.com/craterdog/go-syntax-notation/v6"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	tes "testing"
)

var syntaxFiles = []string{
	"./Syntax.cdsn",
}

func TestRoundTrips(t *tes.T) {
	fmt.Println("Round Trip Tests:")
	for _, syntaxFile := range syntaxFiles {
		fmt.Printf("   %v\n", syntaxFile)
		var bytes, err = osx.ReadFile(syntaxFile)
		if err != nil {
			panic(err)
		}
		var source = string(bytes)
		var syntax = not.ParseSource(source)
		not.ValidateSyntax(syntax)
		var actual = not.FormatSyntax(syntax)
		ass.Equal(t, source, actual)
	}
	fmt.Println("Done.")
}
