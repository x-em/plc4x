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

// BACnetConstructedDataAccessDoorPresentValue is the corresponding interface of BACnetConstructedDataAccessDoorPresentValue
type BACnetConstructedDataAccessDoorPresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetDoorValueTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDoorValueTagged
	// IsBACnetConstructedDataAccessDoorPresentValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAccessDoorPresentValue()
	// CreateBuilder creates a BACnetConstructedDataAccessDoorPresentValueBuilder
	CreateBACnetConstructedDataAccessDoorPresentValueBuilder() BACnetConstructedDataAccessDoorPresentValueBuilder
}

// _BACnetConstructedDataAccessDoorPresentValue is the data-structure of this message
type _BACnetConstructedDataAccessDoorPresentValue struct {
	BACnetConstructedDataContract
	PresentValue BACnetDoorValueTagged
}

var _ BACnetConstructedDataAccessDoorPresentValue = (*_BACnetConstructedDataAccessDoorPresentValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccessDoorPresentValue)(nil)

// NewBACnetConstructedDataAccessDoorPresentValue factory function for _BACnetConstructedDataAccessDoorPresentValue
func NewBACnetConstructedDataAccessDoorPresentValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, presentValue BACnetDoorValueTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessDoorPresentValue {
	if presentValue == nil {
		panic("presentValue of type BACnetDoorValueTagged for BACnetConstructedDataAccessDoorPresentValue must not be nil")
	}
	_result := &_BACnetConstructedDataAccessDoorPresentValue{
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

// BACnetConstructedDataAccessDoorPresentValueBuilder is a builder for BACnetConstructedDataAccessDoorPresentValue
type BACnetConstructedDataAccessDoorPresentValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(presentValue BACnetDoorValueTagged) BACnetConstructedDataAccessDoorPresentValueBuilder
	// WithPresentValue adds PresentValue (property field)
	WithPresentValue(BACnetDoorValueTagged) BACnetConstructedDataAccessDoorPresentValueBuilder
	// WithPresentValueBuilder adds PresentValue (property field) which is build by the builder
	WithPresentValueBuilder(func(BACnetDoorValueTaggedBuilder) BACnetDoorValueTaggedBuilder) BACnetConstructedDataAccessDoorPresentValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataAccessDoorPresentValue or returns an error if something is wrong
	Build() (BACnetConstructedDataAccessDoorPresentValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAccessDoorPresentValue
}

// NewBACnetConstructedDataAccessDoorPresentValueBuilder() creates a BACnetConstructedDataAccessDoorPresentValueBuilder
func NewBACnetConstructedDataAccessDoorPresentValueBuilder() BACnetConstructedDataAccessDoorPresentValueBuilder {
	return &_BACnetConstructedDataAccessDoorPresentValueBuilder{_BACnetConstructedDataAccessDoorPresentValue: new(_BACnetConstructedDataAccessDoorPresentValue)}
}

type _BACnetConstructedDataAccessDoorPresentValueBuilder struct {
	*_BACnetConstructedDataAccessDoorPresentValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAccessDoorPresentValueBuilder) = (*_BACnetConstructedDataAccessDoorPresentValueBuilder)(nil)

func (b *_BACnetConstructedDataAccessDoorPresentValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataAccessDoorPresentValue
}

func (b *_BACnetConstructedDataAccessDoorPresentValueBuilder) WithMandatoryFields(presentValue BACnetDoorValueTagged) BACnetConstructedDataAccessDoorPresentValueBuilder {
	return b.WithPresentValue(presentValue)
}

func (b *_BACnetConstructedDataAccessDoorPresentValueBuilder) WithPresentValue(presentValue BACnetDoorValueTagged) BACnetConstructedDataAccessDoorPresentValueBuilder {
	b.PresentValue = presentValue
	return b
}

func (b *_BACnetConstructedDataAccessDoorPresentValueBuilder) WithPresentValueBuilder(builderSupplier func(BACnetDoorValueTaggedBuilder) BACnetDoorValueTaggedBuilder) BACnetConstructedDataAccessDoorPresentValueBuilder {
	builder := builderSupplier(b.PresentValue.CreateBACnetDoorValueTaggedBuilder())
	var err error
	b.PresentValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDoorValueTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAccessDoorPresentValueBuilder) Build() (BACnetConstructedDataAccessDoorPresentValue, error) {
	if b.PresentValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'presentValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAccessDoorPresentValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataAccessDoorPresentValueBuilder) MustBuild() BACnetConstructedDataAccessDoorPresentValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataAccessDoorPresentValueBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAccessDoorPresentValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAccessDoorPresentValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAccessDoorPresentValueBuilder().(*_BACnetConstructedDataAccessDoorPresentValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAccessDoorPresentValueBuilder creates a BACnetConstructedDataAccessDoorPresentValueBuilder
func (b *_BACnetConstructedDataAccessDoorPresentValue) CreateBACnetConstructedDataAccessDoorPresentValueBuilder() BACnetConstructedDataAccessDoorPresentValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataAccessDoorPresentValueBuilder()
	}
	return &_BACnetConstructedDataAccessDoorPresentValueBuilder{_BACnetConstructedDataAccessDoorPresentValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessDoorPresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCESS_DOOR
}

func (m *_BACnetConstructedDataAccessDoorPresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessDoorPresentValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessDoorPresentValue) GetPresentValue() BACnetDoorValueTagged {
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

func (m *_BACnetConstructedDataAccessDoorPresentValue) GetActualValue() BACnetDoorValueTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDoorValueTagged(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessDoorPresentValue(structType any) BACnetConstructedDataAccessDoorPresentValue {
	if casted, ok := structType.(BACnetConstructedDataAccessDoorPresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessDoorPresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessDoorPresentValue) GetTypeName() string {
	return "BACnetConstructedDataAccessDoorPresentValue"
}

func (m *_BACnetConstructedDataAccessDoorPresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessDoorPresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccessDoorPresentValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccessDoorPresentValue BACnetConstructedDataAccessDoorPresentValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessDoorPresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessDoorPresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	presentValue, err := ReadSimpleField[BACnetDoorValueTagged](ctx, "presentValue", ReadComplex[BACnetDoorValueTagged](BACnetDoorValueTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	actualValue, err := ReadVirtualField[BACnetDoorValueTagged](ctx, "actualValue", (*BACnetDoorValueTagged)(nil), presentValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessDoorPresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessDoorPresentValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccessDoorPresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessDoorPresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessDoorPresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessDoorPresentValue")
		}

		if err := WriteSimpleField[BACnetDoorValueTagged](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetDoorValueTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessDoorPresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessDoorPresentValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessDoorPresentValue) IsBACnetConstructedDataAccessDoorPresentValue() {
}

func (m *_BACnetConstructedDataAccessDoorPresentValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAccessDoorPresentValue) deepCopy() *_BACnetConstructedDataAccessDoorPresentValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAccessDoorPresentValueCopy := &_BACnetConstructedDataAccessDoorPresentValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetDoorValueTagged](m.PresentValue),
	}
	_BACnetConstructedDataAccessDoorPresentValueCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAccessDoorPresentValueCopy
}

func (m *_BACnetConstructedDataAccessDoorPresentValue) String() string {
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