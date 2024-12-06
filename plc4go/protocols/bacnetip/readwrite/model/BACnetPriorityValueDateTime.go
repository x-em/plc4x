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

// BACnetPriorityValueDateTime is the corresponding interface of BACnetPriorityValueDateTime
type BACnetPriorityValueDateTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPriorityValue
	// GetDateTimeValue returns DateTimeValue (property field)
	GetDateTimeValue() BACnetDateTimeEnclosed
	// IsBACnetPriorityValueDateTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPriorityValueDateTime()
	// CreateBuilder creates a BACnetPriorityValueDateTimeBuilder
	CreateBACnetPriorityValueDateTimeBuilder() BACnetPriorityValueDateTimeBuilder
}

// _BACnetPriorityValueDateTime is the data-structure of this message
type _BACnetPriorityValueDateTime struct {
	BACnetPriorityValueContract
	DateTimeValue BACnetDateTimeEnclosed
}

var _ BACnetPriorityValueDateTime = (*_BACnetPriorityValueDateTime)(nil)
var _ BACnetPriorityValueRequirements = (*_BACnetPriorityValueDateTime)(nil)

// NewBACnetPriorityValueDateTime factory function for _BACnetPriorityValueDateTime
func NewBACnetPriorityValueDateTime(peekedTagHeader BACnetTagHeader, dateTimeValue BACnetDateTimeEnclosed, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueDateTime {
	if dateTimeValue == nil {
		panic("dateTimeValue of type BACnetDateTimeEnclosed for BACnetPriorityValueDateTime must not be nil")
	}
	_result := &_BACnetPriorityValueDateTime{
		BACnetPriorityValueContract: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
		DateTimeValue:               dateTimeValue,
	}
	_result.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPriorityValueDateTimeBuilder is a builder for BACnetPriorityValueDateTime
type BACnetPriorityValueDateTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(dateTimeValue BACnetDateTimeEnclosed) BACnetPriorityValueDateTimeBuilder
	// WithDateTimeValue adds DateTimeValue (property field)
	WithDateTimeValue(BACnetDateTimeEnclosed) BACnetPriorityValueDateTimeBuilder
	// WithDateTimeValueBuilder adds DateTimeValue (property field) which is build by the builder
	WithDateTimeValueBuilder(func(BACnetDateTimeEnclosedBuilder) BACnetDateTimeEnclosedBuilder) BACnetPriorityValueDateTimeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetPriorityValueBuilder
	// Build builds the BACnetPriorityValueDateTime or returns an error if something is wrong
	Build() (BACnetPriorityValueDateTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPriorityValueDateTime
}

// NewBACnetPriorityValueDateTimeBuilder() creates a BACnetPriorityValueDateTimeBuilder
func NewBACnetPriorityValueDateTimeBuilder() BACnetPriorityValueDateTimeBuilder {
	return &_BACnetPriorityValueDateTimeBuilder{_BACnetPriorityValueDateTime: new(_BACnetPriorityValueDateTime)}
}

type _BACnetPriorityValueDateTimeBuilder struct {
	*_BACnetPriorityValueDateTime

	parentBuilder *_BACnetPriorityValueBuilder

	err *utils.MultiError
}

var _ (BACnetPriorityValueDateTimeBuilder) = (*_BACnetPriorityValueDateTimeBuilder)(nil)

func (b *_BACnetPriorityValueDateTimeBuilder) setParent(contract BACnetPriorityValueContract) {
	b.BACnetPriorityValueContract = contract
	contract.(*_BACnetPriorityValue)._SubType = b._BACnetPriorityValueDateTime
}

func (b *_BACnetPriorityValueDateTimeBuilder) WithMandatoryFields(dateTimeValue BACnetDateTimeEnclosed) BACnetPriorityValueDateTimeBuilder {
	return b.WithDateTimeValue(dateTimeValue)
}

func (b *_BACnetPriorityValueDateTimeBuilder) WithDateTimeValue(dateTimeValue BACnetDateTimeEnclosed) BACnetPriorityValueDateTimeBuilder {
	b.DateTimeValue = dateTimeValue
	return b
}

func (b *_BACnetPriorityValueDateTimeBuilder) WithDateTimeValueBuilder(builderSupplier func(BACnetDateTimeEnclosedBuilder) BACnetDateTimeEnclosedBuilder) BACnetPriorityValueDateTimeBuilder {
	builder := builderSupplier(b.DateTimeValue.CreateBACnetDateTimeEnclosedBuilder())
	var err error
	b.DateTimeValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDateTimeEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetPriorityValueDateTimeBuilder) Build() (BACnetPriorityValueDateTime, error) {
	if b.DateTimeValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'dateTimeValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPriorityValueDateTime.deepCopy(), nil
}

func (b *_BACnetPriorityValueDateTimeBuilder) MustBuild() BACnetPriorityValueDateTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPriorityValueDateTimeBuilder) Done() BACnetPriorityValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetPriorityValueBuilder().(*_BACnetPriorityValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetPriorityValueDateTimeBuilder) buildForBACnetPriorityValue() (BACnetPriorityValue, error) {
	return b.Build()
}

func (b *_BACnetPriorityValueDateTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPriorityValueDateTimeBuilder().(*_BACnetPriorityValueDateTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPriorityValueDateTimeBuilder creates a BACnetPriorityValueDateTimeBuilder
func (b *_BACnetPriorityValueDateTime) CreateBACnetPriorityValueDateTimeBuilder() BACnetPriorityValueDateTimeBuilder {
	if b == nil {
		return NewBACnetPriorityValueDateTimeBuilder()
	}
	return &_BACnetPriorityValueDateTimeBuilder{_BACnetPriorityValueDateTime: b.deepCopy()}
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

func (m *_BACnetPriorityValueDateTime) GetParent() BACnetPriorityValueContract {
	return m.BACnetPriorityValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueDateTime) GetDateTimeValue() BACnetDateTimeEnclosed {
	return m.DateTimeValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueDateTime(structType any) BACnetPriorityValueDateTime {
	if casted, ok := structType.(BACnetPriorityValueDateTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueDateTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueDateTime) GetTypeName() string {
	return "BACnetPriorityValueDateTime"
}

func (m *_BACnetPriorityValueDateTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPriorityValueContract.(*_BACnetPriorityValue).getLengthInBits(ctx))

	// Simple field (dateTimeValue)
	lengthInBits += m.DateTimeValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPriorityValueDateTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPriorityValueDateTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPriorityValue, objectTypeArgument BACnetObjectType) (__bACnetPriorityValueDateTime BACnetPriorityValueDateTime, err error) {
	m.BACnetPriorityValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueDateTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueDateTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dateTimeValue, err := ReadSimpleField[BACnetDateTimeEnclosed](ctx, "dateTimeValue", ReadComplex[BACnetDateTimeEnclosed](BACnetDateTimeEnclosedParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dateTimeValue' field"))
	}
	m.DateTimeValue = dateTimeValue

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueDateTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueDateTime")
	}

	return m, nil
}

func (m *_BACnetPriorityValueDateTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueDateTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueDateTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueDateTime")
		}

		if err := WriteSimpleField[BACnetDateTimeEnclosed](ctx, "dateTimeValue", m.GetDateTimeValue(), WriteComplex[BACnetDateTimeEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dateTimeValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueDateTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueDateTime")
		}
		return nil
	}
	return m.BACnetPriorityValueContract.(*_BACnetPriorityValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueDateTime) IsBACnetPriorityValueDateTime() {}

func (m *_BACnetPriorityValueDateTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPriorityValueDateTime) deepCopy() *_BACnetPriorityValueDateTime {
	if m == nil {
		return nil
	}
	_BACnetPriorityValueDateTimeCopy := &_BACnetPriorityValueDateTime{
		m.BACnetPriorityValueContract.(*_BACnetPriorityValue).deepCopy(),
		utils.DeepCopy[BACnetDateTimeEnclosed](m.DateTimeValue),
	}
	_BACnetPriorityValueDateTimeCopy.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = m
	return _BACnetPriorityValueDateTimeCopy
}

func (m *_BACnetPriorityValueDateTime) String() string {
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
