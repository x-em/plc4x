/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type UnknownMessage struct {
	UnknownData []byte
	Parent      *KnxNetIpMessage
}

// The corresponding interface
type IUnknownMessage interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *UnknownMessage) MsgType() uint16 {
	return 0x020B
}

func (m *UnknownMessage) InitializeParent(parent *KnxNetIpMessage) {
}

func NewUnknownMessage(unknownData []byte) *KnxNetIpMessage {
	child := &UnknownMessage{
		UnknownData: unknownData,
		Parent:      NewKnxNetIpMessage(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastUnknownMessage(structType interface{}) *UnknownMessage {
	castFunc := func(typ interface{}) *UnknownMessage {
		if casted, ok := typ.(UnknownMessage); ok {
			return &casted
		}
		if casted, ok := typ.(*UnknownMessage); ok {
			return casted
		}
		if casted, ok := typ.(KnxNetIpMessage); ok {
			return CastUnknownMessage(casted.Child)
		}
		if casted, ok := typ.(*KnxNetIpMessage); ok {
			return CastUnknownMessage(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *UnknownMessage) GetTypeName() string {
	return "UnknownMessage"
}

func (m *UnknownMessage) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *UnknownMessage) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Array field
	if len(m.UnknownData) > 0 {
		lengthInBits += 8 * uint16(len(m.UnknownData))
	}

	return lengthInBits
}

func (m *UnknownMessage) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func UnknownMessageParse(readBuffer utils.ReadBuffer, totalLength uint16) (*KnxNetIpMessage, error) {
	if pullErr := readBuffer.PullContext("UnknownMessage"); pullErr != nil {
		return nil, pullErr
	}
	// Byte Array field (unknownData)
	numberOfBytesunknownData := int(uint16(totalLength) - uint16(uint16(6)))
	unknownData, _readArrayErr := readBuffer.ReadByteArray("unknownData", numberOfBytesunknownData)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'unknownData' field")
	}

	if closeErr := readBuffer.CloseContext("UnknownMessage"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &UnknownMessage{
		UnknownData: unknownData,
		Parent:      &KnxNetIpMessage{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *UnknownMessage) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("UnknownMessage"); pushErr != nil {
			return pushErr
		}

		// Array Field (unknownData)
		if m.UnknownData != nil {
			// Byte Array field (unknownData)
			_writeArrayErr := writeBuffer.WriteByteArray("unknownData", m.UnknownData)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'unknownData' field")
			}
		}

		if popErr := writeBuffer.PopContext("UnknownMessage"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *UnknownMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
