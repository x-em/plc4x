/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ExtensionObject is the corresponding interface of ExtensionObject
type ExtensionObject interface {
	ExtensionObjectContract
	ExtensionObjectRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsExtensionObject is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsExtensionObject()
	// CreateBuilder creates a ExtensionObjectBuilder
	CreateExtensionObjectBuilder() ExtensionObjectBuilder
}

// ExtensionObjectContract provides a set of functions which can be overwritten by a sub struct
type ExtensionObjectContract interface {
	// GetTypeId returns TypeId (property field)
	GetTypeId() ExpandedNodeId
	// GetExtensionId returns ExtensionId (virtual field)
	GetExtensionId() int32
	// GetBody returns Body (abstract field)
	GetBody() ExtensionObjectDefinition
	// IsExtensionObject is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsExtensionObject()
	// CreateBuilder creates a ExtensionObjectBuilder
	CreateExtensionObjectBuilder() ExtensionObjectBuilder
}

// ExtensionObjectRequirements provides a set of functions which need to be implemented by a sub struct
type ExtensionObjectRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetIncludeEncodingMask returns IncludeEncodingMask (discriminator field)
	GetIncludeEncodingMask() bool
	// GetBody returns Body (abstract field)
	GetBody() ExtensionObjectDefinition
}

// _ExtensionObject is the data-structure of this message
type _ExtensionObject struct {
	_SubType interface {
		ExtensionObjectContract
		ExtensionObjectRequirements
	}
	TypeId ExpandedNodeId
}

var _ ExtensionObjectContract = (*_ExtensionObject)(nil)

