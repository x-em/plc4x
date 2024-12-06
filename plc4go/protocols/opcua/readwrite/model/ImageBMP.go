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

// ImageBMP is the corresponding interface of ImageBMP
type ImageBMP interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsImageBMP is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsImageBMP()
	// CreateBuilder creates a ImageBMPBuilder
	CreateImageBMPBuilder() ImageBMPBuilder
}

// _ImageBMP is the data-structure of this message
type _ImageBMP struct {
}

var _ ImageBMP = (*_ImageBMP)(nil)

// NewImageBMP factory function for _ImageBMP
func NewImageBMP() *_ImageBMP {
	return &_ImageBMP{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ImageBMPBuilder is a builder for ImageBMP
type ImageBMPBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ImageBMPBuilder
	// Build builds the ImageBMP or returns an error if something is wrong
	Build() (ImageBMP, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ImageBMP
}

// NewImageBMPBuilder() creates a ImageBMPBuilder
func NewImageBMPBuilder() ImageBMPBuilder {
	return &_ImageBMPBuilder{_ImageBMP: new(_ImageBMP)}
}

type _ImageBMPBuilder struct {
	*_ImageBMP

	err *utils.MultiError
}

var _ (ImageBMPBuilder) = (*_ImageBMPBuilder)(nil)

func (b *_ImageBMPBuilder) WithMandatoryFields() ImageBMPBuilder {
	return b
}

func (b *_ImageBMPBuilder) Build() (ImageBMP, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ImageBMP.deepCopy(), nil
}

func (b *_ImageBMPBuilder) MustBuild() ImageBMP {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ImageBMPBuilder) DeepCopy() any {
	_copy := b.CreateImageBMPBuilder().(*_ImageBMPBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateImageBMPBuilder creates a ImageBMPBuilder
func (b *_ImageBMP) CreateImageBMPBuilder() ImageBMPBuilder {
	if b == nil {
		return NewImageBMPBuilder()
	}
	return &_ImageBMPBuilder{_ImageBMP: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastImageBMP(structType any) ImageBMP {
	if casted, ok := structType.(ImageBMP); ok {
		return casted
	}
	if casted, ok := structType.(*ImageBMP); ok {
		return *casted
	}
	return nil
}

func (m *_ImageBMP) GetTypeName() string {
	return "ImageBMP"
}

func (m *_ImageBMP) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_ImageBMP) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ImageBMPParse(ctx context.Context, theBytes []byte) (ImageBMP, error) {
	return ImageBMPParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ImageBMPParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ImageBMP, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ImageBMP, error) {
		return ImageBMPParseWithBuffer(ctx, readBuffer)
	}
}

func ImageBMPParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ImageBMP, error) {
	v, err := (&_ImageBMP{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_ImageBMP) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__imageBMP ImageBMP, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ImageBMP"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ImageBMP")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ImageBMP"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ImageBMP")
	}

	return m, nil
}

func (m *_ImageBMP) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ImageBMP) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ImageBMP"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ImageBMP")
	}

	if popErr := writeBuffer.PopContext("ImageBMP"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ImageBMP")
	}
	return nil
}

func (m *_ImageBMP) IsImageBMP() {}

func (m *_ImageBMP) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ImageBMP) deepCopy() *_ImageBMP {
	if m == nil {
		return nil
	}
	_ImageBMPCopy := &_ImageBMP{}
	return _ImageBMPCopy
}

func (m *_ImageBMP) String() string {
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