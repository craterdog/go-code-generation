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
	stc "strconv"
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

func (v *moduleSynthesizer_) createArgumentCases(
	constructorMethods abs.ListLike[mod.ConstructorMethodLike],
) string {
	var argumentCases string
	var argumentTypes = col.Set[string]()
	var arguments = v.extractArguments(constructorMethods).GetIterator()
	for arguments.HasNext() {
		var argument = arguments.GetNext()
		var argumentType = v.extractType(argument.GetValue())
		if argumentTypes.ContainsValue(argumentType) {
			// This type has already been found.
			continue
		}
		var argumentCase = moduleSynthesizerReference().argumentCase_
		argumentCase = uti.ReplaceAll(
			argumentCase,
			"argumentType",
			argumentType,
		)
		argumentTypes.AddValue(argumentType)
		argumentCases += argumentCase
	}
	return argumentCases
}

func (v *moduleSynthesizer_) createConstructionCases(
	constructorMethods abs.ListLike[mod.ConstructorMethodLike],
) string {
	var constructionCases string
	var methods = constructorMethods.GetIterator()
	for methods.HasNext() {
		var method = methods.GetNext()
		var argumentNames = v.extractArgumentNames(method)
		var argumentTypes = v.extractArgumentTypes(method)
		var argumentAssignments = v.createArgumentAssignments(method)
		var constructionCase = moduleSynthesizerReference().constructionCase_
		constructionCase = uti.ReplaceAll(
			constructionCase,
			"argumentNames",
			argumentNames,
		)
		constructionCase = uti.ReplaceAll(
			constructionCase,
			"argumentTypes",
			argumentTypes,
		)
		constructionCase = uti.ReplaceAll(
			constructionCase,
			"argumentAssignments",
			argumentAssignments,
		)
		constructionCases += constructionCase
	}
	return constructionCases
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
	var argumentCases = v.createArgumentCases(constructorMethods)
	constructorFunction = uti.ReplaceAll(
		constructorFunction,
		"argumentCases",
		argumentCases,
	)
	var constructionCases = v.createConstructionCases(constructorMethods)
	constructorFunction = uti.ReplaceAll(
		constructorFunction,
		"constructionCases",
		constructionCases,
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
	declaration mod.DeclarationLike,
) string {
	var typeAlias string
	var typeName = declaration.GetName()
	if uti.IsDefined(declaration.GetOptionalConstraints()) {
		// Type aliases are not supported for generic types in Go.
		return typeAlias
	}
	typeAlias = moduleSynthesizerReference().typeAlias_
	typeAlias = uti.ReplaceAll(
		typeAlias,
		"typeName",
		typeName,
	)
	return typeAlias
}

func (v *moduleSynthesizer_) createPackageAliases(
	packageName string,
	model mod.ModelLike,
) string {
	var typeAliases string
	var primitiveDeclarations = model.GetPrimitiveDeclarations()
	var typeSection = primitiveDeclarations.GetOptionalTypeSection()
	if uti.IsDefined(typeSection) {
		var types = typeSection.GetTypeDeclarations().GetIterator()
		for types.HasNext() {
			var declaration = types.GetNext().GetDeclaration()
			var typeAlias = v.createTypeAlias(model, declaration)
			typeAliases += typeAlias
		}
	}
	var functionalSection = primitiveDeclarations.GetOptionalFunctionalSection()
	if uti.IsDefined(functionalSection) {
		var functionals = functionalSection.GetFunctionalDeclarations().GetIterator()
		for functionals.HasNext() {
			var declaration = functionals.GetNext().GetDeclaration()
			var typeAlias = v.createTypeAlias(model, declaration)
			typeAliases += typeAlias
		}
	}
	var interfaceDeclarations = model.GetInterfaceDeclarations()
	var instanceSection = interfaceDeclarations.GetInstanceSection()
	var instances = instanceSection.GetInstanceDeclarations().GetIterator()
	for instances.HasNext() {
		var declaration = instances.GetNext().GetDeclaration()
		var typeAlias = v.createTypeAlias(model, declaration)
		typeAliases += typeAlias
	}
	var aspectSection = interfaceDeclarations.GetOptionalAspectSection()
	if uti.IsDefined(aspectSection) {
		var aspects = aspectSection.GetAspectDeclarations().GetIterator()
		for aspects.HasNext() {
			var declaration = aspects.GetNext().GetDeclaration()
			var typeAlias = v.createTypeAlias(model, declaration)
			typeAliases += typeAlias
		}
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
	var packageNames = v.models_.GetKeys().GetIterator()
	for packageNames.HasNext() {
		var packageName = packageNames.GetNext()
		var importedPackage = moduleSynthesizerReference().importedPackage_
		importedPackage = v.replacePackageAttributes(importedPackage, packageName)
		importedPackages += importedPackage
	}
	if sts.Contains(source, "fmt.") && !sts.Contains(importedPackages, "fmt") {
		var packageAlias = moduleSynthesizerReference().packageAlias_
		packageAlias = uti.ReplaceAll(
			packageAlias,
			"packageAcronym",
			"fmt",
		)
		packageAlias = uti.ReplaceAll(
			packageAlias,
			"packagePath",
			"\"fmt\"",
		)
		importedPackages += packageAlias
	}
	if sts.Contains(source, "abs.") && !sts.Contains(importedPackages, "abs") {
		var packageAlias = moduleSynthesizerReference().packageAlias_
		packageAlias = uti.ReplaceAll(
			packageAlias,
			"packageAcronym",
			"abs",
		)
		packageAlias = uti.ReplaceAll(
			packageAlias,
			"packagePath",
			"\"github.com/craterdog/go-collection-framework/v4/collection\"",
		)
		importedPackages += packageAlias
	}
	return importedPackages
}

func (v *moduleSynthesizer_) createArgumentAssignments(
	constructorMethod mod.ConstructorMethodLike,
) string {
	var argumentAssignments string
	var index int
	var parameters = constructorMethod.GetParameters().GetIterator()
	for parameters.HasNext() {
		var parameter = parameters.GetNext()
		var argumentName = parameter.GetName()
		var argumentType = v.extractType(parameter.GetAbstraction())
		var argumentAssignment = moduleSynthesizerReference().argumentAssignment_
		argumentAssignment = uti.ReplaceAll(
			argumentAssignment,
			"argumentName",
			argumentName,
		)
		argumentAssignment = uti.ReplaceAll(
			argumentAssignment,
			"argumentType",
			argumentType,
		)
		argumentAssignment = uti.ReplaceAll(
			argumentAssignment,
			"index",
			stc.Itoa(index),
		)
		argumentAssignments += argumentAssignment
		index++
	}
	return argumentAssignments
}

func (v *moduleSynthesizer_) extractArgumentTypes(
	constructorMethod mod.ConstructorMethodLike,
) string {
	var argumentTypes string
	var parameters = constructorMethod.GetParameters().GetIterator()
	if parameters.IsEmpty() {
		return argumentTypes
	}
	var parameter = parameters.GetNext()
	var argumentType = v.extractType(parameter.GetAbstraction())
	argumentTypes += argumentType
	for parameters.HasNext() {
		parameter = parameters.GetNext()
		argumentType = v.extractType(parameter.GetAbstraction())
		argumentTypes += ", " + argumentType
	}
	return argumentTypes
}

func (v *moduleSynthesizer_) extractArgumentNames(
	constructorMethod mod.ConstructorMethodLike,
) string {
	var argumentNames string
	var parameters = constructorMethod.GetParameters().GetIterator()
	if parameters.IsEmpty() {
		return argumentNames
	}
	for parameters.HasNext() {
		var parameter = parameters.GetNext()
		var argumentName = parameter.GetName()
		argumentNames += "\n\t\t\t\t" + argumentName + ","
	}
	argumentNames += "\n\t\t\t"
	return argumentNames
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
	moduleImports_         string
	packageAlias_          string
	importedPackage_       string
	packageAliases_        string
	typeAlias_             string
	universalConstructors_ string
	constructorFunction_   string
	argumentCase_          string
	constructionCase_      string
	argumentAssignment_    string
}

// Class Reference

func moduleSynthesizerReference() *moduleSynthesizerClass_ {
	return moduleSynthesizerReference_
}

var moduleSynthesizerReference_ = &moduleSynthesizerClass_{
	// Initialize the class constants.
	moduleImports_: `

import (<ImportedPackages>
)`,

	packageAlias_: `
	<~packageAcronym> <packagePath>`,

	importedPackage_: `
	<~packageAcronym> "<moduleName>/<~packageName>"`,

	packageAliases_: `

// <~PackageName>

type (<TypeAliases>
)`,

	typeAlias_: `
	<TypeName> = <~packageAcronym>.<TypeName>`,

	universalConstructors_: `

// <~PackageName><ConstructorFunctions>`,

	constructorFunction_: `

func <~ClassName>(arguments ...any) <~ClassName>Like {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {<ArgumentCases>
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the <~ClassName> constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma.
		argumentTypes = argumentTypes[:length - 1]
	}

	// Call the corresponding constructor.
	var <className_> <~ClassName>Like
	switch argumentTypes {<ConstructionCases>
		default:
			var message = fmt.Sprintf(
				"No <~ClassName> constructor matching the arguments was found: $v\n",
				arguments,
			)
			panic(message)
	}
	return <className_>
}`,

	argumentCase_: `
		case <ArgumentType>:
			argumentTypes += "<ArgumentType>, "`,

	constructionCase_: `
		case "<ArgumentTypes>":<ArgumentAssignments>
			<className_> = <~packageAcronym>.<~ClassName>().Make(<ArgumentNames>)`,

	argumentAssignment_: `
			var <argumentName_> = arguments[<index>].(<ArgumentType>)`,
}
