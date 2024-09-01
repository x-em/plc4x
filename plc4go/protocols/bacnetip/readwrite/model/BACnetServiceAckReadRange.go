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

// BACnetServiceAckReadRange is the corresponding interface of BACnetServiceAckReadRange
type BACnetServiceAckReadRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetPropertyArrayIndex returns PropertyArrayIndex (property field)
	GetPropertyArrayIndex() BACnetContextTagUnsignedInteger
	// GetResultFlags returns ResultFlags (property field)
	GetResultFlags() BACnetResultFlagsTagged
	// GetItemCount returns ItemCount (property field)
	GetItemCount() BACnetContextTagUnsignedInteger
	// GetItemData returns ItemData (property field)
	GetItemData() BACnetConstructedData
	// GetFirstSequenceNumber returns FirstSequenceNumber (property field)
	GetFirstSequenceNumber() BACnetContextTagUnsignedInteger
}

// BACnetServiceAckReadRangeExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckReadRange.
// This is useful for switch cases.
type BACnetServiceAckReadRangeExactly interface {
	BACnetServiceAckReadRange
	isBACnetServiceAckReadRange() bool
}

// _BACnetServiceAckReadRange is the data-structure of this message
type _BACnetServiceAckReadRange struct {
	*_BACnetServiceAck
	ObjectIdentifier    BACnetContextTagObjectIdentifier
	PropertyIdentifier  BACnetPropertyIdentifierTagged
	PropertyArrayIndex  BACnetContextTagUnsignedInteger
	ResultFlags         BACnetResultFlagsTagged
	ItemCount           BACnetContextTagUnsignedInteger
	ItemData            BACnetConstructedData
	FirstSequenceNumber BACnetContextTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckReadRange) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_READ_RANGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckReadRange) InitializeParent(parent BACnetServiceAck) {}

func (m *_BACnetServiceAckReadRange) GetParent() BACnetServiceAck {
	return m._BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckReadRange) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetServiceAckReadRange) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetServiceAckReadRange) GetPropertyArrayIndex() BACnetContextTagUnsignedInteger {
	return m.PropertyArrayIndex
}

func (m *_BACnetServiceAckReadRange) GetResultFlags() BACnetResultFlagsTagged {
	return m.ResultFlags
}

func (m *_BACnetServiceAckReadRange) GetItemCount() BACnetContextTagUnsignedInteger {
	return m.ItemCount
}

func (m *_BACnetServiceAckReadRange) GetItemData() BACnetConstructedData {
	return m.ItemData
}

