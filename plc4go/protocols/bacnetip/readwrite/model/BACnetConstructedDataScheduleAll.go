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

// BACnetConstructedDataScheduleAll is the corresponding interface of BACnetConstructedDataScheduleAll
type BACnetConstructedDataScheduleAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// IsBACnetConstructedDataScheduleAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataScheduleAll()
}

// _BACnetConstructedDataScheduleAll is the data-structure of this message
type _BACnetConstructedDataScheduleAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataScheduleAll = (*_BACnetConstructedDataScheduleAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataScheduleAll)(nil)

// NewBACnetConstructedDataScheduleAll factory function for _BACnetConstructedDataScheduleAll
func NewBACnetConstructedDataScheduleAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataScheduleAll {
	_result := &_BACnetConstructedDataScheduleAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataScheduleAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_SCHEDULE
}

func (m *_BACnetConstructedDataScheduleAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataScheduleAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataScheduleAll(structType any) BACnetConstructedDataScheduleAll {
	if casted, ok := structType.(BACnetConstructedDataScheduleAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataScheduleAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataScheduleAll) GetTypeName() string {
	return "BACnetConstructedDataScheduleAll"
}

func (m *_BACnetConstructedDataScheduleAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataScheduleAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataScheduleAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataScheduleAll BACnetConstructedDataScheduleAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataScheduleAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataScheduleAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataScheduleAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataScheduleAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataScheduleAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataScheduleAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataScheduleAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataScheduleAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataScheduleAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataScheduleAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataScheduleAll) IsBACnetConstructedDataScheduleAll() {}

func (m *_BACnetConstructedDataScheduleAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
