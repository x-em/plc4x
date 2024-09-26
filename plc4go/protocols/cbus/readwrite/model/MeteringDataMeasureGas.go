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

// MeteringDataMeasureGas is the corresponding interface of MeteringDataMeasureGas
type MeteringDataMeasureGas interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MeteringData
	// IsMeteringDataMeasureGas is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMeteringDataMeasureGas()
}

// _MeteringDataMeasureGas is the data-structure of this message
type _MeteringDataMeasureGas struct {
	MeteringDataContract
}

var _ MeteringDataMeasureGas = (*_MeteringDataMeasureGas)(nil)
var _ MeteringDataRequirements = (*_MeteringDataMeasureGas)(nil)

// NewMeteringDataMeasureGas factory function for _MeteringDataMeasureGas
func NewMeteringDataMeasureGas(commandTypeContainer MeteringCommandTypeContainer, argument byte) *_MeteringDataMeasureGas {
	_result := &_MeteringDataMeasureGas{
		MeteringDataContract: NewMeteringData(commandTypeContainer, argument),
	}
	_result.MeteringDataContract.(*_MeteringData)._SubType = _result
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

func (m *_MeteringDataMeasureGas) GetParent() MeteringDataContract {
	return m.MeteringDataContract
}

// Deprecated: use the interface for direct cast
func CastMeteringDataMeasureGas(structType any) MeteringDataMeasureGas {
	if casted, ok := structType.(MeteringDataMeasureGas); ok {
		return casted
	}
	if casted, ok := structType.(*MeteringDataMeasureGas); ok {
		return *casted
	}
	return nil
}

func (m *_MeteringDataMeasureGas) GetTypeName() string {
	return "MeteringDataMeasureGas"
}

func (m *_MeteringDataMeasureGas) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MeteringDataContract.(*_MeteringData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_MeteringDataMeasureGas) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MeteringDataMeasureGas) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MeteringData) (__meteringDataMeasureGas MeteringDataMeasureGas, err error) {
	m.MeteringDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MeteringDataMeasureGas"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeteringDataMeasureGas")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MeteringDataMeasureGas"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeteringDataMeasureGas")
	}

	return m, nil
}

func (m *_MeteringDataMeasureGas) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MeteringDataMeasureGas) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MeteringDataMeasureGas"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MeteringDataMeasureGas")
		}

		if popErr := writeBuffer.PopContext("MeteringDataMeasureGas"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MeteringDataMeasureGas")
		}
		return nil
	}
	return m.MeteringDataContract.(*_MeteringData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MeteringDataMeasureGas) IsMeteringDataMeasureGas() {}

func (m *_MeteringDataMeasureGas) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
