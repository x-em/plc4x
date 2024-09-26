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

// DF1Command is the corresponding interface of DF1Command
type DF1Command interface {
	DF1CommandContract
	DF1CommandRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsDF1Command is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDF1Command()
}

// DF1CommandContract provides a set of functions which can be overwritten by a sub struct
type DF1CommandContract interface {
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetTransactionCounter returns TransactionCounter (property field)
	GetTransactionCounter() uint16
	// IsDF1Command is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDF1Command()
}

// DF1CommandRequirements provides a set of functions which need to be implemented by a sub struct
type DF1CommandRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetCommandCode returns CommandCode (discriminator field)
	GetCommandCode() uint8
}

// _DF1Command is the data-structure of this message
type _DF1Command struct {
	_SubType           DF1Command
	Status             uint8
	TransactionCounter uint16
}

var _ DF1CommandContract = (*_DF1Command)(nil)

// NewDF1Command factory function for _DF1Command
func NewDF1Command(status uint8, transactionCounter uint16) *_DF1Command {
	return &_DF1Command{Status: status, TransactionCounter: transactionCounter}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DF1Command) GetStatus() uint8 {
	return m.Status
}

func (m *_DF1Command) GetTransactionCounter() uint16 {
	return m.TransactionCounter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDF1Command(structType any) DF1Command {
	if casted, ok := structType.(DF1Command); ok {
		return casted
	}
	if casted, ok := structType.(*DF1Command); ok {
		return *casted
	}
	return nil
}

func (m *_DF1Command) GetTypeName() string {
	return "DF1Command"
}

func (m *_DF1Command) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (commandCode)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (transactionCounter)
	lengthInBits += 16

	return lengthInBits
}

func (m *_DF1Command) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func DF1CommandParse[T DF1Command](ctx context.Context, theBytes []byte) (T, error) {
	return DF1CommandParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func DF1CommandParseWithBufferProducer[T DF1Command]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := DF1CommandParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func DF1CommandParseWithBuffer[T DF1Command](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_DF1Command{}).parse(ctx, readBuffer)
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

func (m *_DF1Command) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__dF1Command DF1Command, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1Command"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1Command")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	commandCode, err := ReadDiscriminatorField[uint8](ctx, "commandCode", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandCode' field"))
	}

	status, err := ReadSimpleField(ctx, "status", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	transactionCounter, err := ReadSimpleField(ctx, "transactionCounter", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transactionCounter' field"))
	}
	m.TransactionCounter = transactionCounter

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child DF1Command
	switch {
	case commandCode == 0x01: // DF1UnprotectedReadRequest
		if _child, err = new(_DF1UnprotectedReadRequest).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type DF1UnprotectedReadRequest for type-switch of DF1Command")
		}
	case commandCode == 0x41: // DF1UnprotectedReadResponse
		if _child, err = new(_DF1UnprotectedReadResponse).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type DF1UnprotectedReadResponse for type-switch of DF1Command")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [commandCode=%v]", commandCode)
	}

	if closeErr := readBuffer.CloseContext("DF1Command"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1Command")
	}

	return _child, nil
}

func (pm *_DF1Command) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child DF1Command, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("DF1Command"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DF1Command")
	}

	if err := WriteDiscriminatorField(ctx, "commandCode", m.GetCommandCode(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'commandCode' field")
	}

	if err := WriteSimpleField[uint8](ctx, "status", m.GetStatus(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'status' field")
	}

	if err := WriteSimpleField[uint16](ctx, "transactionCounter", m.GetTransactionCounter(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'transactionCounter' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("DF1Command"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DF1Command")
	}
	return nil
}

func (m *_DF1Command) IsDF1Command() {}
