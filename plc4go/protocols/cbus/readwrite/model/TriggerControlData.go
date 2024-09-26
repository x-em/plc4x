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

// TriggerControlData is the corresponding interface of TriggerControlData
type TriggerControlData interface {
	TriggerControlDataContract
	TriggerControlDataRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsTriggerControlData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTriggerControlData()
}

// TriggerControlDataContract provides a set of functions which can be overwritten by a sub struct
type TriggerControlDataContract interface {
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() TriggerControlCommandTypeContainer
	// GetTriggerGroup returns TriggerGroup (property field)
	GetTriggerGroup() byte
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() TriggerControlCommandType
	// GetIsUnused returns IsUnused (virtual field)
	GetIsUnused() bool
	// IsTriggerControlData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTriggerControlData()
}

// TriggerControlDataRequirements provides a set of functions which need to be implemented by a sub struct
type TriggerControlDataRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetCommandType returns CommandType (discriminator field)
	GetCommandType() TriggerControlCommandType
}

// _TriggerControlData is the data-structure of this message
type _TriggerControlData struct {
	_SubType             TriggerControlData
	CommandTypeContainer TriggerControlCommandTypeContainer
	TriggerGroup         byte
}

var _ TriggerControlDataContract = (*_TriggerControlData)(nil)

// NewTriggerControlData factory function for _TriggerControlData
func NewTriggerControlData(commandTypeContainer TriggerControlCommandTypeContainer, triggerGroup byte) *_TriggerControlData {
	return &_TriggerControlData{CommandTypeContainer: commandTypeContainer, TriggerGroup: triggerGroup}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TriggerControlData) GetCommandTypeContainer() TriggerControlCommandTypeContainer {
	return m.CommandTypeContainer
}

func (m *_TriggerControlData) GetTriggerGroup() byte {
	return m.TriggerGroup
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_TriggerControlData) GetCommandType() TriggerControlCommandType {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return CastTriggerControlCommandType(m.GetCommandTypeContainer().CommandType())
}

func (pm *_TriggerControlData) GetIsUnused() bool {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetTriggerGroup()) > (0xFE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastTriggerControlData(structType any) TriggerControlData {
	if casted, ok := structType.(TriggerControlData); ok {
		return casted
	}
	if casted, ok := structType.(*TriggerControlData); ok {
		return *casted
	}
	return nil
}

func (m *_TriggerControlData) GetTypeName() string {
	return "TriggerControlData"
}

func (m *_TriggerControlData) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (triggerGroup)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_TriggerControlData) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func TriggerControlDataParse[T TriggerControlData](ctx context.Context, theBytes []byte) (T, error) {
	return TriggerControlDataParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func TriggerControlDataParseWithBufferProducer[T TriggerControlData]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := TriggerControlDataParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func TriggerControlDataParseWithBuffer[T TriggerControlData](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_TriggerControlData{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_TriggerControlData) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__triggerControlData TriggerControlData, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TriggerControlData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TriggerControlData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsTriggerControlCommandTypeContainer(ctx, readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "no command type could be found"})
	}

	commandTypeContainer, err := ReadEnumField[TriggerControlCommandTypeContainer](ctx, "commandTypeContainer", "TriggerControlCommandTypeContainer", ReadEnum(TriggerControlCommandTypeContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandTypeContainer' field"))
	}
	m.CommandTypeContainer = commandTypeContainer

	commandType, err := ReadVirtualField[TriggerControlCommandType](ctx, "commandType", (*TriggerControlCommandType)(nil), commandTypeContainer.CommandType())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandType' field"))
	}
	_ = commandType

	triggerGroup, err := ReadSimpleField(ctx, "triggerGroup", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'triggerGroup' field"))
	}
	m.TriggerGroup = triggerGroup

	isUnused, err := ReadVirtualField[bool](ctx, "isUnused", (*bool)(nil), bool((triggerGroup) > (0xFE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isUnused' field"))
	}
	_ = isUnused

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child TriggerControlData
	switch {
	case commandType == TriggerControlCommandType_TRIGGER_EVENT: // TriggerControlDataTriggerEvent
		if _child, err = new(_TriggerControlDataTriggerEvent).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TriggerControlDataTriggerEvent for type-switch of TriggerControlData")
		}
	case commandType == TriggerControlCommandType_TRIGGER_MIN: // TriggerControlDataTriggerMin
		if _child, err = new(_TriggerControlDataTriggerMin).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TriggerControlDataTriggerMin for type-switch of TriggerControlData")
		}
	case commandType == TriggerControlCommandType_TRIGGER_MAX: // TriggerControlDataTriggerMax
		if _child, err = new(_TriggerControlDataTriggerMax).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TriggerControlDataTriggerMax for type-switch of TriggerControlData")
		}
	case commandType == TriggerControlCommandType_INDICATOR_KILL: // TriggerControlDataIndicatorKill
		if _child, err = new(_TriggerControlDataIndicatorKill).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TriggerControlDataIndicatorKill for type-switch of TriggerControlData")
		}
	case commandType == TriggerControlCommandType_LABEL: // TriggerControlDataLabel
		if _child, err = new(_TriggerControlDataLabel).parse(ctx, readBuffer, m, commandTypeContainer); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TriggerControlDataLabel for type-switch of TriggerControlData")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [commandType=%v]", commandType)
	}

	if closeErr := readBuffer.CloseContext("TriggerControlData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TriggerControlData")
	}

	return _child, nil
}

func (pm *_TriggerControlData) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child TriggerControlData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("TriggerControlData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TriggerControlData")
	}

	if err := WriteSimpleEnumField[TriggerControlCommandTypeContainer](ctx, "commandTypeContainer", "TriggerControlCommandTypeContainer", m.GetCommandTypeContainer(), WriteEnum[TriggerControlCommandTypeContainer, uint8](TriggerControlCommandTypeContainer.GetValue, TriggerControlCommandTypeContainer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	commandType := m.GetCommandType()
	_ = commandType
	if _commandTypeErr := writeBuffer.WriteVirtual(ctx, "commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	if err := WriteSimpleField[byte](ctx, "triggerGroup", m.GetTriggerGroup(), WriteByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'triggerGroup' field")
	}
	// Virtual field
	isUnused := m.GetIsUnused()
	_ = isUnused
	if _isUnusedErr := writeBuffer.WriteVirtual(ctx, "isUnused", m.GetIsUnused()); _isUnusedErr != nil {
		return errors.Wrap(_isUnusedErr, "Error serializing 'isUnused' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("TriggerControlData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TriggerControlData")
	}
	return nil
}

func (m *_TriggerControlData) IsTriggerControlData() {}
