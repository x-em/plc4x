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

// LDataFrameACK is the corresponding interface of LDataFrameACK
type LDataFrameACK interface {
	utils.LengthAware
	utils.Serializable
	LDataFrame
}

// LDataFrameACKExactly can be used when we want exactly this type and not a type which fulfills LDataFrameACK.
// This is useful for switch cases.
type LDataFrameACKExactly interface {
	LDataFrameACK
	isLDataFrameACK() bool
}

// _LDataFrameACK is the data-structure of this message
type _LDataFrameACK struct {
	*_LDataFrame
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LDataFrameACK) GetNotAckFrame() bool {
	return bool(false)
}

func (m *_LDataFrameACK) GetPolling() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LDataFrameACK) InitializeParent(parent LDataFrame, frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) {
	m.FrameType = frameType
	m.NotRepeated = notRepeated
	m.Priority = priority
	m.AcknowledgeRequested = acknowledgeRequested
	m.ErrorFlag = errorFlag
}

func (m *_LDataFrameACK) GetParent() LDataFrame {
	return m._LDataFrame
}

// NewLDataFrameACK factory function for _LDataFrameACK
func NewLDataFrameACK(frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) *_LDataFrameACK {
	_result := &_LDataFrameACK{
		_LDataFrame: NewLDataFrame(frameType, notRepeated, priority, acknowledgeRequested, errorFlag),
	}
	_result._LDataFrame._LDataFrameChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLDataFrameACK(structType interface{}) LDataFrameACK {
	if casted, ok := structType.(LDataFrameACK); ok {
		return casted
	}
	if casted, ok := structType.(*LDataFrameACK); ok {
		return *casted
	}
	return nil
}

func (m *_LDataFrameACK) GetTypeName() string {
	return "LDataFrameACK"
}

func (m *_LDataFrameACK) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_LDataFrameACK) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_LDataFrameACK) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func LDataFrameACKParse(theBytes []byte) (LDataFrameACK, error) {
	return LDataFrameACKParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func LDataFrameACKParseWithBuffer(readBuffer utils.ReadBuffer) (LDataFrameACK, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LDataFrameACK"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LDataFrameACK")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("LDataFrameACK"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LDataFrameACK")
	}

	// Create a partially initialized instance
	_child := &_LDataFrameACK{
		_LDataFrame: &_LDataFrame{},
	}
	_child._LDataFrame._LDataFrameChildRequirements = _child
	return _child, nil
}

func (m *_LDataFrameACK) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LDataFrameACK) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LDataFrameACK"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LDataFrameACK")
		}

		if popErr := writeBuffer.PopContext("LDataFrameACK"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LDataFrameACK")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_LDataFrameACK) isLDataFrameACK() bool {
	return true
}

func (m *_LDataFrameACK) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
