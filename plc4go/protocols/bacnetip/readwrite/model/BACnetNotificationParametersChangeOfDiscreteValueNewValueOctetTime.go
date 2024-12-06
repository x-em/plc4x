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

// BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime
type BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetNotificationParametersChangeOfDiscreteValueNewValue
	// GetTimeValue returns TimeValue (property field)
	GetTimeValue() BACnetApplicationTagTime
	// IsBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime()
	// CreateBuilder creates a BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder
	CreateBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder() BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder
}

// _BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime is the data-structure of this message
type _BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime struct {
	BACnetNotificationParametersChangeOfDiscreteValueNewValueContract
	TimeValue BACnetApplicationTagTime
}

var _ BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime = (*_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime)(nil)
var _ BACnetNotificationParametersChangeOfDiscreteValueNewValueRequirements = (*_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime)(nil)

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime factory function for _BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, timeValue BACnetApplicationTagTime, tagNumber uint8) *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime {
	if timeValue == nil {
		panic("timeValue of type BACnetApplicationTagTime for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime must not be nil")
	}
	_result := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime{
		BACnetNotificationParametersChangeOfDiscreteValueNewValueContract: NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
		TimeValue: timeValue,
	}
	_result.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder is a builder for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime
type BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(timeValue BACnetApplicationTagTime) BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder
	// WithTimeValue adds TimeValue (property field)
	WithTimeValue(BACnetApplicationTagTime) BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder
	// WithTimeValueBuilder adds TimeValue (property field) which is build by the builder
	WithTimeValueBuilder(func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder
	// Build builds the BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime or returns an error if something is wrong
	Build() (BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime
}

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder() creates a BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder() BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder {
	return &_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder{_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime: new(_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime)}
}

type _BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder struct {
	*_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime

	parentBuilder *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder

	err *utils.MultiError
}

var _ (BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder) = (*_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder)(nil)

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder) setParent(contract BACnetNotificationParametersChangeOfDiscreteValueNewValueContract) {
	b.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract = contract
	contract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue)._SubType = b._BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder) WithMandatoryFields(timeValue BACnetApplicationTagTime) BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder {
	return b.WithTimeValue(timeValue)
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder) WithTimeValue(timeValue BACnetApplicationTagTime) BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder {
	b.TimeValue = timeValue
	return b
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder) WithTimeValueBuilder(builderSupplier func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder {
	builder := builderSupplier(b.TimeValue.CreateBACnetApplicationTagTimeBuilder())
	var err error
	b.TimeValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagTimeBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder) Build() (BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime, error) {
	if b.TimeValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime.deepCopy(), nil
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder) MustBuild() BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder) Done() BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder().(*_BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder) buildForBACnetNotificationParametersChangeOfDiscreteValueNewValue() (BACnetNotificationParametersChangeOfDiscreteValueNewValue, error) {
	return b.Build()
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder().(*_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder creates a BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder
func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) CreateBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder() BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder {
	if b == nil {
		return NewBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder()
	}
	return &_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeBuilder{_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime: b.deepCopy()}
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

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) GetParent() BACnetNotificationParametersChangeOfDiscreteValueNewValueContract {
	return m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) GetTimeValue() BACnetApplicationTagTime {
	return m.TimeValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime(structType any) BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime"
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue).getLengthInBits(ctx))

	// Simple field (timeValue)
	lengthInBits += m.TimeValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParametersChangeOfDiscreteValueNewValue, tagNumber uint8) (__bACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime, err error) {
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeValue, err := ReadSimpleField[BACnetApplicationTagTime](ctx, "timeValue", ReadComplex[BACnetApplicationTagTime](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagTime](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeValue' field"))
	}
	m.TimeValue = timeValue

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime")
		}

		if err := WriteSimpleField[BACnetApplicationTagTime](ctx, "timeValue", m.GetTimeValue(), WriteComplex[BACnetApplicationTagTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime")
		}
		return nil
	}
	return m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) IsBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime() {
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) deepCopy() *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime {
	if m == nil {
		return nil
	}
	_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeCopy := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime{
		m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagTime](m.TimeValue),
	}
	_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeCopy.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue)._SubType = m
	return _BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeCopy
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) String() string {
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
