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

// NodeIdByteString is the corresponding interface of NodeIdByteString
type NodeIdByteString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	NodeIdTypeDefinition
	// GetNamespaceIndex returns NamespaceIndex (property field)
	GetNamespaceIndex() uint16
	// GetId returns Id (property field)
	GetId() PascalByteString
	// GetIdentifier returns Identifier (virtual field)
	GetIdentifier() string
	// GetNamespace returns Namespace (virtual field)
	GetNamespace() int16
	// IsNodeIdByteString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNodeIdByteString()
	// CreateBuilder creates a NodeIdByteStringBuilder
	CreateNodeIdByteStringBuilder() NodeIdByteStringBuilder
}

// _NodeIdByteString is the data-structure of this message
type _NodeIdByteString struct {
	NodeIdTypeDefinitionContract
	NamespaceIndex uint16
	Id             PascalByteString
}

var _ NodeIdByteString = (*_NodeIdByteString)(nil)
var _ NodeIdTypeDefinitionRequirements = (*_NodeIdByteString)(nil)

// NewNodeIdByteString factory function for _NodeIdByteString
func NewNodeIdByteString(namespaceIndex uint16, id PascalByteString) *_NodeIdByteString {
	if id == nil {
		panic("id of type PascalByteString for NodeIdByteString must not be nil")
	}
	_result := &_NodeIdByteString{
		NodeIdTypeDefinitionContract: NewNodeIdTypeDefinition(),
		NamespaceIndex:               namespaceIndex,
		Id:                           id,
	}
	_result.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NodeIdByteStringBuilder is a builder for NodeIdByteString
type NodeIdByteStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(namespaceIndex uint16, id PascalByteString) NodeIdByteStringBuilder
	// WithNamespaceIndex adds NamespaceIndex (property field)
	WithNamespaceIndex(uint16) NodeIdByteStringBuilder
	// WithId adds Id (property field)
	WithId(PascalByteString) NodeIdByteStringBuilder
	// WithIdBuilder adds Id (property field) which is build by the builder
	WithIdBuilder(func(PascalByteStringBuilder) PascalByteStringBuilder) NodeIdByteStringBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() NodeIdTypeDefinitionBuilder
	// Build builds the NodeIdByteString or returns an error if something is wrong
	Build() (NodeIdByteString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NodeIdByteString
}

// NewNodeIdByteStringBuilder() creates a NodeIdByteStringBuilder
func NewNodeIdByteStringBuilder() NodeIdByteStringBuilder {
	return &_NodeIdByteStringBuilder{_NodeIdByteString: new(_NodeIdByteString)}
}

type _NodeIdByteStringBuilder struct {
	*_NodeIdByteString

	parentBuilder *_NodeIdTypeDefinitionBuilder

	err *utils.MultiError
}

var _ (NodeIdByteStringBuilder) = (*_NodeIdByteStringBuilder)(nil)

func (b *_NodeIdByteStringBuilder) setParent(contract NodeIdTypeDefinitionContract) {
	b.NodeIdTypeDefinitionContract = contract
	contract.(*_NodeIdTypeDefinition)._SubType = b._NodeIdByteString
}

func (b *_NodeIdByteStringBuilder) WithMandatoryFields(namespaceIndex uint16, id PascalByteString) NodeIdByteStringBuilder {
	return b.WithNamespaceIndex(namespaceIndex).WithId(id)
}

func (b *_NodeIdByteStringBuilder) WithNamespaceIndex(namespaceIndex uint16) NodeIdByteStringBuilder {
	b.NamespaceIndex = namespaceIndex
	return b
}

func (b *_NodeIdByteStringBuilder) WithId(id PascalByteString) NodeIdByteStringBuilder {
	b.Id = id
	return b
}

func (b *_NodeIdByteStringBuilder) WithIdBuilder(builderSupplier func(PascalByteStringBuilder) PascalByteStringBuilder) NodeIdByteStringBuilder {
	builder := builderSupplier(b.Id.CreatePascalByteStringBuilder())
	var err error
	b.Id, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalByteStringBuilder failed"))
	}
	return b
}

func (b *_NodeIdByteStringBuilder) Build() (NodeIdByteString, error) {
	if b.Id == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'id' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._NodeIdByteString.deepCopy(), nil
}

func (b *_NodeIdByteStringBuilder) MustBuild() NodeIdByteString {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_NodeIdByteStringBuilder) Done() NodeIdTypeDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewNodeIdTypeDefinitionBuilder().(*_NodeIdTypeDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_NodeIdByteStringBuilder) buildForNodeIdTypeDefinition() (NodeIdTypeDefinition, error) {
	return b.Build()
}

