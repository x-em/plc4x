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

// BACnetReadAccessProperty is the corresponding interface of BACnetReadAccessProperty
type BACnetReadAccessProperty interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() BACnetContextTagUnsignedInteger
	// GetReadResult returns ReadResult (property field)
	GetReadResult() BACnetReadAccessPropertyReadResult
	// IsBACnetReadAccessProperty is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetReadAccessProperty()
	// CreateBuilder creates a BACnetReadAccessPropertyBuilder
	CreateBACnetReadAccessPropertyBuilder() BACnetReadAccessPropertyBuilder
}

// _BACnetReadAccessProperty is the data-structure of this message
type _BACnetReadAccessProperty struct {
	PropertyIdentifier BACnetPropertyIdentifierTagged
	ArrayIndex         BACnetContextTagUnsignedInteger
	ReadResult         BACnetReadAccessPropertyReadResult

	// Arguments.
	ObjectTypeArgument BACnetObjectType
}

var _ BACnetReadAccessProperty = (*_BACnetReadAccessProperty)(nil)

// NewBACnetReadAccessProperty factory function for _BACnetReadAccessProperty
func NewBACnetReadAccessProperty(propertyIdentifier BACnetPropertyIdentifierTagged, arrayIndex BACnetContextTagUnsignedInteger, readResult BACnetReadAccessPropertyReadResult, objectTypeArgument BACnetObjectType) *_BACnetReadAccessProperty {
	if propertyIdentifier == nil {
		panic("propertyIdentifier of type BACnetPropertyIdentifierTagged for BACnetReadAccessProperty must not be nil")
	}
	return &_BACnetReadAccessProperty{PropertyIdentifier: propertyIdentifier, ArrayIndex: arrayIndex, ReadResult: readResult, ObjectTypeArgument: objectTypeArgument}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetReadAccessPropertyBuilder is a builder for BACnetReadAccessProperty
type BACnetReadAccessPropertyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(propertyIdentifier BACnetPropertyIdentifierTagged) BACnetReadAccessPropertyBuilder
	// WithPropertyIdentifier adds PropertyIdentifier (property field)
	WithPropertyIdentifier(BACnetPropertyIdentifierTagged) BACnetReadAccessPropertyBuilder
	// WithPropertyIdentifierBuilder adds PropertyIdentifier (property field) which is build by the builder
	WithPropertyIdentifierBuilder(func(BACnetPropertyIdentifierTaggedBuilder) BACnetPropertyIdentifierTaggedBuilder) BACnetReadAccessPropertyBuilder
	// WithArrayIndex adds ArrayIndex (property field)
	WithOptionalArrayIndex(BACnetContextTagUnsignedInteger) BACnetReadAccessPropertyBuilder
	// WithOptionalArrayIndexBuilder adds ArrayIndex (property field) which is build by the builder
	WithOptionalArrayIndexBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetReadAccessPropertyBuilder
	// WithReadResult adds ReadResult (property field)
	WithOptionalReadResult(BACnetReadAccessPropertyReadResult) BACnetReadAccessPropertyBuilder
	// WithOptionalReadResultBuilder adds ReadResult (property field) which is build by the builder
	WithOptionalReadResultBuilder(func(BACnetReadAccessPropertyReadResultBuilder) BACnetReadAccessPropertyReadResultBuilder) BACnetReadAccessPropertyBuilder
	// WithArgObjectTypeArgument sets a parser argument
	WithArgObjectTypeArgument(BACnetObjectType) BACnetReadAccessPropertyBuilder
	// Build builds the BACnetReadAccessProperty or returns an error if something is wrong
	Build() (BACnetReadAccessProperty, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetReadAccessProperty
}

// NewBACnetReadAccessPropertyBuilder() creates a BACnetReadAccessPropertyBuilder
func NewBACnetReadAccessPropertyBuilder() BACnetReadAccessPropertyBuilder {
	return &_BACnetReadAccessPropertyBuilder{_BACnetReadAccessProperty: new(_BACnetReadAccessProperty)}
}

type _BACnetReadAccessPropertyBuilder struct {
	*_BACnetReadAccessProperty

	err *utils.MultiError
}

var _ (BACnetReadAccessPropertyBuilder) = (*_BACnetReadAccessPropertyBuilder)(nil)

func (b *_BACnetReadAccessPropertyBuilder) WithMandatoryFields(propertyIdentifier BACnetPropertyIdentifierTagged) BACnetReadAccessPropertyBuilder {
	return b.WithPropertyIdentifier(propertyIdentifier)
}

func (b *_BACnetReadAccessPropertyBuilder) WithPropertyIdentifier(propertyIdentifier BACnetPropertyIdentifierTagged) BACnetReadAccessPropertyBuilder {
	b.PropertyIdentifier = propertyIdentifier
	return b
}

func (b *_BACnetReadAccessPropertyBuilder) WithPropertyIdentifierBuilder(builderSupplier func(BACnetPropertyIdentifierTaggedBuilder) BACnetPropertyIdentifierTaggedBuilder) BACnetReadAccessPropertyBuilder {
	builder := builderSupplier(b.PropertyIdentifier.CreateBACnetPropertyIdentifierTaggedBuilder())
	var err error
	b.PropertyIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetPropertyIdentifierTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetReadAccessPropertyBuilder) WithOptionalArrayIndex(arrayIndex BACnetContextTagUnsignedInteger) BACnetReadAccessPropertyBuilder {
	b.ArrayIndex = arrayIndex
	return b
}

func (b *_BACnetReadAccessPropertyBuilder) WithOptionalArrayIndexBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetReadAccessPropertyBuilder {
	builder := builderSupplier(b.ArrayIndex.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.ArrayIndex, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetReadAccessPropertyBuilder) WithOptionalReadResult(readResult BACnetReadAccessPropertyReadResult) BACnetReadAccessPropertyBuilder {
	b.ReadResult = readResult
	return b
}

func (b *_BACnetReadAccessPropertyBuilder) WithOptionalReadResultBuilder(builderSupplier func(BACnetReadAccessPropertyReadResultBuilder) BACnetReadAccessPropertyReadResultBuilder) BACnetReadAccessPropertyBuilder {
	builder := builderSupplier(b.ReadResult.CreateBACnetReadAccessPropertyReadResultBuilder())
	var err error
	b.ReadResult, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetReadAccessPropertyReadResultBuilder failed"))
	}
	return b
}

func (b *_BACnetReadAccessPropertyBuilder) WithArgObjectTypeArgument(objectTypeArgument BACnetObjectType) BACnetReadAccessPropertyBuilder {
	b.ObjectTypeArgument = objectTypeArgument
	return b
}

func (b *_BACnetReadAccessPropertyBuilder) Build() (BACnetReadAccessProperty, error) {
	if b.PropertyIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'propertyIdentifier' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetReadAccessProperty.deepCopy(), nil
}

func (b *_BACnetReadAccessPropertyBuilder) MustBuild() BACnetReadAccessProperty {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetReadAccessPropertyBuilder) DeepCopy() any {
	_copy := b.CreateBACnetReadAccessPropertyBuilder().(*_BACnetReadAccessPropertyBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetReadAccessPropertyBuilder creates a BACnetReadAccessPropertyBuilder
func (b *_BACnetReadAccessProperty) CreateBACnetReadAccessPropertyBuilder() BACnetReadAccessPropertyBuilder {
	if b == nil {
		return NewBACnetReadAccessPropertyBuilder()
	}
	return &_BACnetReadAccessPropertyBuilder{_BACnetReadAccessProperty: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetReadAccessProperty) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetReadAccessProperty) GetArrayIndex() BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

func (m *_BACnetReadAccessProperty) GetReadResult() BACnetReadAccessPropertyReadResult {
	return m.ReadResult
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetReadAccessProperty(structType any) BACnetReadAccessProperty {
	if casted, ok := structType.(BACnetReadAccessProperty); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetReadAccessProperty); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetReadAccessProperty) GetTypeName() string {
	return "BACnetReadAccessProperty"
}

func (m *_BACnetReadAccessProperty) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += m.ArrayIndex.GetLengthInBits(ctx)
	}

	// Optional Field (readResult)
	if m.ReadResult != nil {
		lengthInBits += m.ReadResult.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetReadAccessProperty) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetReadAccessPropertyParse(ctx context.Context, theBytes []byte, objectTypeArgument BACnetObjectType) (BACnetReadAccessProperty, error) {
	return BACnetReadAccessPropertyParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), objectTypeArgument)
}

