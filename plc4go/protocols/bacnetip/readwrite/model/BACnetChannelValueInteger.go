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

// BACnetChannelValueInteger is the corresponding interface of BACnetChannelValueInteger
type BACnetChannelValueInteger interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetChannelValue
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() BACnetApplicationTagSignedInteger
	// IsBACnetChannelValueInteger is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetChannelValueInteger()
	// CreateBuilder creates a BACnetChannelValueIntegerBuilder
	CreateBACnetChannelValueIntegerBuilder() BACnetChannelValueIntegerBuilder
}

// _BACnetChannelValueInteger is the data-structure of this message
type _BACnetChannelValueInteger struct {
	BACnetChannelValueContract
	IntegerValue BACnetApplicationTagSignedInteger
}

var _ BACnetChannelValueInteger = (*_BACnetChannelValueInteger)(nil)
var _ BACnetChannelValueRequirements = (*_BACnetChannelValueInteger)(nil)

// NewBACnetChannelValueInteger factory function for _BACnetChannelValueInteger
func NewBACnetChannelValueInteger(peekedTagHeader BACnetTagHeader, integerValue BACnetApplicationTagSignedInteger) *_BACnetChannelValueInteger {
	if integerValue == nil {
		panic("integerValue of type BACnetApplicationTagSignedInteger for BACnetChannelValueInteger must not be nil")
	}
	_result := &_BACnetChannelValueInteger{
		BACnetChannelValueContract: NewBACnetChannelValue(peekedTagHeader),
		IntegerValue:               integerValue,
	}
	_result.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetChannelValueIntegerBuilder is a builder for BACnetChannelValueInteger
type BACnetChannelValueIntegerBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(integerValue BACnetApplicationTagSignedInteger) BACnetChannelValueIntegerBuilder
	// WithIntegerValue adds IntegerValue (property field)
	WithIntegerValue(BACnetApplicationTagSignedInteger) BACnetChannelValueIntegerBuilder
	// WithIntegerValueBuilder adds IntegerValue (property field) which is build by the builder
	WithIntegerValueBuilder(func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetChannelValueIntegerBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetChannelValueBuilder
	// Build builds the BACnetChannelValueInteger or returns an error if something is wrong
	Build() (BACnetChannelValueInteger, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetChannelValueInteger
}

// NewBACnetChannelValueIntegerBuilder() creates a BACnetChannelValueIntegerBuilder
func NewBACnetChannelValueIntegerBuilder() BACnetChannelValueIntegerBuilder {
	return &_BACnetChannelValueIntegerBuilder{_BACnetChannelValueInteger: new(_BACnetChannelValueInteger)}
}

type _BACnetChannelValueIntegerBuilder struct {
	*_BACnetChannelValueInteger

	parentBuilder *_BACnetChannelValueBuilder

	err *utils.MultiError
}

var _ (BACnetChannelValueIntegerBuilder) = (*_BACnetChannelValueIntegerBuilder)(nil)

func (b *_BACnetChannelValueIntegerBuilder) setParent(contract BACnetChannelValueContract) {
	b.BACnetChannelValueContract = contract
	contract.(*_BACnetChannelValue)._SubType = b._BACnetChannelValueInteger
}

func (b *_BACnetChannelValueIntegerBuilder) WithMandatoryFields(integerValue BACnetApplicationTagSignedInteger) BACnetChannelValueIntegerBuilder {
	return b.WithIntegerValue(integerValue)
}

func (b *_BACnetChannelValueIntegerBuilder) WithIntegerValue(integerValue BACnetApplicationTagSignedInteger) BACnetChannelValueIntegerBuilder {
	b.IntegerValue = integerValue
	return b
}

func (b *_BACnetChannelValueIntegerBuilder) WithIntegerValueBuilder(builderSupplier func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetChannelValueIntegerBuilder {
	builder := builderSupplier(b.IntegerValue.CreateBACnetApplicationTagSignedIntegerBuilder())
	var err error
	b.IntegerValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetChannelValueIntegerBuilder) Build() (BACnetChannelValueInteger, error) {
	if b.IntegerValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'integerValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetChannelValueInteger.deepCopy(), nil
}

func (b *_BACnetChannelValueIntegerBuilder) MustBuild() BACnetChannelValueInteger {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetChannelValueIntegerBuilder) Done() BACnetChannelValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetChannelValueBuilder().(*_BACnetChannelValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetChannelValueIntegerBuilder) buildForBACnetChannelValue() (BACnetChannelValue, error) {
	return b.Build()
}

func (b *_BACnetChannelValueIntegerBuilder) DeepCopy() any {
	_copy := b.CreateBACnetChannelValueIntegerBuilder().(*_BACnetChannelValueIntegerBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetChannelValueIntegerBuilder creates a BACnetChannelValueIntegerBuilder
func (b *_BACnetChannelValueInteger) CreateBACnetChannelValueIntegerBuilder() BACnetChannelValueIntegerBuilder {
	if b == nil {
		return NewBACnetChannelValueIntegerBuilder()
	}
	return &_BACnetChannelValueIntegerBuilder{_BACnetChannelValueInteger: b.deepCopy()}
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

func (m *_BACnetChannelValueInteger) GetParent() BACnetChannelValueContract {
	return m.BACnetChannelValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueInteger) GetIntegerValue() BACnetApplicationTagSignedInteger {
	return m.IntegerValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueInteger(structType any) BACnetChannelValueInteger {
	if casted, ok := structType.(BACnetChannelValueInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueInteger) GetTypeName() string {
	return "BACnetChannelValueInteger"
}

func (m *_BACnetChannelValueInteger) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetChannelValueContract.(*_BACnetChannelValue).getLengthInBits(ctx))

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetChannelValueInteger) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetChannelValueInteger) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetChannelValue) (__bACnetChannelValueInteger BACnetChannelValueInteger, err error) {
	m.BACnetChannelValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	integerValue, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "integerValue", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'integerValue' field"))
	}
	m.IntegerValue = integerValue

	if closeErr := readBuffer.CloseContext("BACnetChannelValueInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueInteger")
	}

	return m, nil
}

func (m *_BACnetChannelValueInteger) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueInteger) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueInteger")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "integerValue", m.GetIntegerValue(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'integerValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueInteger")
		}
		return nil
	}
	return m.BACnetChannelValueContract.(*_BACnetChannelValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetChannelValueInteger) IsBACnetChannelValueInteger() {}

func (m *_BACnetChannelValueInteger) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetChannelValueInteger) deepCopy() *_BACnetChannelValueInteger {
	if m == nil {
		return nil
	}
	_BACnetChannelValueIntegerCopy := &_BACnetChannelValueInteger{
		m.BACnetChannelValueContract.(*_BACnetChannelValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagSignedInteger](m.IntegerValue),
	}
	_BACnetChannelValueIntegerCopy.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = m
	return _BACnetChannelValueIntegerCopy
}

func (m *_BACnetChannelValueInteger) String() string {
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