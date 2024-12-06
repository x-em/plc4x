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

// MediaTransportControlDataCategoryName is the corresponding interface of MediaTransportControlDataCategoryName
type MediaTransportControlDataCategoryName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MediaTransportControlData
	// GetCategoryName returns CategoryName (property field)
	GetCategoryName() string
	// IsMediaTransportControlDataCategoryName is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMediaTransportControlDataCategoryName()
	// CreateBuilder creates a MediaTransportControlDataCategoryNameBuilder
	CreateMediaTransportControlDataCategoryNameBuilder() MediaTransportControlDataCategoryNameBuilder
}

// _MediaTransportControlDataCategoryName is the data-structure of this message
type _MediaTransportControlDataCategoryName struct {
	MediaTransportControlDataContract
	CategoryName string
}

var _ MediaTransportControlDataCategoryName = (*_MediaTransportControlDataCategoryName)(nil)
var _ MediaTransportControlDataRequirements = (*_MediaTransportControlDataCategoryName)(nil)

// NewMediaTransportControlDataCategoryName factory function for _MediaTransportControlDataCategoryName
func NewMediaTransportControlDataCategoryName(commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte, categoryName string) *_MediaTransportControlDataCategoryName {
	_result := &_MediaTransportControlDataCategoryName{
		MediaTransportControlDataContract: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
		CategoryName:                      categoryName,
	}
	_result.MediaTransportControlDataContract.(*_MediaTransportControlData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MediaTransportControlDataCategoryNameBuilder is a builder for MediaTransportControlDataCategoryName
type MediaTransportControlDataCategoryNameBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(categoryName string) MediaTransportControlDataCategoryNameBuilder
	// WithCategoryName adds CategoryName (property field)
	WithCategoryName(string) MediaTransportControlDataCategoryNameBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() MediaTransportControlDataBuilder
	// Build builds the MediaTransportControlDataCategoryName or returns an error if something is wrong
	Build() (MediaTransportControlDataCategoryName, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MediaTransportControlDataCategoryName
}

// NewMediaTransportControlDataCategoryNameBuilder() creates a MediaTransportControlDataCategoryNameBuilder
func NewMediaTransportControlDataCategoryNameBuilder() MediaTransportControlDataCategoryNameBuilder {
	return &_MediaTransportControlDataCategoryNameBuilder{_MediaTransportControlDataCategoryName: new(_MediaTransportControlDataCategoryName)}
}

type _MediaTransportControlDataCategoryNameBuilder struct {
	*_MediaTransportControlDataCategoryName

	parentBuilder *_MediaTransportControlDataBuilder

	err *utils.MultiError
}

var _ (MediaTransportControlDataCategoryNameBuilder) = (*_MediaTransportControlDataCategoryNameBuilder)(nil)

func (b *_MediaTransportControlDataCategoryNameBuilder) setParent(contract MediaTransportControlDataContract) {
	b.MediaTransportControlDataContract = contract
	contract.(*_MediaTransportControlData)._SubType = b._MediaTransportControlDataCategoryName
}

func (b *_MediaTransportControlDataCategoryNameBuilder) WithMandatoryFields(categoryName string) MediaTransportControlDataCategoryNameBuilder {
	return b.WithCategoryName(categoryName)
}

func (b *_MediaTransportControlDataCategoryNameBuilder) WithCategoryName(categoryName string) MediaTransportControlDataCategoryNameBuilder {
	b.CategoryName = categoryName
	return b
}

func (b *_MediaTransportControlDataCategoryNameBuilder) Build() (MediaTransportControlDataCategoryName, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MediaTransportControlDataCategoryName.deepCopy(), nil
}

func (b *_MediaTransportControlDataCategoryNameBuilder) MustBuild() MediaTransportControlDataCategoryName {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_MediaTransportControlDataCategoryNameBuilder) Done() MediaTransportControlDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewMediaTransportControlDataBuilder().(*_MediaTransportControlDataBuilder)
	}
	return b.parentBuilder
}

func (b *_MediaTransportControlDataCategoryNameBuilder) buildForMediaTransportControlData() (MediaTransportControlData, error) {
	return b.Build()
}

func (b *_MediaTransportControlDataCategoryNameBuilder) DeepCopy() any {
	_copy := b.CreateMediaTransportControlDataCategoryNameBuilder().(*_MediaTransportControlDataCategoryNameBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMediaTransportControlDataCategoryNameBuilder creates a MediaTransportControlDataCategoryNameBuilder
func (b *_MediaTransportControlDataCategoryName) CreateMediaTransportControlDataCategoryNameBuilder() MediaTransportControlDataCategoryNameBuilder {
	if b == nil {
		return NewMediaTransportControlDataCategoryNameBuilder()
	}
	return &_MediaTransportControlDataCategoryNameBuilder{_MediaTransportControlDataCategoryName: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataCategoryName) GetParent() MediaTransportControlDataContract {
	return m.MediaTransportControlDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataCategoryName) GetCategoryName() string {
	return m.CategoryName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataCategoryName(structType any) MediaTransportControlDataCategoryName {
	if casted, ok := structType.(MediaTransportControlDataCategoryName); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataCategoryName); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataCategoryName) GetTypeName() string {
	return "MediaTransportControlDataCategoryName"
}

func (m *_MediaTransportControlDataCategoryName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MediaTransportControlDataContract.(*_MediaTransportControlData).getLengthInBits(ctx))

	// Simple field (categoryName)
	lengthInBits += uint16(int32((int32(m.GetCommandTypeContainer().NumBytes()) - int32(int32(1)))) * int32(int32(8)))

	return lengthInBits
}

func (m *_MediaTransportControlDataCategoryName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MediaTransportControlDataCategoryName) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MediaTransportControlData, commandTypeContainer MediaTransportControlCommandTypeContainer) (__mediaTransportControlDataCategoryName MediaTransportControlDataCategoryName, err error) {
	m.MediaTransportControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataCategoryName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataCategoryName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	categoryName, err := ReadSimpleField(ctx, "categoryName", ReadString(readBuffer, uint32(int32((int32(commandTypeContainer.NumBytes())-int32(int32(1))))*int32(int32(8)))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'categoryName' field"))
	}
	m.CategoryName = categoryName

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataCategoryName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataCategoryName")
	}

	return m, nil
}

func (m *_MediaTransportControlDataCategoryName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataCategoryName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataCategoryName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataCategoryName")
		}

		if err := WriteSimpleField[string](ctx, "categoryName", m.GetCategoryName(), WriteString(writeBuffer, int32(int32((int32(m.GetCommandTypeContainer().NumBytes())-int32(int32(1))))*int32(int32(8))))); err != nil {
			return errors.Wrap(err, "Error serializing 'categoryName' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataCategoryName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataCategoryName")
		}
		return nil
	}
	return m.MediaTransportControlDataContract.(*_MediaTransportControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataCategoryName) IsMediaTransportControlDataCategoryName() {}

func (m *_MediaTransportControlDataCategoryName) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MediaTransportControlDataCategoryName) deepCopy() *_MediaTransportControlDataCategoryName {
	if m == nil {
		return nil
	}
	_MediaTransportControlDataCategoryNameCopy := &_MediaTransportControlDataCategoryName{
		m.MediaTransportControlDataContract.(*_MediaTransportControlData).deepCopy(),
		m.CategoryName,
	}
	_MediaTransportControlDataCategoryNameCopy.MediaTransportControlDataContract.(*_MediaTransportControlData)._SubType = m
	return _MediaTransportControlDataCategoryNameCopy
}

func (m *_MediaTransportControlDataCategoryName) String() string {
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