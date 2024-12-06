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

// BACnetServiceAckVTData is the corresponding interface of BACnetServiceAckVTData
type BACnetServiceAckVTData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetServiceAck
	// GetVtSessionIdentifier returns VtSessionIdentifier (property field)
	GetVtSessionIdentifier() BACnetApplicationTagUnsignedInteger
	// GetVtNewData returns VtNewData (property field)
	GetVtNewData() BACnetApplicationTagOctetString
	// GetVtDataFlag returns VtDataFlag (property field)
	GetVtDataFlag() BACnetApplicationTagUnsignedInteger
	// IsBACnetServiceAckVTData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetServiceAckVTData()
	// CreateBuilder creates a BACnetServiceAckVTDataBuilder
	CreateBACnetServiceAckVTDataBuilder() BACnetServiceAckVTDataBuilder
}

// _BACnetServiceAckVTData is the data-structure of this message
type _BACnetServiceAckVTData struct {
	BACnetServiceAckContract
	VtSessionIdentifier BACnetApplicationTagUnsignedInteger
	VtNewData           BACnetApplicationTagOctetString
	VtDataFlag          BACnetApplicationTagUnsignedInteger
}

var _ BACnetServiceAckVTData = (*_BACnetServiceAckVTData)(nil)
var _ BACnetServiceAckRequirements = (*_BACnetServiceAckVTData)(nil)

