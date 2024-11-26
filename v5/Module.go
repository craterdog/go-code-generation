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

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│              This "Module.go" file was automatically generated.              │
│  Updates to any section other than the GLOBAL FUNCTIONS may be overwritten.  │
└──────────────────────────────────────────────────────────────────────────────┘
Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides a universal constructor
for each commonly used class that is exported by the module.  Each constructor
delegates the actual construction process to its corresponding concrete class
declared in the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - github.com/craterdog/go-code-generation/wiki
*/
package module

import (
	fmt "fmt"
	mod "github.com/craterdog/go-class-model/v5"
	ana "github.com/craterdog/go-code-generation/v5/analyzer"
	gen "github.com/craterdog/go-code-generation/v5/generator"
	syn "github.com/craterdog/go-code-generation/v5/synthesizer"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	not "github.com/craterdog/go-syntax-notation/v5"
)

// TYPE ALIASES

// Analyzer

type (
	ModelAnalyzerLike  = ana.ModelAnalyzerLike
	SyntaxAnalyzerLike = ana.SyntaxAnalyzerLike
)

// Generator

type (
	ClassGeneratorLike    = gen.ClassGeneratorLike
	ModuleGeneratorLike   = gen.ModuleGeneratorLike
	PackageGeneratorLike  = gen.PackageGeneratorLike
	TemplateGeneratorLike = gen.TemplateGeneratorLike
)

type (
	ClassTemplateDriven   = gen.ClassTemplateDriven
	ModuleTemplateDriven  = gen.ModuleTemplateDriven
	PackageTemplateDriven = gen.PackageTemplateDriven
)

// Synthesizer

type (
	AstSynthesizerLike       = syn.AstSynthesizerLike
	ClassSynthesizerLike     = syn.ClassSynthesizerLike
	ExampleSynthesizerLike   = syn.ExampleSynthesizerLike
	FormatterSynthesizerLike = syn.FormatterSynthesizerLike
	GrammarSynthesizerLike   = syn.GrammarSynthesizerLike
	ModuleSynthesizerLike    = syn.ModuleSynthesizerLike
	ParserSynthesizerLike    = syn.ParserSynthesizerLike
	ProcessorSynthesizerLike = syn.ProcessorSynthesizerLike
	ScannerSynthesizerLike   = syn.ScannerSynthesizerLike
	TokenSynthesizerLike     = syn.TokenSynthesizerLike
	ValidatorSynthesizerLike = syn.ValidatorSynthesizerLike
	VisitorSynthesizerLike   = syn.VisitorSynthesizerLike
)

// UNIVERSAL CONSTRUCTORS

// Analyzer

