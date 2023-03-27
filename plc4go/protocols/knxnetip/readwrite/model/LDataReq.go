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
	spiContext "github.com/apache/plc4x/plc4go/spi/context"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// LDataReq is the corresponding interface of LDataReq
type LDataReq interface {
	utils.LengthAware
	utils.Serializable
	CEMI
	// GetAdditionalInformationLength returns AdditionalInformationLength (property field)
	GetAdditionalInformationLength() uint8
	// GetAdditionalInformation returns AdditionalInformation (property field)
	GetAdditionalInformation() []CEMIAdditionalInformation
	// GetDataFrame returns DataFrame (property field)
	GetDataFrame() LDataFrame
}

// LDataReqExactly can be used when we want exactly this type and not a type which fulfills LDataReq.
// This is useful for switch cases.
type LDataReqExactly interface {
	LDataReq
	isLDataReq() bool
}

// _LDataReq is the data-structure of this message
type _LDataReq struct {
	*_CEMI
	AdditionalInformationLength uint8
	AdditionalInformation       []CEMIAdditionalInformation
	DataFrame                   LDataFrame
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LDataReq) GetMessageCode() uint8 {
	return 0x11
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LDataReq) InitializeParent(parent CEMI) {}

func (m *_LDataReq) GetParent() CEMI {
	return m._CEMI
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LDataReq) GetAdditionalInformationLength() uint8 {
	return m.AdditionalInformationLength
}

func (m *_LDataReq) GetAdditionalInformation() []CEMIAdditionalInformation {
	return m.AdditionalInformation
}

func (m *_LDataReq) GetDataFrame() LDataFrame {
	return m.DataFrame
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLDataReq factory function for _LDataReq
func NewLDataReq(additionalInformationLength uint8, additionalInformation []CEMIAdditionalInformation, dataFrame LDataFrame, size uint16) *_LDataReq {
	_result := &_LDataReq{
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation:       additionalInformation,
		DataFrame:                   dataFrame,
		_CEMI:                       NewCEMI(size),
	}
	_result._CEMI._CEMIChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLDataReq(structType interface{}) LDataReq {
	if casted, ok := structType.(LDataReq); ok {
		return casted
	}
	if casted, ok := structType.(*LDataReq); ok {
		return *casted
	}
	return nil
}

func (m *_LDataReq) GetTypeName() string {
	return "LDataReq"
}

func (m *_LDataReq) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (additionalInformationLength)
	lengthInBits += 8

	// Array field
	if len(m.AdditionalInformation) > 0 {
		for _, element := range m.AdditionalInformation {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (dataFrame)
	lengthInBits += m.DataFrame.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_LDataReq) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LDataReqParse(theBytes []byte, size uint16) (LDataReq, error) {
	return LDataReqParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), size)
}

func LDataReqParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, size uint16) (LDataReq, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LDataReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LDataReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (additionalInformationLength)
	_additionalInformationLength, _additionalInformationLengthErr := readBuffer.ReadUint8("additionalInformationLength", 8)
	if _additionalInformationLengthErr != nil {
		return nil, errors.Wrap(_additionalInformationLengthErr, "Error parsing 'additionalInformationLength' field of LDataReq")
	}
	additionalInformationLength := _additionalInformationLength

	// Array field (additionalInformation)
	if pullErr := readBuffer.PullContext("additionalInformation", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for additionalInformation")
	}
	// Length array
	var additionalInformation []CEMIAdditionalInformation
	{
		_additionalInformationLength := additionalInformationLength
		_additionalInformationEndPos := positionAware.GetPos() + uint16(_additionalInformationLength)
		for positionAware.GetPos() < _additionalInformationEndPos {
			_item, _err := CEMIAdditionalInformationParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'additionalInformation' field of LDataReq")
			}
			additionalInformation = append(additionalInformation, _item.(CEMIAdditionalInformation))
		}
	}
	if closeErr := readBuffer.CloseContext("additionalInformation", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for additionalInformation")
	}

	// Simple Field (dataFrame)
	if pullErr := readBuffer.PullContext("dataFrame"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dataFrame")
	}
	_dataFrame, _dataFrameErr := LDataFrameParseWithBuffer(ctx, readBuffer)
	if _dataFrameErr != nil {
		return nil, errors.Wrap(_dataFrameErr, "Error parsing 'dataFrame' field of LDataReq")
	}
	dataFrame := _dataFrame.(LDataFrame)
	if closeErr := readBuffer.CloseContext("dataFrame"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dataFrame")
	}

	if closeErr := readBuffer.CloseContext("LDataReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LDataReq")
	}

	// Create a partially initialized instance
	_child := &_LDataReq{
		_CEMI: &_CEMI{
			Size: size,
		},
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation:       additionalInformation,
		DataFrame:                   dataFrame,
	}
	_child._CEMI._CEMIChildRequirements = _child
	return _child, nil
}

func (m *_LDataReq) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LDataReq) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LDataReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LDataReq")
		}

		// Simple Field (additionalInformationLength)
		additionalInformationLength := uint8(m.GetAdditionalInformationLength())
		_additionalInformationLengthErr := writeBuffer.WriteUint8("additionalInformationLength", 8, (additionalInformationLength))
		if _additionalInformationLengthErr != nil {
			return errors.Wrap(_additionalInformationLengthErr, "Error serializing 'additionalInformationLength' field")
		}

		// Array Field (additionalInformation)
		if pushErr := writeBuffer.PushContext("additionalInformation", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for additionalInformation")
		}
		for _curItem, _element := range m.GetAdditionalInformation() {
			_ = _curItem
			arrayCtx := spiContext.CreateArrayContext(ctx, len(m.GetAdditionalInformation()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'additionalInformation' field")
			}
		}
		if popErr := writeBuffer.PopContext("additionalInformation", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for additionalInformation")
		}

		// Simple Field (dataFrame)
		if pushErr := writeBuffer.PushContext("dataFrame"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dataFrame")
		}
		_dataFrameErr := writeBuffer.WriteSerializable(ctx, m.GetDataFrame())
		if popErr := writeBuffer.PopContext("dataFrame"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dataFrame")
		}
		if _dataFrameErr != nil {
			return errors.Wrap(_dataFrameErr, "Error serializing 'dataFrame' field")
		}

		if popErr := writeBuffer.PopContext("LDataReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LDataReq")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LDataReq) isLDataReq() bool {
	return true
}

func (m *_LDataReq) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
