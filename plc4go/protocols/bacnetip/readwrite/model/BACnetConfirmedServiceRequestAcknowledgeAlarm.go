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

// BACnetConfirmedServiceRequestAcknowledgeAlarm is the corresponding interface of BACnetConfirmedServiceRequestAcknowledgeAlarm
type BACnetConfirmedServiceRequestAcknowledgeAlarm interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetAcknowledgingProcessIdentifier returns AcknowledgingProcessIdentifier (property field)
	GetAcknowledgingProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetEventObjectIdentifier returns EventObjectIdentifier (property field)
	GetEventObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetEventStateAcknowledged returns EventStateAcknowledged (property field)
	GetEventStateAcknowledged() BACnetEventStateTagged
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetTimeStampEnclosed
	// GetAcknowledgmentSource returns AcknowledgmentSource (property field)
	GetAcknowledgmentSource() BACnetContextTagCharacterString
	// GetTimeOfAcknowledgment returns TimeOfAcknowledgment (property field)
	GetTimeOfAcknowledgment() BACnetTimeStampEnclosed
	// IsBACnetConfirmedServiceRequestAcknowledgeAlarm is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestAcknowledgeAlarm()
	// CreateBuilder creates a BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder
	CreateBACnetConfirmedServiceRequestAcknowledgeAlarmBuilder() BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder
}

// _BACnetConfirmedServiceRequestAcknowledgeAlarm is the data-structure of this message
type _BACnetConfirmedServiceRequestAcknowledgeAlarm struct {
	BACnetConfirmedServiceRequestContract
	AcknowledgingProcessIdentifier BACnetContextTagUnsignedInteger
	EventObjectIdentifier          BACnetContextTagObjectIdentifier
	EventStateAcknowledged         BACnetEventStateTagged
	Timestamp                      BACnetTimeStampEnclosed
	AcknowledgmentSource           BACnetContextTagCharacterString
	TimeOfAcknowledgment           BACnetTimeStampEnclosed
}

var _ BACnetConfirmedServiceRequestAcknowledgeAlarm = (*_BACnetConfirmedServiceRequestAcknowledgeAlarm)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestAcknowledgeAlarm)(nil)

