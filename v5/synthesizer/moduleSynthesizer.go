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
	mod "github.com/craterdog/go-class-model/v5/ast"
	ana "github.com/craterdog/go-code-generation/v5/analyzer"
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ModuleSynthesizer() ModuleSynthesizerClassLike {
	return moduleSynthesizerReference()
}

// Constructor Methods

func (c *moduleSynthesizerClass_) Make(
	models abs.CatalogLike[string, mod.ModelLike],
) ModuleSynthesizerLike {
	var instance = &moduleSynthesizer_{
		// Initialize the instance attributes.
		models_: models,
	}
	return instance
}

// INSTANCE INTERFACE

// Primary Methods

func (v *moduleSynthesizer_) GetClass() ModuleSynthesizerClassLike {
	return moduleSynthesizerReference()
}

// TemplateDriven Methods

func (v *moduleSynthesizer_) CreateLegalNotice() string {
	var packageNames = v.models_.GetKeys().AsArray()
	var firstPackage = packageNames[0]
	var firstModel = v.models_.GetValue(firstPackage)
	var moduleDeclaration = firstModel.GetModuleDeclaration()
	var legalNotice = moduleDeclaration.GetLegalNotice().GetComment()
	return legalNotice
}

func (v *moduleSynthesizer_) CreateTypeAliases() string {
	var typeAliases string
	var models = v.models_.GetIterator()
	for models.HasNext() {
		var association = models.GetNext()
		var packageName = association.GetKey()
		var model = association.GetValue()
		var packageAliases = v.createPackageAliases(packageName, model)
		typeAliases += packageAliases
	}
	return typeAliases
}

func (v *moduleSynthesizer_) CreateUniversalConstructors() string {
	var universalConstructors string
	var models = v.models_.GetIterator()
	for models.HasNext() {
		var association = models.GetNext()
		var packageName = association.GetKey()
		var model = association.GetValue()
		var constructorFunctions = v.createConstructorFunctions(packageName, model)
		universalConstructors += constructorFunctions
	}
	return universalConstructors
}

func (v *moduleSynthesizer_) PerformGlobalUpdates(
	source string,
) string {
	var moduleImports string
	var importedPackages = v.createImportedPackages(source)
	if uti.IsDefined(importedPackages) {
		moduleImports = moduleSynthesizerReference().moduleImports_
		moduleImports = uti.ReplaceAll(
			moduleImports,
			"importedPackages",
			importedPackages,
		)
	}
	source = uti.ReplaceAll(
		source,
		"moduleImports",
		moduleImports,
	)
	return source
}

// PROTECTED INTERFACE

// Private Methods

func (v *moduleSynthesizer_) createInitializations(
	constructorMethods abs.ListLike[mod.ConstructorMethodLike],
) string {
	var initializations string
	var arguments = v.extractArguments(constructorMethods).GetIterator()
	for arguments.HasNext() {
		var argument = arguments.GetNext()
		var argumentName = argument.GetKey()
		var argumentType = v.extractType(argument.GetValue())
		var initialization = moduleSynthesizerReference().argumentInitialization_
		initialization = uti.ReplaceAll(
			initialization,
			"argumentName",
			argumentName,
		)
		initialization = uti.ReplaceAll(
			initialization,
			"argumentType",
			argumentType,
		)
		initializations += initialization
	}
	return initializations
}

func (v *moduleSynthesizer_) createValidations(
	constructorMethods abs.ListLike[mod.ConstructorMethodLike],
) string {
	var validations string
	var arguments = v.extractArguments(constructorMethods).GetIterator()
	for arguments.HasNext() {
		var argument = arguments.GetNext()
		var argumentName = argument.GetKey()
		var argumentType = v.extractType(argument.GetValue())
		var validation = moduleSynthesizerReference().validationCase_
		validation = uti.ReplaceAll(
			validation,
			"argumentName",
			argumentName,
		)
		validation = uti.ReplaceAll(
			validation,
			"argumentType",
			argumentType,
		)
		validations += validation
	}
	return validations
}

