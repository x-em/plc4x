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

// Constant values.
const EipConnectionRequest_PROTOCOLVERSION uint16 = 0x01
const EipConnectionRequest_FLAGS uint16 = 0x00

// EipConnectionRequest is the corresponding interface of EipConnectionRequest
type EipConnectionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	EipPacket
	// IsEipConnectionRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsEipConnectionRequest()
	// CreateBuilder creates a EipConnectionRequestBuilder
	CreateEipConnectionRequestBuilder() EipConnectionRequestBuilder
}

// _EipConnectionRequest is the data-structure of this message
type _EipConnectionRequest struct {
	EipPacketContract
}

var _ EipConnectionRequest = (*_EipConnectionRequest)(nil)
var _ EipPacketRequirements = (*_EipConnectionRequest)(nil)

// NewEipConnectionRequest factory function for _EipConnectionRequest
func NewEipConnectionRequest(sessionHandle uint32, status uint32, senderContext []byte, options uint32) *_EipConnectionRequest {
	_result := &_EipConnectionRequest{
		EipPacketContract: NewEipPacket(sessionHandle, status, senderContext, options),
	}
	_result.EipPacketContract.(*_EipPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// EipConnectionRequestBuilder is a builder for EipConnectionRequest
type EipConnectionRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() EipConnectionRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() EipPacketBuilder
	// Build builds the EipConnectionRequest or returns an error if something is wrong
	Build() (EipConnectionRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() EipConnectionRequest
}

// NewEipConnectionRequestBuilder() creates a EipConnectionRequestBuilder
func NewEipConnectionRequestBuilder() EipConnectionRequestBuilder {
	return &_EipConnectionRequestBuilder{_EipConnectionRequest: new(_EipConnectionRequest)}
}

type _EipConnectionRequestBuilder struct {
	*_EipConnectionRequest

	parentBuilder *_EipPacketBuilder

	err *utils.MultiError
}

var _ (EipConnectionRequestBuilder) = (*_EipConnectionRequestBuilder)(nil)

func (b *_EipConnectionRequestBuilder) setParent(contract EipPacketContract) {
	b.EipPacketContract = contract
	contract.(*_EipPacket)._SubType = b._EipConnectionRequest
}

func (b *_EipConnectionRequestBuilder) WithMandatoryFields() EipConnectionRequestBuilder {
	return b
}

func (b *_EipConnectionRequestBuilder) Build() (EipConnectionRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._EipConnectionRequest.deepCopy(), nil
}

func (b *_EipConnectionRequestBuilder) MustBuild() EipConnectionRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_EipConnectionRequestBuilder) Done() EipPacketBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewEipPacketBuilder().(*_EipPacketBuilder)
	}
	return b.parentBuilder
}

func (b *_EipConnectionRequestBuilder) buildForEipPacket() (EipPacket, error) {
	return b.Build()
}

func (b *_EipConnectionRequestBuilder) DeepCopy() any {
	_copy := b.CreateEipConnectionRequestBuilder().(*_EipConnectionRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateEipConnectionRequestBuilder creates a EipConnectionRequestBuilder
func (b *_EipConnectionRequest) CreateEipConnectionRequestBuilder() EipConnectionRequestBuilder {
	if b == nil {
		return NewEipConnectionRequestBuilder()
	}
	return &_EipConnectionRequestBuilder{_EipConnectionRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EipConnectionRequest) GetCommand() uint16 {
	return 0x0065
}

func (m *_EipConnectionRequest) GetResponse() bool {
	return bool(false)
}

func (m *_EipConnectionRequest) GetPacketLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EipConnectionRequest) GetParent() EipPacketContract {
	return m.EipPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_EipConnectionRequest) GetProtocolVersion() uint16 {
	return EipConnectionRequest_PROTOCOLVERSION
}

func (m *_EipConnectionRequest) GetFlags() uint16 {
	return EipConnectionRequest_FLAGS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastEipConnectionRequest(structType any) EipConnectionRequest {
	if casted, ok := structType.(EipConnectionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*EipConnectionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_EipConnectionRequest) GetTypeName() string {
	return "EipConnectionRequest"
}

func (m *_EipConnectionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.EipPacketContract.(*_EipPacket).getLengthInBits(ctx))

	// Const Field (protocolVersion)
	lengthInBits += 16

	// Const Field (flags)
	lengthInBits += 16

	return lengthInBits
}

func (m *_EipConnectionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_EipConnectionRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_EipPacket, response bool) (__eipConnectionRequest EipConnectionRequest, err error) {
	m.EipPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EipConnectionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EipConnectionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	protocolVersion, err := ReadConstField[uint16](ctx, "protocolVersion", ReadUnsignedShort(readBuffer, uint8(16)), EipConnectionRequest_PROTOCOLVERSION)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolVersion' field"))
	}
	_ = protocolVersion

	flags, err := ReadConstField[uint16](ctx, "flags", ReadUnsignedShort(readBuffer, uint8(16)), EipConnectionRequest_FLAGS)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'flags' field"))
	}
	_ = flags

	if closeErr := readBuffer.CloseContext("EipConnectionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EipConnectionRequest")
	}

	return m, nil
}

func (m *_EipConnectionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EipConnectionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EipConnectionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EipConnectionRequest")
		}

		if err := WriteConstField(ctx, "protocolVersion", EipConnectionRequest_PROTOCOLVERSION, WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'protocolVersion' field")
		}

		if err := WriteConstField(ctx, "flags", EipConnectionRequest_FLAGS, WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'flags' field")
		}

		if popErr := writeBuffer.PopContext("EipConnectionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EipConnectionRequest")
		}
		return nil
	}
	return m.EipPacketContract.(*_EipPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EipConnectionRequest) IsEipConnectionRequest() {}

func (m *_EipConnectionRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_EipConnectionRequest) deepCopy() *_EipConnectionRequest {
	if m == nil {
		return nil
	}
	_EipConnectionRequestCopy := &_EipConnectionRequest{
		m.EipPacketContract.(*_EipPacket).deepCopy(),
	}
	_EipConnectionRequestCopy.EipPacketContract.(*_EipPacket)._SubType = m
	return _EipConnectionRequestCopy
}

func (m *_EipConnectionRequest) String() string {
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