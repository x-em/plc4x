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

// BACnetConstructedDataBackupAndRestoreState is the corresponding interface of BACnetConstructedDataBackupAndRestoreState
type BACnetConstructedDataBackupAndRestoreState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetBackupAndRestoreState returns BackupAndRestoreState (property field)
	GetBackupAndRestoreState() BACnetBackupStateTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetBackupStateTagged
	// IsBACnetConstructedDataBackupAndRestoreState is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBackupAndRestoreState()
}

// _BACnetConstructedDataBackupAndRestoreState is the data-structure of this message
type _BACnetConstructedDataBackupAndRestoreState struct {
	BACnetConstructedDataContract
	BackupAndRestoreState BACnetBackupStateTagged
}

var _ BACnetConstructedDataBackupAndRestoreState = (*_BACnetConstructedDataBackupAndRestoreState)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBackupAndRestoreState)(nil)

// NewBACnetConstructedDataBackupAndRestoreState factory function for _BACnetConstructedDataBackupAndRestoreState
func NewBACnetConstructedDataBackupAndRestoreState(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, backupAndRestoreState BACnetBackupStateTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBackupAndRestoreState {
	if backupAndRestoreState == nil {
		panic("backupAndRestoreState of type BACnetBackupStateTagged for BACnetConstructedDataBackupAndRestoreState must not be nil")
	}
	_result := &_BACnetConstructedDataBackupAndRestoreState{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		BackupAndRestoreState:         backupAndRestoreState,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBackupAndRestoreState) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBackupAndRestoreState) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACKUP_AND_RESTORE_STATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBackupAndRestoreState) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBackupAndRestoreState) GetBackupAndRestoreState() BACnetBackupStateTagged {
	return m.BackupAndRestoreState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBackupAndRestoreState) GetActualValue() BACnetBackupStateTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetBackupStateTagged(m.GetBackupAndRestoreState())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBackupAndRestoreState(structType any) BACnetConstructedDataBackupAndRestoreState {
	if casted, ok := structType.(BACnetConstructedDataBackupAndRestoreState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBackupAndRestoreState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBackupAndRestoreState) GetTypeName() string {
	return "BACnetConstructedDataBackupAndRestoreState"
}

func (m *_BACnetConstructedDataBackupAndRestoreState) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (backupAndRestoreState)
	lengthInBits += m.BackupAndRestoreState.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBackupAndRestoreState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBackupAndRestoreState) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBackupAndRestoreState BACnetConstructedDataBackupAndRestoreState, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBackupAndRestoreState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBackupAndRestoreState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	backupAndRestoreState, err := ReadSimpleField[BACnetBackupStateTagged](ctx, "backupAndRestoreState", ReadComplex[BACnetBackupStateTagged](BACnetBackupStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'backupAndRestoreState' field"))
	}
	m.BackupAndRestoreState = backupAndRestoreState

	actualValue, err := ReadVirtualField[BACnetBackupStateTagged](ctx, "actualValue", (*BACnetBackupStateTagged)(nil), backupAndRestoreState)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBackupAndRestoreState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBackupAndRestoreState")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBackupAndRestoreState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBackupAndRestoreState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBackupAndRestoreState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBackupAndRestoreState")
		}

		if err := WriteSimpleField[BACnetBackupStateTagged](ctx, "backupAndRestoreState", m.GetBackupAndRestoreState(), WriteComplex[BACnetBackupStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'backupAndRestoreState' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBackupAndRestoreState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBackupAndRestoreState")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBackupAndRestoreState) IsBACnetConstructedDataBackupAndRestoreState() {
}

func (m *_BACnetConstructedDataBackupAndRestoreState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
