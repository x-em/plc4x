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

// BACnetLiftCarCallListFloorList is the corresponding interface of BACnetLiftCarCallListFloorList
type BACnetLiftCarCallListFloorList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetFloorNumbers returns FloorNumbers (property field)
	GetFloorNumbers() []BACnetApplicationTagUnsignedInteger
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetLiftCarCallListFloorList is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLiftCarCallListFloorList()
	// CreateBuilder creates a BACnetLiftCarCallListFloorListBuilder
	CreateBACnetLiftCarCallListFloorListBuilder() BACnetLiftCarCallListFloorListBuilder
}

// _BACnetLiftCarCallListFloorList is the data-structure of this message
type _BACnetLiftCarCallListFloorList struct {
	OpeningTag   BACnetOpeningTag
	FloorNumbers []BACnetApplicationTagUnsignedInteger
	ClosingTag   BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetLiftCarCallListFloorList = (*_BACnetLiftCarCallListFloorList)(nil)

// NewBACnetLiftCarCallListFloorList factory function for _BACnetLiftCarCallListFloorList
func NewBACnetLiftCarCallListFloorList(openingTag BACnetOpeningTag, floorNumbers []BACnetApplicationTagUnsignedInteger, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLiftCarCallListFloorList {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetLiftCarCallListFloorList must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetLiftCarCallListFloorList must not be nil")
	}
	return &_BACnetLiftCarCallListFloorList{OpeningTag: openingTag, FloorNumbers: floorNumbers, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLiftCarCallListFloorListBuilder is a builder for BACnetLiftCarCallListFloorList
type BACnetLiftCarCallListFloorListBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, floorNumbers []BACnetApplicationTagUnsignedInteger, closingTag BACnetClosingTag) BACnetLiftCarCallListFloorListBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetLiftCarCallListFloorListBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetLiftCarCallListFloorListBuilder
	// WithFloorNumbers adds FloorNumbers (property field)
	WithFloorNumbers(...BACnetApplicationTagUnsignedInteger) BACnetLiftCarCallListFloorListBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetLiftCarCallListFloorListBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetLiftCarCallListFloorListBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetLiftCarCallListFloorListBuilder
	// Build builds the BACnetLiftCarCallListFloorList or returns an error if something is wrong
	Build() (BACnetLiftCarCallListFloorList, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLiftCarCallListFloorList
}

// NewBACnetLiftCarCallListFloorListBuilder() creates a BACnetLiftCarCallListFloorListBuilder
func NewBACnetLiftCarCallListFloorListBuilder() BACnetLiftCarCallListFloorListBuilder {
	return &_BACnetLiftCarCallListFloorListBuilder{_BACnetLiftCarCallListFloorList: new(_BACnetLiftCarCallListFloorList)}
}

type _BACnetLiftCarCallListFloorListBuilder struct {
	*_BACnetLiftCarCallListFloorList

	err *utils.MultiError
}

var _ (BACnetLiftCarCallListFloorListBuilder) = (*_BACnetLiftCarCallListFloorListBuilder)(nil)

func (b *_BACnetLiftCarCallListFloorListBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, floorNumbers []BACnetApplicationTagUnsignedInteger, closingTag BACnetClosingTag) BACnetLiftCarCallListFloorListBuilder {
	return b.WithOpeningTag(openingTag).WithFloorNumbers(floorNumbers...).WithClosingTag(closingTag)
}

func (b *_BACnetLiftCarCallListFloorListBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetLiftCarCallListFloorListBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetLiftCarCallListFloorListBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetLiftCarCallListFloorListBuilder {
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

func (b *_BACnetLiftCarCallListFloorListBuilder) WithFloorNumbers(floorNumbers ...BACnetApplicationTagUnsignedInteger) BACnetLiftCarCallListFloorListBuilder {
	b.FloorNumbers = floorNumbers
	return b
}

func (b *_BACnetLiftCarCallListFloorListBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetLiftCarCallListFloorListBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetLiftCarCallListFloorListBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetLiftCarCallListFloorListBuilder {
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

func (b *_BACnetLiftCarCallListFloorListBuilder) WithArgTagNumber(tagNumber uint8) BACnetLiftCarCallListFloorListBuilder {
	b.TagNumber = tagNumber
	return b
}

func (b *_BACnetLiftCarCallListFloorListBuilder) Build() (BACnetLiftCarCallListFloorList, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
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
	return b._BACnetLiftCarCallListFloorList.deepCopy(), nil
}

func (b *_BACnetLiftCarCallListFloorListBuilder) MustBuild() BACnetLiftCarCallListFloorList {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetLiftCarCallListFloorListBuilder) DeepCopy() any {
	_copy := b.CreateBACnetLiftCarCallListFloorListBuilder().(*_BACnetLiftCarCallListFloorListBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetLiftCarCallListFloorListBuilder creates a BACnetLiftCarCallListFloorListBuilder
func (b *_BACnetLiftCarCallListFloorList) CreateBACnetLiftCarCallListFloorListBuilder() BACnetLiftCarCallListFloorListBuilder {
	if b == nil {
		return NewBACnetLiftCarCallListFloorListBuilder()
	}
	return &_BACnetLiftCarCallListFloorListBuilder{_BACnetLiftCarCallListFloorList: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLiftCarCallListFloorList) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetLiftCarCallListFloorList) GetFloorNumbers() []BACnetApplicationTagUnsignedInteger {
	return m.FloorNumbers
}

func (m *_BACnetLiftCarCallListFloorList) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLiftCarCallListFloorList(structType any) BACnetLiftCarCallListFloorList {
	if casted, ok := structType.(BACnetLiftCarCallListFloorList); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLiftCarCallListFloorList); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLiftCarCallListFloorList) GetTypeName() string {
	return "BACnetLiftCarCallListFloorList"
}

func (m *_BACnetLiftCarCallListFloorList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.FloorNumbers) > 0 {
		for _, element := range m.FloorNumbers {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLiftCarCallListFloorList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLiftCarCallListFloorListParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetLiftCarCallListFloorList, error) {
	return BACnetLiftCarCallListFloorListParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetLiftCarCallListFloorListParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLiftCarCallListFloorList, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLiftCarCallListFloorList, error) {
		return BACnetLiftCarCallListFloorListParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetLiftCarCallListFloorListParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetLiftCarCallListFloorList, error) {
	v, err := (&_BACnetLiftCarCallListFloorList{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetLiftCarCallListFloorList) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetLiftCarCallListFloorList BACnetLiftCarCallListFloorList, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLiftCarCallListFloorList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLiftCarCallListFloorList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	floorNumbers, err := ReadTerminatedArrayField[BACnetApplicationTagUnsignedInteger](ctx, "floorNumbers", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'floorNumbers' field"))
	}
	m.FloorNumbers = floorNumbers

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetLiftCarCallListFloorList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLiftCarCallListFloorList")
	}

	return m, nil
}

func (m *_BACnetLiftCarCallListFloorList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLiftCarCallListFloorList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLiftCarCallListFloorList"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLiftCarCallListFloorList")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "floorNumbers", m.GetFloorNumbers(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'floorNumbers' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLiftCarCallListFloorList"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLiftCarCallListFloorList")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetLiftCarCallListFloorList) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetLiftCarCallListFloorList) IsBACnetLiftCarCallListFloorList() {}

func (m *_BACnetLiftCarCallListFloorList) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLiftCarCallListFloorList) deepCopy() *_BACnetLiftCarCallListFloorList {
	if m == nil {
		return nil
	}
	_BACnetLiftCarCallListFloorListCopy := &_BACnetLiftCarCallListFloorList{
		utils.DeepCopy[BACnetOpeningTag](m.OpeningTag),
		utils.DeepCopySlice[BACnetApplicationTagUnsignedInteger, BACnetApplicationTagUnsignedInteger](m.FloorNumbers),
		utils.DeepCopy[BACnetClosingTag](m.ClosingTag),
		m.TagNumber,
	}
	return _BACnetLiftCarCallListFloorListCopy
}

func (m *_BACnetLiftCarCallListFloorList) String() string {
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