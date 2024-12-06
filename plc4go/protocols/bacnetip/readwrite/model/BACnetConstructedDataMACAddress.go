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

// BACnetConstructedDataMACAddress is the corresponding interface of BACnetConstructedDataMACAddress
type BACnetConstructedDataMACAddress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMacAddress returns MacAddress (property field)
	GetMacAddress() BACnetApplicationTagOctetString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagOctetString
	// IsBACnetConstructedDataMACAddress is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMACAddress()
	// CreateBuilder creates a BACnetConstructedDataMACAddressBuilder
	CreateBACnetConstructedDataMACAddressBuilder() BACnetConstructedDataMACAddressBuilder
}

// _BACnetConstructedDataMACAddress is the data-structure of this message
type _BACnetConstructedDataMACAddress struct {
	BACnetConstructedDataContract
	MacAddress BACnetApplicationTagOctetString
}

var _ BACnetConstructedDataMACAddress = (*_BACnetConstructedDataMACAddress)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMACAddress)(nil)

// NewBACnetConstructedDataMACAddress factory function for _BACnetConstructedDataMACAddress
func NewBACnetConstructedDataMACAddress(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, macAddress BACnetApplicationTagOctetString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMACAddress {
	if macAddress == nil {
		panic("macAddress of type BACnetApplicationTagOctetString for BACnetConstructedDataMACAddress must not be nil")
	}
	_result := &_BACnetConstructedDataMACAddress{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MacAddress:                    macAddress,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataMACAddressBuilder is a builder for BACnetConstructedDataMACAddress
type BACnetConstructedDataMACAddressBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(macAddress BACnetApplicationTagOctetString) BACnetConstructedDataMACAddressBuilder
	// WithMacAddress adds MacAddress (property field)
	WithMacAddress(BACnetApplicationTagOctetString) BACnetConstructedDataMACAddressBuilder
	// WithMacAddressBuilder adds MacAddress (property field) which is build by the builder
	WithMacAddressBuilder(func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetConstructedDataMACAddressBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataMACAddress or returns an error if something is wrong
	Build() (BACnetConstructedDataMACAddress, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataMACAddress
}

// NewBACnetConstructedDataMACAddressBuilder() creates a BACnetConstructedDataMACAddressBuilder
func NewBACnetConstructedDataMACAddressBuilder() BACnetConstructedDataMACAddressBuilder {
	return &_BACnetConstructedDataMACAddressBuilder{_BACnetConstructedDataMACAddress: new(_BACnetConstructedDataMACAddress)}
}

type _BACnetConstructedDataMACAddressBuilder struct {
	*_BACnetConstructedDataMACAddress

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataMACAddressBuilder) = (*_BACnetConstructedDataMACAddressBuilder)(nil)

func (b *_BACnetConstructedDataMACAddressBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataMACAddress
}

func (b *_BACnetConstructedDataMACAddressBuilder) WithMandatoryFields(macAddress BACnetApplicationTagOctetString) BACnetConstructedDataMACAddressBuilder {
	return b.WithMacAddress(macAddress)
}

func (b *_BACnetConstructedDataMACAddressBuilder) WithMacAddress(macAddress BACnetApplicationTagOctetString) BACnetConstructedDataMACAddressBuilder {
	b.MacAddress = macAddress
	return b
}

func (b *_BACnetConstructedDataMACAddressBuilder) WithMacAddressBuilder(builderSupplier func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetConstructedDataMACAddressBuilder {
	builder := builderSupplier(b.MacAddress.CreateBACnetApplicationTagOctetStringBuilder())
	var err error
	b.MacAddress, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagOctetStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataMACAddressBuilder) Build() (BACnetConstructedDataMACAddress, error) {
	if b.MacAddress == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'macAddress' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataMACAddress.deepCopy(), nil
}

func (b *_BACnetConstructedDataMACAddressBuilder) MustBuild() BACnetConstructedDataMACAddress {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataMACAddressBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataMACAddressBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataMACAddressBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataMACAddressBuilder().(*_BACnetConstructedDataMACAddressBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataMACAddressBuilder creates a BACnetConstructedDataMACAddressBuilder
func (b *_BACnetConstructedDataMACAddress) CreateBACnetConstructedDataMACAddressBuilder() BACnetConstructedDataMACAddressBuilder {
	if b == nil {
		return NewBACnetConstructedDataMACAddressBuilder()
	}
	return &_BACnetConstructedDataMACAddressBuilder{_BACnetConstructedDataMACAddress: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMACAddress) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMACAddress) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAC_ADDRESS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMACAddress) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMACAddress) GetMacAddress() BACnetApplicationTagOctetString {
	return m.MacAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMACAddress) GetActualValue() BACnetApplicationTagOctetString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagOctetString(m.GetMacAddress())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMACAddress(structType any) BACnetConstructedDataMACAddress {
	if casted, ok := structType.(BACnetConstructedDataMACAddress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMACAddress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMACAddress) GetTypeName() string {
	return "BACnetConstructedDataMACAddress"
}

func (m *_BACnetConstructedDataMACAddress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (macAddress)
	lengthInBits += m.MacAddress.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMACAddress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMACAddress) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMACAddress BACnetConstructedDataMACAddress, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMACAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMACAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	macAddress, err := ReadSimpleField[BACnetApplicationTagOctetString](ctx, "macAddress", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'macAddress' field"))
	}
	m.MacAddress = macAddress

	actualValue, err := ReadVirtualField[BACnetApplicationTagOctetString](ctx, "actualValue", (*BACnetApplicationTagOctetString)(nil), macAddress)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMACAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMACAddress")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMACAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMACAddress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMACAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMACAddress")
		}

		if err := WriteSimpleField[BACnetApplicationTagOctetString](ctx, "macAddress", m.GetMacAddress(), WriteComplex[BACnetApplicationTagOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'macAddress' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMACAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMACAddress")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMACAddress) IsBACnetConstructedDataMACAddress() {}

func (m *_BACnetConstructedDataMACAddress) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataMACAddress) deepCopy() *_BACnetConstructedDataMACAddress {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataMACAddressCopy := &_BACnetConstructedDataMACAddress{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagOctetString](m.MacAddress),
	}
	_BACnetConstructedDataMACAddressCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataMACAddressCopy
}

func (m *_BACnetConstructedDataMACAddress) String() string {
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
