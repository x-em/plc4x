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

// BACnetValueSourceNone is the corresponding interface of BACnetValueSourceNone
type BACnetValueSourceNone interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetValueSource
	// GetNone returns None (property field)
	GetNone() BACnetContextTagNull
	// IsBACnetValueSourceNone is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetValueSourceNone()
}

// _BACnetValueSourceNone is the data-structure of this message
type _BACnetValueSourceNone struct {
	BACnetValueSourceContract
	None BACnetContextTagNull
}

var _ BACnetValueSourceNone = (*_BACnetValueSourceNone)(nil)
var _ BACnetValueSourceRequirements = (*_BACnetValueSourceNone)(nil)

// NewBACnetValueSourceNone factory function for _BACnetValueSourceNone
func NewBACnetValueSourceNone(peekedTagHeader BACnetTagHeader, none BACnetContextTagNull) *_BACnetValueSourceNone {
	if none == nil {
		panic("none of type BACnetContextTagNull for BACnetValueSourceNone must not be nil")
	}
	_result := &_BACnetValueSourceNone{
		BACnetValueSourceContract: NewBACnetValueSource(peekedTagHeader),
		None:                      none,
	}
	_result.BACnetValueSourceContract.(*_BACnetValueSource)._SubType = _result
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

func (m *_BACnetValueSourceNone) GetParent() BACnetValueSourceContract {
	return m.BACnetValueSourceContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetValueSourceNone) GetNone() BACnetContextTagNull {
	return m.None
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetValueSourceNone(structType any) BACnetValueSourceNone {
	if casted, ok := structType.(BACnetValueSourceNone); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetValueSourceNone); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetValueSourceNone) GetTypeName() string {
	return "BACnetValueSourceNone"
}

func (m *_BACnetValueSourceNone) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetValueSourceContract.(*_BACnetValueSource).getLengthInBits(ctx))

	// Simple field (none)
	lengthInBits += m.None.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetValueSourceNone) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetValueSourceNone) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetValueSource) (__bACnetValueSourceNone BACnetValueSourceNone, err error) {
	m.BACnetValueSourceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetValueSourceNone"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetValueSourceNone")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	none, err := ReadSimpleField[BACnetContextTagNull](ctx, "none", ReadComplex[BACnetContextTagNull](BACnetContextTagParseWithBufferProducer[BACnetContextTagNull]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_NULL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'none' field"))
	}
	m.None = none

	if closeErr := readBuffer.CloseContext("BACnetValueSourceNone"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetValueSourceNone")
	}

	return m, nil
}

func (m *_BACnetValueSourceNone) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetValueSourceNone) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetValueSourceNone"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetValueSourceNone")
		}

		if err := WriteSimpleField[BACnetContextTagNull](ctx, "none", m.GetNone(), WriteComplex[BACnetContextTagNull](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'none' field")
		}

		if popErr := writeBuffer.PopContext("BACnetValueSourceNone"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetValueSourceNone")
		}
		return nil
	}
	return m.BACnetValueSourceContract.(*_BACnetValueSource).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetValueSourceNone) IsBACnetValueSourceNone() {}

func (m *_BACnetValueSourceNone) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
