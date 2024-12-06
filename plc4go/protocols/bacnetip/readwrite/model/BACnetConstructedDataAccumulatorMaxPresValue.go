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

// BACnetConstructedDataAccumulatorMaxPresValue is the corresponding interface of BACnetConstructedDataAccumulatorMaxPresValue
type BACnetConstructedDataAccumulatorMaxPresValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMaxPresValue returns MaxPresValue (property field)
	GetMaxPresValue() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataAccumulatorMaxPresValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAccumulatorMaxPresValue()
	// CreateBuilder creates a BACnetConstructedDataAccumulatorMaxPresValueBuilder
	CreateBACnetConstructedDataAccumulatorMaxPresValueBuilder() BACnetConstructedDataAccumulatorMaxPresValueBuilder
}

// _BACnetConstructedDataAccumulatorMaxPresValue is the data-structure of this message
type _BACnetConstructedDataAccumulatorMaxPresValue struct {
	BACnetConstructedDataContract
	MaxPresValue BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataAccumulatorMaxPresValue = (*_BACnetConstructedDataAccumulatorMaxPresValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccumulatorMaxPresValue)(nil)

// NewBACnetConstructedDataAccumulatorMaxPresValue factory function for _BACnetConstructedDataAccumulatorMaxPresValue
func NewBACnetConstructedDataAccumulatorMaxPresValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, maxPresValue BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccumulatorMaxPresValue {
	if maxPresValue == nil {
		panic("maxPresValue of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataAccumulatorMaxPresValue must not be nil")
	}
	_result := &_BACnetConstructedDataAccumulatorMaxPresValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MaxPresValue:                  maxPresValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAccumulatorMaxPresValueBuilder is a builder for BACnetConstructedDataAccumulatorMaxPresValue
type BACnetConstructedDataAccumulatorMaxPresValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(maxPresValue BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAccumulatorMaxPresValueBuilder
	// WithMaxPresValue adds MaxPresValue (property field)
	WithMaxPresValue(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAccumulatorMaxPresValueBuilder
	// WithMaxPresValueBuilder adds MaxPresValue (property field) which is build by the builder
	WithMaxPresValueBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAccumulatorMaxPresValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataAccumulatorMaxPresValue or returns an error if something is wrong
	Build() (BACnetConstructedDataAccumulatorMaxPresValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAccumulatorMaxPresValue
}

// NewBACnetConstructedDataAccumulatorMaxPresValueBuilder() creates a BACnetConstructedDataAccumulatorMaxPresValueBuilder
func NewBACnetConstructedDataAccumulatorMaxPresValueBuilder() BACnetConstructedDataAccumulatorMaxPresValueBuilder {
	return &_BACnetConstructedDataAccumulatorMaxPresValueBuilder{_BACnetConstructedDataAccumulatorMaxPresValue: new(_BACnetConstructedDataAccumulatorMaxPresValue)}
}

type _BACnetConstructedDataAccumulatorMaxPresValueBuilder struct {
	*_BACnetConstructedDataAccumulatorMaxPresValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAccumulatorMaxPresValueBuilder) = (*_BACnetConstructedDataAccumulatorMaxPresValueBuilder)(nil)

func (b *_BACnetConstructedDataAccumulatorMaxPresValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataAccumulatorMaxPresValue
}

func (b *_BACnetConstructedDataAccumulatorMaxPresValueBuilder) WithMandatoryFields(maxPresValue BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAccumulatorMaxPresValueBuilder {
	return b.WithMaxPresValue(maxPresValue)
}

func (b *_BACnetConstructedDataAccumulatorMaxPresValueBuilder) WithMaxPresValue(maxPresValue BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAccumulatorMaxPresValueBuilder {
	b.MaxPresValue = maxPresValue
	return b
}

func (b *_BACnetConstructedDataAccumulatorMaxPresValueBuilder) WithMaxPresValueBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAccumulatorMaxPresValueBuilder {
	builder := builderSupplier(b.MaxPresValue.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.MaxPresValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAccumulatorMaxPresValueBuilder) Build() (BACnetConstructedDataAccumulatorMaxPresValue, error) {
	if b.MaxPresValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'maxPresValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAccumulatorMaxPresValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataAccumulatorMaxPresValueBuilder) MustBuild() BACnetConstructedDataAccumulatorMaxPresValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataAccumulatorMaxPresValueBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAccumulatorMaxPresValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAccumulatorMaxPresValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAccumulatorMaxPresValueBuilder().(*_BACnetConstructedDataAccumulatorMaxPresValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAccumulatorMaxPresValueBuilder creates a BACnetConstructedDataAccumulatorMaxPresValueBuilder
func (b *_BACnetConstructedDataAccumulatorMaxPresValue) CreateBACnetConstructedDataAccumulatorMaxPresValueBuilder() BACnetConstructedDataAccumulatorMaxPresValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataAccumulatorMaxPresValueBuilder()
	}
	return &_BACnetConstructedDataAccumulatorMaxPresValueBuilder{_BACnetConstructedDataAccumulatorMaxPresValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorMaxPresValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCUMULATOR
}

func (m *_BACnetConstructedDataAccumulatorMaxPresValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_PRES_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccumulatorMaxPresValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorMaxPresValue) GetMaxPresValue() BACnetApplicationTagUnsignedInteger {
	return m.MaxPresValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorMaxPresValue) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetMaxPresValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccumulatorMaxPresValue(structType any) BACnetConstructedDataAccumulatorMaxPresValue {
	if casted, ok := structType.(BACnetConstructedDataAccumulatorMaxPresValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccumulatorMaxPresValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccumulatorMaxPresValue) GetTypeName() string {
	return "BACnetConstructedDataAccumulatorMaxPresValue"
}

func (m *_BACnetConstructedDataAccumulatorMaxPresValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (maxPresValue)
	lengthInBits += m.MaxPresValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccumulatorMaxPresValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccumulatorMaxPresValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccumulatorMaxPresValue BACnetConstructedDataAccumulatorMaxPresValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccumulatorMaxPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccumulatorMaxPresValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	maxPresValue, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "maxPresValue", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxPresValue' field"))
	}
	m.MaxPresValue = maxPresValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), maxPresValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccumulatorMaxPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccumulatorMaxPresValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccumulatorMaxPresValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccumulatorMaxPresValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccumulatorMaxPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccumulatorMaxPresValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "maxPresValue", m.GetMaxPresValue(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxPresValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccumulatorMaxPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccumulatorMaxPresValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccumulatorMaxPresValue) IsBACnetConstructedDataAccumulatorMaxPresValue() {
}

func (m *_BACnetConstructedDataAccumulatorMaxPresValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAccumulatorMaxPresValue) deepCopy() *_BACnetConstructedDataAccumulatorMaxPresValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAccumulatorMaxPresValueCopy := &_BACnetConstructedDataAccumulatorMaxPresValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.MaxPresValue),
	}
	_BACnetConstructedDataAccumulatorMaxPresValueCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAccumulatorMaxPresValueCopy
}

func (m *_BACnetConstructedDataAccumulatorMaxPresValue) String() string {
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
