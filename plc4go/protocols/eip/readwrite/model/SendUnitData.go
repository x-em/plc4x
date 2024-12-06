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

// Constant values.
const SendUnitData_INTERFACEHANDLE uint32 = 0x00000000

// SendUnitData is the corresponding interface of SendUnitData
type SendUnitData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	EipPacket
	// GetTimeout returns Timeout (property field)
	GetTimeout() uint16
	// GetTypeIds returns TypeIds (property field)
	GetTypeIds() []TypeId
	// IsSendUnitData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSendUnitData()
	// CreateBuilder creates a SendUnitDataBuilder
	CreateSendUnitDataBuilder() SendUnitDataBuilder
}

// _SendUnitData is the data-structure of this message
type _SendUnitData struct {
	EipPacketContract
	Timeout uint16
	TypeIds []TypeId
}

var _ SendUnitData = (*_SendUnitData)(nil)
var _ EipPacketRequirements = (*_SendUnitData)(nil)

// NewSendUnitData factory function for _SendUnitData
func NewSendUnitData(sessionHandle uint32, status uint32, senderContext []byte, options uint32, timeout uint16, typeIds []TypeId) *_SendUnitData {
	_result := &_SendUnitData{
		EipPacketContract: NewEipPacket(sessionHandle, status, senderContext, options),
		Timeout:           timeout,
		TypeIds:           typeIds,
	}
	_result.EipPacketContract.(*_EipPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SendUnitDataBuilder is a builder for SendUnitData
type SendUnitDataBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(timeout uint16, typeIds []TypeId) SendUnitDataBuilder
	// WithTimeout adds Timeout (property field)
	WithTimeout(uint16) SendUnitDataBuilder
	// WithTypeIds adds TypeIds (property field)
	WithTypeIds(...TypeId) SendUnitDataBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() EipPacketBuilder
	// Build builds the SendUnitData or returns an error if something is wrong
	Build() (SendUnitData, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SendUnitData
}

// NewSendUnitDataBuilder() creates a SendUnitDataBuilder
func NewSendUnitDataBuilder() SendUnitDataBuilder {
	return &_SendUnitDataBuilder{_SendUnitData: new(_SendUnitData)}
}

type _SendUnitDataBuilder struct {
	*_SendUnitData

	parentBuilder *_EipPacketBuilder

	err *utils.MultiError
}

var _ (SendUnitDataBuilder) = (*_SendUnitDataBuilder)(nil)

func (b *_SendUnitDataBuilder) setParent(contract EipPacketContract) {
	b.EipPacketContract = contract
	contract.(*_EipPacket)._SubType = b._SendUnitData
}

func (b *_SendUnitDataBuilder) WithMandatoryFields(timeout uint16, typeIds []TypeId) SendUnitDataBuilder {
	return b.WithTimeout(timeout).WithTypeIds(typeIds...)
}

func (b *_SendUnitDataBuilder) WithTimeout(timeout uint16) SendUnitDataBuilder {
	b.Timeout = timeout
	return b
}

func (b *_SendUnitDataBuilder) WithTypeIds(typeIds ...TypeId) SendUnitDataBuilder {
	b.TypeIds = typeIds
	return b
}

func (b *_SendUnitDataBuilder) Build() (SendUnitData, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SendUnitData.deepCopy(), nil
}

func (b *_SendUnitDataBuilder) MustBuild() SendUnitData {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SendUnitDataBuilder) Done() EipPacketBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewEipPacketBuilder().(*_EipPacketBuilder)
	}
	return b.parentBuilder
}

func (b *_SendUnitDataBuilder) buildForEipPacket() (EipPacket, error) {
	return b.Build()
}

