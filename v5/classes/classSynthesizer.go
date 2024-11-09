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

package classes

import (
	fmt "fmt"
	mod "github.com/craterdog/go-class-model/v5/ast"
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ClassSynthesizer() ClassSynthesizerClassLike {
	return classSynthesizerReference()
}

// Constructor Methods

func (c *classSynthesizerClass_) Make(
	model mod.ModelLike,
	className string,
) ClassSynthesizerLike {
	var instance = &classSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: Analyzer().Make(model, className),
	}
	return instance
}

// INSTANCE INTERFACE

// Primary Methods

func (v *classSynthesizer_) GetClass() ClassSynthesizerClassLike {
	return classSynthesizerReference()
}

// TemplateDriven Methods

func (v *classSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *classSynthesizer_) CreateAccessFunction() string {
	var accessFunction = classSynthesizerReference().accessFunction_
	return accessFunction
}

func (v *classSynthesizer_) CreateConstantMethods() string {
	var methods = v.analyzer_.GetConstantMethods()
	var constantMethods = v.createConstantMethods(methods)
	return constantMethods
}

func (v *classSynthesizer_) CreateConstructorMethods() string {
	var methods = v.analyzer_.GetConstructorMethods()
	var constructorMethods = v.createConstructorMethods(methods)
	return constructorMethods
}

func (v *classSynthesizer_) CreateFunctionMethods() string {
	var methods = v.analyzer_.GetFunctionMethods()
	var functionMethods = v.createFunctionMethods(methods)
	return functionMethods
}

func (v *classSynthesizer_) CreatePrimaryMethods() string {
	var methods = v.analyzer_.GetPrimaryMethods()
	var primaryMethods = v.createPrimaryMethods(methods)
	return primaryMethods
}

func (v *classSynthesizer_) CreateAttributeMethods() string {
	var methods = v.analyzer_.GetAttributeMethods()
	var attributeMethods = v.createAttributeMethods(methods)
	return attributeMethods
}

func (v *classSynthesizer_) CreateAspectMethods() string {
	var declarations = v.analyzer_.GetAspectDeclarations()
	var interfaces = v.analyzer_.GetAspectInterfaces()
	var aspectMethods = v.createAspectInterfaces(declarations, interfaces)
	return aspectMethods
}

func (v *classSynthesizer_) CreatePrivateMethods() string {
	var privateMethods = classSynthesizerReference().privateMethods_
	return privateMethods
}

func (v *classSynthesizer_) CreateInstanceStructure() string {
	var instanceStructure = v.createInstanceStructure()
	return instanceStructure
}

func (v *classSynthesizer_) CreateClassStructure() string {
	var classStructure = v.createClassStructure()
	return classStructure
}

func (v *classSynthesizer_) CreateClassReference() string {
	var classReference = v.createClassReference()
	return classReference
}

func (v *classSynthesizer_) PerformGlobalUpdates(
	source string,
) string {
	var class = v.performGlobalUpdates(source)
	return class
}

// PROTECTED INTERFACE

// Private Methods

func (v *classSynthesizer_) extractAttributeName(
	accessorName string,
) string {
	var attributeName string
	switch {
	case sts.HasPrefix(accessorName, "Get"):
		attributeName = sts.TrimPrefix(accessorName, "Get")
	case sts.HasPrefix(accessorName, "Set"):
		attributeName = sts.TrimPrefix(accessorName, "Set")
	case sts.HasPrefix(accessorName, "Is"):
		attributeName = sts.TrimPrefix(accessorName, "Is")
	case sts.HasPrefix(accessorName, "Was"):
		attributeName = sts.TrimPrefix(accessorName, "Was")
	case sts.HasPrefix(accessorName, "Are"):
		attributeName = sts.TrimPrefix(accessorName, "Are")
	case sts.HasPrefix(accessorName, "Were"):
		attributeName = sts.TrimPrefix(accessorName, "Were")
	case sts.HasPrefix(accessorName, "Has"):
		attributeName = sts.TrimPrefix(accessorName, "Has")
	case sts.HasPrefix(accessorName, "Had"):
		attributeName = sts.TrimPrefix(accessorName, "Had")
	case sts.HasPrefix(accessorName, "Have"):
		attributeName = sts.TrimPrefix(accessorName, "Have")
	default:
		var message = fmt.Sprintf(
			"An unknown accessor name was found: %q",
			accessorName,
		)
		panic(message)
	}
	attributeName = uti.MakeLowerCase(attributeName)
	return attributeName
}

