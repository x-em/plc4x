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

// CALReplyShort is the corresponding interface of CALReplyShort
type CALReplyShort interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CALReply
	// IsCALReplyShort is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCALReplyShort()
}

// _CALReplyShort is the data-structure of this message
type _CALReplyShort struct {
	CALReplyContract
}

var _ CALReplyShort = (*_CALReplyShort)(nil)
var _ CALReplyRequirements = (*_CALReplyShort)(nil)

// NewCALReplyShort factory function for _CALReplyShort
func NewCALReplyShort(calType byte, calData CALData, cBusOptions CBusOptions, requestContext RequestContext) *_CALReplyShort {
	_result := &_CALReplyShort{
		CALReplyContract: NewCALReply(calType, calData, cBusOptions, requestContext),
	}
	_result.CALReplyContract.(*_CALReply)._SubType = _result
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

func (m *_CALReplyShort) GetParent() CALReplyContract {
	return m.CALReplyContract
}

// Deprecated: use the interface for direct cast
func CastCALReplyShort(structType any) CALReplyShort {
	if casted, ok := structType.(CALReplyShort); ok {
		return casted
	}
	if casted, ok := structType.(*CALReplyShort); ok {
		return *casted
	}
	return nil
}

func (m *_CALReplyShort) GetTypeName() string {
	return "CALReplyShort"
}

func (m *_CALReplyShort) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CALReplyContract.(*_CALReply).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_CALReplyShort) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CALReplyShort) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CALReply, cBusOptions CBusOptions, requestContext RequestContext) (__cALReplyShort CALReplyShort, err error) {
	m.CALReplyContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALReplyShort"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALReplyShort")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("CALReplyShort"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALReplyShort")
	}

	return m, nil
}

func (m *_CALReplyShort) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALReplyShort) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALReplyShort"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALReplyShort")
		}

		if popErr := writeBuffer.PopContext("CALReplyShort"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALReplyShort")
		}
		return nil
	}
	return m.CALReplyContract.(*_CALReply).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CALReplyShort) IsCALReplyShort() {}

func (m *_CALReplyShort) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
