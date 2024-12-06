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

// BinaryExtensionObjectWithMask is the corresponding interface of BinaryExtensionObjectWithMask
type BinaryExtensionObjectWithMask interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectWithMask
	// GetBody returns Body (property field)
	GetBody() ExtensionObjectDefinition
	// IsBinaryExtensionObjectWithMask is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBinaryExtensionObjectWithMask()
	// CreateBuilder creates a BinaryExtensionObjectWithMaskBuilder
	CreateBinaryExtensionObjectWithMaskBuilder() BinaryExtensionObjectWithMaskBuilder
}

// _BinaryExtensionObjectWithMask is the data-structure of this message
type _BinaryExtensionObjectWithMask struct {
	ExtensionObjectWithMaskContract
	Body ExtensionObjectDefinition
}

var _ BinaryExtensionObjectWithMask = (*_BinaryExtensionObjectWithMask)(nil)
var _ ExtensionObjectWithMaskRequirements = (*_BinaryExtensionObjectWithMask)(nil)

// NewBinaryExtensionObjectWithMask factory function for _BinaryExtensionObjectWithMask
func NewBinaryExtensionObjectWithMask(typeId ExpandedNodeId, encodingMask ExtensionObjectEncodingMask, body ExtensionObjectDefinition, extensionId int32, includeEncodingMask bool) *_BinaryExtensionObjectWithMask {
	if body == nil {
		panic("body of type ExtensionObjectDefinition for BinaryExtensionObjectWithMask must not be nil")
	}
	_result := &_BinaryExtensionObjectWithMask{
		ExtensionObjectWithMaskContract: NewExtensionObjectWithMask(typeId, encodingMask, extensionId),
		Body:                            body,
	}
	_result.ExtensionObjectWithMaskContract.(*_ExtensionObjectWithMask)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BinaryExtensionObjectWithMaskBuilder is a builder for BinaryExtensionObjectWithMask
type BinaryExtensionObjectWithMaskBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(body ExtensionObjectDefinition) BinaryExtensionObjectWithMaskBuilder
	// WithBody adds Body (property field)
	WithBody(ExtensionObjectDefinition) BinaryExtensionObjectWithMaskBuilder
	// WithBodyBuilder adds Body (property field) which is build by the builder
	WithBodyBuilder(func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) BinaryExtensionObjectWithMaskBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectWithMaskBuilder
	// Build builds the BinaryExtensionObjectWithMask or returns an error if something is wrong
	Build() (BinaryExtensionObjectWithMask, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BinaryExtensionObjectWithMask
}

// NewBinaryExtensionObjectWithMaskBuilder() creates a BinaryExtensionObjectWithMaskBuilder
func NewBinaryExtensionObjectWithMaskBuilder() BinaryExtensionObjectWithMaskBuilder {
	return &_BinaryExtensionObjectWithMaskBuilder{_BinaryExtensionObjectWithMask: new(_BinaryExtensionObjectWithMask)}
}

type _BinaryExtensionObjectWithMaskBuilder struct {
	*_BinaryExtensionObjectWithMask

	parentBuilder *_ExtensionObjectWithMaskBuilder

	err *utils.MultiError
}

var _ (BinaryExtensionObjectWithMaskBuilder) = (*_BinaryExtensionObjectWithMaskBuilder)(nil)

func (b *_BinaryExtensionObjectWithMaskBuilder) setParent(contract ExtensionObjectWithMaskContract) {
	b.ExtensionObjectWithMaskContract = contract
	contract.(*_ExtensionObjectWithMask)._SubType = b._BinaryExtensionObjectWithMask
}

func (b *_BinaryExtensionObjectWithMaskBuilder) WithMandatoryFields(body ExtensionObjectDefinition) BinaryExtensionObjectWithMaskBuilder {
	return b.WithBody(body)
}

func (b *_BinaryExtensionObjectWithMaskBuilder) WithBody(body ExtensionObjectDefinition) BinaryExtensionObjectWithMaskBuilder {
	b.Body = body
	return b
}

func (b *_BinaryExtensionObjectWithMaskBuilder) WithBodyBuilder(builderSupplier func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) BinaryExtensionObjectWithMaskBuilder {
	builder := builderSupplier(b.Body.CreateExtensionObjectDefinitionBuilder())
	var err error
	b.Body, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectDefinitionBuilder failed"))
	}
	return b
}

func (b *_BinaryExtensionObjectWithMaskBuilder) Build() (BinaryExtensionObjectWithMask, error) {
	if b.Body == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'body' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BinaryExtensionObjectWithMask.deepCopy(), nil
}

