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

// AirConditioningDataSetPlantHumidityLevel is the corresponding interface of AirConditioningDataSetPlantHumidityLevel
type AirConditioningDataSetPlantHumidityLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetHumidityModeAndFlags returns HumidityModeAndFlags (property field)
	GetHumidityModeAndFlags() HVACHumidityModeAndFlags
	// GetHumidityType returns HumidityType (property field)
	GetHumidityType() HVACHumidityType
	// GetLevel returns Level (property field)
	GetLevel() HVACHumidity
	// GetRawLevel returns RawLevel (property field)
	GetRawLevel() HVACRawLevels
	// GetAuxLevel returns AuxLevel (property field)
	GetAuxLevel() HVACAuxiliaryLevel
	// IsAirConditioningDataSetPlantHumidityLevel is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAirConditioningDataSetPlantHumidityLevel()
	// CreateBuilder creates a AirConditioningDataSetPlantHumidityLevelBuilder
	CreateAirConditioningDataSetPlantHumidityLevelBuilder() AirConditioningDataSetPlantHumidityLevelBuilder
}

// _AirConditioningDataSetPlantHumidityLevel is the data-structure of this message
type _AirConditioningDataSetPlantHumidityLevel struct {
	AirConditioningDataContract
	ZoneGroup            byte
	ZoneList             HVACZoneList
	HumidityModeAndFlags HVACHumidityModeAndFlags
	HumidityType         HVACHumidityType
	Level                HVACHumidity
	RawLevel             HVACRawLevels
	AuxLevel             HVACAuxiliaryLevel
}

var _ AirConditioningDataSetPlantHumidityLevel = (*_AirConditioningDataSetPlantHumidityLevel)(nil)
var _ AirConditioningDataRequirements = (*_AirConditioningDataSetPlantHumidityLevel)(nil)

