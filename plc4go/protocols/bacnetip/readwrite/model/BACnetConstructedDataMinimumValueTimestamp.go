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

// BACnetConstructedDataMinimumValueTimestamp is the corresponding interface of BACnetConstructedDataMinimumValueTimestamp
type BACnetConstructedDataMinimumValueTimestamp interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMinimumValueTimestamp returns MinimumValueTimestamp (property field)
	GetMinimumValueTimestamp() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
	// IsBACnetConstructedDataMinimumValueTimestamp is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMinimumValueTimestamp()
	// CreateBuilder creates a BACnetConstructedDataMinimumValueTimestampBuilder
	CreateBACnetConstructedDataMinimumValueTimestampBuilder() BACnetConstructedDataMinimumValueTimestampBuilder
}

// _BACnetConstructedDataMinimumValueTimestamp is the data-structure of this message
type _BACnetConstructedDataMinimumValueTimestamp struct {
	BACnetConstructedDataContract
	MinimumValueTimestamp BACnetDateTime
}

var _ BACnetConstructedDataMinimumValueTimestamp = (*_BACnetConstructedDataMinimumValueTimestamp)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMinimumValueTimestamp)(nil)

// NewBACnetConstructedDataMinimumValueTimestamp factory function for _BACnetConstructedDataMinimumValueTimestamp
func NewBACnetConstructedDataMinimumValueTimestamp(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, minimumValueTimestamp BACnetDateTime, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMinimumValueTimestamp {
	if minimumValueTimestamp == nil {
		panic("minimumValueTimestamp of type BACnetDateTime for BACnetConstructedDataMinimumValueTimestamp must not be nil")
	}
	_result := &_BACnetConstructedDataMinimumValueTimestamp{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MinimumValueTimestamp:         minimumValueTimestamp,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataMinimumValueTimestampBuilder is a builder for BACnetConstructedDataMinimumValueTimestamp
type BACnetConstructedDataMinimumValueTimestampBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(minimumValueTimestamp BACnetDateTime) BACnetConstructedDataMinimumValueTimestampBuilder
	// WithMinimumValueTimestamp adds MinimumValueTimestamp (property field)
	WithMinimumValueTimestamp(BACnetDateTime) BACnetConstructedDataMinimumValueTimestampBuilder
	// WithMinimumValueTimestampBuilder adds MinimumValueTimestamp (property field) which is build by the builder
	WithMinimumValueTimestampBuilder(func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataMinimumValueTimestampBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataMinimumValueTimestamp or returns an error if something is wrong
	Build() (BACnetConstructedDataMinimumValueTimestamp, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataMinimumValueTimestamp
}

// NewBACnetConstructedDataMinimumValueTimestampBuilder() creates a BACnetConstructedDataMinimumValueTimestampBuilder
func NewBACnetConstructedDataMinimumValueTimestampBuilder() BACnetConstructedDataMinimumValueTimestampBuilder {
	return &_BACnetConstructedDataMinimumValueTimestampBuilder{_BACnetConstructedDataMinimumValueTimestamp: new(_BACnetConstructedDataMinimumValueTimestamp)}
}

type _BACnetConstructedDataMinimumValueTimestampBuilder struct {
	*_BACnetConstructedDataMinimumValueTimestamp

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataMinimumValueTimestampBuilder) = (*_BACnetConstructedDataMinimumValueTimestampBuilder)(nil)

func (b *_BACnetConstructedDataMinimumValueTimestampBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataMinimumValueTimestamp
}

func (b *_BACnetConstructedDataMinimumValueTimestampBuilder) WithMandatoryFields(minimumValueTimestamp BACnetDateTime) BACnetConstructedDataMinimumValueTimestampBuilder {
	return b.WithMinimumValueTimestamp(minimumValueTimestamp)
}

func (b *_BACnetConstructedDataMinimumValueTimestampBuilder) WithMinimumValueTimestamp(minimumValueTimestamp BACnetDateTime) BACnetConstructedDataMinimumValueTimestampBuilder {
	b.MinimumValueTimestamp = minimumValueTimestamp
	return b
}

func (b *_BACnetConstructedDataMinimumValueTimestampBuilder) WithMinimumValueTimestampBuilder(builderSupplier func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataMinimumValueTimestampBuilder {
	builder := builderSupplier(b.MinimumValueTimestamp.CreateBACnetDateTimeBuilder())
	var err error
	b.MinimumValueTimestamp, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDateTimeBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataMinimumValueTimestampBuilder) Build() (BACnetConstructedDataMinimumValueTimestamp, error) {
	if b.MinimumValueTimestamp == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'minimumValueTimestamp' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataMinimumValueTimestamp.deepCopy(), nil
}

func (b *_BACnetConstructedDataMinimumValueTimestampBuilder) MustBuild() BACnetConstructedDataMinimumValueTimestamp {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataMinimumValueTimestampBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataMinimumValueTimestampBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataMinimumValueTimestampBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataMinimumValueTimestampBuilder().(*_BACnetConstructedDataMinimumValueTimestampBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataMinimumValueTimestampBuilder creates a BACnetConstructedDataMinimumValueTimestampBuilder
func (b *_BACnetConstructedDataMinimumValueTimestamp) CreateBACnetConstructedDataMinimumValueTimestampBuilder() BACnetConstructedDataMinimumValueTimestampBuilder {
	if b == nil {
		return NewBACnetConstructedDataMinimumValueTimestampBuilder()
	}
	return &_BACnetConstructedDataMinimumValueTimestampBuilder{_BACnetConstructedDataMinimumValueTimestamp: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MINIMUM_VALUE_TIMESTAMP
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetMinimumValueTimestamp() BACnetDateTime {
	return m.MinimumValueTimestamp
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetActualValue() BACnetDateTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateTime(m.GetMinimumValueTimestamp())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMinimumValueTimestamp(structType any) BACnetConstructedDataMinimumValueTimestamp {
	if casted, ok := structType.(BACnetConstructedDataMinimumValueTimestamp); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMinimumValueTimestamp); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetTypeName() string {
	return "BACnetConstructedDataMinimumValueTimestamp"
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (minimumValueTimestamp)
	lengthInBits += m.MinimumValueTimestamp.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMinimumValueTimestamp BACnetConstructedDataMinimumValueTimestamp, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMinimumValueTimestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMinimumValueTimestamp")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	minimumValueTimestamp, err := ReadSimpleField[BACnetDateTime](ctx, "minimumValueTimestamp", ReadComplex[BACnetDateTime](BACnetDateTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minimumValueTimestamp' field"))
	}
	m.MinimumValueTimestamp = minimumValueTimestamp

	actualValue, err := ReadVirtualField[BACnetDateTime](ctx, "actualValue", (*BACnetDateTime)(nil), minimumValueTimestamp)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMinimumValueTimestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMinimumValueTimestamp")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMinimumValueTimestamp"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMinimumValueTimestamp")
		}

		if err := WriteSimpleField[BACnetDateTime](ctx, "minimumValueTimestamp", m.GetMinimumValueTimestamp(), WriteComplex[BACnetDateTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'minimumValueTimestamp' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMinimumValueTimestamp"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMinimumValueTimestamp")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) IsBACnetConstructedDataMinimumValueTimestamp() {
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) deepCopy() *_BACnetConstructedDataMinimumValueTimestamp {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataMinimumValueTimestampCopy := &_BACnetConstructedDataMinimumValueTimestamp{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetDateTime](m.MinimumValueTimestamp),
	}
	_BACnetConstructedDataMinimumValueTimestampCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataMinimumValueTimestampCopy
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) String() string {
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