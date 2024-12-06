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

// BACnetConstructedDataStrikeCount is the corresponding interface of BACnetConstructedDataStrikeCount
type BACnetConstructedDataStrikeCount interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetStrikeCount returns StrikeCount (property field)
	GetStrikeCount() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataStrikeCount is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataStrikeCount()
	// CreateBuilder creates a BACnetConstructedDataStrikeCountBuilder
	CreateBACnetConstructedDataStrikeCountBuilder() BACnetConstructedDataStrikeCountBuilder
}

// _BACnetConstructedDataStrikeCount is the data-structure of this message
type _BACnetConstructedDataStrikeCount struct {
	BACnetConstructedDataContract
	StrikeCount BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataStrikeCount = (*_BACnetConstructedDataStrikeCount)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataStrikeCount)(nil)

// NewBACnetConstructedDataStrikeCount factory function for _BACnetConstructedDataStrikeCount
func NewBACnetConstructedDataStrikeCount(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, strikeCount BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataStrikeCount {
	if strikeCount == nil {
		panic("strikeCount of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataStrikeCount must not be nil")
	}
	_result := &_BACnetConstructedDataStrikeCount{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		StrikeCount:                   strikeCount,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataStrikeCountBuilder is a builder for BACnetConstructedDataStrikeCount
type BACnetConstructedDataStrikeCountBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(strikeCount BACnetApplicationTagUnsignedInteger) BACnetConstructedDataStrikeCountBuilder
	// WithStrikeCount adds StrikeCount (property field)
	WithStrikeCount(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataStrikeCountBuilder
	// WithStrikeCountBuilder adds StrikeCount (property field) which is build by the builder
	WithStrikeCountBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataStrikeCountBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataStrikeCount or returns an error if something is wrong
	Build() (BACnetConstructedDataStrikeCount, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataStrikeCount
}

// NewBACnetConstructedDataStrikeCountBuilder() creates a BACnetConstructedDataStrikeCountBuilder
func NewBACnetConstructedDataStrikeCountBuilder() BACnetConstructedDataStrikeCountBuilder {
	return &_BACnetConstructedDataStrikeCountBuilder{_BACnetConstructedDataStrikeCount: new(_BACnetConstructedDataStrikeCount)}
}

type _BACnetConstructedDataStrikeCountBuilder struct {
	*_BACnetConstructedDataStrikeCount

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataStrikeCountBuilder) = (*_BACnetConstructedDataStrikeCountBuilder)(nil)

func (b *_BACnetConstructedDataStrikeCountBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataStrikeCount
}

func (b *_BACnetConstructedDataStrikeCountBuilder) WithMandatoryFields(strikeCount BACnetApplicationTagUnsignedInteger) BACnetConstructedDataStrikeCountBuilder {
	return b.WithStrikeCount(strikeCount)
}

func (b *_BACnetConstructedDataStrikeCountBuilder) WithStrikeCount(strikeCount BACnetApplicationTagUnsignedInteger) BACnetConstructedDataStrikeCountBuilder {
	b.StrikeCount = strikeCount
	return b
}

func (b *_BACnetConstructedDataStrikeCountBuilder) WithStrikeCountBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataStrikeCountBuilder {
	builder := builderSupplier(b.StrikeCount.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.StrikeCount, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataStrikeCountBuilder) Build() (BACnetConstructedDataStrikeCount, error) {
	if b.StrikeCount == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'strikeCount' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataStrikeCount.deepCopy(), nil
}

func (b *_BACnetConstructedDataStrikeCountBuilder) MustBuild() BACnetConstructedDataStrikeCount {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataStrikeCountBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataStrikeCountBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataStrikeCountBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataStrikeCountBuilder().(*_BACnetConstructedDataStrikeCountBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataStrikeCountBuilder creates a BACnetConstructedDataStrikeCountBuilder
func (b *_BACnetConstructedDataStrikeCount) CreateBACnetConstructedDataStrikeCountBuilder() BACnetConstructedDataStrikeCountBuilder {
	if b == nil {
		return NewBACnetConstructedDataStrikeCountBuilder()
	}
	return &_BACnetConstructedDataStrikeCountBuilder{_BACnetConstructedDataStrikeCount: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataStrikeCount) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataStrikeCount) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_STRIKE_COUNT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataStrikeCount) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataStrikeCount) GetStrikeCount() BACnetApplicationTagUnsignedInteger {
	return m.StrikeCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataStrikeCount) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetStrikeCount())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataStrikeCount(structType any) BACnetConstructedDataStrikeCount {
	if casted, ok := structType.(BACnetConstructedDataStrikeCount); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataStrikeCount); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataStrikeCount) GetTypeName() string {
	return "BACnetConstructedDataStrikeCount"
}

func (m *_BACnetConstructedDataStrikeCount) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (strikeCount)
	lengthInBits += m.StrikeCount.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataStrikeCount) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataStrikeCount) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataStrikeCount BACnetConstructedDataStrikeCount, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataStrikeCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataStrikeCount")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	strikeCount, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "strikeCount", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'strikeCount' field"))
	}
	m.StrikeCount = strikeCount

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), strikeCount)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataStrikeCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataStrikeCount")
	}

	return m, nil
}

func (m *_BACnetConstructedDataStrikeCount) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataStrikeCount) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataStrikeCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataStrikeCount")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "strikeCount", m.GetStrikeCount(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'strikeCount' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataStrikeCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataStrikeCount")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataStrikeCount) IsBACnetConstructedDataStrikeCount() {}

func (m *_BACnetConstructedDataStrikeCount) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataStrikeCount) deepCopy() *_BACnetConstructedDataStrikeCount {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataStrikeCountCopy := &_BACnetConstructedDataStrikeCount{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.StrikeCount),
	}
	_BACnetConstructedDataStrikeCountCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataStrikeCountCopy
}

func (m *_BACnetConstructedDataStrikeCount) String() string {
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
