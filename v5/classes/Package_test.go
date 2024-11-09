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

package classes_test

import (
	mod "github.com/craterdog/go-class-model/v5"
	cla "github.com/craterdog/go-code-generation/v5/classes"
	uti "github.com/craterdog/go-missing-utilities/v2"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	sts "strings"
	tes "testing"
)

var directory = "../testdata/"
var moduleName = "github.com/craterdog/go-syntax-notation/v5"
var wikiPath = "github.com/craterdog/go-syntax-notation/wiki"

func TestAstGeneration(t *tes.T) {
	// Validate the class model.
	var packageName = "ast"
	var filename = directory + packageName + "/Package.go"
	var bytes, err = osx.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var model = mod.ParseSource(source)
	mod.ValidateModel(model)
	var actual = mod.FormatModel(model)
	ass.Equal(t, source, actual)
	var generator = cla.Generator().Make()

	// Generate the AST concrete classes.
	var interfaceDeclarations = model.GetInterfaceDeclarations()
	var classSection = interfaceDeclarations.GetClassSection()
	var classDeclarations = classSection.GetClassDeclarations().GetIterator()
	for classDeclarations.HasNext() {
		var classDeclaration = classDeclarations.GetNext()
		var className = classDeclaration.GetDeclaration().GetName()
		className = sts.TrimSuffix(className, "ClassLike")
		className = uti.MakeLowerCase(className)
		filename = directory + packageName + "/" + className + ".go"
		var astSynthesizer = cla.ClassSynthesizer().Make(model, className)
		source = generator.GenerateClass(
			moduleName,
			wikiPath,
			packageName,
			className,
			astSynthesizer,
		)
		bytes = []byte(source)
		err = osx.WriteFile(filename, bytes, 0644)
		if err != nil {
			panic(err)
		}
	}
}
