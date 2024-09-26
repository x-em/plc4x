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

// Constant values.
const AlarmMessageObjectPushType_VARIABLESPEC uint8 = 0x12

// AlarmMessageObjectPushType is the corresponding interface of AlarmMessageObjectPushType
type AlarmMessageObjectPushType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetLengthSpec returns LengthSpec (property field)
	GetLengthSpec() uint8
	// GetSyntaxId returns SyntaxId (property field)
	GetSyntaxId() SyntaxIdType
	// GetNumberOfValues returns NumberOfValues (property field)
	GetNumberOfValues() uint8
	// GetEventId returns EventId (property field)
	GetEventId() uint32
	// GetEventState returns EventState (property field)
	GetEventState() State
	// GetLocalState returns LocalState (property field)
	GetLocalState() State
	// GetAckStateGoing returns AckStateGoing (property field)
	GetAckStateGoing() State
	// GetAckStateComing returns AckStateComing (property field)
	GetAckStateComing() State
	// GetAssociatedValues returns AssociatedValues (property field)
	GetAssociatedValues() []AssociatedValueType
	// IsAlarmMessageObjectPushType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAlarmMessageObjectPushType()
}

// _AlarmMessageObjectPushType is the data-structure of this message
type _AlarmMessageObjectPushType struct {
	LengthSpec       uint8
	SyntaxId         SyntaxIdType
	NumberOfValues   uint8
	EventId          uint32
	EventState       State
	LocalState       State
	AckStateGoing    State
	AckStateComing   State
	AssociatedValues []AssociatedValueType
}

var _ AlarmMessageObjectPushType = (*_AlarmMessageObjectPushType)(nil)

// NewAlarmMessageObjectPushType factory function for _AlarmMessageObjectPushType
func NewAlarmMessageObjectPushType(lengthSpec uint8, syntaxId SyntaxIdType, numberOfValues uint8, eventId uint32, eventState State, localState State, ackStateGoing State, ackStateComing State, AssociatedValues []AssociatedValueType) *_AlarmMessageObjectPushType {
	if eventState == nil {
		panic("eventState of type State for AlarmMessageObjectPushType must not be nil")
	}
	if localState == nil {
		panic("localState of type State for AlarmMessageObjectPushType must not be nil")
	}
	if ackStateGoing == nil {
		panic("ackStateGoing of type State for AlarmMessageObjectPushType must not be nil")
	}
	if ackStateComing == nil {
		panic("ackStateComing of type State for AlarmMessageObjectPushType must not be nil")
	}
	return &_AlarmMessageObjectPushType{LengthSpec: lengthSpec, SyntaxId: syntaxId, NumberOfValues: numberOfValues, EventId: eventId, EventState: eventState, LocalState: localState, AckStateGoing: ackStateGoing, AckStateComing: ackStateComing, AssociatedValues: AssociatedValues}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AlarmMessageObjectPushType) GetLengthSpec() uint8 {
	return m.LengthSpec
}

func (m *_AlarmMessageObjectPushType) GetSyntaxId() SyntaxIdType {
	return m.SyntaxId
}

func (m *_AlarmMessageObjectPushType) GetNumberOfValues() uint8 {
	return m.NumberOfValues
}

func (m *_AlarmMessageObjectPushType) GetEventId() uint32 {
	return m.EventId
}

func (m *_AlarmMessageObjectPushType) GetEventState() State {
	return m.EventState
}

func (m *_AlarmMessageObjectPushType) GetLocalState() State {
	return m.LocalState
}

func (m *_AlarmMessageObjectPushType) GetAckStateGoing() State {
	return m.AckStateGoing
}

func (m *_AlarmMessageObjectPushType) GetAckStateComing() State {
	return m.AckStateComing
}

