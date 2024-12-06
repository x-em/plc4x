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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// AccessControlDataLockAccessPoint is the corresponding interface of AccessControlDataLockAccessPoint
type AccessControlDataLockAccessPoint interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AccessControlData
	// IsAccessControlDataLockAccessPoint is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAccessControlDataLockAccessPoint()
	// CreateBuilder creates a AccessControlDataLockAccessPointBuilder
	CreateAccessControlDataLockAccessPointBuilder() AccessControlDataLockAccessPointBuilder
}

// _AccessControlDataLockAccessPoint is the data-structure of this message
type _AccessControlDataLockAccessPoint struct {
	AccessControlDataContract
}

var _ AccessControlDataLockAccessPoint = (*_AccessControlDataLockAccessPoint)(nil)
var _ AccessControlDataRequirements = (*_AccessControlDataLockAccessPoint)(nil)

// NewAccessControlDataLockAccessPoint factory function for _AccessControlDataLockAccessPoint
func NewAccessControlDataLockAccessPoint(commandTypeContainer AccessControlCommandTypeContainer, networkId byte, accessPointId byte) *_AccessControlDataLockAccessPoint {
	_result := &_AccessControlDataLockAccessPoint{
		AccessControlDataContract: NewAccessControlData(commandTypeContainer, networkId, accessPointId),
	}
	_result.AccessControlDataContract.(*_AccessControlData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AccessControlDataLockAccessPointBuilder is a builder for AccessControlDataLockAccessPoint
type AccessControlDataLockAccessPointBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() AccessControlDataLockAccessPointBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() AccessControlDataBuilder
	// Build builds the AccessControlDataLockAccessPoint or returns an error if something is wrong
	Build() (AccessControlDataLockAccessPoint, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AccessControlDataLockAccessPoint
}

// NewAccessControlDataLockAccessPointBuilder() creates a AccessControlDataLockAccessPointBuilder
func NewAccessControlDataLockAccessPointBuilder() AccessControlDataLockAccessPointBuilder {
	return &_AccessControlDataLockAccessPointBuilder{_AccessControlDataLockAccessPoint: new(_AccessControlDataLockAccessPoint)}
}

type _AccessControlDataLockAccessPointBuilder struct {
	*_AccessControlDataLockAccessPoint

	parentBuilder *_AccessControlDataBuilder

	err *utils.MultiError
}

var _ (AccessControlDataLockAccessPointBuilder) = (*_AccessControlDataLockAccessPointBuilder)(nil)

func (b *_AccessControlDataLockAccessPointBuilder) setParent(contract AccessControlDataContract) {
	b.AccessControlDataContract = contract
	contract.(*_AccessControlData)._SubType = b._AccessControlDataLockAccessPoint
}

func (b *_AccessControlDataLockAccessPointBuilder) WithMandatoryFields() AccessControlDataLockAccessPointBuilder {
	return b
}

func (b *_AccessControlDataLockAccessPointBuilder) Build() (AccessControlDataLockAccessPoint, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AccessControlDataLockAccessPoint.deepCopy(), nil
}

func (b *_AccessControlDataLockAccessPointBuilder) MustBuild() AccessControlDataLockAccessPoint {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AccessControlDataLockAccessPointBuilder) Done() AccessControlDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAccessControlDataBuilder().(*_AccessControlDataBuilder)
	}
	return b.parentBuilder
}

func (b *_AccessControlDataLockAccessPointBuilder) buildForAccessControlData() (AccessControlData, error) {
	return b.Build()
}

func (b *_AccessControlDataLockAccessPointBuilder) DeepCopy() any {
	_copy := b.CreateAccessControlDataLockAccessPointBuilder().(*_AccessControlDataLockAccessPointBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAccessControlDataLockAccessPointBuilder creates a AccessControlDataLockAccessPointBuilder
func (b *_AccessControlDataLockAccessPoint) CreateAccessControlDataLockAccessPointBuilder() AccessControlDataLockAccessPointBuilder {
	if b == nil {
		return NewAccessControlDataLockAccessPointBuilder()
	}
	return &_AccessControlDataLockAccessPointBuilder{_AccessControlDataLockAccessPoint: b.deepCopy()}
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

func (m *_AccessControlDataLockAccessPoint) GetParent() AccessControlDataContract {
	return m.AccessControlDataContract
}

// Deprecated: use the interface for direct cast
func CastAccessControlDataLockAccessPoint(structType any) AccessControlDataLockAccessPoint {
	if casted, ok := structType.(AccessControlDataLockAccessPoint); ok {
		return casted
	}
	if casted, ok := structType.(*AccessControlDataLockAccessPoint); ok {
		return *casted
	}
	return nil
}

func (m *_AccessControlDataLockAccessPoint) GetTypeName() string {
	return "AccessControlDataLockAccessPoint"
}

func (m *_AccessControlDataLockAccessPoint) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AccessControlDataContract.(*_AccessControlData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_AccessControlDataLockAccessPoint) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AccessControlDataLockAccessPoint) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AccessControlData) (__accessControlDataLockAccessPoint AccessControlDataLockAccessPoint, err error) {
	m.AccessControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AccessControlDataLockAccessPoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AccessControlDataLockAccessPoint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AccessControlDataLockAccessPoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AccessControlDataLockAccessPoint")
	}

	return m, nil
}

func (m *_AccessControlDataLockAccessPoint) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AccessControlDataLockAccessPoint) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AccessControlDataLockAccessPoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AccessControlDataLockAccessPoint")
		}

		if popErr := writeBuffer.PopContext("AccessControlDataLockAccessPoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AccessControlDataLockAccessPoint")
		}
		return nil
	}
	return m.AccessControlDataContract.(*_AccessControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AccessControlDataLockAccessPoint) IsAccessControlDataLockAccessPoint() {}

func (m *_AccessControlDataLockAccessPoint) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AccessControlDataLockAccessPoint) deepCopy() *_AccessControlDataLockAccessPoint {
	if m == nil {
		return nil
	}
	_AccessControlDataLockAccessPointCopy := &_AccessControlDataLockAccessPoint{
		m.AccessControlDataContract.(*_AccessControlData).deepCopy(),
	}
	_AccessControlDataLockAccessPointCopy.AccessControlDataContract.(*_AccessControlData)._SubType = m
	return _AccessControlDataLockAccessPointCopy
}

func (m *_AccessControlDataLockAccessPoint) String() string {
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