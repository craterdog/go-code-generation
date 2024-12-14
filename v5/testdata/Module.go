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
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides constructors for each
commonly used class that is exported by the module.  Each constructor delegates
the actual construction process to its corresponding concrete class declared in
the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/craterdog/go-class-model/wiki
*/
package module

import (
	ast "github.com/craterdog/go-class-model/v5/ast"
	gra "github.com/craterdog/go-class-model/v5/grammar"
	abs "github.com/craterdog/go-collection-framework/v5/collection"
)

// TYPE ALIASES

// Ast

type (
	AbstractionLike           = ast.AbstractionLike
	AdditionalArgumentLike    = ast.AdditionalArgumentLike
	AdditionalConstraintLike  = ast.AdditionalConstraintLike
	AdditionalValueLike       = ast.AdditionalValueLike
	ArgumentLike              = ast.ArgumentLike
	ArgumentsLike             = ast.ArgumentsLike
	ArrayLike                 = ast.ArrayLike
	AspectDeclarationLike     = ast.AspectDeclarationLike
	AspectInterfaceLike       = ast.AspectInterfaceLike
	AspectMethodLike          = ast.AspectMethodLike
	AspectSectionLike         = ast.AspectSectionLike
	AspectSubsectionLike      = ast.AspectSubsectionLike
	AttributeMethodLike       = ast.AttributeMethodLike
	AttributeSubsectionLike   = ast.AttributeSubsectionLike
	ChannelLike               = ast.ChannelLike
	ClassDeclarationLike      = ast.ClassDeclarationLike
	ClassMethodsLike          = ast.ClassMethodsLike
	ClassSectionLike          = ast.ClassSectionLike
	ConstantMethodLike        = ast.ConstantMethodLike
	ConstantSubsectionLike    = ast.ConstantSubsectionLike
	ConstraintLike            = ast.ConstraintLike
	ConstraintsLike           = ast.ConstraintsLike
	ConstructorMethodLike     = ast.ConstructorMethodLike
	ConstructorSubsectionLike = ast.ConstructorSubsectionLike
	DeclarationLike           = ast.DeclarationLike
	EnumerationLike           = ast.EnumerationLike
	FunctionMethodLike        = ast.FunctionMethodLike
	FunctionSubsectionLike    = ast.FunctionSubsectionLike
	FunctionalDeclarationLike = ast.FunctionalDeclarationLike
	FunctionalSectionLike     = ast.FunctionalSectionLike
	GetterMethodLike          = ast.GetterMethodLike
	ImportedPackageLike       = ast.ImportedPackageLike
	InstanceDeclarationLike   = ast.InstanceDeclarationLike
	InstanceMethodsLike       = ast.InstanceMethodsLike
	InstanceSectionLike       = ast.InstanceSectionLike
	InterfaceDeclarationsLike = ast.InterfaceDeclarationsLike
	LegalNoticeLike           = ast.LegalNoticeLike
	MapLike                   = ast.MapLike
	MethodLike                = ast.MethodLike
	ModelLike                 = ast.ModelLike
	MultivalueLike            = ast.MultivalueLike
	NoneLike                  = ast.NoneLike
	PackageDeclarationLike    = ast.PackageDeclarationLike
	PackageHeaderLike         = ast.PackageHeaderLike
	PackageImportsLike        = ast.PackageImportsLike
	ParameterLike             = ast.ParameterLike
	PrefixLike                = ast.PrefixLike
	PrimitiveDeclarationsLike = ast.PrimitiveDeclarationsLike
	PrincipalMethodLike       = ast.PrincipalMethodLike
	PrincipalSubsectionLike   = ast.PrincipalSubsectionLike
	ResultLike                = ast.ResultLike
	SetterMethodLike          = ast.SetterMethodLike
	SuffixLike                = ast.SuffixLike
	TypeDeclarationLike       = ast.TypeDeclarationLike
	TypeSectionLike           = ast.TypeSectionLike
	ValueLike                 = ast.ValueLike
)

