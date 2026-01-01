/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
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
	mod "github.com/craterdog/go-class-model/v8"
	ana "github.com/craterdog/go-code-generation/v8/analyzer"
	com "github.com/craterdog/go-essential-composites/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func NodeSynthesizerClass() NodeSynthesizerClassLike {
	return nodeSynthesizerClass()
}

// Constructor Methods

func (c *nodeSynthesizerClass_) NodeSynthesizer(
	model mod.ModelLike,
	className string,
) NodeSynthesizerLike {
	var instance = &nodeSynthesizer_{
		// Initialize the instance attributes.
		packageAnalyzer_: ana.PackageAnalyzerClass().PackageAnalyzer(model),
		classAnalyzer_:   ana.ClassAnalyzerClass().ClassAnalyzer(model, className),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *nodeSynthesizer_) GetClass() NodeSynthesizerClassLike {
	return nodeSynthesizerClass()
}

// TemplateDriven Methods

func (v *nodeSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.packageAnalyzer_.GetLegalNotice()
	return legalNotice
}

func (v *nodeSynthesizer_) CreateWarningMessage() string {
	var class = nodeSynthesizerClass()
	var warningMessage = class.warningMessage_
	return warningMessage
}

func (v *nodeSynthesizer_) CreateImportedPackages() string {
	var class = nodeSynthesizerClass()
	var importedPackages = class.importedPackages_
	return importedPackages
}

func (v *nodeSynthesizer_) CreateAccessFunction() string {
	var class = nodeSynthesizerClass()
	var accessFunction = class.accessFunction_
	return accessFunction
}

func (v *nodeSynthesizer_) CreateConstantMethods() string {
	var constantMethods string
	return constantMethods
}

func (v *nodeSynthesizer_) CreateConstructorMethods() string {
	var methods = v.classAnalyzer_.GetConstructorMethods()
	var constructorMethods = v.createConstructorMethods(methods)
	return constructorMethods
}

func (v *nodeSynthesizer_) CreateFunctionMethods() string {
	var functionMethods string
	return functionMethods
}

func (v *nodeSynthesizer_) CreatePrincipalMethods() string {
	var methods = v.classAnalyzer_.GetPrincipalMethods()
	var principalMethods = v.createPrincipalMethods(methods)
	return principalMethods
}

func (v *nodeSynthesizer_) CreateAttributeMethods() string {
	var methods = v.classAnalyzer_.GetAttributeMethods()
	var attributeMethods = v.createAttributeMethods(methods)
	return attributeMethods
}

func (v *nodeSynthesizer_) CreateAspectInterfaces() string {
	var aspectInterfaces string
	return aspectInterfaces
}

func (v *nodeSynthesizer_) CreatePrivateMethods() string {
	var privateMethods string
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

func (v *nodeSynthesizer_) CreateClass() string {
	var classReference = v.createClass()
	return classReference
}

func (v *nodeSynthesizer_) PerformGlobalUpdates(
	existing string,
	generated string,
) string {
	return generated
}

// PROTECTED INTERFACE

// Private Methods

func (v *nodeSynthesizer_) createAttributeCheck(
	parameter mod.ParameterLike,
) string {
	var attributeCheck string
	var parameterName = parameter.GetName()
	var attributeName = sts.TrimSuffix(parameterName, "_")
	// Ignore optional attributes.
	if !sts.HasPrefix(attributeName, "optional") {
		var class = nodeSynthesizerClass()
		attributeCheck = class.attributeCheck_
		attributeCheck = uti.ReplaceAll(
			attributeCheck,
			"attributeName",
			attributeName,
		)
	}
	return attributeCheck
}

func (v *nodeSynthesizer_) createAttributeChecks(
	sequence com.Sequential[mod.ParameterLike],
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
	var attributes = v.classAnalyzer_.GetAttributes().GetIterator()
	for attributes.HasNext() {
		var attribute = attributes.GetNext()
		var attributeName = attribute.GetKey()
		var attributeType = attribute.GetValue()
		var class = nodeSynthesizerClass()
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

func (v *nodeSynthesizer_) createAttributeInitializations(
	sequence com.Sequential[mod.ParameterLike],
) string {
	var initializations string
	var parameters = sequence.GetIterator()
	for parameters.HasNext() {
		var parameter = parameters.GetNext()
		var parameterName = parameter.GetName()
		var attributeName = sts.TrimSuffix(parameterName, "_")
		if uti.IsDefined(v.classAnalyzer_.GetAttributes().GetValue(attributeName)) {
			var class = nodeSynthesizerClass()
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

func (v *nodeSynthesizer_) createAttributeMethods(
	sequence com.Sequential[mod.AttributeMethodLike],
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
		var class = nodeSynthesizerClass()
		methods = class.attributeMethods_
		methods = uti.ReplaceAll(
			methods,
			"attributeMethods",
			attributeMethods,
		)
	}
	return methods
}

func (v *nodeSynthesizer_) createClass() string {
	var class = nodeSynthesizerClass()
	var classReference = class.classReference_
	return classReference
}

func (v *nodeSynthesizer_) createClassStructure() string {
	var class = nodeSynthesizerClass()
	var classStructure = class.classStructure_
	return classStructure
}

func (v *nodeSynthesizer_) createConstructorMethod(
	constructorMethod mod.ConstructorMethodLike,
) string {
	var methodName = constructorMethod.GetName()
	var parameters string
	var parameterList = constructorMethod.GetOptionalParameterList()
	if uti.IsDefined(parameterList) {
		var constructorParameters = parameterList.GetParameters()
		parameters = v.createParameters(constructorParameters)
	}
	var resultType = v.extractType(constructorMethod.GetAbstraction())
	var instanceInstantiation = v.createInstanceInstantiation(constructorMethod)
	var class = nodeSynthesizerClass()
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

func (v *nodeSynthesizer_) createConstructorMethods(
	sequence com.Sequential[mod.ConstructorMethodLike],
) string {
	var methods string
	if uti.IsDefined(sequence) {
		var constructorMethods string
		var constructors = sequence.GetIterator()
		for constructors.HasNext() {
			var constructorMethod = constructors.GetNext()
			constructorMethods += v.createConstructorMethod(constructorMethod)
		}
		var class = nodeSynthesizerClass()
		methods = class.constructorMethods_
		methods = uti.ReplaceAll(
			methods,
			"constructorMethods",
			constructorMethods,
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
	var class = nodeSynthesizerClass()
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

func (v *nodeSynthesizer_) createInstanceInstantiation(
	constructorMethod mod.ConstructorMethodLike,
) string {
	var attributeChecks string
	var attributeInitializations string
	var parameterList = constructorMethod.GetOptionalParameterList()
	if uti.IsDefined(parameterList) {
		var constructorParameters = parameterList.GetParameters()
		attributeChecks = v.createAttributeChecks(constructorParameters)
		attributeInitializations = v.createAttributeInitializations(
			constructorParameters,
		)
	}
	var class = nodeSynthesizerClass()
	var instantiation = class.structureInstantiation_
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
	return instantiation
}

func (v *nodeSynthesizer_) createInstanceStructure() string {
	var class = nodeSynthesizerClass()
	var structure = class.instanceStructure_
	var attributeDeclarations = v.createAttributeDeclarations()
	structure = uti.ReplaceAll(
		structure,
		"attributeDeclarations",
		attributeDeclarations,
	)
	return structure
}

func (v *nodeSynthesizer_) createParameters(
	sequence com.Sequential[mod.ParameterLike],
) string {
	var methodParameters string
	var parameters = sequence.GetIterator()
	for parameters.HasNext() {
		var parameter = parameters.GetNext()
		var parameterName = parameter.GetName()
		var parameterType = v.extractType(parameter.GetAbstraction())
		var class = nodeSynthesizerClass()
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

func (v *nodeSynthesizer_) createPrincipalMethods(
	sequence com.Sequential[mod.PrincipalMethodLike],
) string {
	var class = nodeSynthesizerClass()
	var implementation = class.principalMethods_
	return implementation
}

func (v *nodeSynthesizer_) createSetterMethod(
	setterMethod mod.SetterMethodLike,
) string {
	var methodName = setterMethod.GetName()
	var attributeName = v.extractAttributeName(methodName)
	var parameter = setterMethod.GetParameter()
	var attributeType = v.extractType(parameter.GetAbstraction())
	var attributeCheck = v.createAttributeCheck(parameter)
	var class = nodeSynthesizerClass()
	var method = class.setterMethod_
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

func (v *nodeSynthesizer_) extractType(
	abstraction mod.AbstractionLike,
) string {
	var abstractType string
	var wrapper = abstraction.GetOptionalWrapper()
	if uti.IsDefined(wrapper) {
		switch actual := wrapper.GetAny().(type) {
		case mod.DotsLike:
			abstractType = "..."
		case mod.StarLike:
			abstractType = "*"
		case mod.ArrayLike:
			abstractType = "[]"
		case mod.ChannelLike:
			abstractType = "chan "
		case mod.MapLike:
			abstractType = "map[" + actual.GetName() + "]"
		}
	}
	switch actual := abstraction.GetType().GetAny().(type) {
	case mod.NamedLike:
		var prefix = actual.GetOptionalPrefix()
		if uti.IsDefined(prefix) {
			abstractType += prefix
		}
		var name = actual.GetName()
		abstractType += name
		var arguments = actual.GetOptionalArguments()
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
	case mod.FunctionalLike:
		abstractType += "func("
		var parameterList = actual.GetOptionalParameterList()
		if uti.IsDefined(parameterList) {
			var parameters = parameterList.GetParameters().GetIterator()
			for parameters.HasNext() {
				var parameter = parameters.GetNext()
				var parameterName = parameter.GetName()
				var parameterType = v.extractType(
					parameter.GetAbstraction(),
				)
				var class = nodeSynthesizerClass()
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
				abstractType += methodParameter
			}
			abstractType += "\n"
		}
		abstractType += ")"
		var result = actual.GetOptionalResult()
		if uti.IsDefined(result) {
			switch actual := result.GetAny().(type) {
			case mod.NoneLike:
			case mod.AbstractionLike:
				abstractType += " " + v.extractType(actual)
			case mod.MultivalueLike:
				abstractType += " ("
				var parameterList = actual.GetParameterList()
				var parameters = parameterList.GetParameters().GetIterator()
				for parameters.HasNext() {
					var parameter = parameters.GetNext()
					var parameterName = parameter.GetName()
					var parameterType = v.extractType(
						parameter.GetAbstraction(),
					)
					var class = nodeSynthesizerClass()
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
					abstractType += methodParameter
				}
				abstractType += "\n)"
			}
		}
	}
	return abstractType
}

// Instance Structure

type nodeSynthesizer_ struct {
	// Declare the instance attributes.
	packageAnalyzer_ ana.PackageAnalyzerLike
	classAnalyzer_   ana.ClassAnalyzerLike
}

// Class Structure

type nodeSynthesizerClass_ struct {
	// Declare the class constants.
	warningMessage_          string
	importedPackages_        string
	accessFunction_          string
	constructorMethods_      string
	constructorMethod_       string
	structureInstantiation_  string
	attributeMethods_        string
	getterMethod_            string
	setterMethod_            string
	principalMethods_        string
	instanceMethod_          string
	methodParameter_         string
	instanceStructure_       string
	attributeDeclaration_    string
	attributeCheck_          string
	attributeInitialization_ string
	classStructure_          string
	classReference_          string
}

// Class Reference

func nodeSynthesizerClass() *nodeSynthesizerClass_ {
	return nodeSynthesizerClassReference_
}

var nodeSynthesizerClassReference_ = &nodeSynthesizerClass_{
	// Initialize the class constants.
	warningMessage_: `
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│              This class file was automatically generated using:              │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
`,

	importedPackages_: `
	com "github.com/craterdog/go-essential-composites/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
`,

	accessFunction_: `
// Access Function

func <~ClassName>Class() <~ClassName>ClassLike {
	return <~className>Class()
}
`,

	constructorMethods_: `
// Constructor Methods
<ConstructorMethods>`,

	constructorMethod_: `
func (c *<~className>Class_) <MethodName>(<Parameters>) <~ClassName>Like {<InstanceInstantiation>}
`,

	structureInstantiation_: `<AttributeChecks>
	var instance = &<~className>_{
		// Initialize the instance attributes.<AttributeInitializations>
	}
	return instance
`,

	attributeMethods_: `
// Attribute Methods
<AttributeMethods>`,

	getterMethod_: `
func (v *<~className>_) <~MethodName>() <AttributeType> {
	return v.<~attributeName>_
}
`,

	setterMethod_: `
func (v *<~className>_) <~MethodName>(
	<attributeName_> <AttributeType>,
) {<AttributeCheck>
	v.<~attributeName>_ = <attributeName_>
}
`,

	principalMethods_: `
// Principal Methods

func (v *<~className>_) GetClass() <~ClassName>ClassLike {
	return <~className>Class()
}
`,

	instanceMethod_: `
func (v *<~className>_) <~MethodName>(<Parameters>) {
	// TBD - Add the method implementation.
}
`,

	methodParameter_: `
	<parameterName_> <ParameterType>,`,

	instanceStructure_: `
// Instance Structure

type <~className>_ struct {
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

type <~className>Class_ struct {
	// Declare the class constants.
}
`,

	classReference_: `
// Class Reference

func <~className>Class() *<~className>Class_ {
	return <~className>ClassReference_
}

var <~className>ClassReference_ = &<~className>Class_{
	// Initialize the class constants.
}
`,
}
