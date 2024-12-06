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

// BrowseRequest is the corresponding interface of BrowseRequest
type BrowseRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() RequestHeader
	// GetView returns View (property field)
	GetView() ViewDescription
	// GetRequestedMaxReferencesPerNode returns RequestedMaxReferencesPerNode (property field)
	GetRequestedMaxReferencesPerNode() uint32
	// GetNodesToBrowse returns NodesToBrowse (property field)
	GetNodesToBrowse() []BrowseDescription
	// IsBrowseRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBrowseRequest()
	// CreateBuilder creates a BrowseRequestBuilder
	CreateBrowseRequestBuilder() BrowseRequestBuilder
}

// _BrowseRequest is the data-structure of this message
type _BrowseRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader                 RequestHeader
	View                          ViewDescription
	RequestedMaxReferencesPerNode uint32
	NodesToBrowse                 []BrowseDescription
}

var _ BrowseRequest = (*_BrowseRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_BrowseRequest)(nil)

// NewBrowseRequest factory function for _BrowseRequest
func NewBrowseRequest(requestHeader RequestHeader, view ViewDescription, requestedMaxReferencesPerNode uint32, nodesToBrowse []BrowseDescription) *_BrowseRequest {
	if requestHeader == nil {
		panic("requestHeader of type RequestHeader for BrowseRequest must not be nil")
	}
	if view == nil {
		panic("view of type ViewDescription for BrowseRequest must not be nil")
	}
	_result := &_BrowseRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		View:                              view,
		RequestedMaxReferencesPerNode:     requestedMaxReferencesPerNode,
		NodesToBrowse:                     nodesToBrowse,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BrowseRequestBuilder is a builder for BrowseRequest
type BrowseRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader RequestHeader, view ViewDescription, requestedMaxReferencesPerNode uint32, nodesToBrowse []BrowseDescription) BrowseRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(RequestHeader) BrowseRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(RequestHeaderBuilder) RequestHeaderBuilder) BrowseRequestBuilder
	// WithView adds View (property field)
	WithView(ViewDescription) BrowseRequestBuilder
	// WithViewBuilder adds View (property field) which is build by the builder
	WithViewBuilder(func(ViewDescriptionBuilder) ViewDescriptionBuilder) BrowseRequestBuilder
	// WithRequestedMaxReferencesPerNode adds RequestedMaxReferencesPerNode (property field)
	WithRequestedMaxReferencesPerNode(uint32) BrowseRequestBuilder
	// WithNodesToBrowse adds NodesToBrowse (property field)
	WithNodesToBrowse(...BrowseDescription) BrowseRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the BrowseRequest or returns an error if something is wrong
	Build() (BrowseRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BrowseRequest
}

// NewBrowseRequestBuilder() creates a BrowseRequestBuilder
func NewBrowseRequestBuilder() BrowseRequestBuilder {
	return &_BrowseRequestBuilder{_BrowseRequest: new(_BrowseRequest)}
}

type _BrowseRequestBuilder struct {
	*_BrowseRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (BrowseRequestBuilder) = (*_BrowseRequestBuilder)(nil)

func (b *_BrowseRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._BrowseRequest
}

func (b *_BrowseRequestBuilder) WithMandatoryFields(requestHeader RequestHeader, view ViewDescription, requestedMaxReferencesPerNode uint32, nodesToBrowse []BrowseDescription) BrowseRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithView(view).WithRequestedMaxReferencesPerNode(requestedMaxReferencesPerNode).WithNodesToBrowse(nodesToBrowse...)
}

func (b *_BrowseRequestBuilder) WithRequestHeader(requestHeader RequestHeader) BrowseRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_BrowseRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(RequestHeaderBuilder) RequestHeaderBuilder) BrowseRequestBuilder {
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

func (b *_BrowseRequestBuilder) WithView(view ViewDescription) BrowseRequestBuilder {
	b.View = view
	return b
}

func (b *_BrowseRequestBuilder) WithViewBuilder(builderSupplier func(ViewDescriptionBuilder) ViewDescriptionBuilder) BrowseRequestBuilder {
	builder := builderSupplier(b.View.CreateViewDescriptionBuilder())
	var err error
	b.View, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ViewDescriptionBuilder failed"))
	}
	return b
}

func (b *_BrowseRequestBuilder) WithRequestedMaxReferencesPerNode(requestedMaxReferencesPerNode uint32) BrowseRequestBuilder {
	b.RequestedMaxReferencesPerNode = requestedMaxReferencesPerNode
	return b
}

func (b *_BrowseRequestBuilder) WithNodesToBrowse(nodesToBrowse ...BrowseDescription) BrowseRequestBuilder {
	b.NodesToBrowse = nodesToBrowse
	return b
}

func (b *_BrowseRequestBuilder) Build() (BrowseRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.View == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'view' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BrowseRequest.deepCopy(), nil
}

