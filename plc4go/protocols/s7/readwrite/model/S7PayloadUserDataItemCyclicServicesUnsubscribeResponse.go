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

// S7PayloadUserDataItemCyclicServicesUnsubscribeResponse is the corresponding interface of S7PayloadUserDataItemCyclicServicesUnsubscribeResponse
type S7PayloadUserDataItemCyclicServicesUnsubscribeResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7PayloadUserDataItem
	// IsS7PayloadUserDataItemCyclicServicesUnsubscribeResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadUserDataItemCyclicServicesUnsubscribeResponse()
	// CreateBuilder creates a S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder
	CreateS7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder() S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder
}

// _S7PayloadUserDataItemCyclicServicesUnsubscribeResponse is the data-structure of this message
type _S7PayloadUserDataItemCyclicServicesUnsubscribeResponse struct {
	S7PayloadUserDataItemContract
}

var _ S7PayloadUserDataItemCyclicServicesUnsubscribeResponse = (*_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse)(nil)

// NewS7PayloadUserDataItemCyclicServicesUnsubscribeResponse factory function for _S7PayloadUserDataItemCyclicServicesUnsubscribeResponse
func NewS7PayloadUserDataItemCyclicServicesUnsubscribeResponse(returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse {
	_result := &_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
	}
	_result.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder is a builder for S7PayloadUserDataItemCyclicServicesUnsubscribeResponse
type S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() S7PayloadUserDataItemBuilder
	// Build builds the S7PayloadUserDataItemCyclicServicesUnsubscribeResponse or returns an error if something is wrong
	Build() (S7PayloadUserDataItemCyclicServicesUnsubscribeResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7PayloadUserDataItemCyclicServicesUnsubscribeResponse
}

// NewS7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder() creates a S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder
func NewS7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder() S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder {
	return &_S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder{_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse: new(_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse)}
}

type _S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder struct {
	*_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse

	parentBuilder *_S7PayloadUserDataItemBuilder

	err *utils.MultiError
}

var _ (S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder) = (*_S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder)(nil)

func (b *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder) setParent(contract S7PayloadUserDataItemContract) {
	b.S7PayloadUserDataItemContract = contract
	contract.(*_S7PayloadUserDataItem)._SubType = b._S7PayloadUserDataItemCyclicServicesUnsubscribeResponse
}

func (b *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder) WithMandatoryFields() S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder {
	return b
}

func (b *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder) Build() (S7PayloadUserDataItemCyclicServicesUnsubscribeResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7PayloadUserDataItemCyclicServicesUnsubscribeResponse.deepCopy(), nil
}

func (b *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder) MustBuild() S7PayloadUserDataItemCyclicServicesUnsubscribeResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder) Done() S7PayloadUserDataItemBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewS7PayloadUserDataItemBuilder().(*_S7PayloadUserDataItemBuilder)
	}
	return b.parentBuilder
}

func (b *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder) buildForS7PayloadUserDataItem() (S7PayloadUserDataItem, error) {
	return b.Build()
}

func (b *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder) DeepCopy() any {
	_copy := b.CreateS7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder().(*_S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder creates a S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder
func (b *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) CreateS7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder() S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder {
	if b == nil {
		return NewS7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder()
	}
	return &_S7PayloadUserDataItemCyclicServicesUnsubscribeResponseBuilder{_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) GetCpuFunctionGroup() uint8 {
	return 0x02
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) GetCpuSubfunction() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCyclicServicesUnsubscribeResponse(structType any) S7PayloadUserDataItemCyclicServicesUnsubscribeResponse {
	if casted, ok := structType.(S7PayloadUserDataItemCyclicServicesUnsubscribeResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCyclicServicesUnsubscribeResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) GetTypeName() string {
	return "S7PayloadUserDataItemCyclicServicesUnsubscribeResponse"
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadUserDataItemCyclicServicesUnsubscribeResponse S7PayloadUserDataItemCyclicServicesUnsubscribeResponse, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCyclicServicesUnsubscribeResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCyclicServicesUnsubscribeResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCyclicServicesUnsubscribeResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCyclicServicesUnsubscribeResponse")
	}

	return m, nil
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCyclicServicesUnsubscribeResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCyclicServicesUnsubscribeResponse")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCyclicServicesUnsubscribeResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCyclicServicesUnsubscribeResponse")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) IsS7PayloadUserDataItemCyclicServicesUnsubscribeResponse() {
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) deepCopy() *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse {
	if m == nil {
		return nil
	}
	_S7PayloadUserDataItemCyclicServicesUnsubscribeResponseCopy := &_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse{
		m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).deepCopy(),
	}
	_S7PayloadUserDataItemCyclicServicesUnsubscribeResponseCopy.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = m
	return _S7PayloadUserDataItemCyclicServicesUnsubscribeResponseCopy
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) String() string {
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
