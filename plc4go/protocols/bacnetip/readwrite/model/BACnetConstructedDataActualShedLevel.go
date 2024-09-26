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

// BACnetConstructedDataActualShedLevel is the corresponding interface of BACnetConstructedDataActualShedLevel
type BACnetConstructedDataActualShedLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetActualShedLevel returns ActualShedLevel (property field)
	GetActualShedLevel() BACnetShedLevel
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetShedLevel
	// IsBACnetConstructedDataActualShedLevel is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataActualShedLevel()
}

// _BACnetConstructedDataActualShedLevel is the data-structure of this message
type _BACnetConstructedDataActualShedLevel struct {
	BACnetConstructedDataContract
	ActualShedLevel BACnetShedLevel
}

var _ BACnetConstructedDataActualShedLevel = (*_BACnetConstructedDataActualShedLevel)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataActualShedLevel)(nil)

// NewBACnetConstructedDataActualShedLevel factory function for _BACnetConstructedDataActualShedLevel
func NewBACnetConstructedDataActualShedLevel(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, actualShedLevel BACnetShedLevel, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataActualShedLevel {
	if actualShedLevel == nil {
		panic("actualShedLevel of type BACnetShedLevel for BACnetConstructedDataActualShedLevel must not be nil")
	}
	_result := &_BACnetConstructedDataActualShedLevel{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ActualShedLevel:               actualShedLevel,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataActualShedLevel) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataActualShedLevel) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACTUAL_SHED_LEVEL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataActualShedLevel) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataActualShedLevel) GetActualShedLevel() BACnetShedLevel {
	return m.ActualShedLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataActualShedLevel) GetActualValue() BACnetShedLevel {
	ctx := context.Background()
	_ = ctx
	return CastBACnetShedLevel(m.GetActualShedLevel())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataActualShedLevel(structType any) BACnetConstructedDataActualShedLevel {
	if casted, ok := structType.(BACnetConstructedDataActualShedLevel); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataActualShedLevel); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataActualShedLevel) GetTypeName() string {
	return "BACnetConstructedDataActualShedLevel"
}

func (m *_BACnetConstructedDataActualShedLevel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (actualShedLevel)
	lengthInBits += m.ActualShedLevel.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataActualShedLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataActualShedLevel) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataActualShedLevel BACnetConstructedDataActualShedLevel, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataActualShedLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataActualShedLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	actualShedLevel, err := ReadSimpleField[BACnetShedLevel](ctx, "actualShedLevel", ReadComplex[BACnetShedLevel](BACnetShedLevelParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualShedLevel' field"))
	}
	m.ActualShedLevel = actualShedLevel

	actualValue, err := ReadVirtualField[BACnetShedLevel](ctx, "actualValue", (*BACnetShedLevel)(nil), actualShedLevel)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataActualShedLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataActualShedLevel")
	}

	return m, nil
}

func (m *_BACnetConstructedDataActualShedLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataActualShedLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataActualShedLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataActualShedLevel")
		}

		if err := WriteSimpleField[BACnetShedLevel](ctx, "actualShedLevel", m.GetActualShedLevel(), WriteComplex[BACnetShedLevel](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'actualShedLevel' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataActualShedLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataActualShedLevel")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataActualShedLevel) IsBACnetConstructedDataActualShedLevel() {}

func (m *_BACnetConstructedDataActualShedLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
