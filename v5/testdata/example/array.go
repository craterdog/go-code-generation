/*
................................................................................
.                  Copyright (c) 2024.  All Rights Reserved.                   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package example

import (
	fmt "fmt"
	syn "sync"
)

// CLASS INTERFACE

// Access Function

func Array[V any]() ArrayClassLike[V] {
	return arrayReference[V]()
}

// Constructor Methods

func (c *arrayClass_[V]) Make(
	intrinsic []V,
) ArrayLike[V] {
	var instance = array_[V](intrinsic)
	return instance
}

func (c *arrayClass_[V]) MakeWithSize(
	size Ordinal,
) ArrayLike[V] {
	var instance ArrayLike[V]
	// TBD - Add the constructor implementation.
	return instance
}

func (c *arrayClass_[V]) MakeFromSequence(
	values Sequential[V],
) ArrayLike[V] {
	var instance ArrayLike[V]
	// TBD - Add the constructor implementation.
	return instance
}

// Constant Methods

// Function Methods

func (c *arrayClass_[V]) Merge(
	first ArrayLike[V],
	second ArrayLike[V],
) ArrayLike[V] {
	var result_ ArrayLike[V]
	// TBD - Add the function implementation.
	return result_
}

// INSTANCE INTERFACE

// Principal Methods

func (v array_[V]) GetClass() ArrayClassLike[V] {
	return arrayReference[V]()
}

func (v array_[V]) GetIntrinsic() []V {
	return []V(v)
}

func (v array_[V]) SortValuesWithRanker(
	ranker RankingFunction[V],
) {
	// TBD - Add the method implementation.
}

// Attribute Methods

// Accessible[V] Methods

func (v array_[V]) GetValue(
	index Ordinal,
) V {
	var result_ V
	// TBD - Add the method implementation.
	return result_
}

func (v array_[V]) GetValues(
	first Ordinal,
	last Ordinal,
) Sequential[V] {
	var result_ Sequential[V]
	// TBD - Add the method implementation.
	return result_
}

// Sequential[V] Methods

func (v array_[V]) IsEmpty() bool {
	var result_ bool
	// TBD - Add the method implementation.
	return result_
}

func (v array_[V]) GetSize() Ordinal {
	var result_ Ordinal
	// TBD - Add the method implementation.
	return result_
}

func (v array_[V]) AsArray() []V {
	var result_ []V
	// TBD - Add the method implementation.
	return result_
}

// Updatable[V] Methods

func (v array_[V]) SetValue(
	index Ordinal,
	value V,
) {
	// TBD - Add the method implementation.
}

func (v array_[V]) SetValues(
	index Ordinal,
	values Sequential[V],
) {
	// TBD - Add the method implementation.
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type array_[V any] []V

// Class Structure

type arrayClass_[V any] struct {
	// Declare the class constants.
}

// Class Reference

var arrayMap_ = map[string]any{}
var arrayMutex_ syn.Mutex

func arrayReference[V any]() *arrayClass_[V] {
	// Generate the name of the bound class type.
	var class *arrayClass_[V]
	var name = fmt.Sprintf("%T", class)

	// Check for an existing bound class type.
	arrayMutex_.Lock()
	var value = arrayMap_[name]
	switch actual := value.(type) {
	case *arrayClass_[V]:
		// This bound class type already exists.
		class = actual
	default:
		// Add a new bound class type.
		class = &arrayClass_[V]{
			// Initialize the class constants.
		}
		arrayMap_[name] = class
	}
	arrayMutex_.Unlock()

	// Return a reference to the bound class type.
	return class
}
