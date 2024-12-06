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

// CipConnectedResponse is the corresponding interface of CipConnectedResponse
type CipConnectedResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CipService
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetAdditionalStatusWords returns AdditionalStatusWords (property field)
	GetAdditionalStatusWords() uint8
	// GetData returns Data (property field)
	GetData() CIPDataConnected
	// IsCipConnectedResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCipConnectedResponse()
	// CreateBuilder creates a CipConnectedResponseBuilder
	CreateCipConnectedResponseBuilder() CipConnectedResponseBuilder
}

// _CipConnectedResponse is the data-structure of this message
type _CipConnectedResponse struct {
	CipServiceContract
	Status                uint8
	AdditionalStatusWords uint8
	Data                  CIPDataConnected
	// Reserved Fields
	reservedField0 *uint8
}

var _ CipConnectedResponse = (*_CipConnectedResponse)(nil)
var _ CipServiceRequirements = (*_CipConnectedResponse)(nil)

// NewCipConnectedResponse factory function for _CipConnectedResponse
func NewCipConnectedResponse(status uint8, additionalStatusWords uint8, data CIPDataConnected, serviceLen uint16) *_CipConnectedResponse {
	_result := &_CipConnectedResponse{
		CipServiceContract:    NewCipService(serviceLen),
		Status:                status,
		AdditionalStatusWords: additionalStatusWords,
		Data:                  data,
	}
	_result.CipServiceContract.(*_CipService)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CipConnectedResponseBuilder is a builder for CipConnectedResponse
type CipConnectedResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(status uint8, additionalStatusWords uint8) CipConnectedResponseBuilder
	// WithStatus adds Status (property field)
	WithStatus(uint8) CipConnectedResponseBuilder
	// WithAdditionalStatusWords adds AdditionalStatusWords (property field)
	WithAdditionalStatusWords(uint8) CipConnectedResponseBuilder
	// WithData adds Data (property field)
	WithOptionalData(CIPDataConnected) CipConnectedResponseBuilder
	// WithOptionalDataBuilder adds Data (property field) which is build by the builder
	WithOptionalDataBuilder(func(CIPDataConnectedBuilder) CIPDataConnectedBuilder) CipConnectedResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CipServiceBuilder
	// Build builds the CipConnectedResponse or returns an error if something is wrong
	Build() (CipConnectedResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CipConnectedResponse
}

// NewCipConnectedResponseBuilder() creates a CipConnectedResponseBuilder
func NewCipConnectedResponseBuilder() CipConnectedResponseBuilder {
	return &_CipConnectedResponseBuilder{_CipConnectedResponse: new(_CipConnectedResponse)}
}

type _CipConnectedResponseBuilder struct {
	*_CipConnectedResponse

	parentBuilder *_CipServiceBuilder

	err *utils.MultiError
}

var _ (CipConnectedResponseBuilder) = (*_CipConnectedResponseBuilder)(nil)

func (b *_CipConnectedResponseBuilder) setParent(contract CipServiceContract) {
	b.CipServiceContract = contract
	contract.(*_CipService)._SubType = b._CipConnectedResponse
}

func (b *_CipConnectedResponseBuilder) WithMandatoryFields(status uint8, additionalStatusWords uint8) CipConnectedResponseBuilder {
	return b.WithStatus(status).WithAdditionalStatusWords(additionalStatusWords)
}

func (b *_CipConnectedResponseBuilder) WithStatus(status uint8) CipConnectedResponseBuilder {
	b.Status = status
	return b
}

func (b *_CipConnectedResponseBuilder) WithAdditionalStatusWords(additionalStatusWords uint8) CipConnectedResponseBuilder {
	b.AdditionalStatusWords = additionalStatusWords
	return b
}

func (b *_CipConnectedResponseBuilder) WithOptionalData(data CIPDataConnected) CipConnectedResponseBuilder {
	b.Data = data
	return b
}

func (b *_CipConnectedResponseBuilder) WithOptionalDataBuilder(builderSupplier func(CIPDataConnectedBuilder) CIPDataConnectedBuilder) CipConnectedResponseBuilder {
	builder := builderSupplier(b.Data.CreateCIPDataConnectedBuilder())
	var err error
	b.Data, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "CIPDataConnectedBuilder failed"))
	}
	return b
}

