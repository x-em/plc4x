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

// BACnetEventParameterSignedOutOfRange is the corresponding interface of BACnetEventParameterSignedOutOfRange
type BACnetEventParameterSignedOutOfRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetLowLimit returns LowLimit (property field)
	GetLowLimit() BACnetContextTagSignedInteger
	// GetHighLimit returns HighLimit (property field)
	GetHighLimit() BACnetContextTagSignedInteger
	// GetDeadband returns Deadband (property field)
	GetDeadband() BACnetContextTagUnsignedInteger
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventParameterSignedOutOfRange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterSignedOutOfRange()
	// CreateBuilder creates a BACnetEventParameterSignedOutOfRangeBuilder
	CreateBACnetEventParameterSignedOutOfRangeBuilder() BACnetEventParameterSignedOutOfRangeBuilder
}

// _BACnetEventParameterSignedOutOfRange is the data-structure of this message
type _BACnetEventParameterSignedOutOfRange struct {
	BACnetEventParameterContract
	OpeningTag BACnetOpeningTag
	TimeDelay  BACnetContextTagUnsignedInteger
	LowLimit   BACnetContextTagSignedInteger
	HighLimit  BACnetContextTagSignedInteger
	Deadband   BACnetContextTagUnsignedInteger
	ClosingTag BACnetClosingTag
}

var _ BACnetEventParameterSignedOutOfRange = (*_BACnetEventParameterSignedOutOfRange)(nil)
var _ BACnetEventParameterRequirements = (*_BACnetEventParameterSignedOutOfRange)(nil)

// NewBACnetEventParameterSignedOutOfRange factory function for _BACnetEventParameterSignedOutOfRange
func NewBACnetEventParameterSignedOutOfRange(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, lowLimit BACnetContextTagSignedInteger, highLimit BACnetContextTagSignedInteger, deadband BACnetContextTagUnsignedInteger, closingTag BACnetClosingTag) *_BACnetEventParameterSignedOutOfRange {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventParameterSignedOutOfRange must not be nil")
	}
	if timeDelay == nil {
		panic("timeDelay of type BACnetContextTagUnsignedInteger for BACnetEventParameterSignedOutOfRange must not be nil")
	}
	if lowLimit == nil {
		panic("lowLimit of type BACnetContextTagSignedInteger for BACnetEventParameterSignedOutOfRange must not be nil")
	}
	if highLimit == nil {
		panic("highLimit of type BACnetContextTagSignedInteger for BACnetEventParameterSignedOutOfRange must not be nil")
	}
	if deadband == nil {
		panic("deadband of type BACnetContextTagUnsignedInteger for BACnetEventParameterSignedOutOfRange must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventParameterSignedOutOfRange must not be nil")
	}
	_result := &_BACnetEventParameterSignedOutOfRange{
		BACnetEventParameterContract: NewBACnetEventParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		TimeDelay:                    timeDelay,
		LowLimit:                     lowLimit,
		HighLimit:                    highLimit,
		Deadband:                     deadband,
		ClosingTag:                   closingTag,
	}
	_result.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventParameterSignedOutOfRangeBuilder is a builder for BACnetEventParameterSignedOutOfRange
type BACnetEventParameterSignedOutOfRangeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, lowLimit BACnetContextTagSignedInteger, highLimit BACnetContextTagSignedInteger, deadband BACnetContextTagUnsignedInteger, closingTag BACnetClosingTag) BACnetEventParameterSignedOutOfRangeBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetEventParameterSignedOutOfRangeBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterSignedOutOfRangeBuilder
	// WithTimeDelay adds TimeDelay (property field)
	WithTimeDelay(BACnetContextTagUnsignedInteger) BACnetEventParameterSignedOutOfRangeBuilder
	// WithTimeDelayBuilder adds TimeDelay (property field) which is build by the builder
	WithTimeDelayBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterSignedOutOfRangeBuilder
	// WithLowLimit adds LowLimit (property field)
	WithLowLimit(BACnetContextTagSignedInteger) BACnetEventParameterSignedOutOfRangeBuilder
	// WithLowLimitBuilder adds LowLimit (property field) which is build by the builder
	WithLowLimitBuilder(func(BACnetContextTagSignedIntegerBuilder) BACnetContextTagSignedIntegerBuilder) BACnetEventParameterSignedOutOfRangeBuilder
	// WithHighLimit adds HighLimit (property field)
	WithHighLimit(BACnetContextTagSignedInteger) BACnetEventParameterSignedOutOfRangeBuilder
	// WithHighLimitBuilder adds HighLimit (property field) which is build by the builder
	WithHighLimitBuilder(func(BACnetContextTagSignedIntegerBuilder) BACnetContextTagSignedIntegerBuilder) BACnetEventParameterSignedOutOfRangeBuilder
	// WithDeadband adds Deadband (property field)
	WithDeadband(BACnetContextTagUnsignedInteger) BACnetEventParameterSignedOutOfRangeBuilder
	// WithDeadbandBuilder adds Deadband (property field) which is build by the builder
	WithDeadbandBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterSignedOutOfRangeBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetEventParameterSignedOutOfRangeBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterSignedOutOfRangeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetEventParameterBuilder
	// Build builds the BACnetEventParameterSignedOutOfRange or returns an error if something is wrong
	Build() (BACnetEventParameterSignedOutOfRange, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventParameterSignedOutOfRange
}

// NewBACnetEventParameterSignedOutOfRangeBuilder() creates a BACnetEventParameterSignedOutOfRangeBuilder
func NewBACnetEventParameterSignedOutOfRangeBuilder() BACnetEventParameterSignedOutOfRangeBuilder {
	return &_BACnetEventParameterSignedOutOfRangeBuilder{_BACnetEventParameterSignedOutOfRange: new(_BACnetEventParameterSignedOutOfRange)}
}

type _BACnetEventParameterSignedOutOfRangeBuilder struct {
	*_BACnetEventParameterSignedOutOfRange

	parentBuilder *_BACnetEventParameterBuilder

	err *utils.MultiError
}

var _ (BACnetEventParameterSignedOutOfRangeBuilder) = (*_BACnetEventParameterSignedOutOfRangeBuilder)(nil)

func (b *_BACnetEventParameterSignedOutOfRangeBuilder) setParent(contract BACnetEventParameterContract) {
	b.BACnetEventParameterContract = contract
	contract.(*_BACnetEventParameter)._SubType = b._BACnetEventParameterSignedOutOfRange
}

func (b *_BACnetEventParameterSignedOutOfRangeBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, lowLimit BACnetContextTagSignedInteger, highLimit BACnetContextTagSignedInteger, deadband BACnetContextTagUnsignedInteger, closingTag BACnetClosingTag) BACnetEventParameterSignedOutOfRangeBuilder {
	return b.WithOpeningTag(openingTag).WithTimeDelay(timeDelay).WithLowLimit(lowLimit).WithHighLimit(highLimit).WithDeadband(deadband).WithClosingTag(closingTag)
}

