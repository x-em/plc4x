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

// SysexCommandSysexRealtime is the corresponding interface of SysexCommandSysexRealtime
type SysexCommandSysexRealtime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SysexCommand
	// IsSysexCommandSysexRealtime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSysexCommandSysexRealtime()
}

// _SysexCommandSysexRealtime is the data-structure of this message
type _SysexCommandSysexRealtime struct {
	SysexCommandContract
}

var _ SysexCommandSysexRealtime = (*_SysexCommandSysexRealtime)(nil)
var _ SysexCommandRequirements = (*_SysexCommandSysexRealtime)(nil)

// NewSysexCommandSysexRealtime factory function for _SysexCommandSysexRealtime
func NewSysexCommandSysexRealtime() *_SysexCommandSysexRealtime {
	_result := &_SysexCommandSysexRealtime{
		SysexCommandContract: NewSysexCommand(),
	}
	_result.SysexCommandContract.(*_SysexCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandSysexRealtime) GetCommandType() uint8 {
	return 0x7F
}

func (m *_SysexCommandSysexRealtime) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandSysexRealtime) GetParent() SysexCommandContract {
	return m.SysexCommandContract
}

// Deprecated: use the interface for direct cast
func CastSysexCommandSysexRealtime(structType any) SysexCommandSysexRealtime {
	if casted, ok := structType.(SysexCommandSysexRealtime); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandSysexRealtime); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandSysexRealtime) GetTypeName() string {
	return "SysexCommandSysexRealtime"
}

func (m *_SysexCommandSysexRealtime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SysexCommandContract.(*_SysexCommand).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SysexCommandSysexRealtime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SysexCommandSysexRealtime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SysexCommand, response bool) (__sysexCommandSysexRealtime SysexCommandSysexRealtime, err error) {
	m.SysexCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandSysexRealtime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandSysexRealtime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandSysexRealtime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandSysexRealtime")
	}

	return m, nil
}

func (m *_SysexCommandSysexRealtime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandSysexRealtime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandSysexRealtime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandSysexRealtime")
		}

		if popErr := writeBuffer.PopContext("SysexCommandSysexRealtime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandSysexRealtime")
		}
		return nil
	}
	return m.SysexCommandContract.(*_SysexCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandSysexRealtime) IsSysexCommandSysexRealtime() {}

func (m *_SysexCommandSysexRealtime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
