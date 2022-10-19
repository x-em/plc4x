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

// AdsInvalidResponse is the corresponding interface of AdsInvalidResponse
type AdsInvalidResponse interface {
	utils.LengthAware
	utils.Serializable
	AmsPacket
}

// AdsInvalidResponseExactly can be used when we want exactly this type and not a type which fulfills AdsInvalidResponse.
// This is useful for switch cases.
type AdsInvalidResponseExactly interface {
	AdsInvalidResponse
	isAdsInvalidResponse() bool
}

// _AdsInvalidResponse is the data-structure of this message
type _AdsInvalidResponse struct {
	*_AmsPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsInvalidResponse) GetCommandId() CommandId {
	return CommandId_INVALID
}

func (m *_AdsInvalidResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsInvalidResponse) InitializeParent(parent AmsPacket, targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) {
	m.TargetAmsNetId = targetAmsNetId
	m.TargetAmsPort = targetAmsPort
	m.SourceAmsNetId = sourceAmsNetId
	m.SourceAmsPort = sourceAmsPort
	m.ErrorCode = errorCode
	m.InvokeId = invokeId
}

func (m *_AdsInvalidResponse) GetParent() AmsPacket {
	return m._AmsPacket
}

// NewAdsInvalidResponse factory function for _AdsInvalidResponse
func NewAdsInvalidResponse(targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) *_AdsInvalidResponse {
	_result := &_AdsInvalidResponse{
		_AmsPacket: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
	}
	_result._AmsPacket._AmsPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsInvalidResponse(structType interface{}) AdsInvalidResponse {
	if casted, ok := structType.(AdsInvalidResponse); ok {
		return casted
	}
	if casted, ok := structType.(*AdsInvalidResponse); ok {
		return *casted
	}
	return nil
}

func (m *_AdsInvalidResponse) GetTypeName() string {
	return "AdsInvalidResponse"
}

func (m *_AdsInvalidResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AdsInvalidResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_AdsInvalidResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsInvalidResponseParse(readBuffer utils.ReadBuffer) (AdsInvalidResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsInvalidResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsInvalidResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AdsInvalidResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsInvalidResponse")
	}

	// Create a partially initialized instance
	_child := &_AdsInvalidResponse{
		_AmsPacket: &_AmsPacket{},
	}
	_child._AmsPacket._AmsPacketChildRequirements = _child
	return _child, nil
}

func (m *_AdsInvalidResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsInvalidResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsInvalidResponse")
		}

		if popErr := writeBuffer.PopContext("AdsInvalidResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsInvalidResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_AdsInvalidResponse) isAdsInvalidResponse() bool {
	return true
}

func (m *_AdsInvalidResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
