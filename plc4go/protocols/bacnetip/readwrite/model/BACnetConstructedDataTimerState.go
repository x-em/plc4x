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

// BACnetConstructedDataTimerState is the corresponding interface of BACnetConstructedDataTimerState
type BACnetConstructedDataTimerState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetTimerState returns TimerState (property field)
	GetTimerState() BACnetTimerStateTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetTimerStateTagged
	// IsBACnetConstructedDataTimerState is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTimerState()
}

// _BACnetConstructedDataTimerState is the data-structure of this message
type _BACnetConstructedDataTimerState struct {
	BACnetConstructedDataContract
	TimerState BACnetTimerStateTagged
}

var _ BACnetConstructedDataTimerState = (*_BACnetConstructedDataTimerState)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTimerState)(nil)

// NewBACnetConstructedDataTimerState factory function for _BACnetConstructedDataTimerState
func NewBACnetConstructedDataTimerState(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, timerState BACnetTimerStateTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimerState {
	if timerState == nil {
		panic("timerState of type BACnetTimerStateTagged for BACnetConstructedDataTimerState must not be nil")
	}
	_result := &_BACnetConstructedDataTimerState{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		TimerState:                    timerState,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimerState) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTimerState) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIMER_STATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimerState) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimerState) GetTimerState() BACnetTimerStateTagged {
	return m.TimerState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTimerState) GetActualValue() BACnetTimerStateTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetTimerStateTagged(m.GetTimerState())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimerState(structType any) BACnetConstructedDataTimerState {
	if casted, ok := structType.(BACnetConstructedDataTimerState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimerState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimerState) GetTypeName() string {
	return "BACnetConstructedDataTimerState"
}

func (m *_BACnetConstructedDataTimerState) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (timerState)
	lengthInBits += m.TimerState.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimerState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTimerState) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTimerState BACnetConstructedDataTimerState, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimerState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimerState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timerState, err := ReadSimpleField[BACnetTimerStateTagged](ctx, "timerState", ReadComplex[BACnetTimerStateTagged](BACnetTimerStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timerState' field"))
	}
	m.TimerState = timerState

	actualValue, err := ReadVirtualField[BACnetTimerStateTagged](ctx, "actualValue", (*BACnetTimerStateTagged)(nil), timerState)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimerState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimerState")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTimerState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimerState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimerState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimerState")
		}

		if err := WriteSimpleField[BACnetTimerStateTagged](ctx, "timerState", m.GetTimerState(), WriteComplex[BACnetTimerStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timerState' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimerState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimerState")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimerState) IsBACnetConstructedDataTimerState() {}

func (m *_BACnetConstructedDataTimerState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
