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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BVLCReadBroadcastDistributionTableAck is the corresponding interface of BVLCReadBroadcastDistributionTableAck
type BVLCReadBroadcastDistributionTableAck interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BVLC
	// GetTable returns Table (property field)
	GetTable() []BVLCBroadcastDistributionTableEntry
	// IsBVLCReadBroadcastDistributionTableAck is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBVLCReadBroadcastDistributionTableAck()
	// CreateBuilder creates a BVLCReadBroadcastDistributionTableAckBuilder
	CreateBVLCReadBroadcastDistributionTableAckBuilder() BVLCReadBroadcastDistributionTableAckBuilder
}

// _BVLCReadBroadcastDistributionTableAck is the data-structure of this message
type _BVLCReadBroadcastDistributionTableAck struct {
	BVLCContract
	Table []BVLCBroadcastDistributionTableEntry

	// Arguments.
	BvlcPayloadLength uint16
}

var _ BVLCReadBroadcastDistributionTableAck = (*_BVLCReadBroadcastDistributionTableAck)(nil)
var _ BVLCRequirements = (*_BVLCReadBroadcastDistributionTableAck)(nil)

// NewBVLCReadBroadcastDistributionTableAck factory function for _BVLCReadBroadcastDistributionTableAck
func NewBVLCReadBroadcastDistributionTableAck(table []BVLCBroadcastDistributionTableEntry, bvlcPayloadLength uint16) *_BVLCReadBroadcastDistributionTableAck {
	_result := &_BVLCReadBroadcastDistributionTableAck{
		BVLCContract: NewBVLC(),
		Table:        table,
	}
	_result.BVLCContract.(*_BVLC)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BVLCReadBroadcastDistributionTableAckBuilder is a builder for BVLCReadBroadcastDistributionTableAck
type BVLCReadBroadcastDistributionTableAckBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(table []BVLCBroadcastDistributionTableEntry) BVLCReadBroadcastDistributionTableAckBuilder
	// WithTable adds Table (property field)
	WithTable(...BVLCBroadcastDistributionTableEntry) BVLCReadBroadcastDistributionTableAckBuilder
	// WithArgBvlcPayloadLength sets a parser argument
	WithArgBvlcPayloadLength(uint16) BVLCReadBroadcastDistributionTableAckBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BVLCBuilder
	// Build builds the BVLCReadBroadcastDistributionTableAck or returns an error if something is wrong
	Build() (BVLCReadBroadcastDistributionTableAck, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BVLCReadBroadcastDistributionTableAck
}

// NewBVLCReadBroadcastDistributionTableAckBuilder() creates a BVLCReadBroadcastDistributionTableAckBuilder
func NewBVLCReadBroadcastDistributionTableAckBuilder() BVLCReadBroadcastDistributionTableAckBuilder {
	return &_BVLCReadBroadcastDistributionTableAckBuilder{_BVLCReadBroadcastDistributionTableAck: new(_BVLCReadBroadcastDistributionTableAck)}
}

type _BVLCReadBroadcastDistributionTableAckBuilder struct {
	*_BVLCReadBroadcastDistributionTableAck

	parentBuilder *_BVLCBuilder

	err *utils.MultiError
}

var _ (BVLCReadBroadcastDistributionTableAckBuilder) = (*_BVLCReadBroadcastDistributionTableAckBuilder)(nil)

func (b *_BVLCReadBroadcastDistributionTableAckBuilder) setParent(contract BVLCContract) {
	b.BVLCContract = contract
	contract.(*_BVLC)._SubType = b._BVLCReadBroadcastDistributionTableAck
}

func (b *_BVLCReadBroadcastDistributionTableAckBuilder) WithMandatoryFields(table []BVLCBroadcastDistributionTableEntry) BVLCReadBroadcastDistributionTableAckBuilder {
	return b.WithTable(table...)
}

func (b *_BVLCReadBroadcastDistributionTableAckBuilder) WithTable(table ...BVLCBroadcastDistributionTableEntry) BVLCReadBroadcastDistributionTableAckBuilder {
	b.Table = table
	return b
}

func (b *_BVLCReadBroadcastDistributionTableAckBuilder) WithArgBvlcPayloadLength(bvlcPayloadLength uint16) BVLCReadBroadcastDistributionTableAckBuilder {
	b.BvlcPayloadLength = bvlcPayloadLength
	return b
}

func (b *_BVLCReadBroadcastDistributionTableAckBuilder) Build() (BVLCReadBroadcastDistributionTableAck, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BVLCReadBroadcastDistributionTableAck.deepCopy(), nil
}

func (b *_BVLCReadBroadcastDistributionTableAckBuilder) MustBuild() BVLCReadBroadcastDistributionTableAck {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BVLCReadBroadcastDistributionTableAckBuilder) Done() BVLCBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBVLCBuilder().(*_BVLCBuilder)
	}
	return b.parentBuilder
}

