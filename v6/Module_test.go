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
	mod "github.com/craterdog/go-class-model/v5"
	gen "github.com/craterdog/go-code-generation/v6"
	fra "github.com/craterdog/go-collection-framework/v5"
	uti "github.com/craterdog/go-missing-utilities/v2"
	not "github.com/craterdog/go-syntax-notation/v6"
	ass "github.com/stretchr/testify/assert"
	sts "strings"
	tes "testing"
)

var directory = "./testdata/"
var moduleName = "github.com/craterdog/go-syntax-notation/v6"
var wikiPath = "https://github.com/craterdog/go-syntax-notation/wiki"

func TestPackageGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)
	var actual = gen.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the AST Package.go file.
	var generated = gen.GenerateAstPackage(
		moduleName,
		wikiPath,
		syntax,
	)
	source = gen.FormatModel(generated)
	var packageName = "ast"
	uti.RemakeDirectory(directory + packageName)
	filename = directory + packageName + "/Package.go"
	uti.WriteFile(filename, source)

	// Generate the grammar Package.go file.
	generated = gen.GenerateGrammarPackage(
		moduleName,
		wikiPath,
		syntax,
	)
	source = gen.FormatModel(generated)
	packageName = "grammar"
	filename = directory + packageName + "/Package.go"
	uti.WriteFile(filename, source)
}

func TestModuleGeneration(t *tes.T) {
	var models = fra.Catalog[string, mod.ModelLike]()

	// Read in the AST package file.
	var packageName = "ast"
	var filename = directory + packageName + "/Package.go"
	var source = uti.ReadFile(filename)
	var model = mod.ParseSource(source)
	models.SetValue(packageName, model)

	// Read in the grammar package file.
	packageName = "grammar"
	filename = directory + packageName + "/Package.go"
	source = uti.ReadFile(filename)
	model = mod.ParseSource(source)
	models.SetValue(packageName, model)

	// Regenerate the module file.
	filename = directory + "Module.go"
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
	var filename = directory + packageName + "/Package.go"
	var source = uti.ReadFile(filename)
	var model = mod.ParseSource(source)
	var actual = gen.FormatModel(model)
	ass.Equal(t, source, actual)

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
		filename = directory + packageName + "/" + className + ".go"
		uti.WriteFile(filename, source)
	}
}

func TestTokenGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)
	var actual = gen.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the token concrete class.
	source = gen.GenerateTokenClass(
		moduleName,
		syntax,
	)
	var packageName = "grammar"
	var className = "token"
	filename = directory + packageName + "/" + className + ".go"
	uti.WriteFile(filename, source)
}

func TestScannerGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)
	var actual = gen.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the scanner concrete class.
	source = gen.GenerateScannerClass(
		moduleName,
		syntax,
	)
	var packageName = "grammar"
	var className = "scanner"
	filename = directory + packageName + "/" + className + ".go"
	uti.WriteFile(filename, source)
}

func TestParserGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)
	var actual = gen.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the parser concrete class.
	source = gen.GenerateParserClass(
		moduleName,
		syntax,
	)
	var packageName = "grammar"
	var className = "parser"
	filename = directory + packageName + "/" + className + ".go"
	uti.WriteFile(filename, source)
}

func TestVisitorGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)
	var actual = gen.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the visitor concrete class.
	source = gen.GenerateVisitorClass(
		moduleName,
		syntax,
	)
	var packageName = "grammar"
	var className = "visitor"
	filename = directory + packageName + "/" + className + ".go"
	uti.WriteFile(filename, source)
}

func TestFormatterGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)
	var actual = gen.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the formatter concrete class.
	var packageName = "grammar"
	var className = "formatter"
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
	var filename = directory + "Syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)
	var actual = gen.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the processor concrete class.
	source = gen.GenerateProcessorClass(
		moduleName,
		syntax,
	)
	var packageName = "grammar"
	var className = "processor"
	filename = directory + packageName + "/" + className + ".go"
	uti.WriteFile(filename, source)
}

func TestValidatorGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)
	var actual = gen.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the validator concrete class.
	var packageName = "grammar"
	var className = "validator"
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
	var packageAssembler = gen.PackageAssembler()
	var packageSynthesizer = gen.PackageSynthesizer()
	var generated = packageAssembler.AssemblePackage(
		moduleName,
		wikiPath,
		packageName,
		"", // There is no existing example package file.
		packageSynthesizer,
	)
	var filename = directory + packageName + "/Package.go"
	uti.WriteFile(filename, generated)

	// Validate the class model.
	var model = mod.ParseSource(generated)
	mod.ValidateModel(model)
	var actual = mod.FormatModel(model)
	ass.Equal(t, generated, actual)

	// Generate the example concrete classes.
	var classAssembler = gen.ClassAssembler()
	var interfaceDeclarations = model.GetInterfaceDeclarations()
	var classSection = interfaceDeclarations.GetClassSection()
	var classDeclarations = classSection.GetClassDeclarations().GetIterator()
	for classDeclarations.HasNext() {
		var classDeclaration = classDeclarations.GetNext()
		var className = classDeclaration.GetDeclaration().GetName()
		className = sts.TrimSuffix(className, "ClassLike")
		className = uti.MakeLowerCase(className)
		var classSynthesizer = gen.ClassSynthesizer(model, className)
		generated = classAssembler.AssembleClass(
			moduleName,
			packageName,
			className,
			"", // There are no existing example class files.
			classSynthesizer,
		)
		filename = directory + packageName + "/" + className + ".go"
		uti.WriteFile(filename, generated)
	}
}
