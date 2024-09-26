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

// NLMVendorProprietaryMessage is the corresponding interface of NLMVendorProprietaryMessage
type NLMVendorProprietaryMessage interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NLM
	// GetVendorId returns VendorId (property field)
	GetVendorId() BACnetVendorId
	// GetProprietaryMessage returns ProprietaryMessage (property field)
	GetProprietaryMessage() []byte
	// IsNLMVendorProprietaryMessage is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNLMVendorProprietaryMessage()
}

// _NLMVendorProprietaryMessage is the data-structure of this message
type _NLMVendorProprietaryMessage struct {
	NLMContract
	VendorId           BACnetVendorId
	ProprietaryMessage []byte
}

var _ NLMVendorProprietaryMessage = (*_NLMVendorProprietaryMessage)(nil)
var _ NLMRequirements = (*_NLMVendorProprietaryMessage)(nil)

// NewNLMVendorProprietaryMessage factory function for _NLMVendorProprietaryMessage
func NewNLMVendorProprietaryMessage(vendorId BACnetVendorId, proprietaryMessage []byte, apduLength uint16) *_NLMVendorProprietaryMessage {
	_result := &_NLMVendorProprietaryMessage{
		NLMContract:        NewNLM(apduLength),
		VendorId:           vendorId,
		ProprietaryMessage: proprietaryMessage,
	}
	_result.NLMContract.(*_NLM)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMVendorProprietaryMessage) GetMessageType() uint8 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMVendorProprietaryMessage) GetParent() NLMContract {
	return m.NLMContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMVendorProprietaryMessage) GetVendorId() BACnetVendorId {
	return m.VendorId
}

func (m *_NLMVendorProprietaryMessage) GetProprietaryMessage() []byte {
	return m.ProprietaryMessage
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNLMVendorProprietaryMessage(structType any) NLMVendorProprietaryMessage {
	if casted, ok := structType.(NLMVendorProprietaryMessage); ok {
		return casted
	}
	if casted, ok := structType.(*NLMVendorProprietaryMessage); ok {
		return *casted
	}
	return nil
}

func (m *_NLMVendorProprietaryMessage) GetTypeName() string {
	return "NLMVendorProprietaryMessage"
}

func (m *_NLMVendorProprietaryMessage) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NLMContract.(*_NLM).getLengthInBits(ctx))

	// Simple field (vendorId)
	lengthInBits += 16

	// Array field
	if len(m.ProprietaryMessage) > 0 {
		lengthInBits += 8 * uint16(len(m.ProprietaryMessage))
	}

	return lengthInBits
}

func (m *_NLMVendorProprietaryMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NLMVendorProprietaryMessage) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NLM, apduLength uint16) (__nLMVendorProprietaryMessage NLMVendorProprietaryMessage, err error) {
	m.NLMContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMVendorProprietaryMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMVendorProprietaryMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	vendorId, err := ReadEnumField[BACnetVendorId](ctx, "vendorId", "BACnetVendorId", ReadEnum(BACnetVendorIdByValue, ReadUnsignedShort(readBuffer, uint8(16))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vendorId' field"))
	}
	m.VendorId = vendorId

	proprietaryMessage, err := readBuffer.ReadByteArray("proprietaryMessage", int(utils.InlineIf((bool((apduLength) > (0))), func() any { return int32((int32(apduLength) - int32(int32(3)))) }, func() any { return int32(int32(0)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryMessage' field"))
	}
	m.ProprietaryMessage = proprietaryMessage

	if closeErr := readBuffer.CloseContext("NLMVendorProprietaryMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMVendorProprietaryMessage")
	}

	return m, nil
}

func (m *_NLMVendorProprietaryMessage) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMVendorProprietaryMessage) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMVendorProprietaryMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMVendorProprietaryMessage")
		}

		if err := WriteSimpleEnumField[BACnetVendorId](ctx, "vendorId", "BACnetVendorId", m.GetVendorId(), WriteEnum[BACnetVendorId, uint16](BACnetVendorId.GetValue, BACnetVendorId.PLC4XEnumName, WriteUnsignedShort(writeBuffer, 16))); err != nil {
			return errors.Wrap(err, "Error serializing 'vendorId' field")
		}

		if err := WriteByteArrayField(ctx, "proprietaryMessage", m.GetProprietaryMessage(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'proprietaryMessage' field")
		}

		if popErr := writeBuffer.PopContext("NLMVendorProprietaryMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMVendorProprietaryMessage")
		}
		return nil
	}
	return m.NLMContract.(*_NLM).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMVendorProprietaryMessage) IsNLMVendorProprietaryMessage() {}

func (m *_NLMVendorProprietaryMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
