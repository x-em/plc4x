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

// S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest is the corresponding interface of S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest
type S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7PayloadUserDataItem
	// IsS7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest()
	// CreateBuilder creates a S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder
	CreateS7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder() S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder
}

// _S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest is the data-structure of this message
type _S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest struct {
	S7PayloadUserDataItemContract
}

var _ S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest = (*_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest)(nil)

// NewS7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest factory function for _S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest
func NewS7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest(returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest {
	_result := &_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
	}
	_result.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder is a builder for S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest
type S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() S7PayloadUserDataItemBuilder
	// Build builds the S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest or returns an error if something is wrong
	Build() (S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest
}

// NewS7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder() creates a S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder
func NewS7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder() S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder {
	return &_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder{_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest: new(_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest)}
}

type _S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder struct {
	*_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest

	parentBuilder *_S7PayloadUserDataItemBuilder

	err *utils.MultiError
}

var _ (S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder) = (*_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder)(nil)

func (b *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder) setParent(contract S7PayloadUserDataItemContract) {
	b.S7PayloadUserDataItemContract = contract
	contract.(*_S7PayloadUserDataItem)._SubType = b._S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest
}

func (b *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder) WithMandatoryFields() S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder {
	return b
}

func (b *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder) Build() (S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest.deepCopy(), nil
}

func (b *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder) MustBuild() S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder) Done() S7PayloadUserDataItemBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewS7PayloadUserDataItemBuilder().(*_S7PayloadUserDataItemBuilder)
	}
	return b.parentBuilder
}

func (b *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder) buildForS7PayloadUserDataItem() (S7PayloadUserDataItem, error) {
	return b.Build()
}

func (b *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder) DeepCopy() any {
	_copy := b.CreateS7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder().(*_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder creates a S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder
func (b *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest) CreateS7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder() S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder {
	if b == nil {
		return NewS7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder()
	}
	return &_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestBuilder{_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest) GetCpuFunctionGroup() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest) GetCpuFunctionType() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest) GetCpuSubfunction() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest(structType any) S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest"
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest")
	}

	return m, nil
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest) IsS7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest() {
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest) deepCopy() *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest {
	if m == nil {
		return nil
	}
	_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestCopy := &_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest{
		m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).deepCopy(),
	}
	_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestCopy.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = m
	return _S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestCopy
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest) String() string {
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