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

// DeleteNodesItem is the corresponding interface of DeleteNodesItem
type DeleteNodesItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNodeId returns NodeId (property field)
	GetNodeId() NodeId
	// GetDeleteTargetReferences returns DeleteTargetReferences (property field)
	GetDeleteTargetReferences() bool
	// IsDeleteNodesItem is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDeleteNodesItem()
}

// _DeleteNodesItem is the data-structure of this message
type _DeleteNodesItem struct {
	ExtensionObjectDefinitionContract
	NodeId                 NodeId
	DeleteTargetReferences bool
	// Reserved Fields
	reservedField0 *uint8
}

var _ DeleteNodesItem = (*_DeleteNodesItem)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DeleteNodesItem)(nil)

// NewDeleteNodesItem factory function for _DeleteNodesItem
func NewDeleteNodesItem(nodeId NodeId, deleteTargetReferences bool) *_DeleteNodesItem {
	if nodeId == nil {
		panic("nodeId of type NodeId for DeleteNodesItem must not be nil")
	}
	_result := &_DeleteNodesItem{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NodeId:                            nodeId,
		DeleteTargetReferences:            deleteTargetReferences,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeleteNodesItem) GetIdentifier() string {
	return "384"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeleteNodesItem) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeleteNodesItem) GetNodeId() NodeId {
	return m.NodeId
}

func (m *_DeleteNodesItem) GetDeleteTargetReferences() bool {
	return m.DeleteTargetReferences
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDeleteNodesItem(structType any) DeleteNodesItem {
	if casted, ok := structType.(DeleteNodesItem); ok {
		return casted
	}
	if casted, ok := structType.(*DeleteNodesItem); ok {
		return *casted
	}
	return nil
}

func (m *_DeleteNodesItem) GetTypeName() string {
	return "DeleteNodesItem"
}

func (m *_DeleteNodesItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (nodeId)
	lengthInBits += m.NodeId.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (deleteTargetReferences)
	lengthInBits += 1

	return lengthInBits
}

func (m *_DeleteNodesItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DeleteNodesItem) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__deleteNodesItem DeleteNodesItem, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeleteNodesItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeleteNodesItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nodeId, err := ReadSimpleField[NodeId](ctx, "nodeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeId' field"))
	}
	m.NodeId = nodeId

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	deleteTargetReferences, err := ReadSimpleField(ctx, "deleteTargetReferences", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deleteTargetReferences' field"))
	}
	m.DeleteTargetReferences = deleteTargetReferences

	if closeErr := readBuffer.CloseContext("DeleteNodesItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeleteNodesItem")
	}

	return m, nil
}

func (m *_DeleteNodesItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeleteNodesItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeleteNodesItem"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeleteNodesItem")
		}

		if err := WriteSimpleField[NodeId](ctx, "nodeId", m.GetNodeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nodeId' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "deleteTargetReferences", m.GetDeleteTargetReferences(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deleteTargetReferences' field")
		}

		if popErr := writeBuffer.PopContext("DeleteNodesItem"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeleteNodesItem")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DeleteNodesItem) IsDeleteNodesItem() {}

func (m *_DeleteNodesItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
