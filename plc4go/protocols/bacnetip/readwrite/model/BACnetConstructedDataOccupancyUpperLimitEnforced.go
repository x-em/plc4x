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

// BACnetConstructedDataOccupancyUpperLimitEnforced is the corresponding interface of BACnetConstructedDataOccupancyUpperLimitEnforced
type BACnetConstructedDataOccupancyUpperLimitEnforced interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetOccupancyUpperLimitEnforced returns OccupancyUpperLimitEnforced (property field)
	GetOccupancyUpperLimitEnforced() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataOccupancyUpperLimitEnforced is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataOccupancyUpperLimitEnforced()
	// CreateBuilder creates a BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder
	CreateBACnetConstructedDataOccupancyUpperLimitEnforcedBuilder() BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder
}

// _BACnetConstructedDataOccupancyUpperLimitEnforced is the data-structure of this message
type _BACnetConstructedDataOccupancyUpperLimitEnforced struct {
	BACnetConstructedDataContract
	OccupancyUpperLimitEnforced BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataOccupancyUpperLimitEnforced = (*_BACnetConstructedDataOccupancyUpperLimitEnforced)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataOccupancyUpperLimitEnforced)(nil)

// NewBACnetConstructedDataOccupancyUpperLimitEnforced factory function for _BACnetConstructedDataOccupancyUpperLimitEnforced
func NewBACnetConstructedDataOccupancyUpperLimitEnforced(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, occupancyUpperLimitEnforced BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOccupancyUpperLimitEnforced {
	if occupancyUpperLimitEnforced == nil {
		panic("occupancyUpperLimitEnforced of type BACnetApplicationTagBoolean for BACnetConstructedDataOccupancyUpperLimitEnforced must not be nil")
	}
	_result := &_BACnetConstructedDataOccupancyUpperLimitEnforced{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		OccupancyUpperLimitEnforced:   occupancyUpperLimitEnforced,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder is a builder for BACnetConstructedDataOccupancyUpperLimitEnforced
type BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(occupancyUpperLimitEnforced BACnetApplicationTagBoolean) BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder
	// WithOccupancyUpperLimitEnforced adds OccupancyUpperLimitEnforced (property field)
	WithOccupancyUpperLimitEnforced(BACnetApplicationTagBoolean) BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder
	// WithOccupancyUpperLimitEnforcedBuilder adds OccupancyUpperLimitEnforced (property field) which is build by the builder
	WithOccupancyUpperLimitEnforcedBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataOccupancyUpperLimitEnforced or returns an error if something is wrong
	Build() (BACnetConstructedDataOccupancyUpperLimitEnforced, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataOccupancyUpperLimitEnforced
}

// NewBACnetConstructedDataOccupancyUpperLimitEnforcedBuilder() creates a BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder
func NewBACnetConstructedDataOccupancyUpperLimitEnforcedBuilder() BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder {
	return &_BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder{_BACnetConstructedDataOccupancyUpperLimitEnforced: new(_BACnetConstructedDataOccupancyUpperLimitEnforced)}
}

type _BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder struct {
	*_BACnetConstructedDataOccupancyUpperLimitEnforced

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder) = (*_BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder)(nil)

func (b *_BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataOccupancyUpperLimitEnforced
}

func (b *_BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder) WithMandatoryFields(occupancyUpperLimitEnforced BACnetApplicationTagBoolean) BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder {
	return b.WithOccupancyUpperLimitEnforced(occupancyUpperLimitEnforced)
}

func (b *_BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder) WithOccupancyUpperLimitEnforced(occupancyUpperLimitEnforced BACnetApplicationTagBoolean) BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder {
	b.OccupancyUpperLimitEnforced = occupancyUpperLimitEnforced
	return b
}

func (b *_BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder) WithOccupancyUpperLimitEnforcedBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder {
	builder := builderSupplier(b.OccupancyUpperLimitEnforced.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.OccupancyUpperLimitEnforced, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder) Build() (BACnetConstructedDataOccupancyUpperLimitEnforced, error) {
	if b.OccupancyUpperLimitEnforced == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'occupancyUpperLimitEnforced' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataOccupancyUpperLimitEnforced.deepCopy(), nil
}

func (b *_BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder) MustBuild() BACnetConstructedDataOccupancyUpperLimitEnforced {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataOccupancyUpperLimitEnforcedBuilder().(*_BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataOccupancyUpperLimitEnforcedBuilder creates a BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder
func (b *_BACnetConstructedDataOccupancyUpperLimitEnforced) CreateBACnetConstructedDataOccupancyUpperLimitEnforcedBuilder() BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder {
	if b == nil {
		return NewBACnetConstructedDataOccupancyUpperLimitEnforcedBuilder()
	}
	return &_BACnetConstructedDataOccupancyUpperLimitEnforcedBuilder{_BACnetConstructedDataOccupancyUpperLimitEnforced: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OCCUPANCY_UPPER_LIMIT_ENFORCED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) GetOccupancyUpperLimitEnforced() BACnetApplicationTagBoolean {
	return m.OccupancyUpperLimitEnforced
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetOccupancyUpperLimitEnforced())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOccupancyUpperLimitEnforced(structType any) BACnetConstructedDataOccupancyUpperLimitEnforced {
	if casted, ok := structType.(BACnetConstructedDataOccupancyUpperLimitEnforced); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOccupancyUpperLimitEnforced); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) GetTypeName() string {
	return "BACnetConstructedDataOccupancyUpperLimitEnforced"
}

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (occupancyUpperLimitEnforced)
	lengthInBits += m.OccupancyUpperLimitEnforced.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataOccupancyUpperLimitEnforced BACnetConstructedDataOccupancyUpperLimitEnforced, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOccupancyUpperLimitEnforced"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOccupancyUpperLimitEnforced")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	occupancyUpperLimitEnforced, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "occupancyUpperLimitEnforced", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'occupancyUpperLimitEnforced' field"))
	}
	m.OccupancyUpperLimitEnforced = occupancyUpperLimitEnforced

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), occupancyUpperLimitEnforced)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOccupancyUpperLimitEnforced"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOccupancyUpperLimitEnforced")
	}

	return m, nil
}

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOccupancyUpperLimitEnforced"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOccupancyUpperLimitEnforced")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "occupancyUpperLimitEnforced", m.GetOccupancyUpperLimitEnforced(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'occupancyUpperLimitEnforced' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOccupancyUpperLimitEnforced"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOccupancyUpperLimitEnforced")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) IsBACnetConstructedDataOccupancyUpperLimitEnforced() {
}

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) deepCopy() *_BACnetConstructedDataOccupancyUpperLimitEnforced {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataOccupancyUpperLimitEnforcedCopy := &_BACnetConstructedDataOccupancyUpperLimitEnforced{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBoolean](m.OccupancyUpperLimitEnforced),
	}
	_BACnetConstructedDataOccupancyUpperLimitEnforcedCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataOccupancyUpperLimitEnforcedCopy
}

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) String() string {
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