// NewAirConditioningDataSetPlantHumidityLevel factory function for _AirConditioningDataSetPlantHumidityLevel
func NewAirConditioningDataSetPlantHumidityLevel(commandTypeContainer AirConditioningCommandTypeContainer, zoneGroup byte, zoneList HVACZoneList, humidityModeAndFlags HVACHumidityModeAndFlags, humidityType HVACHumidityType, level HVACHumidity, rawLevel HVACRawLevels, auxLevel HVACAuxiliaryLevel) *_AirConditioningDataSetPlantHumidityLevel {
	if zoneList == nil {
		panic("zoneList of type HVACZoneList for AirConditioningDataSetPlantHumidityLevel must not be nil")
	}
	if humidityModeAndFlags == nil {
		panic("humidityModeAndFlags of type HVACHumidityModeAndFlags for AirConditioningDataSetPlantHumidityLevel must not be nil")
	}
	_result := &_AirConditioningDataSetPlantHumidityLevel{
		AirConditioningDataContract: NewAirConditioningData(commandTypeContainer),
		ZoneGroup:                   zoneGroup,
		ZoneList:                    zoneList,
		HumidityModeAndFlags:        humidityModeAndFlags,
		HumidityType:                humidityType,
		Level:                       level,
		RawLevel:                    rawLevel,
		AuxLevel:                    auxLevel,
	}
	_result.AirConditioningDataContract.(*_AirConditioningData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AirConditioningDataSetPlantHumidityLevelBuilder is a builder for AirConditioningDataSetPlantHumidityLevel
type AirConditioningDataSetPlantHumidityLevelBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(zoneGroup byte, zoneList HVACZoneList, humidityModeAndFlags HVACHumidityModeAndFlags, humidityType HVACHumidityType) AirConditioningDataSetPlantHumidityLevelBuilder
	// WithZoneGroup adds ZoneGroup (property field)
	WithZoneGroup(byte) AirConditioningDataSetPlantHumidityLevelBuilder
	// WithZoneList adds ZoneList (property field)
	WithZoneList(HVACZoneList) AirConditioningDataSetPlantHumidityLevelBuilder
	// WithZoneListBuilder adds ZoneList (property field) which is build by the builder
	WithZoneListBuilder(func(HVACZoneListBuilder) HVACZoneListBuilder) AirConditioningDataSetPlantHumidityLevelBuilder
	// WithHumidityModeAndFlags adds HumidityModeAndFlags (property field)
	WithHumidityModeAndFlags(HVACHumidityModeAndFlags) AirConditioningDataSetPlantHumidityLevelBuilder
	// WithHumidityModeAndFlagsBuilder adds HumidityModeAndFlags (property field) which is build by the builder
	WithHumidityModeAndFlagsBuilder(func(HVACHumidityModeAndFlagsBuilder) HVACHumidityModeAndFlagsBuilder) AirConditioningDataSetPlantHumidityLevelBuilder
	// WithHumidityType adds HumidityType (property field)
	WithHumidityType(HVACHumidityType) AirConditioningDataSetPlantHumidityLevelBuilder
	// WithLevel adds Level (property field)
	WithOptionalLevel(HVACHumidity) AirConditioningDataSetPlantHumidityLevelBuilder
	// WithOptionalLevelBuilder adds Level (property field) which is build by the builder
	WithOptionalLevelBuilder(func(HVACHumidityBuilder) HVACHumidityBuilder) AirConditioningDataSetPlantHumidityLevelBuilder
	// WithRawLevel adds RawLevel (property field)
	WithOptionalRawLevel(HVACRawLevels) AirConditioningDataSetPlantHumidityLevelBuilder
	// WithOptionalRawLevelBuilder adds RawLevel (property field) which is build by the builder
	WithOptionalRawLevelBuilder(func(HVACRawLevelsBuilder) HVACRawLevelsBuilder) AirConditioningDataSetPlantHumidityLevelBuilder
	// WithAuxLevel adds AuxLevel (property field)
	WithOptionalAuxLevel(HVACAuxiliaryLevel) AirConditioningDataSetPlantHumidityLevelBuilder
	// WithOptionalAuxLevelBuilder adds AuxLevel (property field) which is build by the builder
	WithOptionalAuxLevelBuilder(func(HVACAuxiliaryLevelBuilder) HVACAuxiliaryLevelBuilder) AirConditioningDataSetPlantHumidityLevelBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() AirConditioningDataBuilder
	// Build builds the AirConditioningDataSetPlantHumidityLevel or returns an error if something is wrong
	Build() (AirConditioningDataSetPlantHumidityLevel, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AirConditioningDataSetPlantHumidityLevel
}

// NewAirConditioningDataSetPlantHumidityLevelBuilder() creates a AirConditioningDataSetPlantHumidityLevelBuilder
func NewAirConditioningDataSetPlantHumidityLevelBuilder() AirConditioningDataSetPlantHumidityLevelBuilder {
	return &_AirConditioningDataSetPlantHumidityLevelBuilder{_AirConditioningDataSetPlantHumidityLevel: new(_AirConditioningDataSetPlantHumidityLevel)}
}

type _AirConditioningDataSetPlantHumidityLevelBuilder struct {
	*_AirConditioningDataSetPlantHumidityLevel

	parentBuilder *_AirConditioningDataBuilder

	err *utils.MultiError
}

var _ (AirConditioningDataSetPlantHumidityLevelBuilder) = (*_AirConditioningDataSetPlantHumidityLevelBuilder)(nil)

func (b *_AirConditioningDataSetPlantHumidityLevelBuilder) setParent(contract AirConditioningDataContract) {
	b.AirConditioningDataContract = contract
	contract.(*_AirConditioningData)._SubType = b._AirConditioningDataSetPlantHumidityLevel
}

func (b *_AirConditioningDataSetPlantHumidityLevelBuilder) WithMandatoryFields(zoneGroup byte, zoneList HVACZoneList, humidityModeAndFlags HVACHumidityModeAndFlags, humidityType HVACHumidityType) AirConditioningDataSetPlantHumidityLevelBuilder {
	return b.WithZoneGroup(zoneGroup).WithZoneList(zoneList).WithHumidityModeAndFlags(humidityModeAndFlags).WithHumidityType(humidityType)
}

func (b *_AirConditioningDataSetPlantHumidityLevelBuilder) WithZoneGroup(zoneGroup byte) AirConditioningDataSetPlantHumidityLevelBuilder {
	b.ZoneGroup = zoneGroup
	return b
}

func (b *_AirConditioningDataSetPlantHumidityLevelBuilder) WithZoneList(zoneList HVACZoneList) AirConditioningDataSetPlantHumidityLevelBuilder {
	b.ZoneList = zoneList
	return b
}

func (b *_AirConditioningDataSetPlantHumidityLevelBuilder) WithZoneListBuilder(builderSupplier func(HVACZoneListBuilder) HVACZoneListBuilder) AirConditioningDataSetPlantHumidityLevelBuilder {
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

func (b *_AirConditioningDataSetPlantHumidityLevelBuilder) WithHumidityModeAndFlags(humidityModeAndFlags HVACHumidityModeAndFlags) AirConditioningDataSetPlantHumidityLevelBuilder {
	b.HumidityModeAndFlags = humidityModeAndFlags
	return b
}

func (b *_AirConditioningDataSetPlantHumidityLevelBuilder) WithHumidityModeAndFlagsBuilder(builderSupplier func(HVACHumidityModeAndFlagsBuilder) HVACHumidityModeAndFlagsBuilder) AirConditioningDataSetPlantHumidityLevelBuilder {
	builder := builderSupplier(b.HumidityModeAndFlags.CreateHVACHumidityModeAndFlagsBuilder())
	var err error
	b.HumidityModeAndFlags, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACHumidityModeAndFlagsBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataSetPlantHumidityLevelBuilder) WithHumidityType(humidityType HVACHumidityType) AirConditioningDataSetPlantHumidityLevelBuilder {
	b.HumidityType = humidityType
	return b
}

func (b *_AirConditioningDataSetPlantHumidityLevelBuilder) WithOptionalLevel(level HVACHumidity) AirConditioningDataSetPlantHumidityLevelBuilder {
	b.Level = level
	return b
}

func (b *_AirConditioningDataSetPlantHumidityLevelBuilder) WithOptionalLevelBuilder(builderSupplier func(HVACHumidityBuilder) HVACHumidityBuilder) AirConditioningDataSetPlantHumidityLevelBuilder {
	builder := builderSupplier(b.Level.CreateHVACHumidityBuilder())
	var err error
	b.Level, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACHumidityBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataSetPlantHumidityLevelBuilder) WithOptionalRawLevel(rawLevel HVACRawLevels) AirConditioningDataSetPlantHumidityLevelBuilder {
	b.RawLevel = rawLevel
	return b
}

func (b *_AirConditioningDataSetPlantHumidityLevelBuilder) WithOptionalRawLevelBuilder(builderSupplier func(HVACRawLevelsBuilder) HVACRawLevelsBuilder) AirConditioningDataSetPlantHumidityLevelBuilder {
	builder := builderSupplier(b.RawLevel.CreateHVACRawLevelsBuilder())
	var err error
	b.RawLevel, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACRawLevelsBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataSetPlantHumidityLevelBuilder) WithOptionalAuxLevel(auxLevel HVACAuxiliaryLevel) AirConditioningDataSetPlantHumidityLevelBuilder {
	b.AuxLevel = auxLevel
	return b
}

func (b *_AirConditioningDataSetPlantHumidityLevelBuilder) WithOptionalAuxLevelBuilder(builderSupplier func(HVACAuxiliaryLevelBuilder) HVACAuxiliaryLevelBuilder) AirConditioningDataSetPlantHumidityLevelBuilder {
	builder := builderSupplier(b.AuxLevel.CreateHVACAuxiliaryLevelBuilder())
	var err error
	b.AuxLevel, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACAuxiliaryLevelBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataSetPlantHumidityLevelBuilder) Build() (AirConditioningDataSetPlantHumidityLevel, error) {
	if b.ZoneList == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'zoneList' not set"))
	}
	if b.HumidityModeAndFlags == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'humidityModeAndFlags' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AirConditioningDataSetPlantHumidityLevel.deepCopy(), nil
}

func (b *_AirConditioningDataSetPlantHumidityLevelBuilder) MustBuild() AirConditioningDataSetPlantHumidityLevel {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AirConditioningDataSetPlantHumidityLevelBuilder) Done() AirConditioningDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAirConditioningDataBuilder().(*_AirConditioningDataBuilder)
	}
	return b.parentBuilder
}

