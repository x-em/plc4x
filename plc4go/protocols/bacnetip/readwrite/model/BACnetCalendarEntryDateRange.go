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

// BACnetCalendarEntryDateRange is the corresponding interface of BACnetCalendarEntryDateRange
type BACnetCalendarEntryDateRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetCalendarEntry
	// GetDateRange returns DateRange (property field)
	GetDateRange() BACnetDateRangeEnclosed
	// IsBACnetCalendarEntryDateRange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetCalendarEntryDateRange()
	// CreateBuilder creates a BACnetCalendarEntryDateRangeBuilder
	CreateBACnetCalendarEntryDateRangeBuilder() BACnetCalendarEntryDateRangeBuilder
}

// _BACnetCalendarEntryDateRange is the data-structure of this message
type _BACnetCalendarEntryDateRange struct {
	BACnetCalendarEntryContract
	DateRange BACnetDateRangeEnclosed
}

var _ BACnetCalendarEntryDateRange = (*_BACnetCalendarEntryDateRange)(nil)
var _ BACnetCalendarEntryRequirements = (*_BACnetCalendarEntryDateRange)(nil)

// NewBACnetCalendarEntryDateRange factory function for _BACnetCalendarEntryDateRange
func NewBACnetCalendarEntryDateRange(peekedTagHeader BACnetTagHeader, dateRange BACnetDateRangeEnclosed) *_BACnetCalendarEntryDateRange {
	if dateRange == nil {
		panic("dateRange of type BACnetDateRangeEnclosed for BACnetCalendarEntryDateRange must not be nil")
	}
	_result := &_BACnetCalendarEntryDateRange{
		BACnetCalendarEntryContract: NewBACnetCalendarEntry(peekedTagHeader),
		DateRange:                   dateRange,
	}
	_result.BACnetCalendarEntryContract.(*_BACnetCalendarEntry)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetCalendarEntryDateRangeBuilder is a builder for BACnetCalendarEntryDateRange
type BACnetCalendarEntryDateRangeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(dateRange BACnetDateRangeEnclosed) BACnetCalendarEntryDateRangeBuilder
	// WithDateRange adds DateRange (property field)
	WithDateRange(BACnetDateRangeEnclosed) BACnetCalendarEntryDateRangeBuilder
	// WithDateRangeBuilder adds DateRange (property field) which is build by the builder
	WithDateRangeBuilder(func(BACnetDateRangeEnclosedBuilder) BACnetDateRangeEnclosedBuilder) BACnetCalendarEntryDateRangeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetCalendarEntryBuilder
	// Build builds the BACnetCalendarEntryDateRange or returns an error if something is wrong
	Build() (BACnetCalendarEntryDateRange, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetCalendarEntryDateRange
}

// NewBACnetCalendarEntryDateRangeBuilder() creates a BACnetCalendarEntryDateRangeBuilder
func NewBACnetCalendarEntryDateRangeBuilder() BACnetCalendarEntryDateRangeBuilder {
	return &_BACnetCalendarEntryDateRangeBuilder{_BACnetCalendarEntryDateRange: new(_BACnetCalendarEntryDateRange)}
}

type _BACnetCalendarEntryDateRangeBuilder struct {
	*_BACnetCalendarEntryDateRange

	parentBuilder *_BACnetCalendarEntryBuilder

	err *utils.MultiError
}

var _ (BACnetCalendarEntryDateRangeBuilder) = (*_BACnetCalendarEntryDateRangeBuilder)(nil)

func (b *_BACnetCalendarEntryDateRangeBuilder) setParent(contract BACnetCalendarEntryContract) {
	b.BACnetCalendarEntryContract = contract
	contract.(*_BACnetCalendarEntry)._SubType = b._BACnetCalendarEntryDateRange
}

func (b *_BACnetCalendarEntryDateRangeBuilder) WithMandatoryFields(dateRange BACnetDateRangeEnclosed) BACnetCalendarEntryDateRangeBuilder {
	return b.WithDateRange(dateRange)
}

func (b *_BACnetCalendarEntryDateRangeBuilder) WithDateRange(dateRange BACnetDateRangeEnclosed) BACnetCalendarEntryDateRangeBuilder {
	b.DateRange = dateRange
	return b
}

func (b *_BACnetCalendarEntryDateRangeBuilder) WithDateRangeBuilder(builderSupplier func(BACnetDateRangeEnclosedBuilder) BACnetDateRangeEnclosedBuilder) BACnetCalendarEntryDateRangeBuilder {
	builder := builderSupplier(b.DateRange.CreateBACnetDateRangeEnclosedBuilder())
	var err error
	b.DateRange, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDateRangeEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetCalendarEntryDateRangeBuilder) Build() (BACnetCalendarEntryDateRange, error) {
	if b.DateRange == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'dateRange' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetCalendarEntryDateRange.deepCopy(), nil
}

func (b *_BACnetCalendarEntryDateRangeBuilder) MustBuild() BACnetCalendarEntryDateRange {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetCalendarEntryDateRangeBuilder) Done() BACnetCalendarEntryBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetCalendarEntryBuilder().(*_BACnetCalendarEntryBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetCalendarEntryDateRangeBuilder) buildForBACnetCalendarEntry() (BACnetCalendarEntry, error) {
	return b.Build()
}

func (b *_BACnetCalendarEntryDateRangeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetCalendarEntryDateRangeBuilder().(*_BACnetCalendarEntryDateRangeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetCalendarEntryDateRangeBuilder creates a BACnetCalendarEntryDateRangeBuilder
func (b *_BACnetCalendarEntryDateRange) CreateBACnetCalendarEntryDateRangeBuilder() BACnetCalendarEntryDateRangeBuilder {
	if b == nil {
		return NewBACnetCalendarEntryDateRangeBuilder()
	}
	return &_BACnetCalendarEntryDateRangeBuilder{_BACnetCalendarEntryDateRange: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetCalendarEntryDateRange) GetParent() BACnetCalendarEntryContract {
	return m.BACnetCalendarEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCalendarEntryDateRange) GetDateRange() BACnetDateRangeEnclosed {
	return m.DateRange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetCalendarEntryDateRange(structType any) BACnetCalendarEntryDateRange {
	if casted, ok := structType.(BACnetCalendarEntryDateRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCalendarEntryDateRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCalendarEntryDateRange) GetTypeName() string {
	return "BACnetCalendarEntryDateRange"
}

func (m *_BACnetCalendarEntryDateRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetCalendarEntryContract.(*_BACnetCalendarEntry).getLengthInBits(ctx))

	// Simple field (dateRange)
	lengthInBits += m.DateRange.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetCalendarEntryDateRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetCalendarEntryDateRange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetCalendarEntry) (__bACnetCalendarEntryDateRange BACnetCalendarEntryDateRange, err error) {
	m.BACnetCalendarEntryContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCalendarEntryDateRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCalendarEntryDateRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dateRange, err := ReadSimpleField[BACnetDateRangeEnclosed](ctx, "dateRange", ReadComplex[BACnetDateRangeEnclosed](BACnetDateRangeEnclosedParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dateRange' field"))
	}
	m.DateRange = dateRange

	if closeErr := readBuffer.CloseContext("BACnetCalendarEntryDateRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCalendarEntryDateRange")
	}

	return m, nil
}

func (m *_BACnetCalendarEntryDateRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetCalendarEntryDateRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetCalendarEntryDateRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetCalendarEntryDateRange")
		}

		if err := WriteSimpleField[BACnetDateRangeEnclosed](ctx, "dateRange", m.GetDateRange(), WriteComplex[BACnetDateRangeEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dateRange' field")
		}

		if popErr := writeBuffer.PopContext("BACnetCalendarEntryDateRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetCalendarEntryDateRange")
		}
		return nil
	}
	return m.BACnetCalendarEntryContract.(*_BACnetCalendarEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetCalendarEntryDateRange) IsBACnetCalendarEntryDateRange() {}

func (m *_BACnetCalendarEntryDateRange) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetCalendarEntryDateRange) deepCopy() *_BACnetCalendarEntryDateRange {
	if m == nil {
		return nil
	}
	_BACnetCalendarEntryDateRangeCopy := &_BACnetCalendarEntryDateRange{
		m.BACnetCalendarEntryContract.(*_BACnetCalendarEntry).deepCopy(),
		utils.DeepCopy[BACnetDateRangeEnclosed](m.DateRange),
	}
	_BACnetCalendarEntryDateRangeCopy.BACnetCalendarEntryContract.(*_BACnetCalendarEntry)._SubType = m
	return _BACnetCalendarEntryDateRangeCopy
}

func (m *_BACnetCalendarEntryDateRange) String() string {
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