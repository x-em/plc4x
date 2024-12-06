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

// BACnetNetworkNumberQualityTagged is the corresponding interface of BACnetNetworkNumberQualityTagged
type BACnetNetworkNumberQualityTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetNetworkNumberQuality
	// IsBACnetNetworkNumberQualityTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNetworkNumberQualityTagged()
	// CreateBuilder creates a BACnetNetworkNumberQualityTaggedBuilder
	CreateBACnetNetworkNumberQualityTaggedBuilder() BACnetNetworkNumberQualityTaggedBuilder
}

// _BACnetNetworkNumberQualityTagged is the data-structure of this message
type _BACnetNetworkNumberQualityTagged struct {
	Header BACnetTagHeader
	Value  BACnetNetworkNumberQuality

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetNetworkNumberQualityTagged = (*_BACnetNetworkNumberQualityTagged)(nil)

// NewBACnetNetworkNumberQualityTagged factory function for _BACnetNetworkNumberQualityTagged
func NewBACnetNetworkNumberQualityTagged(header BACnetTagHeader, value BACnetNetworkNumberQuality, tagNumber uint8, tagClass TagClass) *_BACnetNetworkNumberQualityTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetNetworkNumberQualityTagged must not be nil")
	}
	return &_BACnetNetworkNumberQualityTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNetworkNumberQualityTaggedBuilder is a builder for BACnetNetworkNumberQualityTagged
type BACnetNetworkNumberQualityTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetNetworkNumberQuality) BACnetNetworkNumberQualityTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetNetworkNumberQualityTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetNetworkNumberQualityTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetNetworkNumberQuality) BACnetNetworkNumberQualityTaggedBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetNetworkNumberQualityTaggedBuilder
	// WithArgTagClass sets a parser argument
	WithArgTagClass(TagClass) BACnetNetworkNumberQualityTaggedBuilder
	// Build builds the BACnetNetworkNumberQualityTagged or returns an error if something is wrong
	Build() (BACnetNetworkNumberQualityTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNetworkNumberQualityTagged
}

// NewBACnetNetworkNumberQualityTaggedBuilder() creates a BACnetNetworkNumberQualityTaggedBuilder
func NewBACnetNetworkNumberQualityTaggedBuilder() BACnetNetworkNumberQualityTaggedBuilder {
	return &_BACnetNetworkNumberQualityTaggedBuilder{_BACnetNetworkNumberQualityTagged: new(_BACnetNetworkNumberQualityTagged)}
}

type _BACnetNetworkNumberQualityTaggedBuilder struct {
	*_BACnetNetworkNumberQualityTagged

	err *utils.MultiError
}

var _ (BACnetNetworkNumberQualityTaggedBuilder) = (*_BACnetNetworkNumberQualityTaggedBuilder)(nil)

func (b *_BACnetNetworkNumberQualityTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetNetworkNumberQuality) BACnetNetworkNumberQualityTaggedBuilder {
	return b.WithHeader(header).WithValue(value)
}

func (b *_BACnetNetworkNumberQualityTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetNetworkNumberQualityTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetNetworkNumberQualityTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetNetworkNumberQualityTaggedBuilder {
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

func (b *_BACnetNetworkNumberQualityTaggedBuilder) WithValue(value BACnetNetworkNumberQuality) BACnetNetworkNumberQualityTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetNetworkNumberQualityTaggedBuilder) WithArgTagNumber(tagNumber uint8) BACnetNetworkNumberQualityTaggedBuilder {
	b.TagNumber = tagNumber
	return b
}
func (b *_BACnetNetworkNumberQualityTaggedBuilder) WithArgTagClass(tagClass TagClass) BACnetNetworkNumberQualityTaggedBuilder {
	b.TagClass = tagClass
	return b
}

func (b *_BACnetNetworkNumberQualityTaggedBuilder) Build() (BACnetNetworkNumberQualityTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetNetworkNumberQualityTagged.deepCopy(), nil
}

func (b *_BACnetNetworkNumberQualityTaggedBuilder) MustBuild() BACnetNetworkNumberQualityTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetNetworkNumberQualityTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetNetworkNumberQualityTaggedBuilder().(*_BACnetNetworkNumberQualityTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetNetworkNumberQualityTaggedBuilder creates a BACnetNetworkNumberQualityTaggedBuilder
func (b *_BACnetNetworkNumberQualityTagged) CreateBACnetNetworkNumberQualityTaggedBuilder() BACnetNetworkNumberQualityTaggedBuilder {
	if b == nil {
		return NewBACnetNetworkNumberQualityTaggedBuilder()
	}
	return &_BACnetNetworkNumberQualityTaggedBuilder{_BACnetNetworkNumberQualityTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNetworkNumberQualityTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetNetworkNumberQualityTagged) GetValue() BACnetNetworkNumberQuality {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNetworkNumberQualityTagged(structType any) BACnetNetworkNumberQualityTagged {
	if casted, ok := structType.(BACnetNetworkNumberQualityTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNetworkNumberQualityTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNetworkNumberQualityTagged) GetTypeName() string {
	return "BACnetNetworkNumberQualityTagged"
}

func (m *_BACnetNetworkNumberQualityTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetNetworkNumberQualityTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNetworkNumberQualityTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetNetworkNumberQualityTagged, error) {
	return BACnetNetworkNumberQualityTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetNetworkNumberQualityTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNetworkNumberQualityTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNetworkNumberQualityTagged, error) {
		return BACnetNetworkNumberQualityTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetNetworkNumberQualityTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetNetworkNumberQualityTagged, error) {
	v, err := (&_BACnetNetworkNumberQualityTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetNetworkNumberQualityTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetNetworkNumberQualityTagged BACnetNetworkNumberQualityTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNetworkNumberQualityTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNetworkNumberQualityTagged")
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

	value, err := ReadManualField[BACnetNetworkNumberQuality](ctx, "value", readBuffer, EnsureType[BACnetNetworkNumberQuality](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetNetworkNumberQuality_UNKNOWN)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetNetworkNumberQualityTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNetworkNumberQualityTagged")
	}

	return m, nil
}

func (m *_BACnetNetworkNumberQualityTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNetworkNumberQualityTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetNetworkNumberQualityTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetNetworkNumberQualityTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetNetworkNumberQuality](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetNetworkNumberQualityTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetNetworkNumberQualityTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetNetworkNumberQualityTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetNetworkNumberQualityTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetNetworkNumberQualityTagged) IsBACnetNetworkNumberQualityTagged() {}

func (m *_BACnetNetworkNumberQualityTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNetworkNumberQualityTagged) deepCopy() *_BACnetNetworkNumberQualityTagged {
	if m == nil {
		return nil
	}
	_BACnetNetworkNumberQualityTaggedCopy := &_BACnetNetworkNumberQualityTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetNetworkNumberQualityTaggedCopy
}

func (m *_BACnetNetworkNumberQualityTagged) String() string {
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