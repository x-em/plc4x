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

// CALDataWrite is the corresponding interface of CALDataWrite
type CALDataWrite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CALData
	// GetParamNo returns ParamNo (property field)
	GetParamNo() Parameter
	// GetCode returns Code (property field)
	GetCode() byte
	// GetParameterValue returns ParameterValue (property field)
	GetParameterValue() ParameterValue
	// IsCALDataWrite is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCALDataWrite()
	// CreateBuilder creates a CALDataWriteBuilder
	CreateCALDataWriteBuilder() CALDataWriteBuilder
}

// _CALDataWrite is the data-structure of this message
type _CALDataWrite struct {
	CALDataContract
	ParamNo        Parameter
	Code           byte
	ParameterValue ParameterValue
}

var _ CALDataWrite = (*_CALDataWrite)(nil)
var _ CALDataRequirements = (*_CALDataWrite)(nil)

// NewCALDataWrite factory function for _CALDataWrite
func NewCALDataWrite(commandTypeContainer CALCommandTypeContainer, additionalData CALData, paramNo Parameter, code byte, parameterValue ParameterValue, requestContext RequestContext) *_CALDataWrite {
	if parameterValue == nil {
		panic("parameterValue of type ParameterValue for CALDataWrite must not be nil")
	}
	_result := &_CALDataWrite{
		CALDataContract: NewCALData(commandTypeContainer, additionalData, requestContext),
		ParamNo:         paramNo,
		Code:            code,
		ParameterValue:  parameterValue,
	}
	_result.CALDataContract.(*_CALData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CALDataWriteBuilder is a builder for CALDataWrite
type CALDataWriteBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(paramNo Parameter, code byte, parameterValue ParameterValue) CALDataWriteBuilder
	// WithParamNo adds ParamNo (property field)
	WithParamNo(Parameter) CALDataWriteBuilder
	// WithCode adds Code (property field)
	WithCode(byte) CALDataWriteBuilder
	// WithParameterValue adds ParameterValue (property field)
	WithParameterValue(ParameterValue) CALDataWriteBuilder
	// WithParameterValueBuilder adds ParameterValue (property field) which is build by the builder
	WithParameterValueBuilder(func(ParameterValueBuilder) ParameterValueBuilder) CALDataWriteBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CALDataBuilder
	// Build builds the CALDataWrite or returns an error if something is wrong
	Build() (CALDataWrite, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CALDataWrite
}

// NewCALDataWriteBuilder() creates a CALDataWriteBuilder
func NewCALDataWriteBuilder() CALDataWriteBuilder {
	return &_CALDataWriteBuilder{_CALDataWrite: new(_CALDataWrite)}
}

type _CALDataWriteBuilder struct {
	*_CALDataWrite

	parentBuilder *_CALDataBuilder

	err *utils.MultiError
}

var _ (CALDataWriteBuilder) = (*_CALDataWriteBuilder)(nil)

func (b *_CALDataWriteBuilder) setParent(contract CALDataContract) {
	b.CALDataContract = contract
	contract.(*_CALData)._SubType = b._CALDataWrite
}

func (b *_CALDataWriteBuilder) WithMandatoryFields(paramNo Parameter, code byte, parameterValue ParameterValue) CALDataWriteBuilder {
	return b.WithParamNo(paramNo).WithCode(code).WithParameterValue(parameterValue)
}

func (b *_CALDataWriteBuilder) WithParamNo(paramNo Parameter) CALDataWriteBuilder {
	b.ParamNo = paramNo
	return b
}

func (b *_CALDataWriteBuilder) WithCode(code byte) CALDataWriteBuilder {
	b.Code = code
	return b
}

func (b *_CALDataWriteBuilder) WithParameterValue(parameterValue ParameterValue) CALDataWriteBuilder {
	b.ParameterValue = parameterValue
	return b
}

func (b *_CALDataWriteBuilder) WithParameterValueBuilder(builderSupplier func(ParameterValueBuilder) ParameterValueBuilder) CALDataWriteBuilder {
	builder := builderSupplier(b.ParameterValue.CreateParameterValueBuilder())
	var err error
	b.ParameterValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ParameterValueBuilder failed"))
	}
	return b
}

func (b *_CALDataWriteBuilder) Build() (CALDataWrite, error) {
	if b.ParameterValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'parameterValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CALDataWrite.deepCopy(), nil
}

func (b *_CALDataWriteBuilder) MustBuild() CALDataWrite {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CALDataWriteBuilder) Done() CALDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCALDataBuilder().(*_CALDataBuilder)
	}
	return b.parentBuilder
}

