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

// BACnetAddress is the corresponding interface of BACnetAddress
type BACnetAddress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetNetworkNumber returns NetworkNumber (property field)
	GetNetworkNumber() BACnetApplicationTagUnsignedInteger
	// GetMacAddress returns MacAddress (property field)
	GetMacAddress() BACnetApplicationTagOctetString
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// GetIsLocalNetwork returns IsLocalNetwork (virtual field)
	GetIsLocalNetwork() bool
	// GetIsBroadcast returns IsBroadcast (virtual field)
	GetIsBroadcast() bool
	// IsBACnetAddress is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAddress()
}

// _BACnetAddress is the data-structure of this message
type _BACnetAddress struct {
	NetworkNumber BACnetApplicationTagUnsignedInteger
	MacAddress    BACnetApplicationTagOctetString
}

var _ BACnetAddress = (*_BACnetAddress)(nil)

// NewBACnetAddress factory function for _BACnetAddress
func NewBACnetAddress(networkNumber BACnetApplicationTagUnsignedInteger, macAddress BACnetApplicationTagOctetString) *_BACnetAddress {
	if networkNumber == nil {
		panic("networkNumber of type BACnetApplicationTagUnsignedInteger for BACnetAddress must not be nil")
	}
	if macAddress == nil {
		panic("macAddress of type BACnetApplicationTagOctetString for BACnetAddress must not be nil")
	}
	return &_BACnetAddress{NetworkNumber: networkNumber, MacAddress: macAddress}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAddress) GetNetworkNumber() BACnetApplicationTagUnsignedInteger {
	return m.NetworkNumber
}

func (m *_BACnetAddress) GetMacAddress() BACnetApplicationTagOctetString {
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

func (m *_BACnetAddress) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	return uint64(uint64(0))
}

func (m *_BACnetAddress) GetIsLocalNetwork() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetNetworkNumber().GetActualValue()) == (m.GetZero())))
}

func (m *_BACnetAddress) GetIsBroadcast() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetMacAddress().GetActualLength()) == (0)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetAddress(structType any) BACnetAddress {
	if casted, ok := structType.(BACnetAddress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAddress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAddress) GetTypeName() string {
	return "BACnetAddress"
}

func (m *_BACnetAddress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (networkNumber)
	lengthInBits += m.NetworkNumber.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (macAddress)
	lengthInBits += m.MacAddress.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetAddress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAddressParse(ctx context.Context, theBytes []byte) (BACnetAddress, error) {
	return BACnetAddressParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAddressParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAddress, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAddress, error) {
		return BACnetAddressParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetAddressParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAddress, error) {
	v, err := (&_BACnetAddress{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAddress) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetAddress BACnetAddress, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	networkNumber, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "networkNumber", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkNumber' field"))
	}
	m.NetworkNumber = networkNumber

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	isLocalNetwork, err := ReadVirtualField[bool](ctx, "isLocalNetwork", (*bool)(nil), bool((networkNumber.GetActualValue()) == (zero)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isLocalNetwork' field"))
	}
	_ = isLocalNetwork

	macAddress, err := ReadSimpleField[BACnetApplicationTagOctetString](ctx, "macAddress", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'macAddress' field"))
	}
	m.MacAddress = macAddress

	isBroadcast, err := ReadVirtualField[bool](ctx, "isBroadcast", (*bool)(nil), bool((macAddress.GetActualLength()) == (0)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isBroadcast' field"))
	}
	_ = isBroadcast

	if closeErr := readBuffer.CloseContext("BACnetAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAddress")
	}

	return m, nil
}

func (m *_BACnetAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAddress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAddress")
	}

	if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "networkNumber", m.GetNetworkNumber(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'networkNumber' field")
	}
	// Virtual field
	zero := m.GetZero()
	_ = zero
	if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
		return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
	}
	// Virtual field
	isLocalNetwork := m.GetIsLocalNetwork()
	_ = isLocalNetwork
	if _isLocalNetworkErr := writeBuffer.WriteVirtual(ctx, "isLocalNetwork", m.GetIsLocalNetwork()); _isLocalNetworkErr != nil {
		return errors.Wrap(_isLocalNetworkErr, "Error serializing 'isLocalNetwork' field")
	}

	if err := WriteSimpleField[BACnetApplicationTagOctetString](ctx, "macAddress", m.GetMacAddress(), WriteComplex[BACnetApplicationTagOctetString](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'macAddress' field")
	}
	// Virtual field
	isBroadcast := m.GetIsBroadcast()
	_ = isBroadcast
	if _isBroadcastErr := writeBuffer.WriteVirtual(ctx, "isBroadcast", m.GetIsBroadcast()); _isBroadcastErr != nil {
		return errors.Wrap(_isBroadcastErr, "Error serializing 'isBroadcast' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAddress")
	}
	return nil
}

func (m *_BACnetAddress) IsBACnetAddress() {}

func (m *_BACnetAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
