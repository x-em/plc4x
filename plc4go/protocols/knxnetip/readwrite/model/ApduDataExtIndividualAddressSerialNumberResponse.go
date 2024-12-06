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

// ApduDataExtIndividualAddressSerialNumberResponse is the corresponding interface of ApduDataExtIndividualAddressSerialNumberResponse
type ApduDataExtIndividualAddressSerialNumberResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// IsApduDataExtIndividualAddressSerialNumberResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtIndividualAddressSerialNumberResponse()
	// CreateBuilder creates a ApduDataExtIndividualAddressSerialNumberResponseBuilder
	CreateApduDataExtIndividualAddressSerialNumberResponseBuilder() ApduDataExtIndividualAddressSerialNumberResponseBuilder
}

// _ApduDataExtIndividualAddressSerialNumberResponse is the data-structure of this message
type _ApduDataExtIndividualAddressSerialNumberResponse struct {
	ApduDataExtContract
}

var _ ApduDataExtIndividualAddressSerialNumberResponse = (*_ApduDataExtIndividualAddressSerialNumberResponse)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtIndividualAddressSerialNumberResponse)(nil)

// NewApduDataExtIndividualAddressSerialNumberResponse factory function for _ApduDataExtIndividualAddressSerialNumberResponse
func NewApduDataExtIndividualAddressSerialNumberResponse(length uint8) *_ApduDataExtIndividualAddressSerialNumberResponse {
	_result := &_ApduDataExtIndividualAddressSerialNumberResponse{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataExtIndividualAddressSerialNumberResponseBuilder is a builder for ApduDataExtIndividualAddressSerialNumberResponse
type ApduDataExtIndividualAddressSerialNumberResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ApduDataExtIndividualAddressSerialNumberResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ApduDataExtBuilder
	// Build builds the ApduDataExtIndividualAddressSerialNumberResponse or returns an error if something is wrong
	Build() (ApduDataExtIndividualAddressSerialNumberResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataExtIndividualAddressSerialNumberResponse
}

// NewApduDataExtIndividualAddressSerialNumberResponseBuilder() creates a ApduDataExtIndividualAddressSerialNumberResponseBuilder
func NewApduDataExtIndividualAddressSerialNumberResponseBuilder() ApduDataExtIndividualAddressSerialNumberResponseBuilder {
	return &_ApduDataExtIndividualAddressSerialNumberResponseBuilder{_ApduDataExtIndividualAddressSerialNumberResponse: new(_ApduDataExtIndividualAddressSerialNumberResponse)}
}

type _ApduDataExtIndividualAddressSerialNumberResponseBuilder struct {
	*_ApduDataExtIndividualAddressSerialNumberResponse

	parentBuilder *_ApduDataExtBuilder

	err *utils.MultiError
}

var _ (ApduDataExtIndividualAddressSerialNumberResponseBuilder) = (*_ApduDataExtIndividualAddressSerialNumberResponseBuilder)(nil)

func (b *_ApduDataExtIndividualAddressSerialNumberResponseBuilder) setParent(contract ApduDataExtContract) {
	b.ApduDataExtContract = contract
	contract.(*_ApduDataExt)._SubType = b._ApduDataExtIndividualAddressSerialNumberResponse
}

func (b *_ApduDataExtIndividualAddressSerialNumberResponseBuilder) WithMandatoryFields() ApduDataExtIndividualAddressSerialNumberResponseBuilder {
	return b
}

func (b *_ApduDataExtIndividualAddressSerialNumberResponseBuilder) Build() (ApduDataExtIndividualAddressSerialNumberResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ApduDataExtIndividualAddressSerialNumberResponse.deepCopy(), nil
}

func (b *_ApduDataExtIndividualAddressSerialNumberResponseBuilder) MustBuild() ApduDataExtIndividualAddressSerialNumberResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ApduDataExtIndividualAddressSerialNumberResponseBuilder) Done() ApduDataExtBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewApduDataExtBuilder().(*_ApduDataExtBuilder)
	}
	return b.parentBuilder
}

func (b *_ApduDataExtIndividualAddressSerialNumberResponseBuilder) buildForApduDataExt() (ApduDataExt, error) {
	return b.Build()
}

func (b *_ApduDataExtIndividualAddressSerialNumberResponseBuilder) DeepCopy() any {
	_copy := b.CreateApduDataExtIndividualAddressSerialNumberResponseBuilder().(*_ApduDataExtIndividualAddressSerialNumberResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateApduDataExtIndividualAddressSerialNumberResponseBuilder creates a ApduDataExtIndividualAddressSerialNumberResponseBuilder
func (b *_ApduDataExtIndividualAddressSerialNumberResponse) CreateApduDataExtIndividualAddressSerialNumberResponseBuilder() ApduDataExtIndividualAddressSerialNumberResponseBuilder {
	if b == nil {
		return NewApduDataExtIndividualAddressSerialNumberResponseBuilder()
	}
	return &_ApduDataExtIndividualAddressSerialNumberResponseBuilder{_ApduDataExtIndividualAddressSerialNumberResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtIndividualAddressSerialNumberResponse) GetExtApciType() uint8 {
	return 0x1D
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtIndividualAddressSerialNumberResponse) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// Deprecated: use the interface for direct cast
func CastApduDataExtIndividualAddressSerialNumberResponse(structType any) ApduDataExtIndividualAddressSerialNumberResponse {
	if casted, ok := structType.(ApduDataExtIndividualAddressSerialNumberResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtIndividualAddressSerialNumberResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtIndividualAddressSerialNumberResponse) GetTypeName() string {
	return "ApduDataExtIndividualAddressSerialNumberResponse"
}

func (m *_ApduDataExtIndividualAddressSerialNumberResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtIndividualAddressSerialNumberResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtIndividualAddressSerialNumberResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtIndividualAddressSerialNumberResponse ApduDataExtIndividualAddressSerialNumberResponse, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtIndividualAddressSerialNumberResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtIndividualAddressSerialNumberResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtIndividualAddressSerialNumberResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtIndividualAddressSerialNumberResponse")
	}

	return m, nil
}

func (m *_ApduDataExtIndividualAddressSerialNumberResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtIndividualAddressSerialNumberResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtIndividualAddressSerialNumberResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtIndividualAddressSerialNumberResponse")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtIndividualAddressSerialNumberResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtIndividualAddressSerialNumberResponse")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtIndividualAddressSerialNumberResponse) IsApduDataExtIndividualAddressSerialNumberResponse() {
}

func (m *_ApduDataExtIndividualAddressSerialNumberResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtIndividualAddressSerialNumberResponse) deepCopy() *_ApduDataExtIndividualAddressSerialNumberResponse {
	if m == nil {
		return nil
	}
	_ApduDataExtIndividualAddressSerialNumberResponseCopy := &_ApduDataExtIndividualAddressSerialNumberResponse{
		m.ApduDataExtContract.(*_ApduDataExt).deepCopy(),
	}
	_ApduDataExtIndividualAddressSerialNumberResponseCopy.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtIndividualAddressSerialNumberResponseCopy
}

func (m *_ApduDataExtIndividualAddressSerialNumberResponse) String() string {
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
