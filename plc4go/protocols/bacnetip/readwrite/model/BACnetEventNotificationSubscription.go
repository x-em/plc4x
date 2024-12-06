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

// BACnetEventNotificationSubscription is the corresponding interface of BACnetEventNotificationSubscription
type BACnetEventNotificationSubscription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetRecipient returns Recipient (property field)
	GetRecipient() BACnetRecipientEnclosed
	// GetProcessIdentifier returns ProcessIdentifier (property field)
	GetProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetIssueConfirmedNotifications returns IssueConfirmedNotifications (property field)
	GetIssueConfirmedNotifications() BACnetContextTagBoolean
	// GetTimeRemaining returns TimeRemaining (property field)
	GetTimeRemaining() BACnetContextTagUnsignedInteger
	// IsBACnetEventNotificationSubscription is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventNotificationSubscription()
	// CreateBuilder creates a BACnetEventNotificationSubscriptionBuilder
	CreateBACnetEventNotificationSubscriptionBuilder() BACnetEventNotificationSubscriptionBuilder
}

// _BACnetEventNotificationSubscription is the data-structure of this message
type _BACnetEventNotificationSubscription struct {
	Recipient                   BACnetRecipientEnclosed
	ProcessIdentifier           BACnetContextTagUnsignedInteger
	IssueConfirmedNotifications BACnetContextTagBoolean
	TimeRemaining               BACnetContextTagUnsignedInteger
}

var _ BACnetEventNotificationSubscription = (*_BACnetEventNotificationSubscription)(nil)

