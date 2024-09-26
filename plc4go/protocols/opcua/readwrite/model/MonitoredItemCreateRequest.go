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

// MonitoredItemCreateRequest is the corresponding interface of MonitoredItemCreateRequest
type MonitoredItemCreateRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetItemToMonitor returns ItemToMonitor (property field)
	GetItemToMonitor() ExtensionObjectDefinition
	// GetMonitoringMode returns MonitoringMode (property field)
	GetMonitoringMode() MonitoringMode
	// GetRequestedParameters returns RequestedParameters (property field)
	GetRequestedParameters() ExtensionObjectDefinition
	// IsMonitoredItemCreateRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMonitoredItemCreateRequest()
}

// _MonitoredItemCreateRequest is the data-structure of this message
type _MonitoredItemCreateRequest struct {
	ExtensionObjectDefinitionContract
	ItemToMonitor       ExtensionObjectDefinition
	MonitoringMode      MonitoringMode
	RequestedParameters ExtensionObjectDefinition
}

var _ MonitoredItemCreateRequest = (*_MonitoredItemCreateRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_MonitoredItemCreateRequest)(nil)

// NewMonitoredItemCreateRequest factory function for _MonitoredItemCreateRequest
func NewMonitoredItemCreateRequest(itemToMonitor ExtensionObjectDefinition, monitoringMode MonitoringMode, requestedParameters ExtensionObjectDefinition) *_MonitoredItemCreateRequest {
	if itemToMonitor == nil {
		panic("itemToMonitor of type ExtensionObjectDefinition for MonitoredItemCreateRequest must not be nil")
	}
	if requestedParameters == nil {
		panic("requestedParameters of type ExtensionObjectDefinition for MonitoredItemCreateRequest must not be nil")
	}
	_result := &_MonitoredItemCreateRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ItemToMonitor:                     itemToMonitor,
		MonitoringMode:                    monitoringMode,
		RequestedParameters:               requestedParameters,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MonitoredItemCreateRequest) GetIdentifier() string {
	return "745"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MonitoredItemCreateRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoredItemCreateRequest) GetItemToMonitor() ExtensionObjectDefinition {
	return m.ItemToMonitor
}

func (m *_MonitoredItemCreateRequest) GetMonitoringMode() MonitoringMode {
	return m.MonitoringMode
}

func (m *_MonitoredItemCreateRequest) GetRequestedParameters() ExtensionObjectDefinition {
	return m.RequestedParameters
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMonitoredItemCreateRequest(structType any) MonitoredItemCreateRequest {
	if casted, ok := structType.(MonitoredItemCreateRequest); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoredItemCreateRequest); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoredItemCreateRequest) GetTypeName() string {
	return "MonitoredItemCreateRequest"
}

func (m *_MonitoredItemCreateRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (itemToMonitor)
	lengthInBits += m.ItemToMonitor.GetLengthInBits(ctx)

	// Simple field (monitoringMode)
	lengthInBits += 32

	// Simple field (requestedParameters)
	lengthInBits += m.RequestedParameters.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_MonitoredItemCreateRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MonitoredItemCreateRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__monitoredItemCreateRequest MonitoredItemCreateRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MonitoredItemCreateRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoredItemCreateRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	itemToMonitor, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "itemToMonitor", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("628")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemToMonitor' field"))
	}
	m.ItemToMonitor = itemToMonitor

	monitoringMode, err := ReadEnumField[MonitoringMode](ctx, "monitoringMode", "MonitoringMode", ReadEnum(MonitoringModeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoringMode' field"))
	}
	m.MonitoringMode = monitoringMode

	requestedParameters, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestedParameters", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("742")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedParameters' field"))
	}
	m.RequestedParameters = requestedParameters

	if closeErr := readBuffer.CloseContext("MonitoredItemCreateRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoredItemCreateRequest")
	}

	return m, nil
}

func (m *_MonitoredItemCreateRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoredItemCreateRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoredItemCreateRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoredItemCreateRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "itemToMonitor", m.GetItemToMonitor(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemToMonitor' field")
		}

		if err := WriteSimpleEnumField[MonitoringMode](ctx, "monitoringMode", "MonitoringMode", m.GetMonitoringMode(), WriteEnum[MonitoringMode, uint32](MonitoringMode.GetValue, MonitoringMode.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'monitoringMode' field")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestedParameters", m.GetRequestedParameters(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedParameters' field")
		}

		if popErr := writeBuffer.PopContext("MonitoredItemCreateRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoredItemCreateRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MonitoredItemCreateRequest) IsMonitoredItemCreateRequest() {}

func (m *_MonitoredItemCreateRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
