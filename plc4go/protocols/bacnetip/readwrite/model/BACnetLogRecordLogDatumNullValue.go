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

// BACnetLogRecordLogDatumNullValue is the corresponding interface of BACnetLogRecordLogDatumNullValue
type BACnetLogRecordLogDatumNullValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetLogRecordLogDatum
	// GetNullValue returns NullValue (property field)
	GetNullValue() BACnetContextTagNull
	// IsBACnetLogRecordLogDatumNullValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogRecordLogDatumNullValue()
	// CreateBuilder creates a BACnetLogRecordLogDatumNullValueBuilder
	CreateBACnetLogRecordLogDatumNullValueBuilder() BACnetLogRecordLogDatumNullValueBuilder
}

// _BACnetLogRecordLogDatumNullValue is the data-structure of this message
type _BACnetLogRecordLogDatumNullValue struct {
	BACnetLogRecordLogDatumContract
	NullValue BACnetContextTagNull
}

var _ BACnetLogRecordLogDatumNullValue = (*_BACnetLogRecordLogDatumNullValue)(nil)
var _ BACnetLogRecordLogDatumRequirements = (*_BACnetLogRecordLogDatumNullValue)(nil)

// NewBACnetLogRecordLogDatumNullValue factory function for _BACnetLogRecordLogDatumNullValue
func NewBACnetLogRecordLogDatumNullValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, nullValue BACnetContextTagNull, tagNumber uint8) *_BACnetLogRecordLogDatumNullValue {
	if nullValue == nil {
		panic("nullValue of type BACnetContextTagNull for BACnetLogRecordLogDatumNullValue must not be nil")
	}
	_result := &_BACnetLogRecordLogDatumNullValue{
		BACnetLogRecordLogDatumContract: NewBACnetLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
		NullValue:                       nullValue,
	}
	_result.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLogRecordLogDatumNullValueBuilder is a builder for BACnetLogRecordLogDatumNullValue
type BACnetLogRecordLogDatumNullValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(nullValue BACnetContextTagNull) BACnetLogRecordLogDatumNullValueBuilder
	// WithNullValue adds NullValue (property field)
	WithNullValue(BACnetContextTagNull) BACnetLogRecordLogDatumNullValueBuilder
	// WithNullValueBuilder adds NullValue (property field) which is build by the builder
	WithNullValueBuilder(func(BACnetContextTagNullBuilder) BACnetContextTagNullBuilder) BACnetLogRecordLogDatumNullValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetLogRecordLogDatumBuilder
	// Build builds the BACnetLogRecordLogDatumNullValue or returns an error if something is wrong
	Build() (BACnetLogRecordLogDatumNullValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLogRecordLogDatumNullValue
}

// NewBACnetLogRecordLogDatumNullValueBuilder() creates a BACnetLogRecordLogDatumNullValueBuilder
func NewBACnetLogRecordLogDatumNullValueBuilder() BACnetLogRecordLogDatumNullValueBuilder {
	return &_BACnetLogRecordLogDatumNullValueBuilder{_BACnetLogRecordLogDatumNullValue: new(_BACnetLogRecordLogDatumNullValue)}
}

type _BACnetLogRecordLogDatumNullValueBuilder struct {
	*_BACnetLogRecordLogDatumNullValue

	parentBuilder *_BACnetLogRecordLogDatumBuilder

	err *utils.MultiError
}

var _ (BACnetLogRecordLogDatumNullValueBuilder) = (*_BACnetLogRecordLogDatumNullValueBuilder)(nil)

func (b *_BACnetLogRecordLogDatumNullValueBuilder) setParent(contract BACnetLogRecordLogDatumContract) {
	b.BACnetLogRecordLogDatumContract = contract
	contract.(*_BACnetLogRecordLogDatum)._SubType = b._BACnetLogRecordLogDatumNullValue
}

func (b *_BACnetLogRecordLogDatumNullValueBuilder) WithMandatoryFields(nullValue BACnetContextTagNull) BACnetLogRecordLogDatumNullValueBuilder {
	return b.WithNullValue(nullValue)
}

func (b *_BACnetLogRecordLogDatumNullValueBuilder) WithNullValue(nullValue BACnetContextTagNull) BACnetLogRecordLogDatumNullValueBuilder {
	b.NullValue = nullValue
	return b
}

func (b *_BACnetLogRecordLogDatumNullValueBuilder) WithNullValueBuilder(builderSupplier func(BACnetContextTagNullBuilder) BACnetContextTagNullBuilder) BACnetLogRecordLogDatumNullValueBuilder {
	builder := builderSupplier(b.NullValue.CreateBACnetContextTagNullBuilder())
	var err error
	b.NullValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagNullBuilder failed"))
	}
	return b
}

