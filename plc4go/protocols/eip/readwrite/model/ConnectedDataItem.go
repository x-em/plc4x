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

// ConnectedDataItem is the corresponding interface of ConnectedDataItem
type ConnectedDataItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	TypeId
	// GetSequenceCount returns SequenceCount (property field)
	GetSequenceCount() uint16
	// GetService returns Service (property field)
	GetService() CipService
	// IsConnectedDataItem is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsConnectedDataItem()
	// CreateBuilder creates a ConnectedDataItemBuilder
	CreateConnectedDataItemBuilder() ConnectedDataItemBuilder
}

// _ConnectedDataItem is the data-structure of this message
type _ConnectedDataItem struct {
	TypeIdContract
	SequenceCount uint16
	Service       CipService
}

var _ ConnectedDataItem = (*_ConnectedDataItem)(nil)
var _ TypeIdRequirements = (*_ConnectedDataItem)(nil)

// NewConnectedDataItem factory function for _ConnectedDataItem
func NewConnectedDataItem(sequenceCount uint16, service CipService) *_ConnectedDataItem {
	if service == nil {
		panic("service of type CipService for ConnectedDataItem must not be nil")
	}
	_result := &_ConnectedDataItem{
		TypeIdContract: NewTypeId(),
		SequenceCount:  sequenceCount,
		Service:        service,
	}
	_result.TypeIdContract.(*_TypeId)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ConnectedDataItemBuilder is a builder for ConnectedDataItem
type ConnectedDataItemBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(sequenceCount uint16, service CipService) ConnectedDataItemBuilder
	// WithSequenceCount adds SequenceCount (property field)
	WithSequenceCount(uint16) ConnectedDataItemBuilder
	// WithService adds Service (property field)
	WithService(CipService) ConnectedDataItemBuilder
	// WithServiceBuilder adds Service (property field) which is build by the builder
	WithServiceBuilder(func(CipServiceBuilder) CipServiceBuilder) ConnectedDataItemBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() TypeIdBuilder
	// Build builds the ConnectedDataItem or returns an error if something is wrong
	Build() (ConnectedDataItem, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ConnectedDataItem
}

// NewConnectedDataItemBuilder() creates a ConnectedDataItemBuilder
func NewConnectedDataItemBuilder() ConnectedDataItemBuilder {
	return &_ConnectedDataItemBuilder{_ConnectedDataItem: new(_ConnectedDataItem)}
}

type _ConnectedDataItemBuilder struct {
	*_ConnectedDataItem

	parentBuilder *_TypeIdBuilder

	err *utils.MultiError
}

var _ (ConnectedDataItemBuilder) = (*_ConnectedDataItemBuilder)(nil)

func (b *_ConnectedDataItemBuilder) setParent(contract TypeIdContract) {
	b.TypeIdContract = contract
	contract.(*_TypeId)._SubType = b._ConnectedDataItem
}

func (b *_ConnectedDataItemBuilder) WithMandatoryFields(sequenceCount uint16, service CipService) ConnectedDataItemBuilder {
	return b.WithSequenceCount(sequenceCount).WithService(service)
}

func (b *_ConnectedDataItemBuilder) WithSequenceCount(sequenceCount uint16) ConnectedDataItemBuilder {
	b.SequenceCount = sequenceCount
	return b
}

func (b *_ConnectedDataItemBuilder) WithService(service CipService) ConnectedDataItemBuilder {
	b.Service = service
	return b
}

func (b *_ConnectedDataItemBuilder) WithServiceBuilder(builderSupplier func(CipServiceBuilder) CipServiceBuilder) ConnectedDataItemBuilder {
	builder := builderSupplier(b.Service.CreateCipServiceBuilder())
	var err error
	b.Service, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "CipServiceBuilder failed"))
	}
	return b
}

func (b *_ConnectedDataItemBuilder) Build() (ConnectedDataItem, error) {
	if b.Service == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'service' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ConnectedDataItem.deepCopy(), nil
}

