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
	mod "github.com/craterdog/go-class-model/v7"
	gen "github.com/craterdog/go-code-generation/v7"
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	not "github.com/craterdog/go-syntax-notation/v7"
	sts "strings"
	tes "testing"
)

var directory = "./testdata/"
var moduleName = "github.com/craterdog/go-syntax-notation/v7"
var wikiPath = "https://github.com/craterdog/go-syntax-notation/wiki"

func TestPackageGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)

	// Generate the AST package_api.go file.
	var model = gen.GenerateAstPackage(
		moduleName,
		wikiPath,
		syntax,
	)
	mod.ValidateModel(model)
	source = mod.FormatModel(model)
	var packageName = "ast"
	uti.RemakeDirectory(directory + packageName)
	filename = directory + packageName + "/package_api.go"
	uti.WriteFile(filename, source)

	// Generate the grammar package_api.go file.
	model = gen.GenerateGrammarPackage(
		moduleName,
		wikiPath,
		syntax,
	)
	mod.ValidateModel(model)
	source = mod.FormatModel(model)
	packageName = "grammar"
	filename = directory + packageName + "/package_api.go"
	uti.WriteFile(filename, source)
}

func TestModuleGeneration(t *tes.T) {
	var models = fra.Catalog[string, mod.ModelLike]()

	// Read in the AST package file.
	var packageName = "ast"
	var filename = directory + packageName + "/package_api.go"
	var source = uti.ReadFile(filename)
	var model = mod.ParseSource(source)
	models.SetValue(packageName, model)

	// Read in the grammar package file.
	packageName = "grammar"
	filename = directory + packageName + "/package_api.go"
	source = uti.ReadFile(filename)
	model = mod.ParseSource(source)
	models.SetValue(packageName, model)

	// Regenerate the module file.
	filename = directory + "module_api.go"
	var existing = uti.ReadFile(filename)
	source = gen.GenerateModule(
		moduleName,
		wikiPath,
		existing,
		models,
	)
	uti.WriteFile(filename, source)
}

func TestAstGeneration(t *tes.T) {
	// Validate the class model.
	var packageName = "ast"
	var filename = directory + packageName + "/package_api.go"
	var source = uti.ReadFile(filename)
	var model = mod.ParseSource(source)
	mod.ValidateModel(model)

	// Generate the AST concrete classes.
	var interfaceDeclarations = model.GetInterfaceDeclarations()
	var classSection = interfaceDeclarations.GetClassSection()
	var classDeclarations = classSection.GetClassDeclarations().GetIterator()
	for classDeclarations.HasNext() {
		var classDeclaration = classDeclarations.GetNext()
		var className = classDeclaration.GetDeclaration().GetName()
		className = sts.TrimSuffix(className, "ClassLike")
		className = uti.MakeLowerCase(className)
		source = gen.GenerateAstClass(
			moduleName,
			className,
			model,
		)
		className = uti.MakeUpperCase(className)
		filename = directory + packageName + "/" + className + ".go"
		uti.WriteFile(filename, source)
	}
}

func TestTokenGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)

	// Generate the token concrete class.
	source = gen.GenerateTokenClass(
		moduleName,
		syntax,
	)
	var packageName = "grammar"
	var className = "Token"
	filename = directory + packageName + "/" + className + ".go"
	uti.WriteFile(filename, source)
}

func TestScannerGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)

	// Generate the scanner concrete class.
	source = gen.GenerateScannerClass(
		moduleName,
		syntax,
	)
	var packageName = "grammar"
	var className = "Scanner"
	filename = directory + packageName + "/" + className + ".go"
	uti.WriteFile(filename, source)
}

func TestParserGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)

	// Generate the parser concrete class.
	source = gen.GenerateParserClass(
		moduleName,
		syntax,
	)
	var packageName = "grammar"
	var className = "Parser"
	filename = directory + packageName + "/" + className + ".go"
	uti.WriteFile(filename, source)
}

func TestVisitorGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)

	// Generate the visitor concrete class.
	source = gen.GenerateVisitorClass(
		moduleName,
		syntax,
	)
	var packageName = "grammar"
	var className = "Visitor"
	filename = directory + packageName + "/" + className + ".go"
	uti.WriteFile(filename, source)
}

func TestFormatterGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)

	// Generate the formatter concrete class.
	var packageName = "grammar"
	var className = "Formatter"
	filename = directory + packageName + "/" + className + ".go"
	var existing = uti.ReadFile(filename)
	source = gen.GenerateFormatterClass(
		moduleName,
		existing,
		syntax,
	)
	uti.WriteFile(filename, source)
}

func TestProcessorGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)

	// Generate the processor concrete class.
	source = gen.GenerateProcessorClass(
		moduleName,
		syntax,
	)
	var packageName = "grammar"
	var className = "Processor"
	filename = directory + packageName + "/" + className + ".go"
	uti.WriteFile(filename, source)
}

func TestValidatorGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)

	// Generate the validator concrete class.
	var packageName = "grammar"
	var className = "Validator"
	filename = directory + packageName + "/" + className + ".go"
	var existing = uti.ReadFile(filename)
	source = gen.GenerateValidatorClass(
		moduleName,
		existing,
		syntax,
	)
	uti.WriteFile(filename, source)
}

func TestExampleGeneration(t *tes.T) {
	// Generate the example class model.
	var packageName = "example"
	var wikiPath = "https://github.com/craterdog/go-package-example/wiki"
	uti.RemakeDirectory(directory + packageName)
	var model = gen.CreatePackage(
		moduleName,
		wikiPath,
		packageName,
	)
	mod.ValidateModel(model)
	var source = mod.FormatModel(model)
	var filename = directory + packageName + "/package_api.go"
	uti.WriteFile(filename, source)

	// Generate the example concrete classes.
	var interfaceDeclarations = model.GetInterfaceDeclarations()
	var classSection = interfaceDeclarations.GetClassSection()
	var classDeclarations = classSection.GetClassDeclarations().GetIterator()
	for classDeclarations.HasNext() {
		var existing string // There is no existing class file.
		var classDeclaration = classDeclarations.GetNext()
		var className = classDeclaration.GetDeclaration().GetName()
		className = sts.TrimSuffix(className, "ClassLike")
		className = uti.MakeLowerCase(className)
		source = gen.GenerateClass(
			moduleName,
			packageName,
			className,
			existing,
			model,
		)
		className = uti.MakeUpperCase(className)
		filename = directory + packageName + "/" + className + ".go"
		uti.WriteFile(filename, source)
	}
}
