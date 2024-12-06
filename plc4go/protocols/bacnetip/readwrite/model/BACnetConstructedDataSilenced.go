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

// BACnetConstructedDataSilenced is the corresponding interface of BACnetConstructedDataSilenced
type BACnetConstructedDataSilenced interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetSilenced returns Silenced (property field)
	GetSilenced() BACnetSilencedStateTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetSilencedStateTagged
	// IsBACnetConstructedDataSilenced is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataSilenced()
	// CreateBuilder creates a BACnetConstructedDataSilencedBuilder
	CreateBACnetConstructedDataSilencedBuilder() BACnetConstructedDataSilencedBuilder
}

// _BACnetConstructedDataSilenced is the data-structure of this message
type _BACnetConstructedDataSilenced struct {
	BACnetConstructedDataContract
	Silenced BACnetSilencedStateTagged
}

var _ BACnetConstructedDataSilenced = (*_BACnetConstructedDataSilenced)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataSilenced)(nil)

// NewBACnetConstructedDataSilenced factory function for _BACnetConstructedDataSilenced
func NewBACnetConstructedDataSilenced(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, silenced BACnetSilencedStateTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSilenced {
	if silenced == nil {
		panic("silenced of type BACnetSilencedStateTagged for BACnetConstructedDataSilenced must not be nil")
	}
	_result := &_BACnetConstructedDataSilenced{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Silenced:                      silenced,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataSilencedBuilder is a builder for BACnetConstructedDataSilenced
type BACnetConstructedDataSilencedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(silenced BACnetSilencedStateTagged) BACnetConstructedDataSilencedBuilder
	// WithSilenced adds Silenced (property field)
	WithSilenced(BACnetSilencedStateTagged) BACnetConstructedDataSilencedBuilder
	// WithSilencedBuilder adds Silenced (property field) which is build by the builder
	WithSilencedBuilder(func(BACnetSilencedStateTaggedBuilder) BACnetSilencedStateTaggedBuilder) BACnetConstructedDataSilencedBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataSilenced or returns an error if something is wrong
	Build() (BACnetConstructedDataSilenced, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataSilenced
}

// NewBACnetConstructedDataSilencedBuilder() creates a BACnetConstructedDataSilencedBuilder
func NewBACnetConstructedDataSilencedBuilder() BACnetConstructedDataSilencedBuilder {
	return &_BACnetConstructedDataSilencedBuilder{_BACnetConstructedDataSilenced: new(_BACnetConstructedDataSilenced)}
}

type _BACnetConstructedDataSilencedBuilder struct {
	*_BACnetConstructedDataSilenced

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataSilencedBuilder) = (*_BACnetConstructedDataSilencedBuilder)(nil)

func (b *_BACnetConstructedDataSilencedBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataSilenced
}

func (b *_BACnetConstructedDataSilencedBuilder) WithMandatoryFields(silenced BACnetSilencedStateTagged) BACnetConstructedDataSilencedBuilder {
	return b.WithSilenced(silenced)
}

func (b *_BACnetConstructedDataSilencedBuilder) WithSilenced(silenced BACnetSilencedStateTagged) BACnetConstructedDataSilencedBuilder {
	b.Silenced = silenced
	return b
}

func (b *_BACnetConstructedDataSilencedBuilder) WithSilencedBuilder(builderSupplier func(BACnetSilencedStateTaggedBuilder) BACnetSilencedStateTaggedBuilder) BACnetConstructedDataSilencedBuilder {
	builder := builderSupplier(b.Silenced.CreateBACnetSilencedStateTaggedBuilder())
	var err error
	b.Silenced, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetSilencedStateTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataSilencedBuilder) Build() (BACnetConstructedDataSilenced, error) {
	if b.Silenced == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'silenced' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataSilenced.deepCopy(), nil
}

func (b *_BACnetConstructedDataSilencedBuilder) MustBuild() BACnetConstructedDataSilenced {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataSilencedBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataSilencedBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataSilencedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataSilencedBuilder().(*_BACnetConstructedDataSilencedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataSilencedBuilder creates a BACnetConstructedDataSilencedBuilder
func (b *_BACnetConstructedDataSilenced) CreateBACnetConstructedDataSilencedBuilder() BACnetConstructedDataSilencedBuilder {
	if b == nil {
		return NewBACnetConstructedDataSilencedBuilder()
	}
	return &_BACnetConstructedDataSilencedBuilder{_BACnetConstructedDataSilenced: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSilenced) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSilenced) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SILENCED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSilenced) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSilenced) GetSilenced() BACnetSilencedStateTagged {
	return m.Silenced
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSilenced) GetActualValue() BACnetSilencedStateTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetSilencedStateTagged(m.GetSilenced())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSilenced(structType any) BACnetConstructedDataSilenced {
	if casted, ok := structType.(BACnetConstructedDataSilenced); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSilenced); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSilenced) GetTypeName() string {
	return "BACnetConstructedDataSilenced"
}

func (m *_BACnetConstructedDataSilenced) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (silenced)
	lengthInBits += m.Silenced.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataSilenced) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataSilenced) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataSilenced BACnetConstructedDataSilenced, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSilenced"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSilenced")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	silenced, err := ReadSimpleField[BACnetSilencedStateTagged](ctx, "silenced", ReadComplex[BACnetSilencedStateTagged](BACnetSilencedStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'silenced' field"))
	}
	m.Silenced = silenced

	actualValue, err := ReadVirtualField[BACnetSilencedStateTagged](ctx, "actualValue", (*BACnetSilencedStateTagged)(nil), silenced)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSilenced"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSilenced")
	}

	return m, nil
}

func (m *_BACnetConstructedDataSilenced) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSilenced) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSilenced"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSilenced")
		}

		if err := WriteSimpleField[BACnetSilencedStateTagged](ctx, "silenced", m.GetSilenced(), WriteComplex[BACnetSilencedStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'silenced' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSilenced"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSilenced")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSilenced) IsBACnetConstructedDataSilenced() {}

func (m *_BACnetConstructedDataSilenced) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataSilenced) deepCopy() *_BACnetConstructedDataSilenced {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataSilencedCopy := &_BACnetConstructedDataSilenced{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetSilencedStateTagged](m.Silenced),
	}
	_BACnetConstructedDataSilencedCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataSilencedCopy
}

func (m *_BACnetConstructedDataSilenced) String() string {
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
