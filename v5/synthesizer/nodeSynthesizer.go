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
	fmt "fmt"
	mod "github.com/craterdog/go-class-model/v5/ast"
	ana "github.com/craterdog/go-code-generation/v5/analyzer"
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func NodeSynthesizer() NodeSynthesizerClassLike {
	return nodeSynthesizerReference()
}

// Constructor Methods

func (c *nodeSynthesizerClass_) Make(
	model mod.ModelLike,
	className string,
) NodeSynthesizerLike {
	var instance = &nodeSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.ModelAnalyzer().Make(model, className),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *nodeSynthesizer_) GetClass() NodeSynthesizerClassLike {
	return nodeSynthesizerReference()
}

// TemplateDriven Methods

func (v *nodeSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *nodeSynthesizer_) CreateWarningMessage() string {
	var warningMessage = nodeSynthesizerReference().warningMessage_
	return warningMessage
}

func (v *nodeSynthesizer_) CreateAccessFunction() string {
	var accessFunction = nodeSynthesizerReference().accessFunction_
	return accessFunction
}

func (v *nodeSynthesizer_) CreateConstantMethods() string {
	var methods = v.analyzer_.GetConstantMethods()
	var constantMethods = v.createConstantMethods(methods)
	return constantMethods
}

func (v *nodeSynthesizer_) CreateConstructorMethods() string {
	var methods = v.analyzer_.GetConstructorMethods()
	var constructorMethods = v.createConstructorMethods(methods)
	return constructorMethods
}

func (v *nodeSynthesizer_) CreateFunctionMethods() string {
	var methods = v.analyzer_.GetFunctionMethods()
	var functionMethods = v.createFunctionMethods(methods)
	return functionMethods
}

func (v *nodeSynthesizer_) CreatePrincipalMethods() string {
	var methods = v.analyzer_.GetPrincipalMethods()
	var principalMethods = v.createPrincipalMethods(methods)
	return principalMethods
}

func (v *nodeSynthesizer_) CreateAttributeMethods() string {
	var methods = v.analyzer_.GetAttributeMethods()
	var attributeMethods = v.createAttributeMethods(methods)
	return attributeMethods
}

func (v *nodeSynthesizer_) CreateAspectMethods() string {
	var declarations = v.analyzer_.GetAspectDeclarations()
	var interfaces = v.analyzer_.GetAspectInterfaces()
	var aspectMethods = v.createAspectInterfaces(declarations, interfaces)
	return aspectMethods
}

func (v *nodeSynthesizer_) CreatePrivateMethods() string {
	var privateMethods = nodeSynthesizerReference().privateMethods_
	return privateMethods
}

func (v *nodeSynthesizer_) CreateInstanceStructure() string {
	var instanceStructure = v.createInstanceStructure()
	return instanceStructure
}

func (v *nodeSynthesizer_) CreateClassStructure() string {
	var classStructure = v.createClassStructure()
	return classStructure
}

func (v *nodeSynthesizer_) CreateClassReference() string {
	var classReference = v.createClassReference()
	return classReference
}

func (v *nodeSynthesizer_) PerformGlobalUpdates(
	source string,
) string {
	source = v.performGlobalUpdates(source)
	return source
}

// PROTECTED INTERFACE

// Private Methods

