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

// BACnetAccessCredentialDisableTagged is the corresponding interface of BACnetAccessCredentialDisableTagged
type BACnetAccessCredentialDisableTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetAccessCredentialDisable
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetAccessCredentialDisableTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAccessCredentialDisableTagged()
	// CreateBuilder creates a BACnetAccessCredentialDisableTaggedBuilder
	CreateBACnetAccessCredentialDisableTaggedBuilder() BACnetAccessCredentialDisableTaggedBuilder
}

// _BACnetAccessCredentialDisableTagged is the data-structure of this message
type _BACnetAccessCredentialDisableTagged struct {
	Header           BACnetTagHeader
	Value            BACnetAccessCredentialDisable
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetAccessCredentialDisableTagged = (*_BACnetAccessCredentialDisableTagged)(nil)

// NewBACnetAccessCredentialDisableTagged factory function for _BACnetAccessCredentialDisableTagged
func NewBACnetAccessCredentialDisableTagged(header BACnetTagHeader, value BACnetAccessCredentialDisable, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetAccessCredentialDisableTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetAccessCredentialDisableTagged must not be nil")
	}
	return &_BACnetAccessCredentialDisableTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetAccessCredentialDisableTaggedBuilder is a builder for BACnetAccessCredentialDisableTagged
type BACnetAccessCredentialDisableTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetAccessCredentialDisable, proprietaryValue uint32) BACnetAccessCredentialDisableTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetAccessCredentialDisableTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAccessCredentialDisableTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetAccessCredentialDisable) BACnetAccessCredentialDisableTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetAccessCredentialDisableTaggedBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetAccessCredentialDisableTaggedBuilder
	// WithArgTagClass sets a parser argument
	WithArgTagClass(TagClass) BACnetAccessCredentialDisableTaggedBuilder
	// Build builds the BACnetAccessCredentialDisableTagged or returns an error if something is wrong
	Build() (BACnetAccessCredentialDisableTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetAccessCredentialDisableTagged
}

// NewBACnetAccessCredentialDisableTaggedBuilder() creates a BACnetAccessCredentialDisableTaggedBuilder
func NewBACnetAccessCredentialDisableTaggedBuilder() BACnetAccessCredentialDisableTaggedBuilder {
	return &_BACnetAccessCredentialDisableTaggedBuilder{_BACnetAccessCredentialDisableTagged: new(_BACnetAccessCredentialDisableTagged)}
}

type _BACnetAccessCredentialDisableTaggedBuilder struct {
	*_BACnetAccessCredentialDisableTagged

	err *utils.MultiError
}

var _ (BACnetAccessCredentialDisableTaggedBuilder) = (*_BACnetAccessCredentialDisableTaggedBuilder)(nil)

func (b *_BACnetAccessCredentialDisableTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetAccessCredentialDisable, proprietaryValue uint32) BACnetAccessCredentialDisableTaggedBuilder {
	return b.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (b *_BACnetAccessCredentialDisableTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetAccessCredentialDisableTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetAccessCredentialDisableTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAccessCredentialDisableTaggedBuilder {
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

func (b *_BACnetAccessCredentialDisableTaggedBuilder) WithValue(value BACnetAccessCredentialDisable) BACnetAccessCredentialDisableTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetAccessCredentialDisableTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetAccessCredentialDisableTaggedBuilder {
	b.ProprietaryValue = proprietaryValue
	return b
}

func (b *_BACnetAccessCredentialDisableTaggedBuilder) WithArgTagNumber(tagNumber uint8) BACnetAccessCredentialDisableTaggedBuilder {
	b.TagNumber = tagNumber
	return b
}
func (b *_BACnetAccessCredentialDisableTaggedBuilder) WithArgTagClass(tagClass TagClass) BACnetAccessCredentialDisableTaggedBuilder {
	b.TagClass = tagClass
	return b
}

func (b *_BACnetAccessCredentialDisableTaggedBuilder) Build() (BACnetAccessCredentialDisableTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetAccessCredentialDisableTagged.deepCopy(), nil
}

func (b *_BACnetAccessCredentialDisableTaggedBuilder) MustBuild() BACnetAccessCredentialDisableTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetAccessCredentialDisableTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetAccessCredentialDisableTaggedBuilder().(*_BACnetAccessCredentialDisableTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetAccessCredentialDisableTaggedBuilder creates a BACnetAccessCredentialDisableTaggedBuilder
func (b *_BACnetAccessCredentialDisableTagged) CreateBACnetAccessCredentialDisableTaggedBuilder() BACnetAccessCredentialDisableTaggedBuilder {
	if b == nil {
		return NewBACnetAccessCredentialDisableTaggedBuilder()
	}
	return &_BACnetAccessCredentialDisableTaggedBuilder{_BACnetAccessCredentialDisableTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAccessCredentialDisableTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetAccessCredentialDisableTagged) GetValue() BACnetAccessCredentialDisable {
	return m.Value
}

func (m *_BACnetAccessCredentialDisableTagged) GetProprietaryValue() uint32 {
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

func (m *_BACnetAccessCredentialDisableTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetAccessCredentialDisable_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetAccessCredentialDisableTagged(structType any) BACnetAccessCredentialDisableTagged {
	if casted, ok := structType.(BACnetAccessCredentialDisableTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAccessCredentialDisableTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAccessCredentialDisableTagged) GetTypeName() string {
	return "BACnetAccessCredentialDisableTagged"
}

func (m *_BACnetAccessCredentialDisableTagged) GetLengthInBits(ctx context.Context) uint16 {
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

func (m *_BACnetAccessCredentialDisableTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccessCredentialDisableTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetAccessCredentialDisableTagged, error) {
	return BACnetAccessCredentialDisableTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetAccessCredentialDisableTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessCredentialDisableTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessCredentialDisableTagged, error) {
		return BACnetAccessCredentialDisableTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetAccessCredentialDisableTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetAccessCredentialDisableTagged, error) {
	v, err := (&_BACnetAccessCredentialDisableTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAccessCredentialDisableTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetAccessCredentialDisableTagged BACnetAccessCredentialDisableTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAccessCredentialDisableTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAccessCredentialDisableTagged")
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

	value, err := ReadManualField[BACnetAccessCredentialDisable](ctx, "value", readBuffer, EnsureType[BACnetAccessCredentialDisable](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetAccessCredentialDisable_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetAccessCredentialDisable_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetAccessCredentialDisableTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAccessCredentialDisableTagged")
	}

	return m, nil
}

func (m *_BACnetAccessCredentialDisableTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAccessCredentialDisableTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAccessCredentialDisableTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAccessCredentialDisableTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetAccessCredentialDisable](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
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

	if popErr := writeBuffer.PopContext("BACnetAccessCredentialDisableTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAccessCredentialDisableTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAccessCredentialDisableTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetAccessCredentialDisableTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetAccessCredentialDisableTagged) IsBACnetAccessCredentialDisableTagged() {}

func (m *_BACnetAccessCredentialDisableTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetAccessCredentialDisableTagged) deepCopy() *_BACnetAccessCredentialDisableTagged {
	if m == nil {
		return nil
	}
	_BACnetAccessCredentialDisableTaggedCopy := &_BACnetAccessCredentialDisableTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetAccessCredentialDisableTaggedCopy
}

func (m *_BACnetAccessCredentialDisableTagged) String() string {
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
