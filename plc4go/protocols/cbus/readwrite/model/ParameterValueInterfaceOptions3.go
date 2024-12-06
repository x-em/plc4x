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

// ParameterValueInterfaceOptions3 is the corresponding interface of ParameterValueInterfaceOptions3
type ParameterValueInterfaceOptions3 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ParameterValue
	// GetValue returns Value (property field)
	GetValue() InterfaceOptions3
	// GetData returns Data (property field)
	GetData() []byte
	// IsParameterValueInterfaceOptions3 is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsParameterValueInterfaceOptions3()
	// CreateBuilder creates a ParameterValueInterfaceOptions3Builder
	CreateParameterValueInterfaceOptions3Builder() ParameterValueInterfaceOptions3Builder
}

// _ParameterValueInterfaceOptions3 is the data-structure of this message
type _ParameterValueInterfaceOptions3 struct {
	ParameterValueContract
	Value InterfaceOptions3
	Data  []byte
}

var _ ParameterValueInterfaceOptions3 = (*_ParameterValueInterfaceOptions3)(nil)
var _ ParameterValueRequirements = (*_ParameterValueInterfaceOptions3)(nil)

// NewParameterValueInterfaceOptions3 factory function for _ParameterValueInterfaceOptions3
func NewParameterValueInterfaceOptions3(value InterfaceOptions3, data []byte, numBytes uint8) *_ParameterValueInterfaceOptions3 {
	if value == nil {
		panic("value of type InterfaceOptions3 for ParameterValueInterfaceOptions3 must not be nil")
	}
	_result := &_ParameterValueInterfaceOptions3{
		ParameterValueContract: NewParameterValue(numBytes),
		Value:                  value,
		Data:                   data,
	}
	_result.ParameterValueContract.(*_ParameterValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ParameterValueInterfaceOptions3Builder is a builder for ParameterValueInterfaceOptions3
type ParameterValueInterfaceOptions3Builder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value InterfaceOptions3, data []byte) ParameterValueInterfaceOptions3Builder
	// WithValue adds Value (property field)
	WithValue(InterfaceOptions3) ParameterValueInterfaceOptions3Builder
	// WithValueBuilder adds Value (property field) which is build by the builder
	WithValueBuilder(func(InterfaceOptions3Builder) InterfaceOptions3Builder) ParameterValueInterfaceOptions3Builder
	// WithData adds Data (property field)
	WithData(...byte) ParameterValueInterfaceOptions3Builder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ParameterValueBuilder
	// Build builds the ParameterValueInterfaceOptions3 or returns an error if something is wrong
	Build() (ParameterValueInterfaceOptions3, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ParameterValueInterfaceOptions3
}

// NewParameterValueInterfaceOptions3Builder() creates a ParameterValueInterfaceOptions3Builder
func NewParameterValueInterfaceOptions3Builder() ParameterValueInterfaceOptions3Builder {
	return &_ParameterValueInterfaceOptions3Builder{_ParameterValueInterfaceOptions3: new(_ParameterValueInterfaceOptions3)}
}

type _ParameterValueInterfaceOptions3Builder struct {
	*_ParameterValueInterfaceOptions3

	parentBuilder *_ParameterValueBuilder

	err *utils.MultiError
}

var _ (ParameterValueInterfaceOptions3Builder) = (*_ParameterValueInterfaceOptions3Builder)(nil)

func (b *_ParameterValueInterfaceOptions3Builder) setParent(contract ParameterValueContract) {
	b.ParameterValueContract = contract
	contract.(*_ParameterValue)._SubType = b._ParameterValueInterfaceOptions3
}

func (b *_ParameterValueInterfaceOptions3Builder) WithMandatoryFields(value InterfaceOptions3, data []byte) ParameterValueInterfaceOptions3Builder {
	return b.WithValue(value).WithData(data...)
}

func (b *_ParameterValueInterfaceOptions3Builder) WithValue(value InterfaceOptions3) ParameterValueInterfaceOptions3Builder {
	b.Value = value
	return b
}

func (b *_ParameterValueInterfaceOptions3Builder) WithValueBuilder(builderSupplier func(InterfaceOptions3Builder) InterfaceOptions3Builder) ParameterValueInterfaceOptions3Builder {
	builder := builderSupplier(b.Value.CreateInterfaceOptions3Builder())
	var err error
	b.Value, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "InterfaceOptions3Builder failed"))
	}
	return b
}

