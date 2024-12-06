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

// SecurityDataSystemDisarmed is the corresponding interface of SecurityDataSystemDisarmed
type SecurityDataSystemDisarmed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// IsSecurityDataSystemDisarmed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataSystemDisarmed()
	// CreateBuilder creates a SecurityDataSystemDisarmedBuilder
	CreateSecurityDataSystemDisarmedBuilder() SecurityDataSystemDisarmedBuilder
}

// _SecurityDataSystemDisarmed is the data-structure of this message
type _SecurityDataSystemDisarmed struct {
	SecurityDataContract
}

var _ SecurityDataSystemDisarmed = (*_SecurityDataSystemDisarmed)(nil)
var _ SecurityDataRequirements = (*_SecurityDataSystemDisarmed)(nil)

// NewSecurityDataSystemDisarmed factory function for _SecurityDataSystemDisarmed
func NewSecurityDataSystemDisarmed(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataSystemDisarmed {
	_result := &_SecurityDataSystemDisarmed{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataSystemDisarmedBuilder is a builder for SecurityDataSystemDisarmed
type SecurityDataSystemDisarmedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SecurityDataSystemDisarmedBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() SecurityDataBuilder
	// Build builds the SecurityDataSystemDisarmed or returns an error if something is wrong
	Build() (SecurityDataSystemDisarmed, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataSystemDisarmed
}

// NewSecurityDataSystemDisarmedBuilder() creates a SecurityDataSystemDisarmedBuilder
func NewSecurityDataSystemDisarmedBuilder() SecurityDataSystemDisarmedBuilder {
	return &_SecurityDataSystemDisarmedBuilder{_SecurityDataSystemDisarmed: new(_SecurityDataSystemDisarmed)}
}

type _SecurityDataSystemDisarmedBuilder struct {
	*_SecurityDataSystemDisarmed

	parentBuilder *_SecurityDataBuilder

	err *utils.MultiError
}

var _ (SecurityDataSystemDisarmedBuilder) = (*_SecurityDataSystemDisarmedBuilder)(nil)

func (b *_SecurityDataSystemDisarmedBuilder) setParent(contract SecurityDataContract) {
	b.SecurityDataContract = contract
	contract.(*_SecurityData)._SubType = b._SecurityDataSystemDisarmed
}

func (b *_SecurityDataSystemDisarmedBuilder) WithMandatoryFields() SecurityDataSystemDisarmedBuilder {
	return b
}

func (b *_SecurityDataSystemDisarmedBuilder) Build() (SecurityDataSystemDisarmed, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SecurityDataSystemDisarmed.deepCopy(), nil
}

func (b *_SecurityDataSystemDisarmedBuilder) MustBuild() SecurityDataSystemDisarmed {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SecurityDataSystemDisarmedBuilder) Done() SecurityDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewSecurityDataBuilder().(*_SecurityDataBuilder)
	}
	return b.parentBuilder
}

func (b *_SecurityDataSystemDisarmedBuilder) buildForSecurityData() (SecurityData, error) {
	return b.Build()
}

func (b *_SecurityDataSystemDisarmedBuilder) DeepCopy() any {
	_copy := b.CreateSecurityDataSystemDisarmedBuilder().(*_SecurityDataSystemDisarmedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSecurityDataSystemDisarmedBuilder creates a SecurityDataSystemDisarmedBuilder
func (b *_SecurityDataSystemDisarmed) CreateSecurityDataSystemDisarmedBuilder() SecurityDataSystemDisarmedBuilder {
	if b == nil {
		return NewSecurityDataSystemDisarmedBuilder()
	}
	return &_SecurityDataSystemDisarmedBuilder{_SecurityDataSystemDisarmed: b.deepCopy()}
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

func (m *_SecurityDataSystemDisarmed) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// Deprecated: use the interface for direct cast
func CastSecurityDataSystemDisarmed(structType any) SecurityDataSystemDisarmed {
	if casted, ok := structType.(SecurityDataSystemDisarmed); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataSystemDisarmed); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataSystemDisarmed) GetTypeName() string {
	return "SecurityDataSystemDisarmed"
}

func (m *_SecurityDataSystemDisarmed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataSystemDisarmed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataSystemDisarmed) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataSystemDisarmed SecurityDataSystemDisarmed, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataSystemDisarmed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataSystemDisarmed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataSystemDisarmed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataSystemDisarmed")
	}

	return m, nil
}

func (m *_SecurityDataSystemDisarmed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataSystemDisarmed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataSystemDisarmed"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataSystemDisarmed")
		}

		if popErr := writeBuffer.PopContext("SecurityDataSystemDisarmed"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataSystemDisarmed")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataSystemDisarmed) IsSecurityDataSystemDisarmed() {}

func (m *_SecurityDataSystemDisarmed) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataSystemDisarmed) deepCopy() *_SecurityDataSystemDisarmed {
	if m == nil {
		return nil
	}
	_SecurityDataSystemDisarmedCopy := &_SecurityDataSystemDisarmed{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
	}
	_SecurityDataSystemDisarmedCopy.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataSystemDisarmedCopy
}

func (m *_SecurityDataSystemDisarmed) String() string {
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