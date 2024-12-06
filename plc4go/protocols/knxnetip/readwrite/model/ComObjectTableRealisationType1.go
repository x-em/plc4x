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

// ComObjectTableRealisationType1 is the corresponding interface of ComObjectTableRealisationType1
type ComObjectTableRealisationType1 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ComObjectTable
	// GetNumEntries returns NumEntries (property field)
	GetNumEntries() uint8
	// GetRamFlagsTablePointer returns RamFlagsTablePointer (property field)
	GetRamFlagsTablePointer() uint8
	// GetComObjectDescriptors returns ComObjectDescriptors (property field)
	GetComObjectDescriptors() []GroupObjectDescriptorRealisationType1
	// IsComObjectTableRealisationType1 is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsComObjectTableRealisationType1()
	// CreateBuilder creates a ComObjectTableRealisationType1Builder
	CreateComObjectTableRealisationType1Builder() ComObjectTableRealisationType1Builder
}

// _ComObjectTableRealisationType1 is the data-structure of this message
type _ComObjectTableRealisationType1 struct {
	ComObjectTableContract
	NumEntries           uint8
	RamFlagsTablePointer uint8
	ComObjectDescriptors []GroupObjectDescriptorRealisationType1
}

var _ ComObjectTableRealisationType1 = (*_ComObjectTableRealisationType1)(nil)
var _ ComObjectTableRequirements = (*_ComObjectTableRealisationType1)(nil)

// NewComObjectTableRealisationType1 factory function for _ComObjectTableRealisationType1
func NewComObjectTableRealisationType1(numEntries uint8, ramFlagsTablePointer uint8, comObjectDescriptors []GroupObjectDescriptorRealisationType1) *_ComObjectTableRealisationType1 {
	_result := &_ComObjectTableRealisationType1{
		ComObjectTableContract: NewComObjectTable(),
		NumEntries:             numEntries,
		RamFlagsTablePointer:   ramFlagsTablePointer,
		ComObjectDescriptors:   comObjectDescriptors,
	}
	_result.ComObjectTableContract.(*_ComObjectTable)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ComObjectTableRealisationType1Builder is a builder for ComObjectTableRealisationType1
type ComObjectTableRealisationType1Builder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(numEntries uint8, ramFlagsTablePointer uint8, comObjectDescriptors []GroupObjectDescriptorRealisationType1) ComObjectTableRealisationType1Builder
	// WithNumEntries adds NumEntries (property field)
	WithNumEntries(uint8) ComObjectTableRealisationType1Builder
	// WithRamFlagsTablePointer adds RamFlagsTablePointer (property field)
	WithRamFlagsTablePointer(uint8) ComObjectTableRealisationType1Builder
	// WithComObjectDescriptors adds ComObjectDescriptors (property field)
	WithComObjectDescriptors(...GroupObjectDescriptorRealisationType1) ComObjectTableRealisationType1Builder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ComObjectTableBuilder
	// Build builds the ComObjectTableRealisationType1 or returns an error if something is wrong
	Build() (ComObjectTableRealisationType1, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ComObjectTableRealisationType1
}

// NewComObjectTableRealisationType1Builder() creates a ComObjectTableRealisationType1Builder
func NewComObjectTableRealisationType1Builder() ComObjectTableRealisationType1Builder {
	return &_ComObjectTableRealisationType1Builder{_ComObjectTableRealisationType1: new(_ComObjectTableRealisationType1)}
}

type _ComObjectTableRealisationType1Builder struct {
	*_ComObjectTableRealisationType1

	parentBuilder *_ComObjectTableBuilder

	err *utils.MultiError
}

var _ (ComObjectTableRealisationType1Builder) = (*_ComObjectTableRealisationType1Builder)(nil)

func (b *_ComObjectTableRealisationType1Builder) setParent(contract ComObjectTableContract) {
	b.ComObjectTableContract = contract
	contract.(*_ComObjectTable)._SubType = b._ComObjectTableRealisationType1
}

func (b *_ComObjectTableRealisationType1Builder) WithMandatoryFields(numEntries uint8, ramFlagsTablePointer uint8, comObjectDescriptors []GroupObjectDescriptorRealisationType1) ComObjectTableRealisationType1Builder {
	return b.WithNumEntries(numEntries).WithRamFlagsTablePointer(ramFlagsTablePointer).WithComObjectDescriptors(comObjectDescriptors...)
}

func (b *_ComObjectTableRealisationType1Builder) WithNumEntries(numEntries uint8) ComObjectTableRealisationType1Builder {
	b.NumEntries = numEntries
	return b
}

func (b *_ComObjectTableRealisationType1Builder) WithRamFlagsTablePointer(ramFlagsTablePointer uint8) ComObjectTableRealisationType1Builder {
	b.RamFlagsTablePointer = ramFlagsTablePointer
	return b
}

func (b *_ComObjectTableRealisationType1Builder) WithComObjectDescriptors(comObjectDescriptors ...GroupObjectDescriptorRealisationType1) ComObjectTableRealisationType1Builder {
	b.ComObjectDescriptors = comObjectDescriptors
	return b
}

func (b *_ComObjectTableRealisationType1Builder) Build() (ComObjectTableRealisationType1, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ComObjectTableRealisationType1.deepCopy(), nil
}

func (b *_ComObjectTableRealisationType1Builder) MustBuild() ComObjectTableRealisationType1 {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ComObjectTableRealisationType1Builder) Done() ComObjectTableBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewComObjectTableBuilder().(*_ComObjectTableBuilder)
	}
	return b.parentBuilder
}

