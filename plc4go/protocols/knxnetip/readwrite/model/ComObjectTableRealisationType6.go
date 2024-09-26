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

// ComObjectTableRealisationType6 is the corresponding interface of ComObjectTableRealisationType6
type ComObjectTableRealisationType6 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ComObjectTable
	// GetComObjectDescriptors returns ComObjectDescriptors (property field)
	GetComObjectDescriptors() GroupObjectDescriptorRealisationType6
	// IsComObjectTableRealisationType6 is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsComObjectTableRealisationType6()
}

// _ComObjectTableRealisationType6 is the data-structure of this message
type _ComObjectTableRealisationType6 struct {
	ComObjectTableContract
	ComObjectDescriptors GroupObjectDescriptorRealisationType6
}

var _ ComObjectTableRealisationType6 = (*_ComObjectTableRealisationType6)(nil)
var _ ComObjectTableRequirements = (*_ComObjectTableRealisationType6)(nil)

// NewComObjectTableRealisationType6 factory function for _ComObjectTableRealisationType6
func NewComObjectTableRealisationType6(comObjectDescriptors GroupObjectDescriptorRealisationType6) *_ComObjectTableRealisationType6 {
	if comObjectDescriptors == nil {
		panic("comObjectDescriptors of type GroupObjectDescriptorRealisationType6 for ComObjectTableRealisationType6 must not be nil")
	}
	_result := &_ComObjectTableRealisationType6{
		ComObjectTableContract: NewComObjectTable(),
		ComObjectDescriptors:   comObjectDescriptors,
	}
	_result.ComObjectTableContract.(*_ComObjectTable)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ComObjectTableRealisationType6) GetFirmwareType() FirmwareType {
	return FirmwareType_SYSTEM_300
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ComObjectTableRealisationType6) GetParent() ComObjectTableContract {
	return m.ComObjectTableContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ComObjectTableRealisationType6) GetComObjectDescriptors() GroupObjectDescriptorRealisationType6 {
	return m.ComObjectDescriptors
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastComObjectTableRealisationType6(structType any) ComObjectTableRealisationType6 {
	if casted, ok := structType.(ComObjectTableRealisationType6); ok {
		return casted
	}
	if casted, ok := structType.(*ComObjectTableRealisationType6); ok {
		return *casted
	}
	return nil
}

func (m *_ComObjectTableRealisationType6) GetTypeName() string {
	return "ComObjectTableRealisationType6"
}

func (m *_ComObjectTableRealisationType6) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ComObjectTableContract.(*_ComObjectTable).getLengthInBits(ctx))

	// Simple field (comObjectDescriptors)
	lengthInBits += m.ComObjectDescriptors.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ComObjectTableRealisationType6) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ComObjectTableRealisationType6) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ComObjectTable, firmwareType FirmwareType) (__comObjectTableRealisationType6 ComObjectTableRealisationType6, err error) {
	m.ComObjectTableContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ComObjectTableRealisationType6"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ComObjectTableRealisationType6")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	comObjectDescriptors, err := ReadSimpleField[GroupObjectDescriptorRealisationType6](ctx, "comObjectDescriptors", ReadComplex[GroupObjectDescriptorRealisationType6](GroupObjectDescriptorRealisationType6ParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'comObjectDescriptors' field"))
	}
	m.ComObjectDescriptors = comObjectDescriptors

	if closeErr := readBuffer.CloseContext("ComObjectTableRealisationType6"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ComObjectTableRealisationType6")
	}

	return m, nil
}

func (m *_ComObjectTableRealisationType6) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ComObjectTableRealisationType6) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ComObjectTableRealisationType6"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ComObjectTableRealisationType6")
		}

		if err := WriteSimpleField[GroupObjectDescriptorRealisationType6](ctx, "comObjectDescriptors", m.GetComObjectDescriptors(), WriteComplex[GroupObjectDescriptorRealisationType6](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'comObjectDescriptors' field")
		}

		if popErr := writeBuffer.PopContext("ComObjectTableRealisationType6"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ComObjectTableRealisationType6")
		}
		return nil
	}
	return m.ComObjectTableContract.(*_ComObjectTable).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ComObjectTableRealisationType6) IsComObjectTableRealisationType6() {}

func (m *_ComObjectTableRealisationType6) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
