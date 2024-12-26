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

package synthesizer

import (
	ana "github.com/craterdog/go-code-generation/v5/analyzer"
	col "github.com/craterdog/go-collection-framework/v5"
	abs "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	not "github.com/craterdog/go-syntax-notation/v5"
	reg "regexp"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ParserSynthesizerClass() ParserSynthesizerClassLike {
	return parserSynthesizerClassReference()
}

// Constructor Methods

func (c *parserSynthesizerClass_) ParserSynthesizer(
	syntax not.SyntaxLike,
) ParserSynthesizerLike {
	var instance = &parserSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzerClass().SyntaxAnalyzer(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *parserSynthesizer_) GetClass() ParserSynthesizerClassLike {
	return parserSynthesizerClassReference()
}

// TemplateDriven Methods

func (v *parserSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *parserSynthesizer_) CreateWarningMessage() string {
	var class = parserSynthesizerClassReference()
	var warningMessage = class.warningMessage_
	return warningMessage
}

func (v *parserSynthesizer_) CreateAccessFunction() string {
	var class = parserSynthesizerClassReference()
	var accessFunction = class.accessFunction_
	return accessFunction
}

func (v *parserSynthesizer_) CreateConstantMethods() string {
	var constantMethods string
	return constantMethods
}

func (v *parserSynthesizer_) CreateConstructorMethods() string {
	var class = parserSynthesizerClassReference()
	var constructorMethods = class.constructorMethods_
	return constructorMethods
}

func (v *parserSynthesizer_) CreateFunctionMethods() string {
	var functionMethods string
	return functionMethods
}

func (v *parserSynthesizer_) CreatePrincipalMethods() string {
	var class = parserSynthesizerClassReference()
	var principalMethods = class.principalMethods_
	return principalMethods
}

func (v *parserSynthesizer_) CreateAttributeMethods() string {
	var attributeMethods string
	return attributeMethods
}

func (v *parserSynthesizer_) CreateAspectMethods() string {
	var aspectMethods string
	return aspectMethods
}

func (v *parserSynthesizer_) CreatePrivateMethods() string {
	var parseMethods = v.createParseMethods()
	var class = parserSynthesizerClassReference()
	var privateMethods = class.privateMethods_
	privateMethods = uti.ReplaceAll(
		privateMethods,
		"parseMethods",
		parseMethods,
	)
	return privateMethods
}

func (v *parserSynthesizer_) CreateInstanceStructure() string {
	var class = parserSynthesizerClassReference()
	var instanceStructure = class.instanceStructure_
	return instanceStructure
}

func (v *parserSynthesizer_) CreateClassStructure() string {
	var class = parserSynthesizerClassReference()
	var classStructure = class.classStructure_
	return classStructure
}

func (v *parserSynthesizer_) CreateClassReference() string {
	var class = parserSynthesizerClassReference()
	var classReference = class.classReference_
	return classReference
}

func (v *parserSynthesizer_) PerformGlobalUpdates(
	moduleName string,
	packageName string,
	className string,
	existing string,
	generated string,
) string {
	var syntaxMap = v.analyzer_.GetSyntaxMap()
	generated = uti.ReplaceAll(
		generated,
		"syntaxMap",
		syntaxMap,
	)
	var syntaxName = v.analyzer_.GetSyntaxName()
	generated = uti.ReplaceAll(
		generated,
		"syntaxName",
		syntaxName,
	)
	generated = v.updateImportedPackages(moduleName, existing, generated)
	return generated
}

// PROTECTED INTERFACE

// Private Methods

func (v *parserSynthesizer_) createArguments(
	ruleName string,
) (
	arguments string,
) {
	var variableNames = v.analyzer_.GetVariables(ruleName).GetIterator()
	if variableNames.IsEmpty() {
		return arguments
	}

	// Define the first argument.
	if variableNames.GetSize() > 1 {
		// Use the multiline argument style.
		arguments += "\n\t\t"
	}
	var argument = variableNames.GetNext()
	arguments += uti.ReplaceAll(
		"<argument_>",
		"argument",
		argument,
	)

	// Define any additional arguments.
	for variableNames.HasNext() {
		arguments += ",\n\t\t"
		argument = variableNames.GetNext()
		arguments += uti.ReplaceAll(
			"<argument_>",
			"argument",
			argument,
		)
	}
	if variableNames.GetSize() > 1 {
		// Use the multiline argument style.
		arguments += ",\n\t"
	}

	return arguments
}

func (v *parserSynthesizer_) createCardinality(
	cardinality not.CardinalityLike,
	optionalCardinality string,
	requiredCardinality string,
	repeatedCardinality string,
) string {
	var unlimited = "mat.MaxInt"
	var first string
	var last = unlimited
	var actualCardinality = requiredCardinality
	if uti.IsUndefined(cardinality) {
		return actualCardinality
	}
	switch actual := cardinality.GetAny().(type) {
	case not.ConstrainedLike:
		actualCardinality = repeatedCardinality
		switch actual.GetAny().(string) {
		case "?":
			// This is the "{0..1}" case.
			first = "0"
			last = "1"
			actualCardinality = optionalCardinality
		case "*":
			// This is the "{0..}" case.
			first = "0"
		case "+":
			// This is the "{1..}" case.
			first = "1"
		}
	case not.QuantifiedLike:
		first = actual.GetNumber()
		var limit = actual.GetOptionalLimit()
		if uti.IsUndefined(limit) {
			// This is the "{m}" case.
			last = first
		} else {
			last = limit.GetOptionalNumber()
			if uti.IsUndefined(last) {
				// This is the "{m..}" case.
				last = unlimited
			}
			// This is the "{m..n}" case.
		}
	}
	actualCardinality = uti.ReplaceAll(
		actualCardinality,
		"first",
		first,
	)
	actualCardinality = uti.ReplaceAll(
		actualCardinality,
		"last",
		last,
	)
	return actualCardinality
}

func (v *parserSynthesizer_) createDelimiterStep() string {
	var class = parserSynthesizerClassReference()
	var delimiterStep = class.parseDelimiterStep_
	return delimiterStep
}

func (v *parserSynthesizer_) createImportedPath(
	packageAcronym string,
	packagePath string,
) string {
	var class = parserSynthesizerClassReference()
	var importedPath = class.importedPath_
	importedPath = uti.ReplaceAll(
		importedPath,
		"packageAcronym",
		packageAcronym,
	)
	importedPath = uti.ReplaceAll(
		importedPath,
		"packagePath",
		packagePath,
	)
	return importedPath
}

func (v *parserSynthesizer_) createInlineImplementation(
	ruleName string,
) string {
	var class = parserSynthesizerClassReference()
	var implementation = class.declarationStep_
	var terms = v.analyzer_.GetTerms(ruleName).GetIterator()
	var variables = v.analyzer_.GetVariables(ruleName).GetIterator()
	for terms.HasNext() {
		var term = terms.GetNext()
		var parseStep string
		switch actual := term.GetAny().(type) {
		case not.ReferenceLike:
			var cardinality = actual.GetOptionalCardinality()
			var variableName = variables.GetNext()
			var identifier = actual.GetIdentifier().GetAny().(string)
			switch {
			case not.MatchesType(identifier, not.LowercaseToken):
				parseStep = v.createTokenStep(cardinality)
			case not.MatchesType(identifier, not.UppercaseToken):
				parseStep = v.createRuleStep(cardinality)
			}
			parseStep = uti.ReplaceAll(
				parseStep,
				"variableName",
				variableName,
			)
			parseStep = uti.ReplaceAll(
				parseStep,
				"identifier",
				identifier,
			)
		case string:
			var delimiter = actual
			parseStep = v.createDelimiterStep()
			parseStep = uti.ReplaceAll(
				parseStep,
				"delimiter",
				delimiter,
			)
		}
		implementation += parseStep
	}
	var foundStep = class.parseFoundStep_
	var arguments = v.createArguments(ruleName)
	implementation += uti.ReplaceAll(
		foundStep,
		"arguments",
		arguments,
	)
	return implementation
}

func (v *parserSynthesizer_) createMultilineImplementation(
	ruleName string,
) string {
	var implementation string
	var class = parserSynthesizerClassReference()
	var identifiers = v.analyzer_.GetIdentifiers(ruleName).GetIterator()
	for identifiers.HasNext() {
		var identifier = identifiers.GetNext().GetAny().(string)
		var parseCase string
		switch {
		case not.MatchesType(identifier, not.LowercaseToken):
			parseCase = class.parseTokenCase_
		case not.MatchesType(identifier, not.UppercaseToken):
			parseCase = class.parseRuleCase_
		}
		implementation += uti.ReplaceAll(
			parseCase,
			"identifier",
			identifier,
		)
	}
	implementation += class.parseDefaultCase_
	return implementation
}

func (v *parserSynthesizer_) createParseMethod(
	ruleName string,
) string {
	var methodImplementation string
	var references = v.analyzer_.GetReferences(ruleName)
	if uti.IsDefined(references) {
		methodImplementation = v.createInlineImplementation(ruleName)
	} else {
		methodImplementation = v.createMultilineImplementation(ruleName)
	}
	var class = parserSynthesizerClassReference()
	var parseMethod = class.parseMethod_
	parseMethod = uti.ReplaceAll(
		parseMethod,
		"methodImplementation",
		methodImplementation,
	)
	parseMethod = uti.ReplaceAll(
		parseMethod,
		"ruleName",
		ruleName,
	)
	return parseMethod
}

func (v *parserSynthesizer_) createParseMethods() string {
	var parseMethods string
	var ruleNames = v.analyzer_.GetRuleNames().GetIterator()
	for ruleNames.HasNext() {
		var ruleName = ruleNames.GetNext()
		parseMethods += v.createParseMethod(ruleName)
	}
	return parseMethods
}

func (v *parserSynthesizer_) createRuleStep(
	cardinality not.CardinalityLike,
) string {
	var class = parserSynthesizerClassReference()
	var optionalStep = class.parseOptionalRuleStep_
	var requiredStep = class.parseRequiredRuleStep_
	var repeatedStep = class.parseRepeatedRuleStep_
	var ruleStep = v.createCardinality(
		cardinality,
		optionalStep,
		requiredStep,
		repeatedStep,
	)
	return ruleStep
}

func (v *parserSynthesizer_) createTokenStep(
	cardinality not.CardinalityLike,
) string {
	var class = parserSynthesizerClassReference()
	var optionalStep = class.parseOptionalTokenStep_
	var requiredStep = class.parseRequiredTokenStep_
	var repeatedStep = class.parseRepeatedTokenStep_
	var tokenStep = v.createCardinality(
		cardinality,
		optionalStep,
		requiredStep,
		repeatedStep,
	)
	return tokenStep
}

func (v *parserSynthesizer_) extractImportedPackages(
	source string,
) (
	packages abs.CatalogLike[string, string],
) {
	packages = col.Catalog[string, string]()
	var lower_ = `\p{Ll}`
	var digit_ = `\p{Nd}`
	var acronym_ = `(` + lower_ + `(?:` + lower_ + `|` + digit_ + `){2})`
	var white_ = `[ \t\r\n]*`
	var path_ = `("[^"]+")`
	var pattern = `import \((?:.|\r?\n)*?\)`
	var matcher = reg.MustCompile(pattern)
	var imports = matcher.FindString(source)
	var lines = sts.Split(imports, "\n")
	var count = len(lines)
	if count > 2 {
		pattern = acronym_ + white_ + path_
		matcher = reg.MustCompile(pattern)
		lines = lines[1 : count-1]
		for _, line := range lines {
			var matches = matcher.FindStringSubmatch(line)
			var packageAcronym = matches[1]
			var packagePath = matches[2]
			packages.SetValue(packageAcronym, packagePath)
		}
	}
	return
}

func (v *parserSynthesizer_) updateImportedPackages(
	moduleName string,
	existing string,
	generated string,
) string {
	// Seed the imported packages with the most common ones.
	var imports = abs.CatalogClass[string, string]().CatalogFromMap(
		map[string]string{
			"fmt": `"fmt"`,
			"ast": `"<ModuleName>/ast"`,
			"col": `"github.com/craterdog/go-collection-framework/v5"`,
			"abs": `"github.com/craterdog/go-collection-framework/v5/collection"`,
			"uti": `"github.com/craterdog/go-missing-utilities/v2"`,
			"mat": `"math"`,
			"sts": `"strings"`,
		},
	)

	// Add in the imported packages from the existing code.
	imports = abs.CatalogClass[string, string]().Merge(
		imports, v.extractImportedPackages(existing),
	)
	imports.SortValuesWithRanker(
		func(first, second abs.AssociationLike[string, string]) col.Rank {
			var firstValue = first.GetValue()
			var secondValue = second.GetValue()
			switch {
			case firstValue < secondValue:
				return col.LesserRank
			case firstValue > secondValue:
				return col.GreaterRank
			default:
				return col.EqualRank
			}
		},
	)

	// Create imported package statements for each imported package.
	var importedPackages string
	var packages = imports.GetIterator()
	for packages.HasNext() {
		var association = packages.GetNext()
		var packageAcronym = association.GetKey()
		var packagePath = association.GetValue()
		if sts.Contains(generated, packageAcronym+".") {
			// Only import packages that are actually used in the generated code.
			importedPackages += v.createImportedPath(
				packageAcronym,
				packagePath,
			)
		}
	}
	if uti.IsDefined(importedPackages) {
		importedPackages += "\n"
	}

	// Insert the imported packages into the generated code.
	generated = uti.ReplaceAll(
		generated,
		"importedPackages",
		importedPackages,
	)
	return generated
}

// Instance Structure

type parserSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.SyntaxAnalyzerLike
}

// Class Structure

type parserSynthesizerClass_ struct {
	// Declare the class constants.
	warningMessage_         string
	importedPath_           string
	accessFunction_         string
	constructorMethods_     string
	parseMethod_            string
	parseDelimiterStep_     string
	declarationStep_        string
	parseRequiredTokenStep_ string
	parseOptionalTokenStep_ string
	parseRepeatedTokenStep_ string
	parseRequiredRuleStep_  string
	parseOptionalRuleStep_  string
	parseRepeatedRuleStep_  string
	parseFoundStep_         string
	parseTokenCase_         string
	parseRuleCase_          string
	parseDefaultCase_       string
	principalMethods_       string
	privateMethods_         string
	possibleDelimiter_      string
	requiredDelimiter_      string
	instanceStructure_      string
	classStructure_         string
	classReference_         string
}

// Class Reference

func parserSynthesizerClassReference() *parserSynthesizerClass_ {
	return parserSynthesizerClassReference_
}

var parserSynthesizerClassReference_ = &parserSynthesizerClass_{
	// Initialize the class constants.
	warningMessage_: `
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
`,

	importedPath_: `
	<~packageAcronym> <packagePath>`,

	accessFunction_: `
// Access Function

func ParserClass() ParserClassLike {
	return parserClassReference()
}
`,

	constructorMethods_: `
// Constructor Methods

func (c *parserClass_) Parser() ParserLike {
	var instance = &parser_{
		// Initialize the instance attributes.
	}
	return instance
}
`,

	parseMethod_: `
func (v *parser_) parse<~RuleName>() (
	<ruleName_> ast.<~RuleName>Like,
	token TokenLike,
	ok bool,
) {<MethodImplementation>}
`,

	parseDelimiterStep_: `
	// Attempt to parse a single <delimiter> delimiter.
	_, token, ok = v.parseDelimiter(<delimiter>)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single <~RuleName> rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$<~RuleName>", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}
`,

	declarationStep_: `
	var tokens = col.List[TokenLike]()
`,

	parseRequiredTokenStep_: `
	// Attempt to parse a single <~identifier> token.
	var <variableName_> string
	<variableName_>, token, ok = v.parseToken(<~Identifier>Token)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single <~RuleName> rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$<~RuleName>", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}
`,

	parseOptionalTokenStep_: `
	// Attempt to parse an optional <~identifier> token.
	var <variableName_> string
	<variableName_>, token, ok = v.parseToken(<~Identifier>Token)
	if ok && uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}
`,

	parseRepeatedTokenStep_: `
	// Attempt to parse multiple <~identifier> tokens.
	var <variableName_> = col.List[string]()
<~variableName>Loop:
	for count := 0; count < <last>; count++ {
		var <identifier_> string
		<identifier_>, token, ok = v.parseToken(<~Identifier>Token)
		if !ok {
			switch {
			case count >= <first>:
				break <~variableName>Loop
			case uti.IsDefined(tokens):
				// This is not multiple <~identifier> tokens.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$<~RuleName>", token)
				message += "<first> or more <~identifier> tokens are required."
				panic(message)
			}
		}
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
		<variableName_>.AppendValue(<identifier_>)
	}
`,

	parseRequiredRuleStep_: `
	// Attempt to parse a single <~Identifier> rule.
	var <variableName_> ast.<~Identifier>Like
	<variableName_>, token, ok = v.parse<~Identifier>()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single <~RuleName> rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$<~RuleName>", token)
		panic(message)
	}
`,

	parseOptionalRuleStep_: `
	// Attempt to parse an optional <~Identifier> rule.
	var <variableName_> ast.<~Identifier>Like
	<variableName_>, _, ok = v.parse<~Identifier>()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}
`,

	parseRepeatedRuleStep_: `
	// Attempt to parse multiple <~Identifier> rules.
	var <variableName_> = col.List[ast.<~Identifier>Like]()
<~variableName>Loop:
	for count := 0; count < <last>; count++ {
		var <identifier_> ast.<~Identifier>Like
		<identifier_>, token, ok = v.parse<~Identifier>()
		if !ok {
			switch {
			case count >= <first>:
				break <~variableName>Loop
			case uti.IsDefined(tokens):
				// This is not multiple <~Identifier> rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$<~RuleName>", token)
				message += "<first> or more <~Identifier> rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		<variableName_>.AppendValue(<identifier_>)
	}
`,

	parseFoundStep_: `
	// Found a single <~RuleName> rule.
	ok = true
	v.remove(tokens)
	<ruleName_> = ast.<~RuleName>Class().<~RuleName>(<Arguments>)
	return
`,

	parseTokenCase_: `
	// Attempt to parse a single <~identifier> <~RuleName>.
	var <identifier_> string
	<identifier_>, token, ok = v.parseToken(<~Identifier>Token)
	if ok {
		// Found a single <~identifier> <~RuleName>.
		<ruleName_> = ast.<~RuleName>Class().<~RuleName>(<identifier_>)
		return
	}
`,

	parseRuleCase_: `
	// Attempt to parse a single <~Identifier> <~RuleName>.
	var <identifier_> ast.<~Identifier>Like
	<identifier_>, token, ok = v.parse<~Identifier>()
	if ok {
		// Found a single <~Identifier> <~RuleName>.
		<ruleName_> = ast.<~RuleName>Class().<~RuleName>(<identifier_>)
		return
	}
`,

	parseDefaultCase_: `
	// This is not a single <~RuleName> rule.
	return
`,

	principalMethods_: `
// Principal Methods

func (v *parser_) GetClass() ParserClassLike {
	return parserClassReference()
}

func (v *parser_) ParseSource(
	source string,
) ast.<~SyntaxName>Like {
	v.source_ = sts.ReplaceAll(source, "\t", "    ")
	v.tokens_ = col.Queue[TokenLike]()
	v.next_ = col.Stack[TokenLike]()

	// The scanner runs in a separate Go routine.
	ScannerClass().Scanner(v.source_, v.tokens_)

	// Attempt to parse the <~syntaxName>.
	var <syntaxName_>, token, ok = v.parse<~SyntaxName>()
	if !ok {
		var message = v.formatError("$<~SyntaxName>", token)
		panic(message)
	}
	return <syntaxName_>
}
`,

	privateMethods_: `
// Private Methods
<ParseMethods>
func (v *parser_) parseDelimiter(
	expectedValue string,
) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single delimiter.
	value, token, ok = v.parseToken(DelimiterToken)
	if ok {
		if value == expectedValue {
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
	var syntax = parserClassReference().syntax_
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
	tokens abs.Sequential[TokenLike],
) {
	var iterator = tokens.GetIterator()
	for iterator.ToEnd(); iterator.HasPrevious(); {
		var token = iterator.GetPrevious()
		v.next_.AddValue(token)
	}
}

func (v *parser_) remove(
	tokens abs.Sequential[TokenLike],
) {
	// NOTE: This method does nothing but must exist to satisfy the lint
	// check on the generated parser code.
}
`,

	possibleDelimiter_: `
	// Attempt to parse a single "<delimiter>" delimiter.
	_, token, ok = v.parseDelimiter("<delimiter>")
	if !ok {
		// This is not a single <~RuleName> rule.
		return <ruleName_>, token, false
	}
`,

	requiredDelimiter_: `
	// Attempt to parse a single "<delimiter>" delimiter.
	_, token, ok = v.parseDelimiter("<delimiter>")
	if !ok {
		// Encountered a syntax error.
		var message = v.formatError("$<~RuleName>", token)
		panic(message)
	}
`,

	instanceStructure_: `
// Instance Structure

type parser_ struct {
	// Declare the instance attributes.
	source_ string                   // The original source code.
	tokens_ abs.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   abs.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}
`,

	classStructure_: `
// Class Structure

type parserClass_ struct {
	// Declare the class constants.
	syntax_ abs.CatalogLike[string, string]
}
`,

	classReference_: `
// Class Reference

func parserClassReference() *parserClass_ {
	return parserClassReference_
}

var parserClassReference_ = &parserClass_{
	// Initialize the class constants.
	syntax_: col.CatalogFromMap[string, string](
		map[string]string{<SyntaxMap>
		},
	),
}
`,
}
