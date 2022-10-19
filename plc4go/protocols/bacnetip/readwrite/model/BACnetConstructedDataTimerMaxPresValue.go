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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataTimerMaxPresValue is the corresponding interface of BACnetConstructedDataTimerMaxPresValue
type BACnetConstructedDataTimerMaxPresValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMaxPresValue returns MaxPresValue (property field)
	GetMaxPresValue() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataTimerMaxPresValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataTimerMaxPresValue.
// This is useful for switch cases.
type BACnetConstructedDataTimerMaxPresValueExactly interface {
	BACnetConstructedDataTimerMaxPresValue
	isBACnetConstructedDataTimerMaxPresValue() bool
}

// _BACnetConstructedDataTimerMaxPresValue is the data-structure of this message
type _BACnetConstructedDataTimerMaxPresValue struct {
	*_BACnetConstructedData
	MaxPresValue BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimerMaxPresValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TIMER
}

func (m *_BACnetConstructedDataTimerMaxPresValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_PRES_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimerMaxPresValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataTimerMaxPresValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimerMaxPresValue) GetMaxPresValue() BACnetApplicationTagUnsignedInteger {
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

func (m *_BACnetConstructedDataTimerMaxPresValue) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetMaxPresValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTimerMaxPresValue factory function for _BACnetConstructedDataTimerMaxPresValue
func NewBACnetConstructedDataTimerMaxPresValue(maxPresValue BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimerMaxPresValue {
	_result := &_BACnetConstructedDataTimerMaxPresValue{
		MaxPresValue:           maxPresValue,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimerMaxPresValue(structType interface{}) BACnetConstructedDataTimerMaxPresValue {
	if casted, ok := structType.(BACnetConstructedDataTimerMaxPresValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimerMaxPresValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimerMaxPresValue) GetTypeName() string {
	return "BACnetConstructedDataTimerMaxPresValue"
}

func (m *_BACnetConstructedDataTimerMaxPresValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataTimerMaxPresValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (maxPresValue)
	lengthInBits += m.MaxPresValue.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimerMaxPresValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTimerMaxPresValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTimerMaxPresValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimerMaxPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimerMaxPresValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maxPresValue)
	if pullErr := readBuffer.PullContext("maxPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maxPresValue")
	}
	_maxPresValue, _maxPresValueErr := BACnetApplicationTagParse(readBuffer)
	if _maxPresValueErr != nil {
		return nil, errors.Wrap(_maxPresValueErr, "Error parsing 'maxPresValue' field of BACnetConstructedDataTimerMaxPresValue")
	}
	maxPresValue := _maxPresValue.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("maxPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maxPresValue")
	}

	// Virtual field
	_actualValue := maxPresValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimerMaxPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimerMaxPresValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataTimerMaxPresValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		MaxPresValue: maxPresValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataTimerMaxPresValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimerMaxPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimerMaxPresValue")
		}

		// Simple Field (maxPresValue)
		if pushErr := writeBuffer.PushContext("maxPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maxPresValue")
		}
		_maxPresValueErr := writeBuffer.WriteSerializable(m.GetMaxPresValue())
		if popErr := writeBuffer.PopContext("maxPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maxPresValue")
		}
		if _maxPresValueErr != nil {
			return errors.Wrap(_maxPresValueErr, "Error serializing 'maxPresValue' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimerMaxPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimerMaxPresValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimerMaxPresValue) isBACnetConstructedDataTimerMaxPresValue() bool {
	return true
}

func (m *_BACnetConstructedDataTimerMaxPresValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