func (b *_CALDataWriteBuilder) buildForCALData() (CALData, error) {
	return b.Build()
}

func (b *_CALDataWriteBuilder) DeepCopy() any {
	_copy := b.CreateCALDataWriteBuilder().(*_CALDataWriteBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCALDataWriteBuilder creates a CALDataWriteBuilder
func (b *_CALDataWrite) CreateCALDataWriteBuilder() CALDataWriteBuilder {
	if b == nil {
		return NewCALDataWriteBuilder()
	}
	return &_CALDataWriteBuilder{_CALDataWrite: b.deepCopy()}
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

func (m *_CALDataWrite) GetParent() CALDataContract {
	return m.CALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataWrite) GetParamNo() Parameter {
	return m.ParamNo
}

func (m *_CALDataWrite) GetCode() byte {
	return m.Code
}

func (m *_CALDataWrite) GetParameterValue() ParameterValue {
	return m.ParameterValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCALDataWrite(structType any) CALDataWrite {
	if casted, ok := structType.(CALDataWrite); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataWrite); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataWrite) GetTypeName() string {
	return "CALDataWrite"
}

func (m *_CALDataWrite) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CALDataContract.(*_CALData).getLengthInBits(ctx))

	// Simple field (paramNo)
	lengthInBits += 8

	// Simple field (code)
	lengthInBits += 8

	// Simple field (parameterValue)
	lengthInBits += m.ParameterValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CALDataWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CALDataWrite) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CALData, commandTypeContainer CALCommandTypeContainer, requestContext RequestContext) (__cALDataWrite CALDataWrite, err error) {
	m.CALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	paramNo, err := ReadEnumField[Parameter](ctx, "paramNo", "Parameter", ReadEnum(ParameterByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'paramNo' field"))
	}
	m.ParamNo = paramNo

	code, err := ReadSimpleField(ctx, "code", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'code' field"))
	}
	m.Code = code

	parameterValue, err := ReadSimpleField[ParameterValue](ctx, "parameterValue", ReadComplex[ParameterValue](ParameterValueParseWithBufferProducer[ParameterValue]((ParameterType)(paramNo.ParameterType()), (uint8)(uint8(commandTypeContainer.NumBytes())-uint8(uint8(2)))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parameterValue' field"))
	}
	m.ParameterValue = parameterValue

	if closeErr := readBuffer.CloseContext("CALDataWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataWrite")
	}

	return m, nil
}

func (m *_CALDataWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALDataWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataWrite")
		}

		if err := WriteSimpleEnumField[Parameter](ctx, "paramNo", "Parameter", m.GetParamNo(), WriteEnum[Parameter, uint8](Parameter.GetValue, Parameter.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'paramNo' field")
		}

		if err := WriteSimpleField[byte](ctx, "code", m.GetCode(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'code' field")
		}

		if err := WriteSimpleField[ParameterValue](ctx, "parameterValue", m.GetParameterValue(), WriteComplex[ParameterValue](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'parameterValue' field")
		}

		if popErr := writeBuffer.PopContext("CALDataWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataWrite")
		}
		return nil
	}
	return m.CALDataContract.(*_CALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CALDataWrite) IsCALDataWrite() {}

func (m *_CALDataWrite) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CALDataWrite) deepCopy() *_CALDataWrite {
	if m == nil {
		return nil
	}
	_CALDataWriteCopy := &_CALDataWrite{
		m.CALDataContract.(*_CALData).deepCopy(),
		m.ParamNo,
		m.Code,
		utils.DeepCopy[ParameterValue](m.ParameterValue),
	}
	_CALDataWriteCopy.CALDataContract.(*_CALData)._SubType = m
	return _CALDataWriteCopy
}

func (m *_CALDataWrite) String() string {
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
