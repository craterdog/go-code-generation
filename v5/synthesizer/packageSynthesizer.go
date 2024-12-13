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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func PackageSynthesizerClass() PackageSynthesizerClassLike {
	return packageSynthesizerClassReference()
}

// Constructor Methods

func (c *packageSynthesizerClass_) Make() PackageSynthesizerLike {
	var instance = &packageSynthesizer_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *packageSynthesizer_) GetClass() PackageSynthesizerClassLike {
	return packageSynthesizerClassReference()
}

// TemplateDriven Methods

func (v *packageSynthesizer_) CreateLegalNotice() string {
	var class = packageSynthesizerClassReference()
	var legalNotice = class.legalNotice_
	return legalNotice
}

func (v *packageSynthesizer_) CreateWarningMessage() string {
	var class = packageSynthesizerClassReference()
	var warningMessage = class.warningMessage_
	return warningMessage
}

func (v *packageSynthesizer_) CreatePackageDescription() string {
	var class = packageSynthesizerClassReference()
	var packageDescription = class.packageDescription_
	return packageDescription
}

func (v *packageSynthesizer_) CreateTypeDeclarations() string {
	var class = packageSynthesizerClassReference()
	var typeDeclarations = class.typeDeclarations_
	return typeDeclarations
}

func (v *packageSynthesizer_) CreateFunctionalDeclarations() string {
	var class = packageSynthesizerClassReference()
	var functionalDeclarations = class.functionalDeclarations_
	return functionalDeclarations
}

func (v *packageSynthesizer_) CreateClassDeclarations() string {
	var class = packageSynthesizerClassReference()
	var classDeclarations = class.classDeclarations_
	return classDeclarations
}

func (v *packageSynthesizer_) CreateInstanceDeclarations() string {
	var class = packageSynthesizerClassReference()
	var instanceDeclarations = class.instanceDeclarations_
	return instanceDeclarations
}

func (v *packageSynthesizer_) CreateAspectDeclarations() string {
	var class = packageSynthesizerClassReference()
	var aspectDeclarations = class.aspectDeclarations_
	return aspectDeclarations
}

func (v *packageSynthesizer_) PerformGlobalUpdates(
	source string,
) string {
	var class = packageSynthesizerClassReference()
	var importedPackages = class.importedPackages_
	source = uti.ReplaceAll(
		source,
		"importedPackages",
		importedPackages,
	)
	return source
}

// PROTECTED INTERFACE

// Instance Structure

type packageSynthesizer_ struct {
	// Declare the instance attributes.
}

// Class Structure

type packageSynthesizerClass_ struct {
	// Declare the class constants.
	legalNotice_            string
	warningMessage_         string
	packageDescription_     string
	importedPackages_       string
	typeDeclarations_       string
	functionalDeclarations_ string
	classDeclarations_      string
	instanceDeclarations_   string
	aspectDeclarations_     string
}

// Class Reference

func packageSynthesizerClassReference() *packageSynthesizerClass_ {
	return packageSynthesizerClassReference_
}