func ModelAnalyzer(arguments ...any) ModelAnalyzerLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case mod.ModelLike:
			argumentTypes += "mod.ModelLike, "
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the ModelAnalyzer constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ ModelAnalyzerLike
	switch argumentTypes {
	case "mod.ModelLike, string":
		var model = arguments[0].(mod.ModelLike)
		var className = arguments[1].(string)
		instance_ = ana.ModelAnalyzer().Make(
			model,
			className,
		)
	default:
		var message = fmt.Sprintf(
			"No ModelAnalyzer constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func SyntaxAnalyzer(arguments ...any) SyntaxAnalyzerLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case not.SyntaxLike:
			argumentTypes += "not.SyntaxLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the SyntaxAnalyzer constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ SyntaxAnalyzerLike
	switch argumentTypes {
	case "not.SyntaxLike":
		var syntax = arguments[0].(not.SyntaxLike)
		instance_ = ana.SyntaxAnalyzer().Make(
			syntax,
		)
	default:
		var message = fmt.Sprintf(
			"No SyntaxAnalyzer constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

// Generator

func ClassGenerator(arguments ...any) ClassGeneratorLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the ClassGenerator constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ ClassGeneratorLike
	switch argumentTypes {
	case "":
		instance_ = gen.ClassGenerator().Make()
	default:
		var message = fmt.Sprintf(
			"No ClassGenerator constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func ModuleGenerator(arguments ...any) ModuleGeneratorLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the ModuleGenerator constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ ModuleGeneratorLike
	switch argumentTypes {
	case "":
		instance_ = gen.ModuleGenerator().Make()
	default:
		var message = fmt.Sprintf(
			"No ModuleGenerator constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func PackageGenerator(arguments ...any) PackageGeneratorLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the PackageGenerator constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ PackageGeneratorLike
	switch argumentTypes {
	case "":
		instance_ = gen.PackageGenerator().Make()
	default:
		var message = fmt.Sprintf(
			"No PackageGenerator constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func TemplateGenerator(arguments ...any) TemplateGeneratorLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the TemplateGenerator constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ TemplateGeneratorLike
	switch argumentTypes {
	case "":
		instance_ = gen.TemplateGenerator().Make()
	default:
		var message = fmt.Sprintf(
			"No TemplateGenerator constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

// Synthesizer

func AstSynthesizer(arguments ...any) AstSynthesizerLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case not.SyntaxLike:
			argumentTypes += "not.SyntaxLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the AstSynthesizer constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ AstSynthesizerLike
	switch argumentTypes {
	case "not.SyntaxLike":
		var syntax = arguments[0].(not.SyntaxLike)
		instance_ = syn.AstSynthesizer().Make(
			syntax,
		)
	default:
		var message = fmt.Sprintf(
			"No AstSynthesizer constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func ClassSynthesizer(arguments ...any) ClassSynthesizerLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case mod.ModelLike:
			argumentTypes += "mod.ModelLike, "
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the ClassSynthesizer constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ ClassSynthesizerLike
	switch argumentTypes {
	case "mod.ModelLike, string":
		var model = arguments[0].(mod.ModelLike)
		var className = arguments[1].(string)
		instance_ = syn.ClassSynthesizer().Make(
			model,
			className,
		)
	default:
		var message = fmt.Sprintf(
			"No ClassSynthesizer constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func ExampleSynthesizer(arguments ...any) ExampleSynthesizerLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the ExampleSynthesizer constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ ExampleSynthesizerLike
	switch argumentTypes {
	case "":
		instance_ = syn.ExampleSynthesizer().Make()
	default:
		var message = fmt.Sprintf(
			"No ExampleSynthesizer constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func FormatterSynthesizer(arguments ...any) FormatterSynthesizerLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case not.SyntaxLike:
			argumentTypes += "not.SyntaxLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the FormatterSynthesizer constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ FormatterSynthesizerLike
	switch argumentTypes {
	case "not.SyntaxLike":
		var syntax = arguments[0].(not.SyntaxLike)
		instance_ = syn.FormatterSynthesizer().Make(
			syntax,
		)
	default:
		var message = fmt.Sprintf(
			"No FormatterSynthesizer constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func GrammarSynthesizer(arguments ...any) GrammarSynthesizerLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case not.SyntaxLike:
			argumentTypes += "not.SyntaxLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the GrammarSynthesizer constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ GrammarSynthesizerLike
	switch argumentTypes {
	case "not.SyntaxLike":
		var syntax = arguments[0].(not.SyntaxLike)
		instance_ = syn.GrammarSynthesizer().Make(
			syntax,
		)
	default:
		var message = fmt.Sprintf(
			"No GrammarSynthesizer constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func ModuleSynthesizer(arguments ...any) ModuleSynthesizerLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.CatalogLike[string, mod.ModelLike]:
			argumentTypes += "abs.CatalogLike[string, mod.ModelLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the ModuleSynthesizer constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ ModuleSynthesizerLike
	switch argumentTypes {
	case "abs.CatalogLike[string, mod.ModelLike]":
		var models = arguments[0].(abs.CatalogLike[string, mod.ModelLike])
		instance_ = syn.ModuleSynthesizer().Make(
			models,
		)
	default:
		var message = fmt.Sprintf(
			"No ModuleSynthesizer constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func ParserSynthesizer(arguments ...any) ParserSynthesizerLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case not.SyntaxLike:
			argumentTypes += "not.SyntaxLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the ParserSynthesizer constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ ParserSynthesizerLike
	switch argumentTypes {
	case "not.SyntaxLike":
		var syntax = arguments[0].(not.SyntaxLike)
		instance_ = syn.ParserSynthesizer().Make(
			syntax,
		)
	default:
		var message = fmt.Sprintf(
			"No ParserSynthesizer constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func ProcessorSynthesizer(arguments ...any) ProcessorSynthesizerLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case not.SyntaxLike:
			argumentTypes += "not.SyntaxLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the ProcessorSynthesizer constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ ProcessorSynthesizerLike
	switch argumentTypes {
	case "not.SyntaxLike":
		var syntax = arguments[0].(not.SyntaxLike)
		instance_ = syn.ProcessorSynthesizer().Make(
			syntax,
		)
	default:
		var message = fmt.Sprintf(
			"No ProcessorSynthesizer constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func ScannerSynthesizer(arguments ...any) ScannerSynthesizerLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case not.SyntaxLike:
			argumentTypes += "not.SyntaxLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the ScannerSynthesizer constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ ScannerSynthesizerLike
	switch argumentTypes {
	case "not.SyntaxLike":
		var syntax = arguments[0].(not.SyntaxLike)
		instance_ = syn.ScannerSynthesizer().Make(
			syntax,
		)
	default:
		var message = fmt.Sprintf(
			"No ScannerSynthesizer constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func TokenSynthesizer(arguments ...any) TokenSynthesizerLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case not.SyntaxLike:
			argumentTypes += "not.SyntaxLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the TokenSynthesizer constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ TokenSynthesizerLike
	switch argumentTypes {
	case "not.SyntaxLike":
		var syntax = arguments[0].(not.SyntaxLike)
		instance_ = syn.TokenSynthesizer().Make(
			syntax,
		)
	default:
		var message = fmt.Sprintf(
			"No TokenSynthesizer constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func ValidatorSynthesizer(arguments ...any) ValidatorSynthesizerLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case not.SyntaxLike:
			argumentTypes += "not.SyntaxLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the ValidatorSynthesizer constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ ValidatorSynthesizerLike
	switch argumentTypes {
	case "not.SyntaxLike":
		var syntax = arguments[0].(not.SyntaxLike)
		instance_ = syn.ValidatorSynthesizer().Make(
			syntax,
		)
	default:
		var message = fmt.Sprintf(
			"No ValidatorSynthesizer constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func VisitorSynthesizer(arguments ...any) VisitorSynthesizerLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case not.SyntaxLike:
			argumentTypes += "not.SyntaxLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the VisitorSynthesizer constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ VisitorSynthesizerLike
	switch argumentTypes {
	case "not.SyntaxLike":
		var syntax = arguments[0].(not.SyntaxLike)
		instance_ = syn.VisitorSynthesizer().Make(
			syntax,
		)
	default:
		var message = fmt.Sprintf(
			"No VisitorSynthesizer constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

// GLOBAL FUNCTIONS
