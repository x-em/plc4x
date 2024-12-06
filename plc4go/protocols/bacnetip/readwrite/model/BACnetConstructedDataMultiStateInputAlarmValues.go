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

// BACnetConstructedDataMultiStateInputAlarmValues is the corresponding interface of BACnetConstructedDataMultiStateInputAlarmValues
type BACnetConstructedDataMultiStateInputAlarmValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetAlarmValues returns AlarmValues (property field)
	GetAlarmValues() []BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataMultiStateInputAlarmValues is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMultiStateInputAlarmValues()
	// CreateBuilder creates a BACnetConstructedDataMultiStateInputAlarmValuesBuilder
	CreateBACnetConstructedDataMultiStateInputAlarmValuesBuilder() BACnetConstructedDataMultiStateInputAlarmValuesBuilder
}

// _BACnetConstructedDataMultiStateInputAlarmValues is the data-structure of this message
type _BACnetConstructedDataMultiStateInputAlarmValues struct {
	BACnetConstructedDataContract
	AlarmValues []BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataMultiStateInputAlarmValues = (*_BACnetConstructedDataMultiStateInputAlarmValues)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMultiStateInputAlarmValues)(nil)

// NewBACnetConstructedDataMultiStateInputAlarmValues factory function for _BACnetConstructedDataMultiStateInputAlarmValues
func NewBACnetConstructedDataMultiStateInputAlarmValues(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, alarmValues []BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMultiStateInputAlarmValues {
	_result := &_BACnetConstructedDataMultiStateInputAlarmValues{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AlarmValues:                   alarmValues,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataMultiStateInputAlarmValuesBuilder is a builder for BACnetConstructedDataMultiStateInputAlarmValues
type BACnetConstructedDataMultiStateInputAlarmValuesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(alarmValues []BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMultiStateInputAlarmValuesBuilder
	// WithAlarmValues adds AlarmValues (property field)
	WithAlarmValues(...BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMultiStateInputAlarmValuesBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataMultiStateInputAlarmValues or returns an error if something is wrong
	Build() (BACnetConstructedDataMultiStateInputAlarmValues, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataMultiStateInputAlarmValues
}

// NewBACnetConstructedDataMultiStateInputAlarmValuesBuilder() creates a BACnetConstructedDataMultiStateInputAlarmValuesBuilder
func NewBACnetConstructedDataMultiStateInputAlarmValuesBuilder() BACnetConstructedDataMultiStateInputAlarmValuesBuilder {
	return &_BACnetConstructedDataMultiStateInputAlarmValuesBuilder{_BACnetConstructedDataMultiStateInputAlarmValues: new(_BACnetConstructedDataMultiStateInputAlarmValues)}
}

type _BACnetConstructedDataMultiStateInputAlarmValuesBuilder struct {
	*_BACnetConstructedDataMultiStateInputAlarmValues

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataMultiStateInputAlarmValuesBuilder) = (*_BACnetConstructedDataMultiStateInputAlarmValuesBuilder)(nil)

func (b *_BACnetConstructedDataMultiStateInputAlarmValuesBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataMultiStateInputAlarmValues
}

func (b *_BACnetConstructedDataMultiStateInputAlarmValuesBuilder) WithMandatoryFields(alarmValues []BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMultiStateInputAlarmValuesBuilder {
	return b.WithAlarmValues(alarmValues...)
}

func (b *_BACnetConstructedDataMultiStateInputAlarmValuesBuilder) WithAlarmValues(alarmValues ...BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMultiStateInputAlarmValuesBuilder {
	b.AlarmValues = alarmValues
	return b
}

func (b *_BACnetConstructedDataMultiStateInputAlarmValuesBuilder) Build() (BACnetConstructedDataMultiStateInputAlarmValues, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataMultiStateInputAlarmValues.deepCopy(), nil
}

func (b *_BACnetConstructedDataMultiStateInputAlarmValuesBuilder) MustBuild() BACnetConstructedDataMultiStateInputAlarmValues {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataMultiStateInputAlarmValuesBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataMultiStateInputAlarmValuesBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataMultiStateInputAlarmValuesBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataMultiStateInputAlarmValuesBuilder().(*_BACnetConstructedDataMultiStateInputAlarmValuesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataMultiStateInputAlarmValuesBuilder creates a BACnetConstructedDataMultiStateInputAlarmValuesBuilder
func (b *_BACnetConstructedDataMultiStateInputAlarmValues) CreateBACnetConstructedDataMultiStateInputAlarmValuesBuilder() BACnetConstructedDataMultiStateInputAlarmValuesBuilder {
	if b == nil {
		return NewBACnetConstructedDataMultiStateInputAlarmValuesBuilder()
	}
	return &_BACnetConstructedDataMultiStateInputAlarmValuesBuilder{_BACnetConstructedDataMultiStateInputAlarmValues: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMultiStateInputAlarmValues) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_MULTI_STATE_INPUT
}

func (m *_BACnetConstructedDataMultiStateInputAlarmValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALARM_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMultiStateInputAlarmValues) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMultiStateInputAlarmValues) GetAlarmValues() []BACnetApplicationTagUnsignedInteger {
	return m.AlarmValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMultiStateInputAlarmValues(structType any) BACnetConstructedDataMultiStateInputAlarmValues {
	if casted, ok := structType.(BACnetConstructedDataMultiStateInputAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMultiStateInputAlarmValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMultiStateInputAlarmValues) GetTypeName() string {
	return "BACnetConstructedDataMultiStateInputAlarmValues"
}

func (m *_BACnetConstructedDataMultiStateInputAlarmValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.AlarmValues) > 0 {
		for _, element := range m.AlarmValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataMultiStateInputAlarmValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMultiStateInputAlarmValues) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMultiStateInputAlarmValues BACnetConstructedDataMultiStateInputAlarmValues, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMultiStateInputAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMultiStateInputAlarmValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	alarmValues, err := ReadTerminatedArrayField[BACnetApplicationTagUnsignedInteger](ctx, "alarmValues", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alarmValues' field"))
	}
	m.AlarmValues = alarmValues

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMultiStateInputAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMultiStateInputAlarmValues")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMultiStateInputAlarmValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMultiStateInputAlarmValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMultiStateInputAlarmValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMultiStateInputAlarmValues")
		}

		if err := WriteComplexTypeArrayField(ctx, "alarmValues", m.GetAlarmValues(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'alarmValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMultiStateInputAlarmValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMultiStateInputAlarmValues")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMultiStateInputAlarmValues) IsBACnetConstructedDataMultiStateInputAlarmValues() {
}

func (m *_BACnetConstructedDataMultiStateInputAlarmValues) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataMultiStateInputAlarmValues) deepCopy() *_BACnetConstructedDataMultiStateInputAlarmValues {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataMultiStateInputAlarmValuesCopy := &_BACnetConstructedDataMultiStateInputAlarmValues{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetApplicationTagUnsignedInteger, BACnetApplicationTagUnsignedInteger](m.AlarmValues),
	}
	_BACnetConstructedDataMultiStateInputAlarmValuesCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataMultiStateInputAlarmValuesCopy
}

func (m *_BACnetConstructedDataMultiStateInputAlarmValues) String() string {
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