func (b *_BrowseRequestBuilder) MustBuild() BrowseRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BrowseRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_BrowseRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_BrowseRequestBuilder) DeepCopy() any {
	_copy := b.CreateBrowseRequestBuilder().(*_BrowseRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBrowseRequestBuilder creates a BrowseRequestBuilder
func (b *_BrowseRequest) CreateBrowseRequestBuilder() BrowseRequestBuilder {
	if b == nil {
		return NewBrowseRequestBuilder()
	}
	return &_BrowseRequestBuilder{_BrowseRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BrowseRequest) GetExtensionId() int32 {
	return int32(527)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BrowseRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BrowseRequest) GetRequestHeader() RequestHeader {
	return m.RequestHeader
}

func (m *_BrowseRequest) GetView() ViewDescription {
	return m.View
}

func (m *_BrowseRequest) GetRequestedMaxReferencesPerNode() uint32 {
	return m.RequestedMaxReferencesPerNode
}

func (m *_BrowseRequest) GetNodesToBrowse() []BrowseDescription {
	return m.NodesToBrowse
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBrowseRequest(structType any) BrowseRequest {
	if casted, ok := structType.(BrowseRequest); ok {
		return casted
	}
	if casted, ok := structType.(*BrowseRequest); ok {
		return *casted
	}
	return nil
}

func (m *_BrowseRequest) GetTypeName() string {
	return "BrowseRequest"
}

func (m *_BrowseRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (view)
	lengthInBits += m.View.GetLengthInBits(ctx)

	// Simple field (requestedMaxReferencesPerNode)
	lengthInBits += 32

	// Implicit Field (noOfNodesToBrowse)
	lengthInBits += 32

	// Array field
	if len(m.NodesToBrowse) > 0 {
		for _curItem, element := range m.NodesToBrowse {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NodesToBrowse), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_BrowseRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BrowseRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__browseRequest BrowseRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BrowseRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BrowseRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[RequestHeader](ctx, "requestHeader", ReadComplex[RequestHeader](ExtensionObjectDefinitionParseWithBufferProducer[RequestHeader]((int32)(int32(391))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	view, err := ReadSimpleField[ViewDescription](ctx, "view", ReadComplex[ViewDescription](ExtensionObjectDefinitionParseWithBufferProducer[ViewDescription]((int32)(int32(513))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'view' field"))
	}
	m.View = view

	requestedMaxReferencesPerNode, err := ReadSimpleField(ctx, "requestedMaxReferencesPerNode", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedMaxReferencesPerNode' field"))
	}
	m.RequestedMaxReferencesPerNode = requestedMaxReferencesPerNode

	noOfNodesToBrowse, err := ReadImplicitField[int32](ctx, "noOfNodesToBrowse", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfNodesToBrowse' field"))
	}
	_ = noOfNodesToBrowse

	nodesToBrowse, err := ReadCountArrayField[BrowseDescription](ctx, "nodesToBrowse", ReadComplex[BrowseDescription](ExtensionObjectDefinitionParseWithBufferProducer[BrowseDescription]((int32)(int32(516))), readBuffer), uint64(noOfNodesToBrowse))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodesToBrowse' field"))
	}
	m.NodesToBrowse = nodesToBrowse

	if closeErr := readBuffer.CloseContext("BrowseRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BrowseRequest")
	}

	return m, nil
}

func (m *_BrowseRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BrowseRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BrowseRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BrowseRequest")
		}

		if err := WriteSimpleField[RequestHeader](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[RequestHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[ViewDescription](ctx, "view", m.GetView(), WriteComplex[ViewDescription](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'view' field")
		}

		if err := WriteSimpleField[uint32](ctx, "requestedMaxReferencesPerNode", m.GetRequestedMaxReferencesPerNode(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedMaxReferencesPerNode' field")
		}
		noOfNodesToBrowse := int32(utils.InlineIf(bool((m.GetNodesToBrowse()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetNodesToBrowse()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfNodesToBrowse", noOfNodesToBrowse, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfNodesToBrowse' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "nodesToBrowse", m.GetNodesToBrowse(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'nodesToBrowse' field")
		}

		if popErr := writeBuffer.PopContext("BrowseRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BrowseRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BrowseRequest) IsBrowseRequest() {}

func (m *_BrowseRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BrowseRequest) deepCopy() *_BrowseRequest {
	if m == nil {
		return nil
	}
	_BrowseRequestCopy := &_BrowseRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[RequestHeader](m.RequestHeader),
		utils.DeepCopy[ViewDescription](m.View),
		m.RequestedMaxReferencesPerNode,
		utils.DeepCopySlice[BrowseDescription, BrowseDescription](m.NodesToBrowse),
	}
	_BrowseRequestCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _BrowseRequestCopy
}

func (m *_BrowseRequest) String() string {
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
