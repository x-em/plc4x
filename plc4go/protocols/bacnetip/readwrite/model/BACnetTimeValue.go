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

// BACnetTimeValue is the corresponding interface of BACnetTimeValue
type BACnetTimeValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetTimeValue returns TimeValue (property field)
	GetTimeValue() BACnetApplicationTagTime
	// GetValue returns Value (property field)
	GetValue() BACnetConstructedDataElement
	// IsBACnetTimeValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimeValue()
}

// _BACnetTimeValue is the data-structure of this message
type _BACnetTimeValue struct {
	TimeValue BACnetApplicationTagTime
	Value     BACnetConstructedDataElement
}

var _ BACnetTimeValue = (*_BACnetTimeValue)(nil)

// NewBACnetTimeValue factory function for _BACnetTimeValue
func NewBACnetTimeValue(timeValue BACnetApplicationTagTime, value BACnetConstructedDataElement) *_BACnetTimeValue {
	if timeValue == nil {
		panic("timeValue of type BACnetApplicationTagTime for BACnetTimeValue must not be nil")
	}
	if value == nil {
		panic("value of type BACnetConstructedDataElement for BACnetTimeValue must not be nil")
	}
	return &_BACnetTimeValue{TimeValue: timeValue, Value: value}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimeValue) GetTimeValue() BACnetApplicationTagTime {
	return m.TimeValue
}

func (m *_BACnetTimeValue) GetValue() BACnetConstructedDataElement {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTimeValue(structType any) BACnetTimeValue {
	if casted, ok := structType.(BACnetTimeValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimeValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimeValue) GetTypeName() string {
	return "BACnetTimeValue"
}

func (m *_BACnetTimeValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timeValue)
	lengthInBits += m.TimeValue.GetLengthInBits(ctx)

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimeValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTimeValueParse(ctx context.Context, theBytes []byte) (BACnetTimeValue, error) {
	return BACnetTimeValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetTimeValueParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimeValue, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimeValue, error) {
		return BACnetTimeValueParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetTimeValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimeValue, error) {
	v, err := (&_BACnetTimeValue{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetTimeValue) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetTimeValue BACnetTimeValue, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimeValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimeValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeValue, err := ReadSimpleField[BACnetApplicationTagTime](ctx, "timeValue", ReadComplex[BACnetApplicationTagTime](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagTime](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeValue' field"))
	}
	m.TimeValue = timeValue

	value, err := ReadSimpleField[BACnetConstructedDataElement](ctx, "value", ReadComplex[BACnetConstructedDataElement](BACnetConstructedDataElementParseWithBufferProducer((BACnetObjectType)(BACnetObjectType_VENDOR_PROPRIETARY_VALUE), (BACnetPropertyIdentifier)(BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE), (BACnetTagPayloadUnsignedInteger)(nil)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetTimeValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimeValue")
	}

	return m, nil
}

func (m *_BACnetTimeValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimeValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTimeValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTimeValue")
	}

	if err := WriteSimpleField[BACnetApplicationTagTime](ctx, "timeValue", m.GetTimeValue(), WriteComplex[BACnetApplicationTagTime](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'timeValue' field")
	}

	if err := WriteSimpleField[BACnetConstructedDataElement](ctx, "value", m.GetValue(), WriteComplex[BACnetConstructedDataElement](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTimeValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTimeValue")
	}
	return nil
}

func (m *_BACnetTimeValue) IsBACnetTimeValue() {}

func (m *_BACnetTimeValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
