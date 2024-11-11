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
		var classSynthesizer = syn.ClassSynthesizer().Make(model, className)
		source = generator.GenerateClass(
			moduleName,
			wikiPath,
			packageName,
			className,
			classSynthesizer,
		)
		bytes = []byte(source)
		err = osx.WriteFile(filename, bytes, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func TestTokenGeneration(t *tes.T) {
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

	// Generate the token concrete class.
	var generator = gen.ClassGenerator().Make()
	var packageName = "grammar"
	var className = "token"
	filename = directory + packageName + "/" + className + ".go"
	var tokenSynthesizer = syn.TokenSynthesizer().Make(syntax)
	source = generator.GenerateClass(
		moduleName,
		wikiPath,
		packageName,
		className,
		tokenSynthesizer,
	)
	bytes = []byte(source)
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func TestScannerGeneration(t *tes.T) {
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

	// Generate the scanner concrete class.
	var generator = gen.ClassGenerator().Make()
	var packageName = "grammar"
	var className = "scanner"
	filename = directory + packageName + "/" + className + ".go"
	var scannerSynthesizer = syn.ScannerSynthesizer().Make(syntax)
	source = generator.GenerateClass(
		moduleName,
		wikiPath,
		packageName,
		className,
		scannerSynthesizer,
	)
	bytes = []byte(source)
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func TestParserGeneration(t *tes.T) {
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

	// Generate the parser concrete class.
	var generator = gen.ClassGenerator().Make()
	var packageName = "grammar"
	var className = "parser"
	filename = directory + packageName + "/" + className + ".go"
	var parserSynthesizer = syn.ParserSynthesizer().Make(syntax)
	source = generator.GenerateClass(
		moduleName,
		wikiPath,
		packageName,
		className,
		parserSynthesizer,
	)
	bytes = []byte(source)
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func TestFormatterGeneration(t *tes.T) {
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

	// Generate the formatter concrete class.
	var generator = gen.ClassGenerator().Make()
	var packageName = "grammar"
	var className = "formatter"
	filename = directory + packageName + "/" + className + ".go"
	var formatterSynthesizer = syn.FormatterSynthesizer().Make(syntax)
	source = generator.GenerateClass(
		moduleName,
		wikiPath,
		packageName,
		className,
		formatterSynthesizer,
	)
	bytes = []byte(source)
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func TestProcessorGeneration(t *tes.T) {
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

	// Generate the processor concrete class.
	var generator = gen.ClassGenerator().Make()
	var packageName = "grammar"
	var className = "processor"
	filename = directory + packageName + "/" + className + ".go"
	var processorSynthesizer = syn.ProcessorSynthesizer().Make(syntax)
	source = generator.GenerateClass(
		moduleName,
		wikiPath,
		packageName,
		className,
		processorSynthesizer,
	)
	bytes = []byte(source)
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func TestValidatorGeneration(t *tes.T) {
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

	// Generate the validator concrete class.
	var generator = gen.ClassGenerator().Make()
	var packageName = "grammar"
	var className = "validator"
	filename = directory + packageName + "/" + className + ".go"
	var validatorSynthesizer = syn.ValidatorSynthesizer().Make(syntax)
	source = generator.GenerateClass(
		moduleName,
		wikiPath,
		packageName,
		className,
		validatorSynthesizer,
	)
	bytes = []byte(source)
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
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

/*
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
*/
