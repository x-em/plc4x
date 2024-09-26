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

// BACnetConstructedDataMachineRoomID is the corresponding interface of BACnetConstructedDataMachineRoomID
type BACnetConstructedDataMachineRoomID interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMachineRoomId returns MachineRoomId (property field)
	GetMachineRoomId() BACnetApplicationTagObjectIdentifier
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagObjectIdentifier
	// IsBACnetConstructedDataMachineRoomID is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMachineRoomID()
}

// _BACnetConstructedDataMachineRoomID is the data-structure of this message
type _BACnetConstructedDataMachineRoomID struct {
	BACnetConstructedDataContract
	MachineRoomId BACnetApplicationTagObjectIdentifier
}

var _ BACnetConstructedDataMachineRoomID = (*_BACnetConstructedDataMachineRoomID)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMachineRoomID)(nil)

// NewBACnetConstructedDataMachineRoomID factory function for _BACnetConstructedDataMachineRoomID
func NewBACnetConstructedDataMachineRoomID(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, machineRoomId BACnetApplicationTagObjectIdentifier, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMachineRoomID {
	if machineRoomId == nil {
		panic("machineRoomId of type BACnetApplicationTagObjectIdentifier for BACnetConstructedDataMachineRoomID must not be nil")
	}
	_result := &_BACnetConstructedDataMachineRoomID{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MachineRoomId:                 machineRoomId,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMachineRoomID) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMachineRoomID) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MACHINE_ROOM_ID
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMachineRoomID) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMachineRoomID) GetMachineRoomId() BACnetApplicationTagObjectIdentifier {
	return m.MachineRoomId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMachineRoomID) GetActualValue() BACnetApplicationTagObjectIdentifier {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagObjectIdentifier(m.GetMachineRoomId())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMachineRoomID(structType any) BACnetConstructedDataMachineRoomID {
	if casted, ok := structType.(BACnetConstructedDataMachineRoomID); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMachineRoomID); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMachineRoomID) GetTypeName() string {
	return "BACnetConstructedDataMachineRoomID"
}

func (m *_BACnetConstructedDataMachineRoomID) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (machineRoomId)
	lengthInBits += m.MachineRoomId.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMachineRoomID) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMachineRoomID) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMachineRoomID BACnetConstructedDataMachineRoomID, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMachineRoomID"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMachineRoomID")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	machineRoomId, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "machineRoomId", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'machineRoomId' field"))
	}
	m.MachineRoomId = machineRoomId

	actualValue, err := ReadVirtualField[BACnetApplicationTagObjectIdentifier](ctx, "actualValue", (*BACnetApplicationTagObjectIdentifier)(nil), machineRoomId)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMachineRoomID"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMachineRoomID")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMachineRoomID) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMachineRoomID) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMachineRoomID"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMachineRoomID")
		}

		if err := WriteSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "machineRoomId", m.GetMachineRoomId(), WriteComplex[BACnetApplicationTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'machineRoomId' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMachineRoomID"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMachineRoomID")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMachineRoomID) IsBACnetConstructedDataMachineRoomID() {}

func (m *_BACnetConstructedDataMachineRoomID) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
