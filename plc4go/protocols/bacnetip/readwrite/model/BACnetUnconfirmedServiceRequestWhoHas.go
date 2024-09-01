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

// BACnetUnconfirmedServiceRequestWhoHas is the corresponding interface of BACnetUnconfirmedServiceRequestWhoHas
type BACnetUnconfirmedServiceRequestWhoHas interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetUnconfirmedServiceRequest
	// GetDeviceInstanceRangeLowLimit returns DeviceInstanceRangeLowLimit (property field)
	GetDeviceInstanceRangeLowLimit() BACnetContextTagUnsignedInteger
	// GetDeviceInstanceRangeHighLimit returns DeviceInstanceRangeHighLimit (property field)
	GetDeviceInstanceRangeHighLimit() BACnetContextTagUnsignedInteger
	// GetObject returns Object (property field)
	GetObject() BACnetUnconfirmedServiceRequestWhoHasObject
}

// BACnetUnconfirmedServiceRequestWhoHasExactly can be used when we want exactly this type and not a type which fulfills BACnetUnconfirmedServiceRequestWhoHas.
// This is useful for switch cases.
type BACnetUnconfirmedServiceRequestWhoHasExactly interface {
	BACnetUnconfirmedServiceRequestWhoHas
	isBACnetUnconfirmedServiceRequestWhoHas() bool
}

// _BACnetUnconfirmedServiceRequestWhoHas is the data-structure of this message
type _BACnetUnconfirmedServiceRequestWhoHas struct {
	*_BACnetUnconfirmedServiceRequest
	DeviceInstanceRangeLowLimit  BACnetContextTagUnsignedInteger
	DeviceInstanceRangeHighLimit BACnetContextTagUnsignedInteger
	Object                       BACnetUnconfirmedServiceRequestWhoHasObject
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestWhoHas) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_WHO_HAS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestWhoHas) InitializeParent(parent BACnetUnconfirmedServiceRequest) {
}

func (m *_BACnetUnconfirmedServiceRequestWhoHas) GetParent() BACnetUnconfirmedServiceRequest {
	return m._BACnetUnconfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestWhoHas) GetDeviceInstanceRangeLowLimit() BACnetContextTagUnsignedInteger {
	return m.DeviceInstanceRangeLowLimit
}

func (m *_BACnetUnconfirmedServiceRequestWhoHas) GetDeviceInstanceRangeHighLimit() BACnetContextTagUnsignedInteger {
	return m.DeviceInstanceRangeHighLimit
}

func (m *_BACnetUnconfirmedServiceRequestWhoHas) GetObject() BACnetUnconfirmedServiceRequestWhoHasObject {
	return m.Object
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestWhoHas factory function for _BACnetUnconfirmedServiceRequestWhoHas
func NewBACnetUnconfirmedServiceRequestWhoHas(deviceInstanceRangeLowLimit BACnetContextTagUnsignedInteger, deviceInstanceRangeHighLimit BACnetContextTagUnsignedInteger, object BACnetUnconfirmedServiceRequestWhoHasObject, serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequestWhoHas {
	_result := &_BACnetUnconfirmedServiceRequestWhoHas{
		DeviceInstanceRangeLowLimit:      deviceInstanceRangeLowLimit,
		DeviceInstanceRangeHighLimit:     deviceInstanceRangeHighLimit,
		Object:                           object,
		_BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestWhoHas(structType any) BACnetUnconfirmedServiceRequestWhoHas {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestWhoHas); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestWhoHas); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestWhoHas) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestWhoHas"
}

func (m *_BACnetUnconfirmedServiceRequestWhoHas) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Optional Field (deviceInstanceRangeLowLimit)
	if m.DeviceInstanceRangeLowLimit != nil {
		lengthInBits += m.DeviceInstanceRangeLowLimit.GetLengthInBits(ctx)
	}

	// Optional Field (deviceInstanceRangeHighLimit)
	if m.DeviceInstanceRangeHighLimit != nil {
		lengthInBits += m.DeviceInstanceRangeHighLimit.GetLengthInBits(ctx)
	}

	// Simple field (object)
	lengthInBits += m.Object.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestWhoHas) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetUnconfirmedServiceRequestWhoHasParse(ctx context.Context, theBytes []byte, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequestWhoHas, error) {
	return BACnetUnconfirmedServiceRequestWhoHasParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetUnconfirmedServiceRequestWhoHasParseWithBufferProducer(serviceRequestLength uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetUnconfirmedServiceRequestWhoHas, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetUnconfirmedServiceRequestWhoHas, error) {
		return BACnetUnconfirmedServiceRequestWhoHasParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	}
}

func BACnetUnconfirmedServiceRequestWhoHasParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequestWhoHas, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestWhoHas"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestWhoHas")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	_deviceInstanceRangeLowLimit, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "deviceInstanceRangeLowLimit", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceInstanceRangeLowLimit' field"))
	}
	var deviceInstanceRangeLowLimit BACnetContextTagUnsignedInteger
	if _deviceInstanceRangeLowLimit != nil {
		deviceInstanceRangeLowLimit = *_deviceInstanceRangeLowLimit
	}

	_deviceInstanceRangeHighLimit, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "deviceInstanceRangeHighLimit", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), bool((deviceInstanceRangeLowLimit) != (nil)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceInstanceRangeHighLimit' field"))
	}
	var deviceInstanceRangeHighLimit BACnetContextTagUnsignedInteger
	if _deviceInstanceRangeHighLimit != nil {
		deviceInstanceRangeHighLimit = *_deviceInstanceRangeHighLimit
	}

	object, err := ReadSimpleField[BACnetUnconfirmedServiceRequestWhoHasObject](ctx, "object", ReadComplex[BACnetUnconfirmedServiceRequestWhoHasObject](BACnetUnconfirmedServiceRequestWhoHasObjectParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'object' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestWhoHas"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestWhoHas")
	}

	// Create a partially initialized instance
	_child := &_BACnetUnconfirmedServiceRequestWhoHas{
		_BACnetUnconfirmedServiceRequest: &_BACnetUnconfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		DeviceInstanceRangeLowLimit:  deviceInstanceRangeLowLimit,
		DeviceInstanceRangeHighLimit: deviceInstanceRangeHighLimit,
		Object:                       object,
	}
	_child._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetUnconfirmedServiceRequestWhoHas) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestWhoHas) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestWhoHas"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestWhoHas")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "deviceInstanceRangeLowLimit", GetRef(m.GetDeviceInstanceRangeLowLimit()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'deviceInstanceRangeLowLimit' field")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "deviceInstanceRangeHighLimit", GetRef(m.GetDeviceInstanceRangeHighLimit()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'deviceInstanceRangeHighLimit' field")
		}

		// Simple Field (object)
		if pushErr := writeBuffer.PushContext("object"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for object")
		}
		_objectErr := writeBuffer.WriteSerializable(ctx, m.GetObject())
		if popErr := writeBuffer.PopContext("object"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for object")
		}
		if _objectErr != nil {
			return errors.Wrap(_objectErr, "Error serializing 'object' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestWhoHas"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestWhoHas")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestWhoHas) isBACnetUnconfirmedServiceRequestWhoHas() bool {
	return true
}

func (m *_BACnetUnconfirmedServiceRequestWhoHas) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
