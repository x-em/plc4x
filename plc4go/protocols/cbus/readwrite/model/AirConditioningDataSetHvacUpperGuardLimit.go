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

// AirConditioningDataSetHvacUpperGuardLimit is the corresponding interface of AirConditioningDataSetHvacUpperGuardLimit
type AirConditioningDataSetHvacUpperGuardLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetLimit returns Limit (property field)
	GetLimit() HVACTemperature
	// GetHvacModeAndFlags returns HvacModeAndFlags (property field)
	GetHvacModeAndFlags() HVACModeAndFlags
	// IsAirConditioningDataSetHvacUpperGuardLimit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAirConditioningDataSetHvacUpperGuardLimit()
	// CreateBuilder creates a AirConditioningDataSetHvacUpperGuardLimitBuilder
	CreateAirConditioningDataSetHvacUpperGuardLimitBuilder() AirConditioningDataSetHvacUpperGuardLimitBuilder
}

// _AirConditioningDataSetHvacUpperGuardLimit is the data-structure of this message
type _AirConditioningDataSetHvacUpperGuardLimit struct {
	AirConditioningDataContract
	ZoneGroup        byte
	ZoneList         HVACZoneList
	Limit            HVACTemperature
	HvacModeAndFlags HVACModeAndFlags
}

var _ AirConditioningDataSetHvacUpperGuardLimit = (*_AirConditioningDataSetHvacUpperGuardLimit)(nil)
var _ AirConditioningDataRequirements = (*_AirConditioningDataSetHvacUpperGuardLimit)(nil)

