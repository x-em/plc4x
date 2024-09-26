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

// Constant values.
const S7DataAlarmMessage_FUNCTIONID uint8 = 0x00
const S7DataAlarmMessage_NUMBERMESSAGEOBJ uint8 = 0x01

// S7DataAlarmMessage is the corresponding interface of S7DataAlarmMessage
type S7DataAlarmMessage interface {
	S7DataAlarmMessageContract
	S7DataAlarmMessageRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsS7DataAlarmMessage is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7DataAlarmMessage()
}

// S7DataAlarmMessageContract provides a set of functions which can be overwritten by a sub struct
type S7DataAlarmMessageContract interface {
	// IsS7DataAlarmMessage is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7DataAlarmMessage()
}

// S7DataAlarmMessageRequirements provides a set of functions which need to be implemented by a sub struct
type S7DataAlarmMessageRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetCpuFunctionType returns CpuFunctionType (discriminator field)
	GetCpuFunctionType() uint8
}

// _S7DataAlarmMessage is the data-structure of this message
type _S7DataAlarmMessage struct {
	_SubType S7DataAlarmMessage
}

var _ S7DataAlarmMessageContract = (*_S7DataAlarmMessage)(nil)

// NewS7DataAlarmMessage factory function for _S7DataAlarmMessage
func NewS7DataAlarmMessage() *_S7DataAlarmMessage {
	return &_S7DataAlarmMessage{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_S7DataAlarmMessage) GetFunctionId() uint8 {
	return S7DataAlarmMessage_FUNCTIONID
}

func (m *_S7DataAlarmMessage) GetNumberMessageObj() uint8 {
	return S7DataAlarmMessage_NUMBERMESSAGEOBJ
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7DataAlarmMessage(structType any) S7DataAlarmMessage {
	if casted, ok := structType.(S7DataAlarmMessage); ok {
		return casted
	}
	if casted, ok := structType.(*S7DataAlarmMessage); ok {
		return *casted
	}
	return nil
}

func (m *_S7DataAlarmMessage) GetTypeName() string {
	return "S7DataAlarmMessage"
}

func (m *_S7DataAlarmMessage) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (functionId)
	lengthInBits += 8

	// Const Field (numberMessageObj)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7DataAlarmMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func S7DataAlarmMessageParse[T S7DataAlarmMessage](ctx context.Context, theBytes []byte, cpuFunctionType uint8) (T, error) {
	return S7DataAlarmMessageParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), cpuFunctionType)
}

func S7DataAlarmMessageParseWithBufferProducer[T S7DataAlarmMessage](cpuFunctionType uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := S7DataAlarmMessageParseWithBuffer[T](ctx, readBuffer, cpuFunctionType)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func S7DataAlarmMessageParseWithBuffer[T S7DataAlarmMessage](ctx context.Context, readBuffer utils.ReadBuffer, cpuFunctionType uint8) (T, error) {
	v, err := (&_S7DataAlarmMessage{}).parse(ctx, readBuffer, cpuFunctionType)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_S7DataAlarmMessage) parse(ctx context.Context, readBuffer utils.ReadBuffer, cpuFunctionType uint8) (__s7DataAlarmMessage S7DataAlarmMessage, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7DataAlarmMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7DataAlarmMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	functionId, err := ReadConstField[uint8](ctx, "functionId", ReadUnsignedByte(readBuffer, uint8(8)), S7DataAlarmMessage_FUNCTIONID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'functionId' field"))
	}
	_ = functionId

	numberMessageObj, err := ReadConstField[uint8](ctx, "numberMessageObj", ReadUnsignedByte(readBuffer, uint8(8)), S7DataAlarmMessage_NUMBERMESSAGEOBJ)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberMessageObj' field"))
	}
	_ = numberMessageObj

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child S7DataAlarmMessage
	switch {
	case cpuFunctionType == 0x04: // S7MessageObjectRequest
		if _child, err = new(_S7MessageObjectRequest).parse(ctx, readBuffer, m, cpuFunctionType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type S7MessageObjectRequest for type-switch of S7DataAlarmMessage")
		}
	case cpuFunctionType == 0x08: // S7MessageObjectResponse
		if _child, err = new(_S7MessageObjectResponse).parse(ctx, readBuffer, m, cpuFunctionType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type S7MessageObjectResponse for type-switch of S7DataAlarmMessage")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [cpuFunctionType=%v]", cpuFunctionType)
	}

	if closeErr := readBuffer.CloseContext("S7DataAlarmMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7DataAlarmMessage")
	}

	return _child, nil
}

func (pm *_S7DataAlarmMessage) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child S7DataAlarmMessage, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("S7DataAlarmMessage"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for S7DataAlarmMessage")
	}

	if err := WriteConstField(ctx, "functionId", S7DataAlarmMessage_FUNCTIONID, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'functionId' field")
	}

	if err := WriteConstField(ctx, "numberMessageObj", S7DataAlarmMessage_NUMBERMESSAGEOBJ, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'numberMessageObj' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7DataAlarmMessage"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for S7DataAlarmMessage")
	}
	return nil
}

func (m *_S7DataAlarmMessage) IsS7DataAlarmMessage() {}
