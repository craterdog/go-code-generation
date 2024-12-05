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

package module_test

import (
	mod "github.com/craterdog/go-class-model/v5"
	gen "github.com/craterdog/go-code-generation/v5"
	col "github.com/craterdog/go-collection-framework/v4"
	uti "github.com/craterdog/go-missing-utilities/v2"
	not "github.com/craterdog/go-syntax-notation/v5"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	sts "strings"
	tes "testing"
)

var directory = "./testdata/"
var moduleName = "github.com/craterdog/go-class-model/v5"
var wikiPath = "https://github.com/craterdog/go-class-model/wiki"

func TestPackageGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = readFile(filename)
	var syntax = not.ParseSource(source)
	not.ValidateSyntax(syntax)
	var actual = not.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the AST Package.go file.
	var packageName = "ast"
	remakeDirectory(directory + packageName)
	var generator = gen.PackageGenerator()
	var astSynthesizer = gen.AstSynthesizer(syntax)
	source = generator.GeneratePackage(
		moduleName,
		wikiPath,
		packageName,
		astSynthesizer,
	)
	filename = directory + packageName + "/Package.go"
	writeFile(filename, source)

	// Generate the grammar Package.go file.
	packageName = "grammar"
	remakeDirectory(directory + packageName)
	var grammarSynthesizer = gen.GrammarSynthesizer(syntax)
	source = generator.GeneratePackage(
		moduleName,
		wikiPath,
		packageName,
		grammarSynthesizer,
	)
	filename = directory + packageName + "/Package.go"
	writeFile(filename, source)
}

func TestModuleGeneration(t *tes.T) {
	var models = col.Catalog[string, mod.ModelLike]()
	var packages = []string{
		"ast",
		"grammar",
	}
	for _, packageName := range packages {
		var filename = directory + packageName + "/Package.go"
		var source = readFile(filename)
		var model = mod.ParseSource(source)
		models.SetValue(packageName, model)
	}
	var generator = gen.ModuleGenerator()
	var moduleSynthesizer = gen.ModuleSynthesizer(models)
	var source = generator.GenerateModule(
		moduleName,
		wikiPath,
		moduleSynthesizer,
	)
	var filename = directory + "Module.go"
	writeFile(filename, source)
}

func TestAstGeneration(t *tes.T) {
	// Validate the class model.
	var packageName = "ast"
	var filename = directory + packageName + "/Package.go"
	var source = readFile(filename)
	var model = mod.ParseSource(source)
	mod.ValidateModel(model)
	var actual = mod.FormatModel(model)
	ass.Equal(t, source, actual)
	var generator = gen.ClassGenerator()

	// Generate the AST concrete classes.
	var interfaceDeclarations = model.GetInterfaceDeclarations()
	var classSection = interfaceDeclarations.GetClassSection()
	var classDeclarations = classSection.GetClassDeclarations().GetIterator()
	for classDeclarations.HasNext() {
		var classDeclaration = classDeclarations.GetNext()
		var className = classDeclaration.GetDeclaration().GetName()
		className = sts.TrimSuffix(className, "ClassLike")
		className = uti.MakeLowerCase(className)
		var nodeSynthesizer = gen.NodeSynthesizer(model, className)
		source = generator.GenerateClass(
			moduleName,
			packageName,
			className,
			nodeSynthesizer,
		)
		filename = directory + packageName + "/" + className + ".go"
		writeFile(filename, source)
	}
}

func TestTokenGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = readFile(filename)
	var syntax = not.ParseSource(source)
	not.ValidateSyntax(syntax)
	var actual = not.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the token concrete class.
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "token"
	var tokenSynthesizer = gen.TokenSynthesizer(syntax)
	source = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		tokenSynthesizer,
	)
	filename = directory + packageName + "/" + className + ".go"
	writeFile(filename, source)
}

func TestScannerGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = readFile(filename)
	var syntax = not.ParseSource(source)
	not.ValidateSyntax(syntax)
	var actual = not.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the scanner concrete class.
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "scanner"
	var scannerSynthesizer = gen.ScannerSynthesizer(syntax)
	source = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		scannerSynthesizer,
	)
	filename = directory + packageName + "/" + className + ".go"
	writeFile(filename, source)
}

func TestParserGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = readFile(filename)
	var syntax = not.ParseSource(source)
	not.ValidateSyntax(syntax)
	var actual = not.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the parser concrete class.
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "parser"
	var parserSynthesizer = gen.ParserSynthesizer(syntax)
	source = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		parserSynthesizer,
	)
	filename = directory + packageName + "/" + className + ".go"
	writeFile(filename, source)
}

func TestVisitorGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = readFile(filename)
	var syntax = not.ParseSource(source)
	not.ValidateSyntax(syntax)
	var actual = not.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the visitor concrete class.
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "visitor"
	var visitorSynthesizer = gen.VisitorSynthesizer(syntax)
	source = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		visitorSynthesizer,
	)
	filename = directory + packageName + "/" + className + ".go"
	writeFile(filename, source)
}

func TestFormatterGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = readFile(filename)
	var syntax = not.ParseSource(source)
	not.ValidateSyntax(syntax)
	var actual = not.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the formatter concrete class.
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "formatter"
	var formatterSynthesizer = gen.FormatterSynthesizer(syntax)
	source = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		formatterSynthesizer,
	)
	filename = directory + packageName + "/" + className + ".go"
	writeFile(filename, source)
}

func TestProcessorGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = readFile(filename)
	var syntax = not.ParseSource(source)
	not.ValidateSyntax(syntax)
	var actual = not.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the processor concrete class.
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "processor"
	var processorSynthesizer = gen.ProcessorSynthesizer(syntax)
	source = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		processorSynthesizer,
	)
	filename = directory + packageName + "/" + className + ".go"
	writeFile(filename, source)
}

func TestValidatorGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = readFile(filename)
	var syntax = not.ParseSource(source)
	not.ValidateSyntax(syntax)
	var actual = not.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the validator concrete class.
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "validator"
	var validatorSynthesizer = gen.ValidatorSynthesizer(syntax)
	source = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		validatorSynthesizer,
	)
	filename = directory + packageName + "/" + className + ".go"
	writeFile(filename, source)
}

func TestExampleGeneration(t *tes.T) {
	// Generate the example class model.
	var packageName = "example"
	var wikiPath = "https://github.com/craterdog/go-package-example/wiki"
	remakeDirectory(directory + packageName)
	var packageGenerator = gen.PackageGenerator()
	var exampleSynthesizer = gen.ExampleSynthesizer()
	var source = packageGenerator.GeneratePackage(
		moduleName,
		wikiPath,
		packageName,
		exampleSynthesizer,
	)
	var filename = directory + packageName + "/Package.go"
	writeFile(filename, source)

	// Validate the class model.
	var model = mod.ParseSource(source)
	mod.ValidateModel(model)
	var actual = mod.FormatModel(model)
	ass.Equal(t, source, actual)

	// Generate the example concrete classes.
	var classGenerator = gen.ClassGenerator()
	var interfaceDeclarations = model.GetInterfaceDeclarations()
	var classSection = interfaceDeclarations.GetClassSection()
	var classDeclarations = classSection.GetClassDeclarations().GetIterator()
	for classDeclarations.HasNext() {
		var classDeclaration = classDeclarations.GetNext()
		var className = classDeclaration.GetDeclaration().GetName()
		className = sts.TrimSuffix(className, "ClassLike")
		className = uti.MakeLowerCase(className)
		var classSynthesizer = gen.ClassSynthesizer(model, className)
		source = classGenerator.GenerateClass(
			moduleName,
			packageName,
			className,
			classSynthesizer,
		)
		filename = directory + packageName + "/" + className + ".go"
		writeFile(filename, source)
	}
}

func readFile(
	filename string,
) string {
	var bytes, err = osx.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	return source
}

func remakeDirectory(
	directory string,
) {
	var err = osx.RemoveAll(directory)
	if err != nil {
		panic(err)
	}
	err = osx.MkdirAll(directory, 0755)
	if err != nil {
		panic(err)
	}
}

func writeFile(
	filename string,
	source string,
) {
	var bytes = []byte(source)
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
