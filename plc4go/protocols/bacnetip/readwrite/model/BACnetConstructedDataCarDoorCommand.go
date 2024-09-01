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

// BACnetConstructedDataCarDoorCommand is the corresponding interface of BACnetConstructedDataCarDoorCommand
type BACnetConstructedDataCarDoorCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetCarDoorCommand returns CarDoorCommand (property field)
	GetCarDoorCommand() []BACnetLiftCarDoorCommandTagged
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataCarDoorCommandExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCarDoorCommand.
// This is useful for switch cases.
type BACnetConstructedDataCarDoorCommandExactly interface {
	BACnetConstructedDataCarDoorCommand
	isBACnetConstructedDataCarDoorCommand() bool
}

// _BACnetConstructedDataCarDoorCommand is the data-structure of this message
type _BACnetConstructedDataCarDoorCommand struct {
	*_BACnetConstructedData
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	CarDoorCommand       []BACnetLiftCarDoorCommandTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCarDoorCommand) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCarDoorCommand) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CAR_DOOR_COMMAND
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCarDoorCommand) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCarDoorCommand) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCarDoorCommand) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataCarDoorCommand) GetCarDoorCommand() []BACnetLiftCarDoorCommandTagged {
	return m.CarDoorCommand
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCarDoorCommand) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCarDoorCommand factory function for _BACnetConstructedDataCarDoorCommand
func NewBACnetConstructedDataCarDoorCommand(numberOfDataElements BACnetApplicationTagUnsignedInteger, carDoorCommand []BACnetLiftCarDoorCommandTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCarDoorCommand {
	_result := &_BACnetConstructedDataCarDoorCommand{
		NumberOfDataElements:   numberOfDataElements,
		CarDoorCommand:         carDoorCommand,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCarDoorCommand(structType any) BACnetConstructedDataCarDoorCommand {
	if casted, ok := structType.(BACnetConstructedDataCarDoorCommand); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCarDoorCommand); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCarDoorCommand) GetTypeName() string {
	return "BACnetConstructedDataCarDoorCommand"
}

func (m *_BACnetConstructedDataCarDoorCommand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.CarDoorCommand) > 0 {
		for _, element := range m.CarDoorCommand {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataCarDoorCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataCarDoorCommandParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCarDoorCommand, error) {
	return BACnetConstructedDataCarDoorCommandParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataCarDoorCommandParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataCarDoorCommand, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataCarDoorCommand, error) {
		return BACnetConstructedDataCarDoorCommandParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataCarDoorCommandParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCarDoorCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCarDoorCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCarDoorCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
	}

	carDoorCommand, err := ReadTerminatedArrayField[BACnetLiftCarDoorCommandTagged](ctx, "carDoorCommand", ReadComplex[BACnetLiftCarDoorCommandTagged](BACnetLiftCarDoorCommandTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'carDoorCommand' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCarDoorCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCarDoorCommand")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCarDoorCommand{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		NumberOfDataElements: numberOfDataElements,
		CarDoorCommand:       carDoorCommand,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCarDoorCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCarDoorCommand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCarDoorCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCarDoorCommand")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", GetRef(m.GetNumberOfDataElements()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfDataElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "carDoorCommand", m.GetCarDoorCommand(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'carDoorCommand' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCarDoorCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCarDoorCommand")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCarDoorCommand) isBACnetConstructedDataCarDoorCommand() bool {
	return true
}

func (m *_BACnetConstructedDataCarDoorCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
