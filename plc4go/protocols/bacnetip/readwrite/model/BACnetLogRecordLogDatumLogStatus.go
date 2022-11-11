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

// BACnetLogRecordLogDatumLogStatus is the corresponding interface of BACnetLogRecordLogDatumLogStatus
type BACnetLogRecordLogDatumLogStatus interface {
	utils.LengthAware
	utils.Serializable
	BACnetLogRecordLogDatum
	// GetLogStatus returns LogStatus (property field)
	GetLogStatus() BACnetLogStatusTagged
}

// BACnetLogRecordLogDatumLogStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetLogRecordLogDatumLogStatus.
// This is useful for switch cases.
type BACnetLogRecordLogDatumLogStatusExactly interface {
	BACnetLogRecordLogDatumLogStatus
	isBACnetLogRecordLogDatumLogStatus() bool
}

// _BACnetLogRecordLogDatumLogStatus is the data-structure of this message
type _BACnetLogRecordLogDatumLogStatus struct {
	*_BACnetLogRecordLogDatum
	LogStatus BACnetLogStatusTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogRecordLogDatumLogStatus) InitializeParent(parent BACnetLogRecordLogDatum, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetLogRecordLogDatumLogStatus) GetParent() BACnetLogRecordLogDatum {
	return m._BACnetLogRecordLogDatum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecordLogDatumLogStatus) GetLogStatus() BACnetLogStatusTagged {
	return m.LogStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogRecordLogDatumLogStatus factory function for _BACnetLogRecordLogDatumLogStatus
func NewBACnetLogRecordLogDatumLogStatus(logStatus BACnetLogStatusTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLogRecordLogDatumLogStatus {
	_result := &_BACnetLogRecordLogDatumLogStatus{
		LogStatus:                logStatus,
		_BACnetLogRecordLogDatum: NewBACnetLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetLogRecordLogDatum._BACnetLogRecordLogDatumChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogRecordLogDatumLogStatus(structType interface{}) BACnetLogRecordLogDatumLogStatus {
	if casted, ok := structType.(BACnetLogRecordLogDatumLogStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatumLogStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecordLogDatumLogStatus) GetTypeName() string {
	return "BACnetLogRecordLogDatumLogStatus"
}

func (m *_BACnetLogRecordLogDatumLogStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetLogRecordLogDatumLogStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (logStatus)
	lengthInBits += m.LogStatus.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetLogRecordLogDatumLogStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogRecordLogDatumLogStatusParse(theBytes []byte, tagNumber uint8) (BACnetLogRecordLogDatumLogStatus, error) {
	return BACnetLogRecordLogDatumLogStatusParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetLogRecordLogDatumLogStatusParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetLogRecordLogDatumLogStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatumLogStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatumLogStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (logStatus)
	if pullErr := readBuffer.PullContext("logStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for logStatus")
	}
	_logStatus, _logStatusErr := BACnetLogStatusTaggedParseWithBuffer(readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _logStatusErr != nil {
		return nil, errors.Wrap(_logStatusErr, "Error parsing 'logStatus' field of BACnetLogRecordLogDatumLogStatus")
	}
	logStatus := _logStatus.(BACnetLogStatusTagged)
	if closeErr := readBuffer.CloseContext("logStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for logStatus")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatumLogStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatumLogStatus")
	}

	// Create a partially initialized instance
	_child := &_BACnetLogRecordLogDatumLogStatus{
		_BACnetLogRecordLogDatum: &_BACnetLogRecordLogDatum{
			TagNumber: tagNumber,
		},
		LogStatus: logStatus,
	}
	_child._BACnetLogRecordLogDatum._BACnetLogRecordLogDatumChildRequirements = _child
	return _child, nil
}

func (m *_BACnetLogRecordLogDatumLogStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogRecordLogDatumLogStatus) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatumLogStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatumLogStatus")
		}

		// Simple Field (logStatus)
		if pushErr := writeBuffer.PushContext("logStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for logStatus")
		}
		_logStatusErr := writeBuffer.WriteSerializable(m.GetLogStatus())
		if popErr := writeBuffer.PopContext("logStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for logStatus")
		}
		if _logStatusErr != nil {
			return errors.Wrap(_logStatusErr, "Error serializing 'logStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatumLogStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatumLogStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetLogRecordLogDatumLogStatus) isBACnetLogRecordLogDatumLogStatus() bool {
	return true
}

func (m *_BACnetLogRecordLogDatumLogStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
