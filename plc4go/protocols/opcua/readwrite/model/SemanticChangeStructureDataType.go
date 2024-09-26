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

// SemanticChangeStructureDataType is the corresponding interface of SemanticChangeStructureDataType
type SemanticChangeStructureDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetAffected returns Affected (property field)
	GetAffected() NodeId
	// GetAffectedType returns AffectedType (property field)
	GetAffectedType() NodeId
	// IsSemanticChangeStructureDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSemanticChangeStructureDataType()
}

// _SemanticChangeStructureDataType is the data-structure of this message
type _SemanticChangeStructureDataType struct {
	ExtensionObjectDefinitionContract
	Affected     NodeId
	AffectedType NodeId
}

var _ SemanticChangeStructureDataType = (*_SemanticChangeStructureDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_SemanticChangeStructureDataType)(nil)

// NewSemanticChangeStructureDataType factory function for _SemanticChangeStructureDataType
func NewSemanticChangeStructureDataType(affected NodeId, affectedType NodeId) *_SemanticChangeStructureDataType {
	if affected == nil {
		panic("affected of type NodeId for SemanticChangeStructureDataType must not be nil")
	}
	if affectedType == nil {
		panic("affectedType of type NodeId for SemanticChangeStructureDataType must not be nil")
	}
	_result := &_SemanticChangeStructureDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Affected:                          affected,
		AffectedType:                      affectedType,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SemanticChangeStructureDataType) GetIdentifier() string {
	return "899"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SemanticChangeStructureDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SemanticChangeStructureDataType) GetAffected() NodeId {
	return m.Affected
}

func (m *_SemanticChangeStructureDataType) GetAffectedType() NodeId {
	return m.AffectedType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSemanticChangeStructureDataType(structType any) SemanticChangeStructureDataType {
	if casted, ok := structType.(SemanticChangeStructureDataType); ok {
		return casted
	}
	if casted, ok := structType.(*SemanticChangeStructureDataType); ok {
		return *casted
	}
	return nil
}

func (m *_SemanticChangeStructureDataType) GetTypeName() string {
	return "SemanticChangeStructureDataType"
}

func (m *_SemanticChangeStructureDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (affected)
	lengthInBits += m.Affected.GetLengthInBits(ctx)

	// Simple field (affectedType)
	lengthInBits += m.AffectedType.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SemanticChangeStructureDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SemanticChangeStructureDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__semanticChangeStructureDataType SemanticChangeStructureDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SemanticChangeStructureDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SemanticChangeStructureDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	affected, err := ReadSimpleField[NodeId](ctx, "affected", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'affected' field"))
	}
	m.Affected = affected

	affectedType, err := ReadSimpleField[NodeId](ctx, "affectedType", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'affectedType' field"))
	}
	m.AffectedType = affectedType

	if closeErr := readBuffer.CloseContext("SemanticChangeStructureDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SemanticChangeStructureDataType")
	}

	return m, nil
}

func (m *_SemanticChangeStructureDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SemanticChangeStructureDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SemanticChangeStructureDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SemanticChangeStructureDataType")
		}

		if err := WriteSimpleField[NodeId](ctx, "affected", m.GetAffected(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'affected' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "affectedType", m.GetAffectedType(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'affectedType' field")
		}

		if popErr := writeBuffer.PopContext("SemanticChangeStructureDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SemanticChangeStructureDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SemanticChangeStructureDataType) IsSemanticChangeStructureDataType() {}

func (m *_SemanticChangeStructureDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