// NewBACnetEventNotificationSubscription factory function for _BACnetEventNotificationSubscription
func NewBACnetEventNotificationSubscription(recipient BACnetRecipientEnclosed, processIdentifier BACnetContextTagUnsignedInteger, issueConfirmedNotifications BACnetContextTagBoolean, timeRemaining BACnetContextTagUnsignedInteger) *_BACnetEventNotificationSubscription {
	if recipient == nil {
		panic("recipient of type BACnetRecipientEnclosed for BACnetEventNotificationSubscription must not be nil")
	}
	if processIdentifier == nil {
		panic("processIdentifier of type BACnetContextTagUnsignedInteger for BACnetEventNotificationSubscription must not be nil")
	}
	if timeRemaining == nil {
		panic("timeRemaining of type BACnetContextTagUnsignedInteger for BACnetEventNotificationSubscription must not be nil")
	}
	return &_BACnetEventNotificationSubscription{Recipient: recipient, ProcessIdentifier: processIdentifier, IssueConfirmedNotifications: issueConfirmedNotifications, TimeRemaining: timeRemaining}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventNotificationSubscriptionBuilder is a builder for BACnetEventNotificationSubscription
type BACnetEventNotificationSubscriptionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(recipient BACnetRecipientEnclosed, processIdentifier BACnetContextTagUnsignedInteger, timeRemaining BACnetContextTagUnsignedInteger) BACnetEventNotificationSubscriptionBuilder
	// WithRecipient adds Recipient (property field)
	WithRecipient(BACnetRecipientEnclosed) BACnetEventNotificationSubscriptionBuilder
	// WithRecipientBuilder adds Recipient (property field) which is build by the builder
	WithRecipientBuilder(func(BACnetRecipientEnclosedBuilder) BACnetRecipientEnclosedBuilder) BACnetEventNotificationSubscriptionBuilder
	// WithProcessIdentifier adds ProcessIdentifier (property field)
	WithProcessIdentifier(BACnetContextTagUnsignedInteger) BACnetEventNotificationSubscriptionBuilder
	// WithProcessIdentifierBuilder adds ProcessIdentifier (property field) which is build by the builder
	WithProcessIdentifierBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventNotificationSubscriptionBuilder
	// WithIssueConfirmedNotifications adds IssueConfirmedNotifications (property field)
	WithOptionalIssueConfirmedNotifications(BACnetContextTagBoolean) BACnetEventNotificationSubscriptionBuilder
	// WithOptionalIssueConfirmedNotificationsBuilder adds IssueConfirmedNotifications (property field) which is build by the builder
	WithOptionalIssueConfirmedNotificationsBuilder(func(BACnetContextTagBooleanBuilder) BACnetContextTagBooleanBuilder) BACnetEventNotificationSubscriptionBuilder
	// WithTimeRemaining adds TimeRemaining (property field)
	WithTimeRemaining(BACnetContextTagUnsignedInteger) BACnetEventNotificationSubscriptionBuilder
	// WithTimeRemainingBuilder adds TimeRemaining (property field) which is build by the builder
	WithTimeRemainingBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventNotificationSubscriptionBuilder
	// Build builds the BACnetEventNotificationSubscription or returns an error if something is wrong
	Build() (BACnetEventNotificationSubscription, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventNotificationSubscription
}

// NewBACnetEventNotificationSubscriptionBuilder() creates a BACnetEventNotificationSubscriptionBuilder
func NewBACnetEventNotificationSubscriptionBuilder() BACnetEventNotificationSubscriptionBuilder {
	return &_BACnetEventNotificationSubscriptionBuilder{_BACnetEventNotificationSubscription: new(_BACnetEventNotificationSubscription)}
}

type _BACnetEventNotificationSubscriptionBuilder struct {
	*_BACnetEventNotificationSubscription

	err *utils.MultiError
}

var _ (BACnetEventNotificationSubscriptionBuilder) = (*_BACnetEventNotificationSubscriptionBuilder)(nil)

func (b *_BACnetEventNotificationSubscriptionBuilder) WithMandatoryFields(recipient BACnetRecipientEnclosed, processIdentifier BACnetContextTagUnsignedInteger, timeRemaining BACnetContextTagUnsignedInteger) BACnetEventNotificationSubscriptionBuilder {
	return b.WithRecipient(recipient).WithProcessIdentifier(processIdentifier).WithTimeRemaining(timeRemaining)
}

func (b *_BACnetEventNotificationSubscriptionBuilder) WithRecipient(recipient BACnetRecipientEnclosed) BACnetEventNotificationSubscriptionBuilder {
	b.Recipient = recipient
	return b
}

func (b *_BACnetEventNotificationSubscriptionBuilder) WithRecipientBuilder(builderSupplier func(BACnetRecipientEnclosedBuilder) BACnetRecipientEnclosedBuilder) BACnetEventNotificationSubscriptionBuilder {
	builder := builderSupplier(b.Recipient.CreateBACnetRecipientEnclosedBuilder())
	var err error
	b.Recipient, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetRecipientEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetEventNotificationSubscriptionBuilder) WithProcessIdentifier(processIdentifier BACnetContextTagUnsignedInteger) BACnetEventNotificationSubscriptionBuilder {
	b.ProcessIdentifier = processIdentifier
	return b
}

func (b *_BACnetEventNotificationSubscriptionBuilder) WithProcessIdentifierBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventNotificationSubscriptionBuilder {
	builder := builderSupplier(b.ProcessIdentifier.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.ProcessIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetEventNotificationSubscriptionBuilder) WithOptionalIssueConfirmedNotifications(issueConfirmedNotifications BACnetContextTagBoolean) BACnetEventNotificationSubscriptionBuilder {
	b.IssueConfirmedNotifications = issueConfirmedNotifications
	return b
}

func (b *_BACnetEventNotificationSubscriptionBuilder) WithOptionalIssueConfirmedNotificationsBuilder(builderSupplier func(BACnetContextTagBooleanBuilder) BACnetContextTagBooleanBuilder) BACnetEventNotificationSubscriptionBuilder {
	builder := builderSupplier(b.IssueConfirmedNotifications.CreateBACnetContextTagBooleanBuilder())
	var err error
	b.IssueConfirmedNotifications, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetEventNotificationSubscriptionBuilder) WithTimeRemaining(timeRemaining BACnetContextTagUnsignedInteger) BACnetEventNotificationSubscriptionBuilder {
	b.TimeRemaining = timeRemaining
	return b
}

func (b *_BACnetEventNotificationSubscriptionBuilder) WithTimeRemainingBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventNotificationSubscriptionBuilder {
	builder := builderSupplier(b.TimeRemaining.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.TimeRemaining, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetEventNotificationSubscriptionBuilder) Build() (BACnetEventNotificationSubscription, error) {
	if b.Recipient == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'recipient' not set"))
	}
	if b.ProcessIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'processIdentifier' not set"))
	}
	if b.TimeRemaining == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeRemaining' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetEventNotificationSubscription.deepCopy(), nil
}

func (b *_BACnetEventNotificationSubscriptionBuilder) MustBuild() BACnetEventNotificationSubscription {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetEventNotificationSubscriptionBuilder) DeepCopy() any {
	_copy := b.CreateBACnetEventNotificationSubscriptionBuilder().(*_BACnetEventNotificationSubscriptionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetEventNotificationSubscriptionBuilder creates a BACnetEventNotificationSubscriptionBuilder
func (b *_BACnetEventNotificationSubscription) CreateBACnetEventNotificationSubscriptionBuilder() BACnetEventNotificationSubscriptionBuilder {
	if b == nil {
		return NewBACnetEventNotificationSubscriptionBuilder()
	}
	return &_BACnetEventNotificationSubscriptionBuilder{_BACnetEventNotificationSubscription: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventNotificationSubscription) GetRecipient() BACnetRecipientEnclosed {
	return m.Recipient
}

func (m *_BACnetEventNotificationSubscription) GetProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.ProcessIdentifier
}

func (m *_BACnetEventNotificationSubscription) GetIssueConfirmedNotifications() BACnetContextTagBoolean {
	return m.IssueConfirmedNotifications
}

func (m *_BACnetEventNotificationSubscription) GetTimeRemaining() BACnetContextTagUnsignedInteger {
	return m.TimeRemaining
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventNotificationSubscription(structType any) BACnetEventNotificationSubscription {
	if casted, ok := structType.(BACnetEventNotificationSubscription); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventNotificationSubscription); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventNotificationSubscription) GetTypeName() string {
	return "BACnetEventNotificationSubscription"
}

func (m *_BACnetEventNotificationSubscription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (recipient)
	lengthInBits += m.Recipient.GetLengthInBits(ctx)

	// Simple field (processIdentifier)
	lengthInBits += m.ProcessIdentifier.GetLengthInBits(ctx)

	// Optional Field (issueConfirmedNotifications)
	if m.IssueConfirmedNotifications != nil {
		lengthInBits += m.IssueConfirmedNotifications.GetLengthInBits(ctx)
	}

	// Simple field (timeRemaining)
	lengthInBits += m.TimeRemaining.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventNotificationSubscription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventNotificationSubscriptionParse(ctx context.Context, theBytes []byte) (BACnetEventNotificationSubscription, error) {
	return BACnetEventNotificationSubscriptionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetEventNotificationSubscriptionParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventNotificationSubscription, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventNotificationSubscription, error) {
		return BACnetEventNotificationSubscriptionParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetEventNotificationSubscriptionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventNotificationSubscription, error) {
	v, err := (&_BACnetEventNotificationSubscription{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetEventNotificationSubscription) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetEventNotificationSubscription BACnetEventNotificationSubscription, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventNotificationSubscription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventNotificationSubscription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	recipient, err := ReadSimpleField[BACnetRecipientEnclosed](ctx, "recipient", ReadComplex[BACnetRecipientEnclosed](BACnetRecipientEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recipient' field"))
	}
	m.Recipient = recipient

	processIdentifier, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "processIdentifier", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'processIdentifier' field"))
	}
	m.ProcessIdentifier = processIdentifier

	var issueConfirmedNotifications BACnetContextTagBoolean
	_issueConfirmedNotifications, err := ReadOptionalField[BACnetContextTagBoolean](ctx, "issueConfirmedNotifications", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'issueConfirmedNotifications' field"))
	}
	if _issueConfirmedNotifications != nil {
		issueConfirmedNotifications = *_issueConfirmedNotifications
		m.IssueConfirmedNotifications = issueConfirmedNotifications
	}

	timeRemaining, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeRemaining", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeRemaining' field"))
	}
	m.TimeRemaining = timeRemaining

	if closeErr := readBuffer.CloseContext("BACnetEventNotificationSubscription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventNotificationSubscription")
	}

	return m, nil
}

func (m *_BACnetEventNotificationSubscription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventNotificationSubscription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventNotificationSubscription"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventNotificationSubscription")
	}

	if err := WriteSimpleField[BACnetRecipientEnclosed](ctx, "recipient", m.GetRecipient(), WriteComplex[BACnetRecipientEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'recipient' field")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "processIdentifier", m.GetProcessIdentifier(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'processIdentifier' field")
	}

	if err := WriteOptionalField[BACnetContextTagBoolean](ctx, "issueConfirmedNotifications", GetRef(m.GetIssueConfirmedNotifications()), WriteComplex[BACnetContextTagBoolean](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'issueConfirmedNotifications' field")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeRemaining", m.GetTimeRemaining(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'timeRemaining' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventNotificationSubscription"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventNotificationSubscription")
	}
	return nil
}

func (m *_BACnetEventNotificationSubscription) IsBACnetEventNotificationSubscription() {}

func (m *_BACnetEventNotificationSubscription) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventNotificationSubscription) deepCopy() *_BACnetEventNotificationSubscription {
	if m == nil {
		return nil
	}
	_BACnetEventNotificationSubscriptionCopy := &_BACnetEventNotificationSubscription{
		utils.DeepCopy[BACnetRecipientEnclosed](m.Recipient),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.ProcessIdentifier),
		utils.DeepCopy[BACnetContextTagBoolean](m.IssueConfirmedNotifications),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.TimeRemaining),
	}
	return _BACnetEventNotificationSubscriptionCopy
}

func (m *_BACnetEventNotificationSubscription) String() string {
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