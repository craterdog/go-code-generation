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

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func ClassMethods() ClassMethodsClassLike {
	return classMethodsReference()
}

// Constructor Methods

func (c *classMethodsClass_) Make(
	constructorSubsection ConstructorSubsectionLike,
	constantSubsection ConstantSubsectionLike,
	functionSubsection FunctionSubsectionLike,
) ClassMethodsLike {
	if uti.IsUndefined(constructorSubsection) {
		panic("The \"constructorSubsection\" attribute is required by this class.")
	}
	if uti.IsUndefined(constantSubsection) {
		panic("The \"constantSubsection\" attribute is required by this class.")
	}
	if uti.IsUndefined(functionSubsection) {
		panic("The \"functionSubsection\" attribute is required by this class.")
	}
	var instance = &classMethods_{
		// Initialize the instance attributes.
		constructorSubsection_: constructorSubsection,
		constantSubsection_:    constantSubsection,
		functionSubsection_:    functionSubsection,
	}
	return instance

}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *classMethods_) GetClass() ClassMethodsClassLike {
	return classMethodsReference()
}

// Attribute Methods

func (v *classMethods_) GetConstructorSubsection() ConstructorSubsectionLike {
	return v.constructorSubsection_
}

func (v *classMethods_) GetConstantSubsection() ConstantSubsectionLike {
	return v.constantSubsection_
}

func (v *classMethods_) GetFunctionSubsection() FunctionSubsectionLike {
	return v.functionSubsection_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type classMethods_ struct {
	// Declare the instance attributes.
	constructorSubsection_ ConstructorSubsectionLike
	constantSubsection_    ConstantSubsectionLike
	functionSubsection_    FunctionSubsectionLike
}

// Class Structure

type classMethodsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func classMethodsReference() *classMethodsClass_ {
	return classMethodsReference_
}

var classMethodsReference_ = &classMethodsClass_{
	// Initialize the class constants.
}