func (b *_ComObjectTableRealisationType1Builder) buildForComObjectTable() (ComObjectTable, error) {
	return b.Build()
}

func (b *_ComObjectTableRealisationType1Builder) DeepCopy() any {
	_copy := b.CreateComObjectTableRealisationType1Builder().(*_ComObjectTableRealisationType1Builder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateComObjectTableRealisationType1Builder creates a ComObjectTableRealisationType1Builder
func (b *_ComObjectTableRealisationType1) CreateComObjectTableRealisationType1Builder() ComObjectTableRealisationType1Builder {
	if b == nil {
		return NewComObjectTableRealisationType1Builder()
	}
	return &_ComObjectTableRealisationType1Builder{_ComObjectTableRealisationType1: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ComObjectTableRealisationType1) GetFirmwareType() FirmwareType {
	return FirmwareType_SYSTEM_1
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ComObjectTableRealisationType1) GetParent() ComObjectTableContract {
	return m.ComObjectTableContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ComObjectTableRealisationType1) GetNumEntries() uint8 {
	return m.NumEntries
}

func (m *_ComObjectTableRealisationType1) GetRamFlagsTablePointer() uint8 {
	return m.RamFlagsTablePointer
}

func (m *_ComObjectTableRealisationType1) GetComObjectDescriptors() []GroupObjectDescriptorRealisationType1 {
	return m.ComObjectDescriptors
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastComObjectTableRealisationType1(structType any) ComObjectTableRealisationType1 {
	if casted, ok := structType.(ComObjectTableRealisationType1); ok {
		return casted
	}
	if casted, ok := structType.(*ComObjectTableRealisationType1); ok {
		return *casted
	}
	return nil
}

func (m *_ComObjectTableRealisationType1) GetTypeName() string {
	return "ComObjectTableRealisationType1"
}

func (m *_ComObjectTableRealisationType1) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ComObjectTableContract.(*_ComObjectTable).getLengthInBits(ctx))

	// Simple field (numEntries)
	lengthInBits += 8

	// Simple field (ramFlagsTablePointer)
	lengthInBits += 8

	// Array field
	if len(m.ComObjectDescriptors) > 0 {
		for _curItem, element := range m.ComObjectDescriptors {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ComObjectDescriptors), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_ComObjectTableRealisationType1) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ComObjectTableRealisationType1) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ComObjectTable, firmwareType FirmwareType) (__comObjectTableRealisationType1 ComObjectTableRealisationType1, err error) {
	m.ComObjectTableContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ComObjectTableRealisationType1"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ComObjectTableRealisationType1")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numEntries, err := ReadSimpleField(ctx, "numEntries", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numEntries' field"))
	}
	m.NumEntries = numEntries

	ramFlagsTablePointer, err := ReadSimpleField(ctx, "ramFlagsTablePointer", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ramFlagsTablePointer' field"))
	}
	m.RamFlagsTablePointer = ramFlagsTablePointer

	comObjectDescriptors, err := ReadCountArrayField[GroupObjectDescriptorRealisationType1](ctx, "comObjectDescriptors", ReadComplex[GroupObjectDescriptorRealisationType1](GroupObjectDescriptorRealisationType1ParseWithBuffer, readBuffer), uint64(numEntries))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'comObjectDescriptors' field"))
	}
	m.ComObjectDescriptors = comObjectDescriptors

	if closeErr := readBuffer.CloseContext("ComObjectTableRealisationType1"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ComObjectTableRealisationType1")
	}

	return m, nil
}

func (m *_ComObjectTableRealisationType1) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ComObjectTableRealisationType1) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ComObjectTableRealisationType1"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ComObjectTableRealisationType1")
		}

		if err := WriteSimpleField[uint8](ctx, "numEntries", m.GetNumEntries(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'numEntries' field")
		}

		if err := WriteSimpleField[uint8](ctx, "ramFlagsTablePointer", m.GetRamFlagsTablePointer(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'ramFlagsTablePointer' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "comObjectDescriptors", m.GetComObjectDescriptors(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'comObjectDescriptors' field")
		}

		if popErr := writeBuffer.PopContext("ComObjectTableRealisationType1"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ComObjectTableRealisationType1")
		}
		return nil
	}
	return m.ComObjectTableContract.(*_ComObjectTable).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ComObjectTableRealisationType1) IsComObjectTableRealisationType1() {}

func (m *_ComObjectTableRealisationType1) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ComObjectTableRealisationType1) deepCopy() *_ComObjectTableRealisationType1 {
	if m == nil {
		return nil
	}
	_ComObjectTableRealisationType1Copy := &_ComObjectTableRealisationType1{
		m.ComObjectTableContract.(*_ComObjectTable).deepCopy(),
		m.NumEntries,
		m.RamFlagsTablePointer,
		utils.DeepCopySlice[GroupObjectDescriptorRealisationType1, GroupObjectDescriptorRealisationType1](m.ComObjectDescriptors),
	}
	_ComObjectTableRealisationType1Copy.ComObjectTableContract.(*_ComObjectTable)._SubType = m
	return _ComObjectTableRealisationType1Copy
}

func (m *_ComObjectTableRealisationType1) String() string {
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