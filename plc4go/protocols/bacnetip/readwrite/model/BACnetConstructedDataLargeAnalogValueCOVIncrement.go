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

// BACnetConstructedDataLargeAnalogValueCOVIncrement is the corresponding interface of BACnetConstructedDataLargeAnalogValueCOVIncrement
type BACnetConstructedDataLargeAnalogValueCOVIncrement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetCovIncrement returns CovIncrement (property field)
	GetCovIncrement() BACnetApplicationTagDouble
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagDouble
	// IsBACnetConstructedDataLargeAnalogValueCOVIncrement is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLargeAnalogValueCOVIncrement()
	// CreateBuilder creates a BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder
	CreateBACnetConstructedDataLargeAnalogValueCOVIncrementBuilder() BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder
}

// _BACnetConstructedDataLargeAnalogValueCOVIncrement is the data-structure of this message
type _BACnetConstructedDataLargeAnalogValueCOVIncrement struct {
	BACnetConstructedDataContract
	CovIncrement BACnetApplicationTagDouble
}

var _ BACnetConstructedDataLargeAnalogValueCOVIncrement = (*_BACnetConstructedDataLargeAnalogValueCOVIncrement)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLargeAnalogValueCOVIncrement)(nil)

// NewBACnetConstructedDataLargeAnalogValueCOVIncrement factory function for _BACnetConstructedDataLargeAnalogValueCOVIncrement
func NewBACnetConstructedDataLargeAnalogValueCOVIncrement(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, covIncrement BACnetApplicationTagDouble, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLargeAnalogValueCOVIncrement {
	if covIncrement == nil {
		panic("covIncrement of type BACnetApplicationTagDouble for BACnetConstructedDataLargeAnalogValueCOVIncrement must not be nil")
	}
	_result := &_BACnetConstructedDataLargeAnalogValueCOVIncrement{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		CovIncrement:                  covIncrement,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder is a builder for BACnetConstructedDataLargeAnalogValueCOVIncrement
type BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(covIncrement BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder
	// WithCovIncrement adds CovIncrement (property field)
	WithCovIncrement(BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder
	// WithCovIncrementBuilder adds CovIncrement (property field) which is build by the builder
	WithCovIncrementBuilder(func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLargeAnalogValueCOVIncrement or returns an error if something is wrong
	Build() (BACnetConstructedDataLargeAnalogValueCOVIncrement, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLargeAnalogValueCOVIncrement
}

// NewBACnetConstructedDataLargeAnalogValueCOVIncrementBuilder() creates a BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder
func NewBACnetConstructedDataLargeAnalogValueCOVIncrementBuilder() BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder {
	return &_BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder{_BACnetConstructedDataLargeAnalogValueCOVIncrement: new(_BACnetConstructedDataLargeAnalogValueCOVIncrement)}
}

type _BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder struct {
	*_BACnetConstructedDataLargeAnalogValueCOVIncrement

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder) = (*_BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder)(nil)

func (b *_BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataLargeAnalogValueCOVIncrement
}

func (b *_BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder) WithMandatoryFields(covIncrement BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder {
	return b.WithCovIncrement(covIncrement)
}

func (b *_BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder) WithCovIncrement(covIncrement BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder {
	b.CovIncrement = covIncrement
	return b
}

func (b *_BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder) WithCovIncrementBuilder(builderSupplier func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder {
	builder := builderSupplier(b.CovIncrement.CreateBACnetApplicationTagDoubleBuilder())
	var err error
	b.CovIncrement, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagDoubleBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder) Build() (BACnetConstructedDataLargeAnalogValueCOVIncrement, error) {
	if b.CovIncrement == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'covIncrement' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLargeAnalogValueCOVIncrement.deepCopy(), nil
}

func (b *_BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder) MustBuild() BACnetConstructedDataLargeAnalogValueCOVIncrement {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLargeAnalogValueCOVIncrementBuilder().(*_BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLargeAnalogValueCOVIncrementBuilder creates a BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder
func (b *_BACnetConstructedDataLargeAnalogValueCOVIncrement) CreateBACnetConstructedDataLargeAnalogValueCOVIncrementBuilder() BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder {
	if b == nil {
		return NewBACnetConstructedDataLargeAnalogValueCOVIncrementBuilder()
	}
	return &_BACnetConstructedDataLargeAnalogValueCOVIncrementBuilder{_BACnetConstructedDataLargeAnalogValueCOVIncrement: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueCOVIncrement) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LARGE_ANALOG_VALUE
}

func (m *_BACnetConstructedDataLargeAnalogValueCOVIncrement) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_COV_INCREMENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueCOVIncrement) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueCOVIncrement) GetCovIncrement() BACnetApplicationTagDouble {
	return m.CovIncrement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueCOVIncrement) GetActualValue() BACnetApplicationTagDouble {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagDouble(m.GetCovIncrement())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLargeAnalogValueCOVIncrement(structType any) BACnetConstructedDataLargeAnalogValueCOVIncrement {
	if casted, ok := structType.(BACnetConstructedDataLargeAnalogValueCOVIncrement); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLargeAnalogValueCOVIncrement); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLargeAnalogValueCOVIncrement) GetTypeName() string {
	return "BACnetConstructedDataLargeAnalogValueCOVIncrement"
}

func (m *_BACnetConstructedDataLargeAnalogValueCOVIncrement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (covIncrement)
	lengthInBits += m.CovIncrement.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLargeAnalogValueCOVIncrement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLargeAnalogValueCOVIncrement) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLargeAnalogValueCOVIncrement BACnetConstructedDataLargeAnalogValueCOVIncrement, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLargeAnalogValueCOVIncrement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLargeAnalogValueCOVIncrement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	covIncrement, err := ReadSimpleField[BACnetApplicationTagDouble](ctx, "covIncrement", ReadComplex[BACnetApplicationTagDouble](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDouble](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'covIncrement' field"))
	}
	m.CovIncrement = covIncrement

	actualValue, err := ReadVirtualField[BACnetApplicationTagDouble](ctx, "actualValue", (*BACnetApplicationTagDouble)(nil), covIncrement)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLargeAnalogValueCOVIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLargeAnalogValueCOVIncrement")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLargeAnalogValueCOVIncrement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLargeAnalogValueCOVIncrement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLargeAnalogValueCOVIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLargeAnalogValueCOVIncrement")
		}

		if err := WriteSimpleField[BACnetApplicationTagDouble](ctx, "covIncrement", m.GetCovIncrement(), WriteComplex[BACnetApplicationTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'covIncrement' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLargeAnalogValueCOVIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLargeAnalogValueCOVIncrement")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLargeAnalogValueCOVIncrement) IsBACnetConstructedDataLargeAnalogValueCOVIncrement() {
}

func (m *_BACnetConstructedDataLargeAnalogValueCOVIncrement) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLargeAnalogValueCOVIncrement) deepCopy() *_BACnetConstructedDataLargeAnalogValueCOVIncrement {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLargeAnalogValueCOVIncrementCopy := &_BACnetConstructedDataLargeAnalogValueCOVIncrement{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagDouble](m.CovIncrement),
	}
	_BACnetConstructedDataLargeAnalogValueCOVIncrementCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLargeAnalogValueCOVIncrementCopy
}

func (m *_BACnetConstructedDataLargeAnalogValueCOVIncrement) String() string {
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