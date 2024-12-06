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

// ConnectionResponseDataBlockTunnelConnection is the corresponding interface of ConnectionResponseDataBlockTunnelConnection
type ConnectionResponseDataBlockTunnelConnection interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ConnectionResponseDataBlock
	// GetKnxAddress returns KnxAddress (property field)
	GetKnxAddress() KnxAddress
	// IsConnectionResponseDataBlockTunnelConnection is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsConnectionResponseDataBlockTunnelConnection()
	// CreateBuilder creates a ConnectionResponseDataBlockTunnelConnectionBuilder
	CreateConnectionResponseDataBlockTunnelConnectionBuilder() ConnectionResponseDataBlockTunnelConnectionBuilder
}

// _ConnectionResponseDataBlockTunnelConnection is the data-structure of this message
type _ConnectionResponseDataBlockTunnelConnection struct {
	ConnectionResponseDataBlockContract
	KnxAddress KnxAddress
}

var _ ConnectionResponseDataBlockTunnelConnection = (*_ConnectionResponseDataBlockTunnelConnection)(nil)
var _ ConnectionResponseDataBlockRequirements = (*_ConnectionResponseDataBlockTunnelConnection)(nil)

// NewConnectionResponseDataBlockTunnelConnection factory function for _ConnectionResponseDataBlockTunnelConnection
func NewConnectionResponseDataBlockTunnelConnection(knxAddress KnxAddress) *_ConnectionResponseDataBlockTunnelConnection {
	if knxAddress == nil {
		panic("knxAddress of type KnxAddress for ConnectionResponseDataBlockTunnelConnection must not be nil")
	}
	_result := &_ConnectionResponseDataBlockTunnelConnection{
		ConnectionResponseDataBlockContract: NewConnectionResponseDataBlock(),
		KnxAddress:                          knxAddress,
	}
	_result.ConnectionResponseDataBlockContract.(*_ConnectionResponseDataBlock)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ConnectionResponseDataBlockTunnelConnectionBuilder is a builder for ConnectionResponseDataBlockTunnelConnection
type ConnectionResponseDataBlockTunnelConnectionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(knxAddress KnxAddress) ConnectionResponseDataBlockTunnelConnectionBuilder
	// WithKnxAddress adds KnxAddress (property field)
	WithKnxAddress(KnxAddress) ConnectionResponseDataBlockTunnelConnectionBuilder
	// WithKnxAddressBuilder adds KnxAddress (property field) which is build by the builder
	WithKnxAddressBuilder(func(KnxAddressBuilder) KnxAddressBuilder) ConnectionResponseDataBlockTunnelConnectionBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ConnectionResponseDataBlockBuilder
	// Build builds the ConnectionResponseDataBlockTunnelConnection or returns an error if something is wrong
	Build() (ConnectionResponseDataBlockTunnelConnection, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ConnectionResponseDataBlockTunnelConnection
}

// NewConnectionResponseDataBlockTunnelConnectionBuilder() creates a ConnectionResponseDataBlockTunnelConnectionBuilder
func NewConnectionResponseDataBlockTunnelConnectionBuilder() ConnectionResponseDataBlockTunnelConnectionBuilder {
	return &_ConnectionResponseDataBlockTunnelConnectionBuilder{_ConnectionResponseDataBlockTunnelConnection: new(_ConnectionResponseDataBlockTunnelConnection)}
}

type _ConnectionResponseDataBlockTunnelConnectionBuilder struct {
	*_ConnectionResponseDataBlockTunnelConnection

	parentBuilder *_ConnectionResponseDataBlockBuilder

	err *utils.MultiError
}

var _ (ConnectionResponseDataBlockTunnelConnectionBuilder) = (*_ConnectionResponseDataBlockTunnelConnectionBuilder)(nil)

func (b *_ConnectionResponseDataBlockTunnelConnectionBuilder) setParent(contract ConnectionResponseDataBlockContract) {
	b.ConnectionResponseDataBlockContract = contract
	contract.(*_ConnectionResponseDataBlock)._SubType = b._ConnectionResponseDataBlockTunnelConnection
}

func (b *_ConnectionResponseDataBlockTunnelConnectionBuilder) WithMandatoryFields(knxAddress KnxAddress) ConnectionResponseDataBlockTunnelConnectionBuilder {
	return b.WithKnxAddress(knxAddress)
}

func (b *_ConnectionResponseDataBlockTunnelConnectionBuilder) WithKnxAddress(knxAddress KnxAddress) ConnectionResponseDataBlockTunnelConnectionBuilder {
	b.KnxAddress = knxAddress
	return b
}

func (b *_ConnectionResponseDataBlockTunnelConnectionBuilder) WithKnxAddressBuilder(builderSupplier func(KnxAddressBuilder) KnxAddressBuilder) ConnectionResponseDataBlockTunnelConnectionBuilder {
	builder := builderSupplier(b.KnxAddress.CreateKnxAddressBuilder())
	var err error
	b.KnxAddress, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "KnxAddressBuilder failed"))
	}
	return b
}