func (v *moduleSynthesizer_) createConstructions(
	constructorMethods abs.ListLike[mod.ConstructorMethodLike],
) string {
	var constructions string
	return constructions
}

func (v *moduleSynthesizer_) createConstructorFunction(
	model mod.ModelLike,
	class mod.ClassDeclarationLike,
) string {
	var constructorFunction = moduleSynthesizerReference().constructorFunction_
	var className = sts.TrimSuffix(class.GetDeclaration().GetName(), "ClassLike")
	className = uti.MakeLowerCase(className)
	var analyzer = ana.ModelAnalyzer().Make(model, className)
	var constructorMethods = analyzer.GetConstructorMethods()
	var initializations = v.createInitializations(constructorMethods)
	constructorFunction = uti.ReplaceAll(
		constructorFunction,
		"initializations",
		initializations,
	)
	var validations = v.createValidations(constructorMethods)
	constructorFunction = uti.ReplaceAll(
		constructorFunction,
		"validations",
		validations,
	)
	var constructions = v.createConstructions(constructorMethods)
	constructorFunction = uti.ReplaceAll(
		constructorFunction,
		"constructions",
		constructions,
	)
	constructorFunction = uti.ReplaceAll(
		constructorFunction,
		"className",
		className,
	)
	return constructorFunction
}

func (v *moduleSynthesizer_) createConstructorFunctions(
	packageName string,
	model mod.ModelLike,
) string {
	var constructorFunctions string
	var interfaceDeclarations = model.GetInterfaceDeclarations()
	var classSection = interfaceDeclarations.GetClassSection()
	var classes = classSection.GetClassDeclarations().GetIterator()
	for classes.HasNext() {
		var class = classes.GetNext()
		var constructorFunction = v.createConstructorFunction(model, class)
		constructorFunctions += constructorFunction
	}
	var universalConstructors = moduleSynthesizerReference().universalConstructors_
	universalConstructors = uti.ReplaceAll(
		universalConstructors,
		"constructorFunctions",
		constructorFunctions,
	)
	universalConstructors = v.replacePackageAttributes(
		universalConstructors,
		packageName,
	)
	return universalConstructors
}

func (v *moduleSynthesizer_) createTypeAlias(
	model mod.ModelLike,
	class mod.ClassDeclarationLike,
) string {
	var typeAlias string
	var declaration = class.GetDeclaration()
	var className = sts.TrimSuffix(declaration.GetName(), "ClassLike")
	var analyzer = ana.ModelAnalyzer().Make(model, className)
	if analyzer.IsGeneric() {
		// Type aliases are not supported for generic types in Go.
		return typeAlias
	}
	typeAlias = moduleSynthesizerReference().typeAlias_
	typeAlias = uti.ReplaceAll(
		typeAlias,
		"className",
		className,
	)
	return typeAlias
}

func (v *moduleSynthesizer_) createPackageAliases(
	packageName string,
	model mod.ModelLike,
) string {
	var typeAliases string
	var interfaceDeclarations = model.GetInterfaceDeclarations()
	var classSection = interfaceDeclarations.GetClassSection()
	var classes = classSection.GetClassDeclarations().GetIterator()
	for classes.HasNext() {
		var class = classes.GetNext()
		var typeAlias = v.createTypeAlias(model, class)
		typeAliases += typeAlias
	}
	var packageAliases = moduleSynthesizerReference().packageAliases_
	packageAliases = uti.ReplaceAll(
		packageAliases,
		"typeAliases",
		typeAliases,
	)
	packageAliases = v.replacePackageAttributes(
		packageAliases,
		packageName,
	)
	return packageAliases
}

func (v *moduleSynthesizer_) createImportedPackages(
	source string,
) string {
	var importedPackages string
	return importedPackages
}

