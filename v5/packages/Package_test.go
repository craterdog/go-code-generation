/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package packages_test

import (
	pac "github.com/craterdog/go-code-generation/v5/packages"
	not "github.com/craterdog/go-syntax-notation/v5"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	tes "testing"
)

var directory = "../testdata/packages/"
var moduleName = "github.com/craterdog/go-syntax-notation/v5"
var wikiPath = "github.com/craterdog/go-syntax-notation/wiki"

func TestPackageGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var bytes, err = osx.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var syntax = not.ParseSource(source)
	not.ValidateSyntax(syntax)
	var actual = not.FormatSyntax(syntax)
	ass.Equal(t, source, actual)
	var generator = pac.Generator().Make()

	// Generate the AST Package.go file.
	filename = directory + "ast/Package.go"
	var packageName = "ast"
	var astSynthesizer = pac.AstSynthesizer().Make(syntax)
	source = generator.GeneratePackage(
		moduleName,
		wikiPath,
		packageName,
		astSynthesizer,
	)
	bytes = []byte(source)
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}

	// Generate the grammar Package.go file.
	filename = directory + "grammar/Package.go"
	packageName = "grammar"
	var grammarSynthesizer = pac.GrammarSynthesizer().Make(syntax)
	source = generator.GeneratePackage(
		moduleName,
		wikiPath,
		packageName,
		grammarSynthesizer,
	)
	bytes = []byte(source)
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}

	// Generate the example Package.go file.
	filename = directory + "example/Package.go"
	packageName = "example"
	var exampleSynthesizer = pac.ExampleSynthesizer().Make(syntax)
	source = generator.GeneratePackage(
		moduleName,
		wikiPath,
		packageName,
		exampleSynthesizer,
	)
	bytes = []byte(source)
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
