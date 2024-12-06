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

// EventNotificationList is the corresponding interface of EventNotificationList
type EventNotificationList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetEvents returns Events (property field)
	GetEvents() []EventFieldList
	// IsEventNotificationList is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsEventNotificationList()
	// CreateBuilder creates a EventNotificationListBuilder
	CreateEventNotificationListBuilder() EventNotificationListBuilder
}

// _EventNotificationList is the data-structure of this message
type _EventNotificationList struct {
	ExtensionObjectDefinitionContract
	Events []EventFieldList
}

var _ EventNotificationList = (*_EventNotificationList)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_EventNotificationList)(nil)

// NewEventNotificationList factory function for _EventNotificationList
func NewEventNotificationList(events []EventFieldList) *_EventNotificationList {
	_result := &_EventNotificationList{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Events:                            events,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// EventNotificationListBuilder is a builder for EventNotificationList
type EventNotificationListBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(events []EventFieldList) EventNotificationListBuilder
	// WithEvents adds Events (property field)
	WithEvents(...EventFieldList) EventNotificationListBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the EventNotificationList or returns an error if something is wrong
	Build() (EventNotificationList, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() EventNotificationList
}

// NewEventNotificationListBuilder() creates a EventNotificationListBuilder
func NewEventNotificationListBuilder() EventNotificationListBuilder {
	return &_EventNotificationListBuilder{_EventNotificationList: new(_EventNotificationList)}
}

type _EventNotificationListBuilder struct {
	*_EventNotificationList

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (EventNotificationListBuilder) = (*_EventNotificationListBuilder)(nil)

func (b *_EventNotificationListBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._EventNotificationList
}

func (b *_EventNotificationListBuilder) WithMandatoryFields(events []EventFieldList) EventNotificationListBuilder {
	return b.WithEvents(events...)
}

func (b *_EventNotificationListBuilder) WithEvents(events ...EventFieldList) EventNotificationListBuilder {
	b.Events = events
	return b
}

func (b *_EventNotificationListBuilder) Build() (EventNotificationList, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._EventNotificationList.deepCopy(), nil
}

func (b *_EventNotificationListBuilder) MustBuild() EventNotificationList {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_EventNotificationListBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_EventNotificationListBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_EventNotificationListBuilder) DeepCopy() any {
	_copy := b.CreateEventNotificationListBuilder().(*_EventNotificationListBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateEventNotificationListBuilder creates a EventNotificationListBuilder
func (b *_EventNotificationList) CreateEventNotificationListBuilder() EventNotificationListBuilder {
	if b == nil {
		return NewEventNotificationListBuilder()
	}
	return &_EventNotificationListBuilder{_EventNotificationList: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EventNotificationList) GetExtensionId() int32 {
	return int32(916)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EventNotificationList) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EventNotificationList) GetEvents() []EventFieldList {
	return m.Events
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastEventNotificationList(structType any) EventNotificationList {
	if casted, ok := structType.(EventNotificationList); ok {
		return casted
	}
	if casted, ok := structType.(*EventNotificationList); ok {
		return *casted
	}
	return nil
}

func (m *_EventNotificationList) GetTypeName() string {
	return "EventNotificationList"
}

func (m *_EventNotificationList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Implicit Field (noOfEvents)
	lengthInBits += 32

	// Array field
	if len(m.Events) > 0 {
		for _curItem, element := range m.Events {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Events), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_EventNotificationList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_EventNotificationList) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__eventNotificationList EventNotificationList, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EventNotificationList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EventNotificationList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	noOfEvents, err := ReadImplicitField[int32](ctx, "noOfEvents", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfEvents' field"))
	}
	_ = noOfEvents

	events, err := ReadCountArrayField[EventFieldList](ctx, "events", ReadComplex[EventFieldList](ExtensionObjectDefinitionParseWithBufferProducer[EventFieldList]((int32)(int32(919))), readBuffer), uint64(noOfEvents))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'events' field"))
	}
	m.Events = events

	if closeErr := readBuffer.CloseContext("EventNotificationList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EventNotificationList")
	}

	return m, nil
}

func (m *_EventNotificationList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EventNotificationList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EventNotificationList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EventNotificationList")
		}
		noOfEvents := int32(utils.InlineIf(bool((m.GetEvents()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetEvents()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfEvents", noOfEvents, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfEvents' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "events", m.GetEvents(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'events' field")
		}

		if popErr := writeBuffer.PopContext("EventNotificationList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EventNotificationList")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EventNotificationList) IsEventNotificationList() {}

func (m *_EventNotificationList) DeepCopy() any {
	return m.deepCopy()
}

func (m *_EventNotificationList) deepCopy() *_EventNotificationList {
	if m == nil {
		return nil
	}
	_EventNotificationListCopy := &_EventNotificationList{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopySlice[EventFieldList, EventFieldList](m.Events),
	}
	_EventNotificationListCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _EventNotificationListCopy
}

func (m *_EventNotificationList) String() string {
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