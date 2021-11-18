/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const BACnetUnconfirmedServiceRequestIAm_OBJECTIDENTIFIERHEADER uint8 = 0xC4
const BACnetUnconfirmedServiceRequestIAm_MAXIMUMAPDULENGTHACCEPTEDHEADER uint8 = 0x04
const BACnetUnconfirmedServiceRequestIAm_SEGMENTATIONSUPPORTEDHEADER uint8 = 0x91
const BACnetUnconfirmedServiceRequestIAm_VENDORIDHEADER uint8 = 0x21

// The data-structure of this message
type BACnetUnconfirmedServiceRequestIAm struct {
	ObjectType                      uint16
	ObjectInstanceNumber            uint32
	MaximumApduLengthAcceptedLength uint8
	MaximumApduLengthAccepted       []int8
	SegmentationSupported           uint8
	VendorId                        uint8
	Parent                          *BACnetUnconfirmedServiceRequest
}

// The corresponding interface
type IBACnetUnconfirmedServiceRequestIAm interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetUnconfirmedServiceRequestIAm) ServiceChoice() uint8 {
	return 0x00
}

func (m *BACnetUnconfirmedServiceRequestIAm) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func NewBACnetUnconfirmedServiceRequestIAm(objectType uint16, objectInstanceNumber uint32, maximumApduLengthAcceptedLength uint8, maximumApduLengthAccepted []int8, segmentationSupported uint8, vendorId uint8) *BACnetUnconfirmedServiceRequest {
	child := &BACnetUnconfirmedServiceRequestIAm{
		ObjectType:                      objectType,
		ObjectInstanceNumber:            objectInstanceNumber,
		MaximumApduLengthAcceptedLength: maximumApduLengthAcceptedLength,
		MaximumApduLengthAccepted:       maximumApduLengthAccepted,
		SegmentationSupported:           segmentationSupported,
		VendorId:                        vendorId,
		Parent:                          NewBACnetUnconfirmedServiceRequest(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetUnconfirmedServiceRequestIAm(structType interface{}) *BACnetUnconfirmedServiceRequestIAm {
	castFunc := func(typ interface{}) *BACnetUnconfirmedServiceRequestIAm {
		if casted, ok := typ.(BACnetUnconfirmedServiceRequestIAm); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequestIAm); ok {
			return casted
		}
		if casted, ok := typ.(BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestIAm(casted.Child)
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestIAm(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetUnconfirmedServiceRequestIAm) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestIAm"
}

func (m *BACnetUnconfirmedServiceRequestIAm) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestIAm) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Const Field (objectIdentifierHeader)
	lengthInBits += 8

	// Simple field (objectType)
	lengthInBits += 10

	// Simple field (objectInstanceNumber)
	lengthInBits += 22

	// Const Field (maximumApduLengthAcceptedHeader)
	lengthInBits += 5

	// Simple field (maximumApduLengthAcceptedLength)
	lengthInBits += 3

	// Array field
	if len(m.MaximumApduLengthAccepted) > 0 {
		lengthInBits += 8 * uint16(len(m.MaximumApduLengthAccepted))
	}

	// Const Field (segmentationSupportedHeader)
	lengthInBits += 8

	// Simple field (segmentationSupported)
	lengthInBits += 8

	// Const Field (vendorIdHeader)
	lengthInBits += 8

	// Simple field (vendorId)
	lengthInBits += 8

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestIAm) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestIAmParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetUnconfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestIAm"); pullErr != nil {
		return nil, pullErr
	}

	// Const Field (objectIdentifierHeader)
	objectIdentifierHeader, _objectIdentifierHeaderErr := readBuffer.ReadUint8("objectIdentifierHeader", 8)
	if _objectIdentifierHeaderErr != nil {
		return nil, errors.Wrap(_objectIdentifierHeaderErr, "Error parsing 'objectIdentifierHeader' field")
	}
	if objectIdentifierHeader != BACnetUnconfirmedServiceRequestIAm_OBJECTIDENTIFIERHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetUnconfirmedServiceRequestIAm_OBJECTIDENTIFIERHEADER) + " but got " + fmt.Sprintf("%d", objectIdentifierHeader))
	}

	// Simple Field (objectType)
	objectType, _objectTypeErr := readBuffer.ReadUint16("objectType", 10)
	if _objectTypeErr != nil {
		return nil, errors.Wrap(_objectTypeErr, "Error parsing 'objectType' field")
	}

	// Simple Field (objectInstanceNumber)
	objectInstanceNumber, _objectInstanceNumberErr := readBuffer.ReadUint32("objectInstanceNumber", 22)
	if _objectInstanceNumberErr != nil {
		return nil, errors.Wrap(_objectInstanceNumberErr, "Error parsing 'objectInstanceNumber' field")
	}

	// Const Field (maximumApduLengthAcceptedHeader)
	maximumApduLengthAcceptedHeader, _maximumApduLengthAcceptedHeaderErr := readBuffer.ReadUint8("maximumApduLengthAcceptedHeader", 5)
	if _maximumApduLengthAcceptedHeaderErr != nil {
		return nil, errors.Wrap(_maximumApduLengthAcceptedHeaderErr, "Error parsing 'maximumApduLengthAcceptedHeader' field")
	}
	if maximumApduLengthAcceptedHeader != BACnetUnconfirmedServiceRequestIAm_MAXIMUMAPDULENGTHACCEPTEDHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetUnconfirmedServiceRequestIAm_MAXIMUMAPDULENGTHACCEPTEDHEADER) + " but got " + fmt.Sprintf("%d", maximumApduLengthAcceptedHeader))
	}

	// Simple Field (maximumApduLengthAcceptedLength)
	maximumApduLengthAcceptedLength, _maximumApduLengthAcceptedLengthErr := readBuffer.ReadUint8("maximumApduLengthAcceptedLength", 3)
	if _maximumApduLengthAcceptedLengthErr != nil {
		return nil, errors.Wrap(_maximumApduLengthAcceptedLengthErr, "Error parsing 'maximumApduLengthAcceptedLength' field")
	}

	// Array field (maximumApduLengthAccepted)
	if pullErr := readBuffer.PullContext("maximumApduLengthAccepted", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	maximumApduLengthAccepted := make([]int8, maximumApduLengthAcceptedLength)
	for curItem := uint16(0); curItem < uint16(maximumApduLengthAcceptedLength); curItem++ {
		_item, _err := readBuffer.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'maximumApduLengthAccepted' field")
		}
		maximumApduLengthAccepted[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("maximumApduLengthAccepted", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	// Const Field (segmentationSupportedHeader)
	segmentationSupportedHeader, _segmentationSupportedHeaderErr := readBuffer.ReadUint8("segmentationSupportedHeader", 8)
	if _segmentationSupportedHeaderErr != nil {
		return nil, errors.Wrap(_segmentationSupportedHeaderErr, "Error parsing 'segmentationSupportedHeader' field")
	}
	if segmentationSupportedHeader != BACnetUnconfirmedServiceRequestIAm_SEGMENTATIONSUPPORTEDHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetUnconfirmedServiceRequestIAm_SEGMENTATIONSUPPORTEDHEADER) + " but got " + fmt.Sprintf("%d", segmentationSupportedHeader))
	}

	// Simple Field (segmentationSupported)
	segmentationSupported, _segmentationSupportedErr := readBuffer.ReadUint8("segmentationSupported", 8)
	if _segmentationSupportedErr != nil {
		return nil, errors.Wrap(_segmentationSupportedErr, "Error parsing 'segmentationSupported' field")
	}

	// Const Field (vendorIdHeader)
	vendorIdHeader, _vendorIdHeaderErr := readBuffer.ReadUint8("vendorIdHeader", 8)
	if _vendorIdHeaderErr != nil {
		return nil, errors.Wrap(_vendorIdHeaderErr, "Error parsing 'vendorIdHeader' field")
	}
	if vendorIdHeader != BACnetUnconfirmedServiceRequestIAm_VENDORIDHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetUnconfirmedServiceRequestIAm_VENDORIDHEADER) + " but got " + fmt.Sprintf("%d", vendorIdHeader))
	}

	// Simple Field (vendorId)
	vendorId, _vendorIdErr := readBuffer.ReadUint8("vendorId", 8)
	if _vendorIdErr != nil {
		return nil, errors.Wrap(_vendorIdErr, "Error parsing 'vendorId' field")
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestIAm"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestIAm{
		ObjectType:                      objectType,
		ObjectInstanceNumber:            objectInstanceNumber,
		MaximumApduLengthAcceptedLength: maximumApduLengthAcceptedLength,
		MaximumApduLengthAccepted:       maximumApduLengthAccepted,
		SegmentationSupported:           segmentationSupported,
		VendorId:                        vendorId,
		Parent:                          &BACnetUnconfirmedServiceRequest{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetUnconfirmedServiceRequestIAm) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestIAm"); pushErr != nil {
			return pushErr
		}

		// Const Field (objectIdentifierHeader)
		_objectIdentifierHeaderErr := writeBuffer.WriteUint8("objectIdentifierHeader", 8, 0xC4)
		if _objectIdentifierHeaderErr != nil {
			return errors.Wrap(_objectIdentifierHeaderErr, "Error serializing 'objectIdentifierHeader' field")
		}

		// Simple Field (objectType)
		objectType := uint16(m.ObjectType)
		_objectTypeErr := writeBuffer.WriteUint16("objectType", 10, (objectType))
		if _objectTypeErr != nil {
			return errors.Wrap(_objectTypeErr, "Error serializing 'objectType' field")
		}

		// Simple Field (objectInstanceNumber)
		objectInstanceNumber := uint32(m.ObjectInstanceNumber)
		_objectInstanceNumberErr := writeBuffer.WriteUint32("objectInstanceNumber", 22, (objectInstanceNumber))
		if _objectInstanceNumberErr != nil {
			return errors.Wrap(_objectInstanceNumberErr, "Error serializing 'objectInstanceNumber' field")
		}

		// Const Field (maximumApduLengthAcceptedHeader)
		_maximumApduLengthAcceptedHeaderErr := writeBuffer.WriteUint8("maximumApduLengthAcceptedHeader", 5, 0x04)
		if _maximumApduLengthAcceptedHeaderErr != nil {
			return errors.Wrap(_maximumApduLengthAcceptedHeaderErr, "Error serializing 'maximumApduLengthAcceptedHeader' field")
		}

		// Simple Field (maximumApduLengthAcceptedLength)
		maximumApduLengthAcceptedLength := uint8(m.MaximumApduLengthAcceptedLength)
		_maximumApduLengthAcceptedLengthErr := writeBuffer.WriteUint8("maximumApduLengthAcceptedLength", 3, (maximumApduLengthAcceptedLength))
		if _maximumApduLengthAcceptedLengthErr != nil {
			return errors.Wrap(_maximumApduLengthAcceptedLengthErr, "Error serializing 'maximumApduLengthAcceptedLength' field")
		}

		// Array Field (maximumApduLengthAccepted)
		if m.MaximumApduLengthAccepted != nil {
			if pushErr := writeBuffer.PushContext("maximumApduLengthAccepted", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.MaximumApduLengthAccepted {
				_elementErr := writeBuffer.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'maximumApduLengthAccepted' field")
				}
			}
			if popErr := writeBuffer.PopContext("maximumApduLengthAccepted", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		// Const Field (segmentationSupportedHeader)
		_segmentationSupportedHeaderErr := writeBuffer.WriteUint8("segmentationSupportedHeader", 8, 0x91)
		if _segmentationSupportedHeaderErr != nil {
			return errors.Wrap(_segmentationSupportedHeaderErr, "Error serializing 'segmentationSupportedHeader' field")
		}

		// Simple Field (segmentationSupported)
		segmentationSupported := uint8(m.SegmentationSupported)
		_segmentationSupportedErr := writeBuffer.WriteUint8("segmentationSupported", 8, (segmentationSupported))
		if _segmentationSupportedErr != nil {
			return errors.Wrap(_segmentationSupportedErr, "Error serializing 'segmentationSupported' field")
		}

		// Const Field (vendorIdHeader)
		_vendorIdHeaderErr := writeBuffer.WriteUint8("vendorIdHeader", 8, 0x21)
		if _vendorIdHeaderErr != nil {
			return errors.Wrap(_vendorIdHeaderErr, "Error serializing 'vendorIdHeader' field")
		}

		// Simple Field (vendorId)
		vendorId := uint8(m.VendorId)
		_vendorIdErr := writeBuffer.WriteUint8("vendorId", 8, (vendorId))
		if _vendorIdErr != nil {
			return errors.Wrap(_vendorIdErr, "Error serializing 'vendorId' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestIAm"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestIAm) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
