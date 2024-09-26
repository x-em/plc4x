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

// BACnetConstructedDataLastNotifyRecord is the corresponding interface of BACnetConstructedDataLastNotifyRecord
type BACnetConstructedDataLastNotifyRecord interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLastNotifyRecord returns LastNotifyRecord (property field)
	GetLastNotifyRecord() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataLastNotifyRecord is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLastNotifyRecord()
}

// _BACnetConstructedDataLastNotifyRecord is the data-structure of this message
type _BACnetConstructedDataLastNotifyRecord struct {
	BACnetConstructedDataContract
	LastNotifyRecord BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataLastNotifyRecord = (*_BACnetConstructedDataLastNotifyRecord)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLastNotifyRecord)(nil)

// NewBACnetConstructedDataLastNotifyRecord factory function for _BACnetConstructedDataLastNotifyRecord
func NewBACnetConstructedDataLastNotifyRecord(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lastNotifyRecord BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastNotifyRecord {
	if lastNotifyRecord == nil {
		panic("lastNotifyRecord of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataLastNotifyRecord must not be nil")
	}
	_result := &_BACnetConstructedDataLastNotifyRecord{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LastNotifyRecord:              lastNotifyRecord,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastNotifyRecord) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastNotifyRecord) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_NOTIFY_RECORD
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastNotifyRecord) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastNotifyRecord) GetLastNotifyRecord() BACnetApplicationTagUnsignedInteger {
	return m.LastNotifyRecord
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastNotifyRecord) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetLastNotifyRecord())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastNotifyRecord(structType any) BACnetConstructedDataLastNotifyRecord {
	if casted, ok := structType.(BACnetConstructedDataLastNotifyRecord); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastNotifyRecord); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastNotifyRecord) GetTypeName() string {
	return "BACnetConstructedDataLastNotifyRecord"
}

func (m *_BACnetConstructedDataLastNotifyRecord) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lastNotifyRecord)
	lengthInBits += m.LastNotifyRecord.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastNotifyRecord) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLastNotifyRecord) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLastNotifyRecord BACnetConstructedDataLastNotifyRecord, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastNotifyRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastNotifyRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lastNotifyRecord, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "lastNotifyRecord", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastNotifyRecord' field"))
	}
	m.LastNotifyRecord = lastNotifyRecord

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), lastNotifyRecord)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastNotifyRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastNotifyRecord")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLastNotifyRecord) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastNotifyRecord) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastNotifyRecord"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastNotifyRecord")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "lastNotifyRecord", m.GetLastNotifyRecord(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastNotifyRecord' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastNotifyRecord"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastNotifyRecord")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastNotifyRecord) IsBACnetConstructedDataLastNotifyRecord() {}

func (m *_BACnetConstructedDataLastNotifyRecord) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
