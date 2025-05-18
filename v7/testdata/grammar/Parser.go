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
	col "github.com/craterdog/go-collection-framework/v7"
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
	v.tokens_ = col.Queue[TokenLike]()
	v.next_ = col.Stack[TokenLike]()

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

func (v *parser_) parseAdditionalCharacter() (
	additionalCharacter ast.AdditionalCharacterLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Character rule.
	var character ast.CharacterLike
	character, token, ok = v.parseCharacter()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Character rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AdditionalCharacter", token)
		panic(message)
	}

	// Found a single AdditionalCharacter rule.
	ok = true
	v.remove(tokens)
	additionalCharacter = ast.AdditionalCharacterClass().AdditionalCharacter(character)
	return
}

func (v *parser_) parseAdditionalRepetition() (
	additionalRepetition ast.AdditionalRepetitionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Repetition rule.
	var repetition ast.RepetitionLike
	repetition, token, ok = v.parseRepetition()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Repetition rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AdditionalRepetition", token)
		panic(message)
	}

	// Found a single AdditionalRepetition rule.
	ok = true
	v.remove(tokens)
	additionalRepetition = ast.AdditionalRepetitionClass().AdditionalRepetition(repetition)
	return
}

func (v *parser_) parseAllowedCharacters() (
	allowedCharacters ast.AllowedCharactersLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Character rule.
	var character ast.CharacterLike
	character, token, ok = v.parseCharacter()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Character rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AllowedCharacters", token)
		panic(message)
	}

	// Attempt to parse multiple AdditionalCharacter rules.
	var additionalCharacters = col.List[ast.AdditionalCharacterLike]()
additionalCharactersLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var additionalCharacter ast.AdditionalCharacterLike
		additionalCharacter, token, ok = v.parseAdditionalCharacter()
		if !ok {
			switch {
			case count >= 0:
				break additionalCharactersLoop
			case uti.IsDefined(tokens):
				// This is not multiple AdditionalCharacter rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$AllowedCharacters", token)
				message += "0 or more AdditionalCharacter rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		additionalCharacters.AppendValue(additionalCharacter)
	}

	// Found a single AllowedCharacters rule.
	ok = true
	v.remove(tokens)
	allowedCharacters = ast.AllowedCharactersClass().AllowedCharacters(
		character,
		additionalCharacters,
	)
	return
}

