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

// BACnetActionCommand is the corresponding interface of BACnetActionCommand
type BACnetActionCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetDeviceIdentifier returns DeviceIdentifier (property field)
	GetDeviceIdentifier() BACnetContextTagObjectIdentifier
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() BACnetContextTagUnsignedInteger
	// GetPropertyValue returns PropertyValue (property field)
	GetPropertyValue() BACnetConstructedData
	// GetPriority returns Priority (property field)
	GetPriority() BACnetContextTagUnsignedInteger
	// GetPostDelay returns PostDelay (property field)
	GetPostDelay() BACnetContextTagBoolean
	// GetQuitOnFailure returns QuitOnFailure (property field)
	GetQuitOnFailure() BACnetContextTagBoolean
	// GetWriteSuccessful returns WriteSuccessful (property field)
	GetWriteSuccessful() BACnetContextTagBoolean
}

// BACnetActionCommandExactly can be used when we want exactly this type and not a type which fulfills BACnetActionCommand.
// This is useful for switch cases.
type BACnetActionCommandExactly interface {
	BACnetActionCommand
	isBACnetActionCommand() bool
}

// _BACnetActionCommand is the data-structure of this message
type _BACnetActionCommand struct {
	DeviceIdentifier   BACnetContextTagObjectIdentifier
	ObjectIdentifier   BACnetContextTagObjectIdentifier
	PropertyIdentifier BACnetPropertyIdentifierTagged
	ArrayIndex         BACnetContextTagUnsignedInteger
	PropertyValue      BACnetConstructedData
	Priority           BACnetContextTagUnsignedInteger
	PostDelay          BACnetContextTagBoolean
	QuitOnFailure      BACnetContextTagBoolean
	WriteSuccessful    BACnetContextTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetActionCommand) GetDeviceIdentifier() BACnetContextTagObjectIdentifier {
	return m.DeviceIdentifier
}

func (m *_BACnetActionCommand) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetActionCommand) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetActionCommand) GetArrayIndex() BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

func (m *_BACnetActionCommand) GetPropertyValue() BACnetConstructedData {
	return m.PropertyValue
}

func (m *_BACnetActionCommand) GetPriority() BACnetContextTagUnsignedInteger {
	return m.Priority
}

func (m *_BACnetActionCommand) GetPostDelay() BACnetContextTagBoolean {
	return m.PostDelay
}

func (m *_BACnetActionCommand) GetQuitOnFailure() BACnetContextTagBoolean {
	return m.QuitOnFailure
}

func (m *_BACnetActionCommand) GetWriteSuccessful() BACnetContextTagBoolean {
	return m.WriteSuccessful
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetActionCommand factory function for _BACnetActionCommand
func NewBACnetActionCommand(deviceIdentifier BACnetContextTagObjectIdentifier, objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged, arrayIndex BACnetContextTagUnsignedInteger, propertyValue BACnetConstructedData, priority BACnetContextTagUnsignedInteger, postDelay BACnetContextTagBoolean, quitOnFailure BACnetContextTagBoolean, writeSuccessful BACnetContextTagBoolean) *_BACnetActionCommand {
	return &_BACnetActionCommand{DeviceIdentifier: deviceIdentifier, ObjectIdentifier: objectIdentifier, PropertyIdentifier: propertyIdentifier, ArrayIndex: arrayIndex, PropertyValue: propertyValue, Priority: priority, PostDelay: postDelay, QuitOnFailure: quitOnFailure, WriteSuccessful: writeSuccessful}
}

// Deprecated: use the interface for direct cast
func CastBACnetActionCommand(structType any) BACnetActionCommand {
	if casted, ok := structType.(BACnetActionCommand); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetActionCommand); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetActionCommand) GetTypeName() string {
	return "BACnetActionCommand"
}

