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

// BACnetChannelValueUnsigned is the corresponding interface of BACnetChannelValueUnsigned
type BACnetChannelValueUnsigned interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetChannelValue
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetChannelValueUnsigned is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetChannelValueUnsigned()
	// CreateBuilder creates a BACnetChannelValueUnsignedBuilder
	CreateBACnetChannelValueUnsignedBuilder() BACnetChannelValueUnsignedBuilder
}

// _BACnetChannelValueUnsigned is the data-structure of this message
type _BACnetChannelValueUnsigned struct {
	BACnetChannelValueContract
	UnsignedValue BACnetApplicationTagUnsignedInteger
}

var _ BACnetChannelValueUnsigned = (*_BACnetChannelValueUnsigned)(nil)
var _ BACnetChannelValueRequirements = (*_BACnetChannelValueUnsigned)(nil)

// NewBACnetChannelValueUnsigned factory function for _BACnetChannelValueUnsigned
func NewBACnetChannelValueUnsigned(peekedTagHeader BACnetTagHeader, unsignedValue BACnetApplicationTagUnsignedInteger) *_BACnetChannelValueUnsigned {
	if unsignedValue == nil {
		panic("unsignedValue of type BACnetApplicationTagUnsignedInteger for BACnetChannelValueUnsigned must not be nil")
	}
	_result := &_BACnetChannelValueUnsigned{
		BACnetChannelValueContract: NewBACnetChannelValue(peekedTagHeader),
		UnsignedValue:              unsignedValue,
	}
	_result.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetChannelValueUnsignedBuilder is a builder for BACnetChannelValueUnsigned
type BACnetChannelValueUnsignedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(unsignedValue BACnetApplicationTagUnsignedInteger) BACnetChannelValueUnsignedBuilder
	// WithUnsignedValue adds UnsignedValue (property field)
	WithUnsignedValue(BACnetApplicationTagUnsignedInteger) BACnetChannelValueUnsignedBuilder
	// WithUnsignedValueBuilder adds UnsignedValue (property field) which is build by the builder
	WithUnsignedValueBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetChannelValueUnsignedBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetChannelValueBuilder
	// Build builds the BACnetChannelValueUnsigned or returns an error if something is wrong
	Build() (BACnetChannelValueUnsigned, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetChannelValueUnsigned
}

// NewBACnetChannelValueUnsignedBuilder() creates a BACnetChannelValueUnsignedBuilder
func NewBACnetChannelValueUnsignedBuilder() BACnetChannelValueUnsignedBuilder {
	return &_BACnetChannelValueUnsignedBuilder{_BACnetChannelValueUnsigned: new(_BACnetChannelValueUnsigned)}
}

type _BACnetChannelValueUnsignedBuilder struct {
	*_BACnetChannelValueUnsigned

	parentBuilder *_BACnetChannelValueBuilder

	err *utils.MultiError
}

var _ (BACnetChannelValueUnsignedBuilder) = (*_BACnetChannelValueUnsignedBuilder)(nil)

func (b *_BACnetChannelValueUnsignedBuilder) setParent(contract BACnetChannelValueContract) {
	b.BACnetChannelValueContract = contract
	contract.(*_BACnetChannelValue)._SubType = b._BACnetChannelValueUnsigned
}

func (b *_BACnetChannelValueUnsignedBuilder) WithMandatoryFields(unsignedValue BACnetApplicationTagUnsignedInteger) BACnetChannelValueUnsignedBuilder {
	return b.WithUnsignedValue(unsignedValue)
}

func (b *_BACnetChannelValueUnsignedBuilder) WithUnsignedValue(unsignedValue BACnetApplicationTagUnsignedInteger) BACnetChannelValueUnsignedBuilder {
	b.UnsignedValue = unsignedValue
	return b
}

func (b *_BACnetChannelValueUnsignedBuilder) WithUnsignedValueBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetChannelValueUnsignedBuilder {
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

func (b *_BACnetChannelValueUnsignedBuilder) Build() (BACnetChannelValueUnsigned, error) {
	if b.UnsignedValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'unsignedValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetChannelValueUnsigned.deepCopy(), nil
}

func (b *_BACnetChannelValueUnsignedBuilder) MustBuild() BACnetChannelValueUnsigned {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetChannelValueUnsignedBuilder) Done() BACnetChannelValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetChannelValueBuilder().(*_BACnetChannelValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetChannelValueUnsignedBuilder) buildForBACnetChannelValue() (BACnetChannelValue, error) {
	return b.Build()
}

func (b *_BACnetChannelValueUnsignedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetChannelValueUnsignedBuilder().(*_BACnetChannelValueUnsignedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetChannelValueUnsignedBuilder creates a BACnetChannelValueUnsignedBuilder
func (b *_BACnetChannelValueUnsigned) CreateBACnetChannelValueUnsignedBuilder() BACnetChannelValueUnsignedBuilder {
	if b == nil {
		return NewBACnetChannelValueUnsignedBuilder()
	}
	return &_BACnetChannelValueUnsignedBuilder{_BACnetChannelValueUnsigned: b.deepCopy()}
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

func (m *_BACnetChannelValueUnsigned) GetParent() BACnetChannelValueContract {
	return m.BACnetChannelValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueUnsigned) GetUnsignedValue() BACnetApplicationTagUnsignedInteger {
	return m.UnsignedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueUnsigned(structType any) BACnetChannelValueUnsigned {
	if casted, ok := structType.(BACnetChannelValueUnsigned); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueUnsigned); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueUnsigned) GetTypeName() string {
	return "BACnetChannelValueUnsigned"
}

func (m *_BACnetChannelValueUnsigned) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetChannelValueContract.(*_BACnetChannelValue).getLengthInBits(ctx))

	// Simple field (unsignedValue)
	lengthInBits += m.UnsignedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetChannelValueUnsigned) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetChannelValueUnsigned) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetChannelValue) (__bACnetChannelValueUnsigned BACnetChannelValueUnsigned, err error) {
	m.BACnetChannelValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueUnsigned"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueUnsigned")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	unsignedValue, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "unsignedValue", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unsignedValue' field"))
	}
	m.UnsignedValue = unsignedValue

	if closeErr := readBuffer.CloseContext("BACnetChannelValueUnsigned"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueUnsigned")
	}

	return m, nil
}

func (m *_BACnetChannelValueUnsigned) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueUnsigned) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueUnsigned"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueUnsigned")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "unsignedValue", m.GetUnsignedValue(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'unsignedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueUnsigned"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueUnsigned")
		}
		return nil
	}
	return m.BACnetChannelValueContract.(*_BACnetChannelValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetChannelValueUnsigned) IsBACnetChannelValueUnsigned() {}

func (m *_BACnetChannelValueUnsigned) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetChannelValueUnsigned) deepCopy() *_BACnetChannelValueUnsigned {
	if m == nil {
		return nil
	}
	_BACnetChannelValueUnsignedCopy := &_BACnetChannelValueUnsigned{
		m.BACnetChannelValueContract.(*_BACnetChannelValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.UnsignedValue),
	}
	_BACnetChannelValueUnsignedCopy.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = m
	return _BACnetChannelValueUnsignedCopy
}

func (m *_BACnetChannelValueUnsigned) String() string {
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