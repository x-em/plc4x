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

// BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned
type BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameterFaultOutOfRangeMaxNormalValue
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned()
}

// _BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned struct {
	BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract
	UnsignedValue BACnetApplicationTagUnsignedInteger
}

var _ BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned = (*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned)(nil)
var _ BACnetFaultParameterFaultOutOfRangeMaxNormalValueRequirements = (*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned)(nil)

// NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned factory function for _BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned
func NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, unsignedValue BACnetApplicationTagUnsignedInteger, tagNumber uint8) *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned {
	if unsignedValue == nil {
		panic("unsignedValue of type BACnetApplicationTagUnsignedInteger for BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned must not be nil")
	}
	_result := &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned{
		BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract: NewBACnetFaultParameterFaultOutOfRangeMaxNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
		UnsignedValue: unsignedValue,
	}
	_result.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue)._SubType = _result
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

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) GetParent() BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract {
	return m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) GetUnsignedValue() BACnetApplicationTagUnsignedInteger {
	return m.UnsignedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned(structType any) BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned"
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue).getLengthInBits(ctx))

	// Simple field (unsignedValue)
	lengthInBits += m.UnsignedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultOutOfRangeMaxNormalValue, tagNumber uint8) (__bACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned, err error) {
	m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	unsignedValue, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "unsignedValue", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unsignedValue' field"))
	}
	m.UnsignedValue = unsignedValue

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "unsignedValue", m.GetUnsignedValue(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'unsignedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) IsBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned() {
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
