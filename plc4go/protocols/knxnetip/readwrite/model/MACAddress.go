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

// MACAddress is the corresponding interface of MACAddress
type MACAddress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetAddr returns Addr (property field)
	GetAddr() []byte
	// IsMACAddress is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMACAddress()
	// CreateBuilder creates a MACAddressBuilder
	CreateMACAddressBuilder() MACAddressBuilder
}

// _MACAddress is the data-structure of this message
type _MACAddress struct {
	Addr []byte
}

var _ MACAddress = (*_MACAddress)(nil)

// NewMACAddress factory function for _MACAddress
func NewMACAddress(addr []byte) *_MACAddress {
	return &_MACAddress{Addr: addr}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MACAddressBuilder is a builder for MACAddress
type MACAddressBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(addr []byte) MACAddressBuilder
	// WithAddr adds Addr (property field)
	WithAddr(...byte) MACAddressBuilder
	// Build builds the MACAddress or returns an error if something is wrong
	Build() (MACAddress, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MACAddress
}

// NewMACAddressBuilder() creates a MACAddressBuilder
func NewMACAddressBuilder() MACAddressBuilder {
	return &_MACAddressBuilder{_MACAddress: new(_MACAddress)}
}

type _MACAddressBuilder struct {
	*_MACAddress

	err *utils.MultiError
}

var _ (MACAddressBuilder) = (*_MACAddressBuilder)(nil)

func (b *_MACAddressBuilder) WithMandatoryFields(addr []byte) MACAddressBuilder {
	return b.WithAddr(addr...)
}

func (b *_MACAddressBuilder) WithAddr(addr ...byte) MACAddressBuilder {
	b.Addr = addr
	return b
}

func (b *_MACAddressBuilder) Build() (MACAddress, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MACAddress.deepCopy(), nil
}

func (b *_MACAddressBuilder) MustBuild() MACAddress {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_MACAddressBuilder) DeepCopy() any {
	_copy := b.CreateMACAddressBuilder().(*_MACAddressBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMACAddressBuilder creates a MACAddressBuilder
func (b *_MACAddress) CreateMACAddressBuilder() MACAddressBuilder {
	if b == nil {
		return NewMACAddressBuilder()
	}
	return &_MACAddressBuilder{_MACAddress: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MACAddress) GetAddr() []byte {
	return m.Addr
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMACAddress(structType any) MACAddress {
	if casted, ok := structType.(MACAddress); ok {
		return casted
	}
	if casted, ok := structType.(*MACAddress); ok {
		return *casted
	}
	return nil
}

func (m *_MACAddress) GetTypeName() string {
	return "MACAddress"
}

func (m *_MACAddress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Array field
	if len(m.Addr) > 0 {
		lengthInBits += 8 * uint16(len(m.Addr))
	}

	return lengthInBits
}

func (m *_MACAddress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MACAddressParse(ctx context.Context, theBytes []byte) (MACAddress, error) {
	return MACAddressParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func MACAddressParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (MACAddress, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (MACAddress, error) {
		return MACAddressParseWithBuffer(ctx, readBuffer)
	}
}

func MACAddressParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MACAddress, error) {
	v, err := (&_MACAddress{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_MACAddress) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__mACAddress MACAddress, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MACAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MACAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	addr, err := readBuffer.ReadByteArray("addr", int(int32(6)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'addr' field"))
	}
	m.Addr = addr

	if closeErr := readBuffer.CloseContext("MACAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MACAddress")
	}

	return m, nil
}

func (m *_MACAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MACAddress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("MACAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for MACAddress")
	}

	if err := WriteByteArrayField(ctx, "addr", m.GetAddr(), WriteByteArray(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'addr' field")
	}

	if popErr := writeBuffer.PopContext("MACAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for MACAddress")
	}
	return nil
}

func (m *_MACAddress) IsMACAddress() {}

func (m *_MACAddress) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MACAddress) deepCopy() *_MACAddress {
	if m == nil {
		return nil
	}
	_MACAddressCopy := &_MACAddress{
		utils.DeepCopySlice[byte, byte](m.Addr),
	}
	return _MACAddressCopy
}

func (m *_MACAddress) String() string {
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