func (b *_BinaryExtensionObjectWithMaskBuilder) MustBuild() BinaryExtensionObjectWithMask {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BinaryExtensionObjectWithMaskBuilder) Done() ExtensionObjectWithMaskBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectWithMaskBuilder().(*_ExtensionObjectWithMaskBuilder)
	}
	return b.parentBuilder
}

func (b *_BinaryExtensionObjectWithMaskBuilder) buildForExtensionObjectWithMask() (ExtensionObjectWithMask, error) {
	return b.Build()
}

func (b *_BinaryExtensionObjectWithMaskBuilder) DeepCopy() any {
	_copy := b.CreateBinaryExtensionObjectWithMaskBuilder().(*_BinaryExtensionObjectWithMaskBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBinaryExtensionObjectWithMaskBuilder creates a BinaryExtensionObjectWithMaskBuilder
func (b *_BinaryExtensionObjectWithMask) CreateBinaryExtensionObjectWithMaskBuilder() BinaryExtensionObjectWithMaskBuilder {
	if b == nil {
		return NewBinaryExtensionObjectWithMaskBuilder()
	}
	return &_BinaryExtensionObjectWithMaskBuilder{_BinaryExtensionObjectWithMask: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BinaryExtensionObjectWithMask) GetEncodingMaskXmlBody() bool {
	return bool(false)
}

func (m *_BinaryExtensionObjectWithMask) GetEncodingMaskBinaryBody() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BinaryExtensionObjectWithMask) GetParent() ExtensionObjectWithMaskContract {
	return m.ExtensionObjectWithMaskContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BinaryExtensionObjectWithMask) GetBody() ExtensionObjectDefinition {
	return m.Body
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBinaryExtensionObjectWithMask(structType any) BinaryExtensionObjectWithMask {
	if casted, ok := structType.(BinaryExtensionObjectWithMask); ok {
		return casted
	}
	if casted, ok := structType.(*BinaryExtensionObjectWithMask); ok {
		return *casted
	}
	return nil
}

func (m *_BinaryExtensionObjectWithMask) GetTypeName() string {
	return "BinaryExtensionObjectWithMask"
}

func (m *_BinaryExtensionObjectWithMask) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectWithMaskContract.(*_ExtensionObjectWithMask).getLengthInBits(ctx))

	// Implicit Field (bodyLength)
	lengthInBits += 32

	// Simple field (body)
	lengthInBits += m.Body.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BinaryExtensionObjectWithMask) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BinaryExtensionObjectWithMask) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectWithMask, extensionId int32, includeEncodingMask bool) (__binaryExtensionObjectWithMask BinaryExtensionObjectWithMask, err error) {
	m.ExtensionObjectWithMaskContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BinaryExtensionObjectWithMask"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BinaryExtensionObjectWithMask")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bodyLength, err := ReadImplicitField[int32](ctx, "bodyLength", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bodyLength' field"))
	}
	_ = bodyLength

	body, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "body", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((int32)(extensionId)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'body' field"))
	}
	m.Body = body

	if closeErr := readBuffer.CloseContext("BinaryExtensionObjectWithMask"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BinaryExtensionObjectWithMask")
	}

	return m, nil
}

func (m *_BinaryExtensionObjectWithMask) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BinaryExtensionObjectWithMask) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BinaryExtensionObjectWithMask"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BinaryExtensionObjectWithMask")
		}
		bodyLength := int32(utils.InlineIf(bool((m.GetBody()) == (nil)), func() any { return int32(int32(0)) }, func() any { return int32(m.GetBody().GetLengthInBytes(ctx)) }).(int32))
		if err := WriteImplicitField(ctx, "bodyLength", bodyLength, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'bodyLength' field")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "body", m.GetBody(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'body' field")
		}

		if popErr := writeBuffer.PopContext("BinaryExtensionObjectWithMask"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BinaryExtensionObjectWithMask")
		}
		return nil
	}
	return m.ExtensionObjectWithMaskContract.(*_ExtensionObjectWithMask).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BinaryExtensionObjectWithMask) IsBinaryExtensionObjectWithMask() {}

func (m *_BinaryExtensionObjectWithMask) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BinaryExtensionObjectWithMask) deepCopy() *_BinaryExtensionObjectWithMask {
	if m == nil {
		return nil
	}
	_BinaryExtensionObjectWithMaskCopy := &_BinaryExtensionObjectWithMask{
		m.ExtensionObjectWithMaskContract.(*_ExtensionObjectWithMask).deepCopy(),
		utils.DeepCopy[ExtensionObjectDefinition](m.Body),
	}
	_BinaryExtensionObjectWithMaskCopy.ExtensionObjectWithMaskContract.(*_ExtensionObjectWithMask)._SubType = m
	return _BinaryExtensionObjectWithMaskCopy
}

func (m *_BinaryExtensionObjectWithMask) String() string {
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
