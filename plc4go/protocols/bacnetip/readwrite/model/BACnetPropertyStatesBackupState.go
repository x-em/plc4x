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

// BACnetPropertyStatesBackupState is the corresponding interface of BACnetPropertyStatesBackupState
type BACnetPropertyStatesBackupState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetBackupState returns BackupState (property field)
	GetBackupState() BACnetBackupStateTagged
	// IsBACnetPropertyStatesBackupState is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesBackupState()
}

// _BACnetPropertyStatesBackupState is the data-structure of this message
type _BACnetPropertyStatesBackupState struct {
	BACnetPropertyStatesContract
	BackupState BACnetBackupStateTagged
}

var _ BACnetPropertyStatesBackupState = (*_BACnetPropertyStatesBackupState)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesBackupState)(nil)

// NewBACnetPropertyStatesBackupState factory function for _BACnetPropertyStatesBackupState
func NewBACnetPropertyStatesBackupState(peekedTagHeader BACnetTagHeader, backupState BACnetBackupStateTagged) *_BACnetPropertyStatesBackupState {
	if backupState == nil {
		panic("backupState of type BACnetBackupStateTagged for BACnetPropertyStatesBackupState must not be nil")
	}
	_result := &_BACnetPropertyStatesBackupState{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		BackupState:                  backupState,
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

func (m *_BACnetPropertyStatesBackupState) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesBackupState) GetBackupState() BACnetBackupStateTagged {
	return m.BackupState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesBackupState(structType any) BACnetPropertyStatesBackupState {
	if casted, ok := structType.(BACnetPropertyStatesBackupState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesBackupState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesBackupState) GetTypeName() string {
	return "BACnetPropertyStatesBackupState"
}

func (m *_BACnetPropertyStatesBackupState) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (backupState)
	lengthInBits += m.BackupState.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesBackupState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesBackupState) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesBackupState BACnetPropertyStatesBackupState, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesBackupState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesBackupState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	backupState, err := ReadSimpleField[BACnetBackupStateTagged](ctx, "backupState", ReadComplex[BACnetBackupStateTagged](BACnetBackupStateTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'backupState' field"))
	}
	m.BackupState = backupState

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesBackupState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesBackupState")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesBackupState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesBackupState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesBackupState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesBackupState")
		}

		if err := WriteSimpleField[BACnetBackupStateTagged](ctx, "backupState", m.GetBackupState(), WriteComplex[BACnetBackupStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'backupState' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesBackupState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesBackupState")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesBackupState) IsBACnetPropertyStatesBackupState() {}

func (m *_BACnetPropertyStatesBackupState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
