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

// BACnetDestination is the corresponding interface of BACnetDestination
type BACnetDestination interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetValidDays returns ValidDays (property field)
	GetValidDays() BACnetDaysOfWeekTagged
	// GetFromTime returns FromTime (property field)
	GetFromTime() BACnetApplicationTagTime
	// GetToTime returns ToTime (property field)
	GetToTime() BACnetApplicationTagTime
	// GetRecipient returns Recipient (property field)
	GetRecipient() BACnetRecipient
	// GetProcessIdentifier returns ProcessIdentifier (property field)
	GetProcessIdentifier() BACnetApplicationTagUnsignedInteger
	// GetIssueConfirmedNotifications returns IssueConfirmedNotifications (property field)
	GetIssueConfirmedNotifications() BACnetApplicationTagBoolean
	// GetTransitions returns Transitions (property field)
	GetTransitions() BACnetEventTransitionBitsTagged
	// IsBACnetDestination is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetDestination()
	// CreateBuilder creates a BACnetDestinationBuilder
	CreateBACnetDestinationBuilder() BACnetDestinationBuilder
}

// _BACnetDestination is the data-structure of this message
type _BACnetDestination struct {
	ValidDays                   BACnetDaysOfWeekTagged
	FromTime                    BACnetApplicationTagTime
	ToTime                      BACnetApplicationTagTime
	Recipient                   BACnetRecipient
	ProcessIdentifier           BACnetApplicationTagUnsignedInteger
	IssueConfirmedNotifications BACnetApplicationTagBoolean
	Transitions                 BACnetEventTransitionBitsTagged
}

var _ BACnetDestination = (*_BACnetDestination)(nil)

