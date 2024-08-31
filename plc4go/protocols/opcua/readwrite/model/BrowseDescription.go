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

// BrowseDescription is the corresponding interface of BrowseDescription
type BrowseDescription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNodeId returns NodeId (property field)
	GetNodeId() NodeId
	// GetBrowseDirection returns BrowseDirection (property field)
	GetBrowseDirection() BrowseDirection
	// GetReferenceTypeId returns ReferenceTypeId (property field)
	GetReferenceTypeId() NodeId
	// GetIncludeSubtypes returns IncludeSubtypes (property field)
	GetIncludeSubtypes() bool
	// GetNodeClassMask returns NodeClassMask (property field)
	GetNodeClassMask() uint32
	// GetResultMask returns ResultMask (property field)
	GetResultMask() uint32
}

// BrowseDescriptionExactly can be used when we want exactly this type and not a type which fulfills BrowseDescription.
// This is useful for switch cases.
type BrowseDescriptionExactly interface {
	BrowseDescription
	isBrowseDescription() bool
}

// _BrowseDescription is the data-structure of this message
type _BrowseDescription struct {
	*_ExtensionObjectDefinition
	NodeId          NodeId
	BrowseDirection BrowseDirection
	ReferenceTypeId NodeId
	IncludeSubtypes bool
	NodeClassMask   uint32
	ResultMask      uint32
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BrowseDescription) GetIdentifier() string {
	return "516"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BrowseDescription) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_BrowseDescription) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BrowseDescription) GetNodeId() NodeId {
	return m.NodeId
}

func (m *_BrowseDescription) GetBrowseDirection() BrowseDirection {
	return m.BrowseDirection
}

func (m *_BrowseDescription) GetReferenceTypeId() NodeId {
	return m.ReferenceTypeId
}

func (m *_BrowseDescription) GetIncludeSubtypes() bool {
	return m.IncludeSubtypes
}

func (m *_BrowseDescription) GetNodeClassMask() uint32 {
	return m.NodeClassMask
}

