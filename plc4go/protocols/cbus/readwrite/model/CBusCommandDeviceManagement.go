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

// Constant values.
const CBusCommandDeviceManagement_DELIMITER byte = 0x0

// CBusCommandDeviceManagement is the corresponding interface of CBusCommandDeviceManagement
type CBusCommandDeviceManagement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CBusCommand
	// GetParamNo returns ParamNo (property field)
	GetParamNo() Parameter
	// GetParameterValue returns ParameterValue (property field)
	GetParameterValue() byte
	// IsCBusCommandDeviceManagement is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCBusCommandDeviceManagement()
	// CreateBuilder creates a CBusCommandDeviceManagementBuilder
	CreateCBusCommandDeviceManagementBuilder() CBusCommandDeviceManagementBuilder
}

// _CBusCommandDeviceManagement is the data-structure of this message
type _CBusCommandDeviceManagement struct {
	CBusCommandContract
	ParamNo        Parameter
	ParameterValue byte
}

var _ CBusCommandDeviceManagement = (*_CBusCommandDeviceManagement)(nil)
var _ CBusCommandRequirements = (*_CBusCommandDeviceManagement)(nil)

// NewCBusCommandDeviceManagement factory function for _CBusCommandDeviceManagement
func NewCBusCommandDeviceManagement(header CBusHeader, paramNo Parameter, parameterValue byte, cBusOptions CBusOptions) *_CBusCommandDeviceManagement {
	_result := &_CBusCommandDeviceManagement{
		CBusCommandContract: NewCBusCommand(header, cBusOptions),
		ParamNo:             paramNo,
		ParameterValue:      parameterValue,
	}
	_result.CBusCommandContract.(*_CBusCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CBusCommandDeviceManagementBuilder is a builder for CBusCommandDeviceManagement
type CBusCommandDeviceManagementBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(paramNo Parameter, parameterValue byte) CBusCommandDeviceManagementBuilder
	// WithParamNo adds ParamNo (property field)
	WithParamNo(Parameter) CBusCommandDeviceManagementBuilder
	// WithParameterValue adds ParameterValue (property field)
	WithParameterValue(byte) CBusCommandDeviceManagementBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CBusCommandBuilder
	// Build builds the CBusCommandDeviceManagement or returns an error if something is wrong
	Build() (CBusCommandDeviceManagement, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CBusCommandDeviceManagement
}

// NewCBusCommandDeviceManagementBuilder() creates a CBusCommandDeviceManagementBuilder
func NewCBusCommandDeviceManagementBuilder() CBusCommandDeviceManagementBuilder {
	return &_CBusCommandDeviceManagementBuilder{_CBusCommandDeviceManagement: new(_CBusCommandDeviceManagement)}
}

type _CBusCommandDeviceManagementBuilder struct {
	*_CBusCommandDeviceManagement

	parentBuilder *_CBusCommandBuilder

	err *utils.MultiError
}

var _ (CBusCommandDeviceManagementBuilder) = (*_CBusCommandDeviceManagementBuilder)(nil)

func (b *_CBusCommandDeviceManagementBuilder) setParent(contract CBusCommandContract) {
	b.CBusCommandContract = contract
	contract.(*_CBusCommand)._SubType = b._CBusCommandDeviceManagement
}

func (b *_CBusCommandDeviceManagementBuilder) WithMandatoryFields(paramNo Parameter, parameterValue byte) CBusCommandDeviceManagementBuilder {
	return b.WithParamNo(paramNo).WithParameterValue(parameterValue)
}

func (b *_CBusCommandDeviceManagementBuilder) WithParamNo(paramNo Parameter) CBusCommandDeviceManagementBuilder {
	b.ParamNo = paramNo
	return b
}

func (b *_CBusCommandDeviceManagementBuilder) WithParameterValue(parameterValue byte) CBusCommandDeviceManagementBuilder {
	b.ParameterValue = parameterValue
	return b
}

func (b *_CBusCommandDeviceManagementBuilder) Build() (CBusCommandDeviceManagement, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CBusCommandDeviceManagement.deepCopy(), nil
}

func (b *_CBusCommandDeviceManagementBuilder) MustBuild() CBusCommandDeviceManagement {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CBusCommandDeviceManagementBuilder) Done() CBusCommandBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCBusCommandBuilder().(*_CBusCommandBuilder)
	}
	return b.parentBuilder
}

