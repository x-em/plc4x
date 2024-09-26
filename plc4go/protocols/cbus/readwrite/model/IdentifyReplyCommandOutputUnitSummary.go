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

// IdentifyReplyCommandOutputUnitSummary is the corresponding interface of IdentifyReplyCommandOutputUnitSummary
type IdentifyReplyCommandOutputUnitSummary interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetUnitFlags returns UnitFlags (property field)
	GetUnitFlags() IdentifyReplyCommandUnitSummary
	// GetGavStoreEnabledByte1 returns GavStoreEnabledByte1 (property field)
	GetGavStoreEnabledByte1() *byte
	// GetGavStoreEnabledByte2 returns GavStoreEnabledByte2 (property field)
	GetGavStoreEnabledByte2() *byte
	// GetTimeFromLastRecoverOfMainsInSeconds returns TimeFromLastRecoverOfMainsInSeconds (property field)
	GetTimeFromLastRecoverOfMainsInSeconds() uint8
	// IsIdentifyReplyCommandOutputUnitSummary is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsIdentifyReplyCommandOutputUnitSummary()
}

// _IdentifyReplyCommandOutputUnitSummary is the data-structure of this message
type _IdentifyReplyCommandOutputUnitSummary struct {
	IdentifyReplyCommandContract
	UnitFlags                           IdentifyReplyCommandUnitSummary
	GavStoreEnabledByte1                *byte
	GavStoreEnabledByte2                *byte
	TimeFromLastRecoverOfMainsInSeconds uint8
}

var _ IdentifyReplyCommandOutputUnitSummary = (*_IdentifyReplyCommandOutputUnitSummary)(nil)
var _ IdentifyReplyCommandRequirements = (*_IdentifyReplyCommandOutputUnitSummary)(nil)

// NewIdentifyReplyCommandOutputUnitSummary factory function for _IdentifyReplyCommandOutputUnitSummary
func NewIdentifyReplyCommandOutputUnitSummary(unitFlags IdentifyReplyCommandUnitSummary, gavStoreEnabledByte1 *byte, gavStoreEnabledByte2 *byte, timeFromLastRecoverOfMainsInSeconds uint8, numBytes uint8) *_IdentifyReplyCommandOutputUnitSummary {
	if unitFlags == nil {
		panic("unitFlags of type IdentifyReplyCommandUnitSummary for IdentifyReplyCommandOutputUnitSummary must not be nil")
	}
	_result := &_IdentifyReplyCommandOutputUnitSummary{
		IdentifyReplyCommandContract:        NewIdentifyReplyCommand(numBytes),
		UnitFlags:                           unitFlags,
		GavStoreEnabledByte1:                gavStoreEnabledByte1,
		GavStoreEnabledByte2:                gavStoreEnabledByte2,
		TimeFromLastRecoverOfMainsInSeconds: timeFromLastRecoverOfMainsInSeconds,
	}
	_result.IdentifyReplyCommandContract.(*_IdentifyReplyCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandOutputUnitSummary) GetAttribute() Attribute {
	return Attribute_OutputUnitSummary
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandOutputUnitSummary) GetParent() IdentifyReplyCommandContract {
	return m.IdentifyReplyCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandOutputUnitSummary) GetUnitFlags() IdentifyReplyCommandUnitSummary {
	return m.UnitFlags
}

func (m *_IdentifyReplyCommandOutputUnitSummary) GetGavStoreEnabledByte1() *byte {
	return m.GavStoreEnabledByte1
}

func (m *_IdentifyReplyCommandOutputUnitSummary) GetGavStoreEnabledByte2() *byte {
	return m.GavStoreEnabledByte2
}

func (m *_IdentifyReplyCommandOutputUnitSummary) GetTimeFromLastRecoverOfMainsInSeconds() uint8 {
	return m.TimeFromLastRecoverOfMainsInSeconds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandOutputUnitSummary(structType any) IdentifyReplyCommandOutputUnitSummary {
	if casted, ok := structType.(IdentifyReplyCommandOutputUnitSummary); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandOutputUnitSummary); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandOutputUnitSummary) GetTypeName() string {
	return "IdentifyReplyCommandOutputUnitSummary"
}

func (m *_IdentifyReplyCommandOutputUnitSummary) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).getLengthInBits(ctx))

	// Simple field (unitFlags)
	lengthInBits += m.UnitFlags.GetLengthInBits(ctx)

	// Optional Field (gavStoreEnabledByte1)
	if m.GavStoreEnabledByte1 != nil {
		lengthInBits += 8
	}

	// Optional Field (gavStoreEnabledByte2)
	if m.GavStoreEnabledByte2 != nil {
		lengthInBits += 8
	}

	// Simple field (timeFromLastRecoverOfMainsInSeconds)
	lengthInBits += 8

	return lengthInBits
}