func (b *_BVLCReadBroadcastDistributionTableAckBuilder) buildForBVLC() (BVLC, error) {
	return b.Build()
}

func (b *_BVLCReadBroadcastDistributionTableAckBuilder) DeepCopy() any {
	_copy := b.CreateBVLCReadBroadcastDistributionTableAckBuilder().(*_BVLCReadBroadcastDistributionTableAckBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBVLCReadBroadcastDistributionTableAckBuilder creates a BVLCReadBroadcastDistributionTableAckBuilder
func (b *_BVLCReadBroadcastDistributionTableAck) CreateBVLCReadBroadcastDistributionTableAckBuilder() BVLCReadBroadcastDistributionTableAckBuilder {
	if b == nil {
		return NewBVLCReadBroadcastDistributionTableAckBuilder()
	}
	return &_BVLCReadBroadcastDistributionTableAckBuilder{_BVLCReadBroadcastDistributionTableAck: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BVLCReadBroadcastDistributionTableAck) GetBvlcFunction() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BVLCReadBroadcastDistributionTableAck) GetParent() BVLCContract {
	return m.BVLCContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BVLCReadBroadcastDistributionTableAck) GetTable() []BVLCBroadcastDistributionTableEntry {
	return m.Table
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBVLCReadBroadcastDistributionTableAck(structType any) BVLCReadBroadcastDistributionTableAck {
	if casted, ok := structType.(BVLCReadBroadcastDistributionTableAck); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCReadBroadcastDistributionTableAck); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCReadBroadcastDistributionTableAck) GetTypeName() string {
	return "BVLCReadBroadcastDistributionTableAck"
}

func (m *_BVLCReadBroadcastDistributionTableAck) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BVLCContract.(*_BVLC).getLengthInBits(ctx))

	// Array field
	if len(m.Table) > 0 {
		for _, element := range m.Table {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BVLCReadBroadcastDistributionTableAck) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BVLCReadBroadcastDistributionTableAck) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BVLC, bvlcPayloadLength uint16) (__bVLCReadBroadcastDistributionTableAck BVLCReadBroadcastDistributionTableAck, err error) {
	m.BVLCContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCReadBroadcastDistributionTableAck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCReadBroadcastDistributionTableAck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	table, err := ReadLengthArrayField[BVLCBroadcastDistributionTableEntry](ctx, "table", ReadComplex[BVLCBroadcastDistributionTableEntry](BVLCBroadcastDistributionTableEntryParseWithBuffer, readBuffer), int(bvlcPayloadLength), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'table' field"))
	}
	m.Table = table

	if closeErr := readBuffer.CloseContext("BVLCReadBroadcastDistributionTableAck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCReadBroadcastDistributionTableAck")
	}

	return m, nil
}

func (m *_BVLCReadBroadcastDistributionTableAck) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BVLCReadBroadcastDistributionTableAck) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCReadBroadcastDistributionTableAck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BVLCReadBroadcastDistributionTableAck")
		}

		if err := WriteComplexTypeArrayField(ctx, "table", m.GetTable(), writeBuffer, codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'table' field")
		}

		if popErr := writeBuffer.PopContext("BVLCReadBroadcastDistributionTableAck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BVLCReadBroadcastDistributionTableAck")
		}
		return nil
	}
	return m.BVLCContract.(*_BVLC).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BVLCReadBroadcastDistributionTableAck) GetBvlcPayloadLength() uint16 {
	return m.BvlcPayloadLength
}

//
////

func (m *_BVLCReadBroadcastDistributionTableAck) IsBVLCReadBroadcastDistributionTableAck() {}

func (m *_BVLCReadBroadcastDistributionTableAck) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BVLCReadBroadcastDistributionTableAck) deepCopy() *_BVLCReadBroadcastDistributionTableAck {
	if m == nil {
		return nil
	}
	_BVLCReadBroadcastDistributionTableAckCopy := &_BVLCReadBroadcastDistributionTableAck{
		m.BVLCContract.(*_BVLC).deepCopy(),
		utils.DeepCopySlice[BVLCBroadcastDistributionTableEntry, BVLCBroadcastDistributionTableEntry](m.Table),
		m.BvlcPayloadLength,
	}
	_BVLCReadBroadcastDistributionTableAckCopy.BVLCContract.(*_BVLC)._SubType = m
	return _BVLCReadBroadcastDistributionTableAckCopy
}

func (m *_BVLCReadBroadcastDistributionTableAck) String() string {
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