// NewAirConditioningDataSetHvacUpperGuardLimit factory function for _AirConditioningDataSetHvacUpperGuardLimit
func NewAirConditioningDataSetHvacUpperGuardLimit(commandTypeContainer AirConditioningCommandTypeContainer, zoneGroup byte, zoneList HVACZoneList, limit HVACTemperature, hvacModeAndFlags HVACModeAndFlags) *_AirConditioningDataSetHvacUpperGuardLimit {
	if zoneList == nil {
		panic("zoneList of type HVACZoneList for AirConditioningDataSetHvacUpperGuardLimit must not be nil")
	}
	if limit == nil {
		panic("limit of type HVACTemperature for AirConditioningDataSetHvacUpperGuardLimit must not be nil")
	}
	if hvacModeAndFlags == nil {
		panic("hvacModeAndFlags of type HVACModeAndFlags for AirConditioningDataSetHvacUpperGuardLimit must not be nil")
	}
	_result := &_AirConditioningDataSetHvacUpperGuardLimit{
		AirConditioningDataContract: NewAirConditioningData(commandTypeContainer),
		ZoneGroup:                   zoneGroup,
		ZoneList:                    zoneList,
		Limit:                       limit,
		HvacModeAndFlags:            hvacModeAndFlags,
	}
	_result.AirConditioningDataContract.(*_AirConditioningData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AirConditioningDataSetHvacUpperGuardLimitBuilder is a builder for AirConditioningDataSetHvacUpperGuardLimit
type AirConditioningDataSetHvacUpperGuardLimitBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(zoneGroup byte, zoneList HVACZoneList, limit HVACTemperature, hvacModeAndFlags HVACModeAndFlags) AirConditioningDataSetHvacUpperGuardLimitBuilder
	// WithZoneGroup adds ZoneGroup (property field)
	WithZoneGroup(byte) AirConditioningDataSetHvacUpperGuardLimitBuilder
	// WithZoneList adds ZoneList (property field)
	WithZoneList(HVACZoneList) AirConditioningDataSetHvacUpperGuardLimitBuilder
	// WithZoneListBuilder adds ZoneList (property field) which is build by the builder
	WithZoneListBuilder(func(HVACZoneListBuilder) HVACZoneListBuilder) AirConditioningDataSetHvacUpperGuardLimitBuilder
	// WithLimit adds Limit (property field)
	WithLimit(HVACTemperature) AirConditioningDataSetHvacUpperGuardLimitBuilder
	// WithLimitBuilder adds Limit (property field) which is build by the builder
	WithLimitBuilder(func(HVACTemperatureBuilder) HVACTemperatureBuilder) AirConditioningDataSetHvacUpperGuardLimitBuilder
	// WithHvacModeAndFlags adds HvacModeAndFlags (property field)
	WithHvacModeAndFlags(HVACModeAndFlags) AirConditioningDataSetHvacUpperGuardLimitBuilder
	// WithHvacModeAndFlagsBuilder adds HvacModeAndFlags (property field) which is build by the builder
	WithHvacModeAndFlagsBuilder(func(HVACModeAndFlagsBuilder) HVACModeAndFlagsBuilder) AirConditioningDataSetHvacUpperGuardLimitBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() AirConditioningDataBuilder
	// Build builds the AirConditioningDataSetHvacUpperGuardLimit or returns an error if something is wrong
	Build() (AirConditioningDataSetHvacUpperGuardLimit, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AirConditioningDataSetHvacUpperGuardLimit
}

// NewAirConditioningDataSetHvacUpperGuardLimitBuilder() creates a AirConditioningDataSetHvacUpperGuardLimitBuilder
func NewAirConditioningDataSetHvacUpperGuardLimitBuilder() AirConditioningDataSetHvacUpperGuardLimitBuilder {
	return &_AirConditioningDataSetHvacUpperGuardLimitBuilder{_AirConditioningDataSetHvacUpperGuardLimit: new(_AirConditioningDataSetHvacUpperGuardLimit)}
}

type _AirConditioningDataSetHvacUpperGuardLimitBuilder struct {
	*_AirConditioningDataSetHvacUpperGuardLimit

	parentBuilder *_AirConditioningDataBuilder

	err *utils.MultiError
}

var _ (AirConditioningDataSetHvacUpperGuardLimitBuilder) = (*_AirConditioningDataSetHvacUpperGuardLimitBuilder)(nil)

func (b *_AirConditioningDataSetHvacUpperGuardLimitBuilder) setParent(contract AirConditioningDataContract) {
	b.AirConditioningDataContract = contract
	contract.(*_AirConditioningData)._SubType = b._AirConditioningDataSetHvacUpperGuardLimit
}

func (b *_AirConditioningDataSetHvacUpperGuardLimitBuilder) WithMandatoryFields(zoneGroup byte, zoneList HVACZoneList, limit HVACTemperature, hvacModeAndFlags HVACModeAndFlags) AirConditioningDataSetHvacUpperGuardLimitBuilder {
	return b.WithZoneGroup(zoneGroup).WithZoneList(zoneList).WithLimit(limit).WithHvacModeAndFlags(hvacModeAndFlags)
}

func (b *_AirConditioningDataSetHvacUpperGuardLimitBuilder) WithZoneGroup(zoneGroup byte) AirConditioningDataSetHvacUpperGuardLimitBuilder {
	b.ZoneGroup = zoneGroup
	return b
}

func (b *_AirConditioningDataSetHvacUpperGuardLimitBuilder) WithZoneList(zoneList HVACZoneList) AirConditioningDataSetHvacUpperGuardLimitBuilder {
	b.ZoneList = zoneList
	return b
}

func (b *_AirConditioningDataSetHvacUpperGuardLimitBuilder) WithZoneListBuilder(builderSupplier func(HVACZoneListBuilder) HVACZoneListBuilder) AirConditioningDataSetHvacUpperGuardLimitBuilder {
	builder := builderSupplier(b.ZoneList.CreateHVACZoneListBuilder())
	var err error
	b.ZoneList, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACZoneListBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataSetHvacUpperGuardLimitBuilder) WithLimit(limit HVACTemperature) AirConditioningDataSetHvacUpperGuardLimitBuilder {
	b.Limit = limit
	return b
}

func (b *_AirConditioningDataSetHvacUpperGuardLimitBuilder) WithLimitBuilder(builderSupplier func(HVACTemperatureBuilder) HVACTemperatureBuilder) AirConditioningDataSetHvacUpperGuardLimitBuilder {
	builder := builderSupplier(b.Limit.CreateHVACTemperatureBuilder())
	var err error
	b.Limit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACTemperatureBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataSetHvacUpperGuardLimitBuilder) WithHvacModeAndFlags(hvacModeAndFlags HVACModeAndFlags) AirConditioningDataSetHvacUpperGuardLimitBuilder {
	b.HvacModeAndFlags = hvacModeAndFlags
	return b
}

func (b *_AirConditioningDataSetHvacUpperGuardLimitBuilder) WithHvacModeAndFlagsBuilder(builderSupplier func(HVACModeAndFlagsBuilder) HVACModeAndFlagsBuilder) AirConditioningDataSetHvacUpperGuardLimitBuilder {
	builder := builderSupplier(b.HvacModeAndFlags.CreateHVACModeAndFlagsBuilder())
	var err error
	b.HvacModeAndFlags, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACModeAndFlagsBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataSetHvacUpperGuardLimitBuilder) Build() (AirConditioningDataSetHvacUpperGuardLimit, error) {
	if b.ZoneList == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'zoneList' not set"))
	}
	if b.Limit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'limit' not set"))
	}
	if b.HvacModeAndFlags == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'hvacModeAndFlags' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AirConditioningDataSetHvacUpperGuardLimit.deepCopy(), nil
}

