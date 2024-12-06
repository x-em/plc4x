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

// HVACZoneList is the corresponding interface of HVACZoneList
type HVACZoneList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetExpansion returns Expansion (property field)
	GetExpansion() bool
	// GetZone6 returns Zone6 (property field)
	GetZone6() bool
	// GetZone5 returns Zone5 (property field)
	GetZone5() bool
	// GetZone4 returns Zone4 (property field)
	GetZone4() bool
	// GetZone3 returns Zone3 (property field)
	GetZone3() bool
	// GetZone2 returns Zone2 (property field)
	GetZone2() bool
	// GetZone1 returns Zone1 (property field)
	GetZone1() bool
	// GetZone0 returns Zone0 (property field)
	GetZone0() bool
	// GetUnswitchedZone returns UnswitchedZone (virtual field)
	GetUnswitchedZone() bool
	// IsHVACZoneList is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsHVACZoneList()
	// CreateBuilder creates a HVACZoneListBuilder
	CreateHVACZoneListBuilder() HVACZoneListBuilder
}

// _HVACZoneList is the data-structure of this message
type _HVACZoneList struct {
	Expansion bool
	Zone6     bool
	Zone5     bool
	Zone4     bool
	Zone3     bool
	Zone2     bool
	Zone1     bool
	Zone0     bool
}

var _ HVACZoneList = (*_HVACZoneList)(nil)

