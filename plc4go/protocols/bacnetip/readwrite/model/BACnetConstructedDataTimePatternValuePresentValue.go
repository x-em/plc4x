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

// BACnetConstructedDataTimePatternValuePresentValue is the corresponding interface of BACnetConstructedDataTimePatternValuePresentValue
type BACnetConstructedDataTimePatternValuePresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetApplicationTagTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagTime
	// IsBACnetConstructedDataTimePatternValuePresentValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTimePatternValuePresentValue()
	// CreateBuilder creates a BACnetConstructedDataTimePatternValuePresentValueBuilder
	CreateBACnetConstructedDataTimePatternValuePresentValueBuilder() BACnetConstructedDataTimePatternValuePresentValueBuilder
}

// _BACnetConstructedDataTimePatternValuePresentValue is the data-structure of this message
type _BACnetConstructedDataTimePatternValuePresentValue struct {
	BACnetConstructedDataContract
	PresentValue BACnetApplicationTagTime
}

var _ BACnetConstructedDataTimePatternValuePresentValue = (*_BACnetConstructedDataTimePatternValuePresentValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTimePatternValuePresentValue)(nil)

// NewBACnetConstructedDataTimePatternValuePresentValue factory function for _BACnetConstructedDataTimePatternValuePresentValue
func NewBACnetConstructedDataTimePatternValuePresentValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, presentValue BACnetApplicationTagTime, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimePatternValuePresentValue {
	if presentValue == nil {
		panic("presentValue of type BACnetApplicationTagTime for BACnetConstructedDataTimePatternValuePresentValue must not be nil")
	}
	_result := &_BACnetConstructedDataTimePatternValuePresentValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		PresentValue:                  presentValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataTimePatternValuePresentValueBuilder is a builder for BACnetConstructedDataTimePatternValuePresentValue
type BACnetConstructedDataTimePatternValuePresentValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(presentValue BACnetApplicationTagTime) BACnetConstructedDataTimePatternValuePresentValueBuilder
	// WithPresentValue adds PresentValue (property field)
	WithPresentValue(BACnetApplicationTagTime) BACnetConstructedDataTimePatternValuePresentValueBuilder
	// WithPresentValueBuilder adds PresentValue (property field) which is build by the builder
	WithPresentValueBuilder(func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetConstructedDataTimePatternValuePresentValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataTimePatternValuePresentValue or returns an error if something is wrong
	Build() (BACnetConstructedDataTimePatternValuePresentValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataTimePatternValuePresentValue
}

// NewBACnetConstructedDataTimePatternValuePresentValueBuilder() creates a BACnetConstructedDataTimePatternValuePresentValueBuilder
func NewBACnetConstructedDataTimePatternValuePresentValueBuilder() BACnetConstructedDataTimePatternValuePresentValueBuilder {
	return &_BACnetConstructedDataTimePatternValuePresentValueBuilder{_BACnetConstructedDataTimePatternValuePresentValue: new(_BACnetConstructedDataTimePatternValuePresentValue)}
}

type _BACnetConstructedDataTimePatternValuePresentValueBuilder struct {
	*_BACnetConstructedDataTimePatternValuePresentValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataTimePatternValuePresentValueBuilder) = (*_BACnetConstructedDataTimePatternValuePresentValueBuilder)(nil)

func (b *_BACnetConstructedDataTimePatternValuePresentValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataTimePatternValuePresentValue
}

func (b *_BACnetConstructedDataTimePatternValuePresentValueBuilder) WithMandatoryFields(presentValue BACnetApplicationTagTime) BACnetConstructedDataTimePatternValuePresentValueBuilder {
	return b.WithPresentValue(presentValue)
}

func (b *_BACnetConstructedDataTimePatternValuePresentValueBuilder) WithPresentValue(presentValue BACnetApplicationTagTime) BACnetConstructedDataTimePatternValuePresentValueBuilder {
	b.PresentValue = presentValue
	return b
}

func (b *_BACnetConstructedDataTimePatternValuePresentValueBuilder) WithPresentValueBuilder(builderSupplier func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetConstructedDataTimePatternValuePresentValueBuilder {
	builder := builderSupplier(b.PresentValue.CreateBACnetApplicationTagTimeBuilder())
	var err error
	b.PresentValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagTimeBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataTimePatternValuePresentValueBuilder) Build() (BACnetConstructedDataTimePatternValuePresentValue, error) {
	if b.PresentValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'presentValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataTimePatternValuePresentValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataTimePatternValuePresentValueBuilder) MustBuild() BACnetConstructedDataTimePatternValuePresentValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataTimePatternValuePresentValueBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataTimePatternValuePresentValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataTimePatternValuePresentValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataTimePatternValuePresentValueBuilder().(*_BACnetConstructedDataTimePatternValuePresentValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataTimePatternValuePresentValueBuilder creates a BACnetConstructedDataTimePatternValuePresentValueBuilder
func (b *_BACnetConstructedDataTimePatternValuePresentValue) CreateBACnetConstructedDataTimePatternValuePresentValueBuilder() BACnetConstructedDataTimePatternValuePresentValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataTimePatternValuePresentValueBuilder()
	}
	return &_BACnetConstructedDataTimePatternValuePresentValueBuilder{_BACnetConstructedDataTimePatternValuePresentValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimePatternValuePresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TIMEPATTERN_VALUE
}

func (m *_BACnetConstructedDataTimePatternValuePresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimePatternValuePresentValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimePatternValuePresentValue) GetPresentValue() BACnetApplicationTagTime {
	return m.PresentValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTimePatternValuePresentValue) GetActualValue() BACnetApplicationTagTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagTime(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimePatternValuePresentValue(structType any) BACnetConstructedDataTimePatternValuePresentValue {
	if casted, ok := structType.(BACnetConstructedDataTimePatternValuePresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimePatternValuePresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimePatternValuePresentValue) GetTypeName() string {
	return "BACnetConstructedDataTimePatternValuePresentValue"
}

func (m *_BACnetConstructedDataTimePatternValuePresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimePatternValuePresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTimePatternValuePresentValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTimePatternValuePresentValue BACnetConstructedDataTimePatternValuePresentValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimePatternValuePresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimePatternValuePresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	presentValue, err := ReadSimpleField[BACnetApplicationTagTime](ctx, "presentValue", ReadComplex[BACnetApplicationTagTime](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagTime](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagTime](ctx, "actualValue", (*BACnetApplicationTagTime)(nil), presentValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimePatternValuePresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimePatternValuePresentValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTimePatternValuePresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimePatternValuePresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimePatternValuePresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimePatternValuePresentValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagTime](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetApplicationTagTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimePatternValuePresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimePatternValuePresentValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimePatternValuePresentValue) IsBACnetConstructedDataTimePatternValuePresentValue() {
}

func (m *_BACnetConstructedDataTimePatternValuePresentValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataTimePatternValuePresentValue) deepCopy() *_BACnetConstructedDataTimePatternValuePresentValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataTimePatternValuePresentValueCopy := &_BACnetConstructedDataTimePatternValuePresentValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagTime](m.PresentValue),
	}
	_BACnetConstructedDataTimePatternValuePresentValueCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataTimePatternValuePresentValueCopy
}

func (m *_BACnetConstructedDataTimePatternValuePresentValue) String() string {
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