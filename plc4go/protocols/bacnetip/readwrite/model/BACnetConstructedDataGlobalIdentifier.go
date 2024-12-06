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

// BACnetConstructedDataGlobalIdentifier is the corresponding interface of BACnetConstructedDataGlobalIdentifier
type BACnetConstructedDataGlobalIdentifier interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetGlobalIdentifier returns GlobalIdentifier (property field)
	GetGlobalIdentifier() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataGlobalIdentifier is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataGlobalIdentifier()
	// CreateBuilder creates a BACnetConstructedDataGlobalIdentifierBuilder
	CreateBACnetConstructedDataGlobalIdentifierBuilder() BACnetConstructedDataGlobalIdentifierBuilder
}

// _BACnetConstructedDataGlobalIdentifier is the data-structure of this message
type _BACnetConstructedDataGlobalIdentifier struct {
	BACnetConstructedDataContract
	GlobalIdentifier BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataGlobalIdentifier = (*_BACnetConstructedDataGlobalIdentifier)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataGlobalIdentifier)(nil)

// NewBACnetConstructedDataGlobalIdentifier factory function for _BACnetConstructedDataGlobalIdentifier
func NewBACnetConstructedDataGlobalIdentifier(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, globalIdentifier BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataGlobalIdentifier {
	if globalIdentifier == nil {
		panic("globalIdentifier of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataGlobalIdentifier must not be nil")
	}
	_result := &_BACnetConstructedDataGlobalIdentifier{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		GlobalIdentifier:              globalIdentifier,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataGlobalIdentifierBuilder is a builder for BACnetConstructedDataGlobalIdentifier
type BACnetConstructedDataGlobalIdentifierBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(globalIdentifier BACnetApplicationTagUnsignedInteger) BACnetConstructedDataGlobalIdentifierBuilder
	// WithGlobalIdentifier adds GlobalIdentifier (property field)
	WithGlobalIdentifier(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataGlobalIdentifierBuilder
	// WithGlobalIdentifierBuilder adds GlobalIdentifier (property field) which is build by the builder
	WithGlobalIdentifierBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataGlobalIdentifierBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataGlobalIdentifier or returns an error if something is wrong
	Build() (BACnetConstructedDataGlobalIdentifier, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataGlobalIdentifier
}

// NewBACnetConstructedDataGlobalIdentifierBuilder() creates a BACnetConstructedDataGlobalIdentifierBuilder
func NewBACnetConstructedDataGlobalIdentifierBuilder() BACnetConstructedDataGlobalIdentifierBuilder {
	return &_BACnetConstructedDataGlobalIdentifierBuilder{_BACnetConstructedDataGlobalIdentifier: new(_BACnetConstructedDataGlobalIdentifier)}
}

type _BACnetConstructedDataGlobalIdentifierBuilder struct {
	*_BACnetConstructedDataGlobalIdentifier

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataGlobalIdentifierBuilder) = (*_BACnetConstructedDataGlobalIdentifierBuilder)(nil)

func (b *_BACnetConstructedDataGlobalIdentifierBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataGlobalIdentifier
}

func (b *_BACnetConstructedDataGlobalIdentifierBuilder) WithMandatoryFields(globalIdentifier BACnetApplicationTagUnsignedInteger) BACnetConstructedDataGlobalIdentifierBuilder {
	return b.WithGlobalIdentifier(globalIdentifier)
}

func (b *_BACnetConstructedDataGlobalIdentifierBuilder) WithGlobalIdentifier(globalIdentifier BACnetApplicationTagUnsignedInteger) BACnetConstructedDataGlobalIdentifierBuilder {
	b.GlobalIdentifier = globalIdentifier
	return b
}

func (b *_BACnetConstructedDataGlobalIdentifierBuilder) WithGlobalIdentifierBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataGlobalIdentifierBuilder {
	builder := builderSupplier(b.GlobalIdentifier.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.GlobalIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataGlobalIdentifierBuilder) Build() (BACnetConstructedDataGlobalIdentifier, error) {
	if b.GlobalIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'globalIdentifier' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataGlobalIdentifier.deepCopy(), nil
}

func (b *_BACnetConstructedDataGlobalIdentifierBuilder) MustBuild() BACnetConstructedDataGlobalIdentifier {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataGlobalIdentifierBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataGlobalIdentifierBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataGlobalIdentifierBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataGlobalIdentifierBuilder().(*_BACnetConstructedDataGlobalIdentifierBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataGlobalIdentifierBuilder creates a BACnetConstructedDataGlobalIdentifierBuilder
func (b *_BACnetConstructedDataGlobalIdentifier) CreateBACnetConstructedDataGlobalIdentifierBuilder() BACnetConstructedDataGlobalIdentifierBuilder {
	if b == nil {
		return NewBACnetConstructedDataGlobalIdentifierBuilder()
	}
	return &_BACnetConstructedDataGlobalIdentifierBuilder{_BACnetConstructedDataGlobalIdentifier: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataGlobalIdentifier) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataGlobalIdentifier) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_GLOBAL_IDENTIFIER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataGlobalIdentifier) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataGlobalIdentifier) GetGlobalIdentifier() BACnetApplicationTagUnsignedInteger {
	return m.GlobalIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataGlobalIdentifier) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetGlobalIdentifier())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataGlobalIdentifier(structType any) BACnetConstructedDataGlobalIdentifier {
	if casted, ok := structType.(BACnetConstructedDataGlobalIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataGlobalIdentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataGlobalIdentifier) GetTypeName() string {
	return "BACnetConstructedDataGlobalIdentifier"
}

func (m *_BACnetConstructedDataGlobalIdentifier) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (globalIdentifier)
	lengthInBits += m.GlobalIdentifier.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataGlobalIdentifier) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataGlobalIdentifier) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataGlobalIdentifier BACnetConstructedDataGlobalIdentifier, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataGlobalIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataGlobalIdentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	globalIdentifier, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "globalIdentifier", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'globalIdentifier' field"))
	}
	m.GlobalIdentifier = globalIdentifier

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), globalIdentifier)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataGlobalIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataGlobalIdentifier")
	}

	return m, nil
}

func (m *_BACnetConstructedDataGlobalIdentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataGlobalIdentifier) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataGlobalIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataGlobalIdentifier")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "globalIdentifier", m.GetGlobalIdentifier(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'globalIdentifier' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataGlobalIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataGlobalIdentifier")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataGlobalIdentifier) IsBACnetConstructedDataGlobalIdentifier() {}

func (m *_BACnetConstructedDataGlobalIdentifier) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataGlobalIdentifier) deepCopy() *_BACnetConstructedDataGlobalIdentifier {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataGlobalIdentifierCopy := &_BACnetConstructedDataGlobalIdentifier{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.GlobalIdentifier),
	}
	_BACnetConstructedDataGlobalIdentifierCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataGlobalIdentifierCopy
}

func (m *_BACnetConstructedDataGlobalIdentifier) String() string {
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