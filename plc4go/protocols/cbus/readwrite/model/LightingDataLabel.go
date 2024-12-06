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

// LightingDataLabel is the corresponding interface of LightingDataLabel
type LightingDataLabel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	LightingData
	// GetGroup returns Group (property field)
	GetGroup() byte
	// GetLabelOptions returns LabelOptions (property field)
	GetLabelOptions() LightingLabelOptions
	// GetLanguage returns Language (property field)
	GetLanguage() *Language
	// GetData returns Data (property field)
	GetData() []byte
	// IsLightingDataLabel is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLightingDataLabel()
	// CreateBuilder creates a LightingDataLabelBuilder
	CreateLightingDataLabelBuilder() LightingDataLabelBuilder
}

// _LightingDataLabel is the data-structure of this message
type _LightingDataLabel struct {
	LightingDataContract
	Group        byte
	LabelOptions LightingLabelOptions
	Language     *Language
	Data         []byte
}

var _ LightingDataLabel = (*_LightingDataLabel)(nil)
var _ LightingDataRequirements = (*_LightingDataLabel)(nil)

// NewLightingDataLabel factory function for _LightingDataLabel
func NewLightingDataLabel(commandTypeContainer LightingCommandTypeContainer, group byte, labelOptions LightingLabelOptions, language *Language, data []byte) *_LightingDataLabel {
	if labelOptions == nil {
		panic("labelOptions of type LightingLabelOptions for LightingDataLabel must not be nil")
	}
	_result := &_LightingDataLabel{
		LightingDataContract: NewLightingData(commandTypeContainer),
		Group:                group,
		LabelOptions:         labelOptions,
		Language:             language,
		Data:                 data,
	}
	_result.LightingDataContract.(*_LightingData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// LightingDataLabelBuilder is a builder for LightingDataLabel
type LightingDataLabelBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(group byte, labelOptions LightingLabelOptions, data []byte) LightingDataLabelBuilder
	// WithGroup adds Group (property field)
	WithGroup(byte) LightingDataLabelBuilder
	// WithLabelOptions adds LabelOptions (property field)
	WithLabelOptions(LightingLabelOptions) LightingDataLabelBuilder
	// WithLabelOptionsBuilder adds LabelOptions (property field) which is build by the builder
	WithLabelOptionsBuilder(func(LightingLabelOptionsBuilder) LightingLabelOptionsBuilder) LightingDataLabelBuilder
	// WithLanguage adds Language (property field)
	WithOptionalLanguage(Language) LightingDataLabelBuilder
	// WithData adds Data (property field)
	WithData(...byte) LightingDataLabelBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() LightingDataBuilder
	// Build builds the LightingDataLabel or returns an error if something is wrong
	Build() (LightingDataLabel, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() LightingDataLabel
}

// NewLightingDataLabelBuilder() creates a LightingDataLabelBuilder
func NewLightingDataLabelBuilder() LightingDataLabelBuilder {
	return &_LightingDataLabelBuilder{_LightingDataLabel: new(_LightingDataLabel)}
}

type _LightingDataLabelBuilder struct {
	*_LightingDataLabel

	parentBuilder *_LightingDataBuilder

	err *utils.MultiError
}

var _ (LightingDataLabelBuilder) = (*_LightingDataLabelBuilder)(nil)

func (b *_LightingDataLabelBuilder) setParent(contract LightingDataContract) {
	b.LightingDataContract = contract
	contract.(*_LightingData)._SubType = b._LightingDataLabel
}

func (b *_LightingDataLabelBuilder) WithMandatoryFields(group byte, labelOptions LightingLabelOptions, data []byte) LightingDataLabelBuilder {
	return b.WithGroup(group).WithLabelOptions(labelOptions).WithData(data...)
}

func (b *_LightingDataLabelBuilder) WithGroup(group byte) LightingDataLabelBuilder {
	b.Group = group
	return b
}

func (b *_LightingDataLabelBuilder) WithLabelOptions(labelOptions LightingLabelOptions) LightingDataLabelBuilder {
	b.LabelOptions = labelOptions
	return b
}

func (b *_LightingDataLabelBuilder) WithLabelOptionsBuilder(builderSupplier func(LightingLabelOptionsBuilder) LightingLabelOptionsBuilder) LightingDataLabelBuilder {
	builder := builderSupplier(b.LabelOptions.CreateLightingLabelOptionsBuilder())
	var err error
	b.LabelOptions, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LightingLabelOptionsBuilder failed"))
	}
	return b
}

func (b *_LightingDataLabelBuilder) WithOptionalLanguage(language Language) LightingDataLabelBuilder {
	b.Language = &language
	return b
}

func (b *_LightingDataLabelBuilder) WithData(data ...byte) LightingDataLabelBuilder {
	b.Data = data
	return b
}

func (b *_LightingDataLabelBuilder) Build() (LightingDataLabel, error) {
	if b.LabelOptions == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'labelOptions' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._LightingDataLabel.deepCopy(), nil
}

func (b *_LightingDataLabelBuilder) MustBuild() LightingDataLabel {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_LightingDataLabelBuilder) Done() LightingDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewLightingDataBuilder().(*_LightingDataBuilder)
	}
	return b.parentBuilder
}

