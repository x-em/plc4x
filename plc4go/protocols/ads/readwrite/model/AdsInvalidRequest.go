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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// AdsInvalidRequest is the corresponding interface of AdsInvalidRequest
type AdsInvalidRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AmsPacket
	// IsAdsInvalidRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsInvalidRequest()
	// CreateBuilder creates a AdsInvalidRequestBuilder
	CreateAdsInvalidRequestBuilder() AdsInvalidRequestBuilder
}

// _AdsInvalidRequest is the data-structure of this message
type _AdsInvalidRequest struct {
	AmsPacketContract
}

var _ AdsInvalidRequest = (*_AdsInvalidRequest)(nil)
var _ AmsPacketRequirements = (*_AdsInvalidRequest)(nil)

// NewAdsInvalidRequest factory function for _AdsInvalidRequest
func NewAdsInvalidRequest(targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) *_AdsInvalidRequest {
	_result := &_AdsInvalidRequest{
		AmsPacketContract: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
	}
	_result.AmsPacketContract.(*_AmsPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AdsInvalidRequestBuilder is a builder for AdsInvalidRequest
type AdsInvalidRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() AdsInvalidRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() AmsPacketBuilder
	// Build builds the AdsInvalidRequest or returns an error if something is wrong
	Build() (AdsInvalidRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AdsInvalidRequest
}

// NewAdsInvalidRequestBuilder() creates a AdsInvalidRequestBuilder
func NewAdsInvalidRequestBuilder() AdsInvalidRequestBuilder {
	return &_AdsInvalidRequestBuilder{_AdsInvalidRequest: new(_AdsInvalidRequest)}
}

type _AdsInvalidRequestBuilder struct {
	*_AdsInvalidRequest

	parentBuilder *_AmsPacketBuilder

	err *utils.MultiError
}

var _ (AdsInvalidRequestBuilder) = (*_AdsInvalidRequestBuilder)(nil)

func (b *_AdsInvalidRequestBuilder) setParent(contract AmsPacketContract) {
	b.AmsPacketContract = contract
	contract.(*_AmsPacket)._SubType = b._AdsInvalidRequest
}

func (b *_AdsInvalidRequestBuilder) WithMandatoryFields() AdsInvalidRequestBuilder {
	return b
}

func (b *_AdsInvalidRequestBuilder) Build() (AdsInvalidRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AdsInvalidRequest.deepCopy(), nil
}

func (b *_AdsInvalidRequestBuilder) MustBuild() AdsInvalidRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AdsInvalidRequestBuilder) Done() AmsPacketBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAmsPacketBuilder().(*_AmsPacketBuilder)
	}
	return b.parentBuilder
}

func (b *_AdsInvalidRequestBuilder) buildForAmsPacket() (AmsPacket, error) {
	return b.Build()
}

func (b *_AdsInvalidRequestBuilder) DeepCopy() any {
	_copy := b.CreateAdsInvalidRequestBuilder().(*_AdsInvalidRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAdsInvalidRequestBuilder creates a AdsInvalidRequestBuilder
func (b *_AdsInvalidRequest) CreateAdsInvalidRequestBuilder() AdsInvalidRequestBuilder {
	if b == nil {
		return NewAdsInvalidRequestBuilder()
	}
	return &_AdsInvalidRequestBuilder{_AdsInvalidRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsInvalidRequest) GetCommandId() CommandId {
	return CommandId_INVALID
}

func (m *_AdsInvalidRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsInvalidRequest) GetParent() AmsPacketContract {
	return m.AmsPacketContract
}

// Deprecated: use the interface for direct cast
func CastAdsInvalidRequest(structType any) AdsInvalidRequest {
	if casted, ok := structType.(AdsInvalidRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AdsInvalidRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AdsInvalidRequest) GetTypeName() string {
	return "AdsInvalidRequest"
}

func (m *_AdsInvalidRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AmsPacketContract.(*_AmsPacket).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_AdsInvalidRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsInvalidRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AmsPacket) (__adsInvalidRequest AdsInvalidRequest, err error) {
	m.AmsPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsInvalidRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsInvalidRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AdsInvalidRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsInvalidRequest")
	}

	return m, nil
}

func (m *_AdsInvalidRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsInvalidRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsInvalidRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsInvalidRequest")
		}

		if popErr := writeBuffer.PopContext("AdsInvalidRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsInvalidRequest")
		}
		return nil
	}
	return m.AmsPacketContract.(*_AmsPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsInvalidRequest) IsAdsInvalidRequest() {}

func (m *_AdsInvalidRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsInvalidRequest) deepCopy() *_AdsInvalidRequest {
	if m == nil {
		return nil
	}
	_AdsInvalidRequestCopy := &_AdsInvalidRequest{
		m.AmsPacketContract.(*_AmsPacket).deepCopy(),
	}
	_AdsInvalidRequestCopy.AmsPacketContract.(*_AmsPacket)._SubType = m
	return _AdsInvalidRequestCopy
}

func (m *_AdsInvalidRequest) String() string {
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