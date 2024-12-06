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

// BACnetConstructedDataValueChangeTime is the corresponding interface of BACnetConstructedDataValueChangeTime
type BACnetConstructedDataValueChangeTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetValueChangeTime returns ValueChangeTime (property field)
	GetValueChangeTime() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
	// IsBACnetConstructedDataValueChangeTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataValueChangeTime()
	// CreateBuilder creates a BACnetConstructedDataValueChangeTimeBuilder
	CreateBACnetConstructedDataValueChangeTimeBuilder() BACnetConstructedDataValueChangeTimeBuilder
}

// _BACnetConstructedDataValueChangeTime is the data-structure of this message
type _BACnetConstructedDataValueChangeTime struct {
	BACnetConstructedDataContract
	ValueChangeTime BACnetDateTime
}

var _ BACnetConstructedDataValueChangeTime = (*_BACnetConstructedDataValueChangeTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataValueChangeTime)(nil)

// NewBACnetConstructedDataValueChangeTime factory function for _BACnetConstructedDataValueChangeTime
func NewBACnetConstructedDataValueChangeTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, valueChangeTime BACnetDateTime, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataValueChangeTime {
	if valueChangeTime == nil {
		panic("valueChangeTime of type BACnetDateTime for BACnetConstructedDataValueChangeTime must not be nil")
	}
	_result := &_BACnetConstructedDataValueChangeTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ValueChangeTime:               valueChangeTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataValueChangeTimeBuilder is a builder for BACnetConstructedDataValueChangeTime
type BACnetConstructedDataValueChangeTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(valueChangeTime BACnetDateTime) BACnetConstructedDataValueChangeTimeBuilder
	// WithValueChangeTime adds ValueChangeTime (property field)
	WithValueChangeTime(BACnetDateTime) BACnetConstructedDataValueChangeTimeBuilder
	// WithValueChangeTimeBuilder adds ValueChangeTime (property field) which is build by the builder
	WithValueChangeTimeBuilder(func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataValueChangeTimeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataValueChangeTime or returns an error if something is wrong
	Build() (BACnetConstructedDataValueChangeTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataValueChangeTime
}

// NewBACnetConstructedDataValueChangeTimeBuilder() creates a BACnetConstructedDataValueChangeTimeBuilder
func NewBACnetConstructedDataValueChangeTimeBuilder() BACnetConstructedDataValueChangeTimeBuilder {
	return &_BACnetConstructedDataValueChangeTimeBuilder{_BACnetConstructedDataValueChangeTime: new(_BACnetConstructedDataValueChangeTime)}
}

type _BACnetConstructedDataValueChangeTimeBuilder struct {
	*_BACnetConstructedDataValueChangeTime

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataValueChangeTimeBuilder) = (*_BACnetConstructedDataValueChangeTimeBuilder)(nil)

func (b *_BACnetConstructedDataValueChangeTimeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataValueChangeTime
}

func (b *_BACnetConstructedDataValueChangeTimeBuilder) WithMandatoryFields(valueChangeTime BACnetDateTime) BACnetConstructedDataValueChangeTimeBuilder {
	return b.WithValueChangeTime(valueChangeTime)
}

func (b *_BACnetConstructedDataValueChangeTimeBuilder) WithValueChangeTime(valueChangeTime BACnetDateTime) BACnetConstructedDataValueChangeTimeBuilder {
	b.ValueChangeTime = valueChangeTime
	return b
}

func (b *_BACnetConstructedDataValueChangeTimeBuilder) WithValueChangeTimeBuilder(builderSupplier func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataValueChangeTimeBuilder {
	builder := builderSupplier(b.ValueChangeTime.CreateBACnetDateTimeBuilder())
	var err error
	b.ValueChangeTime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDateTimeBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataValueChangeTimeBuilder) Build() (BACnetConstructedDataValueChangeTime, error) {
	if b.ValueChangeTime == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'valueChangeTime' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataValueChangeTime.deepCopy(), nil
}

func (b *_BACnetConstructedDataValueChangeTimeBuilder) MustBuild() BACnetConstructedDataValueChangeTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataValueChangeTimeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataValueChangeTimeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataValueChangeTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataValueChangeTimeBuilder().(*_BACnetConstructedDataValueChangeTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataValueChangeTimeBuilder creates a BACnetConstructedDataValueChangeTimeBuilder
func (b *_BACnetConstructedDataValueChangeTime) CreateBACnetConstructedDataValueChangeTimeBuilder() BACnetConstructedDataValueChangeTimeBuilder {
	if b == nil {
		return NewBACnetConstructedDataValueChangeTimeBuilder()
	}
	return &_BACnetConstructedDataValueChangeTimeBuilder{_BACnetConstructedDataValueChangeTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataValueChangeTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataValueChangeTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VALUE_CHANGE_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataValueChangeTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataValueChangeTime) GetValueChangeTime() BACnetDateTime {
	return m.ValueChangeTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataValueChangeTime) GetActualValue() BACnetDateTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateTime(m.GetValueChangeTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataValueChangeTime(structType any) BACnetConstructedDataValueChangeTime {
	if casted, ok := structType.(BACnetConstructedDataValueChangeTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataValueChangeTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataValueChangeTime) GetTypeName() string {
	return "BACnetConstructedDataValueChangeTime"
}

func (m *_BACnetConstructedDataValueChangeTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (valueChangeTime)
	lengthInBits += m.ValueChangeTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataValueChangeTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataValueChangeTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataValueChangeTime BACnetConstructedDataValueChangeTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataValueChangeTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataValueChangeTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	valueChangeTime, err := ReadSimpleField[BACnetDateTime](ctx, "valueChangeTime", ReadComplex[BACnetDateTime](BACnetDateTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueChangeTime' field"))
	}
	m.ValueChangeTime = valueChangeTime

	actualValue, err := ReadVirtualField[BACnetDateTime](ctx, "actualValue", (*BACnetDateTime)(nil), valueChangeTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataValueChangeTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataValueChangeTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataValueChangeTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataValueChangeTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataValueChangeTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataValueChangeTime")
		}

		if err := WriteSimpleField[BACnetDateTime](ctx, "valueChangeTime", m.GetValueChangeTime(), WriteComplex[BACnetDateTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'valueChangeTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataValueChangeTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataValueChangeTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataValueChangeTime) IsBACnetConstructedDataValueChangeTime() {}

func (m *_BACnetConstructedDataValueChangeTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataValueChangeTime) deepCopy() *_BACnetConstructedDataValueChangeTime {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataValueChangeTimeCopy := &_BACnetConstructedDataValueChangeTime{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetDateTime](m.ValueChangeTime),
	}
	_BACnetConstructedDataValueChangeTimeCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataValueChangeTimeCopy
}

func (m *_BACnetConstructedDataValueChangeTime) String() string {
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