func (b *_ConnectedDataItemBuilder) MustBuild() ConnectedDataItem {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ConnectedDataItemBuilder) Done() TypeIdBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewTypeIdBuilder().(*_TypeIdBuilder)
	}
	return b.parentBuilder
}

func (b *_ConnectedDataItemBuilder) buildForTypeId() (TypeId, error) {
	return b.Build()
}

func (b *_ConnectedDataItemBuilder) DeepCopy() any {
	_copy := b.CreateConnectedDataItemBuilder().(*_ConnectedDataItemBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateConnectedDataItemBuilder creates a ConnectedDataItemBuilder
func (b *_ConnectedDataItem) CreateConnectedDataItemBuilder() ConnectedDataItemBuilder {
	if b == nil {
		return NewConnectedDataItemBuilder()
	}
	return &_ConnectedDataItemBuilder{_ConnectedDataItem: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConnectedDataItem) GetId() uint16 {
	return 0x00B1
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConnectedDataItem) GetParent() TypeIdContract {
	return m.TypeIdContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ConnectedDataItem) GetSequenceCount() uint16 {
	return m.SequenceCount
}

func (m *_ConnectedDataItem) GetService() CipService {
	return m.Service
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastConnectedDataItem(structType any) ConnectedDataItem {
	if casted, ok := structType.(ConnectedDataItem); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectedDataItem); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectedDataItem) GetTypeName() string {
	return "ConnectedDataItem"
}

func (m *_ConnectedDataItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.TypeIdContract.(*_TypeId).getLengthInBits(ctx))

	// Implicit Field (packetSize)
	lengthInBits += 16

	// Simple field (sequenceCount)
	lengthInBits += 16

	// Simple field (service)
	lengthInBits += m.Service.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ConnectedDataItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ConnectedDataItem) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_TypeId) (__connectedDataItem ConnectedDataItem, err error) {
	m.TypeIdContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectedDataItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectedDataItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	packetSize, err := ReadImplicitField[uint16](ctx, "packetSize", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'packetSize' field"))
	}
	_ = packetSize

	sequenceCount, err := ReadSimpleField(ctx, "sequenceCount", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sequenceCount' field"))
	}
	m.SequenceCount = sequenceCount

	service, err := ReadSimpleField[CipService](ctx, "service", ReadComplex[CipService](CipServiceParseWithBufferProducer[CipService]((bool)(bool(true)), (uint16)(uint16(packetSize)-uint16(uint16(2)))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'service' field"))
	}
	m.Service = service

	if closeErr := readBuffer.CloseContext("ConnectedDataItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectedDataItem")
	}

	return m, nil
}

func (m *_ConnectedDataItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConnectedDataItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectedDataItem"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectedDataItem")
		}
		packetSize := uint16(uint16(m.GetService().GetLengthInBytes(ctx)) + uint16(uint16(2)))
		if err := WriteImplicitField(ctx, "packetSize", packetSize, WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'packetSize' field")
		}

		if err := WriteSimpleField[uint16](ctx, "sequenceCount", m.GetSequenceCount(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'sequenceCount' field")
		}

		if err := WriteSimpleField[CipService](ctx, "service", m.GetService(), WriteComplex[CipService](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'service' field")
		}

		if popErr := writeBuffer.PopContext("ConnectedDataItem"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectedDataItem")
		}
		return nil
	}
	return m.TypeIdContract.(*_TypeId).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConnectedDataItem) IsConnectedDataItem() {}

func (m *_ConnectedDataItem) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ConnectedDataItem) deepCopy() *_ConnectedDataItem {
	if m == nil {
		return nil
	}
	_ConnectedDataItemCopy := &_ConnectedDataItem{
		m.TypeIdContract.(*_TypeId).deepCopy(),
		m.SequenceCount,
		utils.DeepCopy[CipService](m.Service),
	}
	_ConnectedDataItemCopy.TypeIdContract.(*_TypeId)._SubType = m
	return _ConnectedDataItemCopy
}

func (m *_ConnectedDataItem) String() string {
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
