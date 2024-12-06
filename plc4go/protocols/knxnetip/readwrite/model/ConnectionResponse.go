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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ConnectionResponse is the corresponding interface of ConnectionResponse
type ConnectionResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	KnxNetIpMessage
	// GetCommunicationChannelId returns CommunicationChannelId (property field)
	GetCommunicationChannelId() uint8
	// GetStatus returns Status (property field)
	GetStatus() Status
	// GetHpaiDataEndpoint returns HpaiDataEndpoint (property field)
	GetHpaiDataEndpoint() HPAIDataEndpoint
	// GetConnectionResponseDataBlock returns ConnectionResponseDataBlock (property field)
	GetConnectionResponseDataBlock() ConnectionResponseDataBlock
	// IsConnectionResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsConnectionResponse()
	// CreateBuilder creates a ConnectionResponseBuilder
	CreateConnectionResponseBuilder() ConnectionResponseBuilder
}

// _ConnectionResponse is the data-structure of this message
type _ConnectionResponse struct {
	KnxNetIpMessageContract
	CommunicationChannelId      uint8
	Status                      Status
	HpaiDataEndpoint            HPAIDataEndpoint
	ConnectionResponseDataBlock ConnectionResponseDataBlock
}

var _ ConnectionResponse = (*_ConnectionResponse)(nil)
var _ KnxNetIpMessageRequirements = (*_ConnectionResponse)(nil)

// NewConnectionResponse factory function for _ConnectionResponse
func NewConnectionResponse(communicationChannelId uint8, status Status, hpaiDataEndpoint HPAIDataEndpoint, connectionResponseDataBlock ConnectionResponseDataBlock) *_ConnectionResponse {
	_result := &_ConnectionResponse{
		KnxNetIpMessageContract:     NewKnxNetIpMessage(),
		CommunicationChannelId:      communicationChannelId,
		Status:                      status,
		HpaiDataEndpoint:            hpaiDataEndpoint,
		ConnectionResponseDataBlock: connectionResponseDataBlock,
	}
	_result.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ConnectionResponseBuilder is a builder for ConnectionResponse
type ConnectionResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(communicationChannelId uint8, status Status) ConnectionResponseBuilder
	// WithCommunicationChannelId adds CommunicationChannelId (property field)
	WithCommunicationChannelId(uint8) ConnectionResponseBuilder
	// WithStatus adds Status (property field)
	WithStatus(Status) ConnectionResponseBuilder
	// WithHpaiDataEndpoint adds HpaiDataEndpoint (property field)
	WithOptionalHpaiDataEndpoint(HPAIDataEndpoint) ConnectionResponseBuilder
	// WithOptionalHpaiDataEndpointBuilder adds HpaiDataEndpoint (property field) which is build by the builder
	WithOptionalHpaiDataEndpointBuilder(func(HPAIDataEndpointBuilder) HPAIDataEndpointBuilder) ConnectionResponseBuilder
	// WithConnectionResponseDataBlock adds ConnectionResponseDataBlock (property field)
	WithOptionalConnectionResponseDataBlock(ConnectionResponseDataBlock) ConnectionResponseBuilder
	// WithOptionalConnectionResponseDataBlockBuilder adds ConnectionResponseDataBlock (property field) which is build by the builder
	WithOptionalConnectionResponseDataBlockBuilder(func(ConnectionResponseDataBlockBuilder) ConnectionResponseDataBlockBuilder) ConnectionResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() KnxNetIpMessageBuilder
	// Build builds the ConnectionResponse or returns an error if something is wrong
	Build() (ConnectionResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ConnectionResponse
}

// NewConnectionResponseBuilder() creates a ConnectionResponseBuilder
func NewConnectionResponseBuilder() ConnectionResponseBuilder {
	return &_ConnectionResponseBuilder{_ConnectionResponse: new(_ConnectionResponse)}
}

type _ConnectionResponseBuilder struct {
	*_ConnectionResponse

	parentBuilder *_KnxNetIpMessageBuilder

	err *utils.MultiError
}

var _ (ConnectionResponseBuilder) = (*_ConnectionResponseBuilder)(nil)

func (b *_ConnectionResponseBuilder) setParent(contract KnxNetIpMessageContract) {
	b.KnxNetIpMessageContract = contract
	contract.(*_KnxNetIpMessage)._SubType = b._ConnectionResponse
}

func (b *_ConnectionResponseBuilder) WithMandatoryFields(communicationChannelId uint8, status Status) ConnectionResponseBuilder {
	return b.WithCommunicationChannelId(communicationChannelId).WithStatus(status)
}

func (b *_ConnectionResponseBuilder) WithCommunicationChannelId(communicationChannelId uint8) ConnectionResponseBuilder {
	b.CommunicationChannelId = communicationChannelId
	return b
}

func (b *_ConnectionResponseBuilder) WithStatus(status Status) ConnectionResponseBuilder {
	b.Status = status
	return b
}

func (b *_ConnectionResponseBuilder) WithOptionalHpaiDataEndpoint(hpaiDataEndpoint HPAIDataEndpoint) ConnectionResponseBuilder {
	b.HpaiDataEndpoint = hpaiDataEndpoint
	return b
}

func (b *_ConnectionResponseBuilder) WithOptionalHpaiDataEndpointBuilder(builderSupplier func(HPAIDataEndpointBuilder) HPAIDataEndpointBuilder) ConnectionResponseBuilder {
	builder := builderSupplier(b.HpaiDataEndpoint.CreateHPAIDataEndpointBuilder())
	var err error
	b.HpaiDataEndpoint, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HPAIDataEndpointBuilder failed"))
	}
	return b
}

