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

// BACnetHostNPort is the corresponding interface of BACnetHostNPort
type BACnetHostNPort interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHost returns Host (property field)
	GetHost() BACnetHostAddressEnclosed
	// GetPort returns Port (property field)
	GetPort() BACnetContextTagUnsignedInteger
	// IsBACnetHostNPort is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetHostNPort()
	// CreateBuilder creates a BACnetHostNPortBuilder
	CreateBACnetHostNPortBuilder() BACnetHostNPortBuilder
}

// _BACnetHostNPort is the data-structure of this message
type _BACnetHostNPort struct {
	Host BACnetHostAddressEnclosed
	Port BACnetContextTagUnsignedInteger
}

var _ BACnetHostNPort = (*_BACnetHostNPort)(nil)

// NewBACnetHostNPort factory function for _BACnetHostNPort
func NewBACnetHostNPort(host BACnetHostAddressEnclosed, port BACnetContextTagUnsignedInteger) *_BACnetHostNPort {
	if host == nil {
		panic("host of type BACnetHostAddressEnclosed for BACnetHostNPort must not be nil")
	}
	if port == nil {
		panic("port of type BACnetContextTagUnsignedInteger for BACnetHostNPort must not be nil")
	}
	return &_BACnetHostNPort{Host: host, Port: port}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetHostNPortBuilder is a builder for BACnetHostNPort
type BACnetHostNPortBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(host BACnetHostAddressEnclosed, port BACnetContextTagUnsignedInteger) BACnetHostNPortBuilder
	// WithHost adds Host (property field)
	WithHost(BACnetHostAddressEnclosed) BACnetHostNPortBuilder
	// WithHostBuilder adds Host (property field) which is build by the builder
	WithHostBuilder(func(BACnetHostAddressEnclosedBuilder) BACnetHostAddressEnclosedBuilder) BACnetHostNPortBuilder
	// WithPort adds Port (property field)
	WithPort(BACnetContextTagUnsignedInteger) BACnetHostNPortBuilder
	// WithPortBuilder adds Port (property field) which is build by the builder
	WithPortBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetHostNPortBuilder
	// Build builds the BACnetHostNPort or returns an error if something is wrong
	Build() (BACnetHostNPort, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetHostNPort
}

// NewBACnetHostNPortBuilder() creates a BACnetHostNPortBuilder
func NewBACnetHostNPortBuilder() BACnetHostNPortBuilder {
	return &_BACnetHostNPortBuilder{_BACnetHostNPort: new(_BACnetHostNPort)}
}

type _BACnetHostNPortBuilder struct {
	*_BACnetHostNPort

	err *utils.MultiError
}

var _ (BACnetHostNPortBuilder) = (*_BACnetHostNPortBuilder)(nil)

func (b *_BACnetHostNPortBuilder) WithMandatoryFields(host BACnetHostAddressEnclosed, port BACnetContextTagUnsignedInteger) BACnetHostNPortBuilder {
	return b.WithHost(host).WithPort(port)
}

func (b *_BACnetHostNPortBuilder) WithHost(host BACnetHostAddressEnclosed) BACnetHostNPortBuilder {
	b.Host = host
	return b
}

func (b *_BACnetHostNPortBuilder) WithHostBuilder(builderSupplier func(BACnetHostAddressEnclosedBuilder) BACnetHostAddressEnclosedBuilder) BACnetHostNPortBuilder {
	builder := builderSupplier(b.Host.CreateBACnetHostAddressEnclosedBuilder())
	var err error
	b.Host, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetHostAddressEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetHostNPortBuilder) WithPort(port BACnetContextTagUnsignedInteger) BACnetHostNPortBuilder {
	b.Port = port
	return b
}

func (b *_BACnetHostNPortBuilder) WithPortBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetHostNPortBuilder {
	builder := builderSupplier(b.Port.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.Port, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetHostNPortBuilder) Build() (BACnetHostNPort, error) {
	if b.Host == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'host' not set"))
	}
	if b.Port == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'port' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetHostNPort.deepCopy(), nil
}

func (b *_BACnetHostNPortBuilder) MustBuild() BACnetHostNPort {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetHostNPortBuilder) DeepCopy() any {
	_copy := b.CreateBACnetHostNPortBuilder().(*_BACnetHostNPortBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetHostNPortBuilder creates a BACnetHostNPortBuilder
func (b *_BACnetHostNPort) CreateBACnetHostNPortBuilder() BACnetHostNPortBuilder {
	if b == nil {
		return NewBACnetHostNPortBuilder()
	}
	return &_BACnetHostNPortBuilder{_BACnetHostNPort: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetHostNPort) GetHost() BACnetHostAddressEnclosed {
	return m.Host
}

func (m *_BACnetHostNPort) GetPort() BACnetContextTagUnsignedInteger {
	return m.Port
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetHostNPort(structType any) BACnetHostNPort {
	if casted, ok := structType.(BACnetHostNPort); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetHostNPort); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetHostNPort) GetTypeName() string {
	return "BACnetHostNPort"
}

func (m *_BACnetHostNPort) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (host)
	lengthInBits += m.Host.GetLengthInBits(ctx)

	// Simple field (port)
	lengthInBits += m.Port.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetHostNPort) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetHostNPortParse(ctx context.Context, theBytes []byte) (BACnetHostNPort, error) {
	return BACnetHostNPortParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetHostNPortParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetHostNPort, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetHostNPort, error) {
		return BACnetHostNPortParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetHostNPortParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetHostNPort, error) {
	v, err := (&_BACnetHostNPort{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetHostNPort) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetHostNPort BACnetHostNPort, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetHostNPort"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetHostNPort")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	host, err := ReadSimpleField[BACnetHostAddressEnclosed](ctx, "host", ReadComplex[BACnetHostAddressEnclosed](BACnetHostAddressEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'host' field"))
	}
	m.Host = host

	port, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "port", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'port' field"))
	}
	m.Port = port

	if closeErr := readBuffer.CloseContext("BACnetHostNPort"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetHostNPort")
	}

	return m, nil
}

func (m *_BACnetHostNPort) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetHostNPort) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetHostNPort"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetHostNPort")
	}

	if err := WriteSimpleField[BACnetHostAddressEnclosed](ctx, "host", m.GetHost(), WriteComplex[BACnetHostAddressEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'host' field")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "port", m.GetPort(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'port' field")
	}

	if popErr := writeBuffer.PopContext("BACnetHostNPort"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetHostNPort")
	}
	return nil
}

func (m *_BACnetHostNPort) IsBACnetHostNPort() {}

func (m *_BACnetHostNPort) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetHostNPort) deepCopy() *_BACnetHostNPort {
	if m == nil {
		return nil
	}
	_BACnetHostNPortCopy := &_BACnetHostNPort{
		utils.DeepCopy[BACnetHostAddressEnclosed](m.Host),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.Port),
	}
	return _BACnetHostNPortCopy
}

func (m *_BACnetHostNPort) String() string {
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