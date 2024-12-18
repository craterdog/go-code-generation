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
	abs "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ModuleSynthesizerClass() ModuleSynthesizerClassLike {
	return moduleSynthesizerClassReference()
}

// Constructor Methods

func (c *moduleSynthesizerClass_) ModuleSynthesizer(
	models abs.CatalogLike[string, mod.ModelLike],
) ModuleSynthesizerLike {
	var instance = &moduleSynthesizer_{
		// Initialize the instance attributes.
		models_: models,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *moduleSynthesizer_) GetClass() ModuleSynthesizerClassLike {
	return moduleSynthesizerClassReference()
}

// TemplateDriven Methods

func (v *moduleSynthesizer_) CreateLegalNotice() string {
	var packageNames = v.models_.GetKeys().AsArray()
	var firstPackage = packageNames[0]
	var firstModel = v.models_.GetValue(firstPackage)
	var packageDeclaration = firstModel.GetPackageDeclaration()
	var legalNotice = packageDeclaration.GetLegalNotice().GetComment()
	return legalNotice
}

func (v *moduleSynthesizer_) CreateWarningMessage() string {
	var class = moduleSynthesizerClassReference()
	var warningMessage = class.warningMessage_
	return warningMessage
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

func (v *moduleSynthesizer_) CreateClassConstructors() string {
	var classConstructors string
	var models = v.models_.GetIterator()
	for models.HasNext() {
		var association = models.GetNext()
		var packageName = association.GetKey()
		var model = association.GetValue()
		var constructorFunctions = v.createConstructorFunctions(packageName, model)
		classConstructors += constructorFunctions
	}
	return classConstructors
}

func (v *moduleSynthesizer_) PerformGlobalUpdates(
	source string,
) string {
	var importedPackages = v.createImportedPackages(source)
	source = uti.ReplaceAll(
		source,
		"importedPackages",
		importedPackages,
	)
	return source
}

// PROTECTED INTERFACE

// Private Methods

func (v *moduleSynthesizer_) createAspectAliases(
	interfaceDeclarations mod.InterfaceDeclarationsLike,
) (
	aspectAliases string,
) {
	var nameAliases string
	var aspectSection = interfaceDeclarations.GetAspectSection()
	var aspects = aspectSection.GetAspectDeclarations().GetIterator()
	for aspects.HasNext() {
		var declaration = aspects.GetNext().GetDeclaration()
		var nameAlias = v.createNameAlias(declaration)
		nameAliases += nameAlias
	}
	if uti.IsDefined(nameAliases) {
		nameAliases += "\n"
		var class = moduleSynthesizerClassReference()
		aspectAliases = class.typeAliases_
		aspectAliases = uti.ReplaceAll(
			aspectAliases,
			"nameAliases",
			nameAliases,
		)
	}
	return
}

func (v *moduleSynthesizer_) createClassAliases(
	interfaceDeclarations mod.InterfaceDeclarationsLike,
) (
	classAliases string,
) {
	var nameAliases string
	var instanceSection = interfaceDeclarations.GetInstanceSection()
	var instances = instanceSection.GetInstanceDeclarations().GetIterator()
	for instances.HasNext() {
		var declaration = instances.GetNext().GetDeclaration()
		nameAliases += v.createNameAlias(declaration)
	}
	if uti.IsDefined(nameAliases) {
		nameAliases += "\n"
		var class = moduleSynthesizerClassReference()
		classAliases = class.typeAliases_
		classAliases = uti.ReplaceAll(
			classAliases,
			"nameAliases",
			nameAliases,
		)
	}
	return
}

func (v *moduleSynthesizer_) createConstructorFunction(
	constructorMethod mod.ConstructorMethodLike,
	className string,
) string {
	var class = moduleSynthesizerClassReference()
	var constructorFunction = class.constructorFunction_
	var methodName = constructorMethod.GetName()
	var parameters = v.extractParameters(constructorMethod)
	constructorFunction = uti.ReplaceAll(
		constructorFunction,
		"methodName",
		methodName,
	)
	constructorFunction = uti.ReplaceAll(
		constructorFunction,
		"parameters",
		parameters,
	)
	var parameterNames = v.extractParameterNames(constructorMethod)
	constructorFunction = uti.ReplaceAll(
		constructorFunction,
		"parameterNames",
		parameterNames,
	)
	return constructorFunction
}

func (v *moduleSynthesizer_) createClassConstructors(
	model mod.ModelLike,
	classDeclaration mod.ClassDeclarationLike,
) string {
	var constructors string
	var className = sts.TrimSuffix(
		classDeclaration.GetDeclaration().GetName(),
		"ClassLike",
	)
	className = uti.MakeLowerCase(className)
	var analyzerClass = ana.ModelAnalyzerClass()
	var analyzer = analyzerClass.ModelAnalyzer(model, className)
	var constructorMethods = analyzer.GetConstructorMethods().GetIterator()
	for constructorMethods.HasNext() {
		var constructorMethod = constructorMethods.GetNext()
		var constructorFunction = v.createConstructorFunction(
			constructorMethod,
			className,
		)
		constructors += constructorFunction
	}
	var class = moduleSynthesizerClassReference()
	var classConstructors = class.classConstructors_
	classConstructors = uti.ReplaceAll(
		classConstructors,
		"constructors",
		constructors,
	)
	classConstructors = uti.ReplaceAll(
		classConstructors,
		"className",
		className,
	)
	var constraints = analyzer.GetTypeConstraints()
	var arguments = analyzer.GetTypeArguments()
	classConstructors = uti.ReplaceAll(
		classConstructors,
		"constraints",
		constraints,
	)
	classConstructors = uti.ReplaceAll(
		classConstructors,
		"arguments",
		arguments,
	)

	return classConstructors
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
		var constructors = v.createClassConstructors(model, class)
		constructorFunctions += constructors
	}
	constructorFunctions = v.replacePackageAttributes(
		constructorFunctions,
		packageName,
	)
	return constructorFunctions
}

func (v *moduleSynthesizer_) createEnumerationAlias(
	name string,
) string {
	var class = moduleSynthesizerClassReference()
	var nameAlias = class.nameAlias_
	nameAlias = uti.ReplaceAll(
		nameAlias,
		"name",
		name,
	)
	return nameAlias
}

func (v *moduleSynthesizer_) createEnumerationAliases(
	enumeration mod.EnumerationLike,
) (
	constAliases string,
) {
	if uti.IsUndefined(enumeration) {
		return
	}
	var value = enumeration.GetValue()
	var name = value.GetName()
	var nameAliases = v.createEnumerationAlias(name)
	var values = enumeration.GetAdditionalValues().GetIterator()
	for values.HasNext() {
		name = values.GetNext().GetName()
		nameAliases += v.createEnumerationAlias(name)
	}
	if uti.IsDefined(nameAliases) {
		nameAliases += "\n"
		var class = moduleSynthesizerClassReference()
		constAliases = class.constAliases_
		constAliases = uti.ReplaceAll(
			constAliases,
			"nameAliases",
			nameAliases,
		)
	}
	return
}

func (v *moduleSynthesizer_) createFunctionalAliases(
	primitiveDeclarations mod.PrimitiveDeclarationsLike,
) (
	functionalAliases string,
) {
	var nameAliases string
	var functionalSection = primitiveDeclarations.GetFunctionalSection()
	var functionals = functionalSection.GetFunctionalDeclarations().GetIterator()
	for functionals.HasNext() {
		var declaration = functionals.GetNext().GetDeclaration()
		var nameAlias = v.createNameAlias(declaration)
		nameAliases += nameAlias
	}
	if uti.IsDefined(nameAliases) {
		nameAliases += "\n"
		var class = moduleSynthesizerClassReference()
		functionalAliases = class.typeAliases_
		functionalAliases = uti.ReplaceAll(
			functionalAliases,
			"nameAliases",
			nameAliases,
		)
	}
	return
}

func (v *moduleSynthesizer_) createImportedPackages(
	source string,
) string {
	var importedPackages string
	var class = moduleSynthesizerClassReference()
	var packageNames = v.models_.GetKeys().GetIterator()
	for packageNames.HasNext() {
		var packageName = packageNames.GetNext()
		var importedPackage = class.importedPackage_
		importedPackage = v.replacePackageAttributes(importedPackage, packageName)
		importedPackages += importedPackage
	}
	if sts.Contains(source, "fmt.") && !sts.Contains(importedPackages, "fmt") {
		var packageAlias = class.packageAlias_
		packageAlias = uti.ReplaceAll(
			packageAlias,
			"packageAcronym",
			"fmt",
		)
		packageAlias = uti.ReplaceAll(
			packageAlias,
			"packagePath",
			"fmt",
		)
		importedPackages += packageAlias
	}
	if sts.Contains(source, "abs.") && !sts.Contains(importedPackages, "abs") {
		var packageAlias = class.packageAlias_
		packageAlias = uti.ReplaceAll(
			packageAlias,
			"packageAcronym",
			"abs",
		)
		packageAlias = uti.ReplaceAll(
			packageAlias,
			"packagePath",
			"github.com/craterdog/go-collection-framework/v5/collection",
		)
		importedPackages += packageAlias
	}
	if uti.IsDefined(importedPackages) {
		importedPackages += "\n"
	}
	return importedPackages
}

func (v *moduleSynthesizer_) createNameAlias(
	declaration mod.DeclarationLike,
) (
	nameAlias string,
) {
	var name = declaration.GetName()
	if uti.IsDefined(declaration.GetOptionalConstraints()) {
		// Type aliases are not supported for generic types in Go.
		return
	}
	var class = moduleSynthesizerClassReference()
	nameAlias = class.nameAlias_
	nameAlias = uti.ReplaceAll(
		nameAlias,
		"name",
		name,
	)
	return
}

func (v *moduleSynthesizer_) createPackageAliases(
	packageName string,
	model mod.ModelLike,
) string {
	var primitiveDeclarations = model.GetPrimitiveDeclarations()
	var typeAliases = v.createTypeAliases(
		primitiveDeclarations,
	)
	typeAliases += v.createFunctionalAliases(
		primitiveDeclarations,
	)
	var interfaceDeclarations = model.GetInterfaceDeclarations()
	typeAliases += v.createClassAliases(
		interfaceDeclarations,
	)
	typeAliases += v.createAspectAliases(
		interfaceDeclarations,
	)
	var class = moduleSynthesizerClassReference()
	var packageAliases = class.packageAliases_
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

func (v *moduleSynthesizer_) createTypeAliases(
	primitiveDeclarations mod.PrimitiveDeclarationsLike,
) (
	typeAliases string,
) {
	var nameAliases string
	var constAliases string
	var typeSection = primitiveDeclarations.GetTypeSection()
	var types = typeSection.GetTypeDeclarations().GetIterator()
	for types.HasNext() {
		var typeDeclaration = types.GetNext()
		var declaration = typeDeclaration.GetDeclaration()
		nameAliases += v.createNameAlias(declaration)
		var enumeration = typeDeclaration.GetOptionalEnumeration()
		constAliases += v.createEnumerationAliases(enumeration)
	}
	if uti.IsDefined(nameAliases) {
		nameAliases += "\n"
		var class = moduleSynthesizerClassReference()
		typeAliases = class.typeAliases_
		typeAliases = uti.ReplaceAll(
			typeAliases,
			"nameAliases",
			nameAliases,
		)
	}
	typeAliases += constAliases
	return
}

func (v *moduleSynthesizer_) extractParameterNames(
	constructorMethod mod.ConstructorMethodLike,
) string {
	var parameterNames string
	var parameters = constructorMethod.GetParameters().GetIterator()
	if parameters.IsEmpty() {
		return parameterNames
	}
	for parameters.HasNext() {
		var parameter = parameters.GetNext()
		var parameterName = parameter.GetName()
		parameterNames += "\n\t\t" + parameterName + ","
	}
	parameterNames += "\n\t"
	return parameterNames
}

func (v *moduleSynthesizer_) extractParameters(
	constructorMethod mod.ConstructorMethodLike,
) string {
	var methodParameters string
	var parameters = constructorMethod.GetParameters().GetIterator()
	for parameters.HasNext() {
		var parameter = parameters.GetNext()
		var parameterName = parameter.GetName()
		var parameterType = v.extractType(parameter.GetAbstraction())
		var class = moduleSynthesizerClassReference()
		var methodParameter = class.methodParameter_
		methodParameter = uti.ReplaceAll(
			methodParameter,
			"parameterName",
			parameterName,
		)
		methodParameter = uti.ReplaceAll(
			methodParameter,
			"parameterType",
			parameterType,
		)
		methodParameters += methodParameter
	}
	if uti.IsDefined(methodParameters) {
		methodParameters += "\n"
	}
	return methodParameters
}

func (v *moduleSynthesizer_) extractType(
	abstraction mod.AbstractionLike,
) string {
	var abstractType string
	var wrapper = abstraction.GetOptionalWrapper()
	if uti.IsDefined(wrapper) {
		switch actual := wrapper.GetAny().(type) {
		case mod.ArrayLike:
			abstractType = "[]"
		case mod.MapLike:
			abstractType = "map[" + actual.GetName() + "]"
		case mod.ChannelLike:
			abstractType = "chan "
		}
	}
	var prefix = abstraction.GetOptionalPrefix()
	if uti.IsDefined(prefix) {
		abstractType += prefix
	}
	var name = abstraction.GetName()
	abstractType += name
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
	warningMessage_      string
	packageAlias_        string
	importedPackage_     string
	packageAliases_      string
	typeAliases_         string
	constAliases_        string
	nameAlias_           string
	classConstructors_   string
	constructorFunction_ string
	methodParameter_     string
}

// Class Reference

func moduleSynthesizerClassReference() *moduleSynthesizerClass_ {
	return moduleSynthesizerClassReference_
}

var moduleSynthesizerClassReference_ = &moduleSynthesizerClass_{
	// Initialize the class constants.
	warningMessage_: `
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│              This "Module.go" file was automatically generated.              │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘
`,

	packageAlias_: `
	<~packageAcronym> "<packagePath>"`,

	importedPackage_: `
	<~packageAcronym> "<moduleName>/<~packageName>"`,

	packageAliases_: `
// <~PackageName><TypeAliases>
`,

	typeAliases_: `

type (<NameAliases>)`,

	constAliases_: `

const (<NameAliases>)`,

	nameAlias_: `
	<Name> = <~packageAcronym>.<Name>`,

	classConstructors_: `
// <~PackageName>/<~ClassName><Constructors>
`,

	constructorFunction_: `

func <MethodName><Constraints>(<Parameters>) <~packageAcronym>.<~ClassName>Like<Arguments> {
	return <~packageAcronym>.<~ClassName>Class<Arguments>().<MethodName>(<ParameterNames>)
}`,

	methodParameter_: `
	<parameterName_> <ParameterType>,`,
}
