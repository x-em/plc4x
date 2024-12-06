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
const S7MessageObjectRequest_VARIABLESPEC uint8 = 0x12
const S7MessageObjectRequest_LENGTH uint8 = 0x08

// S7MessageObjectRequest is the corresponding interface of S7MessageObjectRequest
type S7MessageObjectRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7DataAlarmMessage
	// GetSyntaxId returns SyntaxId (property field)
	GetSyntaxId() SyntaxIdType
	// GetQueryType returns QueryType (property field)
	GetQueryType() QueryType
	// GetAlarmType returns AlarmType (property field)
	GetAlarmType() AlarmType
	// IsS7MessageObjectRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7MessageObjectRequest()
	// CreateBuilder creates a S7MessageObjectRequestBuilder
	CreateS7MessageObjectRequestBuilder() S7MessageObjectRequestBuilder
}

// _S7MessageObjectRequest is the data-structure of this message
type _S7MessageObjectRequest struct {
	S7DataAlarmMessageContract
	SyntaxId  SyntaxIdType
	QueryType QueryType
	AlarmType AlarmType
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *uint8
}

var _ S7MessageObjectRequest = (*_S7MessageObjectRequest)(nil)
var _ S7DataAlarmMessageRequirements = (*_S7MessageObjectRequest)(nil)

// NewS7MessageObjectRequest factory function for _S7MessageObjectRequest
func NewS7MessageObjectRequest(syntaxId SyntaxIdType, queryType QueryType, alarmType AlarmType) *_S7MessageObjectRequest {
	_result := &_S7MessageObjectRequest{
		S7DataAlarmMessageContract: NewS7DataAlarmMessage(),
		SyntaxId:                   syntaxId,
		QueryType:                  queryType,
		AlarmType:                  alarmType,
	}
	_result.S7DataAlarmMessageContract.(*_S7DataAlarmMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7MessageObjectRequestBuilder is a builder for S7MessageObjectRequest
type S7MessageObjectRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(syntaxId SyntaxIdType, queryType QueryType, alarmType AlarmType) S7MessageObjectRequestBuilder
	// WithSyntaxId adds SyntaxId (property field)
	WithSyntaxId(SyntaxIdType) S7MessageObjectRequestBuilder
	// WithQueryType adds QueryType (property field)
	WithQueryType(QueryType) S7MessageObjectRequestBuilder
	// WithAlarmType adds AlarmType (property field)
	WithAlarmType(AlarmType) S7MessageObjectRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() S7DataAlarmMessageBuilder
	// Build builds the S7MessageObjectRequest or returns an error if something is wrong
	Build() (S7MessageObjectRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7MessageObjectRequest
}

// NewS7MessageObjectRequestBuilder() creates a S7MessageObjectRequestBuilder
func NewS7MessageObjectRequestBuilder() S7MessageObjectRequestBuilder {
	return &_S7MessageObjectRequestBuilder{_S7MessageObjectRequest: new(_S7MessageObjectRequest)}
}

type _S7MessageObjectRequestBuilder struct {
	*_S7MessageObjectRequest

	parentBuilder *_S7DataAlarmMessageBuilder

	err *utils.MultiError
}

var _ (S7MessageObjectRequestBuilder) = (*_S7MessageObjectRequestBuilder)(nil)

func (b *_S7MessageObjectRequestBuilder) setParent(contract S7DataAlarmMessageContract) {
	b.S7DataAlarmMessageContract = contract
	contract.(*_S7DataAlarmMessage)._SubType = b._S7MessageObjectRequest
}

func (b *_S7MessageObjectRequestBuilder) WithMandatoryFields(syntaxId SyntaxIdType, queryType QueryType, alarmType AlarmType) S7MessageObjectRequestBuilder {
	return b.WithSyntaxId(syntaxId).WithQueryType(queryType).WithAlarmType(alarmType)
}

func (b *_S7MessageObjectRequestBuilder) WithSyntaxId(syntaxId SyntaxIdType) S7MessageObjectRequestBuilder {
	b.SyntaxId = syntaxId
	return b
}

func (b *_S7MessageObjectRequestBuilder) WithQueryType(queryType QueryType) S7MessageObjectRequestBuilder {
	b.QueryType = queryType
	return b
}

func (b *_S7MessageObjectRequestBuilder) WithAlarmType(alarmType AlarmType) S7MessageObjectRequestBuilder {
	b.AlarmType = alarmType
	return b
}

func (b *_S7MessageObjectRequestBuilder) Build() (S7MessageObjectRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7MessageObjectRequest.deepCopy(), nil
}

func (b *_S7MessageObjectRequestBuilder) MustBuild() S7MessageObjectRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_S7MessageObjectRequestBuilder) Done() S7DataAlarmMessageBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewS7DataAlarmMessageBuilder().(*_S7DataAlarmMessageBuilder)
	}
	return b.parentBuilder
}

func (b *_S7MessageObjectRequestBuilder) buildForS7DataAlarmMessage() (S7DataAlarmMessage, error) {
	return b.Build()
}