func (m *_BACnetServiceAckReadRange) GetFirstSequenceNumber() BACnetContextTagUnsignedInteger {
	return m.FirstSequenceNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckReadRange factory function for _BACnetServiceAckReadRange
func NewBACnetServiceAckReadRange(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged, propertyArrayIndex BACnetContextTagUnsignedInteger, resultFlags BACnetResultFlagsTagged, itemCount BACnetContextTagUnsignedInteger, itemData BACnetConstructedData, firstSequenceNumber BACnetContextTagUnsignedInteger, serviceAckLength uint32) *_BACnetServiceAckReadRange {
	_result := &_BACnetServiceAckReadRange{
		ObjectIdentifier:    objectIdentifier,
		PropertyIdentifier:  propertyIdentifier,
		PropertyArrayIndex:  propertyArrayIndex,
		ResultFlags:         resultFlags,
		ItemCount:           itemCount,
		ItemData:            itemData,
		FirstSequenceNumber: firstSequenceNumber,
		_BACnetServiceAck:   NewBACnetServiceAck(serviceAckLength),
	}
	_result._BACnetServiceAck._BACnetServiceAckChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckReadRange(structType any) BACnetServiceAckReadRange {
	if casted, ok := structType.(BACnetServiceAckReadRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckReadRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckReadRange) GetTypeName() string {
	return "BACnetServiceAckReadRange"
}

func (m *_BACnetServiceAckReadRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (propertyArrayIndex)
	if m.PropertyArrayIndex != nil {
		lengthInBits += m.PropertyArrayIndex.GetLengthInBits(ctx)
	}

	// Simple field (resultFlags)
	lengthInBits += m.ResultFlags.GetLengthInBits(ctx)

	// Simple field (itemCount)
	lengthInBits += m.ItemCount.GetLengthInBits(ctx)

	// Optional Field (itemData)
	if m.ItemData != nil {
		lengthInBits += m.ItemData.GetLengthInBits(ctx)
	}

	// Optional Field (firstSequenceNumber)
	if m.FirstSequenceNumber != nil {
		lengthInBits += m.FirstSequenceNumber.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetServiceAckReadRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetServiceAckReadRangeParse(ctx context.Context, theBytes []byte, serviceAckLength uint32) (BACnetServiceAckReadRange, error) {
	return BACnetServiceAckReadRangeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceAckLength)
}

func BACnetServiceAckReadRangeParseWithBufferProducer(serviceAckLength uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetServiceAckReadRange, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetServiceAckReadRange, error) {
		return BACnetServiceAckReadRangeParseWithBuffer(ctx, readBuffer, serviceAckLength)
	}
}

func BACnetServiceAckReadRangeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceAckLength uint32) (BACnetServiceAckReadRange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckReadRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckReadRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}

	propertyIdentifier, err := ReadSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", ReadComplex[BACnetPropertyIdentifierTagged](BACnetPropertyIdentifierTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyIdentifier' field"))
	}

	_propertyArrayIndex, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "propertyArrayIndex", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyArrayIndex' field"))
	}
	var propertyArrayIndex BACnetContextTagUnsignedInteger
	if _propertyArrayIndex != nil {
		propertyArrayIndex = *_propertyArrayIndex
	}

	resultFlags, err := ReadSimpleField[BACnetResultFlagsTagged](ctx, "resultFlags", ReadComplex[BACnetResultFlagsTagged](BACnetResultFlagsTaggedParseWithBufferProducer((uint8)(uint8(3)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'resultFlags' field"))
	}

	itemCount, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "itemCount", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(4)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemCount' field"))
	}

	_itemData, err := ReadOptionalField[BACnetConstructedData](ctx, "itemData", ReadComplex[BACnetConstructedData](BACnetConstructedDataParseWithBufferProducer[BACnetConstructedData]((uint8)(uint8(5)), (BACnetObjectType)(objectIdentifier.GetObjectType()), (BACnetPropertyIdentifier)(propertyIdentifier.GetValue()), (BACnetTagPayloadUnsignedInteger)((CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool((propertyArrayIndex) != (nil)), func() any { return CastBACnetTagPayloadUnsignedInteger((propertyArrayIndex).GetPayload()) }, func() any { return CastBACnetTagPayloadUnsignedInteger(nil) }))))), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemData' field"))
	}
	var itemData BACnetConstructedData
	if _itemData != nil {
		itemData = *_itemData
	}

	_firstSequenceNumber, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "firstSequenceNumber", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(6)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'firstSequenceNumber' field"))
	}
	var firstSequenceNumber BACnetContextTagUnsignedInteger
	if _firstSequenceNumber != nil {
		firstSequenceNumber = *_firstSequenceNumber
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckReadRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckReadRange")
	}

	// Create a partially initialized instance
	_child := &_BACnetServiceAckReadRange{
		_BACnetServiceAck: &_BACnetServiceAck{
			ServiceAckLength: serviceAckLength,
		},
		ObjectIdentifier:    objectIdentifier,
		PropertyIdentifier:  propertyIdentifier,
		PropertyArrayIndex:  propertyArrayIndex,
		ResultFlags:         resultFlags,
		ItemCount:           itemCount,
		ItemData:            itemData,
		FirstSequenceNumber: firstSequenceNumber,
	}
	_child._BACnetServiceAck._BACnetServiceAckChildRequirements = _child
	return _child, nil
}

func (m *_BACnetServiceAckReadRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckReadRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckReadRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckReadRange")
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

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "propertyArrayIndex", GetRef(m.GetPropertyArrayIndex()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyArrayIndex' field")
		}

		// Simple Field (resultFlags)
		if pushErr := writeBuffer.PushContext("resultFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for resultFlags")
		}
		_resultFlagsErr := writeBuffer.WriteSerializable(ctx, m.GetResultFlags())
		if popErr := writeBuffer.PopContext("resultFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for resultFlags")
		}
		if _resultFlagsErr != nil {
			return errors.Wrap(_resultFlagsErr, "Error serializing 'resultFlags' field")
		}

		// Simple Field (itemCount)
		if pushErr := writeBuffer.PushContext("itemCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for itemCount")
		}
		_itemCountErr := writeBuffer.WriteSerializable(ctx, m.GetItemCount())
		if popErr := writeBuffer.PopContext("itemCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for itemCount")
		}
		if _itemCountErr != nil {
			return errors.Wrap(_itemCountErr, "Error serializing 'itemCount' field")
		}

		if err := WriteOptionalField[BACnetConstructedData](ctx, "itemData", GetRef(m.GetItemData()), WriteComplex[BACnetConstructedData](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'itemData' field")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "firstSequenceNumber", GetRef(m.GetFirstSequenceNumber()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'firstSequenceNumber' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckReadRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckReadRange")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetServiceAckReadRange) isBACnetServiceAckReadRange() bool {
	return true
}

func (m *_BACnetServiceAckReadRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