func (m *_AlarmMessageObjectPushType) GetAssociatedValues() []AssociatedValueType {
	return m.AssociatedValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_AlarmMessageObjectPushType) GetVariableSpec() uint8 {
	return AlarmMessageObjectPushType_VARIABLESPEC
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAlarmMessageObjectPushType(structType any) AlarmMessageObjectPushType {
	if casted, ok := structType.(AlarmMessageObjectPushType); ok {
		return casted
	}
	if casted, ok := structType.(*AlarmMessageObjectPushType); ok {
		return *casted
	}
	return nil
}

func (m *_AlarmMessageObjectPushType) GetTypeName() string {
	return "AlarmMessageObjectPushType"
}

func (m *_AlarmMessageObjectPushType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (variableSpec)
	lengthInBits += 8

	// Simple field (lengthSpec)
	lengthInBits += 8

	// Simple field (syntaxId)
	lengthInBits += 8

	// Simple field (numberOfValues)
	lengthInBits += 8

	// Simple field (eventId)
	lengthInBits += 32

	// Simple field (eventState)
	lengthInBits += m.EventState.GetLengthInBits(ctx)

	// Simple field (localState)
	lengthInBits += m.LocalState.GetLengthInBits(ctx)

	// Simple field (ackStateGoing)
	lengthInBits += m.AckStateGoing.GetLengthInBits(ctx)

	// Simple field (ackStateComing)
	lengthInBits += m.AckStateComing.GetLengthInBits(ctx)

	// Array field
	if len(m.AssociatedValues) > 0 {
		for _curItem, element := range m.AssociatedValues {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.AssociatedValues), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_AlarmMessageObjectPushType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AlarmMessageObjectPushTypeParse(ctx context.Context, theBytes []byte) (AlarmMessageObjectPushType, error) {
	return AlarmMessageObjectPushTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AlarmMessageObjectPushTypeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageObjectPushType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageObjectPushType, error) {
		return AlarmMessageObjectPushTypeParseWithBuffer(ctx, readBuffer)
	}
}

func AlarmMessageObjectPushTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageObjectPushType, error) {
	v, err := (&_AlarmMessageObjectPushType{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_AlarmMessageObjectPushType) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__alarmMessageObjectPushType AlarmMessageObjectPushType, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AlarmMessageObjectPushType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AlarmMessageObjectPushType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	variableSpec, err := ReadConstField[uint8](ctx, "variableSpec", ReadUnsignedByte(readBuffer, uint8(8)), AlarmMessageObjectPushType_VARIABLESPEC)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'variableSpec' field"))
	}
	_ = variableSpec

	lengthSpec, err := ReadSimpleField(ctx, "lengthSpec", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lengthSpec' field"))
	}
	m.LengthSpec = lengthSpec

	syntaxId, err := ReadEnumField[SyntaxIdType](ctx, "syntaxId", "SyntaxIdType", ReadEnum(SyntaxIdTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'syntaxId' field"))
	}
	m.SyntaxId = syntaxId

	numberOfValues, err := ReadSimpleField(ctx, "numberOfValues", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfValues' field"))
	}
	m.NumberOfValues = numberOfValues

	eventId, err := ReadSimpleField(ctx, "eventId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventId' field"))
	}
	m.EventId = eventId

	eventState, err := ReadSimpleField[State](ctx, "eventState", ReadComplex[State](StateParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventState' field"))
	}
	m.EventState = eventState

	localState, err := ReadSimpleField[State](ctx, "localState", ReadComplex[State](StateParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'localState' field"))
	}
	m.LocalState = localState

	ackStateGoing, err := ReadSimpleField[State](ctx, "ackStateGoing", ReadComplex[State](StateParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ackStateGoing' field"))
	}
	m.AckStateGoing = ackStateGoing

	ackStateComing, err := ReadSimpleField[State](ctx, "ackStateComing", ReadComplex[State](StateParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ackStateComing' field"))
	}
	m.AckStateComing = ackStateComing

	AssociatedValues, err := ReadCountArrayField[AssociatedValueType](ctx, "AssociatedValues", ReadComplex[AssociatedValueType](AssociatedValueTypeParseWithBuffer, readBuffer), uint64(numberOfValues))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'AssociatedValues' field"))
	}
	m.AssociatedValues = AssociatedValues

	if closeErr := readBuffer.CloseContext("AlarmMessageObjectPushType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AlarmMessageObjectPushType")
	}

	return m, nil
}

func (m *_AlarmMessageObjectPushType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AlarmMessageObjectPushType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AlarmMessageObjectPushType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AlarmMessageObjectPushType")
	}

	if err := WriteConstField(ctx, "variableSpec", AlarmMessageObjectPushType_VARIABLESPEC, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'variableSpec' field")
	}

	if err := WriteSimpleField[uint8](ctx, "lengthSpec", m.GetLengthSpec(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'lengthSpec' field")
	}

	if err := WriteSimpleEnumField[SyntaxIdType](ctx, "syntaxId", "SyntaxIdType", m.GetSyntaxId(), WriteEnum[SyntaxIdType, uint8](SyntaxIdType.GetValue, SyntaxIdType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'syntaxId' field")
	}

	if err := WriteSimpleField[uint8](ctx, "numberOfValues", m.GetNumberOfValues(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'numberOfValues' field")
	}

	if err := WriteSimpleField[uint32](ctx, "eventId", m.GetEventId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'eventId' field")
	}

	if err := WriteSimpleField[State](ctx, "eventState", m.GetEventState(), WriteComplex[State](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'eventState' field")
	}

	if err := WriteSimpleField[State](ctx, "localState", m.GetLocalState(), WriteComplex[State](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'localState' field")
	}

	if err := WriteSimpleField[State](ctx, "ackStateGoing", m.GetAckStateGoing(), WriteComplex[State](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'ackStateGoing' field")
	}

	if err := WriteSimpleField[State](ctx, "ackStateComing", m.GetAckStateComing(), WriteComplex[State](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'ackStateComing' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "AssociatedValues", m.GetAssociatedValues(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'AssociatedValues' field")
	}

	if popErr := writeBuffer.PopContext("AlarmMessageObjectPushType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AlarmMessageObjectPushType")
	}
	return nil
}

func (m *_AlarmMessageObjectPushType) IsAlarmMessageObjectPushType() {}

func (m *_AlarmMessageObjectPushType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
