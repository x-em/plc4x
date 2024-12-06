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

// BACnetConstructedDataAllWritesSuccessful is the corresponding interface of BACnetConstructedDataAllWritesSuccessful
type BACnetConstructedDataAllWritesSuccessful interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetAllWritesSuccessful returns AllWritesSuccessful (property field)
	GetAllWritesSuccessful() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataAllWritesSuccessful is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAllWritesSuccessful()
	// CreateBuilder creates a BACnetConstructedDataAllWritesSuccessfulBuilder
	CreateBACnetConstructedDataAllWritesSuccessfulBuilder() BACnetConstructedDataAllWritesSuccessfulBuilder
}

// _BACnetConstructedDataAllWritesSuccessful is the data-structure of this message
type _BACnetConstructedDataAllWritesSuccessful struct {
	BACnetConstructedDataContract
	AllWritesSuccessful BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataAllWritesSuccessful = (*_BACnetConstructedDataAllWritesSuccessful)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAllWritesSuccessful)(nil)

// NewBACnetConstructedDataAllWritesSuccessful factory function for _BACnetConstructedDataAllWritesSuccessful
func NewBACnetConstructedDataAllWritesSuccessful(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, allWritesSuccessful BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAllWritesSuccessful {
	if allWritesSuccessful == nil {
		panic("allWritesSuccessful of type BACnetApplicationTagBoolean for BACnetConstructedDataAllWritesSuccessful must not be nil")
	}
	_result := &_BACnetConstructedDataAllWritesSuccessful{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AllWritesSuccessful:           allWritesSuccessful,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAllWritesSuccessfulBuilder is a builder for BACnetConstructedDataAllWritesSuccessful
type BACnetConstructedDataAllWritesSuccessfulBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(allWritesSuccessful BACnetApplicationTagBoolean) BACnetConstructedDataAllWritesSuccessfulBuilder
	// WithAllWritesSuccessful adds AllWritesSuccessful (property field)
	WithAllWritesSuccessful(BACnetApplicationTagBoolean) BACnetConstructedDataAllWritesSuccessfulBuilder
	// WithAllWritesSuccessfulBuilder adds AllWritesSuccessful (property field) which is build by the builder
	WithAllWritesSuccessfulBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataAllWritesSuccessfulBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataAllWritesSuccessful or returns an error if something is wrong
	Build() (BACnetConstructedDataAllWritesSuccessful, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAllWritesSuccessful
}

// NewBACnetConstructedDataAllWritesSuccessfulBuilder() creates a BACnetConstructedDataAllWritesSuccessfulBuilder
func NewBACnetConstructedDataAllWritesSuccessfulBuilder() BACnetConstructedDataAllWritesSuccessfulBuilder {
	return &_BACnetConstructedDataAllWritesSuccessfulBuilder{_BACnetConstructedDataAllWritesSuccessful: new(_BACnetConstructedDataAllWritesSuccessful)}
}

type _BACnetConstructedDataAllWritesSuccessfulBuilder struct {
	*_BACnetConstructedDataAllWritesSuccessful

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAllWritesSuccessfulBuilder) = (*_BACnetConstructedDataAllWritesSuccessfulBuilder)(nil)

func (b *_BACnetConstructedDataAllWritesSuccessfulBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataAllWritesSuccessful
}

func (b *_BACnetConstructedDataAllWritesSuccessfulBuilder) WithMandatoryFields(allWritesSuccessful BACnetApplicationTagBoolean) BACnetConstructedDataAllWritesSuccessfulBuilder {
	return b.WithAllWritesSuccessful(allWritesSuccessful)
}

func (b *_BACnetConstructedDataAllWritesSuccessfulBuilder) WithAllWritesSuccessful(allWritesSuccessful BACnetApplicationTagBoolean) BACnetConstructedDataAllWritesSuccessfulBuilder {
	b.AllWritesSuccessful = allWritesSuccessful
	return b
}

func (b *_BACnetConstructedDataAllWritesSuccessfulBuilder) WithAllWritesSuccessfulBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataAllWritesSuccessfulBuilder {
	builder := builderSupplier(b.AllWritesSuccessful.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.AllWritesSuccessful, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAllWritesSuccessfulBuilder) Build() (BACnetConstructedDataAllWritesSuccessful, error) {
	if b.AllWritesSuccessful == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'allWritesSuccessful' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAllWritesSuccessful.deepCopy(), nil
}

func (b *_BACnetConstructedDataAllWritesSuccessfulBuilder) MustBuild() BACnetConstructedDataAllWritesSuccessful {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataAllWritesSuccessfulBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAllWritesSuccessfulBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAllWritesSuccessfulBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAllWritesSuccessfulBuilder().(*_BACnetConstructedDataAllWritesSuccessfulBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAllWritesSuccessfulBuilder creates a BACnetConstructedDataAllWritesSuccessfulBuilder
func (b *_BACnetConstructedDataAllWritesSuccessful) CreateBACnetConstructedDataAllWritesSuccessfulBuilder() BACnetConstructedDataAllWritesSuccessfulBuilder {
	if b == nil {
		return NewBACnetConstructedDataAllWritesSuccessfulBuilder()
	}
	return &_BACnetConstructedDataAllWritesSuccessfulBuilder{_BACnetConstructedDataAllWritesSuccessful: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAllWritesSuccessful) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAllWritesSuccessful) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL_WRITES_SUCCESSFUL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAllWritesSuccessful) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAllWritesSuccessful) GetAllWritesSuccessful() BACnetApplicationTagBoolean {
	return m.AllWritesSuccessful
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAllWritesSuccessful) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetAllWritesSuccessful())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAllWritesSuccessful(structType any) BACnetConstructedDataAllWritesSuccessful {
	if casted, ok := structType.(BACnetConstructedDataAllWritesSuccessful); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAllWritesSuccessful); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAllWritesSuccessful) GetTypeName() string {
	return "BACnetConstructedDataAllWritesSuccessful"
}

func (m *_BACnetConstructedDataAllWritesSuccessful) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (allWritesSuccessful)
	lengthInBits += m.AllWritesSuccessful.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAllWritesSuccessful) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAllWritesSuccessful) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAllWritesSuccessful BACnetConstructedDataAllWritesSuccessful, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAllWritesSuccessful"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAllWritesSuccessful")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	allWritesSuccessful, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "allWritesSuccessful", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'allWritesSuccessful' field"))
	}
	m.AllWritesSuccessful = allWritesSuccessful

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), allWritesSuccessful)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAllWritesSuccessful"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAllWritesSuccessful")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAllWritesSuccessful) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAllWritesSuccessful) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAllWritesSuccessful"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAllWritesSuccessful")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "allWritesSuccessful", m.GetAllWritesSuccessful(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'allWritesSuccessful' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAllWritesSuccessful"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAllWritesSuccessful")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAllWritesSuccessful) IsBACnetConstructedDataAllWritesSuccessful() {}

func (m *_BACnetConstructedDataAllWritesSuccessful) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAllWritesSuccessful) deepCopy() *_BACnetConstructedDataAllWritesSuccessful {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAllWritesSuccessfulCopy := &_BACnetConstructedDataAllWritesSuccessful{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBoolean](m.AllWritesSuccessful),
	}
	_BACnetConstructedDataAllWritesSuccessfulCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAllWritesSuccessfulCopy
}

func (m *_BACnetConstructedDataAllWritesSuccessful) String() string {
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