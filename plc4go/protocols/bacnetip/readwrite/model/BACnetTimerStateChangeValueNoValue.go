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

// BACnetTimerStateChangeValueNoValue is the corresponding interface of BACnetTimerStateChangeValueNoValue
type BACnetTimerStateChangeValueNoValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetTimerStateChangeValue
	// GetNoValue returns NoValue (property field)
	GetNoValue() BACnetContextTagNull
	// IsBACnetTimerStateChangeValueNoValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimerStateChangeValueNoValue()
}

// _BACnetTimerStateChangeValueNoValue is the data-structure of this message
type _BACnetTimerStateChangeValueNoValue struct {
	BACnetTimerStateChangeValueContract
	NoValue BACnetContextTagNull
}

var _ BACnetTimerStateChangeValueNoValue = (*_BACnetTimerStateChangeValueNoValue)(nil)
var _ BACnetTimerStateChangeValueRequirements = (*_BACnetTimerStateChangeValueNoValue)(nil)

// NewBACnetTimerStateChangeValueNoValue factory function for _BACnetTimerStateChangeValueNoValue
func NewBACnetTimerStateChangeValueNoValue(peekedTagHeader BACnetTagHeader, noValue BACnetContextTagNull, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueNoValue {
	if noValue == nil {
		panic("noValue of type BACnetContextTagNull for BACnetTimerStateChangeValueNoValue must not be nil")
	}
	_result := &_BACnetTimerStateChangeValueNoValue{
		BACnetTimerStateChangeValueContract: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
		NoValue:                             noValue,
	}
	_result.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = _result
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

func (m *_BACnetTimerStateChangeValueNoValue) GetParent() BACnetTimerStateChangeValueContract {
	return m.BACnetTimerStateChangeValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueNoValue) GetNoValue() BACnetContextTagNull {
	return m.NoValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueNoValue(structType any) BACnetTimerStateChangeValueNoValue {
	if casted, ok := structType.(BACnetTimerStateChangeValueNoValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueNoValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueNoValue) GetTypeName() string {
	return "BACnetTimerStateChangeValueNoValue"
}

func (m *_BACnetTimerStateChangeValueNoValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).getLengthInBits(ctx))

	// Simple field (noValue)
	lengthInBits += m.NoValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueNoValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetTimerStateChangeValueNoValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetTimerStateChangeValue, objectTypeArgument BACnetObjectType) (__bACnetTimerStateChangeValueNoValue BACnetTimerStateChangeValueNoValue, err error) {
	m.BACnetTimerStateChangeValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueNoValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueNoValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	noValue, err := ReadSimpleField[BACnetContextTagNull](ctx, "noValue", ReadComplex[BACnetContextTagNull](BACnetContextTagParseWithBufferProducer[BACnetContextTagNull]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_NULL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noValue' field"))
	}
	m.NoValue = noValue

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueNoValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueNoValue")
	}

	return m, nil
}

func (m *_BACnetTimerStateChangeValueNoValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerStateChangeValueNoValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueNoValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueNoValue")
		}

		if err := WriteSimpleField[BACnetContextTagNull](ctx, "noValue", m.GetNoValue(), WriteComplex[BACnetContextTagNull](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'noValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueNoValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueNoValue")
		}
		return nil
	}
	return m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueNoValue) IsBACnetTimerStateChangeValueNoValue() {}

func (m *_BACnetTimerStateChangeValueNoValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
