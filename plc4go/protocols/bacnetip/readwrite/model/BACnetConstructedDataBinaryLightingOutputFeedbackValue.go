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

// BACnetConstructedDataBinaryLightingOutputFeedbackValue is the corresponding interface of BACnetConstructedDataBinaryLightingOutputFeedbackValue
type BACnetConstructedDataBinaryLightingOutputFeedbackValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFeedbackValue returns FeedbackValue (property field)
	GetFeedbackValue() BACnetBinaryLightingPVTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetBinaryLightingPVTagged
	// IsBACnetConstructedDataBinaryLightingOutputFeedbackValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBinaryLightingOutputFeedbackValue()
}

// _BACnetConstructedDataBinaryLightingOutputFeedbackValue is the data-structure of this message
type _BACnetConstructedDataBinaryLightingOutputFeedbackValue struct {
	BACnetConstructedDataContract
	FeedbackValue BACnetBinaryLightingPVTagged
}

var _ BACnetConstructedDataBinaryLightingOutputFeedbackValue = (*_BACnetConstructedDataBinaryLightingOutputFeedbackValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBinaryLightingOutputFeedbackValue)(nil)

// NewBACnetConstructedDataBinaryLightingOutputFeedbackValue factory function for _BACnetConstructedDataBinaryLightingOutputFeedbackValue
func NewBACnetConstructedDataBinaryLightingOutputFeedbackValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, feedbackValue BACnetBinaryLightingPVTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBinaryLightingOutputFeedbackValue {
	if feedbackValue == nil {
		panic("feedbackValue of type BACnetBinaryLightingPVTagged for BACnetConstructedDataBinaryLightingOutputFeedbackValue must not be nil")
	}
	_result := &_BACnetConstructedDataBinaryLightingOutputFeedbackValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FeedbackValue:                 feedbackValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_BINARY_LIGHTING_OUTPUT
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FEEDBACK_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) GetFeedbackValue() BACnetBinaryLightingPVTagged {
	return m.FeedbackValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) GetActualValue() BACnetBinaryLightingPVTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetBinaryLightingPVTagged(m.GetFeedbackValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBinaryLightingOutputFeedbackValue(structType any) BACnetConstructedDataBinaryLightingOutputFeedbackValue {
	if casted, ok := structType.(BACnetConstructedDataBinaryLightingOutputFeedbackValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBinaryLightingOutputFeedbackValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) GetTypeName() string {
	return "BACnetConstructedDataBinaryLightingOutputFeedbackValue"
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (feedbackValue)
	lengthInBits += m.FeedbackValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBinaryLightingOutputFeedbackValue BACnetConstructedDataBinaryLightingOutputFeedbackValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBinaryLightingOutputFeedbackValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBinaryLightingOutputFeedbackValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	feedbackValue, err := ReadSimpleField[BACnetBinaryLightingPVTagged](ctx, "feedbackValue", ReadComplex[BACnetBinaryLightingPVTagged](BACnetBinaryLightingPVTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'feedbackValue' field"))
	}
	m.FeedbackValue = feedbackValue

	actualValue, err := ReadVirtualField[BACnetBinaryLightingPVTagged](ctx, "actualValue", (*BACnetBinaryLightingPVTagged)(nil), feedbackValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBinaryLightingOutputFeedbackValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBinaryLightingOutputFeedbackValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBinaryLightingOutputFeedbackValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBinaryLightingOutputFeedbackValue")
		}

		if err := WriteSimpleField[BACnetBinaryLightingPVTagged](ctx, "feedbackValue", m.GetFeedbackValue(), WriteComplex[BACnetBinaryLightingPVTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'feedbackValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBinaryLightingOutputFeedbackValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBinaryLightingOutputFeedbackValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) IsBACnetConstructedDataBinaryLightingOutputFeedbackValue() {
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
