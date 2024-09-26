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

// CipWriteRequest is the corresponding interface of CipWriteRequest
type CipWriteRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
	// GetTag returns Tag (property field)
	GetTag() []byte
	// GetDataType returns DataType (property field)
	GetDataType() CIPDataTypeCode
	// GetElementNb returns ElementNb (property field)
	GetElementNb() uint16
	// GetData returns Data (property field)
	GetData() []byte
	// IsCipWriteRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCipWriteRequest()
}

// _CipWriteRequest is the data-structure of this message
type _CipWriteRequest struct {
	CipServiceContract
	Tag       []byte
	DataType  CIPDataTypeCode
	ElementNb uint16
	Data      []byte
}

var _ CipWriteRequest = (*_CipWriteRequest)(nil)
var _ CipServiceRequirements = (*_CipWriteRequest)(nil)

// NewCipWriteRequest factory function for _CipWriteRequest
func NewCipWriteRequest(tag []byte, dataType CIPDataTypeCode, elementNb uint16, data []byte, serviceLen uint16) *_CipWriteRequest {
	_result := &_CipWriteRequest{
		CipServiceContract: NewCipService(serviceLen),
		Tag:                tag,
		DataType:           dataType,
		ElementNb:          elementNb,
		Data:               data,
	}
	_result.CipServiceContract.(*_CipService)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipWriteRequest) GetService() uint8 {
	return 0x4D
}

func (m *_CipWriteRequest) GetResponse() bool {
	return bool(false)
}

func (m *_CipWriteRequest) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipWriteRequest) GetParent() CipServiceContract {
	return m.CipServiceContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipWriteRequest) GetTag() []byte {
	return m.Tag
}

func (m *_CipWriteRequest) GetDataType() CIPDataTypeCode {
	return m.DataType
}

func (m *_CipWriteRequest) GetElementNb() uint16 {
	return m.ElementNb
}

func (m *_CipWriteRequest) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCipWriteRequest(structType any) CipWriteRequest {
	if casted, ok := structType.(CipWriteRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CipWriteRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CipWriteRequest) GetTypeName() string {
	return "CipWriteRequest"
}

func (m *_CipWriteRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	// Implicit Field (requestPathSize)
	lengthInBits += 8

	// Array field
	if len(m.Tag) > 0 {
		lengthInBits += 8 * uint16(len(m.Tag))
	}

	// Simple field (dataType)
	lengthInBits += 16

	// Simple field (elementNb)
	lengthInBits += 16

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_CipWriteRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CipWriteRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__cipWriteRequest CipWriteRequest, err error) {
	m.CipServiceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipWriteRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipWriteRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestPathSize, err := ReadImplicitField[uint8](ctx, "requestPathSize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestPathSize' field"))
	}
	_ = requestPathSize

	tag, err := readBuffer.ReadByteArray("tag", int(int32(requestPathSize)*int32(int32(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tag' field"))
	}
	m.Tag = tag

	dataType, err := ReadEnumField[CIPDataTypeCode](ctx, "dataType", "CIPDataTypeCode", ReadEnum(CIPDataTypeCodeByValue, ReadUnsignedShort(readBuffer, uint8(16))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataType' field"))
	}
	m.DataType = dataType

	elementNb, err := ReadSimpleField(ctx, "elementNb", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'elementNb' field"))
	}
	m.ElementNb = elementNb

	data, err := readBuffer.ReadByteArray("data", int(int32(dataType.Size())*int32(elementNb)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("CipWriteRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipWriteRequest")
	}

	return m, nil
}

func (m *_CipWriteRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipWriteRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipWriteRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipWriteRequest")
		}
		requestPathSize := uint8(uint8(uint8(len(m.GetTag()))) / uint8(uint8(2)))
		if err := WriteImplicitField(ctx, "requestPathSize", requestPathSize, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestPathSize' field")
		}

		if err := WriteByteArrayField(ctx, "tag", m.GetTag(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'tag' field")
		}

		if err := WriteSimpleEnumField[CIPDataTypeCode](ctx, "dataType", "CIPDataTypeCode", m.GetDataType(), WriteEnum[CIPDataTypeCode, uint16](CIPDataTypeCode.GetValue, CIPDataTypeCode.PLC4XEnumName, WriteUnsignedShort(writeBuffer, 16))); err != nil {
			return errors.Wrap(err, "Error serializing 'dataType' field")
		}

		if err := WriteSimpleField[uint16](ctx, "elementNb", m.GetElementNb(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'elementNb' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("CipWriteRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipWriteRequest")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipWriteRequest) IsCipWriteRequest() {}

func (m *_CipWriteRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