func (b *_SendUnitDataBuilder) DeepCopy() any {
	_copy := b.CreateSendUnitDataBuilder().(*_SendUnitDataBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSendUnitDataBuilder creates a SendUnitDataBuilder
func (b *_SendUnitData) CreateSendUnitDataBuilder() SendUnitDataBuilder {
	if b == nil {
		return NewSendUnitDataBuilder()
	}
	return &_SendUnitDataBuilder{_SendUnitData: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SendUnitData) GetCommand() uint16 {
	return 0x0070
}

func (m *_SendUnitData) GetResponse() bool {
	return false
}

func (m *_SendUnitData) GetPacketLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SendUnitData) GetParent() EipPacketContract {
	return m.EipPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SendUnitData) GetTimeout() uint16 {
	return m.Timeout
}

func (m *_SendUnitData) GetTypeIds() []TypeId {
	return m.TypeIds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_SendUnitData) GetInterfaceHandle() uint32 {
	return SendUnitData_INTERFACEHANDLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSendUnitData(structType any) SendUnitData {
	if casted, ok := structType.(SendUnitData); ok {
		return casted
	}
	if casted, ok := structType.(*SendUnitData); ok {
		return *casted
	}
	return nil
}

func (m *_SendUnitData) GetTypeName() string {
	return "SendUnitData"
}

func (m *_SendUnitData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.EipPacketContract.(*_EipPacket).getLengthInBits(ctx))

	// Const Field (interfaceHandle)
	lengthInBits += 32

	// Simple field (timeout)
	lengthInBits += 16

	// Implicit Field (typeIdCount)
	lengthInBits += 16

	// Array field
	if len(m.TypeIds) > 0 {
		for _curItem, element := range m.TypeIds {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.TypeIds), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_SendUnitData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SendUnitData) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_EipPacket, response bool) (__sendUnitData SendUnitData, err error) {
	m.EipPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SendUnitData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SendUnitData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	interfaceHandle, err := ReadConstField[uint32](ctx, "interfaceHandle", ReadUnsignedInt(readBuffer, uint8(32)), SendUnitData_INTERFACEHANDLE)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'interfaceHandle' field"))
	}
	_ = interfaceHandle

	timeout, err := ReadSimpleField(ctx, "timeout", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeout' field"))
	}
	m.Timeout = timeout

	typeIdCount, err := ReadImplicitField[uint16](ctx, "typeIdCount", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'typeIdCount' field"))
	}
	_ = typeIdCount

	typeIds, err := ReadCountArrayField[TypeId](ctx, "typeIds", ReadComplex[TypeId](TypeIdParseWithBuffer, readBuffer), uint64(typeIdCount))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'typeIds' field"))
	}
	m.TypeIds = typeIds

	if closeErr := readBuffer.CloseContext("SendUnitData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SendUnitData")
	}

	return m, nil
}

func (m *_SendUnitData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SendUnitData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SendUnitData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SendUnitData")
		}

		if err := WriteConstField(ctx, "interfaceHandle", SendUnitData_INTERFACEHANDLE, WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'interfaceHandle' field")
		}

		if err := WriteSimpleField[uint16](ctx, "timeout", m.GetTimeout(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeout' field")
		}
		typeIdCount := uint16(uint16(len(m.GetTypeIds())))
		if err := WriteImplicitField(ctx, "typeIdCount", typeIdCount, WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'typeIdCount' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "typeIds", m.GetTypeIds(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'typeIds' field")
		}

		if popErr := writeBuffer.PopContext("SendUnitData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SendUnitData")
		}
		return nil
	}
	return m.EipPacketContract.(*_EipPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SendUnitData) IsSendUnitData() {}

func (m *_SendUnitData) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SendUnitData) deepCopy() *_SendUnitData {
	if m == nil {
		return nil
	}
	_SendUnitDataCopy := &_SendUnitData{
		m.EipPacketContract.(*_EipPacket).deepCopy(),
		m.Timeout,
		utils.DeepCopySlice[TypeId, TypeId](m.TypeIds),
	}
	_SendUnitDataCopy.EipPacketContract.(*_EipPacket)._SubType = m
	return _SendUnitDataCopy
}

func (m *_SendUnitData) String() string {
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