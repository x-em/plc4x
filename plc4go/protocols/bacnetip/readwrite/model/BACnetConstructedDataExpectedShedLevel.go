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

// BACnetConstructedDataExpectedShedLevel is the corresponding interface of BACnetConstructedDataExpectedShedLevel
type BACnetConstructedDataExpectedShedLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetExpectedShedLevel returns ExpectedShedLevel (property field)
	GetExpectedShedLevel() BACnetShedLevel
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetShedLevel
	// IsBACnetConstructedDataExpectedShedLevel is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataExpectedShedLevel()
}

// _BACnetConstructedDataExpectedShedLevel is the data-structure of this message
type _BACnetConstructedDataExpectedShedLevel struct {
	BACnetConstructedDataContract
	ExpectedShedLevel BACnetShedLevel
}

var _ BACnetConstructedDataExpectedShedLevel = (*_BACnetConstructedDataExpectedShedLevel)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataExpectedShedLevel)(nil)

// NewBACnetConstructedDataExpectedShedLevel factory function for _BACnetConstructedDataExpectedShedLevel
func NewBACnetConstructedDataExpectedShedLevel(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, expectedShedLevel BACnetShedLevel, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataExpectedShedLevel {
	if expectedShedLevel == nil {
		panic("expectedShedLevel of type BACnetShedLevel for BACnetConstructedDataExpectedShedLevel must not be nil")
	}
	_result := &_BACnetConstructedDataExpectedShedLevel{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ExpectedShedLevel:             expectedShedLevel,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataExpectedShedLevel) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataExpectedShedLevel) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EXPECTED_SHED_LEVEL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataExpectedShedLevel) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataExpectedShedLevel) GetExpectedShedLevel() BACnetShedLevel {
	return m.ExpectedShedLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataExpectedShedLevel) GetActualValue() BACnetShedLevel {
	ctx := context.Background()
	_ = ctx
	return CastBACnetShedLevel(m.GetExpectedShedLevel())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataExpectedShedLevel(structType any) BACnetConstructedDataExpectedShedLevel {
	if casted, ok := structType.(BACnetConstructedDataExpectedShedLevel); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataExpectedShedLevel); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataExpectedShedLevel) GetTypeName() string {
	return "BACnetConstructedDataExpectedShedLevel"
}

func (m *_BACnetConstructedDataExpectedShedLevel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (expectedShedLevel)
	lengthInBits += m.ExpectedShedLevel.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataExpectedShedLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataExpectedShedLevel) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataExpectedShedLevel BACnetConstructedDataExpectedShedLevel, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataExpectedShedLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataExpectedShedLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	expectedShedLevel, err := ReadSimpleField[BACnetShedLevel](ctx, "expectedShedLevel", ReadComplex[BACnetShedLevel](BACnetShedLevelParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'expectedShedLevel' field"))
	}
	m.ExpectedShedLevel = expectedShedLevel

	actualValue, err := ReadVirtualField[BACnetShedLevel](ctx, "actualValue", (*BACnetShedLevel)(nil), expectedShedLevel)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataExpectedShedLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataExpectedShedLevel")
	}

	return m, nil
}

func (m *_BACnetConstructedDataExpectedShedLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataExpectedShedLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataExpectedShedLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataExpectedShedLevel")
		}

		if err := WriteSimpleField[BACnetShedLevel](ctx, "expectedShedLevel", m.GetExpectedShedLevel(), WriteComplex[BACnetShedLevel](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'expectedShedLevel' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataExpectedShedLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataExpectedShedLevel")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataExpectedShedLevel) IsBACnetConstructedDataExpectedShedLevel() {}

func (m *_BACnetConstructedDataExpectedShedLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