func (v *nodeSynthesizer_) extractAttributeName(
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

func (v *nodeSynthesizer_) extractConcreteMappings(
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

func (v *nodeSynthesizer_) extractType(
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

func (v *nodeSynthesizer_) createAspectInterface(
	sequence abs.Sequential[mod.AspectDeclarationLike],
	aspectType mod.AbstractionLike,
) string {
	var methods string
	if uti.IsDefined(sequence) {
		var aspectDeclarations = sequence.GetIterator()
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
	var aspectInterface = nodeSynthesizerReference().aspectInterface_
	aspectInterface = uti.ReplaceAll(
		aspectInterface,
		"aspectType",
		v.extractType(aspectType),
	)
	aspectInterface = uti.ReplaceAll(
		aspectInterface,
		"aspectMethods",
		methods,
	)
	return aspectInterface
}

func (v *nodeSynthesizer_) createAspectInterfaces(
	declarations abs.Sequential[mod.AspectDeclarationLike],
	interfaces abs.Sequential[mod.AspectInterfaceLike],
) string {
	var aspectInterfaces string
	if uti.IsDefined(interfaces) {
		var aspects = interfaces.GetIterator()
		for aspects.HasNext() {
			var aspectInterface = aspects.GetNext()
			var aspectType = aspectInterface.GetAbstraction()
			aspectInterfaces += v.createAspectInterface(declarations, aspectType)
		}
	}
	return aspectInterfaces
}

func (v *nodeSynthesizer_) createAspectMethod(
	aspectType mod.AbstractionLike,
	method mod.MethodLike,
	mappings abs.CatalogLike[string, mod.AbstractionLike],
) string {
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
	var aspectMethod = nodeSynthesizerReference().instanceMethod_
	if uti.IsDefined(resultType) {
		aspectMethod = nodeSynthesizerReference().instanceFunction_
		aspectMethod = uti.ReplaceAll(
			aspectMethod,
			"resultType",
			resultType,
		)
	}
	aspectMethod = uti.ReplaceAll(
		aspectMethod,
		"methodName",
		methodName,
	)
	aspectMethod = uti.ReplaceAll(
		aspectMethod,
		"parameters",
		parameters,
	)
	return aspectMethod
}

func (v *nodeSynthesizer_) createAspectMethods(
	aspectType mod.AbstractionLike,
	aspectDeclaration mod.AspectDeclarationLike,
	mappings abs.CatalogLike[string, mod.AbstractionLike],
) string {
	var aspectMethods string
	var methods = aspectDeclaration.GetAspectMethods().GetIterator()
	for methods.HasNext() {
		var aspectMethod = methods.GetNext()
		var method = aspectMethod.GetMethod()
		aspectMethods += v.createAspectMethod(
			aspectType,
			method,
			mappings,
		)
	}
	return aspectMethods
}

func (v *nodeSynthesizer_) createAttributeCheck(
	parameter mod.ParameterLike,
) string {
	var attributeCheck string
	var parameterName = parameter.GetName()
	var attributeName = sts.TrimSuffix(parameterName, "_")
	// Ignore optional attributes.
	if !sts.HasPrefix(attributeName, "optional") {
		attributeCheck = nodeSynthesizerReference().attributeCheck_
		attributeCheck = uti.ReplaceAll(
			attributeCheck,
			"attributeName",
			attributeName,
		)
	}
	return attributeCheck
}

func (v *nodeSynthesizer_) createAttributeChecks(
	sequence abs.Sequential[mod.ParameterLike],
) string {
	var attributeChecks string
	var parameters = sequence.GetIterator()
	for parameters.HasNext() {
		var parameter = parameters.GetNext()
		var attributeCheck = v.createAttributeCheck(parameter)
		attributeChecks += attributeCheck
	}
	return attributeChecks
}

func (v *nodeSynthesizer_) createAttributeDeclarations() string {
	var declarations string
	var attributes = v.analyzer_.GetAttributes().GetIterator()
	for attributes.HasNext() {
		var attribute = attributes.GetNext()
		var attributeName = attribute.GetKey()
		var attributeType = attribute.GetValue()
		var declaration = nodeSynthesizerReference().attributeDeclaration_
		declaration = uti.ReplaceAll(
			declaration,
			"attributeName",
			attributeName,
		)
		declaration = uti.ReplaceAll(
			declaration,
			"attributeType",
			attributeType,
		)
		declarations += declaration
	}
	return declarations
}

func (v *nodeSynthesizer_) createAttributeInitializations(
	sequence abs.Sequential[mod.ParameterLike],
) string {
	var initializations string
	var parameters = sequence.GetIterator()
	for parameters.HasNext() {
		var parameter = parameters.GetNext()
		var parameterName = parameter.GetName()
		var attributeName = sts.TrimSuffix(parameterName, "_")
		if uti.IsDefined(v.analyzer_.GetAttributes().GetValue(attributeName)) {
			var initialization = nodeSynthesizerReference().attributeInitialization_
			initialization = uti.ReplaceAll(
				initialization,
				"attributeName",
				attributeName,
			)
			initializations += initialization
		}
	}
	return initializations
}

func (v *nodeSynthesizer_) createAttributeMethods(
	sequence abs.Sequential[mod.AttributeMethodLike],
) string {
	var methods string
	if uti.IsDefined(sequence) {
		var attributeMethods string
		var attributes = sequence.GetIterator()
		for attributes.HasNext() {
			var attributeMethod string
			var attribute = attributes.GetNext()
			switch actual := attribute.GetAny().(type) {
			case mod.GetterMethodLike:
				attributeMethod = v.createGetterMethod(actual)
			case mod.SetterMethodLike:
				attributeMethod = v.createSetterMethod(actual)
			}
			attributeMethods += attributeMethod
		}
		methods = nodeSynthesizerReference().attributeMethods_
		methods = uti.ReplaceAll(
			methods,
			"attributeMethods",
			attributeMethods,
		)
	}
	return methods
}

func (v *nodeSynthesizer_) createClassReference() string {
	var classReference = nodeSynthesizerReference().classReference_
	if v.analyzer_.IsGeneric() {
		classReference = nodeSynthesizerReference().classMap_
	}
	var constantInitializations = v.createConstantInitializations()
	classReference = uti.ReplaceAll(
		classReference,
		"constantInitializations",
		constantInitializations,
	)
	return classReference
}

func (v *nodeSynthesizer_) createClassStructure() string {
	var classStructure = nodeSynthesizerReference().classStructure_
	var constantDeclarations = v.createConstantDeclarations()
	classStructure = uti.ReplaceAll(
		classStructure,
		"constantDeclarations",
		constantDeclarations,
	)
	return classStructure
}

func (v *nodeSynthesizer_) createConstantDeclarations() string {
	var declarations string
	var constants = v.analyzer_.GetConstants().GetIterator()
	for constants.HasNext() {
		var constant = constants.GetNext()
		var constantName = constant.GetKey()
		var constantType = constant.GetValue()
		var declaration = nodeSynthesizerReference().constantDeclaration_
		declaration = uti.ReplaceAll(
			declaration,
			"constantName",
			constantName,
		)
		declaration = uti.ReplaceAll(
			declaration,
			"constantType",
			constantType,
		)
		declarations += declaration
	}
	return declarations
}

func (v *nodeSynthesizer_) createConstantInitializations() string {
	var initializations string
	var constants = v.analyzer_.GetConstants().GetIterator()
	for constants.HasNext() {
		var constant = constants.GetNext()
		var constantName = constant.GetKey()
		var initialization = nodeSynthesizerReference().constantInitialization_
		initialization = uti.ReplaceAll(
			initialization,
			"constantName",
			constantName,
		)
		initializations += initialization
	}
	return initializations
}

func (v *nodeSynthesizer_) createConstantMethod(
	constantMethod mod.ConstantMethodLike,
) string {
	var methodName = constantMethod.GetName()
	var resultType = v.extractType(constantMethod.GetAbstraction())
	var method = nodeSynthesizerReference().constantMethod_
	method = uti.ReplaceAll(
		method,
		"methodName",
		methodName,
	)
	method = uti.ReplaceAll(
		method,
		"resultType",
		resultType,
	)
	return method
}

func (v *nodeSynthesizer_) createConstantMethods(
	sequence abs.Sequential[mod.ConstantMethodLike],
) string {
	var methods string
	if uti.IsDefined(sequence) {
		var constantMethods string
		var constants = sequence.GetIterator()
		for constants.HasNext() {
			var constantMethod = constants.GetNext()
			constantMethods += v.createConstantMethod(constantMethod)
		}
		methods = nodeSynthesizerReference().constantMethods_
		methods = uti.ReplaceAll(
			methods,
			"constantMethods",
			constantMethods,
		)
	}
	return methods
}

func (v *nodeSynthesizer_) createConstructorMethod(
	constructorMethod mod.ConstructorMethodLike,
) string {
	var methodName = constructorMethod.GetName()
	var constructorParameters = constructorMethod.GetParameters()
	var parameters = v.createParameters(constructorParameters)
	var resultType = v.extractType(constructorMethod.GetAbstraction())
	var instanceInstantiation = v.createInstanceInstantiation(constructorMethod)
	var method = nodeSynthesizerReference().constructorMethod_
	method = uti.ReplaceAll(
		method,
		"methodName",
		methodName,
	)
	method = uti.ReplaceAll(
		method,
		"parameters",
		parameters,
	)
	method = uti.ReplaceAll(
		method,
		"resultType",
		resultType,
	)
	method = uti.ReplaceAll(
		method,
		"instanceInstantiation",
		instanceInstantiation,
	)
	return method
}

func (v *nodeSynthesizer_) createConstructorMethods(
	sequence abs.Sequential[mod.ConstructorMethodLike],
) string {
	var methods string
	if uti.IsDefined(sequence) {
		var constructorMethods string
		var constructors = sequence.GetIterator()
		for constructors.HasNext() {
			var constructorMethod = constructors.GetNext()
			constructorMethods += v.createConstructorMethod(constructorMethod)
		}
		methods = nodeSynthesizerReference().constructorMethods_
		methods = uti.ReplaceAll(
			methods,
			"constructorMethods",
			constructorMethods,
		)
	}
	return methods
}

func (v *nodeSynthesizer_) createFunctionMethod(
	functionMethod mod.FunctionMethodLike,
) string {
	var methodName = functionMethod.GetName()
	var parameters = v.createParameters(functionMethod.GetParameters())
	var resultType = v.createResult(functionMethod.GetResult())
	var method = nodeSynthesizerReference().functionMethod_
	method = uti.ReplaceAll(
		method,
		"methodName",
		methodName,
	)
	method = uti.ReplaceAll(
		method,
		"parameters",
		parameters,
	)
	method = uti.ReplaceAll(
		method,
		"resultType",
		resultType,
	)
	return method
}

func (v *nodeSynthesizer_) createFunctionMethods(
	sequence abs.Sequential[mod.FunctionMethodLike],
) string {
	var methods string
	if uti.IsDefined(sequence) {
		var functionMethods string
		var functions = sequence.GetIterator()
		for functions.HasNext() {
			var functionMethod = functions.GetNext()
			functionMethods += v.createFunctionMethod(functionMethod)
		}
		methods = nodeSynthesizerReference().functionMethods_
		methods = uti.ReplaceAll(
			methods,
			"functionMethods",
			functionMethods,
		)
	}
	return methods
}

func (v *nodeSynthesizer_) createGetterMethod(
	getterMethod mod.GetterMethodLike,
) string {
	var methodName = getterMethod.GetName()
	var attributeName = v.extractAttributeName(methodName)
	var attributeType = v.extractType(getterMethod.GetAbstraction())
	var method = nodeSynthesizerReference().getterMethod_
	method = uti.ReplaceAll(
		method,
		"methodName",
		methodName,
	)
	method = uti.ReplaceAll(
		method,
		"attributeName",
		attributeName,
	)
	method = uti.ReplaceAll(
		method,
		"attributeType",
		attributeType,
	)
	return method
}

func (v *nodeSynthesizer_) createInstanceInstantiation(
	constructorMethod mod.ConstructorMethodLike,
) string {
	var methodName = constructorMethod.GetName()
	var instantiation = nodeSynthesizerReference().instanceInstantiation_
	if v.analyzer_.IsIntrinsic() {
		if methodName == "Make" {
			instantiation = nodeSynthesizerReference().intrinsicInstantiation_
		}
	} else {
		if methodName == "Make" || sts.HasPrefix(methodName, "MakeWith") {
			instantiation = nodeSynthesizerReference().structureInstantiation_
			var constructorParameters = constructorMethod.GetParameters()
			var attributeChecks = v.createAttributeChecks(constructorParameters)
			var attributeInitializations = v.createAttributeInitializations(
				constructorParameters,
			)
			instantiation = uti.ReplaceAll(
				instantiation,
				"attributeChecks",
				attributeChecks,
			)
			instantiation = uti.ReplaceAll(
				instantiation,
				"attributeInitializations",
				attributeInitializations,
			)
		}
	}
	return instantiation
}

func (v *nodeSynthesizer_) createInstanceStructure() string {
	var structure string
	if v.analyzer_.IsIntrinsic() {
		var intrinsicType = v.extractType(v.analyzer_.GetIntrinsicType())
		structure = nodeSynthesizerReference().instanceIntrinsic_
		structure = uti.ReplaceAll(
			structure,
			"intrinsicType",
			intrinsicType,
		)
	} else {
		structure = nodeSynthesizerReference().instanceStructure_
		var attributeDeclarations = v.createAttributeDeclarations()
		structure = uti.ReplaceAll(
			structure,
			"attributeDeclarations",
			attributeDeclarations,
		)
	}
	return structure
}

func (v *nodeSynthesizer_) createIntrinsicMethod() string {
	var method string
	if v.analyzer_.IsIntrinsic() {
		var intrinsicType = v.extractType(v.analyzer_.GetIntrinsicType())
		method = nodeSynthesizerReference().intrinsicMethod_
		method = uti.ReplaceAll(
			method,
			"intrinsicType",
			intrinsicType,
		)
	}
	return method
}

func (v *nodeSynthesizer_) createImportedPackages(
	source string,
) string {
	var importedPackages string
	var packages = v.analyzer_.GetImportedPackages().GetIterator()
	for packages.HasNext() {
		var association = packages.GetNext()
		var packagePath = association.GetKey()
		var packageAcronym = association.GetValue()
		var prefix = packageAcronym + "."
		if sts.Contains(source, prefix) {
			var packageAlias = nodeSynthesizerReference().packageAlias_
			packageAlias = uti.ReplaceAll(
				packageAlias,
				"packageAcronym",
				packageAcronym,
			)
			packageAlias = uti.ReplaceAll(
				packageAlias,
				"packagePath",
				packagePath,
			)
			importedPackages += packageAlias
		}
	}
	if uti.IsDefined(importedPackages) {
		importedPackages += "\n"
	}
	return importedPackages
}

func (v *nodeSynthesizer_) createParameters(
	sequence abs.Sequential[mod.ParameterLike],
) string {
	var methodParameters string
	var parameters = sequence.GetIterator()
	for parameters.HasNext() {
		var parameter = parameters.GetNext()
		var parameterName = parameter.GetName()
		var parameterType = v.extractType(parameter.GetAbstraction())
		var methodParameter = nodeSynthesizerReference().methodParameter_
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

func (v *nodeSynthesizer_) createPrincipalMethod(
	method mod.MethodLike,
) string {
	var methodName = method.GetName()
	var parameters = v.createParameters(method.GetParameters())
	var resultType = v.createResult(method.GetOptionalResult())
	var principalMethod = nodeSynthesizerReference().instanceMethod_
	if uti.IsDefined(resultType) {
		principalMethod = nodeSynthesizerReference().instanceFunction_
		principalMethod = uti.ReplaceAll(
			principalMethod,
			"resultType",
			resultType,
		)
	}
	principalMethod = uti.ReplaceAll(
		principalMethod,
		"methodName",
		methodName,
	)
	principalMethod = uti.ReplaceAll(
		principalMethod,
		"parameters",
		parameters,
	)
	return principalMethod
}

func (v *nodeSynthesizer_) createPrincipalMethods(
	sequence abs.Sequential[mod.PrincipalMethodLike],
) string {
	var principalMethods string
	var methods = sequence.GetIterator()
	for methods.HasNext() {
		var method = methods.GetNext().GetMethod()
		if method.GetName() == "GetClass" ||
			method.GetName() == "GetIntrinsic" {
			continue
		}
		var principalMethod = v.createPrincipalMethod(method)
		principalMethods += principalMethod
	}
	var implementation = nodeSynthesizerReference().principalMethods_
	implementation = uti.ReplaceAll(
		implementation,
		"principalMethods",
		principalMethods,
	)
	var intrinsicMethod = v.createIntrinsicMethod()
	implementation = uti.ReplaceAll(
		implementation,
		"intrinsicMethod",
		intrinsicMethod,
	)
	return implementation
}

func (v *nodeSynthesizer_) createResult(
	result mod.ResultLike,
) string {
	var results string
	if uti.IsDefined(result) {
		switch actual := result.GetAny().(type) {
		case mod.AbstractionLike:
			results = v.extractType(actual)
		case mod.MultivalueLike:
			results = "(" + v.createParameters(actual.GetParameters()) + "\n)"
		}
	}
	return results
}

func (v *nodeSynthesizer_) createSetterMethod(
	setterMethod mod.SetterMethodLike,
) string {
	var methodName = setterMethod.GetName()
	var attributeName = v.extractAttributeName(methodName)
	var parameter = setterMethod.GetParameter()
	var attributeType = v.extractType(parameter.GetAbstraction())
	var attributeCheck = v.createAttributeCheck(parameter)
	var method = nodeSynthesizerReference().setterMethod_
	method = uti.ReplaceAll(
		method,
		"methodName",
		methodName,
	)
	method = uti.ReplaceAll(
		method,
		"attributeName",
		attributeName,
	)
	method = uti.ReplaceAll(
		method,
		"attributeType",
		attributeType,
	)
	method = uti.ReplaceAll(
		method,
		"attributeCheck",
		attributeCheck,
	)
	return method
}

func (v *nodeSynthesizer_) performGlobalUpdates(
	source string,
) string {
	// Update the class imports.
	var importedPackages = v.createImportedPackages(source)
	source = uti.ReplaceAll(
		source,
		"importedPackages",
		importedPackages,
	)

	// Update the instance method targets to "by value" if necessary.
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

func (v *nodeSynthesizer_) replaceAbstractionType(
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

func (v *nodeSynthesizer_) replaceArgumentType(
	argument mod.ArgumentLike,
	mappings abs.CatalogLike[string, mod.AbstractionLike],
) mod.ArgumentLike {
	var abstraction = argument.GetAbstraction()
	abstraction = v.replaceAbstractionType(abstraction, mappings)
	argument = mod.Argument().Make(abstraction)
	return argument
}

func (v *nodeSynthesizer_) replaceArgumentTypes(
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

func (v *nodeSynthesizer_) replaceMultivalueTypes(
	parameterized mod.MultivalueLike,
	mappings abs.CatalogLike[string, mod.AbstractionLike],
) mod.MultivalueLike {
	var parameters = parameterized.GetParameters()
	var replacedParameters = v.replaceParameterTypes(parameters, mappings)
	parameterized = mod.Multivalue().Make(replacedParameters)
	return parameterized
}

func (v *nodeSynthesizer_) replaceParameterType(
	parameter mod.ParameterLike,
	mappings abs.CatalogLike[string, mod.AbstractionLike],
) mod.ParameterLike {
	var parameterName = parameter.GetName()
	var abstraction = parameter.GetAbstraction()
	abstraction = v.replaceAbstractionType(abstraction, mappings)
	parameter = mod.Parameter().Make(parameterName, abstraction)
	return parameter
}

func (v *nodeSynthesizer_) replaceParameterTypes(
	sequence abs.Sequential[mod.ParameterLike],
	mappings abs.CatalogLike[string, mod.AbstractionLike],
) abs.Sequential[mod.ParameterLike] {
	var replacedParameters = col.List[mod.ParameterLike]()
	var parameters = sequence.GetIterator()
	for parameters.HasNext() {
		var parameter = parameters.GetNext()
		parameter = v.replaceParameterType(parameter, mappings)
		replacedParameters.AppendValue(parameter)
	}
	return replacedParameters
}

func (v *nodeSynthesizer_) replacePrefixType(
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

func (v *nodeSynthesizer_) replaceResultType(
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

type nodeSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.ModelAnalyzerLike
}

// Class Structure

type nodeSynthesizerClass_ struct {
	// Declare the class constants.
	warningMessage_          string
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
	principalMethods_        string
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

func nodeSynthesizerReference() *nodeSynthesizerClass_ {
	return nodeSynthesizerReference_
}

var nodeSynthesizerReference_ = &nodeSynthesizerClass_{
	// Initialize the class constants.
	warningMessage_: `
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
`,

	packageAlias_: `
	<~packageAcronym> "<packagePath>"`,

	accessFunction_: `
// Access Function

func <~ClassName><Constraints>() <~ClassName>ClassLike<Arguments> {
	return <~className>Reference<Arguments>()
}
`,

	constructorMethods_: `
// Constructor Methods
<ConstructorMethods>`,

	constructorMethod_: `
func (c *<~className>Class_<Arguments>) <MethodName>(<Parameters>) <~ClassName>Like<Arguments> {<InstanceInstantiation>}
`,

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
// Constant Methods
<ConstantMethods>`,

	constantMethod_: `
func (c *<~className>Class_<Arguments>) <~MethodName>() <ResultType> {
	return c.<~methodName>_
}
`,

	functionMethods_: `
// Function Methods
<FunctionMethods>`,

	functionMethod_: `
func (c *<~className>Class_<Arguments>) <~MethodName>(<Parameters>) <ResultType> {
	var result_ <ResultType>
	// TBD - Add the function implementation.
	return result_
}
`,

	attributeMethods_: `
// Attribute Methods
<AttributeMethods>`,

	getterMethod_: `
func (v <*><~className>_<Arguments>) <~MethodName>() <AttributeType> {
	return v.<~attributeName>_
}
`,

	setterMethod_: `
func (v <*><~className>_<Arguments>) <~MethodName>(
	<attributeName_> <AttributeType>,
) {<AttributeCheck>
	v.<~attributeName>_ = <attributeName_>
}
`,

	aspectInterface_: `
// <AspectType> Methods
<AspectMethods>`,

	principalMethods_: `
// Principal Methods

func (v <*><~className>_<Arguments>) GetClass() <~ClassName>ClassLike<Arguments> {
	return <~className>Reference<Arguments>()
}
<IntrinsicMethod><PrincipalMethods>`,

	intrinsicMethod_: `
func (v <*><~className>_<Arguments>) GetIntrinsic() <IntrinsicType> {
	return <IntrinsicType>(v)
}
`,

	instanceMethod_: `
func (v <*><~className>_<Arguments>) <~MethodName>(<Parameters>) {
	// TBD - Add the method implementation.
}
`,

	instanceFunction_: `
func (v <*><~className>_<Arguments>) <~MethodName>(<Parameters>) <ResultType> {
	var result_ <ResultType>
	// TBD - Add the method implementation.
	return result_
}
`,

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
}
`,

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
}
`,

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
}
`,

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
}
`,
}
