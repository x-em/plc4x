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

// BACnetUnconfirmedServiceRequestUnconfirmedEventNotification is the corresponding interface of BACnetUnconfirmedServiceRequestUnconfirmedEventNotification
type BACnetUnconfirmedServiceRequestUnconfirmedEventNotification interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetUnconfirmedServiceRequest
	// GetProcessIdentifier returns ProcessIdentifier (property field)
	GetProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetInitiatingDeviceIdentifier returns InitiatingDeviceIdentifier (property field)
	GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier
	// GetEventObjectIdentifier returns EventObjectIdentifier (property field)
	GetEventObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetTimeStampEnclosed
	// GetNotificationClass returns NotificationClass (property field)
	GetNotificationClass() BACnetContextTagUnsignedInteger
	// GetPriority returns Priority (property field)
	GetPriority() BACnetContextTagUnsignedInteger
	// GetEventType returns EventType (property field)
	GetEventType() BACnetEventTypeTagged
	// GetMessageText returns MessageText (property field)
	GetMessageText() BACnetContextTagCharacterString
	// GetNotifyType returns NotifyType (property field)
	GetNotifyType() BACnetNotifyTypeTagged
	// GetAckRequired returns AckRequired (property field)
	GetAckRequired() BACnetContextTagBoolean
	// GetFromState returns FromState (property field)
	GetFromState() BACnetEventStateTagged
	// GetToState returns ToState (property field)
	GetToState() BACnetEventStateTagged
	// GetEventValues returns EventValues (property field)
	GetEventValues() BACnetNotificationParameters
}

// BACnetUnconfirmedServiceRequestUnconfirmedEventNotificationExactly can be used when we want exactly this type and not a type which fulfills BACnetUnconfirmedServiceRequestUnconfirmedEventNotification.
// This is useful for switch cases.
type BACnetUnconfirmedServiceRequestUnconfirmedEventNotificationExactly interface {
	BACnetUnconfirmedServiceRequestUnconfirmedEventNotification
	isBACnetUnconfirmedServiceRequestUnconfirmedEventNotification() bool
}

