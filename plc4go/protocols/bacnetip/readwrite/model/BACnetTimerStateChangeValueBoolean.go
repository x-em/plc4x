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

// BACnetTimerStateChangeValueBoolean is the corresponding interface of BACnetTimerStateChangeValueBoolean
type BACnetTimerStateChangeValueBoolean interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetTimerStateChangeValue
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() BACnetApplicationTagBoolean
	// IsBACnetTimerStateChangeValueBoolean is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimerStateChangeValueBoolean()
	// CreateBuilder creates a BACnetTimerStateChangeValueBooleanBuilder
	CreateBACnetTimerStateChangeValueBooleanBuilder() BACnetTimerStateChangeValueBooleanBuilder
}

// _BACnetTimerStateChangeValueBoolean is the data-structure of this message
type _BACnetTimerStateChangeValueBoolean struct {
	BACnetTimerStateChangeValueContract
	BooleanValue BACnetApplicationTagBoolean
}

var _ BACnetTimerStateChangeValueBoolean = (*_BACnetTimerStateChangeValueBoolean)(nil)
var _ BACnetTimerStateChangeValueRequirements = (*_BACnetTimerStateChangeValueBoolean)(nil)

// NewBACnetTimerStateChangeValueBoolean factory function for _BACnetTimerStateChangeValueBoolean
func NewBACnetTimerStateChangeValueBoolean(peekedTagHeader BACnetTagHeader, booleanValue BACnetApplicationTagBoolean, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueBoolean {
	if booleanValue == nil {
		panic("booleanValue of type BACnetApplicationTagBoolean for BACnetTimerStateChangeValueBoolean must not be nil")
	}
	_result := &_BACnetTimerStateChangeValueBoolean{
		BACnetTimerStateChangeValueContract: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
		BooleanValue:                        booleanValue,
	}
	_result.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTimerStateChangeValueBooleanBuilder is a builder for BACnetTimerStateChangeValueBoolean
type BACnetTimerStateChangeValueBooleanBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(booleanValue BACnetApplicationTagBoolean) BACnetTimerStateChangeValueBooleanBuilder
	// WithBooleanValue adds BooleanValue (property field)
	WithBooleanValue(BACnetApplicationTagBoolean) BACnetTimerStateChangeValueBooleanBuilder
	// WithBooleanValueBuilder adds BooleanValue (property field) which is build by the builder
	WithBooleanValueBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetTimerStateChangeValueBooleanBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetTimerStateChangeValueBuilder
	// Build builds the BACnetTimerStateChangeValueBoolean or returns an error if something is wrong
	Build() (BACnetTimerStateChangeValueBoolean, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTimerStateChangeValueBoolean
}

// NewBACnetTimerStateChangeValueBooleanBuilder() creates a BACnetTimerStateChangeValueBooleanBuilder
func NewBACnetTimerStateChangeValueBooleanBuilder() BACnetTimerStateChangeValueBooleanBuilder {
	return &_BACnetTimerStateChangeValueBooleanBuilder{_BACnetTimerStateChangeValueBoolean: new(_BACnetTimerStateChangeValueBoolean)}
}

type _BACnetTimerStateChangeValueBooleanBuilder struct {
	*_BACnetTimerStateChangeValueBoolean

	parentBuilder *_BACnetTimerStateChangeValueBuilder

	err *utils.MultiError
}

var _ (BACnetTimerStateChangeValueBooleanBuilder) = (*_BACnetTimerStateChangeValueBooleanBuilder)(nil)

func (b *_BACnetTimerStateChangeValueBooleanBuilder) setParent(contract BACnetTimerStateChangeValueContract) {
	b.BACnetTimerStateChangeValueContract = contract
	contract.(*_BACnetTimerStateChangeValue)._SubType = b._BACnetTimerStateChangeValueBoolean
}

func (b *_BACnetTimerStateChangeValueBooleanBuilder) WithMandatoryFields(booleanValue BACnetApplicationTagBoolean) BACnetTimerStateChangeValueBooleanBuilder {
	return b.WithBooleanValue(booleanValue)
}

func (b *_BACnetTimerStateChangeValueBooleanBuilder) WithBooleanValue(booleanValue BACnetApplicationTagBoolean) BACnetTimerStateChangeValueBooleanBuilder {
	b.BooleanValue = booleanValue
	return b
}

func (b *_BACnetTimerStateChangeValueBooleanBuilder) WithBooleanValueBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetTimerStateChangeValueBooleanBuilder {
	builder := builderSupplier(b.BooleanValue.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.BooleanValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetTimerStateChangeValueBooleanBuilder) Build() (BACnetTimerStateChangeValueBoolean, error) {
	if b.BooleanValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'booleanValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetTimerStateChangeValueBoolean.deepCopy(), nil
}

func (b *_BACnetTimerStateChangeValueBooleanBuilder) MustBuild() BACnetTimerStateChangeValueBoolean {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetTimerStateChangeValueBooleanBuilder) Done() BACnetTimerStateChangeValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetTimerStateChangeValueBuilder().(*_BACnetTimerStateChangeValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetTimerStateChangeValueBooleanBuilder) buildForBACnetTimerStateChangeValue() (BACnetTimerStateChangeValue, error) {
	return b.Build()
}

func (b *_BACnetTimerStateChangeValueBooleanBuilder) DeepCopy() any {
	_copy := b.CreateBACnetTimerStateChangeValueBooleanBuilder().(*_BACnetTimerStateChangeValueBooleanBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetTimerStateChangeValueBooleanBuilder creates a BACnetTimerStateChangeValueBooleanBuilder
func (b *_BACnetTimerStateChangeValueBoolean) CreateBACnetTimerStateChangeValueBooleanBuilder() BACnetTimerStateChangeValueBooleanBuilder {
	if b == nil {
		return NewBACnetTimerStateChangeValueBooleanBuilder()
	}
	return &_BACnetTimerStateChangeValueBooleanBuilder{_BACnetTimerStateChangeValueBoolean: b.deepCopy()}
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

func (m *_BACnetTimerStateChangeValueBoolean) GetParent() BACnetTimerStateChangeValueContract {
	return m.BACnetTimerStateChangeValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueBoolean) GetBooleanValue() BACnetApplicationTagBoolean {
	return m.BooleanValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueBoolean(structType any) BACnetTimerStateChangeValueBoolean {
	if casted, ok := structType.(BACnetTimerStateChangeValueBoolean); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueBoolean); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueBoolean) GetTypeName() string {
	return "BACnetTimerStateChangeValueBoolean"
}

func (m *_BACnetTimerStateChangeValueBoolean) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).getLengthInBits(ctx))

	// Simple field (booleanValue)
	lengthInBits += m.BooleanValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueBoolean) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetTimerStateChangeValueBoolean) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetTimerStateChangeValue, objectTypeArgument BACnetObjectType) (__bACnetTimerStateChangeValueBoolean BACnetTimerStateChangeValueBoolean, err error) {
	m.BACnetTimerStateChangeValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueBoolean"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueBoolean")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	booleanValue, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "booleanValue", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'booleanValue' field"))
	}
	m.BooleanValue = booleanValue

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueBoolean"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueBoolean")
	}

	return m, nil
}

func (m *_BACnetTimerStateChangeValueBoolean) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerStateChangeValueBoolean) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueBoolean"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueBoolean")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "booleanValue", m.GetBooleanValue(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'booleanValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueBoolean"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueBoolean")
		}
		return nil
	}
	return m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueBoolean) IsBACnetTimerStateChangeValueBoolean() {}

func (m *_BACnetTimerStateChangeValueBoolean) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTimerStateChangeValueBoolean) deepCopy() *_BACnetTimerStateChangeValueBoolean {
	if m == nil {
		return nil
	}
	_BACnetTimerStateChangeValueBooleanCopy := &_BACnetTimerStateChangeValueBoolean{
		m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBoolean](m.BooleanValue),
	}
	_BACnetTimerStateChangeValueBooleanCopy.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = m
	return _BACnetTimerStateChangeValueBooleanCopy
}

func (m *_BACnetTimerStateChangeValueBoolean) String() string {
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