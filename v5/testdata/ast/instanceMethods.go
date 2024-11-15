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

func InstanceMethods() InstanceMethodsClassLike {
	return instanceMethodsReference()
}

// Constructor Methods

func (c *instanceMethodsClass_) Make(
	primarySubsection PrimarySubsectionLike,
	attributeSubsection AttributeSubsectionLike,
	aspectSubsection AspectSubsectionLike,
) InstanceMethodsLike {
	if uti.IsUndefined(primarySubsection) {
		panic("The \"primarySubsection\" attribute is required by this class.")
	}
	if uti.IsUndefined(attributeSubsection) {
		panic("The \"attributeSubsection\" attribute is required by this class.")
	}
	if uti.IsUndefined(aspectSubsection) {
		panic("The \"aspectSubsection\" attribute is required by this class.")
	}
	var instance = &instanceMethods_{
		// Initialize the instance attributes.
		primarySubsection_:   primarySubsection,
		attributeSubsection_: attributeSubsection,
		aspectSubsection_:    aspectSubsection,
	}
	return instance

}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *instanceMethods_) GetClass() InstanceMethodsClassLike {
	return instanceMethodsReference()
}

// Attribute Methods

func (v *instanceMethods_) GetPrimarySubsection() PrimarySubsectionLike {
	return v.primarySubsection_
}

func (v *instanceMethods_) GetAttributeSubsection() AttributeSubsectionLike {
	return v.attributeSubsection_
}

func (v *instanceMethods_) GetAspectSubsection() AspectSubsectionLike {
	return v.aspectSubsection_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type instanceMethods_ struct {
	// Declare the instance attributes.
	primarySubsection_   PrimarySubsectionLike
	attributeSubsection_ AttributeSubsectionLike
	aspectSubsection_    AspectSubsectionLike
}

// Class Structure

type instanceMethodsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func instanceMethodsReference() *instanceMethodsClass_ {
	return instanceMethodsReference_
}

var instanceMethodsReference_ = &instanceMethodsClass_{
	// Initialize the class constants.
}
