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

// CustomManufacturer is the corresponding interface of CustomManufacturer
type CustomManufacturer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetCustomString returns CustomString (property field)
	GetCustomString() string
	// IsCustomManufacturer is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCustomManufacturer()
	// CreateBuilder creates a CustomManufacturerBuilder
	CreateCustomManufacturerBuilder() CustomManufacturerBuilder
}

// _CustomManufacturer is the data-structure of this message
type _CustomManufacturer struct {
	CustomString string

	// Arguments.
	NumBytes uint8
}

var _ CustomManufacturer = (*_CustomManufacturer)(nil)

// NewCustomManufacturer factory function for _CustomManufacturer
func NewCustomManufacturer(customString string, numBytes uint8) *_CustomManufacturer {
	return &_CustomManufacturer{CustomString: customString, NumBytes: numBytes}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CustomManufacturerBuilder is a builder for CustomManufacturer
type CustomManufacturerBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(customString string) CustomManufacturerBuilder
	// WithCustomString adds CustomString (property field)
	WithCustomString(string) CustomManufacturerBuilder
	// WithArgNumBytes sets a parser argument
	WithArgNumBytes(uint8) CustomManufacturerBuilder
	// Build builds the CustomManufacturer or returns an error if something is wrong
	Build() (CustomManufacturer, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CustomManufacturer
}

// NewCustomManufacturerBuilder() creates a CustomManufacturerBuilder
func NewCustomManufacturerBuilder() CustomManufacturerBuilder {
	return &_CustomManufacturerBuilder{_CustomManufacturer: new(_CustomManufacturer)}
}

type _CustomManufacturerBuilder struct {
	*_CustomManufacturer

	err *utils.MultiError
}

var _ (CustomManufacturerBuilder) = (*_CustomManufacturerBuilder)(nil)

func (b *_CustomManufacturerBuilder) WithMandatoryFields(customString string) CustomManufacturerBuilder {
	return b.WithCustomString(customString)
}

func (b *_CustomManufacturerBuilder) WithCustomString(customString string) CustomManufacturerBuilder {
	b.CustomString = customString
	return b
}

func (b *_CustomManufacturerBuilder) WithArgNumBytes(numBytes uint8) CustomManufacturerBuilder {
	b.NumBytes = numBytes
	return b
}

func (b *_CustomManufacturerBuilder) Build() (CustomManufacturer, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CustomManufacturer.deepCopy(), nil
}

func (b *_CustomManufacturerBuilder) MustBuild() CustomManufacturer {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CustomManufacturerBuilder) DeepCopy() any {
	_copy := b.CreateCustomManufacturerBuilder().(*_CustomManufacturerBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCustomManufacturerBuilder creates a CustomManufacturerBuilder
func (b *_CustomManufacturer) CreateCustomManufacturerBuilder() CustomManufacturerBuilder {
	if b == nil {
		return NewCustomManufacturerBuilder()
	}
	return &_CustomManufacturerBuilder{_CustomManufacturer: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CustomManufacturer) GetCustomString() string {
	return m.CustomString
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCustomManufacturer(structType any) CustomManufacturer {
	if casted, ok := structType.(CustomManufacturer); ok {
		return casted
	}
	if casted, ok := structType.(*CustomManufacturer); ok {
		return *casted
	}
	return nil
}

func (m *_CustomManufacturer) GetTypeName() string {
	return "CustomManufacturer"
}

func (m *_CustomManufacturer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (customString)
	lengthInBits += uint16(int32(int32(8)) * int32(m.GetNumBytes()))

	return lengthInBits
}

func (m *_CustomManufacturer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CustomManufacturerParse(ctx context.Context, theBytes []byte, numBytes uint8) (CustomManufacturer, error) {
	return CustomManufacturerParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), numBytes)
}

func CustomManufacturerParseWithBufferProducer(numBytes uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (CustomManufacturer, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (CustomManufacturer, error) {
		return CustomManufacturerParseWithBuffer(ctx, readBuffer, numBytes)
	}
}

func CustomManufacturerParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, numBytes uint8) (CustomManufacturer, error) {
	v, err := (&_CustomManufacturer{NumBytes: numBytes}).parse(ctx, readBuffer, numBytes)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_CustomManufacturer) parse(ctx context.Context, readBuffer utils.ReadBuffer, numBytes uint8) (__customManufacturer CustomManufacturer, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CustomManufacturer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CustomManufacturer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	customString, err := ReadSimpleField(ctx, "customString", ReadString(readBuffer, uint32(int32(int32(8))*int32(numBytes))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'customString' field"))
	}
	m.CustomString = customString

	if closeErr := readBuffer.CloseContext("CustomManufacturer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CustomManufacturer")
	}

	return m, nil
}

func (m *_CustomManufacturer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CustomManufacturer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CustomManufacturer"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CustomManufacturer")
	}

	if err := WriteSimpleField[string](ctx, "customString", m.GetCustomString(), WriteString(writeBuffer, int32(int32(int32(8))*int32(m.GetNumBytes())))); err != nil {
		return errors.Wrap(err, "Error serializing 'customString' field")
	}

	if popErr := writeBuffer.PopContext("CustomManufacturer"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CustomManufacturer")
	}
	return nil
}

////
// Arguments Getter

func (m *_CustomManufacturer) GetNumBytes() uint8 {
	return m.NumBytes
}

//
////

func (m *_CustomManufacturer) IsCustomManufacturer() {}

func (m *_CustomManufacturer) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CustomManufacturer) deepCopy() *_CustomManufacturer {
	if m == nil {
		return nil
	}
	_CustomManufacturerCopy := &_CustomManufacturer{
		m.CustomString,
		m.NumBytes,
	}
	return _CustomManufacturerCopy
}

func (m *_CustomManufacturer) String() string {
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