func (v *classSynthesizer_) extractConcreteMappings(
	constraints mod.ConstraintsLike,
	arguments mod.ArgumentsLike,
) abs.CatalogLike[string, mod.AbstractionLike] {
	// Create the mappings catalog.
	var mappings = col.Catalog[string, mod.AbstractionLike]()
	if uti.IsUndefined(constraints) || uti.IsUndefined(arguments) {
		return mappings
	}

	// Map the name of the first constraint to its concrete type.
	var constraint = constraints.GetConstraint()
	var constraintName = constraint.GetName()
	var argument = arguments.GetArgument()
	var concreteType = argument.GetAbstraction()
	mappings.SetValue(constraintName, concreteType)

	// Map the name of the additional constraints to their concrete types.
	var additionalConstraints = constraints.GetAdditionalConstraints().GetIterator()
	var additionalArguments = arguments.GetAdditionalArguments().GetIterator()
	for additionalConstraints.HasNext() {
		constraint = additionalConstraints.GetNext().GetConstraint()
		constraintName = constraint.GetName()
		argument = additionalArguments.GetNext().GetArgument()
		concreteType = argument.GetAbstraction()
		mappings.SetValue(constraintName, concreteType)
	}

	return mappings
}

func (v *classSynthesizer_) extractType(
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

func (v *classSynthesizer_) createAspectInterface(
	declarations abs.Sequential[mod.AspectDeclarationLike],
	aspectInterface mod.AspectInterfaceLike,
) (
	implementation string,
) {
	var methods string
	var aspectType = aspectInterface.GetAbstraction()
	if uti.IsDefined(declarations) {
		var aspectDeclarations = declarations.GetIterator()
		for aspectDeclarations.HasNext() {
			var aspectDeclaration = aspectDeclarations.GetNext()
			var declaration = aspectDeclaration.GetDeclaration()
			var constraints = declaration.GetOptionalConstraints()
			var arguments = aspectType.GetOptionalArguments()
			if uti.IsUndefined(aspectType.GetOptionalSuffix()) &&
				declaration.GetName() == aspectType.GetName() {
				var mappings = v.extractConcreteMappings(constraints, arguments)
				methods = v.createAspectMethods(
					aspectType,
					aspectDeclaration,
					mappings,
				)
			}
		}
	}
	implementation = classSynthesizerReference().aspectInterface_
	implementation = uti.ReplaceAll(
		implementation,
		"aspectType",
		v.extractType(aspectType),
	)
	implementation = uti.ReplaceAll(
		implementation,
		"aspectMethods",
		methods,
	)
	return implementation
}

func (v *classSynthesizer_) createAspectInterfaces(
	declarations abs.Sequential[mod.AspectDeclarationLike],
	interfaces abs.Sequential[mod.AspectInterfaceLike],
) (
	implementation string,
) {
	if uti.IsDefined(interfaces) {
		var aspectInterfaces = interfaces.GetIterator()
		for aspectInterfaces.HasNext() {
			var aspectInterface = aspectInterfaces.GetNext()
			implementation += v.createAspectInterface(declarations, aspectInterface)
		}
	}
	return implementation
}

func (v *classSynthesizer_) createAspectMethod(
	aspectType mod.AbstractionLike,
	aspectMethod mod.AspectMethodLike,
	mappings abs.CatalogLike[string, mod.AbstractionLike],
) (
	implementation string,
) {
	var method = aspectMethod.GetMethod()
	var methodName = method.GetName()
	var methodParameters = method.GetParameters()
	var methodResult = method.GetOptionalResult()
	if mappings.GetSize() > 0 {
		methodParameters = v.replaceParameterTypes(method.GetParameters(), mappings)
		if uti.IsDefined(methodResult) {
			methodResult = v.replaceResultType(method.GetOptionalResult(), mappings)
		}
	}
	var parameters = v.createParameters(methodParameters)
	var resultType = v.createResult(methodResult)
	implementation = classSynthesizerReference().instanceMethod_
	if uti.IsDefined(resultType) {
		implementation = classSynthesizerReference().instanceFunction_
		implementation = uti.ReplaceAll(implementation, "resultType", resultType)
	}
	implementation = uti.ReplaceAll(implementation, "methodName", methodName)
	implementation = uti.ReplaceAll(implementation, "parameters", parameters)
	return implementation
}

func (v *classSynthesizer_) createAspectMethods(
	aspectType mod.AbstractionLike,
	aspectDeclaration mod.AspectDeclarationLike,
	mappings abs.CatalogLike[string, mod.AbstractionLike],
) (
	implementation string,
) {
	var aspectMethods = aspectDeclaration.GetAspectMethods().GetIterator()
	for aspectMethods.HasNext() {
		var aspectMethod = aspectMethods.GetNext()
		implementation += v.createAspectMethod(
			aspectType,
			aspectMethod,
			mappings,
		)
	}
	return implementation
}

func (v *classSynthesizer_) createAttributeCheck(
	parameter mod.ParameterLike,
) (
	implementation string,
) {
	var parameterName = parameter.GetName()
	var attributeName = sts.TrimSuffix(parameterName, "_")
	// Ignore optional attributes.
	if !sts.HasPrefix(attributeName, "optional") {
		var template = classSynthesizerReference().attributeCheck_
		template = uti.ReplaceAll(template, "attributeName", attributeName)
		implementation += template
	}
	return implementation
}

func (v *classSynthesizer_) createAttributeChecks(
	parameters abs.Sequential[mod.ParameterLike],
) (
	implementation string,
) {
	var iterator = parameters.GetIterator()
	for iterator.HasNext() {
		var parameter = iterator.GetNext()
		var attributeCheck = v.createAttributeCheck(parameter)
		implementation += attributeCheck
	}
	return implementation
}

func (v *classSynthesizer_) createAttributeDeclarations() (
	implementation string,
) {
	var attributes = v.analyzer_.GetAttributes().GetIterator()
	for attributes.HasNext() {
		var attribute = attributes.GetNext()
		var attributeName = attribute.GetKey()
		var attributeType = attribute.GetValue()
		var declaration = classSynthesizerReference().attributeDeclaration_
		declaration = uti.ReplaceAll(declaration, "attributeName", attributeName)
		declaration = uti.ReplaceAll(declaration, "attributeType", attributeType)
		implementation += declaration
	}
	return implementation
}

func (v *classSynthesizer_) createAttributeInitializations(
	parameters abs.Sequential[mod.ParameterLike],
) (
	implementation string,
) {
	var iterator = parameters.GetIterator()
	for iterator.HasNext() {
		var parameter = iterator.GetNext()
		var parameterName = parameter.GetName()
		var attributeName = sts.TrimSuffix(parameterName, "_")
		if uti.IsDefined(v.analyzer_.GetAttributes().GetValue(attributeName)) {
			var template = classSynthesizerReference().attributeInitialization_
			template = uti.ReplaceAll(template, "attributeName", attributeName)
			implementation += template
		}
	}
	return implementation
}

func (v *classSynthesizer_) createAttributeMethods(
	attributes abs.Sequential[mod.AttributeMethodLike],
) (
	implementation string,
) {
	if uti.IsDefined(attributes) {
		var methods string
		var attributeMethods = attributes.GetIterator()
		for attributeMethods.HasNext() {
			var method string
			var attributeMethod = attributeMethods.GetNext()
			switch actual := attributeMethod.GetAny().(type) {
			case mod.GetterMethodLike:
				method = v.createGetterMethod(actual)
			case mod.SetterMethodLike:
				method = v.createSetterMethod(actual)
			}
			methods += method
		}
		implementation = classSynthesizerReference().attributeMethods_
		implementation = uti.ReplaceAll(implementation, "attributeMethods", methods)
	}
	return implementation
}

func (v *classSynthesizer_) createClassReference() (
	implementation string,
) {
	implementation = classSynthesizerReference().classReference_
	if v.analyzer_.IsGeneric() {
		implementation = classSynthesizerReference().classMap_
	}
	var constantInitializations = v.createConstantInitializations()
	implementation = uti.ReplaceAll(
		implementation,
		"constantInitializations",
		constantInitializations,
	)
	return implementation
}

func (v *classSynthesizer_) createClassStructure() (
	implementation string,
) {
	implementation = classSynthesizerReference().classStructure_
	var constantDeclarations = v.createConstantDeclarations()
	implementation = uti.ReplaceAll(
		implementation,
		"constantDeclarations",
		constantDeclarations,
	)
	return implementation
}

func (v *classSynthesizer_) createConstantDeclarations() (
	implementation string,
) {
	var constants = v.analyzer_.GetConstants().GetIterator()
	for constants.HasNext() {
		var constant = constants.GetNext()
		var constantName = constant.GetKey()
		var constantType = constant.GetValue()
		var declaration = classSynthesizerReference().constantDeclaration_
		declaration = uti.ReplaceAll(declaration, "constantName", constantName)
		declaration = uti.ReplaceAll(declaration, "constantType", constantType)
		implementation += declaration
	}
	return implementation
}

func (v *classSynthesizer_) createConstantInitializations() (
	implementation string,
) {
	var constants = v.analyzer_.GetConstants().GetIterator()
	for constants.HasNext() {
		var constant = constants.GetNext()
		var constantName = constant.GetKey()
		var initialization = classSynthesizerReference().constantInitialization_
		initialization = uti.ReplaceAll(initialization, "constantName", constantName)
		implementation += initialization
	}
	return implementation
}

func (v *classSynthesizer_) createConstantMethod(
	constantMethod mod.ConstantMethodLike,
) (
	implementation string,
) {
	var methodName = constantMethod.GetName()
	var resultType = v.extractType(constantMethod.GetAbstraction())
	implementation = classSynthesizerReference().constantMethod_
	implementation = uti.ReplaceAll(implementation, "methodName", methodName)
	implementation = uti.ReplaceAll(implementation, "resultType", resultType)
	return implementation
}

func (v *classSynthesizer_) createConstantMethods(
	constants abs.Sequential[mod.ConstantMethodLike],
) (
	implementation string,
) {
	if uti.IsDefined(constants) {
		var methods string
		var constantMethods = constants.GetIterator()
		for constantMethods.HasNext() {
			var constantMethod = constantMethods.GetNext()
			methods += v.createConstantMethod(constantMethod)
		}
		implementation = classSynthesizerReference().constantMethods_
		implementation = uti.ReplaceAll(implementation, "constantMethods", methods)
	}
	return implementation
}

func (v *classSynthesizer_) createConstructorMethod(
	constructorMethod mod.ConstructorMethodLike,
) (
	implementation string,
) {
	var methodName = constructorMethod.GetName()
	var constructorParameters = constructorMethod.GetParameters()
	var parameters = v.createParameters(constructorParameters)
	var resultType = v.extractType(constructorMethod.GetAbstraction())
	var instanceInstantiation = v.createInstanceInstantiation(constructorMethod)
	implementation = classSynthesizerReference().constructorMethod_
	implementation = uti.ReplaceAll(
		implementation,
		"methodName",
		methodName,
	)
	implementation = uti.ReplaceAll(
		implementation,
		"parameters",
		parameters,
	)
	implementation = uti.ReplaceAll(
		implementation,
		"resultType",
		resultType,
	)
	implementation = uti.ReplaceAll(
		implementation,
		"instanceInstantiation",
		instanceInstantiation,
	)
	return implementation
}

func (v *classSynthesizer_) createConstructorMethods(
	constructors abs.Sequential[mod.ConstructorMethodLike],
) (
	implementation string,
) {
	if uti.IsDefined(constructors) {
		var methods string
		var constructorMethods = constructors.GetIterator()
		for constructorMethods.HasNext() {
			var constructorMethod = constructorMethods.GetNext()
			methods += v.createConstructorMethod(constructorMethod)
		}
		implementation = classSynthesizerReference().constructorMethods_
		implementation = uti.ReplaceAll(implementation, "constructorMethods", methods)
	}
	return implementation
}

func (v *classSynthesizer_) createFunctionMethod(
	functionMethod mod.FunctionMethodLike,
) (
	implementation string,
) {
	var methodName = functionMethod.GetName()
	var parameters = v.createParameters(functionMethod.GetParameters())
	var resultType = v.createResult(functionMethod.GetResult())
	implementation = classSynthesizerReference().functionMethod_
	implementation = uti.ReplaceAll(implementation, "methodName", methodName)
	implementation = uti.ReplaceAll(implementation, "parameters", parameters)
	implementation = uti.ReplaceAll(implementation, "resultType", resultType)
	return implementation
}

func (v *classSynthesizer_) createFunctionMethods(
	functions abs.Sequential[mod.FunctionMethodLike],
) (
	implementation string,
) {
	if uti.IsDefined(functions) {
		var methods string
		var functionMethods = functions.GetIterator()
		for functionMethods.HasNext() {
			var functionMethod = functionMethods.GetNext()
			methods += v.createFunctionMethod(functionMethod)
		}
		implementation = classSynthesizerReference().functionMethods_
		implementation = uti.ReplaceAll(implementation, "functionMethods", methods)
	}
	return implementation
}

func (v *classSynthesizer_) createGetterMethod(
	getterMethod mod.GetterMethodLike,
) (
	implementation string,
) {
	var methodName = getterMethod.GetName()
	var attributeName = v.extractAttributeName(methodName)
	var attributeType = v.extractType(getterMethod.GetAbstraction())
	implementation = classSynthesizerReference().getterMethod_
	implementation = uti.ReplaceAll(implementation, "methodName", methodName)
	implementation = uti.ReplaceAll(implementation, "attributeName", attributeName)
	implementation = uti.ReplaceAll(implementation, "attributeType", attributeType)
	return implementation
}

func (v *classSynthesizer_) createInstanceInstantiation(
	constructorMethod mod.ConstructorMethodLike,
) (
	implementation string,
) {
	var methodName = constructorMethod.GetName()
	implementation = classSynthesizerReference().instanceInstantiation_
	if v.analyzer_.IsIntrinsic() {
		if methodName == "Make" {
			implementation = classSynthesizerReference().intrinsicInstantiation_
		}
	} else {
		if methodName == "Make" || sts.HasPrefix(methodName, "MakeWith") {
			implementation = classSynthesizerReference().structureInstantiation_
			var constructorParameters = constructorMethod.GetParameters()
			var attributeChecks = v.createAttributeChecks(constructorParameters)
			var attributeInitializations = v.createAttributeInitializations(
				constructorParameters,
			)
			implementation = uti.ReplaceAll(
				implementation,
				"attributeChecks",
				attributeChecks,
			)
			implementation = uti.ReplaceAll(
				implementation,
				"attributeInitializations",
				attributeInitializations,
			)
		}
	}
	return implementation
}

func (v *classSynthesizer_) createInstanceStructure() (
	implementation string,
) {
	if v.analyzer_.IsIntrinsic() {
		var intrinsicType = v.extractType(v.analyzer_.GetIntrinsicType())
		implementation = classSynthesizerReference().instanceIntrinsic_
		implementation = uti.ReplaceAll(
			implementation,
			"intrinsicType",
			intrinsicType,
		)
	} else {
		implementation = classSynthesizerReference().instanceStructure_
		var attributeDeclarations = v.createAttributeDeclarations()
		implementation = uti.ReplaceAll(
			implementation,
			"attributeDeclarations",
			attributeDeclarations,
		)
	}
	return implementation
}

func (v *classSynthesizer_) createIntrinsicMethod() (
	implementation string,
) {
	if v.analyzer_.IsIntrinsic() {
		var intrinsicType = v.extractType(v.analyzer_.GetIntrinsicType())
		implementation = classSynthesizerReference().intrinsicMethod_
		implementation = uti.ReplaceAll(
			implementation,
			"intrinsicType",
			intrinsicType,
		)
	}
	return implementation
}

func (v *classSynthesizer_) createImportedPackages(
	moduleImports mod.ModuleImportsLike,
	class string,
) (
	implementation string,
) {
	if uti.IsDefined(moduleImports) {
		var importedPackages = moduleImports.GetImportedPackages().GetIterator()
		for importedPackages.HasNext() {
			var importedPackage = importedPackages.GetNext()
			var packageName = importedPackage.GetName()
			var prefix = packageName + "."
			var packagePath = importedPackage.GetPath()
			if sts.Contains(class, prefix) && !sts.Contains(implementation, prefix) {
				var packageAlias = classSynthesizerReference().packageAlias_
				packageAlias = uti.ReplaceAll(packageAlias, "packageName", packageName)
				packageAlias = uti.ReplaceAll(packageAlias, "packagePath", packagePath)
				implementation += packageAlias
			}
		}
	}
	if sts.Contains(class, "fmt.") && !sts.Contains(implementation, "fmt") {
		var packageAlias = classSynthesizerReference().packageAlias_
		packageAlias = uti.ReplaceAll(packageAlias, "packageName", "fmt")
		packageAlias = uti.ReplaceAll(packageAlias, "packagePath", "\"fmt\"")
		implementation += packageAlias
	}
	if sts.Contains(class, "uti.") && !sts.Contains(implementation, "uti") {
		var packageAlias = classSynthesizerReference().packageAlias_
		packageAlias = uti.ReplaceAll(packageAlias, "packageName", "uti")
		packageAlias = uti.ReplaceAll(packageAlias, "packagePath", "\"github.com/craterdog/go-missing-utilities/v2\"")
		implementation += packageAlias
	}
	if sts.Contains(class, "col.") && !sts.Contains(implementation, "col") {
		var packageAlias = classSynthesizerReference().packageAlias_
		packageAlias = uti.ReplaceAll(packageAlias, "packageName", "col")
		packageAlias = uti.ReplaceAll(packageAlias, "packagePath", "\"github.com/craterdog/go-collection-framework/v4\"")
		implementation += packageAlias
	}
	if sts.Contains(class, "abs.") && !sts.Contains(implementation, "abs") {
		var packageAlias = classSynthesizerReference().packageAlias_
		packageAlias = uti.ReplaceAll(packageAlias, "packageName", "abs")
		packageAlias = uti.ReplaceAll(packageAlias, "packagePath", "\"github.com/craterdog/go-collection-framework/v4/collection\"")
		implementation += packageAlias
	}
	if sts.Contains(class, "syn.") && !sts.Contains(implementation, "syn") {
		var packageAlias = classSynthesizerReference().packageAlias_
		packageAlias = uti.ReplaceAll(packageAlias, "packageName", "syn")
		packageAlias = uti.ReplaceAll(packageAlias, "packagePath", "\"sync\"")
		implementation += packageAlias
	}
	if uti.IsDefined(implementation) {
		implementation += "\n"
	}
	return implementation
}

func (v *classSynthesizer_) createModuleImports(
	imports mod.ModuleImportsLike,
	source string,
) (
	implementation string,
) {
	var importedPackages = v.createImportedPackages(imports, source)
	if uti.IsDefined(importedPackages) {
		implementation = classSynthesizerReference().moduleImports_
		implementation = uti.ReplaceAll(
			implementation,
			"importedPackages",
			importedPackages,
		)
	}
	return implementation
}

func (v *classSynthesizer_) createParameters(
	parameters abs.Sequential[mod.ParameterLike],
) (
	implementation string,
) {
	var iterator = parameters.GetIterator()
	for iterator.HasNext() {
		var parameter = iterator.GetNext()
		var parameterName = parameter.GetName()
		var parameterType = v.extractType(parameter.GetAbstraction())
		var template = classSynthesizerReference().methodParameter_
		template = uti.ReplaceAll(template, "parameterName", parameterName)
		template = uti.ReplaceAll(template, "parameterType", parameterType)
		implementation += template
	}
	if uti.IsDefined(implementation) {
		implementation += "\n"
	}
	return implementation
}

func (v *classSynthesizer_) createPrimaryMethod(
	primaryMethod mod.PrimaryMethodLike,
) (
	implementation string,
) {
	var method = primaryMethod.GetMethod()
	var methodName = method.GetName()
	var parameters = v.createParameters(method.GetParameters())
	var resultType = v.createResult(method.GetOptionalResult())
	implementation = classSynthesizerReference().instanceMethod_
	if uti.IsDefined(resultType) {
		implementation = classSynthesizerReference().instanceFunction_
		implementation = uti.ReplaceAll(implementation, "resultType", resultType)
	}
	implementation = uti.ReplaceAll(implementation, "methodName", methodName)
	implementation = uti.ReplaceAll(implementation, "parameters", parameters)
	return implementation
}

func (v *classSynthesizer_) createPrimaryMethods(
	primaries abs.Sequential[mod.PrimaryMethodLike],
) (
	implementation string,
) {
	var methods string
	var primaryMethods = primaries.GetIterator()
	for primaryMethods.HasNext() {
		var primaryMethod = primaryMethods.GetNext()
		if primaryMethod.GetMethod().GetName() == "GetClass" ||
			primaryMethod.GetMethod().GetName() == "GetIntrinsic" {
			continue
		}
		methods += v.createPrimaryMethod(primaryMethod)
	}
	implementation = classSynthesizerReference().primaryMethods_
	implementation = uti.ReplaceAll(
		implementation,
		"primaryMethods",
		methods,
	)
	var intrinsicMethod = v.createIntrinsicMethod()
	implementation = uti.ReplaceAll(
		implementation,
		"intrinsicMethod",
		intrinsicMethod,
	)
	return implementation
}

func (v *classSynthesizer_) createResult(
	result mod.ResultLike,
) (
	implementation string,
) {
	if uti.IsDefined(result) {
		switch actual := result.GetAny().(type) {
		case mod.AbstractionLike:
			implementation = v.extractType(actual)
		case mod.MultivalueLike:
			implementation = "(" + v.createParameters(actual.GetParameters()) + "\n)"
		}
	}
	return implementation
}

func (v *classSynthesizer_) createSetterMethod(
	setterMethod mod.SetterMethodLike,
) (
	implementation string,
) {
	var methodName = setterMethod.GetName()
	var attributeName = v.extractAttributeName(methodName)
	var parameter = setterMethod.GetParameter()
	var attributeType = v.extractType(parameter.GetAbstraction())
	var attributeCheck = v.createAttributeCheck(parameter)
	implementation = classSynthesizerReference().setterMethod_
	implementation = uti.ReplaceAll(implementation, "methodName", methodName)
	implementation = uti.ReplaceAll(implementation, "attributeName", attributeName)
	implementation = uti.ReplaceAll(implementation, "attributeType", attributeType)
	implementation = uti.ReplaceAll(implementation, "attributeCheck", attributeCheck)
	return implementation
}

func (v *classSynthesizer_) performGlobalUpdates(
	source string,
) string {
	// Set the module imports.
	var imports = v.analyzer_.GetModuleImports()
	var moduleImports = v.createModuleImports(imports, source)
	source = uti.ReplaceAll(
		source,
		"moduleImports",
		moduleImports,
	)

	// Set the instance method targets to "by value" if necessary.
	var star = "*"
	if v.analyzer_.IsIntrinsic() {
		star = ""
	}
	source = uti.ReplaceAll(
		source,
		"*",
		star,
	)

	// Insert any generics.
	var constraints = v.analyzer_.GetTypeConstraints()
	var arguments = v.analyzer_.GetTypeArguments()
	source = uti.ReplaceAll(
		source,
		"constraints",
		constraints,
	)
	source = uti.ReplaceAll(
		source,
		"arguments",
		arguments,
	)

	return source
}

func (v *classSynthesizer_) replaceAbstractionType(
	abstraction mod.AbstractionLike,
	mappings abs.CatalogLike[string, mod.AbstractionLike],
) mod.AbstractionLike {
	// Replace the generic type in a prefix with the concrete type.
	var prefix = abstraction.GetOptionalPrefix()
	if uti.IsDefined(prefix) {
		prefix = v.replacePrefixType(prefix, mappings)
	}

	// Replace the generic types in a sequence of arguments with concrete types.
	var arguments = abstraction.GetOptionalArguments()
	if uti.IsDefined(arguments) {
		arguments = v.replaceArgumentTypes(arguments, mappings)
	}

	// Replace a non-suffixed generic type with its concrete type.
	var typeName = abstraction.GetName()
	var suffix = abstraction.GetOptionalSuffix()
	if uti.IsUndefined(suffix) {
		var concreteType = mappings.GetValue(typeName)
		if uti.IsDefined(concreteType) {
			suffix = concreteType.GetOptionalSuffix()
			typeName = concreteType.GetName()
			arguments = concreteType.GetOptionalArguments()
		}
	}

	// Recreate the abstraction using its updated types.
	abstraction = mod.Abstraction().Make(
		prefix,
		typeName,
		suffix,
		arguments,
	)

	return abstraction
}

func (v *classSynthesizer_) replaceArgumentType(
	argument mod.ArgumentLike,
	mappings abs.CatalogLike[string, mod.AbstractionLike],
) mod.ArgumentLike {
	var abstraction = argument.GetAbstraction()
	abstraction = v.replaceAbstractionType(abstraction, mappings)
	argument = mod.Argument().Make(abstraction)
	return argument
}

func (v *classSynthesizer_) replaceArgumentTypes(
	arguments mod.ArgumentsLike,
	mappings abs.CatalogLike[string, mod.AbstractionLike],
) mod.ArgumentsLike {
	// Replace the generic type of the first argument with its concrete type.
	var argument = arguments.GetArgument()
	argument = v.replaceArgumentType(argument, mappings)

	// Replace the generic types of any additional arguments with concrete types.
	var additionalArguments = col.List[mod.AdditionalArgumentLike]()
	var iterator = arguments.GetAdditionalArguments().GetIterator()
	for iterator.HasNext() {
		var additionalArgument = iterator.GetNext()
		var argument = additionalArgument.GetArgument()
		argument = v.replaceArgumentType(argument, mappings)
		additionalArgument = mod.AdditionalArgument().Make(argument)
		additionalArguments.AppendValue(additionalArgument)
	}

	// Construct the updated sequence of arguments.
	arguments = mod.Arguments().Make(argument, additionalArguments)
	return arguments
}

func (v *classSynthesizer_) replaceMultivalueTypes(
	parameterized mod.MultivalueLike,
	mappings abs.CatalogLike[string, mod.AbstractionLike],
) mod.MultivalueLike {
	var parameters = parameterized.GetParameters()
	var replacedParameters = v.replaceParameterTypes(parameters, mappings)
	parameterized = mod.Multivalue().Make(replacedParameters)
	return parameterized
}

func (v *classSynthesizer_) replaceParameterType(
	parameter mod.ParameterLike,
	mappings abs.CatalogLike[string, mod.AbstractionLike],
) mod.ParameterLike {
	var parameterName = parameter.GetName()
	var abstraction = parameter.GetAbstraction()
	abstraction = v.replaceAbstractionType(abstraction, mappings)
	parameter = mod.Parameter().Make(parameterName, abstraction)
	return parameter
}

func (v *classSynthesizer_) replaceParameterTypes(
	parameters abs.Sequential[mod.ParameterLike],
	mappings abs.CatalogLike[string, mod.AbstractionLike],
) abs.Sequential[mod.ParameterLike] {
	var replacedParameters = col.List[mod.ParameterLike]()
	var iterator = parameters.GetIterator()
	for iterator.HasNext() {
		var parameter = iterator.GetNext()
		parameter = v.replaceParameterType(parameter, mappings)
		replacedParameters.AppendValue(parameter)
	}
	return replacedParameters
}

func (v *classSynthesizer_) replacePrefixType(
	prefix mod.PrefixLike,
	mappings abs.CatalogLike[string, mod.AbstractionLike],
) mod.PrefixLike {
	switch actual := prefix.GetAny().(type) {
	case mod.MapLike:
		// eg. map[K]V -> map[string]int
		var typeName = actual.GetName()
		var concreteType = mappings.GetValue(typeName)
		typeName = concreteType.GetName()
		var map_ = mod.Map().Make(typeName)
		prefix = mod.Prefix().Make(map_)
	default:
		// Ignore the rest since they don't contain any generic types.
	}
	return prefix
}

func (v *classSynthesizer_) replaceResultType(
	result mod.ResultLike,
	mappings abs.CatalogLike[string, mod.AbstractionLike],
) mod.ResultLike {
	if uti.IsUndefined(result) {
		return result
	}
	switch actual := result.GetAny().(type) {
	case mod.NoneLike:
	case mod.AbstractionLike:
		var abstraction = actual
		abstraction = v.replaceAbstractionType(abstraction, mappings)
		result = mod.Result().Make(abstraction)
	case mod.MultivalueLike:
		var parameterized = actual
		parameterized = v.replaceMultivalueTypes(parameterized, mappings)
		result = mod.Result().Make(parameterized)
	default:
		var message = fmt.Sprintf(
			"An unknown result type was found: %T",
			actual,
		)
		panic(message)
	}
	return result
}

// Instance Structure

type classSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ AnalyzerLike
}

