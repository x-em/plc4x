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

// BACnetFaultParameterFaultLifeSafety is the corresponding interface of BACnetFaultParameterFaultLifeSafety
type BACnetFaultParameterFaultLifeSafety interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfFaultValues returns ListOfFaultValues (property field)
	GetListOfFaultValues() BACnetFaultParameterFaultLifeSafetyListOfFaultValues
	// GetModePropertyReference returns ModePropertyReference (property field)
	GetModePropertyReference() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetFaultParameterFaultLifeSafety is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultLifeSafety()
}

// _BACnetFaultParameterFaultLifeSafety is the data-structure of this message
type _BACnetFaultParameterFaultLifeSafety struct {
	BACnetFaultParameterContract
	OpeningTag            BACnetOpeningTag
	ListOfFaultValues     BACnetFaultParameterFaultLifeSafetyListOfFaultValues
	ModePropertyReference BACnetDeviceObjectPropertyReferenceEnclosed
	ClosingTag            BACnetClosingTag
}

var _ BACnetFaultParameterFaultLifeSafety = (*_BACnetFaultParameterFaultLifeSafety)(nil)
var _ BACnetFaultParameterRequirements = (*_BACnetFaultParameterFaultLifeSafety)(nil)

// NewBACnetFaultParameterFaultLifeSafety factory function for _BACnetFaultParameterFaultLifeSafety
func NewBACnetFaultParameterFaultLifeSafety(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, listOfFaultValues BACnetFaultParameterFaultLifeSafetyListOfFaultValues, modePropertyReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag) *_BACnetFaultParameterFaultLifeSafety {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetFaultParameterFaultLifeSafety must not be nil")
	}
	if listOfFaultValues == nil {
		panic("listOfFaultValues of type BACnetFaultParameterFaultLifeSafetyListOfFaultValues for BACnetFaultParameterFaultLifeSafety must not be nil")
	}
	if modePropertyReference == nil {
		panic("modePropertyReference of type BACnetDeviceObjectPropertyReferenceEnclosed for BACnetFaultParameterFaultLifeSafety must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetFaultParameterFaultLifeSafety must not be nil")
	}
	_result := &_BACnetFaultParameterFaultLifeSafety{
		BACnetFaultParameterContract: NewBACnetFaultParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		ListOfFaultValues:            listOfFaultValues,
		ModePropertyReference:        modePropertyReference,
		ClosingTag:                   closingTag,
	}
	_result.BACnetFaultParameterContract.(*_BACnetFaultParameter)._SubType = _result
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

func (m *_BACnetFaultParameterFaultLifeSafety) GetParent() BACnetFaultParameterContract {
	return m.BACnetFaultParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultLifeSafety) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultLifeSafety) GetListOfFaultValues() BACnetFaultParameterFaultLifeSafetyListOfFaultValues {
	return m.ListOfFaultValues
}

func (m *_BACnetFaultParameterFaultLifeSafety) GetModePropertyReference() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.ModePropertyReference
}

func (m *_BACnetFaultParameterFaultLifeSafety) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultLifeSafety(structType any) BACnetFaultParameterFaultLifeSafety {
	if casted, ok := structType.(BACnetFaultParameterFaultLifeSafety); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultLifeSafety); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultLifeSafety) GetTypeName() string {
	return "BACnetFaultParameterFaultLifeSafety"
}

func (m *_BACnetFaultParameterFaultLifeSafety) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterContract.(*_BACnetFaultParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (listOfFaultValues)
	lengthInBits += m.ListOfFaultValues.GetLengthInBits(ctx)

	// Simple field (modePropertyReference)
	lengthInBits += m.ModePropertyReference.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultLifeSafety) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultLifeSafety) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameter) (__bACnetFaultParameterFaultLifeSafety BACnetFaultParameterFaultLifeSafety, err error) {
	m.BACnetFaultParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultLifeSafety"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultLifeSafety")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(3))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listOfFaultValues, err := ReadSimpleField[BACnetFaultParameterFaultLifeSafetyListOfFaultValues](ctx, "listOfFaultValues", ReadComplex[BACnetFaultParameterFaultLifeSafetyListOfFaultValues](BACnetFaultParameterFaultLifeSafetyListOfFaultValuesParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfFaultValues' field"))
	}
	m.ListOfFaultValues = listOfFaultValues

	modePropertyReference, err := ReadSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "modePropertyReference", ReadComplex[BACnetDeviceObjectPropertyReferenceEnclosed](BACnetDeviceObjectPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'modePropertyReference' field"))
	}
	m.ModePropertyReference = modePropertyReference

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(3))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultLifeSafety"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultLifeSafety")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultLifeSafety) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultLifeSafety) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultLifeSafety"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultLifeSafety")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetFaultParameterFaultLifeSafetyListOfFaultValues](ctx, "listOfFaultValues", m.GetListOfFaultValues(), WriteComplex[BACnetFaultParameterFaultLifeSafetyListOfFaultValues](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfFaultValues' field")
		}

		if err := WriteSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "modePropertyReference", m.GetModePropertyReference(), WriteComplex[BACnetDeviceObjectPropertyReferenceEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'modePropertyReference' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultLifeSafety"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultLifeSafety")
		}
		return nil
	}
	return m.BACnetFaultParameterContract.(*_BACnetFaultParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultLifeSafety) IsBACnetFaultParameterFaultLifeSafety() {}

func (m *_BACnetFaultParameterFaultLifeSafety) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
