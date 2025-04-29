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

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│              This "Module.go" file was automatically generated.              │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides constructors for each
commonly used class that is exported by the module.  Each constructor delegates
the actual construction process to its corresponding concrete class declared in
the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/craterdog/go-code-generation/wiki
*/
package module

import (
	mod "github.com/craterdog/go-class-model/v5"
	ana "github.com/craterdog/go-code-generation/v6/analyzer"
	ass "github.com/craterdog/go-code-generation/v6/assembler"
	syn "github.com/craterdog/go-code-generation/v6/synthesizer"
	col "github.com/craterdog/go-collection-framework/v5/collection"
	not "github.com/craterdog/go-syntax-notation/v6"
)

// TYPE ALIASES

// Analyzer

type (
	ClassAnalyzerClassLike   = ana.ClassAnalyzerClassLike
	PackageAnalyzerClassLike = ana.PackageAnalyzerClassLike
	SyntaxAnalyzerClassLike  = ana.SyntaxAnalyzerClassLike
)

type (
	ClassAnalyzerLike   = ana.ClassAnalyzerLike
	PackageAnalyzerLike = ana.PackageAnalyzerLike
	SyntaxAnalyzerLike  = ana.SyntaxAnalyzerLike
)

// Assembler

type (
	ClassAssemblerClassLike   = ass.ClassAssemblerClassLike
	ModuleAssemblerClassLike  = ass.ModuleAssemblerClassLike
	PackageAssemblerClassLike = ass.PackageAssemblerClassLike
)

type (
	ClassAssemblerLike   = ass.ClassAssemblerLike
	ModuleAssemblerLike  = ass.ModuleAssemblerLike
	PackageAssemblerLike = ass.PackageAssemblerLike
)

type (
	ClassTemplateDriven   = ass.ClassTemplateDriven
	ModuleTemplateDriven  = ass.ModuleTemplateDriven
	PackageTemplateDriven = ass.PackageTemplateDriven
)

// Synthesizer

type (
	AstSynthesizerClassLike       = syn.AstSynthesizerClassLike
	ClassSynthesizerClassLike     = syn.ClassSynthesizerClassLike
	FormatterSynthesizerClassLike = syn.FormatterSynthesizerClassLike
	GrammarSynthesizerClassLike   = syn.GrammarSynthesizerClassLike
	ModuleSynthesizerClassLike    = syn.ModuleSynthesizerClassLike
	NodeSynthesizerClassLike      = syn.NodeSynthesizerClassLike
	PackageSynthesizerClassLike   = syn.PackageSynthesizerClassLike
	ParserSynthesizerClassLike    = syn.ParserSynthesizerClassLike
	ProcessorSynthesizerClassLike = syn.ProcessorSynthesizerClassLike
	ScannerSynthesizerClassLike   = syn.ScannerSynthesizerClassLike
	TokenSynthesizerClassLike     = syn.TokenSynthesizerClassLike
	ValidatorSynthesizerClassLike = syn.ValidatorSynthesizerClassLike
	VisitorSynthesizerClassLike   = syn.VisitorSynthesizerClassLike
)

type (
	AstSynthesizerLike       = syn.AstSynthesizerLike
	ClassSynthesizerLike     = syn.ClassSynthesizerLike
	FormatterSynthesizerLike = syn.FormatterSynthesizerLike
	GrammarSynthesizerLike   = syn.GrammarSynthesizerLike
	ModuleSynthesizerLike    = syn.ModuleSynthesizerLike
	NodeSynthesizerLike      = syn.NodeSynthesizerLike
	PackageSynthesizerLike   = syn.PackageSynthesizerLike
	ParserSynthesizerLike    = syn.ParserSynthesizerLike
	ProcessorSynthesizerLike = syn.ProcessorSynthesizerLike
	ScannerSynthesizerLike   = syn.ScannerSynthesizerLike
	TokenSynthesizerLike     = syn.TokenSynthesizerLike
	ValidatorSynthesizerLike = syn.ValidatorSynthesizerLike
	VisitorSynthesizerLike   = syn.VisitorSynthesizerLike
)

// CLASS CONSTRUCTORS

// Analyzer/ClassAnalyzer

func ClassAnalyzer(
	model mod.ModelLike,
	className string,
) ana.ClassAnalyzerLike {
	return ana.ClassAnalyzerClass().ClassAnalyzer(
		model,
		className,
	)
}

// Analyzer/PackageAnalyzer

func PackageAnalyzer(
	model mod.ModelLike,
) ana.PackageAnalyzerLike {
	return ana.PackageAnalyzerClass().PackageAnalyzer(
		model,
	)
}

// Analyzer/SyntaxAnalyzer

func SyntaxAnalyzer(
	syntax not.SyntaxLike,
) ana.SyntaxAnalyzerLike {
	return ana.SyntaxAnalyzerClass().SyntaxAnalyzer(
		syntax,
	)
}

