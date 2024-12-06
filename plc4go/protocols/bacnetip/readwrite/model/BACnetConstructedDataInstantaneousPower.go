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

// BACnetConstructedDataInstantaneousPower is the corresponding interface of BACnetConstructedDataInstantaneousPower
type BACnetConstructedDataInstantaneousPower interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetInstantaneousPower returns InstantaneousPower (property field)
	GetInstantaneousPower() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataInstantaneousPower is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataInstantaneousPower()
	// CreateBuilder creates a BACnetConstructedDataInstantaneousPowerBuilder
	CreateBACnetConstructedDataInstantaneousPowerBuilder() BACnetConstructedDataInstantaneousPowerBuilder
}

// _BACnetConstructedDataInstantaneousPower is the data-structure of this message
type _BACnetConstructedDataInstantaneousPower struct {
	BACnetConstructedDataContract
	InstantaneousPower BACnetApplicationTagReal
}

var _ BACnetConstructedDataInstantaneousPower = (*_BACnetConstructedDataInstantaneousPower)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataInstantaneousPower)(nil)

// NewBACnetConstructedDataInstantaneousPower factory function for _BACnetConstructedDataInstantaneousPower
func NewBACnetConstructedDataInstantaneousPower(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, instantaneousPower BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataInstantaneousPower {
	if instantaneousPower == nil {
		panic("instantaneousPower of type BACnetApplicationTagReal for BACnetConstructedDataInstantaneousPower must not be nil")
	}
	_result := &_BACnetConstructedDataInstantaneousPower{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		InstantaneousPower:            instantaneousPower,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataInstantaneousPowerBuilder is a builder for BACnetConstructedDataInstantaneousPower
type BACnetConstructedDataInstantaneousPowerBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(instantaneousPower BACnetApplicationTagReal) BACnetConstructedDataInstantaneousPowerBuilder
	// WithInstantaneousPower adds InstantaneousPower (property field)
	WithInstantaneousPower(BACnetApplicationTagReal) BACnetConstructedDataInstantaneousPowerBuilder
	// WithInstantaneousPowerBuilder adds InstantaneousPower (property field) which is build by the builder
	WithInstantaneousPowerBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataInstantaneousPowerBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataInstantaneousPower or returns an error if something is wrong
	Build() (BACnetConstructedDataInstantaneousPower, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataInstantaneousPower
}

// NewBACnetConstructedDataInstantaneousPowerBuilder() creates a BACnetConstructedDataInstantaneousPowerBuilder
func NewBACnetConstructedDataInstantaneousPowerBuilder() BACnetConstructedDataInstantaneousPowerBuilder {
	return &_BACnetConstructedDataInstantaneousPowerBuilder{_BACnetConstructedDataInstantaneousPower: new(_BACnetConstructedDataInstantaneousPower)}
}

type _BACnetConstructedDataInstantaneousPowerBuilder struct {
	*_BACnetConstructedDataInstantaneousPower

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataInstantaneousPowerBuilder) = (*_BACnetConstructedDataInstantaneousPowerBuilder)(nil)

func (b *_BACnetConstructedDataInstantaneousPowerBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataInstantaneousPower
}

func (b *_BACnetConstructedDataInstantaneousPowerBuilder) WithMandatoryFields(instantaneousPower BACnetApplicationTagReal) BACnetConstructedDataInstantaneousPowerBuilder {
	return b.WithInstantaneousPower(instantaneousPower)
}

func (b *_BACnetConstructedDataInstantaneousPowerBuilder) WithInstantaneousPower(instantaneousPower BACnetApplicationTagReal) BACnetConstructedDataInstantaneousPowerBuilder {
	b.InstantaneousPower = instantaneousPower
	return b
}

func (b *_BACnetConstructedDataInstantaneousPowerBuilder) WithInstantaneousPowerBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataInstantaneousPowerBuilder {
	builder := builderSupplier(b.InstantaneousPower.CreateBACnetApplicationTagRealBuilder())
	var err error
	b.InstantaneousPower, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataInstantaneousPowerBuilder) Build() (BACnetConstructedDataInstantaneousPower, error) {
	if b.InstantaneousPower == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'instantaneousPower' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataInstantaneousPower.deepCopy(), nil
}

func (b *_BACnetConstructedDataInstantaneousPowerBuilder) MustBuild() BACnetConstructedDataInstantaneousPower {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataInstantaneousPowerBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataInstantaneousPowerBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataInstantaneousPowerBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataInstantaneousPowerBuilder().(*_BACnetConstructedDataInstantaneousPowerBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataInstantaneousPowerBuilder creates a BACnetConstructedDataInstantaneousPowerBuilder
func (b *_BACnetConstructedDataInstantaneousPower) CreateBACnetConstructedDataInstantaneousPowerBuilder() BACnetConstructedDataInstantaneousPowerBuilder {
	if b == nil {
		return NewBACnetConstructedDataInstantaneousPowerBuilder()
	}
	return &_BACnetConstructedDataInstantaneousPowerBuilder{_BACnetConstructedDataInstantaneousPower: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataInstantaneousPower) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataInstantaneousPower) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INSTANTANEOUS_POWER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataInstantaneousPower) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataInstantaneousPower) GetInstantaneousPower() BACnetApplicationTagReal {
	return m.InstantaneousPower
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataInstantaneousPower) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetInstantaneousPower())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataInstantaneousPower(structType any) BACnetConstructedDataInstantaneousPower {
	if casted, ok := structType.(BACnetConstructedDataInstantaneousPower); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataInstantaneousPower); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataInstantaneousPower) GetTypeName() string {
	return "BACnetConstructedDataInstantaneousPower"
}

func (m *_BACnetConstructedDataInstantaneousPower) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (instantaneousPower)
	lengthInBits += m.InstantaneousPower.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataInstantaneousPower) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataInstantaneousPower) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataInstantaneousPower BACnetConstructedDataInstantaneousPower, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataInstantaneousPower"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataInstantaneousPower")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	instantaneousPower, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "instantaneousPower", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'instantaneousPower' field"))
	}
	m.InstantaneousPower = instantaneousPower

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), instantaneousPower)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataInstantaneousPower"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataInstantaneousPower")
	}

	return m, nil
}

func (m *_BACnetConstructedDataInstantaneousPower) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataInstantaneousPower) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataInstantaneousPower"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataInstantaneousPower")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "instantaneousPower", m.GetInstantaneousPower(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'instantaneousPower' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataInstantaneousPower"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataInstantaneousPower")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataInstantaneousPower) IsBACnetConstructedDataInstantaneousPower() {}

func (m *_BACnetConstructedDataInstantaneousPower) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataInstantaneousPower) deepCopy() *_BACnetConstructedDataInstantaneousPower {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataInstantaneousPowerCopy := &_BACnetConstructedDataInstantaneousPower{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagReal](m.InstantaneousPower),
	}
	_BACnetConstructedDataInstantaneousPowerCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataInstantaneousPowerCopy
}

func (m *_BACnetConstructedDataInstantaneousPower) String() string {
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