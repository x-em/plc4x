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

// BACnetLogDataLogDataEntryBitStringValue is the corresponding interface of BACnetLogDataLogDataEntryBitStringValue
type BACnetLogDataLogDataEntryBitStringValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetLogDataLogDataEntry
	// GetBitStringValue returns BitStringValue (property field)
	GetBitStringValue() BACnetContextTagBitString
}

// BACnetLogDataLogDataEntryBitStringValueExactly can be used when we want exactly this type and not a type which fulfills BACnetLogDataLogDataEntryBitStringValue.
// This is useful for switch cases.
type BACnetLogDataLogDataEntryBitStringValueExactly interface {
	BACnetLogDataLogDataEntryBitStringValue
	isBACnetLogDataLogDataEntryBitStringValue() bool
}

// _BACnetLogDataLogDataEntryBitStringValue is the data-structure of this message
type _BACnetLogDataLogDataEntryBitStringValue struct {
	*_BACnetLogDataLogDataEntry
	BitStringValue BACnetContextTagBitString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogDataLogDataEntryBitStringValue) InitializeParent(parent BACnetLogDataLogDataEntry, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetLogDataLogDataEntryBitStringValue) GetParent() BACnetLogDataLogDataEntry {
	return m._BACnetLogDataLogDataEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogDataEntryBitStringValue) GetBitStringValue() BACnetContextTagBitString {
	return m.BitStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogDataLogDataEntryBitStringValue factory function for _BACnetLogDataLogDataEntryBitStringValue
func NewBACnetLogDataLogDataEntryBitStringValue(bitStringValue BACnetContextTagBitString, peekedTagHeader BACnetTagHeader) *_BACnetLogDataLogDataEntryBitStringValue {
	_result := &_BACnetLogDataLogDataEntryBitStringValue{
		BitStringValue:             bitStringValue,
		_BACnetLogDataLogDataEntry: NewBACnetLogDataLogDataEntry(peekedTagHeader),
	}
	_result._BACnetLogDataLogDataEntry._BACnetLogDataLogDataEntryChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogDataEntryBitStringValue(structType interface{}) BACnetLogDataLogDataEntryBitStringValue {
	if casted, ok := structType.(BACnetLogDataLogDataEntryBitStringValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntryBitStringValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogDataEntryBitStringValue) GetTypeName() string {
	return "BACnetLogDataLogDataEntryBitStringValue"
}

func (m *_BACnetLogDataLogDataEntryBitStringValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetLogDataLogDataEntryBitStringValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (bitStringValue)
	lengthInBits += m.BitStringValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetLogDataLogDataEntryBitStringValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogDataLogDataEntryBitStringValueParse(readBuffer utils.ReadBuffer) (BACnetLogDataLogDataEntryBitStringValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataEntryBitStringValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataEntryBitStringValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (bitStringValue)
	if pullErr := readBuffer.PullContext("bitStringValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bitStringValue")
	}
	_bitStringValue, _bitStringValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(5)), BACnetDataType(BACnetDataType_BIT_STRING))
	if _bitStringValueErr != nil {
		return nil, errors.Wrap(_bitStringValueErr, "Error parsing 'bitStringValue' field of BACnetLogDataLogDataEntryBitStringValue")
	}
	bitStringValue := _bitStringValue.(BACnetContextTagBitString)
	if closeErr := readBuffer.CloseContext("bitStringValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bitStringValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataEntryBitStringValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataEntryBitStringValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetLogDataLogDataEntryBitStringValue{
		_BACnetLogDataLogDataEntry: &_BACnetLogDataLogDataEntry{},
		BitStringValue:             bitStringValue,
	}
	_child._BACnetLogDataLogDataEntry._BACnetLogDataLogDataEntryChildRequirements = _child
	return _child, nil
}

func (m *_BACnetLogDataLogDataEntryBitStringValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataEntryBitStringValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataEntryBitStringValue")
		}

		// Simple Field (bitStringValue)
		if pushErr := writeBuffer.PushContext("bitStringValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for bitStringValue")
		}
		_bitStringValueErr := writeBuffer.WriteSerializable(m.GetBitStringValue())
		if popErr := writeBuffer.PopContext("bitStringValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for bitStringValue")
		}
		if _bitStringValueErr != nil {
			return errors.Wrap(_bitStringValueErr, "Error serializing 'bitStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataEntryBitStringValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataEntryBitStringValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogDataEntryBitStringValue) isBACnetLogDataLogDataEntryBitStringValue() bool {
	return true
}

func (m *_BACnetLogDataLogDataEntryBitStringValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
