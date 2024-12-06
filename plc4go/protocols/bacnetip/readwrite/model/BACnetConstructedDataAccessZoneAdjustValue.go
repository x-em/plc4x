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

// BACnetConstructedDataAccessZoneAdjustValue is the corresponding interface of BACnetConstructedDataAccessZoneAdjustValue
type BACnetConstructedDataAccessZoneAdjustValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetAdjustValue returns AdjustValue (property field)
	GetAdjustValue() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
	// IsBACnetConstructedDataAccessZoneAdjustValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAccessZoneAdjustValue()
	// CreateBuilder creates a BACnetConstructedDataAccessZoneAdjustValueBuilder
	CreateBACnetConstructedDataAccessZoneAdjustValueBuilder() BACnetConstructedDataAccessZoneAdjustValueBuilder
}

// _BACnetConstructedDataAccessZoneAdjustValue is the data-structure of this message
type _BACnetConstructedDataAccessZoneAdjustValue struct {
	BACnetConstructedDataContract
	AdjustValue BACnetApplicationTagSignedInteger
}

var _ BACnetConstructedDataAccessZoneAdjustValue = (*_BACnetConstructedDataAccessZoneAdjustValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccessZoneAdjustValue)(nil)

// NewBACnetConstructedDataAccessZoneAdjustValue factory function for _BACnetConstructedDataAccessZoneAdjustValue
func NewBACnetConstructedDataAccessZoneAdjustValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, adjustValue BACnetApplicationTagSignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessZoneAdjustValue {
	if adjustValue == nil {
		panic("adjustValue of type BACnetApplicationTagSignedInteger for BACnetConstructedDataAccessZoneAdjustValue must not be nil")
	}
	_result := &_BACnetConstructedDataAccessZoneAdjustValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AdjustValue:                   adjustValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAccessZoneAdjustValueBuilder is a builder for BACnetConstructedDataAccessZoneAdjustValue
type BACnetConstructedDataAccessZoneAdjustValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(adjustValue BACnetApplicationTagSignedInteger) BACnetConstructedDataAccessZoneAdjustValueBuilder
	// WithAdjustValue adds AdjustValue (property field)
	WithAdjustValue(BACnetApplicationTagSignedInteger) BACnetConstructedDataAccessZoneAdjustValueBuilder
	// WithAdjustValueBuilder adds AdjustValue (property field) which is build by the builder
	WithAdjustValueBuilder(func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConstructedDataAccessZoneAdjustValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataAccessZoneAdjustValue or returns an error if something is wrong
	Build() (BACnetConstructedDataAccessZoneAdjustValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAccessZoneAdjustValue
}

// NewBACnetConstructedDataAccessZoneAdjustValueBuilder() creates a BACnetConstructedDataAccessZoneAdjustValueBuilder
func NewBACnetConstructedDataAccessZoneAdjustValueBuilder() BACnetConstructedDataAccessZoneAdjustValueBuilder {
	return &_BACnetConstructedDataAccessZoneAdjustValueBuilder{_BACnetConstructedDataAccessZoneAdjustValue: new(_BACnetConstructedDataAccessZoneAdjustValue)}
}

type _BACnetConstructedDataAccessZoneAdjustValueBuilder struct {
	*_BACnetConstructedDataAccessZoneAdjustValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAccessZoneAdjustValueBuilder) = (*_BACnetConstructedDataAccessZoneAdjustValueBuilder)(nil)

func (b *_BACnetConstructedDataAccessZoneAdjustValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataAccessZoneAdjustValue
}

func (b *_BACnetConstructedDataAccessZoneAdjustValueBuilder) WithMandatoryFields(adjustValue BACnetApplicationTagSignedInteger) BACnetConstructedDataAccessZoneAdjustValueBuilder {
	return b.WithAdjustValue(adjustValue)
}

func (b *_BACnetConstructedDataAccessZoneAdjustValueBuilder) WithAdjustValue(adjustValue BACnetApplicationTagSignedInteger) BACnetConstructedDataAccessZoneAdjustValueBuilder {
	b.AdjustValue = adjustValue
	return b
}

func (b *_BACnetConstructedDataAccessZoneAdjustValueBuilder) WithAdjustValueBuilder(builderSupplier func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConstructedDataAccessZoneAdjustValueBuilder {
	builder := builderSupplier(b.AdjustValue.CreateBACnetApplicationTagSignedIntegerBuilder())
	var err error
	b.AdjustValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAccessZoneAdjustValueBuilder) Build() (BACnetConstructedDataAccessZoneAdjustValue, error) {
	if b.AdjustValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'adjustValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAccessZoneAdjustValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataAccessZoneAdjustValueBuilder) MustBuild() BACnetConstructedDataAccessZoneAdjustValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataAccessZoneAdjustValueBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAccessZoneAdjustValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAccessZoneAdjustValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAccessZoneAdjustValueBuilder().(*_BACnetConstructedDataAccessZoneAdjustValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAccessZoneAdjustValueBuilder creates a BACnetConstructedDataAccessZoneAdjustValueBuilder
func (b *_BACnetConstructedDataAccessZoneAdjustValue) CreateBACnetConstructedDataAccessZoneAdjustValueBuilder() BACnetConstructedDataAccessZoneAdjustValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataAccessZoneAdjustValueBuilder()
	}
	return &_BACnetConstructedDataAccessZoneAdjustValueBuilder{_BACnetConstructedDataAccessZoneAdjustValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCESS_ZONE
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ADJUST_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetAdjustValue() BACnetApplicationTagSignedInteger {
	return m.AdjustValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetActualValue() BACnetApplicationTagSignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagSignedInteger(m.GetAdjustValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessZoneAdjustValue(structType any) BACnetConstructedDataAccessZoneAdjustValue {
	if casted, ok := structType.(BACnetConstructedDataAccessZoneAdjustValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessZoneAdjustValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetTypeName() string {
	return "BACnetConstructedDataAccessZoneAdjustValue"
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (adjustValue)
	lengthInBits += m.AdjustValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccessZoneAdjustValue BACnetConstructedDataAccessZoneAdjustValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessZoneAdjustValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessZoneAdjustValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	adjustValue, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "adjustValue", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'adjustValue' field"))
	}
	m.AdjustValue = adjustValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagSignedInteger](ctx, "actualValue", (*BACnetApplicationTagSignedInteger)(nil), adjustValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessZoneAdjustValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessZoneAdjustValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessZoneAdjustValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessZoneAdjustValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "adjustValue", m.GetAdjustValue(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'adjustValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessZoneAdjustValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessZoneAdjustValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) IsBACnetConstructedDataAccessZoneAdjustValue() {
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) deepCopy() *_BACnetConstructedDataAccessZoneAdjustValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAccessZoneAdjustValueCopy := &_BACnetConstructedDataAccessZoneAdjustValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagSignedInteger](m.AdjustValue),
	}
	_BACnetConstructedDataAccessZoneAdjustValueCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAccessZoneAdjustValueCopy
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) String() string {
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