func (b *_AirConditioningDataSetPlantHumidityLevelBuilder) buildForAirConditioningData() (AirConditioningData, error) {
	return b.Build()
}

func (b *_AirConditioningDataSetPlantHumidityLevelBuilder) DeepCopy() any {
	_copy := b.CreateAirConditioningDataSetPlantHumidityLevelBuilder().(*_AirConditioningDataSetPlantHumidityLevelBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAirConditioningDataSetPlantHumidityLevelBuilder creates a AirConditioningDataSetPlantHumidityLevelBuilder
func (b *_AirConditioningDataSetPlantHumidityLevel) CreateAirConditioningDataSetPlantHumidityLevelBuilder() AirConditioningDataSetPlantHumidityLevelBuilder {
	if b == nil {
		return NewAirConditioningDataSetPlantHumidityLevelBuilder()
	}
	return &_AirConditioningDataSetPlantHumidityLevelBuilder{_AirConditioningDataSetPlantHumidityLevel: b.deepCopy()}
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

func (m *_AirConditioningDataSetPlantHumidityLevel) GetParent() AirConditioningDataContract {
	return m.AirConditioningDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataSetPlantHumidityLevel) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataSetPlantHumidityLevel) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataSetPlantHumidityLevel) GetHumidityModeAndFlags() HVACHumidityModeAndFlags {
	return m.HumidityModeAndFlags
}

func (m *_AirConditioningDataSetPlantHumidityLevel) GetHumidityType() HVACHumidityType {
	return m.HumidityType
}

func (m *_AirConditioningDataSetPlantHumidityLevel) GetLevel() HVACHumidity {
	return m.Level
}

func (m *_AirConditioningDataSetPlantHumidityLevel) GetRawLevel() HVACRawLevels {
	return m.RawLevel
}

func (m *_AirConditioningDataSetPlantHumidityLevel) GetAuxLevel() HVACAuxiliaryLevel {
	return m.AuxLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAirConditioningDataSetPlantHumidityLevel(structType any) AirConditioningDataSetPlantHumidityLevel {
	if casted, ok := structType.(AirConditioningDataSetPlantHumidityLevel); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataSetPlantHumidityLevel); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataSetPlantHumidityLevel) GetTypeName() string {
	return "AirConditioningDataSetPlantHumidityLevel"
}

func (m *_AirConditioningDataSetPlantHumidityLevel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AirConditioningDataContract.(*_AirConditioningData).getLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (humidityModeAndFlags)
	lengthInBits += m.HumidityModeAndFlags.GetLengthInBits(ctx)

	// Simple field (humidityType)
	lengthInBits += 8

	// Optional Field (level)
	if m.Level != nil {
		lengthInBits += m.Level.GetLengthInBits(ctx)
	}

	// Optional Field (rawLevel)
	if m.RawLevel != nil {
		lengthInBits += m.RawLevel.GetLengthInBits(ctx)
	}

	// Optional Field (auxLevel)
	if m.AuxLevel != nil {
		lengthInBits += m.AuxLevel.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_AirConditioningDataSetPlantHumidityLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AirConditioningDataSetPlantHumidityLevel) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AirConditioningData) (__airConditioningDataSetPlantHumidityLevel AirConditioningDataSetPlantHumidityLevel, err error) {
	m.AirConditioningDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataSetPlantHumidityLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataSetPlantHumidityLevel")
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

	humidityModeAndFlags, err := ReadSimpleField[HVACHumidityModeAndFlags](ctx, "humidityModeAndFlags", ReadComplex[HVACHumidityModeAndFlags](HVACHumidityModeAndFlagsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'humidityModeAndFlags' field"))
	}
	m.HumidityModeAndFlags = humidityModeAndFlags

	humidityType, err := ReadEnumField[HVACHumidityType](ctx, "humidityType", "HVACHumidityType", ReadEnum(HVACHumidityTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'humidityType' field"))
	}
	m.HumidityType = humidityType

	var level HVACHumidity
	_level, err := ReadOptionalField[HVACHumidity](ctx, "level", ReadComplex[HVACHumidity](HVACHumidityParseWithBuffer, readBuffer), humidityModeAndFlags.GetIsLevelHumidity())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'level' field"))
	}
	if _level != nil {
		level = *_level
		m.Level = level
	}

	var rawLevel HVACRawLevels
	_rawLevel, err := ReadOptionalField[HVACRawLevels](ctx, "rawLevel", ReadComplex[HVACRawLevels](HVACRawLevelsParseWithBuffer, readBuffer), humidityModeAndFlags.GetIsLevelRaw())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rawLevel' field"))
	}
	if _rawLevel != nil {
		rawLevel = *_rawLevel
		m.RawLevel = rawLevel
	}

	var auxLevel HVACAuxiliaryLevel
	_auxLevel, err := ReadOptionalField[HVACAuxiliaryLevel](ctx, "auxLevel", ReadComplex[HVACAuxiliaryLevel](HVACAuxiliaryLevelParseWithBuffer, readBuffer), humidityModeAndFlags.GetIsAuxLevelUsed())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'auxLevel' field"))
	}
	if _auxLevel != nil {
		auxLevel = *_auxLevel
		m.AuxLevel = auxLevel
	}

	if closeErr := readBuffer.CloseContext("AirConditioningDataSetPlantHumidityLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataSetPlantHumidityLevel")
	}

	return m, nil
}

