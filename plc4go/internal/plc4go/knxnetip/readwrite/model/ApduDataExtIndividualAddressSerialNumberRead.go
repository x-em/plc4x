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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type ApduDataExtIndividualAddressSerialNumberRead struct {
	Parent *ApduDataExt
}

// The corresponding interface
type IApduDataExtIndividualAddressSerialNumberRead interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtIndividualAddressSerialNumberRead) ExtApciType() uint8 {
	return 0x1C
}

func (m *ApduDataExtIndividualAddressSerialNumberRead) InitializeParent(parent *ApduDataExt) {
}

func NewApduDataExtIndividualAddressSerialNumberRead() *ApduDataExt {
	child := &ApduDataExtIndividualAddressSerialNumberRead{
		Parent: NewApduDataExt(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduDataExtIndividualAddressSerialNumberRead(structType interface{}) *ApduDataExtIndividualAddressSerialNumberRead {
	castFunc := func(typ interface{}) *ApduDataExtIndividualAddressSerialNumberRead {
		if casted, ok := typ.(ApduDataExtIndividualAddressSerialNumberRead); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataExtIndividualAddressSerialNumberRead); ok {
			return casted
		}
		if casted, ok := typ.(ApduDataExt); ok {
			return CastApduDataExtIndividualAddressSerialNumberRead(casted.Child)
		}
		if casted, ok := typ.(*ApduDataExt); ok {
			return CastApduDataExtIndividualAddressSerialNumberRead(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataExtIndividualAddressSerialNumberRead) GetTypeName() string {
	return "ApduDataExtIndividualAddressSerialNumberRead"
}

func (m *ApduDataExtIndividualAddressSerialNumberRead) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataExtIndividualAddressSerialNumberRead) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtIndividualAddressSerialNumberRead) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataExtIndividualAddressSerialNumberReadParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExt, error) {
	if pullErr := readBuffer.PullContext("ApduDataExtIndividualAddressSerialNumberRead"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("ApduDataExtIndividualAddressSerialNumberRead"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataExtIndividualAddressSerialNumberRead{
		Parent: &ApduDataExt{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduDataExtIndividualAddressSerialNumberRead) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtIndividualAddressSerialNumberRead"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataExtIndividualAddressSerialNumberRead"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtIndividualAddressSerialNumberRead) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
