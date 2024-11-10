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

package generator_test

import (
	mod "github.com/craterdog/go-class-model/v5"
	gen "github.com/craterdog/go-code-generation/v5/generator"
	syn "github.com/craterdog/go-code-generation/v5/synthesizer"
	uti "github.com/craterdog/go-missing-utilities/v2"
	not "github.com/craterdog/go-syntax-notation/v5"
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
	var generator = gen.ClassGenerator().Make()

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
		var astSynthesizer = syn.ClassSynthesizer().Make(model, className)
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

func TestExampleGeneration(t *tes.T) {
	// Validate the class model.
	var packageName = "example"
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
	var generator = gen.ClassGenerator().Make()

	// Generate the example concrete classes.
	var interfaceDeclarations = model.GetInterfaceDeclarations()
	var classSection = interfaceDeclarations.GetClassSection()
	var classDeclarations = classSection.GetClassDeclarations().GetIterator()
	for classDeclarations.HasNext() {
		var classDeclaration = classDeclarations.GetNext()
		var className = classDeclaration.GetDeclaration().GetName()
		className = sts.TrimSuffix(className, "ClassLike")
		className = uti.MakeLowerCase(className)
		filename = directory + packageName + "/" + className + ".go"
		var exampleSynthesizer = syn.ClassSynthesizer().Make(model, className)
		source = generator.GenerateClass(
			moduleName,
			wikiPath,
			packageName,
			className,
			exampleSynthesizer,
		)
		bytes = []byte(source)
		err = osx.WriteFile(filename, bytes, 0644)
		if err != nil {
			panic(err)
		}
	}
}

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
	var generator = gen.PackageGenerator().Make()

	// Generate the AST Package.go file.
	filename = directory + "ast/Package.go"
	var packageName = "ast"
	var astSynthesizer = syn.AstSynthesizer().Make(syntax)
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
	var grammarSynthesizer = syn.GrammarSynthesizer().Make(syntax)
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
	var exampleSynthesizer = syn.ExampleSynthesizer().Make()
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