func (b *_ConnectionResponseBuilder) WithOptionalConnectionResponseDataBlock(connectionResponseDataBlock ConnectionResponseDataBlock) ConnectionResponseBuilder {
	b.ConnectionResponseDataBlock = connectionResponseDataBlock
	return b
}

func (b *_ConnectionResponseBuilder) WithOptionalConnectionResponseDataBlockBuilder(builderSupplier func(ConnectionResponseDataBlockBuilder) ConnectionResponseDataBlockBuilder) ConnectionResponseBuilder {
	builder := builderSupplier(b.ConnectionResponseDataBlock.CreateConnectionResponseDataBlockBuilder())
	var err error
	b.ConnectionResponseDataBlock, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ConnectionResponseDataBlockBuilder failed"))
	}
	return b
}

func (b *_ConnectionResponseBuilder) Build() (ConnectionResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ConnectionResponse.deepCopy(), nil
}

func (b *_ConnectionResponseBuilder) MustBuild() ConnectionResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ConnectionResponseBuilder) Done() KnxNetIpMessageBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewKnxNetIpMessageBuilder().(*_KnxNetIpMessageBuilder)
	}
	return b.parentBuilder
}

func (b *_ConnectionResponseBuilder) buildForKnxNetIpMessage() (KnxNetIpMessage, error) {
	return b.Build()
}

