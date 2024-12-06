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

// ModbusPDUDiagnosticRequest is the corresponding interface of ModbusPDUDiagnosticRequest
type ModbusPDUDiagnosticRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusPDU
	// GetSubFunction returns SubFunction (property field)
	GetSubFunction() uint16
	// GetData returns Data (property field)
	GetData() uint16
	// IsModbusPDUDiagnosticRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUDiagnosticRequest()
	// CreateBuilder creates a ModbusPDUDiagnosticRequestBuilder
	CreateModbusPDUDiagnosticRequestBuilder() ModbusPDUDiagnosticRequestBuilder
}

// _ModbusPDUDiagnosticRequest is the data-structure of this message
type _ModbusPDUDiagnosticRequest struct {
	ModbusPDUContract
	SubFunction uint16
	Data        uint16
}

var _ ModbusPDUDiagnosticRequest = (*_ModbusPDUDiagnosticRequest)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUDiagnosticRequest)(nil)

// NewModbusPDUDiagnosticRequest factory function for _ModbusPDUDiagnosticRequest
func NewModbusPDUDiagnosticRequest(subFunction uint16, data uint16) *_ModbusPDUDiagnosticRequest {
	_result := &_ModbusPDUDiagnosticRequest{
		ModbusPDUContract: NewModbusPDU(),
		SubFunction:       subFunction,
		Data:              data,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusPDUDiagnosticRequestBuilder is a builder for ModbusPDUDiagnosticRequest
type ModbusPDUDiagnosticRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(subFunction uint16, data uint16) ModbusPDUDiagnosticRequestBuilder
	// WithSubFunction adds SubFunction (property field)
	WithSubFunction(uint16) ModbusPDUDiagnosticRequestBuilder
	// WithData adds Data (property field)
	WithData(uint16) ModbusPDUDiagnosticRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ModbusPDUBuilder
	// Build builds the ModbusPDUDiagnosticRequest or returns an error if something is wrong
	Build() (ModbusPDUDiagnosticRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusPDUDiagnosticRequest
}

// NewModbusPDUDiagnosticRequestBuilder() creates a ModbusPDUDiagnosticRequestBuilder
func NewModbusPDUDiagnosticRequestBuilder() ModbusPDUDiagnosticRequestBuilder {
	return &_ModbusPDUDiagnosticRequestBuilder{_ModbusPDUDiagnosticRequest: new(_ModbusPDUDiagnosticRequest)}
}

type _ModbusPDUDiagnosticRequestBuilder struct {
	*_ModbusPDUDiagnosticRequest

	parentBuilder *_ModbusPDUBuilder

	err *utils.MultiError
}

var _ (ModbusPDUDiagnosticRequestBuilder) = (*_ModbusPDUDiagnosticRequestBuilder)(nil)

func (b *_ModbusPDUDiagnosticRequestBuilder) setParent(contract ModbusPDUContract) {
	b.ModbusPDUContract = contract
	contract.(*_ModbusPDU)._SubType = b._ModbusPDUDiagnosticRequest
}

func (b *_ModbusPDUDiagnosticRequestBuilder) WithMandatoryFields(subFunction uint16, data uint16) ModbusPDUDiagnosticRequestBuilder {
	return b.WithSubFunction(subFunction).WithData(data)
}

func (b *_ModbusPDUDiagnosticRequestBuilder) WithSubFunction(subFunction uint16) ModbusPDUDiagnosticRequestBuilder {
	b.SubFunction = subFunction
	return b
}

func (b *_ModbusPDUDiagnosticRequestBuilder) WithData(data uint16) ModbusPDUDiagnosticRequestBuilder {
	b.Data = data
	return b
}

func (b *_ModbusPDUDiagnosticRequestBuilder) Build() (ModbusPDUDiagnosticRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModbusPDUDiagnosticRequest.deepCopy(), nil
}

func (b *_ModbusPDUDiagnosticRequestBuilder) MustBuild() ModbusPDUDiagnosticRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ModbusPDUDiagnosticRequestBuilder) Done() ModbusPDUBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewModbusPDUBuilder().(*_ModbusPDUBuilder)
	}
	return b.parentBuilder
}

func (b *_ModbusPDUDiagnosticRequestBuilder) buildForModbusPDU() (ModbusPDU, error) {
	return b.Build()
}

func (b *_ModbusPDUDiagnosticRequestBuilder) DeepCopy() any {
	_copy := b.CreateModbusPDUDiagnosticRequestBuilder().(*_ModbusPDUDiagnosticRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModbusPDUDiagnosticRequestBuilder creates a ModbusPDUDiagnosticRequestBuilder
func (b *_ModbusPDUDiagnosticRequest) CreateModbusPDUDiagnosticRequestBuilder() ModbusPDUDiagnosticRequestBuilder {
	if b == nil {
		return NewModbusPDUDiagnosticRequestBuilder()
	}
	return &_ModbusPDUDiagnosticRequestBuilder{_ModbusPDUDiagnosticRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUDiagnosticRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUDiagnosticRequest) GetFunctionFlag() uint8 {
	return 0x08
}

func (m *_ModbusPDUDiagnosticRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUDiagnosticRequest) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUDiagnosticRequest) GetSubFunction() uint16 {
	return m.SubFunction
}

func (m *_ModbusPDUDiagnosticRequest) GetData() uint16 {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUDiagnosticRequest(structType any) ModbusPDUDiagnosticRequest {
	if casted, ok := structType.(ModbusPDUDiagnosticRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUDiagnosticRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUDiagnosticRequest) GetTypeName() string {
	return "ModbusPDUDiagnosticRequest"
}

func (m *_ModbusPDUDiagnosticRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Simple field (subFunction)
	lengthInBits += 16

	// Simple field (data)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUDiagnosticRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUDiagnosticRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUDiagnosticRequest ModbusPDUDiagnosticRequest, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUDiagnosticRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUDiagnosticRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	subFunction, err := ReadSimpleField(ctx, "subFunction", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subFunction' field"))
	}
	m.SubFunction = subFunction

	data, err := ReadSimpleField(ctx, "data", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("ModbusPDUDiagnosticRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUDiagnosticRequest")
	}

	return m, nil
}

func (m *_ModbusPDUDiagnosticRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUDiagnosticRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUDiagnosticRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUDiagnosticRequest")
		}

		if err := WriteSimpleField[uint16](ctx, "subFunction", m.GetSubFunction(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'subFunction' field")
		}

		if err := WriteSimpleField[uint16](ctx, "data", m.GetData(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUDiagnosticRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUDiagnosticRequest")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUDiagnosticRequest) IsModbusPDUDiagnosticRequest() {}

func (m *_ModbusPDUDiagnosticRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUDiagnosticRequest) deepCopy() *_ModbusPDUDiagnosticRequest {
	if m == nil {
		return nil
	}
	_ModbusPDUDiagnosticRequestCopy := &_ModbusPDUDiagnosticRequest{
		m.ModbusPDUContract.(*_ModbusPDU).deepCopy(),
		m.SubFunction,
		m.Data,
	}
	_ModbusPDUDiagnosticRequestCopy.ModbusPDUContract.(*_ModbusPDU)._SubType = m
	return _ModbusPDUDiagnosticRequestCopy
}

func (m *_ModbusPDUDiagnosticRequest) String() string {
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
