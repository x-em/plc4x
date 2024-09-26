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

// BACnetPriorityValueBitString is the corresponding interface of BACnetPriorityValueBitString
type BACnetPriorityValueBitString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPriorityValue
	// GetBitStringValue returns BitStringValue (property field)
	GetBitStringValue() BACnetApplicationTagBitString
	// IsBACnetPriorityValueBitString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPriorityValueBitString()
}

// _BACnetPriorityValueBitString is the data-structure of this message
type _BACnetPriorityValueBitString struct {
	BACnetPriorityValueContract
	BitStringValue BACnetApplicationTagBitString
}

var _ BACnetPriorityValueBitString = (*_BACnetPriorityValueBitString)(nil)
var _ BACnetPriorityValueRequirements = (*_BACnetPriorityValueBitString)(nil)

// NewBACnetPriorityValueBitString factory function for _BACnetPriorityValueBitString
func NewBACnetPriorityValueBitString(peekedTagHeader BACnetTagHeader, bitStringValue BACnetApplicationTagBitString, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueBitString {
	if bitStringValue == nil {
		panic("bitStringValue of type BACnetApplicationTagBitString for BACnetPriorityValueBitString must not be nil")
	}
	_result := &_BACnetPriorityValueBitString{
		BACnetPriorityValueContract: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
		BitStringValue:              bitStringValue,
	}
	_result.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = _result
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

func (m *_BACnetPriorityValueBitString) GetParent() BACnetPriorityValueContract {
	return m.BACnetPriorityValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueBitString) GetBitStringValue() BACnetApplicationTagBitString {
	return m.BitStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueBitString(structType any) BACnetPriorityValueBitString {
	if casted, ok := structType.(BACnetPriorityValueBitString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueBitString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueBitString) GetTypeName() string {
	return "BACnetPriorityValueBitString"
}

func (m *_BACnetPriorityValueBitString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPriorityValueContract.(*_BACnetPriorityValue).getLengthInBits(ctx))

	// Simple field (bitStringValue)
	lengthInBits += m.BitStringValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPriorityValueBitString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPriorityValueBitString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPriorityValue, objectTypeArgument BACnetObjectType) (__bACnetPriorityValueBitString BACnetPriorityValueBitString, err error) {
	m.BACnetPriorityValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueBitString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueBitString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bitStringValue, err := ReadSimpleField[BACnetApplicationTagBitString](ctx, "bitStringValue", ReadComplex[BACnetApplicationTagBitString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBitString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bitStringValue' field"))
	}
	m.BitStringValue = bitStringValue

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueBitString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueBitString")
	}

	return m, nil
}

func (m *_BACnetPriorityValueBitString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueBitString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueBitString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueBitString")
		}

		if err := WriteSimpleField[BACnetApplicationTagBitString](ctx, "bitStringValue", m.GetBitStringValue(), WriteComplex[BACnetApplicationTagBitString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'bitStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueBitString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueBitString")
		}
		return nil
	}
	return m.BACnetPriorityValueContract.(*_BACnetPriorityValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueBitString) IsBACnetPriorityValueBitString() {}

func (m *_BACnetPriorityValueBitString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
