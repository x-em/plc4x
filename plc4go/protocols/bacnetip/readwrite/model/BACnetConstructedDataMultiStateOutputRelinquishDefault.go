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

// BACnetConstructedDataMultiStateOutputRelinquishDefault is the corresponding interface of BACnetConstructedDataMultiStateOutputRelinquishDefault
type BACnetConstructedDataMultiStateOutputRelinquishDefault interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetRelinquishDefault returns RelinquishDefault (property field)
	GetRelinquishDefault() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataMultiStateOutputRelinquishDefault is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMultiStateOutputRelinquishDefault()
	// CreateBuilder creates a BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder
	CreateBACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder() BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder
}

// _BACnetConstructedDataMultiStateOutputRelinquishDefault is the data-structure of this message
type _BACnetConstructedDataMultiStateOutputRelinquishDefault struct {
	BACnetConstructedDataContract
	RelinquishDefault BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataMultiStateOutputRelinquishDefault = (*_BACnetConstructedDataMultiStateOutputRelinquishDefault)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMultiStateOutputRelinquishDefault)(nil)

// NewBACnetConstructedDataMultiStateOutputRelinquishDefault factory function for _BACnetConstructedDataMultiStateOutputRelinquishDefault
func NewBACnetConstructedDataMultiStateOutputRelinquishDefault(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, relinquishDefault BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMultiStateOutputRelinquishDefault {
	if relinquishDefault == nil {
		panic("relinquishDefault of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataMultiStateOutputRelinquishDefault must not be nil")
	}
	_result := &_BACnetConstructedDataMultiStateOutputRelinquishDefault{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		RelinquishDefault:             relinquishDefault,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder is a builder for BACnetConstructedDataMultiStateOutputRelinquishDefault
type BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(relinquishDefault BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder
	// WithRelinquishDefault adds RelinquishDefault (property field)
	WithRelinquishDefault(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder
	// WithRelinquishDefaultBuilder adds RelinquishDefault (property field) which is build by the builder
	WithRelinquishDefaultBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataMultiStateOutputRelinquishDefault or returns an error if something is wrong
	Build() (BACnetConstructedDataMultiStateOutputRelinquishDefault, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataMultiStateOutputRelinquishDefault
}

// NewBACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder() creates a BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder
func NewBACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder() BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder {
	return &_BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder{_BACnetConstructedDataMultiStateOutputRelinquishDefault: new(_BACnetConstructedDataMultiStateOutputRelinquishDefault)}
}

type _BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder struct {
	*_BACnetConstructedDataMultiStateOutputRelinquishDefault

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder) = (*_BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder)(nil)

func (b *_BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataMultiStateOutputRelinquishDefault
}

func (b *_BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder) WithMandatoryFields(relinquishDefault BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder {
	return b.WithRelinquishDefault(relinquishDefault)
}

func (b *_BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder) WithRelinquishDefault(relinquishDefault BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder {
	b.RelinquishDefault = relinquishDefault
	return b
}

func (b *_BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder) WithRelinquishDefaultBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder {
	builder := builderSupplier(b.RelinquishDefault.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.RelinquishDefault, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder) Build() (BACnetConstructedDataMultiStateOutputRelinquishDefault, error) {
	if b.RelinquishDefault == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'relinquishDefault' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataMultiStateOutputRelinquishDefault.deepCopy(), nil
}

func (b *_BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder) MustBuild() BACnetConstructedDataMultiStateOutputRelinquishDefault {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder().(*_BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder creates a BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder
func (b *_BACnetConstructedDataMultiStateOutputRelinquishDefault) CreateBACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder() BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder {
	if b == nil {
		return NewBACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder()
	}
	return &_BACnetConstructedDataMultiStateOutputRelinquishDefaultBuilder{_BACnetConstructedDataMultiStateOutputRelinquishDefault: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMultiStateOutputRelinquishDefault) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_MULTI_STATE_OUTPUT
}

func (m *_BACnetConstructedDataMultiStateOutputRelinquishDefault) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RELINQUISH_DEFAULT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMultiStateOutputRelinquishDefault) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMultiStateOutputRelinquishDefault) GetRelinquishDefault() BACnetApplicationTagUnsignedInteger {
	return m.RelinquishDefault
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMultiStateOutputRelinquishDefault) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetRelinquishDefault())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMultiStateOutputRelinquishDefault(structType any) BACnetConstructedDataMultiStateOutputRelinquishDefault {
	if casted, ok := structType.(BACnetConstructedDataMultiStateOutputRelinquishDefault); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMultiStateOutputRelinquishDefault); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMultiStateOutputRelinquishDefault) GetTypeName() string {
	return "BACnetConstructedDataMultiStateOutputRelinquishDefault"
}

func (m *_BACnetConstructedDataMultiStateOutputRelinquishDefault) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (relinquishDefault)
	lengthInBits += m.RelinquishDefault.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMultiStateOutputRelinquishDefault) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMultiStateOutputRelinquishDefault) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMultiStateOutputRelinquishDefault BACnetConstructedDataMultiStateOutputRelinquishDefault, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMultiStateOutputRelinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMultiStateOutputRelinquishDefault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	relinquishDefault, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "relinquishDefault", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'relinquishDefault' field"))
	}
	m.RelinquishDefault = relinquishDefault

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), relinquishDefault)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMultiStateOutputRelinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMultiStateOutputRelinquishDefault")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMultiStateOutputRelinquishDefault) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMultiStateOutputRelinquishDefault) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMultiStateOutputRelinquishDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMultiStateOutputRelinquishDefault")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "relinquishDefault", m.GetRelinquishDefault(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'relinquishDefault' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMultiStateOutputRelinquishDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMultiStateOutputRelinquishDefault")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMultiStateOutputRelinquishDefault) IsBACnetConstructedDataMultiStateOutputRelinquishDefault() {
}

func (m *_BACnetConstructedDataMultiStateOutputRelinquishDefault) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataMultiStateOutputRelinquishDefault) deepCopy() *_BACnetConstructedDataMultiStateOutputRelinquishDefault {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataMultiStateOutputRelinquishDefaultCopy := &_BACnetConstructedDataMultiStateOutputRelinquishDefault{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.RelinquishDefault),
	}
	_BACnetConstructedDataMultiStateOutputRelinquishDefaultCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataMultiStateOutputRelinquishDefaultCopy
}

func (m *_BACnetConstructedDataMultiStateOutputRelinquishDefault) String() string {
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
