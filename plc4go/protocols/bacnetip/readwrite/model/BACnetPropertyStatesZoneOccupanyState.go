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

// BACnetPropertyStatesZoneOccupanyState is the corresponding interface of BACnetPropertyStatesZoneOccupanyState
type BACnetPropertyStatesZoneOccupanyState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetZoneOccupanyState returns ZoneOccupanyState (property field)
	GetZoneOccupanyState() BACnetAccessZoneOccupancyStateTagged
	// IsBACnetPropertyStatesZoneOccupanyState is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesZoneOccupanyState()
}

// _BACnetPropertyStatesZoneOccupanyState is the data-structure of this message
type _BACnetPropertyStatesZoneOccupanyState struct {
	BACnetPropertyStatesContract
	ZoneOccupanyState BACnetAccessZoneOccupancyStateTagged
}

var _ BACnetPropertyStatesZoneOccupanyState = (*_BACnetPropertyStatesZoneOccupanyState)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesZoneOccupanyState)(nil)

// NewBACnetPropertyStatesZoneOccupanyState factory function for _BACnetPropertyStatesZoneOccupanyState
func NewBACnetPropertyStatesZoneOccupanyState(peekedTagHeader BACnetTagHeader, zoneOccupanyState BACnetAccessZoneOccupancyStateTagged) *_BACnetPropertyStatesZoneOccupanyState {
	if zoneOccupanyState == nil {
		panic("zoneOccupanyState of type BACnetAccessZoneOccupancyStateTagged for BACnetPropertyStatesZoneOccupanyState must not be nil")
	}
	_result := &_BACnetPropertyStatesZoneOccupanyState{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		ZoneOccupanyState:            zoneOccupanyState,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesZoneOccupanyState) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesZoneOccupanyState) GetZoneOccupanyState() BACnetAccessZoneOccupancyStateTagged {
	return m.ZoneOccupanyState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesZoneOccupanyState(structType any) BACnetPropertyStatesZoneOccupanyState {
	if casted, ok := structType.(BACnetPropertyStatesZoneOccupanyState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesZoneOccupanyState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesZoneOccupanyState) GetTypeName() string {
	return "BACnetPropertyStatesZoneOccupanyState"
}

func (m *_BACnetPropertyStatesZoneOccupanyState) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (zoneOccupanyState)
	lengthInBits += m.ZoneOccupanyState.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesZoneOccupanyState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesZoneOccupanyState) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesZoneOccupanyState BACnetPropertyStatesZoneOccupanyState, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesZoneOccupanyState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesZoneOccupanyState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneOccupanyState, err := ReadSimpleField[BACnetAccessZoneOccupancyStateTagged](ctx, "zoneOccupanyState", ReadComplex[BACnetAccessZoneOccupancyStateTagged](BACnetAccessZoneOccupancyStateTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneOccupanyState' field"))
	}
	m.ZoneOccupanyState = zoneOccupanyState

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesZoneOccupanyState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesZoneOccupanyState")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesZoneOccupanyState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesZoneOccupanyState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesZoneOccupanyState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesZoneOccupanyState")
		}

		if err := WriteSimpleField[BACnetAccessZoneOccupancyStateTagged](ctx, "zoneOccupanyState", m.GetZoneOccupanyState(), WriteComplex[BACnetAccessZoneOccupancyStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneOccupanyState' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesZoneOccupanyState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesZoneOccupanyState")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesZoneOccupanyState) IsBACnetPropertyStatesZoneOccupanyState() {}

func (m *_BACnetPropertyStatesZoneOccupanyState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