func (b *_ParameterValueInterfaceOptions3Builder) WithData(data ...byte) ParameterValueInterfaceOptions3Builder {
	b.Data = data
	return b
}

func (b *_ParameterValueInterfaceOptions3Builder) Build() (ParameterValueInterfaceOptions3, error) {
	if b.Value == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'value' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ParameterValueInterfaceOptions3.deepCopy(), nil
}

func (b *_ParameterValueInterfaceOptions3Builder) MustBuild() ParameterValueInterfaceOptions3 {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ParameterValueInterfaceOptions3Builder) Done() ParameterValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewParameterValueBuilder().(*_ParameterValueBuilder)
	}
	return b.parentBuilder
}

func (b *_ParameterValueInterfaceOptions3Builder) buildForParameterValue() (ParameterValue, error) {
	return b.Build()
}

func (b *_ParameterValueInterfaceOptions3Builder) DeepCopy() any {
	_copy := b.CreateParameterValueInterfaceOptions3Builder().(*_ParameterValueInterfaceOptions3Builder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateParameterValueInterfaceOptions3Builder creates a ParameterValueInterfaceOptions3Builder
func (b *_ParameterValueInterfaceOptions3) CreateParameterValueInterfaceOptions3Builder() ParameterValueInterfaceOptions3Builder {
	if b == nil {
		return NewParameterValueInterfaceOptions3Builder()
	}
	return &_ParameterValueInterfaceOptions3Builder{_ParameterValueInterfaceOptions3: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ParameterValueInterfaceOptions3) GetParameterType() ParameterType {
	return ParameterType_INTERFACE_OPTIONS_3
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ParameterValueInterfaceOptions3) GetParent() ParameterValueContract {
	return m.ParameterValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ParameterValueInterfaceOptions3) GetValue() InterfaceOptions3 {
	return m.Value
}

func (m *_ParameterValueInterfaceOptions3) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastParameterValueInterfaceOptions3(structType any) ParameterValueInterfaceOptions3 {
	if casted, ok := structType.(ParameterValueInterfaceOptions3); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterValueInterfaceOptions3); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterValueInterfaceOptions3) GetTypeName() string {
	return "ParameterValueInterfaceOptions3"
}

func (m *_ParameterValueInterfaceOptions3) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ParameterValueContract.(*_ParameterValue).getLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ParameterValueInterfaceOptions3) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ParameterValueInterfaceOptions3) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ParameterValue, parameterType ParameterType, numBytes uint8) (__parameterValueInterfaceOptions3 ParameterValueInterfaceOptions3, err error) {
	m.ParameterValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterValueInterfaceOptions3"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterValueInterfaceOptions3")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((numBytes) >= (1))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "InterfaceOptions3 has exactly one byte"})
	}

	value, err := ReadSimpleField[InterfaceOptions3](ctx, "value", ReadComplex[InterfaceOptions3](InterfaceOptions3ParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	data, err := readBuffer.ReadByteArray("data", int(int32(numBytes)-int32(int32(1))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("ParameterValueInterfaceOptions3"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterValueInterfaceOptions3")
	}

	return m, nil
}

func (m *_ParameterValueInterfaceOptions3) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ParameterValueInterfaceOptions3) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ParameterValueInterfaceOptions3"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ParameterValueInterfaceOptions3")
		}

		if err := WriteSimpleField[InterfaceOptions3](ctx, "value", m.GetValue(), WriteComplex[InterfaceOptions3](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ParameterValueInterfaceOptions3"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ParameterValueInterfaceOptions3")
		}
		return nil
	}
	return m.ParameterValueContract.(*_ParameterValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ParameterValueInterfaceOptions3) IsParameterValueInterfaceOptions3() {}

func (m *_ParameterValueInterfaceOptions3) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ParameterValueInterfaceOptions3) deepCopy() *_ParameterValueInterfaceOptions3 {
	if m == nil {
		return nil
	}
	_ParameterValueInterfaceOptions3Copy := &_ParameterValueInterfaceOptions3{
		m.ParameterValueContract.(*_ParameterValue).deepCopy(),
		utils.DeepCopy[InterfaceOptions3](m.Value),
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	_ParameterValueInterfaceOptions3Copy.ParameterValueContract.(*_ParameterValue)._SubType = m
	return _ParameterValueInterfaceOptions3Copy
}

func (m *_ParameterValueInterfaceOptions3) String() string {
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