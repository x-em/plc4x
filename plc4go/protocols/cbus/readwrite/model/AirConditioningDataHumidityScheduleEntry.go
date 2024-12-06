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

// AirConditioningDataHumidityScheduleEntry is the corresponding interface of AirConditioningDataHumidityScheduleEntry
type AirConditioningDataHumidityScheduleEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetEntry returns Entry (property field)
	GetEntry() uint8
	// GetFormat returns Format (property field)
	GetFormat() byte
	// GetHumidityModeAndFlags returns HumidityModeAndFlags (property field)
	GetHumidityModeAndFlags() HVACHumidityModeAndFlags
	// GetStartTime returns StartTime (property field)
	GetStartTime() HVACStartTime
	// GetLevel returns Level (property field)
	GetLevel() HVACHumidity
	// GetRawLevel returns RawLevel (property field)
	GetRawLevel() HVACRawLevels
	// IsAirConditioningDataHumidityScheduleEntry is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAirConditioningDataHumidityScheduleEntry()
	// CreateBuilder creates a AirConditioningDataHumidityScheduleEntryBuilder
	CreateAirConditioningDataHumidityScheduleEntryBuilder() AirConditioningDataHumidityScheduleEntryBuilder
}

// _AirConditioningDataHumidityScheduleEntry is the data-structure of this message
type _AirConditioningDataHumidityScheduleEntry struct {
	AirConditioningDataContract
	ZoneGroup            byte
	ZoneList             HVACZoneList
	Entry                uint8
	Format               byte
	HumidityModeAndFlags HVACHumidityModeAndFlags
	StartTime            HVACStartTime
	Level                HVACHumidity
	RawLevel             HVACRawLevels
}

var _ AirConditioningDataHumidityScheduleEntry = (*_AirConditioningDataHumidityScheduleEntry)(nil)
var _ AirConditioningDataRequirements = (*_AirConditioningDataHumidityScheduleEntry)(nil)

