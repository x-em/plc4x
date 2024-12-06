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
	utils.Copyable
	APDU
	// GetOriginalInvokeId returns OriginalInvokeId (property field)
	GetOriginalInvokeId() uint8
	// GetRejectReason returns RejectReason (property field)
	GetRejectReason() BACnetRejectReasonTagged
	// IsAPDUReject is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAPDUReject()
	// CreateBuilder creates a APDURejectBuilder
	CreateAPDURejectBuilder() APDURejectBuilder
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
/////////////////////// Builder
///////////////////////

// APDURejectBuilder is a builder for APDUReject
type APDURejectBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(originalInvokeId uint8, rejectReason BACnetRejectReasonTagged) APDURejectBuilder
	// WithOriginalInvokeId adds OriginalInvokeId (property field)
	WithOriginalInvokeId(uint8) APDURejectBuilder
	// WithRejectReason adds RejectReason (property field)
	WithRejectReason(BACnetRejectReasonTagged) APDURejectBuilder
	// WithRejectReasonBuilder adds RejectReason (property field) which is build by the builder
	WithRejectReasonBuilder(func(BACnetRejectReasonTaggedBuilder) BACnetRejectReasonTaggedBuilder) APDURejectBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() APDUBuilder
	// Build builds the APDUReject or returns an error if something is wrong
	Build() (APDUReject, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() APDUReject
}

// NewAPDURejectBuilder() creates a APDURejectBuilder
func NewAPDURejectBuilder() APDURejectBuilder {
	return &_APDURejectBuilder{_APDUReject: new(_APDUReject)}
}

type _APDURejectBuilder struct {
	*_APDUReject

	parentBuilder *_APDUBuilder

	err *utils.MultiError
}

var _ (APDURejectBuilder) = (*_APDURejectBuilder)(nil)

func (b *_APDURejectBuilder) setParent(contract APDUContract) {
	b.APDUContract = contract
	contract.(*_APDU)._SubType = b._APDUReject
}

func (b *_APDURejectBuilder) WithMandatoryFields(originalInvokeId uint8, rejectReason BACnetRejectReasonTagged) APDURejectBuilder {
	return b.WithOriginalInvokeId(originalInvokeId).WithRejectReason(rejectReason)
}

func (b *_APDURejectBuilder) WithOriginalInvokeId(originalInvokeId uint8) APDURejectBuilder {
	b.OriginalInvokeId = originalInvokeId
	return b
}

func (b *_APDURejectBuilder) WithRejectReason(rejectReason BACnetRejectReasonTagged) APDURejectBuilder {
	b.RejectReason = rejectReason
	return b
}

func (b *_APDURejectBuilder) WithRejectReasonBuilder(builderSupplier func(BACnetRejectReasonTaggedBuilder) BACnetRejectReasonTaggedBuilder) APDURejectBuilder {
	builder := builderSupplier(b.RejectReason.CreateBACnetRejectReasonTaggedBuilder())
	var err error
	b.RejectReason, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetRejectReasonTaggedBuilder failed"))
	}
	return b
}

func (b *_APDURejectBuilder) Build() (APDUReject, error) {
	if b.RejectReason == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'rejectReason' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._APDUReject.deepCopy(), nil
}

func (b *_APDURejectBuilder) MustBuild() APDUReject {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_APDURejectBuilder) Done() APDUBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAPDUBuilder().(*_APDUBuilder)
	}
	return b.parentBuilder
}

func (b *_APDURejectBuilder) buildForAPDU() (APDU, error) {
	return b.Build()
}

func (b *_APDURejectBuilder) DeepCopy() any {
	_copy := b.CreateAPDURejectBuilder().(*_APDURejectBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAPDURejectBuilder creates a APDURejectBuilder
func (b *_APDUReject) CreateAPDURejectBuilder() APDURejectBuilder {
	if b == nil {
		return NewAPDURejectBuilder()
	}
	return &_APDURejectBuilder{_APDUReject: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

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

func (m *_APDUReject) DeepCopy() any {
	return m.deepCopy()
}

func (m *_APDUReject) deepCopy() *_APDUReject {
	if m == nil {
		return nil
	}
	_APDURejectCopy := &_APDUReject{
		m.APDUContract.(*_APDU).deepCopy(),
		m.OriginalInvokeId,
		utils.DeepCopy[BACnetRejectReasonTagged](m.RejectReason),
		m.reservedField0,
	}
	_APDURejectCopy.APDUContract.(*_APDU)._SubType = m
	return _APDURejectCopy
}

func (m *_APDUReject) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}