func BACnetReadAccessPropertyParseWithBufferProducer(objectTypeArgument BACnetObjectType) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetReadAccessProperty, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetReadAccessProperty, error) {
		return BACnetReadAccessPropertyParseWithBuffer(ctx, readBuffer, objectTypeArgument)
	}
}

func BACnetReadAccessPropertyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetReadAccessProperty, error) {
	v, err := (&_BACnetReadAccessProperty{ObjectTypeArgument: objectTypeArgument}).parse(ctx, readBuffer, objectTypeArgument)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetReadAccessProperty) parse(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (__bACnetReadAccessProperty BACnetReadAccessProperty, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetReadAccessProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetReadAccessProperty")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	propertyIdentifier, err := ReadSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", ReadComplex[BACnetPropertyIdentifierTagged](BACnetPropertyIdentifierTaggedParseWithBufferProducer((uint8)(uint8(2)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyIdentifier' field"))
	}
	m.PropertyIdentifier = propertyIdentifier

	var arrayIndex BACnetContextTagUnsignedInteger
	_arrayIndex, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "arrayIndex", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayIndex' field"))
	}
	if _arrayIndex != nil {
		arrayIndex = *_arrayIndex
		m.ArrayIndex = arrayIndex
	}

	var readResult BACnetReadAccessPropertyReadResult
	_readResult, err := ReadOptionalField[BACnetReadAccessPropertyReadResult](ctx, "readResult", ReadComplex[BACnetReadAccessPropertyReadResult](BACnetReadAccessPropertyReadResultParseWithBufferProducer((BACnetObjectType)(objectTypeArgument), (BACnetPropertyIdentifier)(propertyIdentifier.GetValue()), (BACnetTagPayloadUnsignedInteger)((CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool((arrayIndex) != (nil)), func() any { return CastBACnetTagPayloadUnsignedInteger((arrayIndex).GetPayload()) }, func() any { return CastBACnetTagPayloadUnsignedInteger(nil) }))))), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'readResult' field"))
	}
	if _readResult != nil {
		readResult = *_readResult
		m.ReadResult = readResult
	}

	if closeErr := readBuffer.CloseContext("BACnetReadAccessProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetReadAccessProperty")
	}

	return m, nil
}

