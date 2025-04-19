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
	gen "github.com/craterdog/go-code-generation/v6/generator"
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

// Generator

type (
	ClassGeneratorClassLike   = gen.ClassGeneratorClassLike
	ModuleGeneratorClassLike  = gen.ModuleGeneratorClassLike
	PackageGeneratorClassLike = gen.PackageGeneratorClassLike
)

type (
	ClassGeneratorLike   = gen.ClassGeneratorLike
	ModuleGeneratorLike  = gen.ModuleGeneratorLike
	PackageGeneratorLike = gen.PackageGeneratorLike
)

type (
	ClassTemplateDriven   = gen.ClassTemplateDriven
	ModuleTemplateDriven  = gen.ModuleTemplateDriven
	PackageTemplateDriven = gen.PackageTemplateDriven
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

// Generator/ClassGenerator

func ClassGenerator() gen.ClassGeneratorLike {
	return gen.ClassGeneratorClass().ClassGenerator()
}

// Generator/ModuleGenerator

func ModuleGenerator() gen.ModuleGeneratorLike {
	return gen.ModuleGeneratorClass().ModuleGenerator()
}

// Generator/PackageGenerator

func PackageGenerator() gen.PackageGeneratorLike {
	return gen.PackageGeneratorClass().PackageGenerator()
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
