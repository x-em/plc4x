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

// HPAIControlEndpoint is the corresponding interface of HPAIControlEndpoint
type HPAIControlEndpoint interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHostProtocolCode returns HostProtocolCode (property field)
	GetHostProtocolCode() HostProtocolCode
	// GetIpAddress returns IpAddress (property field)
	GetIpAddress() IPAddress
	// GetIpPort returns IpPort (property field)
	GetIpPort() uint16
	// IsHPAIControlEndpoint is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsHPAIControlEndpoint()
	// CreateBuilder creates a HPAIControlEndpointBuilder
	CreateHPAIControlEndpointBuilder() HPAIControlEndpointBuilder
}

// _HPAIControlEndpoint is the data-structure of this message
type _HPAIControlEndpoint struct {
	HostProtocolCode HostProtocolCode
	IpAddress        IPAddress
	IpPort           uint16
}

var _ HPAIControlEndpoint = (*_HPAIControlEndpoint)(nil)

// NewHPAIControlEndpoint factory function for _HPAIControlEndpoint
func NewHPAIControlEndpoint(hostProtocolCode HostProtocolCode, ipAddress IPAddress, ipPort uint16) *_HPAIControlEndpoint {
	if ipAddress == nil {
		panic("ipAddress of type IPAddress for HPAIControlEndpoint must not be nil")
	}
	return &_HPAIControlEndpoint{HostProtocolCode: hostProtocolCode, IpAddress: ipAddress, IpPort: ipPort}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// HPAIControlEndpointBuilder is a builder for HPAIControlEndpoint
type HPAIControlEndpointBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(hostProtocolCode HostProtocolCode, ipAddress IPAddress, ipPort uint16) HPAIControlEndpointBuilder
	// WithHostProtocolCode adds HostProtocolCode (property field)
	WithHostProtocolCode(HostProtocolCode) HPAIControlEndpointBuilder
	// WithIpAddress adds IpAddress (property field)
	WithIpAddress(IPAddress) HPAIControlEndpointBuilder
	// WithIpAddressBuilder adds IpAddress (property field) which is build by the builder
	WithIpAddressBuilder(func(IPAddressBuilder) IPAddressBuilder) HPAIControlEndpointBuilder
	// WithIpPort adds IpPort (property field)
	WithIpPort(uint16) HPAIControlEndpointBuilder
	// Build builds the HPAIControlEndpoint or returns an error if something is wrong
	Build() (HPAIControlEndpoint, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() HPAIControlEndpoint
}

// NewHPAIControlEndpointBuilder() creates a HPAIControlEndpointBuilder
func NewHPAIControlEndpointBuilder() HPAIControlEndpointBuilder {
	return &_HPAIControlEndpointBuilder{_HPAIControlEndpoint: new(_HPAIControlEndpoint)}
}

type _HPAIControlEndpointBuilder struct {
	*_HPAIControlEndpoint

	err *utils.MultiError
}

var _ (HPAIControlEndpointBuilder) = (*_HPAIControlEndpointBuilder)(nil)

func (b *_HPAIControlEndpointBuilder) WithMandatoryFields(hostProtocolCode HostProtocolCode, ipAddress IPAddress, ipPort uint16) HPAIControlEndpointBuilder {
	return b.WithHostProtocolCode(hostProtocolCode).WithIpAddress(ipAddress).WithIpPort(ipPort)
}

func (b *_HPAIControlEndpointBuilder) WithHostProtocolCode(hostProtocolCode HostProtocolCode) HPAIControlEndpointBuilder {
	b.HostProtocolCode = hostProtocolCode
	return b
}

func (b *_HPAIControlEndpointBuilder) WithIpAddress(ipAddress IPAddress) HPAIControlEndpointBuilder {
	b.IpAddress = ipAddress
	return b
}

func (b *_HPAIControlEndpointBuilder) WithIpAddressBuilder(builderSupplier func(IPAddressBuilder) IPAddressBuilder) HPAIControlEndpointBuilder {
	builder := builderSupplier(b.IpAddress.CreateIPAddressBuilder())
	var err error
	b.IpAddress, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "IPAddressBuilder failed"))
	}
	return b
}

func (b *_HPAIControlEndpointBuilder) WithIpPort(ipPort uint16) HPAIControlEndpointBuilder {
	b.IpPort = ipPort
	return b
}

func (b *_HPAIControlEndpointBuilder) Build() (HPAIControlEndpoint, error) {
	if b.IpAddress == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'ipAddress' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._HPAIControlEndpoint.deepCopy(), nil
}