func (v *moduleSynthesizer_) extractArguments(
	constructorMethods abs.ListLike[mod.ConstructorMethodLike],
) abs.CatalogLike[string, mod.AbstractionLike] {
	var arguments = col.Catalog[string, mod.AbstractionLike]()
	var methods = constructorMethods.GetIterator()
	for methods.HasNext() {
		var method = methods.GetNext()
		var parameters = method.GetParameters().GetIterator()
		for parameters.HasNext() {
			var parameter = parameters.GetNext()
			var argumentName = parameter.GetName()
			var argumentType = parameter.GetAbstraction()
			arguments.SetValue(argumentName, argumentType)
		}
	}
	return arguments
}

func (v *moduleSynthesizer_) extractType(
	abstraction mod.AbstractionLike,
) string {
	var abstractType string
	var prefix = abstraction.GetOptionalPrefix()
	if uti.IsDefined(prefix) {
		switch actual := prefix.GetAny().(type) {
		case mod.ArrayLike:
			abstractType = "[]"
		case mod.MapLike:
			abstractType = "map[" + actual.GetName() + "]"
		case mod.ChannelLike:
			abstractType = "chan "
		}
	}
	var name = abstraction.GetName()
	abstractType += name
	var suffix = abstraction.GetOptionalSuffix()
	if uti.IsDefined(suffix) {
		abstractType += "." + suffix.GetName()
	}
	var arguments = abstraction.GetOptionalArguments()
	if uti.IsDefined(arguments) {
		var argument = v.extractType(arguments.GetArgument().GetAbstraction())
		abstractType += "[" + argument
		var additionalArguments = arguments.GetAdditionalArguments().GetIterator()
		for additionalArguments.HasNext() {
			var additionalArgument = additionalArguments.GetNext().GetArgument()
			argument = v.extractType(additionalArgument.GetAbstraction())
			abstractType += ", " + argument
		}
		abstractType += "]"
	}
	return abstractType
}

func (v *moduleSynthesizer_) replacePackageAttributes(
	source string,
	packageName string,
) string {
	source = uti.ReplaceAll(
		source,
		"packageName",
		packageName,
	)
	var packageAcronym = packageName[0:3]
	source = uti.ReplaceAll(
		source,
		"packageAcronym",
		packageAcronym,
	)
	return source
}

// Instance Structure

type moduleSynthesizer_ struct {
	// Declare the instance attributes.
	models_ abs.CatalogLike[string, mod.ModelLike]
}

// Class Structure

type moduleSynthesizerClass_ struct {
	// Declare the class constants.
	moduleImports_          string
	packageAlias_           string
	packageAliases_         string
	typeAlias_              string
	universalConstructors_  string
	constructorFunction_    string
	argumentInitialization_ string
	validationCase_         string
	constructionCase_       string
}

// Class Reference

func moduleSynthesizerReference() *moduleSynthesizerClass_ {
	return moduleSynthesizerReference_
}

var moduleSynthesizerReference_ = &moduleSynthesizerClass_{
	// Initialize the class constants.
	moduleImports_: `

import (<ImportedPackages>)`,

	packageAlias_: `
	<~packageAcronym> <packagePath>`,

	packageAliases_: `

// <~PackageName>

type (<TypeAliases>
)`,

	typeAlias_: `
	<~ClassName>Like = <~packageAcronym>.<~ClassName>Like`,

	universalConstructors_: `

// <~PackageName><ConstructorFunctions>`,

	constructorFunction_: `

func <~ClassName>(arguments ...any) <~ClassName>Like {
	// Initialize the possible arguments.<Initializations>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<Validations>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the <~ClassName> constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var <className_> <~ClassName>Like
	switch {<Constructions>
		default:
			var message = fmt.Sprintf(
				"A <~ClassName> constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return <className_>
}`,

	argumentInitialization_: `
	var <argumentName_> <ArgumentType>`,

	validationCase_: `
		case <ArgumentType>:
			<argumentName_> = actual`,

	constructionCase_: `
	var <className_> = <~packageAcronym>.<~ClassName>().Make(<Arguments>)`,
}
