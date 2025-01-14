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

// BACnetConstructedDataUTCTimeSynchronizationRecipients is the corresponding interface of BACnetConstructedDataUTCTimeSynchronizationRecipients
type BACnetConstructedDataUTCTimeSynchronizationRecipients interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetUtcTimeSynchronizationRecipients returns UtcTimeSynchronizationRecipients (property field)
	GetUtcTimeSynchronizationRecipients() []BACnetRecipient
	// IsBACnetConstructedDataUTCTimeSynchronizationRecipients is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataUTCTimeSynchronizationRecipients()
	// CreateBuilder creates a BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder
	CreateBACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder() BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder
}

// _BACnetConstructedDataUTCTimeSynchronizationRecipients is the data-structure of this message
type _BACnetConstructedDataUTCTimeSynchronizationRecipients struct {
	BACnetConstructedDataContract
	UtcTimeSynchronizationRecipients []BACnetRecipient
}

var _ BACnetConstructedDataUTCTimeSynchronizationRecipients = (*_BACnetConstructedDataUTCTimeSynchronizationRecipients)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataUTCTimeSynchronizationRecipients)(nil)

// NewBACnetConstructedDataUTCTimeSynchronizationRecipients factory function for _BACnetConstructedDataUTCTimeSynchronizationRecipients
func NewBACnetConstructedDataUTCTimeSynchronizationRecipients(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, utcTimeSynchronizationRecipients []BACnetRecipient, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataUTCTimeSynchronizationRecipients {
	_result := &_BACnetConstructedDataUTCTimeSynchronizationRecipients{
		BACnetConstructedDataContract:    NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		UtcTimeSynchronizationRecipients: utcTimeSynchronizationRecipients,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder is a builder for BACnetConstructedDataUTCTimeSynchronizationRecipients
type BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(utcTimeSynchronizationRecipients []BACnetRecipient) BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder
	// WithUtcTimeSynchronizationRecipients adds UtcTimeSynchronizationRecipients (property field)
	WithUtcTimeSynchronizationRecipients(...BACnetRecipient) BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataUTCTimeSynchronizationRecipients or returns an error if something is wrong
	Build() (BACnetConstructedDataUTCTimeSynchronizationRecipients, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataUTCTimeSynchronizationRecipients
}

// NewBACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder() creates a BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder
func NewBACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder() BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder {
	return &_BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder{_BACnetConstructedDataUTCTimeSynchronizationRecipients: new(_BACnetConstructedDataUTCTimeSynchronizationRecipients)}
}

type _BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder struct {
	*_BACnetConstructedDataUTCTimeSynchronizationRecipients

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder) = (*_BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder)(nil)

func (b *_BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataUTCTimeSynchronizationRecipients
}

func (b *_BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder) WithMandatoryFields(utcTimeSynchronizationRecipients []BACnetRecipient) BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder {
	return b.WithUtcTimeSynchronizationRecipients(utcTimeSynchronizationRecipients...)
}

func (b *_BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder) WithUtcTimeSynchronizationRecipients(utcTimeSynchronizationRecipients ...BACnetRecipient) BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder {
	b.UtcTimeSynchronizationRecipients = utcTimeSynchronizationRecipients
	return b
}

func (b *_BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder) Build() (BACnetConstructedDataUTCTimeSynchronizationRecipients, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataUTCTimeSynchronizationRecipients.deepCopy(), nil
}

func (b *_BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder) MustBuild() BACnetConstructedDataUTCTimeSynchronizationRecipients {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder().(*_BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder creates a BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder
func (b *_BACnetConstructedDataUTCTimeSynchronizationRecipients) CreateBACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder() BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder {
	if b == nil {
		return NewBACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder()
	}
	return &_BACnetConstructedDataUTCTimeSynchronizationRecipientsBuilder{_BACnetConstructedDataUTCTimeSynchronizationRecipients: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_UTC_TIME_SYNCHRONIZATION_RECIPIENTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) GetUtcTimeSynchronizationRecipients() []BACnetRecipient {
	return m.UtcTimeSynchronizationRecipients
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataUTCTimeSynchronizationRecipients(structType any) BACnetConstructedDataUTCTimeSynchronizationRecipients {
	if casted, ok := structType.(BACnetConstructedDataUTCTimeSynchronizationRecipients); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUTCTimeSynchronizationRecipients); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) GetTypeName() string {
	return "BACnetConstructedDataUTCTimeSynchronizationRecipients"
}

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.UtcTimeSynchronizationRecipients) > 0 {
		for _, element := range m.UtcTimeSynchronizationRecipients {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataUTCTimeSynchronizationRecipients BACnetConstructedDataUTCTimeSynchronizationRecipients, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUTCTimeSynchronizationRecipients"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataUTCTimeSynchronizationRecipients")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	utcTimeSynchronizationRecipients, err := ReadTerminatedArrayField[BACnetRecipient](ctx, "utcTimeSynchronizationRecipients", ReadComplex[BACnetRecipient](BACnetRecipientParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'utcTimeSynchronizationRecipients' field"))
	}
	m.UtcTimeSynchronizationRecipients = utcTimeSynchronizationRecipients

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUTCTimeSynchronizationRecipients"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataUTCTimeSynchronizationRecipients")
	}

	return m, nil
}

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUTCTimeSynchronizationRecipients"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataUTCTimeSynchronizationRecipients")
		}

		if err := WriteComplexTypeArrayField(ctx, "utcTimeSynchronizationRecipients", m.GetUtcTimeSynchronizationRecipients(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'utcTimeSynchronizationRecipients' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUTCTimeSynchronizationRecipients"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataUTCTimeSynchronizationRecipients")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) IsBACnetConstructedDataUTCTimeSynchronizationRecipients() {
}

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) deepCopy() *_BACnetConstructedDataUTCTimeSynchronizationRecipients {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataUTCTimeSynchronizationRecipientsCopy := &_BACnetConstructedDataUTCTimeSynchronizationRecipients{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetRecipient, BACnetRecipient](m.UtcTimeSynchronizationRecipients),
	}
	_BACnetConstructedDataUTCTimeSynchronizationRecipientsCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataUTCTimeSynchronizationRecipientsCopy
}

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) String() string {
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