func (b *_NodeIdByteStringBuilder) DeepCopy() any {
	_copy := b.CreateNodeIdByteStringBuilder().(*_NodeIdByteStringBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateNodeIdByteStringBuilder creates a NodeIdByteStringBuilder
func (b *_NodeIdByteString) CreateNodeIdByteStringBuilder() NodeIdByteStringBuilder {
	if b == nil {
		return NewNodeIdByteStringBuilder()
	}
	return &_NodeIdByteStringBuilder{_NodeIdByteString: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NodeIdByteString) GetNodeType() NodeIdType {
	return NodeIdType_nodeIdTypeByteString
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NodeIdByteString) GetParent() NodeIdTypeDefinitionContract {
	return m.NodeIdTypeDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NodeIdByteString) GetNamespaceIndex() uint16 {
	return m.NamespaceIndex
}

func (m *_NodeIdByteString) GetId() PascalByteString {
	return m.Id
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_NodeIdByteString) GetIdentifier() string {
	ctx := context.Background()
	_ = ctx
	return fmt.Sprintf("%v", m.GetId().GetStringValue())
}

func (m *_NodeIdByteString) GetNamespace() int16 {
	ctx := context.Background()
	_ = ctx
	return int16(m.GetNamespaceIndex())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNodeIdByteString(structType any) NodeIdByteString {
	if casted, ok := structType.(NodeIdByteString); ok {
		return casted
	}
	if casted, ok := structType.(*NodeIdByteString); ok {
		return *casted
	}
	return nil
}

func (m *_NodeIdByteString) GetTypeName() string {
	return "NodeIdByteString"
}

func (m *_NodeIdByteString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition).getLengthInBits(ctx))

	// Simple field (namespaceIndex)
	lengthInBits += 16

	// Simple field (id)
	lengthInBits += m.Id.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_NodeIdByteString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NodeIdByteString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NodeIdTypeDefinition) (__nodeIdByteString NodeIdByteString, err error) {
	m.NodeIdTypeDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NodeIdByteString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeIdByteString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	namespaceIndex, err := ReadSimpleField(ctx, "namespaceIndex", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaceIndex' field"))
	}
	m.NamespaceIndex = namespaceIndex

	id, err := ReadSimpleField[PascalByteString](ctx, "id", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'id' field"))
	}
	m.Id = id

	identifier, err := ReadVirtualField[string](ctx, "identifier", (*string)(nil), id.GetStringValue())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'identifier' field"))
	}
	_ = identifier

	namespace, err := ReadVirtualField[int16](ctx, "namespace", (*int16)(nil), namespaceIndex)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespace' field"))
	}
	_ = namespace

	if closeErr := readBuffer.CloseContext("NodeIdByteString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeIdByteString")
	}

	return m, nil
}

func (m *_NodeIdByteString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NodeIdByteString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NodeIdByteString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NodeIdByteString")
		}

		if err := WriteSimpleField[uint16](ctx, "namespaceIndex", m.GetNamespaceIndex(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'namespaceIndex' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "id", m.GetId(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'id' field")
		}
		// Virtual field
		identifier := m.GetIdentifier()
		_ = identifier
		if _identifierErr := writeBuffer.WriteVirtual(ctx, "identifier", m.GetIdentifier()); _identifierErr != nil {
			return errors.Wrap(_identifierErr, "Error serializing 'identifier' field")
		}
		// Virtual field
		namespace := m.GetNamespace()
		_ = namespace
		if _namespaceErr := writeBuffer.WriteVirtual(ctx, "namespace", m.GetNamespace()); _namespaceErr != nil {
			return errors.Wrap(_namespaceErr, "Error serializing 'namespace' field")
		}

		if popErr := writeBuffer.PopContext("NodeIdByteString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NodeIdByteString")
		}
		return nil
	}
	return m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NodeIdByteString) IsNodeIdByteString() {}

func (m *_NodeIdByteString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NodeIdByteString) deepCopy() *_NodeIdByteString {
	if m == nil {
		return nil
	}
	_NodeIdByteStringCopy := &_NodeIdByteString{
		m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition).deepCopy(),
		m.NamespaceIndex,
		utils.DeepCopy[PascalByteString](m.Id),
	}
	_NodeIdByteStringCopy.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition)._SubType = m
	return _NodeIdByteStringCopy
}

func (m *_NodeIdByteString) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}