func (b *_S7MessageObjectRequestBuilder) DeepCopy() any {
	_copy := b.CreateS7MessageObjectRequestBuilder().(*_S7MessageObjectRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7MessageObjectRequestBuilder creates a S7MessageObjectRequestBuilder
func (b *_S7MessageObjectRequest) CreateS7MessageObjectRequestBuilder() S7MessageObjectRequestBuilder {
	if b == nil {
		return NewS7MessageObjectRequestBuilder()
	}
	return &_S7MessageObjectRequestBuilder{_S7MessageObjectRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7MessageObjectRequest) GetCpuFunctionType() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7MessageObjectRequest) GetParent() S7DataAlarmMessageContract {
	return m.S7DataAlarmMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7MessageObjectRequest) GetSyntaxId() SyntaxIdType {
	return m.SyntaxId
}

func (m *_S7MessageObjectRequest) GetQueryType() QueryType {
	return m.QueryType
}

func (m *_S7MessageObjectRequest) GetAlarmType() AlarmType {
	return m.AlarmType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_S7MessageObjectRequest) GetVariableSpec() uint8 {
	return S7MessageObjectRequest_VARIABLESPEC
}

func (m *_S7MessageObjectRequest) GetLength() uint8 {
	return S7MessageObjectRequest_LENGTH
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7MessageObjectRequest(structType any) S7MessageObjectRequest {
	if casted, ok := structType.(S7MessageObjectRequest); ok {
		return casted
	}
	if casted, ok := structType.(*S7MessageObjectRequest); ok {
		return *casted
	}
	return nil
}

func (m *_S7MessageObjectRequest) GetTypeName() string {
	return "S7MessageObjectRequest"
}

func (m *_S7MessageObjectRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7DataAlarmMessageContract.(*_S7DataAlarmMessage).getLengthInBits(ctx))

	// Const Field (variableSpec)
	lengthInBits += 8

	// Const Field (length)
	lengthInBits += 8

	// Simple field (syntaxId)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (queryType)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (alarmType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7MessageObjectRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7MessageObjectRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7DataAlarmMessage, cpuFunctionType uint8) (__s7MessageObjectRequest S7MessageObjectRequest, err error) {
	m.S7DataAlarmMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7MessageObjectRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7MessageObjectRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	variableSpec, err := ReadConstField[uint8](ctx, "variableSpec", ReadUnsignedByte(readBuffer, uint8(8)), S7MessageObjectRequest_VARIABLESPEC)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'variableSpec' field"))
	}
	_ = variableSpec

	length, err := ReadConstField[uint8](ctx, "length", ReadUnsignedByte(readBuffer, uint8(8)), S7MessageObjectRequest_LENGTH)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'length' field"))
	}
	_ = length

	syntaxId, err := ReadEnumField[SyntaxIdType](ctx, "syntaxId", "SyntaxIdType", ReadEnum(SyntaxIdTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'syntaxId' field"))
	}
	m.SyntaxId = syntaxId

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	queryType, err := ReadEnumField[QueryType](ctx, "queryType", "QueryType", ReadEnum(QueryTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'queryType' field"))
	}
	m.QueryType = queryType

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x34))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField1 = reservedField1

	alarmType, err := ReadEnumField[AlarmType](ctx, "alarmType", "AlarmType", ReadEnum(AlarmTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alarmType' field"))
	}
	m.AlarmType = alarmType

	if closeErr := readBuffer.CloseContext("S7MessageObjectRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7MessageObjectRequest")
	}

	return m, nil
}

func (m *_S7MessageObjectRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7MessageObjectRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7MessageObjectRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7MessageObjectRequest")
		}

		if err := WriteConstField(ctx, "variableSpec", S7MessageObjectRequest_VARIABLESPEC, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'variableSpec' field")
		}

		if err := WriteConstField(ctx, "length", S7MessageObjectRequest_LENGTH, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'length' field")
		}

		if err := WriteSimpleEnumField[SyntaxIdType](ctx, "syntaxId", "SyntaxIdType", m.GetSyntaxId(), WriteEnum[SyntaxIdType, uint8](SyntaxIdType.GetValue, SyntaxIdType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'syntaxId' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleEnumField[QueryType](ctx, "queryType", "QueryType", m.GetQueryType(), WriteEnum[QueryType, uint8](QueryType.GetValue, QueryType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'queryType' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x34), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 2")
		}

		if err := WriteSimpleEnumField[AlarmType](ctx, "alarmType", "AlarmType", m.GetAlarmType(), WriteEnum[AlarmType, uint8](AlarmType.GetValue, AlarmType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'alarmType' field")
		}

		if popErr := writeBuffer.PopContext("S7MessageObjectRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7MessageObjectRequest")
		}
		return nil
	}
	return m.S7DataAlarmMessageContract.(*_S7DataAlarmMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7MessageObjectRequest) IsS7MessageObjectRequest() {}

func (m *_S7MessageObjectRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7MessageObjectRequest) deepCopy() *_S7MessageObjectRequest {
	if m == nil {
		return nil
	}
	_S7MessageObjectRequestCopy := &_S7MessageObjectRequest{
		m.S7DataAlarmMessageContract.(*_S7DataAlarmMessage).deepCopy(),
		m.SyntaxId,
		m.QueryType,
		m.AlarmType,
		m.reservedField0,
		m.reservedField1,
	}
	_S7MessageObjectRequestCopy.S7DataAlarmMessageContract.(*_S7DataAlarmMessage)._SubType = m
	return _S7MessageObjectRequestCopy
}

func (m *_S7MessageObjectRequest) String() string {
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