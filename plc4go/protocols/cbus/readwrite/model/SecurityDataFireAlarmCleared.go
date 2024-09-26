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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// SecurityDataFireAlarmCleared is the corresponding interface of SecurityDataFireAlarmCleared
type SecurityDataFireAlarmCleared interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
	// IsSecurityDataFireAlarmCleared is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataFireAlarmCleared()
}

// _SecurityDataFireAlarmCleared is the data-structure of this message
type _SecurityDataFireAlarmCleared struct {
	SecurityDataContract
}

var _ SecurityDataFireAlarmCleared = (*_SecurityDataFireAlarmCleared)(nil)
var _ SecurityDataRequirements = (*_SecurityDataFireAlarmCleared)(nil)

// NewSecurityDataFireAlarmCleared factory function for _SecurityDataFireAlarmCleared
func NewSecurityDataFireAlarmCleared(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataFireAlarmCleared {
	_result := &_SecurityDataFireAlarmCleared{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataFireAlarmCleared) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// Deprecated: use the interface for direct cast
func CastSecurityDataFireAlarmCleared(structType any) SecurityDataFireAlarmCleared {
	if casted, ok := structType.(SecurityDataFireAlarmCleared); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataFireAlarmCleared); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataFireAlarmCleared) GetTypeName() string {
	return "SecurityDataFireAlarmCleared"
}

func (m *_SecurityDataFireAlarmCleared) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataFireAlarmCleared) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataFireAlarmCleared) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataFireAlarmCleared SecurityDataFireAlarmCleared, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataFireAlarmCleared"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataFireAlarmCleared")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataFireAlarmCleared"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataFireAlarmCleared")
	}

	return m, nil
}

func (m *_SecurityDataFireAlarmCleared) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataFireAlarmCleared) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataFireAlarmCleared"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataFireAlarmCleared")
		}

		if popErr := writeBuffer.PopContext("SecurityDataFireAlarmCleared"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataFireAlarmCleared")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataFireAlarmCleared) IsSecurityDataFireAlarmCleared() {}

func (m *_SecurityDataFireAlarmCleared) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