// NewHVACZoneList factory function for _HVACZoneList
func NewHVACZoneList(expansion bool, zone6 bool, zone5 bool, zone4 bool, zone3 bool, zone2 bool, zone1 bool, zone0 bool) *_HVACZoneList {
	return &_HVACZoneList{Expansion: expansion, Zone6: zone6, Zone5: zone5, Zone4: zone4, Zone3: zone3, Zone2: zone2, Zone1: zone1, Zone0: zone0}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// HVACZoneListBuilder is a builder for HVACZoneList
type HVACZoneListBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(expansion bool, zone6 bool, zone5 bool, zone4 bool, zone3 bool, zone2 bool, zone1 bool, zone0 bool) HVACZoneListBuilder
	// WithExpansion adds Expansion (property field)
	WithExpansion(bool) HVACZoneListBuilder
	// WithZone6 adds Zone6 (property field)
	WithZone6(bool) HVACZoneListBuilder
	// WithZone5 adds Zone5 (property field)
	WithZone5(bool) HVACZoneListBuilder
	// WithZone4 adds Zone4 (property field)
	WithZone4(bool) HVACZoneListBuilder
	// WithZone3 adds Zone3 (property field)
	WithZone3(bool) HVACZoneListBuilder
	// WithZone2 adds Zone2 (property field)
	WithZone2(bool) HVACZoneListBuilder
	// WithZone1 adds Zone1 (property field)
	WithZone1(bool) HVACZoneListBuilder
	// WithZone0 adds Zone0 (property field)
	WithZone0(bool) HVACZoneListBuilder
	// Build builds the HVACZoneList or returns an error if something is wrong
	Build() (HVACZoneList, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() HVACZoneList
}

// NewHVACZoneListBuilder() creates a HVACZoneListBuilder
func NewHVACZoneListBuilder() HVACZoneListBuilder {
	return &_HVACZoneListBuilder{_HVACZoneList: new(_HVACZoneList)}
}

type _HVACZoneListBuilder struct {
	*_HVACZoneList

	err *utils.MultiError
}

var _ (HVACZoneListBuilder) = (*_HVACZoneListBuilder)(nil)

func (b *_HVACZoneListBuilder) WithMandatoryFields(expansion bool, zone6 bool, zone5 bool, zone4 bool, zone3 bool, zone2 bool, zone1 bool, zone0 bool) HVACZoneListBuilder {
	return b.WithExpansion(expansion).WithZone6(zone6).WithZone5(zone5).WithZone4(zone4).WithZone3(zone3).WithZone2(zone2).WithZone1(zone1).WithZone0(zone0)
}

func (b *_HVACZoneListBuilder) WithExpansion(expansion bool) HVACZoneListBuilder {
	b.Expansion = expansion
	return b
}

func (b *_HVACZoneListBuilder) WithZone6(zone6 bool) HVACZoneListBuilder {
	b.Zone6 = zone6
	return b
}

func (b *_HVACZoneListBuilder) WithZone5(zone5 bool) HVACZoneListBuilder {
	b.Zone5 = zone5
	return b
}

func (b *_HVACZoneListBuilder) WithZone4(zone4 bool) HVACZoneListBuilder {
	b.Zone4 = zone4
	return b
}

func (b *_HVACZoneListBuilder) WithZone3(zone3 bool) HVACZoneListBuilder {
	b.Zone3 = zone3
	return b
}

func (b *_HVACZoneListBuilder) WithZone2(zone2 bool) HVACZoneListBuilder {
	b.Zone2 = zone2
	return b
}

func (b *_HVACZoneListBuilder) WithZone1(zone1 bool) HVACZoneListBuilder {
	b.Zone1 = zone1
	return b
}

func (b *_HVACZoneListBuilder) WithZone0(zone0 bool) HVACZoneListBuilder {
	b.Zone0 = zone0
	return b
}

func (b *_HVACZoneListBuilder) Build() (HVACZoneList, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._HVACZoneList.deepCopy(), nil
}

func (b *_HVACZoneListBuilder) MustBuild() HVACZoneList {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_HVACZoneListBuilder) DeepCopy() any {
	_copy := b.CreateHVACZoneListBuilder().(*_HVACZoneListBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateHVACZoneListBuilder creates a HVACZoneListBuilder
func (b *_HVACZoneList) CreateHVACZoneListBuilder() HVACZoneListBuilder {
	if b == nil {
		return NewHVACZoneListBuilder()
	}
	return &_HVACZoneListBuilder{_HVACZoneList: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HVACZoneList) GetExpansion() bool {
	return m.Expansion
}

func (m *_HVACZoneList) GetZone6() bool {
	return m.Zone6
}

func (m *_HVACZoneList) GetZone5() bool {
	return m.Zone5
}

func (m *_HVACZoneList) GetZone4() bool {
	return m.Zone4
}

func (m *_HVACZoneList) GetZone3() bool {
	return m.Zone3
}

func (m *_HVACZoneList) GetZone2() bool {
	return m.Zone2
}

func (m *_HVACZoneList) GetZone1() bool {
	return m.Zone1
}

func (m *_HVACZoneList) GetZone0() bool {
	return m.Zone0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_HVACZoneList) GetUnswitchedZone() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetZone0())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastHVACZoneList(structType any) HVACZoneList {
	if casted, ok := structType.(HVACZoneList); ok {
		return casted
	}
	if casted, ok := structType.(*HVACZoneList); ok {
		return *casted
	}
	return nil
}

func (m *_HVACZoneList) GetTypeName() string {
	return "HVACZoneList"
}

func (m *_HVACZoneList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (expansion)
	lengthInBits += 1

	// Simple field (zone6)
	lengthInBits += 1

	// Simple field (zone5)
	lengthInBits += 1

	// Simple field (zone4)
	lengthInBits += 1

	// Simple field (zone3)
	lengthInBits += 1

	// Simple field (zone2)
	lengthInBits += 1

	// Simple field (zone1)
	lengthInBits += 1

	// Simple field (zone0)
	lengthInBits += 1

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_HVACZoneList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HVACZoneListParse(ctx context.Context, theBytes []byte) (HVACZoneList, error) {
	return HVACZoneListParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func HVACZoneListParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (HVACZoneList, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (HVACZoneList, error) {
		return HVACZoneListParseWithBuffer(ctx, readBuffer)
	}
}

func HVACZoneListParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HVACZoneList, error) {
	v, err := (&_HVACZoneList{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_HVACZoneList) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__hVACZoneList HVACZoneList, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HVACZoneList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HVACZoneList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	expansion, err := ReadSimpleField(ctx, "expansion", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'expansion' field"))
	}
	m.Expansion = expansion

	zone6, err := ReadSimpleField(ctx, "zone6", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zone6' field"))
	}
	m.Zone6 = zone6

	zone5, err := ReadSimpleField(ctx, "zone5", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zone5' field"))
	}
	m.Zone5 = zone5

	zone4, err := ReadSimpleField(ctx, "zone4", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zone4' field"))
	}
	m.Zone4 = zone4

	zone3, err := ReadSimpleField(ctx, "zone3", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zone3' field"))
	}
	m.Zone3 = zone3

	zone2, err := ReadSimpleField(ctx, "zone2", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zone2' field"))
	}
	m.Zone2 = zone2

	zone1, err := ReadSimpleField(ctx, "zone1", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zone1' field"))
	}
	m.Zone1 = zone1

	zone0, err := ReadSimpleField(ctx, "zone0", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zone0' field"))
	}
	m.Zone0 = zone0

	unswitchedZone, err := ReadVirtualField[bool](ctx, "unswitchedZone", (*bool)(nil), zone0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unswitchedZone' field"))
	}
	_ = unswitchedZone

	if closeErr := readBuffer.CloseContext("HVACZoneList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HVACZoneList")
	}

	return m, nil
}

func (m *_HVACZoneList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HVACZoneList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("HVACZoneList"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for HVACZoneList")
	}

	if err := WriteSimpleField[bool](ctx, "expansion", m.GetExpansion(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'expansion' field")
	}

	if err := WriteSimpleField[bool](ctx, "zone6", m.GetZone6(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'zone6' field")
	}

	if err := WriteSimpleField[bool](ctx, "zone5", m.GetZone5(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'zone5' field")
	}

	if err := WriteSimpleField[bool](ctx, "zone4", m.GetZone4(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'zone4' field")
	}

	if err := WriteSimpleField[bool](ctx, "zone3", m.GetZone3(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'zone3' field")
	}

	if err := WriteSimpleField[bool](ctx, "zone2", m.GetZone2(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'zone2' field")
	}

	if err := WriteSimpleField[bool](ctx, "zone1", m.GetZone1(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'zone1' field")
	}

	if err := WriteSimpleField[bool](ctx, "zone0", m.GetZone0(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'zone0' field")
	}
	// Virtual field
	unswitchedZone := m.GetUnswitchedZone()
	_ = unswitchedZone
	if _unswitchedZoneErr := writeBuffer.WriteVirtual(ctx, "unswitchedZone", m.GetUnswitchedZone()); _unswitchedZoneErr != nil {
		return errors.Wrap(_unswitchedZoneErr, "Error serializing 'unswitchedZone' field")
	}

	if popErr := writeBuffer.PopContext("HVACZoneList"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for HVACZoneList")
	}
	return nil
}

func (m *_HVACZoneList) IsHVACZoneList() {}

func (m *_HVACZoneList) DeepCopy() any {
	return m.deepCopy()
}

func (m *_HVACZoneList) deepCopy() *_HVACZoneList {
	if m == nil {
		return nil
	}
	_HVACZoneListCopy := &_HVACZoneList{
		m.Expansion,
		m.Zone6,
		m.Zone5,
		m.Zone4,
		m.Zone3,
		m.Zone2,
		m.Zone1,
		m.Zone0,
	}
	return _HVACZoneListCopy
}

func (m *_HVACZoneList) String() string {
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
