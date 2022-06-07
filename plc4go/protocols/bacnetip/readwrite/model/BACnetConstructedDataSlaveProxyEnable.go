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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataSlaveProxyEnable is the data-structure of this message
type BACnetConstructedDataSlaveProxyEnable struct {
	*BACnetConstructedData
	SlaveProxyEnable *BACnetApplicationTagBoolean

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataSlaveProxyEnable is the corresponding interface of BACnetConstructedDataSlaveProxyEnable
type IBACnetConstructedDataSlaveProxyEnable interface {
	IBACnetConstructedData
	// GetSlaveProxyEnable returns SlaveProxyEnable (property field)
	GetSlaveProxyEnable() *BACnetApplicationTagBoolean
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataSlaveProxyEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataSlaveProxyEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SLAVE_PROXY_ENABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataSlaveProxyEnable) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataSlaveProxyEnable) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataSlaveProxyEnable) GetSlaveProxyEnable() *BACnetApplicationTagBoolean {
	return m.SlaveProxyEnable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSlaveProxyEnable factory function for BACnetConstructedDataSlaveProxyEnable
func NewBACnetConstructedDataSlaveProxyEnable(slaveProxyEnable *BACnetApplicationTagBoolean, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataSlaveProxyEnable {
	_result := &BACnetConstructedDataSlaveProxyEnable{
		SlaveProxyEnable:      slaveProxyEnable,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataSlaveProxyEnable(structType interface{}) *BACnetConstructedDataSlaveProxyEnable {
	if casted, ok := structType.(BACnetConstructedDataSlaveProxyEnable); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSlaveProxyEnable); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataSlaveProxyEnable(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataSlaveProxyEnable(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataSlaveProxyEnable) GetTypeName() string {
	return "BACnetConstructedDataSlaveProxyEnable"
}

func (m *BACnetConstructedDataSlaveProxyEnable) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataSlaveProxyEnable) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (slaveProxyEnable)
	lengthInBits += m.SlaveProxyEnable.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataSlaveProxyEnable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataSlaveProxyEnableParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataSlaveProxyEnable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSlaveProxyEnable"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (slaveProxyEnable)
	if pullErr := readBuffer.PullContext("slaveProxyEnable"); pullErr != nil {
		return nil, pullErr
	}
	_slaveProxyEnable, _slaveProxyEnableErr := BACnetApplicationTagParse(readBuffer)
	if _slaveProxyEnableErr != nil {
		return nil, errors.Wrap(_slaveProxyEnableErr, "Error parsing 'slaveProxyEnable' field")
	}
	slaveProxyEnable := CastBACnetApplicationTagBoolean(_slaveProxyEnable)
	if closeErr := readBuffer.CloseContext("slaveProxyEnable"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSlaveProxyEnable"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataSlaveProxyEnable{
		SlaveProxyEnable:      CastBACnetApplicationTagBoolean(slaveProxyEnable),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataSlaveProxyEnable) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSlaveProxyEnable"); pushErr != nil {
			return pushErr
		}

		// Simple Field (slaveProxyEnable)
		if pushErr := writeBuffer.PushContext("slaveProxyEnable"); pushErr != nil {
			return pushErr
		}
		_slaveProxyEnableErr := m.SlaveProxyEnable.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("slaveProxyEnable"); popErr != nil {
			return popErr
		}
		if _slaveProxyEnableErr != nil {
			return errors.Wrap(_slaveProxyEnableErr, "Error serializing 'slaveProxyEnable' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSlaveProxyEnable"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataSlaveProxyEnable) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
