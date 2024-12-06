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
	utils.Copyable
	MeteringData
	// IsMeteringDataMeasureGas is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMeteringDataMeasureGas()
	// CreateBuilder creates a MeteringDataMeasureGasBuilder
	CreateMeteringDataMeasureGasBuilder() MeteringDataMeasureGasBuilder
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
/////////////////////// Builder
///////////////////////

// MeteringDataMeasureGasBuilder is a builder for MeteringDataMeasureGas
type MeteringDataMeasureGasBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() MeteringDataMeasureGasBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() MeteringDataBuilder
	// Build builds the MeteringDataMeasureGas or returns an error if something is wrong
	Build() (MeteringDataMeasureGas, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MeteringDataMeasureGas
}

// NewMeteringDataMeasureGasBuilder() creates a MeteringDataMeasureGasBuilder
func NewMeteringDataMeasureGasBuilder() MeteringDataMeasureGasBuilder {
	return &_MeteringDataMeasureGasBuilder{_MeteringDataMeasureGas: new(_MeteringDataMeasureGas)}
}

type _MeteringDataMeasureGasBuilder struct {
	*_MeteringDataMeasureGas

	parentBuilder *_MeteringDataBuilder

	err *utils.MultiError
}

var _ (MeteringDataMeasureGasBuilder) = (*_MeteringDataMeasureGasBuilder)(nil)

func (b *_MeteringDataMeasureGasBuilder) setParent(contract MeteringDataContract) {
	b.MeteringDataContract = contract
	contract.(*_MeteringData)._SubType = b._MeteringDataMeasureGas
}

func (b *_MeteringDataMeasureGasBuilder) WithMandatoryFields() MeteringDataMeasureGasBuilder {
	return b
}

func (b *_MeteringDataMeasureGasBuilder) Build() (MeteringDataMeasureGas, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MeteringDataMeasureGas.deepCopy(), nil
}

func (b *_MeteringDataMeasureGasBuilder) MustBuild() MeteringDataMeasureGas {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_MeteringDataMeasureGasBuilder) Done() MeteringDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewMeteringDataBuilder().(*_MeteringDataBuilder)
	}
	return b.parentBuilder
}

func (b *_MeteringDataMeasureGasBuilder) buildForMeteringData() (MeteringData, error) {
	return b.Build()
}

func (b *_MeteringDataMeasureGasBuilder) DeepCopy() any {
	_copy := b.CreateMeteringDataMeasureGasBuilder().(*_MeteringDataMeasureGasBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMeteringDataMeasureGasBuilder creates a MeteringDataMeasureGasBuilder
func (b *_MeteringDataMeasureGas) CreateMeteringDataMeasureGasBuilder() MeteringDataMeasureGasBuilder {
	if b == nil {
		return NewMeteringDataMeasureGasBuilder()
	}
	return &_MeteringDataMeasureGasBuilder{_MeteringDataMeasureGas: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

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

func (m *_MeteringDataMeasureGas) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MeteringDataMeasureGas) deepCopy() *_MeteringDataMeasureGas {
	if m == nil {
		return nil
	}
	_MeteringDataMeasureGasCopy := &_MeteringDataMeasureGas{
		m.MeteringDataContract.(*_MeteringData).deepCopy(),
	}
	_MeteringDataMeasureGasCopy.MeteringDataContract.(*_MeteringData)._SubType = m
	return _MeteringDataMeasureGasCopy
}

func (m *_MeteringDataMeasureGas) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}