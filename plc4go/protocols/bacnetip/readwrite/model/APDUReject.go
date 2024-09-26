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

// APDUReject is the corresponding interface of APDUReject
type APDUReject interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	APDU
	// GetOriginalInvokeId returns OriginalInvokeId (property field)
	GetOriginalInvokeId() uint8
	// GetRejectReason returns RejectReason (property field)
	GetRejectReason() BACnetRejectReasonTagged
	// IsAPDUReject is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAPDUReject()
}

// _APDUReject is the data-structure of this message
type _APDUReject struct {
	APDUContract
	OriginalInvokeId uint8
	RejectReason     BACnetRejectReasonTagged
	// Reserved Fields
	reservedField0 *uint8
}

var _ APDUReject = (*_APDUReject)(nil)
var _ APDURequirements = (*_APDUReject)(nil)

// NewAPDUReject factory function for _APDUReject
func NewAPDUReject(originalInvokeId uint8, rejectReason BACnetRejectReasonTagged, apduLength uint16) *_APDUReject {
	if rejectReason == nil {
		panic("rejectReason of type BACnetRejectReasonTagged for APDUReject must not be nil")
	}
	_result := &_APDUReject{
		APDUContract:     NewAPDU(apduLength),
		OriginalInvokeId: originalInvokeId,
		RejectReason:     rejectReason,
	}
	_result.APDUContract.(*_APDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUReject) GetApduType() ApduType {
	return ApduType_REJECT_PDU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUReject) GetParent() APDUContract {
	return m.APDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUReject) GetOriginalInvokeId() uint8 {
	return m.OriginalInvokeId
}

func (m *_APDUReject) GetRejectReason() BACnetRejectReasonTagged {
	return m.RejectReason
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAPDUReject(structType any) APDUReject {
	if casted, ok := structType.(APDUReject); ok {
		return casted
	}
	if casted, ok := structType.(*APDUReject); ok {
		return *casted
	}
	return nil
}

func (m *_APDUReject) GetTypeName() string {
	return "APDUReject"
}

func (m *_APDUReject) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.APDUContract.(*_APDU).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Simple field (rejectReason)
	lengthInBits += m.RejectReason.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_APDUReject) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_APDUReject) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_APDU, apduLength uint16) (__aPDUReject APDUReject, err error) {
	m.APDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUReject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUReject")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(4)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	originalInvokeId, err := ReadSimpleField(ctx, "originalInvokeId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originalInvokeId' field"))
	}
	m.OriginalInvokeId = originalInvokeId

	rejectReason, err := ReadSimpleField[BACnetRejectReasonTagged](ctx, "rejectReason", ReadComplex[BACnetRejectReasonTagged](BACnetRejectReasonTaggedParseWithBufferProducer((uint32)(uint32(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rejectReason' field"))
	}
	m.RejectReason = rejectReason

	if closeErr := readBuffer.CloseContext("APDUReject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUReject")
	}

	return m, nil
}

func (m *_APDUReject) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_APDUReject) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUReject"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUReject")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[uint8](ctx, "originalInvokeId", m.GetOriginalInvokeId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'originalInvokeId' field")
		}

		if err := WriteSimpleField[BACnetRejectReasonTagged](ctx, "rejectReason", m.GetRejectReason(), WriteComplex[BACnetRejectReasonTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'rejectReason' field")
		}

		if popErr := writeBuffer.PopContext("APDUReject"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUReject")
		}
		return nil
	}
	return m.APDUContract.(*_APDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_APDUReject) IsAPDUReject() {}

func (m *_APDUReject) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