func (b *_CBusCommandDeviceManagementBuilder) buildForCBusCommand() (CBusCommand, error) {
	return b.Build()
}

func (b *_CBusCommandDeviceManagementBuilder) DeepCopy() any {
	_copy := b.CreateCBusCommandDeviceManagementBuilder().(*_CBusCommandDeviceManagementBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCBusCommandDeviceManagementBuilder creates a CBusCommandDeviceManagementBuilder
func (b *_CBusCommandDeviceManagement) CreateCBusCommandDeviceManagementBuilder() CBusCommandDeviceManagementBuilder {
	if b == nil {
		return NewCBusCommandDeviceManagementBuilder()
	}
	return &_CBusCommandDeviceManagementBuilder{_CBusCommandDeviceManagement: b.deepCopy()}
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

func (m *_CBusCommandDeviceManagement) GetParent() CBusCommandContract {
	return m.CBusCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusCommandDeviceManagement) GetParamNo() Parameter {
	return m.ParamNo
}

func (m *_CBusCommandDeviceManagement) GetParameterValue() byte {
	return m.ParameterValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_CBusCommandDeviceManagement) GetDelimiter() byte {
	return CBusCommandDeviceManagement_DELIMITER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCBusCommandDeviceManagement(structType any) CBusCommandDeviceManagement {
	if casted, ok := structType.(CBusCommandDeviceManagement); ok {
		return casted
	}
	if casted, ok := structType.(*CBusCommandDeviceManagement); ok {
		return *casted
	}
	return nil
}

func (m *_CBusCommandDeviceManagement) GetTypeName() string {
	return "CBusCommandDeviceManagement"
}

func (m *_CBusCommandDeviceManagement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CBusCommandContract.(*_CBusCommand).getLengthInBits(ctx))

	// Simple field (paramNo)
	lengthInBits += 8

	// Const Field (delimiter)
	lengthInBits += 8

	// Simple field (parameterValue)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CBusCommandDeviceManagement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CBusCommandDeviceManagement) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CBusCommand, cBusOptions CBusOptions) (__cBusCommandDeviceManagement CBusCommandDeviceManagement, err error) {
	m.CBusCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusCommandDeviceManagement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusCommandDeviceManagement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	paramNo, err := ReadEnumField[Parameter](ctx, "paramNo", "Parameter", ReadEnum(ParameterByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'paramNo' field"))
	}
	m.ParamNo = paramNo

	delimiter, err := ReadConstField[byte](ctx, "delimiter", ReadByte(readBuffer, 8), CBusCommandDeviceManagement_DELIMITER)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'delimiter' field"))
	}
	_ = delimiter

	parameterValue, err := ReadSimpleField(ctx, "parameterValue", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parameterValue' field"))
	}
	m.ParameterValue = parameterValue

	if closeErr := readBuffer.CloseContext("CBusCommandDeviceManagement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusCommandDeviceManagement")
	}

	return m, nil
}

func (m *_CBusCommandDeviceManagement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CBusCommandDeviceManagement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusCommandDeviceManagement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusCommandDeviceManagement")
		}

		if err := WriteSimpleEnumField[Parameter](ctx, "paramNo", "Parameter", m.GetParamNo(), WriteEnum[Parameter, uint8](Parameter.GetValue, Parameter.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'paramNo' field")
		}

		if err := WriteConstField(ctx, "delimiter", CBusCommandDeviceManagement_DELIMITER, WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'delimiter' field")
		}

		if err := WriteSimpleField[byte](ctx, "parameterValue", m.GetParameterValue(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'parameterValue' field")
		}

		if popErr := writeBuffer.PopContext("CBusCommandDeviceManagement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusCommandDeviceManagement")
		}
		return nil
	}
	return m.CBusCommandContract.(*_CBusCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CBusCommandDeviceManagement) IsCBusCommandDeviceManagement() {}

func (m *_CBusCommandDeviceManagement) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CBusCommandDeviceManagement) deepCopy() *_CBusCommandDeviceManagement {
	if m == nil {
		return nil
	}
	_CBusCommandDeviceManagementCopy := &_CBusCommandDeviceManagement{
		m.CBusCommandContract.(*_CBusCommand).deepCopy(),
		m.ParamNo,
		m.ParameterValue,
	}
	_CBusCommandDeviceManagementCopy.CBusCommandContract.(*_CBusCommand)._SubType = m
	return _CBusCommandDeviceManagementCopy
}

func (m *_CBusCommandDeviceManagement) String() string {
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