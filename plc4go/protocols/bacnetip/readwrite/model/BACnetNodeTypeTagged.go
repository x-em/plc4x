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

// BACnetNodeTypeTagged is the corresponding interface of BACnetNodeTypeTagged
type BACnetNodeTypeTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetNodeType
	// IsBACnetNodeTypeTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNodeTypeTagged()
	// CreateBuilder creates a BACnetNodeTypeTaggedBuilder
	CreateBACnetNodeTypeTaggedBuilder() BACnetNodeTypeTaggedBuilder
}

// _BACnetNodeTypeTagged is the data-structure of this message
type _BACnetNodeTypeTagged struct {
	Header BACnetTagHeader
	Value  BACnetNodeType

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetNodeTypeTagged = (*_BACnetNodeTypeTagged)(nil)

// NewBACnetNodeTypeTagged factory function for _BACnetNodeTypeTagged
func NewBACnetNodeTypeTagged(header BACnetTagHeader, value BACnetNodeType, tagNumber uint8, tagClass TagClass) *_BACnetNodeTypeTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetNodeTypeTagged must not be nil")
	}
	return &_BACnetNodeTypeTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNodeTypeTaggedBuilder is a builder for BACnetNodeTypeTagged
type BACnetNodeTypeTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetNodeType) BACnetNodeTypeTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetNodeTypeTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetNodeTypeTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetNodeType) BACnetNodeTypeTaggedBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetNodeTypeTaggedBuilder
	// WithArgTagClass sets a parser argument
	WithArgTagClass(TagClass) BACnetNodeTypeTaggedBuilder
	// Build builds the BACnetNodeTypeTagged or returns an error if something is wrong
	Build() (BACnetNodeTypeTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNodeTypeTagged
}

// NewBACnetNodeTypeTaggedBuilder() creates a BACnetNodeTypeTaggedBuilder
func NewBACnetNodeTypeTaggedBuilder() BACnetNodeTypeTaggedBuilder {
	return &_BACnetNodeTypeTaggedBuilder{_BACnetNodeTypeTagged: new(_BACnetNodeTypeTagged)}
}

type _BACnetNodeTypeTaggedBuilder struct {
	*_BACnetNodeTypeTagged

	err *utils.MultiError
}

var _ (BACnetNodeTypeTaggedBuilder) = (*_BACnetNodeTypeTaggedBuilder)(nil)

func (b *_BACnetNodeTypeTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetNodeType) BACnetNodeTypeTaggedBuilder {
	return b.WithHeader(header).WithValue(value)
}

func (b *_BACnetNodeTypeTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetNodeTypeTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetNodeTypeTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetNodeTypeTaggedBuilder {
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

func (b *_BACnetNodeTypeTaggedBuilder) WithValue(value BACnetNodeType) BACnetNodeTypeTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetNodeTypeTaggedBuilder) WithArgTagNumber(tagNumber uint8) BACnetNodeTypeTaggedBuilder {
	b.TagNumber = tagNumber
	return b
}
func (b *_BACnetNodeTypeTaggedBuilder) WithArgTagClass(tagClass TagClass) BACnetNodeTypeTaggedBuilder {
	b.TagClass = tagClass
	return b
}

func (b *_BACnetNodeTypeTaggedBuilder) Build() (BACnetNodeTypeTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetNodeTypeTagged.deepCopy(), nil
}

func (b *_BACnetNodeTypeTaggedBuilder) MustBuild() BACnetNodeTypeTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetNodeTypeTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetNodeTypeTaggedBuilder().(*_BACnetNodeTypeTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetNodeTypeTaggedBuilder creates a BACnetNodeTypeTaggedBuilder
func (b *_BACnetNodeTypeTagged) CreateBACnetNodeTypeTaggedBuilder() BACnetNodeTypeTaggedBuilder {
	if b == nil {
		return NewBACnetNodeTypeTaggedBuilder()
	}
	return &_BACnetNodeTypeTaggedBuilder{_BACnetNodeTypeTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNodeTypeTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetNodeTypeTagged) GetValue() BACnetNodeType {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNodeTypeTagged(structType any) BACnetNodeTypeTagged {
	if casted, ok := structType.(BACnetNodeTypeTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNodeTypeTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNodeTypeTagged) GetTypeName() string {
	return "BACnetNodeTypeTagged"
}

func (m *_BACnetNodeTypeTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetNodeTypeTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNodeTypeTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetNodeTypeTagged, error) {
	return BACnetNodeTypeTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetNodeTypeTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNodeTypeTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNodeTypeTagged, error) {
		return BACnetNodeTypeTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetNodeTypeTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetNodeTypeTagged, error) {
	v, err := (&_BACnetNodeTypeTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetNodeTypeTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetNodeTypeTagged BACnetNodeTypeTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNodeTypeTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNodeTypeTagged")
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

	value, err := ReadManualField[BACnetNodeType](ctx, "value", readBuffer, EnsureType[BACnetNodeType](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetNodeType_UNKNOWN)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetNodeTypeTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNodeTypeTagged")
	}

	return m, nil
}

func (m *_BACnetNodeTypeTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNodeTypeTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetNodeTypeTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetNodeTypeTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetNodeType](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetNodeTypeTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetNodeTypeTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetNodeTypeTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetNodeTypeTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetNodeTypeTagged) IsBACnetNodeTypeTagged() {}

func (m *_BACnetNodeTypeTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNodeTypeTagged) deepCopy() *_BACnetNodeTypeTagged {
	if m == nil {
		return nil
	}
	_BACnetNodeTypeTaggedCopy := &_BACnetNodeTypeTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetNodeTypeTaggedCopy
}

func (m *_BACnetNodeTypeTagged) String() string {
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