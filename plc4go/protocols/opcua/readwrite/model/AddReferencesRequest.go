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

// AddReferencesRequest is the corresponding interface of AddReferencesRequest
type AddReferencesRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfReferencesToAdd returns NoOfReferencesToAdd (property field)
	GetNoOfReferencesToAdd() int32
	// GetReferencesToAdd returns ReferencesToAdd (property field)
	GetReferencesToAdd() []ExtensionObjectDefinition
	// IsAddReferencesRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAddReferencesRequest()
}

// _AddReferencesRequest is the data-structure of this message
type _AddReferencesRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader       ExtensionObjectDefinition
	NoOfReferencesToAdd int32
	ReferencesToAdd     []ExtensionObjectDefinition
}

var _ AddReferencesRequest = (*_AddReferencesRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_AddReferencesRequest)(nil)

// NewAddReferencesRequest factory function for _AddReferencesRequest
func NewAddReferencesRequest(requestHeader ExtensionObjectDefinition, noOfReferencesToAdd int32, referencesToAdd []ExtensionObjectDefinition) *_AddReferencesRequest {
	if requestHeader == nil {
		panic("requestHeader of type ExtensionObjectDefinition for AddReferencesRequest must not be nil")
	}
	_result := &_AddReferencesRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		NoOfReferencesToAdd:               noOfReferencesToAdd,
		ReferencesToAdd:                   referencesToAdd,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AddReferencesRequest) GetIdentifier() string {
	return "494"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AddReferencesRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AddReferencesRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_AddReferencesRequest) GetNoOfReferencesToAdd() int32 {
	return m.NoOfReferencesToAdd
}

func (m *_AddReferencesRequest) GetReferencesToAdd() []ExtensionObjectDefinition {
	return m.ReferencesToAdd
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAddReferencesRequest(structType any) AddReferencesRequest {
	if casted, ok := structType.(AddReferencesRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AddReferencesRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AddReferencesRequest) GetTypeName() string {
	return "AddReferencesRequest"
}

func (m *_AddReferencesRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfReferencesToAdd)
	lengthInBits += 32

	// Array field
	if len(m.ReferencesToAdd) > 0 {
		for _curItem, element := range m.ReferencesToAdd {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ReferencesToAdd), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_AddReferencesRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AddReferencesRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__addReferencesRequest AddReferencesRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AddReferencesRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AddReferencesRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	noOfReferencesToAdd, err := ReadSimpleField(ctx, "noOfReferencesToAdd", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfReferencesToAdd' field"))
	}
	m.NoOfReferencesToAdd = noOfReferencesToAdd

	referencesToAdd, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "referencesToAdd", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("381")), readBuffer), uint64(noOfReferencesToAdd))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referencesToAdd' field"))
	}
	m.ReferencesToAdd = referencesToAdd

	if closeErr := readBuffer.CloseContext("AddReferencesRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AddReferencesRequest")
	}

	return m, nil
}

func (m *_AddReferencesRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AddReferencesRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AddReferencesRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AddReferencesRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfReferencesToAdd", m.GetNoOfReferencesToAdd(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfReferencesToAdd' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "referencesToAdd", m.GetReferencesToAdd(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'referencesToAdd' field")
		}

		if popErr := writeBuffer.PopContext("AddReferencesRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AddReferencesRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AddReferencesRequest) IsAddReferencesRequest() {}

func (m *_AddReferencesRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
