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

// BACnetChannelValueNull is the corresponding interface of BACnetChannelValueNull
type BACnetChannelValueNull interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetChannelValue
	// GetNullValue returns NullValue (property field)
	GetNullValue() BACnetApplicationTagNull
	// IsBACnetChannelValueNull is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetChannelValueNull()
	// CreateBuilder creates a BACnetChannelValueNullBuilder
	CreateBACnetChannelValueNullBuilder() BACnetChannelValueNullBuilder
}

// _BACnetChannelValueNull is the data-structure of this message
type _BACnetChannelValueNull struct {
	BACnetChannelValueContract
	NullValue BACnetApplicationTagNull
}

var _ BACnetChannelValueNull = (*_BACnetChannelValueNull)(nil)
var _ BACnetChannelValueRequirements = (*_BACnetChannelValueNull)(nil)

// NewBACnetChannelValueNull factory function for _BACnetChannelValueNull
func NewBACnetChannelValueNull(peekedTagHeader BACnetTagHeader, nullValue BACnetApplicationTagNull) *_BACnetChannelValueNull {
	if nullValue == nil {
		panic("nullValue of type BACnetApplicationTagNull for BACnetChannelValueNull must not be nil")
	}
	_result := &_BACnetChannelValueNull{
		BACnetChannelValueContract: NewBACnetChannelValue(peekedTagHeader),
		NullValue:                  nullValue,
	}
	_result.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetChannelValueNullBuilder is a builder for BACnetChannelValueNull
type BACnetChannelValueNullBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(nullValue BACnetApplicationTagNull) BACnetChannelValueNullBuilder
	// WithNullValue adds NullValue (property field)
	WithNullValue(BACnetApplicationTagNull) BACnetChannelValueNullBuilder
	// WithNullValueBuilder adds NullValue (property field) which is build by the builder
	WithNullValueBuilder(func(BACnetApplicationTagNullBuilder) BACnetApplicationTagNullBuilder) BACnetChannelValueNullBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetChannelValueBuilder
	// Build builds the BACnetChannelValueNull or returns an error if something is wrong
	Build() (BACnetChannelValueNull, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetChannelValueNull
}

// NewBACnetChannelValueNullBuilder() creates a BACnetChannelValueNullBuilder
func NewBACnetChannelValueNullBuilder() BACnetChannelValueNullBuilder {
	return &_BACnetChannelValueNullBuilder{_BACnetChannelValueNull: new(_BACnetChannelValueNull)}
}

type _BACnetChannelValueNullBuilder struct {
	*_BACnetChannelValueNull

	parentBuilder *_BACnetChannelValueBuilder

	err *utils.MultiError
}

var _ (BACnetChannelValueNullBuilder) = (*_BACnetChannelValueNullBuilder)(nil)

func (b *_BACnetChannelValueNullBuilder) setParent(contract BACnetChannelValueContract) {
	b.BACnetChannelValueContract = contract
	contract.(*_BACnetChannelValue)._SubType = b._BACnetChannelValueNull
}

func (b *_BACnetChannelValueNullBuilder) WithMandatoryFields(nullValue BACnetApplicationTagNull) BACnetChannelValueNullBuilder {
	return b.WithNullValue(nullValue)
}

func (b *_BACnetChannelValueNullBuilder) WithNullValue(nullValue BACnetApplicationTagNull) BACnetChannelValueNullBuilder {
	b.NullValue = nullValue
	return b
}

func (b *_BACnetChannelValueNullBuilder) WithNullValueBuilder(builderSupplier func(BACnetApplicationTagNullBuilder) BACnetApplicationTagNullBuilder) BACnetChannelValueNullBuilder {
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

func (b *_BACnetChannelValueNullBuilder) Build() (BACnetChannelValueNull, error) {
	if b.NullValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'nullValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetChannelValueNull.deepCopy(), nil
}

func (b *_BACnetChannelValueNullBuilder) MustBuild() BACnetChannelValueNull {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetChannelValueNullBuilder) Done() BACnetChannelValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetChannelValueBuilder().(*_BACnetChannelValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetChannelValueNullBuilder) buildForBACnetChannelValue() (BACnetChannelValue, error) {
	return b.Build()
}

func (b *_BACnetChannelValueNullBuilder) DeepCopy() any {
	_copy := b.CreateBACnetChannelValueNullBuilder().(*_BACnetChannelValueNullBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetChannelValueNullBuilder creates a BACnetChannelValueNullBuilder
func (b *_BACnetChannelValueNull) CreateBACnetChannelValueNullBuilder() BACnetChannelValueNullBuilder {
	if b == nil {
		return NewBACnetChannelValueNullBuilder()
	}
	return &_BACnetChannelValueNullBuilder{_BACnetChannelValueNull: b.deepCopy()}
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

func (m *_BACnetChannelValueNull) GetParent() BACnetChannelValueContract {
	return m.BACnetChannelValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueNull) GetNullValue() BACnetApplicationTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueNull(structType any) BACnetChannelValueNull {
	if casted, ok := structType.(BACnetChannelValueNull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueNull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueNull) GetTypeName() string {
	return "BACnetChannelValueNull"
}

func (m *_BACnetChannelValueNull) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetChannelValueContract.(*_BACnetChannelValue).getLengthInBits(ctx))

	// Simple field (nullValue)
	lengthInBits += m.NullValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetChannelValueNull) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetChannelValueNull) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetChannelValue) (__bACnetChannelValueNull BACnetChannelValueNull, err error) {
	m.BACnetChannelValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nullValue, err := ReadSimpleField[BACnetApplicationTagNull](ctx, "nullValue", ReadComplex[BACnetApplicationTagNull](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagNull](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nullValue' field"))
	}
	m.NullValue = nullValue

	if closeErr := readBuffer.CloseContext("BACnetChannelValueNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueNull")
	}

	return m, nil
}

func (m *_BACnetChannelValueNull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueNull) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueNull")
		}

		if err := WriteSimpleField[BACnetApplicationTagNull](ctx, "nullValue", m.GetNullValue(), WriteComplex[BACnetApplicationTagNull](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nullValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueNull")
		}
		return nil
	}
	return m.BACnetChannelValueContract.(*_BACnetChannelValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetChannelValueNull) IsBACnetChannelValueNull() {}

func (m *_BACnetChannelValueNull) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetChannelValueNull) deepCopy() *_BACnetChannelValueNull {
	if m == nil {
		return nil
	}
	_BACnetChannelValueNullCopy := &_BACnetChannelValueNull{
		m.BACnetChannelValueContract.(*_BACnetChannelValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagNull](m.NullValue),
	}
	_BACnetChannelValueNullCopy.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = m
	return _BACnetChannelValueNullCopy
}

func (m *_BACnetChannelValueNull) String() string {
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