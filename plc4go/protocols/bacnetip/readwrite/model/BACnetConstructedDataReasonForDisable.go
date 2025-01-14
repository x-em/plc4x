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

// BACnetConstructedDataReasonForDisable is the corresponding interface of BACnetConstructedDataReasonForDisable
type BACnetConstructedDataReasonForDisable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetReasonForDisable returns ReasonForDisable (property field)
	GetReasonForDisable() []BACnetAccessCredentialDisableReasonTagged
	// IsBACnetConstructedDataReasonForDisable is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataReasonForDisable()
	// CreateBuilder creates a BACnetConstructedDataReasonForDisableBuilder
	CreateBACnetConstructedDataReasonForDisableBuilder() BACnetConstructedDataReasonForDisableBuilder
}

// _BACnetConstructedDataReasonForDisable is the data-structure of this message
type _BACnetConstructedDataReasonForDisable struct {
	BACnetConstructedDataContract
	ReasonForDisable []BACnetAccessCredentialDisableReasonTagged
}

var _ BACnetConstructedDataReasonForDisable = (*_BACnetConstructedDataReasonForDisable)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataReasonForDisable)(nil)

// NewBACnetConstructedDataReasonForDisable factory function for _BACnetConstructedDataReasonForDisable
func NewBACnetConstructedDataReasonForDisable(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, reasonForDisable []BACnetAccessCredentialDisableReasonTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataReasonForDisable {
	_result := &_BACnetConstructedDataReasonForDisable{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ReasonForDisable:              reasonForDisable,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataReasonForDisableBuilder is a builder for BACnetConstructedDataReasonForDisable
type BACnetConstructedDataReasonForDisableBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(reasonForDisable []BACnetAccessCredentialDisableReasonTagged) BACnetConstructedDataReasonForDisableBuilder
	// WithReasonForDisable adds ReasonForDisable (property field)
	WithReasonForDisable(...BACnetAccessCredentialDisableReasonTagged) BACnetConstructedDataReasonForDisableBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataReasonForDisable or returns an error if something is wrong
	Build() (BACnetConstructedDataReasonForDisable, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataReasonForDisable
}

// NewBACnetConstructedDataReasonForDisableBuilder() creates a BACnetConstructedDataReasonForDisableBuilder
func NewBACnetConstructedDataReasonForDisableBuilder() BACnetConstructedDataReasonForDisableBuilder {
	return &_BACnetConstructedDataReasonForDisableBuilder{_BACnetConstructedDataReasonForDisable: new(_BACnetConstructedDataReasonForDisable)}
}

type _BACnetConstructedDataReasonForDisableBuilder struct {
	*_BACnetConstructedDataReasonForDisable

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataReasonForDisableBuilder) = (*_BACnetConstructedDataReasonForDisableBuilder)(nil)

func (b *_BACnetConstructedDataReasonForDisableBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataReasonForDisable
}

func (b *_BACnetConstructedDataReasonForDisableBuilder) WithMandatoryFields(reasonForDisable []BACnetAccessCredentialDisableReasonTagged) BACnetConstructedDataReasonForDisableBuilder {
	return b.WithReasonForDisable(reasonForDisable...)
}

func (b *_BACnetConstructedDataReasonForDisableBuilder) WithReasonForDisable(reasonForDisable ...BACnetAccessCredentialDisableReasonTagged) BACnetConstructedDataReasonForDisableBuilder {
	b.ReasonForDisable = reasonForDisable
	return b
}

func (b *_BACnetConstructedDataReasonForDisableBuilder) Build() (BACnetConstructedDataReasonForDisable, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataReasonForDisable.deepCopy(), nil
}

func (b *_BACnetConstructedDataReasonForDisableBuilder) MustBuild() BACnetConstructedDataReasonForDisable {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataReasonForDisableBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataReasonForDisableBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataReasonForDisableBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataReasonForDisableBuilder().(*_BACnetConstructedDataReasonForDisableBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataReasonForDisableBuilder creates a BACnetConstructedDataReasonForDisableBuilder
func (b *_BACnetConstructedDataReasonForDisable) CreateBACnetConstructedDataReasonForDisableBuilder() BACnetConstructedDataReasonForDisableBuilder {
	if b == nil {
		return NewBACnetConstructedDataReasonForDisableBuilder()
	}
	return &_BACnetConstructedDataReasonForDisableBuilder{_BACnetConstructedDataReasonForDisable: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataReasonForDisable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataReasonForDisable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_REASON_FOR_DISABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataReasonForDisable) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataReasonForDisable) GetReasonForDisable() []BACnetAccessCredentialDisableReasonTagged {
	return m.ReasonForDisable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataReasonForDisable(structType any) BACnetConstructedDataReasonForDisable {
	if casted, ok := structType.(BACnetConstructedDataReasonForDisable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataReasonForDisable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataReasonForDisable) GetTypeName() string {
	return "BACnetConstructedDataReasonForDisable"
}

func (m *_BACnetConstructedDataReasonForDisable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.ReasonForDisable) > 0 {
		for _, element := range m.ReasonForDisable {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataReasonForDisable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataReasonForDisable) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataReasonForDisable BACnetConstructedDataReasonForDisable, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataReasonForDisable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataReasonForDisable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reasonForDisable, err := ReadTerminatedArrayField[BACnetAccessCredentialDisableReasonTagged](ctx, "reasonForDisable", ReadComplex[BACnetAccessCredentialDisableReasonTagged](BACnetAccessCredentialDisableReasonTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reasonForDisable' field"))
	}
	m.ReasonForDisable = reasonForDisable

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataReasonForDisable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataReasonForDisable")
	}

	return m, nil
}

func (m *_BACnetConstructedDataReasonForDisable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataReasonForDisable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataReasonForDisable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataReasonForDisable")
		}

		if err := WriteComplexTypeArrayField(ctx, "reasonForDisable", m.GetReasonForDisable(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'reasonForDisable' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataReasonForDisable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataReasonForDisable")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataReasonForDisable) IsBACnetConstructedDataReasonForDisable() {}

func (m *_BACnetConstructedDataReasonForDisable) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataReasonForDisable) deepCopy() *_BACnetConstructedDataReasonForDisable {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataReasonForDisableCopy := &_BACnetConstructedDataReasonForDisable{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetAccessCredentialDisableReasonTagged, BACnetAccessCredentialDisableReasonTagged](m.ReasonForDisable),
	}
	_BACnetConstructedDataReasonForDisableCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataReasonForDisableCopy
}

func (m *_BACnetConstructedDataReasonForDisable) String() string {
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
