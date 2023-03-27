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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataBACnetIPv6UDPPort is the corresponding interface of BACnetConstructedDataBACnetIPv6UDPPort
type BACnetConstructedDataBACnetIPv6UDPPort interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetIpv6UdpPort returns Ipv6UdpPort (property field)
	GetIpv6UdpPort() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataBACnetIPv6UDPPortExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBACnetIPv6UDPPort.
// This is useful for switch cases.
type BACnetConstructedDataBACnetIPv6UDPPortExactly interface {
	BACnetConstructedDataBACnetIPv6UDPPort
	isBACnetConstructedDataBACnetIPv6UDPPort() bool
}

// _BACnetConstructedDataBACnetIPv6UDPPort is the data-structure of this message
type _BACnetConstructedDataBACnetIPv6UDPPort struct {
	*_BACnetConstructedData
	Ipv6UdpPort BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACNET_IPV6_UDP_PORT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) GetIpv6UdpPort() BACnetApplicationTagUnsignedInteger {
	return m.Ipv6UdpPort
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetIpv6UdpPort())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBACnetIPv6UDPPort factory function for _BACnetConstructedDataBACnetIPv6UDPPort
func NewBACnetConstructedDataBACnetIPv6UDPPort(ipv6UdpPort BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBACnetIPv6UDPPort {
	_result := &_BACnetConstructedDataBACnetIPv6UDPPort{
		Ipv6UdpPort:            ipv6UdpPort,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBACnetIPv6UDPPort(structType interface{}) BACnetConstructedDataBACnetIPv6UDPPort {
	if casted, ok := structType.(BACnetConstructedDataBACnetIPv6UDPPort); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBACnetIPv6UDPPort); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) GetTypeName() string {
	return "BACnetConstructedDataBACnetIPv6UDPPort"
}

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (ipv6UdpPort)
	lengthInBits += m.Ipv6UdpPort.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataBACnetIPv6UDPPortParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBACnetIPv6UDPPort, error) {
	return BACnetConstructedDataBACnetIPv6UDPPortParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataBACnetIPv6UDPPortParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBACnetIPv6UDPPort, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBACnetIPv6UDPPort"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBACnetIPv6UDPPort")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipv6UdpPort)
	if pullErr := readBuffer.PullContext("ipv6UdpPort"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ipv6UdpPort")
	}
	_ipv6UdpPort, _ipv6UdpPortErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _ipv6UdpPortErr != nil {
		return nil, errors.Wrap(_ipv6UdpPortErr, "Error parsing 'ipv6UdpPort' field of BACnetConstructedDataBACnetIPv6UDPPort")
	}
	ipv6UdpPort := _ipv6UdpPort.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("ipv6UdpPort"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ipv6UdpPort")
	}

	// Virtual field
	_actualValue := ipv6UdpPort
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBACnetIPv6UDPPort"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBACnetIPv6UDPPort")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBACnetIPv6UDPPort{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		Ipv6UdpPort: ipv6UdpPort,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBACnetIPv6UDPPort"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBACnetIPv6UDPPort")
		}

		// Simple Field (ipv6UdpPort)
		if pushErr := writeBuffer.PushContext("ipv6UdpPort"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ipv6UdpPort")
		}
		_ipv6UdpPortErr := writeBuffer.WriteSerializable(ctx, m.GetIpv6UdpPort())
		if popErr := writeBuffer.PopContext("ipv6UdpPort"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ipv6UdpPort")
		}
		if _ipv6UdpPortErr != nil {
			return errors.Wrap(_ipv6UdpPortErr, "Error serializing 'ipv6UdpPort' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBACnetIPv6UDPPort"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBACnetIPv6UDPPort")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) isBACnetConstructedDataBACnetIPv6UDPPort() bool {
	return true
}

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
