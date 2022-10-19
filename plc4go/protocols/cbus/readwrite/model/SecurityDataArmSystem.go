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

// SecurityDataArmSystem is the corresponding interface of SecurityDataArmSystem
type SecurityDataArmSystem interface {
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetArmMode returns ArmMode (property field)
	GetArmMode() byte
	// GetIsReserved returns IsReserved (virtual field)
	GetIsReserved() bool
	// GetIsArmToAwayMode returns IsArmToAwayMode (virtual field)
	GetIsArmToAwayMode() bool
	// GetIsArmToNightMode returns IsArmToNightMode (virtual field)
	GetIsArmToNightMode() bool
	// GetIsArmToDayMode returns IsArmToDayMode (virtual field)
	GetIsArmToDayMode() bool
	// GetIsArmToVacationMode returns IsArmToVacationMode (virtual field)
	GetIsArmToVacationMode() bool
	// GetIsArmToHighestLevelOfProtection returns IsArmToHighestLevelOfProtection (virtual field)
	GetIsArmToHighestLevelOfProtection() bool
}

// SecurityDataArmSystemExactly can be used when we want exactly this type and not a type which fulfills SecurityDataArmSystem.
// This is useful for switch cases.
type SecurityDataArmSystemExactly interface {
	SecurityDataArmSystem
	isSecurityDataArmSystem() bool
}

// _SecurityDataArmSystem is the data-structure of this message
type _SecurityDataArmSystem struct {
	*_SecurityData
	ArmMode byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataArmSystem) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataArmSystem) GetParent() SecurityData {
	return m._SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataArmSystem) GetArmMode() byte {
	return m.ArmMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_SecurityDataArmSystem) GetIsReserved() bool {
	return bool(bool(bool((m.GetArmMode()) == (0x00))) || bool((bool(bool((m.GetArmMode()) >= (0x05))) && bool(bool((m.GetArmMode()) <= (0xFE))))))
}

func (m *_SecurityDataArmSystem) GetIsArmToAwayMode() bool {
	return bool(bool((m.GetArmMode()) == (0x01)))
}

func (m *_SecurityDataArmSystem) GetIsArmToNightMode() bool {
	return bool(bool((m.GetArmMode()) == (0x02)))
}

func (m *_SecurityDataArmSystem) GetIsArmToDayMode() bool {
	return bool(bool((m.GetArmMode()) == (0x03)))
}

func (m *_SecurityDataArmSystem) GetIsArmToVacationMode() bool {
	return bool(bool((m.GetArmMode()) == (0x04)))
}

func (m *_SecurityDataArmSystem) GetIsArmToHighestLevelOfProtection() bool {
	return bool(bool((m.GetArmMode()) > (0xFE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityDataArmSystem factory function for _SecurityDataArmSystem
func NewSecurityDataArmSystem(armMode byte, commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataArmSystem {
	_result := &_SecurityDataArmSystem{
		ArmMode:       armMode,
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataArmSystem(structType interface{}) SecurityDataArmSystem {
	if casted, ok := structType.(SecurityDataArmSystem); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataArmSystem); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataArmSystem) GetTypeName() string {
	return "SecurityDataArmSystem"
}

func (m *_SecurityDataArmSystem) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SecurityDataArmSystem) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (armMode)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_SecurityDataArmSystem) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SecurityDataArmSystemParse(readBuffer utils.ReadBuffer) (SecurityDataArmSystem, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataArmSystem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataArmSystem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (armMode)
	_armMode, _armModeErr := readBuffer.ReadByte("armMode")
	if _armModeErr != nil {
		return nil, errors.Wrap(_armModeErr, "Error parsing 'armMode' field of SecurityDataArmSystem")
	}
	armMode := _armMode

	// Virtual field
	_isReserved := bool(bool((armMode) == (0x00))) || bool((bool(bool((armMode) >= (0x05))) && bool(bool((armMode) <= (0xFE)))))
	isReserved := bool(_isReserved)
	_ = isReserved

	// Virtual field
	_isArmToAwayMode := bool((armMode) == (0x01))
	isArmToAwayMode := bool(_isArmToAwayMode)
	_ = isArmToAwayMode

	// Virtual field
	_isArmToNightMode := bool((armMode) == (0x02))
	isArmToNightMode := bool(_isArmToNightMode)
	_ = isArmToNightMode

	// Virtual field
	_isArmToDayMode := bool((armMode) == (0x03))
	isArmToDayMode := bool(_isArmToDayMode)
	_ = isArmToDayMode

	// Virtual field
	_isArmToVacationMode := bool((armMode) == (0x04))
	isArmToVacationMode := bool(_isArmToVacationMode)
	_ = isArmToVacationMode

	// Virtual field
	_isArmToHighestLevelOfProtection := bool((armMode) > (0xFE))
	isArmToHighestLevelOfProtection := bool(_isArmToHighestLevelOfProtection)
	_ = isArmToHighestLevelOfProtection

	if closeErr := readBuffer.CloseContext("SecurityDataArmSystem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataArmSystem")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataArmSystem{
		_SecurityData: &_SecurityData{},
		ArmMode:       armMode,
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataArmSystem) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataArmSystem"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataArmSystem")
		}

		// Simple Field (armMode)
		armMode := byte(m.GetArmMode())
		_armModeErr := writeBuffer.WriteByte("armMode", (armMode))
		if _armModeErr != nil {
			return errors.Wrap(_armModeErr, "Error serializing 'armMode' field")
		}
		// Virtual field
		if _isReservedErr := writeBuffer.WriteVirtual("isReserved", m.GetIsReserved()); _isReservedErr != nil {
			return errors.Wrap(_isReservedErr, "Error serializing 'isReserved' field")
		}
		// Virtual field
		if _isArmToAwayModeErr := writeBuffer.WriteVirtual("isArmToAwayMode", m.GetIsArmToAwayMode()); _isArmToAwayModeErr != nil {
			return errors.Wrap(_isArmToAwayModeErr, "Error serializing 'isArmToAwayMode' field")
		}
		// Virtual field
		if _isArmToNightModeErr := writeBuffer.WriteVirtual("isArmToNightMode", m.GetIsArmToNightMode()); _isArmToNightModeErr != nil {
			return errors.Wrap(_isArmToNightModeErr, "Error serializing 'isArmToNightMode' field")
		}
		// Virtual field
		if _isArmToDayModeErr := writeBuffer.WriteVirtual("isArmToDayMode", m.GetIsArmToDayMode()); _isArmToDayModeErr != nil {
			return errors.Wrap(_isArmToDayModeErr, "Error serializing 'isArmToDayMode' field")
		}
		// Virtual field
		if _isArmToVacationModeErr := writeBuffer.WriteVirtual("isArmToVacationMode", m.GetIsArmToVacationMode()); _isArmToVacationModeErr != nil {
			return errors.Wrap(_isArmToVacationModeErr, "Error serializing 'isArmToVacationMode' field")
		}
		// Virtual field
		if _isArmToHighestLevelOfProtectionErr := writeBuffer.WriteVirtual("isArmToHighestLevelOfProtection", m.GetIsArmToHighestLevelOfProtection()); _isArmToHighestLevelOfProtectionErr != nil {
			return errors.Wrap(_isArmToHighestLevelOfProtectionErr, "Error serializing 'isArmToHighestLevelOfProtection' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataArmSystem"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataArmSystem")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SecurityDataArmSystem) isSecurityDataArmSystem() bool {
	return true
}

func (m *_SecurityDataArmSystem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
