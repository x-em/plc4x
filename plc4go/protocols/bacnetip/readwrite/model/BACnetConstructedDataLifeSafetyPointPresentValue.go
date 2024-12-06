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

// BACnetConstructedDataLifeSafetyPointPresentValue is the corresponding interface of BACnetConstructedDataLifeSafetyPointPresentValue
type BACnetConstructedDataLifeSafetyPointPresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetLifeSafetyStateTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLifeSafetyStateTagged
	// IsBACnetConstructedDataLifeSafetyPointPresentValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLifeSafetyPointPresentValue()
	// CreateBuilder creates a BACnetConstructedDataLifeSafetyPointPresentValueBuilder
	CreateBACnetConstructedDataLifeSafetyPointPresentValueBuilder() BACnetConstructedDataLifeSafetyPointPresentValueBuilder
}

// _BACnetConstructedDataLifeSafetyPointPresentValue is the data-structure of this message
type _BACnetConstructedDataLifeSafetyPointPresentValue struct {
	BACnetConstructedDataContract
	PresentValue BACnetLifeSafetyStateTagged
}

var _ BACnetConstructedDataLifeSafetyPointPresentValue = (*_BACnetConstructedDataLifeSafetyPointPresentValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLifeSafetyPointPresentValue)(nil)

// NewBACnetConstructedDataLifeSafetyPointPresentValue factory function for _BACnetConstructedDataLifeSafetyPointPresentValue
func NewBACnetConstructedDataLifeSafetyPointPresentValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, presentValue BACnetLifeSafetyStateTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLifeSafetyPointPresentValue {
	if presentValue == nil {
		panic("presentValue of type BACnetLifeSafetyStateTagged for BACnetConstructedDataLifeSafetyPointPresentValue must not be nil")
	}
	_result := &_BACnetConstructedDataLifeSafetyPointPresentValue{
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

// BACnetConstructedDataLifeSafetyPointPresentValueBuilder is a builder for BACnetConstructedDataLifeSafetyPointPresentValue
type BACnetConstructedDataLifeSafetyPointPresentValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(presentValue BACnetLifeSafetyStateTagged) BACnetConstructedDataLifeSafetyPointPresentValueBuilder
	// WithPresentValue adds PresentValue (property field)
	WithPresentValue(BACnetLifeSafetyStateTagged) BACnetConstructedDataLifeSafetyPointPresentValueBuilder
	// WithPresentValueBuilder adds PresentValue (property field) which is build by the builder
	WithPresentValueBuilder(func(BACnetLifeSafetyStateTaggedBuilder) BACnetLifeSafetyStateTaggedBuilder) BACnetConstructedDataLifeSafetyPointPresentValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLifeSafetyPointPresentValue or returns an error if something is wrong
	Build() (BACnetConstructedDataLifeSafetyPointPresentValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLifeSafetyPointPresentValue
}

// NewBACnetConstructedDataLifeSafetyPointPresentValueBuilder() creates a BACnetConstructedDataLifeSafetyPointPresentValueBuilder
func NewBACnetConstructedDataLifeSafetyPointPresentValueBuilder() BACnetConstructedDataLifeSafetyPointPresentValueBuilder {
	return &_BACnetConstructedDataLifeSafetyPointPresentValueBuilder{_BACnetConstructedDataLifeSafetyPointPresentValue: new(_BACnetConstructedDataLifeSafetyPointPresentValue)}
}

type _BACnetConstructedDataLifeSafetyPointPresentValueBuilder struct {
	*_BACnetConstructedDataLifeSafetyPointPresentValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLifeSafetyPointPresentValueBuilder) = (*_BACnetConstructedDataLifeSafetyPointPresentValueBuilder)(nil)

func (b *_BACnetConstructedDataLifeSafetyPointPresentValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataLifeSafetyPointPresentValue
}

func (b *_BACnetConstructedDataLifeSafetyPointPresentValueBuilder) WithMandatoryFields(presentValue BACnetLifeSafetyStateTagged) BACnetConstructedDataLifeSafetyPointPresentValueBuilder {
	return b.WithPresentValue(presentValue)
}

func (b *_BACnetConstructedDataLifeSafetyPointPresentValueBuilder) WithPresentValue(presentValue BACnetLifeSafetyStateTagged) BACnetConstructedDataLifeSafetyPointPresentValueBuilder {
	b.PresentValue = presentValue
	return b
}

func (b *_BACnetConstructedDataLifeSafetyPointPresentValueBuilder) WithPresentValueBuilder(builderSupplier func(BACnetLifeSafetyStateTaggedBuilder) BACnetLifeSafetyStateTaggedBuilder) BACnetConstructedDataLifeSafetyPointPresentValueBuilder {
	builder := builderSupplier(b.PresentValue.CreateBACnetLifeSafetyStateTaggedBuilder())
	var err error
	b.PresentValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLifeSafetyStateTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLifeSafetyPointPresentValueBuilder) Build() (BACnetConstructedDataLifeSafetyPointPresentValue, error) {
	if b.PresentValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'presentValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLifeSafetyPointPresentValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataLifeSafetyPointPresentValueBuilder) MustBuild() BACnetConstructedDataLifeSafetyPointPresentValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLifeSafetyPointPresentValueBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLifeSafetyPointPresentValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLifeSafetyPointPresentValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLifeSafetyPointPresentValueBuilder().(*_BACnetConstructedDataLifeSafetyPointPresentValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLifeSafetyPointPresentValueBuilder creates a BACnetConstructedDataLifeSafetyPointPresentValueBuilder
func (b *_BACnetConstructedDataLifeSafetyPointPresentValue) CreateBACnetConstructedDataLifeSafetyPointPresentValueBuilder() BACnetConstructedDataLifeSafetyPointPresentValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataLifeSafetyPointPresentValueBuilder()
	}
	return &_BACnetConstructedDataLifeSafetyPointPresentValueBuilder{_BACnetConstructedDataLifeSafetyPointPresentValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LIFE_SAFETY_POINT
}

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) GetPresentValue() BACnetLifeSafetyStateTagged {
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

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) GetActualValue() BACnetLifeSafetyStateTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLifeSafetyStateTagged(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLifeSafetyPointPresentValue(structType any) BACnetConstructedDataLifeSafetyPointPresentValue {
	if casted, ok := structType.(BACnetConstructedDataLifeSafetyPointPresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLifeSafetyPointPresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) GetTypeName() string {
	return "BACnetConstructedDataLifeSafetyPointPresentValue"
}

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLifeSafetyPointPresentValue BACnetConstructedDataLifeSafetyPointPresentValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLifeSafetyPointPresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLifeSafetyPointPresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	presentValue, err := ReadSimpleField[BACnetLifeSafetyStateTagged](ctx, "presentValue", ReadComplex[BACnetLifeSafetyStateTagged](BACnetLifeSafetyStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	actualValue, err := ReadVirtualField[BACnetLifeSafetyStateTagged](ctx, "actualValue", (*BACnetLifeSafetyStateTagged)(nil), presentValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLifeSafetyPointPresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLifeSafetyPointPresentValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLifeSafetyPointPresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLifeSafetyPointPresentValue")
		}

		if err := WriteSimpleField[BACnetLifeSafetyStateTagged](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetLifeSafetyStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLifeSafetyPointPresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLifeSafetyPointPresentValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) IsBACnetConstructedDataLifeSafetyPointPresentValue() {
}

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) deepCopy() *_BACnetConstructedDataLifeSafetyPointPresentValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLifeSafetyPointPresentValueCopy := &_BACnetConstructedDataLifeSafetyPointPresentValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetLifeSafetyStateTagged](m.PresentValue),
	}
	_BACnetConstructedDataLifeSafetyPointPresentValueCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLifeSafetyPointPresentValueCopy
}

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) String() string {
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