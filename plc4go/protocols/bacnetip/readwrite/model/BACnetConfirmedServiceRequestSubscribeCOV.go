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

// BACnetConfirmedServiceRequestSubscribeCOV is the corresponding interface of BACnetConfirmedServiceRequestSubscribeCOV
type BACnetConfirmedServiceRequestSubscribeCOV interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetSubscriberProcessIdentifier returns SubscriberProcessIdentifier (property field)
	GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetMonitoredObjectIdentifier returns MonitoredObjectIdentifier (property field)
	GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetIssueConfirmed returns IssueConfirmed (property field)
	GetIssueConfirmed() BACnetContextTagBoolean
	// GetLifetimeInSeconds returns LifetimeInSeconds (property field)
	GetLifetimeInSeconds() BACnetContextTagUnsignedInteger
	// IsBACnetConfirmedServiceRequestSubscribeCOV is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestSubscribeCOV()
	// CreateBuilder creates a BACnetConfirmedServiceRequestSubscribeCOVBuilder
	CreateBACnetConfirmedServiceRequestSubscribeCOVBuilder() BACnetConfirmedServiceRequestSubscribeCOVBuilder
}

// _BACnetConfirmedServiceRequestSubscribeCOV is the data-structure of this message
type _BACnetConfirmedServiceRequestSubscribeCOV struct {
	BACnetConfirmedServiceRequestContract
	SubscriberProcessIdentifier BACnetContextTagUnsignedInteger
	MonitoredObjectIdentifier   BACnetContextTagObjectIdentifier
	IssueConfirmed              BACnetContextTagBoolean
	LifetimeInSeconds           BACnetContextTagUnsignedInteger
}

var _ BACnetConfirmedServiceRequestSubscribeCOV = (*_BACnetConfirmedServiceRequestSubscribeCOV)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestSubscribeCOV)(nil)

