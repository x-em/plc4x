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

// BACnetConstructedDataOctetStringValueRelinquishDefault is the corresponding interface of BACnetConstructedDataOctetStringValueRelinquishDefault
type BACnetConstructedDataOctetStringValueRelinquishDefault interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetRelinquishDefault returns RelinquishDefault (property field)
	GetRelinquishDefault() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
	// IsBACnetConstructedDataOctetStringValueRelinquishDefault is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataOctetStringValueRelinquishDefault()
	// CreateBuilder creates a BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder
	CreateBACnetConstructedDataOctetStringValueRelinquishDefaultBuilder() BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder
}

// _BACnetConstructedDataOctetStringValueRelinquishDefault is the data-structure of this message
type _BACnetConstructedDataOctetStringValueRelinquishDefault struct {
	BACnetConstructedDataContract
	RelinquishDefault BACnetApplicationTagSignedInteger
}

var _ BACnetConstructedDataOctetStringValueRelinquishDefault = (*_BACnetConstructedDataOctetStringValueRelinquishDefault)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataOctetStringValueRelinquishDefault)(nil)

// NewBACnetConstructedDataOctetStringValueRelinquishDefault factory function for _BACnetConstructedDataOctetStringValueRelinquishDefault
func NewBACnetConstructedDataOctetStringValueRelinquishDefault(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, relinquishDefault BACnetApplicationTagSignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOctetStringValueRelinquishDefault {
	if relinquishDefault == nil {
		panic("relinquishDefault of type BACnetApplicationTagSignedInteger for BACnetConstructedDataOctetStringValueRelinquishDefault must not be nil")
	}
	_result := &_BACnetConstructedDataOctetStringValueRelinquishDefault{
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

// BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder is a builder for BACnetConstructedDataOctetStringValueRelinquishDefault
type BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(relinquishDefault BACnetApplicationTagSignedInteger) BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder
	// WithRelinquishDefault adds RelinquishDefault (property field)
	WithRelinquishDefault(BACnetApplicationTagSignedInteger) BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder
	// WithRelinquishDefaultBuilder adds RelinquishDefault (property field) which is build by the builder
	WithRelinquishDefaultBuilder(func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataOctetStringValueRelinquishDefault or returns an error if something is wrong
	Build() (BACnetConstructedDataOctetStringValueRelinquishDefault, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataOctetStringValueRelinquishDefault
}

// NewBACnetConstructedDataOctetStringValueRelinquishDefaultBuilder() creates a BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder
func NewBACnetConstructedDataOctetStringValueRelinquishDefaultBuilder() BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder {
	return &_BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder{_BACnetConstructedDataOctetStringValueRelinquishDefault: new(_BACnetConstructedDataOctetStringValueRelinquishDefault)}
}

type _BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder struct {
	*_BACnetConstructedDataOctetStringValueRelinquishDefault

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder) = (*_BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder)(nil)

func (b *_BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataOctetStringValueRelinquishDefault
}

func (b *_BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder) WithMandatoryFields(relinquishDefault BACnetApplicationTagSignedInteger) BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder {
	return b.WithRelinquishDefault(relinquishDefault)
}

func (b *_BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder) WithRelinquishDefault(relinquishDefault BACnetApplicationTagSignedInteger) BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder {
	b.RelinquishDefault = relinquishDefault
	return b
}

func (b *_BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder) WithRelinquishDefaultBuilder(builderSupplier func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder {
	builder := builderSupplier(b.RelinquishDefault.CreateBACnetApplicationTagSignedIntegerBuilder())
	var err error
	b.RelinquishDefault, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder) Build() (BACnetConstructedDataOctetStringValueRelinquishDefault, error) {
	if b.RelinquishDefault == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'relinquishDefault' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataOctetStringValueRelinquishDefault.deepCopy(), nil
}

func (b *_BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder) MustBuild() BACnetConstructedDataOctetStringValueRelinquishDefault {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataOctetStringValueRelinquishDefaultBuilder().(*_BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataOctetStringValueRelinquishDefaultBuilder creates a BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder
func (b *_BACnetConstructedDataOctetStringValueRelinquishDefault) CreateBACnetConstructedDataOctetStringValueRelinquishDefaultBuilder() BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder {
	if b == nil {
		return NewBACnetConstructedDataOctetStringValueRelinquishDefaultBuilder()
	}
	return &_BACnetConstructedDataOctetStringValueRelinquishDefaultBuilder{_BACnetConstructedDataOctetStringValueRelinquishDefault: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOctetStringValueRelinquishDefault) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_OCTETSTRING_VALUE
}

func (m *_BACnetConstructedDataOctetStringValueRelinquishDefault) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RELINQUISH_DEFAULT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOctetStringValueRelinquishDefault) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataOctetStringValueRelinquishDefault) GetRelinquishDefault() BACnetApplicationTagSignedInteger {
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

func (m *_BACnetConstructedDataOctetStringValueRelinquishDefault) GetActualValue() BACnetApplicationTagSignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagSignedInteger(m.GetRelinquishDefault())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOctetStringValueRelinquishDefault(structType any) BACnetConstructedDataOctetStringValueRelinquishDefault {
	if casted, ok := structType.(BACnetConstructedDataOctetStringValueRelinquishDefault); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOctetStringValueRelinquishDefault); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOctetStringValueRelinquishDefault) GetTypeName() string {
	return "BACnetConstructedDataOctetStringValueRelinquishDefault"
}

func (m *_BACnetConstructedDataOctetStringValueRelinquishDefault) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (relinquishDefault)
	lengthInBits += m.RelinquishDefault.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataOctetStringValueRelinquishDefault) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataOctetStringValueRelinquishDefault) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataOctetStringValueRelinquishDefault BACnetConstructedDataOctetStringValueRelinquishDefault, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOctetStringValueRelinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOctetStringValueRelinquishDefault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	relinquishDefault, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "relinquishDefault", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'relinquishDefault' field"))
	}
	m.RelinquishDefault = relinquishDefault

	actualValue, err := ReadVirtualField[BACnetApplicationTagSignedInteger](ctx, "actualValue", (*BACnetApplicationTagSignedInteger)(nil), relinquishDefault)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOctetStringValueRelinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOctetStringValueRelinquishDefault")
	}

	return m, nil
}

func (m *_BACnetConstructedDataOctetStringValueRelinquishDefault) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataOctetStringValueRelinquishDefault) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOctetStringValueRelinquishDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOctetStringValueRelinquishDefault")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "relinquishDefault", m.GetRelinquishDefault(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'relinquishDefault' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOctetStringValueRelinquishDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOctetStringValueRelinquishDefault")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOctetStringValueRelinquishDefault) IsBACnetConstructedDataOctetStringValueRelinquishDefault() {
}

func (m *_BACnetConstructedDataOctetStringValueRelinquishDefault) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataOctetStringValueRelinquishDefault) deepCopy() *_BACnetConstructedDataOctetStringValueRelinquishDefault {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataOctetStringValueRelinquishDefaultCopy := &_BACnetConstructedDataOctetStringValueRelinquishDefault{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagSignedInteger](m.RelinquishDefault),
	}
	_BACnetConstructedDataOctetStringValueRelinquishDefaultCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataOctetStringValueRelinquishDefaultCopy
}

func (m *_BACnetConstructedDataOctetStringValueRelinquishDefault) String() string {
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