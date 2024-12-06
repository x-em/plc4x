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

// BACnetConstructedDataLastRestartReason is the corresponding interface of BACnetConstructedDataLastRestartReason
type BACnetConstructedDataLastRestartReason interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLastRestartReason returns LastRestartReason (property field)
	GetLastRestartReason() BACnetRestartReasonTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetRestartReasonTagged
	// IsBACnetConstructedDataLastRestartReason is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLastRestartReason()
	// CreateBuilder creates a BACnetConstructedDataLastRestartReasonBuilder
	CreateBACnetConstructedDataLastRestartReasonBuilder() BACnetConstructedDataLastRestartReasonBuilder
}

// _BACnetConstructedDataLastRestartReason is the data-structure of this message
type _BACnetConstructedDataLastRestartReason struct {
	BACnetConstructedDataContract
	LastRestartReason BACnetRestartReasonTagged
}

var _ BACnetConstructedDataLastRestartReason = (*_BACnetConstructedDataLastRestartReason)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLastRestartReason)(nil)

// NewBACnetConstructedDataLastRestartReason factory function for _BACnetConstructedDataLastRestartReason
func NewBACnetConstructedDataLastRestartReason(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lastRestartReason BACnetRestartReasonTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastRestartReason {
	if lastRestartReason == nil {
		panic("lastRestartReason of type BACnetRestartReasonTagged for BACnetConstructedDataLastRestartReason must not be nil")
	}
	_result := &_BACnetConstructedDataLastRestartReason{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LastRestartReason:             lastRestartReason,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLastRestartReasonBuilder is a builder for BACnetConstructedDataLastRestartReason
type BACnetConstructedDataLastRestartReasonBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lastRestartReason BACnetRestartReasonTagged) BACnetConstructedDataLastRestartReasonBuilder
	// WithLastRestartReason adds LastRestartReason (property field)
	WithLastRestartReason(BACnetRestartReasonTagged) BACnetConstructedDataLastRestartReasonBuilder
	// WithLastRestartReasonBuilder adds LastRestartReason (property field) which is build by the builder
	WithLastRestartReasonBuilder(func(BACnetRestartReasonTaggedBuilder) BACnetRestartReasonTaggedBuilder) BACnetConstructedDataLastRestartReasonBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLastRestartReason or returns an error if something is wrong
	Build() (BACnetConstructedDataLastRestartReason, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLastRestartReason
}

// NewBACnetConstructedDataLastRestartReasonBuilder() creates a BACnetConstructedDataLastRestartReasonBuilder
func NewBACnetConstructedDataLastRestartReasonBuilder() BACnetConstructedDataLastRestartReasonBuilder {
	return &_BACnetConstructedDataLastRestartReasonBuilder{_BACnetConstructedDataLastRestartReason: new(_BACnetConstructedDataLastRestartReason)}
}

type _BACnetConstructedDataLastRestartReasonBuilder struct {
	*_BACnetConstructedDataLastRestartReason

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLastRestartReasonBuilder) = (*_BACnetConstructedDataLastRestartReasonBuilder)(nil)

func (b *_BACnetConstructedDataLastRestartReasonBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataLastRestartReason
}

func (b *_BACnetConstructedDataLastRestartReasonBuilder) WithMandatoryFields(lastRestartReason BACnetRestartReasonTagged) BACnetConstructedDataLastRestartReasonBuilder {
	return b.WithLastRestartReason(lastRestartReason)
}

func (b *_BACnetConstructedDataLastRestartReasonBuilder) WithLastRestartReason(lastRestartReason BACnetRestartReasonTagged) BACnetConstructedDataLastRestartReasonBuilder {
	b.LastRestartReason = lastRestartReason
	return b
}

func (b *_BACnetConstructedDataLastRestartReasonBuilder) WithLastRestartReasonBuilder(builderSupplier func(BACnetRestartReasonTaggedBuilder) BACnetRestartReasonTaggedBuilder) BACnetConstructedDataLastRestartReasonBuilder {
	builder := builderSupplier(b.LastRestartReason.CreateBACnetRestartReasonTaggedBuilder())
	var err error
	b.LastRestartReason, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetRestartReasonTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLastRestartReasonBuilder) Build() (BACnetConstructedDataLastRestartReason, error) {
	if b.LastRestartReason == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lastRestartReason' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLastRestartReason.deepCopy(), nil
}

func (b *_BACnetConstructedDataLastRestartReasonBuilder) MustBuild() BACnetConstructedDataLastRestartReason {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLastRestartReasonBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLastRestartReasonBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLastRestartReasonBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLastRestartReasonBuilder().(*_BACnetConstructedDataLastRestartReasonBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLastRestartReasonBuilder creates a BACnetConstructedDataLastRestartReasonBuilder
func (b *_BACnetConstructedDataLastRestartReason) CreateBACnetConstructedDataLastRestartReasonBuilder() BACnetConstructedDataLastRestartReasonBuilder {
	if b == nil {
		return NewBACnetConstructedDataLastRestartReasonBuilder()
	}
	return &_BACnetConstructedDataLastRestartReasonBuilder{_BACnetConstructedDataLastRestartReason: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastRestartReason) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastRestartReason) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_RESTART_REASON
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastRestartReason) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastRestartReason) GetLastRestartReason() BACnetRestartReasonTagged {
	return m.LastRestartReason
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastRestartReason) GetActualValue() BACnetRestartReasonTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetRestartReasonTagged(m.GetLastRestartReason())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastRestartReason(structType any) BACnetConstructedDataLastRestartReason {
	if casted, ok := structType.(BACnetConstructedDataLastRestartReason); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastRestartReason); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastRestartReason) GetTypeName() string {
	return "BACnetConstructedDataLastRestartReason"
}

func (m *_BACnetConstructedDataLastRestartReason) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lastRestartReason)
	lengthInBits += m.LastRestartReason.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastRestartReason) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLastRestartReason) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLastRestartReason BACnetConstructedDataLastRestartReason, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastRestartReason"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastRestartReason")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lastRestartReason, err := ReadSimpleField[BACnetRestartReasonTagged](ctx, "lastRestartReason", ReadComplex[BACnetRestartReasonTagged](BACnetRestartReasonTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastRestartReason' field"))
	}
	m.LastRestartReason = lastRestartReason

	actualValue, err := ReadVirtualField[BACnetRestartReasonTagged](ctx, "actualValue", (*BACnetRestartReasonTagged)(nil), lastRestartReason)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastRestartReason"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastRestartReason")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLastRestartReason) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastRestartReason) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastRestartReason"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastRestartReason")
		}

		if err := WriteSimpleField[BACnetRestartReasonTagged](ctx, "lastRestartReason", m.GetLastRestartReason(), WriteComplex[BACnetRestartReasonTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastRestartReason' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastRestartReason"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastRestartReason")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastRestartReason) IsBACnetConstructedDataLastRestartReason() {}

func (m *_BACnetConstructedDataLastRestartReason) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLastRestartReason) deepCopy() *_BACnetConstructedDataLastRestartReason {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLastRestartReasonCopy := &_BACnetConstructedDataLastRestartReason{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetRestartReasonTagged](m.LastRestartReason),
	}
	_BACnetConstructedDataLastRestartReasonCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLastRestartReasonCopy
}

func (m *_BACnetConstructedDataLastRestartReason) String() string {
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