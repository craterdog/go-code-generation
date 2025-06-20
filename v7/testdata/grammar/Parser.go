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
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	fmt "fmt"
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	ast "github.com/craterdog/go-syntax-notation/v7/ast"
	mat "math"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ParserClass() ParserClassLike {
	return parserClass()
}

// Constructor Methods

func (c *parserClass_) Parser() ParserLike {
	var instance = &parser_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *parser_) GetClass() ParserClassLike {
	return parserClass()
}

func (v *parser_) ParseSource(
	source string,
) ast.SyntaxLike {
	v.source_ = sts.ReplaceAll(source, "\t", "    ")
	v.tokens_ = fra.Queue[TokenLike]()
	v.next_ = fra.Stack[TokenLike]()

	// The scanner runs in a separate Go routine.
	ScannerClass().Scanner(v.source_, v.tokens_)

	// Attempt to parse the syntax.
	var syntax, token, ok = v.parseSyntax()
	if !ok || !v.tokens_.IsEmpty() {
		var message = v.formatError("$Syntax", token)
		panic(message)
	}
	return syntax
}

// PROTECTED INTERFACE

// Private Methods

func (v *parser_) parseAlternativeSequence() (
	alternativeSequence ast.AlternativeSequenceLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "|" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("|")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AlternativeSequence rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AlternativeSequence", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Sequence rule.
	var sequence ast.SequenceLike
	sequence, token, ok = v.parseSequence()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Sequence rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AlternativeSequence", token)
		panic(message)
	}

	// Found a single AlternativeSequence rule.
	ok = true
	v.remove(tokens)
	alternativeSequence = ast.AlternativeSequenceClass().AlternativeSequence(
		delimiter,
		sequence,
	)
	return
}

func (v *parser_) parseAlternatives() (
	alternatives ast.AlternativesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Sequence rule.
	var sequence ast.SequenceLike
	sequence, token, ok = v.parseSequence()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Sequence rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Alternatives", token)
		panic(message)
	}

	// Attempt to parse multiple AlternativeSequence rules.
	var alternativeSequences = fra.List[ast.AlternativeSequenceLike]()
alternativeSequencesLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var alternativeSequence ast.AlternativeSequenceLike
		alternativeSequence, token, ok = v.parseAlternativeSequence()
		if !ok {
			switch {
			case count_ >= 0:
				break alternativeSequencesLoop
			case uti.IsDefined(tokens):
				// This is not multiple AlternativeSequence rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Alternatives", token)
				message += "0 or more AlternativeSequence rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		alternativeSequences.AppendValue(alternativeSequence)
	}

	// Found a single Alternatives rule.
	ok = true
	v.remove(tokens)
	alternatives = ast.AlternativesClass().Alternatives(
		sequence,
		alternativeSequences,
	)
	return
}

func (v *parser_) parseCardinality() (
	cardinality ast.CardinalityLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Constrained Cardinality.
	var constrained ast.ConstrainedLike
	constrained, token, ok = v.parseConstrained()
	if ok {
		// Found a single Constrained Cardinality.
		cardinality = ast.CardinalityClass().Cardinality(constrained)
		return
	}

	// Attempt to parse a single Quantified Cardinality.
	var quantified ast.QuantifiedLike
	quantified, token, ok = v.parseQuantified()
	if ok {
		// Found a single Quantified Cardinality.
		cardinality = ast.CardinalityClass().Cardinality(quantified)
		return
	}

	// This is not a single Cardinality rule.
	return
}

func (v *parser_) parseCharacter() (
	character ast.CharacterLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Implicit Character.
	var implicit ast.ImplicitLike
	implicit, token, ok = v.parseImplicit()
	if ok {
		// Found a single Implicit Character.
		character = ast.CharacterClass().Character(implicit)
		return
	}

	// Attempt to parse a single Explicit Character.
	var explicit ast.ExplicitLike
	explicit, token, ok = v.parseExplicit()
	if ok {
		// Found a single Explicit Character.
		character = ast.CharacterClass().Character(explicit)
		return
	}

	// This is not a single Character rule.
	return
}

