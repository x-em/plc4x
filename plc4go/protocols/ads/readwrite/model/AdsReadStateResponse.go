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

// AdsReadStateResponse is the corresponding interface of AdsReadStateResponse
type AdsReadStateResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AmsPacket
	// GetResult returns Result (property field)
	GetResult() ReturnCode
	// GetAdsState returns AdsState (property field)
	GetAdsState() uint16
	// GetDeviceState returns DeviceState (property field)
	GetDeviceState() uint16
	// IsAdsReadStateResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsReadStateResponse()
	// CreateBuilder creates a AdsReadStateResponseBuilder
	CreateAdsReadStateResponseBuilder() AdsReadStateResponseBuilder
}

// _AdsReadStateResponse is the data-structure of this message
type _AdsReadStateResponse struct {
	AmsPacketContract
	Result      ReturnCode
	AdsState    uint16
	DeviceState uint16
}

var _ AdsReadStateResponse = (*_AdsReadStateResponse)(nil)
var _ AmsPacketRequirements = (*_AdsReadStateResponse)(nil)

// NewAdsReadStateResponse factory function for _AdsReadStateResponse
func NewAdsReadStateResponse(targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32, result ReturnCode, adsState uint16, deviceState uint16) *_AdsReadStateResponse {
	_result := &_AdsReadStateResponse{
		AmsPacketContract: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
		Result:            result,
		AdsState:          adsState,
		DeviceState:       deviceState,
	}
	_result.AmsPacketContract.(*_AmsPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AdsReadStateResponseBuilder is a builder for AdsReadStateResponse
type AdsReadStateResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(result ReturnCode, adsState uint16, deviceState uint16) AdsReadStateResponseBuilder
	// WithResult adds Result (property field)
	WithResult(ReturnCode) AdsReadStateResponseBuilder
	// WithAdsState adds AdsState (property field)
	WithAdsState(uint16) AdsReadStateResponseBuilder
	// WithDeviceState adds DeviceState (property field)
	WithDeviceState(uint16) AdsReadStateResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() AmsPacketBuilder
	// Build builds the AdsReadStateResponse or returns an error if something is wrong
	Build() (AdsReadStateResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AdsReadStateResponse
}

// NewAdsReadStateResponseBuilder() creates a AdsReadStateResponseBuilder
func NewAdsReadStateResponseBuilder() AdsReadStateResponseBuilder {
	return &_AdsReadStateResponseBuilder{_AdsReadStateResponse: new(_AdsReadStateResponse)}
}

type _AdsReadStateResponseBuilder struct {
	*_AdsReadStateResponse

	parentBuilder *_AmsPacketBuilder

	err *utils.MultiError
}

var _ (AdsReadStateResponseBuilder) = (*_AdsReadStateResponseBuilder)(nil)

func (b *_AdsReadStateResponseBuilder) setParent(contract AmsPacketContract) {
	b.AmsPacketContract = contract
	contract.(*_AmsPacket)._SubType = b._AdsReadStateResponse
}

func (b *_AdsReadStateResponseBuilder) WithMandatoryFields(result ReturnCode, adsState uint16, deviceState uint16) AdsReadStateResponseBuilder {
	return b.WithResult(result).WithAdsState(adsState).WithDeviceState(deviceState)
}

func (b *_AdsReadStateResponseBuilder) WithResult(result ReturnCode) AdsReadStateResponseBuilder {
	b.Result = result
	return b
}

func (b *_AdsReadStateResponseBuilder) WithAdsState(adsState uint16) AdsReadStateResponseBuilder {
	b.AdsState = adsState
	return b
}

func (b *_AdsReadStateResponseBuilder) WithDeviceState(deviceState uint16) AdsReadStateResponseBuilder {
	b.DeviceState = deviceState
	return b
}

func (b *_AdsReadStateResponseBuilder) Build() (AdsReadStateResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AdsReadStateResponse.deepCopy(), nil
}

func (b *_AdsReadStateResponseBuilder) MustBuild() AdsReadStateResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AdsReadStateResponseBuilder) Done() AmsPacketBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAmsPacketBuilder().(*_AmsPacketBuilder)
	}
	return b.parentBuilder
}

