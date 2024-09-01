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

// BACnetConstructedDataBitStringValueAlarmValues is the corresponding interface of BACnetConstructedDataBitStringValueAlarmValues
type BACnetConstructedDataBitStringValueAlarmValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetAlarmValues returns AlarmValues (property field)
	GetAlarmValues() []BACnetApplicationTagBitString
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataBitStringValueAlarmValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBitStringValueAlarmValues.
// This is useful for switch cases.
type BACnetConstructedDataBitStringValueAlarmValuesExactly interface {
	BACnetConstructedDataBitStringValueAlarmValues
	isBACnetConstructedDataBitStringValueAlarmValues() bool
}

// _BACnetConstructedDataBitStringValueAlarmValues is the data-structure of this message
type _BACnetConstructedDataBitStringValueAlarmValues struct {
	*_BACnetConstructedData
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	AlarmValues          []BACnetApplicationTagBitString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBitStringValueAlarmValues) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_BITSTRING_VALUE
}

func (m *_BACnetConstructedDataBitStringValueAlarmValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALARM_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBitStringValueAlarmValues) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBitStringValueAlarmValues) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBitStringValueAlarmValues) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataBitStringValueAlarmValues) GetAlarmValues() []BACnetApplicationTagBitString {
	return m.AlarmValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBitStringValueAlarmValues) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBitStringValueAlarmValues factory function for _BACnetConstructedDataBitStringValueAlarmValues
func NewBACnetConstructedDataBitStringValueAlarmValues(numberOfDataElements BACnetApplicationTagUnsignedInteger, alarmValues []BACnetApplicationTagBitString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBitStringValueAlarmValues {
	_result := &_BACnetConstructedDataBitStringValueAlarmValues{
		NumberOfDataElements:   numberOfDataElements,
		AlarmValues:            alarmValues,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBitStringValueAlarmValues(structType any) BACnetConstructedDataBitStringValueAlarmValues {
	if casted, ok := structType.(BACnetConstructedDataBitStringValueAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBitStringValueAlarmValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBitStringValueAlarmValues) GetTypeName() string {
	return "BACnetConstructedDataBitStringValueAlarmValues"
}

func (m *_BACnetConstructedDataBitStringValueAlarmValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.AlarmValues) > 0 {
		for _, element := range m.AlarmValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataBitStringValueAlarmValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataBitStringValueAlarmValuesParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBitStringValueAlarmValues, error) {
	return BACnetConstructedDataBitStringValueAlarmValuesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataBitStringValueAlarmValuesParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataBitStringValueAlarmValues, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataBitStringValueAlarmValues, error) {
		return BACnetConstructedDataBitStringValueAlarmValuesParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataBitStringValueAlarmValuesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBitStringValueAlarmValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBitStringValueAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBitStringValueAlarmValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
	}

	alarmValues, err := ReadTerminatedArrayField[BACnetApplicationTagBitString](ctx, "alarmValues", ReadComplex[BACnetApplicationTagBitString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBitString](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alarmValues' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBitStringValueAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBitStringValueAlarmValues")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBitStringValueAlarmValues{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		NumberOfDataElements: numberOfDataElements,
		AlarmValues:          alarmValues,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBitStringValueAlarmValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBitStringValueAlarmValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBitStringValueAlarmValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBitStringValueAlarmValues")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", GetRef(m.GetNumberOfDataElements()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfDataElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "alarmValues", m.GetAlarmValues(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'alarmValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBitStringValueAlarmValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBitStringValueAlarmValues")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBitStringValueAlarmValues) isBACnetConstructedDataBitStringValueAlarmValues() bool {
	return true
}

func (m *_BACnetConstructedDataBitStringValueAlarmValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
