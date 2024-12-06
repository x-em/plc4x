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

// BACnetTimerStateChangeValue is the corresponding interface of BACnetTimerStateChangeValue
type BACnetTimerStateChangeValue interface {
	BACnetTimerStateChangeValueContract
	BACnetTimerStateChangeValueRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBACnetTimerStateChangeValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimerStateChangeValue()
	// CreateBuilder creates a BACnetTimerStateChangeValueBuilder
	CreateBACnetTimerStateChangeValueBuilder() BACnetTimerStateChangeValueBuilder
}

// BACnetTimerStateChangeValueContract provides a set of functions which can be overwritten by a sub struct
type BACnetTimerStateChangeValueContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetPeekedIsContextTag returns PeekedIsContextTag (virtual field)
	GetPeekedIsContextTag() bool
	// GetObjectTypeArgument() returns a parser argument
	GetObjectTypeArgument() BACnetObjectType
	// IsBACnetTimerStateChangeValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimerStateChangeValue()
	// CreateBuilder creates a BACnetTimerStateChangeValueBuilder
	CreateBACnetTimerStateChangeValueBuilder() BACnetTimerStateChangeValueBuilder
}

// BACnetTimerStateChangeValueRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetTimerStateChangeValueRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedIsContextTag returns PeekedIsContextTag (discriminator field)
	GetPeekedIsContextTag() bool
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetTimerStateChangeValue is the data-structure of this message
type _BACnetTimerStateChangeValue struct {
	_SubType interface {
		BACnetTimerStateChangeValueContract
		BACnetTimerStateChangeValueRequirements
	}
	PeekedTagHeader BACnetTagHeader

	// Arguments.
	ObjectTypeArgument BACnetObjectType
}

var _ BACnetTimerStateChangeValueContract = (*_BACnetTimerStateChangeValue)(nil)

