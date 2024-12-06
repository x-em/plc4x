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

// BACnetAccessEventTagged is the corresponding interface of BACnetAccessEventTagged
type BACnetAccessEventTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetAccessEvent
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetAccessEventTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAccessEventTagged()
	// CreateBuilder creates a BACnetAccessEventTaggedBuilder
	CreateBACnetAccessEventTaggedBuilder() BACnetAccessEventTaggedBuilder
}

// _BACnetAccessEventTagged is the data-structure of this message
type _BACnetAccessEventTagged struct {
	Header           BACnetTagHeader
	Value            BACnetAccessEvent
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetAccessEventTagged = (*_BACnetAccessEventTagged)(nil)

// NewBACnetAccessEventTagged factory function for _BACnetAccessEventTagged
func NewBACnetAccessEventTagged(header BACnetTagHeader, value BACnetAccessEvent, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetAccessEventTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetAccessEventTagged must not be nil")
	}
	return &_BACnetAccessEventTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetAccessEventTaggedBuilder is a builder for BACnetAccessEventTagged
type BACnetAccessEventTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetAccessEvent, proprietaryValue uint32) BACnetAccessEventTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetAccessEventTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAccessEventTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetAccessEvent) BACnetAccessEventTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetAccessEventTaggedBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetAccessEventTaggedBuilder
	// WithArgTagClass sets a parser argument
	WithArgTagClass(TagClass) BACnetAccessEventTaggedBuilder
	// Build builds the BACnetAccessEventTagged or returns an error if something is wrong
	Build() (BACnetAccessEventTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetAccessEventTagged
}

// NewBACnetAccessEventTaggedBuilder() creates a BACnetAccessEventTaggedBuilder
func NewBACnetAccessEventTaggedBuilder() BACnetAccessEventTaggedBuilder {
	return &_BACnetAccessEventTaggedBuilder{_BACnetAccessEventTagged: new(_BACnetAccessEventTagged)}
}

type _BACnetAccessEventTaggedBuilder struct {
	*_BACnetAccessEventTagged

	err *utils.MultiError
}

var _ (BACnetAccessEventTaggedBuilder) = (*_BACnetAccessEventTaggedBuilder)(nil)

func (b *_BACnetAccessEventTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetAccessEvent, proprietaryValue uint32) BACnetAccessEventTaggedBuilder {
	return b.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (b *_BACnetAccessEventTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetAccessEventTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetAccessEventTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAccessEventTaggedBuilder {
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

func (b *_BACnetAccessEventTaggedBuilder) WithValue(value BACnetAccessEvent) BACnetAccessEventTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetAccessEventTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetAccessEventTaggedBuilder {
	b.ProprietaryValue = proprietaryValue
	return b
}

func (b *_BACnetAccessEventTaggedBuilder) WithArgTagNumber(tagNumber uint8) BACnetAccessEventTaggedBuilder {
	b.TagNumber = tagNumber
	return b
}
func (b *_BACnetAccessEventTaggedBuilder) WithArgTagClass(tagClass TagClass) BACnetAccessEventTaggedBuilder {
	b.TagClass = tagClass
	return b
}

func (b *_BACnetAccessEventTaggedBuilder) Build() (BACnetAccessEventTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetAccessEventTagged.deepCopy(), nil
}

func (b *_BACnetAccessEventTaggedBuilder) MustBuild() BACnetAccessEventTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetAccessEventTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetAccessEventTaggedBuilder().(*_BACnetAccessEventTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetAccessEventTaggedBuilder creates a BACnetAccessEventTaggedBuilder
func (b *_BACnetAccessEventTagged) CreateBACnetAccessEventTaggedBuilder() BACnetAccessEventTaggedBuilder {
	if b == nil {
		return NewBACnetAccessEventTaggedBuilder()
	}
	return &_BACnetAccessEventTaggedBuilder{_BACnetAccessEventTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAccessEventTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetAccessEventTagged) GetValue() BACnetAccessEvent {
	return m.Value
}

func (m *_BACnetAccessEventTagged) GetProprietaryValue() uint32 {
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

func (m *_BACnetAccessEventTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetAccessEvent_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetAccessEventTagged(structType any) BACnetAccessEventTagged {
	if casted, ok := structType.(BACnetAccessEventTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAccessEventTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAccessEventTagged) GetTypeName() string {
	return "BACnetAccessEventTagged"
}

func (m *_BACnetAccessEventTagged) GetLengthInBits(ctx context.Context) uint16 {
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

func (m *_BACnetAccessEventTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccessEventTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetAccessEventTagged, error) {
	return BACnetAccessEventTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetAccessEventTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessEventTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessEventTagged, error) {
		return BACnetAccessEventTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetAccessEventTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetAccessEventTagged, error) {
	v, err := (&_BACnetAccessEventTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAccessEventTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetAccessEventTagged BACnetAccessEventTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAccessEventTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAccessEventTagged")
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

	value, err := ReadManualField[BACnetAccessEvent](ctx, "value", readBuffer, EnsureType[BACnetAccessEvent](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetAccessEvent_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetAccessEvent_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetAccessEventTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAccessEventTagged")
	}

	return m, nil
}

func (m *_BACnetAccessEventTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAccessEventTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAccessEventTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAccessEventTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetAccessEvent](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
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

	if popErr := writeBuffer.PopContext("BACnetAccessEventTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAccessEventTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAccessEventTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetAccessEventTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetAccessEventTagged) IsBACnetAccessEventTagged() {}

func (m *_BACnetAccessEventTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetAccessEventTagged) deepCopy() *_BACnetAccessEventTagged {
	if m == nil {
		return nil
	}
	_BACnetAccessEventTaggedCopy := &_BACnetAccessEventTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetAccessEventTaggedCopy
}

func (m *_BACnetAccessEventTagged) String() string {
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