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
	mod "github.com/craterdog/go-class-model/v7"
	ana "github.com/craterdog/go-code-generation/v7/analyzer"
	com "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	reg "regexp"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ModuleSynthesizerClass() ModuleSynthesizerClassLike {
	return moduleSynthesizerClass()
}

// Constructor Methods

func (c *moduleSynthesizerClass_) ModuleSynthesizer(
	moduleName string,
	models com.CatalogLike[string, mod.ModelLike],
) ModuleSynthesizerLike {
	var instance = &moduleSynthesizer_{
		// Initialize the instance attributes.
		moduleName_: moduleName,
		models_:     models,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *moduleSynthesizer_) GetClass() ModuleSynthesizerClassLike {
	return moduleSynthesizerClass()
}

// TemplateDriven Methods

func (v *moduleSynthesizer_) CreateLegalNotice() string {
	var models = v.models_.GetIterator()
	var model = models.GetNext().GetValue()
	var packageDeclaration = model.GetPackageDeclaration()
	var legalNotice = packageDeclaration.GetLegalNotice().GetComment()
	return legalNotice
}

func (v *moduleSynthesizer_) CreateWarningMessage() string {
	var class = moduleSynthesizerClass()
	var warningMessage = class.warningMessage_
	return warningMessage
}

func (v *moduleSynthesizer_) CreateImportedPackages() string {
	var class = moduleSynthesizerClass()
	var importedPackages = class.importedPackages_
	var models = v.models_.GetIterator()
	for models.HasNext() {
		var association = models.GetNext()
		var packageName = association.GetKey()
		var packageAcronym = packageName[0:3]
		var packagePath = v.moduleName_ + "/" + packageName
		importedPackages += "\t" + packageAcronym + " \"" + packagePath + "\"\n"
	}
	return importedPackages
}

func (v *moduleSynthesizer_) CreateTypeAliases() string {
	var typeAliases string
	var models = v.models_.GetIterator()
	for models.HasNext() {
		var association = models.GetNext()
		var packageName = association.GetKey()
		var model = association.GetValue()
		var packageAliases = v.createPackageAliases(
			packageName,
			model,
		)
		typeAliases += packageAliases
	}
	return typeAliases
}

func (v *moduleSynthesizer_) CreateClassAccessors() string {
	var classAccessors string
	var models = v.models_.GetIterator()
	for models.HasNext() {
		var association = models.GetNext()
		var packageName = association.GetKey()
		var model = association.GetValue()
		classAccessors += v.createAccessFunctions(
			packageName,
			model,
		)
	}
	return classAccessors
}

func (v *moduleSynthesizer_) PerformGlobalUpdates(
	existing string,
	generated string,
) string {
	if uti.IsDefined(existing) {
		generated = v.preserveExistingCode(existing, generated)
	}
	return generated
}

// PROTECTED INTERFACE

// Private Methods

func (v *moduleSynthesizer_) createAspectAliases(
	aspectDeclarations com.ListLike[mod.AspectDeclarationLike],
	model mod.ModelLike,
) (
	aspectAliases string,
) {
	var nameAliases string
	var aspects = aspectDeclarations.GetIterator()
	for aspects.HasNext() {
		var declaration = aspects.GetNext().GetDeclaration()
		var nameAlias = v.createNameAlias(declaration, model)
		nameAliases += nameAlias
	}
	if uti.IsDefined(nameAliases) {
		nameAliases += "\n"
		var class = moduleSynthesizerClass()
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
	classDeclarations com.ListLike[mod.ClassDeclarationLike],
	model mod.ModelLike,
) (
	classAliases string,
) {
	var nameAliases string
	var classes = classDeclarations.GetIterator()
	for classes.HasNext() {
		var declaration = classes.GetNext().GetDeclaration()
		nameAliases += v.createNameAlias(declaration, model)
	}
	if uti.IsDefined(nameAliases) {
		nameAliases += "\n"
		var class = moduleSynthesizerClass()
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
	model mod.ModelLike,
) string {
	var class = moduleSynthesizerClass()
	var constructorFunction = class.constructorFunction_
	var methodName = constructorMethod.GetName()
	var parameters = v.extractParameters(constructorMethod, model)
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

func (v *moduleSynthesizer_) createClassFunctions(
	model mod.ModelLike,
	classDeclaration mod.ClassDeclarationLike,
) string {
	var class = moduleSynthesizerClass()
	var className = sts.TrimSuffix(
		classDeclaration.GetDeclaration().GetName(),
		"ClassLike",
	)
	className = uti.MakeLowerCase(className)
	var classFunctions = class.classFunctions_
	var analyzerClass = ana.ClassAnalyzerClass()
	var analyzer = analyzerClass.ClassAnalyzer(model, className)
	var constructorMethods = analyzer.GetConstructorMethods().GetIterator()
	for constructorMethods.HasNext() {
		var constructorMethod = constructorMethods.GetNext()
		var constructorFunction = v.createConstructorFunction(
			constructorMethod,
			className,
			model,
		)
		classFunctions += constructorFunction
	}
	classFunctions = uti.ReplaceAll(
		classFunctions,
		"className",
		className,
	)
	var constraints = analyzer.GetTypeConstraints()
	classFunctions = uti.ReplaceAll(
		classFunctions,
		"constraints",
		constraints,
	)
	var arguments = analyzer.GetTypeArguments()
	classFunctions = uti.ReplaceAll(
		classFunctions,
		"arguments",
		arguments,
	)

	return classFunctions
}

func (v *moduleSynthesizer_) createAccessFunctions(
	packageName string,
	model mod.ModelLike,
) string {
	var class = moduleSynthesizerClass()
	var accessFunctions = class.accessFunctions_
	var interfaceDeclarations = model.GetInterfaceDeclarations()
	var classSection = interfaceDeclarations.GetClassSection()
	var classes = classSection.GetClassDeclarations().GetIterator()
	for classes.HasNext() {
		var class = classes.GetNext()
		accessFunctions += v.createClassFunctions(model, class)
	}
	accessFunctions = v.replacePackageAttributes(
		accessFunctions,
		packageName,
	)
	return accessFunctions
}

func (v *moduleSynthesizer_) createEnumeratedAliases(
	enumeratedValues com.ListLike[string],
	model mod.ModelLike,
) (
	enumeratedAliases string,
) {
	var constantAliases string
	var class = moduleSynthesizerClass()
	var constants = enumeratedValues.GetIterator()
	for constants.HasNext() {
		var constant = constants.GetNext()
		var constantAlias = class.constantAlias_
		constantAlias = uti.ReplaceAll(
			constantAlias,
			"constant",
			constant,
		)
		constantAliases += constantAlias
	}
	if uti.IsDefined(constantAliases) {
		constantAliases += "\n"
		enumeratedAliases = class.enumeratedAliases_
		enumeratedAliases = uti.ReplaceAll(
			enumeratedAliases,
			"constantAliases",
			constantAliases,
		)
	}
	return
}

func (v *moduleSynthesizer_) createFunctionalAliases(
	functionalDeclarations com.ListLike[mod.FunctionalDeclarationLike],
	model mod.ModelLike,
) (
	functionalAliases string,
) {
	var nameAliases string
	var functionals = functionalDeclarations.GetIterator()
	for functionals.HasNext() {
		var declaration = functionals.GetNext().GetDeclaration()
		var nameAlias = v.createNameAlias(declaration, model)
		nameAliases += nameAlias
	}
	if uti.IsDefined(nameAliases) {
		nameAliases += "\n"
		var class = moduleSynthesizerClass()
		functionalAliases = class.typeAliases_
		functionalAliases = uti.ReplaceAll(
			functionalAliases,
			"nameAliases",
			nameAliases,
		)
	}
	return
}

func (v *moduleSynthesizer_) createInstanceAliases(
	instanceDeclarations com.ListLike[mod.InstanceDeclarationLike],
	model mod.ModelLike,
) (
	instanceAliases string,
) {
	var nameAliases string
	var instances = instanceDeclarations.GetIterator()
	for instances.HasNext() {
		var declaration = instances.GetNext().GetDeclaration()
		nameAliases += v.createNameAlias(declaration, model)
	}
	if uti.IsDefined(nameAliases) {
		nameAliases += "\n"
		var class = moduleSynthesizerClass()
		instanceAliases = class.typeAliases_
		instanceAliases = uti.ReplaceAll(
			instanceAliases,
			"nameAliases",
			nameAliases,
		)
	}
	return
}

func (v *moduleSynthesizer_) createNameAlias(
	declaration mod.DeclarationLike,
	model mod.ModelLike,
) (
	nameAlias string,
) {
	var class = moduleSynthesizerClass()
	nameAlias = class.nameAlias_
	var name = declaration.GetName()
	nameAlias = uti.ReplaceAll(
		nameAlias,
		"name",
		name,
	)
	var constraints string
	var arguments string
	var optionalConstraints = declaration.GetOptionalConstraints()
	if uti.IsDefined(optionalConstraints) {
		constraints = "["
		arguments = "["
		var constraint = optionalConstraints.GetConstraint()
		var constraintName = constraint.GetName()
		var constraintType = v.extractType(
			constraint.GetAbstraction(),
			model,
		)
		constraints += constraintName + " " + constraintType
		arguments += constraintName
		var additionalConstraints = optionalConstraints.GetAdditionalConstraints().GetIterator()
		for additionalConstraints.HasNext() {
			constraint = additionalConstraints.GetNext().GetConstraint()
			constraintName = constraint.GetName()
			constraintType = v.extractType(
				constraint.GetAbstraction(),
				model,
			)
			constraints += ", " + constraintName + " " + constraintType
			arguments += ", " + constraintName
		}
		constraints += "]"
		arguments += "]"
	}
	nameAlias = uti.ReplaceAll(
		nameAlias,
		"constraints",
		constraints,
	)
	nameAlias = uti.ReplaceAll(
		nameAlias,
		"arguments",
		arguments,
	)
	return
}

func (v *moduleSynthesizer_) createPackageAliases(
	packageName string,
	model mod.ModelLike,
) string {
	var analyzer = ana.PackageAnalyzerClass().PackageAnalyzer(model)
	var typeAliases = v.createTypeAliases(
		analyzer.GetTypeDeclarations(),
		model,
	)
	typeAliases += v.createEnumeratedAliases(
		analyzer.GetEnumeratedValues(),
		model,
	)

	typeAliases += v.createFunctionalAliases(
		analyzer.GetFunctionalDeclarations(),
		model,
	)
	typeAliases += v.createClassAliases(
		analyzer.GetClassDeclarations(),
		model,
	)
	typeAliases += v.createInstanceAliases(
		analyzer.GetInstanceDeclarations(),
		model,
	)
	typeAliases += v.createAspectAliases(
		analyzer.GetAspectDeclarations(),
		model,
	)
	var class = moduleSynthesizerClass()
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
	typeDeclarations com.ListLike[mod.TypeDeclarationLike],
	model mod.ModelLike,
) (
	typeAliases string,
) {
	var nameAliases string
	var types = typeDeclarations.GetIterator()
	for types.HasNext() {
		var typeDeclaration = types.GetNext().GetDeclaration()
		nameAliases += v.createNameAlias(typeDeclaration, model)
	}
	if uti.IsDefined(nameAliases) {
		nameAliases += "\n"
		var class = moduleSynthesizerClass()
		typeAliases = class.typeAliases_
		typeAliases = uti.ReplaceAll(
			typeAliases,
			"nameAliases",
			nameAliases,
		)
	}
	return
}

func (v *moduleSynthesizer_) extractParameterNames(
	constructorMethod mod.ConstructorMethodLike,
) string {
	var parameterNames string
	var parameterList = constructorMethod.GetOptionalParameterList()
	if uti.IsUndefined(parameterList) {
		return parameterNames
	}
	var parameters = parameterList.GetParameters().GetIterator()
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
	model mod.ModelLike,
) string {
	var methodParameters string
	var parameterList = constructorMethod.GetOptionalParameterList()
	if uti.IsUndefined(parameterList) {
		return methodParameters
	}
	var parameters = parameterList.GetParameters().GetIterator()
	for parameters.HasNext() {
		var parameter = parameters.GetNext()
		var parameterName = parameter.GetName()
		var parameterType = v.extractType(
			parameter.GetAbstraction(),
			model,
		)
		var class = moduleSynthesizerClass()
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
	model mod.ModelLike,
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
	var analyzer = ana.PackageAnalyzerClass().PackageAnalyzer(model)
	var packageAcronym = analyzer.GetPackageName()[0:3] + "."
	var name = abstraction.GetName()
	var prefix = abstraction.GetOptionalPrefix()
	if uti.IsUndefined(prefix) {
		var typeDeclarations = analyzer.GetTypeDeclarations()
		var types = typeDeclarations.GetIterator()
		for types.HasNext() {
			var typeName = types.GetNext().GetDeclaration().GetName()
			if typeName == name {
				prefix = packageAcronym
				break
			}
		}
		var enumeratedValues = analyzer.GetEnumeratedValues()
		var enumerations = enumeratedValues.GetIterator()
		for enumerations.HasNext() {
			var enumerationName = enumerations.GetNext()
			if enumerationName == name {
				prefix = packageAcronym
				break
			}
		}
		var functionalDeclarations = analyzer.GetFunctionalDeclarations()
		var functionals = functionalDeclarations.GetIterator()
		for functionals.HasNext() {
			var functionalName = functionals.GetNext().GetDeclaration().GetName()
			if functionalName == name {
				prefix = packageAcronym
				break
			}
		}
		var classDeclarations = analyzer.GetClassDeclarations()
		var classes = classDeclarations.GetIterator()
		for classes.HasNext() {
			var className = classes.GetNext().GetDeclaration().GetName()
			if className == name {
				prefix = packageAcronym
				break
			}
		}
		var instanceDeclarations = analyzer.GetInstanceDeclarations()
		var instances = instanceDeclarations.GetIterator()
		for instances.HasNext() {
			var instanceName = instances.GetNext().GetDeclaration().GetName()
			if instanceName == name {
				prefix = packageAcronym
				break
			}
		}
		var aspectDeclarations = analyzer.GetAspectDeclarations()
		var aspects = aspectDeclarations.GetIterator()
		for aspects.HasNext() {
			var aspectName = aspects.GetNext().GetDeclaration().GetName()
			if aspectName == name {
				prefix = packageAcronym
				break
			}
		}
	}
	abstractType += prefix
	abstractType += name
	var arguments = abstraction.GetOptionalArguments()
	if uti.IsDefined(arguments) {
		var argument = v.extractType(
			arguments.GetArgument().GetAbstraction(),
			model,
		)
		abstractType += "[" + argument
		var additionalArguments = arguments.GetAdditionalArguments().GetIterator()
		for additionalArguments.HasNext() {
			var additionalArgument = additionalArguments.GetNext().GetArgument()
			argument = v.extractType(
				additionalArgument.GetAbstraction(),
				model,
			)
			abstractType += ", " + argument
		}
		abstractType += "]"
	}
	return abstractType
}

func (v *moduleSynthesizer_) preserveExistingCode(
	existing string,
	generated string,
) string {
	// Preserve the module description.
	var pattern = `└──────────────────────────────────────────────────────────────────────────────┘(.|\r?\n)+package module`
	generated = v.replacePattern(pattern, existing, generated)

	// Preserve the global functions.
	pattern = `// GLOBAL FUNCTIONS(.|\r?\n)*`
	generated = v.replacePattern(pattern, existing, generated)
	return generated
}

func (v *moduleSynthesizer_) replacePackageAttributes(
	generated string,
	packageName string,
) string {
	generated = uti.ReplaceAll(
		generated,
		"packageName",
		packageName,
	)
	var packageAcronym = packageName[0:3]
	generated = uti.ReplaceAll(
		generated,
		"packageAcronym",
		packageAcronym,
	)
	return generated
}

func (v *moduleSynthesizer_) replacePattern(
	pattern string,
	existing string,
	generated string,
) string {
	var matcher = reg.MustCompile(pattern)
	var existingPattern = matcher.FindString(existing)
	var generatedPattern = matcher.FindString(generated)
	generated = sts.ReplaceAll(
		generated,
		generatedPattern,
		existingPattern,
	)
	return generated
}

// Instance Structure

type moduleSynthesizer_ struct {
	// Declare the instance attributes.
	moduleName_ string
	models_     com.CatalogLike[string, mod.ModelLike]
}

// Class Structure

type moduleSynthesizerClass_ struct {
	// Declare the class constants.
	warningMessage_      string
	importedPackages_    string
	packageAliases_      string
	typeAliases_         string
	enumeratedAliases_   string
	constantAlias_       string
	nameAlias_           string
	accessFunctions_     string
	classFunctions_      string
	constructorFunction_ string
	methodParameter_     string
}

// Class Reference

func moduleSynthesizerClass() *moduleSynthesizerClass_ {
	return moduleSynthesizerClassReference_
}

var moduleSynthesizerClassReference_ = &moduleSynthesizerClass_{
	// Initialize the class constants.
	warningMessage_: `
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│             This "module_api.go" file was automatically generated.           │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘
`,

	importedPackages_: `
	fmt "fmt"
	com "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	ref "reflect"
	sts "strings"
`,

	packageAliases_: `
// <~PackageName><TypeAliases>
`,

	typeAliases_: `

type (<NameAliases>)`,

	enumeratedAliases_: `

const (<ConstantAliases>)`,

	constantAlias_: `
	<Constant> = <~packageAcronym>.<Constant>`,

	nameAlias_: `
	<Name><Constraints> = <~packageAcronym>.<Name><Arguments>`,

	accessFunctions_: `
// <~PackageName>
`,

	classFunctions_: `

func <~ClassName>Class<Constraints>() <~ClassName>ClassLike<Arguments> {
	return <~packageAcronym>.<~ClassName>Class<Arguments>()
}`,

	constructorFunction_: `

func <MethodName><Constraints>(<Parameters>) <~ClassName>Like<Arguments> {
	return <~ClassName>Class<Arguments>().<MethodName>(<ParameterNames>)
}`,

	methodParameter_: `
	<parameterName_> <ParameterType>,`,
}
