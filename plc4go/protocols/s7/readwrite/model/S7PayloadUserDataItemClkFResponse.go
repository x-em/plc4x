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

// S7PayloadUserDataItemClkFResponse is the corresponding interface of S7PayloadUserDataItemClkFResponse
type S7PayloadUserDataItemClkFResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7PayloadUserDataItem
	// GetRes returns Res (property field)
	GetRes() uint8
	// GetYear1 returns Year1 (property field)
	GetYear1() uint8
	// GetTimeStamp returns TimeStamp (property field)
	GetTimeStamp() DateAndTime
	// IsS7PayloadUserDataItemClkFResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadUserDataItemClkFResponse()
	// CreateBuilder creates a S7PayloadUserDataItemClkFResponseBuilder
	CreateS7PayloadUserDataItemClkFResponseBuilder() S7PayloadUserDataItemClkFResponseBuilder
}

// _S7PayloadUserDataItemClkFResponse is the data-structure of this message
type _S7PayloadUserDataItemClkFResponse struct {
	S7PayloadUserDataItemContract
	Res       uint8
	Year1     uint8
	TimeStamp DateAndTime
}

var _ S7PayloadUserDataItemClkFResponse = (*_S7PayloadUserDataItemClkFResponse)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadUserDataItemClkFResponse)(nil)

// NewS7PayloadUserDataItemClkFResponse factory function for _S7PayloadUserDataItemClkFResponse
func NewS7PayloadUserDataItemClkFResponse(returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16, res uint8, year1 uint8, timeStamp DateAndTime) *_S7PayloadUserDataItemClkFResponse {
	if timeStamp == nil {
		panic("timeStamp of type DateAndTime for S7PayloadUserDataItemClkFResponse must not be nil")
	}
	_result := &_S7PayloadUserDataItemClkFResponse{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
		Res:                           res,
		Year1:                         year1,
		TimeStamp:                     timeStamp,
	}
	_result.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7PayloadUserDataItemClkFResponseBuilder is a builder for S7PayloadUserDataItemClkFResponse
type S7PayloadUserDataItemClkFResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(res uint8, year1 uint8, timeStamp DateAndTime) S7PayloadUserDataItemClkFResponseBuilder
	// WithRes adds Res (property field)
	WithRes(uint8) S7PayloadUserDataItemClkFResponseBuilder
	// WithYear1 adds Year1 (property field)
	WithYear1(uint8) S7PayloadUserDataItemClkFResponseBuilder
	// WithTimeStamp adds TimeStamp (property field)
	WithTimeStamp(DateAndTime) S7PayloadUserDataItemClkFResponseBuilder
	// WithTimeStampBuilder adds TimeStamp (property field) which is build by the builder
	WithTimeStampBuilder(func(DateAndTimeBuilder) DateAndTimeBuilder) S7PayloadUserDataItemClkFResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() S7PayloadUserDataItemBuilder
	// Build builds the S7PayloadUserDataItemClkFResponse or returns an error if something is wrong
	Build() (S7PayloadUserDataItemClkFResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7PayloadUserDataItemClkFResponse
}

// NewS7PayloadUserDataItemClkFResponseBuilder() creates a S7PayloadUserDataItemClkFResponseBuilder
func NewS7PayloadUserDataItemClkFResponseBuilder() S7PayloadUserDataItemClkFResponseBuilder {
	return &_S7PayloadUserDataItemClkFResponseBuilder{_S7PayloadUserDataItemClkFResponse: new(_S7PayloadUserDataItemClkFResponse)}
}

type _S7PayloadUserDataItemClkFResponseBuilder struct {
	*_S7PayloadUserDataItemClkFResponse

	parentBuilder *_S7PayloadUserDataItemBuilder

	err *utils.MultiError
}

var _ (S7PayloadUserDataItemClkFResponseBuilder) = (*_S7PayloadUserDataItemClkFResponseBuilder)(nil)

func (b *_S7PayloadUserDataItemClkFResponseBuilder) setParent(contract S7PayloadUserDataItemContract) {
	b.S7PayloadUserDataItemContract = contract
	contract.(*_S7PayloadUserDataItem)._SubType = b._S7PayloadUserDataItemClkFResponse
}

func (b *_S7PayloadUserDataItemClkFResponseBuilder) WithMandatoryFields(res uint8, year1 uint8, timeStamp DateAndTime) S7PayloadUserDataItemClkFResponseBuilder {
	return b.WithRes(res).WithYear1(year1).WithTimeStamp(timeStamp)
}

func (b *_S7PayloadUserDataItemClkFResponseBuilder) WithRes(res uint8) S7PayloadUserDataItemClkFResponseBuilder {
	b.Res = res
	return b
}

func (b *_S7PayloadUserDataItemClkFResponseBuilder) WithYear1(year1 uint8) S7PayloadUserDataItemClkFResponseBuilder {
	b.Year1 = year1
	return b
}

func (b *_S7PayloadUserDataItemClkFResponseBuilder) WithTimeStamp(timeStamp DateAndTime) S7PayloadUserDataItemClkFResponseBuilder {
	b.TimeStamp = timeStamp
	return b
}

func (b *_S7PayloadUserDataItemClkFResponseBuilder) WithTimeStampBuilder(builderSupplier func(DateAndTimeBuilder) DateAndTimeBuilder) S7PayloadUserDataItemClkFResponseBuilder {
	builder := builderSupplier(b.TimeStamp.CreateDateAndTimeBuilder())
	var err error
	b.TimeStamp, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "DateAndTimeBuilder failed"))
	}
	return b
}

