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

// BACnetRelationshipTagged is the corresponding interface of BACnetRelationshipTagged
type BACnetRelationshipTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetRelationship
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetRelationshipTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetRelationshipTagged()
	// CreateBuilder creates a BACnetRelationshipTaggedBuilder
	CreateBACnetRelationshipTaggedBuilder() BACnetRelationshipTaggedBuilder
}

// _BACnetRelationshipTagged is the data-structure of this message
type _BACnetRelationshipTagged struct {
	Header           BACnetTagHeader
	Value            BACnetRelationship
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetRelationshipTagged = (*_BACnetRelationshipTagged)(nil)

// NewBACnetRelationshipTagged factory function for _BACnetRelationshipTagged
func NewBACnetRelationshipTagged(header BACnetTagHeader, value BACnetRelationship, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetRelationshipTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetRelationshipTagged must not be nil")
	}
	return &_BACnetRelationshipTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetRelationshipTaggedBuilder is a builder for BACnetRelationshipTagged
type BACnetRelationshipTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetRelationship, proprietaryValue uint32) BACnetRelationshipTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetRelationshipTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetRelationshipTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetRelationship) BACnetRelationshipTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetRelationshipTaggedBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetRelationshipTaggedBuilder
	// WithArgTagClass sets a parser argument
	WithArgTagClass(TagClass) BACnetRelationshipTaggedBuilder
	// Build builds the BACnetRelationshipTagged or returns an error if something is wrong
	Build() (BACnetRelationshipTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetRelationshipTagged
}

// NewBACnetRelationshipTaggedBuilder() creates a BACnetRelationshipTaggedBuilder
func NewBACnetRelationshipTaggedBuilder() BACnetRelationshipTaggedBuilder {
	return &_BACnetRelationshipTaggedBuilder{_BACnetRelationshipTagged: new(_BACnetRelationshipTagged)}
}

type _BACnetRelationshipTaggedBuilder struct {
	*_BACnetRelationshipTagged

	err *utils.MultiError
}

var _ (BACnetRelationshipTaggedBuilder) = (*_BACnetRelationshipTaggedBuilder)(nil)

func (b *_BACnetRelationshipTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetRelationship, proprietaryValue uint32) BACnetRelationshipTaggedBuilder {
	return b.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (b *_BACnetRelationshipTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetRelationshipTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetRelationshipTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetRelationshipTaggedBuilder {
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

func (b *_BACnetRelationshipTaggedBuilder) WithValue(value BACnetRelationship) BACnetRelationshipTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetRelationshipTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetRelationshipTaggedBuilder {
	b.ProprietaryValue = proprietaryValue
	return b
}

func (b *_BACnetRelationshipTaggedBuilder) WithArgTagNumber(tagNumber uint8) BACnetRelationshipTaggedBuilder {
	b.TagNumber = tagNumber
	return b
}
func (b *_BACnetRelationshipTaggedBuilder) WithArgTagClass(tagClass TagClass) BACnetRelationshipTaggedBuilder {
	b.TagClass = tagClass
	return b
}

func (b *_BACnetRelationshipTaggedBuilder) Build() (BACnetRelationshipTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetRelationshipTagged.deepCopy(), nil
}

func (b *_BACnetRelationshipTaggedBuilder) MustBuild() BACnetRelationshipTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetRelationshipTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetRelationshipTaggedBuilder().(*_BACnetRelationshipTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetRelationshipTaggedBuilder creates a BACnetRelationshipTaggedBuilder
func (b *_BACnetRelationshipTagged) CreateBACnetRelationshipTaggedBuilder() BACnetRelationshipTaggedBuilder {
	if b == nil {
		return NewBACnetRelationshipTaggedBuilder()
	}
	return &_BACnetRelationshipTaggedBuilder{_BACnetRelationshipTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetRelationshipTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetRelationshipTagged) GetValue() BACnetRelationship {
	return m.Value
}

func (m *_BACnetRelationshipTagged) GetProprietaryValue() uint32 {
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

func (m *_BACnetRelationshipTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetRelationship_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetRelationshipTagged(structType any) BACnetRelationshipTagged {
	if casted, ok := structType.(BACnetRelationshipTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetRelationshipTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetRelationshipTagged) GetTypeName() string {
	return "BACnetRelationshipTagged"
}

func (m *_BACnetRelationshipTagged) GetLengthInBits(ctx context.Context) uint16 {
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

func (m *_BACnetRelationshipTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetRelationshipTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetRelationshipTagged, error) {
	return BACnetRelationshipTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetRelationshipTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRelationshipTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRelationshipTagged, error) {
		return BACnetRelationshipTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetRelationshipTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetRelationshipTagged, error) {
	v, err := (&_BACnetRelationshipTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetRelationshipTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetRelationshipTagged BACnetRelationshipTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetRelationshipTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRelationshipTagged")
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

	value, err := ReadManualField[BACnetRelationship](ctx, "value", readBuffer, EnsureType[BACnetRelationship](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetRelationship_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetRelationship_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetRelationshipTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRelationshipTagged")
	}

	return m, nil
}

func (m *_BACnetRelationshipTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetRelationshipTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetRelationshipTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetRelationshipTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetRelationship](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
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

	if popErr := writeBuffer.PopContext("BACnetRelationshipTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetRelationshipTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetRelationshipTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetRelationshipTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetRelationshipTagged) IsBACnetRelationshipTagged() {}

func (m *_BACnetRelationshipTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetRelationshipTagged) deepCopy() *_BACnetRelationshipTagged {
	if m == nil {
		return nil
	}
	_BACnetRelationshipTaggedCopy := &_BACnetRelationshipTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetRelationshipTaggedCopy
}

func (m *_BACnetRelationshipTagged) String() string {
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