func (b *_BACnetLogRecordLogDatumNullValueBuilder) Build() (BACnetLogRecordLogDatumNullValue, error) {
	if b.NullValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'nullValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetLogRecordLogDatumNullValue.deepCopy(), nil
}

func (b *_BACnetLogRecordLogDatumNullValueBuilder) MustBuild() BACnetLogRecordLogDatumNullValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetLogRecordLogDatumNullValueBuilder) Done() BACnetLogRecordLogDatumBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetLogRecordLogDatumBuilder().(*_BACnetLogRecordLogDatumBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetLogRecordLogDatumNullValueBuilder) buildForBACnetLogRecordLogDatum() (BACnetLogRecordLogDatum, error) {
	return b.Build()
}

func (b *_BACnetLogRecordLogDatumNullValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetLogRecordLogDatumNullValueBuilder().(*_BACnetLogRecordLogDatumNullValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetLogRecordLogDatumNullValueBuilder creates a BACnetLogRecordLogDatumNullValueBuilder
func (b *_BACnetLogRecordLogDatumNullValue) CreateBACnetLogRecordLogDatumNullValueBuilder() BACnetLogRecordLogDatumNullValueBuilder {
	if b == nil {
		return NewBACnetLogRecordLogDatumNullValueBuilder()
	}
	return &_BACnetLogRecordLogDatumNullValueBuilder{_BACnetLogRecordLogDatumNullValue: b.deepCopy()}
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

func (m *_BACnetLogRecordLogDatumNullValue) GetParent() BACnetLogRecordLogDatumContract {
	return m.BACnetLogRecordLogDatumContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecordLogDatumNullValue) GetNullValue() BACnetContextTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLogRecordLogDatumNullValue(structType any) BACnetLogRecordLogDatumNullValue {
	if casted, ok := structType.(BACnetLogRecordLogDatumNullValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatumNullValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecordLogDatumNullValue) GetTypeName() string {
	return "BACnetLogRecordLogDatumNullValue"
}

func (m *_BACnetLogRecordLogDatumNullValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).getLengthInBits(ctx))

	// Simple field (nullValue)
	lengthInBits += m.NullValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogRecordLogDatumNullValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogRecordLogDatumNullValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogRecordLogDatum, tagNumber uint8) (__bACnetLogRecordLogDatumNullValue BACnetLogRecordLogDatumNullValue, err error) {
	m.BACnetLogRecordLogDatumContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatumNullValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatumNullValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nullValue, err := ReadSimpleField[BACnetContextTagNull](ctx, "nullValue", ReadComplex[BACnetContextTagNull](BACnetContextTagParseWithBufferProducer[BACnetContextTagNull]((uint8)(uint8(7)), (BACnetDataType)(BACnetDataType_NULL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nullValue' field"))
	}
	m.NullValue = nullValue

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatumNullValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatumNullValue")
	}

	return m, nil
}

func (m *_BACnetLogRecordLogDatumNullValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogRecordLogDatumNullValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatumNullValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatumNullValue")
		}

		if err := WriteSimpleField[BACnetContextTagNull](ctx, "nullValue", m.GetNullValue(), WriteComplex[BACnetContextTagNull](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nullValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatumNullValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatumNullValue")
		}
		return nil
	}
	return m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogRecordLogDatumNullValue) IsBACnetLogRecordLogDatumNullValue() {}

func (m *_BACnetLogRecordLogDatumNullValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLogRecordLogDatumNullValue) deepCopy() *_BACnetLogRecordLogDatumNullValue {
	if m == nil {
		return nil
	}
	_BACnetLogRecordLogDatumNullValueCopy := &_BACnetLogRecordLogDatumNullValue{
		m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).deepCopy(),
		utils.DeepCopy[BACnetContextTagNull](m.NullValue),
	}
	_BACnetLogRecordLogDatumNullValueCopy.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum)._SubType = m
	return _BACnetLogRecordLogDatumNullValueCopy
}

func (m *_BACnetLogRecordLogDatumNullValue) String() string {
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