// NewBACnetServiceAckVTData factory function for _BACnetServiceAckVTData
func NewBACnetServiceAckVTData(vtSessionIdentifier BACnetApplicationTagUnsignedInteger, vtNewData BACnetApplicationTagOctetString, vtDataFlag BACnetApplicationTagUnsignedInteger, serviceAckLength uint32) *_BACnetServiceAckVTData {
	if vtSessionIdentifier == nil {
		panic("vtSessionIdentifier of type BACnetApplicationTagUnsignedInteger for BACnetServiceAckVTData must not be nil")
	}
	if vtNewData == nil {
		panic("vtNewData of type BACnetApplicationTagOctetString for BACnetServiceAckVTData must not be nil")
	}
	if vtDataFlag == nil {
		panic("vtDataFlag of type BACnetApplicationTagUnsignedInteger for BACnetServiceAckVTData must not be nil")
	}
	_result := &_BACnetServiceAckVTData{
		BACnetServiceAckContract: NewBACnetServiceAck(serviceAckLength),
		VtSessionIdentifier:      vtSessionIdentifier,
		VtNewData:                vtNewData,
		VtDataFlag:               vtDataFlag,
	}
	_result.BACnetServiceAckContract.(*_BACnetServiceAck)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetServiceAckVTDataBuilder is a builder for BACnetServiceAckVTData
type BACnetServiceAckVTDataBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(vtSessionIdentifier BACnetApplicationTagUnsignedInteger, vtNewData BACnetApplicationTagOctetString, vtDataFlag BACnetApplicationTagUnsignedInteger) BACnetServiceAckVTDataBuilder
	// WithVtSessionIdentifier adds VtSessionIdentifier (property field)
	WithVtSessionIdentifier(BACnetApplicationTagUnsignedInteger) BACnetServiceAckVTDataBuilder
	// WithVtSessionIdentifierBuilder adds VtSessionIdentifier (property field) which is build by the builder
	WithVtSessionIdentifierBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetServiceAckVTDataBuilder
	// WithVtNewData adds VtNewData (property field)
	WithVtNewData(BACnetApplicationTagOctetString) BACnetServiceAckVTDataBuilder
	// WithVtNewDataBuilder adds VtNewData (property field) which is build by the builder
	WithVtNewDataBuilder(func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetServiceAckVTDataBuilder
	// WithVtDataFlag adds VtDataFlag (property field)
	WithVtDataFlag(BACnetApplicationTagUnsignedInteger) BACnetServiceAckVTDataBuilder
	// WithVtDataFlagBuilder adds VtDataFlag (property field) which is build by the builder
	WithVtDataFlagBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetServiceAckVTDataBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetServiceAckBuilder
	// Build builds the BACnetServiceAckVTData or returns an error if something is wrong
	Build() (BACnetServiceAckVTData, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetServiceAckVTData
}

// NewBACnetServiceAckVTDataBuilder() creates a BACnetServiceAckVTDataBuilder
func NewBACnetServiceAckVTDataBuilder() BACnetServiceAckVTDataBuilder {
	return &_BACnetServiceAckVTDataBuilder{_BACnetServiceAckVTData: new(_BACnetServiceAckVTData)}
}

type _BACnetServiceAckVTDataBuilder struct {
	*_BACnetServiceAckVTData

	parentBuilder *_BACnetServiceAckBuilder

	err *utils.MultiError
}

var _ (BACnetServiceAckVTDataBuilder) = (*_BACnetServiceAckVTDataBuilder)(nil)

func (b *_BACnetServiceAckVTDataBuilder) setParent(contract BACnetServiceAckContract) {
	b.BACnetServiceAckContract = contract
	contract.(*_BACnetServiceAck)._SubType = b._BACnetServiceAckVTData
}

func (b *_BACnetServiceAckVTDataBuilder) WithMandatoryFields(vtSessionIdentifier BACnetApplicationTagUnsignedInteger, vtNewData BACnetApplicationTagOctetString, vtDataFlag BACnetApplicationTagUnsignedInteger) BACnetServiceAckVTDataBuilder {
	return b.WithVtSessionIdentifier(vtSessionIdentifier).WithVtNewData(vtNewData).WithVtDataFlag(vtDataFlag)
}

func (b *_BACnetServiceAckVTDataBuilder) WithVtSessionIdentifier(vtSessionIdentifier BACnetApplicationTagUnsignedInteger) BACnetServiceAckVTDataBuilder {
	b.VtSessionIdentifier = vtSessionIdentifier
	return b
}

func (b *_BACnetServiceAckVTDataBuilder) WithVtSessionIdentifierBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetServiceAckVTDataBuilder {
	builder := builderSupplier(b.VtSessionIdentifier.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.VtSessionIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetServiceAckVTDataBuilder) WithVtNewData(vtNewData BACnetApplicationTagOctetString) BACnetServiceAckVTDataBuilder {
	b.VtNewData = vtNewData
	return b
}

func (b *_BACnetServiceAckVTDataBuilder) WithVtNewDataBuilder(builderSupplier func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetServiceAckVTDataBuilder {
	builder := builderSupplier(b.VtNewData.CreateBACnetApplicationTagOctetStringBuilder())
	var err error
	b.VtNewData, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagOctetStringBuilder failed"))
	}
	return b
}

func (b *_BACnetServiceAckVTDataBuilder) WithVtDataFlag(vtDataFlag BACnetApplicationTagUnsignedInteger) BACnetServiceAckVTDataBuilder {
	b.VtDataFlag = vtDataFlag
	return b
}

func (b *_BACnetServiceAckVTDataBuilder) WithVtDataFlagBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetServiceAckVTDataBuilder {
	builder := builderSupplier(b.VtDataFlag.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.VtDataFlag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetServiceAckVTDataBuilder) Build() (BACnetServiceAckVTData, error) {
	if b.VtSessionIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'vtSessionIdentifier' not set"))
	}
	if b.VtNewData == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'vtNewData' not set"))
	}
	if b.VtDataFlag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'vtDataFlag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetServiceAckVTData.deepCopy(), nil
}

func (b *_BACnetServiceAckVTDataBuilder) MustBuild() BACnetServiceAckVTData {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetServiceAckVTDataBuilder) Done() BACnetServiceAckBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetServiceAckBuilder().(*_BACnetServiceAckBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetServiceAckVTDataBuilder) buildForBACnetServiceAck() (BACnetServiceAck, error) {
	return b.Build()
}

func (b *_BACnetServiceAckVTDataBuilder) DeepCopy() any {
	_copy := b.CreateBACnetServiceAckVTDataBuilder().(*_BACnetServiceAckVTDataBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetServiceAckVTDataBuilder creates a BACnetServiceAckVTDataBuilder
func (b *_BACnetServiceAckVTData) CreateBACnetServiceAckVTDataBuilder() BACnetServiceAckVTDataBuilder {
	if b == nil {
		return NewBACnetServiceAckVTDataBuilder()
	}
	return &_BACnetServiceAckVTDataBuilder{_BACnetServiceAckVTData: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckVTData) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_VT_DATA
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckVTData) GetParent() BACnetServiceAckContract {
	return m.BACnetServiceAckContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckVTData) GetVtSessionIdentifier() BACnetApplicationTagUnsignedInteger {
	return m.VtSessionIdentifier
}

func (m *_BACnetServiceAckVTData) GetVtNewData() BACnetApplicationTagOctetString {
	return m.VtNewData
}

func (m *_BACnetServiceAckVTData) GetVtDataFlag() BACnetApplicationTagUnsignedInteger {
	return m.VtDataFlag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckVTData(structType any) BACnetServiceAckVTData {
	if casted, ok := structType.(BACnetServiceAckVTData); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckVTData); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckVTData) GetTypeName() string {
	return "BACnetServiceAckVTData"
}

func (m *_BACnetServiceAckVTData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetServiceAckContract.(*_BACnetServiceAck).getLengthInBits(ctx))

	// Simple field (vtSessionIdentifier)
	lengthInBits += m.VtSessionIdentifier.GetLengthInBits(ctx)

	// Simple field (vtNewData)
	lengthInBits += m.VtNewData.GetLengthInBits(ctx)

	// Simple field (vtDataFlag)
	lengthInBits += m.VtDataFlag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetServiceAckVTData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetServiceAckVTData) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetServiceAck, serviceAckLength uint32) (__bACnetServiceAckVTData BACnetServiceAckVTData, err error) {
	m.BACnetServiceAckContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckVTData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckVTData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	vtSessionIdentifier, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "vtSessionIdentifier", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vtSessionIdentifier' field"))
	}
	m.VtSessionIdentifier = vtSessionIdentifier

	vtNewData, err := ReadSimpleField[BACnetApplicationTagOctetString](ctx, "vtNewData", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vtNewData' field"))
	}
	m.VtNewData = vtNewData

	vtDataFlag, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "vtDataFlag", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vtDataFlag' field"))
	}
	m.VtDataFlag = vtDataFlag

	if closeErr := readBuffer.CloseContext("BACnetServiceAckVTData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckVTData")
	}

	return m, nil
}

