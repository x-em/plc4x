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

// RedundantServerDataType is the corresponding interface of RedundantServerDataType
type RedundantServerDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetServerId returns ServerId (property field)
	GetServerId() PascalString
	// GetServiceLevel returns ServiceLevel (property field)
	GetServiceLevel() uint8
	// GetServerState returns ServerState (property field)
	GetServerState() ServerState
	// IsRedundantServerDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRedundantServerDataType()
}

// _RedundantServerDataType is the data-structure of this message
type _RedundantServerDataType struct {
	ExtensionObjectDefinitionContract
	ServerId     PascalString
	ServiceLevel uint8
	ServerState  ServerState
}

var _ RedundantServerDataType = (*_RedundantServerDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_RedundantServerDataType)(nil)

// NewRedundantServerDataType factory function for _RedundantServerDataType
func NewRedundantServerDataType(serverId PascalString, serviceLevel uint8, serverState ServerState) *_RedundantServerDataType {
	if serverId == nil {
		panic("serverId of type PascalString for RedundantServerDataType must not be nil")
	}
	_result := &_RedundantServerDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ServerId:                          serverId,
		ServiceLevel:                      serviceLevel,
		ServerState:                       serverState,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RedundantServerDataType) GetIdentifier() string {
	return "855"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RedundantServerDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RedundantServerDataType) GetServerId() PascalString {
	return m.ServerId
}

func (m *_RedundantServerDataType) GetServiceLevel() uint8 {
	return m.ServiceLevel
}

func (m *_RedundantServerDataType) GetServerState() ServerState {
	return m.ServerState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastRedundantServerDataType(structType any) RedundantServerDataType {
	if casted, ok := structType.(RedundantServerDataType); ok {
		return casted
	}
	if casted, ok := structType.(*RedundantServerDataType); ok {
		return *casted
	}
	return nil
}

func (m *_RedundantServerDataType) GetTypeName() string {
	return "RedundantServerDataType"
}

func (m *_RedundantServerDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (serverId)
	lengthInBits += m.ServerId.GetLengthInBits(ctx)

	// Simple field (serviceLevel)
	lengthInBits += 8

	// Simple field (serverState)
	lengthInBits += 32

	return lengthInBits
}

func (m *_RedundantServerDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RedundantServerDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__redundantServerDataType RedundantServerDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RedundantServerDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RedundantServerDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	serverId, err := ReadSimpleField[PascalString](ctx, "serverId", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serverId' field"))
	}
	m.ServerId = serverId

	serviceLevel, err := ReadSimpleField(ctx, "serviceLevel", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceLevel' field"))
	}
	m.ServiceLevel = serviceLevel

	serverState, err := ReadEnumField[ServerState](ctx, "serverState", "ServerState", ReadEnum(ServerStateByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serverState' field"))
	}
	m.ServerState = serverState

	if closeErr := readBuffer.CloseContext("RedundantServerDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RedundantServerDataType")
	}

	return m, nil
}

func (m *_RedundantServerDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RedundantServerDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RedundantServerDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RedundantServerDataType")
		}

		if err := WriteSimpleField[PascalString](ctx, "serverId", m.GetServerId(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'serverId' field")
		}

		if err := WriteSimpleField[uint8](ctx, "serviceLevel", m.GetServiceLevel(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'serviceLevel' field")
		}

		if err := WriteSimpleEnumField[ServerState](ctx, "serverState", "ServerState", m.GetServerState(), WriteEnum[ServerState, uint32](ServerState.GetValue, ServerState.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'serverState' field")
		}

		if popErr := writeBuffer.PopContext("RedundantServerDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RedundantServerDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RedundantServerDataType) IsRedundantServerDataType() {}

func (m *_RedundantServerDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
