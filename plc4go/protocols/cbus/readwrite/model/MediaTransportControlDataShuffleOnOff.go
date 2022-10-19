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

// MediaTransportControlDataShuffleOnOff is the corresponding interface of MediaTransportControlDataShuffleOnOff
type MediaTransportControlDataShuffleOnOff interface {
	utils.LengthAware
	utils.Serializable
	MediaTransportControlData
	// GetState returns State (property field)
	GetState() byte
	// GetIsOff returns IsOff (virtual field)
	GetIsOff() bool
	// GetIsOn returns IsOn (virtual field)
	GetIsOn() bool
}

// MediaTransportControlDataShuffleOnOffExactly can be used when we want exactly this type and not a type which fulfills MediaTransportControlDataShuffleOnOff.
// This is useful for switch cases.
type MediaTransportControlDataShuffleOnOffExactly interface {
	MediaTransportControlDataShuffleOnOff
	isMediaTransportControlDataShuffleOnOff() bool
}

// _MediaTransportControlDataShuffleOnOff is the data-structure of this message
type _MediaTransportControlDataShuffleOnOff struct {
	*_MediaTransportControlData
	State byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataShuffleOnOff) InitializeParent(parent MediaTransportControlData, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.MediaLinkGroup = mediaLinkGroup
}

func (m *_MediaTransportControlDataShuffleOnOff) GetParent() MediaTransportControlData {
	return m._MediaTransportControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataShuffleOnOff) GetState() byte {
	return m.State
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MediaTransportControlDataShuffleOnOff) GetIsOff() bool {
	return bool(bool((m.GetState()) == (0x00)))
}

func (m *_MediaTransportControlDataShuffleOnOff) GetIsOn() bool {
	return bool(bool((m.GetState()) > (0xFE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMediaTransportControlDataShuffleOnOff factory function for _MediaTransportControlDataShuffleOnOff
func NewMediaTransportControlDataShuffleOnOff(state byte, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlDataShuffleOnOff {
	_result := &_MediaTransportControlDataShuffleOnOff{
		State:                      state,
		_MediaTransportControlData: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
	}
	_result._MediaTransportControlData._MediaTransportControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataShuffleOnOff(structType interface{}) MediaTransportControlDataShuffleOnOff {
	if casted, ok := structType.(MediaTransportControlDataShuffleOnOff); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataShuffleOnOff); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataShuffleOnOff) GetTypeName() string {
	return "MediaTransportControlDataShuffleOnOff"
}

func (m *_MediaTransportControlDataShuffleOnOff) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_MediaTransportControlDataShuffleOnOff) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (state)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_MediaTransportControlDataShuffleOnOff) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MediaTransportControlDataShuffleOnOffParse(readBuffer utils.ReadBuffer) (MediaTransportControlDataShuffleOnOff, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataShuffleOnOff"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataShuffleOnOff")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (state)
	_state, _stateErr := readBuffer.ReadByte("state")
	if _stateErr != nil {
		return nil, errors.Wrap(_stateErr, "Error parsing 'state' field of MediaTransportControlDataShuffleOnOff")
	}
	state := _state

	// Virtual field
	_isOff := bool((state) == (0x00))
	isOff := bool(_isOff)
	_ = isOff

	// Virtual field
	_isOn := bool((state) > (0xFE))
	isOn := bool(_isOn)
	_ = isOn

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataShuffleOnOff"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataShuffleOnOff")
	}

	// Create a partially initialized instance
	_child := &_MediaTransportControlDataShuffleOnOff{
		_MediaTransportControlData: &_MediaTransportControlData{},
		State:                      state,
	}
	_child._MediaTransportControlData._MediaTransportControlDataChildRequirements = _child
	return _child, nil
}

func (m *_MediaTransportControlDataShuffleOnOff) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataShuffleOnOff"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataShuffleOnOff")
		}

		// Simple Field (state)
		state := byte(m.GetState())
		_stateErr := writeBuffer.WriteByte("state", (state))
		if _stateErr != nil {
			return errors.Wrap(_stateErr, "Error serializing 'state' field")
		}
		// Virtual field
		if _isOffErr := writeBuffer.WriteVirtual("isOff", m.GetIsOff()); _isOffErr != nil {
			return errors.Wrap(_isOffErr, "Error serializing 'isOff' field")
		}
		// Virtual field
		if _isOnErr := writeBuffer.WriteVirtual("isOn", m.GetIsOn()); _isOnErr != nil {
			return errors.Wrap(_isOnErr, "Error serializing 'isOn' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataShuffleOnOff"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataShuffleOnOff")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataShuffleOnOff) isMediaTransportControlDataShuffleOnOff() bool {
	return true
}

func (m *_MediaTransportControlDataShuffleOnOff) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
