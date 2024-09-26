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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ApduDataExtGroupPropertyValueWrite is the corresponding interface of ApduDataExtGroupPropertyValueWrite
type ApduDataExtGroupPropertyValueWrite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduDataExt
	// IsApduDataExtGroupPropertyValueWrite is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtGroupPropertyValueWrite()
}

// _ApduDataExtGroupPropertyValueWrite is the data-structure of this message
type _ApduDataExtGroupPropertyValueWrite struct {
	ApduDataExtContract
}

var _ ApduDataExtGroupPropertyValueWrite = (*_ApduDataExtGroupPropertyValueWrite)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtGroupPropertyValueWrite)(nil)

// NewApduDataExtGroupPropertyValueWrite factory function for _ApduDataExtGroupPropertyValueWrite
func NewApduDataExtGroupPropertyValueWrite(length uint8) *_ApduDataExtGroupPropertyValueWrite {
	_result := &_ApduDataExtGroupPropertyValueWrite{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtGroupPropertyValueWrite) GetExtApciType() uint8 {
	return 0x2A
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtGroupPropertyValueWrite) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// Deprecated: use the interface for direct cast
func CastApduDataExtGroupPropertyValueWrite(structType any) ApduDataExtGroupPropertyValueWrite {
	if casted, ok := structType.(ApduDataExtGroupPropertyValueWrite); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtGroupPropertyValueWrite); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtGroupPropertyValueWrite) GetTypeName() string {
	return "ApduDataExtGroupPropertyValueWrite"
}

func (m *_ApduDataExtGroupPropertyValueWrite) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtGroupPropertyValueWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtGroupPropertyValueWrite) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtGroupPropertyValueWrite ApduDataExtGroupPropertyValueWrite, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtGroupPropertyValueWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtGroupPropertyValueWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtGroupPropertyValueWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtGroupPropertyValueWrite")
	}

	return m, nil
}

func (m *_ApduDataExtGroupPropertyValueWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtGroupPropertyValueWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtGroupPropertyValueWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtGroupPropertyValueWrite")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtGroupPropertyValueWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtGroupPropertyValueWrite")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtGroupPropertyValueWrite) IsApduDataExtGroupPropertyValueWrite() {}

func (m *_ApduDataExtGroupPropertyValueWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
