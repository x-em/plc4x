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

// S7PayloadWriteVarResponse is the corresponding interface of S7PayloadWriteVarResponse
type S7PayloadWriteVarResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7Payload
	// GetItems returns Items (property field)
	GetItems() []S7VarPayloadStatusItem
	// IsS7PayloadWriteVarResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadWriteVarResponse()
}

// _S7PayloadWriteVarResponse is the data-structure of this message
type _S7PayloadWriteVarResponse struct {
	S7PayloadContract
	Items []S7VarPayloadStatusItem
}

var _ S7PayloadWriteVarResponse = (*_S7PayloadWriteVarResponse)(nil)
var _ S7PayloadRequirements = (*_S7PayloadWriteVarResponse)(nil)

// NewS7PayloadWriteVarResponse factory function for _S7PayloadWriteVarResponse
func NewS7PayloadWriteVarResponse(items []S7VarPayloadStatusItem, parameter S7Parameter) *_S7PayloadWriteVarResponse {
	_result := &_S7PayloadWriteVarResponse{
		S7PayloadContract: NewS7Payload(parameter),
		Items:             items,
	}
	_result.S7PayloadContract.(*_S7Payload)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadWriteVarResponse) GetParameterParameterType() uint8 {
	return 0x05
}

func (m *_S7PayloadWriteVarResponse) GetMessageType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadWriteVarResponse) GetParent() S7PayloadContract {
	return m.S7PayloadContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadWriteVarResponse) GetItems() []S7VarPayloadStatusItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7PayloadWriteVarResponse(structType any) S7PayloadWriteVarResponse {
	if casted, ok := structType.(S7PayloadWriteVarResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadWriteVarResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadWriteVarResponse) GetTypeName() string {
	return "S7PayloadWriteVarResponse"
}

func (m *_S7PayloadWriteVarResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadContract.(*_S7Payload).getLengthInBits(ctx))

	// Array field
	if len(m.Items) > 0 {
		for _curItem, element := range m.Items {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Items), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_S7PayloadWriteVarResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadWriteVarResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7Payload, messageType uint8, parameter S7Parameter) (__s7PayloadWriteVarResponse S7PayloadWriteVarResponse, err error) {
	m.S7PayloadContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadWriteVarResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadWriteVarResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	items, err := ReadCountArrayField[S7VarPayloadStatusItem](ctx, "items", ReadComplex[S7VarPayloadStatusItem](S7VarPayloadStatusItemParseWithBuffer, readBuffer), uint64(CastS7ParameterWriteVarResponse(parameter).GetNumItems()))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}
	m.Items = items

	if closeErr := readBuffer.CloseContext("S7PayloadWriteVarResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadWriteVarResponse")
	}

	return m, nil
}

func (m *_S7PayloadWriteVarResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadWriteVarResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadWriteVarResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadWriteVarResponse")
		}

		if err := WriteComplexTypeArrayField(ctx, "items", m.GetItems(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'items' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadWriteVarResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadWriteVarResponse")
		}
		return nil
	}
	return m.S7PayloadContract.(*_S7Payload).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadWriteVarResponse) IsS7PayloadWriteVarResponse() {}

func (m *_S7PayloadWriteVarResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