func (m *_IdentifyReplyCommandOutputUnitSummary) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_IdentifyReplyCommandOutputUnitSummary) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_IdentifyReplyCommand, attribute Attribute, numBytes uint8) (__identifyReplyCommandOutputUnitSummary IdentifyReplyCommandOutputUnitSummary, err error) {
	m.IdentifyReplyCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandOutputUnitSummary"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandOutputUnitSummary")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	unitFlags, err := ReadSimpleField[IdentifyReplyCommandUnitSummary](ctx, "unitFlags", ReadComplex[IdentifyReplyCommandUnitSummary](IdentifyReplyCommandUnitSummaryParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unitFlags' field"))
	}
	m.UnitFlags = unitFlags

	var gavStoreEnabledByte1 *byte
	gavStoreEnabledByte1, err = ReadOptionalField[byte](ctx, "gavStoreEnabledByte1", ReadByte(readBuffer, 8), bool((numBytes) > (1)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'gavStoreEnabledByte1' field"))
	}
	m.GavStoreEnabledByte1 = gavStoreEnabledByte1

	var gavStoreEnabledByte2 *byte
	gavStoreEnabledByte2, err = ReadOptionalField[byte](ctx, "gavStoreEnabledByte2", ReadByte(readBuffer, 8), bool((numBytes) > (2)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'gavStoreEnabledByte2' field"))
	}
	m.GavStoreEnabledByte2 = gavStoreEnabledByte2

	timeFromLastRecoverOfMainsInSeconds, err := ReadSimpleField(ctx, "timeFromLastRecoverOfMainsInSeconds", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeFromLastRecoverOfMainsInSeconds' field"))
	}
	m.TimeFromLastRecoverOfMainsInSeconds = timeFromLastRecoverOfMainsInSeconds

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandOutputUnitSummary"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandOutputUnitSummary")
	}

	return m, nil
}

func (m *_IdentifyReplyCommandOutputUnitSummary) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandOutputUnitSummary) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandOutputUnitSummary"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandOutputUnitSummary")
		}

		if err := WriteSimpleField[IdentifyReplyCommandUnitSummary](ctx, "unitFlags", m.GetUnitFlags(), WriteComplex[IdentifyReplyCommandUnitSummary](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'unitFlags' field")
		}

		if err := WriteOptionalField[byte](ctx, "gavStoreEnabledByte1", m.GetGavStoreEnabledByte1(), WriteByte(writeBuffer, 8), true); err != nil {
			return errors.Wrap(err, "Error serializing 'gavStoreEnabledByte1' field")
		}

		if err := WriteOptionalField[byte](ctx, "gavStoreEnabledByte2", m.GetGavStoreEnabledByte2(), WriteByte(writeBuffer, 8), true); err != nil {
			return errors.Wrap(err, "Error serializing 'gavStoreEnabledByte2' field")
		}

		if err := WriteSimpleField[uint8](ctx, "timeFromLastRecoverOfMainsInSeconds", m.GetTimeFromLastRecoverOfMainsInSeconds(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeFromLastRecoverOfMainsInSeconds' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandOutputUnitSummary"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandOutputUnitSummary")
		}
		return nil
	}
	return m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandOutputUnitSummary) IsIdentifyReplyCommandOutputUnitSummary() {}

func (m *_IdentifyReplyCommandOutputUnitSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