// Assembler/ClassAssembler

func ClassAssembler() ass.ClassAssemblerLike {
	return ass.ClassAssemblerClass().ClassAssembler()
}

// Assembler/ModuleAssembler

func ModuleAssembler() ass.ModuleAssemblerLike {
	return ass.ModuleAssemblerClass().ModuleAssembler()
}

// Assembler/PackageAssembler

func PackageAssembler() ass.PackageAssemblerLike {
	return ass.PackageAssemblerClass().PackageAssembler()
}

// Synthesizer/AstSynthesizer

func AstSynthesizer(
	syntax not.SyntaxLike,
) syn.AstSynthesizerLike {
	return syn.AstSynthesizerClass().AstSynthesizer(
		syntax,
	)
}

// Synthesizer/ClassSynthesizer

func ClassSynthesizer(
	model mod.ModelLike,
	className string,
) syn.ClassSynthesizerLike {
	return syn.ClassSynthesizerClass().ClassSynthesizer(
		model,
		className,
	)
}

// Synthesizer/FormatterSynthesizer

func FormatterSynthesizer(
	syntax not.SyntaxLike,
) syn.FormatterSynthesizerLike {
	return syn.FormatterSynthesizerClass().FormatterSynthesizer(
		syntax,
	)
}

// Synthesizer/GrammarSynthesizer

func GrammarSynthesizer(
	syntax not.SyntaxLike,
) syn.GrammarSynthesizerLike {
	return syn.GrammarSynthesizerClass().GrammarSynthesizer(
		syntax,
	)
}

// Synthesizer/ModuleSynthesizer

func ModuleSynthesizer(
	moduleName string,
	models col.CatalogLike[string, mod.ModelLike],
) syn.ModuleSynthesizerLike {
	return syn.ModuleSynthesizerClass().ModuleSynthesizer(
		moduleName,
		models,
	)
}

// Synthesizer/NodeSynthesizer

func NodeSynthesizer(
	model mod.ModelLike,
	className string,
) syn.NodeSynthesizerLike {
	return syn.NodeSynthesizerClass().NodeSynthesizer(
		model,
		className,
	)
}

// Synthesizer/PackageSynthesizer

func PackageSynthesizer() syn.PackageSynthesizerLike {
	return syn.PackageSynthesizerClass().PackageSynthesizer()
}

// Synthesizer/ParserSynthesizer

func ParserSynthesizer(
	syntax not.SyntaxLike,
) syn.ParserSynthesizerLike {
	return syn.ParserSynthesizerClass().ParserSynthesizer(
		syntax,
	)
}

// Synthesizer/ProcessorSynthesizer

func ProcessorSynthesizer(
	syntax not.SyntaxLike,
) syn.ProcessorSynthesizerLike {
	return syn.ProcessorSynthesizerClass().ProcessorSynthesizer(
		syntax,
	)
}

// Synthesizer/ScannerSynthesizer

func ScannerSynthesizer(
	syntax not.SyntaxLike,
) syn.ScannerSynthesizerLike {
	return syn.ScannerSynthesizerClass().ScannerSynthesizer(
		syntax,
	)
}

// Synthesizer/TokenSynthesizer

func TokenSynthesizer(
	syntax not.SyntaxLike,
) syn.TokenSynthesizerLike {
	return syn.TokenSynthesizerClass().TokenSynthesizer(
		syntax,
	)
}

// Synthesizer/ValidatorSynthesizer

func ValidatorSynthesizer(
	syntax not.SyntaxLike,
) syn.ValidatorSynthesizerLike {
	return syn.ValidatorSynthesizerClass().ValidatorSynthesizer(
		syntax,
	)
}

// Synthesizer/VisitorSynthesizer

func VisitorSynthesizer(
	syntax not.SyntaxLike,
) syn.VisitorSynthesizerLike {
	return syn.VisitorSynthesizerClass().VisitorSynthesizer(
		syntax,
	)
}

// GLOBAL FUNCTIONS

func CreatePackage(
	moduleName string,
	wikiPath string,
	packageName string,
) mod.ModelLike {
	var existing = "" // There is no existing package.
	var assembler = PackageAssembler()
	var synthesizer = PackageSynthesizer()
	var source = assembler.AssemblePackage(
		moduleName,
		wikiPath,
		packageName,
		existing,
		synthesizer,
	)
	var model = mod.ParseSource(source)
	return model
}

func GenerateAstPackage(
	moduleName string,
	wikiPath string,
	syntax not.SyntaxLike,
) mod.ModelLike {
	var packageName = "ast"
	var existing = "" // There is no existing AST package.
	var assembler = PackageAssembler()
	var synthesizer = AstSynthesizer(syntax)
	var source = assembler.AssemblePackage(
		moduleName,
		wikiPath,
		packageName,
		existing,
		synthesizer,
	)
	var model = mod.ParseSource(source)
	return model
}