func (m *_BACnetReadAccessProperty) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetReadAccessProperty) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetReadAccessProperty"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetReadAccessProperty")
	}

	if err := WriteSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", m.GetPropertyIdentifier(), WriteComplex[BACnetPropertyIdentifierTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'propertyIdentifier' field")
	}

	if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "arrayIndex", GetRef(m.GetArrayIndex()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'arrayIndex' field")
	}

	if err := WriteOptionalField[BACnetReadAccessPropertyReadResult](ctx, "readResult", GetRef(m.GetReadResult()), WriteComplex[BACnetReadAccessPropertyReadResult](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'readResult' field")
	}

	if popErr := writeBuffer.PopContext("BACnetReadAccessProperty"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetReadAccessProperty")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetReadAccessProperty) GetObjectTypeArgument() BACnetObjectType {
	return m.ObjectTypeArgument
}

//
////

func (m *_BACnetReadAccessProperty) IsBACnetReadAccessProperty() {}

func (m *_BACnetReadAccessProperty) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetReadAccessProperty) deepCopy() *_BACnetReadAccessProperty {
	if m == nil {
		return nil
	}
	_BACnetReadAccessPropertyCopy := &_BACnetReadAccessProperty{
		utils.DeepCopy[BACnetPropertyIdentifierTagged](m.PropertyIdentifier),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.ArrayIndex),
		utils.DeepCopy[BACnetReadAccessPropertyReadResult](m.ReadResult),
		m.ObjectTypeArgument,
	}
	return _BACnetReadAccessPropertyCopy
}

func (m *_BACnetReadAccessProperty) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}