// NewBACnetTimerStateChangeValue factory function for _BACnetTimerStateChangeValue
func NewBACnetTimerStateChangeValue(peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValue {
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetTimerStateChangeValue must not be nil")
	}
	return &_BACnetTimerStateChangeValue{PeekedTagHeader: peekedTagHeader, ObjectTypeArgument: objectTypeArgument}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTimerStateChangeValueBuilder is a builder for BACnetTimerStateChangeValue
type BACnetTimerStateChangeValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(peekedTagHeader BACnetTagHeader) BACnetTimerStateChangeValueBuilder
	// WithPeekedTagHeader adds PeekedTagHeader (property field)
	WithPeekedTagHeader(BACnetTagHeader) BACnetTimerStateChangeValueBuilder
	// WithPeekedTagHeaderBuilder adds PeekedTagHeader (property field) which is build by the builder
	WithPeekedTagHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetTimerStateChangeValueBuilder
	// WithArgObjectTypeArgument sets a parser argument
	WithArgObjectTypeArgument(BACnetObjectType) BACnetTimerStateChangeValueBuilder
	// AsBACnetTimerStateChangeValueNull converts this build to a subType of BACnetTimerStateChangeValue. It is always possible to return to current builder using Done()
	AsBACnetTimerStateChangeValueNull() BACnetTimerStateChangeValueNullBuilder
	// AsBACnetTimerStateChangeValueBoolean converts this build to a subType of BACnetTimerStateChangeValue. It is always possible to return to current builder using Done()
	AsBACnetTimerStateChangeValueBoolean() BACnetTimerStateChangeValueBooleanBuilder
	// AsBACnetTimerStateChangeValueUnsigned converts this build to a subType of BACnetTimerStateChangeValue. It is always possible to return to current builder using Done()
	AsBACnetTimerStateChangeValueUnsigned() BACnetTimerStateChangeValueUnsignedBuilder
	// AsBACnetTimerStateChangeValueInteger converts this build to a subType of BACnetTimerStateChangeValue. It is always possible to return to current builder using Done()
	AsBACnetTimerStateChangeValueInteger() BACnetTimerStateChangeValueIntegerBuilder
	// AsBACnetTimerStateChangeValueReal converts this build to a subType of BACnetTimerStateChangeValue. It is always possible to return to current builder using Done()
	AsBACnetTimerStateChangeValueReal() BACnetTimerStateChangeValueRealBuilder
	// AsBACnetTimerStateChangeValueDouble converts this build to a subType of BACnetTimerStateChangeValue. It is always possible to return to current builder using Done()
	AsBACnetTimerStateChangeValueDouble() BACnetTimerStateChangeValueDoubleBuilder
	// AsBACnetTimerStateChangeValueOctetString converts this build to a subType of BACnetTimerStateChangeValue. It is always possible to return to current builder using Done()
	AsBACnetTimerStateChangeValueOctetString() BACnetTimerStateChangeValueOctetStringBuilder
	// AsBACnetTimerStateChangeValueCharacterString converts this build to a subType of BACnetTimerStateChangeValue. It is always possible to return to current builder using Done()
	AsBACnetTimerStateChangeValueCharacterString() BACnetTimerStateChangeValueCharacterStringBuilder
	// AsBACnetTimerStateChangeValueBitString converts this build to a subType of BACnetTimerStateChangeValue. It is always possible to return to current builder using Done()
	AsBACnetTimerStateChangeValueBitString() BACnetTimerStateChangeValueBitStringBuilder
	// AsBACnetTimerStateChangeValueEnumerated converts this build to a subType of BACnetTimerStateChangeValue. It is always possible to return to current builder using Done()
	AsBACnetTimerStateChangeValueEnumerated() BACnetTimerStateChangeValueEnumeratedBuilder
	// AsBACnetTimerStateChangeValueDate converts this build to a subType of BACnetTimerStateChangeValue. It is always possible to return to current builder using Done()
	AsBACnetTimerStateChangeValueDate() BACnetTimerStateChangeValueDateBuilder
	// AsBACnetTimerStateChangeValueTime converts this build to a subType of BACnetTimerStateChangeValue. It is always possible to return to current builder using Done()
	AsBACnetTimerStateChangeValueTime() BACnetTimerStateChangeValueTimeBuilder
	// AsBACnetTimerStateChangeValueObjectidentifier converts this build to a subType of BACnetTimerStateChangeValue. It is always possible to return to current builder using Done()
	AsBACnetTimerStateChangeValueObjectidentifier() BACnetTimerStateChangeValueObjectidentifierBuilder
	// AsBACnetTimerStateChangeValueNoValue converts this build to a subType of BACnetTimerStateChangeValue. It is always possible to return to current builder using Done()
	AsBACnetTimerStateChangeValueNoValue() BACnetTimerStateChangeValueNoValueBuilder
	// AsBACnetTimerStateChangeValueConstructedValue converts this build to a subType of BACnetTimerStateChangeValue. It is always possible to return to current builder using Done()
	AsBACnetTimerStateChangeValueConstructedValue() BACnetTimerStateChangeValueConstructedValueBuilder
	// AsBACnetTimerStateChangeValueDateTime converts this build to a subType of BACnetTimerStateChangeValue. It is always possible to return to current builder using Done()
	AsBACnetTimerStateChangeValueDateTime() BACnetTimerStateChangeValueDateTimeBuilder
	// AsBACnetTimerStateChangeValueLightingCommand converts this build to a subType of BACnetTimerStateChangeValue. It is always possible to return to current builder using Done()
	AsBACnetTimerStateChangeValueLightingCommand() BACnetTimerStateChangeValueLightingCommandBuilder
	// Build builds the BACnetTimerStateChangeValue or returns an error if something is wrong
	PartialBuild() (BACnetTimerStateChangeValueContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() BACnetTimerStateChangeValueContract
	// Build builds the BACnetTimerStateChangeValue or returns an error if something is wrong
	Build() (BACnetTimerStateChangeValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTimerStateChangeValue
}

// NewBACnetTimerStateChangeValueBuilder() creates a BACnetTimerStateChangeValueBuilder
func NewBACnetTimerStateChangeValueBuilder() BACnetTimerStateChangeValueBuilder {
	return &_BACnetTimerStateChangeValueBuilder{_BACnetTimerStateChangeValue: new(_BACnetTimerStateChangeValue)}
}

type _BACnetTimerStateChangeValueChildBuilder interface {
	utils.Copyable
	setParent(BACnetTimerStateChangeValueContract)
	buildForBACnetTimerStateChangeValue() (BACnetTimerStateChangeValue, error)
}

type _BACnetTimerStateChangeValueBuilder struct {
	*_BACnetTimerStateChangeValue

	childBuilder _BACnetTimerStateChangeValueChildBuilder

	err *utils.MultiError
}

var _ (BACnetTimerStateChangeValueBuilder) = (*_BACnetTimerStateChangeValueBuilder)(nil)

func (b *_BACnetTimerStateChangeValueBuilder) WithMandatoryFields(peekedTagHeader BACnetTagHeader) BACnetTimerStateChangeValueBuilder {
	return b.WithPeekedTagHeader(peekedTagHeader)
}

func (b *_BACnetTimerStateChangeValueBuilder) WithPeekedTagHeader(peekedTagHeader BACnetTagHeader) BACnetTimerStateChangeValueBuilder {
	b.PeekedTagHeader = peekedTagHeader
	return b
}

func (b *_BACnetTimerStateChangeValueBuilder) WithPeekedTagHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetTimerStateChangeValueBuilder {
	builder := builderSupplier(b.PeekedTagHeader.CreateBACnetTagHeaderBuilder())
	var err error
	b.PeekedTagHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return b
}

func (b *_BACnetTimerStateChangeValueBuilder) WithArgObjectTypeArgument(objectTypeArgument BACnetObjectType) BACnetTimerStateChangeValueBuilder {
	b.ObjectTypeArgument = objectTypeArgument
	return b
}

func (b *_BACnetTimerStateChangeValueBuilder) PartialBuild() (BACnetTimerStateChangeValueContract, error) {
	if b.PeekedTagHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'peekedTagHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetTimerStateChangeValue.deepCopy(), nil
}

func (b *_BACnetTimerStateChangeValueBuilder) PartialMustBuild() BACnetTimerStateChangeValueContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetTimerStateChangeValueBuilder) AsBACnetTimerStateChangeValueNull() BACnetTimerStateChangeValueNullBuilder {
	if cb, ok := b.childBuilder.(BACnetTimerStateChangeValueNullBuilder); ok {
		return cb
	}
	cb := NewBACnetTimerStateChangeValueNullBuilder().(*_BACnetTimerStateChangeValueNullBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetTimerStateChangeValueBuilder) AsBACnetTimerStateChangeValueBoolean() BACnetTimerStateChangeValueBooleanBuilder {
	if cb, ok := b.childBuilder.(BACnetTimerStateChangeValueBooleanBuilder); ok {
		return cb
	}
	cb := NewBACnetTimerStateChangeValueBooleanBuilder().(*_BACnetTimerStateChangeValueBooleanBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetTimerStateChangeValueBuilder) AsBACnetTimerStateChangeValueUnsigned() BACnetTimerStateChangeValueUnsignedBuilder {
	if cb, ok := b.childBuilder.(BACnetTimerStateChangeValueUnsignedBuilder); ok {
		return cb
	}
	cb := NewBACnetTimerStateChangeValueUnsignedBuilder().(*_BACnetTimerStateChangeValueUnsignedBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetTimerStateChangeValueBuilder) AsBACnetTimerStateChangeValueInteger() BACnetTimerStateChangeValueIntegerBuilder {
	if cb, ok := b.childBuilder.(BACnetTimerStateChangeValueIntegerBuilder); ok {
		return cb
	}
	cb := NewBACnetTimerStateChangeValueIntegerBuilder().(*_BACnetTimerStateChangeValueIntegerBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetTimerStateChangeValueBuilder) AsBACnetTimerStateChangeValueReal() BACnetTimerStateChangeValueRealBuilder {
	if cb, ok := b.childBuilder.(BACnetTimerStateChangeValueRealBuilder); ok {
		return cb
	}
	cb := NewBACnetTimerStateChangeValueRealBuilder().(*_BACnetTimerStateChangeValueRealBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetTimerStateChangeValueBuilder) AsBACnetTimerStateChangeValueDouble() BACnetTimerStateChangeValueDoubleBuilder {
	if cb, ok := b.childBuilder.(BACnetTimerStateChangeValueDoubleBuilder); ok {
		return cb
	}
	cb := NewBACnetTimerStateChangeValueDoubleBuilder().(*_BACnetTimerStateChangeValueDoubleBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetTimerStateChangeValueBuilder) AsBACnetTimerStateChangeValueOctetString() BACnetTimerStateChangeValueOctetStringBuilder {
	if cb, ok := b.childBuilder.(BACnetTimerStateChangeValueOctetStringBuilder); ok {
		return cb
	}
	cb := NewBACnetTimerStateChangeValueOctetStringBuilder().(*_BACnetTimerStateChangeValueOctetStringBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetTimerStateChangeValueBuilder) AsBACnetTimerStateChangeValueCharacterString() BACnetTimerStateChangeValueCharacterStringBuilder {
	if cb, ok := b.childBuilder.(BACnetTimerStateChangeValueCharacterStringBuilder); ok {
		return cb
	}
	cb := NewBACnetTimerStateChangeValueCharacterStringBuilder().(*_BACnetTimerStateChangeValueCharacterStringBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetTimerStateChangeValueBuilder) AsBACnetTimerStateChangeValueBitString() BACnetTimerStateChangeValueBitStringBuilder {
	if cb, ok := b.childBuilder.(BACnetTimerStateChangeValueBitStringBuilder); ok {
		return cb
	}
	cb := NewBACnetTimerStateChangeValueBitStringBuilder().(*_BACnetTimerStateChangeValueBitStringBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetTimerStateChangeValueBuilder) AsBACnetTimerStateChangeValueEnumerated() BACnetTimerStateChangeValueEnumeratedBuilder {
	if cb, ok := b.childBuilder.(BACnetTimerStateChangeValueEnumeratedBuilder); ok {
		return cb
	}
	cb := NewBACnetTimerStateChangeValueEnumeratedBuilder().(*_BACnetTimerStateChangeValueEnumeratedBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetTimerStateChangeValueBuilder) AsBACnetTimerStateChangeValueDate() BACnetTimerStateChangeValueDateBuilder {
	if cb, ok := b.childBuilder.(BACnetTimerStateChangeValueDateBuilder); ok {
		return cb
	}
	cb := NewBACnetTimerStateChangeValueDateBuilder().(*_BACnetTimerStateChangeValueDateBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetTimerStateChangeValueBuilder) AsBACnetTimerStateChangeValueTime() BACnetTimerStateChangeValueTimeBuilder {
	if cb, ok := b.childBuilder.(BACnetTimerStateChangeValueTimeBuilder); ok {
		return cb
	}
	cb := NewBACnetTimerStateChangeValueTimeBuilder().(*_BACnetTimerStateChangeValueTimeBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetTimerStateChangeValueBuilder) AsBACnetTimerStateChangeValueObjectidentifier() BACnetTimerStateChangeValueObjectidentifierBuilder {
	if cb, ok := b.childBuilder.(BACnetTimerStateChangeValueObjectidentifierBuilder); ok {
		return cb
	}
	cb := NewBACnetTimerStateChangeValueObjectidentifierBuilder().(*_BACnetTimerStateChangeValueObjectidentifierBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetTimerStateChangeValueBuilder) AsBACnetTimerStateChangeValueNoValue() BACnetTimerStateChangeValueNoValueBuilder {
	if cb, ok := b.childBuilder.(BACnetTimerStateChangeValueNoValueBuilder); ok {
		return cb
	}
	cb := NewBACnetTimerStateChangeValueNoValueBuilder().(*_BACnetTimerStateChangeValueNoValueBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetTimerStateChangeValueBuilder) AsBACnetTimerStateChangeValueConstructedValue() BACnetTimerStateChangeValueConstructedValueBuilder {
	if cb, ok := b.childBuilder.(BACnetTimerStateChangeValueConstructedValueBuilder); ok {
		return cb
	}
	cb := NewBACnetTimerStateChangeValueConstructedValueBuilder().(*_BACnetTimerStateChangeValueConstructedValueBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetTimerStateChangeValueBuilder) AsBACnetTimerStateChangeValueDateTime() BACnetTimerStateChangeValueDateTimeBuilder {
	if cb, ok := b.childBuilder.(BACnetTimerStateChangeValueDateTimeBuilder); ok {
		return cb
	}
	cb := NewBACnetTimerStateChangeValueDateTimeBuilder().(*_BACnetTimerStateChangeValueDateTimeBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetTimerStateChangeValueBuilder) AsBACnetTimerStateChangeValueLightingCommand() BACnetTimerStateChangeValueLightingCommandBuilder {
	if cb, ok := b.childBuilder.(BACnetTimerStateChangeValueLightingCommandBuilder); ok {
		return cb
	}
	cb := NewBACnetTimerStateChangeValueLightingCommandBuilder().(*_BACnetTimerStateChangeValueLightingCommandBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetTimerStateChangeValueBuilder) Build() (BACnetTimerStateChangeValue, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForBACnetTimerStateChangeValue()
}

func (b *_BACnetTimerStateChangeValueBuilder) MustBuild() BACnetTimerStateChangeValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetTimerStateChangeValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetTimerStateChangeValueBuilder().(*_BACnetTimerStateChangeValueBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_BACnetTimerStateChangeValueChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetTimerStateChangeValueBuilder creates a BACnetTimerStateChangeValueBuilder
func (b *_BACnetTimerStateChangeValue) CreateBACnetTimerStateChangeValueBuilder() BACnetTimerStateChangeValueBuilder {
	if b == nil {
		return NewBACnetTimerStateChangeValueBuilder()
	}
	return &_BACnetTimerStateChangeValueBuilder{_BACnetTimerStateChangeValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValue) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetTimerStateChangeValue) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

func (pm *_BACnetTimerStateChangeValue) GetPeekedIsContextTag() bool {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetPeekedTagHeader().GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValue(structType any) BACnetTimerStateChangeValue {
	if casted, ok := structType.(BACnetTimerStateChangeValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValue) GetTypeName() string {
	return "BACnetTimerStateChangeValue"
}

func (m *_BACnetTimerStateChangeValue) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValue) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_BACnetTimerStateChangeValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetTimerStateChangeValueParse[T BACnetTimerStateChangeValue](ctx context.Context, theBytes []byte, objectTypeArgument BACnetObjectType) (T, error) {
	return BACnetTimerStateChangeValueParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), objectTypeArgument)
}

func BACnetTimerStateChangeValueParseWithBufferProducer[T BACnetTimerStateChangeValue](objectTypeArgument BACnetObjectType) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetTimerStateChangeValueParseWithBuffer[T](ctx, readBuffer, objectTypeArgument)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetTimerStateChangeValueParseWithBuffer[T BACnetTimerStateChangeValue](ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (T, error) {
	v, err := (&_BACnetTimerStateChangeValue{ObjectTypeArgument: objectTypeArgument}).parse(ctx, readBuffer, objectTypeArgument)
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

func (m *_BACnetTimerStateChangeValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (__bACnetTimerStateChangeValue BACnetTimerStateChangeValue, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	peekedIsContextTag, err := ReadVirtualField[bool](ctx, "peekedIsContextTag", (*bool)(nil), bool((peekedTagHeader.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedIsContextTag' field"))
	}
	_ = peekedIsContextTag

	// Validation
	if !(bool((!(peekedIsContextTag))) || bool((bool(bool(peekedIsContextTag) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x6)))) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x7)))))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "unexpected opening or closing tag"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetTimerStateChangeValue
	switch {
	case peekedTagNumber == 0x0 && peekedIsContextTag == bool(false): // BACnetTimerStateChangeValueNull
		if _child, err = new(_BACnetTimerStateChangeValueNull).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetTimerStateChangeValueNull for type-switch of BACnetTimerStateChangeValue")
		}
	case peekedTagNumber == 0x1 && peekedIsContextTag == bool(false): // BACnetTimerStateChangeValueBoolean
		if _child, err = new(_BACnetTimerStateChangeValueBoolean).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetTimerStateChangeValueBoolean for type-switch of BACnetTimerStateChangeValue")
		}
	case peekedTagNumber == 0x2 && peekedIsContextTag == bool(false): // BACnetTimerStateChangeValueUnsigned
		if _child, err = new(_BACnetTimerStateChangeValueUnsigned).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetTimerStateChangeValueUnsigned for type-switch of BACnetTimerStateChangeValue")
		}
	case peekedTagNumber == 0x3 && peekedIsContextTag == bool(false): // BACnetTimerStateChangeValueInteger
		if _child, err = new(_BACnetTimerStateChangeValueInteger).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetTimerStateChangeValueInteger for type-switch of BACnetTimerStateChangeValue")
		}
	case peekedTagNumber == 0x4 && peekedIsContextTag == bool(false): // BACnetTimerStateChangeValueReal
		if _child, err = new(_BACnetTimerStateChangeValueReal).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetTimerStateChangeValueReal for type-switch of BACnetTimerStateChangeValue")
		}
	case peekedTagNumber == 0x5 && peekedIsContextTag == bool(false): // BACnetTimerStateChangeValueDouble
		if _child, err = new(_BACnetTimerStateChangeValueDouble).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetTimerStateChangeValueDouble for type-switch of BACnetTimerStateChangeValue")
		}
	case peekedTagNumber == 0x6 && peekedIsContextTag == bool(false): // BACnetTimerStateChangeValueOctetString
		if _child, err = new(_BACnetTimerStateChangeValueOctetString).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetTimerStateChangeValueOctetString for type-switch of BACnetTimerStateChangeValue")
		}
	case peekedTagNumber == 0x7 && peekedIsContextTag == bool(false): // BACnetTimerStateChangeValueCharacterString
		if _child, err = new(_BACnetTimerStateChangeValueCharacterString).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetTimerStateChangeValueCharacterString for type-switch of BACnetTimerStateChangeValue")
		}
	case peekedTagNumber == 0x8 && peekedIsContextTag == bool(false): // BACnetTimerStateChangeValueBitString
		if _child, err = new(_BACnetTimerStateChangeValueBitString).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetTimerStateChangeValueBitString for type-switch of BACnetTimerStateChangeValue")
		}
	case peekedTagNumber == 0x9 && peekedIsContextTag == bool(false): // BACnetTimerStateChangeValueEnumerated
		if _child, err = new(_BACnetTimerStateChangeValueEnumerated).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetTimerStateChangeValueEnumerated for type-switch of BACnetTimerStateChangeValue")
		}
	case peekedTagNumber == 0xA && peekedIsContextTag == bool(false): // BACnetTimerStateChangeValueDate
		if _child, err = new(_BACnetTimerStateChangeValueDate).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetTimerStateChangeValueDate for type-switch of BACnetTimerStateChangeValue")
		}
	case peekedTagNumber == 0xB && peekedIsContextTag == bool(false): // BACnetTimerStateChangeValueTime
		if _child, err = new(_BACnetTimerStateChangeValueTime).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetTimerStateChangeValueTime for type-switch of BACnetTimerStateChangeValue")
		}
	case peekedTagNumber == 0xC && peekedIsContextTag == bool(false): // BACnetTimerStateChangeValueObjectidentifier
		if _child, err = new(_BACnetTimerStateChangeValueObjectidentifier).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetTimerStateChangeValueObjectidentifier for type-switch of BACnetTimerStateChangeValue")
		}
	case peekedTagNumber == uint8(0) && peekedIsContextTag == bool(true): // BACnetTimerStateChangeValueNoValue
		if _child, err = new(_BACnetTimerStateChangeValueNoValue).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetTimerStateChangeValueNoValue for type-switch of BACnetTimerStateChangeValue")
		}
	case peekedTagNumber == uint8(1) && peekedIsContextTag == bool(true): // BACnetTimerStateChangeValueConstructedValue
		if _child, err = new(_BACnetTimerStateChangeValueConstructedValue).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetTimerStateChangeValueConstructedValue for type-switch of BACnetTimerStateChangeValue")
		}
	case peekedTagNumber == uint8(2) && peekedIsContextTag == bool(true): // BACnetTimerStateChangeValueDateTime
		if _child, err = new(_BACnetTimerStateChangeValueDateTime).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetTimerStateChangeValueDateTime for type-switch of BACnetTimerStateChangeValue")
		}
	case peekedTagNumber == uint8(3) && peekedIsContextTag == bool(true): // BACnetTimerStateChangeValueLightingCommand
		if _child, err = new(_BACnetTimerStateChangeValueLightingCommand).parse(ctx, readBuffer, m, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetTimerStateChangeValueLightingCommand for type-switch of BACnetTimerStateChangeValue")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v, peekedIsContextTag=%v]", peekedTagNumber, peekedIsContextTag)
	}

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValue")
	}

	return _child, nil
}

func (pm *_BACnetTimerStateChangeValue) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetTimerStateChangeValue, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValue")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}
	// Virtual field
	peekedIsContextTag := m.GetPeekedIsContextTag()
	_ = peekedIsContextTag
	if _peekedIsContextTagErr := writeBuffer.WriteVirtual(ctx, "peekedIsContextTag", m.GetPeekedIsContextTag()); _peekedIsContextTagErr != nil {
		return errors.Wrap(_peekedIsContextTagErr, "Error serializing 'peekedIsContextTag' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValue")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetTimerStateChangeValue) GetObjectTypeArgument() BACnetObjectType {
	return m.ObjectTypeArgument
}

//
////

func (m *_BACnetTimerStateChangeValue) IsBACnetTimerStateChangeValue() {}

func (m *_BACnetTimerStateChangeValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTimerStateChangeValue) deepCopy() *_BACnetTimerStateChangeValue {
	if m == nil {
		return nil
	}
	_BACnetTimerStateChangeValueCopy := &_BACnetTimerStateChangeValue{
		nil, // will be set by child
		utils.DeepCopy[BACnetTagHeader](m.PeekedTagHeader),
		m.ObjectTypeArgument,
	}
	return _BACnetTimerStateChangeValueCopy
}