// Grammar

type (
	TokenType = gra.TokenType
)

const (
	ErrorToken     = gra.ErrorToken
	CommentToken   = gra.CommentToken
	DelimiterToken = gra.DelimiterToken
	NameToken      = gra.NameToken
	NewlineToken   = gra.NewlineToken
	PathToken      = gra.PathToken
	SpaceToken     = gra.SpaceToken
)

type (
	FormatterLike = gra.FormatterLike
	ParserLike    = gra.ParserLike
	ProcessorLike = gra.ProcessorLike
	ScannerLike   = gra.ScannerLike
	TokenLike     = gra.TokenLike
	ValidatorLike = gra.ValidatorLike
	VisitorLike   = gra.VisitorLike
)

type (
	Methodical = gra.Methodical
)

// CLASS CONSTRUCTORS

// ast/Abstraction

func Abstraction(
	optionalPrefix PrefixLike,
	name string,
	optionalSuffix SuffixLike,
	optionalArguments ArgumentsLike,
) AbstractionLike {
	return ast.AbstractionClass().Abstraction(
		optionalPrefix,
		name,
		optionalSuffix,
		optionalArguments,
	)
}

// ast/AdditionalArgument

func AdditionalArgument(
	argument ArgumentLike,
) AdditionalArgumentLike {
	return ast.AdditionalArgumentClass().AdditionalArgument(
		argument,
	)
}

// ast/AdditionalConstraint

func AdditionalConstraint(
	constraint ConstraintLike,
) AdditionalConstraintLike {
	return ast.AdditionalConstraintClass().AdditionalConstraint(
		constraint,
	)
}

// ast/AdditionalValue

func AdditionalValue(
	name string,
) AdditionalValueLike {
	return ast.AdditionalValueClass().AdditionalValue(
		name,
	)
}

// ast/Argument

func Argument(
	abstraction AbstractionLike,
) ArgumentLike {
	return ast.ArgumentClass().Argument(
		abstraction,
	)
}

// ast/Arguments

func Arguments(
	argument ArgumentLike,
	additionalArguments abs.Sequential[AdditionalArgumentLike],
) ArgumentsLike {
	return ast.ArgumentsClass().Arguments(
		argument,
		additionalArguments,
	)
}

// ast/Array

func Array() ArrayLike {
	return ast.ArrayClass().Array()
}

// ast/AspectDeclaration

func AspectDeclaration(
	declaration DeclarationLike,
	aspectMethods abs.Sequential[AspectMethodLike],
) AspectDeclarationLike {
	return ast.AspectDeclarationClass().AspectDeclaration(
		declaration,
		aspectMethods,
	)
}

// ast/AspectInterface

func AspectInterface(
	abstraction AbstractionLike,
) AspectInterfaceLike {
	return ast.AspectInterfaceClass().AspectInterface(
		abstraction,
	)
}

// ast/AspectMethod

func AspectMethod(
	method MethodLike,
) AspectMethodLike {
	return ast.AspectMethodClass().AspectMethod(
		method,
	)
}

// ast/AspectSection

func AspectSection(
	aspectDeclarations abs.Sequential[AspectDeclarationLike],
) AspectSectionLike {
	return ast.AspectSectionClass().AspectSection(
		aspectDeclarations,
	)
}

// ast/AspectSubsection

func AspectSubsection(
	aspectInterfaces abs.Sequential[AspectInterfaceLike],
) AspectSubsectionLike {
	return ast.AspectSubsectionClass().AspectSubsection(
		aspectInterfaces,
	)
}

// ast/AttributeMethod

func AttributeMethod(
	any_ any,
) AttributeMethodLike {
	return ast.AttributeMethodClass().AttributeMethod(
		any_,
	)
}

// ast/AttributeSubsection