// NewExtensionObject factory function for _ExtensionObject
func NewExtensionObject(typeId ExpandedNodeId) *_ExtensionObject {
	if typeId == nil {
		panic("typeId of type ExpandedNodeId for ExtensionObject must not be nil")
	}
	return &_ExtensionObject{TypeId: typeId}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ExtensionObjectBuilder is a builder for ExtensionObject
type ExtensionObjectBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(typeId ExpandedNodeId) ExtensionObjectBuilder
	// WithTypeId adds TypeId (property field)
	WithTypeId(ExpandedNodeId) ExtensionObjectBuilder
	// WithTypeIdBuilder adds TypeId (property field) which is build by the builder
	WithTypeIdBuilder(func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) ExtensionObjectBuilder
	// AsRootExtensionObject converts this build to a subType of ExtensionObject. It is always possible to return to current builder using Done()
	AsRootExtensionObject() RootExtensionObjectBuilder
	// AsExtensionObjectWithMask converts this build to a subType of ExtensionObject. It is always possible to return to current builder using Done()
	AsExtensionObjectWithMask() ExtensionObjectWithMaskBuilder
	// Build builds the ExtensionObject or returns an error if something is wrong
	PartialBuild() (ExtensionObjectContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() ExtensionObjectContract
	// Build builds the ExtensionObject or returns an error if something is wrong
	Build() (ExtensionObject, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ExtensionObject
}

// NewExtensionObjectBuilder() creates a ExtensionObjectBuilder
func NewExtensionObjectBuilder() ExtensionObjectBuilder {
	return &_ExtensionObjectBuilder{_ExtensionObject: new(_ExtensionObject)}
}

type _ExtensionObjectChildBuilder interface {
	utils.Copyable
	setParent(ExtensionObjectContract)
	buildForExtensionObject() (ExtensionObject, error)
}

type _ExtensionObjectBuilder struct {
	*_ExtensionObject

	childBuilder _ExtensionObjectChildBuilder

	err *utils.MultiError
}

var _ (ExtensionObjectBuilder) = (*_ExtensionObjectBuilder)(nil)

func (b *_ExtensionObjectBuilder) WithMandatoryFields(typeId ExpandedNodeId) ExtensionObjectBuilder {
	return b.WithTypeId(typeId)
}

func (b *_ExtensionObjectBuilder) WithTypeId(typeId ExpandedNodeId) ExtensionObjectBuilder {
	b.TypeId = typeId
	return b
}

func (b *_ExtensionObjectBuilder) WithTypeIdBuilder(builderSupplier func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) ExtensionObjectBuilder {
	builder := builderSupplier(b.TypeId.CreateExpandedNodeIdBuilder())
	var err error
	b.TypeId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExpandedNodeIdBuilder failed"))
	}
	return b
}

func (b *_ExtensionObjectBuilder) PartialBuild() (ExtensionObjectContract, error) {
	if b.TypeId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'typeId' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ExtensionObject.deepCopy(), nil
}

func (b *_ExtensionObjectBuilder) PartialMustBuild() ExtensionObjectContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ExtensionObjectBuilder) AsRootExtensionObject() RootExtensionObjectBuilder {
	if cb, ok := b.childBuilder.(RootExtensionObjectBuilder); ok {
		return cb
	}
	cb := NewRootExtensionObjectBuilder().(*_RootExtensionObjectBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ExtensionObjectBuilder) AsExtensionObjectWithMask() ExtensionObjectWithMaskBuilder {
	if cb, ok := b.childBuilder.(ExtensionObjectWithMaskBuilder); ok {
		return cb
	}
	cb := NewExtensionObjectWithMaskBuilder().(*_ExtensionObjectWithMaskBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ExtensionObjectBuilder) Build() (ExtensionObject, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForExtensionObject()
}

func (b *_ExtensionObjectBuilder) MustBuild() ExtensionObject {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ExtensionObjectBuilder) DeepCopy() any {
	_copy := b.CreateExtensionObjectBuilder().(*_ExtensionObjectBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_ExtensionObjectChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateExtensionObjectBuilder creates a ExtensionObjectBuilder
func (b *_ExtensionObject) CreateExtensionObjectBuilder() ExtensionObjectBuilder {
	if b == nil {
		return NewExtensionObjectBuilder()
	}
	return &_ExtensionObjectBuilder{_ExtensionObject: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ExtensionObject) GetTypeId() ExpandedNodeId {
	return m.TypeId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_ExtensionObject) GetExtensionId() int32 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return int32(utils.InlineIf(bool((m.GetTypeId()) == (nil)), func() any { return int32(int32(0)) }, func() any { return int32(ExtensionId(ctx, m.GetTypeId())) }).(int32))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for abstract fields.
///////////////////////

func (m *_ExtensionObject) GetBody() ExtensionObjectDefinition {
	return m._SubType.GetBody()
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastExtensionObject(structType any) ExtensionObject {
	if casted, ok := structType.(ExtensionObject); ok {
		return casted
	}
	if casted, ok := structType.(*ExtensionObject); ok {
		return *casted
	}
	return nil
}

func (m *_ExtensionObject) GetTypeName() string {
	return "ExtensionObject"
}

func (m *_ExtensionObject) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (typeId)
	lengthInBits += m.TypeId.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_ExtensionObject) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_ExtensionObject) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func ExtensionObjectParse[T ExtensionObject](ctx context.Context, theBytes []byte, includeEncodingMask bool) (T, error) {
	return ExtensionObjectParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), includeEncodingMask)
}

func ExtensionObjectParseWithBufferProducer[T ExtensionObject](includeEncodingMask bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := ExtensionObjectParseWithBuffer[T](ctx, readBuffer, includeEncodingMask)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func ExtensionObjectParseWithBuffer[T ExtensionObject](ctx context.Context, readBuffer utils.ReadBuffer, includeEncodingMask bool) (T, error) {
	v, err := (&_ExtensionObject{}).parse(ctx, readBuffer, includeEncodingMask)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_ExtensionObject) parse(ctx context.Context, readBuffer utils.ReadBuffer, includeEncodingMask bool) (__extensionObject ExtensionObject, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ExtensionObject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ExtensionObject")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	typeId, err := ReadSimpleField[ExpandedNodeId](ctx, "typeId", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'typeId' field"))
	}
	m.TypeId = typeId

	extensionId, err := ReadVirtualField[int32](ctx, "extensionId", (*int32)(nil), utils.InlineIf(bool((typeId) == (nil)), func() any { return int32(int32(0)) }, func() any { return int32(ExtensionId(ctx, typeId)) }).(int32))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'extensionId' field"))
	}
	_ = extensionId

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child ExtensionObject
	switch {
	case includeEncodingMask == bool(false): // RootExtensionObject
		if _child, err = new(_RootExtensionObject).parse(ctx, readBuffer, m, extensionId, includeEncodingMask); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type RootExtensionObject for type-switch of ExtensionObject")
		}
	case includeEncodingMask == bool(true): // ExtensionObjectWithMask
		if _child, err = new(_ExtensionObjectWithMask).parse(ctx, readBuffer, m, extensionId, includeEncodingMask); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ExtensionObjectWithMask for type-switch of ExtensionObject")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [includeEncodingMask=%v]", includeEncodingMask)
	}

	if closeErr := readBuffer.CloseContext("ExtensionObject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ExtensionObject")
	}

	return _child, nil
}

func (pm *_ExtensionObject) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ExtensionObject, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ExtensionObject"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ExtensionObject")
	}

	if err := WriteSimpleField[ExpandedNodeId](ctx, "typeId", m.GetTypeId(), WriteComplex[ExpandedNodeId](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'typeId' field")
	}
	// Virtual field
	extensionId := m.GetExtensionId()
	_ = extensionId
	if _extensionIdErr := writeBuffer.WriteVirtual(ctx, "extensionId", m.GetExtensionId()); _extensionIdErr != nil {
		return errors.Wrap(_extensionIdErr, "Error serializing 'extensionId' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ExtensionObject"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ExtensionObject")
	}
	return nil
}

func (m *_ExtensionObject) IsExtensionObject() {}

func (m *_ExtensionObject) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ExtensionObject) deepCopy() *_ExtensionObject {
	if m == nil {
		return nil
	}
	_ExtensionObjectCopy := &_ExtensionObject{
		nil, // will be set by child
		utils.DeepCopy[ExpandedNodeId](m.TypeId),
	}
	return _ExtensionObjectCopy
}