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

// ServiceCounterDataType is the corresponding interface of ServiceCounterDataType
type ServiceCounterDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetTotalCount returns TotalCount (property field)
	GetTotalCount() uint32
	// GetErrorCount returns ErrorCount (property field)
	GetErrorCount() uint32
	// IsServiceCounterDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsServiceCounterDataType()
}

// _ServiceCounterDataType is the data-structure of this message
type _ServiceCounterDataType struct {
	ExtensionObjectDefinitionContract
	TotalCount uint32
	ErrorCount uint32
}

var _ ServiceCounterDataType = (*_ServiceCounterDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ServiceCounterDataType)(nil)

// NewServiceCounterDataType factory function for _ServiceCounterDataType
func NewServiceCounterDataType(totalCount uint32, errorCount uint32) *_ServiceCounterDataType {
	_result := &_ServiceCounterDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		TotalCount:                        totalCount,
		ErrorCount:                        errorCount,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ServiceCounterDataType) GetIdentifier() string {
	return "873"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ServiceCounterDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ServiceCounterDataType) GetTotalCount() uint32 {
	return m.TotalCount
}

func (m *_ServiceCounterDataType) GetErrorCount() uint32 {
	return m.ErrorCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastServiceCounterDataType(structType any) ServiceCounterDataType {
	if casted, ok := structType.(ServiceCounterDataType); ok {
		return casted
	}
	if casted, ok := structType.(*ServiceCounterDataType); ok {
		return *casted
	}
	return nil
}

func (m *_ServiceCounterDataType) GetTypeName() string {
	return "ServiceCounterDataType"
}

func (m *_ServiceCounterDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (totalCount)
	lengthInBits += 32

	// Simple field (errorCount)
	lengthInBits += 32

	return lengthInBits
}

func (m *_ServiceCounterDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ServiceCounterDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__serviceCounterDataType ServiceCounterDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ServiceCounterDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ServiceCounterDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	totalCount, err := ReadSimpleField(ctx, "totalCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'totalCount' field"))
	}
	m.TotalCount = totalCount

	errorCount, err := ReadSimpleField(ctx, "errorCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorCount' field"))
	}
	m.ErrorCount = errorCount

	if closeErr := readBuffer.CloseContext("ServiceCounterDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ServiceCounterDataType")
	}

	return m, nil
}

func (m *_ServiceCounterDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ServiceCounterDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ServiceCounterDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ServiceCounterDataType")
		}

		if err := WriteSimpleField[uint32](ctx, "totalCount", m.GetTotalCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'totalCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "errorCount", m.GetErrorCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'errorCount' field")
		}

		if popErr := writeBuffer.PopContext("ServiceCounterDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ServiceCounterDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ServiceCounterDataType) IsServiceCounterDataType() {}

func (m *_ServiceCounterDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
