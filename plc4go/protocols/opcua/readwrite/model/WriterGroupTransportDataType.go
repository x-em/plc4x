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

// WriterGroupTransportDataType is the corresponding interface of WriterGroupTransportDataType
type WriterGroupTransportDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// IsWriterGroupTransportDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsWriterGroupTransportDataType()
}

// _WriterGroupTransportDataType is the data-structure of this message
type _WriterGroupTransportDataType struct {
	ExtensionObjectDefinitionContract
}

var _ WriterGroupTransportDataType = (*_WriterGroupTransportDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_WriterGroupTransportDataType)(nil)

// NewWriterGroupTransportDataType factory function for _WriterGroupTransportDataType
func NewWriterGroupTransportDataType() *_WriterGroupTransportDataType {
	_result := &_WriterGroupTransportDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_WriterGroupTransportDataType) GetIdentifier() string {
	return "15613"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_WriterGroupTransportDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// Deprecated: use the interface for direct cast
func CastWriterGroupTransportDataType(structType any) WriterGroupTransportDataType {
	if casted, ok := structType.(WriterGroupTransportDataType); ok {
		return casted
	}
	if casted, ok := structType.(*WriterGroupTransportDataType); ok {
		return *casted
	}
	return nil
}

func (m *_WriterGroupTransportDataType) GetTypeName() string {
	return "WriterGroupTransportDataType"
}

func (m *_WriterGroupTransportDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_WriterGroupTransportDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_WriterGroupTransportDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__writerGroupTransportDataType WriterGroupTransportDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("WriterGroupTransportDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for WriterGroupTransportDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("WriterGroupTransportDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for WriterGroupTransportDataType")
	}

	return m, nil
}

func (m *_WriterGroupTransportDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_WriterGroupTransportDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("WriterGroupTransportDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for WriterGroupTransportDataType")
		}

		if popErr := writeBuffer.PopContext("WriterGroupTransportDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for WriterGroupTransportDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_WriterGroupTransportDataType) IsWriterGroupTransportDataType() {}

func (m *_WriterGroupTransportDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