func AttributeSubsection(
	attributeMethods abs.Sequential[AttributeMethodLike],
) AttributeSubsectionLike {
	return ast.AttributeSubsectionClass().AttributeSubsection(
		attributeMethods,
	)
}

// ast/Channel

func Channel() ChannelLike {
	return ast.ChannelClass().Channel()
}

// ast/ClassDeclaration

func ClassDeclaration(
	declaration DeclarationLike,
	classMethods ClassMethodsLike,
) ClassDeclarationLike {
	return ast.ClassDeclarationClass().ClassDeclaration(
		declaration,
		classMethods,
	)
}

// ast/ClassMethods

func ClassMethods(
	constructorSubsection ConstructorSubsectionLike,
	optionalConstantSubsection ConstantSubsectionLike,
	optionalFunctionSubsection FunctionSubsectionLike,
) ClassMethodsLike {
	return ast.ClassMethodsClass().ClassMethods(
		constructorSubsection,
		optionalConstantSubsection,
		optionalFunctionSubsection,
	)
}

// ast/ClassSection

func ClassSection(
	classDeclarations abs.Sequential[ClassDeclarationLike],
) ClassSectionLike {
	return ast.ClassSectionClass().ClassSection(
		classDeclarations,
	)
}

// ast/ConstantMethod

func ConstantMethod(
	name string,
	abstraction AbstractionLike,
) ConstantMethodLike {
	return ast.ConstantMethodClass().ConstantMethod(
		name,
		abstraction,
	)
}

// ast/ConstantSubsection

func ConstantSubsection(
	constantMethods abs.Sequential[ConstantMethodLike],
) ConstantSubsectionLike {
	return ast.ConstantSubsectionClass().ConstantSubsection(
		constantMethods,
	)
}

// ast/Constraint

func Constraint(
	name string,
	abstraction AbstractionLike,
) ConstraintLike {
	return ast.ConstraintClass().Constraint(
		name,
		abstraction,
	)
}

// ast/Constraints

func Constraints(
	constraint ConstraintLike,
	additionalConstraints abs.Sequential[AdditionalConstraintLike],
) ConstraintsLike {
	return ast.ConstraintsClass().Constraints(
		constraint,
		additionalConstraints,
	)
}

// ast/ConstructorMethod

func ConstructorMethod(
	name string,
	parameters abs.Sequential[ParameterLike],
	abstraction AbstractionLike,
) ConstructorMethodLike {
	return ast.ConstructorMethodClass().ConstructorMethod(
		name,
		parameters,
		abstraction,
	)
}

// ast/ConstructorSubsection

func ConstructorSubsection(
	constructorMethods abs.Sequential[ConstructorMethodLike],
) ConstructorSubsectionLike {
	return ast.ConstructorSubsectionClass().ConstructorSubsection(
		constructorMethods,
	)
}

// ast/Declaration

func Declaration(
	comment string,
	name string,
	optionalConstraints ConstraintsLike,
) DeclarationLike {
	return ast.DeclarationClass().Declaration(
		comment,
		name,
		optionalConstraints,
	)
}

// ast/Enumeration

func Enumeration(
	value ValueLike,
	additionalValues abs.Sequential[AdditionalValueLike],
) EnumerationLike {
	return ast.EnumerationClass().Enumeration(
		value,
		additionalValues,
	)
}

// ast/FunctionMethod

func FunctionMethod(
	name string,
	parameters abs.Sequential[ParameterLike],
	result ResultLike,
) FunctionMethodLike {
	return ast.FunctionMethodClass().FunctionMethod(
		name,
		parameters,
		result,
	)
}

// ast/FunctionSubsection

func FunctionSubsection(
	functionMethods abs.Sequential[FunctionMethodLike],
) FunctionSubsectionLike {
	return ast.FunctionSubsectionClass().FunctionSubsection(
		functionMethods,
	)
}

// ast/FunctionalDeclaration

