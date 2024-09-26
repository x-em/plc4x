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

// BACnetLogDataLogDataEntryIntegerValue is the corresponding interface of BACnetLogDataLogDataEntryIntegerValue
type BACnetLogDataLogDataEntryIntegerValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetLogDataLogDataEntry
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() BACnetContextTagSignedInteger
	// IsBACnetLogDataLogDataEntryIntegerValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogDataLogDataEntryIntegerValue()
}

// _BACnetLogDataLogDataEntryIntegerValue is the data-structure of this message
type _BACnetLogDataLogDataEntryIntegerValue struct {
	BACnetLogDataLogDataEntryContract
	IntegerValue BACnetContextTagSignedInteger
}

var _ BACnetLogDataLogDataEntryIntegerValue = (*_BACnetLogDataLogDataEntryIntegerValue)(nil)
var _ BACnetLogDataLogDataEntryRequirements = (*_BACnetLogDataLogDataEntryIntegerValue)(nil)

// NewBACnetLogDataLogDataEntryIntegerValue factory function for _BACnetLogDataLogDataEntryIntegerValue
func NewBACnetLogDataLogDataEntryIntegerValue(peekedTagHeader BACnetTagHeader, integerValue BACnetContextTagSignedInteger) *_BACnetLogDataLogDataEntryIntegerValue {
	if integerValue == nil {
		panic("integerValue of type BACnetContextTagSignedInteger for BACnetLogDataLogDataEntryIntegerValue must not be nil")
	}
	_result := &_BACnetLogDataLogDataEntryIntegerValue{
		BACnetLogDataLogDataEntryContract: NewBACnetLogDataLogDataEntry(peekedTagHeader),
		IntegerValue:                      integerValue,
	}
	_result.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry)._SubType = _result
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

func (m *_BACnetLogDataLogDataEntryIntegerValue) GetParent() BACnetLogDataLogDataEntryContract {
	return m.BACnetLogDataLogDataEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogDataEntryIntegerValue) GetIntegerValue() BACnetContextTagSignedInteger {
	return m.IntegerValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogDataEntryIntegerValue(structType any) BACnetLogDataLogDataEntryIntegerValue {
	if casted, ok := structType.(BACnetLogDataLogDataEntryIntegerValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntryIntegerValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogDataEntryIntegerValue) GetTypeName() string {
	return "BACnetLogDataLogDataEntryIntegerValue"
}

func (m *_BACnetLogDataLogDataEntryIntegerValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry).getLengthInBits(ctx))

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogDataLogDataEntryIntegerValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogDataLogDataEntryIntegerValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogDataLogDataEntry) (__bACnetLogDataLogDataEntryIntegerValue BACnetLogDataLogDataEntryIntegerValue, err error) {
	m.BACnetLogDataLogDataEntryContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataEntryIntegerValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataEntryIntegerValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	integerValue, err := ReadSimpleField[BACnetContextTagSignedInteger](ctx, "integerValue", ReadComplex[BACnetContextTagSignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagSignedInteger]((uint8)(uint8(4)), (BACnetDataType)(BACnetDataType_SIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'integerValue' field"))
	}
	m.IntegerValue = integerValue

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataEntryIntegerValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataEntryIntegerValue")
	}

	return m, nil
}

func (m *_BACnetLogDataLogDataEntryIntegerValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogDataLogDataEntryIntegerValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataEntryIntegerValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataEntryIntegerValue")
		}

		if err := WriteSimpleField[BACnetContextTagSignedInteger](ctx, "integerValue", m.GetIntegerValue(), WriteComplex[BACnetContextTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'integerValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataEntryIntegerValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataEntryIntegerValue")
		}
		return nil
	}
	return m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogDataEntryIntegerValue) IsBACnetLogDataLogDataEntryIntegerValue() {}

func (m *_BACnetLogDataLogDataEntryIntegerValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
