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

// BACnetConstructedDataLowDiffLimit is the corresponding interface of BACnetConstructedDataLowDiffLimit
type BACnetConstructedDataLowDiffLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLowDiffLimit returns LowDiffLimit (property field)
	GetLowDiffLimit() BACnetOptionalREAL
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetOptionalREAL
	// IsBACnetConstructedDataLowDiffLimit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLowDiffLimit()
	// CreateBuilder creates a BACnetConstructedDataLowDiffLimitBuilder
	CreateBACnetConstructedDataLowDiffLimitBuilder() BACnetConstructedDataLowDiffLimitBuilder
}

// _BACnetConstructedDataLowDiffLimit is the data-structure of this message
type _BACnetConstructedDataLowDiffLimit struct {
	BACnetConstructedDataContract
	LowDiffLimit BACnetOptionalREAL
}

var _ BACnetConstructedDataLowDiffLimit = (*_BACnetConstructedDataLowDiffLimit)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLowDiffLimit)(nil)

// NewBACnetConstructedDataLowDiffLimit factory function for _BACnetConstructedDataLowDiffLimit
func NewBACnetConstructedDataLowDiffLimit(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lowDiffLimit BACnetOptionalREAL, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLowDiffLimit {
	if lowDiffLimit == nil {
		panic("lowDiffLimit of type BACnetOptionalREAL for BACnetConstructedDataLowDiffLimit must not be nil")
	}
	_result := &_BACnetConstructedDataLowDiffLimit{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LowDiffLimit:                  lowDiffLimit,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLowDiffLimitBuilder is a builder for BACnetConstructedDataLowDiffLimit
type BACnetConstructedDataLowDiffLimitBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lowDiffLimit BACnetOptionalREAL) BACnetConstructedDataLowDiffLimitBuilder
	// WithLowDiffLimit adds LowDiffLimit (property field)
	WithLowDiffLimit(BACnetOptionalREAL) BACnetConstructedDataLowDiffLimitBuilder
	// WithLowDiffLimitBuilder adds LowDiffLimit (property field) which is build by the builder
	WithLowDiffLimitBuilder(func(BACnetOptionalREALBuilder) BACnetOptionalREALBuilder) BACnetConstructedDataLowDiffLimitBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLowDiffLimit or returns an error if something is wrong
	Build() (BACnetConstructedDataLowDiffLimit, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLowDiffLimit
}

// NewBACnetConstructedDataLowDiffLimitBuilder() creates a BACnetConstructedDataLowDiffLimitBuilder
func NewBACnetConstructedDataLowDiffLimitBuilder() BACnetConstructedDataLowDiffLimitBuilder {
	return &_BACnetConstructedDataLowDiffLimitBuilder{_BACnetConstructedDataLowDiffLimit: new(_BACnetConstructedDataLowDiffLimit)}
}

type _BACnetConstructedDataLowDiffLimitBuilder struct {
	*_BACnetConstructedDataLowDiffLimit

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLowDiffLimitBuilder) = (*_BACnetConstructedDataLowDiffLimitBuilder)(nil)

func (b *_BACnetConstructedDataLowDiffLimitBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataLowDiffLimit
}

func (b *_BACnetConstructedDataLowDiffLimitBuilder) WithMandatoryFields(lowDiffLimit BACnetOptionalREAL) BACnetConstructedDataLowDiffLimitBuilder {
	return b.WithLowDiffLimit(lowDiffLimit)
}

func (b *_BACnetConstructedDataLowDiffLimitBuilder) WithLowDiffLimit(lowDiffLimit BACnetOptionalREAL) BACnetConstructedDataLowDiffLimitBuilder {
	b.LowDiffLimit = lowDiffLimit
	return b
}

func (b *_BACnetConstructedDataLowDiffLimitBuilder) WithLowDiffLimitBuilder(builderSupplier func(BACnetOptionalREALBuilder) BACnetOptionalREALBuilder) BACnetConstructedDataLowDiffLimitBuilder {
	builder := builderSupplier(b.LowDiffLimit.CreateBACnetOptionalREALBuilder())
	var err error
	b.LowDiffLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOptionalREALBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLowDiffLimitBuilder) Build() (BACnetConstructedDataLowDiffLimit, error) {
	if b.LowDiffLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lowDiffLimit' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLowDiffLimit.deepCopy(), nil
}

func (b *_BACnetConstructedDataLowDiffLimitBuilder) MustBuild() BACnetConstructedDataLowDiffLimit {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLowDiffLimitBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLowDiffLimitBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLowDiffLimitBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLowDiffLimitBuilder().(*_BACnetConstructedDataLowDiffLimitBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLowDiffLimitBuilder creates a BACnetConstructedDataLowDiffLimitBuilder
func (b *_BACnetConstructedDataLowDiffLimit) CreateBACnetConstructedDataLowDiffLimitBuilder() BACnetConstructedDataLowDiffLimitBuilder {
	if b == nil {
		return NewBACnetConstructedDataLowDiffLimitBuilder()
	}
	return &_BACnetConstructedDataLowDiffLimitBuilder{_BACnetConstructedDataLowDiffLimit: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLowDiffLimit) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLowDiffLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOW_DIFF_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLowDiffLimit) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLowDiffLimit) GetLowDiffLimit() BACnetOptionalREAL {
	return m.LowDiffLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLowDiffLimit) GetActualValue() BACnetOptionalREAL {
	ctx := context.Background()
	_ = ctx
	return CastBACnetOptionalREAL(m.GetLowDiffLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLowDiffLimit(structType any) BACnetConstructedDataLowDiffLimit {
	if casted, ok := structType.(BACnetConstructedDataLowDiffLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLowDiffLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLowDiffLimit) GetTypeName() string {
	return "BACnetConstructedDataLowDiffLimit"
}

func (m *_BACnetConstructedDataLowDiffLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lowDiffLimit)
	lengthInBits += m.LowDiffLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLowDiffLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLowDiffLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLowDiffLimit BACnetConstructedDataLowDiffLimit, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLowDiffLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLowDiffLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lowDiffLimit, err := ReadSimpleField[BACnetOptionalREAL](ctx, "lowDiffLimit", ReadComplex[BACnetOptionalREAL](BACnetOptionalREALParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lowDiffLimit' field"))
	}
	m.LowDiffLimit = lowDiffLimit

	actualValue, err := ReadVirtualField[BACnetOptionalREAL](ctx, "actualValue", (*BACnetOptionalREAL)(nil), lowDiffLimit)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLowDiffLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLowDiffLimit")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLowDiffLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLowDiffLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLowDiffLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLowDiffLimit")
		}

		if err := WriteSimpleField[BACnetOptionalREAL](ctx, "lowDiffLimit", m.GetLowDiffLimit(), WriteComplex[BACnetOptionalREAL](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lowDiffLimit' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLowDiffLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLowDiffLimit")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLowDiffLimit) IsBACnetConstructedDataLowDiffLimit() {}

func (m *_BACnetConstructedDataLowDiffLimit) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLowDiffLimit) deepCopy() *_BACnetConstructedDataLowDiffLimit {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLowDiffLimitCopy := &_BACnetConstructedDataLowDiffLimit{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetOptionalREAL](m.LowDiffLimit),
	}
	_BACnetConstructedDataLowDiffLimitCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLowDiffLimitCopy
}

func (m *_BACnetConstructedDataLowDiffLimit) String() string {
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