func (b *_BACnetEventParameterSignedOutOfRangeBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetEventParameterSignedOutOfRangeBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetEventParameterSignedOutOfRangeBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterSignedOutOfRangeBuilder {
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

func (b *_BACnetEventParameterSignedOutOfRangeBuilder) WithTimeDelay(timeDelay BACnetContextTagUnsignedInteger) BACnetEventParameterSignedOutOfRangeBuilder {
	b.TimeDelay = timeDelay
	return b
}

func (b *_BACnetEventParameterSignedOutOfRangeBuilder) WithTimeDelayBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterSignedOutOfRangeBuilder {
	builder := builderSupplier(b.TimeDelay.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.TimeDelay, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterSignedOutOfRangeBuilder) WithLowLimit(lowLimit BACnetContextTagSignedInteger) BACnetEventParameterSignedOutOfRangeBuilder {
	b.LowLimit = lowLimit
	return b
}

func (b *_BACnetEventParameterSignedOutOfRangeBuilder) WithLowLimitBuilder(builderSupplier func(BACnetContextTagSignedIntegerBuilder) BACnetContextTagSignedIntegerBuilder) BACnetEventParameterSignedOutOfRangeBuilder {
	builder := builderSupplier(b.LowLimit.CreateBACnetContextTagSignedIntegerBuilder())
	var err error
	b.LowLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterSignedOutOfRangeBuilder) WithHighLimit(highLimit BACnetContextTagSignedInteger) BACnetEventParameterSignedOutOfRangeBuilder {
	b.HighLimit = highLimit
	return b
}

func (b *_BACnetEventParameterSignedOutOfRangeBuilder) WithHighLimitBuilder(builderSupplier func(BACnetContextTagSignedIntegerBuilder) BACnetContextTagSignedIntegerBuilder) BACnetEventParameterSignedOutOfRangeBuilder {
	builder := builderSupplier(b.HighLimit.CreateBACnetContextTagSignedIntegerBuilder())
	var err error
	b.HighLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterSignedOutOfRangeBuilder) WithDeadband(deadband BACnetContextTagUnsignedInteger) BACnetEventParameterSignedOutOfRangeBuilder {
	b.Deadband = deadband
	return b
}

func (b *_BACnetEventParameterSignedOutOfRangeBuilder) WithDeadbandBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterSignedOutOfRangeBuilder {
	builder := builderSupplier(b.Deadband.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.Deadband, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterSignedOutOfRangeBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetEventParameterSignedOutOfRangeBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetEventParameterSignedOutOfRangeBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterSignedOutOfRangeBuilder {
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

func (b *_BACnetEventParameterSignedOutOfRangeBuilder) Build() (BACnetEventParameterSignedOutOfRange, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.TimeDelay == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeDelay' not set"))
	}
	if b.LowLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lowLimit' not set"))
	}
	if b.HighLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'highLimit' not set"))
	}
	if b.Deadband == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'deadband' not set"))
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
	return b._BACnetEventParameterSignedOutOfRange.deepCopy(), nil
}

func (b *_BACnetEventParameterSignedOutOfRangeBuilder) MustBuild() BACnetEventParameterSignedOutOfRange {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetEventParameterSignedOutOfRangeBuilder) Done() BACnetEventParameterBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetEventParameterBuilder().(*_BACnetEventParameterBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetEventParameterSignedOutOfRangeBuilder) buildForBACnetEventParameter() (BACnetEventParameter, error) {
	return b.Build()
}

func (b *_BACnetEventParameterSignedOutOfRangeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetEventParameterSignedOutOfRangeBuilder().(*_BACnetEventParameterSignedOutOfRangeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetEventParameterSignedOutOfRangeBuilder creates a BACnetEventParameterSignedOutOfRangeBuilder
func (b *_BACnetEventParameterSignedOutOfRange) CreateBACnetEventParameterSignedOutOfRangeBuilder() BACnetEventParameterSignedOutOfRangeBuilder {
	if b == nil {
		return NewBACnetEventParameterSignedOutOfRangeBuilder()
	}
	return &_BACnetEventParameterSignedOutOfRangeBuilder{_BACnetEventParameterSignedOutOfRange: b.deepCopy()}
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

func (m *_BACnetEventParameterSignedOutOfRange) GetParent() BACnetEventParameterContract {
	return m.BACnetEventParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterSignedOutOfRange) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterSignedOutOfRange) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterSignedOutOfRange) GetLowLimit() BACnetContextTagSignedInteger {
	return m.LowLimit
}

func (m *_BACnetEventParameterSignedOutOfRange) GetHighLimit() BACnetContextTagSignedInteger {
	return m.HighLimit
}

func (m *_BACnetEventParameterSignedOutOfRange) GetDeadband() BACnetContextTagUnsignedInteger {
	return m.Deadband
}

func (m *_BACnetEventParameterSignedOutOfRange) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterSignedOutOfRange(structType any) BACnetEventParameterSignedOutOfRange {
	if casted, ok := structType.(BACnetEventParameterSignedOutOfRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterSignedOutOfRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterSignedOutOfRange) GetTypeName() string {
	return "BACnetEventParameterSignedOutOfRange"
}

func (m *_BACnetEventParameterSignedOutOfRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventParameterContract.(*_BACnetEventParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits(ctx)

	// Simple field (lowLimit)
	lengthInBits += m.LowLimit.GetLengthInBits(ctx)

	// Simple field (highLimit)
	lengthInBits += m.HighLimit.GetLengthInBits(ctx)

	// Simple field (deadband)
	lengthInBits += m.Deadband.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterSignedOutOfRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventParameterSignedOutOfRange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventParameter) (__bACnetEventParameterSignedOutOfRange BACnetEventParameterSignedOutOfRange, err error) {
	m.BACnetEventParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterSignedOutOfRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterSignedOutOfRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(15))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	timeDelay, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeDelay' field"))
	}
	m.TimeDelay = timeDelay

	lowLimit, err := ReadSimpleField[BACnetContextTagSignedInteger](ctx, "lowLimit", ReadComplex[BACnetContextTagSignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagSignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_SIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lowLimit' field"))
	}
	m.LowLimit = lowLimit

	highLimit, err := ReadSimpleField[BACnetContextTagSignedInteger](ctx, "highLimit", ReadComplex[BACnetContextTagSignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagSignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_SIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'highLimit' field"))
	}
	m.HighLimit = highLimit

	deadband, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "deadband", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deadband' field"))
	}
	m.Deadband = deadband

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(15))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterSignedOutOfRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterSignedOutOfRange")
	}

	return m, nil
}

func (m *_BACnetEventParameterSignedOutOfRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterSignedOutOfRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterSignedOutOfRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterSignedOutOfRange")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", m.GetTimeDelay(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeDelay' field")
		}

		if err := WriteSimpleField[BACnetContextTagSignedInteger](ctx, "lowLimit", m.GetLowLimit(), WriteComplex[BACnetContextTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lowLimit' field")
		}

		if err := WriteSimpleField[BACnetContextTagSignedInteger](ctx, "highLimit", m.GetHighLimit(), WriteComplex[BACnetContextTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'highLimit' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "deadband", m.GetDeadband(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deadband' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterSignedOutOfRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterSignedOutOfRange")
		}
		return nil
	}
	return m.BACnetEventParameterContract.(*_BACnetEventParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterSignedOutOfRange) IsBACnetEventParameterSignedOutOfRange() {}

func (m *_BACnetEventParameterSignedOutOfRange) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventParameterSignedOutOfRange) deepCopy() *_BACnetEventParameterSignedOutOfRange {
	if m == nil {
		return nil
	}
	_BACnetEventParameterSignedOutOfRangeCopy := &_BACnetEventParameterSignedOutOfRange{
		m.BACnetEventParameterContract.(*_BACnetEventParameter).deepCopy(),
		utils.DeepCopy[BACnetOpeningTag](m.OpeningTag),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.TimeDelay),
		utils.DeepCopy[BACnetContextTagSignedInteger](m.LowLimit),
		utils.DeepCopy[BACnetContextTagSignedInteger](m.HighLimit),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.Deadband),
		utils.DeepCopy[BACnetClosingTag](m.ClosingTag),
	}
	_BACnetEventParameterSignedOutOfRangeCopy.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = m
	return _BACnetEventParameterSignedOutOfRangeCopy
}

func (m *_BACnetEventParameterSignedOutOfRange) String() string {
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