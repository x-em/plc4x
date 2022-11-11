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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataDoNotHide is the corresponding interface of BACnetConstructedDataDoNotHide
type BACnetConstructedDataDoNotHide interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDoNotHide returns DoNotHide (property field)
	GetDoNotHide() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataDoNotHideExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDoNotHide.
// This is useful for switch cases.
type BACnetConstructedDataDoNotHideExactly interface {
	BACnetConstructedDataDoNotHide
	isBACnetConstructedDataDoNotHide() bool
}

// _BACnetConstructedDataDoNotHide is the data-structure of this message
type _BACnetConstructedDataDoNotHide struct {
	*_BACnetConstructedData
	DoNotHide BACnetApplicationTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDoNotHide) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDoNotHide) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DO_NOT_HIDE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDoNotHide) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDoNotHide) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDoNotHide) GetDoNotHide() BACnetApplicationTagBoolean {
	return m.DoNotHide
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDoNotHide) GetActualValue() BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetDoNotHide())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDoNotHide factory function for _BACnetConstructedDataDoNotHide
func NewBACnetConstructedDataDoNotHide(doNotHide BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDoNotHide {
	_result := &_BACnetConstructedDataDoNotHide{
		DoNotHide:              doNotHide,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDoNotHide(structType interface{}) BACnetConstructedDataDoNotHide {
	if casted, ok := structType.(BACnetConstructedDataDoNotHide); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDoNotHide); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDoNotHide) GetTypeName() string {
	return "BACnetConstructedDataDoNotHide"
}

func (m *_BACnetConstructedDataDoNotHide) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataDoNotHide) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (doNotHide)
	lengthInBits += m.DoNotHide.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDoNotHide) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDoNotHideParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDoNotHide, error) {
	return BACnetConstructedDataDoNotHideParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataDoNotHideParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDoNotHide, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDoNotHide"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDoNotHide")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (doNotHide)
	if pullErr := readBuffer.PullContext("doNotHide"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for doNotHide")
	}
	_doNotHide, _doNotHideErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _doNotHideErr != nil {
		return nil, errors.Wrap(_doNotHideErr, "Error parsing 'doNotHide' field of BACnetConstructedDataDoNotHide")
	}
	doNotHide := _doNotHide.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("doNotHide"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for doNotHide")
	}

	// Virtual field
	_actualValue := doNotHide
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDoNotHide"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDoNotHide")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDoNotHide{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		DoNotHide: doNotHide,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDoNotHide) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDoNotHide) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDoNotHide"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDoNotHide")
		}

		// Simple Field (doNotHide)
		if pushErr := writeBuffer.PushContext("doNotHide"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for doNotHide")
		}
		_doNotHideErr := writeBuffer.WriteSerializable(m.GetDoNotHide())
		if popErr := writeBuffer.PopContext("doNotHide"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for doNotHide")
		}
		if _doNotHideErr != nil {
			return errors.Wrap(_doNotHideErr, "Error serializing 'doNotHide' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDoNotHide"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDoNotHide")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDoNotHide) isBACnetConstructedDataDoNotHide() bool {
	return true
}

func (m *_BACnetConstructedDataDoNotHide) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
