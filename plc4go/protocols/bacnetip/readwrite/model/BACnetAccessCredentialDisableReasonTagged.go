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

// BACnetAccessCredentialDisableReasonTagged is the corresponding interface of BACnetAccessCredentialDisableReasonTagged
type BACnetAccessCredentialDisableReasonTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetAccessCredentialDisableReason
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetAccessCredentialDisableReasonTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAccessCredentialDisableReasonTagged()
	// CreateBuilder creates a BACnetAccessCredentialDisableReasonTaggedBuilder
	CreateBACnetAccessCredentialDisableReasonTaggedBuilder() BACnetAccessCredentialDisableReasonTaggedBuilder
}

// _BACnetAccessCredentialDisableReasonTagged is the data-structure of this message
type _BACnetAccessCredentialDisableReasonTagged struct {
	Header           BACnetTagHeader
	Value            BACnetAccessCredentialDisableReason
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetAccessCredentialDisableReasonTagged = (*_BACnetAccessCredentialDisableReasonTagged)(nil)

// NewBACnetAccessCredentialDisableReasonTagged factory function for _BACnetAccessCredentialDisableReasonTagged
func NewBACnetAccessCredentialDisableReasonTagged(header BACnetTagHeader, value BACnetAccessCredentialDisableReason, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetAccessCredentialDisableReasonTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetAccessCredentialDisableReasonTagged must not be nil")
	}
	return &_BACnetAccessCredentialDisableReasonTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetAccessCredentialDisableReasonTaggedBuilder is a builder for BACnetAccessCredentialDisableReasonTagged
type BACnetAccessCredentialDisableReasonTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetAccessCredentialDisableReason, proprietaryValue uint32) BACnetAccessCredentialDisableReasonTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetAccessCredentialDisableReasonTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAccessCredentialDisableReasonTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetAccessCredentialDisableReason) BACnetAccessCredentialDisableReasonTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetAccessCredentialDisableReasonTaggedBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetAccessCredentialDisableReasonTaggedBuilder
	// WithArgTagClass sets a parser argument
	WithArgTagClass(TagClass) BACnetAccessCredentialDisableReasonTaggedBuilder
	// Build builds the BACnetAccessCredentialDisableReasonTagged or returns an error if something is wrong
	Build() (BACnetAccessCredentialDisableReasonTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetAccessCredentialDisableReasonTagged
}

// NewBACnetAccessCredentialDisableReasonTaggedBuilder() creates a BACnetAccessCredentialDisableReasonTaggedBuilder
func NewBACnetAccessCredentialDisableReasonTaggedBuilder() BACnetAccessCredentialDisableReasonTaggedBuilder {
	return &_BACnetAccessCredentialDisableReasonTaggedBuilder{_BACnetAccessCredentialDisableReasonTagged: new(_BACnetAccessCredentialDisableReasonTagged)}
}

type _BACnetAccessCredentialDisableReasonTaggedBuilder struct {
	*_BACnetAccessCredentialDisableReasonTagged

	err *utils.MultiError
}

var _ (BACnetAccessCredentialDisableReasonTaggedBuilder) = (*_BACnetAccessCredentialDisableReasonTaggedBuilder)(nil)

