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

// BACnetConfirmedServiceRequestSubscribeCOVProperty is the corresponding interface of BACnetConfirmedServiceRequestSubscribeCOVProperty
type BACnetConfirmedServiceRequestSubscribeCOVProperty interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetSubscriberProcessIdentifier returns SubscriberProcessIdentifier (property field)
	GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetMonitoredObjectIdentifier returns MonitoredObjectIdentifier (property field)
	GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetIssueConfirmedNotifications returns IssueConfirmedNotifications (property field)
	GetIssueConfirmedNotifications() BACnetContextTagBoolean
	// GetLifetime returns Lifetime (property field)
	GetLifetime() BACnetContextTagUnsignedInteger
	// GetMonitoredPropertyIdentifier returns MonitoredPropertyIdentifier (property field)
	GetMonitoredPropertyIdentifier() BACnetPropertyReferenceEnclosed
	// GetCovIncrement returns CovIncrement (property field)
	GetCovIncrement() BACnetContextTagReal
	// IsBACnetConfirmedServiceRequestSubscribeCOVProperty is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestSubscribeCOVProperty()
}

// _BACnetConfirmedServiceRequestSubscribeCOVProperty is the data-structure of this message
type _BACnetConfirmedServiceRequestSubscribeCOVProperty struct {
	BACnetConfirmedServiceRequestContract
	SubscriberProcessIdentifier BACnetContextTagUnsignedInteger
	MonitoredObjectIdentifier   BACnetContextTagObjectIdentifier
	IssueConfirmedNotifications BACnetContextTagBoolean
	Lifetime                    BACnetContextTagUnsignedInteger
	MonitoredPropertyIdentifier BACnetPropertyReferenceEnclosed
	CovIncrement                BACnetContextTagReal
}

var _ BACnetConfirmedServiceRequestSubscribeCOVProperty = (*_BACnetConfirmedServiceRequestSubscribeCOVProperty)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestSubscribeCOVProperty)(nil)