func (b *_HPAIControlEndpointBuilder) MustBuild() HPAIControlEndpoint {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_HPAIControlEndpointBuilder) DeepCopy() any {
	_copy := b.CreateHPAIControlEndpointBuilder().(*_HPAIControlEndpointBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateHPAIControlEndpointBuilder creates a HPAIControlEndpointBuilder
func (b *_HPAIControlEndpoint) CreateHPAIControlEndpointBuilder() HPAIControlEndpointBuilder {
	if b == nil {
		return NewHPAIControlEndpointBuilder()
	}
	return &_HPAIControlEndpointBuilder{_HPAIControlEndpoint: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HPAIControlEndpoint) GetHostProtocolCode() HostProtocolCode {
	return m.HostProtocolCode
}

func (m *_HPAIControlEndpoint) GetIpAddress() IPAddress {
	return m.IpAddress
}

func (m *_HPAIControlEndpoint) GetIpPort() uint16 {
	return m.IpPort
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastHPAIControlEndpoint(structType any) HPAIControlEndpoint {
	if casted, ok := structType.(HPAIControlEndpoint); ok {
		return casted
	}
	if casted, ok := structType.(*HPAIControlEndpoint); ok {
		return *casted
	}
	return nil
}

func (m *_HPAIControlEndpoint) GetTypeName() string {
	return "HPAIControlEndpoint"
}

func (m *_HPAIControlEndpoint) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (hostProtocolCode)
	lengthInBits += 8

	// Simple field (ipAddress)
	lengthInBits += m.IpAddress.GetLengthInBits(ctx)

	// Simple field (ipPort)
	lengthInBits += 16

	return lengthInBits
}

func (m *_HPAIControlEndpoint) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HPAIControlEndpointParse(ctx context.Context, theBytes []byte) (HPAIControlEndpoint, error) {
	return HPAIControlEndpointParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func HPAIControlEndpointParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (HPAIControlEndpoint, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (HPAIControlEndpoint, error) {
		return HPAIControlEndpointParseWithBuffer(ctx, readBuffer)
	}
}

func HPAIControlEndpointParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HPAIControlEndpoint, error) {
	v, err := (&_HPAIControlEndpoint{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_HPAIControlEndpoint) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__hPAIControlEndpoint HPAIControlEndpoint, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HPAIControlEndpoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HPAIControlEndpoint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	structureLength, err := ReadImplicitField[uint8](ctx, "structureLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'structureLength' field"))
	}
	_ = structureLength

	hostProtocolCode, err := ReadEnumField[HostProtocolCode](ctx, "hostProtocolCode", "HostProtocolCode", ReadEnum(HostProtocolCodeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hostProtocolCode' field"))
	}
	m.HostProtocolCode = hostProtocolCode

	ipAddress, err := ReadSimpleField[IPAddress](ctx, "ipAddress", ReadComplex[IPAddress](IPAddressParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipAddress' field"))
	}
	m.IpAddress = ipAddress

	ipPort, err := ReadSimpleField(ctx, "ipPort", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipPort' field"))
	}
	m.IpPort = ipPort

	if closeErr := readBuffer.CloseContext("HPAIControlEndpoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HPAIControlEndpoint")
	}

	return m, nil
}

func (m *_HPAIControlEndpoint) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HPAIControlEndpoint) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("HPAIControlEndpoint"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for HPAIControlEndpoint")
	}
	structureLength := uint8(uint8(m.GetLengthInBytes(ctx)))
	if err := WriteImplicitField(ctx, "structureLength", structureLength, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'structureLength' field")
	}

	if err := WriteSimpleEnumField[HostProtocolCode](ctx, "hostProtocolCode", "HostProtocolCode", m.GetHostProtocolCode(), WriteEnum[HostProtocolCode, uint8](HostProtocolCode.GetValue, HostProtocolCode.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'hostProtocolCode' field")
	}

	if err := WriteSimpleField[IPAddress](ctx, "ipAddress", m.GetIpAddress(), WriteComplex[IPAddress](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'ipAddress' field")
	}

	if err := WriteSimpleField[uint16](ctx, "ipPort", m.GetIpPort(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'ipPort' field")
	}

	if popErr := writeBuffer.PopContext("HPAIControlEndpoint"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for HPAIControlEndpoint")
	}
	return nil
}

func (m *_HPAIControlEndpoint) IsHPAIControlEndpoint() {}

func (m *_HPAIControlEndpoint) DeepCopy() any {
	return m.deepCopy()
}

func (m *_HPAIControlEndpoint) deepCopy() *_HPAIControlEndpoint {
	if m == nil {
		return nil
	}
	_HPAIControlEndpointCopy := &_HPAIControlEndpoint{
		m.HostProtocolCode,
		utils.DeepCopy[IPAddress](m.IpAddress),
		m.IpPort,
	}
	return _HPAIControlEndpointCopy
}

func (m *_HPAIControlEndpoint) String() string {
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
