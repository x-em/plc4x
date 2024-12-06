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

// BACnetConstructedDataStopTime is the corresponding interface of BACnetConstructedDataStopTime
type BACnetConstructedDataStopTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetStopTime returns StopTime (property field)
	GetStopTime() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
	// IsBACnetConstructedDataStopTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataStopTime()
	// CreateBuilder creates a BACnetConstructedDataStopTimeBuilder
	CreateBACnetConstructedDataStopTimeBuilder() BACnetConstructedDataStopTimeBuilder
}

// _BACnetConstructedDataStopTime is the data-structure of this message
type _BACnetConstructedDataStopTime struct {
	BACnetConstructedDataContract
	StopTime BACnetDateTime
}

var _ BACnetConstructedDataStopTime = (*_BACnetConstructedDataStopTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataStopTime)(nil)

// NewBACnetConstructedDataStopTime factory function for _BACnetConstructedDataStopTime
func NewBACnetConstructedDataStopTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, stopTime BACnetDateTime, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataStopTime {
	if stopTime == nil {
		panic("stopTime of type BACnetDateTime for BACnetConstructedDataStopTime must not be nil")
	}
	_result := &_BACnetConstructedDataStopTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		StopTime:                      stopTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataStopTimeBuilder is a builder for BACnetConstructedDataStopTime
type BACnetConstructedDataStopTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(stopTime BACnetDateTime) BACnetConstructedDataStopTimeBuilder
	// WithStopTime adds StopTime (property field)
	WithStopTime(BACnetDateTime) BACnetConstructedDataStopTimeBuilder
	// WithStopTimeBuilder adds StopTime (property field) which is build by the builder
	WithStopTimeBuilder(func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataStopTimeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataStopTime or returns an error if something is wrong
	Build() (BACnetConstructedDataStopTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataStopTime
}

// NewBACnetConstructedDataStopTimeBuilder() creates a BACnetConstructedDataStopTimeBuilder
func NewBACnetConstructedDataStopTimeBuilder() BACnetConstructedDataStopTimeBuilder {
	return &_BACnetConstructedDataStopTimeBuilder{_BACnetConstructedDataStopTime: new(_BACnetConstructedDataStopTime)}
}

type _BACnetConstructedDataStopTimeBuilder struct {
	*_BACnetConstructedDataStopTime

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataStopTimeBuilder) = (*_BACnetConstructedDataStopTimeBuilder)(nil)

func (b *_BACnetConstructedDataStopTimeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataStopTime
}

func (b *_BACnetConstructedDataStopTimeBuilder) WithMandatoryFields(stopTime BACnetDateTime) BACnetConstructedDataStopTimeBuilder {
	return b.WithStopTime(stopTime)
}

func (b *_BACnetConstructedDataStopTimeBuilder) WithStopTime(stopTime BACnetDateTime) BACnetConstructedDataStopTimeBuilder {
	b.StopTime = stopTime
	return b
}

func (b *_BACnetConstructedDataStopTimeBuilder) WithStopTimeBuilder(builderSupplier func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataStopTimeBuilder {
	builder := builderSupplier(b.StopTime.CreateBACnetDateTimeBuilder())
	var err error
	b.StopTime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDateTimeBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataStopTimeBuilder) Build() (BACnetConstructedDataStopTime, error) {
	if b.StopTime == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'stopTime' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataStopTime.deepCopy(), nil
}

func (b *_BACnetConstructedDataStopTimeBuilder) MustBuild() BACnetConstructedDataStopTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataStopTimeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataStopTimeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataStopTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataStopTimeBuilder().(*_BACnetConstructedDataStopTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataStopTimeBuilder creates a BACnetConstructedDataStopTimeBuilder
func (b *_BACnetConstructedDataStopTime) CreateBACnetConstructedDataStopTimeBuilder() BACnetConstructedDataStopTimeBuilder {
	if b == nil {
		return NewBACnetConstructedDataStopTimeBuilder()
	}
	return &_BACnetConstructedDataStopTimeBuilder{_BACnetConstructedDataStopTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataStopTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataStopTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_STOP_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataStopTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataStopTime) GetStopTime() BACnetDateTime {
	return m.StopTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataStopTime) GetActualValue() BACnetDateTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateTime(m.GetStopTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataStopTime(structType any) BACnetConstructedDataStopTime {
	if casted, ok := structType.(BACnetConstructedDataStopTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataStopTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataStopTime) GetTypeName() string {
	return "BACnetConstructedDataStopTime"
}

func (m *_BACnetConstructedDataStopTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (stopTime)
	lengthInBits += m.StopTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataStopTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataStopTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataStopTime BACnetConstructedDataStopTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataStopTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataStopTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	stopTime, err := ReadSimpleField[BACnetDateTime](ctx, "stopTime", ReadComplex[BACnetDateTime](BACnetDateTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'stopTime' field"))
	}
	m.StopTime = stopTime

	actualValue, err := ReadVirtualField[BACnetDateTime](ctx, "actualValue", (*BACnetDateTime)(nil), stopTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataStopTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataStopTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataStopTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataStopTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataStopTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataStopTime")
		}

		if err := WriteSimpleField[BACnetDateTime](ctx, "stopTime", m.GetStopTime(), WriteComplex[BACnetDateTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'stopTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataStopTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataStopTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataStopTime) IsBACnetConstructedDataStopTime() {}

func (m *_BACnetConstructedDataStopTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataStopTime) deepCopy() *_BACnetConstructedDataStopTime {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataStopTimeCopy := &_BACnetConstructedDataStopTime{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetDateTime](m.StopTime),
	}
	_BACnetConstructedDataStopTimeCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataStopTimeCopy
}

func (m *_BACnetConstructedDataStopTime) String() string {
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