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

// BACnetLightingOperationTagged is the corresponding interface of BACnetLightingOperationTagged
type BACnetLightingOperationTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetLightingOperation
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetLightingOperationTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLightingOperationTagged()
	// CreateBuilder creates a BACnetLightingOperationTaggedBuilder
	CreateBACnetLightingOperationTaggedBuilder() BACnetLightingOperationTaggedBuilder
}

// _BACnetLightingOperationTagged is the data-structure of this message
type _BACnetLightingOperationTagged struct {
	Header           BACnetTagHeader
	Value            BACnetLightingOperation
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetLightingOperationTagged = (*_BACnetLightingOperationTagged)(nil)

// NewBACnetLightingOperationTagged factory function for _BACnetLightingOperationTagged
func NewBACnetLightingOperationTagged(header BACnetTagHeader, value BACnetLightingOperation, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetLightingOperationTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetLightingOperationTagged must not be nil")
	}
	return &_BACnetLightingOperationTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLightingOperationTaggedBuilder is a builder for BACnetLightingOperationTagged
type BACnetLightingOperationTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetLightingOperation, proprietaryValue uint32) BACnetLightingOperationTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetLightingOperationTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetLightingOperationTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetLightingOperation) BACnetLightingOperationTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetLightingOperationTaggedBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetLightingOperationTaggedBuilder
	// WithArgTagClass sets a parser argument
	WithArgTagClass(TagClass) BACnetLightingOperationTaggedBuilder
	// Build builds the BACnetLightingOperationTagged or returns an error if something is wrong
	Build() (BACnetLightingOperationTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLightingOperationTagged
}

// NewBACnetLightingOperationTaggedBuilder() creates a BACnetLightingOperationTaggedBuilder
func NewBACnetLightingOperationTaggedBuilder() BACnetLightingOperationTaggedBuilder {
	return &_BACnetLightingOperationTaggedBuilder{_BACnetLightingOperationTagged: new(_BACnetLightingOperationTagged)}
}

type _BACnetLightingOperationTaggedBuilder struct {
	*_BACnetLightingOperationTagged

	err *utils.MultiError
}

var _ (BACnetLightingOperationTaggedBuilder) = (*_BACnetLightingOperationTaggedBuilder)(nil)

func (b *_BACnetLightingOperationTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetLightingOperation, proprietaryValue uint32) BACnetLightingOperationTaggedBuilder {
	return b.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (b *_BACnetLightingOperationTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetLightingOperationTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetLightingOperationTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetLightingOperationTaggedBuilder {
	builder := builderSupplier(b.Header.CreateBACnetTagHeaderBuilder())
	var err error
	b.Header, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return b
}

func (b *_BACnetLightingOperationTaggedBuilder) WithValue(value BACnetLightingOperation) BACnetLightingOperationTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetLightingOperationTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetLightingOperationTaggedBuilder {
	b.ProprietaryValue = proprietaryValue
	return b
}

func (b *_BACnetLightingOperationTaggedBuilder) WithArgTagNumber(tagNumber uint8) BACnetLightingOperationTaggedBuilder {
	b.TagNumber = tagNumber
	return b
}
func (b *_BACnetLightingOperationTaggedBuilder) WithArgTagClass(tagClass TagClass) BACnetLightingOperationTaggedBuilder {
	b.TagClass = tagClass
	return b
}

func (b *_BACnetLightingOperationTaggedBuilder) Build() (BACnetLightingOperationTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetLightingOperationTagged.deepCopy(), nil
}

func (b *_BACnetLightingOperationTaggedBuilder) MustBuild() BACnetLightingOperationTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetLightingOperationTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetLightingOperationTaggedBuilder().(*_BACnetLightingOperationTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetLightingOperationTaggedBuilder creates a BACnetLightingOperationTaggedBuilder
func (b *_BACnetLightingOperationTagged) CreateBACnetLightingOperationTaggedBuilder() BACnetLightingOperationTaggedBuilder {
	if b == nil {
		return NewBACnetLightingOperationTaggedBuilder()
	}
	return &_BACnetLightingOperationTaggedBuilder{_BACnetLightingOperationTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLightingOperationTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetLightingOperationTagged) GetValue() BACnetLightingOperation {
	return m.Value
}

func (m *_BACnetLightingOperationTagged) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetLightingOperationTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetLightingOperation_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLightingOperationTagged(structType any) BACnetLightingOperationTagged {
	if casted, ok := structType.(BACnetLightingOperationTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLightingOperationTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLightingOperationTagged) GetTypeName() string {
	return "BACnetLightingOperationTagged"
}

func (m *_BACnetLightingOperationTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32(int32(0)) }, func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }, func() any { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_BACnetLightingOperationTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLightingOperationTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetLightingOperationTagged, error) {
	return BACnetLightingOperationTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetLightingOperationTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLightingOperationTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLightingOperationTagged, error) {
		return BACnetLightingOperationTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetLightingOperationTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetLightingOperationTagged, error) {
	v, err := (&_BACnetLightingOperationTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetLightingOperationTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetLightingOperationTagged BACnetLightingOperationTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLightingOperationTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLightingOperationTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	value, err := ReadManualField[BACnetLightingOperation](ctx, "value", readBuffer, EnsureType[BACnetLightingOperation](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetLightingOperation_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetLightingOperation_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetLightingOperationTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLightingOperationTagged")
	}

	return m, nil
}

func (m *_BACnetLightingOperationTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLightingOperationTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLightingOperationTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLightingOperationTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetLightingOperation](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}
	// Virtual field
	isProprietary := m.GetIsProprietary()
	_ = isProprietary
	if _isProprietaryErr := writeBuffer.WriteVirtual(ctx, "isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	if err := WriteManualField[uint32](ctx, "proprietaryValue", func(ctx context.Context) error {
		return WriteProprietaryEnumGeneric(ctx, writeBuffer, m.GetProprietaryValue(), m.GetIsProprietary())
	}, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'proprietaryValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLightingOperationTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLightingOperationTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetLightingOperationTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetLightingOperationTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetLightingOperationTagged) IsBACnetLightingOperationTagged() {}

func (m *_BACnetLightingOperationTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLightingOperationTagged) deepCopy() *_BACnetLightingOperationTagged {
	if m == nil {
		return nil
	}
	_BACnetLightingOperationTaggedCopy := &_BACnetLightingOperationTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetLightingOperationTaggedCopy
}

func (m *_BACnetLightingOperationTagged) String() string {
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