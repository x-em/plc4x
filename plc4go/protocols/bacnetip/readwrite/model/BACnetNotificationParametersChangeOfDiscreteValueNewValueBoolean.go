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

// BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean
type BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetNotificationParametersChangeOfDiscreteValueNewValue
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() BACnetApplicationTagBoolean
	// IsBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean()
	// CreateBuilder creates a BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder
	CreateBACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder() BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder
}

// _BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean is the data-structure of this message
type _BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean struct {
	BACnetNotificationParametersChangeOfDiscreteValueNewValueContract
	BooleanValue BACnetApplicationTagBoolean
}

var _ BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean = (*_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean)(nil)
var _ BACnetNotificationParametersChangeOfDiscreteValueNewValueRequirements = (*_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean)(nil)

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean factory function for _BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, booleanValue BACnetApplicationTagBoolean, tagNumber uint8) *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean {
	if booleanValue == nil {
		panic("booleanValue of type BACnetApplicationTagBoolean for BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean must not be nil")
	}
	_result := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean{
		BACnetNotificationParametersChangeOfDiscreteValueNewValueContract: NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
		BooleanValue: booleanValue,
	}
	_result.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder is a builder for BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean
type BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(booleanValue BACnetApplicationTagBoolean) BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder
	// WithBooleanValue adds BooleanValue (property field)
	WithBooleanValue(BACnetApplicationTagBoolean) BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder
	// WithBooleanValueBuilder adds BooleanValue (property field) which is build by the builder
	WithBooleanValueBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder
	// Build builds the BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean or returns an error if something is wrong
	Build() (BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean
}

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder() creates a BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder() BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder {
	return &_BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder{_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean: new(_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean)}
}

type _BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder struct {
	*_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean

	parentBuilder *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder

	err *utils.MultiError
}

var _ (BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder) = (*_BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder)(nil)

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder) setParent(contract BACnetNotificationParametersChangeOfDiscreteValueNewValueContract) {
	b.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract = contract
	contract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue)._SubType = b._BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder) WithMandatoryFields(booleanValue BACnetApplicationTagBoolean) BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder {
	return b.WithBooleanValue(booleanValue)
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder) WithBooleanValue(booleanValue BACnetApplicationTagBoolean) BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder {
	b.BooleanValue = booleanValue
	return b
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder) WithBooleanValueBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder {
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

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder) Build() (BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean, error) {
	if b.BooleanValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'booleanValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean.deepCopy(), nil
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder) MustBuild() BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder) Done() BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder().(*_BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder) buildForBACnetNotificationParametersChangeOfDiscreteValueNewValue() (BACnetNotificationParametersChangeOfDiscreteValueNewValue, error) {
	return b.Build()
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder) DeepCopy() any {
	_copy := b.CreateBACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder().(*_BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder creates a BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder
func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) CreateBACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder() BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder {
	if b == nil {
		return NewBACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder()
	}
	return &_BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanBuilder{_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean: b.deepCopy()}
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

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetParent() BACnetNotificationParametersChangeOfDiscreteValueNewValueContract {
	return m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetBooleanValue() BACnetApplicationTagBoolean {
	return m.BooleanValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean(structType any) BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue).getLengthInBits(ctx))

	// Simple field (booleanValue)
	lengthInBits += m.BooleanValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParametersChangeOfDiscreteValueNewValue, tagNumber uint8) (__bACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean, err error) {
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	booleanValue, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "booleanValue", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'booleanValue' field"))
	}
	m.BooleanValue = booleanValue

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "booleanValue", m.GetBooleanValue(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'booleanValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean")
		}
		return nil
	}
	return m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) IsBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean() {
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) deepCopy() *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean {
	if m == nil {
		return nil
	}
	_BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanCopy := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean{
		m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBoolean](m.BooleanValue),
	}
	_BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanCopy.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue)._SubType = m
	return _BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanCopy
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) String() string {
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
