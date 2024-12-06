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

// KnxNetIpDeviceManagement is the corresponding interface of KnxNetIpDeviceManagement
type KnxNetIpDeviceManagement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ServiceId
	// GetVersion returns Version (property field)
	GetVersion() uint8
	// IsKnxNetIpDeviceManagement is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsKnxNetIpDeviceManagement()
	// CreateBuilder creates a KnxNetIpDeviceManagementBuilder
	CreateKnxNetIpDeviceManagementBuilder() KnxNetIpDeviceManagementBuilder
}

// _KnxNetIpDeviceManagement is the data-structure of this message
type _KnxNetIpDeviceManagement struct {
	ServiceIdContract
	Version uint8
}

var _ KnxNetIpDeviceManagement = (*_KnxNetIpDeviceManagement)(nil)
var _ ServiceIdRequirements = (*_KnxNetIpDeviceManagement)(nil)

// NewKnxNetIpDeviceManagement factory function for _KnxNetIpDeviceManagement
func NewKnxNetIpDeviceManagement(version uint8) *_KnxNetIpDeviceManagement {
	_result := &_KnxNetIpDeviceManagement{
		ServiceIdContract: NewServiceId(),
		Version:           version,
	}
	_result.ServiceIdContract.(*_ServiceId)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// KnxNetIpDeviceManagementBuilder is a builder for KnxNetIpDeviceManagement
type KnxNetIpDeviceManagementBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(version uint8) KnxNetIpDeviceManagementBuilder
	// WithVersion adds Version (property field)
	WithVersion(uint8) KnxNetIpDeviceManagementBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ServiceIdBuilder
	// Build builds the KnxNetIpDeviceManagement or returns an error if something is wrong
	Build() (KnxNetIpDeviceManagement, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() KnxNetIpDeviceManagement
}

// NewKnxNetIpDeviceManagementBuilder() creates a KnxNetIpDeviceManagementBuilder
func NewKnxNetIpDeviceManagementBuilder() KnxNetIpDeviceManagementBuilder {
	return &_KnxNetIpDeviceManagementBuilder{_KnxNetIpDeviceManagement: new(_KnxNetIpDeviceManagement)}
}

type _KnxNetIpDeviceManagementBuilder struct {
	*_KnxNetIpDeviceManagement

	parentBuilder *_ServiceIdBuilder

	err *utils.MultiError
}

var _ (KnxNetIpDeviceManagementBuilder) = (*_KnxNetIpDeviceManagementBuilder)(nil)

func (b *_KnxNetIpDeviceManagementBuilder) setParent(contract ServiceIdContract) {
	b.ServiceIdContract = contract
	contract.(*_ServiceId)._SubType = b._KnxNetIpDeviceManagement
}

func (b *_KnxNetIpDeviceManagementBuilder) WithMandatoryFields(version uint8) KnxNetIpDeviceManagementBuilder {
	return b.WithVersion(version)
}

func (b *_KnxNetIpDeviceManagementBuilder) WithVersion(version uint8) KnxNetIpDeviceManagementBuilder {
	b.Version = version
	return b
}

func (b *_KnxNetIpDeviceManagementBuilder) Build() (KnxNetIpDeviceManagement, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._KnxNetIpDeviceManagement.deepCopy(), nil
}

func (b *_KnxNetIpDeviceManagementBuilder) MustBuild() KnxNetIpDeviceManagement {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_KnxNetIpDeviceManagementBuilder) Done() ServiceIdBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewServiceIdBuilder().(*_ServiceIdBuilder)
	}
	return b.parentBuilder
}

func (b *_KnxNetIpDeviceManagementBuilder) buildForServiceId() (ServiceId, error) {
	return b.Build()
}

func (b *_KnxNetIpDeviceManagementBuilder) DeepCopy() any {
	_copy := b.CreateKnxNetIpDeviceManagementBuilder().(*_KnxNetIpDeviceManagementBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateKnxNetIpDeviceManagementBuilder creates a KnxNetIpDeviceManagementBuilder
func (b *_KnxNetIpDeviceManagement) CreateKnxNetIpDeviceManagementBuilder() KnxNetIpDeviceManagementBuilder {
	if b == nil {
		return NewKnxNetIpDeviceManagementBuilder()
	}
	return &_KnxNetIpDeviceManagementBuilder{_KnxNetIpDeviceManagement: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_KnxNetIpDeviceManagement) GetServiceType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_KnxNetIpDeviceManagement) GetParent() ServiceIdContract {
	return m.ServiceIdContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KnxNetIpDeviceManagement) GetVersion() uint8 {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastKnxNetIpDeviceManagement(structType any) KnxNetIpDeviceManagement {
	if casted, ok := structType.(KnxNetIpDeviceManagement); ok {
		return casted
	}
	if casted, ok := structType.(*KnxNetIpDeviceManagement); ok {
		return *casted
	}
	return nil
}

func (m *_KnxNetIpDeviceManagement) GetTypeName() string {
	return "KnxNetIpDeviceManagement"
}

func (m *_KnxNetIpDeviceManagement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ServiceIdContract.(*_ServiceId).getLengthInBits(ctx))

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *_KnxNetIpDeviceManagement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_KnxNetIpDeviceManagement) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ServiceId) (__knxNetIpDeviceManagement KnxNetIpDeviceManagement, err error) {
	m.ServiceIdContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxNetIpDeviceManagement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxNetIpDeviceManagement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	version, err := ReadSimpleField(ctx, "version", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'version' field"))
	}
	m.Version = version

	if closeErr := readBuffer.CloseContext("KnxNetIpDeviceManagement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxNetIpDeviceManagement")
	}

	return m, nil
}

func (m *_KnxNetIpDeviceManagement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_KnxNetIpDeviceManagement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetIpDeviceManagement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KnxNetIpDeviceManagement")
		}

		if err := WriteSimpleField[uint8](ctx, "version", m.GetVersion(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetIpDeviceManagement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KnxNetIpDeviceManagement")
		}
		return nil
	}
	return m.ServiceIdContract.(*_ServiceId).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_KnxNetIpDeviceManagement) IsKnxNetIpDeviceManagement() {}

func (m *_KnxNetIpDeviceManagement) DeepCopy() any {
	return m.deepCopy()
}

func (m *_KnxNetIpDeviceManagement) deepCopy() *_KnxNetIpDeviceManagement {
	if m == nil {
		return nil
	}
	_KnxNetIpDeviceManagementCopy := &_KnxNetIpDeviceManagement{
		m.ServiceIdContract.(*_ServiceId).deepCopy(),
		m.Version,
	}
	_KnxNetIpDeviceManagementCopy.ServiceIdContract.(*_ServiceId)._SubType = m
	return _KnxNetIpDeviceManagementCopy
}

func (m *_KnxNetIpDeviceManagement) String() string {
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