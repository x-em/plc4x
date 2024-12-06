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

// BACnetConstructedDataMaxInfoFrames is the corresponding interface of BACnetConstructedDataMaxInfoFrames
type BACnetConstructedDataMaxInfoFrames interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMaxInfoFrames returns MaxInfoFrames (property field)
	GetMaxInfoFrames() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataMaxInfoFrames is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMaxInfoFrames()
	// CreateBuilder creates a BACnetConstructedDataMaxInfoFramesBuilder
	CreateBACnetConstructedDataMaxInfoFramesBuilder() BACnetConstructedDataMaxInfoFramesBuilder
}

// _BACnetConstructedDataMaxInfoFrames is the data-structure of this message
type _BACnetConstructedDataMaxInfoFrames struct {
	BACnetConstructedDataContract
	MaxInfoFrames BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataMaxInfoFrames = (*_BACnetConstructedDataMaxInfoFrames)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMaxInfoFrames)(nil)

// NewBACnetConstructedDataMaxInfoFrames factory function for _BACnetConstructedDataMaxInfoFrames
func NewBACnetConstructedDataMaxInfoFrames(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, maxInfoFrames BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMaxInfoFrames {
	if maxInfoFrames == nil {
		panic("maxInfoFrames of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataMaxInfoFrames must not be nil")
	}
	_result := &_BACnetConstructedDataMaxInfoFrames{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MaxInfoFrames:                 maxInfoFrames,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataMaxInfoFramesBuilder is a builder for BACnetConstructedDataMaxInfoFrames
type BACnetConstructedDataMaxInfoFramesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(maxInfoFrames BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMaxInfoFramesBuilder
	// WithMaxInfoFrames adds MaxInfoFrames (property field)
	WithMaxInfoFrames(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMaxInfoFramesBuilder
	// WithMaxInfoFramesBuilder adds MaxInfoFrames (property field) which is build by the builder
	WithMaxInfoFramesBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataMaxInfoFramesBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataMaxInfoFrames or returns an error if something is wrong
	Build() (BACnetConstructedDataMaxInfoFrames, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataMaxInfoFrames
}

// NewBACnetConstructedDataMaxInfoFramesBuilder() creates a BACnetConstructedDataMaxInfoFramesBuilder
func NewBACnetConstructedDataMaxInfoFramesBuilder() BACnetConstructedDataMaxInfoFramesBuilder {
	return &_BACnetConstructedDataMaxInfoFramesBuilder{_BACnetConstructedDataMaxInfoFrames: new(_BACnetConstructedDataMaxInfoFrames)}
}

type _BACnetConstructedDataMaxInfoFramesBuilder struct {
	*_BACnetConstructedDataMaxInfoFrames

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataMaxInfoFramesBuilder) = (*_BACnetConstructedDataMaxInfoFramesBuilder)(nil)

func (b *_BACnetConstructedDataMaxInfoFramesBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataMaxInfoFrames
}

func (b *_BACnetConstructedDataMaxInfoFramesBuilder) WithMandatoryFields(maxInfoFrames BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMaxInfoFramesBuilder {
	return b.WithMaxInfoFrames(maxInfoFrames)
}

func (b *_BACnetConstructedDataMaxInfoFramesBuilder) WithMaxInfoFrames(maxInfoFrames BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMaxInfoFramesBuilder {
	b.MaxInfoFrames = maxInfoFrames
	return b
}

func (b *_BACnetConstructedDataMaxInfoFramesBuilder) WithMaxInfoFramesBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataMaxInfoFramesBuilder {
	builder := builderSupplier(b.MaxInfoFrames.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.MaxInfoFrames, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataMaxInfoFramesBuilder) Build() (BACnetConstructedDataMaxInfoFrames, error) {
	if b.MaxInfoFrames == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'maxInfoFrames' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataMaxInfoFrames.deepCopy(), nil
}

func (b *_BACnetConstructedDataMaxInfoFramesBuilder) MustBuild() BACnetConstructedDataMaxInfoFrames {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataMaxInfoFramesBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataMaxInfoFramesBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataMaxInfoFramesBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataMaxInfoFramesBuilder().(*_BACnetConstructedDataMaxInfoFramesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataMaxInfoFramesBuilder creates a BACnetConstructedDataMaxInfoFramesBuilder
func (b *_BACnetConstructedDataMaxInfoFrames) CreateBACnetConstructedDataMaxInfoFramesBuilder() BACnetConstructedDataMaxInfoFramesBuilder {
	if b == nil {
		return NewBACnetConstructedDataMaxInfoFramesBuilder()
	}
	return &_BACnetConstructedDataMaxInfoFramesBuilder{_BACnetConstructedDataMaxInfoFrames: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMaxInfoFrames) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMaxInfoFrames) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_INFO_FRAMES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMaxInfoFrames) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMaxInfoFrames) GetMaxInfoFrames() BACnetApplicationTagUnsignedInteger {
	return m.MaxInfoFrames
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMaxInfoFrames) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetMaxInfoFrames())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMaxInfoFrames(structType any) BACnetConstructedDataMaxInfoFrames {
	if casted, ok := structType.(BACnetConstructedDataMaxInfoFrames); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaxInfoFrames); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMaxInfoFrames) GetTypeName() string {
	return "BACnetConstructedDataMaxInfoFrames"
}

func (m *_BACnetConstructedDataMaxInfoFrames) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (maxInfoFrames)
	lengthInBits += m.MaxInfoFrames.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMaxInfoFrames) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMaxInfoFrames) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMaxInfoFrames BACnetConstructedDataMaxInfoFrames, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaxInfoFrames"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaxInfoFrames")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	maxInfoFrames, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "maxInfoFrames", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxInfoFrames' field"))
	}
	m.MaxInfoFrames = maxInfoFrames

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), maxInfoFrames)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaxInfoFrames"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaxInfoFrames")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMaxInfoFrames) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMaxInfoFrames) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaxInfoFrames"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaxInfoFrames")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "maxInfoFrames", m.GetMaxInfoFrames(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxInfoFrames' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaxInfoFrames"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaxInfoFrames")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMaxInfoFrames) IsBACnetConstructedDataMaxInfoFrames() {}

func (m *_BACnetConstructedDataMaxInfoFrames) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataMaxInfoFrames) deepCopy() *_BACnetConstructedDataMaxInfoFrames {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataMaxInfoFramesCopy := &_BACnetConstructedDataMaxInfoFrames{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.MaxInfoFrames),
	}
	_BACnetConstructedDataMaxInfoFramesCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataMaxInfoFramesCopy
}

func (m *_BACnetConstructedDataMaxInfoFrames) String() string {
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