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

// OpenChannelMessageResponse is the corresponding interface of OpenChannelMessageResponse
type OpenChannelMessageResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	OpenChannelMessage
	// GetSecureChannelId returns SecureChannelId (property field)
	GetSecureChannelId() int32
	// GetSecurityPolicyUri returns SecurityPolicyUri (property field)
	GetSecurityPolicyUri() PascalString
	// GetSenderCertificate returns SenderCertificate (property field)
	GetSenderCertificate() PascalByteString
	// GetReceiverCertificateThumbprint returns ReceiverCertificateThumbprint (property field)
	GetReceiverCertificateThumbprint() PascalByteString
	// IsOpenChannelMessageResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsOpenChannelMessageResponse()
}

// _OpenChannelMessageResponse is the data-structure of this message
type _OpenChannelMessageResponse struct {
	OpenChannelMessageContract
	SecureChannelId               int32
	SecurityPolicyUri             PascalString
	SenderCertificate             PascalByteString
	ReceiverCertificateThumbprint PascalByteString
}

var _ OpenChannelMessageResponse = (*_OpenChannelMessageResponse)(nil)
var _ OpenChannelMessageRequirements = (*_OpenChannelMessageResponse)(nil)

// NewOpenChannelMessageResponse factory function for _OpenChannelMessageResponse
func NewOpenChannelMessageResponse(secureChannelId int32, securityPolicyUri PascalString, senderCertificate PascalByteString, receiverCertificateThumbprint PascalByteString) *_OpenChannelMessageResponse {
	if securityPolicyUri == nil {
		panic("securityPolicyUri of type PascalString for OpenChannelMessageResponse must not be nil")
	}
	if senderCertificate == nil {
		panic("senderCertificate of type PascalByteString for OpenChannelMessageResponse must not be nil")
	}
	if receiverCertificateThumbprint == nil {
		panic("receiverCertificateThumbprint of type PascalByteString for OpenChannelMessageResponse must not be nil")
	}
	_result := &_OpenChannelMessageResponse{
		OpenChannelMessageContract:    NewOpenChannelMessage(),
		SecureChannelId:               secureChannelId,
		SecurityPolicyUri:             securityPolicyUri,
		SenderCertificate:             senderCertificate,
		ReceiverCertificateThumbprint: receiverCertificateThumbprint,
	}
	_result.OpenChannelMessageContract.(*_OpenChannelMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OpenChannelMessageResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OpenChannelMessageResponse) GetParent() OpenChannelMessageContract {
	return m.OpenChannelMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpenChannelMessageResponse) GetSecureChannelId() int32 {
	return m.SecureChannelId
}

func (m *_OpenChannelMessageResponse) GetSecurityPolicyUri() PascalString {
	return m.SecurityPolicyUri
}

func (m *_OpenChannelMessageResponse) GetSenderCertificate() PascalByteString {
	return m.SenderCertificate
}

func (m *_OpenChannelMessageResponse) GetReceiverCertificateThumbprint() PascalByteString {
	return m.ReceiverCertificateThumbprint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastOpenChannelMessageResponse(structType any) OpenChannelMessageResponse {
	if casted, ok := structType.(OpenChannelMessageResponse); ok {
		return casted
	}
	if casted, ok := structType.(*OpenChannelMessageResponse); ok {
		return *casted
	}
	return nil
}

func (m *_OpenChannelMessageResponse) GetTypeName() string {
	return "OpenChannelMessageResponse"
}

func (m *_OpenChannelMessageResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.OpenChannelMessageContract.(*_OpenChannelMessage).getLengthInBits(ctx))

	// Simple field (secureChannelId)
	lengthInBits += 32

	// Simple field (securityPolicyUri)
	lengthInBits += m.SecurityPolicyUri.GetLengthInBits(ctx)

	// Simple field (senderCertificate)
	lengthInBits += m.SenderCertificate.GetLengthInBits(ctx)

	// Simple field (receiverCertificateThumbprint)
	lengthInBits += m.ReceiverCertificateThumbprint.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_OpenChannelMessageResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_OpenChannelMessageResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_OpenChannelMessage, response bool) (__openChannelMessageResponse OpenChannelMessageResponse, err error) {
	m.OpenChannelMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("OpenChannelMessageResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpenChannelMessageResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	secureChannelId, err := ReadSimpleField(ctx, "secureChannelId", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'secureChannelId' field"))
	}
	m.SecureChannelId = secureChannelId

	securityPolicyUri, err := ReadSimpleField[PascalString](ctx, "securityPolicyUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityPolicyUri' field"))
	}
	m.SecurityPolicyUri = securityPolicyUri

	senderCertificate, err := ReadSimpleField[PascalByteString](ctx, "senderCertificate", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'senderCertificate' field"))
	}
	m.SenderCertificate = senderCertificate

	receiverCertificateThumbprint, err := ReadSimpleField[PascalByteString](ctx, "receiverCertificateThumbprint", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'receiverCertificateThumbprint' field"))
	}
	m.ReceiverCertificateThumbprint = receiverCertificateThumbprint

	if closeErr := readBuffer.CloseContext("OpenChannelMessageResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpenChannelMessageResponse")
	}

	return m, nil
}

func (m *_OpenChannelMessageResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpenChannelMessageResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OpenChannelMessageResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OpenChannelMessageResponse")
		}

		if err := WriteSimpleField[int32](ctx, "secureChannelId", m.GetSecureChannelId(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'secureChannelId' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "securityPolicyUri", m.GetSecurityPolicyUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'securityPolicyUri' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "senderCertificate", m.GetSenderCertificate(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'senderCertificate' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "receiverCertificateThumbprint", m.GetReceiverCertificateThumbprint(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'receiverCertificateThumbprint' field")
		}

		if popErr := writeBuffer.PopContext("OpenChannelMessageResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OpenChannelMessageResponse")
		}
		return nil
	}
	return m.OpenChannelMessageContract.(*_OpenChannelMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_OpenChannelMessageResponse) IsOpenChannelMessageResponse() {}

func (m *_OpenChannelMessageResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
