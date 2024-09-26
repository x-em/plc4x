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

// ActivateSessionRequest is the corresponding interface of ActivateSessionRequest
type ActivateSessionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetClientSignature returns ClientSignature (property field)
	GetClientSignature() ExtensionObjectDefinition
	// GetNoOfClientSoftwareCertificates returns NoOfClientSoftwareCertificates (property field)
	GetNoOfClientSoftwareCertificates() int32
	// GetClientSoftwareCertificates returns ClientSoftwareCertificates (property field)
	GetClientSoftwareCertificates() []ExtensionObjectDefinition
	// GetNoOfLocaleIds returns NoOfLocaleIds (property field)
	GetNoOfLocaleIds() int32
	// GetLocaleIds returns LocaleIds (property field)
	GetLocaleIds() []PascalString
	// GetUserIdentityToken returns UserIdentityToken (property field)
	GetUserIdentityToken() ExtensionObject
	// GetUserTokenSignature returns UserTokenSignature (property field)
	GetUserTokenSignature() ExtensionObjectDefinition
	// IsActivateSessionRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsActivateSessionRequest()
}

// _ActivateSessionRequest is the data-structure of this message
type _ActivateSessionRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader                  ExtensionObjectDefinition
	ClientSignature                ExtensionObjectDefinition
	NoOfClientSoftwareCertificates int32
	ClientSoftwareCertificates     []ExtensionObjectDefinition
	NoOfLocaleIds                  int32
	LocaleIds                      []PascalString
	UserIdentityToken              ExtensionObject
	UserTokenSignature             ExtensionObjectDefinition
}

var _ ActivateSessionRequest = (*_ActivateSessionRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ActivateSessionRequest)(nil)

// NewActivateSessionRequest factory function for _ActivateSessionRequest
func NewActivateSessionRequest(requestHeader ExtensionObjectDefinition, clientSignature ExtensionObjectDefinition, noOfClientSoftwareCertificates int32, clientSoftwareCertificates []ExtensionObjectDefinition, noOfLocaleIds int32, localeIds []PascalString, userIdentityToken ExtensionObject, userTokenSignature ExtensionObjectDefinition) *_ActivateSessionRequest {
	if requestHeader == nil {
		panic("requestHeader of type ExtensionObjectDefinition for ActivateSessionRequest must not be nil")
	}
	if clientSignature == nil {
		panic("clientSignature of type ExtensionObjectDefinition for ActivateSessionRequest must not be nil")
	}
	if userIdentityToken == nil {
		panic("userIdentityToken of type ExtensionObject for ActivateSessionRequest must not be nil")
	}
	if userTokenSignature == nil {
		panic("userTokenSignature of type ExtensionObjectDefinition for ActivateSessionRequest must not be nil")
	}
	_result := &_ActivateSessionRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		ClientSignature:                   clientSignature,
		NoOfClientSoftwareCertificates:    noOfClientSoftwareCertificates,
		ClientSoftwareCertificates:        clientSoftwareCertificates,
		NoOfLocaleIds:                     noOfLocaleIds,
		LocaleIds:                         localeIds,
		UserIdentityToken:                 userIdentityToken,
		UserTokenSignature:                userTokenSignature,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ActivateSessionRequest) GetIdentifier() string {
	return "467"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ActivateSessionRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ActivateSessionRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_ActivateSessionRequest) GetClientSignature() ExtensionObjectDefinition {
	return m.ClientSignature
}

func (m *_ActivateSessionRequest) GetNoOfClientSoftwareCertificates() int32 {
	return m.NoOfClientSoftwareCertificates
}

func (m *_ActivateSessionRequest) GetClientSoftwareCertificates() []ExtensionObjectDefinition {
	return m.ClientSoftwareCertificates
}

func (m *_ActivateSessionRequest) GetNoOfLocaleIds() int32 {
	return m.NoOfLocaleIds
}