// NewBACnetConfirmedServiceRequestAcknowledgeAlarm factory function for _BACnetConfirmedServiceRequestAcknowledgeAlarm
func NewBACnetConfirmedServiceRequestAcknowledgeAlarm(acknowledgingProcessIdentifier BACnetContextTagUnsignedInteger, eventObjectIdentifier BACnetContextTagObjectIdentifier, eventStateAcknowledged BACnetEventStateTagged, timestamp BACnetTimeStampEnclosed, acknowledgmentSource BACnetContextTagCharacterString, timeOfAcknowledgment BACnetTimeStampEnclosed, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestAcknowledgeAlarm {
	if acknowledgingProcessIdentifier == nil {
		panic("acknowledgingProcessIdentifier of type BACnetContextTagUnsignedInteger for BACnetConfirmedServiceRequestAcknowledgeAlarm must not be nil")
	}
	if eventObjectIdentifier == nil {
		panic("eventObjectIdentifier of type BACnetContextTagObjectIdentifier for BACnetConfirmedServiceRequestAcknowledgeAlarm must not be nil")
	}
	if eventStateAcknowledged == nil {
		panic("eventStateAcknowledged of type BACnetEventStateTagged for BACnetConfirmedServiceRequestAcknowledgeAlarm must not be nil")
	}
	if timestamp == nil {
		panic("timestamp of type BACnetTimeStampEnclosed for BACnetConfirmedServiceRequestAcknowledgeAlarm must not be nil")
	}
	if acknowledgmentSource == nil {
		panic("acknowledgmentSource of type BACnetContextTagCharacterString for BACnetConfirmedServiceRequestAcknowledgeAlarm must not be nil")
	}
	if timeOfAcknowledgment == nil {
		panic("timeOfAcknowledgment of type BACnetTimeStampEnclosed for BACnetConfirmedServiceRequestAcknowledgeAlarm must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestAcknowledgeAlarm{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		AcknowledgingProcessIdentifier:        acknowledgingProcessIdentifier,
		EventObjectIdentifier:                 eventObjectIdentifier,
		EventStateAcknowledged:                eventStateAcknowledged,
		Timestamp:                             timestamp,
		AcknowledgmentSource:                  acknowledgmentSource,
		TimeOfAcknowledgment:                  timeOfAcknowledgment,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder is a builder for BACnetConfirmedServiceRequestAcknowledgeAlarm
type BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(acknowledgingProcessIdentifier BACnetContextTagUnsignedInteger, eventObjectIdentifier BACnetContextTagObjectIdentifier, eventStateAcknowledged BACnetEventStateTagged, timestamp BACnetTimeStampEnclosed, acknowledgmentSource BACnetContextTagCharacterString, timeOfAcknowledgment BACnetTimeStampEnclosed) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder
	// WithAcknowledgingProcessIdentifier adds AcknowledgingProcessIdentifier (property field)
	WithAcknowledgingProcessIdentifier(BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder
	// WithAcknowledgingProcessIdentifierBuilder adds AcknowledgingProcessIdentifier (property field) which is build by the builder
	WithAcknowledgingProcessIdentifierBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder
	// WithEventObjectIdentifier adds EventObjectIdentifier (property field)
	WithEventObjectIdentifier(BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder
	// WithEventObjectIdentifierBuilder adds EventObjectIdentifier (property field) which is build by the builder
	WithEventObjectIdentifierBuilder(func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder
	// WithEventStateAcknowledged adds EventStateAcknowledged (property field)
	WithEventStateAcknowledged(BACnetEventStateTagged) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder
	// WithEventStateAcknowledgedBuilder adds EventStateAcknowledged (property field) which is build by the builder
	WithEventStateAcknowledgedBuilder(func(BACnetEventStateTaggedBuilder) BACnetEventStateTaggedBuilder) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder
	// WithTimestamp adds Timestamp (property field)
	WithTimestamp(BACnetTimeStampEnclosed) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder
	// WithTimestampBuilder adds Timestamp (property field) which is build by the builder
	WithTimestampBuilder(func(BACnetTimeStampEnclosedBuilder) BACnetTimeStampEnclosedBuilder) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder
	// WithAcknowledgmentSource adds AcknowledgmentSource (property field)
	WithAcknowledgmentSource(BACnetContextTagCharacterString) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder
	// WithAcknowledgmentSourceBuilder adds AcknowledgmentSource (property field) which is build by the builder
	WithAcknowledgmentSourceBuilder(func(BACnetContextTagCharacterStringBuilder) BACnetContextTagCharacterStringBuilder) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder
	// WithTimeOfAcknowledgment adds TimeOfAcknowledgment (property field)
	WithTimeOfAcknowledgment(BACnetTimeStampEnclosed) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder
	// WithTimeOfAcknowledgmentBuilder adds TimeOfAcknowledgment (property field) which is build by the builder
	WithTimeOfAcknowledgmentBuilder(func(BACnetTimeStampEnclosedBuilder) BACnetTimeStampEnclosedBuilder) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConfirmedServiceRequestBuilder
	// Build builds the BACnetConfirmedServiceRequestAcknowledgeAlarm or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestAcknowledgeAlarm, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestAcknowledgeAlarm
}

// NewBACnetConfirmedServiceRequestAcknowledgeAlarmBuilder() creates a BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder
func NewBACnetConfirmedServiceRequestAcknowledgeAlarmBuilder() BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder {
	return &_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder{_BACnetConfirmedServiceRequestAcknowledgeAlarm: new(_BACnetConfirmedServiceRequestAcknowledgeAlarm)}
}

type _BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder struct {
	*_BACnetConfirmedServiceRequestAcknowledgeAlarm

	parentBuilder *_BACnetConfirmedServiceRequestBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder) = (*_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder) setParent(contract BACnetConfirmedServiceRequestContract) {
	b.BACnetConfirmedServiceRequestContract = contract
	contract.(*_BACnetConfirmedServiceRequest)._SubType = b._BACnetConfirmedServiceRequestAcknowledgeAlarm
}

func (b *_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder) WithMandatoryFields(acknowledgingProcessIdentifier BACnetContextTagUnsignedInteger, eventObjectIdentifier BACnetContextTagObjectIdentifier, eventStateAcknowledged BACnetEventStateTagged, timestamp BACnetTimeStampEnclosed, acknowledgmentSource BACnetContextTagCharacterString, timeOfAcknowledgment BACnetTimeStampEnclosed) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder {
	return b.WithAcknowledgingProcessIdentifier(acknowledgingProcessIdentifier).WithEventObjectIdentifier(eventObjectIdentifier).WithEventStateAcknowledged(eventStateAcknowledged).WithTimestamp(timestamp).WithAcknowledgmentSource(acknowledgmentSource).WithTimeOfAcknowledgment(timeOfAcknowledgment)
}

func (b *_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder) WithAcknowledgingProcessIdentifier(acknowledgingProcessIdentifier BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder {
	b.AcknowledgingProcessIdentifier = acknowledgingProcessIdentifier
	return b
}

func (b *_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder) WithAcknowledgingProcessIdentifierBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder {
	builder := builderSupplier(b.AcknowledgingProcessIdentifier.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.AcknowledgingProcessIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder) WithEventObjectIdentifier(eventObjectIdentifier BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder {
	b.EventObjectIdentifier = eventObjectIdentifier
	return b
}

func (b *_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder) WithEventObjectIdentifierBuilder(builderSupplier func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder {
	builder := builderSupplier(b.EventObjectIdentifier.CreateBACnetContextTagObjectIdentifierBuilder())
	var err error
	b.EventObjectIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagObjectIdentifierBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder) WithEventStateAcknowledged(eventStateAcknowledged BACnetEventStateTagged) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder {
	b.EventStateAcknowledged = eventStateAcknowledged
	return b
}

func (b *_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder) WithEventStateAcknowledgedBuilder(builderSupplier func(BACnetEventStateTaggedBuilder) BACnetEventStateTaggedBuilder) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder {
	builder := builderSupplier(b.EventStateAcknowledged.CreateBACnetEventStateTaggedBuilder())
	var err error
	b.EventStateAcknowledged, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetEventStateTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder) WithTimestamp(timestamp BACnetTimeStampEnclosed) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder {
	b.Timestamp = timestamp
	return b
}

func (b *_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder) WithTimestampBuilder(builderSupplier func(BACnetTimeStampEnclosedBuilder) BACnetTimeStampEnclosedBuilder) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder {
	builder := builderSupplier(b.Timestamp.CreateBACnetTimeStampEnclosedBuilder())
	var err error
	b.Timestamp, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTimeStampEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder) WithAcknowledgmentSource(acknowledgmentSource BACnetContextTagCharacterString) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder {
	b.AcknowledgmentSource = acknowledgmentSource
	return b
}

func (b *_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder) WithAcknowledgmentSourceBuilder(builderSupplier func(BACnetContextTagCharacterStringBuilder) BACnetContextTagCharacterStringBuilder) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder {
	builder := builderSupplier(b.AcknowledgmentSource.CreateBACnetContextTagCharacterStringBuilder())
	var err error
	b.AcknowledgmentSource, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder) WithTimeOfAcknowledgment(timeOfAcknowledgment BACnetTimeStampEnclosed) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder {
	b.TimeOfAcknowledgment = timeOfAcknowledgment
	return b
}

func (b *_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder) WithTimeOfAcknowledgmentBuilder(builderSupplier func(BACnetTimeStampEnclosedBuilder) BACnetTimeStampEnclosedBuilder) BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder {
	builder := builderSupplier(b.TimeOfAcknowledgment.CreateBACnetTimeStampEnclosedBuilder())
	var err error
	b.TimeOfAcknowledgment, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTimeStampEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder) Build() (BACnetConfirmedServiceRequestAcknowledgeAlarm, error) {
	if b.AcknowledgingProcessIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'acknowledgingProcessIdentifier' not set"))
	}
	if b.EventObjectIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'eventObjectIdentifier' not set"))
	}
	if b.EventStateAcknowledged == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'eventStateAcknowledged' not set"))
	}
	if b.Timestamp == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timestamp' not set"))
	}
	if b.AcknowledgmentSource == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'acknowledgmentSource' not set"))
	}
	if b.TimeOfAcknowledgment == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeOfAcknowledgment' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestAcknowledgeAlarm.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder) MustBuild() BACnetConfirmedServiceRequestAcknowledgeAlarm {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder) Done() BACnetConfirmedServiceRequestBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConfirmedServiceRequestBuilder().(*_BACnetConfirmedServiceRequestBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder) buildForBACnetConfirmedServiceRequest() (BACnetConfirmedServiceRequest, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestAcknowledgeAlarmBuilder().(*_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestAcknowledgeAlarmBuilder creates a BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder
func (b *_BACnetConfirmedServiceRequestAcknowledgeAlarm) CreateBACnetConfirmedServiceRequestAcknowledgeAlarmBuilder() BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestAcknowledgeAlarmBuilder()
	}
	return &_BACnetConfirmedServiceRequestAcknowledgeAlarmBuilder{_BACnetConfirmedServiceRequestAcknowledgeAlarm: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_ACKNOWLEDGE_ALARM
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetAcknowledgingProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.AcknowledgingProcessIdentifier
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetEventObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.EventObjectIdentifier
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetEventStateAcknowledged() BACnetEventStateTagged {
	return m.EventStateAcknowledged
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetTimestamp() BACnetTimeStampEnclosed {
	return m.Timestamp
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetAcknowledgmentSource() BACnetContextTagCharacterString {
	return m.AcknowledgmentSource
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetTimeOfAcknowledgment() BACnetTimeStampEnclosed {
	return m.TimeOfAcknowledgment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestAcknowledgeAlarm(structType any) BACnetConfirmedServiceRequestAcknowledgeAlarm {
	if casted, ok := structType.(BACnetConfirmedServiceRequestAcknowledgeAlarm); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestAcknowledgeAlarm); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetTypeName() string {
	return "BACnetConfirmedServiceRequestAcknowledgeAlarm"
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (acknowledgingProcessIdentifier)
	lengthInBits += m.AcknowledgingProcessIdentifier.GetLengthInBits(ctx)

	// Simple field (eventObjectIdentifier)
	lengthInBits += m.EventObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (eventStateAcknowledged)
	lengthInBits += m.EventStateAcknowledged.GetLengthInBits(ctx)

	// Simple field (timestamp)
	lengthInBits += m.Timestamp.GetLengthInBits(ctx)

	// Simple field (acknowledgmentSource)
	lengthInBits += m.AcknowledgmentSource.GetLengthInBits(ctx)

	// Simple field (timeOfAcknowledgment)
	lengthInBits += m.TimeOfAcknowledgment.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestAcknowledgeAlarm BACnetConfirmedServiceRequestAcknowledgeAlarm, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestAcknowledgeAlarm"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestAcknowledgeAlarm")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	acknowledgingProcessIdentifier, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "acknowledgingProcessIdentifier", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'acknowledgingProcessIdentifier' field"))
	}
	m.AcknowledgingProcessIdentifier = acknowledgingProcessIdentifier

	eventObjectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "eventObjectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventObjectIdentifier' field"))
	}
	m.EventObjectIdentifier = eventObjectIdentifier

	eventStateAcknowledged, err := ReadSimpleField[BACnetEventStateTagged](ctx, "eventStateAcknowledged", ReadComplex[BACnetEventStateTagged](BACnetEventStateTaggedParseWithBufferProducer((uint8)(uint8(2)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventStateAcknowledged' field"))
	}
	m.EventStateAcknowledged = eventStateAcknowledged

	timestamp, err := ReadSimpleField[BACnetTimeStampEnclosed](ctx, "timestamp", ReadComplex[BACnetTimeStampEnclosed](BACnetTimeStampEnclosedParseWithBufferProducer((uint8)(uint8(3))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestamp' field"))
	}
	m.Timestamp = timestamp

	acknowledgmentSource, err := ReadSimpleField[BACnetContextTagCharacterString](ctx, "acknowledgmentSource", ReadComplex[BACnetContextTagCharacterString](BACnetContextTagParseWithBufferProducer[BACnetContextTagCharacterString]((uint8)(uint8(4)), (BACnetDataType)(BACnetDataType_CHARACTER_STRING)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'acknowledgmentSource' field"))
	}
	m.AcknowledgmentSource = acknowledgmentSource

	timeOfAcknowledgment, err := ReadSimpleField[BACnetTimeStampEnclosed](ctx, "timeOfAcknowledgment", ReadComplex[BACnetTimeStampEnclosed](BACnetTimeStampEnclosedParseWithBufferProducer((uint8)(uint8(5))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeOfAcknowledgment' field"))
	}
	m.TimeOfAcknowledgment = timeOfAcknowledgment

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestAcknowledgeAlarm"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestAcknowledgeAlarm")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestAcknowledgeAlarm"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestAcknowledgeAlarm")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "acknowledgingProcessIdentifier", m.GetAcknowledgingProcessIdentifier(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'acknowledgingProcessIdentifier' field")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "eventObjectIdentifier", m.GetEventObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventObjectIdentifier' field")
		}

		if err := WriteSimpleField[BACnetEventStateTagged](ctx, "eventStateAcknowledged", m.GetEventStateAcknowledged(), WriteComplex[BACnetEventStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventStateAcknowledged' field")
		}

		if err := WriteSimpleField[BACnetTimeStampEnclosed](ctx, "timestamp", m.GetTimestamp(), WriteComplex[BACnetTimeStampEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timestamp' field")
		}

		if err := WriteSimpleField[BACnetContextTagCharacterString](ctx, "acknowledgmentSource", m.GetAcknowledgmentSource(), WriteComplex[BACnetContextTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'acknowledgmentSource' field")
		}

		if err := WriteSimpleField[BACnetTimeStampEnclosed](ctx, "timeOfAcknowledgment", m.GetTimeOfAcknowledgment(), WriteComplex[BACnetTimeStampEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeOfAcknowledgment' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestAcknowledgeAlarm"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestAcknowledgeAlarm")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) IsBACnetConfirmedServiceRequestAcknowledgeAlarm() {
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) deepCopy() *_BACnetConfirmedServiceRequestAcknowledgeAlarm {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestAcknowledgeAlarmCopy := &_BACnetConfirmedServiceRequestAcknowledgeAlarm{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.AcknowledgingProcessIdentifier),
		utils.DeepCopy[BACnetContextTagObjectIdentifier](m.EventObjectIdentifier),
		utils.DeepCopy[BACnetEventStateTagged](m.EventStateAcknowledged),
		utils.DeepCopy[BACnetTimeStampEnclosed](m.Timestamp),
		utils.DeepCopy[BACnetContextTagCharacterString](m.AcknowledgmentSource),
		utils.DeepCopy[BACnetTimeStampEnclosed](m.TimeOfAcknowledgment),
	}
	_BACnetConfirmedServiceRequestAcknowledgeAlarmCopy.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestAcknowledgeAlarmCopy
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) String() string {
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
