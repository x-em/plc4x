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

// DurationString is the corresponding interface of DurationString
type DurationString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsDurationString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDurationString()
	// CreateBuilder creates a DurationStringBuilder
	CreateDurationStringBuilder() DurationStringBuilder
}

// _DurationString is the data-structure of this message
type _DurationString struct {
}

var _ DurationString = (*_DurationString)(nil)

// NewDurationString factory function for _DurationString
func NewDurationString() *_DurationString {
	return &_DurationString{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DurationStringBuilder is a builder for DurationString
type DurationStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() DurationStringBuilder
	// Build builds the DurationString or returns an error if something is wrong
	Build() (DurationString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DurationString
}

// NewDurationStringBuilder() creates a DurationStringBuilder
func NewDurationStringBuilder() DurationStringBuilder {
	return &_DurationStringBuilder{_DurationString: new(_DurationString)}
}

type _DurationStringBuilder struct {
	*_DurationString

	err *utils.MultiError
}

var _ (DurationStringBuilder) = (*_DurationStringBuilder)(nil)

func (b *_DurationStringBuilder) WithMandatoryFields() DurationStringBuilder {
	return b
}

func (b *_DurationStringBuilder) Build() (DurationString, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DurationString.deepCopy(), nil
}

func (b *_DurationStringBuilder) MustBuild() DurationString {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DurationStringBuilder) DeepCopy() any {
	_copy := b.CreateDurationStringBuilder().(*_DurationStringBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDurationStringBuilder creates a DurationStringBuilder
func (b *_DurationString) CreateDurationStringBuilder() DurationStringBuilder {
	if b == nil {
		return NewDurationStringBuilder()
	}
	return &_DurationStringBuilder{_DurationString: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDurationString(structType any) DurationString {
	if casted, ok := structType.(DurationString); ok {
		return casted
	}
	if casted, ok := structType.(*DurationString); ok {
		return *casted
	}
	return nil
}

func (m *_DurationString) GetTypeName() string {
	return "DurationString"
}

func (m *_DurationString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_DurationString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DurationStringParse(ctx context.Context, theBytes []byte) (DurationString, error) {
	return DurationStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DurationStringParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (DurationString, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (DurationString, error) {
		return DurationStringParseWithBuffer(ctx, readBuffer)
	}
}

func DurationStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DurationString, error) {
	v, err := (&_DurationString{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_DurationString) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__durationString DurationString, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DurationString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DurationString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("DurationString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DurationString")
	}

	return m, nil
}

func (m *_DurationString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DurationString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("DurationString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DurationString")
	}

	if popErr := writeBuffer.PopContext("DurationString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DurationString")
	}
	return nil
}

func (m *_DurationString) IsDurationString() {}

func (m *_DurationString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DurationString) deepCopy() *_DurationString {
	if m == nil {
		return nil
	}
	_DurationStringCopy := &_DurationString{}
	return _DurationStringCopy
}

func (m *_DurationString) String() string {
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