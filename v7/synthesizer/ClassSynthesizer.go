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
	fmt "fmt"
	mod "github.com/craterdog/go-class-model/v7/ast"
	ana "github.com/craterdog/go-code-generation/v7/analyzer"
	col "github.com/craterdog/go-collection-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ClassSynthesizerClass() ClassSynthesizerClassLike {
	return classSynthesizerClass()
}

// Constructor Methods

func (c *classSynthesizerClass_) ClassSynthesizer(
	model mod.ModelLike,
	className string,
) ClassSynthesizerLike {
	var instance = &classSynthesizer_{
		// Initialize the instance attributes.
		packageAnalyzer_: ana.PackageAnalyzerClass().PackageAnalyzer(model),
		classAnalyzer_:   ana.ClassAnalyzerClass().ClassAnalyzer(model, className),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *classSynthesizer_) GetClass() ClassSynthesizerClassLike {
	return classSynthesizerClass()
}

// TemplateDriven Methods

func (v *classSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.packageAnalyzer_.GetLegalNotice()
	return legalNotice
}

func (v *classSynthesizer_) CreateWarningMessage() string {
	var class = classSynthesizerClass()
	return class.warningMessage_
}

func (v *classSynthesizer_) CreateImportedPackages() string {
	var class = classSynthesizerClass()
	var importedPackages = class.importedPackages_
	var packages = v.packageAnalyzer_.GetImportedPackages().GetIterator()
	for packages.HasNext() {
		var association = packages.GetNext()
		var packageAcronym = association.GetKey()
		var packagePath = association.GetValue()
		var importedPath = class.importedPath_
		importedPath = uti.ReplaceAll(
			importedPath,
			"packageAcronym",
			packageAcronym,
		)
		importedPath = uti.ReplaceAll(
			importedPath,
			"packagePath",
			packagePath,
		)
		importedPackages += importedPath
	}
	if uti.IsDefined(importedPackages) {
		importedPackages += "\n"
	}
	return importedPackages
}

func (v *classSynthesizer_) CreateAccessFunction() string {
	var class = classSynthesizerClass()
	return class.accessFunction_
}

func (v *classSynthesizer_) CreateConstantMethods() string {
	var methods = v.classAnalyzer_.GetConstantMethods()
	var constantMethods = v.createConstantMethods(methods)
	return constantMethods
}

func (v *classSynthesizer_) CreateConstructorMethods() string {
	var methods = v.classAnalyzer_.GetConstructorMethods()
	var constructorMethods = v.createConstructorMethods(methods)
	return constructorMethods
}

func (v *classSynthesizer_) CreateFunctionMethods() string {
	var methods = v.classAnalyzer_.GetFunctionMethods()
	var functionMethods = v.createFunctionMethods(methods)
	return functionMethods
}

func (v *classSynthesizer_) CreatePrincipalMethods() string {
	var methods = v.classAnalyzer_.GetPrincipalMethods()
	var principalMethods = v.createPrincipalMethods(methods)
	return principalMethods
}

func (v *classSynthesizer_) CreateAttributeMethods() string {
	var methods = v.classAnalyzer_.GetAttributeMethods()
	var attributeMethods = v.createAttributeMethods(methods)
	return attributeMethods
}

func (v *classSynthesizer_) CreateAspectInterfaces() string {
	var declarations = v.packageAnalyzer_.GetAspectDeclarations()
	var interfaces = v.classAnalyzer_.GetAspectInterfaces()
	var aspectInterfaces = v.createAspectInterfaces(declarations, interfaces)
	return aspectInterfaces
}

func (v *classSynthesizer_) CreatePrivateMethods() string {
	var class = classSynthesizerClass()
	return class.privateMethods_
}

func (v *classSynthesizer_) CreateInstanceStructure() string {
	var instanceStructure = v.createInstanceStructure()
	return instanceStructure
}

func (v *classSynthesizer_) CreateClassStructure() string {
	var classStructure = v.createClassStructure()
	return classStructure
}

func (v *classSynthesizer_) CreateClass() string {
	var classReference = v.createClass()
	return classReference
}

func (v *classSynthesizer_) PerformGlobalUpdates(
	existing string,
	generated string,
) string {
	// Update the instance method targets to "by value" if necessary.
	var star = "*"
	if v.classAnalyzer_.IsIntrinsic() {
		star = ""
	}
	generated = uti.ReplaceAll(
		generated,
		"*",
		star,
	)

	// Insert any generics.
	var constraints = v.classAnalyzer_.GetTypeConstraints()
	var arguments = v.classAnalyzer_.GetTypeArguments()
	generated = uti.ReplaceAll(
		generated,
		"constraints",
		constraints,
	)
	generated = uti.ReplaceAll(
		generated,
		"arguments",
		arguments,
	)

	return generated
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

func (v *classSynthesizer_) extractMappings(
	constraints mod.ConstraintsLike,
	arguments mod.ArgumentsLike,
) col.CatalogLike[string, mod.AbstractionLike] {
	// Create the mappings catalog.
	var mappings = col.Catalog[string, mod.AbstractionLike]()
	if uti.IsUndefined(constraints) || uti.IsUndefined(arguments) {
		return mappings
	}

	// Map the name of the first constraint to its mapped type.
	var constraint = constraints.GetConstraint()
	var constraintName = constraint.GetName()
	var argument = arguments.GetArgument()
	var mappedType = argument.GetAbstraction()
	mappings.SetValue(constraintName, mappedType)

	// Map the name of the additional constraints to their mapped types.
	var additionalConstraints = constraints.GetAdditionalConstraints().GetIterator()
	var additionalArguments = arguments.GetAdditionalArguments().GetIterator()
	for additionalConstraints.HasNext() {
		constraint = additionalConstraints.GetNext().GetConstraint()
		constraintName = constraint.GetName()
		argument = additionalArguments.GetNext().GetArgument()
		mappedType = argument.GetAbstraction()
		mappings.SetValue(constraintName, mappedType)
	}

	return mappings
}

func (v *classSynthesizer_) extractType(
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

func (v *classSynthesizer_) createAspectInterface(
	aspectDeclarations col.ListLike[mod.AspectDeclarationLike],
	aspectType mod.AbstractionLike,
) string {
	var class = classSynthesizerClass()
	var aspectInterface = class.aspectInterface_
	aspectInterface = uti.ReplaceAll(
		aspectInterface,
		"aspectType",
		v.extractType(aspectType),
	)
	var aspectMethods string
	if uti.IsDefined(aspectType.GetOptionalPrefix()) {
		// Ignore the methods for aspect types defined in other packages.
		aspectInterface = uti.ReplaceAll(
			aspectInterface,
			"aspectMethods",
			aspectMethods,
		)
		return aspectInterface
	}
	var iterator = aspectDeclarations.GetIterator()
	for iterator.HasNext() {
		var aspectDeclaration = iterator.GetNext()
		var declaration = aspectDeclaration.GetDeclaration()
		var constraints = declaration.GetOptionalConstraints()
		var arguments = aspectType.GetOptionalArguments()
		if declaration.GetName() == aspectType.GetName() {
			var mappings = v.extractMappings(constraints, arguments)
			aspectMethods = v.createAspectMethods(
				aspectDeclaration,
				mappings,
			)
			break
		}
	}
	aspectInterface = uti.ReplaceAll(
		aspectInterface,
		"aspectMethods",
		aspectMethods,
	)
	return aspectInterface
}

func (v *classSynthesizer_) createAspectInterfaces(
	declarations col.ListLike[mod.AspectDeclarationLike],
	interfaces col.ListLike[mod.AspectInterfaceLike],
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

func (v *classSynthesizer_) createAspectMethod(
	method mod.MethodLike,
	mappings col.CatalogLike[string, mod.AbstractionLike],
) string {
	// Determine any instance method parameters.
	var methodParameters col.ListLike[mod.ParameterLike]
	var parameterList = method.GetOptionalParameterList()
	if uti.IsDefined(parameterList) {
		methodParameters = parameterList.GetParameters()
	} else {
		methodParameters = col.List[mod.ParameterLike]()
	}

	// Determine the type of method template based on its result type.
	var aspectMethod string
	var methodResult = method.GetResult()
	var class = classSynthesizerClass()
	switch methodResult.GetAny().(type) {
	case mod.NoneLike:
		// The method return has no value.
		aspectMethod = class.instanceMethod_
	case mod.AbstractionLike:
		// The method returns a single value.
		aspectMethod = class.instanceFunction_
	case mod.MultivalueLike:
		// The method returns multiple values by name.
		aspectMethod = class.instanceMultiFunction_
	}

	// Replace any generic types with their corresponding mapped types.
	if mappings.GetSize() > 0 {
		methodParameters = v.replaceParameterTypes(methodParameters, mappings)
		methodResult = v.replaceResultType(methodResult, mappings)
	}

	// Plug the values determined above into the code template.
	var methodName = method.GetName()
	aspectMethod = uti.ReplaceAll(
		aspectMethod,
		"methodName",
		methodName,
	)
	var parameters = v.createParameters(methodParameters)
	aspectMethod = uti.ReplaceAll(
		aspectMethod,
		"parameters",
		parameters,
	)
	var resultType = v.createResult(methodResult)
	aspectMethod = uti.ReplaceAll(
		aspectMethod,
		"resultType",
		resultType,
	)
	return aspectMethod
}

func (v *classSynthesizer_) createAspectMethods(
	aspectDeclaration mod.AspectDeclarationLike,
	mappings col.CatalogLike[string, mod.AbstractionLike],
) string {
	var aspectMethods string
	var iterator = aspectDeclaration.GetAspectMethods().GetIterator()
	for iterator.HasNext() {
		var aspectMethod = iterator.GetNext()
		var method = aspectMethod.GetMethod()
		aspectMethods += v.createAspectMethod(
			method,
			mappings,
		)
	}
	return aspectMethods
}

func (v *classSynthesizer_) createAttributeCheck(
	parameter mod.ParameterLike,
) string {
	var attributeCheck string
	var parameterName = parameter.GetName()
	var attributeName = sts.TrimSuffix(parameterName, "_")
	// Ignore optional attributes.
	if !sts.HasPrefix(attributeName, "optional") {
		var class = classSynthesizerClass()
		attributeCheck = class.attributeCheck_
		attributeCheck = uti.ReplaceAll(
			attributeCheck,
			"attributeName",
			attributeName,
		)
	}
	return attributeCheck
}

func (v *classSynthesizer_) createAttributeChecks(
	list col.ListLike[mod.ParameterLike],
) string {
	var attributeChecks string
	var parameters = list.GetIterator()
	for parameters.HasNext() {
		var parameter = parameters.GetNext()
		var attributeCheck = v.createAttributeCheck(parameter)
		attributeChecks += attributeCheck
	}
	return attributeChecks
}

func (v *classSynthesizer_) createAttributeDeclarations() string {
	var declarations string
	var attributes = v.classAnalyzer_.GetAttributes().GetIterator()
	for attributes.HasNext() {
		var attribute = attributes.GetNext()
		var attributeName = attribute.GetKey()
		var attributeType = attribute.GetValue()
		var class = classSynthesizerClass()
		var declaration = class.attributeDeclaration_
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

func (v *classSynthesizer_) createAttributeInitializations(
	list col.ListLike[mod.ParameterLike],
) string {
	var initializations string
	var parameters = list.GetIterator()
	for parameters.HasNext() {
		var parameter = parameters.GetNext()
		var parameterName = parameter.GetName()
		var attributeName = sts.TrimSuffix(parameterName, "_")
		if uti.IsDefined(v.classAnalyzer_.GetAttributes().GetValue(attributeName)) {
			var class = classSynthesizerClass()
			var initialization = class.attributeInitialization_
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

func (v *classSynthesizer_) createAttributeMethods(
	list col.ListLike[mod.AttributeMethodLike],
) string {
	var methods string
	if uti.IsDefined(list) {
		var attributeMethods string
		var attributes = list.GetIterator()
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
		var class = classSynthesizerClass()
		methods = class.attributeMethods_
		methods = uti.ReplaceAll(
			methods,
			"attributeMethods",
			attributeMethods,
		)
	}
	return methods
}

func (v *classSynthesizer_) createClass() string {
	var class = classSynthesizerClass()
	var classReference = class.classReference_
	var constantInitialization = class.constantInitialization_
	if v.classAnalyzer_.IsGeneric() {
		classReference = class.classMap_
		constantInitialization = class.mapInitialization_
	}
	var constantInitializations = v.createConstantInitializations(
		constantInitialization,
	)
	classReference = uti.ReplaceAll(
		classReference,
		"constantInitializations",
		constantInitializations,
	)
	return classReference
}

func (v *classSynthesizer_) createClassStructure() string {
	var class = classSynthesizerClass()
	var classStructure = class.classStructure_
	var constantDeclarations = v.createConstantDeclarations()
	classStructure = uti.ReplaceAll(
		classStructure,
		"constantDeclarations",
		constantDeclarations,
	)
	return classStructure
}

func (v *classSynthesizer_) createConstantDeclarations() string {
	var declarations string
	var constants = v.classAnalyzer_.GetConstants().GetIterator()
	for constants.HasNext() {
		var constant = constants.GetNext()
		var constantName = constant.GetKey()
		var constantType = constant.GetValue()
		var class = classSynthesizerClass()
		var declaration = class.constantDeclaration_
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

func (v *classSynthesizer_) createConstantInitializations(
	intializationTemplate string,
) string {
	var constantInitializations string
	var constants = v.classAnalyzer_.GetConstants().GetIterator()
	for constants.HasNext() {
		var constant = constants.GetNext()
		var constantName = constant.GetKey()
		var constantInitialization = uti.ReplaceAll(
			intializationTemplate,
			"constantName",
			constantName,
		)
		constantInitializations += constantInitialization
	}
	return constantInitializations
}

func (v *classSynthesizer_) createConstantMethod(
	constantMethod mod.ConstantMethodLike,
) string {
	var methodName = constantMethod.GetName()
	var resultType = v.extractType(constantMethod.GetAbstraction())
	var class = classSynthesizerClass()
	var method = class.constantMethod_
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

func (v *classSynthesizer_) createConstantMethods(
	list col.ListLike[mod.ConstantMethodLike],
) string {
	var methods string
	if uti.IsDefined(list) {
		var constantMethods string
		var constants = list.GetIterator()
		for constants.HasNext() {
			var constantMethod = constants.GetNext()
			constantMethods += v.createConstantMethod(constantMethod)
		}
		var class = classSynthesizerClass()
		methods = class.constantMethods_
		methods = uti.ReplaceAll(
			methods,
			"constantMethods",
			constantMethods,
		)
	}
	return methods
}

func (v *classSynthesizer_) createConstructorMethod(
	constructorMethod mod.ConstructorMethodLike,
) string {
	var methodName = constructorMethod.GetName()
	var constructorParameters col.ListLike[mod.ParameterLike]
	var parameterList = constructorMethod.GetOptionalParameterList()
	if uti.IsDefined(parameterList) {
		constructorParameters = parameterList.GetParameters()
	} else {
		constructorParameters = col.List[mod.ParameterLike]()
	}
	var parameters = v.createParameters(constructorParameters)
	var resultType = v.extractType(constructorMethod.GetAbstraction())
	var instanceInstantiation = v.createInstanceInstantiation(constructorMethod)
	var class = classSynthesizerClass()
	var method = class.constructorMethod_
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

func (v *classSynthesizer_) createConstructorMethods(
	list col.ListLike[mod.ConstructorMethodLike],
) string {
	var methods string
	if uti.IsDefined(list) {
		var constructorMethods string
		var constructors = list.GetIterator()
		for constructors.HasNext() {
			var constructorMethod = constructors.GetNext()
			constructorMethods += v.createConstructorMethod(constructorMethod)
		}
		var class = classSynthesizerClass()
		methods = class.constructorMethods_
		methods = uti.ReplaceAll(
			methods,
			"constructorMethods",
			constructorMethods,
		)
	}
	return methods
}

func (v *classSynthesizer_) createFunctionMethod(
	method mod.FunctionMethodLike,
) string {
	var functionParameters col.ListLike[mod.ParameterLike]
	var parameterList = method.GetOptionalParameterList()
	if uti.IsDefined(parameterList) {
		functionParameters = parameterList.GetParameters()
	} else {
		functionParameters = col.List[mod.ParameterLike]()
	}
	var class = classSynthesizerClass()
	var functionMethod = class.functionMethod_
	var methodName = method.GetName()
	functionMethod = uti.ReplaceAll(
		functionMethod,
		"methodName",
		methodName,
	)
	var parameters = v.createParameters(functionParameters)
	functionMethod = uti.ReplaceAll(
		functionMethod,
		"parameters",
		parameters,
	)
	var resultType = v.createResult(method.GetResult())
	functionMethod = uti.ReplaceAll(
		functionMethod,
		"resultType",
		resultType,
	)
	return functionMethod
}

func (v *classSynthesizer_) createFunctionMethods(
	list col.ListLike[mod.FunctionMethodLike],
) string {
	var methods string
	if uti.IsDefined(list) {
		var functionMethods string
		var functions = list.GetIterator()
		for functions.HasNext() {
			var functionMethod = functions.GetNext()
			functionMethods += v.createFunctionMethod(functionMethod)
		}
		var class = classSynthesizerClass()
		methods = class.functionMethods_
		methods = uti.ReplaceAll(
			methods,
			"functionMethods",
			functionMethods,
		)
	}
	return methods
}

func (v *classSynthesizer_) createGetterMethod(
	getterMethod mod.GetterMethodLike,
) string {
	var methodName = getterMethod.GetName()
	var attributeName = v.extractAttributeName(methodName)
	var attributeType = v.extractType(getterMethod.GetAbstraction())
	var class = classSynthesizerClass()
	var method = class.getterMethod_
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

func (v *classSynthesizer_) createInstanceInstantiation(
	constructorMethod mod.ConstructorMethodLike,
) string {
	var methodName = constructorMethod.GetName()
	var className = constructorMethod.GetAbstraction().GetName()
	className = sts.TrimSuffix(className, "Like")
	var class = classSynthesizerClass()
	var instantiation = class.instanceInstantiation_
	if v.classAnalyzer_.IsIntrinsic() {
		if methodName == className {
			instantiation = class.intrinsicInstantiation_
		}
	} else {
		if methodName == className || sts.HasPrefix(methodName, className+"With") {
			instantiation = class.structureInstantiation_
			var constructorParameters col.ListLike[mod.ParameterLike]
			var parameterList = constructorMethod.GetOptionalParameterList()
			if uti.IsDefined(parameterList) {
				constructorParameters = parameterList.GetParameters()
			} else {
				constructorParameters = col.List[mod.ParameterLike]()
			}
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

func (v *classSynthesizer_) createInstanceStructure() string {
	var structure string
	var class = classSynthesizerClass()
	if v.classAnalyzer_.IsIntrinsic() {
		var intrinsicType = v.extractType(v.classAnalyzer_.GetIntrinsicType())
		structure = class.instanceIntrinsic_
		structure = uti.ReplaceAll(
			structure,
			"intrinsicType",
			intrinsicType,
		)
	} else {
		structure = class.instanceStructure_
		var attributeDeclarations = v.createAttributeDeclarations()
		structure = uti.ReplaceAll(
			structure,
			"attributeDeclarations",
			attributeDeclarations,
		)
	}
	return structure
}

func (v *classSynthesizer_) createIntrinsicMethod() string {
	var method string
	if v.classAnalyzer_.IsIntrinsic() {
		var intrinsicType = v.extractType(v.classAnalyzer_.GetIntrinsicType())
		var class = classSynthesizerClass()
		method = class.intrinsicMethod_
		method = uti.ReplaceAll(
			method,
			"intrinsicType",
			intrinsicType,
		)
	}
	return method
}

func (v *classSynthesizer_) createParameters(
	list col.ListLike[mod.ParameterLike],
) string {
	var methodParameters string
	var parameters = list.GetIterator()
	for parameters.HasNext() {
		var parameter = parameters.GetNext()
		var parameterName = parameter.GetName()
		var parameterType = v.extractType(parameter.GetAbstraction())
		var class = classSynthesizerClass()
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

func (v *classSynthesizer_) createPrincipalMethod(
	method mod.MethodLike,
) string {
	// Determine any instance method parameters.
	var methodParameters col.ListLike[mod.ParameterLike]
	var parameterList = method.GetOptionalParameterList()
	if uti.IsDefined(parameterList) {
		methodParameters = parameterList.GetParameters()
	} else {
		methodParameters = col.List[mod.ParameterLike]()
	}

	// Determine the type of method template based on its result type.
	var principalMethod string
	var methodResult = method.GetResult()
	var class = classSynthesizerClass()
	switch methodResult.GetAny().(type) {
	case mod.NoneLike:
		// The method return has no value.
		principalMethod = class.instanceMethod_
	case mod.AbstractionLike:
		// The method returns a single value.
		principalMethod = class.instanceFunction_
	case mod.MultivalueLike:
		// The method returns multiple values by name.
		principalMethod = class.instanceMultiFunction_
	}

	// Plug the values determined above into the code template.
	var methodName = method.GetName()
	principalMethod = uti.ReplaceAll(
		principalMethod,
		"methodName",
		methodName,
	)
	var parameters = v.createParameters(methodParameters)
	principalMethod = uti.ReplaceAll(
		principalMethod,
		"parameters",
		parameters,
	)
	var resultType = v.createResult(methodResult)
	principalMethod = uti.ReplaceAll(
		principalMethod,
		"resultType",
		resultType,
	)
	return principalMethod
}

func (v *classSynthesizer_) createPrincipalMethods(
	list col.ListLike[mod.PrincipalMethodLike],
) string {
	var principalMethods string
	var methods = list.GetIterator()
	for methods.HasNext() {
		var method = methods.GetNext().GetMethod()
		if method.GetName() == "GetClass" ||
			method.GetName() == "GetIntrinsic" {
			continue
		}
		var principalMethod = v.createPrincipalMethod(method)
		principalMethods += principalMethod
	}
	var class = classSynthesizerClass()
	var implementation = class.principalMethods_
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

func (v *classSynthesizer_) createResult(
	result mod.ResultLike,
) string {
	var results string
	switch actual := result.GetAny().(type) {
	case mod.AbstractionLike:
		results = v.extractType(actual)
	case mod.MultivalueLike:
		var parameterList = actual.GetParameterList()
		var parameters = parameterList.GetParameters()
		results = "(" + v.createParameters(parameters) + ")"
	}
	return results
}

func (v *classSynthesizer_) createSetterMethod(
	setterMethod mod.SetterMethodLike,
) string {
	var methodName = setterMethod.GetName()
	var attributeName = v.extractAttributeName(methodName)
	var parameter = setterMethod.GetParameter()
	var attributeType = v.extractType(parameter.GetAbstraction())
	var attributeCheck = v.createAttributeCheck(parameter)
	var method = classSynthesizerClass().setterMethod_
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

func (v *classSynthesizer_) replaceAbstractionType(
	abstraction mod.AbstractionLike,
	mappings col.CatalogLike[string, mod.AbstractionLike],
) mod.AbstractionLike {
	// Check for external abstraction type.
	var prefix = abstraction.GetOptionalPrefix()
	if uti.IsDefined(prefix) {
		// The mappings cannot include mappings from external packages.
		return abstraction
	}

	// Check for wrapper mapping.
	var wrapper = abstraction.GetOptionalWrapper()
	wrapper = v.replaceWrapperType(wrapper, mappings)

	// Check for name mapping.
	var name = abstraction.GetName()
	name = v.replaceNameType(name, mappings)

	// Check for argument type mappings.
	var arguments = abstraction.GetOptionalArguments()
	arguments = v.replaceArgumentTypes(arguments, mappings)

	// Recreate the abstraction using its updated types.
	abstraction = mod.AbstractionClass().Abstraction(
		wrapper,
		prefix,
		name,
		arguments,
	)
	return abstraction
}

func (v *classSynthesizer_) replaceArgumentType(
	argument mod.ArgumentLike,
	mappings col.CatalogLike[string, mod.AbstractionLike],
) mod.ArgumentLike {
	var abstraction = argument.GetAbstraction()
	abstraction = v.replaceAbstractionType(abstraction, mappings)
	argument = mod.ArgumentClass().Argument(abstraction)
	return argument
}

func (v *classSynthesizer_) replaceArgumentTypes(
	arguments mod.ArgumentsLike,
	mappings col.CatalogLike[string, mod.AbstractionLike],
) mod.ArgumentsLike {
	// Check for arguments.
	if uti.IsUndefined(arguments) {
		return arguments
	}

	// Replace the generic type of the first argument with its mapped type.
	var argument = arguments.GetArgument()
	argument = v.replaceArgumentType(argument, mappings)

	// Replace the generic types of any additional arguments with mapped types.
	var additionalArguments = col.List[mod.AdditionalArgumentLike]()
	var iterator = arguments.GetAdditionalArguments().GetIterator()
	for iterator.HasNext() {
		var additionalArgument = iterator.GetNext()
		var argument = additionalArgument.GetArgument()
		argument = v.replaceArgumentType(argument, mappings)
		additionalArgument = mod.AdditionalArgumentClass().AdditionalArgument(
			",",
			argument,
		)
		additionalArguments.AppendValue(additionalArgument)
	}

	// Construct the updated list of arguments.
	arguments = mod.ArgumentsClass().Arguments(
		"[",
		argument,
		additionalArguments,
		"]",
	)
	return arguments
}

func (v *classSynthesizer_) replaceMultivalueTypes(
	parameterized mod.MultivalueLike,
	mappings col.CatalogLike[string, mod.AbstractionLike],
) mod.MultivalueLike {
	var parameterList = parameterized.GetParameterList()
	var parameters = parameterList.GetParameters()
	var replacedParameters = v.replaceParameterTypes(parameters, mappings)
	parameterList = mod.ParameterListClass().ParameterList(replacedParameters)
	parameterized = mod.MultivalueClass().Multivalue(
		"(",
		parameterList,
		")",
	)
	return parameterized
}

func (v *classSynthesizer_) replaceParameterType(
	parameter mod.ParameterLike,
	mappings col.CatalogLike[string, mod.AbstractionLike],
) mod.ParameterLike {
	var parameterName = parameter.GetName()
	var abstraction = parameter.GetAbstraction()
	abstraction = v.replaceAbstractionType(abstraction, mappings)
	parameter = mod.ParameterClass().Parameter(
		parameterName,
		abstraction,
		",",
	)
	return parameter
}

func (v *classSynthesizer_) replaceParameterTypes(
	list col.ListLike[mod.ParameterLike],
	mappings col.CatalogLike[string, mod.AbstractionLike],
) col.ListLike[mod.ParameterLike] {
	var replacedParameters = col.List[mod.ParameterLike]()
	var parameters = list.GetIterator()
	for parameters.HasNext() {
		var parameter = parameters.GetNext()
		parameter = v.replaceParameterType(parameter, mappings)
		replacedParameters.AppendValue(parameter)
	}
	return replacedParameters
}

func (v *classSynthesizer_) replaceResultType(
	result mod.ResultLike,
	mappings col.CatalogLike[string, mod.AbstractionLike],
) mod.ResultLike {
	switch actual := result.GetAny().(type) {
	case mod.NoneLike:
	case mod.AbstractionLike:
		var abstraction = actual
		abstraction = v.replaceAbstractionType(abstraction, mappings)
		result = mod.ResultClass().Result(abstraction)
	case mod.MultivalueLike:
		var parameterized = actual
		parameterized = v.replaceMultivalueTypes(parameterized, mappings)
		result = mod.ResultClass().Result(parameterized)
	default:
		var message = fmt.Sprintf(
			"An unknown result type was found: %T",
			actual,
		)
		panic(message)
	}
	return result
}

func (v *classSynthesizer_) replaceNameType(
	name string,
	mappings col.CatalogLike[string, mod.AbstractionLike],
) string {
	var mappedType = mappings.GetValue(name)
	if uti.IsDefined(mappedType) {
		name = v.extractType(mappedType)
	}
	return name
}

func (v *classSynthesizer_) replaceWrapperType(
	wrapper mod.WrapperLike,
	mappings col.CatalogLike[string, mod.AbstractionLike],
) mod.WrapperLike {
	if uti.IsUndefined(wrapper) {
		return wrapper
	}
	switch actual := wrapper.GetAny().(type) {
	case mod.ArrayLike:
		// Example: [] -> []
	case mod.MapLike:
		// Example: map[K] -> map[string]
		var typeName = actual.GetName()
		var mappedType = mappings.GetValue(typeName)
		if uti.IsDefined(mappedType) {
			typeName = mappedType.GetName()
			var map_ = mod.MapClass().Map(
				"map",
				"[",
				typeName,
				"]",
			)
			wrapper = mod.WrapperClass().Wrapper(map_)
		}
	case mod.ChannelLike:
		// Example: chan -> chan
	}
	return wrapper
}

// Instance Structure

type classSynthesizer_ struct {
	// Declare the instance attributes.
	packageAnalyzer_ ana.PackageAnalyzerLike
	classAnalyzer_   ana.ClassAnalyzerLike
}

// Class Structure

type classSynthesizerClass_ struct {
	// Declare the class constants.
	warningMessage_          string
	importedPackages_        string
	importedPath_            string
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
	instanceMultiFunction_   string
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
	mapInitialization_       string
	classReference_          string
	classMap_                string
}

// Class Reference

func classSynthesizerClass() *classSynthesizerClass_ {
	return classSynthesizerClassReference_
}

var classSynthesizerClassReference_ = &classSynthesizerClass_{
	// Initialize the class constants.
	warningMessage_: ``,

	importedPackages_: ``,

	importedPath_: `
	<~packageAcronym> <packagePath>`,

	accessFunction_: `
// Access Function

func <~ClassName>Class<Constraints>() <~ClassName>ClassLike<Arguments> {
	return <~className>Class<Arguments>()
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
	return <~className>Class<Arguments>()
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

	instanceMultiFunction_: `
func (v <*><~className>_<Arguments>) <~MethodName>(<Parameters>) <ResultType> {
	// TBD - Add the method implementation.
	return
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

	mapInitialization_: `
			// <~constantName>_: constantValue,`,

	classReference_: `
// Class Reference

func <~className>Class() *<~className>Class_ {
	return <~className>ClassReference_
}

var <~className>ClassReference_ = &<~className>Class_{
	// Initialize the class constants.<ConstantInitializations>
}
`,

	classMap_: `
// Class Reference

var <~className>Map_ = map[string]any{}
var <~className>Mutex_ syn.Mutex

func <~className>Class<Constraints>() *<~className>Class_<Arguments> {
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
