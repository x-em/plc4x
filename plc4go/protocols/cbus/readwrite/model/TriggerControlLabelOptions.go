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

// TriggerControlLabelOptions is the corresponding interface of TriggerControlLabelOptions
type TriggerControlLabelOptions interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetLabelFlavour returns LabelFlavour (property field)
	GetLabelFlavour() TriggerControlLabelFlavour
	// GetLabelType returns LabelType (property field)
	GetLabelType() TriggerControlLabelType
	// IsTriggerControlLabelOptions is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTriggerControlLabelOptions()
	// CreateBuilder creates a TriggerControlLabelOptionsBuilder
	CreateTriggerControlLabelOptionsBuilder() TriggerControlLabelOptionsBuilder
}

// _TriggerControlLabelOptions is the data-structure of this message
type _TriggerControlLabelOptions struct {
	LabelFlavour TriggerControlLabelFlavour
	LabelType    TriggerControlLabelType
	// Reserved Fields
	reservedField0 *bool
	reservedField1 *bool
	reservedField2 *bool
	reservedField3 *bool
}

var _ TriggerControlLabelOptions = (*_TriggerControlLabelOptions)(nil)

// NewTriggerControlLabelOptions factory function for _TriggerControlLabelOptions
func NewTriggerControlLabelOptions(labelFlavour TriggerControlLabelFlavour, labelType TriggerControlLabelType) *_TriggerControlLabelOptions {
	return &_TriggerControlLabelOptions{LabelFlavour: labelFlavour, LabelType: labelType}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// TriggerControlLabelOptionsBuilder is a builder for TriggerControlLabelOptions
type TriggerControlLabelOptionsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(labelFlavour TriggerControlLabelFlavour, labelType TriggerControlLabelType) TriggerControlLabelOptionsBuilder
	// WithLabelFlavour adds LabelFlavour (property field)
	WithLabelFlavour(TriggerControlLabelFlavour) TriggerControlLabelOptionsBuilder
	// WithLabelType adds LabelType (property field)
	WithLabelType(TriggerControlLabelType) TriggerControlLabelOptionsBuilder
	// Build builds the TriggerControlLabelOptions or returns an error if something is wrong
	Build() (TriggerControlLabelOptions, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() TriggerControlLabelOptions
}

// NewTriggerControlLabelOptionsBuilder() creates a TriggerControlLabelOptionsBuilder
func NewTriggerControlLabelOptionsBuilder() TriggerControlLabelOptionsBuilder {
	return &_TriggerControlLabelOptionsBuilder{_TriggerControlLabelOptions: new(_TriggerControlLabelOptions)}
}

type _TriggerControlLabelOptionsBuilder struct {
	*_TriggerControlLabelOptions

	err *utils.MultiError
}

var _ (TriggerControlLabelOptionsBuilder) = (*_TriggerControlLabelOptionsBuilder)(nil)

func (b *_TriggerControlLabelOptionsBuilder) WithMandatoryFields(labelFlavour TriggerControlLabelFlavour, labelType TriggerControlLabelType) TriggerControlLabelOptionsBuilder {
	return b.WithLabelFlavour(labelFlavour).WithLabelType(labelType)
}

func (b *_TriggerControlLabelOptionsBuilder) WithLabelFlavour(labelFlavour TriggerControlLabelFlavour) TriggerControlLabelOptionsBuilder {
	b.LabelFlavour = labelFlavour
	return b
}

func (b *_TriggerControlLabelOptionsBuilder) WithLabelType(labelType TriggerControlLabelType) TriggerControlLabelOptionsBuilder {
	b.LabelType = labelType
	return b
}

func (b *_TriggerControlLabelOptionsBuilder) Build() (TriggerControlLabelOptions, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._TriggerControlLabelOptions.deepCopy(), nil
}

func (b *_TriggerControlLabelOptionsBuilder) MustBuild() TriggerControlLabelOptions {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_TriggerControlLabelOptionsBuilder) DeepCopy() any {
	_copy := b.CreateTriggerControlLabelOptionsBuilder().(*_TriggerControlLabelOptionsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateTriggerControlLabelOptionsBuilder creates a TriggerControlLabelOptionsBuilder
func (b *_TriggerControlLabelOptions) CreateTriggerControlLabelOptionsBuilder() TriggerControlLabelOptionsBuilder {
	if b == nil {
		return NewTriggerControlLabelOptionsBuilder()
	}
	return &_TriggerControlLabelOptionsBuilder{_TriggerControlLabelOptions: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TriggerControlLabelOptions) GetLabelFlavour() TriggerControlLabelFlavour {
	return m.LabelFlavour
}

func (m *_TriggerControlLabelOptions) GetLabelType() TriggerControlLabelType {
	return m.LabelType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastTriggerControlLabelOptions(structType any) TriggerControlLabelOptions {
	if casted, ok := structType.(TriggerControlLabelOptions); ok {
		return casted
	}
	if casted, ok := structType.(*TriggerControlLabelOptions); ok {
		return *casted
	}
	return nil
}

func (m *_TriggerControlLabelOptions) GetTypeName() string {
	return "TriggerControlLabelOptions"
}

func (m *_TriggerControlLabelOptions) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (labelFlavour)
	lengthInBits += 2

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (labelType)
	lengthInBits += 2

	// Reserved Field (reserved)
	lengthInBits += 1

	return lengthInBits
}

func (m *_TriggerControlLabelOptions) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TriggerControlLabelOptionsParse(ctx context.Context, theBytes []byte) (TriggerControlLabelOptions, error) {
	return TriggerControlLabelOptionsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TriggerControlLabelOptionsParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (TriggerControlLabelOptions, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (TriggerControlLabelOptions, error) {
		return TriggerControlLabelOptionsParseWithBuffer(ctx, readBuffer)
	}
}

func TriggerControlLabelOptionsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TriggerControlLabelOptions, error) {
	v, err := (&_TriggerControlLabelOptions{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_TriggerControlLabelOptions) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__triggerControlLabelOptions TriggerControlLabelOptions, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TriggerControlLabelOptions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TriggerControlLabelOptions")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	labelFlavour, err := ReadEnumField[TriggerControlLabelFlavour](ctx, "labelFlavour", "TriggerControlLabelFlavour", ReadEnum(TriggerControlLabelFlavourByValue, ReadUnsignedByte(readBuffer, uint8(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'labelFlavour' field"))
	}
	m.LabelFlavour = labelFlavour

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField1 = reservedField1

	reservedField2, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField2 = reservedField2

	labelType, err := ReadEnumField[TriggerControlLabelType](ctx, "labelType", "TriggerControlLabelType", ReadEnum(TriggerControlLabelTypeByValue, ReadUnsignedByte(readBuffer, uint8(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'labelType' field"))
	}
	m.LabelType = labelType

	reservedField3, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField3 = reservedField3

	if closeErr := readBuffer.CloseContext("TriggerControlLabelOptions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TriggerControlLabelOptions")
	}

	return m, nil
}

func (m *_TriggerControlLabelOptions) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TriggerControlLabelOptions) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("TriggerControlLabelOptions"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TriggerControlLabelOptions")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if err := WriteSimpleEnumField[TriggerControlLabelFlavour](ctx, "labelFlavour", "TriggerControlLabelFlavour", m.GetLabelFlavour(), WriteEnum[TriggerControlLabelFlavour, uint8](TriggerControlLabelFlavour.GetValue, TriggerControlLabelFlavour.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 2))); err != nil {
		return errors.Wrap(err, "Error serializing 'labelFlavour' field")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 2")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 3")
	}

	if err := WriteSimpleEnumField[TriggerControlLabelType](ctx, "labelType", "TriggerControlLabelType", m.GetLabelType(), WriteEnum[TriggerControlLabelType, uint8](TriggerControlLabelType.GetValue, TriggerControlLabelType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 2))); err != nil {
		return errors.Wrap(err, "Error serializing 'labelType' field")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 4")
	}

	if popErr := writeBuffer.PopContext("TriggerControlLabelOptions"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TriggerControlLabelOptions")
	}
	return nil
}

func (m *_TriggerControlLabelOptions) IsTriggerControlLabelOptions() {}

func (m *_TriggerControlLabelOptions) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TriggerControlLabelOptions) deepCopy() *_TriggerControlLabelOptions {
	if m == nil {
		return nil
	}
	_TriggerControlLabelOptionsCopy := &_TriggerControlLabelOptions{
		m.LabelFlavour,
		m.LabelType,
		m.reservedField0,
		m.reservedField1,
		m.reservedField2,
		m.reservedField3,
	}
	return _TriggerControlLabelOptionsCopy
}

func (m *_TriggerControlLabelOptions) String() string {
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