func (v *parser_) parseAlternativeSequence() (
	alternativeSequence ast.AlternativeSequenceLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

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
	var tokens = col.List[TokenLike]()

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
	var alternativeSequences = col.List[ast.AlternativeSequenceLike]()
alternativeSequencesLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var alternativeSequence ast.AlternativeSequenceLike
		alternativeSequence, token, ok = v.parseAlternativeSequence()
		if !ok {
			switch {
			case count >= 0:
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
	// Attempt to parse a single LiteralValueAlternatives Definition.
	var literalValueAlternatives ast.LiteralValueAlternativesLike
	literalValueAlternatives, token, ok = v.parseLiteralValueAlternatives()
	if ok {
		// Found a single LiteralValueAlternatives Definition.
		definition = ast.DefinitionClass().Definition(literalValueAlternatives)
		return
	}

	// Attempt to parse a single TokenNameAlternatives Definition.
	var tokenNameAlternatives ast.TokenNameAlternativesLike
	tokenNameAlternatives, token, ok = v.parseTokenNameAlternatives()
	if ok {
		// Found a single TokenNameAlternatives Definition.
		definition = ast.DefinitionClass().Definition(tokenNameAlternatives)
		return
	}

	// Attempt to parse a single RuleNameAlternatives Definition.
	var ruleNameAlternatives ast.RuleNameAlternativesLike
	ruleNameAlternatives, token, ok = v.parseRuleNameAlternatives()
	if ok {
		// Found a single RuleNameAlternatives Definition.
		definition = ast.DefinitionClass().Definition(ruleNameAlternatives)
		return
	}

	// Attempt to parse a single RuleTermSequence Definition.
	var ruleTermSequence ast.RuleTermSequenceLike
	ruleTermSequence, token, ok = v.parseRuleTermSequence()
	if ok {
		// Found a single RuleTermSequence Definition.
		definition = ast.DefinitionClass().Definition(ruleTermSequence)
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
	var tokens = col.List[TokenLike]()

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
	var tokens = col.List[TokenLike]()

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

func (v *parser_) parseExtent() (
	extent ast.ExtentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

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
	var tokens = col.List[TokenLike]()

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

	// Attempt to parse a single AllowedCharacters rule.
	var allowedCharacters ast.AllowedCharactersLike
	allowedCharacters, token, ok = v.parseAllowedCharacters()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single AllowedCharacters rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Filter", token)
		panic(message)
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
		allowedCharacters,
		delimiter2,
	)
	return
}

func (v *parser_) parseGroup() (
	group ast.GroupLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

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
	var tokens = col.List[TokenLike]()

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

func (v *parser_) parseLimit() (
	limit ast.LimitLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

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

func (v *parser_) parseLiteralValue() (
	literalValue ast.LiteralValueLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

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

func (v *parser_) parseLiteralValueAlternatives() (
	literalValueAlternatives ast.LiteralValueAlternativesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse multiple LiteralValue rules.
	var literalValues = col.List[ast.LiteralValueLike]()
literalValuesLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var literalValue ast.LiteralValueLike
		literalValue, token, ok = v.parseLiteralValue()
		if !ok {
			switch {
			case count >= 1:
				break literalValuesLoop
			case uti.IsDefined(tokens):
				// This is not multiple LiteralValue rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$LiteralValueAlternatives", token)
				message += "1 or more LiteralValue rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		literalValues.AppendValue(literalValue)
	}

	// Found a single LiteralValueAlternatives rule.
	ok = true
	v.remove(tokens)
	literalValueAlternatives = ast.LiteralValueAlternativesClass().LiteralValueAlternatives(literalValues)
	return
}

func (v *parser_) parseNotice() (
	notice ast.NoticeLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

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
			var message = v.formatError("$Notice", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

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
			var message = v.formatError("$Notice", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Notice rule.
	ok = true
	v.remove(tokens)
	notice = ast.NoticeClass().Notice(
		comment,
		newline,
	)
	return
}

func (v *parser_) parsePattern() (
	pattern ast.PatternLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

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
	var tokens = col.List[TokenLike]()

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
	var tokens = col.List[TokenLike]()

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
	var tokens = col.List[TokenLike]()

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

func (v *parser_) parseRuleName() (
	ruleName ast.RuleNameLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

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

func (v *parser_) parseRuleNameAlternatives() (
	ruleNameAlternatives ast.RuleNameAlternativesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse multiple RuleName rules.
	var ruleNames = col.List[ast.RuleNameLike]()
ruleNamesLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var ruleName ast.RuleNameLike
		ruleName, token, ok = v.parseRuleName()
		if !ok {
			switch {
			case count >= 1:
				break ruleNamesLoop
			case uti.IsDefined(tokens):
				// This is not multiple RuleName rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$RuleNameAlternatives", token)
				message += "1 or more RuleName rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		ruleNames.AppendValue(ruleName)
	}

	// Found a single RuleNameAlternatives rule.
	ok = true
	v.remove(tokens)
	ruleNameAlternatives = ast.RuleNameAlternativesClass().RuleNameAlternatives(ruleNames)
	return
}

func (v *parser_) parseRuleTerm() (
	ruleTerm ast.RuleTermLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

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

func (v *parser_) parseRuleTermSequence() (
	ruleTermSequence ast.RuleTermSequenceLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse multiple RuleTerm rules.
	var ruleTerms = col.List[ast.RuleTermLike]()
ruleTermsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var ruleTerm ast.RuleTermLike
		ruleTerm, token, ok = v.parseRuleTerm()
		if !ok {
			switch {
			case count >= 1:
				break ruleTermsLoop
			case uti.IsDefined(tokens):
				// This is not multiple RuleTerm rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$RuleTermSequence", token)
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

	// Found a single RuleTermSequence rule.
	ok = true
	v.remove(tokens)
	ruleTermSequence = ast.RuleTermSequenceClass().RuleTermSequence(
		ruleTerms,
		optionalNote,
	)
	return
}

func (v *parser_) parseSequence() (
	sequence ast.SequenceLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Repetition rule.
	var repetition ast.RepetitionLike
	repetition, token, ok = v.parseRepetition()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Repetition rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Sequence", token)
		panic(message)
	}

	// Attempt to parse multiple AdditionalRepetition rules.
	var additionalRepetitions = col.List[ast.AdditionalRepetitionLike]()
additionalRepetitionsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var additionalRepetition ast.AdditionalRepetitionLike
		additionalRepetition, token, ok = v.parseAdditionalRepetition()
		if !ok {
			switch {
			case count >= 0:
				break additionalRepetitionsLoop
			case uti.IsDefined(tokens):
				// This is not multiple AdditionalRepetition rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Sequence", token)
				message += "0 or more AdditionalRepetition rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		additionalRepetitions.AppendValue(additionalRepetition)
	}

	// Found a single Sequence rule.
	ok = true
	v.remove(tokens)
	sequence = ast.SequenceClass().Sequence(
		repetition,
		additionalRepetitions,
	)
	return
}

func (v *parser_) parseSyntax() (
	syntax ast.SyntaxLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Notice rule.
	var notice ast.NoticeLike
	notice, token, ok = v.parseNotice()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Notice rule.
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
	var rules = col.List[ast.RuleLike]()
rulesLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var rule ast.RuleLike
		rule, token, ok = v.parseRule()
		if !ok {
			switch {
			case count >= 1:
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
	var expressions = col.List[ast.ExpressionLike]()
expressionsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var expression ast.ExpressionLike
		expression, token, ok = v.parseExpression()
		if !ok {
			switch {
			case count >= 1:
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

	// Found a single Syntax rule.
	ok = true
	v.remove(tokens)
	syntax = ast.SyntaxClass().Syntax(
		notice,
		comment1,
		rules,
		comment2,
		expressions,
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

	// Attempt to parse a single lowercase Text.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken)
	if ok {
		// Found a single lowercase Text.
		text = ast.TextClass().Text(lowercase)
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

	// This is not a single Text rule.
	return
}

func (v *parser_) parseTokenName() (
	tokenName ast.TokenNameLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

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
			var message = v.formatError("$TokenName", token)
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
			var message = v.formatError("$TokenName", token)
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

	// Found a single TokenName rule.
	ok = true
	v.remove(tokens)
	tokenName = ast.TokenNameClass().TokenName(
		newline,
		lowercase,
		optionalNote,
	)
	return
}

func (v *parser_) parseTokenNameAlternatives() (
	tokenNameAlternatives ast.TokenNameAlternativesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse multiple TokenName rules.
	var tokenNames = col.List[ast.TokenNameLike]()
tokenNamesLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var tokenName ast.TokenNameLike
		tokenName, token, ok = v.parseTokenName()
		if !ok {
			switch {
			case count >= 1:
				break tokenNamesLoop
			case uti.IsDefined(tokens):
				// This is not multiple TokenName rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$TokenNameAlternatives", token)
				message += "1 or more TokenName rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		tokenNames.AppendValue(tokenName)
	}

	// Found a single TokenNameAlternatives rule.
	ok = true
	v.remove(tokens)
	tokenNameAlternatives = ast.TokenNameAlternativesClass().TokenNameAlternatives(tokenNames)
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
	var tokens = col.List[TokenLike]()
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
	tokens col.Sequential[TokenLike],
) {
	var iterator = tokens.GetIterator()
	for iterator.ToEnd(); iterator.HasPrevious(); {
		var token = iterator.GetPrevious()
		v.next_.AddValue(token)
	}
}

func (v *parser_) remove(
	tokens col.Sequential[TokenLike],
) {
	// NOTE: This method does nothing but must exist to satisfy the lint
	// check on the generated parser code.
}

// Instance Structure

type parser_ struct {
	// Declare the instance attributes.
	source_ string                   // The original source code.
	tokens_ col.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   col.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Class Structure

type parserClass_ struct {
	// Declare the class constants.
	syntax_ col.CatalogLike[string, string]
}

// Class Reference

func parserClass() *parserClass_ {
	return parserClassReference_
}

var parserClassReference_ = &parserClass_{
	// Initialize the class constants.
	syntax_: col.CatalogFromMap[string, string](
		map[string]string{
			"$Syntax": `Notice comment Rule+ comment Expression+`,
			"$Notice": `comment newline`,
			"$Rule":   `"$" uppercase ":" Definition`,
			"$Definition": `
    LiteralValueAlternatives
    TokenNameAlternatives
    RuleNameAlternatives
    RuleTermSequence  ! This must be last since it skips newlines.`,
			"$LiteralValueAlternatives": `LiteralValue+`,
			"$LiteralValue":             `newline literal note?`,
			"$TokenNameAlternatives":    `TokenName+`,
			"$TokenName":                `newline lowercase note?`,
			"$RuleNameAlternatives":     `RuleName+`,
			"$RuleName":                 `newline uppercase note?`,
			"$RuleTermSequence":         `RuleTerm+ note?`,
			"$RuleTerm":                 `Component Cardinality?  ! The default cardinality is one.`,
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
			"$Quantified":           `"{" number Limit? "}"  ! A quantified range of numbers is inclusive.`,
			"$Limit":                `".." number?  ! A quantified range may have no upper limit.`,
			"$Expression":           `"$" lowercase ":" Pattern`,
			"$Pattern":              `Alternatives note?`,
			"$Alternatives":         `Sequence AlternativeSequence*`,
			"$AlternativeSequence":  `"|" Sequence`,
			"$Sequence":             `Repetition AdditionalRepetition*`,
			"$AdditionalRepetition": `Repetition`,
			"$Repetition":           `Element Cardinality?  ! The default cardinality is one.`,
			"$Element": `
    Group
    Filter
    Text`,
			"$Group":               `"(" Alternatives ")"`,
			"$Filter":              `"~"? "[" AllowedCharacters "]"`,
			"$AllowedCharacters":   `Character AdditionalCharacter*`,
			"$AdditionalCharacter": `Character`,
			"$Character": `
    Implicit
    Explicit`,
			"$Implicit": `intrinsic`,
			"$Explicit": `glyph Extent?  ! An explicit range of glyphs is inclusive.`,
			"$Extent":   `".." glyph`,
			"$Text": `
    glyph
    literal
    lowercase
    intrinsic`,
		},
	),
}