func (b *_AdsReadStateResponseBuilder) buildForAmsPacket() (AmsPacket, error) {
	return b.Build()
}

func (b *_AdsReadStateResponseBuilder) DeepCopy() any {
	_copy := b.CreateAdsReadStateResponseBuilder().(*_AdsReadStateResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAdsReadStateResponseBuilder creates a AdsReadStateResponseBuilder
func (b *_AdsReadStateResponse) CreateAdsReadStateResponseBuilder() AdsReadStateResponseBuilder {
	if b == nil {
		return NewAdsReadStateResponseBuilder()
	}
	return &_AdsReadStateResponseBuilder{_AdsReadStateResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsReadStateResponse) GetCommandId() CommandId {
	return CommandId_ADS_READ_STATE
}

func (m *_AdsReadStateResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsReadStateResponse) GetParent() AmsPacketContract {
	return m.AmsPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsReadStateResponse) GetResult() ReturnCode {
	return m.Result
}

func (m *_AdsReadStateResponse) GetAdsState() uint16 {
	return m.AdsState
}

func (m *_AdsReadStateResponse) GetDeviceState() uint16 {
	return m.DeviceState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsReadStateResponse(structType any) AdsReadStateResponse {
	if casted, ok := structType.(AdsReadStateResponse); ok {
		return casted
	}
	if casted, ok := structType.(*AdsReadStateResponse); ok {
		return *casted
	}
	return nil
}

func (m *_AdsReadStateResponse) GetTypeName() string {
	return "AdsReadStateResponse"
}

func (m *_AdsReadStateResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AmsPacketContract.(*_AmsPacket).getLengthInBits(ctx))

	// Simple field (result)
	lengthInBits += 32

	// Simple field (adsState)
	lengthInBits += 16

	// Simple field (deviceState)
	lengthInBits += 16

	return lengthInBits
}

func (m *_AdsReadStateResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsReadStateResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AmsPacket) (__adsReadStateResponse AdsReadStateResponse, err error) {
	m.AmsPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsReadStateResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsReadStateResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	result, err := ReadEnumField[ReturnCode](ctx, "result", "ReturnCode", ReadEnum(ReturnCodeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'result' field"))
	}
	m.Result = result

	adsState, err := ReadSimpleField(ctx, "adsState", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'adsState' field"))
	}
	m.AdsState = adsState

	deviceState, err := ReadSimpleField(ctx, "deviceState", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceState' field"))
	}
	m.DeviceState = deviceState

	if closeErr := readBuffer.CloseContext("AdsReadStateResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsReadStateResponse")
	}

	return m, nil
}

func (m *_AdsReadStateResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsReadStateResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadStateResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsReadStateResponse")
		}

		if err := WriteSimpleEnumField[ReturnCode](ctx, "result", "ReturnCode", m.GetResult(), WriteEnum[ReturnCode, uint32](ReturnCode.GetValue, ReturnCode.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'result' field")
		}

		if err := WriteSimpleField[uint16](ctx, "adsState", m.GetAdsState(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'adsState' field")
		}

		if err := WriteSimpleField[uint16](ctx, "deviceState", m.GetDeviceState(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'deviceState' field")
		}

		if popErr := writeBuffer.PopContext("AdsReadStateResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsReadStateResponse")
		}
		return nil
	}
	return m.AmsPacketContract.(*_AmsPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsReadStateResponse) IsAdsReadStateResponse() {}

func (m *_AdsReadStateResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsReadStateResponse) deepCopy() *_AdsReadStateResponse {
	if m == nil {
		return nil
	}
	_AdsReadStateResponseCopy := &_AdsReadStateResponse{
		m.AmsPacketContract.(*_AmsPacket).deepCopy(),
		m.Result,
		m.AdsState,
		m.DeviceState,
	}
	_AdsReadStateResponseCopy.AmsPacketContract.(*_AmsPacket)._SubType = m
	return _AdsReadStateResponseCopy
}

func (m *_AdsReadStateResponse) String() string {
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