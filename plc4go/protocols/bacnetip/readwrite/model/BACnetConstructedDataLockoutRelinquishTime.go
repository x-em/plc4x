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

// BACnetConstructedDataLockoutRelinquishTime is the corresponding interface of BACnetConstructedDataLockoutRelinquishTime
type BACnetConstructedDataLockoutRelinquishTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLockoutRelinquishTime returns LockoutRelinquishTime (property field)
	GetLockoutRelinquishTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataLockoutRelinquishTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLockoutRelinquishTime()
}

// _BACnetConstructedDataLockoutRelinquishTime is the data-structure of this message
type _BACnetConstructedDataLockoutRelinquishTime struct {
	BACnetConstructedDataContract
	LockoutRelinquishTime BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataLockoutRelinquishTime = (*_BACnetConstructedDataLockoutRelinquishTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLockoutRelinquishTime)(nil)

// NewBACnetConstructedDataLockoutRelinquishTime factory function for _BACnetConstructedDataLockoutRelinquishTime
func NewBACnetConstructedDataLockoutRelinquishTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lockoutRelinquishTime BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLockoutRelinquishTime {
	if lockoutRelinquishTime == nil {
		panic("lockoutRelinquishTime of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataLockoutRelinquishTime must not be nil")
	}
	_result := &_BACnetConstructedDataLockoutRelinquishTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LockoutRelinquishTime:         lockoutRelinquishTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOCKOUT_RELINQUISH_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetLockoutRelinquishTime() BACnetApplicationTagUnsignedInteger {
	return m.LockoutRelinquishTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetLockoutRelinquishTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLockoutRelinquishTime(structType any) BACnetConstructedDataLockoutRelinquishTime {
	if casted, ok := structType.(BACnetConstructedDataLockoutRelinquishTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLockoutRelinquishTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetTypeName() string {
	return "BACnetConstructedDataLockoutRelinquishTime"
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lockoutRelinquishTime)
	lengthInBits += m.LockoutRelinquishTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLockoutRelinquishTime BACnetConstructedDataLockoutRelinquishTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLockoutRelinquishTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLockoutRelinquishTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lockoutRelinquishTime, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "lockoutRelinquishTime", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lockoutRelinquishTime' field"))
	}
	m.LockoutRelinquishTime = lockoutRelinquishTime

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), lockoutRelinquishTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLockoutRelinquishTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLockoutRelinquishTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLockoutRelinquishTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLockoutRelinquishTime")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "lockoutRelinquishTime", m.GetLockoutRelinquishTime(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lockoutRelinquishTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLockoutRelinquishTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLockoutRelinquishTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) IsBACnetConstructedDataLockoutRelinquishTime() {
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