// NewAirConditioningDataHumidityScheduleEntry factory function for _AirConditioningDataHumidityScheduleEntry
func NewAirConditioningDataHumidityScheduleEntry(commandTypeContainer AirConditioningCommandTypeContainer, zoneGroup byte, zoneList HVACZoneList, entry uint8, format byte, humidityModeAndFlags HVACHumidityModeAndFlags, startTime HVACStartTime, level HVACHumidity, rawLevel HVACRawLevels) *_AirConditioningDataHumidityScheduleEntry {
	if zoneList == nil {
		panic("zoneList of type HVACZoneList for AirConditioningDataHumidityScheduleEntry must not be nil")
	}
	if humidityModeAndFlags == nil {
		panic("humidityModeAndFlags of type HVACHumidityModeAndFlags for AirConditioningDataHumidityScheduleEntry must not be nil")
	}
	if startTime == nil {
		panic("startTime of type HVACStartTime for AirConditioningDataHumidityScheduleEntry must not be nil")
	}
	_result := &_AirConditioningDataHumidityScheduleEntry{
		AirConditioningDataContract: NewAirConditioningData(commandTypeContainer),
		ZoneGroup:                   zoneGroup,
		ZoneList:                    zoneList,
		Entry:                       entry,
		Format:                      format,
		HumidityModeAndFlags:        humidityModeAndFlags,
		StartTime:                   startTime,
		Level:                       level,
		RawLevel:                    rawLevel,
	}
	_result.AirConditioningDataContract.(*_AirConditioningData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AirConditioningDataHumidityScheduleEntryBuilder is a builder for AirConditioningDataHumidityScheduleEntry
type AirConditioningDataHumidityScheduleEntryBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(zoneGroup byte, zoneList HVACZoneList, entry uint8, format byte, humidityModeAndFlags HVACHumidityModeAndFlags, startTime HVACStartTime) AirConditioningDataHumidityScheduleEntryBuilder
	// WithZoneGroup adds ZoneGroup (property field)
	WithZoneGroup(byte) AirConditioningDataHumidityScheduleEntryBuilder
	// WithZoneList adds ZoneList (property field)
	WithZoneList(HVACZoneList) AirConditioningDataHumidityScheduleEntryBuilder
	// WithZoneListBuilder adds ZoneList (property field) which is build by the builder
	WithZoneListBuilder(func(HVACZoneListBuilder) HVACZoneListBuilder) AirConditioningDataHumidityScheduleEntryBuilder
	// WithEntry adds Entry (property field)
	WithEntry(uint8) AirConditioningDataHumidityScheduleEntryBuilder
	// WithFormat adds Format (property field)
	WithFormat(byte) AirConditioningDataHumidityScheduleEntryBuilder
	// WithHumidityModeAndFlags adds HumidityModeAndFlags (property field)
	WithHumidityModeAndFlags(HVACHumidityModeAndFlags) AirConditioningDataHumidityScheduleEntryBuilder
	// WithHumidityModeAndFlagsBuilder adds HumidityModeAndFlags (property field) which is build by the builder
	WithHumidityModeAndFlagsBuilder(func(HVACHumidityModeAndFlagsBuilder) HVACHumidityModeAndFlagsBuilder) AirConditioningDataHumidityScheduleEntryBuilder
	// WithStartTime adds StartTime (property field)
	WithStartTime(HVACStartTime) AirConditioningDataHumidityScheduleEntryBuilder
	// WithStartTimeBuilder adds StartTime (property field) which is build by the builder
	WithStartTimeBuilder(func(HVACStartTimeBuilder) HVACStartTimeBuilder) AirConditioningDataHumidityScheduleEntryBuilder
	// WithLevel adds Level (property field)
	WithOptionalLevel(HVACHumidity) AirConditioningDataHumidityScheduleEntryBuilder
	// WithOptionalLevelBuilder adds Level (property field) which is build by the builder
	WithOptionalLevelBuilder(func(HVACHumidityBuilder) HVACHumidityBuilder) AirConditioningDataHumidityScheduleEntryBuilder
	// WithRawLevel adds RawLevel (property field)
	WithOptionalRawLevel(HVACRawLevels) AirConditioningDataHumidityScheduleEntryBuilder
	// WithOptionalRawLevelBuilder adds RawLevel (property field) which is build by the builder
	WithOptionalRawLevelBuilder(func(HVACRawLevelsBuilder) HVACRawLevelsBuilder) AirConditioningDataHumidityScheduleEntryBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() AirConditioningDataBuilder
	// Build builds the AirConditioningDataHumidityScheduleEntry or returns an error if something is wrong
	Build() (AirConditioningDataHumidityScheduleEntry, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AirConditioningDataHumidityScheduleEntry
}

// NewAirConditioningDataHumidityScheduleEntryBuilder() creates a AirConditioningDataHumidityScheduleEntryBuilder
func NewAirConditioningDataHumidityScheduleEntryBuilder() AirConditioningDataHumidityScheduleEntryBuilder {
	return &_AirConditioningDataHumidityScheduleEntryBuilder{_AirConditioningDataHumidityScheduleEntry: new(_AirConditioningDataHumidityScheduleEntry)}
}

type _AirConditioningDataHumidityScheduleEntryBuilder struct {
	*_AirConditioningDataHumidityScheduleEntry

	parentBuilder *_AirConditioningDataBuilder

	err *utils.MultiError
}

var _ (AirConditioningDataHumidityScheduleEntryBuilder) = (*_AirConditioningDataHumidityScheduleEntryBuilder)(nil)

func (b *_AirConditioningDataHumidityScheduleEntryBuilder) setParent(contract AirConditioningDataContract) {
	b.AirConditioningDataContract = contract
	contract.(*_AirConditioningData)._SubType = b._AirConditioningDataHumidityScheduleEntry
}

func (b *_AirConditioningDataHumidityScheduleEntryBuilder) WithMandatoryFields(zoneGroup byte, zoneList HVACZoneList, entry uint8, format byte, humidityModeAndFlags HVACHumidityModeAndFlags, startTime HVACStartTime) AirConditioningDataHumidityScheduleEntryBuilder {
	return b.WithZoneGroup(zoneGroup).WithZoneList(zoneList).WithEntry(entry).WithFormat(format).WithHumidityModeAndFlags(humidityModeAndFlags).WithStartTime(startTime)
}

func (b *_AirConditioningDataHumidityScheduleEntryBuilder) WithZoneGroup(zoneGroup byte) AirConditioningDataHumidityScheduleEntryBuilder {
	b.ZoneGroup = zoneGroup
	return b
}

func (b *_AirConditioningDataHumidityScheduleEntryBuilder) WithZoneList(zoneList HVACZoneList) AirConditioningDataHumidityScheduleEntryBuilder {
	b.ZoneList = zoneList
	return b
}

func (b *_AirConditioningDataHumidityScheduleEntryBuilder) WithZoneListBuilder(builderSupplier func(HVACZoneListBuilder) HVACZoneListBuilder) AirConditioningDataHumidityScheduleEntryBuilder {
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

func (b *_AirConditioningDataHumidityScheduleEntryBuilder) WithEntry(entry uint8) AirConditioningDataHumidityScheduleEntryBuilder {
	b.Entry = entry
	return b
}

func (b *_AirConditioningDataHumidityScheduleEntryBuilder) WithFormat(format byte) AirConditioningDataHumidityScheduleEntryBuilder {
	b.Format = format
	return b
}

func (b *_AirConditioningDataHumidityScheduleEntryBuilder) WithHumidityModeAndFlags(humidityModeAndFlags HVACHumidityModeAndFlags) AirConditioningDataHumidityScheduleEntryBuilder {
	b.HumidityModeAndFlags = humidityModeAndFlags
	return b
}

func (b *_AirConditioningDataHumidityScheduleEntryBuilder) WithHumidityModeAndFlagsBuilder(builderSupplier func(HVACHumidityModeAndFlagsBuilder) HVACHumidityModeAndFlagsBuilder) AirConditioningDataHumidityScheduleEntryBuilder {
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

func (b *_AirConditioningDataHumidityScheduleEntryBuilder) WithStartTime(startTime HVACStartTime) AirConditioningDataHumidityScheduleEntryBuilder {
	b.StartTime = startTime
	return b
}

func (b *_AirConditioningDataHumidityScheduleEntryBuilder) WithStartTimeBuilder(builderSupplier func(HVACStartTimeBuilder) HVACStartTimeBuilder) AirConditioningDataHumidityScheduleEntryBuilder {
	builder := builderSupplier(b.StartTime.CreateHVACStartTimeBuilder())
	var err error
	b.StartTime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACStartTimeBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataHumidityScheduleEntryBuilder) WithOptionalLevel(level HVACHumidity) AirConditioningDataHumidityScheduleEntryBuilder {
	b.Level = level
	return b
}

func (b *_AirConditioningDataHumidityScheduleEntryBuilder) WithOptionalLevelBuilder(builderSupplier func(HVACHumidityBuilder) HVACHumidityBuilder) AirConditioningDataHumidityScheduleEntryBuilder {
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

func (b *_AirConditioningDataHumidityScheduleEntryBuilder) WithOptionalRawLevel(rawLevel HVACRawLevels) AirConditioningDataHumidityScheduleEntryBuilder {
	b.RawLevel = rawLevel
	return b
}

func (b *_AirConditioningDataHumidityScheduleEntryBuilder) WithOptionalRawLevelBuilder(builderSupplier func(HVACRawLevelsBuilder) HVACRawLevelsBuilder) AirConditioningDataHumidityScheduleEntryBuilder {
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

func (b *_AirConditioningDataHumidityScheduleEntryBuilder) Build() (AirConditioningDataHumidityScheduleEntry, error) {
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
	if b.StartTime == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'startTime' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AirConditioningDataHumidityScheduleEntry.deepCopy(), nil
}

func (b *_AirConditioningDataHumidityScheduleEntryBuilder) MustBuild() AirConditioningDataHumidityScheduleEntry {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AirConditioningDataHumidityScheduleEntryBuilder) Done() AirConditioningDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAirConditioningDataBuilder().(*_AirConditioningDataBuilder)
	}
	return b.parentBuilder
}

func (b *_AirConditioningDataHumidityScheduleEntryBuilder) buildForAirConditioningData() (AirConditioningData, error) {
	return b.Build()
}

func (b *_AirConditioningDataHumidityScheduleEntryBuilder) DeepCopy() any {
	_copy := b.CreateAirConditioningDataHumidityScheduleEntryBuilder().(*_AirConditioningDataHumidityScheduleEntryBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAirConditioningDataHumidityScheduleEntryBuilder creates a AirConditioningDataHumidityScheduleEntryBuilder
func (b *_AirConditioningDataHumidityScheduleEntry) CreateAirConditioningDataHumidityScheduleEntryBuilder() AirConditioningDataHumidityScheduleEntryBuilder {
	if b == nil {
		return NewAirConditioningDataHumidityScheduleEntryBuilder()
	}
	return &_AirConditioningDataHumidityScheduleEntryBuilder{_AirConditioningDataHumidityScheduleEntry: b.deepCopy()}
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

func (m *_AirConditioningDataHumidityScheduleEntry) GetParent() AirConditioningDataContract {
	return m.AirConditioningDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataHumidityScheduleEntry) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataHumidityScheduleEntry) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataHumidityScheduleEntry) GetEntry() uint8 {
	return m.Entry
}

func (m *_AirConditioningDataHumidityScheduleEntry) GetFormat() byte {
	return m.Format
}

func (m *_AirConditioningDataHumidityScheduleEntry) GetHumidityModeAndFlags() HVACHumidityModeAndFlags {
	return m.HumidityModeAndFlags
}

func (m *_AirConditioningDataHumidityScheduleEntry) GetStartTime() HVACStartTime {
	return m.StartTime
}

func (m *_AirConditioningDataHumidityScheduleEntry) GetLevel() HVACHumidity {
	return m.Level
}

func (m *_AirConditioningDataHumidityScheduleEntry) GetRawLevel() HVACRawLevels {
	return m.RawLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAirConditioningDataHumidityScheduleEntry(structType any) AirConditioningDataHumidityScheduleEntry {
	if casted, ok := structType.(AirConditioningDataHumidityScheduleEntry); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataHumidityScheduleEntry); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataHumidityScheduleEntry) GetTypeName() string {
	return "AirConditioningDataHumidityScheduleEntry"
}

func (m *_AirConditioningDataHumidityScheduleEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AirConditioningDataContract.(*_AirConditioningData).getLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (entry)
	lengthInBits += 8

	// Simple field (format)
	lengthInBits += 8

	// Simple field (humidityModeAndFlags)
	lengthInBits += m.HumidityModeAndFlags.GetLengthInBits(ctx)

	// Simple field (startTime)
	lengthInBits += m.StartTime.GetLengthInBits(ctx)

	// Optional Field (level)
	if m.Level != nil {
		lengthInBits += m.Level.GetLengthInBits(ctx)
	}

	// Optional Field (rawLevel)
	if m.RawLevel != nil {
		lengthInBits += m.RawLevel.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_AirConditioningDataHumidityScheduleEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AirConditioningDataHumidityScheduleEntry) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AirConditioningData) (__airConditioningDataHumidityScheduleEntry AirConditioningDataHumidityScheduleEntry, err error) {
	m.AirConditioningDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataHumidityScheduleEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataHumidityScheduleEntry")
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

	entry, err := ReadSimpleField(ctx, "entry", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'entry' field"))
	}
	m.Entry = entry

	format, err := ReadSimpleField(ctx, "format", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'format' field"))
	}
	m.Format = format

	humidityModeAndFlags, err := ReadSimpleField[HVACHumidityModeAndFlags](ctx, "humidityModeAndFlags", ReadComplex[HVACHumidityModeAndFlags](HVACHumidityModeAndFlagsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'humidityModeAndFlags' field"))
	}
	m.HumidityModeAndFlags = humidityModeAndFlags

	startTime, err := ReadSimpleField[HVACStartTime](ctx, "startTime", ReadComplex[HVACStartTime](HVACStartTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startTime' field"))
	}
	m.StartTime = startTime

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

	if closeErr := readBuffer.CloseContext("AirConditioningDataHumidityScheduleEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataHumidityScheduleEntry")
	}

	return m, nil
}

func (m *_AirConditioningDataHumidityScheduleEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataHumidityScheduleEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataHumidityScheduleEntry"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataHumidityScheduleEntry")
		}

		if err := WriteSimpleField[byte](ctx, "zoneGroup", m.GetZoneGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneGroup' field")
		}

		if err := WriteSimpleField[HVACZoneList](ctx, "zoneList", m.GetZoneList(), WriteComplex[HVACZoneList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneList' field")
		}

		if err := WriteSimpleField[uint8](ctx, "entry", m.GetEntry(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'entry' field")
		}

		if err := WriteSimpleField[byte](ctx, "format", m.GetFormat(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'format' field")
		}

		if err := WriteSimpleField[HVACHumidityModeAndFlags](ctx, "humidityModeAndFlags", m.GetHumidityModeAndFlags(), WriteComplex[HVACHumidityModeAndFlags](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'humidityModeAndFlags' field")
		}

		if err := WriteSimpleField[HVACStartTime](ctx, "startTime", m.GetStartTime(), WriteComplex[HVACStartTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'startTime' field")
		}

		if err := WriteOptionalField[HVACHumidity](ctx, "level", GetRef(m.GetLevel()), WriteComplex[HVACHumidity](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'level' field")
		}

		if err := WriteOptionalField[HVACRawLevels](ctx, "rawLevel", GetRef(m.GetRawLevel()), WriteComplex[HVACRawLevels](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'rawLevel' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataHumidityScheduleEntry"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataHumidityScheduleEntry")
		}
		return nil
	}
	return m.AirConditioningDataContract.(*_AirConditioningData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataHumidityScheduleEntry) IsAirConditioningDataHumidityScheduleEntry() {}

func (m *_AirConditioningDataHumidityScheduleEntry) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AirConditioningDataHumidityScheduleEntry) deepCopy() *_AirConditioningDataHumidityScheduleEntry {
	if m == nil {
		return nil
	}
	_AirConditioningDataHumidityScheduleEntryCopy := &_AirConditioningDataHumidityScheduleEntry{
		m.AirConditioningDataContract.(*_AirConditioningData).deepCopy(),
		m.ZoneGroup,
		utils.DeepCopy[HVACZoneList](m.ZoneList),
		m.Entry,
		m.Format,
		utils.DeepCopy[HVACHumidityModeAndFlags](m.HumidityModeAndFlags),
		utils.DeepCopy[HVACStartTime](m.StartTime),
		utils.DeepCopy[HVACHumidity](m.Level),
		utils.DeepCopy[HVACRawLevels](m.RawLevel),
	}
	_AirConditioningDataHumidityScheduleEntryCopy.AirConditioningDataContract.(*_AirConditioningData)._SubType = m
	return _AirConditioningDataHumidityScheduleEntryCopy
}

func (m *_AirConditioningDataHumidityScheduleEntry) String() string {
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