/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
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
│         This "module_api.go" file was automatically generated using:         │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
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
	mod "github.com/craterdog/go-class-model/v8"
	ana "github.com/craterdog/go-code-generation/v8/analyzer"
	ass "github.com/craterdog/go-code-generation/v8/assembler"
	syn "github.com/craterdog/go-code-generation/v8/synthesizer"
	com "github.com/craterdog/go-essential-composites/v8"
	not "github.com/craterdog/go-syntax-notation/v8"
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

// CLASS ACCESSORS

// Analyzer

func ClassAnalyzerClass() ClassAnalyzerClassLike {
	return ana.ClassAnalyzerClass()
}

func ClassAnalyzer(
	model mod.ModelLike,
	className string,
) ClassAnalyzerLike {
	return ClassAnalyzerClass().ClassAnalyzer(
		model,
		className,
	)
}

func PackageAnalyzerClass() PackageAnalyzerClassLike {
	return ana.PackageAnalyzerClass()
}

func PackageAnalyzer(
	model mod.ModelLike,
) PackageAnalyzerLike {
	return PackageAnalyzerClass().PackageAnalyzer(
		model,
	)
}

func SyntaxAnalyzerClass() SyntaxAnalyzerClassLike {
	return ana.SyntaxAnalyzerClass()
}

func SyntaxAnalyzer(
	syntax not.SyntaxLike,
) SyntaxAnalyzerLike {
	return SyntaxAnalyzerClass().SyntaxAnalyzer(
		syntax,
	)
}

// Assembler

func ClassAssemblerClass() ClassAssemblerClassLike {
	return ass.ClassAssemblerClass()
}

func ClassAssembler() ClassAssemblerLike {
	return ClassAssemblerClass().ClassAssembler()
}

func ModuleAssemblerClass() ModuleAssemblerClassLike {
	return ass.ModuleAssemblerClass()
}

func ModuleAssembler() ModuleAssemblerLike {
	return ModuleAssemblerClass().ModuleAssembler()
}

func PackageAssemblerClass() PackageAssemblerClassLike {
	return ass.PackageAssemblerClass()
}

func PackageAssembler() PackageAssemblerLike {
	return PackageAssemblerClass().PackageAssembler()
}

// Synthesizer

func AstSynthesizerClass() AstSynthesizerClassLike {
	return syn.AstSynthesizerClass()
}

func AstSynthesizer(
	syntax not.SyntaxLike,
) AstSynthesizerLike {
	return AstSynthesizerClass().AstSynthesizer(
		syntax,
	)
}

func ClassSynthesizerClass() ClassSynthesizerClassLike {
	return syn.ClassSynthesizerClass()
}

func ClassSynthesizer(
	model mod.ModelLike,
	className string,
) ClassSynthesizerLike {
	return ClassSynthesizerClass().ClassSynthesizer(
		model,
		className,
	)
}

func FormatterSynthesizerClass() FormatterSynthesizerClassLike {
	return syn.FormatterSynthesizerClass()
}

func FormatterSynthesizer(
	syntax not.SyntaxLike,
) FormatterSynthesizerLike {
	return FormatterSynthesizerClass().FormatterSynthesizer(
		syntax,
	)
}

func GrammarSynthesizerClass() GrammarSynthesizerClassLike {
	return syn.GrammarSynthesizerClass()
}

func GrammarSynthesizer(
	syntax not.SyntaxLike,
) GrammarSynthesizerLike {
	return GrammarSynthesizerClass().GrammarSynthesizer(
		syntax,
	)
}

func ModuleSynthesizerClass() ModuleSynthesizerClassLike {
	return syn.ModuleSynthesizerClass()
}

func ModuleSynthesizer(
	moduleName string,
	models com.CatalogLike[string, mod.ModelLike],
) ModuleSynthesizerLike {
	return ModuleSynthesizerClass().ModuleSynthesizer(
		moduleName,
		models,
	)
}

func NodeSynthesizerClass() NodeSynthesizerClassLike {
	return syn.NodeSynthesizerClass()
}

func NodeSynthesizer(
	model mod.ModelLike,
	className string,
) NodeSynthesizerLike {
	return NodeSynthesizerClass().NodeSynthesizer(
		model,
		className,
	)
}

func PackageSynthesizerClass() PackageSynthesizerClassLike {
	return syn.PackageSynthesizerClass()
}

func PackageSynthesizer() PackageSynthesizerLike {
	return PackageSynthesizerClass().PackageSynthesizer()
}

func ParserSynthesizerClass() ParserSynthesizerClassLike {
	return syn.ParserSynthesizerClass()
}

func ParserSynthesizer(
	syntax not.SyntaxLike,
) ParserSynthesizerLike {
	return ParserSynthesizerClass().ParserSynthesizer(
		syntax,
	)
}

func ProcessorSynthesizerClass() ProcessorSynthesizerClassLike {
	return syn.ProcessorSynthesizerClass()
}

func ProcessorSynthesizer(
	syntax not.SyntaxLike,
) ProcessorSynthesizerLike {
	return ProcessorSynthesizerClass().ProcessorSynthesizer(
		syntax,
	)
}

func ScannerSynthesizerClass() ScannerSynthesizerClassLike {
	return syn.ScannerSynthesizerClass()
}

func ScannerSynthesizer(
	syntax not.SyntaxLike,
) ScannerSynthesizerLike {
	return ScannerSynthesizerClass().ScannerSynthesizer(
		syntax,
	)
}

func TokenSynthesizerClass() TokenSynthesizerClassLike {
	return syn.TokenSynthesizerClass()
}

func TokenSynthesizer(
	syntax not.SyntaxLike,
) TokenSynthesizerLike {
	return TokenSynthesizerClass().TokenSynthesizer(
		syntax,
	)
}

func ValidatorSynthesizerClass() ValidatorSynthesizerClassLike {
	return syn.ValidatorSynthesizerClass()
}

func ValidatorSynthesizer(
	syntax not.SyntaxLike,
) ValidatorSynthesizerLike {
	return ValidatorSynthesizerClass().ValidatorSynthesizer(
		syntax,
	)
}

func VisitorSynthesizerClass() VisitorSynthesizerClassLike {
	return syn.VisitorSynthesizerClass()
}

func VisitorSynthesizer(
	syntax not.SyntaxLike,
) VisitorSynthesizerLike {
	return VisitorSynthesizerClass().VisitorSynthesizer(
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
	models com.CatalogLike[string, mod.ModelLike],
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
