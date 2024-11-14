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
Package "module" defines type aliases for the commonly used types defined in the
packages contained in this module.  It also provides a universal constructor for
each commonly used class that is exported by the module.  Each constructor
delegates the actual construction process to its corresponding concrete class
defined in the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - github.com/craterdog/go-syntax-notation/wiki
*/
package module

import (
	fmt "fmt"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-syntax-notation/v5/ast"
	gra "github.com/craterdog/go-syntax-notation/v5/grammar"
)

// TYPE ALIASES

// Ast

type (
	AlternativeLike = ast.AlternativeLike
	CardinalityLike = ast.CardinalityLike
	CharacterLike   = ast.CharacterLike
	ConstrainedLike = ast.ConstrainedLike
	DefinitionLike  = ast.DefinitionLike
	ElementLike     = ast.ElementLike
	ExplicitLike    = ast.ExplicitLike
	ExpressionLike  = ast.ExpressionLike
	ExtentLike      = ast.ExtentLike
	FilterLike      = ast.FilterLike
	GroupLike       = ast.GroupLike
	IdentifierLike  = ast.IdentifierLike
	InlineLike      = ast.InlineLike
	LimitLike       = ast.LimitLike
	LineLike        = ast.LineLike
	MultilineLike   = ast.MultilineLike
	NoticeLike      = ast.NoticeLike
	OptionLike      = ast.OptionLike
	PatternLike     = ast.PatternLike
	QuantifiedLike  = ast.QuantifiedLike
	ReferenceLike   = ast.ReferenceLike
	RepetitionLike  = ast.RepetitionLike
	RuleLike        = ast.RuleLike
	SyntaxLike      = ast.SyntaxLike
	TermLike        = ast.TermLike
	TextLike        = ast.TextLike
)

// Grammar

type (
	TokenType     = gra.TokenType
	FormatterLike = gra.FormatterLike
	ParserLike    = gra.ParserLike
	ProcessorLike = gra.ProcessorLike
	ScannerLike   = gra.ScannerLike
	TokenLike     = gra.TokenLike
	ValidatorLike = gra.ValidatorLike
	VisitorLike   = gra.VisitorLike
	Methodical    = gra.Methodical
)

// UNIVERSAL CONSTRUCTORS

// Ast