func (b *_ConnectionResponseBuilder) DeepCopy() any {
	_copy := b.CreateConnectionResponseBuilder().(*_ConnectionResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateConnectionResponseBuilder creates a ConnectionResponseBuilder
func (b *_ConnectionResponse) CreateConnectionResponseBuilder() ConnectionResponseBuilder {
	if b == nil {
		return NewConnectionResponseBuilder()
	}
	return &_ConnectionResponseBuilder{_ConnectionResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConnectionResponse) GetMsgType() uint16 {
	return 0x0206
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConnectionResponse) GetParent() KnxNetIpMessageContract {
	return m.KnxNetIpMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ConnectionResponse) GetCommunicationChannelId() uint8 {
	return m.CommunicationChannelId
}

func (m *_ConnectionResponse) GetStatus() Status {
	return m.Status
}

func (m *_ConnectionResponse) GetHpaiDataEndpoint() HPAIDataEndpoint {
	return m.HpaiDataEndpoint
}

func (m *_ConnectionResponse) GetConnectionResponseDataBlock() ConnectionResponseDataBlock {
	return m.ConnectionResponseDataBlock
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastConnectionResponse(structType any) ConnectionResponse {
	if casted, ok := structType.(ConnectionResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectionResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectionResponse) GetTypeName() string {
	return "ConnectionResponse"
}

func (m *_ConnectionResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.KnxNetIpMessageContract.(*_KnxNetIpMessage).getLengthInBits(ctx))

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Optional Field (hpaiDataEndpoint)
	if m.HpaiDataEndpoint != nil {
		lengthInBits += m.HpaiDataEndpoint.GetLengthInBits(ctx)
	}

	// Optional Field (connectionResponseDataBlock)
	if m.ConnectionResponseDataBlock != nil {
		lengthInBits += m.ConnectionResponseDataBlock.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_ConnectionResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ConnectionResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_KnxNetIpMessage) (__connectionResponse ConnectionResponse, err error) {
	m.KnxNetIpMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectionResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	communicationChannelId, err := ReadSimpleField(ctx, "communicationChannelId", ReadUnsignedByte(readBuffer, uint8(8)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'communicationChannelId' field"))
	}
	m.CommunicationChannelId = communicationChannelId

	status, err := ReadEnumField[Status](ctx, "status", "Status", ReadEnum(StatusByValue, ReadUnsignedByte(readBuffer, uint8(8))), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	var hpaiDataEndpoint HPAIDataEndpoint
	_hpaiDataEndpoint, err := ReadOptionalField[HPAIDataEndpoint](ctx, "hpaiDataEndpoint", ReadComplex[HPAIDataEndpoint](HPAIDataEndpointParseWithBuffer, readBuffer), bool((status) == (Status_NO_ERROR)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hpaiDataEndpoint' field"))
	}
	if _hpaiDataEndpoint != nil {
		hpaiDataEndpoint = *_hpaiDataEndpoint
		m.HpaiDataEndpoint = hpaiDataEndpoint
	}

	var connectionResponseDataBlock ConnectionResponseDataBlock
	_connectionResponseDataBlock, err := ReadOptionalField[ConnectionResponseDataBlock](ctx, "connectionResponseDataBlock", ReadComplex[ConnectionResponseDataBlock](ConnectionResponseDataBlockParseWithBuffer, readBuffer), bool((status) == (Status_NO_ERROR)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionResponseDataBlock' field"))
	}
	if _connectionResponseDataBlock != nil {
		connectionResponseDataBlock = *_connectionResponseDataBlock
		m.ConnectionResponseDataBlock = connectionResponseDataBlock
	}

	if closeErr := readBuffer.CloseContext("ConnectionResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionResponse")
	}

	return m, nil
}

func (m *_ConnectionResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConnectionResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectionResponse")
		}

		if err := WriteSimpleField[uint8](ctx, "communicationChannelId", m.GetCommunicationChannelId(), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'communicationChannelId' field")
		}

		if err := WriteSimpleEnumField[Status](ctx, "status", "Status", m.GetStatus(), WriteEnum[Status, uint8](Status.GetValue, Status.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8)), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'status' field")
		}

		if err := WriteOptionalField[HPAIDataEndpoint](ctx, "hpaiDataEndpoint", GetRef(m.GetHpaiDataEndpoint()), WriteComplex[HPAIDataEndpoint](writeBuffer), true, codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'hpaiDataEndpoint' field")
		}

		if err := WriteOptionalField[ConnectionResponseDataBlock](ctx, "connectionResponseDataBlock", GetRef(m.GetConnectionResponseDataBlock()), WriteComplex[ConnectionResponseDataBlock](writeBuffer), true, codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'connectionResponseDataBlock' field")
		}

		if popErr := writeBuffer.PopContext("ConnectionResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectionResponse")
		}
		return nil
	}
	return m.KnxNetIpMessageContract.(*_KnxNetIpMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConnectionResponse) IsConnectionResponse() {}

func (m *_ConnectionResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ConnectionResponse) deepCopy() *_ConnectionResponse {
	if m == nil {
		return nil
	}
	_ConnectionResponseCopy := &_ConnectionResponse{
		m.KnxNetIpMessageContract.(*_KnxNetIpMessage).deepCopy(),
		m.CommunicationChannelId,
		m.Status,
		utils.DeepCopy[HPAIDataEndpoint](m.HpaiDataEndpoint),
		utils.DeepCopy[ConnectionResponseDataBlock](m.ConnectionResponseDataBlock),
	}
	_ConnectionResponseCopy.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = m
	return _ConnectionResponseCopy
}

func (m *_ConnectionResponse) String() string {
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