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

// BACnetEventParameterDoubleOutOfRange is the corresponding interface of BACnetEventParameterDoubleOutOfRange
type BACnetEventParameterDoubleOutOfRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetLowLimit returns LowLimit (property field)
	GetLowLimit() BACnetContextTagDouble
	// GetHighLimit returns HighLimit (property field)
	GetHighLimit() BACnetContextTagDouble
	// GetDeadband returns Deadband (property field)
	GetDeadband() BACnetContextTagDouble
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventParameterDoubleOutOfRange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterDoubleOutOfRange()
}

// _BACnetEventParameterDoubleOutOfRange is the data-structure of this message
type _BACnetEventParameterDoubleOutOfRange struct {
	BACnetEventParameterContract
	OpeningTag BACnetOpeningTag
	TimeDelay  BACnetContextTagUnsignedInteger
	LowLimit   BACnetContextTagDouble
	HighLimit  BACnetContextTagDouble
	Deadband   BACnetContextTagDouble
	ClosingTag BACnetClosingTag
}

var _ BACnetEventParameterDoubleOutOfRange = (*_BACnetEventParameterDoubleOutOfRange)(nil)
var _ BACnetEventParameterRequirements = (*_BACnetEventParameterDoubleOutOfRange)(nil)

// NewBACnetEventParameterDoubleOutOfRange factory function for _BACnetEventParameterDoubleOutOfRange
func NewBACnetEventParameterDoubleOutOfRange(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, lowLimit BACnetContextTagDouble, highLimit BACnetContextTagDouble, deadband BACnetContextTagDouble, closingTag BACnetClosingTag) *_BACnetEventParameterDoubleOutOfRange {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventParameterDoubleOutOfRange must not be nil")
	}
	if timeDelay == nil {
		panic("timeDelay of type BACnetContextTagUnsignedInteger for BACnetEventParameterDoubleOutOfRange must not be nil")
	}
	if lowLimit == nil {
		panic("lowLimit of type BACnetContextTagDouble for BACnetEventParameterDoubleOutOfRange must not be nil")
	}
	if highLimit == nil {
		panic("highLimit of type BACnetContextTagDouble for BACnetEventParameterDoubleOutOfRange must not be nil")
	}
	if deadband == nil {
		panic("deadband of type BACnetContextTagDouble for BACnetEventParameterDoubleOutOfRange must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventParameterDoubleOutOfRange must not be nil")
	}
	_result := &_BACnetEventParameterDoubleOutOfRange{
		BACnetEventParameterContract: NewBACnetEventParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		TimeDelay:                    timeDelay,
		LowLimit:                     lowLimit,
		HighLimit:                    highLimit,
		Deadband:                     deadband,
		ClosingTag:                   closingTag,
	}
	_result.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = _result
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

func (m *_BACnetEventParameterDoubleOutOfRange) GetParent() BACnetEventParameterContract {
	return m.BACnetEventParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterDoubleOutOfRange) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterDoubleOutOfRange) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterDoubleOutOfRange) GetLowLimit() BACnetContextTagDouble {
	return m.LowLimit
}

func (m *_BACnetEventParameterDoubleOutOfRange) GetHighLimit() BACnetContextTagDouble {
	return m.HighLimit
}

func (m *_BACnetEventParameterDoubleOutOfRange) GetDeadband() BACnetContextTagDouble {
	return m.Deadband
}

func (m *_BACnetEventParameterDoubleOutOfRange) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterDoubleOutOfRange(structType any) BACnetEventParameterDoubleOutOfRange {
	if casted, ok := structType.(BACnetEventParameterDoubleOutOfRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterDoubleOutOfRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterDoubleOutOfRange) GetTypeName() string {
	return "BACnetEventParameterDoubleOutOfRange"
}

func (m *_BACnetEventParameterDoubleOutOfRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventParameterContract.(*_BACnetEventParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits(ctx)

	// Simple field (lowLimit)
	lengthInBits += m.LowLimit.GetLengthInBits(ctx)

	// Simple field (highLimit)
	lengthInBits += m.HighLimit.GetLengthInBits(ctx)

	// Simple field (deadband)
	lengthInBits += m.Deadband.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterDoubleOutOfRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventParameterDoubleOutOfRange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventParameter) (__bACnetEventParameterDoubleOutOfRange BACnetEventParameterDoubleOutOfRange, err error) {
	m.BACnetEventParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterDoubleOutOfRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterDoubleOutOfRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(14))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	timeDelay, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeDelay' field"))
	}
	m.TimeDelay = timeDelay

	lowLimit, err := ReadSimpleField[BACnetContextTagDouble](ctx, "lowLimit", ReadComplex[BACnetContextTagDouble](BACnetContextTagParseWithBufferProducer[BACnetContextTagDouble]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_DOUBLE)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lowLimit' field"))
	}
	m.LowLimit = lowLimit

	highLimit, err := ReadSimpleField[BACnetContextTagDouble](ctx, "highLimit", ReadComplex[BACnetContextTagDouble](BACnetContextTagParseWithBufferProducer[BACnetContextTagDouble]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_DOUBLE)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'highLimit' field"))
	}
	m.HighLimit = highLimit

	deadband, err := ReadSimpleField[BACnetContextTagDouble](ctx, "deadband", ReadComplex[BACnetContextTagDouble](BACnetContextTagParseWithBufferProducer[BACnetContextTagDouble]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_DOUBLE)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deadband' field"))
	}
	m.Deadband = deadband

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(14))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterDoubleOutOfRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterDoubleOutOfRange")
	}

	return m, nil
}

func (m *_BACnetEventParameterDoubleOutOfRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterDoubleOutOfRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterDoubleOutOfRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterDoubleOutOfRange")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", m.GetTimeDelay(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeDelay' field")
		}

		if err := WriteSimpleField[BACnetContextTagDouble](ctx, "lowLimit", m.GetLowLimit(), WriteComplex[BACnetContextTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lowLimit' field")
		}

		if err := WriteSimpleField[BACnetContextTagDouble](ctx, "highLimit", m.GetHighLimit(), WriteComplex[BACnetContextTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'highLimit' field")
		}

		if err := WriteSimpleField[BACnetContextTagDouble](ctx, "deadband", m.GetDeadband(), WriteComplex[BACnetContextTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deadband' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterDoubleOutOfRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterDoubleOutOfRange")
		}
		return nil
	}
	return m.BACnetEventParameterContract.(*_BACnetEventParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterDoubleOutOfRange) IsBACnetEventParameterDoubleOutOfRange() {}

func (m *_BACnetEventParameterDoubleOutOfRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
