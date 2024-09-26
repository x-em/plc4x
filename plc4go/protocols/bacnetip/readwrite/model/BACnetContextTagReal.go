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

// BACnetContextTagReal is the corresponding interface of BACnetContextTagReal
type BACnetContextTagReal interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() float32
	// IsBACnetContextTagReal is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetContextTagReal()
}

// _BACnetContextTagReal is the data-structure of this message
type _BACnetContextTagReal struct {
	BACnetContextTagContract
	Payload BACnetTagPayloadReal
}

var _ BACnetContextTagReal = (*_BACnetContextTagReal)(nil)
var _ BACnetContextTagRequirements = (*_BACnetContextTagReal)(nil)

// NewBACnetContextTagReal factory function for _BACnetContextTagReal
func NewBACnetContextTagReal(header BACnetTagHeader, payload BACnetTagPayloadReal, tagNumberArgument uint8) *_BACnetContextTagReal {
	if payload == nil {
		panic("payload of type BACnetTagPayloadReal for BACnetContextTagReal must not be nil")
	}
	_result := &_BACnetContextTagReal{
		BACnetContextTagContract: NewBACnetContextTag(header, tagNumberArgument),
		Payload:                  payload,
	}
	_result.BACnetContextTagContract.(*_BACnetContextTag)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagReal) GetDataType() BACnetDataType {
	return BACnetDataType_REAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagReal) GetParent() BACnetContextTagContract {
	return m.BACnetContextTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagReal) GetPayload() BACnetTagPayloadReal {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetContextTagReal) GetActualValue() float32 {
	ctx := context.Background()
	_ = ctx
	return float32(m.GetPayload().GetValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetContextTagReal(structType any) BACnetContextTagReal {
	if casted, ok := structType.(BACnetContextTagReal); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagReal); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagReal) GetTypeName() string {
	return "BACnetContextTagReal"
}

func (m *_BACnetContextTagReal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetContextTagContract.(*_BACnetContextTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetContextTagReal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetContextTagReal) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetContextTag, tagNumberArgument uint8, dataType BACnetDataType) (__bACnetContextTagReal BACnetContextTagReal, err error) {
	m.BACnetContextTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagReal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagReal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadReal](ctx, "payload", ReadComplex[BACnetTagPayloadReal](BACnetTagPayloadRealParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	actualValue, err := ReadVirtualField[float32](ctx, "actualValue", (*float32)(nil), payload.GetValue())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetContextTagReal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagReal")
	}

	return m, nil
}

func (m *_BACnetContextTagReal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetContextTagReal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagReal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagReal")
		}

		if err := WriteSimpleField[BACnetTagPayloadReal](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagReal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagReal")
		}
		return nil
	}
	return m.BACnetContextTagContract.(*_BACnetContextTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetContextTagReal) IsBACnetContextTagReal() {}

func (m *_BACnetContextTagReal) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
