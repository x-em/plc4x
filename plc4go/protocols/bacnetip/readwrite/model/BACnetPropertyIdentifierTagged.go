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

// BACnetPropertyIdentifierTagged is the corresponding interface of BACnetPropertyIdentifierTagged
type BACnetPropertyIdentifierTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetPropertyIdentifier
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetPropertyIdentifierTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyIdentifierTagged()
	// CreateBuilder creates a BACnetPropertyIdentifierTaggedBuilder
	CreateBACnetPropertyIdentifierTaggedBuilder() BACnetPropertyIdentifierTaggedBuilder
}

// _BACnetPropertyIdentifierTagged is the data-structure of this message
type _BACnetPropertyIdentifierTagged struct {
	Header           BACnetTagHeader
	Value            BACnetPropertyIdentifier
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetPropertyIdentifierTagged = (*_BACnetPropertyIdentifierTagged)(nil)

// NewBACnetPropertyIdentifierTagged factory function for _BACnetPropertyIdentifierTagged
func NewBACnetPropertyIdentifierTagged(header BACnetTagHeader, value BACnetPropertyIdentifier, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetPropertyIdentifierTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetPropertyIdentifierTagged must not be nil")
	}
	return &_BACnetPropertyIdentifierTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyIdentifierTaggedBuilder is a builder for BACnetPropertyIdentifierTagged
type BACnetPropertyIdentifierTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetPropertyIdentifier, proprietaryValue uint32) BACnetPropertyIdentifierTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetPropertyIdentifierTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetPropertyIdentifierTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetPropertyIdentifier) BACnetPropertyIdentifierTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetPropertyIdentifierTaggedBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetPropertyIdentifierTaggedBuilder
	// WithArgTagClass sets a parser argument
	WithArgTagClass(TagClass) BACnetPropertyIdentifierTaggedBuilder
	// Build builds the BACnetPropertyIdentifierTagged or returns an error if something is wrong
	Build() (BACnetPropertyIdentifierTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyIdentifierTagged
}

// NewBACnetPropertyIdentifierTaggedBuilder() creates a BACnetPropertyIdentifierTaggedBuilder
func NewBACnetPropertyIdentifierTaggedBuilder() BACnetPropertyIdentifierTaggedBuilder {
	return &_BACnetPropertyIdentifierTaggedBuilder{_BACnetPropertyIdentifierTagged: new(_BACnetPropertyIdentifierTagged)}
}

type _BACnetPropertyIdentifierTaggedBuilder struct {
	*_BACnetPropertyIdentifierTagged

	err *utils.MultiError
}

var _ (BACnetPropertyIdentifierTaggedBuilder) = (*_BACnetPropertyIdentifierTaggedBuilder)(nil)

func (b *_BACnetPropertyIdentifierTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetPropertyIdentifier, proprietaryValue uint32) BACnetPropertyIdentifierTaggedBuilder {
	return b.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (b *_BACnetPropertyIdentifierTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetPropertyIdentifierTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetPropertyIdentifierTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetPropertyIdentifierTaggedBuilder {
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

func (b *_BACnetPropertyIdentifierTaggedBuilder) WithValue(value BACnetPropertyIdentifier) BACnetPropertyIdentifierTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetPropertyIdentifierTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetPropertyIdentifierTaggedBuilder {
	b.ProprietaryValue = proprietaryValue
	return b
}

func (b *_BACnetPropertyIdentifierTaggedBuilder) WithArgTagNumber(tagNumber uint8) BACnetPropertyIdentifierTaggedBuilder {
	b.TagNumber = tagNumber
	return b
}
func (b *_BACnetPropertyIdentifierTaggedBuilder) WithArgTagClass(tagClass TagClass) BACnetPropertyIdentifierTaggedBuilder {
	b.TagClass = tagClass
	return b
}

func (b *_BACnetPropertyIdentifierTaggedBuilder) Build() (BACnetPropertyIdentifierTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyIdentifierTagged.deepCopy(), nil
}

func (b *_BACnetPropertyIdentifierTaggedBuilder) MustBuild() BACnetPropertyIdentifierTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPropertyIdentifierTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyIdentifierTaggedBuilder().(*_BACnetPropertyIdentifierTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyIdentifierTaggedBuilder creates a BACnetPropertyIdentifierTaggedBuilder
func (b *_BACnetPropertyIdentifierTagged) CreateBACnetPropertyIdentifierTaggedBuilder() BACnetPropertyIdentifierTaggedBuilder {
	if b == nil {
		return NewBACnetPropertyIdentifierTaggedBuilder()
	}
	return &_BACnetPropertyIdentifierTaggedBuilder{_BACnetPropertyIdentifierTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyIdentifierTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetPropertyIdentifierTagged) GetValue() BACnetPropertyIdentifier {
	return m.Value
}

func (m *_BACnetPropertyIdentifierTagged) GetProprietaryValue() uint32 {
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

func (m *_BACnetPropertyIdentifierTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyIdentifierTagged(structType any) BACnetPropertyIdentifierTagged {
	if casted, ok := structType.(BACnetPropertyIdentifierTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyIdentifierTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyIdentifierTagged) GetTypeName() string {
	return "BACnetPropertyIdentifierTagged"
}

func (m *_BACnetPropertyIdentifierTagged) GetLengthInBits(ctx context.Context) uint16 {
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

func (m *_BACnetPropertyIdentifierTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyIdentifierTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetPropertyIdentifierTagged, error) {
	return BACnetPropertyIdentifierTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetPropertyIdentifierTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyIdentifierTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyIdentifierTagged, error) {
		return BACnetPropertyIdentifierTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetPropertyIdentifierTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetPropertyIdentifierTagged, error) {
	v, err := (&_BACnetPropertyIdentifierTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetPropertyIdentifierTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetPropertyIdentifierTagged BACnetPropertyIdentifierTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyIdentifierTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyIdentifierTagged")
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

	value, err := ReadManualField[BACnetPropertyIdentifier](ctx, "value", readBuffer, EnsureType[BACnetPropertyIdentifier](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetPropertyIdentifierTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyIdentifierTagged")
	}

	return m, nil
}

func (m *_BACnetPropertyIdentifierTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyIdentifierTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetPropertyIdentifierTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPropertyIdentifierTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetPropertyIdentifier](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
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

	if popErr := writeBuffer.PopContext("BACnetPropertyIdentifierTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPropertyIdentifierTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetPropertyIdentifierTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetPropertyIdentifierTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetPropertyIdentifierTagged) IsBACnetPropertyIdentifierTagged() {}

func (m *_BACnetPropertyIdentifierTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyIdentifierTagged) deepCopy() *_BACnetPropertyIdentifierTagged {
	if m == nil {
		return nil
	}
	_BACnetPropertyIdentifierTaggedCopy := &_BACnetPropertyIdentifierTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetPropertyIdentifierTaggedCopy
}

func (m *_BACnetPropertyIdentifierTagged) String() string {
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
