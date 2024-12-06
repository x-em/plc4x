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

// VariantInt32 is the corresponding interface of VariantInt32
type VariantInt32 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	Variant
	// GetArrayLength returns ArrayLength (property field)
	GetArrayLength() *int32
	// GetValue returns Value (property field)
	GetValue() []int32
	// IsVariantInt32 is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsVariantInt32()
	// CreateBuilder creates a VariantInt32Builder
	CreateVariantInt32Builder() VariantInt32Builder
}

// _VariantInt32 is the data-structure of this message
type _VariantInt32 struct {
	VariantContract
	ArrayLength *int32
	Value       []int32
}

var _ VariantInt32 = (*_VariantInt32)(nil)
var _ VariantRequirements = (*_VariantInt32)(nil)

// NewVariantInt32 factory function for _VariantInt32
func NewVariantInt32(arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool, arrayLength *int32, value []int32) *_VariantInt32 {
	_result := &_VariantInt32{
		VariantContract: NewVariant(arrayLengthSpecified, arrayDimensionsSpecified, noOfArrayDimensions, arrayDimensions),
		ArrayLength:     arrayLength,
		Value:           value,
	}
	_result.VariantContract.(*_Variant)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// VariantInt32Builder is a builder for VariantInt32
type VariantInt32Builder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value []int32) VariantInt32Builder
	// WithArrayLength adds ArrayLength (property field)
	WithOptionalArrayLength(int32) VariantInt32Builder
	// WithValue adds Value (property field)
	WithValue(...int32) VariantInt32Builder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() VariantBuilder
	// Build builds the VariantInt32 or returns an error if something is wrong
	Build() (VariantInt32, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() VariantInt32
}

// NewVariantInt32Builder() creates a VariantInt32Builder
func NewVariantInt32Builder() VariantInt32Builder {
	return &_VariantInt32Builder{_VariantInt32: new(_VariantInt32)}
}

type _VariantInt32Builder struct {
	*_VariantInt32

	parentBuilder *_VariantBuilder

	err *utils.MultiError
}

var _ (VariantInt32Builder) = (*_VariantInt32Builder)(nil)

func (b *_VariantInt32Builder) setParent(contract VariantContract) {
	b.VariantContract = contract
	contract.(*_Variant)._SubType = b._VariantInt32
}

func (b *_VariantInt32Builder) WithMandatoryFields(value []int32) VariantInt32Builder {
	return b.WithValue(value...)
}

func (b *_VariantInt32Builder) WithOptionalArrayLength(arrayLength int32) VariantInt32Builder {
	b.ArrayLength = &arrayLength
	return b
}

func (b *_VariantInt32Builder) WithValue(value ...int32) VariantInt32Builder {
	b.Value = value
	return b
}

func (b *_VariantInt32Builder) Build() (VariantInt32, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._VariantInt32.deepCopy(), nil
}

func (b *_VariantInt32Builder) MustBuild() VariantInt32 {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_VariantInt32Builder) Done() VariantBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewVariantBuilder().(*_VariantBuilder)
	}
	return b.parentBuilder
}

func (b *_VariantInt32Builder) buildForVariant() (Variant, error) {
	return b.Build()
}

func (b *_VariantInt32Builder) DeepCopy() any {
	_copy := b.CreateVariantInt32Builder().(*_VariantInt32Builder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateVariantInt32Builder creates a VariantInt32Builder
func (b *_VariantInt32) CreateVariantInt32Builder() VariantInt32Builder {
	if b == nil {
		return NewVariantInt32Builder()
	}
	return &_VariantInt32Builder{_VariantInt32: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_VariantInt32) GetVariantType() uint8 {
	return uint8(6)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_VariantInt32) GetParent() VariantContract {
	return m.VariantContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VariantInt32) GetArrayLength() *int32 {
	return m.ArrayLength
}

func (m *_VariantInt32) GetValue() []int32 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastVariantInt32(structType any) VariantInt32 {
	if casted, ok := structType.(VariantInt32); ok {
		return casted
	}
	if casted, ok := structType.(*VariantInt32); ok {
		return *casted
	}
	return nil
}

func (m *_VariantInt32) GetTypeName() string {
	return "VariantInt32"
}

func (m *_VariantInt32) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.VariantContract.(*_Variant).getLengthInBits(ctx))

	// Optional Field (arrayLength)
	if m.ArrayLength != nil {
		lengthInBits += 32
	}

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 32 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_VariantInt32) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_VariantInt32) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Variant, arrayLengthSpecified bool) (__variantInt32 VariantInt32, err error) {
	m.VariantContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("VariantInt32"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VariantInt32")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var arrayLength *int32
	arrayLength, err = ReadOptionalField[int32](ctx, "arrayLength", ReadSignedInt(readBuffer, uint8(32)), arrayLengthSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayLength' field"))
	}
	m.ArrayLength = arrayLength

	value, err := ReadCountArrayField[int32](ctx, "value", ReadSignedInt(readBuffer, uint8(32)), uint64(utils.InlineIf(bool((arrayLength) == (nil)), func() any { return int32(int32(1)) }, func() any { return int32((*arrayLength)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("VariantInt32"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VariantInt32")
	}

	return m, nil
}

func (m *_VariantInt32) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VariantInt32) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VariantInt32"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for VariantInt32")
		}

		if err := WriteOptionalField[int32](ctx, "arrayLength", m.GetArrayLength(), WriteSignedInt(writeBuffer, 32), true); err != nil {
			return errors.Wrap(err, "Error serializing 'arrayLength' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "value", m.GetValue(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("VariantInt32"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for VariantInt32")
		}
		return nil
	}
	return m.VariantContract.(*_Variant).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_VariantInt32) IsVariantInt32() {}

func (m *_VariantInt32) DeepCopy() any {
	return m.deepCopy()
}

func (m *_VariantInt32) deepCopy() *_VariantInt32 {
	if m == nil {
		return nil
	}
	_VariantInt32Copy := &_VariantInt32{
		m.VariantContract.(*_Variant).deepCopy(),
		utils.CopyPtr[int32](m.ArrayLength),
		utils.DeepCopySlice[int32, int32](m.Value),
	}
	_VariantInt32Copy.VariantContract.(*_Variant)._SubType = m
	return _VariantInt32Copy
}

func (m *_VariantInt32) String() string {
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