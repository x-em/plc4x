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

// BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter is the corresponding interface of BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter
type BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetMinPriority returns MinPriority (property field)
	GetMinPriority() BACnetContextTagUnsignedInteger
	// GetMaxPriority returns MaxPriority (property field)
	GetMaxPriority() BACnetContextTagUnsignedInteger
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter()
}

// _BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter is the data-structure of this message
type _BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter struct {
	OpeningTag  BACnetOpeningTag
	MinPriority BACnetContextTagUnsignedInteger
	MaxPriority BACnetContextTagUnsignedInteger
	ClosingTag  BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter = (*_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter)(nil)

// NewBACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter factory function for _BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter
func NewBACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter(openingTag BACnetOpeningTag, minPriority BACnetContextTagUnsignedInteger, maxPriority BACnetContextTagUnsignedInteger, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter must not be nil")
	}
	if minPriority == nil {
		panic("minPriority of type BACnetContextTagUnsignedInteger for BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter must not be nil")
	}
	if maxPriority == nil {
		panic("maxPriority of type BACnetContextTagUnsignedInteger for BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter must not be nil")
	}
	return &_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter{OpeningTag: openingTag, MinPriority: minPriority, MaxPriority: maxPriority, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) GetMinPriority() BACnetContextTagUnsignedInteger {
	return m.MinPriority
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) GetMaxPriority() BACnetContextTagUnsignedInteger {
	return m.MaxPriority
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter(structType any) BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter {
	if casted, ok := structType.(BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) GetTypeName() string {
	return "BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter"
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (minPriority)
	lengthInBits += m.MinPriority.GetLengthInBits(ctx)

	// Simple field (maxPriority)
	lengthInBits += m.MaxPriority.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilterParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter, error) {
	return BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilterParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilterParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter, error) {
		return BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilterParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilterParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter, error) {
	v, err := (&_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	minPriority, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "minPriority", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minPriority' field"))
	}
	m.MinPriority = minPriority

	maxPriority, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "maxPriority", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxPriority' field"))
	}
	m.MaxPriority = maxPriority

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "minPriority", m.GetMinPriority(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'minPriority' field")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "maxPriority", m.GetMaxPriority(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'maxPriority' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) IsBACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter() {
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
