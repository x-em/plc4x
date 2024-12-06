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

// BACnetOptionalUnsignedNull is the corresponding interface of BACnetOptionalUnsignedNull
type BACnetOptionalUnsignedNull interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetOptionalUnsigned
	// GetNullValue returns NullValue (property field)
	GetNullValue() BACnetApplicationTagNull
	// IsBACnetOptionalUnsignedNull is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetOptionalUnsignedNull()
	// CreateBuilder creates a BACnetOptionalUnsignedNullBuilder
	CreateBACnetOptionalUnsignedNullBuilder() BACnetOptionalUnsignedNullBuilder
}

// _BACnetOptionalUnsignedNull is the data-structure of this message
type _BACnetOptionalUnsignedNull struct {
	BACnetOptionalUnsignedContract
	NullValue BACnetApplicationTagNull
}

var _ BACnetOptionalUnsignedNull = (*_BACnetOptionalUnsignedNull)(nil)
var _ BACnetOptionalUnsignedRequirements = (*_BACnetOptionalUnsignedNull)(nil)

// NewBACnetOptionalUnsignedNull factory function for _BACnetOptionalUnsignedNull
func NewBACnetOptionalUnsignedNull(peekedTagHeader BACnetTagHeader, nullValue BACnetApplicationTagNull) *_BACnetOptionalUnsignedNull {
	if nullValue == nil {
		panic("nullValue of type BACnetApplicationTagNull for BACnetOptionalUnsignedNull must not be nil")
	}
	_result := &_BACnetOptionalUnsignedNull{
		BACnetOptionalUnsignedContract: NewBACnetOptionalUnsigned(peekedTagHeader),
		NullValue:                      nullValue,
	}
	_result.BACnetOptionalUnsignedContract.(*_BACnetOptionalUnsigned)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetOptionalUnsignedNullBuilder is a builder for BACnetOptionalUnsignedNull
type BACnetOptionalUnsignedNullBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(nullValue BACnetApplicationTagNull) BACnetOptionalUnsignedNullBuilder
	// WithNullValue adds NullValue (property field)
	WithNullValue(BACnetApplicationTagNull) BACnetOptionalUnsignedNullBuilder
	// WithNullValueBuilder adds NullValue (property field) which is build by the builder
	WithNullValueBuilder(func(BACnetApplicationTagNullBuilder) BACnetApplicationTagNullBuilder) BACnetOptionalUnsignedNullBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetOptionalUnsignedBuilder
	// Build builds the BACnetOptionalUnsignedNull or returns an error if something is wrong
	Build() (BACnetOptionalUnsignedNull, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetOptionalUnsignedNull
}

// NewBACnetOptionalUnsignedNullBuilder() creates a BACnetOptionalUnsignedNullBuilder
func NewBACnetOptionalUnsignedNullBuilder() BACnetOptionalUnsignedNullBuilder {
	return &_BACnetOptionalUnsignedNullBuilder{_BACnetOptionalUnsignedNull: new(_BACnetOptionalUnsignedNull)}
}

type _BACnetOptionalUnsignedNullBuilder struct {
	*_BACnetOptionalUnsignedNull

	parentBuilder *_BACnetOptionalUnsignedBuilder

	err *utils.MultiError
}

var _ (BACnetOptionalUnsignedNullBuilder) = (*_BACnetOptionalUnsignedNullBuilder)(nil)

func (b *_BACnetOptionalUnsignedNullBuilder) setParent(contract BACnetOptionalUnsignedContract) {
	b.BACnetOptionalUnsignedContract = contract
	contract.(*_BACnetOptionalUnsigned)._SubType = b._BACnetOptionalUnsignedNull
}

func (b *_BACnetOptionalUnsignedNullBuilder) WithMandatoryFields(nullValue BACnetApplicationTagNull) BACnetOptionalUnsignedNullBuilder {
	return b.WithNullValue(nullValue)
}

func (b *_BACnetOptionalUnsignedNullBuilder) WithNullValue(nullValue BACnetApplicationTagNull) BACnetOptionalUnsignedNullBuilder {
	b.NullValue = nullValue
	return b
}

func (b *_BACnetOptionalUnsignedNullBuilder) WithNullValueBuilder(builderSupplier func(BACnetApplicationTagNullBuilder) BACnetApplicationTagNullBuilder) BACnetOptionalUnsignedNullBuilder {
	builder := builderSupplier(b.NullValue.CreateBACnetApplicationTagNullBuilder())
	var err error
	b.NullValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagNullBuilder failed"))
	}
	return b
}

