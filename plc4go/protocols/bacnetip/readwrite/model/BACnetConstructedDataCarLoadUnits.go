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

// BACnetConstructedDataCarLoadUnits is the corresponding interface of BACnetConstructedDataCarLoadUnits
type BACnetConstructedDataCarLoadUnits interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetUnits returns Units (property field)
	GetUnits() BACnetEngineeringUnitsTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEngineeringUnitsTagged
	// IsBACnetConstructedDataCarLoadUnits is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataCarLoadUnits()
	// CreateBuilder creates a BACnetConstructedDataCarLoadUnitsBuilder
	CreateBACnetConstructedDataCarLoadUnitsBuilder() BACnetConstructedDataCarLoadUnitsBuilder
}

// _BACnetConstructedDataCarLoadUnits is the data-structure of this message
type _BACnetConstructedDataCarLoadUnits struct {
	BACnetConstructedDataContract
	Units BACnetEngineeringUnitsTagged
}

var _ BACnetConstructedDataCarLoadUnits = (*_BACnetConstructedDataCarLoadUnits)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCarLoadUnits)(nil)

// NewBACnetConstructedDataCarLoadUnits factory function for _BACnetConstructedDataCarLoadUnits
func NewBACnetConstructedDataCarLoadUnits(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, units BACnetEngineeringUnitsTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCarLoadUnits {
	if units == nil {
		panic("units of type BACnetEngineeringUnitsTagged for BACnetConstructedDataCarLoadUnits must not be nil")
	}
	_result := &_BACnetConstructedDataCarLoadUnits{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Units:                         units,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataCarLoadUnitsBuilder is a builder for BACnetConstructedDataCarLoadUnits
type BACnetConstructedDataCarLoadUnitsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(units BACnetEngineeringUnitsTagged) BACnetConstructedDataCarLoadUnitsBuilder
	// WithUnits adds Units (property field)
	WithUnits(BACnetEngineeringUnitsTagged) BACnetConstructedDataCarLoadUnitsBuilder
	// WithUnitsBuilder adds Units (property field) which is build by the builder
	WithUnitsBuilder(func(BACnetEngineeringUnitsTaggedBuilder) BACnetEngineeringUnitsTaggedBuilder) BACnetConstructedDataCarLoadUnitsBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataCarLoadUnits or returns an error if something is wrong
	Build() (BACnetConstructedDataCarLoadUnits, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataCarLoadUnits
}

// NewBACnetConstructedDataCarLoadUnitsBuilder() creates a BACnetConstructedDataCarLoadUnitsBuilder
func NewBACnetConstructedDataCarLoadUnitsBuilder() BACnetConstructedDataCarLoadUnitsBuilder {
	return &_BACnetConstructedDataCarLoadUnitsBuilder{_BACnetConstructedDataCarLoadUnits: new(_BACnetConstructedDataCarLoadUnits)}
}

type _BACnetConstructedDataCarLoadUnitsBuilder struct {
	*_BACnetConstructedDataCarLoadUnits

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataCarLoadUnitsBuilder) = (*_BACnetConstructedDataCarLoadUnitsBuilder)(nil)

func (b *_BACnetConstructedDataCarLoadUnitsBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataCarLoadUnits
}

func (b *_BACnetConstructedDataCarLoadUnitsBuilder) WithMandatoryFields(units BACnetEngineeringUnitsTagged) BACnetConstructedDataCarLoadUnitsBuilder {
	return b.WithUnits(units)
}

func (b *_BACnetConstructedDataCarLoadUnitsBuilder) WithUnits(units BACnetEngineeringUnitsTagged) BACnetConstructedDataCarLoadUnitsBuilder {
	b.Units = units
	return b
}

func (b *_BACnetConstructedDataCarLoadUnitsBuilder) WithUnitsBuilder(builderSupplier func(BACnetEngineeringUnitsTaggedBuilder) BACnetEngineeringUnitsTaggedBuilder) BACnetConstructedDataCarLoadUnitsBuilder {
	builder := builderSupplier(b.Units.CreateBACnetEngineeringUnitsTaggedBuilder())
	var err error
	b.Units, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetEngineeringUnitsTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataCarLoadUnitsBuilder) Build() (BACnetConstructedDataCarLoadUnits, error) {
	if b.Units == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'units' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataCarLoadUnits.deepCopy(), nil
}

func (b *_BACnetConstructedDataCarLoadUnitsBuilder) MustBuild() BACnetConstructedDataCarLoadUnits {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataCarLoadUnitsBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataCarLoadUnitsBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataCarLoadUnitsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataCarLoadUnitsBuilder().(*_BACnetConstructedDataCarLoadUnitsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataCarLoadUnitsBuilder creates a BACnetConstructedDataCarLoadUnitsBuilder
func (b *_BACnetConstructedDataCarLoadUnits) CreateBACnetConstructedDataCarLoadUnitsBuilder() BACnetConstructedDataCarLoadUnitsBuilder {
	if b == nil {
		return NewBACnetConstructedDataCarLoadUnitsBuilder()
	}
	return &_BACnetConstructedDataCarLoadUnitsBuilder{_BACnetConstructedDataCarLoadUnits: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCarLoadUnits) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCarLoadUnits) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CAR_LOAD_UNITS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCarLoadUnits) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCarLoadUnits) GetUnits() BACnetEngineeringUnitsTagged {
	return m.Units
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCarLoadUnits) GetActualValue() BACnetEngineeringUnitsTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetEngineeringUnitsTagged(m.GetUnits())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCarLoadUnits(structType any) BACnetConstructedDataCarLoadUnits {
	if casted, ok := structType.(BACnetConstructedDataCarLoadUnits); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCarLoadUnits); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCarLoadUnits) GetTypeName() string {
	return "BACnetConstructedDataCarLoadUnits"
}

func (m *_BACnetConstructedDataCarLoadUnits) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (units)
	lengthInBits += m.Units.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCarLoadUnits) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCarLoadUnits) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCarLoadUnits BACnetConstructedDataCarLoadUnits, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCarLoadUnits"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCarLoadUnits")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	units, err := ReadSimpleField[BACnetEngineeringUnitsTagged](ctx, "units", ReadComplex[BACnetEngineeringUnitsTagged](BACnetEngineeringUnitsTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'units' field"))
	}
	m.Units = units

	actualValue, err := ReadVirtualField[BACnetEngineeringUnitsTagged](ctx, "actualValue", (*BACnetEngineeringUnitsTagged)(nil), units)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCarLoadUnits"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCarLoadUnits")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCarLoadUnits) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCarLoadUnits) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCarLoadUnits"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCarLoadUnits")
		}

		if err := WriteSimpleField[BACnetEngineeringUnitsTagged](ctx, "units", m.GetUnits(), WriteComplex[BACnetEngineeringUnitsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'units' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCarLoadUnits"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCarLoadUnits")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCarLoadUnits) IsBACnetConstructedDataCarLoadUnits() {}

func (m *_BACnetConstructedDataCarLoadUnits) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataCarLoadUnits) deepCopy() *_BACnetConstructedDataCarLoadUnits {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataCarLoadUnitsCopy := &_BACnetConstructedDataCarLoadUnits{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetEngineeringUnitsTagged](m.Units),
	}
	_BACnetConstructedDataCarLoadUnitsCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataCarLoadUnitsCopy
}

func (m *_BACnetConstructedDataCarLoadUnits) String() string {
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