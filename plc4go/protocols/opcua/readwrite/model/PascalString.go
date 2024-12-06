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

// PascalString is the corresponding interface of PascalString
type PascalString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetStringValue returns StringValue (property field)
	GetStringValue() *string
	// GetStringLength returns StringLength (virtual field)
	GetStringLength() int32
	// IsPascalString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPascalString()
	// CreateBuilder creates a PascalStringBuilder
	CreatePascalStringBuilder() PascalStringBuilder
}

// _PascalString is the data-structure of this message
type _PascalString struct {
	StringValue *string
}

var _ PascalString = (*_PascalString)(nil)

// NewPascalString factory function for _PascalString
func NewPascalString(stringValue *string) *_PascalString {
	return &_PascalString{StringValue: stringValue}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// PascalStringBuilder is a builder for PascalString
type PascalStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() PascalStringBuilder
	// WithStringValue adds StringValue (property field)
	WithOptionalStringValue(string) PascalStringBuilder
	// Build builds the PascalString or returns an error if something is wrong
	Build() (PascalString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() PascalString
}

// NewPascalStringBuilder() creates a PascalStringBuilder
func NewPascalStringBuilder() PascalStringBuilder {
	return &_PascalStringBuilder{_PascalString: new(_PascalString)}
}

type _PascalStringBuilder struct {
	*_PascalString

	err *utils.MultiError
}

var _ (PascalStringBuilder) = (*_PascalStringBuilder)(nil)

func (b *_PascalStringBuilder) WithMandatoryFields() PascalStringBuilder {
	return b
}

func (b *_PascalStringBuilder) WithOptionalStringValue(stringValue string) PascalStringBuilder {
	b.StringValue = &stringValue
	return b
}

func (b *_PascalStringBuilder) Build() (PascalString, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._PascalString.deepCopy(), nil
}

func (b *_PascalStringBuilder) MustBuild() PascalString {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_PascalStringBuilder) DeepCopy() any {
	_copy := b.CreatePascalStringBuilder().(*_PascalStringBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreatePascalStringBuilder creates a PascalStringBuilder
func (b *_PascalString) CreatePascalStringBuilder() PascalStringBuilder {
	if b == nil {
		return NewPascalStringBuilder()
	}
	return &_PascalStringBuilder{_PascalString: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PascalString) GetStringValue() *string {
	return m.StringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_PascalString) GetStringLength() int32 {
	ctx := context.Background()
	_ = ctx
	stringValue := m.GetStringValue()
	_ = stringValue
	return int32(PascalLengthToUtf8Length(ctx, Utf8LengthToPascalLength(ctx, (*m.GetStringValue()))))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastPascalString(structType any) PascalString {
	if casted, ok := structType.(PascalString); ok {
		return casted
	}
	if casted, ok := structType.(*PascalString); ok {
		return *casted
	}
	return nil
}

func (m *_PascalString) GetTypeName() string {
	return "PascalString"
}

func (m *_PascalString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (sLength)
	lengthInBits += 32

	// A virtual field doesn't have any in- or output.

	// Optional Field (stringValue)
	if m.StringValue != nil {
		lengthInBits += 0
	}

	return lengthInBits
}

func (m *_PascalString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PascalStringParse(ctx context.Context, theBytes []byte) (PascalString, error) {
	return PascalStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func PascalStringParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (PascalString, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (PascalString, error) {
		return PascalStringParseWithBuffer(ctx, readBuffer)
	}
}

func PascalStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (PascalString, error) {
	v, err := (&_PascalString{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_PascalString) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__pascalString PascalString, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PascalString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PascalString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	sLength, err := ReadImplicitField[int32](ctx, "sLength", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sLength' field"))
	}
	_ = sLength

	stringLength, err := ReadVirtualField[int32](ctx, "stringLength", (*int32)(nil), PascalLengthToUtf8Length(ctx, sLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'stringLength' field"))
	}
	_ = stringLength

	var stringValue *string
	stringValue, err = ReadOptionalField[string](ctx, "stringValue", ReadString(readBuffer, uint32(int32(stringLength)*int32(int32(8)))), bool((sLength) != (-(1))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'stringValue' field"))
	}
	m.StringValue = stringValue

	if closeErr := readBuffer.CloseContext("PascalString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PascalString")
	}

	return m, nil
}

func (m *_PascalString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PascalString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("PascalString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for PascalString")
	}
	sLength := int32(Utf8LengthToPascalLength(ctx, (*m.GetStringValue())))
	if err := WriteImplicitField(ctx, "sLength", sLength, WriteSignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'sLength' field")
	}
	// Virtual field
	stringLength := m.GetStringLength()
	_ = stringLength
	if _stringLengthErr := writeBuffer.WriteVirtual(ctx, "stringLength", m.GetStringLength()); _stringLengthErr != nil {
		return errors.Wrap(_stringLengthErr, "Error serializing 'stringLength' field")
	}

	if err := WriteOptionalField[string](ctx, "stringValue", m.GetStringValue(), WriteString(writeBuffer, int32(int32(m.GetStringLength())*int32(int32(8)))), true); err != nil {
		return errors.Wrap(err, "Error serializing 'stringValue' field")
	}

	if popErr := writeBuffer.PopContext("PascalString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for PascalString")
	}
	return nil
}

func (m *_PascalString) IsPascalString() {}

func (m *_PascalString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_PascalString) deepCopy() *_PascalString {
	if m == nil {
		return nil
	}
	_PascalStringCopy := &_PascalString{
		utils.CopyPtr[string](m.StringValue),
	}
	return _PascalStringCopy
}

func (m *_PascalString) String() string {
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