func Alternative(arguments ...any) AlternativeLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case OptionLike:
			argumentTypes += "OptionLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Alternative constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var alternative AlternativeLike
	switch argumentTypes {
	case "OptionLike":
		var option = arguments[0].(OptionLike)
		alternative = ast.Alternative().Make(
			option,
		)
	default:
		var message = fmt.Sprintf(
			"No Alternative constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return alternative
}

func Cardinality(arguments ...any) CardinalityLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			argumentTypes += "any, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Cardinality constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var cardinality CardinalityLike
	switch argumentTypes {
	case "any":
		var any_ = arguments[0].(any)
		cardinality = ast.Cardinality().Make(
			any_,
		)
	default:
		var message = fmt.Sprintf(
			"No Cardinality constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return cardinality
}

func Character(arguments ...any) CharacterLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			argumentTypes += "any, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Character constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var character CharacterLike
	switch argumentTypes {
	case "any":
		var any_ = arguments[0].(any)
		character = ast.Character().Make(
			any_,
		)
	default:
		var message = fmt.Sprintf(
			"No Character constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return character
}

func Constrained(arguments ...any) ConstrainedLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			argumentTypes += "any, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Constrained constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var constrained ConstrainedLike
	switch argumentTypes {
	case "any":
		var any_ = arguments[0].(any)
		constrained = ast.Constrained().Make(
			any_,
		)
	default:
		var message = fmt.Sprintf(
			"No Constrained constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return constrained
}

func Definition(arguments ...any) DefinitionLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			argumentTypes += "any, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Definition constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var definition DefinitionLike
	switch argumentTypes {
	case "any":
		var any_ = arguments[0].(any)
		definition = ast.Definition().Make(
			any_,
		)
	default:
		var message = fmt.Sprintf(
			"No Definition constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return definition
}

func Element(arguments ...any) ElementLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			argumentTypes += "any, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Element constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var element ElementLike
	switch argumentTypes {
	case "any":
		var any_ = arguments[0].(any)
		element = ast.Element().Make(
			any_,
		)
	default:
		var message = fmt.Sprintf(
			"No Element constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return element
}

func Explicit(arguments ...any) ExplicitLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case ExtentLike:
			argumentTypes += "ExtentLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Explicit constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var explicit ExplicitLike
	switch argumentTypes {
	case "string, ExtentLike":
		var glyph = arguments[0].(string)
		var optionalExtent = arguments[1].(ExtentLike)
		explicit = ast.Explicit().Make(
			glyph,
			optionalExtent,
		)
	default:
		var message = fmt.Sprintf(
			"No Explicit constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return explicit
}

func Expression(arguments ...any) ExpressionLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case PatternLike:
			argumentTypes += "PatternLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Expression constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var expression ExpressionLike
	switch argumentTypes {
	case "string, PatternLike, string":
		var lowercase = arguments[0].(string)
		var pattern = arguments[1].(PatternLike)
		var optionalNote = arguments[2].(string)
		expression = ast.Expression().Make(
			lowercase,
			pattern,
			optionalNote,
		)
	default:
		var message = fmt.Sprintf(
			"No Expression constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return expression
}

func Extent(arguments ...any) ExtentLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Extent constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var extent ExtentLike
	switch argumentTypes {
	case "string":
		var glyph = arguments[0].(string)
		extent = ast.Extent().Make(
			glyph,
		)
	default:
		var message = fmt.Sprintf(
			"No Extent constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return extent
}

func Filter(arguments ...any) FilterLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case abs.Sequential[CharacterLike]:
			argumentTypes += "abs.Sequential[CharacterLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Filter constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var filter FilterLike
	switch argumentTypes {
	case "string, abs.Sequential[CharacterLike]":
		var optionalExcluded = arguments[0].(string)
		var characters = arguments[1].(abs.Sequential[CharacterLike])
		filter = ast.Filter().Make(
			optionalExcluded,
			characters,
		)
	default:
		var message = fmt.Sprintf(
			"No Filter constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return filter
}

func Group(arguments ...any) GroupLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case PatternLike:
			argumentTypes += "PatternLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Group constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var group GroupLike
	switch argumentTypes {
	case "PatternLike":
		var pattern = arguments[0].(PatternLike)
		group = ast.Group().Make(
			pattern,
		)
	default:
		var message = fmt.Sprintf(
			"No Group constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return group
}

func Identifier(arguments ...any) IdentifierLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			argumentTypes += "any, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Identifier constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var identifier IdentifierLike
	switch argumentTypes {
	case "any":
		var any_ = arguments[0].(any)
		identifier = ast.Identifier().Make(
			any_,
		)
	default:
		var message = fmt.Sprintf(
			"No Identifier constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return identifier
}

func Inline(arguments ...any) InlineLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[TermLike]:
			argumentTypes += "abs.Sequential[TermLike], "
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Inline constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var inline InlineLike
	switch argumentTypes {
	case "abs.Sequential[TermLike], string":
		var terms = arguments[0].(abs.Sequential[TermLike])
		var optionalNote = arguments[1].(string)
		inline = ast.Inline().Make(
			terms,
			optionalNote,
		)
	default:
		var message = fmt.Sprintf(
			"No Inline constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return inline
}

func Limit(arguments ...any) LimitLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Limit constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var limit LimitLike
	switch argumentTypes {
	case "string":
		var optionalNumber = arguments[0].(string)
		limit = ast.Limit().Make(
			optionalNumber,
		)
	default:
		var message = fmt.Sprintf(
			"No Limit constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return limit
}

func Line(arguments ...any) LineLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case IdentifierLike:
			argumentTypes += "IdentifierLike, "
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Line constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var line LineLike
	switch argumentTypes {
	case "IdentifierLike, string":
		var identifier = arguments[0].(IdentifierLike)
		var optionalNote = arguments[1].(string)
		line = ast.Line().Make(
			identifier,
			optionalNote,
		)
	default:
		var message = fmt.Sprintf(
			"No Line constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return line
}

func Multiline(arguments ...any) MultilineLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[LineLike]:
			argumentTypes += "abs.Sequential[LineLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Multiline constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var multiline MultilineLike
	switch argumentTypes {
	case "abs.Sequential[LineLike]":
		var lines = arguments[0].(abs.Sequential[LineLike])
		multiline = ast.Multiline().Make(
			lines,
		)
	default:
		var message = fmt.Sprintf(
			"No Multiline constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return multiline
}

func Notice(arguments ...any) NoticeLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Notice constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var notice NoticeLike
	switch argumentTypes {
	case "string":
		var comment = arguments[0].(string)
		notice = ast.Notice().Make(
			comment,
		)
	default:
		var message = fmt.Sprintf(
			"No Notice constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return notice
}

func Option(arguments ...any) OptionLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[RepetitionLike]:
			argumentTypes += "abs.Sequential[RepetitionLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Option constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var option OptionLike
	switch argumentTypes {
	case "abs.Sequential[RepetitionLike]":
		var repetitions = arguments[0].(abs.Sequential[RepetitionLike])
		option = ast.Option().Make(
			repetitions,
		)
	default:
		var message = fmt.Sprintf(
			"No Option constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return option
}

func Pattern(arguments ...any) PatternLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case OptionLike:
			argumentTypes += "OptionLike, "
		case abs.Sequential[AlternativeLike]:
			argumentTypes += "abs.Sequential[AlternativeLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Pattern constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var pattern PatternLike
	switch argumentTypes {
	case "OptionLike, abs.Sequential[AlternativeLike]":
		var option = arguments[0].(OptionLike)
		var alternatives = arguments[1].(abs.Sequential[AlternativeLike])
		pattern = ast.Pattern().Make(
			option,
			alternatives,
		)
	default:
		var message = fmt.Sprintf(
			"No Pattern constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return pattern
}

func Quantified(arguments ...any) QuantifiedLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case LimitLike:
			argumentTypes += "LimitLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Quantified constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var quantified QuantifiedLike
	switch argumentTypes {
	case "string, LimitLike":
		var number = arguments[0].(string)
		var optionalLimit = arguments[1].(LimitLike)
		quantified = ast.Quantified().Make(
			number,
			optionalLimit,
		)
	default:
		var message = fmt.Sprintf(
			"No Quantified constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return quantified
}

func Reference(arguments ...any) ReferenceLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case IdentifierLike:
			argumentTypes += "IdentifierLike, "
		case CardinalityLike:
			argumentTypes += "CardinalityLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Reference constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var reference ReferenceLike
	switch argumentTypes {
	case "IdentifierLike, CardinalityLike":
		var identifier = arguments[0].(IdentifierLike)
		var optionalCardinality = arguments[1].(CardinalityLike)
		reference = ast.Reference().Make(
			identifier,
			optionalCardinality,
		)
	default:
		var message = fmt.Sprintf(
			"No Reference constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return reference
}

func Repetition(arguments ...any) RepetitionLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ElementLike:
			argumentTypes += "ElementLike, "
		case CardinalityLike:
			argumentTypes += "CardinalityLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Repetition constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var repetition RepetitionLike
	switch argumentTypes {
	case "ElementLike, CardinalityLike":
		var element = arguments[0].(ElementLike)
		var optionalCardinality = arguments[1].(CardinalityLike)
		repetition = ast.Repetition().Make(
			element,
			optionalCardinality,
		)
	default:
		var message = fmt.Sprintf(
			"No Repetition constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return repetition
}

func Rule(arguments ...any) RuleLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case DefinitionLike:
			argumentTypes += "DefinitionLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Rule constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var rule RuleLike
	switch argumentTypes {
	case "string, DefinitionLike":
		var uppercase = arguments[0].(string)
		var definition = arguments[1].(DefinitionLike)
		rule = ast.Rule().Make(
			uppercase,
			definition,
		)
	default:
		var message = fmt.Sprintf(
			"No Rule constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return rule
}

func Syntax(arguments ...any) SyntaxLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case NoticeLike:
			argumentTypes += "NoticeLike, "
		case string:
			argumentTypes += "string, "
		case abs.Sequential[RuleLike]:
			argumentTypes += "abs.Sequential[RuleLike], "
		case abs.Sequential[ExpressionLike]:
			argumentTypes += "abs.Sequential[ExpressionLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Syntax constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var syntax SyntaxLike
	switch argumentTypes {
	case "NoticeLike, string, abs.Sequential[RuleLike], string, abs.Sequential[ExpressionLike]":
		var notice = arguments[0].(NoticeLike)
		var comment1 = arguments[1].(string)
		var rules = arguments[2].(abs.Sequential[RuleLike])
		var comment2 = arguments[3].(string)
		var expressions = arguments[4].(abs.Sequential[ExpressionLike])
		syntax = ast.Syntax().Make(
			notice,
			comment1,
			rules,
			comment2,
			expressions,
		)
	default:
		var message = fmt.Sprintf(
			"No Syntax constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return syntax
}

func Term(arguments ...any) TermLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			argumentTypes += "any, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Term constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var term TermLike
	switch argumentTypes {
	case "any":
		var any_ = arguments[0].(any)
		term = ast.Term().Make(
			any_,
		)
	default:
		var message = fmt.Sprintf(
			"No Term constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return term
}

func Text(arguments ...any) TextLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			argumentTypes += "any, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Text constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var text TextLike
	switch argumentTypes {
	case "any":
		var any_ = arguments[0].(any)
		text = ast.Text().Make(
			any_,
		)
	default:
		var message = fmt.Sprintf(
			"No Text constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return text
}

// Grammar

func Formatter(arguments ...any) FormatterLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Formatter constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var formatter FormatterLike
	switch argumentTypes {
	case "":
		formatter = gra.Formatter().Make()
	default:
		var message = fmt.Sprintf(
			"No Formatter constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return formatter
}

func Parser(arguments ...any) ParserLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Parser constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var parser ParserLike
	switch argumentTypes {
	case "":
		parser = gra.Parser().Make()
	default:
		var message = fmt.Sprintf(
			"No Parser constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return parser
}

func Processor(arguments ...any) ProcessorLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Processor constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var processor ProcessorLike
	switch argumentTypes {
	case "":
		processor = gra.Processor().Make()
	default:
		var message = fmt.Sprintf(
			"No Processor constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return processor
}

func Scanner(arguments ...any) ScannerLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case abs.QueueLike[TokenLike]:
			argumentTypes += "abs.QueueLike[TokenLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Scanner constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var scanner ScannerLike
	switch argumentTypes {
	case "string, abs.QueueLike[TokenLike]":
		var source = arguments[0].(string)
		var tokens = arguments[1].(abs.QueueLike[TokenLike])
		scanner = gra.Scanner().Make(
			source,
			tokens,
		)
	default:
		var message = fmt.Sprintf(
			"No Scanner constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return scanner
}

func Token(arguments ...any) TokenLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case uint:
			argumentTypes += "uint, "
		case TokenType:
			argumentTypes += "TokenType, "
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Token constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var token TokenLike
	switch argumentTypes {
	case "uint, uint, TokenType, string":
		var line = arguments[0].(uint)
		var position = arguments[1].(uint)
		var type_ = arguments[2].(TokenType)
		var value = arguments[3].(string)
		token = gra.Token().Make(
			line,
			position,
			type_,
			value,
		)
	default:
		var message = fmt.Sprintf(
			"No Token constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return token
}

func Validator(arguments ...any) ValidatorLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Validator constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var validator ValidatorLike
	switch argumentTypes {
	case "":
		validator = gra.Validator().Make()
	default:
		var message = fmt.Sprintf(
			"No Validator constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return validator
}

func Visitor(arguments ...any) VisitorLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case Methodical:
			argumentTypes += "Methodical, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Visitor constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length-1]
	}

	// Call the corresponding constructor.
	var visitor VisitorLike
	switch argumentTypes {
	case "Methodical":
		var processor = arguments[0].(Methodical)
		visitor = gra.Visitor().Make(
			processor,
		)
	default:
		var message = fmt.Sprintf(
			"No Visitor constructor matching the arguments was found: $v\n",
			arguments,
		)
		panic(message)
	}
	return visitor
}