func (b *_LightingDataLabelBuilder) buildForLightingData() (LightingData, error) {
	return b.Build()
}

func (b *_LightingDataLabelBuilder) DeepCopy() any {
	_copy := b.CreateLightingDataLabelBuilder().(*_LightingDataLabelBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateLightingDataLabelBuilder creates a LightingDataLabelBuilder
func (b *_LightingDataLabel) CreateLightingDataLabelBuilder() LightingDataLabelBuilder {
	if b == nil {
		return NewLightingDataLabelBuilder()
	}
	return &_LightingDataLabelBuilder{_LightingDataLabel: b.deepCopy()}
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

func (m *_LightingDataLabel) GetParent() LightingDataContract {
	return m.LightingDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LightingDataLabel) GetGroup() byte {
	return m.Group
}

func (m *_LightingDataLabel) GetLabelOptions() LightingLabelOptions {
	return m.LabelOptions
}

func (m *_LightingDataLabel) GetLanguage() *Language {
	return m.Language
}

func (m *_LightingDataLabel) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastLightingDataLabel(structType any) LightingDataLabel {
	if casted, ok := structType.(LightingDataLabel); ok {
		return casted
	}
	if casted, ok := structType.(*LightingDataLabel); ok {
		return *casted
	}
	return nil
}

func (m *_LightingDataLabel) GetTypeName() string {
	return "LightingDataLabel"
}

func (m *_LightingDataLabel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.LightingDataContract.(*_LightingData).getLengthInBits(ctx))

	// Simple field (group)
	lengthInBits += 8

	// Simple field (labelOptions)
	lengthInBits += m.LabelOptions.GetLengthInBits(ctx)

	// Optional Field (language)
	if m.Language != nil {
		lengthInBits += 8
	}

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_LightingDataLabel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LightingDataLabel) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_LightingData, commandTypeContainer LightingCommandTypeContainer) (__lightingDataLabel LightingDataLabel, err error) {
	m.LightingDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LightingDataLabel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LightingDataLabel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	group, err := ReadSimpleField(ctx, "group", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'group' field"))
	}
	m.Group = group

	labelOptions, err := ReadSimpleField[LightingLabelOptions](ctx, "labelOptions", ReadComplex[LightingLabelOptions](LightingLabelOptionsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'labelOptions' field"))
	}
	m.LabelOptions = labelOptions

	var language *Language
	language, err = ReadOptionalField[Language](ctx, "language", ReadEnum(LanguageByValue, ReadUnsignedByte(readBuffer, uint8(8))), bool((labelOptions.GetLabelType()) != (LightingLabelType_LOAD_DYNAMIC_ICON)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'language' field"))
	}
	m.Language = language

	data, err := readBuffer.ReadByteArray("data", int((int32(commandTypeContainer.NumBytes()) - int32((utils.InlineIf((bool((labelOptions.GetLabelType()) != (LightingLabelType_LOAD_DYNAMIC_ICON))), func() any { return int32((int32(3))) }, func() any { return int32((int32(2))) }).(int32))))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("LightingDataLabel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LightingDataLabel")
	}

	return m, nil
}

func (m *_LightingDataLabel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LightingDataLabel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LightingDataLabel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LightingDataLabel")
		}

		if err := WriteSimpleField[byte](ctx, "group", m.GetGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'group' field")
		}

		if err := WriteSimpleField[LightingLabelOptions](ctx, "labelOptions", m.GetLabelOptions(), WriteComplex[LightingLabelOptions](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'labelOptions' field")
		}

		if err := WriteOptionalEnumField[Language](ctx, "language", "Language", m.GetLanguage(), WriteEnum[Language, uint8](Language.GetValue, Language.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8)), bool((m.GetLabelOptions().GetLabelType()) != (LightingLabelType_LOAD_DYNAMIC_ICON))); err != nil {
			return errors.Wrap(err, "Error serializing 'language' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("LightingDataLabel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LightingDataLabel")
		}
		return nil
	}
	return m.LightingDataContract.(*_LightingData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LightingDataLabel) IsLightingDataLabel() {}

func (m *_LightingDataLabel) DeepCopy() any {
	return m.deepCopy()
}

func (m *_LightingDataLabel) deepCopy() *_LightingDataLabel {
	if m == nil {
		return nil
	}
	_LightingDataLabelCopy := &_LightingDataLabel{
		m.LightingDataContract.(*_LightingData).deepCopy(),
		m.Group,
		utils.DeepCopy[LightingLabelOptions](m.LabelOptions),
		utils.CopyPtr[Language](m.Language),
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	_LightingDataLabelCopy.LightingDataContract.(*_LightingData)._SubType = m
	return _LightingDataLabelCopy
}

func (m *_LightingDataLabel) String() string {
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
