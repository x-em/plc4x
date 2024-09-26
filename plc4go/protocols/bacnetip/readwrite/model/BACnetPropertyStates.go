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

// BACnetPropertyStates is the corresponding interface of BACnetPropertyStates
type BACnetPropertyStates interface {
	BACnetPropertyStatesContract
	BACnetPropertyStatesRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsBACnetPropertyStates is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStates()
}

// BACnetPropertyStatesContract provides a set of functions which can be overwritten by a sub struct
type BACnetPropertyStatesContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// IsBACnetPropertyStates is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStates()
}

// BACnetPropertyStatesRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetPropertyStatesRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetPropertyStates is the data-structure of this message
type _BACnetPropertyStates struct {
	_SubType        BACnetPropertyStates
	PeekedTagHeader BACnetTagHeader
}

var _ BACnetPropertyStatesContract = (*_BACnetPropertyStates)(nil)

// NewBACnetPropertyStates factory function for _BACnetPropertyStates
func NewBACnetPropertyStates(peekedTagHeader BACnetTagHeader) *_BACnetPropertyStates {
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetPropertyStates must not be nil")
	}
	return &_BACnetPropertyStates{PeekedTagHeader: peekedTagHeader}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStates) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetPropertyStates) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStates(structType any) BACnetPropertyStates {
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStates) GetTypeName() string {
	return "BACnetPropertyStates"
}

