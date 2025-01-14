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

// BACnetLogDataLogDataEntryEnumeratedValue is the corresponding interface of BACnetLogDataLogDataEntryEnumeratedValue
type BACnetLogDataLogDataEntryEnumeratedValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetLogDataLogDataEntry
	// GetEnumeratedValue returns EnumeratedValue (property field)
	GetEnumeratedValue() BACnetContextTagEnumerated
	// IsBACnetLogDataLogDataEntryEnumeratedValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogDataLogDataEntryEnumeratedValue()
	// CreateBuilder creates a BACnetLogDataLogDataEntryEnumeratedValueBuilder
	CreateBACnetLogDataLogDataEntryEnumeratedValueBuilder() BACnetLogDataLogDataEntryEnumeratedValueBuilder
}

// _BACnetLogDataLogDataEntryEnumeratedValue is the data-structure of this message
type _BACnetLogDataLogDataEntryEnumeratedValue struct {
	BACnetLogDataLogDataEntryContract
	EnumeratedValue BACnetContextTagEnumerated
}

var _ BACnetLogDataLogDataEntryEnumeratedValue = (*_BACnetLogDataLogDataEntryEnumeratedValue)(nil)
var _ BACnetLogDataLogDataEntryRequirements = (*_BACnetLogDataLogDataEntryEnumeratedValue)(nil)

// NewBACnetLogDataLogDataEntryEnumeratedValue factory function for _BACnetLogDataLogDataEntryEnumeratedValue
func NewBACnetLogDataLogDataEntryEnumeratedValue(peekedTagHeader BACnetTagHeader, enumeratedValue BACnetContextTagEnumerated) *_BACnetLogDataLogDataEntryEnumeratedValue {
	if enumeratedValue == nil {
		panic("enumeratedValue of type BACnetContextTagEnumerated for BACnetLogDataLogDataEntryEnumeratedValue must not be nil")
	}
	_result := &_BACnetLogDataLogDataEntryEnumeratedValue{
		BACnetLogDataLogDataEntryContract: NewBACnetLogDataLogDataEntry(peekedTagHeader),
		EnumeratedValue:                   enumeratedValue,
	}
	_result.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLogDataLogDataEntryEnumeratedValueBuilder is a builder for BACnetLogDataLogDataEntryEnumeratedValue
type BACnetLogDataLogDataEntryEnumeratedValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(enumeratedValue BACnetContextTagEnumerated) BACnetLogDataLogDataEntryEnumeratedValueBuilder
	// WithEnumeratedValue adds EnumeratedValue (property field)
	WithEnumeratedValue(BACnetContextTagEnumerated) BACnetLogDataLogDataEntryEnumeratedValueBuilder
	// WithEnumeratedValueBuilder adds EnumeratedValue (property field) which is build by the builder
	WithEnumeratedValueBuilder(func(BACnetContextTagEnumeratedBuilder) BACnetContextTagEnumeratedBuilder) BACnetLogDataLogDataEntryEnumeratedValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetLogDataLogDataEntryBuilder
	// Build builds the BACnetLogDataLogDataEntryEnumeratedValue or returns an error if something is wrong
	Build() (BACnetLogDataLogDataEntryEnumeratedValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLogDataLogDataEntryEnumeratedValue
}

// NewBACnetLogDataLogDataEntryEnumeratedValueBuilder() creates a BACnetLogDataLogDataEntryEnumeratedValueBuilder
func NewBACnetLogDataLogDataEntryEnumeratedValueBuilder() BACnetLogDataLogDataEntryEnumeratedValueBuilder {
	return &_BACnetLogDataLogDataEntryEnumeratedValueBuilder{_BACnetLogDataLogDataEntryEnumeratedValue: new(_BACnetLogDataLogDataEntryEnumeratedValue)}
}

type _BACnetLogDataLogDataEntryEnumeratedValueBuilder struct {
	*_BACnetLogDataLogDataEntryEnumeratedValue

	parentBuilder *_BACnetLogDataLogDataEntryBuilder

	err *utils.MultiError
}

var _ (BACnetLogDataLogDataEntryEnumeratedValueBuilder) = (*_BACnetLogDataLogDataEntryEnumeratedValueBuilder)(nil)

func (b *_BACnetLogDataLogDataEntryEnumeratedValueBuilder) setParent(contract BACnetLogDataLogDataEntryContract) {
	b.BACnetLogDataLogDataEntryContract = contract
	contract.(*_BACnetLogDataLogDataEntry)._SubType = b._BACnetLogDataLogDataEntryEnumeratedValue
}

func (b *_BACnetLogDataLogDataEntryEnumeratedValueBuilder) WithMandatoryFields(enumeratedValue BACnetContextTagEnumerated) BACnetLogDataLogDataEntryEnumeratedValueBuilder {
	return b.WithEnumeratedValue(enumeratedValue)
}

func (b *_BACnetLogDataLogDataEntryEnumeratedValueBuilder) WithEnumeratedValue(enumeratedValue BACnetContextTagEnumerated) BACnetLogDataLogDataEntryEnumeratedValueBuilder {
	b.EnumeratedValue = enumeratedValue
	return b
}

func (b *_BACnetLogDataLogDataEntryEnumeratedValueBuilder) WithEnumeratedValueBuilder(builderSupplier func(BACnetContextTagEnumeratedBuilder) BACnetContextTagEnumeratedBuilder) BACnetLogDataLogDataEntryEnumeratedValueBuilder {
	builder := builderSupplier(b.EnumeratedValue.CreateBACnetContextTagEnumeratedBuilder())
	var err error
	b.EnumeratedValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagEnumeratedBuilder failed"))
	}
	return b
}