func (b *_AirConditioningDataSetHvacUpperGuardLimitBuilder) MustBuild() AirConditioningDataSetHvacUpperGuardLimit {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AirConditioningDataSetHvacUpperGuardLimitBuilder) Done() AirConditioningDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAirConditioningDataBuilder().(*_AirConditioningDataBuilder)
	}
	return b.parentBuilder
}

func (b *_AirConditioningDataSetHvacUpperGuardLimitBuilder) buildForAirConditioningData() (AirConditioningData, error) {
	return b.Build()
}

func (b *_AirConditioningDataSetHvacUpperGuardLimitBuilder) DeepCopy() any {
	_copy := b.CreateAirConditioningDataSetHvacUpperGuardLimitBuilder().(*_AirConditioningDataSetHvacUpperGuardLimitBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAirConditioningDataSetHvacUpperGuardLimitBuilder creates a AirConditioningDataSetHvacUpperGuardLimitBuilder
func (b *_AirConditioningDataSetHvacUpperGuardLimit) CreateAirConditioningDataSetHvacUpperGuardLimitBuilder() AirConditioningDataSetHvacUpperGuardLimitBuilder {
	if b == nil {
		return NewAirConditioningDataSetHvacUpperGuardLimitBuilder()
	}
	return &_AirConditioningDataSetHvacUpperGuardLimitBuilder{_AirConditioningDataSetHvacUpperGuardLimit: b.deepCopy()}
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

func (m *_AirConditioningDataSetHvacUpperGuardLimit) GetParent() AirConditioningDataContract {
	return m.AirConditioningDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataSetHvacUpperGuardLimit) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataSetHvacUpperGuardLimit) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataSetHvacUpperGuardLimit) GetLimit() HVACTemperature {
	return m.Limit
}

func (m *_AirConditioningDataSetHvacUpperGuardLimit) GetHvacModeAndFlags() HVACModeAndFlags {
	return m.HvacModeAndFlags
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAirConditioningDataSetHvacUpperGuardLimit(structType any) AirConditioningDataSetHvacUpperGuardLimit {
	if casted, ok := structType.(AirConditioningDataSetHvacUpperGuardLimit); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataSetHvacUpperGuardLimit); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataSetHvacUpperGuardLimit) GetTypeName() string {
	return "AirConditioningDataSetHvacUpperGuardLimit"
}

func (m *_AirConditioningDataSetHvacUpperGuardLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AirConditioningDataContract.(*_AirConditioningData).getLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (limit)
	lengthInBits += m.Limit.GetLengthInBits(ctx)

	// Simple field (hvacModeAndFlags)
	lengthInBits += m.HvacModeAndFlags.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AirConditioningDataSetHvacUpperGuardLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AirConditioningDataSetHvacUpperGuardLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AirConditioningData) (__airConditioningDataSetHvacUpperGuardLimit AirConditioningDataSetHvacUpperGuardLimit, err error) {
	m.AirConditioningDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataSetHvacUpperGuardLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataSetHvacUpperGuardLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneGroup, err := ReadSimpleField(ctx, "zoneGroup", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneGroup' field"))
	}
	m.ZoneGroup = zoneGroup

	zoneList, err := ReadSimpleField[HVACZoneList](ctx, "zoneList", ReadComplex[HVACZoneList](HVACZoneListParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneList' field"))
	}
	m.ZoneList = zoneList

	limit, err := ReadSimpleField[HVACTemperature](ctx, "limit", ReadComplex[HVACTemperature](HVACTemperatureParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'limit' field"))
	}
	m.Limit = limit

	hvacModeAndFlags, err := ReadSimpleField[HVACModeAndFlags](ctx, "hvacModeAndFlags", ReadComplex[HVACModeAndFlags](HVACModeAndFlagsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hvacModeAndFlags' field"))
	}
	m.HvacModeAndFlags = hvacModeAndFlags

	if closeErr := readBuffer.CloseContext("AirConditioningDataSetHvacUpperGuardLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataSetHvacUpperGuardLimit")
	}

	return m, nil
}

func (m *_AirConditioningDataSetHvacUpperGuardLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataSetHvacUpperGuardLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataSetHvacUpperGuardLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataSetHvacUpperGuardLimit")
		}

		if err := WriteSimpleField[byte](ctx, "zoneGroup", m.GetZoneGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneGroup' field")
		}

		if err := WriteSimpleField[HVACZoneList](ctx, "zoneList", m.GetZoneList(), WriteComplex[HVACZoneList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneList' field")
		}

		if err := WriteSimpleField[HVACTemperature](ctx, "limit", m.GetLimit(), WriteComplex[HVACTemperature](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'limit' field")
		}

		if err := WriteSimpleField[HVACModeAndFlags](ctx, "hvacModeAndFlags", m.GetHvacModeAndFlags(), WriteComplex[HVACModeAndFlags](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'hvacModeAndFlags' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataSetHvacUpperGuardLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataSetHvacUpperGuardLimit")
		}
		return nil
	}
	return m.AirConditioningDataContract.(*_AirConditioningData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataSetHvacUpperGuardLimit) IsAirConditioningDataSetHvacUpperGuardLimit() {}

func (m *_AirConditioningDataSetHvacUpperGuardLimit) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AirConditioningDataSetHvacUpperGuardLimit) deepCopy() *_AirConditioningDataSetHvacUpperGuardLimit {
	if m == nil {
		return nil
	}
	_AirConditioningDataSetHvacUpperGuardLimitCopy := &_AirConditioningDataSetHvacUpperGuardLimit{
		m.AirConditioningDataContract.(*_AirConditioningData).deepCopy(),
		m.ZoneGroup,
		utils.DeepCopy[HVACZoneList](m.ZoneList),
		utils.DeepCopy[HVACTemperature](m.Limit),
		utils.DeepCopy[HVACModeAndFlags](m.HvacModeAndFlags),
	}
	_AirConditioningDataSetHvacUpperGuardLimitCopy.AirConditioningDataContract.(*_AirConditioningData)._SubType = m
	return _AirConditioningDataSetHvacUpperGuardLimitCopy
}

func (m *_AirConditioningDataSetHvacUpperGuardLimit) String() string {
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