func (b *_BACnetOptionalUnsignedNullBuilder) Build() (BACnetOptionalUnsignedNull, error) {
	if b.NullValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'nullValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetOptionalUnsignedNull.deepCopy(), nil
}

func (b *_BACnetOptionalUnsignedNullBuilder) MustBuild() BACnetOptionalUnsignedNull {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetOptionalUnsignedNullBuilder) Done() BACnetOptionalUnsignedBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetOptionalUnsignedBuilder().(*_BACnetOptionalUnsignedBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetOptionalUnsignedNullBuilder) buildForBACnetOptionalUnsigned() (BACnetOptionalUnsigned, error) {
	return b.Build()
}

func (b *_BACnetOptionalUnsignedNullBuilder) DeepCopy() any {
	_copy := b.CreateBACnetOptionalUnsignedNullBuilder().(*_BACnetOptionalUnsignedNullBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetOptionalUnsignedNullBuilder creates a BACnetOptionalUnsignedNullBuilder
func (b *_BACnetOptionalUnsignedNull) CreateBACnetOptionalUnsignedNullBuilder() BACnetOptionalUnsignedNullBuilder {
	if b == nil {
		return NewBACnetOptionalUnsignedNullBuilder()
	}
	return &_BACnetOptionalUnsignedNullBuilder{_BACnetOptionalUnsignedNull: b.deepCopy()}
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

func (m *_BACnetOptionalUnsignedNull) GetParent() BACnetOptionalUnsignedContract {
	return m.BACnetOptionalUnsignedContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOptionalUnsignedNull) GetNullValue() BACnetApplicationTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetOptionalUnsignedNull(structType any) BACnetOptionalUnsignedNull {
	if casted, ok := structType.(BACnetOptionalUnsignedNull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOptionalUnsignedNull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOptionalUnsignedNull) GetTypeName() string {
	return "BACnetOptionalUnsignedNull"
}

func (m *_BACnetOptionalUnsignedNull) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetOptionalUnsignedContract.(*_BACnetOptionalUnsigned).getLengthInBits(ctx))

	// Simple field (nullValue)
	lengthInBits += m.NullValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetOptionalUnsignedNull) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetOptionalUnsignedNull) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetOptionalUnsigned) (__bACnetOptionalUnsignedNull BACnetOptionalUnsignedNull, err error) {
	m.BACnetOptionalUnsignedContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalUnsignedNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalUnsignedNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nullValue, err := ReadSimpleField[BACnetApplicationTagNull](ctx, "nullValue", ReadComplex[BACnetApplicationTagNull](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagNull](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nullValue' field"))
	}
	m.NullValue = nullValue

	if closeErr := readBuffer.CloseContext("BACnetOptionalUnsignedNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalUnsignedNull")
	}

	return m, nil
}

func (m *_BACnetOptionalUnsignedNull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetOptionalUnsignedNull) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetOptionalUnsignedNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetOptionalUnsignedNull")
		}

		if err := WriteSimpleField[BACnetApplicationTagNull](ctx, "nullValue", m.GetNullValue(), WriteComplex[BACnetApplicationTagNull](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nullValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetOptionalUnsignedNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetOptionalUnsignedNull")
		}
		return nil
	}
	return m.BACnetOptionalUnsignedContract.(*_BACnetOptionalUnsigned).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetOptionalUnsignedNull) IsBACnetOptionalUnsignedNull() {}

func (m *_BACnetOptionalUnsignedNull) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetOptionalUnsignedNull) deepCopy() *_BACnetOptionalUnsignedNull {
	if m == nil {
		return nil
	}
	_BACnetOptionalUnsignedNullCopy := &_BACnetOptionalUnsignedNull{
		m.BACnetOptionalUnsignedContract.(*_BACnetOptionalUnsigned).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagNull](m.NullValue),
	}
	_BACnetOptionalUnsignedNullCopy.BACnetOptionalUnsignedContract.(*_BACnetOptionalUnsigned)._SubType = m
	return _BACnetOptionalUnsignedNullCopy
}

func (m *_BACnetOptionalUnsignedNull) String() string {
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