// NewBACnetConfirmedServiceRequestSubscribeCOV factory function for _BACnetConfirmedServiceRequestSubscribeCOV
func NewBACnetConfirmedServiceRequestSubscribeCOV(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, monitoredObjectIdentifier BACnetContextTagObjectIdentifier, issueConfirmed BACnetContextTagBoolean, lifetimeInSeconds BACnetContextTagUnsignedInteger, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestSubscribeCOV {
	if subscriberProcessIdentifier == nil {
		panic("subscriberProcessIdentifier of type BACnetContextTagUnsignedInteger for BACnetConfirmedServiceRequestSubscribeCOV must not be nil")
	}
	if monitoredObjectIdentifier == nil {
		panic("monitoredObjectIdentifier of type BACnetContextTagObjectIdentifier for BACnetConfirmedServiceRequestSubscribeCOV must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestSubscribeCOV{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		SubscriberProcessIdentifier:           subscriberProcessIdentifier,
		MonitoredObjectIdentifier:             monitoredObjectIdentifier,
		IssueConfirmed:                        issueConfirmed,
		LifetimeInSeconds:                     lifetimeInSeconds,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestSubscribeCOVBuilder is a builder for BACnetConfirmedServiceRequestSubscribeCOV
type BACnetConfirmedServiceRequestSubscribeCOVBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, monitoredObjectIdentifier BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestSubscribeCOVBuilder
	// WithSubscriberProcessIdentifier adds SubscriberProcessIdentifier (property field)
	WithSubscriberProcessIdentifier(BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestSubscribeCOVBuilder
	// WithSubscriberProcessIdentifierBuilder adds SubscriberProcessIdentifier (property field) which is build by the builder
	WithSubscriberProcessIdentifierBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestSubscribeCOVBuilder
	// WithMonitoredObjectIdentifier adds MonitoredObjectIdentifier (property field)
	WithMonitoredObjectIdentifier(BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestSubscribeCOVBuilder
	// WithMonitoredObjectIdentifierBuilder adds MonitoredObjectIdentifier (property field) which is build by the builder
	WithMonitoredObjectIdentifierBuilder(func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestSubscribeCOVBuilder
	// WithIssueConfirmed adds IssueConfirmed (property field)
	WithOptionalIssueConfirmed(BACnetContextTagBoolean) BACnetConfirmedServiceRequestSubscribeCOVBuilder
	// WithOptionalIssueConfirmedBuilder adds IssueConfirmed (property field) which is build by the builder
	WithOptionalIssueConfirmedBuilder(func(BACnetContextTagBooleanBuilder) BACnetContextTagBooleanBuilder) BACnetConfirmedServiceRequestSubscribeCOVBuilder
	// WithLifetimeInSeconds adds LifetimeInSeconds (property field)
	WithOptionalLifetimeInSeconds(BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestSubscribeCOVBuilder
	// WithOptionalLifetimeInSecondsBuilder adds LifetimeInSeconds (property field) which is build by the builder
	WithOptionalLifetimeInSecondsBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestSubscribeCOVBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConfirmedServiceRequestBuilder
	// Build builds the BACnetConfirmedServiceRequestSubscribeCOV or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestSubscribeCOV, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestSubscribeCOV
}

// NewBACnetConfirmedServiceRequestSubscribeCOVBuilder() creates a BACnetConfirmedServiceRequestSubscribeCOVBuilder
func NewBACnetConfirmedServiceRequestSubscribeCOVBuilder() BACnetConfirmedServiceRequestSubscribeCOVBuilder {
	return &_BACnetConfirmedServiceRequestSubscribeCOVBuilder{_BACnetConfirmedServiceRequestSubscribeCOV: new(_BACnetConfirmedServiceRequestSubscribeCOV)}
}

type _BACnetConfirmedServiceRequestSubscribeCOVBuilder struct {
	*_BACnetConfirmedServiceRequestSubscribeCOV

	parentBuilder *_BACnetConfirmedServiceRequestBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestSubscribeCOVBuilder) = (*_BACnetConfirmedServiceRequestSubscribeCOVBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestSubscribeCOVBuilder) setParent(contract BACnetConfirmedServiceRequestContract) {
	b.BACnetConfirmedServiceRequestContract = contract
	contract.(*_BACnetConfirmedServiceRequest)._SubType = b._BACnetConfirmedServiceRequestSubscribeCOV
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVBuilder) WithMandatoryFields(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, monitoredObjectIdentifier BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestSubscribeCOVBuilder {
	return b.WithSubscriberProcessIdentifier(subscriberProcessIdentifier).WithMonitoredObjectIdentifier(monitoredObjectIdentifier)
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVBuilder) WithSubscriberProcessIdentifier(subscriberProcessIdentifier BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestSubscribeCOVBuilder {
	b.SubscriberProcessIdentifier = subscriberProcessIdentifier
	return b
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVBuilder) WithSubscriberProcessIdentifierBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestSubscribeCOVBuilder {
	builder := builderSupplier(b.SubscriberProcessIdentifier.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.SubscriberProcessIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVBuilder) WithMonitoredObjectIdentifier(monitoredObjectIdentifier BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestSubscribeCOVBuilder {
	b.MonitoredObjectIdentifier = monitoredObjectIdentifier
	return b
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVBuilder) WithMonitoredObjectIdentifierBuilder(builderSupplier func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestSubscribeCOVBuilder {
	builder := builderSupplier(b.MonitoredObjectIdentifier.CreateBACnetContextTagObjectIdentifierBuilder())
	var err error
	b.MonitoredObjectIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagObjectIdentifierBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVBuilder) WithOptionalIssueConfirmed(issueConfirmed BACnetContextTagBoolean) BACnetConfirmedServiceRequestSubscribeCOVBuilder {
	b.IssueConfirmed = issueConfirmed
	return b
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVBuilder) WithOptionalIssueConfirmedBuilder(builderSupplier func(BACnetContextTagBooleanBuilder) BACnetContextTagBooleanBuilder) BACnetConfirmedServiceRequestSubscribeCOVBuilder {
	builder := builderSupplier(b.IssueConfirmed.CreateBACnetContextTagBooleanBuilder())
	var err error
	b.IssueConfirmed, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVBuilder) WithOptionalLifetimeInSeconds(lifetimeInSeconds BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestSubscribeCOVBuilder {
	b.LifetimeInSeconds = lifetimeInSeconds
	return b
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVBuilder) WithOptionalLifetimeInSecondsBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestSubscribeCOVBuilder {
	builder := builderSupplier(b.LifetimeInSeconds.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.LifetimeInSeconds, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVBuilder) Build() (BACnetConfirmedServiceRequestSubscribeCOV, error) {
	if b.SubscriberProcessIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'subscriberProcessIdentifier' not set"))
	}
	if b.MonitoredObjectIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'monitoredObjectIdentifier' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestSubscribeCOV.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVBuilder) MustBuild() BACnetConfirmedServiceRequestSubscribeCOV {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVBuilder) Done() BACnetConfirmedServiceRequestBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConfirmedServiceRequestBuilder().(*_BACnetConfirmedServiceRequestBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVBuilder) buildForBACnetConfirmedServiceRequest() (BACnetConfirmedServiceRequest, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestSubscribeCOVBuilder().(*_BACnetConfirmedServiceRequestSubscribeCOVBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestSubscribeCOVBuilder creates a BACnetConfirmedServiceRequestSubscribeCOVBuilder
func (b *_BACnetConfirmedServiceRequestSubscribeCOV) CreateBACnetConfirmedServiceRequestSubscribeCOVBuilder() BACnetConfirmedServiceRequestSubscribeCOVBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestSubscribeCOVBuilder()
	}
	return &_BACnetConfirmedServiceRequestSubscribeCOVBuilder{_BACnetConfirmedServiceRequestSubscribeCOV: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_SUBSCRIBE_COV
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.SubscriberProcessIdentifier
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.MonitoredObjectIdentifier
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) GetIssueConfirmed() BACnetContextTagBoolean {
	return m.IssueConfirmed
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) GetLifetimeInSeconds() BACnetContextTagUnsignedInteger {
	return m.LifetimeInSeconds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestSubscribeCOV(structType any) BACnetConfirmedServiceRequestSubscribeCOV {
	if casted, ok := structType.(BACnetConfirmedServiceRequestSubscribeCOV); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestSubscribeCOV); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) GetTypeName() string {
	return "BACnetConfirmedServiceRequestSubscribeCOV"
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += m.SubscriberProcessIdentifier.GetLengthInBits(ctx)

	// Simple field (monitoredObjectIdentifier)
	lengthInBits += m.MonitoredObjectIdentifier.GetLengthInBits(ctx)

	// Optional Field (issueConfirmed)
	if m.IssueConfirmed != nil {
		lengthInBits += m.IssueConfirmed.GetLengthInBits(ctx)
	}

	// Optional Field (lifetimeInSeconds)
	if m.LifetimeInSeconds != nil {
		lengthInBits += m.LifetimeInSeconds.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestSubscribeCOV BACnetConfirmedServiceRequestSubscribeCOV, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestSubscribeCOV"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestSubscribeCOV")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	subscriberProcessIdentifier, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "subscriberProcessIdentifier", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriberProcessIdentifier' field"))
	}
	m.SubscriberProcessIdentifier = subscriberProcessIdentifier

	monitoredObjectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "monitoredObjectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredObjectIdentifier' field"))
	}
	m.MonitoredObjectIdentifier = monitoredObjectIdentifier

	var issueConfirmed BACnetContextTagBoolean
	_issueConfirmed, err := ReadOptionalField[BACnetContextTagBoolean](ctx, "issueConfirmed", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'issueConfirmed' field"))
	}
	if _issueConfirmed != nil {
		issueConfirmed = *_issueConfirmed
		m.IssueConfirmed = issueConfirmed
	}

	var lifetimeInSeconds BACnetContextTagUnsignedInteger
	_lifetimeInSeconds, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "lifetimeInSeconds", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lifetimeInSeconds' field"))
	}
	if _lifetimeInSeconds != nil {
		lifetimeInSeconds = *_lifetimeInSeconds
		m.LifetimeInSeconds = lifetimeInSeconds
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestSubscribeCOV"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestSubscribeCOV")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestSubscribeCOV"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestSubscribeCOV")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "subscriberProcessIdentifier", m.GetSubscriberProcessIdentifier(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriberProcessIdentifier' field")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "monitoredObjectIdentifier", m.GetMonitoredObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'monitoredObjectIdentifier' field")
		}

		if err := WriteOptionalField[BACnetContextTagBoolean](ctx, "issueConfirmed", GetRef(m.GetIssueConfirmed()), WriteComplex[BACnetContextTagBoolean](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'issueConfirmed' field")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "lifetimeInSeconds", GetRef(m.GetLifetimeInSeconds()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'lifetimeInSeconds' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestSubscribeCOV"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestSubscribeCOV")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) IsBACnetConfirmedServiceRequestSubscribeCOV() {}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) deepCopy() *_BACnetConfirmedServiceRequestSubscribeCOV {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestSubscribeCOVCopy := &_BACnetConfirmedServiceRequestSubscribeCOV{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.SubscriberProcessIdentifier),
		utils.DeepCopy[BACnetContextTagObjectIdentifier](m.MonitoredObjectIdentifier),
		utils.DeepCopy[BACnetContextTagBoolean](m.IssueConfirmed),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.LifetimeInSeconds),
	}
	_BACnetConfirmedServiceRequestSubscribeCOVCopy.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestSubscribeCOVCopy
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) String() string {
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