func (m *_BrowseDescription) GetResultMask() uint32 {
	return m.ResultMask
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBrowseDescription factory function for _BrowseDescription
func NewBrowseDescription(nodeId NodeId, browseDirection BrowseDirection, referenceTypeId NodeId, includeSubtypes bool, nodeClassMask uint32, resultMask uint32) *_BrowseDescription {
	_result := &_BrowseDescription{
		NodeId:                     nodeId,
		BrowseDirection:            browseDirection,
		ReferenceTypeId:            referenceTypeId,
		IncludeSubtypes:            includeSubtypes,
		NodeClassMask:              nodeClassMask,
		ResultMask:                 resultMask,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBrowseDescription(structType any) BrowseDescription {
	if casted, ok := structType.(BrowseDescription); ok {
		return casted
	}
	if casted, ok := structType.(*BrowseDescription); ok {
		return *casted
	}
	return nil
}

func (m *_BrowseDescription) GetTypeName() string {
	return "BrowseDescription"
}

func (m *_BrowseDescription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (nodeId)
	lengthInBits += m.NodeId.GetLengthInBits(ctx)

	// Simple field (browseDirection)
	lengthInBits += 32

	// Simple field (referenceTypeId)
	lengthInBits += m.ReferenceTypeId.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (includeSubtypes)
	lengthInBits += 1

	// Simple field (nodeClassMask)
	lengthInBits += 32

	// Simple field (resultMask)
	lengthInBits += 32

	return lengthInBits
}

func (m *_BrowseDescription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BrowseDescriptionParse(ctx context.Context, theBytes []byte, identifier string) (BrowseDescription, error) {
	return BrowseDescriptionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func BrowseDescriptionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (BrowseDescription, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BrowseDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BrowseDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nodeId)
	if pullErr := readBuffer.PullContext("nodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nodeId")
	}
	_nodeId, _nodeIdErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _nodeIdErr != nil {
		return nil, errors.Wrap(_nodeIdErr, "Error parsing 'nodeId' field of BrowseDescription")
	}
	nodeId := _nodeId.(NodeId)
	if closeErr := readBuffer.CloseContext("nodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nodeId")
	}

	// Simple Field (browseDirection)
	if pullErr := readBuffer.PullContext("browseDirection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for browseDirection")
	}
	_browseDirection, _browseDirectionErr := BrowseDirectionParseWithBuffer(ctx, readBuffer)
	if _browseDirectionErr != nil {
		return nil, errors.Wrap(_browseDirectionErr, "Error parsing 'browseDirection' field of BrowseDescription")
	}
	browseDirection := _browseDirection
	if closeErr := readBuffer.CloseContext("browseDirection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for browseDirection")
	}

	// Simple Field (referenceTypeId)
	if pullErr := readBuffer.PullContext("referenceTypeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for referenceTypeId")
	}
	_referenceTypeId, _referenceTypeIdErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _referenceTypeIdErr != nil {
		return nil, errors.Wrap(_referenceTypeIdErr, "Error parsing 'referenceTypeId' field of BrowseDescription")
	}
	referenceTypeId := _referenceTypeId.(NodeId)
	if closeErr := readBuffer.CloseContext("referenceTypeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for referenceTypeId")
	}

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, 7), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	// Simple Field (includeSubtypes)
	_includeSubtypes, _includeSubtypesErr := /*TODO: migrate me*/ readBuffer.ReadBit("includeSubtypes")
	if _includeSubtypesErr != nil {
		return nil, errors.Wrap(_includeSubtypesErr, "Error parsing 'includeSubtypes' field of BrowseDescription")
	}
	includeSubtypes := _includeSubtypes

	// Simple Field (nodeClassMask)
	_nodeClassMask, _nodeClassMaskErr := /*TODO: migrate me*/ readBuffer.ReadUint32("nodeClassMask", 32)
	if _nodeClassMaskErr != nil {
		return nil, errors.Wrap(_nodeClassMaskErr, "Error parsing 'nodeClassMask' field of BrowseDescription")
	}
	nodeClassMask := _nodeClassMask

	// Simple Field (resultMask)
	_resultMask, _resultMaskErr := /*TODO: migrate me*/ readBuffer.ReadUint32("resultMask", 32)
	if _resultMaskErr != nil {
		return nil, errors.Wrap(_resultMaskErr, "Error parsing 'resultMask' field of BrowseDescription")
	}
	resultMask := _resultMask

	if closeErr := readBuffer.CloseContext("BrowseDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BrowseDescription")
	}

	// Create a partially initialized instance
	_child := &_BrowseDescription{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		NodeId:                     nodeId,
		BrowseDirection:            browseDirection,
		ReferenceTypeId:            referenceTypeId,
		IncludeSubtypes:            includeSubtypes,
		NodeClassMask:              nodeClassMask,
		ResultMask:                 resultMask,
		reservedField0:             reservedField0,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_BrowseDescription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BrowseDescription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BrowseDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BrowseDescription")
		}

		// Simple Field (nodeId)
		if pushErr := writeBuffer.PushContext("nodeId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nodeId")
		}
		_nodeIdErr := writeBuffer.WriteSerializable(ctx, m.GetNodeId())
		if popErr := writeBuffer.PopContext("nodeId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nodeId")
		}
		if _nodeIdErr != nil {
			return errors.Wrap(_nodeIdErr, "Error serializing 'nodeId' field")
		}

		// Simple Field (browseDirection)
		if pushErr := writeBuffer.PushContext("browseDirection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for browseDirection")
		}
		_browseDirectionErr := writeBuffer.WriteSerializable(ctx, m.GetBrowseDirection())
		if popErr := writeBuffer.PopContext("browseDirection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for browseDirection")
		}
		if _browseDirectionErr != nil {
			return errors.Wrap(_browseDirectionErr, "Error serializing 'browseDirection' field")
		}

		// Simple Field (referenceTypeId)
		if pushErr := writeBuffer.PushContext("referenceTypeId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for referenceTypeId")
		}
		_referenceTypeIdErr := writeBuffer.WriteSerializable(ctx, m.GetReferenceTypeId())
		if popErr := writeBuffer.PopContext("referenceTypeId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for referenceTypeId")
		}
		if _referenceTypeIdErr != nil {
			return errors.Wrap(_referenceTypeIdErr, "Error serializing 'referenceTypeId' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := /*TODO: migrate me*/ writeBuffer.WriteUint8("reserved", 7, uint8(reserved))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (includeSubtypes)
		includeSubtypes := bool(m.GetIncludeSubtypes())
		_includeSubtypesErr := /*TODO: migrate me*/ writeBuffer.WriteBit("includeSubtypes", (includeSubtypes))
		if _includeSubtypesErr != nil {
			return errors.Wrap(_includeSubtypesErr, "Error serializing 'includeSubtypes' field")
		}

		// Simple Field (nodeClassMask)
		nodeClassMask := uint32(m.GetNodeClassMask())
		_nodeClassMaskErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("nodeClassMask", 32, uint32((nodeClassMask)))
		if _nodeClassMaskErr != nil {
			return errors.Wrap(_nodeClassMaskErr, "Error serializing 'nodeClassMask' field")
		}

		// Simple Field (resultMask)
		resultMask := uint32(m.GetResultMask())
		_resultMaskErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("resultMask", 32, uint32((resultMask)))
		if _resultMaskErr != nil {
			return errors.Wrap(_resultMaskErr, "Error serializing 'resultMask' field")
		}

		if popErr := writeBuffer.PopContext("BrowseDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BrowseDescription")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BrowseDescription) isBrowseDescription() bool {
	return true
}

func (m *_BrowseDescription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