// NewBACnetConfirmedServiceRequestSubscribeCOVProperty factory function for _BACnetConfirmedServiceRequestSubscribeCOVProperty
func NewBACnetConfirmedServiceRequestSubscribeCOVProperty(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, monitoredObjectIdentifier BACnetContextTagObjectIdentifier, issueConfirmedNotifications BACnetContextTagBoolean, lifetime BACnetContextTagUnsignedInteger, monitoredPropertyIdentifier BACnetPropertyReferenceEnclosed, covIncrement BACnetContextTagReal, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestSubscribeCOVProperty {
	if subscriberProcessIdentifier == nil {
		panic("subscriberProcessIdentifier of type BACnetContextTagUnsignedInteger for BACnetConfirmedServiceRequestSubscribeCOVProperty must not be nil")
	}
	if monitoredObjectIdentifier == nil {
		panic("monitoredObjectIdentifier of type BACnetContextTagObjectIdentifier for BACnetConfirmedServiceRequestSubscribeCOVProperty must not be nil")
	}
	if monitoredPropertyIdentifier == nil {
		panic("monitoredPropertyIdentifier of type BACnetPropertyReferenceEnclosed for BACnetConfirmedServiceRequestSubscribeCOVProperty must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestSubscribeCOVProperty{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		SubscriberProcessIdentifier:           subscriberProcessIdentifier,
		MonitoredObjectIdentifier:             monitoredObjectIdentifier,
		IssueConfirmedNotifications:           issueConfirmedNotifications,
		Lifetime:                              lifetime,
		MonitoredPropertyIdentifier:           monitoredPropertyIdentifier,
		CovIncrement:                          covIncrement,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.SubscriberProcessIdentifier
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.MonitoredObjectIdentifier
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetIssueConfirmedNotifications() BACnetContextTagBoolean {
	return m.IssueConfirmedNotifications
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetLifetime() BACnetContextTagUnsignedInteger {
	return m.Lifetime
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetMonitoredPropertyIdentifier() BACnetPropertyReferenceEnclosed {
	return m.MonitoredPropertyIdentifier
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetCovIncrement() BACnetContextTagReal {
	return m.CovIncrement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestSubscribeCOVProperty(structType any) BACnetConfirmedServiceRequestSubscribeCOVProperty {
	if casted, ok := structType.(BACnetConfirmedServiceRequestSubscribeCOVProperty); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestSubscribeCOVProperty); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetTypeName() string {
	return "BACnetConfirmedServiceRequestSubscribeCOVProperty"
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += m.SubscriberProcessIdentifier.GetLengthInBits(ctx)

	// Simple field (monitoredObjectIdentifier)
	lengthInBits += m.MonitoredObjectIdentifier.GetLengthInBits(ctx)

	// Optional Field (issueConfirmedNotifications)
	if m.IssueConfirmedNotifications != nil {
		lengthInBits += m.IssueConfirmedNotifications.GetLengthInBits(ctx)
	}

	// Optional Field (lifetime)
	if m.Lifetime != nil {
		lengthInBits += m.Lifetime.GetLengthInBits(ctx)
	}

	// Simple field (monitoredPropertyIdentifier)
	lengthInBits += m.MonitoredPropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (covIncrement)
	if m.CovIncrement != nil {
		lengthInBits += m.CovIncrement.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestSubscribeCOVProperty BACnetConfirmedServiceRequestSubscribeCOVProperty, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestSubscribeCOVProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestSubscribeCOVProperty")
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

	var issueConfirmedNotifications BACnetContextTagBoolean
	_issueConfirmedNotifications, err := ReadOptionalField[BACnetContextTagBoolean](ctx, "issueConfirmedNotifications", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'issueConfirmedNotifications' field"))
	}
	if _issueConfirmedNotifications != nil {
		issueConfirmedNotifications = *_issueConfirmedNotifications
		m.IssueConfirmedNotifications = issueConfirmedNotifications
	}

	var lifetime BACnetContextTagUnsignedInteger
	_lifetime, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "lifetime", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lifetime' field"))
	}
	if _lifetime != nil {
		lifetime = *_lifetime
		m.Lifetime = lifetime
	}

	monitoredPropertyIdentifier, err := ReadSimpleField[BACnetPropertyReferenceEnclosed](ctx, "monitoredPropertyIdentifier", ReadComplex[BACnetPropertyReferenceEnclosed](BACnetPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(4))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredPropertyIdentifier' field"))
	}
	m.MonitoredPropertyIdentifier = monitoredPropertyIdentifier

	var covIncrement BACnetContextTagReal
	_covIncrement, err := ReadOptionalField[BACnetContextTagReal](ctx, "covIncrement", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(5)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'covIncrement' field"))
	}
	if _covIncrement != nil {
		covIncrement = *_covIncrement
		m.CovIncrement = covIncrement
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestSubscribeCOVProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestSubscribeCOVProperty")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestSubscribeCOVProperty"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestSubscribeCOVProperty")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "subscriberProcessIdentifier", m.GetSubscriberProcessIdentifier(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriberProcessIdentifier' field")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "monitoredObjectIdentifier", m.GetMonitoredObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'monitoredObjectIdentifier' field")
		}

		if err := WriteOptionalField[BACnetContextTagBoolean](ctx, "issueConfirmedNotifications", GetRef(m.GetIssueConfirmedNotifications()), WriteComplex[BACnetContextTagBoolean](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'issueConfirmedNotifications' field")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "lifetime", GetRef(m.GetLifetime()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'lifetime' field")
		}

		if err := WriteSimpleField[BACnetPropertyReferenceEnclosed](ctx, "monitoredPropertyIdentifier", m.GetMonitoredPropertyIdentifier(), WriteComplex[BACnetPropertyReferenceEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'monitoredPropertyIdentifier' field")
		}

		if err := WriteOptionalField[BACnetContextTagReal](ctx, "covIncrement", GetRef(m.GetCovIncrement()), WriteComplex[BACnetContextTagReal](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'covIncrement' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestSubscribeCOVProperty"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestSubscribeCOVProperty")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) IsBACnetConfirmedServiceRequestSubscribeCOVProperty() {
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