func (m *_BACnetActionCommand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Optional Field (deviceIdentifier)
	if m.DeviceIdentifier != nil {
		lengthInBits += m.DeviceIdentifier.GetLengthInBits(ctx)
	}

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += m.ArrayIndex.GetLengthInBits(ctx)
	}

	// Optional Field (propertyValue)
	if m.PropertyValue != nil {
		lengthInBits += m.PropertyValue.GetLengthInBits(ctx)
	}

	// Optional Field (priority)
	if m.Priority != nil {
		lengthInBits += m.Priority.GetLengthInBits(ctx)
	}

	// Optional Field (postDelay)
	if m.PostDelay != nil {
		lengthInBits += m.PostDelay.GetLengthInBits(ctx)
	}

	// Simple field (quitOnFailure)
	lengthInBits += m.QuitOnFailure.GetLengthInBits(ctx)

	// Simple field (writeSuccessful)
	lengthInBits += m.WriteSuccessful.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetActionCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetActionCommandParse(ctx context.Context, theBytes []byte) (BACnetActionCommand, error) {
	return BACnetActionCommandParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetActionCommandParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetActionCommand, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetActionCommand, error) {
		return BACnetActionCommandParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetActionCommandParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetActionCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetActionCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetActionCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	_deviceIdentifier, err := ReadOptionalField[BACnetContextTagObjectIdentifier](ctx, "deviceIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceIdentifier' field"))
	}
	var deviceIdentifier BACnetContextTagObjectIdentifier
	if _deviceIdentifier != nil {
		deviceIdentifier = *_deviceIdentifier
	}

	objectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}

	propertyIdentifier, err := ReadSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", ReadComplex[BACnetPropertyIdentifierTagged](BACnetPropertyIdentifierTaggedParseWithBufferProducer((uint8)(uint8(2)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyIdentifier' field"))
	}

	_arrayIndex, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "arrayIndex", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayIndex' field"))
	}
	var arrayIndex BACnetContextTagUnsignedInteger
	if _arrayIndex != nil {
		arrayIndex = *_arrayIndex
	}

	_propertyValue, err := ReadOptionalField[BACnetConstructedData](ctx, "propertyValue", ReadComplex[BACnetConstructedData](BACnetConstructedDataParseWithBufferProducer[BACnetConstructedData]((uint8)(uint8(4)), (BACnetObjectType)(objectIdentifier.GetObjectType()), (BACnetPropertyIdentifier)(propertyIdentifier.GetValue()), (BACnetTagPayloadUnsignedInteger)((CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool((arrayIndex) != (nil)), func() any { return CastBACnetTagPayloadUnsignedInteger((arrayIndex).GetPayload()) }, func() any { return CastBACnetTagPayloadUnsignedInteger(nil) }))))), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyValue' field"))
	}
	var propertyValue BACnetConstructedData
	if _propertyValue != nil {
		propertyValue = *_propertyValue
	}

	_priority, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "priority", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(5)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}
	var priority BACnetContextTagUnsignedInteger
	if _priority != nil {
		priority = *_priority
	}

	_postDelay, err := ReadOptionalField[BACnetContextTagBoolean](ctx, "postDelay", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(6)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'postDelay' field"))
	}
	var postDelay BACnetContextTagBoolean
	if _postDelay != nil {
		postDelay = *_postDelay
	}

	quitOnFailure, err := ReadSimpleField[BACnetContextTagBoolean](ctx, "quitOnFailure", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(7)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'quitOnFailure' field"))
	}

	writeSuccessful, err := ReadSimpleField[BACnetContextTagBoolean](ctx, "writeSuccessful", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(8)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeSuccessful' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetActionCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetActionCommand")
	}

	// Create the instance
	return &_BACnetActionCommand{
		DeviceIdentifier:   deviceIdentifier,
		ObjectIdentifier:   objectIdentifier,
		PropertyIdentifier: propertyIdentifier,
		ArrayIndex:         arrayIndex,
		PropertyValue:      propertyValue,
		Priority:           priority,
		PostDelay:          postDelay,
		QuitOnFailure:      quitOnFailure,
		WriteSuccessful:    writeSuccessful,
	}, nil
}

func (m *_BACnetActionCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetActionCommand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetActionCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetActionCommand")
	}

	if err := WriteOptionalField[BACnetContextTagObjectIdentifier](ctx, "deviceIdentifier", GetRef(m.GetDeviceIdentifier()), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'deviceIdentifier' field")
	}

	// Simple Field (objectIdentifier)
	if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for objectIdentifier")
	}
	_objectIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetObjectIdentifier())
	if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for objectIdentifier")
	}
	if _objectIdentifierErr != nil {
		return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
	}

	// Simple Field (propertyIdentifier)
	if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for propertyIdentifier")
	}
	_propertyIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetPropertyIdentifier())
	if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for propertyIdentifier")
	}
	if _propertyIdentifierErr != nil {
		return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
	}

	if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "arrayIndex", GetRef(m.GetArrayIndex()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'arrayIndex' field")
	}

	if err := WriteOptionalField[BACnetConstructedData](ctx, "propertyValue", GetRef(m.GetPropertyValue()), WriteComplex[BACnetConstructedData](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'propertyValue' field")
	}

	if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "priority", GetRef(m.GetPriority()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'priority' field")
	}

	if err := WriteOptionalField[BACnetContextTagBoolean](ctx, "postDelay", GetRef(m.GetPostDelay()), WriteComplex[BACnetContextTagBoolean](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'postDelay' field")
	}

	// Simple Field (quitOnFailure)
	if pushErr := writeBuffer.PushContext("quitOnFailure"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for quitOnFailure")
	}
	_quitOnFailureErr := writeBuffer.WriteSerializable(ctx, m.GetQuitOnFailure())
	if popErr := writeBuffer.PopContext("quitOnFailure"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for quitOnFailure")
	}
	if _quitOnFailureErr != nil {
		return errors.Wrap(_quitOnFailureErr, "Error serializing 'quitOnFailure' field")
	}

	// Simple Field (writeSuccessful)
	if pushErr := writeBuffer.PushContext("writeSuccessful"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for writeSuccessful")
	}
	_writeSuccessfulErr := writeBuffer.WriteSerializable(ctx, m.GetWriteSuccessful())
	if popErr := writeBuffer.PopContext("writeSuccessful"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for writeSuccessful")
	}
	if _writeSuccessfulErr != nil {
		return errors.Wrap(_writeSuccessfulErr, "Error serializing 'writeSuccessful' field")
	}

	if popErr := writeBuffer.PopContext("BACnetActionCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetActionCommand")
	}
	return nil
}

func (m *_BACnetActionCommand) isBACnetActionCommand() bool {
	return true
}

func (m *_BACnetActionCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