func (v *parser_) parseComponent() (
	component ast.ComponentLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single literal Component.
	var literal string
	literal, token, ok = v.parseToken(LiteralToken)
	if ok {
		// Found a single literal Component.
		component = ast.ComponentClass().Component(literal)
		return
	}

	// Attempt to parse a single lowercase Component.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken)
	if ok {
		// Found a single lowercase Component.
		component = ast.ComponentClass().Component(lowercase)
		return
	}

	// Attempt to parse a single uppercase Component.
	var uppercase string
	uppercase, token, ok = v.parseToken(UppercaseToken)
	if ok {
		// Found a single uppercase Component.
		component = ast.ComponentClass().Component(uppercase)
		return
	}

	// This is not a single Component rule.
	return
}

func (v *parser_) parseConstrained() (
	constrained ast.ConstrainedLike,
	token TokenLike,
	ok bool,
) {
	var delimiter string

	// Attempt to parse a single "?" delimiter.
	delimiter, token, ok = v.parseDelimiter("?")
	if ok {
		// Found a single "?" delimiter.
		constrained = ast.ConstrainedClass().Constrained(delimiter)
		return
	}

	// Attempt to parse a single "*" delimiter.
	delimiter, token, ok = v.parseDelimiter("*")
	if ok {
		// Found a single "*" delimiter.
		constrained = ast.ConstrainedClass().Constrained(delimiter)
		return
	}

	// Attempt to parse a single "+" delimiter.
	delimiter, token, ok = v.parseDelimiter("+")
	if ok {
		// Found a single "+" delimiter.
		constrained = ast.ConstrainedClass().Constrained(delimiter)
		return
	}

	// This is not a single Constrained rule.
	return
}

func (v *parser_) parseDefinition() (
	definition ast.DefinitionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single LiteralAlternatives Definition.
	var literalAlternatives ast.LiteralAlternativesLike
	literalAlternatives, token, ok = v.parseLiteralAlternatives()
	if ok {
		// Found a single LiteralAlternatives Definition.
		definition = ast.DefinitionClass().Definition(literalAlternatives)
		return
	}

	// Attempt to parse a single ExpressionAlternatives Definition.
	var expressionAlternatives ast.ExpressionAlternativesLike
	expressionAlternatives, token, ok = v.parseExpressionAlternatives()
	if ok {
		// Found a single ExpressionAlternatives Definition.
		definition = ast.DefinitionClass().Definition(expressionAlternatives)
		return
	}

	// Attempt to parse a single RuleAlternatives Definition.
	var ruleAlternatives ast.RuleAlternativesLike
	ruleAlternatives, token, ok = v.parseRuleAlternatives()
	if ok {
		// Found a single RuleAlternatives Definition.
		definition = ast.DefinitionClass().Definition(ruleAlternatives)
		return
	}

	// Attempt to parse a single TermSequence Definition.
	var termSequence ast.TermSequenceLike
	termSequence, token, ok = v.parseTermSequence()
	if ok {
		// Found a single TermSequence Definition.
		definition = ast.DefinitionClass().Definition(termSequence)
		return
	}

	// This is not a single Definition rule.
	return
}

func (v *parser_) parseElement() (
	element ast.ElementLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Group Element.
	var group ast.GroupLike
	group, token, ok = v.parseGroup()
	if ok {
		// Found a single Group Element.
		element = ast.ElementClass().Element(group)
		return
	}

	// Attempt to parse a single Filter Element.
	var filter ast.FilterLike
	filter, token, ok = v.parseFilter()
	if ok {
		// Found a single Filter Element.
		element = ast.ElementClass().Element(filter)
		return
	}

	// Attempt to parse a single Text Element.
	var text ast.TextLike
	text, token, ok = v.parseText()
	if ok {
		// Found a single Text Element.
		element = ast.ElementClass().Element(text)
		return
	}

	// This is not a single Element rule.
	return
}

