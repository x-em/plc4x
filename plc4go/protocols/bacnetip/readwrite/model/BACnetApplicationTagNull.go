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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetApplicationTagNull is the corresponding interface of BACnetApplicationTagNull
type BACnetApplicationTagNull interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetApplicationTag
	// IsBACnetApplicationTagNull is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetApplicationTagNull()
}

// _BACnetApplicationTagNull is the data-structure of this message
type _BACnetApplicationTagNull struct {
	BACnetApplicationTagContract
}

var _ BACnetApplicationTagNull = (*_BACnetApplicationTagNull)(nil)
var _ BACnetApplicationTagRequirements = (*_BACnetApplicationTagNull)(nil)

// NewBACnetApplicationTagNull factory function for _BACnetApplicationTagNull
func NewBACnetApplicationTagNull(header BACnetTagHeader) *_BACnetApplicationTagNull {
	_result := &_BACnetApplicationTagNull{
		BACnetApplicationTagContract: NewBACnetApplicationTag(header),
	}
	_result.BACnetApplicationTagContract.(*_BACnetApplicationTag)._SubType = _result
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

func (m *_BACnetApplicationTagNull) GetParent() BACnetApplicationTagContract {
	return m.BACnetApplicationTagContract
}

// Deprecated: use the interface for direct cast
func CastBACnetApplicationTagNull(structType any) BACnetApplicationTagNull {
	if casted, ok := structType.(BACnetApplicationTagNull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetApplicationTagNull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetApplicationTagNull) GetTypeName() string {
	return "BACnetApplicationTagNull"
}

func (m *_BACnetApplicationTagNull) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetApplicationTagContract.(*_BACnetApplicationTag).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetApplicationTagNull) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetApplicationTagNull) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetApplicationTag) (__bACnetApplicationTagNull BACnetApplicationTagNull, err error) {
	m.BACnetApplicationTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagNull")
	}

	return m, nil
}

func (m *_BACnetApplicationTagNull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetApplicationTagNull) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagNull")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagNull")
		}
		return nil
	}
	return m.BACnetApplicationTagContract.(*_BACnetApplicationTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetApplicationTagNull) IsBACnetApplicationTagNull() {}

func (m *_BACnetApplicationTagNull) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
