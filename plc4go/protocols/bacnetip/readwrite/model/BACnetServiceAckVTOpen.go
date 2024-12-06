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

// BACnetServiceAckVTOpen is the corresponding interface of BACnetServiceAckVTOpen
type BACnetServiceAckVTOpen interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetServiceAck
	// GetRemoteVtSessionIdentifier returns RemoteVtSessionIdentifier (property field)
	GetRemoteVtSessionIdentifier() BACnetApplicationTagUnsignedInteger
	// IsBACnetServiceAckVTOpen is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetServiceAckVTOpen()
	// CreateBuilder creates a BACnetServiceAckVTOpenBuilder
	CreateBACnetServiceAckVTOpenBuilder() BACnetServiceAckVTOpenBuilder
}

// _BACnetServiceAckVTOpen is the data-structure of this message
type _BACnetServiceAckVTOpen struct {
	BACnetServiceAckContract
	RemoteVtSessionIdentifier BACnetApplicationTagUnsignedInteger
}

var _ BACnetServiceAckVTOpen = (*_BACnetServiceAckVTOpen)(nil)
var _ BACnetServiceAckRequirements = (*_BACnetServiceAckVTOpen)(nil)

// NewBACnetServiceAckVTOpen factory function for _BACnetServiceAckVTOpen
func NewBACnetServiceAckVTOpen(remoteVtSessionIdentifier BACnetApplicationTagUnsignedInteger, serviceAckLength uint32) *_BACnetServiceAckVTOpen {
	if remoteVtSessionIdentifier == nil {
		panic("remoteVtSessionIdentifier of type BACnetApplicationTagUnsignedInteger for BACnetServiceAckVTOpen must not be nil")
	}
	_result := &_BACnetServiceAckVTOpen{
		BACnetServiceAckContract:  NewBACnetServiceAck(serviceAckLength),
		RemoteVtSessionIdentifier: remoteVtSessionIdentifier,
	}
	_result.BACnetServiceAckContract.(*_BACnetServiceAck)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetServiceAckVTOpenBuilder is a builder for BACnetServiceAckVTOpen
type BACnetServiceAckVTOpenBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(remoteVtSessionIdentifier BACnetApplicationTagUnsignedInteger) BACnetServiceAckVTOpenBuilder
	// WithRemoteVtSessionIdentifier adds RemoteVtSessionIdentifier (property field)
	WithRemoteVtSessionIdentifier(BACnetApplicationTagUnsignedInteger) BACnetServiceAckVTOpenBuilder
	// WithRemoteVtSessionIdentifierBuilder adds RemoteVtSessionIdentifier (property field) which is build by the builder
	WithRemoteVtSessionIdentifierBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetServiceAckVTOpenBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetServiceAckBuilder
	// Build builds the BACnetServiceAckVTOpen or returns an error if something is wrong
	Build() (BACnetServiceAckVTOpen, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetServiceAckVTOpen
}

// NewBACnetServiceAckVTOpenBuilder() creates a BACnetServiceAckVTOpenBuilder
func NewBACnetServiceAckVTOpenBuilder() BACnetServiceAckVTOpenBuilder {
	return &_BACnetServiceAckVTOpenBuilder{_BACnetServiceAckVTOpen: new(_BACnetServiceAckVTOpen)}
}

type _BACnetServiceAckVTOpenBuilder struct {
	*_BACnetServiceAckVTOpen

	parentBuilder *_BACnetServiceAckBuilder

	err *utils.MultiError
}

var _ (BACnetServiceAckVTOpenBuilder) = (*_BACnetServiceAckVTOpenBuilder)(nil)

func (b *_BACnetServiceAckVTOpenBuilder) setParent(contract BACnetServiceAckContract) {
	b.BACnetServiceAckContract = contract
	contract.(*_BACnetServiceAck)._SubType = b._BACnetServiceAckVTOpen
}

func (b *_BACnetServiceAckVTOpenBuilder) WithMandatoryFields(remoteVtSessionIdentifier BACnetApplicationTagUnsignedInteger) BACnetServiceAckVTOpenBuilder {
	return b.WithRemoteVtSessionIdentifier(remoteVtSessionIdentifier)
}

func (b *_BACnetServiceAckVTOpenBuilder) WithRemoteVtSessionIdentifier(remoteVtSessionIdentifier BACnetApplicationTagUnsignedInteger) BACnetServiceAckVTOpenBuilder {
	b.RemoteVtSessionIdentifier = remoteVtSessionIdentifier
	return b
}

func (b *_BACnetServiceAckVTOpenBuilder) WithRemoteVtSessionIdentifierBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetServiceAckVTOpenBuilder {
	builder := builderSupplier(b.RemoteVtSessionIdentifier.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.RemoteVtSessionIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetServiceAckVTOpenBuilder) Build() (BACnetServiceAckVTOpen, error) {
	if b.RemoteVtSessionIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'remoteVtSessionIdentifier' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetServiceAckVTOpen.deepCopy(), nil
}

func (b *_BACnetServiceAckVTOpenBuilder) MustBuild() BACnetServiceAckVTOpen {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetServiceAckVTOpenBuilder) Done() BACnetServiceAckBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetServiceAckBuilder().(*_BACnetServiceAckBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetServiceAckVTOpenBuilder) buildForBACnetServiceAck() (BACnetServiceAck, error) {
	return b.Build()
}

func (b *_BACnetServiceAckVTOpenBuilder) DeepCopy() any {
	_copy := b.CreateBACnetServiceAckVTOpenBuilder().(*_BACnetServiceAckVTOpenBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetServiceAckVTOpenBuilder creates a BACnetServiceAckVTOpenBuilder
func (b *_BACnetServiceAckVTOpen) CreateBACnetServiceAckVTOpenBuilder() BACnetServiceAckVTOpenBuilder {
	if b == nil {
		return NewBACnetServiceAckVTOpenBuilder()
	}
	return &_BACnetServiceAckVTOpenBuilder{_BACnetServiceAckVTOpen: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckVTOpen) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_VT_OPEN
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckVTOpen) GetParent() BACnetServiceAckContract {
	return m.BACnetServiceAckContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckVTOpen) GetRemoteVtSessionIdentifier() BACnetApplicationTagUnsignedInteger {
	return m.RemoteVtSessionIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckVTOpen(structType any) BACnetServiceAckVTOpen {
	if casted, ok := structType.(BACnetServiceAckVTOpen); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckVTOpen); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckVTOpen) GetTypeName() string {
	return "BACnetServiceAckVTOpen"
}

func (m *_BACnetServiceAckVTOpen) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetServiceAckContract.(*_BACnetServiceAck).getLengthInBits(ctx))

	// Simple field (remoteVtSessionIdentifier)
	lengthInBits += m.RemoteVtSessionIdentifier.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetServiceAckVTOpen) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetServiceAckVTOpen) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetServiceAck, serviceAckLength uint32) (__bACnetServiceAckVTOpen BACnetServiceAckVTOpen, err error) {
	m.BACnetServiceAckContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckVTOpen"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckVTOpen")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	remoteVtSessionIdentifier, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "remoteVtSessionIdentifier", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'remoteVtSessionIdentifier' field"))
	}
	m.RemoteVtSessionIdentifier = remoteVtSessionIdentifier

	if closeErr := readBuffer.CloseContext("BACnetServiceAckVTOpen"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckVTOpen")
	}

	return m, nil
}

func (m *_BACnetServiceAckVTOpen) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckVTOpen) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckVTOpen"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckVTOpen")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "remoteVtSessionIdentifier", m.GetRemoteVtSessionIdentifier(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'remoteVtSessionIdentifier' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckVTOpen"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckVTOpen")
		}
		return nil
	}
	return m.BACnetServiceAckContract.(*_BACnetServiceAck).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetServiceAckVTOpen) IsBACnetServiceAckVTOpen() {}

func (m *_BACnetServiceAckVTOpen) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetServiceAckVTOpen) deepCopy() *_BACnetServiceAckVTOpen {
	if m == nil {
		return nil
	}
	_BACnetServiceAckVTOpenCopy := &_BACnetServiceAckVTOpen{
		m.BACnetServiceAckContract.(*_BACnetServiceAck).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.RemoteVtSessionIdentifier),
	}
	_BACnetServiceAckVTOpenCopy.BACnetServiceAckContract.(*_BACnetServiceAck)._SubType = m
	return _BACnetServiceAckVTOpenCopy
}

func (m *_BACnetServiceAckVTOpen) String() string {
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