// NewBACnetDestination factory function for _BACnetDestination
func NewBACnetDestination(validDays BACnetDaysOfWeekTagged, fromTime BACnetApplicationTagTime, toTime BACnetApplicationTagTime, recipient BACnetRecipient, processIdentifier BACnetApplicationTagUnsignedInteger, issueConfirmedNotifications BACnetApplicationTagBoolean, transitions BACnetEventTransitionBitsTagged) *_BACnetDestination {
	if validDays == nil {
		panic("validDays of type BACnetDaysOfWeekTagged for BACnetDestination must not be nil")
	}
	if fromTime == nil {
		panic("fromTime of type BACnetApplicationTagTime for BACnetDestination must not be nil")
	}
	if toTime == nil {
		panic("toTime of type BACnetApplicationTagTime for BACnetDestination must not be nil")
	}
	if recipient == nil {
		panic("recipient of type BACnetRecipient for BACnetDestination must not be nil")
	}
	if processIdentifier == nil {
		panic("processIdentifier of type BACnetApplicationTagUnsignedInteger for BACnetDestination must not be nil")
	}
	if issueConfirmedNotifications == nil {
		panic("issueConfirmedNotifications of type BACnetApplicationTagBoolean for BACnetDestination must not be nil")
	}
	if transitions == nil {
		panic("transitions of type BACnetEventTransitionBitsTagged for BACnetDestination must not be nil")
	}
	return &_BACnetDestination{ValidDays: validDays, FromTime: fromTime, ToTime: toTime, Recipient: recipient, ProcessIdentifier: processIdentifier, IssueConfirmedNotifications: issueConfirmedNotifications, Transitions: transitions}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetDestinationBuilder is a builder for BACnetDestination
type BACnetDestinationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(validDays BACnetDaysOfWeekTagged, fromTime BACnetApplicationTagTime, toTime BACnetApplicationTagTime, recipient BACnetRecipient, processIdentifier BACnetApplicationTagUnsignedInteger, issueConfirmedNotifications BACnetApplicationTagBoolean, transitions BACnetEventTransitionBitsTagged) BACnetDestinationBuilder
	// WithValidDays adds ValidDays (property field)
	WithValidDays(BACnetDaysOfWeekTagged) BACnetDestinationBuilder
	// WithValidDaysBuilder adds ValidDays (property field) which is build by the builder
	WithValidDaysBuilder(func(BACnetDaysOfWeekTaggedBuilder) BACnetDaysOfWeekTaggedBuilder) BACnetDestinationBuilder
	// WithFromTime adds FromTime (property field)
	WithFromTime(BACnetApplicationTagTime) BACnetDestinationBuilder
	// WithFromTimeBuilder adds FromTime (property field) which is build by the builder
	WithFromTimeBuilder(func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetDestinationBuilder
	// WithToTime adds ToTime (property field)
	WithToTime(BACnetApplicationTagTime) BACnetDestinationBuilder
	// WithToTimeBuilder adds ToTime (property field) which is build by the builder
	WithToTimeBuilder(func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetDestinationBuilder
	// WithRecipient adds Recipient (property field)
	WithRecipient(BACnetRecipient) BACnetDestinationBuilder
	// WithRecipientBuilder adds Recipient (property field) which is build by the builder
	WithRecipientBuilder(func(BACnetRecipientBuilder) BACnetRecipientBuilder) BACnetDestinationBuilder
	// WithProcessIdentifier adds ProcessIdentifier (property field)
	WithProcessIdentifier(BACnetApplicationTagUnsignedInteger) BACnetDestinationBuilder
	// WithProcessIdentifierBuilder adds ProcessIdentifier (property field) which is build by the builder
	WithProcessIdentifierBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetDestinationBuilder
	// WithIssueConfirmedNotifications adds IssueConfirmedNotifications (property field)
	WithIssueConfirmedNotifications(BACnetApplicationTagBoolean) BACnetDestinationBuilder
	// WithIssueConfirmedNotificationsBuilder adds IssueConfirmedNotifications (property field) which is build by the builder
	WithIssueConfirmedNotificationsBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetDestinationBuilder
	// WithTransitions adds Transitions (property field)
	WithTransitions(BACnetEventTransitionBitsTagged) BACnetDestinationBuilder
	// WithTransitionsBuilder adds Transitions (property field) which is build by the builder
	WithTransitionsBuilder(func(BACnetEventTransitionBitsTaggedBuilder) BACnetEventTransitionBitsTaggedBuilder) BACnetDestinationBuilder
	// Build builds the BACnetDestination or returns an error if something is wrong
	Build() (BACnetDestination, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetDestination
}

// NewBACnetDestinationBuilder() creates a BACnetDestinationBuilder
func NewBACnetDestinationBuilder() BACnetDestinationBuilder {
	return &_BACnetDestinationBuilder{_BACnetDestination: new(_BACnetDestination)}
}

type _BACnetDestinationBuilder struct {
	*_BACnetDestination

	err *utils.MultiError
}

var _ (BACnetDestinationBuilder) = (*_BACnetDestinationBuilder)(nil)

func (b *_BACnetDestinationBuilder) WithMandatoryFields(validDays BACnetDaysOfWeekTagged, fromTime BACnetApplicationTagTime, toTime BACnetApplicationTagTime, recipient BACnetRecipient, processIdentifier BACnetApplicationTagUnsignedInteger, issueConfirmedNotifications BACnetApplicationTagBoolean, transitions BACnetEventTransitionBitsTagged) BACnetDestinationBuilder {
	return b.WithValidDays(validDays).WithFromTime(fromTime).WithToTime(toTime).WithRecipient(recipient).WithProcessIdentifier(processIdentifier).WithIssueConfirmedNotifications(issueConfirmedNotifications).WithTransitions(transitions)
}

func (b *_BACnetDestinationBuilder) WithValidDays(validDays BACnetDaysOfWeekTagged) BACnetDestinationBuilder {
	b.ValidDays = validDays
	return b
}

func (b *_BACnetDestinationBuilder) WithValidDaysBuilder(builderSupplier func(BACnetDaysOfWeekTaggedBuilder) BACnetDaysOfWeekTaggedBuilder) BACnetDestinationBuilder {
	builder := builderSupplier(b.ValidDays.CreateBACnetDaysOfWeekTaggedBuilder())
	var err error
	b.ValidDays, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDaysOfWeekTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetDestinationBuilder) WithFromTime(fromTime BACnetApplicationTagTime) BACnetDestinationBuilder {
	b.FromTime = fromTime
	return b
}

func (b *_BACnetDestinationBuilder) WithFromTimeBuilder(builderSupplier func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetDestinationBuilder {
	builder := builderSupplier(b.FromTime.CreateBACnetApplicationTagTimeBuilder())
	var err error
	b.FromTime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagTimeBuilder failed"))
	}
	return b
}

func (b *_BACnetDestinationBuilder) WithToTime(toTime BACnetApplicationTagTime) BACnetDestinationBuilder {
	b.ToTime = toTime
	return b
}

func (b *_BACnetDestinationBuilder) WithToTimeBuilder(builderSupplier func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetDestinationBuilder {
	builder := builderSupplier(b.ToTime.CreateBACnetApplicationTagTimeBuilder())
	var err error
	b.ToTime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagTimeBuilder failed"))
	}
	return b
}

func (b *_BACnetDestinationBuilder) WithRecipient(recipient BACnetRecipient) BACnetDestinationBuilder {
	b.Recipient = recipient
	return b
}

func (b *_BACnetDestinationBuilder) WithRecipientBuilder(builderSupplier func(BACnetRecipientBuilder) BACnetRecipientBuilder) BACnetDestinationBuilder {
	builder := builderSupplier(b.Recipient.CreateBACnetRecipientBuilder())
	var err error
	b.Recipient, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetRecipientBuilder failed"))
	}
	return b
}

func (b *_BACnetDestinationBuilder) WithProcessIdentifier(processIdentifier BACnetApplicationTagUnsignedInteger) BACnetDestinationBuilder {
	b.ProcessIdentifier = processIdentifier
	return b
}

func (b *_BACnetDestinationBuilder) WithProcessIdentifierBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetDestinationBuilder {
	builder := builderSupplier(b.ProcessIdentifier.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.ProcessIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetDestinationBuilder) WithIssueConfirmedNotifications(issueConfirmedNotifications BACnetApplicationTagBoolean) BACnetDestinationBuilder {
	b.IssueConfirmedNotifications = issueConfirmedNotifications
	return b
}

func (b *_BACnetDestinationBuilder) WithIssueConfirmedNotificationsBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetDestinationBuilder {
	builder := builderSupplier(b.IssueConfirmedNotifications.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.IssueConfirmedNotifications, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetDestinationBuilder) WithTransitions(transitions BACnetEventTransitionBitsTagged) BACnetDestinationBuilder {
	b.Transitions = transitions
	return b
}

func (b *_BACnetDestinationBuilder) WithTransitionsBuilder(builderSupplier func(BACnetEventTransitionBitsTaggedBuilder) BACnetEventTransitionBitsTaggedBuilder) BACnetDestinationBuilder {
	builder := builderSupplier(b.Transitions.CreateBACnetEventTransitionBitsTaggedBuilder())
	var err error
	b.Transitions, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetEventTransitionBitsTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetDestinationBuilder) Build() (BACnetDestination, error) {
	if b.ValidDays == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'validDays' not set"))
	}
	if b.FromTime == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'fromTime' not set"))
	}
	if b.ToTime == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'toTime' not set"))
	}
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
	if b.IssueConfirmedNotifications == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'issueConfirmedNotifications' not set"))
	}
	if b.Transitions == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'transitions' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetDestination.deepCopy(), nil
}

