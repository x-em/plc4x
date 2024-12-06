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

// CallMethodRequest is the corresponding interface of CallMethodRequest
type CallMethodRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetObjectId returns ObjectId (property field)
	GetObjectId() NodeId
	// GetMethodId returns MethodId (property field)
	GetMethodId() NodeId
	// GetInputArguments returns InputArguments (property field)
	GetInputArguments() []Variant
	// IsCallMethodRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCallMethodRequest()
	// CreateBuilder creates a CallMethodRequestBuilder
	CreateCallMethodRequestBuilder() CallMethodRequestBuilder
}

// _CallMethodRequest is the data-structure of this message
type _CallMethodRequest struct {
	ExtensionObjectDefinitionContract
	ObjectId       NodeId
	MethodId       NodeId
	InputArguments []Variant
}

var _ CallMethodRequest = (*_CallMethodRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CallMethodRequest)(nil)

// NewCallMethodRequest factory function for _CallMethodRequest
func NewCallMethodRequest(objectId NodeId, methodId NodeId, inputArguments []Variant) *_CallMethodRequest {
	if objectId == nil {
		panic("objectId of type NodeId for CallMethodRequest must not be nil")
	}
	if methodId == nil {
		panic("methodId of type NodeId for CallMethodRequest must not be nil")
	}
	_result := &_CallMethodRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ObjectId:                          objectId,
		MethodId:                          methodId,
		InputArguments:                    inputArguments,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CallMethodRequestBuilder is a builder for CallMethodRequest
type CallMethodRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(objectId NodeId, methodId NodeId, inputArguments []Variant) CallMethodRequestBuilder
	// WithObjectId adds ObjectId (property field)
	WithObjectId(NodeId) CallMethodRequestBuilder
	// WithObjectIdBuilder adds ObjectId (property field) which is build by the builder
	WithObjectIdBuilder(func(NodeIdBuilder) NodeIdBuilder) CallMethodRequestBuilder
	// WithMethodId adds MethodId (property field)
	WithMethodId(NodeId) CallMethodRequestBuilder
	// WithMethodIdBuilder adds MethodId (property field) which is build by the builder
	WithMethodIdBuilder(func(NodeIdBuilder) NodeIdBuilder) CallMethodRequestBuilder
	// WithInputArguments adds InputArguments (property field)
	WithInputArguments(...Variant) CallMethodRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the CallMethodRequest or returns an error if something is wrong
	Build() (CallMethodRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CallMethodRequest
}

// NewCallMethodRequestBuilder() creates a CallMethodRequestBuilder
func NewCallMethodRequestBuilder() CallMethodRequestBuilder {
	return &_CallMethodRequestBuilder{_CallMethodRequest: new(_CallMethodRequest)}
}

type _CallMethodRequestBuilder struct {
	*_CallMethodRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (CallMethodRequestBuilder) = (*_CallMethodRequestBuilder)(nil)

func (b *_CallMethodRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._CallMethodRequest
}

func (b *_CallMethodRequestBuilder) WithMandatoryFields(objectId NodeId, methodId NodeId, inputArguments []Variant) CallMethodRequestBuilder {
	return b.WithObjectId(objectId).WithMethodId(methodId).WithInputArguments(inputArguments...)
}

func (b *_CallMethodRequestBuilder) WithObjectId(objectId NodeId) CallMethodRequestBuilder {
	b.ObjectId = objectId
	return b
}

func (b *_CallMethodRequestBuilder) WithObjectIdBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) CallMethodRequestBuilder {
	builder := builderSupplier(b.ObjectId.CreateNodeIdBuilder())
	var err error
	b.ObjectId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_CallMethodRequestBuilder) WithMethodId(methodId NodeId) CallMethodRequestBuilder {
	b.MethodId = methodId
	return b
}

func (b *_CallMethodRequestBuilder) WithMethodIdBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) CallMethodRequestBuilder {
	builder := builderSupplier(b.MethodId.CreateNodeIdBuilder())
	var err error
	b.MethodId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_CallMethodRequestBuilder) WithInputArguments(inputArguments ...Variant) CallMethodRequestBuilder {
	b.InputArguments = inputArguments
	return b
}

