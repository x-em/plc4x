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

// BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned
type BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameterFaultOutOfRangeMaxNormalValue
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned()
	// CreateBuilder creates a BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder
	CreateBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder() BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder
}

// _BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned struct {
	BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract
	UnsignedValue BACnetApplicationTagUnsignedInteger
}

var _ BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned = (*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned)(nil)
var _ BACnetFaultParameterFaultOutOfRangeMaxNormalValueRequirements = (*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned)(nil)

// NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned factory function for _BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned
func NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, unsignedValue BACnetApplicationTagUnsignedInteger, tagNumber uint8) *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned {
	if unsignedValue == nil {
		panic("unsignedValue of type BACnetApplicationTagUnsignedInteger for BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned must not be nil")
	}
	_result := &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned{
		BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract: NewBACnetFaultParameterFaultOutOfRangeMaxNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
		UnsignedValue: unsignedValue,
	}
	_result.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder is a builder for BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned
type BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(unsignedValue BACnetApplicationTagUnsignedInteger) BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder
	// WithUnsignedValue adds UnsignedValue (property field)
	WithUnsignedValue(BACnetApplicationTagUnsignedInteger) BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder
	// WithUnsignedValueBuilder adds UnsignedValue (property field) which is build by the builder
	WithUnsignedValueBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetFaultParameterFaultOutOfRangeMaxNormalValueBuilder
	// Build builds the BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned or returns an error if something is wrong
	Build() (BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned
}

// NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder() creates a BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder
func NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder() BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder {
	return &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder{_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned: new(_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned)}
}

type _BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder struct {
	*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned

	parentBuilder *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueBuilder

	err *utils.MultiError
}

var _ (BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder) = (*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder)(nil)

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder) setParent(contract BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract) {
	b.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract = contract
	contract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue)._SubType = b._BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder) WithMandatoryFields(unsignedValue BACnetApplicationTagUnsignedInteger) BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder {
	return b.WithUnsignedValue(unsignedValue)
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder) WithUnsignedValue(unsignedValue BACnetApplicationTagUnsignedInteger) BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder {
	b.UnsignedValue = unsignedValue
	return b
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder) WithUnsignedValueBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder {
	builder := builderSupplier(b.UnsignedValue.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.UnsignedValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder) Build() (BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned, error) {
	if b.UnsignedValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'unsignedValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned.deepCopy(), nil
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder) MustBuild() BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder) Done() BACnetFaultParameterFaultOutOfRangeMaxNormalValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueBuilder().(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder) buildForBACnetFaultParameterFaultOutOfRangeMaxNormalValue() (BACnetFaultParameterFaultOutOfRangeMaxNormalValue, error) {
	return b.Build()
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder().(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder creates a BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder
func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) CreateBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder() BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder {
	if b == nil {
		return NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder()
	}
	return &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedBuilder{_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned: b.deepCopy()}
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

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) GetParent() BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract {
	return m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) GetUnsignedValue() BACnetApplicationTagUnsignedInteger {
	return m.UnsignedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned(structType any) BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned"
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue).getLengthInBits(ctx))

	// Simple field (unsignedValue)
	lengthInBits += m.UnsignedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultOutOfRangeMaxNormalValue, tagNumber uint8) (__bACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned, err error) {
	m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	unsignedValue, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "unsignedValue", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unsignedValue' field"))
	}
	m.UnsignedValue = unsignedValue

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "unsignedValue", m.GetUnsignedValue(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'unsignedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) IsBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned() {
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) deepCopy() *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedCopy := &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned{
		m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.UnsignedValue),
	}
	_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedCopy.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue)._SubType = m
	return _BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedCopy
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) String() string {
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
