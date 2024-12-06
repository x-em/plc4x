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

// BACnetConstructedDataLowLimit is the corresponding interface of BACnetConstructedDataLowLimit
type BACnetConstructedDataLowLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLowLimit returns LowLimit (property field)
	GetLowLimit() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataLowLimit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLowLimit()
	// CreateBuilder creates a BACnetConstructedDataLowLimitBuilder
	CreateBACnetConstructedDataLowLimitBuilder() BACnetConstructedDataLowLimitBuilder
}

// _BACnetConstructedDataLowLimit is the data-structure of this message
type _BACnetConstructedDataLowLimit struct {
	BACnetConstructedDataContract
	LowLimit BACnetApplicationTagReal
}

var _ BACnetConstructedDataLowLimit = (*_BACnetConstructedDataLowLimit)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLowLimit)(nil)

// NewBACnetConstructedDataLowLimit factory function for _BACnetConstructedDataLowLimit
func NewBACnetConstructedDataLowLimit(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lowLimit BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLowLimit {
	if lowLimit == nil {
		panic("lowLimit of type BACnetApplicationTagReal for BACnetConstructedDataLowLimit must not be nil")
	}
	_result := &_BACnetConstructedDataLowLimit{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LowLimit:                      lowLimit,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLowLimitBuilder is a builder for BACnetConstructedDataLowLimit
type BACnetConstructedDataLowLimitBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lowLimit BACnetApplicationTagReal) BACnetConstructedDataLowLimitBuilder
	// WithLowLimit adds LowLimit (property field)
	WithLowLimit(BACnetApplicationTagReal) BACnetConstructedDataLowLimitBuilder
	// WithLowLimitBuilder adds LowLimit (property field) which is build by the builder
	WithLowLimitBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataLowLimitBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLowLimit or returns an error if something is wrong
	Build() (BACnetConstructedDataLowLimit, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLowLimit
}

// NewBACnetConstructedDataLowLimitBuilder() creates a BACnetConstructedDataLowLimitBuilder
func NewBACnetConstructedDataLowLimitBuilder() BACnetConstructedDataLowLimitBuilder {
	return &_BACnetConstructedDataLowLimitBuilder{_BACnetConstructedDataLowLimit: new(_BACnetConstructedDataLowLimit)}
}

type _BACnetConstructedDataLowLimitBuilder struct {
	*_BACnetConstructedDataLowLimit

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLowLimitBuilder) = (*_BACnetConstructedDataLowLimitBuilder)(nil)

func (b *_BACnetConstructedDataLowLimitBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataLowLimit
}

func (b *_BACnetConstructedDataLowLimitBuilder) WithMandatoryFields(lowLimit BACnetApplicationTagReal) BACnetConstructedDataLowLimitBuilder {
	return b.WithLowLimit(lowLimit)
}

func (b *_BACnetConstructedDataLowLimitBuilder) WithLowLimit(lowLimit BACnetApplicationTagReal) BACnetConstructedDataLowLimitBuilder {
	b.LowLimit = lowLimit
	return b
}

func (b *_BACnetConstructedDataLowLimitBuilder) WithLowLimitBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataLowLimitBuilder {
	builder := builderSupplier(b.LowLimit.CreateBACnetApplicationTagRealBuilder())
	var err error
	b.LowLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLowLimitBuilder) Build() (BACnetConstructedDataLowLimit, error) {
	if b.LowLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lowLimit' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLowLimit.deepCopy(), nil
}

func (b *_BACnetConstructedDataLowLimitBuilder) MustBuild() BACnetConstructedDataLowLimit {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLowLimitBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLowLimitBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLowLimitBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLowLimitBuilder().(*_BACnetConstructedDataLowLimitBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLowLimitBuilder creates a BACnetConstructedDataLowLimitBuilder
func (b *_BACnetConstructedDataLowLimit) CreateBACnetConstructedDataLowLimitBuilder() BACnetConstructedDataLowLimitBuilder {
	if b == nil {
		return NewBACnetConstructedDataLowLimitBuilder()
	}
	return &_BACnetConstructedDataLowLimitBuilder{_BACnetConstructedDataLowLimit: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLowLimit) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLowLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOW_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLowLimit) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLowLimit) GetLowLimit() BACnetApplicationTagReal {
	return m.LowLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLowLimit) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetLowLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLowLimit(structType any) BACnetConstructedDataLowLimit {
	if casted, ok := structType.(BACnetConstructedDataLowLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLowLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLowLimit) GetTypeName() string {
	return "BACnetConstructedDataLowLimit"
}

func (m *_BACnetConstructedDataLowLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lowLimit)
	lengthInBits += m.LowLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLowLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLowLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLowLimit BACnetConstructedDataLowLimit, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLowLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lowLimit, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "lowLimit", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lowLimit' field"))
	}
	m.LowLimit = lowLimit

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), lowLimit)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLowLimit")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLowLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLowLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLowLimit")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "lowLimit", m.GetLowLimit(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lowLimit' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLowLimit")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLowLimit) IsBACnetConstructedDataLowLimit() {}

func (m *_BACnetConstructedDataLowLimit) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLowLimit) deepCopy() *_BACnetConstructedDataLowLimit {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLowLimitCopy := &_BACnetConstructedDataLowLimit{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagReal](m.LowLimit),
	}
	_BACnetConstructedDataLowLimitCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLowLimitCopy
}

func (m *_BACnetConstructedDataLowLimit) String() string {
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