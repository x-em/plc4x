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

// BACnetPropertyStatesLiftGroupMode is the corresponding interface of BACnetPropertyStatesLiftGroupMode
type BACnetPropertyStatesLiftGroupMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetLiftGroupMode returns LiftGroupMode (property field)
	GetLiftGroupMode() BACnetLiftGroupModeTagged
	// IsBACnetPropertyStatesLiftGroupMode is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesLiftGroupMode()
}

// _BACnetPropertyStatesLiftGroupMode is the data-structure of this message
type _BACnetPropertyStatesLiftGroupMode struct {
	BACnetPropertyStatesContract
	LiftGroupMode BACnetLiftGroupModeTagged
}

var _ BACnetPropertyStatesLiftGroupMode = (*_BACnetPropertyStatesLiftGroupMode)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesLiftGroupMode)(nil)

// NewBACnetPropertyStatesLiftGroupMode factory function for _BACnetPropertyStatesLiftGroupMode
func NewBACnetPropertyStatesLiftGroupMode(peekedTagHeader BACnetTagHeader, liftGroupMode BACnetLiftGroupModeTagged) *_BACnetPropertyStatesLiftGroupMode {
	if liftGroupMode == nil {
		panic("liftGroupMode of type BACnetLiftGroupModeTagged for BACnetPropertyStatesLiftGroupMode must not be nil")
	}
	_result := &_BACnetPropertyStatesLiftGroupMode{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		LiftGroupMode:                liftGroupMode,
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

func (m *_BACnetPropertyStatesLiftGroupMode) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLiftGroupMode) GetLiftGroupMode() BACnetLiftGroupModeTagged {
	return m.LiftGroupMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLiftGroupMode(structType any) BACnetPropertyStatesLiftGroupMode {
	if casted, ok := structType.(BACnetPropertyStatesLiftGroupMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLiftGroupMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLiftGroupMode) GetTypeName() string {
	return "BACnetPropertyStatesLiftGroupMode"
}

func (m *_BACnetPropertyStatesLiftGroupMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (liftGroupMode)
	lengthInBits += m.LiftGroupMode.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesLiftGroupMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesLiftGroupMode) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesLiftGroupMode BACnetPropertyStatesLiftGroupMode, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLiftGroupMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLiftGroupMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	liftGroupMode, err := ReadSimpleField[BACnetLiftGroupModeTagged](ctx, "liftGroupMode", ReadComplex[BACnetLiftGroupModeTagged](BACnetLiftGroupModeTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'liftGroupMode' field"))
	}
	m.LiftGroupMode = liftGroupMode

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLiftGroupMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLiftGroupMode")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesLiftGroupMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLiftGroupMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLiftGroupMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLiftGroupMode")
		}

		if err := WriteSimpleField[BACnetLiftGroupModeTagged](ctx, "liftGroupMode", m.GetLiftGroupMode(), WriteComplex[BACnetLiftGroupModeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'liftGroupMode' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLiftGroupMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLiftGroupMode")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLiftGroupMode) IsBACnetPropertyStatesLiftGroupMode() {}

func (m *_BACnetPropertyStatesLiftGroupMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