func (m *_ActivateSessionRequest) GetLocaleIds() []PascalString {
	return m.LocaleIds
}

func (m *_ActivateSessionRequest) GetUserIdentityToken() ExtensionObject {
	return m.UserIdentityToken
}

func (m *_ActivateSessionRequest) GetUserTokenSignature() ExtensionObjectDefinition {
	return m.UserTokenSignature
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastActivateSessionRequest(structType any) ActivateSessionRequest {
	if casted, ok := structType.(ActivateSessionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ActivateSessionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ActivateSessionRequest) GetTypeName() string {
	return "ActivateSessionRequest"
}

func (m *_ActivateSessionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (clientSignature)
	lengthInBits += m.ClientSignature.GetLengthInBits(ctx)

	// Simple field (noOfClientSoftwareCertificates)
	lengthInBits += 32

	// Array field
	if len(m.ClientSoftwareCertificates) > 0 {
		for _curItem, element := range m.ClientSoftwareCertificates {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ClientSoftwareCertificates), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfLocaleIds)
	lengthInBits += 32

	// Array field
	if len(m.LocaleIds) > 0 {
		for _curItem, element := range m.LocaleIds {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LocaleIds), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (userIdentityToken)
	lengthInBits += m.UserIdentityToken.GetLengthInBits(ctx)

	// Simple field (userTokenSignature)
	lengthInBits += m.UserTokenSignature.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ActivateSessionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ActivateSessionRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__activateSessionRequest ActivateSessionRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ActivateSessionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ActivateSessionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	clientSignature, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "clientSignature", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("458")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'clientSignature' field"))
	}
	m.ClientSignature = clientSignature

	noOfClientSoftwareCertificates, err := ReadSimpleField(ctx, "noOfClientSoftwareCertificates", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfClientSoftwareCertificates' field"))
	}
	m.NoOfClientSoftwareCertificates = noOfClientSoftwareCertificates

	clientSoftwareCertificates, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "clientSoftwareCertificates", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("346")), readBuffer), uint64(noOfClientSoftwareCertificates))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'clientSoftwareCertificates' field"))
	}
	m.ClientSoftwareCertificates = clientSoftwareCertificates

	noOfLocaleIds, err := ReadSimpleField(ctx, "noOfLocaleIds", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfLocaleIds' field"))
	}
	m.NoOfLocaleIds = noOfLocaleIds

	localeIds, err := ReadCountArrayField[PascalString](ctx, "localeIds", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfLocaleIds))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'localeIds' field"))
	}
	m.LocaleIds = localeIds

	userIdentityToken, err := ReadSimpleField[ExtensionObject](ctx, "userIdentityToken", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userIdentityToken' field"))
	}
	m.UserIdentityToken = userIdentityToken

	userTokenSignature, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "userTokenSignature", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("458")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userTokenSignature' field"))
	}
	m.UserTokenSignature = userTokenSignature

	if closeErr := readBuffer.CloseContext("ActivateSessionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ActivateSessionRequest")
	}

	return m, nil
}

func (m *_ActivateSessionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ActivateSessionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ActivateSessionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ActivateSessionRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "clientSignature", m.GetClientSignature(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'clientSignature' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfClientSoftwareCertificates", m.GetNoOfClientSoftwareCertificates(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfClientSoftwareCertificates' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "clientSoftwareCertificates", m.GetClientSoftwareCertificates(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'clientSoftwareCertificates' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfLocaleIds", m.GetNoOfLocaleIds(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfLocaleIds' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "localeIds", m.GetLocaleIds(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'localeIds' field")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "userIdentityToken", m.GetUserIdentityToken(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'userIdentityToken' field")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "userTokenSignature", m.GetUserTokenSignature(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'userTokenSignature' field")
		}

		if popErr := writeBuffer.PopContext("ActivateSessionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ActivateSessionRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ActivateSessionRequest) IsActivateSessionRequest() {}

func (m *_ActivateSessionRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
