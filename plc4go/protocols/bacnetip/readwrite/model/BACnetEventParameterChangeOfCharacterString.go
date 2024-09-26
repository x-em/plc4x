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

// BACnetEventParameterChangeOfCharacterString is the corresponding interface of BACnetEventParameterChangeOfCharacterString
type BACnetEventParameterChangeOfCharacterString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetListOfAlarmValues returns ListOfAlarmValues (property field)
	GetListOfAlarmValues() BACnetEventParameterChangeOfCharacterStringListOfAlarmValues
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventParameterChangeOfCharacterString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterChangeOfCharacterString()
}

// _BACnetEventParameterChangeOfCharacterString is the data-structure of this message
type _BACnetEventParameterChangeOfCharacterString struct {
	BACnetEventParameterContract
	OpeningTag        BACnetOpeningTag
	TimeDelay         BACnetContextTagUnsignedInteger
	ListOfAlarmValues BACnetEventParameterChangeOfCharacterStringListOfAlarmValues
	ClosingTag        BACnetClosingTag
}

var _ BACnetEventParameterChangeOfCharacterString = (*_BACnetEventParameterChangeOfCharacterString)(nil)
var _ BACnetEventParameterRequirements = (*_BACnetEventParameterChangeOfCharacterString)(nil)

// NewBACnetEventParameterChangeOfCharacterString factory function for _BACnetEventParameterChangeOfCharacterString
func NewBACnetEventParameterChangeOfCharacterString(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, listOfAlarmValues BACnetEventParameterChangeOfCharacterStringListOfAlarmValues, closingTag BACnetClosingTag) *_BACnetEventParameterChangeOfCharacterString {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventParameterChangeOfCharacterString must not be nil")
	}
	if timeDelay == nil {
		panic("timeDelay of type BACnetContextTagUnsignedInteger for BACnetEventParameterChangeOfCharacterString must not be nil")
	}
	if listOfAlarmValues == nil {
		panic("listOfAlarmValues of type BACnetEventParameterChangeOfCharacterStringListOfAlarmValues for BACnetEventParameterChangeOfCharacterString must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventParameterChangeOfCharacterString must not be nil")
	}
	_result := &_BACnetEventParameterChangeOfCharacterString{
		BACnetEventParameterContract: NewBACnetEventParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		TimeDelay:                    timeDelay,
		ListOfAlarmValues:            listOfAlarmValues,
		ClosingTag:                   closingTag,
	}
	_result.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = _result
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

func (m *_BACnetEventParameterChangeOfCharacterString) GetParent() BACnetEventParameterContract {
	return m.BACnetEventParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfCharacterString) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfCharacterString) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterChangeOfCharacterString) GetListOfAlarmValues() BACnetEventParameterChangeOfCharacterStringListOfAlarmValues {
	return m.ListOfAlarmValues
}

func (m *_BACnetEventParameterChangeOfCharacterString) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfCharacterString(structType any) BACnetEventParameterChangeOfCharacterString {
	if casted, ok := structType.(BACnetEventParameterChangeOfCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfCharacterString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfCharacterString) GetTypeName() string {
	return "BACnetEventParameterChangeOfCharacterString"
}

func (m *_BACnetEventParameterChangeOfCharacterString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventParameterContract.(*_BACnetEventParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits(ctx)

	// Simple field (listOfAlarmValues)
	lengthInBits += m.ListOfAlarmValues.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfCharacterString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventParameterChangeOfCharacterString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventParameter) (__bACnetEventParameterChangeOfCharacterString BACnetEventParameterChangeOfCharacterString, err error) {
	m.BACnetEventParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfCharacterString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfCharacterString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(17))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	timeDelay, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeDelay' field"))
	}
	m.TimeDelay = timeDelay

	listOfAlarmValues, err := ReadSimpleField[BACnetEventParameterChangeOfCharacterStringListOfAlarmValues](ctx, "listOfAlarmValues", ReadComplex[BACnetEventParameterChangeOfCharacterStringListOfAlarmValues](BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfAlarmValues' field"))
	}
	m.ListOfAlarmValues = listOfAlarmValues

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(17))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfCharacterString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfCharacterString")
	}

	return m, nil
}

func (m *_BACnetEventParameterChangeOfCharacterString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfCharacterString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfCharacterString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfCharacterString")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", m.GetTimeDelay(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeDelay' field")
		}

		if err := WriteSimpleField[BACnetEventParameterChangeOfCharacterStringListOfAlarmValues](ctx, "listOfAlarmValues", m.GetListOfAlarmValues(), WriteComplex[BACnetEventParameterChangeOfCharacterStringListOfAlarmValues](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfAlarmValues' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfCharacterString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfCharacterString")
		}
		return nil
	}
	return m.BACnetEventParameterContract.(*_BACnetEventParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterChangeOfCharacterString) IsBACnetEventParameterChangeOfCharacterString() {
}

func (m *_BACnetEventParameterChangeOfCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
