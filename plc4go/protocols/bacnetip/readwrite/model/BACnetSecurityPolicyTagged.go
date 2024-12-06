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

// BACnetSecurityPolicyTagged is the corresponding interface of BACnetSecurityPolicyTagged
type BACnetSecurityPolicyTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetSecurityPolicy
	// IsBACnetSecurityPolicyTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetSecurityPolicyTagged()
	// CreateBuilder creates a BACnetSecurityPolicyTaggedBuilder
	CreateBACnetSecurityPolicyTaggedBuilder() BACnetSecurityPolicyTaggedBuilder
}

// _BACnetSecurityPolicyTagged is the data-structure of this message
type _BACnetSecurityPolicyTagged struct {
	Header BACnetTagHeader
	Value  BACnetSecurityPolicy

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetSecurityPolicyTagged = (*_BACnetSecurityPolicyTagged)(nil)

// NewBACnetSecurityPolicyTagged factory function for _BACnetSecurityPolicyTagged
func NewBACnetSecurityPolicyTagged(header BACnetTagHeader, value BACnetSecurityPolicy, tagNumber uint8, tagClass TagClass) *_BACnetSecurityPolicyTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetSecurityPolicyTagged must not be nil")
	}
	return &_BACnetSecurityPolicyTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetSecurityPolicyTaggedBuilder is a builder for BACnetSecurityPolicyTagged
type BACnetSecurityPolicyTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetSecurityPolicy) BACnetSecurityPolicyTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetSecurityPolicyTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetSecurityPolicyTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetSecurityPolicy) BACnetSecurityPolicyTaggedBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetSecurityPolicyTaggedBuilder
	// WithArgTagClass sets a parser argument
	WithArgTagClass(TagClass) BACnetSecurityPolicyTaggedBuilder
	// Build builds the BACnetSecurityPolicyTagged or returns an error if something is wrong
	Build() (BACnetSecurityPolicyTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetSecurityPolicyTagged
}

// NewBACnetSecurityPolicyTaggedBuilder() creates a BACnetSecurityPolicyTaggedBuilder
func NewBACnetSecurityPolicyTaggedBuilder() BACnetSecurityPolicyTaggedBuilder {
	return &_BACnetSecurityPolicyTaggedBuilder{_BACnetSecurityPolicyTagged: new(_BACnetSecurityPolicyTagged)}
}

type _BACnetSecurityPolicyTaggedBuilder struct {
	*_BACnetSecurityPolicyTagged

	err *utils.MultiError
}

var _ (BACnetSecurityPolicyTaggedBuilder) = (*_BACnetSecurityPolicyTaggedBuilder)(nil)

func (b *_BACnetSecurityPolicyTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetSecurityPolicy) BACnetSecurityPolicyTaggedBuilder {
	return b.WithHeader(header).WithValue(value)
}

func (b *_BACnetSecurityPolicyTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetSecurityPolicyTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetSecurityPolicyTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetSecurityPolicyTaggedBuilder {
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

func (b *_BACnetSecurityPolicyTaggedBuilder) WithValue(value BACnetSecurityPolicy) BACnetSecurityPolicyTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetSecurityPolicyTaggedBuilder) WithArgTagNumber(tagNumber uint8) BACnetSecurityPolicyTaggedBuilder {
	b.TagNumber = tagNumber
	return b
}
func (b *_BACnetSecurityPolicyTaggedBuilder) WithArgTagClass(tagClass TagClass) BACnetSecurityPolicyTaggedBuilder {
	b.TagClass = tagClass
	return b
}

func (b *_BACnetSecurityPolicyTaggedBuilder) Build() (BACnetSecurityPolicyTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetSecurityPolicyTagged.deepCopy(), nil
}

func (b *_BACnetSecurityPolicyTaggedBuilder) MustBuild() BACnetSecurityPolicyTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetSecurityPolicyTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetSecurityPolicyTaggedBuilder().(*_BACnetSecurityPolicyTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetSecurityPolicyTaggedBuilder creates a BACnetSecurityPolicyTaggedBuilder
func (b *_BACnetSecurityPolicyTagged) CreateBACnetSecurityPolicyTaggedBuilder() BACnetSecurityPolicyTaggedBuilder {
	if b == nil {
		return NewBACnetSecurityPolicyTaggedBuilder()
	}
	return &_BACnetSecurityPolicyTaggedBuilder{_BACnetSecurityPolicyTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetSecurityPolicyTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetSecurityPolicyTagged) GetValue() BACnetSecurityPolicy {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetSecurityPolicyTagged(structType any) BACnetSecurityPolicyTagged {
	if casted, ok := structType.(BACnetSecurityPolicyTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetSecurityPolicyTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetSecurityPolicyTagged) GetTypeName() string {
	return "BACnetSecurityPolicyTagged"
}

func (m *_BACnetSecurityPolicyTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetSecurityPolicyTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetSecurityPolicyTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetSecurityPolicyTagged, error) {
	return BACnetSecurityPolicyTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetSecurityPolicyTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSecurityPolicyTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSecurityPolicyTagged, error) {
		return BACnetSecurityPolicyTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetSecurityPolicyTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetSecurityPolicyTagged, error) {
	v, err := (&_BACnetSecurityPolicyTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetSecurityPolicyTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetSecurityPolicyTagged BACnetSecurityPolicyTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetSecurityPolicyTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetSecurityPolicyTagged")
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

	value, err := ReadManualField[BACnetSecurityPolicy](ctx, "value", readBuffer, EnsureType[BACnetSecurityPolicy](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetSecurityPolicy_PLAIN_NON_TRUSTED)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetSecurityPolicyTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetSecurityPolicyTagged")
	}

	return m, nil
}

func (m *_BACnetSecurityPolicyTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetSecurityPolicyTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetSecurityPolicyTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetSecurityPolicyTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetSecurityPolicy](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetSecurityPolicyTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetSecurityPolicyTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetSecurityPolicyTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetSecurityPolicyTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetSecurityPolicyTagged) IsBACnetSecurityPolicyTagged() {}

func (m *_BACnetSecurityPolicyTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetSecurityPolicyTagged) deepCopy() *_BACnetSecurityPolicyTagged {
	if m == nil {
		return nil
	}
	_BACnetSecurityPolicyTaggedCopy := &_BACnetSecurityPolicyTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetSecurityPolicyTaggedCopy
}

func (m *_BACnetSecurityPolicyTagged) String() string {
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