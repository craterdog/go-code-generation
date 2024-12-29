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
	gen "github.com/craterdog/go-code-generation/v5"
	fra "github.com/craterdog/go-collection-framework/v5"
	uti "github.com/craterdog/go-missing-utilities/v2"
	not "github.com/craterdog/go-syntax-notation/v5"
	ass "github.com/stretchr/testify/assert"
	sts "strings"
	tes "testing"
)

var directory = "./testdata/"
var moduleName = "github.com/craterdog/go-class-model/v5"
var wikiPath = "https://github.com/craterdog/go-class-model/wiki"

func TestPackageGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)
	not.ValidateSyntax(syntax)
	var actual = not.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the AST Package.go file.
	var packageName = "ast"
	uti.RemakeDirectory(directory + packageName)
	var generator = gen.PackageGenerator()
	var astSynthesizer = gen.AstSynthesizer(syntax)
	var generated = generator.GeneratePackage(
		moduleName,
		wikiPath,
		packageName,
		"", // There is no pre-existing AST package file.
		astSynthesizer,
	)
	filename = directory + packageName + "/Package.go"
	uti.WriteFile(filename, generated)

	// Generate the grammar Package.go file.
	packageName = "grammar"
	filename = directory + packageName + "/Package.go"
	var existing = uti.ReadFile(filename)
	var grammarSynthesizer = gen.GrammarSynthesizer(syntax)
	generated = generator.GeneratePackage(
		moduleName,
		wikiPath,
		packageName,
		existing,
		grammarSynthesizer,
	)
	uti.WriteFile(filename, generated)
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
	var generator = gen.ModuleGenerator()
	var moduleSynthesizer = gen.ModuleSynthesizer(models)
	var generated = generator.GenerateModule(
		moduleName,
		wikiPath,
		existing,
		moduleSynthesizer,
	)
	uti.WriteFile(filename, generated)
}

func TestAstGeneration(t *tes.T) {
	// Validate the class model.
	var packageName = "ast"
	var filename = directory + packageName + "/Package.go"
	var source = uti.ReadFile(filename)
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
		var generated = generator.GenerateClass(
			moduleName,
			packageName,
			className,
			"", // There is no pre-existing AST class file.
			nodeSynthesizer,
		)
		filename = directory + packageName + "/" + className + ".go"
		uti.WriteFile(filename, generated)
	}
}

func TestTokenGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)
	not.ValidateSyntax(syntax)
	var actual = not.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the token concrete class.
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "token"
	filename = directory + packageName + "/" + className + ".go"
	var existing = uti.ReadFile(filename)
	var tokenSynthesizer = gen.TokenSynthesizer(syntax)
	var generated = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		existing,
		tokenSynthesizer,
	)
	uti.WriteFile(filename, generated)
}

func TestScannerGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)
	not.ValidateSyntax(syntax)
	var actual = not.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the scanner concrete class.
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "scanner"
	filename = directory + packageName + "/" + className + ".go"
	var existing = uti.ReadFile(filename)
	var scannerSynthesizer = gen.ScannerSynthesizer(syntax)
	var generated = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		existing,
		scannerSynthesizer,
	)
	uti.WriteFile(filename, generated)
}

func TestParserGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)
	not.ValidateSyntax(syntax)
	var actual = not.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the parser concrete class.
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "parser"
	filename = directory + packageName + "/" + className + ".go"
	var existing = uti.ReadFile(filename)
	var parserSynthesizer = gen.ParserSynthesizer(syntax)
	var generated = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		existing,
		parserSynthesizer,
	)
	uti.WriteFile(filename, generated)
}

func TestVisitorGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)
	not.ValidateSyntax(syntax)
	var actual = not.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the visitor concrete class.
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "visitor"
	filename = directory + packageName + "/" + className + ".go"
	var existing = uti.ReadFile(filename)
	var visitorSynthesizer = gen.VisitorSynthesizer(syntax)
	var generated = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		existing,
		visitorSynthesizer,
	)
	uti.WriteFile(filename, generated)
}

func TestFormatterGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)
	not.ValidateSyntax(syntax)
	var actual = not.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the formatter concrete class.
	var packageName = "grammar"
	var className = "formatter"
	filename = directory + packageName + "/" + className + ".go"
	var existing = uti.ReadFile(filename)
	var generator = gen.ClassGenerator()
	var formatterSynthesizer = gen.FormatterSynthesizer(syntax)
	var generated = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		existing,
		formatterSynthesizer,
	)
	uti.WriteFile(filename, generated)
}

func TestProcessorGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)
	not.ValidateSyntax(syntax)
	var actual = not.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the processor concrete class.
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "processor"
	filename = directory + packageName + "/" + className + ".go"
	var existing = uti.ReadFile(filename)
	var processorSynthesizer = gen.ProcessorSynthesizer(syntax)
	var generated = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		existing,
		processorSynthesizer,
	)
	uti.WriteFile(filename, generated)
}

func TestValidatorGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var source = uti.ReadFile(filename)
	var syntax = not.ParseSource(source)
	not.ValidateSyntax(syntax)
	var actual = not.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the validator concrete class.
	var packageName = "grammar"
	var className = "validator"
	filename = directory + packageName + "/" + className + ".go"
	var existing = uti.ReadFile(filename)
	var generator = gen.ClassGenerator()
	var validatorSynthesizer = gen.ValidatorSynthesizer(syntax)
	var generated = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		existing,
		validatorSynthesizer,
	)
	uti.WriteFile(filename, generated)
}

func TestExampleGeneration(t *tes.T) {
	// Generate the example class model.
	var packageName = "example"
	var wikiPath = "https://github.com/craterdog/go-package-example/wiki"
	uti.RemakeDirectory(directory + packageName)
	var packageGenerator = gen.PackageGenerator()
	var packageSynthesizer = gen.PackageSynthesizer()
	var generated = packageGenerator.GeneratePackage(
		moduleName,
		wikiPath,
		packageName,
		"", // There is no pre-existing example package file.
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
		generated = classGenerator.GenerateClass(
			moduleName,
			packageName,
			className,
			"", // There are no pre-existing example class files.
			classSynthesizer,
		)
		filename = directory + packageName + "/" + className + ".go"
		uti.WriteFile(filename, generated)
	}
}
