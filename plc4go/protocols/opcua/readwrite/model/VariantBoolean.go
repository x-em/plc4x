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

// VariantBoolean is the corresponding interface of VariantBoolean
type VariantBoolean interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	Variant
	// GetArrayLength returns ArrayLength (property field)
	GetArrayLength() *int32
	// GetValue returns Value (property field)
	GetValue() []byte
	// IsVariantBoolean is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsVariantBoolean()
	// CreateBuilder creates a VariantBooleanBuilder
	CreateVariantBooleanBuilder() VariantBooleanBuilder
}

// _VariantBoolean is the data-structure of this message
type _VariantBoolean struct {
	VariantContract
	ArrayLength *int32
	Value       []byte
}

var _ VariantBoolean = (*_VariantBoolean)(nil)
var _ VariantRequirements = (*_VariantBoolean)(nil)

// NewVariantBoolean factory function for _VariantBoolean
func NewVariantBoolean(arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool, arrayLength *int32, value []byte) *_VariantBoolean {
	_result := &_VariantBoolean{
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

// VariantBooleanBuilder is a builder for VariantBoolean
type VariantBooleanBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value []byte) VariantBooleanBuilder
	// WithArrayLength adds ArrayLength (property field)
	WithOptionalArrayLength(int32) VariantBooleanBuilder
	// WithValue adds Value (property field)
	WithValue(...byte) VariantBooleanBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() VariantBuilder
	// Build builds the VariantBoolean or returns an error if something is wrong
	Build() (VariantBoolean, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() VariantBoolean
}

// NewVariantBooleanBuilder() creates a VariantBooleanBuilder
func NewVariantBooleanBuilder() VariantBooleanBuilder {
	return &_VariantBooleanBuilder{_VariantBoolean: new(_VariantBoolean)}
}

type _VariantBooleanBuilder struct {
	*_VariantBoolean

	parentBuilder *_VariantBuilder

	err *utils.MultiError
}

var _ (VariantBooleanBuilder) = (*_VariantBooleanBuilder)(nil)

func (b *_VariantBooleanBuilder) setParent(contract VariantContract) {
	b.VariantContract = contract
	contract.(*_Variant)._SubType = b._VariantBoolean
}

func (b *_VariantBooleanBuilder) WithMandatoryFields(value []byte) VariantBooleanBuilder {
	return b.WithValue(value...)
}

func (b *_VariantBooleanBuilder) WithOptionalArrayLength(arrayLength int32) VariantBooleanBuilder {
	b.ArrayLength = &arrayLength
	return b
}

func (b *_VariantBooleanBuilder) WithValue(value ...byte) VariantBooleanBuilder {
	b.Value = value
	return b
}

func (b *_VariantBooleanBuilder) Build() (VariantBoolean, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._VariantBoolean.deepCopy(), nil
}

func (b *_VariantBooleanBuilder) MustBuild() VariantBoolean {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_VariantBooleanBuilder) Done() VariantBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewVariantBuilder().(*_VariantBuilder)
	}
	return b.parentBuilder
}

func (b *_VariantBooleanBuilder) buildForVariant() (Variant, error) {
	return b.Build()
}

func (b *_VariantBooleanBuilder) DeepCopy() any {
	_copy := b.CreateVariantBooleanBuilder().(*_VariantBooleanBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateVariantBooleanBuilder creates a VariantBooleanBuilder
func (b *_VariantBoolean) CreateVariantBooleanBuilder() VariantBooleanBuilder {
	if b == nil {
		return NewVariantBooleanBuilder()
	}
	return &_VariantBooleanBuilder{_VariantBoolean: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_VariantBoolean) GetVariantType() uint8 {
	return uint8(1)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_VariantBoolean) GetParent() VariantContract {
	return m.VariantContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VariantBoolean) GetArrayLength() *int32 {
	return m.ArrayLength
}

func (m *_VariantBoolean) GetValue() []byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastVariantBoolean(structType any) VariantBoolean {
	if casted, ok := structType.(VariantBoolean); ok {
		return casted
	}
	if casted, ok := structType.(*VariantBoolean); ok {
		return *casted
	}
	return nil
}

func (m *_VariantBoolean) GetTypeName() string {
	return "VariantBoolean"
}

func (m *_VariantBoolean) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.VariantContract.(*_Variant).getLengthInBits(ctx))

	// Optional Field (arrayLength)
	if m.ArrayLength != nil {
		lengthInBits += 32
	}

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_VariantBoolean) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_VariantBoolean) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Variant, arrayLengthSpecified bool) (__variantBoolean VariantBoolean, err error) {
	m.VariantContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("VariantBoolean"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VariantBoolean")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var arrayLength *int32
	arrayLength, err = ReadOptionalField[int32](ctx, "arrayLength", ReadSignedInt(readBuffer, uint8(32)), arrayLengthSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayLength' field"))
	}
	m.ArrayLength = arrayLength

	value, err := readBuffer.ReadByteArray("value", int(utils.InlineIf(bool((arrayLength) == (nil)), func() any { return int32(int32(1)) }, func() any { return int32((*arrayLength)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("VariantBoolean"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VariantBoolean")
	}

	return m, nil
}

func (m *_VariantBoolean) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VariantBoolean) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VariantBoolean"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for VariantBoolean")
		}

		if err := WriteOptionalField[int32](ctx, "arrayLength", m.GetArrayLength(), WriteSignedInt(writeBuffer, 32), true); err != nil {
			return errors.Wrap(err, "Error serializing 'arrayLength' field")
		}

		if err := WriteByteArrayField(ctx, "value", m.GetValue(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("VariantBoolean"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for VariantBoolean")
		}
		return nil
	}
	return m.VariantContract.(*_Variant).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_VariantBoolean) IsVariantBoolean() {}

func (m *_VariantBoolean) DeepCopy() any {
	return m.deepCopy()
}

func (m *_VariantBoolean) deepCopy() *_VariantBoolean {
	if m == nil {
		return nil
	}
	_VariantBooleanCopy := &_VariantBoolean{
		m.VariantContract.(*_Variant).deepCopy(),
		utils.CopyPtr[int32](m.ArrayLength),
		utils.DeepCopySlice[byte, byte](m.Value),
	}
	_VariantBooleanCopy.VariantContract.(*_Variant)._SubType = m
	return _VariantBooleanCopy
}

func (m *_VariantBoolean) String() string {
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