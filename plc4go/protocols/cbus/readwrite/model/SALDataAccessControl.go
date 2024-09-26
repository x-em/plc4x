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

// SALDataAccessControl is the corresponding interface of SALDataAccessControl
type SALDataAccessControl interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SALData
	// GetAccessControlData returns AccessControlData (property field)
	GetAccessControlData() AccessControlData
	// IsSALDataAccessControl is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSALDataAccessControl()
}

// _SALDataAccessControl is the data-structure of this message
type _SALDataAccessControl struct {
	SALDataContract
	AccessControlData AccessControlData
}

var _ SALDataAccessControl = (*_SALDataAccessControl)(nil)
var _ SALDataRequirements = (*_SALDataAccessControl)(nil)

// NewSALDataAccessControl factory function for _SALDataAccessControl
func NewSALDataAccessControl(salData SALData, accessControlData AccessControlData) *_SALDataAccessControl {
	if accessControlData == nil {
		panic("accessControlData of type AccessControlData for SALDataAccessControl must not be nil")
	}
	_result := &_SALDataAccessControl{
		SALDataContract:   NewSALData(salData),
		AccessControlData: accessControlData,
	}
	_result.SALDataContract.(*_SALData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataAccessControl) GetApplicationId() ApplicationId {
	return ApplicationId_ACCESS_CONTROL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataAccessControl) GetParent() SALDataContract {
	return m.SALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataAccessControl) GetAccessControlData() AccessControlData {
	return m.AccessControlData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSALDataAccessControl(structType any) SALDataAccessControl {
	if casted, ok := structType.(SALDataAccessControl); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataAccessControl); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataAccessControl) GetTypeName() string {
	return "SALDataAccessControl"
}

func (m *_SALDataAccessControl) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	// Simple field (accessControlData)
	lengthInBits += m.AccessControlData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataAccessControl) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataAccessControl) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataAccessControl SALDataAccessControl, err error) {
	m.SALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataAccessControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataAccessControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	accessControlData, err := ReadSimpleField[AccessControlData](ctx, "accessControlData", ReadComplex[AccessControlData](AccessControlDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accessControlData' field"))
	}
	m.AccessControlData = accessControlData

	if closeErr := readBuffer.CloseContext("SALDataAccessControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataAccessControl")
	}

	return m, nil
}

func (m *_SALDataAccessControl) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataAccessControl) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataAccessControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataAccessControl")
		}

		if err := WriteSimpleField[AccessControlData](ctx, "accessControlData", m.GetAccessControlData(), WriteComplex[AccessControlData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'accessControlData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataAccessControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataAccessControl")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataAccessControl) IsSALDataAccessControl() {}

func (m *_SALDataAccessControl) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
