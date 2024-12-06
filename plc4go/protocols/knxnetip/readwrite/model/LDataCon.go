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

// LDataCon is the corresponding interface of LDataCon
type LDataCon interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CEMI
	// GetAdditionalInformationLength returns AdditionalInformationLength (property field)
	GetAdditionalInformationLength() uint8
	// GetAdditionalInformation returns AdditionalInformation (property field)
	GetAdditionalInformation() []CEMIAdditionalInformation
	// GetDataFrame returns DataFrame (property field)
	GetDataFrame() LDataFrame
	// IsLDataCon is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLDataCon()
	// CreateBuilder creates a LDataConBuilder
	CreateLDataConBuilder() LDataConBuilder
}

// _LDataCon is the data-structure of this message
type _LDataCon struct {
	CEMIContract
	AdditionalInformationLength uint8
	AdditionalInformation       []CEMIAdditionalInformation
	DataFrame                   LDataFrame
}

var _ LDataCon = (*_LDataCon)(nil)
var _ CEMIRequirements = (*_LDataCon)(nil)

// NewLDataCon factory function for _LDataCon
func NewLDataCon(additionalInformationLength uint8, additionalInformation []CEMIAdditionalInformation, dataFrame LDataFrame, size uint16) *_LDataCon {
	if dataFrame == nil {
		panic("dataFrame of type LDataFrame for LDataCon must not be nil")
	}
	_result := &_LDataCon{
		CEMIContract:                NewCEMI(size),
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation:       additionalInformation,
		DataFrame:                   dataFrame,
	}
	_result.CEMIContract.(*_CEMI)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// LDataConBuilder is a builder for LDataCon
type LDataConBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(additionalInformationLength uint8, additionalInformation []CEMIAdditionalInformation, dataFrame LDataFrame) LDataConBuilder
	// WithAdditionalInformationLength adds AdditionalInformationLength (property field)
	WithAdditionalInformationLength(uint8) LDataConBuilder
	// WithAdditionalInformation adds AdditionalInformation (property field)
	WithAdditionalInformation(...CEMIAdditionalInformation) LDataConBuilder
	// WithDataFrame adds DataFrame (property field)
	WithDataFrame(LDataFrame) LDataConBuilder
	// WithDataFrameBuilder adds DataFrame (property field) which is build by the builder
	WithDataFrameBuilder(func(LDataFrameBuilder) LDataFrameBuilder) LDataConBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CEMIBuilder
	// Build builds the LDataCon or returns an error if something is wrong
	Build() (LDataCon, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() LDataCon
}

// NewLDataConBuilder() creates a LDataConBuilder
func NewLDataConBuilder() LDataConBuilder {
	return &_LDataConBuilder{_LDataCon: new(_LDataCon)}
}

type _LDataConBuilder struct {
	*_LDataCon

	parentBuilder *_CEMIBuilder

	err *utils.MultiError
}

var _ (LDataConBuilder) = (*_LDataConBuilder)(nil)

func (b *_LDataConBuilder) setParent(contract CEMIContract) {
	b.CEMIContract = contract
	contract.(*_CEMI)._SubType = b._LDataCon
}

func (b *_LDataConBuilder) WithMandatoryFields(additionalInformationLength uint8, additionalInformation []CEMIAdditionalInformation, dataFrame LDataFrame) LDataConBuilder {
	return b.WithAdditionalInformationLength(additionalInformationLength).WithAdditionalInformation(additionalInformation...).WithDataFrame(dataFrame)
}

func (b *_LDataConBuilder) WithAdditionalInformationLength(additionalInformationLength uint8) LDataConBuilder {
	b.AdditionalInformationLength = additionalInformationLength
	return b
}

func (b *_LDataConBuilder) WithAdditionalInformation(additionalInformation ...CEMIAdditionalInformation) LDataConBuilder {
	b.AdditionalInformation = additionalInformation
	return b
}

func (b *_LDataConBuilder) WithDataFrame(dataFrame LDataFrame) LDataConBuilder {
	b.DataFrame = dataFrame
	return b
}

func (b *_LDataConBuilder) WithDataFrameBuilder(builderSupplier func(LDataFrameBuilder) LDataFrameBuilder) LDataConBuilder {
	builder := builderSupplier(b.DataFrame.CreateLDataFrameBuilder())
	var err error
	b.DataFrame, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LDataFrameBuilder failed"))
	}
	return b
}