func (v *parser_) parseExplicit() (
	explicit ast.ExplicitLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single glyph token.
	var glyph string
	glyph, token, ok = v.parseToken(GlyphToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single glyph token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Explicit", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional Extent rule.
	var optionalExtent ast.ExtentLike
	optionalExtent, _, ok = v.parseExtent()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single Explicit rule.
	ok = true
	v.remove(tokens)
	explicit = ast.ExplicitClass().Explicit(
		glyph,
		optionalExtent,
	)
	return
}

func (v *parser_) parseExpression() (
	expression ast.ExpressionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "$" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("$")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Expression rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Expression", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single lowercase token.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single lowercase token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Expression", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ":" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter(":")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Expression rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Expression", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Pattern rule.
	var pattern ast.PatternLike
	pattern, token, ok = v.parsePattern()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Pattern rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Expression", token)
		panic(message)
	}

	// Found a single Expression rule.
	ok = true
	v.remove(tokens)
	expression = ast.ExpressionClass().Expression(
		delimiter1,
		lowercase,
		delimiter2,
		pattern,
	)
	return
}

func (v *parser_) parseExpressionAlternatives() (
	expressionAlternatives ast.ExpressionAlternativesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse multiple ExpressionName rules.
	var expressionNames = fra.List[ast.ExpressionNameLike]()
expressionNamesLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var expressionName ast.ExpressionNameLike
		expressionName, token, ok = v.parseExpressionName()
		if !ok {
			switch {
			case count_ >= 1:
				break expressionNamesLoop
			case uti.IsDefined(tokens):
				// This is not multiple ExpressionName rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$ExpressionAlternatives", token)
				message += "1 or more ExpressionName rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		expressionNames.AppendValue(expressionName)
	}

	// Found a single ExpressionAlternatives rule.
	ok = true
	v.remove(tokens)
	expressionAlternatives = ast.ExpressionAlternativesClass().ExpressionAlternatives(expressionNames)
	return
}

func (v *parser_) parseExpressionName() (
	expressionName ast.ExpressionNameLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single newline token.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single newline token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExpressionName", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single lowercase token.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single lowercase token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExpressionName", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, token, ok = v.parseToken(NoteToken)
	if ok {
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	} else {
		optionalNote = "" // Reset this to undefined.
	}

	// Found a single ExpressionName rule.
	ok = true
	v.remove(tokens)
	expressionName = ast.ExpressionNameClass().ExpressionName(
		newline,
		lowercase,
		optionalNote,
	)
	return
}

func (v *parser_) parseExtent() (
	extent ast.ExtentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single ".." literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Extent rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Extent", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single glyph token.
	var glyph string
	glyph, token, ok = v.parseToken(GlyphToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single glyph token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Extent", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Extent rule.
	ok = true
	v.remove(tokens)
	extent = ast.ExtentClass().Extent(
		delimiter,
		glyph,
	)
	return
}

func (v *parser_) parseFilter() (
	filter ast.FilterLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse an optional "~" literal.
	var optionalDelimiter string
	optionalDelimiter, token, ok = v.parseDelimiter("~")
	if ok {
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	} else {
		optionalDelimiter = "" // Reset this to undefined.
	}

	// Attempt to parse a single "[" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Filter rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Filter", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple Character rules.
	var characters = fra.List[ast.CharacterLike]()
charactersLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var character ast.CharacterLike
		character, token, ok = v.parseCharacter()
		if !ok {
			switch {
			case count_ >= 1:
				break charactersLoop
			case uti.IsDefined(tokens):
				// This is not multiple Character rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Filter", token)
				message += "1 or more Character rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		characters.AppendValue(character)
	}

	// Attempt to parse a single "]" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Filter rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Filter", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Filter rule.
	ok = true
	v.remove(tokens)
	filter = ast.FilterClass().Filter(
		optionalDelimiter,
		delimiter1,
		characters,
		delimiter2,
	)
	return
}

func (v *parser_) parseFragment() (
	fragment ast.FragmentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "$" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("$")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Fragment rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Fragment", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single allcaps token.
	var allcaps string
	allcaps, token, ok = v.parseToken(AllcapsToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single allcaps token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Fragment", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ":" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter(":")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Fragment rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Fragment", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Pattern rule.
	var pattern ast.PatternLike
	pattern, token, ok = v.parsePattern()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Pattern rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Fragment", token)
		panic(message)
	}

	// Found a single Fragment rule.
	ok = true
	v.remove(tokens)
	fragment = ast.FragmentClass().Fragment(
		delimiter1,
		allcaps,
		delimiter2,
		pattern,
	)
	return
}

func (v *parser_) parseGroup() (
	group ast.GroupLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Group rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Group", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Alternatives rule.
	var alternatives ast.AlternativesLike
	alternatives, token, ok = v.parseAlternatives()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Alternatives rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Group", token)
		panic(message)
	}

	// Attempt to parse a single ")" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Group rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Group", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Group rule.
	ok = true
	v.remove(tokens)
	group = ast.GroupClass().Group(
		delimiter1,
		alternatives,
		delimiter2,
	)
	return
}

func (v *parser_) parseImplicit() (
	implicit ast.ImplicitLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single intrinsic token.
	var intrinsic string
	intrinsic, token, ok = v.parseToken(IntrinsicToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single intrinsic token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Implicit", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Implicit rule.
	ok = true
	v.remove(tokens)
	implicit = ast.ImplicitClass().Implicit(intrinsic)
	return
}

func (v *parser_) parseLegalNotice() (
	legalNotice ast.LegalNoticeLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single comment token.
	var comment string
	comment, token, ok = v.parseToken(CommentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single comment token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$LegalNotice", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single LegalNotice rule.
	ok = true
	v.remove(tokens)
	legalNotice = ast.LegalNoticeClass().LegalNotice(comment)
	return
}

func (v *parser_) parseLimit() (
	limit ast.LimitLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single ".." literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Limit rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Limit", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional number token.
	var optionalNumber string
	optionalNumber, token, ok = v.parseToken(NumberToken)
	if ok {
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	} else {
		optionalNumber = "" // Reset this to undefined.
	}

	// Found a single Limit rule.
	ok = true
	v.remove(tokens)
	limit = ast.LimitClass().Limit(
		delimiter,
		optionalNumber,
	)
	return
}

func (v *parser_) parseLiteralAlternatives() (
	literalAlternatives ast.LiteralAlternativesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse multiple LiteralValue rules.
	var literalValues = fra.List[ast.LiteralValueLike]()
literalValuesLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var literalValue ast.LiteralValueLike
		literalValue, token, ok = v.parseLiteralValue()
		if !ok {
			switch {
			case count_ >= 1:
				break literalValuesLoop
			case uti.IsDefined(tokens):
				// This is not multiple LiteralValue rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$LiteralAlternatives", token)
				message += "1 or more LiteralValue rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		literalValues.AppendValue(literalValue)
	}

	// Found a single LiteralAlternatives rule.
	ok = true
	v.remove(tokens)
	literalAlternatives = ast.LiteralAlternativesClass().LiteralAlternatives(literalValues)
	return
}

func (v *parser_) parseLiteralValue() (
	literalValue ast.LiteralValueLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single newline token.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single newline token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$LiteralValue", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single literal token.
	var literal string
	literal, token, ok = v.parseToken(LiteralToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single literal token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$LiteralValue", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, token, ok = v.parseToken(NoteToken)
	if ok {
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	} else {
		optionalNote = "" // Reset this to undefined.
	}

	// Found a single LiteralValue rule.
	ok = true
	v.remove(tokens)
	literalValue = ast.LiteralValueClass().LiteralValue(
		newline,
		literal,
		optionalNote,
	)
	return
}

func (v *parser_) parsePattern() (
	pattern ast.PatternLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Alternatives rule.
	var alternatives ast.AlternativesLike
	alternatives, token, ok = v.parseAlternatives()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Alternatives rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Pattern", token)
		panic(message)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, token, ok = v.parseToken(NoteToken)
	if ok {
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	} else {
		optionalNote = "" // Reset this to undefined.
	}

	// Found a single Pattern rule.
	ok = true
	v.remove(tokens)
	pattern = ast.PatternClass().Pattern(
		alternatives,
		optionalNote,
	)
	return
}

func (v *parser_) parseQuantified() (
	quantified ast.QuantifiedLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "{" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("{")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Quantified rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Quantified", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single number token.
	var number string
	number, token, ok = v.parseToken(NumberToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single number token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Quantified", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional Limit rule.
	var optionalLimit ast.LimitLike
	optionalLimit, _, ok = v.parseLimit()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Attempt to parse a single "}" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("}")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Quantified rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Quantified", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Quantified rule.
	ok = true
	v.remove(tokens)
	quantified = ast.QuantifiedClass().Quantified(
		delimiter1,
		number,
		optionalLimit,
		delimiter2,
	)
	return
}

func (v *parser_) parseRepetition() (
	repetition ast.RepetitionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Element rule.
	var element ast.ElementLike
	element, token, ok = v.parseElement()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Element rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Repetition", token)
		panic(message)
	}

	// Attempt to parse an optional Cardinality rule.
	var optionalCardinality ast.CardinalityLike
	optionalCardinality, _, ok = v.parseCardinality()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single Repetition rule.
	ok = true
	v.remove(tokens)
	repetition = ast.RepetitionClass().Repetition(
		element,
		optionalCardinality,
	)
	return
}

func (v *parser_) parseRule() (
	rule ast.RuleLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "$" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("$")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Rule rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Rule", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single uppercase token.
	var uppercase string
	uppercase, token, ok = v.parseToken(UppercaseToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single uppercase token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Rule", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ":" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter(":")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Rule rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Rule", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Definition rule.
	var definition ast.DefinitionLike
	definition, token, ok = v.parseDefinition()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Definition rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Rule", token)
		panic(message)
	}

	// Found a single Rule rule.
	ok = true
	v.remove(tokens)
	rule = ast.RuleClass().Rule(
		delimiter1,
		uppercase,
		delimiter2,
		definition,
	)
	return
}

func (v *parser_) parseRuleAlternatives() (
	ruleAlternatives ast.RuleAlternativesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse multiple RuleName rules.
	var ruleNames = fra.List[ast.RuleNameLike]()
ruleNamesLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var ruleName ast.RuleNameLike
		ruleName, token, ok = v.parseRuleName()
		if !ok {
			switch {
			case count_ >= 1:
				break ruleNamesLoop
			case uti.IsDefined(tokens):
				// This is not multiple RuleName rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$RuleAlternatives", token)
				message += "1 or more RuleName rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		ruleNames.AppendValue(ruleName)
	}

	// Found a single RuleAlternatives rule.
	ok = true
	v.remove(tokens)
	ruleAlternatives = ast.RuleAlternativesClass().RuleAlternatives(ruleNames)
	return
}

func (v *parser_) parseRuleName() (
	ruleName ast.RuleNameLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single newline token.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single newline token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$RuleName", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single uppercase token.
	var uppercase string
	uppercase, token, ok = v.parseToken(UppercaseToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single uppercase token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$RuleName", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, token, ok = v.parseToken(NoteToken)
	if ok {
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	} else {
		optionalNote = "" // Reset this to undefined.
	}

	// Found a single RuleName rule.
	ok = true
	v.remove(tokens)
	ruleName = ast.RuleNameClass().RuleName(
		newline,
		uppercase,
		optionalNote,
	)
	return
}

func (v *parser_) parseRuleTerm() (
	ruleTerm ast.RuleTermLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Component rule.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Component rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$RuleTerm", token)
		panic(message)
	}

	// Attempt to parse an optional Cardinality rule.
	var optionalCardinality ast.CardinalityLike
	optionalCardinality, _, ok = v.parseCardinality()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single RuleTerm rule.
	ok = true
	v.remove(tokens)
	ruleTerm = ast.RuleTermClass().RuleTerm(
		component,
		optionalCardinality,
	)
	return
}

func (v *parser_) parseSequence() (
	sequence ast.SequenceLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse multiple Repetition rules.
	var repetitions = fra.List[ast.RepetitionLike]()
repetitionsLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var repetition ast.RepetitionLike
		repetition, token, ok = v.parseRepetition()
		if !ok {
			switch {
			case count_ >= 1:
				break repetitionsLoop
			case uti.IsDefined(tokens):
				// This is not multiple Repetition rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Sequence", token)
				message += "1 or more Repetition rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		repetitions.AppendValue(repetition)
	}

	// Found a single Sequence rule.
	ok = true
	v.remove(tokens)
	sequence = ast.SequenceClass().Sequence(repetitions)
	return
}

func (v *parser_) parseSyntax() (
	syntax ast.SyntaxLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single LegalNotice rule.
	var legalNotice ast.LegalNoticeLike
	legalNotice, token, ok = v.parseLegalNotice()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single LegalNotice rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Syntax", token)
		panic(message)
	}

	// Attempt to parse a single comment token.
	var comment1 string
	comment1, token, ok = v.parseToken(CommentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single comment token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Syntax", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple Rule rules.
	var rules = fra.List[ast.RuleLike]()
rulesLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var rule ast.RuleLike
		rule, token, ok = v.parseRule()
		if !ok {
			switch {
			case count_ >= 1:
				break rulesLoop
			case uti.IsDefined(tokens):
				// This is not multiple Rule rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Syntax", token)
				message += "1 or more Rule rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		rules.AppendValue(rule)
	}

	// Attempt to parse a single comment token.
	var comment2 string
	comment2, token, ok = v.parseToken(CommentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single comment token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Syntax", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple Expression rules.
	var expressions = fra.List[ast.ExpressionLike]()
expressionsLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var expression ast.ExpressionLike
		expression, token, ok = v.parseExpression()
		if !ok {
			switch {
			case count_ >= 1:
				break expressionsLoop
			case uti.IsDefined(tokens):
				// This is not multiple Expression rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Syntax", token)
				message += "1 or more Expression rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		expressions.AppendValue(expression)
	}

	// Attempt to parse a single comment token.
	var comment3 string
	comment3, token, ok = v.parseToken(CommentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single comment token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Syntax", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple Fragment rules.
	var fragments = fra.List[ast.FragmentLike]()
fragmentsLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var fragment ast.FragmentLike
		fragment, token, ok = v.parseFragment()
		if !ok {
			switch {
			case count_ >= 1:
				break fragmentsLoop
			case uti.IsDefined(tokens):
				// This is not multiple Fragment rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Syntax", token)
				message += "1 or more Fragment rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		fragments.AppendValue(fragment)
	}

	// Found a single Syntax rule.
	ok = true
	v.remove(tokens)
	syntax = ast.SyntaxClass().Syntax(
		legalNotice,
		comment1,
		rules,
		comment2,
		expressions,
		comment3,
		fragments,
	)
	return
}

func (v *parser_) parseTermSequence() (
	termSequence ast.TermSequenceLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse multiple RuleTerm rules.
	var ruleTerms = fra.List[ast.RuleTermLike]()
ruleTermsLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var ruleTerm ast.RuleTermLike
		ruleTerm, token, ok = v.parseRuleTerm()
		if !ok {
			switch {
			case count_ >= 1:
				break ruleTermsLoop
			case uti.IsDefined(tokens):
				// This is not multiple RuleTerm rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$TermSequence", token)
				message += "1 or more RuleTerm rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		ruleTerms.AppendValue(ruleTerm)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, token, ok = v.parseToken(NoteToken)
	if ok {
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	} else {
		optionalNote = "" // Reset this to undefined.
	}

	// Found a single TermSequence rule.
	ok = true
	v.remove(tokens)
	termSequence = ast.TermSequenceClass().TermSequence(
		ruleTerms,
		optionalNote,
	)
	return
}

func (v *parser_) parseText() (
	text ast.TextLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single glyph Text.
	var glyph string
	glyph, token, ok = v.parseToken(GlyphToken)
	if ok {
		// Found a single glyph Text.
		text = ast.TextClass().Text(glyph)
		return
	}

	// Attempt to parse a single literal Text.
	var literal string
	literal, token, ok = v.parseToken(LiteralToken)
	if ok {
		// Found a single literal Text.
		text = ast.TextClass().Text(literal)
		return
	}

	// Attempt to parse a single intrinsic Text.
	var intrinsic string
	intrinsic, token, ok = v.parseToken(IntrinsicToken)
	if ok {
		// Found a single intrinsic Text.
		text = ast.TextClass().Text(intrinsic)
		return
	}

	// Attempt to parse a single lowercase Text.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken)
	if ok {
		// Found a single lowercase Text.
		text = ast.TextClass().Text(lowercase)
		return
	}

	// Attempt to parse a single allcaps Text.
	var allcaps string
	allcaps, token, ok = v.parseToken(AllcapsToken)
	if ok {
		// Found a single allcaps Text.
		text = ast.TextClass().Text(allcaps)
		return
	}

	// This is not a single Text rule.
	return
}

func (v *parser_) parseDelimiter(
	literal string,
) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single delimiter.
	value, token, ok = v.parseToken(DelimiterToken)
	if ok {
		if value == literal {
			// Found the desired delimiter.
			return
		}
		v.next_.AddValue(token)
		ok = false
	}

	// This is not the desired delimiter.
	return
}

func (v *parser_) parseToken(
	tokenType TokenType,
) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a specific token type.
	var tokens = fra.List[TokenLike]()
	token = v.getNextToken()
	for token != nil {
		tokens.AppendValue(token)
		switch token.GetType() {
		case tokenType:
			// Found the desired token type.
			value = token.GetValue()
			ok = true
			return
		case SpaceToken, NewlineToken:
			// Ignore any unrequested whitespace.
			token = v.getNextToken()
		default:
			// This is not the desired token type.
			v.putBack(tokens)
			return
		}
	}

	// We are at the end-of-file marker.
	return
}

func (v *parser_) formatError(
	ruleName string,
	token TokenLike,
) string {
	// Format the error message.
	var message = fmt.Sprintf(
		"An unexpected token was received by the parser: %v\n",
		ScannerClass().FormatToken(token),
	)
	var line = token.GetLine()
	var lines = sts.Split(v.source_, "\n")

	// Append the source lines with the error in it.
	message += "\033[36m"
	for index := line - 3; index < line; index++ {
		if index > 1 {
			message += fmt.Sprintf("%04d: ", index) + string(lines[index-1]) + "\n"
		}
	}
	message += fmt.Sprintf("%04d: ", line) + string(lines[line-1]) + "\n"

	// Append an arrow pointing to the error.
	message += " \033[32m>>>─"
	var count uint
	for count < token.GetPosition() {
		message += "─"
		count++
	}
	message += "⌃\033[36m\n"

	// Append the following source line for context.
	if line < uint(len(lines)) {
		message += fmt.Sprintf("%04d: ", line+1) + string(lines[line]) + "\n"
	}
	message += "\033[0m\n"
	if uti.IsDefined(ruleName) {
		message += "Was expecting:\n"
		message += fmt.Sprintf(
			"  \033[32m%v: \033[33m%v\033[0m\n\n",
			ruleName,
			v.getDefinition(ruleName),
		)
	}
	return message
}

func (v *parser_) getDefinition(
	ruleName string,
) string {
	var syntax = parserClass().syntax_
	var definition = syntax.GetValue(ruleName)
	return definition
}

func (v *parser_) getNextToken() TokenLike {
	// Check for any read, but unprocessed tokens.
	if !v.next_.IsEmpty() {
		return v.next_.RemoveLast()
	}

	// Read a new token from the token stream.
	var token, ok = v.tokens_.RemoveFirst() // This will wait for a token.
	if !ok {
		// The token channel has been closed.
		return nil
	}

	// Check for an error token.
	if token.GetType() == ErrorToken {
		var message = v.formatError("", token)
		panic(message)
	}

	return token
}

func (v *parser_) putBack(
	tokens fra.Sequential[TokenLike],
) {
	var iterator = tokens.GetIterator()
	for iterator.ToEnd(); iterator.HasPrevious(); {
		var token = iterator.GetPrevious()
		v.next_.AddValue(token)
	}
}

// NOTE:
// This method does nothing but must exist to satisfy the lint check on the
// generated parser code.  The generated code must call this method is some
// cases to make it look that the tokens variable is being used somewhere.
func (v *parser_) remove(
	tokens fra.Sequential[TokenLike],
) {
}

// Instance Structure

type parser_ struct {
	// Declare the instance attributes.
	source_ string                   // The original source code.
	tokens_ fra.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   fra.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Class Structure

type parserClass_ struct {
	// Declare the class constants.
	syntax_ fra.CatalogLike[string, string]
}

// Class Reference

func parserClass() *parserClass_ {
	return parserClassReference_
}

var parserClassReference_ = &parserClass_{
	// Initialize the class constants.
	syntax_: fra.CatalogFromMap[string, string](
		map[string]string{
			"$Syntax":      `LegalNotice comment Rule+ comment Expression+ comment Fragment+`,
			"$LegalNotice": `comment`,
			"$Rule":        `"$" uppercase ":" Definition`,
			"$Definition": `
    LiteralAlternatives
    ExpressionAlternatives
    RuleAlternatives
    TermSequence  ! This must be last since it skips newlines.`,
			"$LiteralAlternatives":    `LiteralValue+`,
			"$LiteralValue":           `newline literal note?`,
			"$ExpressionAlternatives": `ExpressionName+`,
			"$ExpressionName":         `newline lowercase note?`,
			"$RuleAlternatives":       `RuleName+`,
			"$RuleName":               `newline uppercase note?`,
			"$TermSequence":           `RuleTerm+ note?`,
			"$RuleTerm":               `Component Cardinality?  ! The default cardinality is one.`,
			"$Component": `
    literal
    lowercase
    uppercase`,
			"$Cardinality": `
    Constrained
    Quantified`,
			"$Constrained": `
    "?"  ! Zero or one.
    "*"  ! Zero or more.
    "+"  ! One or more.`,
			"$Quantified":          `"{" number Limit? "}"  ! A quantified range of numbers is inclusive.`,
			"$Limit":               `".." number?  ! A quantified range may have no upper limit.`,
			"$Expression":          `"$" lowercase ":" Pattern`,
			"$Fragment":            `"$" allcaps ":" Pattern`,
			"$Pattern":             `Alternatives note?`,
			"$Alternatives":        `Sequence AlternativeSequence*`,
			"$AlternativeSequence": `"|" Sequence`,
			"$Sequence":            `Repetition+`,
			"$Repetition":          `Element Cardinality?  ! The default cardinality is exactly one.`,
			"$Element": `
    Group
    Filter
    Text`,
			"$Group":  `"(" Alternatives ")"`,
			"$Filter": `"~"? "[" Character+ "]"`,
			"$Character": `
    Implicit
    Explicit`,
			"$Implicit": `intrinsic`,
			"$Explicit": `glyph Extent?  ! An explicit range of glyphs is inclusive.`,
			"$Extent":   `".." glyph`,
			"$Text": `
    glyph
    literal
    intrinsic
    lowercase
    allcaps`,
		},
	),
}
