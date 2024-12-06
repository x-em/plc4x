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

// BACnetDoorStatusTagged is the corresponding interface of BACnetDoorStatusTagged
type BACnetDoorStatusTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetDoorStatus
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetDoorStatusTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetDoorStatusTagged()
	// CreateBuilder creates a BACnetDoorStatusTaggedBuilder
	CreateBACnetDoorStatusTaggedBuilder() BACnetDoorStatusTaggedBuilder
}

// _BACnetDoorStatusTagged is the data-structure of this message
type _BACnetDoorStatusTagged struct {
	Header           BACnetTagHeader
	Value            BACnetDoorStatus
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetDoorStatusTagged = (*_BACnetDoorStatusTagged)(nil)

// NewBACnetDoorStatusTagged factory function for _BACnetDoorStatusTagged
func NewBACnetDoorStatusTagged(header BACnetTagHeader, value BACnetDoorStatus, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetDoorStatusTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetDoorStatusTagged must not be nil")
	}
	return &_BACnetDoorStatusTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetDoorStatusTaggedBuilder is a builder for BACnetDoorStatusTagged
type BACnetDoorStatusTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetDoorStatus, proprietaryValue uint32) BACnetDoorStatusTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetDoorStatusTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetDoorStatusTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetDoorStatus) BACnetDoorStatusTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetDoorStatusTaggedBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetDoorStatusTaggedBuilder
	// WithArgTagClass sets a parser argument
	WithArgTagClass(TagClass) BACnetDoorStatusTaggedBuilder
	// Build builds the BACnetDoorStatusTagged or returns an error if something is wrong
	Build() (BACnetDoorStatusTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetDoorStatusTagged
}

// NewBACnetDoorStatusTaggedBuilder() creates a BACnetDoorStatusTaggedBuilder
func NewBACnetDoorStatusTaggedBuilder() BACnetDoorStatusTaggedBuilder {
	return &_BACnetDoorStatusTaggedBuilder{_BACnetDoorStatusTagged: new(_BACnetDoorStatusTagged)}
}

type _BACnetDoorStatusTaggedBuilder struct {
	*_BACnetDoorStatusTagged

	err *utils.MultiError
}

var _ (BACnetDoorStatusTaggedBuilder) = (*_BACnetDoorStatusTaggedBuilder)(nil)

func (b *_BACnetDoorStatusTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetDoorStatus, proprietaryValue uint32) BACnetDoorStatusTaggedBuilder {
	return b.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (b *_BACnetDoorStatusTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetDoorStatusTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetDoorStatusTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetDoorStatusTaggedBuilder {
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

func (b *_BACnetDoorStatusTaggedBuilder) WithValue(value BACnetDoorStatus) BACnetDoorStatusTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetDoorStatusTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetDoorStatusTaggedBuilder {
	b.ProprietaryValue = proprietaryValue
	return b
}

func (b *_BACnetDoorStatusTaggedBuilder) WithArgTagNumber(tagNumber uint8) BACnetDoorStatusTaggedBuilder {
	b.TagNumber = tagNumber
	return b
}
func (b *_BACnetDoorStatusTaggedBuilder) WithArgTagClass(tagClass TagClass) BACnetDoorStatusTaggedBuilder {
	b.TagClass = tagClass
	return b
}

func (b *_BACnetDoorStatusTaggedBuilder) Build() (BACnetDoorStatusTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetDoorStatusTagged.deepCopy(), nil
}

func (b *_BACnetDoorStatusTaggedBuilder) MustBuild() BACnetDoorStatusTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetDoorStatusTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetDoorStatusTaggedBuilder().(*_BACnetDoorStatusTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetDoorStatusTaggedBuilder creates a BACnetDoorStatusTaggedBuilder
func (b *_BACnetDoorStatusTagged) CreateBACnetDoorStatusTaggedBuilder() BACnetDoorStatusTaggedBuilder {
	if b == nil {
		return NewBACnetDoorStatusTaggedBuilder()
	}
	return &_BACnetDoorStatusTaggedBuilder{_BACnetDoorStatusTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDoorStatusTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetDoorStatusTagged) GetValue() BACnetDoorStatus {
	return m.Value
}

func (m *_BACnetDoorStatusTagged) GetProprietaryValue() uint32 {
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

func (m *_BACnetDoorStatusTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetDoorStatus_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetDoorStatusTagged(structType any) BACnetDoorStatusTagged {
	if casted, ok := structType.(BACnetDoorStatusTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDoorStatusTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDoorStatusTagged) GetTypeName() string {
	return "BACnetDoorStatusTagged"
}

func (m *_BACnetDoorStatusTagged) GetLengthInBits(ctx context.Context) uint16 {
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

func (m *_BACnetDoorStatusTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDoorStatusTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetDoorStatusTagged, error) {
	return BACnetDoorStatusTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetDoorStatusTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDoorStatusTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDoorStatusTagged, error) {
		return BACnetDoorStatusTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetDoorStatusTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetDoorStatusTagged, error) {
	v, err := (&_BACnetDoorStatusTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetDoorStatusTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetDoorStatusTagged BACnetDoorStatusTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDoorStatusTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDoorStatusTagged")
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

	value, err := ReadManualField[BACnetDoorStatus](ctx, "value", readBuffer, EnsureType[BACnetDoorStatus](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetDoorStatus_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetDoorStatus_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetDoorStatusTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDoorStatusTagged")
	}

	return m, nil
}

func (m *_BACnetDoorStatusTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDoorStatusTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetDoorStatusTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDoorStatusTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetDoorStatus](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
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

	if popErr := writeBuffer.PopContext("BACnetDoorStatusTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDoorStatusTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetDoorStatusTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetDoorStatusTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetDoorStatusTagged) IsBACnetDoorStatusTagged() {}

func (m *_BACnetDoorStatusTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetDoorStatusTagged) deepCopy() *_BACnetDoorStatusTagged {
	if m == nil {
		return nil
	}
	_BACnetDoorStatusTaggedCopy := &_BACnetDoorStatusTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetDoorStatusTaggedCopy
}

func (m *_BACnetDoorStatusTagged) String() string {
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