func FunctionalDeclaration(
	declaration DeclarationLike,
	parameters abs.Sequential[ParameterLike],
	result ResultLike,
) FunctionalDeclarationLike {
	return ast.FunctionalDeclarationClass().FunctionalDeclaration(
		declaration,
		parameters,
		result,
	)
}

// ast/FunctionalSection

func FunctionalSection(
	functionalDeclarations abs.Sequential[FunctionalDeclarationLike],
) FunctionalSectionLike {
	return ast.FunctionalSectionClass().FunctionalSection(
		functionalDeclarations,
	)
}

// ast/GetterMethod

func GetterMethod(
	name string,
	abstraction AbstractionLike,
) GetterMethodLike {
	return ast.GetterMethodClass().GetterMethod(
		name,
		abstraction,
	)
}

// ast/ImportedPackage

func ImportedPackage(
	name string,
	path string,
) ImportedPackageLike {
	return ast.ImportedPackageClass().ImportedPackage(
		name,
		path,
	)
}

// ast/InstanceDeclaration

func InstanceDeclaration(
	declaration DeclarationLike,
	instanceMethods InstanceMethodsLike,
) InstanceDeclarationLike {
	return ast.InstanceDeclarationClass().InstanceDeclaration(
		declaration,
		instanceMethods,
	)
}

// ast/InstanceMethods

func InstanceMethods(
	principalSubsection PrincipalSubsectionLike,
	optionalAttributeSubsection AttributeSubsectionLike,
	optionalAspectSubsection AspectSubsectionLike,
) InstanceMethodsLike {
	return ast.InstanceMethodsClass().InstanceMethods(
		principalSubsection,
		optionalAttributeSubsection,
		optionalAspectSubsection,
	)
}

// ast/InstanceSection

func InstanceSection(
	instanceDeclarations abs.Sequential[InstanceDeclarationLike],
) InstanceSectionLike {
	return ast.InstanceSectionClass().InstanceSection(
		instanceDeclarations,
	)
}

// ast/InterfaceDeclarations

func InterfaceDeclarations(
	classSection ClassSectionLike,
	instanceSection InstanceSectionLike,
	aspectSection AspectSectionLike,
) InterfaceDeclarationsLike {
	return ast.InterfaceDeclarationsClass().InterfaceDeclarations(
		classSection,
		instanceSection,
		aspectSection,
	)
}

// ast/LegalNotice

func LegalNotice(
	comment string,
) LegalNoticeLike {
	return ast.LegalNoticeClass().LegalNotice(
		comment,
	)
}

// ast/Map

func Map(
	name string,
) MapLike {
	return ast.MapClass().Map(
		name,
	)
}

// ast/Method

func Method(
	name string,
	parameters abs.Sequential[ParameterLike],
	optionalResult ResultLike,
) MethodLike {
	return ast.MethodClass().Method(
		name,
		parameters,
		optionalResult,
	)
}

// ast/Model

func Model(
	packageDeclaration PackageDeclarationLike,
	primitiveDeclarations PrimitiveDeclarationsLike,
	interfaceDeclarations InterfaceDeclarationsLike,
) ModelLike {
	return ast.ModelClass().Model(
		packageDeclaration,
		primitiveDeclarations,
		interfaceDeclarations,
	)
}

// ast/Multivalue

func Multivalue(
	parameters abs.Sequential[ParameterLike],
) MultivalueLike {
	return ast.MultivalueClass().Multivalue(
		parameters,
	)
}

// ast/None

func None(
	newline string,
) NoneLike {
	return ast.NoneClass().None(
		newline,
	)
}

// ast/PackageDeclaration

func PackageDeclaration(
	legalNotice LegalNoticeLike,
	packageHeader PackageHeaderLike,
	packageImports PackageImportsLike,
) PackageDeclarationLike {
	return ast.PackageDeclarationClass().PackageDeclaration(
		legalNotice,
		packageHeader,
		packageImports,
	)
}