func (b *_ConnectionResponseDataBlockTunnelConnectionBuilder) Build() (ConnectionResponseDataBlockTunnelConnection, error) {
	if b.KnxAddress == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'knxAddress' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ConnectionResponseDataBlockTunnelConnection.deepCopy(), nil
}

func (b *_ConnectionResponseDataBlockTunnelConnectionBuilder) MustBuild() ConnectionResponseDataBlockTunnelConnection {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ConnectionResponseDataBlockTunnelConnectionBuilder) Done() ConnectionResponseDataBlockBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewConnectionResponseDataBlockBuilder().(*_ConnectionResponseDataBlockBuilder)
	}
	return b.parentBuilder
}

func (b *_ConnectionResponseDataBlockTunnelConnectionBuilder) buildForConnectionResponseDataBlock() (ConnectionResponseDataBlock, error) {
	return b.Build()
}

func (b *_ConnectionResponseDataBlockTunnelConnectionBuilder) DeepCopy() any {
	_copy := b.CreateConnectionResponseDataBlockTunnelConnectionBuilder().(*_ConnectionResponseDataBlockTunnelConnectionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateConnectionResponseDataBlockTunnelConnectionBuilder creates a ConnectionResponseDataBlockTunnelConnectionBuilder
func (b *_ConnectionResponseDataBlockTunnelConnection) CreateConnectionResponseDataBlockTunnelConnectionBuilder() ConnectionResponseDataBlockTunnelConnectionBuilder {
	if b == nil {
		return NewConnectionResponseDataBlockTunnelConnectionBuilder()
	}
	return &_ConnectionResponseDataBlockTunnelConnectionBuilder{_ConnectionResponseDataBlockTunnelConnection: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConnectionResponseDataBlockTunnelConnection) GetConnectionType() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConnectionResponseDataBlockTunnelConnection) GetParent() ConnectionResponseDataBlockContract {
	return m.ConnectionResponseDataBlockContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ConnectionResponseDataBlockTunnelConnection) GetKnxAddress() KnxAddress {
	return m.KnxAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastConnectionResponseDataBlockTunnelConnection(structType any) ConnectionResponseDataBlockTunnelConnection {
	if casted, ok := structType.(ConnectionResponseDataBlockTunnelConnection); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectionResponseDataBlockTunnelConnection); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectionResponseDataBlockTunnelConnection) GetTypeName() string {
	return "ConnectionResponseDataBlockTunnelConnection"
}

func (m *_ConnectionResponseDataBlockTunnelConnection) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ConnectionResponseDataBlockContract.(*_ConnectionResponseDataBlock).getLengthInBits(ctx))

	// Simple field (knxAddress)
	lengthInBits += m.KnxAddress.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ConnectionResponseDataBlockTunnelConnection) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ConnectionResponseDataBlockTunnelConnection) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ConnectionResponseDataBlock) (__connectionResponseDataBlockTunnelConnection ConnectionResponseDataBlockTunnelConnection, err error) {
	m.ConnectionResponseDataBlockContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectionResponseDataBlockTunnelConnection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionResponseDataBlockTunnelConnection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	knxAddress, err := ReadSimpleField[KnxAddress](ctx, "knxAddress", ReadComplex[KnxAddress](KnxAddressParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'knxAddress' field"))
	}
	m.KnxAddress = knxAddress

	if closeErr := readBuffer.CloseContext("ConnectionResponseDataBlockTunnelConnection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionResponseDataBlockTunnelConnection")
	}

	return m, nil
}

func (m *_ConnectionResponseDataBlockTunnelConnection) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConnectionResponseDataBlockTunnelConnection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionResponseDataBlockTunnelConnection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectionResponseDataBlockTunnelConnection")
		}

		if err := WriteSimpleField[KnxAddress](ctx, "knxAddress", m.GetKnxAddress(), WriteComplex[KnxAddress](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'knxAddress' field")
		}

		if popErr := writeBuffer.PopContext("ConnectionResponseDataBlockTunnelConnection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectionResponseDataBlockTunnelConnection")
		}
		return nil
	}
	return m.ConnectionResponseDataBlockContract.(*_ConnectionResponseDataBlock).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConnectionResponseDataBlockTunnelConnection) IsConnectionResponseDataBlockTunnelConnection() {
}

func (m *_ConnectionResponseDataBlockTunnelConnection) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ConnectionResponseDataBlockTunnelConnection) deepCopy() *_ConnectionResponseDataBlockTunnelConnection {
	if m == nil {
		return nil
	}
	_ConnectionResponseDataBlockTunnelConnectionCopy := &_ConnectionResponseDataBlockTunnelConnection{
		m.ConnectionResponseDataBlockContract.(*_ConnectionResponseDataBlock).deepCopy(),
		utils.DeepCopy[KnxAddress](m.KnxAddress),
	}
	_ConnectionResponseDataBlockTunnelConnectionCopy.ConnectionResponseDataBlockContract.(*_ConnectionResponseDataBlock)._SubType = m
	return _ConnectionResponseDataBlockTunnelConnectionCopy
}

func (m *_ConnectionResponseDataBlockTunnelConnection) String() string {
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