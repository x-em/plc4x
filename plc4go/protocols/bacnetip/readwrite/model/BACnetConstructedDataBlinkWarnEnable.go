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

// BACnetConstructedDataBlinkWarnEnable is the corresponding interface of BACnetConstructedDataBlinkWarnEnable
type BACnetConstructedDataBlinkWarnEnable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetBlinkWarnEnable returns BlinkWarnEnable (property field)
	GetBlinkWarnEnable() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataBlinkWarnEnable is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBlinkWarnEnable()
}

// _BACnetConstructedDataBlinkWarnEnable is the data-structure of this message
type _BACnetConstructedDataBlinkWarnEnable struct {
	BACnetConstructedDataContract
	BlinkWarnEnable BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataBlinkWarnEnable = (*_BACnetConstructedDataBlinkWarnEnable)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBlinkWarnEnable)(nil)

// NewBACnetConstructedDataBlinkWarnEnable factory function for _BACnetConstructedDataBlinkWarnEnable
func NewBACnetConstructedDataBlinkWarnEnable(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, blinkWarnEnable BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBlinkWarnEnable {
	if blinkWarnEnable == nil {
		panic("blinkWarnEnable of type BACnetApplicationTagBoolean for BACnetConstructedDataBlinkWarnEnable must not be nil")
	}
	_result := &_BACnetConstructedDataBlinkWarnEnable{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		BlinkWarnEnable:               blinkWarnEnable,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBlinkWarnEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBlinkWarnEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BLINK_WARN_ENABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBlinkWarnEnable) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBlinkWarnEnable) GetBlinkWarnEnable() BACnetApplicationTagBoolean {
	return m.BlinkWarnEnable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBlinkWarnEnable) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetBlinkWarnEnable())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBlinkWarnEnable(structType any) BACnetConstructedDataBlinkWarnEnable {
	if casted, ok := structType.(BACnetConstructedDataBlinkWarnEnable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBlinkWarnEnable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBlinkWarnEnable) GetTypeName() string {
	return "BACnetConstructedDataBlinkWarnEnable"
}

func (m *_BACnetConstructedDataBlinkWarnEnable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (blinkWarnEnable)
	lengthInBits += m.BlinkWarnEnable.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBlinkWarnEnable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBlinkWarnEnable) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBlinkWarnEnable BACnetConstructedDataBlinkWarnEnable, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBlinkWarnEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBlinkWarnEnable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	blinkWarnEnable, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "blinkWarnEnable", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'blinkWarnEnable' field"))
	}
	m.BlinkWarnEnable = blinkWarnEnable

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), blinkWarnEnable)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBlinkWarnEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBlinkWarnEnable")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBlinkWarnEnable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBlinkWarnEnable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBlinkWarnEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBlinkWarnEnable")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "blinkWarnEnable", m.GetBlinkWarnEnable(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'blinkWarnEnable' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBlinkWarnEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBlinkWarnEnable")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBlinkWarnEnable) IsBACnetConstructedDataBlinkWarnEnable() {}

func (m *_BACnetConstructedDataBlinkWarnEnable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
