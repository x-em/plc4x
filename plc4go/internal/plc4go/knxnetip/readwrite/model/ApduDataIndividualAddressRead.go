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
type ApduDataIndividualAddressRead struct {
	Parent *ApduData
}

// The corresponding interface
type IApduDataIndividualAddressRead interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataIndividualAddressRead) ApciType() uint8 {
	return 0x4
}

func (m *ApduDataIndividualAddressRead) InitializeParent(parent *ApduData) {
}

func NewApduDataIndividualAddressRead() *ApduData {
	child := &ApduDataIndividualAddressRead{
		Parent: NewApduData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduDataIndividualAddressRead(structType interface{}) *ApduDataIndividualAddressRead {
	castFunc := func(typ interface{}) *ApduDataIndividualAddressRead {
		if casted, ok := typ.(ApduDataIndividualAddressRead); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataIndividualAddressRead); ok {
			return casted
		}
		if casted, ok := typ.(ApduData); ok {
			return CastApduDataIndividualAddressRead(casted.Child)
		}
		if casted, ok := typ.(*ApduData); ok {
			return CastApduDataIndividualAddressRead(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataIndividualAddressRead) GetTypeName() string {
	return "ApduDataIndividualAddressRead"
}

func (m *ApduDataIndividualAddressRead) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataIndividualAddressRead) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataIndividualAddressRead) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataIndividualAddressReadParse(readBuffer utils.ReadBuffer, dataLength uint8) (*ApduData, error) {
	if pullErr := readBuffer.PullContext("ApduDataIndividualAddressRead"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("ApduDataIndividualAddressRead"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataIndividualAddressRead{
		Parent: &ApduData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduDataIndividualAddressRead) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataIndividualAddressRead"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataIndividualAddressRead"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataIndividualAddressRead) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