func (b *_BACnetDestinationBuilder) MustBuild() BACnetDestination {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetDestinationBuilder) DeepCopy() any {
	_copy := b.CreateBACnetDestinationBuilder().(*_BACnetDestinationBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetDestinationBuilder creates a BACnetDestinationBuilder
func (b *_BACnetDestination) CreateBACnetDestinationBuilder() BACnetDestinationBuilder {
	if b == nil {
		return NewBACnetDestinationBuilder()
	}
	return &_BACnetDestinationBuilder{_BACnetDestination: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDestination) GetValidDays() BACnetDaysOfWeekTagged {
	return m.ValidDays
}

func (m *_BACnetDestination) GetFromTime() BACnetApplicationTagTime {
	return m.FromTime
}

func (m *_BACnetDestination) GetToTime() BACnetApplicationTagTime {
	return m.ToTime
}

func (m *_BACnetDestination) GetRecipient() BACnetRecipient {
	return m.Recipient
}

func (m *_BACnetDestination) GetProcessIdentifier() BACnetApplicationTagUnsignedInteger {
	return m.ProcessIdentifier
}

func (m *_BACnetDestination) GetIssueConfirmedNotifications() BACnetApplicationTagBoolean {
	return m.IssueConfirmedNotifications
}

func (m *_BACnetDestination) GetTransitions() BACnetEventTransitionBitsTagged {
	return m.Transitions
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetDestination(structType any) BACnetDestination {
	if casted, ok := structType.(BACnetDestination); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDestination); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDestination) GetTypeName() string {
	return "BACnetDestination"
}

func (m *_BACnetDestination) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (validDays)
	lengthInBits += m.ValidDays.GetLengthInBits(ctx)

	// Simple field (fromTime)
	lengthInBits += m.FromTime.GetLengthInBits(ctx)

	// Simple field (toTime)
	lengthInBits += m.ToTime.GetLengthInBits(ctx)

	// Simple field (recipient)
	lengthInBits += m.Recipient.GetLengthInBits(ctx)

	// Simple field (processIdentifier)
	lengthInBits += m.ProcessIdentifier.GetLengthInBits(ctx)

	// Simple field (issueConfirmedNotifications)
	lengthInBits += m.IssueConfirmedNotifications.GetLengthInBits(ctx)

	// Simple field (transitions)
	lengthInBits += m.Transitions.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetDestination) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDestinationParse(ctx context.Context, theBytes []byte) (BACnetDestination, error) {
	return BACnetDestinationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetDestinationParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDestination, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDestination, error) {
		return BACnetDestinationParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetDestinationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDestination, error) {
	v, err := (&_BACnetDestination{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetDestination) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetDestination BACnetDestination, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDestination"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDestination")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	validDays, err := ReadSimpleField[BACnetDaysOfWeekTagged](ctx, "validDays", ReadComplex[BACnetDaysOfWeekTagged](BACnetDaysOfWeekTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'validDays' field"))
	}
	m.ValidDays = validDays

	fromTime, err := ReadSimpleField[BACnetApplicationTagTime](ctx, "fromTime", ReadComplex[BACnetApplicationTagTime](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagTime](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fromTime' field"))
	}
	m.FromTime = fromTime

	toTime, err := ReadSimpleField[BACnetApplicationTagTime](ctx, "toTime", ReadComplex[BACnetApplicationTagTime](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagTime](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toTime' field"))
	}
	m.ToTime = toTime

	recipient, err := ReadSimpleField[BACnetRecipient](ctx, "recipient", ReadComplex[BACnetRecipient](BACnetRecipientParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recipient' field"))
	}
	m.Recipient = recipient

	processIdentifier, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "processIdentifier", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'processIdentifier' field"))
	}
	m.ProcessIdentifier = processIdentifier

	issueConfirmedNotifications, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "issueConfirmedNotifications", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'issueConfirmedNotifications' field"))
	}
	m.IssueConfirmedNotifications = issueConfirmedNotifications

	transitions, err := ReadSimpleField[BACnetEventTransitionBitsTagged](ctx, "transitions", ReadComplex[BACnetEventTransitionBitsTagged](BACnetEventTransitionBitsTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transitions' field"))
	}
	m.Transitions = transitions

	if closeErr := readBuffer.CloseContext("BACnetDestination"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDestination")
	}

	return m, nil
}

func (m *_BACnetDestination) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDestination) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetDestination"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDestination")
	}

	if err := WriteSimpleField[BACnetDaysOfWeekTagged](ctx, "validDays", m.GetValidDays(), WriteComplex[BACnetDaysOfWeekTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'validDays' field")
	}

	if err := WriteSimpleField[BACnetApplicationTagTime](ctx, "fromTime", m.GetFromTime(), WriteComplex[BACnetApplicationTagTime](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'fromTime' field")
	}

	if err := WriteSimpleField[BACnetApplicationTagTime](ctx, "toTime", m.GetToTime(), WriteComplex[BACnetApplicationTagTime](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'toTime' field")
	}

	if err := WriteSimpleField[BACnetRecipient](ctx, "recipient", m.GetRecipient(), WriteComplex[BACnetRecipient](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'recipient' field")
	}

	if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "processIdentifier", m.GetProcessIdentifier(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'processIdentifier' field")
	}

	if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "issueConfirmedNotifications", m.GetIssueConfirmedNotifications(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'issueConfirmedNotifications' field")
	}

	if err := WriteSimpleField[BACnetEventTransitionBitsTagged](ctx, "transitions", m.GetTransitions(), WriteComplex[BACnetEventTransitionBitsTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'transitions' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDestination"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDestination")
	}
	return nil
}

func (m *_BACnetDestination) IsBACnetDestination() {}

func (m *_BACnetDestination) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetDestination) deepCopy() *_BACnetDestination {
	if m == nil {
		return nil
	}
	_BACnetDestinationCopy := &_BACnetDestination{
		utils.DeepCopy[BACnetDaysOfWeekTagged](m.ValidDays),
		utils.DeepCopy[BACnetApplicationTagTime](m.FromTime),
		utils.DeepCopy[BACnetApplicationTagTime](m.ToTime),
		utils.DeepCopy[BACnetRecipient](m.Recipient),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.ProcessIdentifier),
		utils.DeepCopy[BACnetApplicationTagBoolean](m.IssueConfirmedNotifications),
		utils.DeepCopy[BACnetEventTransitionBitsTagged](m.Transitions),
	}
	return _BACnetDestinationCopy
}

func (m *_BACnetDestination) String() string {
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