func (b *_BACnetLogDataLogDataEntryEnumeratedValueBuilder) Build() (BACnetLogDataLogDataEntryEnumeratedValue, error) {
	if b.EnumeratedValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'enumeratedValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetLogDataLogDataEntryEnumeratedValue.deepCopy(), nil
}

func (b *_BACnetLogDataLogDataEntryEnumeratedValueBuilder) MustBuild() BACnetLogDataLogDataEntryEnumeratedValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetLogDataLogDataEntryEnumeratedValueBuilder) Done() BACnetLogDataLogDataEntryBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetLogDataLogDataEntryBuilder().(*_BACnetLogDataLogDataEntryBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetLogDataLogDataEntryEnumeratedValueBuilder) buildForBACnetLogDataLogDataEntry() (BACnetLogDataLogDataEntry, error) {
	return b.Build()
}

func (b *_BACnetLogDataLogDataEntryEnumeratedValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetLogDataLogDataEntryEnumeratedValueBuilder().(*_BACnetLogDataLogDataEntryEnumeratedValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetLogDataLogDataEntryEnumeratedValueBuilder creates a BACnetLogDataLogDataEntryEnumeratedValueBuilder
func (b *_BACnetLogDataLogDataEntryEnumeratedValue) CreateBACnetLogDataLogDataEntryEnumeratedValueBuilder() BACnetLogDataLogDataEntryEnumeratedValueBuilder {
	if b == nil {
		return NewBACnetLogDataLogDataEntryEnumeratedValueBuilder()
	}
	return &_BACnetLogDataLogDataEntryEnumeratedValueBuilder{_BACnetLogDataLogDataEntryEnumeratedValue: b.deepCopy()}
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

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) GetParent() BACnetLogDataLogDataEntryContract {
	return m.BACnetLogDataLogDataEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) GetEnumeratedValue() BACnetContextTagEnumerated {
	return m.EnumeratedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogDataEntryEnumeratedValue(structType any) BACnetLogDataLogDataEntryEnumeratedValue {
	if casted, ok := structType.(BACnetLogDataLogDataEntryEnumeratedValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntryEnumeratedValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) GetTypeName() string {
	return "BACnetLogDataLogDataEntryEnumeratedValue"
}

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry).getLengthInBits(ctx))

	// Simple field (enumeratedValue)
	lengthInBits += m.EnumeratedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogDataLogDataEntry) (__bACnetLogDataLogDataEntryEnumeratedValue BACnetLogDataLogDataEntryEnumeratedValue, err error) {
	m.BACnetLogDataLogDataEntryContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataEntryEnumeratedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataEntryEnumeratedValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	enumeratedValue, err := ReadSimpleField[BACnetContextTagEnumerated](ctx, "enumeratedValue", ReadComplex[BACnetContextTagEnumerated](BACnetContextTagParseWithBufferProducer[BACnetContextTagEnumerated]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_ENUMERATED)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enumeratedValue' field"))
	}
	m.EnumeratedValue = enumeratedValue

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataEntryEnumeratedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataEntryEnumeratedValue")
	}

	return m, nil
}

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataEntryEnumeratedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataEntryEnumeratedValue")
		}

		if err := WriteSimpleField[BACnetContextTagEnumerated](ctx, "enumeratedValue", m.GetEnumeratedValue(), WriteComplex[BACnetContextTagEnumerated](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'enumeratedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataEntryEnumeratedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataEntryEnumeratedValue")
		}
		return nil
	}
	return m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) IsBACnetLogDataLogDataEntryEnumeratedValue() {}

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) deepCopy() *_BACnetLogDataLogDataEntryEnumeratedValue {
	if m == nil {
		return nil
	}
	_BACnetLogDataLogDataEntryEnumeratedValueCopy := &_BACnetLogDataLogDataEntryEnumeratedValue{
		m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry).deepCopy(),
		utils.DeepCopy[BACnetContextTagEnumerated](m.EnumeratedValue),
	}
	_BACnetLogDataLogDataEntryEnumeratedValueCopy.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry)._SubType = m
	return _BACnetLogDataLogDataEntryEnumeratedValueCopy
}

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) String() string {
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
