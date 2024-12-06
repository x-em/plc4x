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

// SALDataTriggerControl is the corresponding interface of SALDataTriggerControl
type SALDataTriggerControl interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SALData
	// GetTriggerControlData returns TriggerControlData (property field)
	GetTriggerControlData() TriggerControlData
	// IsSALDataTriggerControl is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSALDataTriggerControl()
	// CreateBuilder creates a SALDataTriggerControlBuilder
	CreateSALDataTriggerControlBuilder() SALDataTriggerControlBuilder
}

// _SALDataTriggerControl is the data-structure of this message
type _SALDataTriggerControl struct {
	SALDataContract
	TriggerControlData TriggerControlData
}

var _ SALDataTriggerControl = (*_SALDataTriggerControl)(nil)
var _ SALDataRequirements = (*_SALDataTriggerControl)(nil)

// NewSALDataTriggerControl factory function for _SALDataTriggerControl
func NewSALDataTriggerControl(salData SALData, triggerControlData TriggerControlData) *_SALDataTriggerControl {
	if triggerControlData == nil {
		panic("triggerControlData of type TriggerControlData for SALDataTriggerControl must not be nil")
	}
	_result := &_SALDataTriggerControl{
		SALDataContract:    NewSALData(salData),
		TriggerControlData: triggerControlData,
	}
	_result.SALDataContract.(*_SALData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SALDataTriggerControlBuilder is a builder for SALDataTriggerControl
type SALDataTriggerControlBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(triggerControlData TriggerControlData) SALDataTriggerControlBuilder
	// WithTriggerControlData adds TriggerControlData (property field)
	WithTriggerControlData(TriggerControlData) SALDataTriggerControlBuilder
	// WithTriggerControlDataBuilder adds TriggerControlData (property field) which is build by the builder
	WithTriggerControlDataBuilder(func(TriggerControlDataBuilder) TriggerControlDataBuilder) SALDataTriggerControlBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() SALDataBuilder
	// Build builds the SALDataTriggerControl or returns an error if something is wrong
	Build() (SALDataTriggerControl, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SALDataTriggerControl
}

// NewSALDataTriggerControlBuilder() creates a SALDataTriggerControlBuilder
func NewSALDataTriggerControlBuilder() SALDataTriggerControlBuilder {
	return &_SALDataTriggerControlBuilder{_SALDataTriggerControl: new(_SALDataTriggerControl)}
}

type _SALDataTriggerControlBuilder struct {
	*_SALDataTriggerControl

	parentBuilder *_SALDataBuilder

	err *utils.MultiError
}

var _ (SALDataTriggerControlBuilder) = (*_SALDataTriggerControlBuilder)(nil)

func (b *_SALDataTriggerControlBuilder) setParent(contract SALDataContract) {
	b.SALDataContract = contract
	contract.(*_SALData)._SubType = b._SALDataTriggerControl
}

func (b *_SALDataTriggerControlBuilder) WithMandatoryFields(triggerControlData TriggerControlData) SALDataTriggerControlBuilder {
	return b.WithTriggerControlData(triggerControlData)
}

func (b *_SALDataTriggerControlBuilder) WithTriggerControlData(triggerControlData TriggerControlData) SALDataTriggerControlBuilder {
	b.TriggerControlData = triggerControlData
	return b
}

func (b *_SALDataTriggerControlBuilder) WithTriggerControlDataBuilder(builderSupplier func(TriggerControlDataBuilder) TriggerControlDataBuilder) SALDataTriggerControlBuilder {
	builder := builderSupplier(b.TriggerControlData.CreateTriggerControlDataBuilder())
	var err error
	b.TriggerControlData, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "TriggerControlDataBuilder failed"))
	}
	return b
}

func (b *_SALDataTriggerControlBuilder) Build() (SALDataTriggerControl, error) {
	if b.TriggerControlData == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'triggerControlData' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SALDataTriggerControl.deepCopy(), nil
}

func (b *_SALDataTriggerControlBuilder) MustBuild() SALDataTriggerControl {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SALDataTriggerControlBuilder) Done() SALDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewSALDataBuilder().(*_SALDataBuilder)
	}
	return b.parentBuilder
}

func (b *_SALDataTriggerControlBuilder) buildForSALData() (SALData, error) {
	return b.Build()
}

func (b *_SALDataTriggerControlBuilder) DeepCopy() any {
	_copy := b.CreateSALDataTriggerControlBuilder().(*_SALDataTriggerControlBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSALDataTriggerControlBuilder creates a SALDataTriggerControlBuilder
func (b *_SALDataTriggerControl) CreateSALDataTriggerControlBuilder() SALDataTriggerControlBuilder {
	if b == nil {
		return NewSALDataTriggerControlBuilder()
	}
	return &_SALDataTriggerControlBuilder{_SALDataTriggerControl: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataTriggerControl) GetApplicationId() ApplicationId {
	return ApplicationId_TRIGGER_CONTROL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataTriggerControl) GetParent() SALDataContract {
	return m.SALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataTriggerControl) GetTriggerControlData() TriggerControlData {
	return m.TriggerControlData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSALDataTriggerControl(structType any) SALDataTriggerControl {
	if casted, ok := structType.(SALDataTriggerControl); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataTriggerControl); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataTriggerControl) GetTypeName() string {
	return "SALDataTriggerControl"
}

func (m *_SALDataTriggerControl) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	// Simple field (triggerControlData)
	lengthInBits += m.TriggerControlData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataTriggerControl) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataTriggerControl) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataTriggerControl SALDataTriggerControl, err error) {
	m.SALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataTriggerControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataTriggerControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	triggerControlData, err := ReadSimpleField[TriggerControlData](ctx, "triggerControlData", ReadComplex[TriggerControlData](TriggerControlDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'triggerControlData' field"))
	}
	m.TriggerControlData = triggerControlData

	if closeErr := readBuffer.CloseContext("SALDataTriggerControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataTriggerControl")
	}

	return m, nil
}

func (m *_SALDataTriggerControl) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataTriggerControl) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataTriggerControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataTriggerControl")
		}

		if err := WriteSimpleField[TriggerControlData](ctx, "triggerControlData", m.GetTriggerControlData(), WriteComplex[TriggerControlData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'triggerControlData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataTriggerControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataTriggerControl")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataTriggerControl) IsSALDataTriggerControl() {}

func (m *_SALDataTriggerControl) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SALDataTriggerControl) deepCopy() *_SALDataTriggerControl {
	if m == nil {
		return nil
	}
	_SALDataTriggerControlCopy := &_SALDataTriggerControl{
		m.SALDataContract.(*_SALData).deepCopy(),
		utils.DeepCopy[TriggerControlData](m.TriggerControlData),
	}
	_SALDataTriggerControlCopy.SALDataContract.(*_SALData)._SubType = m
	return _SALDataTriggerControlCopy
}

func (m *_SALDataTriggerControl) String() string {
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