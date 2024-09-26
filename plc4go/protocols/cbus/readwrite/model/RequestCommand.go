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
const RequestCommand_INITIATOR byte = 0x5C

// RequestCommand is the corresponding interface of RequestCommand
type RequestCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	Request
	// GetCbusCommand returns CbusCommand (property field)
	GetCbusCommand() CBusCommand
	// GetChksum returns Chksum (property field)
	GetChksum() Checksum
	// GetAlpha returns Alpha (property field)
	GetAlpha() Alpha
	// GetCbusCommandDecoded returns CbusCommandDecoded (virtual field)
	GetCbusCommandDecoded() CBusCommand
	// GetChksumDecoded returns ChksumDecoded (virtual field)
	GetChksumDecoded() Checksum
	// IsRequestCommand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRequestCommand()
}

// _RequestCommand is the data-structure of this message
type _RequestCommand struct {
	RequestContract
	CbusCommand CBusCommand
	Chksum      Checksum
	Alpha       Alpha
}

var _ RequestCommand = (*_RequestCommand)(nil)
var _ RequestRequirements = (*_RequestCommand)(nil)

// NewRequestCommand factory function for _RequestCommand
func NewRequestCommand(peekedByte RequestType, startingCR *RequestType, resetMode *RequestType, secondPeek RequestType, termination RequestTermination, cbusCommand CBusCommand, chksum Checksum, alpha Alpha, cBusOptions CBusOptions) *_RequestCommand {
	_result := &_RequestCommand{
		RequestContract: NewRequest(peekedByte, startingCR, resetMode, secondPeek, termination, cBusOptions),
		CbusCommand:     cbusCommand,
		Chksum:          chksum,
		Alpha:           alpha,
	}
	_result.RequestContract.(*_Request)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RequestCommand) GetParent() RequestContract {
	return m.RequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RequestCommand) GetCbusCommand() CBusCommand {
	return m.CbusCommand
}

func (m *_RequestCommand) GetChksum() Checksum {
	return m.Chksum
}

func (m *_RequestCommand) GetAlpha() Alpha {
	return m.Alpha
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_RequestCommand) GetCbusCommandDecoded() CBusCommand {
	ctx := context.Background()
	_ = ctx
	alpha := m.GetAlpha()
	_ = alpha
	return CastCBusCommand(m.GetCbusCommand())
}

func (m *_RequestCommand) GetChksumDecoded() Checksum {
	ctx := context.Background()
	_ = ctx
	alpha := m.GetAlpha()
	_ = alpha
	return CastChecksum(m.GetChksum())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_RequestCommand) GetInitiator() byte {
	return RequestCommand_INITIATOR
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastRequestCommand(structType any) RequestCommand {
	if casted, ok := structType.(RequestCommand); ok {
		return casted
	}
	if casted, ok := structType.(*RequestCommand); ok {
		return *casted
	}
	return nil
}

func (m *_RequestCommand) GetTypeName() string {
	return "RequestCommand"
}

func (m *_RequestCommand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.RequestContract.(*_Request).getLengthInBits(ctx))

	// Const Field (initiator)
	lengthInBits += 8

	// Manual Field (cbusCommand)
	lengthInBits += uint16(int32((int32(m.GetCbusCommand().GetLengthInBytes(ctx)) * int32(int32(2)))) * int32(int32(8)))

	// A virtual field doesn't have any in- or output.

	// Manual Field (chksum)
	lengthInBits += uint16(utils.InlineIf((m.GetCBusOptions().GetSrchk()), func() any { return int32((int32(16))) }, func() any { return int32((int32(0))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Optional Field (alpha)
	if m.Alpha != nil {
		lengthInBits += m.Alpha.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_RequestCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RequestCommand) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Request, cBusOptions CBusOptions) (__requestCommand RequestCommand, err error) {
	m.RequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RequestCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RequestCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	initiator, err := ReadConstField[byte](ctx, "initiator", ReadByte(readBuffer, 8), RequestCommand_INITIATOR)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'initiator' field"))
	}
	_ = initiator

	cbusCommand, err := ReadManualField[CBusCommand](ctx, "cbusCommand", readBuffer, EnsureType[CBusCommand](ReadCBusCommand(ctx, readBuffer, cBusOptions, cBusOptions.GetSrchk())))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cbusCommand' field"))
	}
	m.CbusCommand = cbusCommand

	cbusCommandDecoded, err := ReadVirtualField[CBusCommand](ctx, "cbusCommandDecoded", (*CBusCommand)(nil), cbusCommand)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cbusCommandDecoded' field"))
	}
	_ = cbusCommandDecoded

	chksum, err := ReadManualField[Checksum](ctx, "chksum", readBuffer, EnsureType[Checksum](ReadAndValidateChecksum(ctx, readBuffer, cbusCommand, cBusOptions.GetSrchk())))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'chksum' field"))
	}
	m.Chksum = chksum

	chksumDecoded, err := ReadVirtualField[Checksum](ctx, "chksumDecoded", (*Checksum)(nil), chksum)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'chksumDecoded' field"))
	}
	_ = chksumDecoded

	var alpha Alpha
	_alpha, err := ReadOptionalField[Alpha](ctx, "alpha", ReadComplex[Alpha](AlphaParseWithBuffer, readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alpha' field"))
	}
	if _alpha != nil {
		alpha = *_alpha
		m.Alpha = alpha
	}

	if closeErr := readBuffer.CloseContext("RequestCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RequestCommand")
	}

	return m, nil
}

func (m *_RequestCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RequestCommand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RequestCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RequestCommand")
		}

		if err := WriteConstField(ctx, "initiator", RequestCommand_INITIATOR, WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'initiator' field")
		}

		if err := WriteManualField[CBusCommand](ctx, "cbusCommand", func(ctx context.Context) error { return WriteCBusCommand(ctx, writeBuffer, m.GetCbusCommand()) }, writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'cbusCommand' field")
		}
		// Virtual field
		cbusCommandDecoded := m.GetCbusCommandDecoded()
		_ = cbusCommandDecoded
		if _cbusCommandDecodedErr := writeBuffer.WriteVirtual(ctx, "cbusCommandDecoded", m.GetCbusCommandDecoded()); _cbusCommandDecodedErr != nil {
			return errors.Wrap(_cbusCommandDecodedErr, "Error serializing 'cbusCommandDecoded' field")
		}

		if err := WriteManualField[Checksum](ctx, "chksum", func(ctx context.Context) error {
			return CalculateChecksum(ctx, writeBuffer, m.GetCbusCommand(), m.GetCBusOptions().GetSrchk())
		}, writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'chksum' field")
		}
		// Virtual field
		chksumDecoded := m.GetChksumDecoded()
		_ = chksumDecoded
		if _chksumDecodedErr := writeBuffer.WriteVirtual(ctx, "chksumDecoded", m.GetChksumDecoded()); _chksumDecodedErr != nil {
			return errors.Wrap(_chksumDecodedErr, "Error serializing 'chksumDecoded' field")
		}

		if err := WriteOptionalField[Alpha](ctx, "alpha", GetRef(m.GetAlpha()), WriteComplex[Alpha](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'alpha' field")
		}

		if popErr := writeBuffer.PopContext("RequestCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RequestCommand")
		}
		return nil
	}
	return m.RequestContract.(*_Request).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RequestCommand) IsRequestCommand() {}

func (m *_RequestCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