func (b *_CallMethodRequestBuilder) Build() (CallMethodRequest, error) {
	if b.ObjectId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'objectId' not set"))
	}
	if b.MethodId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'methodId' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CallMethodRequest.deepCopy(), nil
}

func (b *_CallMethodRequestBuilder) MustBuild() CallMethodRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CallMethodRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_CallMethodRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_CallMethodRequestBuilder) DeepCopy() any {
	_copy := b.CreateCallMethodRequestBuilder().(*_CallMethodRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCallMethodRequestBuilder creates a CallMethodRequestBuilder
func (b *_CallMethodRequest) CreateCallMethodRequestBuilder() CallMethodRequestBuilder {
	if b == nil {
		return NewCallMethodRequestBuilder()
	}
	return &_CallMethodRequestBuilder{_CallMethodRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CallMethodRequest) GetExtensionId() int32 {
	return int32(706)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CallMethodRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CallMethodRequest) GetObjectId() NodeId {
	return m.ObjectId
}

func (m *_CallMethodRequest) GetMethodId() NodeId {
	return m.MethodId
}

func (m *_CallMethodRequest) GetInputArguments() []Variant {
	return m.InputArguments
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCallMethodRequest(structType any) CallMethodRequest {
	if casted, ok := structType.(CallMethodRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CallMethodRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CallMethodRequest) GetTypeName() string {
	return "CallMethodRequest"
}

func (m *_CallMethodRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (objectId)
	lengthInBits += m.ObjectId.GetLengthInBits(ctx)

	// Simple field (methodId)
	lengthInBits += m.MethodId.GetLengthInBits(ctx)

	// Implicit Field (noOfInputArguments)
	lengthInBits += 32

	// Array field
	if len(m.InputArguments) > 0 {
		for _curItem, element := range m.InputArguments {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.InputArguments), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_CallMethodRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CallMethodRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__callMethodRequest CallMethodRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CallMethodRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CallMethodRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectId, err := ReadSimpleField[NodeId](ctx, "objectId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectId' field"))
	}
	m.ObjectId = objectId

	methodId, err := ReadSimpleField[NodeId](ctx, "methodId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'methodId' field"))
	}
	m.MethodId = methodId

	noOfInputArguments, err := ReadImplicitField[int32](ctx, "noOfInputArguments", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfInputArguments' field"))
	}
	_ = noOfInputArguments

	inputArguments, err := ReadCountArrayField[Variant](ctx, "inputArguments", ReadComplex[Variant](VariantParseWithBuffer, readBuffer), uint64(noOfInputArguments))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'inputArguments' field"))
	}
	m.InputArguments = inputArguments

	if closeErr := readBuffer.CloseContext("CallMethodRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CallMethodRequest")
	}

	return m, nil
}

func (m *_CallMethodRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CallMethodRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CallMethodRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CallMethodRequest")
		}

		if err := WriteSimpleField[NodeId](ctx, "objectId", m.GetObjectId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectId' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "methodId", m.GetMethodId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'methodId' field")
		}
		noOfInputArguments := int32(utils.InlineIf(bool((m.GetInputArguments()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetInputArguments()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfInputArguments", noOfInputArguments, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfInputArguments' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "inputArguments", m.GetInputArguments(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'inputArguments' field")
		}

		if popErr := writeBuffer.PopContext("CallMethodRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CallMethodRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CallMethodRequest) IsCallMethodRequest() {}

func (m *_CallMethodRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CallMethodRequest) deepCopy() *_CallMethodRequest {
	if m == nil {
		return nil
	}
	_CallMethodRequestCopy := &_CallMethodRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[NodeId](m.ObjectId),
		utils.DeepCopy[NodeId](m.MethodId),
		utils.DeepCopySlice[Variant, Variant](m.InputArguments),
	}
	_CallMethodRequestCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _CallMethodRequestCopy
}

func (m *_CallMethodRequest) String() string {
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