func GenerateGrammarPackage(
	moduleName string,
	wikiPath string,
	syntax not.SyntaxLike,
) mod.ModelLike {
	var packageName = "grammar"
	var existing = "" // There is no existing grammar package.
	var assembler = PackageAssembler()
	var synthesizer = GrammarSynthesizer(syntax)
	var source = assembler.AssemblePackage(
		moduleName,
		wikiPath,
		packageName,
		existing,
		synthesizer,
	)
	var model = mod.ParseSource(source)
	return model
}

func GenerateModule(
	moduleName string,
	wikiPath string,
	existing string,
	models col.CatalogLike[string, mod.ModelLike],
) string {
	var assembler = ModuleAssembler()
	var synthesizer = ModuleSynthesizer(
		moduleName,
		models,
	)
	var source = assembler.AssembleModule(
		moduleName,
		wikiPath,
		existing,
		synthesizer,
	)
	return source
}

func GenerateClass(
	moduleName string,
	packageName string,
	className string,
	existing string,
	model mod.ModelLike,
) string {
	var assembler = ClassAssembler()
	var synthesizer = ClassSynthesizer(
		model,
		className,
	)
	var source = assembler.AssembleClass(
		moduleName,
		packageName,
		className,
		existing,
		synthesizer,
	)
	return source
}

func GenerateAstClass(
	moduleName string,
	className string,
	model mod.ModelLike,
) string {
	var packageName = "ast"
	var existing = "" // There is no existing AST class.
	var assembler = ClassAssembler()
	var synthesizer = NodeSynthesizer(model, className)
	var source = assembler.AssembleClass(
		moduleName,
		packageName,
		className,
		existing,
		synthesizer,
	)
	return source
}

func GenerateTokenClass(
	moduleName string,
	syntax not.SyntaxLike,
) string {
	var assembler = ClassAssembler()
	var packageName = "grammar"
	var className = "token"
	var existing = "" // There is no existing token class.
	var synthesizer = TokenSynthesizer(syntax)
	var source = assembler.AssembleClass(
		moduleName,
		packageName,
		className,
		existing,
		synthesizer,
	)
	return source
}

func GenerateScannerClass(
	moduleName string,
	syntax not.SyntaxLike,
) string {
	var assembler = ClassAssembler()
	var packageName = "grammar"
	var className = "scanner"
	var existing = "" // There is no existing scanner class.
	var synthesizer = ScannerSynthesizer(syntax)
	var source = assembler.AssembleClass(
		moduleName,
		packageName,
		className,
		existing,
		synthesizer,
	)
	return source
}

func GenerateParserClass(
	moduleName string,
	syntax not.SyntaxLike,
) string {
	var assembler = ClassAssembler()
	var packageName = "grammar"
	var className = "parser"
	var existing = "" // There is no existing parser class.
	var synthesizer = ParserSynthesizer(syntax)
	var source = assembler.AssembleClass(
		moduleName,
		packageName,
		className,
		existing,
		synthesizer,
	)
	return source
}

func GenerateFormatterClass(
	moduleName string,
	existing string,
	syntax not.SyntaxLike,
) string {
	var assembler = ClassAssembler()
	var packageName = "grammar"
	var className = "formatter"
	var synthesizer = FormatterSynthesizer(syntax)
	var source = assembler.AssembleClass(
		moduleName,
		packageName,
		className,
		existing,
		synthesizer,
	)
	return source
}

func GenerateValidatorClass(
	moduleName string,
	existing string,
	syntax not.SyntaxLike,
) string {
	var assembler = ClassAssembler()
	var packageName = "grammar"
	var className = "validator"
	var synthesizer = ValidatorSynthesizer(syntax)
	var source = assembler.AssembleClass(
		moduleName,
		packageName,
		className,
		existing,
		synthesizer,
	)
	return source
}

func GenerateVisitorClass(
	moduleName string,
	syntax not.SyntaxLike,
) string {
	var assembler = ClassAssembler()
	var packageName = "grammar"
	var className = "visitor"
	var existing = "" // There is no existing visitor class.
	var synthesizer = VisitorSynthesizer(syntax)
	var source = assembler.AssembleClass(
		moduleName,
		packageName,
		className,
		existing,
		synthesizer,
	)
	return source
}

func GenerateProcessorClass(
	moduleName string,
	syntax not.SyntaxLike,
) string {
	var assembler = ClassAssembler()
	var packageName = "grammar"
	var className = "processor"
	var existing = "" // There is no existing processor class.
	var synthesizer = ProcessorSynthesizer(syntax)
	var source = assembler.AssembleClass(
		moduleName,
		packageName,
		className,
		existing,
		synthesizer,
	)
	return source
}
