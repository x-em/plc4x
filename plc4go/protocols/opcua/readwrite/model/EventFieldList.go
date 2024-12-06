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

// EventFieldList is the corresponding interface of EventFieldList
type EventFieldList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetClientHandle returns ClientHandle (property field)
	GetClientHandle() uint32
	// GetEventFields returns EventFields (property field)
	GetEventFields() []Variant
	// IsEventFieldList is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsEventFieldList()
	// CreateBuilder creates a EventFieldListBuilder
	CreateEventFieldListBuilder() EventFieldListBuilder
}

// _EventFieldList is the data-structure of this message
type _EventFieldList struct {
	ExtensionObjectDefinitionContract
	ClientHandle uint32
	EventFields  []Variant
}

var _ EventFieldList = (*_EventFieldList)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_EventFieldList)(nil)

// NewEventFieldList factory function for _EventFieldList
func NewEventFieldList(clientHandle uint32, eventFields []Variant) *_EventFieldList {
	_result := &_EventFieldList{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ClientHandle:                      clientHandle,
		EventFields:                       eventFields,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// EventFieldListBuilder is a builder for EventFieldList
type EventFieldListBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(clientHandle uint32, eventFields []Variant) EventFieldListBuilder
	// WithClientHandle adds ClientHandle (property field)
	WithClientHandle(uint32) EventFieldListBuilder
	// WithEventFields adds EventFields (property field)
	WithEventFields(...Variant) EventFieldListBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the EventFieldList or returns an error if something is wrong
	Build() (EventFieldList, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() EventFieldList
}

// NewEventFieldListBuilder() creates a EventFieldListBuilder
func NewEventFieldListBuilder() EventFieldListBuilder {
	return &_EventFieldListBuilder{_EventFieldList: new(_EventFieldList)}
}

type _EventFieldListBuilder struct {
	*_EventFieldList

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (EventFieldListBuilder) = (*_EventFieldListBuilder)(nil)

func (b *_EventFieldListBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._EventFieldList
}

func (b *_EventFieldListBuilder) WithMandatoryFields(clientHandle uint32, eventFields []Variant) EventFieldListBuilder {
	return b.WithClientHandle(clientHandle).WithEventFields(eventFields...)
}

func (b *_EventFieldListBuilder) WithClientHandle(clientHandle uint32) EventFieldListBuilder {
	b.ClientHandle = clientHandle
	return b
}

func (b *_EventFieldListBuilder) WithEventFields(eventFields ...Variant) EventFieldListBuilder {
	b.EventFields = eventFields
	return b
}

func (b *_EventFieldListBuilder) Build() (EventFieldList, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._EventFieldList.deepCopy(), nil
}

func (b *_EventFieldListBuilder) MustBuild() EventFieldList {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_EventFieldListBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_EventFieldListBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_EventFieldListBuilder) DeepCopy() any {
	_copy := b.CreateEventFieldListBuilder().(*_EventFieldListBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateEventFieldListBuilder creates a EventFieldListBuilder
func (b *_EventFieldList) CreateEventFieldListBuilder() EventFieldListBuilder {
	if b == nil {
		return NewEventFieldListBuilder()
	}
	return &_EventFieldListBuilder{_EventFieldList: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EventFieldList) GetExtensionId() int32 {
	return int32(919)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EventFieldList) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EventFieldList) GetClientHandle() uint32 {
	return m.ClientHandle
}

func (m *_EventFieldList) GetEventFields() []Variant {
	return m.EventFields
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastEventFieldList(structType any) EventFieldList {
	if casted, ok := structType.(EventFieldList); ok {
		return casted
	}
	if casted, ok := structType.(*EventFieldList); ok {
		return *casted
	}
	return nil
}

func (m *_EventFieldList) GetTypeName() string {
	return "EventFieldList"
}

func (m *_EventFieldList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (clientHandle)
	lengthInBits += 32

	// Implicit Field (noOfEventFields)
	lengthInBits += 32

	// Array field
	if len(m.EventFields) > 0 {
		for _curItem, element := range m.EventFields {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.EventFields), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_EventFieldList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_EventFieldList) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__eventFieldList EventFieldList, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EventFieldList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EventFieldList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	clientHandle, err := ReadSimpleField(ctx, "clientHandle", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'clientHandle' field"))
	}
	m.ClientHandle = clientHandle

	noOfEventFields, err := ReadImplicitField[int32](ctx, "noOfEventFields", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfEventFields' field"))
	}
	_ = noOfEventFields

	eventFields, err := ReadCountArrayField[Variant](ctx, "eventFields", ReadComplex[Variant](VariantParseWithBuffer, readBuffer), uint64(noOfEventFields))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventFields' field"))
	}
	m.EventFields = eventFields

	if closeErr := readBuffer.CloseContext("EventFieldList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EventFieldList")
	}

	return m, nil
}

func (m *_EventFieldList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EventFieldList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EventFieldList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EventFieldList")
		}

		if err := WriteSimpleField[uint32](ctx, "clientHandle", m.GetClientHandle(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'clientHandle' field")
		}
		noOfEventFields := int32(utils.InlineIf(bool((m.GetEventFields()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetEventFields()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfEventFields", noOfEventFields, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfEventFields' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "eventFields", m.GetEventFields(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'eventFields' field")
		}

		if popErr := writeBuffer.PopContext("EventFieldList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EventFieldList")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EventFieldList) IsEventFieldList() {}

func (m *_EventFieldList) DeepCopy() any {
	return m.deepCopy()
}

func (m *_EventFieldList) deepCopy() *_EventFieldList {
	if m == nil {
		return nil
	}
	_EventFieldListCopy := &_EventFieldList{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.ClientHandle,
		utils.DeepCopySlice[Variant, Variant](m.EventFields),
	}
	_EventFieldListCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _EventFieldListCopy
}

func (m *_EventFieldList) String() string {
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