var packageSynthesizerClassReference_ = &packageSynthesizerClass_{
	// Initialize the class constants.
	legalNotice_: `/*
................................................................................
.                  Copyright (c) 2024.  All Rights Reserved.                   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/
`,

	warningMessage_: ``,

	packageDescription_: `
Package "example" provides...`,

	importedPackages_: ``,

	typeDeclarations_: `
/*
Identifier is a constrained type representing a generic identifier that can be
used to label things.  An Identifier cannot contain any whitespace and should
match the following regular expression: [a-zA-Z][a-zA-Z0-9]*
*/
type Identifier string

/*
Ordinal is a constrained type representing an ordinal number in the range
[1..∞).  The value 0 is used to represent infinity.
*/
type Ordinal uint64

/*
Rank is a constrained type representing the possible relative ranking of two
values.
*/
type Rank uint8

const (
	LesserRank Rank = iota
	EqualRank
	GreaterRank
)

/*
Units is a constrained type representing the possible units for an angle.
*/
type Units uint8

const (
	Degrees Units = iota
	Radians
	Gradians
)
`,

	functionalDeclarations_: `
/*
RankingFunction[V any] is a functional type that declares the signature for any
function that can determine the relative ranking of two values.
*/
type RankingFunction[V any] func(
	first V,
	second V,
) Rank
`,

	classDeclarations_: `
/*
AngleClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
angle-like class.
*/
type AngleClassLike interface {
	// Constructor Methods
	Angle(
		intrinsic float64,
	) AngleLike
	AngleFromString(
		value string,
	) AngleLike

	// Constant Methods
	Pi() AngleLike
	Tau() AngleLike

	// Function Methods
	Sine(
		angle AngleLike,
	) float64
	Cosine(
		angle AngleLike,
	) float64
	Tangent(
		angle AngleLike,
	) float64
}

/*
ArrayClassLike[V any] is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete array-like class.
*/
type ArrayClassLike[V any] interface {
	// Constructor Methods
	Array(
		intrinsic []V,
	) ArrayLike[V]
	ArrayWithSize(
		size Ordinal,
	) ArrayLike[V]
	ArrayFromSequence(
		values Sequential[V],
	) ArrayLike[V]

	// Function Methods
	Merge(
		first ArrayLike[V],
		second ArrayLike[V],
	) ArrayLike[V]
}

/*
AssociationClassLike[K comparable, V any] is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete association-like class.
*/
type AssociationClassLike[K comparable, V any] interface {
	// Constructor Methods
	Association(
		key K,
		value V,
	) AssociationLike[K, V]
}

/*
CatalogClassLike[V any] is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete catalog-like class.

The following functions are supported:

Extract() returns a new catalog containing only the associations that are in the
specified catalog that have the specified keys.  The associations in the
resulting catalog will be in the same order as the specified keys.

Merge() returns a new catalog containing all of the associations that are in
the specified catalogs in the order that they appear in each catalog.  If a
key is present in both catalogs, the value of the key from the second
catalog takes precedence.
*/
type CatalogClassLike[V any] interface {
	// Constructor Methods
	Catalog() CatalogLike[V]
	CatalogFromArray(
		associations []AssociationLike[Identifier, V],
	) CatalogLike[V]
	CatalogFromMap(
		associations map[Identifier]V,
	) CatalogLike[V]
	CatalogFromSequence(
		associations Sequential[AssociationLike[Identifier, V]],
	) CatalogLike[V]

	// Constant Methods
	Ranker() RankingFunction[AssociationLike[Identifier, V]]

	// Function Methods
	Extract(
		catalog CatalogLike[V],
		keys Sequential[Identifier],
	) CatalogLike[V]
	Merge(
		first CatalogLike[V],
		second CatalogLike[V],
	) CatalogLike[V]
}
`,

	instanceDeclarations_: `
/*
AngleLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete angle-like class.
*/
type AngleLike interface {
	// Principal Methods
	GetClass() AngleClassLike
	GetIntrinsic() float64
	IsZero() bool

	// Aspect Interfaces
	Angular
}

/*
ArrayLike[V any] is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete array-like class.

An array-like class maintains a fixed length indexed sequence of values.  Each
value is associated with an implicit positive integer index. An array-like class
uses ORDINAL based indexing rather than the more common—and nonsensical—ZERO
based indexing scheme.
*/
type ArrayLike[V any] interface {
	// Principal Methods
	GetClass() ArrayClassLike[V]
	GetIntrinsic() []V
	SortValuesWithRanker(
		ranker RankingFunction[V],
	)

	// Aspect Interfaces
	Accessible[V]
	Sequential[V]
	Updatable[V]
}

/*
AssociationLike[K comparable, V any] is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete association-like class.
*/
type AssociationLike[K comparable, V any] interface {
	// Principal Methods
	GetClass() AssociationClassLike[K, V]

	// Attribute Methods
	GetKey() K
	GetValue() V
	SetValue(
		value V,
	)
}

/*
CatalogLike[V any] is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete catalog-like class.
*/
type CatalogLike[V any] interface {
	// Principal Methods
	GetClass() CatalogClassLike[V]
	SortValues()

	// Aspect Interfaces
	Associative[Identifier, V]
	Sequential[AssociationLike[Identifier, V]]
}
`,

	aspectDeclarations_: `
/*
Accessible[V any] is an aspect interface that declares a set of method
signatures that must be supported by each instance of an accessible concrete
class.  The values in an accessible class are accessed using indices. The
indices of an accessible class are ORDINAL rather than ZERO based—which never
really made sense except for pointer offsets:

|      1           2           3             N      |
|  [value 1] . [value 2] . [value 3] ... [value N]  |

Because the indices are ordinal the range of allowed indices for a collection
is [1..N] instead of [0..N-1].
*/
type Accessible[V any] interface {
	GetValue(
		index Ordinal,
	) V
	GetValues(
		first Ordinal,
		last Ordinal,
	) Sequential[V]
}

/*
Angular is an aspect interface that declares a set of method signatures that
must be supported by each instance of an angular concrete class.
*/
type Angular interface {
	AsNormalized() AngleLike
	InUnits(
		units Units,
	) float64
	GetParts() (
		x float64,
		y float64,
	)
}

/*
Associative[K comparable, V any] declares the set of method signatures that
must be supported by all sequences of key-value associations.
*/
type Associative[K comparable, V any] interface {
	GetKeys() Sequential[K]
	GetValue(
		key K,
	) V
	SetValue(
		key K,
		value V,
	)
}

/*
Continuous is an aspect interface that declares a set of method signatures
that must be supported by each instance of a continuous concrete class.
*/
type Continuous interface {
	IsZero() bool
	IsDiscrete() bool
	IsInfinity() bool
}

/*
Sequential[V any] is an aspect interface that declares a set of method
signatures that must be supported by each instance of a sequential concrete
class.
*/
type Sequential[V any] interface {
	IsEmpty() bool
	GetSize() Ordinal
	AsArray() []V
}

/*
Updatable[V any] is an aspect interface that declares a set of method
signatures that must be supported by each instance of an updatable concrete
class.  It uses the same indexing scheme as the Accessible[V any] interface.
*/
type Updatable[V any] interface {
	SetValue(
		index Ordinal,
		value V,
	)
	SetValues(
		index Ordinal,
		values Sequential[V],
	)
}
`,
}