// _BACnetUnconfirmedServiceRequestUnconfirmedEventNotification is the data-structure of this message
type _BACnetUnconfirmedServiceRequestUnconfirmedEventNotification struct {
	*_BACnetUnconfirmedServiceRequest
	ProcessIdentifier          BACnetContextTagUnsignedInteger
	InitiatingDeviceIdentifier BACnetContextTagObjectIdentifier
	EventObjectIdentifier      BACnetContextTagObjectIdentifier
	Timestamp                  BACnetTimeStampEnclosed
	NotificationClass          BACnetContextTagUnsignedInteger
	Priority                   BACnetContextTagUnsignedInteger
	EventType                  BACnetEventTypeTagged
	MessageText                BACnetContextTagCharacterString
	NotifyType                 BACnetNotifyTypeTagged
	AckRequired                BACnetContextTagBoolean
	FromState                  BACnetEventStateTagged
	ToState                    BACnetEventStateTagged
	EventValues                BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_UNCONFIRMED_EVENT_NOTIFICATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) InitializeParent(parent BACnetUnconfirmedServiceRequest) {
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetParent() BACnetUnconfirmedServiceRequest {
	return m._BACnetUnconfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.ProcessIdentifier
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier {
	return m.InitiatingDeviceIdentifier
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetEventObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.EventObjectIdentifier
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetTimestamp() BACnetTimeStampEnclosed {
	return m.Timestamp
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetNotificationClass() BACnetContextTagUnsignedInteger {
	return m.NotificationClass
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetPriority() BACnetContextTagUnsignedInteger {
	return m.Priority
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetEventType() BACnetEventTypeTagged {
	return m.EventType
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetMessageText() BACnetContextTagCharacterString {
	return m.MessageText
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetNotifyType() BACnetNotifyTypeTagged {
	return m.NotifyType
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetAckRequired() BACnetContextTagBoolean {
	return m.AckRequired
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetFromState() BACnetEventStateTagged {
	return m.FromState
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetToState() BACnetEventStateTagged {
	return m.ToState
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetEventValues() BACnetNotificationParameters {
	return m.EventValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestUnconfirmedEventNotification factory function for _BACnetUnconfirmedServiceRequestUnconfirmedEventNotification
func NewBACnetUnconfirmedServiceRequestUnconfirmedEventNotification(processIdentifier BACnetContextTagUnsignedInteger, initiatingDeviceIdentifier BACnetContextTagObjectIdentifier, eventObjectIdentifier BACnetContextTagObjectIdentifier, timestamp BACnetTimeStampEnclosed, notificationClass BACnetContextTagUnsignedInteger, priority BACnetContextTagUnsignedInteger, eventType BACnetEventTypeTagged, messageText BACnetContextTagCharacterString, notifyType BACnetNotifyTypeTagged, ackRequired BACnetContextTagBoolean, fromState BACnetEventStateTagged, toState BACnetEventStateTagged, eventValues BACnetNotificationParameters, serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification {
	_result := &_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification{
		ProcessIdentifier:                processIdentifier,
		InitiatingDeviceIdentifier:       initiatingDeviceIdentifier,
		EventObjectIdentifier:            eventObjectIdentifier,
		Timestamp:                        timestamp,
		NotificationClass:                notificationClass,
		Priority:                         priority,
		EventType:                        eventType,
		MessageText:                      messageText,
		NotifyType:                       notifyType,
		AckRequired:                      ackRequired,
		FromState:                        fromState,
		ToState:                          toState,
		EventValues:                      eventValues,
		_BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestUnconfirmedEventNotification(structType any) BACnetUnconfirmedServiceRequestUnconfirmedEventNotification {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestUnconfirmedEventNotification); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestUnconfirmedEventNotification); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUnconfirmedEventNotification"
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (processIdentifier)
	lengthInBits += m.ProcessIdentifier.GetLengthInBits(ctx)

	// Simple field (initiatingDeviceIdentifier)
	lengthInBits += m.InitiatingDeviceIdentifier.GetLengthInBits(ctx)

	// Simple field (eventObjectIdentifier)
	lengthInBits += m.EventObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (timestamp)
	lengthInBits += m.Timestamp.GetLengthInBits(ctx)

	// Simple field (notificationClass)
	lengthInBits += m.NotificationClass.GetLengthInBits(ctx)

	// Simple field (priority)
	lengthInBits += m.Priority.GetLengthInBits(ctx)

	// Simple field (eventType)
	lengthInBits += m.EventType.GetLengthInBits(ctx)

	// Optional Field (messageText)
	if m.MessageText != nil {
		lengthInBits += m.MessageText.GetLengthInBits(ctx)
	}

	// Simple field (notifyType)
	lengthInBits += m.NotifyType.GetLengthInBits(ctx)

	// Optional Field (ackRequired)
	if m.AckRequired != nil {
		lengthInBits += m.AckRequired.GetLengthInBits(ctx)
	}

	// Optional Field (fromState)
	if m.FromState != nil {
		lengthInBits += m.FromState.GetLengthInBits(ctx)
	}

	// Simple field (toState)
	lengthInBits += m.ToState.GetLengthInBits(ctx)

	// Optional Field (eventValues)
	if m.EventValues != nil {
		lengthInBits += m.EventValues.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetUnconfirmedServiceRequestUnconfirmedEventNotificationParse(ctx context.Context, theBytes []byte, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequestUnconfirmedEventNotification, error) {
	return BACnetUnconfirmedServiceRequestUnconfirmedEventNotificationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetUnconfirmedServiceRequestUnconfirmedEventNotificationParseWithBufferProducer(serviceRequestLength uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetUnconfirmedServiceRequestUnconfirmedEventNotification, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetUnconfirmedServiceRequestUnconfirmedEventNotification, error) {
		return BACnetUnconfirmedServiceRequestUnconfirmedEventNotificationParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	}
}

func BACnetUnconfirmedServiceRequestUnconfirmedEventNotificationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequestUnconfirmedEventNotification, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestUnconfirmedEventNotification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestUnconfirmedEventNotification")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	processIdentifier, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "processIdentifier", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'processIdentifier' field"))
	}

	initiatingDeviceIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "initiatingDeviceIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'initiatingDeviceIdentifier' field"))
	}

	eventObjectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "eventObjectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventObjectIdentifier' field"))
	}

	timestamp, err := ReadSimpleField[BACnetTimeStampEnclosed](ctx, "timestamp", ReadComplex[BACnetTimeStampEnclosed](BACnetTimeStampEnclosedParseWithBufferProducer((uint8)(uint8(3))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestamp' field"))
	}

	notificationClass, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "notificationClass", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(4)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notificationClass' field"))
	}

	priority, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "priority", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(5)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}

	eventType, err := ReadSimpleField[BACnetEventTypeTagged](ctx, "eventType", ReadComplex[BACnetEventTypeTagged](BACnetEventTypeTaggedParseWithBufferProducer((uint8)(uint8(6)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventType' field"))
	}

	_messageText, err := ReadOptionalField[BACnetContextTagCharacterString](ctx, "messageText", ReadComplex[BACnetContextTagCharacterString](BACnetContextTagParseWithBufferProducer[BACnetContextTagCharacterString]((uint8)(uint8(7)), (BACnetDataType)(BACnetDataType_CHARACTER_STRING)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageText' field"))
	}
	var messageText BACnetContextTagCharacterString
	if _messageText != nil {
		messageText = *_messageText
	}

	notifyType, err := ReadSimpleField[BACnetNotifyTypeTagged](ctx, "notifyType", ReadComplex[BACnetNotifyTypeTagged](BACnetNotifyTypeTaggedParseWithBufferProducer((uint8)(uint8(8)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notifyType' field"))
	}

	_ackRequired, err := ReadOptionalField[BACnetContextTagBoolean](ctx, "ackRequired", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(9)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ackRequired' field"))
	}
	var ackRequired BACnetContextTagBoolean
	if _ackRequired != nil {
		ackRequired = *_ackRequired
	}

	_fromState, err := ReadOptionalField[BACnetEventStateTagged](ctx, "fromState", ReadComplex[BACnetEventStateTagged](BACnetEventStateTaggedParseWithBufferProducer((uint8)(uint8(10)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fromState' field"))
	}
	var fromState BACnetEventStateTagged
	if _fromState != nil {
		fromState = *_fromState
	}

	toState, err := ReadSimpleField[BACnetEventStateTagged](ctx, "toState", ReadComplex[BACnetEventStateTagged](BACnetEventStateTaggedParseWithBufferProducer((uint8)(uint8(11)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toState' field"))
	}

	_eventValues, err := ReadOptionalField[BACnetNotificationParameters](ctx, "eventValues", ReadComplex[BACnetNotificationParameters](BACnetNotificationParametersParseWithBufferProducer[BACnetNotificationParameters]((uint8)(uint8(12)), (BACnetObjectType)(eventObjectIdentifier.GetObjectType())), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventValues' field"))
	}
	var eventValues BACnetNotificationParameters
	if _eventValues != nil {
		eventValues = *_eventValues
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestUnconfirmedEventNotification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestUnconfirmedEventNotification")
	}

	// Create a partially initialized instance
	_child := &_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification{
		_BACnetUnconfirmedServiceRequest: &_BACnetUnconfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		ProcessIdentifier:          processIdentifier,
		InitiatingDeviceIdentifier: initiatingDeviceIdentifier,
		EventObjectIdentifier:      eventObjectIdentifier,
		Timestamp:                  timestamp,
		NotificationClass:          notificationClass,
		Priority:                   priority,
		EventType:                  eventType,
		MessageText:                messageText,
		NotifyType:                 notifyType,
		AckRequired:                ackRequired,
		FromState:                  fromState,
		ToState:                    toState,
		EventValues:                eventValues,
	}
	_child._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestUnconfirmedEventNotification"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestUnconfirmedEventNotification")
		}

		// Simple Field (processIdentifier)
		if pushErr := writeBuffer.PushContext("processIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for processIdentifier")
		}
		_processIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetProcessIdentifier())
		if popErr := writeBuffer.PopContext("processIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for processIdentifier")
		}
		if _processIdentifierErr != nil {
			return errors.Wrap(_processIdentifierErr, "Error serializing 'processIdentifier' field")
		}

		// Simple Field (initiatingDeviceIdentifier)
		if pushErr := writeBuffer.PushContext("initiatingDeviceIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for initiatingDeviceIdentifier")
		}
		_initiatingDeviceIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetInitiatingDeviceIdentifier())
		if popErr := writeBuffer.PopContext("initiatingDeviceIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for initiatingDeviceIdentifier")
		}
		if _initiatingDeviceIdentifierErr != nil {
			return errors.Wrap(_initiatingDeviceIdentifierErr, "Error serializing 'initiatingDeviceIdentifier' field")
		}

		// Simple Field (eventObjectIdentifier)
		if pushErr := writeBuffer.PushContext("eventObjectIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for eventObjectIdentifier")
		}
		_eventObjectIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetEventObjectIdentifier())
		if popErr := writeBuffer.PopContext("eventObjectIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for eventObjectIdentifier")
		}
		if _eventObjectIdentifierErr != nil {
			return errors.Wrap(_eventObjectIdentifierErr, "Error serializing 'eventObjectIdentifier' field")
		}

		// Simple Field (timestamp)
		if pushErr := writeBuffer.PushContext("timestamp"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timestamp")
		}
		_timestampErr := writeBuffer.WriteSerializable(ctx, m.GetTimestamp())
		if popErr := writeBuffer.PopContext("timestamp"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timestamp")
		}
		if _timestampErr != nil {
			return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
		}

		// Simple Field (notificationClass)
		if pushErr := writeBuffer.PushContext("notificationClass"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for notificationClass")
		}
		_notificationClassErr := writeBuffer.WriteSerializable(ctx, m.GetNotificationClass())
		if popErr := writeBuffer.PopContext("notificationClass"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for notificationClass")
		}
		if _notificationClassErr != nil {
			return errors.Wrap(_notificationClassErr, "Error serializing 'notificationClass' field")
		}

		// Simple Field (priority)
		if pushErr := writeBuffer.PushContext("priority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for priority")
		}
		_priorityErr := writeBuffer.WriteSerializable(ctx, m.GetPriority())
		if popErr := writeBuffer.PopContext("priority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for priority")
		}
		if _priorityErr != nil {
			return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
		}

		// Simple Field (eventType)
		if pushErr := writeBuffer.PushContext("eventType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for eventType")
		}
		_eventTypeErr := writeBuffer.WriteSerializable(ctx, m.GetEventType())
		if popErr := writeBuffer.PopContext("eventType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for eventType")
		}
		if _eventTypeErr != nil {
			return errors.Wrap(_eventTypeErr, "Error serializing 'eventType' field")
		}

		if err := WriteOptionalField[BACnetContextTagCharacterString](ctx, "messageText", GetRef(m.GetMessageText()), WriteComplex[BACnetContextTagCharacterString](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'messageText' field")
		}

		// Simple Field (notifyType)
		if pushErr := writeBuffer.PushContext("notifyType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for notifyType")
		}
		_notifyTypeErr := writeBuffer.WriteSerializable(ctx, m.GetNotifyType())
		if popErr := writeBuffer.PopContext("notifyType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for notifyType")
		}
		if _notifyTypeErr != nil {
			return errors.Wrap(_notifyTypeErr, "Error serializing 'notifyType' field")
		}

		if err := WriteOptionalField[BACnetContextTagBoolean](ctx, "ackRequired", GetRef(m.GetAckRequired()), WriteComplex[BACnetContextTagBoolean](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'ackRequired' field")
		}

		if err := WriteOptionalField[BACnetEventStateTagged](ctx, "fromState", GetRef(m.GetFromState()), WriteComplex[BACnetEventStateTagged](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'fromState' field")
		}

		// Simple Field (toState)
		if pushErr := writeBuffer.PushContext("toState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for toState")
		}
		_toStateErr := writeBuffer.WriteSerializable(ctx, m.GetToState())
		if popErr := writeBuffer.PopContext("toState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for toState")
		}
		if _toStateErr != nil {
			return errors.Wrap(_toStateErr, "Error serializing 'toState' field")
		}

		if err := WriteOptionalField[BACnetNotificationParameters](ctx, "eventValues", GetRef(m.GetEventValues()), WriteComplex[BACnetNotificationParameters](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'eventValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestUnconfirmedEventNotification"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestUnconfirmedEventNotification")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) isBACnetUnconfirmedServiceRequestUnconfirmedEventNotification() bool {
	return true
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
