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

package synthesizer

import (
	ana "github.com/craterdog/go-code-generation/v6/analyzer"
	not "github.com/craterdog/go-syntax-notation/v6"
)

// CLASS INTERFACE

// Access Function

func TokenSynthesizerClass() TokenSynthesizerClassLike {
	return tokenSynthesizerClass()
}

// Constructor Methods

func (c *tokenSynthesizerClass_) TokenSynthesizer(
	syntax not.SyntaxLike,
) TokenSynthesizerLike {
	var instance = &tokenSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzerClass().SyntaxAnalyzer(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *tokenSynthesizer_) GetClass() TokenSynthesizerClassLike {
	return tokenSynthesizerClass()
}

// TemplateDriven Methods

func (v *tokenSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *tokenSynthesizer_) CreateWarningMessage() string {
	var class = tokenSynthesizerClass()
	var warningMessage = class.warningMessage_
	return warningMessage
}

func (v *tokenSynthesizer_) CreateImportedPackages() string {
	var class = tokenSynthesizerClass()
	var importedPackages = class.importedPackages_
	return importedPackages
}

func (v *tokenSynthesizer_) CreateAccessFunction() string {
	var class = tokenSynthesizerClass()
	var accessFunction = class.accessFunction_
	return accessFunction
}

func (v *tokenSynthesizer_) CreateConstantMethods() string {
	var constantMethods string
	return constantMethods
}

func (v *tokenSynthesizer_) CreateConstructorMethods() string {
	var class = tokenSynthesizerClass()
	var constructorMethods = class.constructorMethods_
	return constructorMethods
}

func (v *tokenSynthesizer_) CreateFunctionMethods() string {
	var functionMethods string
	return functionMethods
}

func (v *tokenSynthesizer_) CreatePrincipalMethods() string {
	var class = tokenSynthesizerClass()
	var principalMethods = class.principalMethods_
	return principalMethods
}

func (v *tokenSynthesizer_) CreateAttributeMethods() string {
	var class = tokenSynthesizerClass()
	var attributeMethods = class.attributeMethods_
	return attributeMethods
}

func (v *tokenSynthesizer_) CreateAspectMethods() string {
	var aspectMethods string
	return aspectMethods
}

func (v *tokenSynthesizer_) CreatePrivateMethods() string {
	var class = tokenSynthesizerClass()
	var privateMethods = class.privateMethods_
	return privateMethods
}

func (v *tokenSynthesizer_) CreateInstanceStructure() string {
	var class = tokenSynthesizerClass()
	var instanceStructure = class.instanceStructure_
	return instanceStructure
}

func (v *tokenSynthesizer_) CreateClassStructure() string {
	var class = tokenSynthesizerClass()
	var classStructure = class.classStructure_
	return classStructure
}

func (v *tokenSynthesizer_) CreateClass() string {
	var class = tokenSynthesizerClass()
	var classReference = class.classReference_
	return classReference
}

func (v *tokenSynthesizer_) PerformGlobalUpdates(
	existing string,
	generated string,
) string {
	return generated
}

// PROTECTED INTERFACE

// Instance Structure

type tokenSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.SyntaxAnalyzerLike
}

// Class Structure

type tokenSynthesizerClass_ struct {
	// Declare the class constants.
	warningMessage_     string
	importedPackages_   string
	accessFunction_     string
	constructorMethods_ string
	principalMethods_   string
	attributeMethods_   string
	privateMethods_     string
	instanceStructure_  string
	classStructure_     string
	classReference_     string
}

// Class Reference

func tokenSynthesizerClass() *tokenSynthesizerClass_ {
	return tokenSynthesizerClassReference_
}

var tokenSynthesizerClassReference_ = &tokenSynthesizerClass_{
	// Initialize the class constants.
	warningMessage_: `
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
`,

	importedPackages_: `
	uti "github.com/craterdog/go-missing-utilities/v2"
`,

	accessFunction_: `
// Access Function

func TokenClass() TokenClassLike {
	return tokenClass()
}
`,

	constructorMethods_: `
// Constructor Methods

func (c *tokenClass_) Token(
	line uint,
	position uint,
	type_ TokenType,
	value string,
) TokenLike {
	if uti.IsUndefined(line) {
		panic("The \"line\" attribute is required by this class.")
	}
	if uti.IsUndefined(position) {
		panic("The \"position\" attribute is required by this class.")
	}
	if uti.IsUndefined(type_) {
		panic("The \"type\" attribute is required by this class.")
	}
	if uti.IsUndefined(value) {
		panic("The \"value\" attribute is required by this class.")
	}
	var instance = &token_{
		// Initialize the instance attributes.
		line_:     line,
		position_: position,
		type_:     type_,
		value_:    value,
	}
	return instance
}
`,

	principalMethods_: `
// Principal Methods

func (v *token_) GetClass() TokenClassLike {
	return tokenClass()
}
`,

	attributeMethods_: `
// Attribute Methods

func (v *token_) GetLine() uint {
	return v.line_
}

func (v *token_) GetPosition() uint {
	return v.position_
}

func (v *token_) GetType() TokenType {
	return v.type_
}

func (v *token_) GetValue() string {
	return v.value_
}
`,

	privateMethods_: `
// Private Methods
`,

	instanceStructure_: `
// Instance Structure

type token_ struct {
	// Declare the instance attributes.
	line_     uint
	position_ uint
	type_     TokenType
	value_    string
}
`,

	classStructure_: `
// Class Structure

type tokenClass_ struct {
	// Declare the class constants.
}
`,

	classReference_: `
// Class Reference

func tokenClass() *tokenClass_ {
	return tokenClassReference_
}

var tokenClassReference_ = &tokenClass_{
	// Initialize the class constants.
}
`,
}
