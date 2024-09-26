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

// BACnetConfirmedServiceRequestRequestKey is the corresponding interface of BACnetConfirmedServiceRequestRequestKey
type BACnetConfirmedServiceRequestRequestKey interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetBytesOfRemovedService returns BytesOfRemovedService (property field)
	GetBytesOfRemovedService() []byte
	// IsBACnetConfirmedServiceRequestRequestKey is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestRequestKey()
}

// _BACnetConfirmedServiceRequestRequestKey is the data-structure of this message
type _BACnetConfirmedServiceRequestRequestKey struct {
	BACnetConfirmedServiceRequestContract
	BytesOfRemovedService []byte

	// Arguments.
	ServiceRequestPayloadLength uint32
}

var _ BACnetConfirmedServiceRequestRequestKey = (*_BACnetConfirmedServiceRequestRequestKey)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestRequestKey)(nil)

// NewBACnetConfirmedServiceRequestRequestKey factory function for _BACnetConfirmedServiceRequestRequestKey
func NewBACnetConfirmedServiceRequestRequestKey(bytesOfRemovedService []byte, serviceRequestPayloadLength uint32, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestRequestKey {
	_result := &_BACnetConfirmedServiceRequestRequestKey{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		BytesOfRemovedService:                 bytesOfRemovedService,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestRequestKey) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_REQUEST_KEY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestRequestKey) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestRequestKey) GetBytesOfRemovedService() []byte {
	return m.BytesOfRemovedService
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestRequestKey(structType any) BACnetConfirmedServiceRequestRequestKey {
	if casted, ok := structType.(BACnetConfirmedServiceRequestRequestKey); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestRequestKey); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestRequestKey) GetTypeName() string {
	return "BACnetConfirmedServiceRequestRequestKey"
}

func (m *_BACnetConfirmedServiceRequestRequestKey) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Array field
	if len(m.BytesOfRemovedService) > 0 {
		lengthInBits += 8 * uint16(len(m.BytesOfRemovedService))
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestRequestKey) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestRequestKey) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestPayloadLength uint32, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestRequestKey BACnetConfirmedServiceRequestRequestKey, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestRequestKey"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestRequestKey")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bytesOfRemovedService, err := readBuffer.ReadByteArray("bytesOfRemovedService", int(serviceRequestPayloadLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bytesOfRemovedService' field"))
	}
	m.BytesOfRemovedService = bytesOfRemovedService

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestRequestKey"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestRequestKey")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestRequestKey) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestRequestKey) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestRequestKey"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestRequestKey")
		}

		if err := WriteByteArrayField(ctx, "bytesOfRemovedService", m.GetBytesOfRemovedService(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'bytesOfRemovedService' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestRequestKey"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestRequestKey")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestRequestKey) GetServiceRequestPayloadLength() uint32 {
	return m.ServiceRequestPayloadLength
}

//
////

func (m *_BACnetConfirmedServiceRequestRequestKey) IsBACnetConfirmedServiceRequestRequestKey() {}

func (m *_BACnetConfirmedServiceRequestRequestKey) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
