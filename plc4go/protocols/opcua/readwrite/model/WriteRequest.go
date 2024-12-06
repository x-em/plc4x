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

// WriteRequest is the corresponding interface of WriteRequest
type WriteRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() RequestHeader
	// GetNodesToWrite returns NodesToWrite (property field)
	GetNodesToWrite() []WriteValue
	// IsWriteRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsWriteRequest()
	// CreateBuilder creates a WriteRequestBuilder
	CreateWriteRequestBuilder() WriteRequestBuilder
}

// _WriteRequest is the data-structure of this message
type _WriteRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader RequestHeader
	NodesToWrite  []WriteValue
}

var _ WriteRequest = (*_WriteRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_WriteRequest)(nil)

// NewWriteRequest factory function for _WriteRequest
func NewWriteRequest(requestHeader RequestHeader, nodesToWrite []WriteValue) *_WriteRequest {
	if requestHeader == nil {
		panic("requestHeader of type RequestHeader for WriteRequest must not be nil")
	}
	_result := &_WriteRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		NodesToWrite:                      nodesToWrite,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// WriteRequestBuilder is a builder for WriteRequest
type WriteRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader RequestHeader, nodesToWrite []WriteValue) WriteRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(RequestHeader) WriteRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(RequestHeaderBuilder) RequestHeaderBuilder) WriteRequestBuilder
	// WithNodesToWrite adds NodesToWrite (property field)
	WithNodesToWrite(...WriteValue) WriteRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the WriteRequest or returns an error if something is wrong
	Build() (WriteRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() WriteRequest
}

// NewWriteRequestBuilder() creates a WriteRequestBuilder
func NewWriteRequestBuilder() WriteRequestBuilder {
	return &_WriteRequestBuilder{_WriteRequest: new(_WriteRequest)}
}

type _WriteRequestBuilder struct {
	*_WriteRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (WriteRequestBuilder) = (*_WriteRequestBuilder)(nil)

func (b *_WriteRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._WriteRequest
}

func (b *_WriteRequestBuilder) WithMandatoryFields(requestHeader RequestHeader, nodesToWrite []WriteValue) WriteRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithNodesToWrite(nodesToWrite...)
}

func (b *_WriteRequestBuilder) WithRequestHeader(requestHeader RequestHeader) WriteRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_WriteRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(RequestHeaderBuilder) RequestHeaderBuilder) WriteRequestBuilder {
	builder := builderSupplier(b.RequestHeader.CreateRequestHeaderBuilder())
	var err error
	b.RequestHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "RequestHeaderBuilder failed"))
	}
	return b
}

func (b *_WriteRequestBuilder) WithNodesToWrite(nodesToWrite ...WriteValue) WriteRequestBuilder {
	b.NodesToWrite = nodesToWrite
	return b
}

func (b *_WriteRequestBuilder) Build() (WriteRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._WriteRequest.deepCopy(), nil
}

func (b *_WriteRequestBuilder) MustBuild() WriteRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_WriteRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_WriteRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_WriteRequestBuilder) DeepCopy() any {
	_copy := b.CreateWriteRequestBuilder().(*_WriteRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateWriteRequestBuilder creates a WriteRequestBuilder
func (b *_WriteRequest) CreateWriteRequestBuilder() WriteRequestBuilder {
	if b == nil {
		return NewWriteRequestBuilder()
	}
	return &_WriteRequestBuilder{_WriteRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_WriteRequest) GetExtensionId() int32 {
	return int32(673)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_WriteRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_WriteRequest) GetRequestHeader() RequestHeader {
	return m.RequestHeader
}

func (m *_WriteRequest) GetNodesToWrite() []WriteValue {
	return m.NodesToWrite
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastWriteRequest(structType any) WriteRequest {
	if casted, ok := structType.(WriteRequest); ok {
		return casted
	}
	if casted, ok := structType.(*WriteRequest); ok {
		return *casted
	}
	return nil
}

func (m *_WriteRequest) GetTypeName() string {
	return "WriteRequest"
}

func (m *_WriteRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Implicit Field (noOfNodesToWrite)
	lengthInBits += 32

	// Array field
	if len(m.NodesToWrite) > 0 {
		for _curItem, element := range m.NodesToWrite {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NodesToWrite), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_WriteRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_WriteRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__writeRequest WriteRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("WriteRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for WriteRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[RequestHeader](ctx, "requestHeader", ReadComplex[RequestHeader](ExtensionObjectDefinitionParseWithBufferProducer[RequestHeader]((int32)(int32(391))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	noOfNodesToWrite, err := ReadImplicitField[int32](ctx, "noOfNodesToWrite", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfNodesToWrite' field"))
	}
	_ = noOfNodesToWrite

	nodesToWrite, err := ReadCountArrayField[WriteValue](ctx, "nodesToWrite", ReadComplex[WriteValue](ExtensionObjectDefinitionParseWithBufferProducer[WriteValue]((int32)(int32(670))), readBuffer), uint64(noOfNodesToWrite))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodesToWrite' field"))
	}
	m.NodesToWrite = nodesToWrite

	if closeErr := readBuffer.CloseContext("WriteRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for WriteRequest")
	}

	return m, nil
}

func (m *_WriteRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_WriteRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("WriteRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for WriteRequest")
		}

		if err := WriteSimpleField[RequestHeader](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[RequestHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}
		noOfNodesToWrite := int32(utils.InlineIf(bool((m.GetNodesToWrite()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetNodesToWrite()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfNodesToWrite", noOfNodesToWrite, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfNodesToWrite' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "nodesToWrite", m.GetNodesToWrite(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'nodesToWrite' field")
		}

		if popErr := writeBuffer.PopContext("WriteRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for WriteRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_WriteRequest) IsWriteRequest() {}

func (m *_WriteRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_WriteRequest) deepCopy() *_WriteRequest {
	if m == nil {
		return nil
	}
	_WriteRequestCopy := &_WriteRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[RequestHeader](m.RequestHeader),
		utils.DeepCopySlice[WriteValue, WriteValue](m.NodesToWrite),
	}
	_WriteRequestCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _WriteRequestCopy
}

func (m *_WriteRequest) String() string {
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