func (m *_BACnetServiceAckVTData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckVTData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckVTData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckVTData")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "vtSessionIdentifier", m.GetVtSessionIdentifier(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'vtSessionIdentifier' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagOctetString](ctx, "vtNewData", m.GetVtNewData(), WriteComplex[BACnetApplicationTagOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'vtNewData' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "vtDataFlag", m.GetVtDataFlag(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'vtDataFlag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckVTData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckVTData")
		}
		return nil
	}
	return m.BACnetServiceAckContract.(*_BACnetServiceAck).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetServiceAckVTData) IsBACnetServiceAckVTData() {}

func (m *_BACnetServiceAckVTData) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetServiceAckVTData) deepCopy() *_BACnetServiceAckVTData {
	if m == nil {
		return nil
	}
	_BACnetServiceAckVTDataCopy := &_BACnetServiceAckVTData{
		m.BACnetServiceAckContract.(*_BACnetServiceAck).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.VtSessionIdentifier),
		utils.DeepCopy[BACnetApplicationTagOctetString](m.VtNewData),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.VtDataFlag),
	}
	_BACnetServiceAckVTDataCopy.BACnetServiceAckContract.(*_BACnetServiceAck)._SubType = m
	return _BACnetServiceAckVTDataCopy
}

func (m *_BACnetServiceAckVTData) String() string {
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