func (b *_S7PayloadUserDataItemClkFResponseBuilder) Build() (S7PayloadUserDataItemClkFResponse, error) {
	if b.TimeStamp == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeStamp' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7PayloadUserDataItemClkFResponse.deepCopy(), nil
}

func (b *_S7PayloadUserDataItemClkFResponseBuilder) MustBuild() S7PayloadUserDataItemClkFResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_S7PayloadUserDataItemClkFResponseBuilder) Done() S7PayloadUserDataItemBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewS7PayloadUserDataItemBuilder().(*_S7PayloadUserDataItemBuilder)
	}
	return b.parentBuilder
}

func (b *_S7PayloadUserDataItemClkFResponseBuilder) buildForS7PayloadUserDataItem() (S7PayloadUserDataItem, error) {
	return b.Build()
}

func (b *_S7PayloadUserDataItemClkFResponseBuilder) DeepCopy() any {
	_copy := b.CreateS7PayloadUserDataItemClkFResponseBuilder().(*_S7PayloadUserDataItemClkFResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7PayloadUserDataItemClkFResponseBuilder creates a S7PayloadUserDataItemClkFResponseBuilder
func (b *_S7PayloadUserDataItemClkFResponse) CreateS7PayloadUserDataItemClkFResponseBuilder() S7PayloadUserDataItemClkFResponseBuilder {
	if b == nil {
		return NewS7PayloadUserDataItemClkFResponseBuilder()
	}
	return &_S7PayloadUserDataItemClkFResponseBuilder{_S7PayloadUserDataItemClkFResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemClkFResponse) GetCpuFunctionGroup() uint8 {
	return 0x07
}

func (m *_S7PayloadUserDataItemClkFResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *_S7PayloadUserDataItemClkFResponse) GetCpuSubfunction() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemClkFResponse) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemClkFResponse) GetRes() uint8 {
	return m.Res
}

func (m *_S7PayloadUserDataItemClkFResponse) GetYear1() uint8 {
	return m.Year1
}

func (m *_S7PayloadUserDataItemClkFResponse) GetTimeStamp() DateAndTime {
	return m.TimeStamp
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemClkFResponse(structType any) S7PayloadUserDataItemClkFResponse {
	if casted, ok := structType.(S7PayloadUserDataItemClkFResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemClkFResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemClkFResponse) GetTypeName() string {
	return "S7PayloadUserDataItemClkFResponse"
}

func (m *_S7PayloadUserDataItemClkFResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	// Simple field (res)
	lengthInBits += 8

	// Simple field (year1)
	lengthInBits += 8

	// Simple field (timeStamp)
	lengthInBits += m.TimeStamp.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_S7PayloadUserDataItemClkFResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserDataItemClkFResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, dataLength uint16, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadUserDataItemClkFResponse S7PayloadUserDataItemClkFResponse, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemClkFResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemClkFResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	res, err := ReadSimpleField(ctx, "res", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'res' field"))
	}
	m.Res = res

	year1, err := ReadSimpleField(ctx, "year1", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'year1' field"))
	}
	m.Year1 = year1

	timeStamp, err := ReadSimpleField[DateAndTime](ctx, "timeStamp", ReadComplex[DateAndTime](DateAndTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeStamp' field"))
	}
	m.TimeStamp = timeStamp

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemClkFResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemClkFResponse")
	}

	return m, nil
}

func (m *_S7PayloadUserDataItemClkFResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemClkFResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemClkFResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemClkFResponse")
		}

		if err := WriteSimpleField[uint8](ctx, "res", m.GetRes(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'res' field")
		}

		if err := WriteSimpleField[uint8](ctx, "year1", m.GetYear1(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'year1' field")
		}

		if err := WriteSimpleField[DateAndTime](ctx, "timeStamp", m.GetTimeStamp(), WriteComplex[DateAndTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeStamp' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemClkFResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemClkFResponse")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemClkFResponse) IsS7PayloadUserDataItemClkFResponse() {}

func (m *_S7PayloadUserDataItemClkFResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7PayloadUserDataItemClkFResponse) deepCopy() *_S7PayloadUserDataItemClkFResponse {
	if m == nil {
		return nil
	}
	_S7PayloadUserDataItemClkFResponseCopy := &_S7PayloadUserDataItemClkFResponse{
		m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).deepCopy(),
		m.Res,
		m.Year1,
		utils.DeepCopy[DateAndTime](m.TimeStamp),
	}
	_S7PayloadUserDataItemClkFResponseCopy.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = m
	return _S7PayloadUserDataItemClkFResponseCopy
}

func (m *_S7PayloadUserDataItemClkFResponse) String() string {
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