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

// BACnetConstructedDataLastAccessEvent is the corresponding interface of BACnetConstructedDataLastAccessEvent
type BACnetConstructedDataLastAccessEvent interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLastAccessEvent returns LastAccessEvent (property field)
	GetLastAccessEvent() BACnetAccessEventTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAccessEventTagged
	// IsBACnetConstructedDataLastAccessEvent is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLastAccessEvent()
}

// _BACnetConstructedDataLastAccessEvent is the data-structure of this message
type _BACnetConstructedDataLastAccessEvent struct {
	BACnetConstructedDataContract
	LastAccessEvent BACnetAccessEventTagged
}

var _ BACnetConstructedDataLastAccessEvent = (*_BACnetConstructedDataLastAccessEvent)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLastAccessEvent)(nil)

// NewBACnetConstructedDataLastAccessEvent factory function for _BACnetConstructedDataLastAccessEvent
func NewBACnetConstructedDataLastAccessEvent(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lastAccessEvent BACnetAccessEventTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastAccessEvent {
	if lastAccessEvent == nil {
		panic("lastAccessEvent of type BACnetAccessEventTagged for BACnetConstructedDataLastAccessEvent must not be nil")
	}
	_result := &_BACnetConstructedDataLastAccessEvent{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LastAccessEvent:               lastAccessEvent,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastAccessEvent) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastAccessEvent) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_ACCESS_EVENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastAccessEvent) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastAccessEvent) GetLastAccessEvent() BACnetAccessEventTagged {
	return m.LastAccessEvent
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastAccessEvent) GetActualValue() BACnetAccessEventTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetAccessEventTagged(m.GetLastAccessEvent())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastAccessEvent(structType any) BACnetConstructedDataLastAccessEvent {
	if casted, ok := structType.(BACnetConstructedDataLastAccessEvent); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastAccessEvent); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastAccessEvent) GetTypeName() string {
	return "BACnetConstructedDataLastAccessEvent"
}

func (m *_BACnetConstructedDataLastAccessEvent) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lastAccessEvent)
	lengthInBits += m.LastAccessEvent.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastAccessEvent) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLastAccessEvent) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLastAccessEvent BACnetConstructedDataLastAccessEvent, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastAccessEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastAccessEvent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lastAccessEvent, err := ReadSimpleField[BACnetAccessEventTagged](ctx, "lastAccessEvent", ReadComplex[BACnetAccessEventTagged](BACnetAccessEventTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastAccessEvent' field"))
	}
	m.LastAccessEvent = lastAccessEvent

	actualValue, err := ReadVirtualField[BACnetAccessEventTagged](ctx, "actualValue", (*BACnetAccessEventTagged)(nil), lastAccessEvent)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastAccessEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastAccessEvent")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLastAccessEvent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastAccessEvent) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastAccessEvent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastAccessEvent")
		}

		if err := WriteSimpleField[BACnetAccessEventTagged](ctx, "lastAccessEvent", m.GetLastAccessEvent(), WriteComplex[BACnetAccessEventTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastAccessEvent' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastAccessEvent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastAccessEvent")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastAccessEvent) IsBACnetConstructedDataLastAccessEvent() {}

func (m *_BACnetConstructedDataLastAccessEvent) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