// ast/PackageHeader

func PackageHeader(
	comment string,
	name string,
) PackageHeaderLike {
	return ast.PackageHeaderClass().PackageHeader(
		comment,
		name,
	)
}

// ast/PackageImports

func PackageImports(
	importedPackages abs.Sequential[ImportedPackageLike],
) PackageImportsLike {
	return ast.PackageImportsClass().PackageImports(
		importedPackages,
	)
}

// ast/Parameter

func Parameter(
	name string,
	abstraction AbstractionLike,
) ParameterLike {
	return ast.ParameterClass().Parameter(
		name,
		abstraction,
	)
}

// ast/Prefix

func Prefix(
	any_ any,
) PrefixLike {
	return ast.PrefixClass().Prefix(
		any_,
	)
}

// ast/PrimitiveDeclarations

func PrimitiveDeclarations(
	typeSection TypeSectionLike,
	functionalSection FunctionalSectionLike,
) PrimitiveDeclarationsLike {
	return ast.PrimitiveDeclarationsClass().PrimitiveDeclarations(
		typeSection,
		functionalSection,
	)
}

// ast/PrincipalMethod

func PrincipalMethod(
	method MethodLike,
) PrincipalMethodLike {
	return ast.PrincipalMethodClass().PrincipalMethod(
		method,
	)
}

// ast/PrincipalSubsection

func PrincipalSubsection(
	principalMethods abs.Sequential[PrincipalMethodLike],
) PrincipalSubsectionLike {
	return ast.PrincipalSubsectionClass().PrincipalSubsection(
		principalMethods,
	)
}

// ast/Result

func Result(
	any_ any,
) ResultLike {
	return ast.ResultClass().Result(
		any_,
	)
}

// ast/SetterMethod

func SetterMethod(
	name string,
	parameter ParameterLike,
) SetterMethodLike {
	return ast.SetterMethodClass().SetterMethod(
		name,
		parameter,
	)
}

// ast/Suffix

func Suffix(
	name string,
) SuffixLike {
	return ast.SuffixClass().Suffix(
		name,
	)
}

// ast/TypeDeclaration

func TypeDeclaration(
	declaration DeclarationLike,
	abstraction AbstractionLike,
	optionalEnumeration EnumerationLike,
) TypeDeclarationLike {
	return ast.TypeDeclarationClass().TypeDeclaration(
		declaration,
		abstraction,
		optionalEnumeration,
	)
}

// ast/TypeSection

func TypeSection(
	typeDeclarations abs.Sequential[TypeDeclarationLike],
) TypeSectionLike {
	return ast.TypeSectionClass().TypeSection(
		typeDeclarations,
	)
}

// ast/Value

func Value(
	name string,
	abstraction AbstractionLike,
) ValueLike {
	return ast.ValueClass().Value(
		name,
		abstraction,
	)
}

// grammar/Formatter

func Formatter() FormatterLike {
	return gra.FormatterClass().Formatter()
}

// grammar/Parser

func Parser() ParserLike {
	return gra.ParserClass().Parser()
}

// grammar/Processor

func Processor() ProcessorLike {
	return gra.ProcessorClass().Processor()
}

// grammar/Scanner

func Scanner(
	source string,
	tokens abs.QueueLike[TokenLike],
) ScannerLike {
	return gra.ScannerClass().Scanner(
		source,
		tokens,
	)
}

// grammar/Token

func Token(
	line uint,
	position uint,
	type_ TokenType,
	value string,
) TokenLike {
	return gra.TokenClass().Token(
		line,
		position,
		type_,
		value,
	)
}

// grammar/Validator

func Validator() ValidatorLike {
	return gra.ValidatorClass().Validator()
}

// grammar/Visitor

func Visitor(
	processor Methodical,
) VisitorLike {
	return gra.VisitorClass().Visitor(
		processor,
	)
}

// GLOBAL FUNCTIONS
