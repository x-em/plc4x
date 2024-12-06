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

// DecimalString is the corresponding interface of DecimalString
type DecimalString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsDecimalString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDecimalString()
	// CreateBuilder creates a DecimalStringBuilder
	CreateDecimalStringBuilder() DecimalStringBuilder
}

// _DecimalString is the data-structure of this message
type _DecimalString struct {
}

var _ DecimalString = (*_DecimalString)(nil)

// NewDecimalString factory function for _DecimalString
func NewDecimalString() *_DecimalString {
	return &_DecimalString{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DecimalStringBuilder is a builder for DecimalString
type DecimalStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() DecimalStringBuilder
	// Build builds the DecimalString or returns an error if something is wrong
	Build() (DecimalString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DecimalString
}

// NewDecimalStringBuilder() creates a DecimalStringBuilder
func NewDecimalStringBuilder() DecimalStringBuilder {
	return &_DecimalStringBuilder{_DecimalString: new(_DecimalString)}
}

type _DecimalStringBuilder struct {
	*_DecimalString

	err *utils.MultiError
}

var _ (DecimalStringBuilder) = (*_DecimalStringBuilder)(nil)

func (b *_DecimalStringBuilder) WithMandatoryFields() DecimalStringBuilder {
	return b
}

func (b *_DecimalStringBuilder) Build() (DecimalString, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DecimalString.deepCopy(), nil
}

func (b *_DecimalStringBuilder) MustBuild() DecimalString {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DecimalStringBuilder) DeepCopy() any {
	_copy := b.CreateDecimalStringBuilder().(*_DecimalStringBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDecimalStringBuilder creates a DecimalStringBuilder
func (b *_DecimalString) CreateDecimalStringBuilder() DecimalStringBuilder {
	if b == nil {
		return NewDecimalStringBuilder()
	}
	return &_DecimalStringBuilder{_DecimalString: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDecimalString(structType any) DecimalString {
	if casted, ok := structType.(DecimalString); ok {
		return casted
	}
	if casted, ok := structType.(*DecimalString); ok {
		return *casted
	}
	return nil
}

func (m *_DecimalString) GetTypeName() string {
	return "DecimalString"
}

func (m *_DecimalString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_DecimalString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DecimalStringParse(ctx context.Context, theBytes []byte) (DecimalString, error) {
	return DecimalStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DecimalStringParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (DecimalString, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (DecimalString, error) {
		return DecimalStringParseWithBuffer(ctx, readBuffer)
	}
}

func DecimalStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DecimalString, error) {
	v, err := (&_DecimalString{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_DecimalString) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__decimalString DecimalString, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DecimalString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DecimalString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("DecimalString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DecimalString")
	}

	return m, nil
}

func (m *_DecimalString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DecimalString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("DecimalString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DecimalString")
	}

	if popErr := writeBuffer.PopContext("DecimalString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DecimalString")
	}
	return nil
}

func (m *_DecimalString) IsDecimalString() {}

func (m *_DecimalString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DecimalString) deepCopy() *_DecimalString {
	if m == nil {
		return nil
	}
	_DecimalStringCopy := &_DecimalString{}
	return _DecimalStringCopy
}

func (m *_DecimalString) String() string {
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