// Class Structure

type classSynthesizerClass_ struct {
	// Declare the class constants.
	moduleImports_           string
	packageAlias_            string
	accessFunction_          string
	constructorMethods_      string
	constructorMethod_       string
	instanceInstantiation_   string
	intrinsicInstantiation_  string
	structureInstantiation_  string
	constantMethods_         string
	constantMethod_          string
	functionMethods_         string
	functionMethod_          string
	attributeMethods_        string
	getterMethod_            string
	setterMethod_            string
	aspectInterface_         string
	primaryMethods_          string
	intrinsicMethod_         string
	instanceMethod_          string
	instanceFunction_        string
	methodParameter_         string
	aspectMethods_           string
	privateMethods_          string
	instanceIntrinsic_       string
	instanceStructure_       string
	attributeDeclaration_    string
	attributeCheck_          string
	attributeInitialization_ string
	classStructure_          string
	constantDeclaration_     string
	constantInitialization_  string
	classReference_          string
	classMap_                string
}

// Class Reference

func classSynthesizerReference() *classSynthesizerClass_ {
	return classSynthesizerReference_
}

var classSynthesizerReference_ = &classSynthesizerClass_{
	// Initialize the class constants.
	moduleImports_: `

import (<ImportedPackages>)`,

	packageAlias_: `
	<~packageName> <packagePath>`,

	accessFunction_: `
// Access Function

func <~ClassName><Constraints>() <~ClassName>ClassLike<Arguments> {
	return <~className>Reference<Arguments>()
}`,

	constructorMethods_: `
// Constructor Methods<ConstructorMethods>`,

	constructorMethod_: `

func (c *<~className>Class_<Arguments>) <MethodName>(<Parameters>) <~ClassName>Like<Arguments> {<InstanceInstantiation>
}`,

	instanceInstantiation_: `
	var instance <~ClassName>Like<Arguments>
	// TBD - Add the constructor implementation.
	return instance
`,

	intrinsicInstantiation_: `
	var instance = <~className>_<Arguments>(intrinsic)
	return instance
`,

	structureInstantiation_: `<AttributeChecks>
	var instance = &<~className>_<Arguments>{
		// Initialize the instance attributes.<AttributeInitializations>
	}
	return instance
`,

	constantMethods_: `

// Constant Methods<ConstantMethods>`,

	constantMethod_: `

func (c *<~className>Class_<Arguments>) <~MethodName>() <ResultType> {
	return c.<~methodName>_
}`,

	functionMethods_: `

// Function Methods<FunctionMethods>`,

	functionMethod_: `

func (c *<~className>Class_<Arguments>) <~MethodName>(<Parameters>) <ResultType> {
	var result_ <ResultType>
	// TBD - Add the function implementation.
	return result_
}`,

	attributeMethods_: `

// Attribute Methods<AttributeMethods>`,

	getterMethod_: `

func (v <*><~className>_<Arguments>) <~MethodName>() <AttributeType> {
	return v.<~attributeName>_
}`,

	setterMethod_: `

func (v <*><~className>_<Arguments>) <~MethodName>(
	<attributeName_> <AttributeType>,
) {<AttributeCheck>
	v.<~attributeName>_ = <attributeName_>
}`,

	aspectInterface_: `

// <AspectType> Methods<AspectMethods>`,

	primaryMethods_: `
// Primary Methods

func (v <*><~className>_<Arguments>) GetClass() <~ClassName>ClassLike<Arguments> {
	return <~className>Reference<Arguments>()
}<IntrinsicMethod><PrimaryMethods>`,

	intrinsicMethod_: `

func (v <*><~className>_<Arguments>) GetIntrinsic() <IntrinsicType> {
	return <IntrinsicType>(v)
}`,

	instanceMethod_: `

func (v <*><~className>_<Arguments>) <~MethodName>(<Parameters>) {
	// TBD - Add the method implementation.
}`,

	instanceFunction_: `

func (v <*><~className>_<Arguments>) <~MethodName>(<Parameters>) <ResultType> {
	var result_ <ResultType>
	// TBD - Add the method implementation.
	return result_
}`,

	methodParameter_: `
	<parameterName_> <ParameterType>,`,

	aspectMethods_: `

// Aspect Methods

`,

	privateMethods_: `
// Private Methods

`,

	instanceIntrinsic_: `

// Instance Structure

type <~className>_<Constraints> <IntrinsicType>
`,

	instanceStructure_: `

// Instance Structure

type <~className>_<Constraints> struct {
	// Declare the instance attributes.<AttributeDeclarations>
}`,

	attributeDeclaration_: `
	<~attributeName>_ <AttributeType>`,

	attributeCheck_: `
	if uti.IsUndefined(<attributeName_>) {
		panic("The \"<~attributeName>\" attribute is required by this class.")
	}`,

	attributeInitialization_: `
		<~attributeName>_: <attributeName_>,`,

	classStructure_: `

// Class Structure

type <~className>Class_<Constraints> struct {
	// Declare the class constants.<ConstantDeclarations>
}`,

	constantDeclaration_: `
	<~constantName>_ <ConstantType>`,

	constantInitialization_: `
	// <~constantName>_: constantValue,`,

	classReference_: `

// Class Reference

func <~className>Reference() *<~className>Class_ {
	return <~className>Reference_
}

var <~className>Reference_ = &<~className>Class_{
	// Initialize the class constants.<ConstantInitializations>
}`,

	classMap_: `

// Class Reference

var <~className>Map_ = map[string]any{}
var <~className>Mutex_ syn.Mutex

func <~className>Reference<Constraints>() *<~className>Class_<Arguments> {
	// Generate the name of the bound class type.
	var class *<className>Class_<Arguments>
	var name = fmt.Sprintf("%T", class)

	// Check for an existing bound class type.
	<className>Mutex_.Lock()
	var value = <className>Map_[name]
	switch actual := value.(type) {
	case *<className>Class_<Arguments>:
		// This bound class type already exists.
		class = actual
	default:
		// Add a new bound class type.
		class = &<className>Class_<Arguments>{
			// Initialize the class constants.<ConstantInitializations>
		}
		<className>Map_[name] = class
	}
	<className>Mutex_.Unlock()

	// Return a reference to the bound class type.
	return class
}`,
}
