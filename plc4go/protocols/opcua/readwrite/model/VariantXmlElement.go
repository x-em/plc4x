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

// VariantXmlElement is the corresponding interface of VariantXmlElement
type VariantXmlElement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	Variant
	// GetArrayLength returns ArrayLength (property field)
	GetArrayLength() *int32
	// GetValue returns Value (property field)
	GetValue() []PascalString
	// IsVariantXmlElement is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsVariantXmlElement()
	// CreateBuilder creates a VariantXmlElementBuilder
	CreateVariantXmlElementBuilder() VariantXmlElementBuilder
}

// _VariantXmlElement is the data-structure of this message
type _VariantXmlElement struct {
	VariantContract
	ArrayLength *int32
	Value       []PascalString
}

var _ VariantXmlElement = (*_VariantXmlElement)(nil)
var _ VariantRequirements = (*_VariantXmlElement)(nil)

// NewVariantXmlElement factory function for _VariantXmlElement
func NewVariantXmlElement(arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool, arrayLength *int32, value []PascalString) *_VariantXmlElement {
	_result := &_VariantXmlElement{
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

// VariantXmlElementBuilder is a builder for VariantXmlElement
type VariantXmlElementBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value []PascalString) VariantXmlElementBuilder
	// WithArrayLength adds ArrayLength (property field)
	WithOptionalArrayLength(int32) VariantXmlElementBuilder
	// WithValue adds Value (property field)
	WithValue(...PascalString) VariantXmlElementBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() VariantBuilder
	// Build builds the VariantXmlElement or returns an error if something is wrong
	Build() (VariantXmlElement, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() VariantXmlElement
}

// NewVariantXmlElementBuilder() creates a VariantXmlElementBuilder
func NewVariantXmlElementBuilder() VariantXmlElementBuilder {
	return &_VariantXmlElementBuilder{_VariantXmlElement: new(_VariantXmlElement)}
}

type _VariantXmlElementBuilder struct {
	*_VariantXmlElement

	parentBuilder *_VariantBuilder

	err *utils.MultiError
}

var _ (VariantXmlElementBuilder) = (*_VariantXmlElementBuilder)(nil)

func (b *_VariantXmlElementBuilder) setParent(contract VariantContract) {
	b.VariantContract = contract
	contract.(*_Variant)._SubType = b._VariantXmlElement
}

func (b *_VariantXmlElementBuilder) WithMandatoryFields(value []PascalString) VariantXmlElementBuilder {
	return b.WithValue(value...)
}

func (b *_VariantXmlElementBuilder) WithOptionalArrayLength(arrayLength int32) VariantXmlElementBuilder {
	b.ArrayLength = &arrayLength
	return b
}

func (b *_VariantXmlElementBuilder) WithValue(value ...PascalString) VariantXmlElementBuilder {
	b.Value = value
	return b
}

func (b *_VariantXmlElementBuilder) Build() (VariantXmlElement, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._VariantXmlElement.deepCopy(), nil
}

func (b *_VariantXmlElementBuilder) MustBuild() VariantXmlElement {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_VariantXmlElementBuilder) Done() VariantBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewVariantBuilder().(*_VariantBuilder)
	}
	return b.parentBuilder
}

func (b *_VariantXmlElementBuilder) buildForVariant() (Variant, error) {
	return b.Build()
}

func (b *_VariantXmlElementBuilder) DeepCopy() any {
	_copy := b.CreateVariantXmlElementBuilder().(*_VariantXmlElementBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateVariantXmlElementBuilder creates a VariantXmlElementBuilder
func (b *_VariantXmlElement) CreateVariantXmlElementBuilder() VariantXmlElementBuilder {
	if b == nil {
		return NewVariantXmlElementBuilder()
	}
	return &_VariantXmlElementBuilder{_VariantXmlElement: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_VariantXmlElement) GetVariantType() uint8 {
	return uint8(16)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_VariantXmlElement) GetParent() VariantContract {
	return m.VariantContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VariantXmlElement) GetArrayLength() *int32 {
	return m.ArrayLength
}

func (m *_VariantXmlElement) GetValue() []PascalString {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastVariantXmlElement(structType any) VariantXmlElement {
	if casted, ok := structType.(VariantXmlElement); ok {
		return casted
	}
	if casted, ok := structType.(*VariantXmlElement); ok {
		return *casted
	}
	return nil
}

func (m *_VariantXmlElement) GetTypeName() string {
	return "VariantXmlElement"
}

func (m *_VariantXmlElement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.VariantContract.(*_Variant).getLengthInBits(ctx))

	// Optional Field (arrayLength)
	if m.ArrayLength != nil {
		lengthInBits += 32
	}

	// Array field
	if len(m.Value) > 0 {
		for _curItem, element := range m.Value {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Value), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_VariantXmlElement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_VariantXmlElement) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Variant, arrayLengthSpecified bool) (__variantXmlElement VariantXmlElement, err error) {
	m.VariantContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("VariantXmlElement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VariantXmlElement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var arrayLength *int32
	arrayLength, err = ReadOptionalField[int32](ctx, "arrayLength", ReadSignedInt(readBuffer, uint8(32)), arrayLengthSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayLength' field"))
	}
	m.ArrayLength = arrayLength

	value, err := ReadCountArrayField[PascalString](ctx, "value", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(utils.InlineIf(bool((arrayLength) == (nil)), func() any { return int32(int32(1)) }, func() any { return int32((*arrayLength)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("VariantXmlElement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VariantXmlElement")
	}

	return m, nil
}

func (m *_VariantXmlElement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VariantXmlElement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VariantXmlElement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for VariantXmlElement")
		}

		if err := WriteOptionalField[int32](ctx, "arrayLength", m.GetArrayLength(), WriteSignedInt(writeBuffer, 32), true); err != nil {
			return errors.Wrap(err, "Error serializing 'arrayLength' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "value", m.GetValue(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("VariantXmlElement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for VariantXmlElement")
		}
		return nil
	}
	return m.VariantContract.(*_Variant).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_VariantXmlElement) IsVariantXmlElement() {}

func (m *_VariantXmlElement) DeepCopy() any {
	return m.deepCopy()
}

func (m *_VariantXmlElement) deepCopy() *_VariantXmlElement {
	if m == nil {
		return nil
	}
	_VariantXmlElementCopy := &_VariantXmlElement{
		m.VariantContract.(*_Variant).deepCopy(),
		utils.CopyPtr[int32](m.ArrayLength),
		utils.DeepCopySlice[PascalString, PascalString](m.Value),
	}
	_VariantXmlElementCopy.VariantContract.(*_Variant)._SubType = m
	return _VariantXmlElementCopy
}

func (m *_VariantXmlElement) String() string {
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