func (b *_BACnetAccessCredentialDisableReasonTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetAccessCredentialDisableReason, proprietaryValue uint32) BACnetAccessCredentialDisableReasonTaggedBuilder {
	return b.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (b *_BACnetAccessCredentialDisableReasonTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetAccessCredentialDisableReasonTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetAccessCredentialDisableReasonTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAccessCredentialDisableReasonTaggedBuilder {
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

func (b *_BACnetAccessCredentialDisableReasonTaggedBuilder) WithValue(value BACnetAccessCredentialDisableReason) BACnetAccessCredentialDisableReasonTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetAccessCredentialDisableReasonTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetAccessCredentialDisableReasonTaggedBuilder {
	b.ProprietaryValue = proprietaryValue
	return b
}

func (b *_BACnetAccessCredentialDisableReasonTaggedBuilder) WithArgTagNumber(tagNumber uint8) BACnetAccessCredentialDisableReasonTaggedBuilder {
	b.TagNumber = tagNumber
	return b
}
func (b *_BACnetAccessCredentialDisableReasonTaggedBuilder) WithArgTagClass(tagClass TagClass) BACnetAccessCredentialDisableReasonTaggedBuilder {
	b.TagClass = tagClass
	return b
}

func (b *_BACnetAccessCredentialDisableReasonTaggedBuilder) Build() (BACnetAccessCredentialDisableReasonTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetAccessCredentialDisableReasonTagged.deepCopy(), nil
}

func (b *_BACnetAccessCredentialDisableReasonTaggedBuilder) MustBuild() BACnetAccessCredentialDisableReasonTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetAccessCredentialDisableReasonTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetAccessCredentialDisableReasonTaggedBuilder().(*_BACnetAccessCredentialDisableReasonTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetAccessCredentialDisableReasonTaggedBuilder creates a BACnetAccessCredentialDisableReasonTaggedBuilder
func (b *_BACnetAccessCredentialDisableReasonTagged) CreateBACnetAccessCredentialDisableReasonTaggedBuilder() BACnetAccessCredentialDisableReasonTaggedBuilder {
	if b == nil {
		return NewBACnetAccessCredentialDisableReasonTaggedBuilder()
	}
	return &_BACnetAccessCredentialDisableReasonTaggedBuilder{_BACnetAccessCredentialDisableReasonTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAccessCredentialDisableReasonTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetAccessCredentialDisableReasonTagged) GetValue() BACnetAccessCredentialDisableReason {
	return m.Value
}

func (m *_BACnetAccessCredentialDisableReasonTagged) GetProprietaryValue() uint32 {
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

func (m *_BACnetAccessCredentialDisableReasonTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetAccessCredentialDisableReason_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetAccessCredentialDisableReasonTagged(structType any) BACnetAccessCredentialDisableReasonTagged {
	if casted, ok := structType.(BACnetAccessCredentialDisableReasonTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAccessCredentialDisableReasonTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAccessCredentialDisableReasonTagged) GetTypeName() string {
	return "BACnetAccessCredentialDisableReasonTagged"
}

func (m *_BACnetAccessCredentialDisableReasonTagged) GetLengthInBits(ctx context.Context) uint16 {
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

func (m *_BACnetAccessCredentialDisableReasonTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccessCredentialDisableReasonTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetAccessCredentialDisableReasonTagged, error) {
	return BACnetAccessCredentialDisableReasonTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetAccessCredentialDisableReasonTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessCredentialDisableReasonTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessCredentialDisableReasonTagged, error) {
		return BACnetAccessCredentialDisableReasonTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetAccessCredentialDisableReasonTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetAccessCredentialDisableReasonTagged, error) {
	v, err := (&_BACnetAccessCredentialDisableReasonTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAccessCredentialDisableReasonTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetAccessCredentialDisableReasonTagged BACnetAccessCredentialDisableReasonTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAccessCredentialDisableReasonTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAccessCredentialDisableReasonTagged")
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

	value, err := ReadManualField[BACnetAccessCredentialDisableReason](ctx, "value", readBuffer, EnsureType[BACnetAccessCredentialDisableReason](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetAccessCredentialDisableReason_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetAccessCredentialDisableReason_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetAccessCredentialDisableReasonTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAccessCredentialDisableReasonTagged")
	}

	return m, nil
}

func (m *_BACnetAccessCredentialDisableReasonTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAccessCredentialDisableReasonTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAccessCredentialDisableReasonTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAccessCredentialDisableReasonTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetAccessCredentialDisableReason](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
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

	if popErr := writeBuffer.PopContext("BACnetAccessCredentialDisableReasonTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAccessCredentialDisableReasonTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAccessCredentialDisableReasonTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetAccessCredentialDisableReasonTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetAccessCredentialDisableReasonTagged) IsBACnetAccessCredentialDisableReasonTagged() {}

func (m *_BACnetAccessCredentialDisableReasonTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetAccessCredentialDisableReasonTagged) deepCopy() *_BACnetAccessCredentialDisableReasonTagged {
	if m == nil {
		return nil
	}
	_BACnetAccessCredentialDisableReasonTaggedCopy := &_BACnetAccessCredentialDisableReasonTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetAccessCredentialDisableReasonTaggedCopy
}

func (m *_BACnetAccessCredentialDisableReasonTagged) String() string {
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