func (m *_AirConditioningDataSetPlantHumidityLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataSetPlantHumidityLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataSetPlantHumidityLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataSetPlantHumidityLevel")
		}

		if err := WriteSimpleField[byte](ctx, "zoneGroup", m.GetZoneGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneGroup' field")
		}

		if err := WriteSimpleField[HVACZoneList](ctx, "zoneList", m.GetZoneList(), WriteComplex[HVACZoneList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneList' field")
		}

		if err := WriteSimpleField[HVACHumidityModeAndFlags](ctx, "humidityModeAndFlags", m.GetHumidityModeAndFlags(), WriteComplex[HVACHumidityModeAndFlags](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'humidityModeAndFlags' field")
		}

		if err := WriteSimpleEnumField[HVACHumidityType](ctx, "humidityType", "HVACHumidityType", m.GetHumidityType(), WriteEnum[HVACHumidityType, uint8](HVACHumidityType.GetValue, HVACHumidityType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'humidityType' field")
		}

		if err := WriteOptionalField[HVACHumidity](ctx, "level", GetRef(m.GetLevel()), WriteComplex[HVACHumidity](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'level' field")
		}

		if err := WriteOptionalField[HVACRawLevels](ctx, "rawLevel", GetRef(m.GetRawLevel()), WriteComplex[HVACRawLevels](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'rawLevel' field")
		}

		if err := WriteOptionalField[HVACAuxiliaryLevel](ctx, "auxLevel", GetRef(m.GetAuxLevel()), WriteComplex[HVACAuxiliaryLevel](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'auxLevel' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataSetPlantHumidityLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataSetPlantHumidityLevel")
		}
		return nil
	}
	return m.AirConditioningDataContract.(*_AirConditioningData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataSetPlantHumidityLevel) IsAirConditioningDataSetPlantHumidityLevel() {}

func (m *_AirConditioningDataSetPlantHumidityLevel) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AirConditioningDataSetPlantHumidityLevel) deepCopy() *_AirConditioningDataSetPlantHumidityLevel {
	if m == nil {
		return nil
	}
	_AirConditioningDataSetPlantHumidityLevelCopy := &_AirConditioningDataSetPlantHumidityLevel{
		m.AirConditioningDataContract.(*_AirConditioningData).deepCopy(),
		m.ZoneGroup,
		utils.DeepCopy[HVACZoneList](m.ZoneList),
		utils.DeepCopy[HVACHumidityModeAndFlags](m.HumidityModeAndFlags),
		m.HumidityType,
		utils.DeepCopy[HVACHumidity](m.Level),
		utils.DeepCopy[HVACRawLevels](m.RawLevel),
		utils.DeepCopy[HVACAuxiliaryLevel](m.AuxLevel),
	}
	_AirConditioningDataSetPlantHumidityLevelCopy.AirConditioningDataContract.(*_AirConditioningData)._SubType = m
	return _AirConditioningDataSetPlantHumidityLevelCopy
}

func (m *_AirConditioningDataSetPlantHumidityLevel) String() string {
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