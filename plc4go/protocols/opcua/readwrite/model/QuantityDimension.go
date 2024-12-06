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

// QuantityDimension is the corresponding interface of QuantityDimension
type QuantityDimension interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetMassExponent returns MassExponent (property field)
	GetMassExponent() int8
	// GetLengthExponent returns LengthExponent (property field)
	GetLengthExponent() int8
	// GetTimeExponent returns TimeExponent (property field)
	GetTimeExponent() int8
	// GetElectricCurrentExponent returns ElectricCurrentExponent (property field)
	GetElectricCurrentExponent() int8
	// GetAmountOfSubstanceExponent returns AmountOfSubstanceExponent (property field)
	GetAmountOfSubstanceExponent() int8
	// GetLuminousIntensityExponent returns LuminousIntensityExponent (property field)
	GetLuminousIntensityExponent() int8
	// GetAbsoluteTemperatureExponent returns AbsoluteTemperatureExponent (property field)
	GetAbsoluteTemperatureExponent() int8
	// GetDimensionlessExponent returns DimensionlessExponent (property field)
	GetDimensionlessExponent() int8
	// IsQuantityDimension is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsQuantityDimension()
	// CreateBuilder creates a QuantityDimensionBuilder
	CreateQuantityDimensionBuilder() QuantityDimensionBuilder
}

// _QuantityDimension is the data-structure of this message
type _QuantityDimension struct {
	ExtensionObjectDefinitionContract
	MassExponent                int8
	LengthExponent              int8
	TimeExponent                int8
	ElectricCurrentExponent     int8
	AmountOfSubstanceExponent   int8
	LuminousIntensityExponent   int8
	AbsoluteTemperatureExponent int8
	DimensionlessExponent       int8
}

var _ QuantityDimension = (*_QuantityDimension)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_QuantityDimension)(nil)

