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

// BACnetConstructedDataLastRestoreTime is the corresponding interface of BACnetConstructedDataLastRestoreTime
type BACnetConstructedDataLastRestoreTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLastRestoreTime returns LastRestoreTime (property field)
	GetLastRestoreTime() BACnetTimeStamp
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetTimeStamp
	// IsBACnetConstructedDataLastRestoreTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLastRestoreTime()
	// CreateBuilder creates a BACnetConstructedDataLastRestoreTimeBuilder
	CreateBACnetConstructedDataLastRestoreTimeBuilder() BACnetConstructedDataLastRestoreTimeBuilder
}

// _BACnetConstructedDataLastRestoreTime is the data-structure of this message
type _BACnetConstructedDataLastRestoreTime struct {
	BACnetConstructedDataContract
	LastRestoreTime BACnetTimeStamp
}

var _ BACnetConstructedDataLastRestoreTime = (*_BACnetConstructedDataLastRestoreTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLastRestoreTime)(nil)

// NewBACnetConstructedDataLastRestoreTime factory function for _BACnetConstructedDataLastRestoreTime
func NewBACnetConstructedDataLastRestoreTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lastRestoreTime BACnetTimeStamp, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastRestoreTime {
	if lastRestoreTime == nil {
		panic("lastRestoreTime of type BACnetTimeStamp for BACnetConstructedDataLastRestoreTime must not be nil")
	}
	_result := &_BACnetConstructedDataLastRestoreTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LastRestoreTime:               lastRestoreTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLastRestoreTimeBuilder is a builder for BACnetConstructedDataLastRestoreTime
type BACnetConstructedDataLastRestoreTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lastRestoreTime BACnetTimeStamp) BACnetConstructedDataLastRestoreTimeBuilder
	// WithLastRestoreTime adds LastRestoreTime (property field)
	WithLastRestoreTime(BACnetTimeStamp) BACnetConstructedDataLastRestoreTimeBuilder
	// WithLastRestoreTimeBuilder adds LastRestoreTime (property field) which is build by the builder
	WithLastRestoreTimeBuilder(func(BACnetTimeStampBuilder) BACnetTimeStampBuilder) BACnetConstructedDataLastRestoreTimeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLastRestoreTime or returns an error if something is wrong
	Build() (BACnetConstructedDataLastRestoreTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLastRestoreTime
}

// NewBACnetConstructedDataLastRestoreTimeBuilder() creates a BACnetConstructedDataLastRestoreTimeBuilder
func NewBACnetConstructedDataLastRestoreTimeBuilder() BACnetConstructedDataLastRestoreTimeBuilder {
	return &_BACnetConstructedDataLastRestoreTimeBuilder{_BACnetConstructedDataLastRestoreTime: new(_BACnetConstructedDataLastRestoreTime)}
}

type _BACnetConstructedDataLastRestoreTimeBuilder struct {
	*_BACnetConstructedDataLastRestoreTime

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLastRestoreTimeBuilder) = (*_BACnetConstructedDataLastRestoreTimeBuilder)(nil)

func (b *_BACnetConstructedDataLastRestoreTimeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataLastRestoreTime
}

func (b *_BACnetConstructedDataLastRestoreTimeBuilder) WithMandatoryFields(lastRestoreTime BACnetTimeStamp) BACnetConstructedDataLastRestoreTimeBuilder {
	return b.WithLastRestoreTime(lastRestoreTime)
}

func (b *_BACnetConstructedDataLastRestoreTimeBuilder) WithLastRestoreTime(lastRestoreTime BACnetTimeStamp) BACnetConstructedDataLastRestoreTimeBuilder {
	b.LastRestoreTime = lastRestoreTime
	return b
}

func (b *_BACnetConstructedDataLastRestoreTimeBuilder) WithLastRestoreTimeBuilder(builderSupplier func(BACnetTimeStampBuilder) BACnetTimeStampBuilder) BACnetConstructedDataLastRestoreTimeBuilder {
	builder := builderSupplier(b.LastRestoreTime.CreateBACnetTimeStampBuilder())
	var err error
	b.LastRestoreTime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTimeStampBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLastRestoreTimeBuilder) Build() (BACnetConstructedDataLastRestoreTime, error) {
	if b.LastRestoreTime == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lastRestoreTime' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLastRestoreTime.deepCopy(), nil
}

func (b *_BACnetConstructedDataLastRestoreTimeBuilder) MustBuild() BACnetConstructedDataLastRestoreTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLastRestoreTimeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLastRestoreTimeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLastRestoreTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLastRestoreTimeBuilder().(*_BACnetConstructedDataLastRestoreTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLastRestoreTimeBuilder creates a BACnetConstructedDataLastRestoreTimeBuilder
func (b *_BACnetConstructedDataLastRestoreTime) CreateBACnetConstructedDataLastRestoreTimeBuilder() BACnetConstructedDataLastRestoreTimeBuilder {
	if b == nil {
		return NewBACnetConstructedDataLastRestoreTimeBuilder()
	}
	return &_BACnetConstructedDataLastRestoreTimeBuilder{_BACnetConstructedDataLastRestoreTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastRestoreTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastRestoreTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_RESTORE_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastRestoreTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastRestoreTime) GetLastRestoreTime() BACnetTimeStamp {
	return m.LastRestoreTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastRestoreTime) GetActualValue() BACnetTimeStamp {
	ctx := context.Background()
	_ = ctx
	return CastBACnetTimeStamp(m.GetLastRestoreTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastRestoreTime(structType any) BACnetConstructedDataLastRestoreTime {
	if casted, ok := structType.(BACnetConstructedDataLastRestoreTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastRestoreTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastRestoreTime) GetTypeName() string {
	return "BACnetConstructedDataLastRestoreTime"
}

func (m *_BACnetConstructedDataLastRestoreTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lastRestoreTime)
	lengthInBits += m.LastRestoreTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastRestoreTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLastRestoreTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLastRestoreTime BACnetConstructedDataLastRestoreTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastRestoreTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastRestoreTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lastRestoreTime, err := ReadSimpleField[BACnetTimeStamp](ctx, "lastRestoreTime", ReadComplex[BACnetTimeStamp](BACnetTimeStampParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastRestoreTime' field"))
	}
	m.LastRestoreTime = lastRestoreTime

	actualValue, err := ReadVirtualField[BACnetTimeStamp](ctx, "actualValue", (*BACnetTimeStamp)(nil), lastRestoreTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastRestoreTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastRestoreTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLastRestoreTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastRestoreTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastRestoreTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastRestoreTime")
		}

		if err := WriteSimpleField[BACnetTimeStamp](ctx, "lastRestoreTime", m.GetLastRestoreTime(), WriteComplex[BACnetTimeStamp](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastRestoreTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastRestoreTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastRestoreTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastRestoreTime) IsBACnetConstructedDataLastRestoreTime() {}

func (m *_BACnetConstructedDataLastRestoreTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLastRestoreTime) deepCopy() *_BACnetConstructedDataLastRestoreTime {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLastRestoreTimeCopy := &_BACnetConstructedDataLastRestoreTime{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetTimeStamp](m.LastRestoreTime),
	}
	_BACnetConstructedDataLastRestoreTimeCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLastRestoreTimeCopy
}

func (m *_BACnetConstructedDataLastRestoreTime) String() string {
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
