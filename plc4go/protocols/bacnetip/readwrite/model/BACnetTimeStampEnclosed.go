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

// BACnetTimeStampEnclosed is the corresponding interface of BACnetTimeStampEnclosed
type BACnetTimeStampEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetTimeStamp
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetTimeStampEnclosed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimeStampEnclosed()
	// CreateBuilder creates a BACnetTimeStampEnclosedBuilder
	CreateBACnetTimeStampEnclosedBuilder() BACnetTimeStampEnclosedBuilder
}

// _BACnetTimeStampEnclosed is the data-structure of this message
type _BACnetTimeStampEnclosed struct {
	OpeningTag BACnetOpeningTag
	Timestamp  BACnetTimeStamp
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetTimeStampEnclosed = (*_BACnetTimeStampEnclosed)(nil)

// NewBACnetTimeStampEnclosed factory function for _BACnetTimeStampEnclosed
func NewBACnetTimeStampEnclosed(openingTag BACnetOpeningTag, timestamp BACnetTimeStamp, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetTimeStampEnclosed {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetTimeStampEnclosed must not be nil")
	}
	if timestamp == nil {
		panic("timestamp of type BACnetTimeStamp for BACnetTimeStampEnclosed must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetTimeStampEnclosed must not be nil")
	}
	return &_BACnetTimeStampEnclosed{OpeningTag: openingTag, Timestamp: timestamp, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTimeStampEnclosedBuilder is a builder for BACnetTimeStampEnclosed
type BACnetTimeStampEnclosedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, timestamp BACnetTimeStamp, closingTag BACnetClosingTag) BACnetTimeStampEnclosedBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetTimeStampEnclosedBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetTimeStampEnclosedBuilder
	// WithTimestamp adds Timestamp (property field)
	WithTimestamp(BACnetTimeStamp) BACnetTimeStampEnclosedBuilder
	// WithTimestampBuilder adds Timestamp (property field) which is build by the builder
	WithTimestampBuilder(func(BACnetTimeStampBuilder) BACnetTimeStampBuilder) BACnetTimeStampEnclosedBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetTimeStampEnclosedBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetTimeStampEnclosedBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetTimeStampEnclosedBuilder
	// Build builds the BACnetTimeStampEnclosed or returns an error if something is wrong
	Build() (BACnetTimeStampEnclosed, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTimeStampEnclosed
}

// NewBACnetTimeStampEnclosedBuilder() creates a BACnetTimeStampEnclosedBuilder
func NewBACnetTimeStampEnclosedBuilder() BACnetTimeStampEnclosedBuilder {
	return &_BACnetTimeStampEnclosedBuilder{_BACnetTimeStampEnclosed: new(_BACnetTimeStampEnclosed)}
}

type _BACnetTimeStampEnclosedBuilder struct {
	*_BACnetTimeStampEnclosed

	err *utils.MultiError
}

var _ (BACnetTimeStampEnclosedBuilder) = (*_BACnetTimeStampEnclosedBuilder)(nil)

func (b *_BACnetTimeStampEnclosedBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, timestamp BACnetTimeStamp, closingTag BACnetClosingTag) BACnetTimeStampEnclosedBuilder {
	return b.WithOpeningTag(openingTag).WithTimestamp(timestamp).WithClosingTag(closingTag)
}

func (b *_BACnetTimeStampEnclosedBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetTimeStampEnclosedBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetTimeStampEnclosedBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetTimeStampEnclosedBuilder {
	builder := builderSupplier(b.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.OpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetTimeStampEnclosedBuilder) WithTimestamp(timestamp BACnetTimeStamp) BACnetTimeStampEnclosedBuilder {
	b.Timestamp = timestamp
	return b
}

func (b *_BACnetTimeStampEnclosedBuilder) WithTimestampBuilder(builderSupplier func(BACnetTimeStampBuilder) BACnetTimeStampBuilder) BACnetTimeStampEnclosedBuilder {
	builder := builderSupplier(b.Timestamp.CreateBACnetTimeStampBuilder())
	var err error
	b.Timestamp, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTimeStampBuilder failed"))
	}
	return b
}

func (b *_BACnetTimeStampEnclosedBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetTimeStampEnclosedBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetTimeStampEnclosedBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetTimeStampEnclosedBuilder {
	builder := builderSupplier(b.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.ClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetTimeStampEnclosedBuilder) WithArgTagNumber(tagNumber uint8) BACnetTimeStampEnclosedBuilder {
	b.TagNumber = tagNumber
	return b
}

func (b *_BACnetTimeStampEnclosedBuilder) Build() (BACnetTimeStampEnclosed, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.Timestamp == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timestamp' not set"))
	}
	if b.ClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetTimeStampEnclosed.deepCopy(), nil
}

func (b *_BACnetTimeStampEnclosedBuilder) MustBuild() BACnetTimeStampEnclosed {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetTimeStampEnclosedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetTimeStampEnclosedBuilder().(*_BACnetTimeStampEnclosedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetTimeStampEnclosedBuilder creates a BACnetTimeStampEnclosedBuilder
func (b *_BACnetTimeStampEnclosed) CreateBACnetTimeStampEnclosedBuilder() BACnetTimeStampEnclosedBuilder {
	if b == nil {
		return NewBACnetTimeStampEnclosedBuilder()
	}
	return &_BACnetTimeStampEnclosedBuilder{_BACnetTimeStampEnclosed: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimeStampEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetTimeStampEnclosed) GetTimestamp() BACnetTimeStamp {
	return m.Timestamp
}

func (m *_BACnetTimeStampEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTimeStampEnclosed(structType any) BACnetTimeStampEnclosed {
	if casted, ok := structType.(BACnetTimeStampEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimeStampEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimeStampEnclosed) GetTypeName() string {
	return "BACnetTimeStampEnclosed"
}

func (m *_BACnetTimeStampEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (timestamp)
	lengthInBits += m.Timestamp.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimeStampEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTimeStampEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetTimeStampEnclosed, error) {
	return BACnetTimeStampEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetTimeStampEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimeStampEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimeStampEnclosed, error) {
		return BACnetTimeStampEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetTimeStampEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetTimeStampEnclosed, error) {
	v, err := (&_BACnetTimeStampEnclosed{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetTimeStampEnclosed) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetTimeStampEnclosed BACnetTimeStampEnclosed, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimeStampEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimeStampEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	timestamp, err := ReadSimpleField[BACnetTimeStamp](ctx, "timestamp", ReadComplex[BACnetTimeStamp](BACnetTimeStampParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestamp' field"))
	}
	m.Timestamp = timestamp

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetTimeStampEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimeStampEnclosed")
	}

	return m, nil
}

func (m *_BACnetTimeStampEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimeStampEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTimeStampEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTimeStampEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetTimeStamp](ctx, "timestamp", m.GetTimestamp(), WriteComplex[BACnetTimeStamp](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'timestamp' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTimeStampEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTimeStampEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetTimeStampEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetTimeStampEnclosed) IsBACnetTimeStampEnclosed() {}

func (m *_BACnetTimeStampEnclosed) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTimeStampEnclosed) deepCopy() *_BACnetTimeStampEnclosed {
	if m == nil {
		return nil
	}
	_BACnetTimeStampEnclosedCopy := &_BACnetTimeStampEnclosed{
		utils.DeepCopy[BACnetOpeningTag](m.OpeningTag),
		utils.DeepCopy[BACnetTimeStamp](m.Timestamp),
		utils.DeepCopy[BACnetClosingTag](m.ClosingTag),
		m.TagNumber,
	}
	return _BACnetTimeStampEnclosedCopy
}

func (m *_BACnetTimeStampEnclosed) String() string {
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