// NewQuantityDimension factory function for _QuantityDimension
func NewQuantityDimension(massExponent int8, lengthExponent int8, timeExponent int8, electricCurrentExponent int8, amountOfSubstanceExponent int8, luminousIntensityExponent int8, absoluteTemperatureExponent int8, dimensionlessExponent int8) *_QuantityDimension {
	_result := &_QuantityDimension{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		MassExponent:                      massExponent,
		LengthExponent:                    lengthExponent,
		TimeExponent:                      timeExponent,
		ElectricCurrentExponent:           electricCurrentExponent,
		AmountOfSubstanceExponent:         amountOfSubstanceExponent,
		LuminousIntensityExponent:         luminousIntensityExponent,
		AbsoluteTemperatureExponent:       absoluteTemperatureExponent,
		DimensionlessExponent:             dimensionlessExponent,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// QuantityDimensionBuilder is a builder for QuantityDimension
type QuantityDimensionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(massExponent int8, lengthExponent int8, timeExponent int8, electricCurrentExponent int8, amountOfSubstanceExponent int8, luminousIntensityExponent int8, absoluteTemperatureExponent int8, dimensionlessExponent int8) QuantityDimensionBuilder
	// WithMassExponent adds MassExponent (property field)
	WithMassExponent(int8) QuantityDimensionBuilder
	// WithLengthExponent adds LengthExponent (property field)
	WithLengthExponent(int8) QuantityDimensionBuilder
	// WithTimeExponent adds TimeExponent (property field)
	WithTimeExponent(int8) QuantityDimensionBuilder
	// WithElectricCurrentExponent adds ElectricCurrentExponent (property field)
	WithElectricCurrentExponent(int8) QuantityDimensionBuilder
	// WithAmountOfSubstanceExponent adds AmountOfSubstanceExponent (property field)
	WithAmountOfSubstanceExponent(int8) QuantityDimensionBuilder
	// WithLuminousIntensityExponent adds LuminousIntensityExponent (property field)
	WithLuminousIntensityExponent(int8) QuantityDimensionBuilder
	// WithAbsoluteTemperatureExponent adds AbsoluteTemperatureExponent (property field)
	WithAbsoluteTemperatureExponent(int8) QuantityDimensionBuilder
	// WithDimensionlessExponent adds DimensionlessExponent (property field)
	WithDimensionlessExponent(int8) QuantityDimensionBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the QuantityDimension or returns an error if something is wrong
	Build() (QuantityDimension, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() QuantityDimension
}

// NewQuantityDimensionBuilder() creates a QuantityDimensionBuilder
func NewQuantityDimensionBuilder() QuantityDimensionBuilder {
	return &_QuantityDimensionBuilder{_QuantityDimension: new(_QuantityDimension)}
}

type _QuantityDimensionBuilder struct {
	*_QuantityDimension

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (QuantityDimensionBuilder) = (*_QuantityDimensionBuilder)(nil)

func (b *_QuantityDimensionBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._QuantityDimension
}

func (b *_QuantityDimensionBuilder) WithMandatoryFields(massExponent int8, lengthExponent int8, timeExponent int8, electricCurrentExponent int8, amountOfSubstanceExponent int8, luminousIntensityExponent int8, absoluteTemperatureExponent int8, dimensionlessExponent int8) QuantityDimensionBuilder {
	return b.WithMassExponent(massExponent).WithLengthExponent(lengthExponent).WithTimeExponent(timeExponent).WithElectricCurrentExponent(electricCurrentExponent).WithAmountOfSubstanceExponent(amountOfSubstanceExponent).WithLuminousIntensityExponent(luminousIntensityExponent).WithAbsoluteTemperatureExponent(absoluteTemperatureExponent).WithDimensionlessExponent(dimensionlessExponent)
}

func (b *_QuantityDimensionBuilder) WithMassExponent(massExponent int8) QuantityDimensionBuilder {
	b.MassExponent = massExponent
	return b
}

func (b *_QuantityDimensionBuilder) WithLengthExponent(lengthExponent int8) QuantityDimensionBuilder {
	b.LengthExponent = lengthExponent
	return b
}

func (b *_QuantityDimensionBuilder) WithTimeExponent(timeExponent int8) QuantityDimensionBuilder {
	b.TimeExponent = timeExponent
	return b
}

func (b *_QuantityDimensionBuilder) WithElectricCurrentExponent(electricCurrentExponent int8) QuantityDimensionBuilder {
	b.ElectricCurrentExponent = electricCurrentExponent
	return b
}

func (b *_QuantityDimensionBuilder) WithAmountOfSubstanceExponent(amountOfSubstanceExponent int8) QuantityDimensionBuilder {
	b.AmountOfSubstanceExponent = amountOfSubstanceExponent
	return b
}

func (b *_QuantityDimensionBuilder) WithLuminousIntensityExponent(luminousIntensityExponent int8) QuantityDimensionBuilder {
	b.LuminousIntensityExponent = luminousIntensityExponent
	return b
}

func (b *_QuantityDimensionBuilder) WithAbsoluteTemperatureExponent(absoluteTemperatureExponent int8) QuantityDimensionBuilder {
	b.AbsoluteTemperatureExponent = absoluteTemperatureExponent
	return b
}

func (b *_QuantityDimensionBuilder) WithDimensionlessExponent(dimensionlessExponent int8) QuantityDimensionBuilder {
	b.DimensionlessExponent = dimensionlessExponent
	return b
}

func (b *_QuantityDimensionBuilder) Build() (QuantityDimension, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._QuantityDimension.deepCopy(), nil
}

func (b *_QuantityDimensionBuilder) MustBuild() QuantityDimension {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_QuantityDimensionBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_QuantityDimensionBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_QuantityDimensionBuilder) DeepCopy() any {
	_copy := b.CreateQuantityDimensionBuilder().(*_QuantityDimensionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateQuantityDimensionBuilder creates a QuantityDimensionBuilder
func (b *_QuantityDimension) CreateQuantityDimensionBuilder() QuantityDimensionBuilder {
	if b == nil {
		return NewQuantityDimensionBuilder()
	}
	return &_QuantityDimensionBuilder{_QuantityDimension: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_QuantityDimension) GetExtensionId() int32 {
	return int32(32440)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_QuantityDimension) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_QuantityDimension) GetMassExponent() int8 {
	return m.MassExponent
}

func (m *_QuantityDimension) GetLengthExponent() int8 {
	return m.LengthExponent
}

func (m *_QuantityDimension) GetTimeExponent() int8 {
	return m.TimeExponent
}

func (m *_QuantityDimension) GetElectricCurrentExponent() int8 {
	return m.ElectricCurrentExponent
}

func (m *_QuantityDimension) GetAmountOfSubstanceExponent() int8 {
	return m.AmountOfSubstanceExponent
}

func (m *_QuantityDimension) GetLuminousIntensityExponent() int8 {
	return m.LuminousIntensityExponent
}

func (m *_QuantityDimension) GetAbsoluteTemperatureExponent() int8 {
	return m.AbsoluteTemperatureExponent
}

func (m *_QuantityDimension) GetDimensionlessExponent() int8 {
	return m.DimensionlessExponent
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastQuantityDimension(structType any) QuantityDimension {
	if casted, ok := structType.(QuantityDimension); ok {
		return casted
	}
	if casted, ok := structType.(*QuantityDimension); ok {
		return *casted
	}
	return nil
}

func (m *_QuantityDimension) GetTypeName() string {
	return "QuantityDimension"
}

func (m *_QuantityDimension) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (massExponent)
	lengthInBits += 8

	// Simple field (lengthExponent)
	lengthInBits += 8

	// Simple field (timeExponent)
	lengthInBits += 8

	// Simple field (electricCurrentExponent)
	lengthInBits += 8

	// Simple field (amountOfSubstanceExponent)
	lengthInBits += 8

	// Simple field (luminousIntensityExponent)
	lengthInBits += 8

	// Simple field (absoluteTemperatureExponent)
	lengthInBits += 8

	// Simple field (dimensionlessExponent)
	lengthInBits += 8

	return lengthInBits
}

func (m *_QuantityDimension) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_QuantityDimension) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__quantityDimension QuantityDimension, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("QuantityDimension"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for QuantityDimension")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	massExponent, err := ReadSimpleField(ctx, "massExponent", ReadSignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'massExponent' field"))
	}
	m.MassExponent = massExponent

	lengthExponent, err := ReadSimpleField(ctx, "lengthExponent", ReadSignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lengthExponent' field"))
	}
	m.LengthExponent = lengthExponent

	timeExponent, err := ReadSimpleField(ctx, "timeExponent", ReadSignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeExponent' field"))
	}
	m.TimeExponent = timeExponent

	electricCurrentExponent, err := ReadSimpleField(ctx, "electricCurrentExponent", ReadSignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'electricCurrentExponent' field"))
	}
	m.ElectricCurrentExponent = electricCurrentExponent

	amountOfSubstanceExponent, err := ReadSimpleField(ctx, "amountOfSubstanceExponent", ReadSignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'amountOfSubstanceExponent' field"))
	}
	m.AmountOfSubstanceExponent = amountOfSubstanceExponent

	luminousIntensityExponent, err := ReadSimpleField(ctx, "luminousIntensityExponent", ReadSignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'luminousIntensityExponent' field"))
	}
	m.LuminousIntensityExponent = luminousIntensityExponent

	absoluteTemperatureExponent, err := ReadSimpleField(ctx, "absoluteTemperatureExponent", ReadSignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'absoluteTemperatureExponent' field"))
	}
	m.AbsoluteTemperatureExponent = absoluteTemperatureExponent

	dimensionlessExponent, err := ReadSimpleField(ctx, "dimensionlessExponent", ReadSignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dimensionlessExponent' field"))
	}
	m.DimensionlessExponent = dimensionlessExponent

	if closeErr := readBuffer.CloseContext("QuantityDimension"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for QuantityDimension")
	}

	return m, nil
}

func (m *_QuantityDimension) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_QuantityDimension) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("QuantityDimension"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for QuantityDimension")
		}

		if err := WriteSimpleField[int8](ctx, "massExponent", m.GetMassExponent(), WriteSignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'massExponent' field")
		}

		if err := WriteSimpleField[int8](ctx, "lengthExponent", m.GetLengthExponent(), WriteSignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'lengthExponent' field")
		}

		if err := WriteSimpleField[int8](ctx, "timeExponent", m.GetTimeExponent(), WriteSignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeExponent' field")
		}

		if err := WriteSimpleField[int8](ctx, "electricCurrentExponent", m.GetElectricCurrentExponent(), WriteSignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'electricCurrentExponent' field")
		}

		if err := WriteSimpleField[int8](ctx, "amountOfSubstanceExponent", m.GetAmountOfSubstanceExponent(), WriteSignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'amountOfSubstanceExponent' field")
		}

		if err := WriteSimpleField[int8](ctx, "luminousIntensityExponent", m.GetLuminousIntensityExponent(), WriteSignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'luminousIntensityExponent' field")
		}

		if err := WriteSimpleField[int8](ctx, "absoluteTemperatureExponent", m.GetAbsoluteTemperatureExponent(), WriteSignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'absoluteTemperatureExponent' field")
		}

		if err := WriteSimpleField[int8](ctx, "dimensionlessExponent", m.GetDimensionlessExponent(), WriteSignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'dimensionlessExponent' field")
		}

		if popErr := writeBuffer.PopContext("QuantityDimension"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for QuantityDimension")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_QuantityDimension) IsQuantityDimension() {}

func (m *_QuantityDimension) DeepCopy() any {
	return m.deepCopy()
}

func (m *_QuantityDimension) deepCopy() *_QuantityDimension {
	if m == nil {
		return nil
	}
	_QuantityDimensionCopy := &_QuantityDimension{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.MassExponent,
		m.LengthExponent,
		m.TimeExponent,
		m.ElectricCurrentExponent,
		m.AmountOfSubstanceExponent,
		m.LuminousIntensityExponent,
		m.AbsoluteTemperatureExponent,
		m.DimensionlessExponent,
	}
	_QuantityDimensionCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _QuantityDimensionCopy
}

func (m *_QuantityDimension) String() string {
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