func (m *_BACnetPropertyStates) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetPropertyStates) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesParse[T BACnetPropertyStates](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetPropertyStatesParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetPropertyStatesParseWithBufferProducer[T BACnetPropertyStates]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetPropertyStatesParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetPropertyStatesParseWithBuffer[T BACnetPropertyStates](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetPropertyStates{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_BACnetPropertyStates) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetPropertyStates BACnetPropertyStates, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStates"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStates")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetPropertyStates
	switch {
	case peekedTagNumber == uint8(0): // BACnetPropertyStatesBoolean
		if _child, err = new(_BACnetPropertyStatesBoolean).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesBoolean for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(1): // BACnetPropertyStatesBinaryValue
		if _child, err = new(_BACnetPropertyStatesBinaryValue).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesBinaryValue for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(2): // BACnetPropertyStatesEventType
		if _child, err = new(_BACnetPropertyStatesEventType).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesEventType for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(3): // BACnetPropertyStatesPolarity
		if _child, err = new(_BACnetPropertyStatesPolarity).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesPolarity for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(4): // BACnetPropertyStatesProgramChange
		if _child, err = new(_BACnetPropertyStatesProgramChange).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesProgramChange for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(5): // BACnetPropertyStatesProgramChange
		if _child, err = new(_BACnetPropertyStatesProgramChange).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesProgramChange for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(6): // BACnetPropertyStatesReasonForHalt
		if _child, err = new(_BACnetPropertyStatesReasonForHalt).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesReasonForHalt for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(7): // BACnetPropertyStatesReliability
		if _child, err = new(_BACnetPropertyStatesReliability).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesReliability for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(8): // BACnetPropertyStatesState
		if _child, err = new(_BACnetPropertyStatesState).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesState for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(9): // BACnetPropertyStatesSystemStatus
		if _child, err = new(_BACnetPropertyStatesSystemStatus).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesSystemStatus for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(10): // BACnetPropertyStatesUnits
		if _child, err = new(_BACnetPropertyStatesUnits).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesUnits for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(11): // BACnetPropertyStatesExtendedValue
		if _child, err = new(_BACnetPropertyStatesExtendedValue).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesExtendedValue for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(12): // BACnetPropertyStatesLifeSafetyMode
		if _child, err = new(_BACnetPropertyStatesLifeSafetyMode).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLifeSafetyMode for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(13): // BACnetPropertyStatesLifeSafetyState
		if _child, err = new(_BACnetPropertyStatesLifeSafetyState).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLifeSafetyState for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(14): // BACnetPropertyStatesRestartReason
		if _child, err = new(_BACnetPropertyStatesRestartReason).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesRestartReason for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(15): // BACnetPropertyStatesDoorAlarmState
		if _child, err = new(_BACnetPropertyStatesDoorAlarmState).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesDoorAlarmState for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(16): // BACnetPropertyStatesAction
		if _child, err = new(_BACnetPropertyStatesAction).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesAction for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(17): // BACnetPropertyStatesDoorSecuredStatus
		if _child, err = new(_BACnetPropertyStatesDoorSecuredStatus).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesDoorSecuredStatus for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(18): // BACnetPropertyStatesDoorStatus
		if _child, err = new(_BACnetPropertyStatesDoorStatus).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesDoorStatus for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(19): // BACnetPropertyStatesDoorValue
		if _child, err = new(_BACnetPropertyStatesDoorValue).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesDoorValue for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(20): // BACnetPropertyStatesFileAccessMethod
		if _child, err = new(_BACnetPropertyStatesFileAccessMethod).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesFileAccessMethod for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(21): // BACnetPropertyStatesLockStatus
		if _child, err = new(_BACnetPropertyStatesLockStatus).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLockStatus for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(22): // BACnetPropertyStatesLifeSafetyOperations
		if _child, err = new(_BACnetPropertyStatesLifeSafetyOperations).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLifeSafetyOperations for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(23): // BACnetPropertyStatesMaintenance
		if _child, err = new(_BACnetPropertyStatesMaintenance).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesMaintenance for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(24): // BACnetPropertyStatesNodeType
		if _child, err = new(_BACnetPropertyStatesNodeType).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesNodeType for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(25): // BACnetPropertyStatesNotifyType
		if _child, err = new(_BACnetPropertyStatesNotifyType).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesNotifyType for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(26): // BACnetPropertyStatesSecurityLevel
		if _child, err = new(_BACnetPropertyStatesSecurityLevel).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesSecurityLevel for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(27): // BACnetPropertyStatesShedState
		if _child, err = new(_BACnetPropertyStatesShedState).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesShedState for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(28): // BACnetPropertyStatesSilencedState
		if _child, err = new(_BACnetPropertyStatesSilencedState).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesSilencedState for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(30): // BACnetPropertyStatesAccessEvent
		if _child, err = new(_BACnetPropertyStatesAccessEvent).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesAccessEvent for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(31): // BACnetPropertyStatesZoneOccupanyState
		if _child, err = new(_BACnetPropertyStatesZoneOccupanyState).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesZoneOccupanyState for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(32): // BACnetPropertyStatesAccessCredentialDisableReason
		if _child, err = new(_BACnetPropertyStatesAccessCredentialDisableReason).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesAccessCredentialDisableReason for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(33): // BACnetPropertyStatesAccessCredentialDisable
		if _child, err = new(_BACnetPropertyStatesAccessCredentialDisable).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesAccessCredentialDisable for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(34): // BACnetPropertyStatesAuthenticationStatus
		if _child, err = new(_BACnetPropertyStatesAuthenticationStatus).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesAuthenticationStatus for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(36): // BACnetPropertyStatesBackupState
		if _child, err = new(_BACnetPropertyStatesBackupState).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesBackupState for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(37): // BACnetPropertyStatesWriteStatus
		if _child, err = new(_BACnetPropertyStatesWriteStatus).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesWriteStatus for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(38): // BACnetPropertyStatesLightningInProgress
		if _child, err = new(_BACnetPropertyStatesLightningInProgress).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLightningInProgress for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(39): // BACnetPropertyStatesLightningOperation
		if _child, err = new(_BACnetPropertyStatesLightningOperation).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLightningOperation for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(40): // BACnetPropertyStatesLightningTransition
		if _child, err = new(_BACnetPropertyStatesLightningTransition).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLightningTransition for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(41): // BACnetPropertyStatesIntegerValue
		if _child, err = new(_BACnetPropertyStatesIntegerValue).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesIntegerValue for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(42): // BACnetPropertyStatesBinaryLightningValue
		if _child, err = new(_BACnetPropertyStatesBinaryLightningValue).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesBinaryLightningValue for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(43): // BACnetPropertyStatesTimerState
		if _child, err = new(_BACnetPropertyStatesTimerState).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesTimerState for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(44): // BACnetPropertyStatesTimerTransition
		if _child, err = new(_BACnetPropertyStatesTimerTransition).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesTimerTransition for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(45): // BACnetPropertyStatesBacnetIpMode
		if _child, err = new(_BACnetPropertyStatesBacnetIpMode).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesBacnetIpMode for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(46): // BACnetPropertyStatesNetworkPortCommand
		if _child, err = new(_BACnetPropertyStatesNetworkPortCommand).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesNetworkPortCommand for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(47): // BACnetPropertyStatesNetworkType
		if _child, err = new(_BACnetPropertyStatesNetworkType).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesNetworkType for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(48): // BACnetPropertyStatesNetworkNumberQuality
		if _child, err = new(_BACnetPropertyStatesNetworkNumberQuality).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesNetworkNumberQuality for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(49): // BACnetPropertyStatesEscalatorOperationDirection
		if _child, err = new(_BACnetPropertyStatesEscalatorOperationDirection).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesEscalatorOperationDirection for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(50): // BACnetPropertyStatesEscalatorFault
		if _child, err = new(_BACnetPropertyStatesEscalatorFault).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesEscalatorFault for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(51): // BACnetPropertyStatesEscalatorMode
		if _child, err = new(_BACnetPropertyStatesEscalatorMode).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesEscalatorMode for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(52): // BACnetPropertyStatesLiftCarDirection
		if _child, err = new(_BACnetPropertyStatesLiftCarDirection).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLiftCarDirection for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(53): // BACnetPropertyStatesLiftCarDoorCommand
		if _child, err = new(_BACnetPropertyStatesLiftCarDoorCommand).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLiftCarDoorCommand for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(54): // BACnetPropertyStatesLiftCarDriveStatus
		if _child, err = new(_BACnetPropertyStatesLiftCarDriveStatus).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLiftCarDriveStatus for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(55): // BACnetPropertyStatesLiftCarMode
		if _child, err = new(_BACnetPropertyStatesLiftCarMode).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLiftCarMode for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(56): // BACnetPropertyStatesLiftGroupMode
		if _child, err = new(_BACnetPropertyStatesLiftGroupMode).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLiftGroupMode for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(57): // BACnetPropertyStatesLiftFault
		if _child, err = new(_BACnetPropertyStatesLiftFault).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLiftFault for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(58): // BACnetPropertyStatesProtocolLevel
		if _child, err = new(_BACnetPropertyStatesProtocolLevel).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesProtocolLevel for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(63): // BACnetPropertyStatesExtendedValue
		if _child, err = new(_BACnetPropertyStatesExtendedValue).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesExtendedValue for type-switch of BACnetPropertyStates")
		}
	case true: // BACnetPropertyStateActionUnknown
		if _child, err = new(_BACnetPropertyStateActionUnknown).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStateActionUnknown for type-switch of BACnetPropertyStates")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStates"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStates")
	}

	return _child, nil
}

func (pm *_BACnetPropertyStates) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetPropertyStates, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetPropertyStates"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStates")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetPropertyStates"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPropertyStates")
	}
	return nil
}

func (m *_BACnetPropertyStates) IsBACnetPropertyStates() {}