func (b *_CipConnectedResponseBuilder) Build() (CipConnectedResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CipConnectedResponse.deepCopy(), nil
}

func (b *_CipConnectedResponseBuilder) MustBuild() CipConnectedResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CipConnectedResponseBuilder) Done() CipServiceBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCipServiceBuilder().(*_CipServiceBuilder)
	}
	return b.parentBuilder
}

func (b *_CipConnectedResponseBuilder) buildForCipService() (CipService, error) {
	return b.Build()
}

func (b *_CipConnectedResponseBuilder) DeepCopy() any {
	_copy := b.CreateCipConnectedResponseBuilder().(*_CipConnectedResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCipConnectedResponseBuilder creates a CipConnectedResponseBuilder
func (b *_CipConnectedResponse) CreateCipConnectedResponseBuilder() CipConnectedResponseBuilder {
	if b == nil {
		return NewCipConnectedResponseBuilder()
	}
	return &_CipConnectedResponseBuilder{_CipConnectedResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipConnectedResponse) GetService() uint8 {
	return 0x52
}

func (m *_CipConnectedResponse) GetResponse() bool {
	return bool(true)
}

func (m *_CipConnectedResponse) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipConnectedResponse) GetParent() CipServiceContract {
	return m.CipServiceContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipConnectedResponse) GetStatus() uint8 {
	return m.Status
}

func (m *_CipConnectedResponse) GetAdditionalStatusWords() uint8 {
	return m.AdditionalStatusWords
}

func (m *_CipConnectedResponse) GetData() CIPDataConnected {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCipConnectedResponse(structType any) CipConnectedResponse {
	if casted, ok := structType.(CipConnectedResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CipConnectedResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CipConnectedResponse) GetTypeName() string {
	return "CipConnectedResponse"
}

func (m *_CipConnectedResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (additionalStatusWords)
	lengthInBits += 8

	// Optional Field (data)
	if m.Data != nil {
		lengthInBits += m.Data.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_CipConnectedResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CipConnectedResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__cipConnectedResponse CipConnectedResponse, err error) {
	m.CipServiceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipConnectedResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipConnectedResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	status, err := ReadSimpleField(ctx, "status", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	additionalStatusWords, err := ReadSimpleField(ctx, "additionalStatusWords", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'additionalStatusWords' field"))
	}
	m.AdditionalStatusWords = additionalStatusWords

	var data CIPDataConnected
	_data, err := ReadOptionalField[CIPDataConnected](ctx, "data", ReadComplex[CIPDataConnected](CIPDataConnectedParseWithBuffer, readBuffer), bool(((serviceLen)-(4)) > (0)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	if _data != nil {
		data = *_data
		m.Data = data
	}

	if closeErr := readBuffer.CloseContext("CipConnectedResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipConnectedResponse")
	}

	return m, nil
}

func (m *_CipConnectedResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipConnectedResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipConnectedResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipConnectedResponse")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[uint8](ctx, "status", m.GetStatus(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'status' field")
		}

		if err := WriteSimpleField[uint8](ctx, "additionalStatusWords", m.GetAdditionalStatusWords(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'additionalStatusWords' field")
		}

		if err := WriteOptionalField[CIPDataConnected](ctx, "data", GetRef(m.GetData()), WriteComplex[CIPDataConnected](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("CipConnectedResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipConnectedResponse")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipConnectedResponse) IsCipConnectedResponse() {}

func (m *_CipConnectedResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CipConnectedResponse) deepCopy() *_CipConnectedResponse {
	if m == nil {
		return nil
	}
	_CipConnectedResponseCopy := &_CipConnectedResponse{
		m.CipServiceContract.(*_CipService).deepCopy(),
		m.Status,
		m.AdditionalStatusWords,
		utils.DeepCopy[CIPDataConnected](m.Data),
		m.reservedField0,
	}
	_CipConnectedResponseCopy.CipServiceContract.(*_CipService)._SubType = m
	return _CipConnectedResponseCopy
}

func (m *_CipConnectedResponse) String() string {
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