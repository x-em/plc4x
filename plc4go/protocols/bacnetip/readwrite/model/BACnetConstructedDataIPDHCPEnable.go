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

// BACnetConstructedDataIPDHCPEnable is the corresponding interface of BACnetConstructedDataIPDHCPEnable
type BACnetConstructedDataIPDHCPEnable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetIpDhcpEnable returns IpDhcpEnable (property field)
	GetIpDhcpEnable() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataIPDHCPEnable is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIPDHCPEnable()
	// CreateBuilder creates a BACnetConstructedDataIPDHCPEnableBuilder
	CreateBACnetConstructedDataIPDHCPEnableBuilder() BACnetConstructedDataIPDHCPEnableBuilder
}

// _BACnetConstructedDataIPDHCPEnable is the data-structure of this message
type _BACnetConstructedDataIPDHCPEnable struct {
	BACnetConstructedDataContract
	IpDhcpEnable BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataIPDHCPEnable = (*_BACnetConstructedDataIPDHCPEnable)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIPDHCPEnable)(nil)

// NewBACnetConstructedDataIPDHCPEnable factory function for _BACnetConstructedDataIPDHCPEnable
func NewBACnetConstructedDataIPDHCPEnable(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, ipDhcpEnable BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPDHCPEnable {
	if ipDhcpEnable == nil {
		panic("ipDhcpEnable of type BACnetApplicationTagBoolean for BACnetConstructedDataIPDHCPEnable must not be nil")
	}
	_result := &_BACnetConstructedDataIPDHCPEnable{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		IpDhcpEnable:                  ipDhcpEnable,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIPDHCPEnableBuilder is a builder for BACnetConstructedDataIPDHCPEnable
type BACnetConstructedDataIPDHCPEnableBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(ipDhcpEnable BACnetApplicationTagBoolean) BACnetConstructedDataIPDHCPEnableBuilder
	// WithIpDhcpEnable adds IpDhcpEnable (property field)
	WithIpDhcpEnable(BACnetApplicationTagBoolean) BACnetConstructedDataIPDHCPEnableBuilder
	// WithIpDhcpEnableBuilder adds IpDhcpEnable (property field) which is build by the builder
	WithIpDhcpEnableBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataIPDHCPEnableBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataIPDHCPEnable or returns an error if something is wrong
	Build() (BACnetConstructedDataIPDHCPEnable, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIPDHCPEnable
}

// NewBACnetConstructedDataIPDHCPEnableBuilder() creates a BACnetConstructedDataIPDHCPEnableBuilder
func NewBACnetConstructedDataIPDHCPEnableBuilder() BACnetConstructedDataIPDHCPEnableBuilder {
	return &_BACnetConstructedDataIPDHCPEnableBuilder{_BACnetConstructedDataIPDHCPEnable: new(_BACnetConstructedDataIPDHCPEnable)}
}

type _BACnetConstructedDataIPDHCPEnableBuilder struct {
	*_BACnetConstructedDataIPDHCPEnable

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataIPDHCPEnableBuilder) = (*_BACnetConstructedDataIPDHCPEnableBuilder)(nil)

func (b *_BACnetConstructedDataIPDHCPEnableBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataIPDHCPEnable
}

func (b *_BACnetConstructedDataIPDHCPEnableBuilder) WithMandatoryFields(ipDhcpEnable BACnetApplicationTagBoolean) BACnetConstructedDataIPDHCPEnableBuilder {
	return b.WithIpDhcpEnable(ipDhcpEnable)
}

func (b *_BACnetConstructedDataIPDHCPEnableBuilder) WithIpDhcpEnable(ipDhcpEnable BACnetApplicationTagBoolean) BACnetConstructedDataIPDHCPEnableBuilder {
	b.IpDhcpEnable = ipDhcpEnable
	return b
}

func (b *_BACnetConstructedDataIPDHCPEnableBuilder) WithIpDhcpEnableBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataIPDHCPEnableBuilder {
	builder := builderSupplier(b.IpDhcpEnable.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.IpDhcpEnable, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataIPDHCPEnableBuilder) Build() (BACnetConstructedDataIPDHCPEnable, error) {
	if b.IpDhcpEnable == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'ipDhcpEnable' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataIPDHCPEnable.deepCopy(), nil
}

func (b *_BACnetConstructedDataIPDHCPEnableBuilder) MustBuild() BACnetConstructedDataIPDHCPEnable {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataIPDHCPEnableBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataIPDHCPEnableBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataIPDHCPEnableBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataIPDHCPEnableBuilder().(*_BACnetConstructedDataIPDHCPEnableBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataIPDHCPEnableBuilder creates a BACnetConstructedDataIPDHCPEnableBuilder
func (b *_BACnetConstructedDataIPDHCPEnable) CreateBACnetConstructedDataIPDHCPEnableBuilder() BACnetConstructedDataIPDHCPEnableBuilder {
	if b == nil {
		return NewBACnetConstructedDataIPDHCPEnableBuilder()
	}
	return &_BACnetConstructedDataIPDHCPEnableBuilder{_BACnetConstructedDataIPDHCPEnable: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPDHCPEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPDHCPEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IP_DHCP_ENABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPDHCPEnable) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPDHCPEnable) GetIpDhcpEnable() BACnetApplicationTagBoolean {
	return m.IpDhcpEnable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPDHCPEnable) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetIpDhcpEnable())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPDHCPEnable(structType any) BACnetConstructedDataIPDHCPEnable {
	if casted, ok := structType.(BACnetConstructedDataIPDHCPEnable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPDHCPEnable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPDHCPEnable) GetTypeName() string {
	return "BACnetConstructedDataIPDHCPEnable"
}

func (m *_BACnetConstructedDataIPDHCPEnable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (ipDhcpEnable)
	lengthInBits += m.IpDhcpEnable.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPDHCPEnable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIPDHCPEnable) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIPDHCPEnable BACnetConstructedDataIPDHCPEnable, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPDHCPEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPDHCPEnable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ipDhcpEnable, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "ipDhcpEnable", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipDhcpEnable' field"))
	}
	m.IpDhcpEnable = ipDhcpEnable

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), ipDhcpEnable)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPDHCPEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPDHCPEnable")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIPDHCPEnable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPDHCPEnable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPDHCPEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPDHCPEnable")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "ipDhcpEnable", m.GetIpDhcpEnable(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ipDhcpEnable' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPDHCPEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPDHCPEnable")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPDHCPEnable) IsBACnetConstructedDataIPDHCPEnable() {}

func (m *_BACnetConstructedDataIPDHCPEnable) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIPDHCPEnable) deepCopy() *_BACnetConstructedDataIPDHCPEnable {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIPDHCPEnableCopy := &_BACnetConstructedDataIPDHCPEnable{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBoolean](m.IpDhcpEnable),
	}
	_BACnetConstructedDataIPDHCPEnableCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIPDHCPEnableCopy
}

func (m *_BACnetConstructedDataIPDHCPEnable) String() string {
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