func (b *_LDataConBuilder) Build() (LDataCon, error) {
	if b.DataFrame == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'dataFrame' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._LDataCon.deepCopy(), nil
}

func (b *_LDataConBuilder) MustBuild() LDataCon {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_LDataConBuilder) Done() CEMIBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCEMIBuilder().(*_CEMIBuilder)
	}
	return b.parentBuilder
}

func (b *_LDataConBuilder) buildForCEMI() (CEMI, error) {
	return b.Build()
}

func (b *_LDataConBuilder) DeepCopy() any {
	_copy := b.CreateLDataConBuilder().(*_LDataConBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateLDataConBuilder creates a LDataConBuilder
func (b *_LDataCon) CreateLDataConBuilder() LDataConBuilder {
	if b == nil {
		return NewLDataConBuilder()
	}
	return &_LDataConBuilder{_LDataCon: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LDataCon) GetMessageCode() uint8 {
	return 0x2E
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LDataCon) GetParent() CEMIContract {
	return m.CEMIContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LDataCon) GetAdditionalInformationLength() uint8 {
	return m.AdditionalInformationLength
}

func (m *_LDataCon) GetAdditionalInformation() []CEMIAdditionalInformation {
	return m.AdditionalInformation
}

func (m *_LDataCon) GetDataFrame() LDataFrame {
	return m.DataFrame
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastLDataCon(structType any) LDataCon {
	if casted, ok := structType.(LDataCon); ok {
		return casted
	}
	if casted, ok := structType.(*LDataCon); ok {
		return *casted
	}
	return nil
}

func (m *_LDataCon) GetTypeName() string {
	return "LDataCon"
}

func (m *_LDataCon) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	// Simple field (additionalInformationLength)
	lengthInBits += 8

	// Array field
	if len(m.AdditionalInformation) > 0 {
		for _, element := range m.AdditionalInformation {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (dataFrame)
	lengthInBits += m.DataFrame.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_LDataCon) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LDataCon) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__lDataCon LDataCon, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LDataCon"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LDataCon")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	additionalInformationLength, err := ReadSimpleField(ctx, "additionalInformationLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'additionalInformationLength' field"))
	}
	m.AdditionalInformationLength = additionalInformationLength

	additionalInformation, err := ReadLengthArrayField[CEMIAdditionalInformation](ctx, "additionalInformation", ReadComplex[CEMIAdditionalInformation](CEMIAdditionalInformationParseWithBuffer, readBuffer), int(additionalInformationLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'additionalInformation' field"))
	}
	m.AdditionalInformation = additionalInformation

	dataFrame, err := ReadSimpleField[LDataFrame](ctx, "dataFrame", ReadComplex[LDataFrame](LDataFrameParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataFrame' field"))
	}
	m.DataFrame = dataFrame

	if closeErr := readBuffer.CloseContext("LDataCon"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LDataCon")
	}

	return m, nil
}

func (m *_LDataCon) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LDataCon) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LDataCon"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LDataCon")
		}

		if err := WriteSimpleField[uint8](ctx, "additionalInformationLength", m.GetAdditionalInformationLength(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'additionalInformationLength' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "additionalInformation", m.GetAdditionalInformation(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'additionalInformation' field")
		}

		if err := WriteSimpleField[LDataFrame](ctx, "dataFrame", m.GetDataFrame(), WriteComplex[LDataFrame](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataFrame' field")
		}

		if popErr := writeBuffer.PopContext("LDataCon"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LDataCon")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LDataCon) IsLDataCon() {}

func (m *_LDataCon) DeepCopy() any {
	return m.deepCopy()
}

func (m *_LDataCon) deepCopy() *_LDataCon {
	if m == nil {
		return nil
	}
	_LDataConCopy := &_LDataCon{
		m.CEMIContract.(*_CEMI).deepCopy(),
		m.AdditionalInformationLength,
		utils.DeepCopySlice[CEMIAdditionalInformation, CEMIAdditionalInformation](m.AdditionalInformation),
		utils.DeepCopy[LDataFrame](m.DataFrame),
	}
	_LDataConCopy.CEMIContract.(*_CEMI)._SubType = m
	return _LDataConCopy
}

func (m *_LDataCon) String() string {
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