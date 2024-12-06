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

// BACnetConstructedDataCurrentCommandPriority is the corresponding interface of BACnetConstructedDataCurrentCommandPriority
type BACnetConstructedDataCurrentCommandPriority interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetCurrentCommandPriority returns CurrentCommandPriority (property field)
	GetCurrentCommandPriority() BACnetOptionalUnsigned
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetOptionalUnsigned
	// IsBACnetConstructedDataCurrentCommandPriority is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataCurrentCommandPriority()
	// CreateBuilder creates a BACnetConstructedDataCurrentCommandPriorityBuilder
	CreateBACnetConstructedDataCurrentCommandPriorityBuilder() BACnetConstructedDataCurrentCommandPriorityBuilder
}

// _BACnetConstructedDataCurrentCommandPriority is the data-structure of this message
type _BACnetConstructedDataCurrentCommandPriority struct {
	BACnetConstructedDataContract
	CurrentCommandPriority BACnetOptionalUnsigned
}

var _ BACnetConstructedDataCurrentCommandPriority = (*_BACnetConstructedDataCurrentCommandPriority)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCurrentCommandPriority)(nil)

// NewBACnetConstructedDataCurrentCommandPriority factory function for _BACnetConstructedDataCurrentCommandPriority
func NewBACnetConstructedDataCurrentCommandPriority(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, currentCommandPriority BACnetOptionalUnsigned, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCurrentCommandPriority {
	if currentCommandPriority == nil {
		panic("currentCommandPriority of type BACnetOptionalUnsigned for BACnetConstructedDataCurrentCommandPriority must not be nil")
	}
	_result := &_BACnetConstructedDataCurrentCommandPriority{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		CurrentCommandPriority:        currentCommandPriority,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataCurrentCommandPriorityBuilder is a builder for BACnetConstructedDataCurrentCommandPriority
type BACnetConstructedDataCurrentCommandPriorityBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(currentCommandPriority BACnetOptionalUnsigned) BACnetConstructedDataCurrentCommandPriorityBuilder
	// WithCurrentCommandPriority adds CurrentCommandPriority (property field)
	WithCurrentCommandPriority(BACnetOptionalUnsigned) BACnetConstructedDataCurrentCommandPriorityBuilder
	// WithCurrentCommandPriorityBuilder adds CurrentCommandPriority (property field) which is build by the builder
	WithCurrentCommandPriorityBuilder(func(BACnetOptionalUnsignedBuilder) BACnetOptionalUnsignedBuilder) BACnetConstructedDataCurrentCommandPriorityBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataCurrentCommandPriority or returns an error if something is wrong
	Build() (BACnetConstructedDataCurrentCommandPriority, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataCurrentCommandPriority
}

// NewBACnetConstructedDataCurrentCommandPriorityBuilder() creates a BACnetConstructedDataCurrentCommandPriorityBuilder
func NewBACnetConstructedDataCurrentCommandPriorityBuilder() BACnetConstructedDataCurrentCommandPriorityBuilder {
	return &_BACnetConstructedDataCurrentCommandPriorityBuilder{_BACnetConstructedDataCurrentCommandPriority: new(_BACnetConstructedDataCurrentCommandPriority)}
}

type _BACnetConstructedDataCurrentCommandPriorityBuilder struct {
	*_BACnetConstructedDataCurrentCommandPriority

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataCurrentCommandPriorityBuilder) = (*_BACnetConstructedDataCurrentCommandPriorityBuilder)(nil)

func (b *_BACnetConstructedDataCurrentCommandPriorityBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataCurrentCommandPriority
}

func (b *_BACnetConstructedDataCurrentCommandPriorityBuilder) WithMandatoryFields(currentCommandPriority BACnetOptionalUnsigned) BACnetConstructedDataCurrentCommandPriorityBuilder {
	return b.WithCurrentCommandPriority(currentCommandPriority)
}

func (b *_BACnetConstructedDataCurrentCommandPriorityBuilder) WithCurrentCommandPriority(currentCommandPriority BACnetOptionalUnsigned) BACnetConstructedDataCurrentCommandPriorityBuilder {
	b.CurrentCommandPriority = currentCommandPriority
	return b
}

func (b *_BACnetConstructedDataCurrentCommandPriorityBuilder) WithCurrentCommandPriorityBuilder(builderSupplier func(BACnetOptionalUnsignedBuilder) BACnetOptionalUnsignedBuilder) BACnetConstructedDataCurrentCommandPriorityBuilder {
	builder := builderSupplier(b.CurrentCommandPriority.CreateBACnetOptionalUnsignedBuilder())
	var err error
	b.CurrentCommandPriority, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOptionalUnsignedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataCurrentCommandPriorityBuilder) Build() (BACnetConstructedDataCurrentCommandPriority, error) {
	if b.CurrentCommandPriority == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'currentCommandPriority' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataCurrentCommandPriority.deepCopy(), nil
}

func (b *_BACnetConstructedDataCurrentCommandPriorityBuilder) MustBuild() BACnetConstructedDataCurrentCommandPriority {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataCurrentCommandPriorityBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataCurrentCommandPriorityBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataCurrentCommandPriorityBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataCurrentCommandPriorityBuilder().(*_BACnetConstructedDataCurrentCommandPriorityBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataCurrentCommandPriorityBuilder creates a BACnetConstructedDataCurrentCommandPriorityBuilder
func (b *_BACnetConstructedDataCurrentCommandPriority) CreateBACnetConstructedDataCurrentCommandPriorityBuilder() BACnetConstructedDataCurrentCommandPriorityBuilder {
	if b == nil {
		return NewBACnetConstructedDataCurrentCommandPriorityBuilder()
	}
	return &_BACnetConstructedDataCurrentCommandPriorityBuilder{_BACnetConstructedDataCurrentCommandPriority: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCurrentCommandPriority) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCurrentCommandPriority) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CURRENT_COMMAND_PRIORITY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCurrentCommandPriority) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCurrentCommandPriority) GetCurrentCommandPriority() BACnetOptionalUnsigned {
	return m.CurrentCommandPriority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCurrentCommandPriority) GetActualValue() BACnetOptionalUnsigned {
	ctx := context.Background()
	_ = ctx
	return CastBACnetOptionalUnsigned(m.GetCurrentCommandPriority())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCurrentCommandPriority(structType any) BACnetConstructedDataCurrentCommandPriority {
	if casted, ok := structType.(BACnetConstructedDataCurrentCommandPriority); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCurrentCommandPriority); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCurrentCommandPriority) GetTypeName() string {
	return "BACnetConstructedDataCurrentCommandPriority"
}

func (m *_BACnetConstructedDataCurrentCommandPriority) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (currentCommandPriority)
	lengthInBits += m.CurrentCommandPriority.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCurrentCommandPriority) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCurrentCommandPriority) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCurrentCommandPriority BACnetConstructedDataCurrentCommandPriority, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCurrentCommandPriority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCurrentCommandPriority")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	currentCommandPriority, err := ReadSimpleField[BACnetOptionalUnsigned](ctx, "currentCommandPriority", ReadComplex[BACnetOptionalUnsigned](BACnetOptionalUnsignedParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'currentCommandPriority' field"))
	}
	m.CurrentCommandPriority = currentCommandPriority

	actualValue, err := ReadVirtualField[BACnetOptionalUnsigned](ctx, "actualValue", (*BACnetOptionalUnsigned)(nil), currentCommandPriority)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCurrentCommandPriority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCurrentCommandPriority")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCurrentCommandPriority) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCurrentCommandPriority) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCurrentCommandPriority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCurrentCommandPriority")
		}

		if err := WriteSimpleField[BACnetOptionalUnsigned](ctx, "currentCommandPriority", m.GetCurrentCommandPriority(), WriteComplex[BACnetOptionalUnsigned](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'currentCommandPriority' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCurrentCommandPriority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCurrentCommandPriority")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCurrentCommandPriority) IsBACnetConstructedDataCurrentCommandPriority() {
}

func (m *_BACnetConstructedDataCurrentCommandPriority) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataCurrentCommandPriority) deepCopy() *_BACnetConstructedDataCurrentCommandPriority {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataCurrentCommandPriorityCopy := &_BACnetConstructedDataCurrentCommandPriority{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetOptionalUnsigned](m.CurrentCommandPriority),
	}
	_BACnetConstructedDataCurrentCommandPriorityCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataCurrentCommandPriorityCopy
}

func (m *_BACnetConstructedDataCurrentCommandPriority) String() string {
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