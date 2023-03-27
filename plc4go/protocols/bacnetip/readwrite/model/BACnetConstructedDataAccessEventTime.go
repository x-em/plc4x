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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataAccessEventTime is the corresponding interface of BACnetConstructedDataAccessEventTime
type BACnetConstructedDataAccessEventTime interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAccessEventTime returns AccessEventTime (property field)
	GetAccessEventTime() BACnetTimeStamp
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetTimeStamp
}

// BACnetConstructedDataAccessEventTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAccessEventTime.
// This is useful for switch cases.
type BACnetConstructedDataAccessEventTimeExactly interface {
	BACnetConstructedDataAccessEventTime
	isBACnetConstructedDataAccessEventTime() bool
}

// _BACnetConstructedDataAccessEventTime is the data-structure of this message
type _BACnetConstructedDataAccessEventTime struct {
	*_BACnetConstructedData
	AccessEventTime BACnetTimeStamp
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessEventTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAccessEventTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCESS_EVENT_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessEventTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAccessEventTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessEventTime) GetAccessEventTime() BACnetTimeStamp {
	return m.AccessEventTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccessEventTime) GetActualValue() BACnetTimeStamp {
	ctx := context.Background()
	_ = ctx
	return CastBACnetTimeStamp(m.GetAccessEventTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccessEventTime factory function for _BACnetConstructedDataAccessEventTime
func NewBACnetConstructedDataAccessEventTime(accessEventTime BACnetTimeStamp, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessEventTime {
	_result := &_BACnetConstructedDataAccessEventTime{
		AccessEventTime:        accessEventTime,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessEventTime(structType interface{}) BACnetConstructedDataAccessEventTime {
	if casted, ok := structType.(BACnetConstructedDataAccessEventTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessEventTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessEventTime) GetTypeName() string {
	return "BACnetConstructedDataAccessEventTime"
}

func (m *_BACnetConstructedDataAccessEventTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (accessEventTime)
	lengthInBits += m.AccessEventTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessEventTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataAccessEventTimeParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccessEventTime, error) {
	return BACnetConstructedDataAccessEventTimeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAccessEventTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccessEventTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessEventTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessEventTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (accessEventTime)
	if pullErr := readBuffer.PullContext("accessEventTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accessEventTime")
	}
	_accessEventTime, _accessEventTimeErr := BACnetTimeStampParseWithBuffer(ctx, readBuffer)
	if _accessEventTimeErr != nil {
		return nil, errors.Wrap(_accessEventTimeErr, "Error parsing 'accessEventTime' field of BACnetConstructedDataAccessEventTime")
	}
	accessEventTime := _accessEventTime.(BACnetTimeStamp)
	if closeErr := readBuffer.CloseContext("accessEventTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accessEventTime")
	}

	// Virtual field
	_actualValue := accessEventTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessEventTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessEventTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAccessEventTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		AccessEventTime: accessEventTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAccessEventTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessEventTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessEventTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessEventTime")
		}

		// Simple Field (accessEventTime)
		if pushErr := writeBuffer.PushContext("accessEventTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for accessEventTime")
		}
		_accessEventTimeErr := writeBuffer.WriteSerializable(ctx, m.GetAccessEventTime())
		if popErr := writeBuffer.PopContext("accessEventTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for accessEventTime")
		}
		if _accessEventTimeErr != nil {
			return errors.Wrap(_accessEventTimeErr, "Error serializing 'accessEventTime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessEventTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessEventTime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessEventTime) isBACnetConstructedDataAccessEventTime() bool {
	return true
}

func (m *_BACnetConstructedDataAccessEventTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
