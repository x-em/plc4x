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

// BACnetTagPayloadDouble is the corresponding interface of BACnetTagPayloadDouble
type BACnetTagPayloadDouble interface {
	utils.LengthAware
	utils.Serializable
	// GetValue returns Value (property field)
	GetValue() float64
}

// BACnetTagPayloadDoubleExactly can be used when we want exactly this type and not a type which fulfills BACnetTagPayloadDouble.
// This is useful for switch cases.
type BACnetTagPayloadDoubleExactly interface {
	BACnetTagPayloadDouble
	isBACnetTagPayloadDouble() bool
}

// _BACnetTagPayloadDouble is the data-structure of this message
type _BACnetTagPayloadDouble struct {
	Value float64
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTagPayloadDouble) GetValue() float64 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTagPayloadDouble factory function for _BACnetTagPayloadDouble
func NewBACnetTagPayloadDouble(value float64) *_BACnetTagPayloadDouble {
	return &_BACnetTagPayloadDouble{Value: value}
}

// Deprecated: use the interface for direct cast
func CastBACnetTagPayloadDouble(structType interface{}) BACnetTagPayloadDouble {
	if casted, ok := structType.(BACnetTagPayloadDouble); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagPayloadDouble); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagPayloadDouble) GetTypeName() string {
	return "BACnetTagPayloadDouble"
}

func (m *_BACnetTagPayloadDouble) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetTagPayloadDouble) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (value)
	lengthInBits += 64

	return lengthInBits
}

func (m *_BACnetTagPayloadDouble) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTagPayloadDoubleParse(theBytes []byte) (BACnetTagPayloadDouble, error) {
	return BACnetTagPayloadDoubleParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetTagPayloadDoubleParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetTagPayloadDouble, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagPayloadDouble"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadDouble")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadFloat64("value", 64)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetTagPayloadDouble")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadDouble"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadDouble")
	}

	// Create the instance
	return &_BACnetTagPayloadDouble{
		Value: value,
	}, nil
}

func (m *_BACnetTagPayloadDouble) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTagPayloadDouble) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadDouble"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadDouble")
	}

	// Simple Field (value)
	value := float64(m.GetValue())
	_valueErr := writeBuffer.WriteFloat64("value", 64, (value))
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadDouble"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadDouble")
	}
	return nil
}

func (m *_BACnetTagPayloadDouble) isBACnetTagPayloadDouble() bool {
	return true
}

func (m *_BACnetTagPayloadDouble) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
