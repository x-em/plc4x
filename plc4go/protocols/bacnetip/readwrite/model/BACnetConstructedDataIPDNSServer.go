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

// BACnetConstructedDataIPDNSServer is the corresponding interface of BACnetConstructedDataIPDNSServer
type BACnetConstructedDataIPDNSServer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetIpDnsServer returns IpDnsServer (property field)
	GetIpDnsServer() []BACnetApplicationTagOctetString
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataIPDNSServer is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIPDNSServer()
	// CreateBuilder creates a BACnetConstructedDataIPDNSServerBuilder
	CreateBACnetConstructedDataIPDNSServerBuilder() BACnetConstructedDataIPDNSServerBuilder
}

// _BACnetConstructedDataIPDNSServer is the data-structure of this message
type _BACnetConstructedDataIPDNSServer struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	IpDnsServer          []BACnetApplicationTagOctetString
}

var _ BACnetConstructedDataIPDNSServer = (*_BACnetConstructedDataIPDNSServer)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIPDNSServer)(nil)

// NewBACnetConstructedDataIPDNSServer factory function for _BACnetConstructedDataIPDNSServer
func NewBACnetConstructedDataIPDNSServer(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, ipDnsServer []BACnetApplicationTagOctetString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPDNSServer {
	_result := &_BACnetConstructedDataIPDNSServer{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		IpDnsServer:                   ipDnsServer,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIPDNSServerBuilder is a builder for BACnetConstructedDataIPDNSServer
type BACnetConstructedDataIPDNSServerBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(ipDnsServer []BACnetApplicationTagOctetString) BACnetConstructedDataIPDNSServerBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIPDNSServerBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataIPDNSServerBuilder
	// WithIpDnsServer adds IpDnsServer (property field)
	WithIpDnsServer(...BACnetApplicationTagOctetString) BACnetConstructedDataIPDNSServerBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataIPDNSServer or returns an error if something is wrong
	Build() (BACnetConstructedDataIPDNSServer, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIPDNSServer
}

// NewBACnetConstructedDataIPDNSServerBuilder() creates a BACnetConstructedDataIPDNSServerBuilder
func NewBACnetConstructedDataIPDNSServerBuilder() BACnetConstructedDataIPDNSServerBuilder {
	return &_BACnetConstructedDataIPDNSServerBuilder{_BACnetConstructedDataIPDNSServer: new(_BACnetConstructedDataIPDNSServer)}
}

type _BACnetConstructedDataIPDNSServerBuilder struct {
	*_BACnetConstructedDataIPDNSServer

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataIPDNSServerBuilder) = (*_BACnetConstructedDataIPDNSServerBuilder)(nil)

func (b *_BACnetConstructedDataIPDNSServerBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataIPDNSServer
}

func (b *_BACnetConstructedDataIPDNSServerBuilder) WithMandatoryFields(ipDnsServer []BACnetApplicationTagOctetString) BACnetConstructedDataIPDNSServerBuilder {
	return b.WithIpDnsServer(ipDnsServer...)
}

func (b *_BACnetConstructedDataIPDNSServerBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIPDNSServerBuilder {
	b.NumberOfDataElements = numberOfDataElements
	return b
}

func (b *_BACnetConstructedDataIPDNSServerBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataIPDNSServerBuilder {
	builder := builderSupplier(b.NumberOfDataElements.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.NumberOfDataElements, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataIPDNSServerBuilder) WithIpDnsServer(ipDnsServer ...BACnetApplicationTagOctetString) BACnetConstructedDataIPDNSServerBuilder {
	b.IpDnsServer = ipDnsServer
	return b
}

func (b *_BACnetConstructedDataIPDNSServerBuilder) Build() (BACnetConstructedDataIPDNSServer, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataIPDNSServer.deepCopy(), nil
}

func (b *_BACnetConstructedDataIPDNSServerBuilder) MustBuild() BACnetConstructedDataIPDNSServer {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataIPDNSServerBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataIPDNSServerBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataIPDNSServerBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataIPDNSServerBuilder().(*_BACnetConstructedDataIPDNSServerBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataIPDNSServerBuilder creates a BACnetConstructedDataIPDNSServerBuilder
func (b *_BACnetConstructedDataIPDNSServer) CreateBACnetConstructedDataIPDNSServerBuilder() BACnetConstructedDataIPDNSServerBuilder {
	if b == nil {
		return NewBACnetConstructedDataIPDNSServerBuilder()
	}
	return &_BACnetConstructedDataIPDNSServerBuilder{_BACnetConstructedDataIPDNSServer: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPDNSServer) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPDNSServer) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IP_DNS_SERVER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPDNSServer) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPDNSServer) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataIPDNSServer) GetIpDnsServer() []BACnetApplicationTagOctetString {
	return m.IpDnsServer
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPDNSServer) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.GetNumberOfDataElements()
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPDNSServer(structType any) BACnetConstructedDataIPDNSServer {
	if casted, ok := structType.(BACnetConstructedDataIPDNSServer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPDNSServer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPDNSServer) GetTypeName() string {
	return "BACnetConstructedDataIPDNSServer"
}

func (m *_BACnetConstructedDataIPDNSServer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.IpDnsServer) > 0 {
		for _, element := range m.IpDnsServer {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataIPDNSServer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIPDNSServer) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIPDNSServer BACnetConstructedDataIPDNSServer, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPDNSServer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPDNSServer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
		m.NumberOfDataElements = numberOfDataElements
	}

	ipDnsServer, err := ReadTerminatedArrayField[BACnetApplicationTagOctetString](ctx, "ipDnsServer", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipDnsServer' field"))
	}
	m.IpDnsServer = ipDnsServer

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPDNSServer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPDNSServer")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIPDNSServer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPDNSServer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPDNSServer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPDNSServer")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", GetRef(m.GetNumberOfDataElements()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfDataElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "ipDnsServer", m.GetIpDnsServer(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'ipDnsServer' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPDNSServer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPDNSServer")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPDNSServer) IsBACnetConstructedDataIPDNSServer() {}

func (m *_BACnetConstructedDataIPDNSServer) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIPDNSServer) deepCopy() *_BACnetConstructedDataIPDNSServer {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIPDNSServerCopy := &_BACnetConstructedDataIPDNSServer{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.NumberOfDataElements),
		utils.DeepCopySlice[BACnetApplicationTagOctetString, BACnetApplicationTagOctetString](m.IpDnsServer),
	}
	_BACnetConstructedDataIPDNSServerCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIPDNSServerCopy
}

func (m *_BACnetConstructedDataIPDNSServer) String() string {
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