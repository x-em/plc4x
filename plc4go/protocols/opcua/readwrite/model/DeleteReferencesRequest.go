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

// DeleteReferencesRequest is the corresponding interface of DeleteReferencesRequest
type DeleteReferencesRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() RequestHeader
	// GetReferencesToDelete returns ReferencesToDelete (property field)
	GetReferencesToDelete() []DeleteReferencesItem
	// IsDeleteReferencesRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDeleteReferencesRequest()
	// CreateBuilder creates a DeleteReferencesRequestBuilder
	CreateDeleteReferencesRequestBuilder() DeleteReferencesRequestBuilder
}

// _DeleteReferencesRequest is the data-structure of this message
type _DeleteReferencesRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader      RequestHeader
	ReferencesToDelete []DeleteReferencesItem
}

var _ DeleteReferencesRequest = (*_DeleteReferencesRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DeleteReferencesRequest)(nil)

// NewDeleteReferencesRequest factory function for _DeleteReferencesRequest
func NewDeleteReferencesRequest(requestHeader RequestHeader, referencesToDelete []DeleteReferencesItem) *_DeleteReferencesRequest {
	if requestHeader == nil {
		panic("requestHeader of type RequestHeader for DeleteReferencesRequest must not be nil")
	}
	_result := &_DeleteReferencesRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		ReferencesToDelete:                referencesToDelete,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DeleteReferencesRequestBuilder is a builder for DeleteReferencesRequest
type DeleteReferencesRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader RequestHeader, referencesToDelete []DeleteReferencesItem) DeleteReferencesRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(RequestHeader) DeleteReferencesRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(RequestHeaderBuilder) RequestHeaderBuilder) DeleteReferencesRequestBuilder
	// WithReferencesToDelete adds ReferencesToDelete (property field)
	WithReferencesToDelete(...DeleteReferencesItem) DeleteReferencesRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the DeleteReferencesRequest or returns an error if something is wrong
	Build() (DeleteReferencesRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DeleteReferencesRequest
}

// NewDeleteReferencesRequestBuilder() creates a DeleteReferencesRequestBuilder
func NewDeleteReferencesRequestBuilder() DeleteReferencesRequestBuilder {
	return &_DeleteReferencesRequestBuilder{_DeleteReferencesRequest: new(_DeleteReferencesRequest)}
}

type _DeleteReferencesRequestBuilder struct {
	*_DeleteReferencesRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (DeleteReferencesRequestBuilder) = (*_DeleteReferencesRequestBuilder)(nil)

func (b *_DeleteReferencesRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._DeleteReferencesRequest
}

func (b *_DeleteReferencesRequestBuilder) WithMandatoryFields(requestHeader RequestHeader, referencesToDelete []DeleteReferencesItem) DeleteReferencesRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithReferencesToDelete(referencesToDelete...)
}

func (b *_DeleteReferencesRequestBuilder) WithRequestHeader(requestHeader RequestHeader) DeleteReferencesRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_DeleteReferencesRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(RequestHeaderBuilder) RequestHeaderBuilder) DeleteReferencesRequestBuilder {
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

func (b *_DeleteReferencesRequestBuilder) WithReferencesToDelete(referencesToDelete ...DeleteReferencesItem) DeleteReferencesRequestBuilder {
	b.ReferencesToDelete = referencesToDelete
	return b
}

func (b *_DeleteReferencesRequestBuilder) Build() (DeleteReferencesRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DeleteReferencesRequest.deepCopy(), nil
}

func (b *_DeleteReferencesRequestBuilder) MustBuild() DeleteReferencesRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DeleteReferencesRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_DeleteReferencesRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_DeleteReferencesRequestBuilder) DeepCopy() any {
	_copy := b.CreateDeleteReferencesRequestBuilder().(*_DeleteReferencesRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDeleteReferencesRequestBuilder creates a DeleteReferencesRequestBuilder
func (b *_DeleteReferencesRequest) CreateDeleteReferencesRequestBuilder() DeleteReferencesRequestBuilder {
	if b == nil {
		return NewDeleteReferencesRequestBuilder()
	}
	return &_DeleteReferencesRequestBuilder{_DeleteReferencesRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeleteReferencesRequest) GetExtensionId() int32 {
	return int32(506)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeleteReferencesRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeleteReferencesRequest) GetRequestHeader() RequestHeader {
	return m.RequestHeader
}

func (m *_DeleteReferencesRequest) GetReferencesToDelete() []DeleteReferencesItem {
	return m.ReferencesToDelete
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDeleteReferencesRequest(structType any) DeleteReferencesRequest {
	if casted, ok := structType.(DeleteReferencesRequest); ok {
		return casted
	}
	if casted, ok := structType.(*DeleteReferencesRequest); ok {
		return *casted
	}
	return nil
}

func (m *_DeleteReferencesRequest) GetTypeName() string {
	return "DeleteReferencesRequest"
}

func (m *_DeleteReferencesRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Implicit Field (noOfReferencesToDelete)
	lengthInBits += 32

	// Array field
	if len(m.ReferencesToDelete) > 0 {
		for _curItem, element := range m.ReferencesToDelete {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ReferencesToDelete), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_DeleteReferencesRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DeleteReferencesRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__deleteReferencesRequest DeleteReferencesRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeleteReferencesRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeleteReferencesRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[RequestHeader](ctx, "requestHeader", ReadComplex[RequestHeader](ExtensionObjectDefinitionParseWithBufferProducer[RequestHeader]((int32)(int32(391))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	noOfReferencesToDelete, err := ReadImplicitField[int32](ctx, "noOfReferencesToDelete", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfReferencesToDelete' field"))
	}
	_ = noOfReferencesToDelete

	referencesToDelete, err := ReadCountArrayField[DeleteReferencesItem](ctx, "referencesToDelete", ReadComplex[DeleteReferencesItem](ExtensionObjectDefinitionParseWithBufferProducer[DeleteReferencesItem]((int32)(int32(387))), readBuffer), uint64(noOfReferencesToDelete))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referencesToDelete' field"))
	}
	m.ReferencesToDelete = referencesToDelete

	if closeErr := readBuffer.CloseContext("DeleteReferencesRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeleteReferencesRequest")
	}

	return m, nil
}

func (m *_DeleteReferencesRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeleteReferencesRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeleteReferencesRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeleteReferencesRequest")
		}

		if err := WriteSimpleField[RequestHeader](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[RequestHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}
		noOfReferencesToDelete := int32(utils.InlineIf(bool((m.GetReferencesToDelete()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetReferencesToDelete()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfReferencesToDelete", noOfReferencesToDelete, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfReferencesToDelete' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "referencesToDelete", m.GetReferencesToDelete(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'referencesToDelete' field")
		}

		if popErr := writeBuffer.PopContext("DeleteReferencesRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeleteReferencesRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DeleteReferencesRequest) IsDeleteReferencesRequest() {}

func (m *_DeleteReferencesRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DeleteReferencesRequest) deepCopy() *_DeleteReferencesRequest {
	if m == nil {
		return nil
	}
	_DeleteReferencesRequestCopy := &_DeleteReferencesRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[RequestHeader](m.RequestHeader),
		utils.DeepCopySlice[DeleteReferencesItem, DeleteReferencesItem](m.ReferencesToDelete),
	}
	_DeleteReferencesRequestCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _DeleteReferencesRequestCopy
}

func (m *_DeleteReferencesRequest) String() string {
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
