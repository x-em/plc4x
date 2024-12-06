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

// BACnetConstructedDataLargeAnalogValueResolution is the corresponding interface of BACnetConstructedDataLargeAnalogValueResolution
type BACnetConstructedDataLargeAnalogValueResolution interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetResolution returns Resolution (property field)
	GetResolution() BACnetApplicationTagDouble
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagDouble
	// IsBACnetConstructedDataLargeAnalogValueResolution is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLargeAnalogValueResolution()
	// CreateBuilder creates a BACnetConstructedDataLargeAnalogValueResolutionBuilder
	CreateBACnetConstructedDataLargeAnalogValueResolutionBuilder() BACnetConstructedDataLargeAnalogValueResolutionBuilder
}

// _BACnetConstructedDataLargeAnalogValueResolution is the data-structure of this message
type _BACnetConstructedDataLargeAnalogValueResolution struct {
	BACnetConstructedDataContract
	Resolution BACnetApplicationTagDouble
}

var _ BACnetConstructedDataLargeAnalogValueResolution = (*_BACnetConstructedDataLargeAnalogValueResolution)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLargeAnalogValueResolution)(nil)

// NewBACnetConstructedDataLargeAnalogValueResolution factory function for _BACnetConstructedDataLargeAnalogValueResolution
func NewBACnetConstructedDataLargeAnalogValueResolution(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, resolution BACnetApplicationTagDouble, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLargeAnalogValueResolution {
	if resolution == nil {
		panic("resolution of type BACnetApplicationTagDouble for BACnetConstructedDataLargeAnalogValueResolution must not be nil")
	}
	_result := &_BACnetConstructedDataLargeAnalogValueResolution{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Resolution:                    resolution,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLargeAnalogValueResolutionBuilder is a builder for BACnetConstructedDataLargeAnalogValueResolution
type BACnetConstructedDataLargeAnalogValueResolutionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(resolution BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueResolutionBuilder
	// WithResolution adds Resolution (property field)
	WithResolution(BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueResolutionBuilder
	// WithResolutionBuilder adds Resolution (property field) which is build by the builder
	WithResolutionBuilder(func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetConstructedDataLargeAnalogValueResolutionBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLargeAnalogValueResolution or returns an error if something is wrong
	Build() (BACnetConstructedDataLargeAnalogValueResolution, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLargeAnalogValueResolution
}

// NewBACnetConstructedDataLargeAnalogValueResolutionBuilder() creates a BACnetConstructedDataLargeAnalogValueResolutionBuilder
func NewBACnetConstructedDataLargeAnalogValueResolutionBuilder() BACnetConstructedDataLargeAnalogValueResolutionBuilder {
	return &_BACnetConstructedDataLargeAnalogValueResolutionBuilder{_BACnetConstructedDataLargeAnalogValueResolution: new(_BACnetConstructedDataLargeAnalogValueResolution)}
}

type _BACnetConstructedDataLargeAnalogValueResolutionBuilder struct {
	*_BACnetConstructedDataLargeAnalogValueResolution

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLargeAnalogValueResolutionBuilder) = (*_BACnetConstructedDataLargeAnalogValueResolutionBuilder)(nil)

func (b *_BACnetConstructedDataLargeAnalogValueResolutionBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataLargeAnalogValueResolution
}

func (b *_BACnetConstructedDataLargeAnalogValueResolutionBuilder) WithMandatoryFields(resolution BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueResolutionBuilder {
	return b.WithResolution(resolution)
}

func (b *_BACnetConstructedDataLargeAnalogValueResolutionBuilder) WithResolution(resolution BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueResolutionBuilder {
	b.Resolution = resolution
	return b
}

func (b *_BACnetConstructedDataLargeAnalogValueResolutionBuilder) WithResolutionBuilder(builderSupplier func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetConstructedDataLargeAnalogValueResolutionBuilder {
	builder := builderSupplier(b.Resolution.CreateBACnetApplicationTagDoubleBuilder())
	var err error
	b.Resolution, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagDoubleBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLargeAnalogValueResolutionBuilder) Build() (BACnetConstructedDataLargeAnalogValueResolution, error) {
	if b.Resolution == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'resolution' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLargeAnalogValueResolution.deepCopy(), nil
}

func (b *_BACnetConstructedDataLargeAnalogValueResolutionBuilder) MustBuild() BACnetConstructedDataLargeAnalogValueResolution {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLargeAnalogValueResolutionBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLargeAnalogValueResolutionBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLargeAnalogValueResolutionBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLargeAnalogValueResolutionBuilder().(*_BACnetConstructedDataLargeAnalogValueResolutionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLargeAnalogValueResolutionBuilder creates a BACnetConstructedDataLargeAnalogValueResolutionBuilder
func (b *_BACnetConstructedDataLargeAnalogValueResolution) CreateBACnetConstructedDataLargeAnalogValueResolutionBuilder() BACnetConstructedDataLargeAnalogValueResolutionBuilder {
	if b == nil {
		return NewBACnetConstructedDataLargeAnalogValueResolutionBuilder()
	}
	return &_BACnetConstructedDataLargeAnalogValueResolutionBuilder{_BACnetConstructedDataLargeAnalogValueResolution: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueResolution) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LARGE_ANALOG_VALUE
}

func (m *_BACnetConstructedDataLargeAnalogValueResolution) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RESOLUTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueResolution) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueResolution) GetResolution() BACnetApplicationTagDouble {
	return m.Resolution
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueResolution) GetActualValue() BACnetApplicationTagDouble {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagDouble(m.GetResolution())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLargeAnalogValueResolution(structType any) BACnetConstructedDataLargeAnalogValueResolution {
	if casted, ok := structType.(BACnetConstructedDataLargeAnalogValueResolution); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLargeAnalogValueResolution); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLargeAnalogValueResolution) GetTypeName() string {
	return "BACnetConstructedDataLargeAnalogValueResolution"
}

func (m *_BACnetConstructedDataLargeAnalogValueResolution) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (resolution)
	lengthInBits += m.Resolution.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLargeAnalogValueResolution) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLargeAnalogValueResolution) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLargeAnalogValueResolution BACnetConstructedDataLargeAnalogValueResolution, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLargeAnalogValueResolution"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLargeAnalogValueResolution")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	resolution, err := ReadSimpleField[BACnetApplicationTagDouble](ctx, "resolution", ReadComplex[BACnetApplicationTagDouble](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDouble](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'resolution' field"))
	}
	m.Resolution = resolution

	actualValue, err := ReadVirtualField[BACnetApplicationTagDouble](ctx, "actualValue", (*BACnetApplicationTagDouble)(nil), resolution)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLargeAnalogValueResolution"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLargeAnalogValueResolution")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLargeAnalogValueResolution) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLargeAnalogValueResolution) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLargeAnalogValueResolution"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLargeAnalogValueResolution")
		}

		if err := WriteSimpleField[BACnetApplicationTagDouble](ctx, "resolution", m.GetResolution(), WriteComplex[BACnetApplicationTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'resolution' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLargeAnalogValueResolution"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLargeAnalogValueResolution")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLargeAnalogValueResolution) IsBACnetConstructedDataLargeAnalogValueResolution() {
}

func (m *_BACnetConstructedDataLargeAnalogValueResolution) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLargeAnalogValueResolution) deepCopy() *_BACnetConstructedDataLargeAnalogValueResolution {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLargeAnalogValueResolutionCopy := &_BACnetConstructedDataLargeAnalogValueResolution{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagDouble](m.Resolution),
	}
	_BACnetConstructedDataLargeAnalogValueResolutionCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLargeAnalogValueResolutionCopy
}

func (m *_BACnetConstructedDataLargeAnalogValueResolution) String() string {
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
