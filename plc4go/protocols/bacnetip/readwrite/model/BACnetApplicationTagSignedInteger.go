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

// BACnetApplicationTagSignedInteger is the corresponding interface of BACnetApplicationTagSignedInteger
type BACnetApplicationTagSignedInteger interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() uint64
	// IsBACnetApplicationTagSignedInteger is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetApplicationTagSignedInteger()
	// CreateBuilder creates a BACnetApplicationTagSignedIntegerBuilder
	CreateBACnetApplicationTagSignedIntegerBuilder() BACnetApplicationTagSignedIntegerBuilder
}

// _BACnetApplicationTagSignedInteger is the data-structure of this message
type _BACnetApplicationTagSignedInteger struct {
	BACnetApplicationTagContract
	Payload BACnetTagPayloadSignedInteger
}

var _ BACnetApplicationTagSignedInteger = (*_BACnetApplicationTagSignedInteger)(nil)
var _ BACnetApplicationTagRequirements = (*_BACnetApplicationTagSignedInteger)(nil)

// NewBACnetApplicationTagSignedInteger factory function for _BACnetApplicationTagSignedInteger
func NewBACnetApplicationTagSignedInteger(header BACnetTagHeader, payload BACnetTagPayloadSignedInteger) *_BACnetApplicationTagSignedInteger {
	if payload == nil {
		panic("payload of type BACnetTagPayloadSignedInteger for BACnetApplicationTagSignedInteger must not be nil")
	}
	_result := &_BACnetApplicationTagSignedInteger{
		BACnetApplicationTagContract: NewBACnetApplicationTag(header),
		Payload:                      payload,
	}
	_result.BACnetApplicationTagContract.(*_BACnetApplicationTag)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetApplicationTagSignedIntegerBuilder is a builder for BACnetApplicationTagSignedInteger
type BACnetApplicationTagSignedIntegerBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(payload BACnetTagPayloadSignedInteger) BACnetApplicationTagSignedIntegerBuilder
	// WithPayload adds Payload (property field)
	WithPayload(BACnetTagPayloadSignedInteger) BACnetApplicationTagSignedIntegerBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(BACnetTagPayloadSignedIntegerBuilder) BACnetTagPayloadSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetApplicationTagBuilder
	// Build builds the BACnetApplicationTagSignedInteger or returns an error if something is wrong
	Build() (BACnetApplicationTagSignedInteger, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetApplicationTagSignedInteger
}

// NewBACnetApplicationTagSignedIntegerBuilder() creates a BACnetApplicationTagSignedIntegerBuilder
func NewBACnetApplicationTagSignedIntegerBuilder() BACnetApplicationTagSignedIntegerBuilder {
	return &_BACnetApplicationTagSignedIntegerBuilder{_BACnetApplicationTagSignedInteger: new(_BACnetApplicationTagSignedInteger)}
}

type _BACnetApplicationTagSignedIntegerBuilder struct {
	*_BACnetApplicationTagSignedInteger

	parentBuilder *_BACnetApplicationTagBuilder

	err *utils.MultiError
}

var _ (BACnetApplicationTagSignedIntegerBuilder) = (*_BACnetApplicationTagSignedIntegerBuilder)(nil)

func (b *_BACnetApplicationTagSignedIntegerBuilder) setParent(contract BACnetApplicationTagContract) {
	b.BACnetApplicationTagContract = contract
	contract.(*_BACnetApplicationTag)._SubType = b._BACnetApplicationTagSignedInteger
}

func (b *_BACnetApplicationTagSignedIntegerBuilder) WithMandatoryFields(payload BACnetTagPayloadSignedInteger) BACnetApplicationTagSignedIntegerBuilder {
	return b.WithPayload(payload)
}

func (b *_BACnetApplicationTagSignedIntegerBuilder) WithPayload(payload BACnetTagPayloadSignedInteger) BACnetApplicationTagSignedIntegerBuilder {
	b.Payload = payload
	return b
}

func (b *_BACnetApplicationTagSignedIntegerBuilder) WithPayloadBuilder(builderSupplier func(BACnetTagPayloadSignedIntegerBuilder) BACnetTagPayloadSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder {
	builder := builderSupplier(b.Payload.CreateBACnetTagPayloadSignedIntegerBuilder())
	var err error
	b.Payload, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagPayloadSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetApplicationTagSignedIntegerBuilder) Build() (BACnetApplicationTagSignedInteger, error) {
	if b.Payload == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetApplicationTagSignedInteger.deepCopy(), nil
}

func (b *_BACnetApplicationTagSignedIntegerBuilder) MustBuild() BACnetApplicationTagSignedInteger {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetApplicationTagSignedIntegerBuilder) Done() BACnetApplicationTagBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetApplicationTagBuilder().(*_BACnetApplicationTagBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetApplicationTagSignedIntegerBuilder) buildForBACnetApplicationTag() (BACnetApplicationTag, error) {
	return b.Build()
}

func (b *_BACnetApplicationTagSignedIntegerBuilder) DeepCopy() any {
	_copy := b.CreateBACnetApplicationTagSignedIntegerBuilder().(*_BACnetApplicationTagSignedIntegerBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetApplicationTagSignedIntegerBuilder creates a BACnetApplicationTagSignedIntegerBuilder
func (b *_BACnetApplicationTagSignedInteger) CreateBACnetApplicationTagSignedIntegerBuilder() BACnetApplicationTagSignedIntegerBuilder {
	if b == nil {
		return NewBACnetApplicationTagSignedIntegerBuilder()
	}
	return &_BACnetApplicationTagSignedIntegerBuilder{_BACnetApplicationTagSignedInteger: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetApplicationTagSignedInteger) GetParent() BACnetApplicationTagContract {
	return m.BACnetApplicationTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetApplicationTagSignedInteger) GetPayload() BACnetTagPayloadSignedInteger {
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

func (m *_BACnetApplicationTagSignedInteger) GetActualValue() uint64 {
	ctx := context.Background()
	_ = ctx
	return uint64(m.GetPayload().GetActualValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetApplicationTagSignedInteger(structType any) BACnetApplicationTagSignedInteger {
	if casted, ok := structType.(BACnetApplicationTagSignedInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetApplicationTagSignedInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetApplicationTagSignedInteger) GetTypeName() string {
	return "BACnetApplicationTagSignedInteger"
}

func (m *_BACnetApplicationTagSignedInteger) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetApplicationTagContract.(*_BACnetApplicationTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetApplicationTagSignedInteger) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetApplicationTagSignedInteger) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetApplicationTag, header BACnetTagHeader) (__bACnetApplicationTagSignedInteger BACnetApplicationTagSignedInteger, err error) {
	m.BACnetApplicationTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagSignedInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagSignedInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadSignedInteger](ctx, "payload", ReadComplex[BACnetTagPayloadSignedInteger](BACnetTagPayloadSignedIntegerParseWithBufferProducer((uint32)(header.GetActualLength())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	actualValue, err := ReadVirtualField[uint64](ctx, "actualValue", (*uint64)(nil), payload.GetActualValue())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagSignedInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagSignedInteger")
	}

	return m, nil
}

func (m *_BACnetApplicationTagSignedInteger) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetApplicationTagSignedInteger) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagSignedInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagSignedInteger")
		}

		if err := WriteSimpleField[BACnetTagPayloadSignedInteger](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagSignedInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagSignedInteger")
		}
		return nil
	}
	return m.BACnetApplicationTagContract.(*_BACnetApplicationTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetApplicationTagSignedInteger) IsBACnetApplicationTagSignedInteger() {}

func (m *_BACnetApplicationTagSignedInteger) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetApplicationTagSignedInteger) deepCopy() *_BACnetApplicationTagSignedInteger {
	if m == nil {
		return nil
	}
	_BACnetApplicationTagSignedIntegerCopy := &_BACnetApplicationTagSignedInteger{
		m.BACnetApplicationTagContract.(*_BACnetApplicationTag).deepCopy(),
		utils.DeepCopy[BACnetTagPayloadSignedInteger](m.Payload),
	}
	_BACnetApplicationTagSignedIntegerCopy.BACnetApplicationTagContract.(*_BACnetApplicationTag)._SubType = m
	return _BACnetApplicationTagSignedIntegerCopy
}

func (m *_BACnetApplicationTagSignedInteger) String() string {
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