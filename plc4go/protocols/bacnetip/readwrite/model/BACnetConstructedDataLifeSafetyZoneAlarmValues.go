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

// BACnetConstructedDataLifeSafetyZoneAlarmValues is the corresponding interface of BACnetConstructedDataLifeSafetyZoneAlarmValues
type BACnetConstructedDataLifeSafetyZoneAlarmValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetAlarmValues returns AlarmValues (property field)
	GetAlarmValues() []BACnetLifeSafetyStateTagged
	// IsBACnetConstructedDataLifeSafetyZoneAlarmValues is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLifeSafetyZoneAlarmValues()
	// CreateBuilder creates a BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder
	CreateBACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder() BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder
}

// _BACnetConstructedDataLifeSafetyZoneAlarmValues is the data-structure of this message
type _BACnetConstructedDataLifeSafetyZoneAlarmValues struct {
	BACnetConstructedDataContract
	AlarmValues []BACnetLifeSafetyStateTagged
}

var _ BACnetConstructedDataLifeSafetyZoneAlarmValues = (*_BACnetConstructedDataLifeSafetyZoneAlarmValues)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLifeSafetyZoneAlarmValues)(nil)

// NewBACnetConstructedDataLifeSafetyZoneAlarmValues factory function for _BACnetConstructedDataLifeSafetyZoneAlarmValues
func NewBACnetConstructedDataLifeSafetyZoneAlarmValues(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, alarmValues []BACnetLifeSafetyStateTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLifeSafetyZoneAlarmValues {
	_result := &_BACnetConstructedDataLifeSafetyZoneAlarmValues{
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

// BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder is a builder for BACnetConstructedDataLifeSafetyZoneAlarmValues
type BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(alarmValues []BACnetLifeSafetyStateTagged) BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder
	// WithAlarmValues adds AlarmValues (property field)
	WithAlarmValues(...BACnetLifeSafetyStateTagged) BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLifeSafetyZoneAlarmValues or returns an error if something is wrong
	Build() (BACnetConstructedDataLifeSafetyZoneAlarmValues, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLifeSafetyZoneAlarmValues
}

// NewBACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder() creates a BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder
func NewBACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder() BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder {
	return &_BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder{_BACnetConstructedDataLifeSafetyZoneAlarmValues: new(_BACnetConstructedDataLifeSafetyZoneAlarmValues)}
}

type _BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder struct {
	*_BACnetConstructedDataLifeSafetyZoneAlarmValues

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder) = (*_BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder)(nil)

func (b *_BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataLifeSafetyZoneAlarmValues
}

func (b *_BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder) WithMandatoryFields(alarmValues []BACnetLifeSafetyStateTagged) BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder {
	return b.WithAlarmValues(alarmValues...)
}

func (b *_BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder) WithAlarmValues(alarmValues ...BACnetLifeSafetyStateTagged) BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder {
	b.AlarmValues = alarmValues
	return b
}

func (b *_BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder) Build() (BACnetConstructedDataLifeSafetyZoneAlarmValues, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLifeSafetyZoneAlarmValues.deepCopy(), nil
}

func (b *_BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder) MustBuild() BACnetConstructedDataLifeSafetyZoneAlarmValues {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder().(*_BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder creates a BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder
func (b *_BACnetConstructedDataLifeSafetyZoneAlarmValues) CreateBACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder() BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder {
	if b == nil {
		return NewBACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder()
	}
	return &_BACnetConstructedDataLifeSafetyZoneAlarmValuesBuilder{_BACnetConstructedDataLifeSafetyZoneAlarmValues: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyZoneAlarmValues) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LIFE_SAFETY_ZONE
}

func (m *_BACnetConstructedDataLifeSafetyZoneAlarmValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALARM_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLifeSafetyZoneAlarmValues) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyZoneAlarmValues) GetAlarmValues() []BACnetLifeSafetyStateTagged {
	return m.AlarmValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLifeSafetyZoneAlarmValues(structType any) BACnetConstructedDataLifeSafetyZoneAlarmValues {
	if casted, ok := structType.(BACnetConstructedDataLifeSafetyZoneAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLifeSafetyZoneAlarmValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLifeSafetyZoneAlarmValues) GetTypeName() string {
	return "BACnetConstructedDataLifeSafetyZoneAlarmValues"
}

func (m *_BACnetConstructedDataLifeSafetyZoneAlarmValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.AlarmValues) > 0 {
		for _, element := range m.AlarmValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataLifeSafetyZoneAlarmValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLifeSafetyZoneAlarmValues) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLifeSafetyZoneAlarmValues BACnetConstructedDataLifeSafetyZoneAlarmValues, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLifeSafetyZoneAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLifeSafetyZoneAlarmValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	alarmValues, err := ReadTerminatedArrayField[BACnetLifeSafetyStateTagged](ctx, "alarmValues", ReadComplex[BACnetLifeSafetyStateTagged](BACnetLifeSafetyStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alarmValues' field"))
	}
	m.AlarmValues = alarmValues

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLifeSafetyZoneAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLifeSafetyZoneAlarmValues")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLifeSafetyZoneAlarmValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLifeSafetyZoneAlarmValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLifeSafetyZoneAlarmValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLifeSafetyZoneAlarmValues")
		}

		if err := WriteComplexTypeArrayField(ctx, "alarmValues", m.GetAlarmValues(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'alarmValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLifeSafetyZoneAlarmValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLifeSafetyZoneAlarmValues")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLifeSafetyZoneAlarmValues) IsBACnetConstructedDataLifeSafetyZoneAlarmValues() {
}

func (m *_BACnetConstructedDataLifeSafetyZoneAlarmValues) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLifeSafetyZoneAlarmValues) deepCopy() *_BACnetConstructedDataLifeSafetyZoneAlarmValues {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLifeSafetyZoneAlarmValuesCopy := &_BACnetConstructedDataLifeSafetyZoneAlarmValues{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetLifeSafetyStateTagged, BACnetLifeSafetyStateTagged](m.AlarmValues),
	}
	_BACnetConstructedDataLifeSafetyZoneAlarmValuesCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLifeSafetyZoneAlarmValuesCopy
}

func (m *_BACnetConstructedDataLifeSafetyZoneAlarmValues) String() string {
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