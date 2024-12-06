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

// BACnetFaultParameterFaultExtendedParametersEntryTime is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryTime
type BACnetFaultParameterFaultExtendedParametersEntryTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameterFaultExtendedParametersEntry
	// GetTimeValue returns TimeValue (property field)
	GetTimeValue() BACnetApplicationTagTime
	// IsBACnetFaultParameterFaultExtendedParametersEntryTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultExtendedParametersEntryTime()
	// CreateBuilder creates a BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder
	CreateBACnetFaultParameterFaultExtendedParametersEntryTimeBuilder() BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder
}

// _BACnetFaultParameterFaultExtendedParametersEntryTime is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParametersEntryTime struct {
	BACnetFaultParameterFaultExtendedParametersEntryContract
	TimeValue BACnetApplicationTagTime
}

var _ BACnetFaultParameterFaultExtendedParametersEntryTime = (*_BACnetFaultParameterFaultExtendedParametersEntryTime)(nil)
var _ BACnetFaultParameterFaultExtendedParametersEntryRequirements = (*_BACnetFaultParameterFaultExtendedParametersEntryTime)(nil)

// NewBACnetFaultParameterFaultExtendedParametersEntryTime factory function for _BACnetFaultParameterFaultExtendedParametersEntryTime
func NewBACnetFaultParameterFaultExtendedParametersEntryTime(peekedTagHeader BACnetTagHeader, timeValue BACnetApplicationTagTime) *_BACnetFaultParameterFaultExtendedParametersEntryTime {
	if timeValue == nil {
		panic("timeValue of type BACnetApplicationTagTime for BACnetFaultParameterFaultExtendedParametersEntryTime must not be nil")
	}
	_result := &_BACnetFaultParameterFaultExtendedParametersEntryTime{
		BACnetFaultParameterFaultExtendedParametersEntryContract: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
		TimeValue: timeValue,
	}
	_result.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder is a builder for BACnetFaultParameterFaultExtendedParametersEntryTime
type BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(timeValue BACnetApplicationTagTime) BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder
	// WithTimeValue adds TimeValue (property field)
	WithTimeValue(BACnetApplicationTagTime) BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder
	// WithTimeValueBuilder adds TimeValue (property field) which is build by the builder
	WithTimeValueBuilder(func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetFaultParameterFaultExtendedParametersEntryBuilder
	// Build builds the BACnetFaultParameterFaultExtendedParametersEntryTime or returns an error if something is wrong
	Build() (BACnetFaultParameterFaultExtendedParametersEntryTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetFaultParameterFaultExtendedParametersEntryTime
}

// NewBACnetFaultParameterFaultExtendedParametersEntryTimeBuilder() creates a BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder
func NewBACnetFaultParameterFaultExtendedParametersEntryTimeBuilder() BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder {
	return &_BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder{_BACnetFaultParameterFaultExtendedParametersEntryTime: new(_BACnetFaultParameterFaultExtendedParametersEntryTime)}
}

type _BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder struct {
	*_BACnetFaultParameterFaultExtendedParametersEntryTime

	parentBuilder *_BACnetFaultParameterFaultExtendedParametersEntryBuilder

	err *utils.MultiError
}

var _ (BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder) = (*_BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder)(nil)

func (b *_BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder) setParent(contract BACnetFaultParameterFaultExtendedParametersEntryContract) {
	b.BACnetFaultParameterFaultExtendedParametersEntryContract = contract
	contract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = b._BACnetFaultParameterFaultExtendedParametersEntryTime
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder) WithMandatoryFields(timeValue BACnetApplicationTagTime) BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder {
	return b.WithTimeValue(timeValue)
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder) WithTimeValue(timeValue BACnetApplicationTagTime) BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder {
	b.TimeValue = timeValue
	return b
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder) WithTimeValueBuilder(builderSupplier func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder {
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

func (b *_BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder) Build() (BACnetFaultParameterFaultExtendedParametersEntryTime, error) {
	if b.TimeValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetFaultParameterFaultExtendedParametersEntryTime.deepCopy(), nil
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder) MustBuild() BACnetFaultParameterFaultExtendedParametersEntryTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder) Done() BACnetFaultParameterFaultExtendedParametersEntryBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetFaultParameterFaultExtendedParametersEntryBuilder().(*_BACnetFaultParameterFaultExtendedParametersEntryBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder) buildForBACnetFaultParameterFaultExtendedParametersEntry() (BACnetFaultParameterFaultExtendedParametersEntry, error) {
	return b.Build()
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetFaultParameterFaultExtendedParametersEntryTimeBuilder().(*_BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetFaultParameterFaultExtendedParametersEntryTimeBuilder creates a BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder
func (b *_BACnetFaultParameterFaultExtendedParametersEntryTime) CreateBACnetFaultParameterFaultExtendedParametersEntryTimeBuilder() BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder {
	if b == nil {
		return NewBACnetFaultParameterFaultExtendedParametersEntryTimeBuilder()
	}
	return &_BACnetFaultParameterFaultExtendedParametersEntryTimeBuilder{_BACnetFaultParameterFaultExtendedParametersEntryTime: b.deepCopy()}
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

func (m *_BACnetFaultParameterFaultExtendedParametersEntryTime) GetParent() BACnetFaultParameterFaultExtendedParametersEntryContract {
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryTime) GetTimeValue() BACnetApplicationTagTime {
	return m.TimeValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParametersEntryTime(structType any) BACnetFaultParameterFaultExtendedParametersEntryTime {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryTime) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryTime"
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).getLengthInBits(ctx))

	// Simple field (timeValue)
	lengthInBits += m.TimeValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultExtendedParametersEntry) (__bACnetFaultParameterFaultExtendedParametersEntryTime BACnetFaultParameterFaultExtendedParametersEntryTime, err error) {
	m.BACnetFaultParameterFaultExtendedParametersEntryContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntryTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeValue, err := ReadSimpleField[BACnetApplicationTagTime](ctx, "timeValue", ReadComplex[BACnetApplicationTagTime](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagTime](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeValue' field"))
	}
	m.TimeValue = timeValue

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntryTime")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntryTime")
		}

		if err := WriteSimpleField[BACnetApplicationTagTime](ctx, "timeValue", m.GetTimeValue(), WriteComplex[BACnetApplicationTagTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntryTime")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryTime) IsBACnetFaultParameterFaultExtendedParametersEntryTime() {
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryTime) deepCopy() *_BACnetFaultParameterFaultExtendedParametersEntryTime {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultExtendedParametersEntryTimeCopy := &_BACnetFaultParameterFaultExtendedParametersEntryTime{
		m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagTime](m.TimeValue),
	}
	_BACnetFaultParameterFaultExtendedParametersEntryTimeCopy.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = m
	return _BACnetFaultParameterFaultExtendedParametersEntryTimeCopy
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryTime) String() string {
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