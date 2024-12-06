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

// Index is the corresponding interface of Index
type Index interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsIndex is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsIndex()
	// CreateBuilder creates a IndexBuilder
	CreateIndexBuilder() IndexBuilder
}

// _Index is the data-structure of this message
type _Index struct {
}

var _ Index = (*_Index)(nil)

// NewIndex factory function for _Index
func NewIndex() *_Index {
	return &_Index{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// IndexBuilder is a builder for Index
type IndexBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() IndexBuilder
	// Build builds the Index or returns an error if something is wrong
	Build() (Index, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() Index
}

// NewIndexBuilder() creates a IndexBuilder
func NewIndexBuilder() IndexBuilder {
	return &_IndexBuilder{_Index: new(_Index)}
}

type _IndexBuilder struct {
	*_Index

	err *utils.MultiError
}

var _ (IndexBuilder) = (*_IndexBuilder)(nil)

func (b *_IndexBuilder) WithMandatoryFields() IndexBuilder {
	return b
}

func (b *_IndexBuilder) Build() (Index, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._Index.deepCopy(), nil
}

func (b *_IndexBuilder) MustBuild() Index {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_IndexBuilder) DeepCopy() any {
	_copy := b.CreateIndexBuilder().(*_IndexBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateIndexBuilder creates a IndexBuilder
func (b *_Index) CreateIndexBuilder() IndexBuilder {
	if b == nil {
		return NewIndexBuilder()
	}
	return &_IndexBuilder{_Index: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastIndex(structType any) Index {
	if casted, ok := structType.(Index); ok {
		return casted
	}
	if casted, ok := structType.(*Index); ok {
		return *casted
	}
	return nil
}

func (m *_Index) GetTypeName() string {
	return "Index"
}

func (m *_Index) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_Index) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func IndexParse(ctx context.Context, theBytes []byte) (Index, error) {
	return IndexParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func IndexParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (Index, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (Index, error) {
		return IndexParseWithBuffer(ctx, readBuffer)
	}
}

func IndexParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (Index, error) {
	v, err := (&_Index{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_Index) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__index Index, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Index"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Index")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("Index"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Index")
	}

	return m, nil
}

func (m *_Index) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Index) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("Index"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Index")
	}

	if popErr := writeBuffer.PopContext("Index"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Index")
	}
	return nil
}

func (m *_Index) IsIndex() {}

func (m *_Index) DeepCopy() any {
	return m.deepCopy()
}

func (m *_Index) deepCopy() *_Index {
	if m == nil {
		return nil
	}
	_IndexCopy := &_Index{}
	return _IndexCopy
}

func (m *_Index) String() string {
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
