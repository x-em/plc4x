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

// BACnetConstructedDataAnalogOutputInterfaceValue is the corresponding interface of BACnetConstructedDataAnalogOutputInterfaceValue
type BACnetConstructedDataAnalogOutputInterfaceValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetInterfaceValue returns InterfaceValue (property field)
	GetInterfaceValue() BACnetOptionalREAL
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetOptionalREAL
	// IsBACnetConstructedDataAnalogOutputInterfaceValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAnalogOutputInterfaceValue()
	// CreateBuilder creates a BACnetConstructedDataAnalogOutputInterfaceValueBuilder
	CreateBACnetConstructedDataAnalogOutputInterfaceValueBuilder() BACnetConstructedDataAnalogOutputInterfaceValueBuilder
}

// _BACnetConstructedDataAnalogOutputInterfaceValue is the data-structure of this message
type _BACnetConstructedDataAnalogOutputInterfaceValue struct {
	BACnetConstructedDataContract
	InterfaceValue BACnetOptionalREAL
}

var _ BACnetConstructedDataAnalogOutputInterfaceValue = (*_BACnetConstructedDataAnalogOutputInterfaceValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAnalogOutputInterfaceValue)(nil)

// NewBACnetConstructedDataAnalogOutputInterfaceValue factory function for _BACnetConstructedDataAnalogOutputInterfaceValue
func NewBACnetConstructedDataAnalogOutputInterfaceValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, interfaceValue BACnetOptionalREAL, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAnalogOutputInterfaceValue {
	if interfaceValue == nil {
		panic("interfaceValue of type BACnetOptionalREAL for BACnetConstructedDataAnalogOutputInterfaceValue must not be nil")
	}
	_result := &_BACnetConstructedDataAnalogOutputInterfaceValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		InterfaceValue:                interfaceValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAnalogOutputInterfaceValueBuilder is a builder for BACnetConstructedDataAnalogOutputInterfaceValue
type BACnetConstructedDataAnalogOutputInterfaceValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(interfaceValue BACnetOptionalREAL) BACnetConstructedDataAnalogOutputInterfaceValueBuilder
	// WithInterfaceValue adds InterfaceValue (property field)
	WithInterfaceValue(BACnetOptionalREAL) BACnetConstructedDataAnalogOutputInterfaceValueBuilder
	// WithInterfaceValueBuilder adds InterfaceValue (property field) which is build by the builder
	WithInterfaceValueBuilder(func(BACnetOptionalREALBuilder) BACnetOptionalREALBuilder) BACnetConstructedDataAnalogOutputInterfaceValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataAnalogOutputInterfaceValue or returns an error if something is wrong
	Build() (BACnetConstructedDataAnalogOutputInterfaceValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAnalogOutputInterfaceValue
}

// NewBACnetConstructedDataAnalogOutputInterfaceValueBuilder() creates a BACnetConstructedDataAnalogOutputInterfaceValueBuilder
func NewBACnetConstructedDataAnalogOutputInterfaceValueBuilder() BACnetConstructedDataAnalogOutputInterfaceValueBuilder {
	return &_BACnetConstructedDataAnalogOutputInterfaceValueBuilder{_BACnetConstructedDataAnalogOutputInterfaceValue: new(_BACnetConstructedDataAnalogOutputInterfaceValue)}
}

type _BACnetConstructedDataAnalogOutputInterfaceValueBuilder struct {
	*_BACnetConstructedDataAnalogOutputInterfaceValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAnalogOutputInterfaceValueBuilder) = (*_BACnetConstructedDataAnalogOutputInterfaceValueBuilder)(nil)

func (b *_BACnetConstructedDataAnalogOutputInterfaceValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataAnalogOutputInterfaceValue
}

func (b *_BACnetConstructedDataAnalogOutputInterfaceValueBuilder) WithMandatoryFields(interfaceValue BACnetOptionalREAL) BACnetConstructedDataAnalogOutputInterfaceValueBuilder {
	return b.WithInterfaceValue(interfaceValue)
}

func (b *_BACnetConstructedDataAnalogOutputInterfaceValueBuilder) WithInterfaceValue(interfaceValue BACnetOptionalREAL) BACnetConstructedDataAnalogOutputInterfaceValueBuilder {
	b.InterfaceValue = interfaceValue
	return b
}

func (b *_BACnetConstructedDataAnalogOutputInterfaceValueBuilder) WithInterfaceValueBuilder(builderSupplier func(BACnetOptionalREALBuilder) BACnetOptionalREALBuilder) BACnetConstructedDataAnalogOutputInterfaceValueBuilder {
	builder := builderSupplier(b.InterfaceValue.CreateBACnetOptionalREALBuilder())
	var err error
	b.InterfaceValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOptionalREALBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAnalogOutputInterfaceValueBuilder) Build() (BACnetConstructedDataAnalogOutputInterfaceValue, error) {
	if b.InterfaceValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'interfaceValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAnalogOutputInterfaceValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataAnalogOutputInterfaceValueBuilder) MustBuild() BACnetConstructedDataAnalogOutputInterfaceValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataAnalogOutputInterfaceValueBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAnalogOutputInterfaceValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAnalogOutputInterfaceValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAnalogOutputInterfaceValueBuilder().(*_BACnetConstructedDataAnalogOutputInterfaceValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAnalogOutputInterfaceValueBuilder creates a BACnetConstructedDataAnalogOutputInterfaceValueBuilder
func (b *_BACnetConstructedDataAnalogOutputInterfaceValue) CreateBACnetConstructedDataAnalogOutputInterfaceValueBuilder() BACnetConstructedDataAnalogOutputInterfaceValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataAnalogOutputInterfaceValueBuilder()
	}
	return &_BACnetConstructedDataAnalogOutputInterfaceValueBuilder{_BACnetConstructedDataAnalogOutputInterfaceValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAnalogOutputInterfaceValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ANALOG_OUTPUT
}

func (m *_BACnetConstructedDataAnalogOutputInterfaceValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INTERFACE_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAnalogOutputInterfaceValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogOutputInterfaceValue) GetInterfaceValue() BACnetOptionalREAL {
	return m.InterfaceValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogOutputInterfaceValue) GetActualValue() BACnetOptionalREAL {
	ctx := context.Background()
	_ = ctx
	return CastBACnetOptionalREAL(m.GetInterfaceValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAnalogOutputInterfaceValue(structType any) BACnetConstructedDataAnalogOutputInterfaceValue {
	if casted, ok := structType.(BACnetConstructedDataAnalogOutputInterfaceValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAnalogOutputInterfaceValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAnalogOutputInterfaceValue) GetTypeName() string {
	return "BACnetConstructedDataAnalogOutputInterfaceValue"
}

func (m *_BACnetConstructedDataAnalogOutputInterfaceValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (interfaceValue)
	lengthInBits += m.InterfaceValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAnalogOutputInterfaceValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAnalogOutputInterfaceValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAnalogOutputInterfaceValue BACnetConstructedDataAnalogOutputInterfaceValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAnalogOutputInterfaceValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAnalogOutputInterfaceValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	interfaceValue, err := ReadSimpleField[BACnetOptionalREAL](ctx, "interfaceValue", ReadComplex[BACnetOptionalREAL](BACnetOptionalREALParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'interfaceValue' field"))
	}
	m.InterfaceValue = interfaceValue

	actualValue, err := ReadVirtualField[BACnetOptionalREAL](ctx, "actualValue", (*BACnetOptionalREAL)(nil), interfaceValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAnalogOutputInterfaceValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAnalogOutputInterfaceValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAnalogOutputInterfaceValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAnalogOutputInterfaceValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAnalogOutputInterfaceValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAnalogOutputInterfaceValue")
		}

		if err := WriteSimpleField[BACnetOptionalREAL](ctx, "interfaceValue", m.GetInterfaceValue(), WriteComplex[BACnetOptionalREAL](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'interfaceValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAnalogOutputInterfaceValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAnalogOutputInterfaceValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAnalogOutputInterfaceValue) IsBACnetConstructedDataAnalogOutputInterfaceValue() {
}

func (m *_BACnetConstructedDataAnalogOutputInterfaceValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAnalogOutputInterfaceValue) deepCopy() *_BACnetConstructedDataAnalogOutputInterfaceValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAnalogOutputInterfaceValueCopy := &_BACnetConstructedDataAnalogOutputInterfaceValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetOptionalREAL](m.InterfaceValue),
	}
	_BACnetConstructedDataAnalogOutputInterfaceValueCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAnalogOutputInterfaceValueCopy
}

func (m *_BACnetConstructedDataAnalogOutputInterfaceValue) String() string {
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
