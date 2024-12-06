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

// BACnetConstructedDataLastCommandTime is the corresponding interface of BACnetConstructedDataLastCommandTime
type BACnetConstructedDataLastCommandTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLastCommandTime returns LastCommandTime (property field)
	GetLastCommandTime() BACnetTimeStamp
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetTimeStamp
	// IsBACnetConstructedDataLastCommandTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLastCommandTime()
	// CreateBuilder creates a BACnetConstructedDataLastCommandTimeBuilder
	CreateBACnetConstructedDataLastCommandTimeBuilder() BACnetConstructedDataLastCommandTimeBuilder
}

// _BACnetConstructedDataLastCommandTime is the data-structure of this message
type _BACnetConstructedDataLastCommandTime struct {
	BACnetConstructedDataContract
	LastCommandTime BACnetTimeStamp
}

var _ BACnetConstructedDataLastCommandTime = (*_BACnetConstructedDataLastCommandTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLastCommandTime)(nil)

// NewBACnetConstructedDataLastCommandTime factory function for _BACnetConstructedDataLastCommandTime
func NewBACnetConstructedDataLastCommandTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lastCommandTime BACnetTimeStamp, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastCommandTime {
	if lastCommandTime == nil {
		panic("lastCommandTime of type BACnetTimeStamp for BACnetConstructedDataLastCommandTime must not be nil")
	}
	_result := &_BACnetConstructedDataLastCommandTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LastCommandTime:               lastCommandTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLastCommandTimeBuilder is a builder for BACnetConstructedDataLastCommandTime
type BACnetConstructedDataLastCommandTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lastCommandTime BACnetTimeStamp) BACnetConstructedDataLastCommandTimeBuilder
	// WithLastCommandTime adds LastCommandTime (property field)
	WithLastCommandTime(BACnetTimeStamp) BACnetConstructedDataLastCommandTimeBuilder
	// WithLastCommandTimeBuilder adds LastCommandTime (property field) which is build by the builder
	WithLastCommandTimeBuilder(func(BACnetTimeStampBuilder) BACnetTimeStampBuilder) BACnetConstructedDataLastCommandTimeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLastCommandTime or returns an error if something is wrong
	Build() (BACnetConstructedDataLastCommandTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLastCommandTime
}

// NewBACnetConstructedDataLastCommandTimeBuilder() creates a BACnetConstructedDataLastCommandTimeBuilder
func NewBACnetConstructedDataLastCommandTimeBuilder() BACnetConstructedDataLastCommandTimeBuilder {
	return &_BACnetConstructedDataLastCommandTimeBuilder{_BACnetConstructedDataLastCommandTime: new(_BACnetConstructedDataLastCommandTime)}
}

type _BACnetConstructedDataLastCommandTimeBuilder struct {
	*_BACnetConstructedDataLastCommandTime

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLastCommandTimeBuilder) = (*_BACnetConstructedDataLastCommandTimeBuilder)(nil)

func (b *_BACnetConstructedDataLastCommandTimeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataLastCommandTime
}

func (b *_BACnetConstructedDataLastCommandTimeBuilder) WithMandatoryFields(lastCommandTime BACnetTimeStamp) BACnetConstructedDataLastCommandTimeBuilder {
	return b.WithLastCommandTime(lastCommandTime)
}

func (b *_BACnetConstructedDataLastCommandTimeBuilder) WithLastCommandTime(lastCommandTime BACnetTimeStamp) BACnetConstructedDataLastCommandTimeBuilder {
	b.LastCommandTime = lastCommandTime
	return b
}

func (b *_BACnetConstructedDataLastCommandTimeBuilder) WithLastCommandTimeBuilder(builderSupplier func(BACnetTimeStampBuilder) BACnetTimeStampBuilder) BACnetConstructedDataLastCommandTimeBuilder {
	builder := builderSupplier(b.LastCommandTime.CreateBACnetTimeStampBuilder())
	var err error
	b.LastCommandTime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTimeStampBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLastCommandTimeBuilder) Build() (BACnetConstructedDataLastCommandTime, error) {
	if b.LastCommandTime == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lastCommandTime' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLastCommandTime.deepCopy(), nil
}

func (b *_BACnetConstructedDataLastCommandTimeBuilder) MustBuild() BACnetConstructedDataLastCommandTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLastCommandTimeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLastCommandTimeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLastCommandTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLastCommandTimeBuilder().(*_BACnetConstructedDataLastCommandTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLastCommandTimeBuilder creates a BACnetConstructedDataLastCommandTimeBuilder
func (b *_BACnetConstructedDataLastCommandTime) CreateBACnetConstructedDataLastCommandTimeBuilder() BACnetConstructedDataLastCommandTimeBuilder {
	if b == nil {
		return NewBACnetConstructedDataLastCommandTimeBuilder()
	}
	return &_BACnetConstructedDataLastCommandTimeBuilder{_BACnetConstructedDataLastCommandTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastCommandTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastCommandTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_COMMAND_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastCommandTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastCommandTime) GetLastCommandTime() BACnetTimeStamp {
	return m.LastCommandTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastCommandTime) GetActualValue() BACnetTimeStamp {
	ctx := context.Background()
	_ = ctx
	return CastBACnetTimeStamp(m.GetLastCommandTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastCommandTime(structType any) BACnetConstructedDataLastCommandTime {
	if casted, ok := structType.(BACnetConstructedDataLastCommandTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastCommandTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastCommandTime) GetTypeName() string {
	return "BACnetConstructedDataLastCommandTime"
}

func (m *_BACnetConstructedDataLastCommandTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lastCommandTime)
	lengthInBits += m.LastCommandTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastCommandTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLastCommandTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLastCommandTime BACnetConstructedDataLastCommandTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastCommandTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastCommandTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lastCommandTime, err := ReadSimpleField[BACnetTimeStamp](ctx, "lastCommandTime", ReadComplex[BACnetTimeStamp](BACnetTimeStampParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastCommandTime' field"))
	}
	m.LastCommandTime = lastCommandTime

	actualValue, err := ReadVirtualField[BACnetTimeStamp](ctx, "actualValue", (*BACnetTimeStamp)(nil), lastCommandTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastCommandTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastCommandTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLastCommandTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastCommandTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastCommandTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastCommandTime")
		}

		if err := WriteSimpleField[BACnetTimeStamp](ctx, "lastCommandTime", m.GetLastCommandTime(), WriteComplex[BACnetTimeStamp](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastCommandTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastCommandTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastCommandTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastCommandTime) IsBACnetConstructedDataLastCommandTime() {}

func (m *_BACnetConstructedDataLastCommandTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLastCommandTime) deepCopy() *_BACnetConstructedDataLastCommandTime {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLastCommandTimeCopy := &_BACnetConstructedDataLastCommandTime{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetTimeStamp](m.LastCommandTime),
	}
	_BACnetConstructedDataLastCommandTimeCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLastCommandTimeCopy
}

func (m *_BACnetConstructedDataLastCommandTime) String() string {
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