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
	byt "bytes"
	mod "github.com/craterdog/go-class-model/v5"
	gen "github.com/craterdog/go-code-generation/v5/generator"
	syn "github.com/craterdog/go-code-generation/v5/synthesizer"
	col "github.com/craterdog/go-collection-framework/v4"
	uti "github.com/craterdog/go-missing-utilities/v2"
	not "github.com/craterdog/go-syntax-notation/v5"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	reg "regexp"
	sts "strings"
	tes "testing"
)

var directory = "../testdata/"
var moduleName = "github.com/craterdog/go-class-model/v5"
var wikiPath = "github.com/craterdog/go-class-model/wiki"

func TestPackageGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var bytes, err = osx.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = not.Parser()
	var syntax = parser.ParseSource(source)
	var validator = not.Validator()
	validator.ValidateSyntax(syntax)
	var formatter = not.Formatter()
	var actual = formatter.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the AST Package.go file.
	var generator = gen.PackageGenerator().Make()
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
}

func TestModuleGeneration(t *tes.T) {
	var models = col.Catalog[string, mod.ModelLike]()
	var packages = []string{
		"ast",
		"grammar",
	}
	for _, packageName := range packages {
		var filename = directory + packageName + "/Package.go"
		var bytes, err = osx.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		var source = string(bytes)
		var parser = mod.Parser()
		var model = parser.ParseSource(source)
		models.SetValue(packageName, model)
	}
	var generator = gen.ModuleGenerator().Make()
	var moduleSynthesizer = syn.ModuleSynthesizer().Make(models)
	var source = generator.GenerateModule(
		moduleName,
		wikiPath,
		moduleSynthesizer,
	)
	var filename = directory + "Module.go"
	var bytes = replaceGlobalFunctions(filename, source)
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func TestAstGeneration(t *tes.T) {
	// Validate the class model.
	var packageName = "ast"
	var filename = directory + packageName + "/Package.go"
	var bytes, err = osx.ReadFile(filename)
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
	var parser = not.Parser()
	var syntax = parser.ParseSource(source)
	var validator = not.Validator()
	validator.ValidateSyntax(syntax)
	var formatter = not.Formatter()
	var actual = formatter.FormatSyntax(syntax)
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
	var parser = not.Parser()
	var syntax = parser.ParseSource(source)
	var validator = not.Validator()
	validator.ValidateSyntax(syntax)
	var formatter = not.Formatter()
	var actual = formatter.FormatSyntax(syntax)
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
	var parser = not.Parser()
	var syntax = parser.ParseSource(source)
	var validator = not.Validator()
	validator.ValidateSyntax(syntax)
	var formatter = not.Formatter()
	var actual = formatter.FormatSyntax(syntax)
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

func TestVisitorGeneration(t *tes.T) {
	// Validate the language grammar.
	var filename = directory + "Syntax.cdsn"
	var bytes, err = osx.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = not.Parser()
	var syntax = parser.ParseSource(source)
	var validator = not.Validator()
	validator.ValidateSyntax(syntax)
	var formatter = not.Formatter()
	var actual = formatter.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the visitor concrete class.
	var generator = gen.ClassGenerator().Make()
	var packageName = "grammar"
	var className = "visitor"
	filename = directory + packageName + "/" + className + ".go"
	var visitorSynthesizer = syn.VisitorSynthesizer().Make(syntax)
	source = generator.GenerateClass(
		moduleName,
		wikiPath,
		packageName,
		className,
		visitorSynthesizer,
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
	var parser = not.Parser()
	var syntax = parser.ParseSource(source)
	var validator = not.Validator()
	validator.ValidateSyntax(syntax)
	var formatter = not.Formatter()
	var actual = formatter.FormatSyntax(syntax)
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
	bytes = replaceMethodicalMethods(filename, source)
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
	var parser = not.Parser()
	var syntax = parser.ParseSource(source)
	var validator = not.Validator()
	validator.ValidateSyntax(syntax)
	var formatter = not.Formatter()
	var actual = formatter.FormatSyntax(syntax)
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
	var parser = not.Parser()
	var syntax = parser.ParseSource(source)
	var validator = not.Validator()
	validator.ValidateSyntax(syntax)
	var formatter = not.Formatter()
	var actual = formatter.FormatSyntax(syntax)
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
	bytes = replaceMethodicalMethods(filename, source)
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
	var parser = mod.Parser()
	var model = parser.ParseSource(source)
	var validator = mod.Validator()
	validator.ValidateModel(model)
	var formatter = mod.Formatter()
	var actual = formatter.FormatModel(model)
	ass.Equal(t, source, actual)

	// Generate the example concrete classes.
	var generator = gen.ClassGenerator().Make()
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

func replaceGlobalFunctions(
	filename string,
	source string,
) []byte {
	var bytes, err = osx.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var matcher = reg.MustCompile(
		`// GLOBAL FUNCTIONS(.|\r?\n)*`,
	)
	var original = matcher.Find(bytes)
	bytes = []byte(source)
	var generated = matcher.Find(bytes)
	bytes = byt.ReplaceAll(bytes, generated, original)
	return bytes
}

func replaceMethodicalMethods(
	filename string,
	source string,
) []byte {
	var bytes, err = osx.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var matcher = reg.MustCompile(
		`// Methodical Methods(.|\r?\n)+// PROTECTED INTERFACE`,
	)
	var original = matcher.Find(bytes)
	bytes = []byte(source)
	var generated = matcher.Find(bytes)
	bytes = byt.ReplaceAll(bytes, generated, original)
	return bytes
}
