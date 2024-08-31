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

// StatusRequestLevel is the corresponding interface of StatusRequestLevel
type StatusRequestLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	StatusRequest
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
	// GetStartingGroupAddressLabel returns StartingGroupAddressLabel (property field)
	GetStartingGroupAddressLabel() byte
}

// StatusRequestLevelExactly can be used when we want exactly this type and not a type which fulfills StatusRequestLevel.
// This is useful for switch cases.
type StatusRequestLevelExactly interface {
	StatusRequestLevel
	isStatusRequestLevel() bool
}

// _StatusRequestLevel is the data-structure of this message
type _StatusRequestLevel struct {
	*_StatusRequest
	Application               ApplicationIdContainer
	StartingGroupAddressLabel byte
	// Reserved Fields
	reservedField0 *byte
	reservedField1 *byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_StatusRequestLevel) InitializeParent(parent StatusRequest, statusType byte) {
	m.StatusType = statusType
}

func (m *_StatusRequestLevel) GetParent() StatusRequest {
	return m._StatusRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_StatusRequestLevel) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *_StatusRequestLevel) GetStartingGroupAddressLabel() byte {
	return m.StartingGroupAddressLabel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewStatusRequestLevel factory function for _StatusRequestLevel
func NewStatusRequestLevel(application ApplicationIdContainer, startingGroupAddressLabel byte, statusType byte) *_StatusRequestLevel {
	_result := &_StatusRequestLevel{
		Application:               application,
		StartingGroupAddressLabel: startingGroupAddressLabel,
		_StatusRequest:            NewStatusRequest(statusType),
	}
	_result._StatusRequest._StatusRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastStatusRequestLevel(structType any) StatusRequestLevel {
	if casted, ok := structType.(StatusRequestLevel); ok {
		return casted
	}
	if casted, ok := structType.(*StatusRequestLevel); ok {
		return *casted
	}
	return nil
}

func (m *_StatusRequestLevel) GetTypeName() string {
	return "StatusRequestLevel"
}

func (m *_StatusRequestLevel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (application)
	lengthInBits += 8

	// Simple field (startingGroupAddressLabel)
	lengthInBits += 8

	return lengthInBits
}

func (m *_StatusRequestLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func StatusRequestLevelParse(ctx context.Context, theBytes []byte) (StatusRequestLevel, error) {
	return StatusRequestLevelParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func StatusRequestLevelParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (StatusRequestLevel, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("StatusRequestLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for StatusRequestLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadByte(readBuffer, 8), byte(0x73))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadByte(readBuffer, 8), byte(0x07))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	// Simple Field (application)
	if pullErr := readBuffer.PullContext("application"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for application")
	}
	_application, _applicationErr := ApplicationIdContainerParseWithBuffer(ctx, readBuffer)
	if _applicationErr != nil {
		return nil, errors.Wrap(_applicationErr, "Error parsing 'application' field of StatusRequestLevel")
	}
	application := _application
	if closeErr := readBuffer.CloseContext("application"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for application")
	}

	// Simple Field (startingGroupAddressLabel)
	_startingGroupAddressLabel, _startingGroupAddressLabelErr := /*TODO: migrate me*/ readBuffer.ReadByte("startingGroupAddressLabel")
	if _startingGroupAddressLabelErr != nil {
		return nil, errors.Wrap(_startingGroupAddressLabelErr, "Error parsing 'startingGroupAddressLabel' field of StatusRequestLevel")
	}
	startingGroupAddressLabel := _startingGroupAddressLabel

	// Validation
	if !(bool(bool(bool(bool(bool(bool(bool(bool((startingGroupAddressLabel) == (0x00))) || bool(bool((startingGroupAddressLabel) == (0x20)))) || bool(bool((startingGroupAddressLabel) == (0x40)))) || bool(bool((startingGroupAddressLabel) == (0x60)))) || bool(bool((startingGroupAddressLabel) == (0x80)))) || bool(bool((startingGroupAddressLabel) == (0xA0)))) || bool(bool((startingGroupAddressLabel) == (0xC0)))) || bool(bool((startingGroupAddressLabel) == (0xE0)))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "invalid label"})
	}

	if closeErr := readBuffer.CloseContext("StatusRequestLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for StatusRequestLevel")
	}

	// Create a partially initialized instance
	_child := &_StatusRequestLevel{
		_StatusRequest:            &_StatusRequest{},
		Application:               application,
		StartingGroupAddressLabel: startingGroupAddressLabel,
		reservedField0:            reservedField0,
		reservedField1:            reservedField1,
	}
	_child._StatusRequest._StatusRequestChildRequirements = _child
	return _child, nil
}

func (m *_StatusRequestLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_StatusRequestLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("StatusRequestLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for StatusRequestLevel")
		}

		// Reserved Field (reserved)
		{
			var reserved byte = byte(0x73)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": byte(0x73),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := /*TODO: migrate me*/ writeBuffer.WriteByte("reserved", reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			var reserved byte = byte(0x07)
			if m.reservedField1 != nil {
				log.Info().Fields(map[string]any{
					"expected value": byte(0x07),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField1
			}
			_err := /*TODO: migrate me*/ writeBuffer.WriteByte("reserved", reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (application)
		if pushErr := writeBuffer.PushContext("application"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for application")
		}
		_applicationErr := writeBuffer.WriteSerializable(ctx, m.GetApplication())
		if popErr := writeBuffer.PopContext("application"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for application")
		}
		if _applicationErr != nil {
			return errors.Wrap(_applicationErr, "Error serializing 'application' field")
		}

		// Simple Field (startingGroupAddressLabel)
		startingGroupAddressLabel := byte(m.GetStartingGroupAddressLabel())
		_startingGroupAddressLabelErr := /*TODO: migrate me*/ writeBuffer.WriteByte("startingGroupAddressLabel", (startingGroupAddressLabel))
		if _startingGroupAddressLabelErr != nil {
			return errors.Wrap(_startingGroupAddressLabelErr, "Error serializing 'startingGroupAddressLabel' field")
		}

		if popErr := writeBuffer.PopContext("StatusRequestLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for StatusRequestLevel")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_StatusRequestLevel) isStatusRequestLevel() bool {
	return true
}

func (m *_StatusRequestLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
