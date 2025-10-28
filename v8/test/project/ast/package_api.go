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
│         This "package_api.go" file was automatically generated using:        │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘

Package "ast" provides the abstract syntax tree (AST) classes for this module
based on the "syntax.cdsn" grammar for the module.  Each AST class manages the
attributes associated with its corresponding rule definition found in the
grammar.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-example-syntax/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package ast

import (
	fra "github.com/craterdog/go-essential-composites/v8"
)

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
AdditionalClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete additional-like class.
*/
type AdditionalClassLike interface {
	// Constructor Methods
	Additional(
		delimiter string,
		component ComponentLike,
	) AdditionalLike
}

/*
ComponentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete component-like class.
*/
type ComponentClassLike interface {
	// Constructor Methods
	Component(
		any_ any,
	) ComponentLike
}

/*
DocumentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete document-like class.
*/
type DocumentClassLike interface {
	// Constructor Methods
	Document(
		components fra.Sequential[ComponentLike],
	) DocumentLike
}

/*
IntrinsicClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete intrinsic-like class.
*/
type IntrinsicClassLike interface {
	// Constructor Methods
	Intrinsic(
		any_ any,
	) IntrinsicLike
}

/*
KeywordClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete keyword-like class.
*/
type KeywordClassLike interface {
	// Constructor Methods
	Keyword(
		any_ any,
	) KeywordLike
}

/*
ListClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete list-like class.
*/
type ListClassLike interface {
	// Constructor Methods
	List(
		delimiter1 string,
		component ComponentLike,
		additionals fra.Sequential[AdditionalLike],
		delimiter2 string,
	) ListLike
}

// INSTANCE DECLARATIONS

/*
AdditionalLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete additional-like class.
*/
type AdditionalLike interface {
	// Principal Methods
	GetClass() AdditionalClassLike

	// Attribute Methods
	GetDelimiter() string
	GetComponent() ComponentLike
}

/*
ComponentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete component-like class.
*/
type ComponentLike interface {
	// Principal Methods
	GetClass() ComponentClassLike

	// Attribute Methods
	GetAny() any
}

/*
DocumentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete document-like class.
*/
type DocumentLike interface {
	// Principal Methods
	GetClass() DocumentClassLike

	// Attribute Methods
	GetComponents() fra.Sequential[ComponentLike]
}

/*
IntrinsicLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete intrinsic-like class.
*/
type IntrinsicLike interface {
	// Principal Methods
	GetClass() IntrinsicClassLike

	// Attribute Methods
	GetAny() any
}

/*
KeywordLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete keyword-like class.
*/
type KeywordLike interface {
	// Principal Methods
	GetClass() KeywordClassLike

	// Attribute Methods
	GetAny() any
}

/*
ListLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete list-like class.
*/
type ListLike interface {
	// Principal Methods
	GetClass() ListClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetComponent() ComponentLike
	GetAdditionals() fra.Sequential[AdditionalLike]
	GetDelimiter2() string
}

// ASPECT DECLARATIONS
