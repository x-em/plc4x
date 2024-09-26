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

// RegisterNodesResponse is the corresponding interface of RegisterNodesResponse
type RegisterNodesResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetNoOfRegisteredNodeIds returns NoOfRegisteredNodeIds (property field)
	GetNoOfRegisteredNodeIds() int32
	// GetRegisteredNodeIds returns RegisteredNodeIds (property field)
	GetRegisteredNodeIds() []NodeId
	// IsRegisterNodesResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRegisterNodesResponse()
}

// _RegisterNodesResponse is the data-structure of this message
type _RegisterNodesResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader        ExtensionObjectDefinition
	NoOfRegisteredNodeIds int32
	RegisteredNodeIds     []NodeId
}

var _ RegisterNodesResponse = (*_RegisterNodesResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_RegisterNodesResponse)(nil)

// NewRegisterNodesResponse factory function for _RegisterNodesResponse
func NewRegisterNodesResponse(responseHeader ExtensionObjectDefinition, noOfRegisteredNodeIds int32, registeredNodeIds []NodeId) *_RegisterNodesResponse {
	if responseHeader == nil {
		panic("responseHeader of type ExtensionObjectDefinition for RegisterNodesResponse must not be nil")
	}
	_result := &_RegisterNodesResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		NoOfRegisteredNodeIds:             noOfRegisteredNodeIds,
		RegisteredNodeIds:                 registeredNodeIds,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RegisterNodesResponse) GetIdentifier() string {
	return "563"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RegisterNodesResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RegisterNodesResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_RegisterNodesResponse) GetNoOfRegisteredNodeIds() int32 {
	return m.NoOfRegisteredNodeIds
}

func (m *_RegisterNodesResponse) GetRegisteredNodeIds() []NodeId {
	return m.RegisteredNodeIds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastRegisterNodesResponse(structType any) RegisterNodesResponse {
	if casted, ok := structType.(RegisterNodesResponse); ok {
		return casted
	}
	if casted, ok := structType.(*RegisterNodesResponse); ok {
		return *casted
	}
	return nil
}

func (m *_RegisterNodesResponse) GetTypeName() string {
	return "RegisterNodesResponse"
}

func (m *_RegisterNodesResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (noOfRegisteredNodeIds)
	lengthInBits += 32

	// Array field
	if len(m.RegisteredNodeIds) > 0 {
		for _curItem, element := range m.RegisteredNodeIds {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.RegisteredNodeIds), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_RegisterNodesResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RegisterNodesResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__registerNodesResponse RegisterNodesResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RegisterNodesResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RegisterNodesResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("394")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	noOfRegisteredNodeIds, err := ReadSimpleField(ctx, "noOfRegisteredNodeIds", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfRegisteredNodeIds' field"))
	}
	m.NoOfRegisteredNodeIds = noOfRegisteredNodeIds

	registeredNodeIds, err := ReadCountArrayField[NodeId](ctx, "registeredNodeIds", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer), uint64(noOfRegisteredNodeIds))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'registeredNodeIds' field"))
	}
	m.RegisteredNodeIds = registeredNodeIds

	if closeErr := readBuffer.CloseContext("RegisterNodesResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RegisterNodesResponse")
	}

	return m, nil
}

func (m *_RegisterNodesResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RegisterNodesResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RegisterNodesResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RegisterNodesResponse")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfRegisteredNodeIds", m.GetNoOfRegisteredNodeIds(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfRegisteredNodeIds' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "registeredNodeIds", m.GetRegisteredNodeIds(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'registeredNodeIds' field")
		}

		if popErr := writeBuffer.PopContext("RegisterNodesResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RegisterNodesResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RegisterNodesResponse) IsRegisterNodesResponse() {}

func (m *_RegisterNodesResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
