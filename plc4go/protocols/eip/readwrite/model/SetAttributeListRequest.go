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

// SetAttributeListRequest is the corresponding interface of SetAttributeListRequest
type SetAttributeListRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CipService
	// IsSetAttributeListRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSetAttributeListRequest()
	// CreateBuilder creates a SetAttributeListRequestBuilder
	CreateSetAttributeListRequestBuilder() SetAttributeListRequestBuilder
}

// _SetAttributeListRequest is the data-structure of this message
type _SetAttributeListRequest struct {
	CipServiceContract
}

var _ SetAttributeListRequest = (*_SetAttributeListRequest)(nil)
var _ CipServiceRequirements = (*_SetAttributeListRequest)(nil)

// NewSetAttributeListRequest factory function for _SetAttributeListRequest
func NewSetAttributeListRequest(serviceLen uint16) *_SetAttributeListRequest {
	_result := &_SetAttributeListRequest{
		CipServiceContract: NewCipService(serviceLen),
	}
	_result.CipServiceContract.(*_CipService)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SetAttributeListRequestBuilder is a builder for SetAttributeListRequest
type SetAttributeListRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SetAttributeListRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CipServiceBuilder
	// Build builds the SetAttributeListRequest or returns an error if something is wrong
	Build() (SetAttributeListRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SetAttributeListRequest
}

// NewSetAttributeListRequestBuilder() creates a SetAttributeListRequestBuilder
func NewSetAttributeListRequestBuilder() SetAttributeListRequestBuilder {
	return &_SetAttributeListRequestBuilder{_SetAttributeListRequest: new(_SetAttributeListRequest)}
}

type _SetAttributeListRequestBuilder struct {
	*_SetAttributeListRequest

	parentBuilder *_CipServiceBuilder

	err *utils.MultiError
}

var _ (SetAttributeListRequestBuilder) = (*_SetAttributeListRequestBuilder)(nil)

func (b *_SetAttributeListRequestBuilder) setParent(contract CipServiceContract) {
	b.CipServiceContract = contract
	contract.(*_CipService)._SubType = b._SetAttributeListRequest
}

func (b *_SetAttributeListRequestBuilder) WithMandatoryFields() SetAttributeListRequestBuilder {
	return b
}

func (b *_SetAttributeListRequestBuilder) Build() (SetAttributeListRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SetAttributeListRequest.deepCopy(), nil
}

func (b *_SetAttributeListRequestBuilder) MustBuild() SetAttributeListRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SetAttributeListRequestBuilder) Done() CipServiceBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCipServiceBuilder().(*_CipServiceBuilder)
	}
	return b.parentBuilder
}

func (b *_SetAttributeListRequestBuilder) buildForCipService() (CipService, error) {
	return b.Build()
}

func (b *_SetAttributeListRequestBuilder) DeepCopy() any {
	_copy := b.CreateSetAttributeListRequestBuilder().(*_SetAttributeListRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSetAttributeListRequestBuilder creates a SetAttributeListRequestBuilder
func (b *_SetAttributeListRequest) CreateSetAttributeListRequestBuilder() SetAttributeListRequestBuilder {
	if b == nil {
		return NewSetAttributeListRequestBuilder()
	}
	return &_SetAttributeListRequestBuilder{_SetAttributeListRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SetAttributeListRequest) GetService() uint8 {
	return 0x04
}

func (m *_SetAttributeListRequest) GetResponse() bool {
	return bool(false)
}

func (m *_SetAttributeListRequest) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SetAttributeListRequest) GetParent() CipServiceContract {
	return m.CipServiceContract
}

// Deprecated: use the interface for direct cast
func CastSetAttributeListRequest(structType any) SetAttributeListRequest {
	if casted, ok := structType.(SetAttributeListRequest); ok {
		return casted
	}
	if casted, ok := structType.(*SetAttributeListRequest); ok {
		return *casted
	}
	return nil
}

func (m *_SetAttributeListRequest) GetTypeName() string {
	return "SetAttributeListRequest"
}

func (m *_SetAttributeListRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SetAttributeListRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SetAttributeListRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__setAttributeListRequest SetAttributeListRequest, err error) {
	m.CipServiceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SetAttributeListRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SetAttributeListRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SetAttributeListRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SetAttributeListRequest")
	}

	return m, nil
}

func (m *_SetAttributeListRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SetAttributeListRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SetAttributeListRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SetAttributeListRequest")
		}

		if popErr := writeBuffer.PopContext("SetAttributeListRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SetAttributeListRequest")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SetAttributeListRequest) IsSetAttributeListRequest() {}

func (m *_SetAttributeListRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SetAttributeListRequest) deepCopy() *_SetAttributeListRequest {
	if m == nil {
		return nil
	}
	_SetAttributeListRequestCopy := &_SetAttributeListRequest{
		m.CipServiceContract.(*_CipService).deepCopy(),
	}
	_SetAttributeListRequestCopy.CipServiceContract.(*_CipService)._SubType = m
	return _SetAttributeListRequestCopy
}

func (m *_SetAttributeListRequest) String() string {
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