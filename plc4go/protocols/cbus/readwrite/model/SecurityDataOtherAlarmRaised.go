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

// SecurityDataOtherAlarmRaised is the corresponding interface of SecurityDataOtherAlarmRaised
type SecurityDataOtherAlarmRaised interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// IsSecurityDataOtherAlarmRaised is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataOtherAlarmRaised()
	// CreateBuilder creates a SecurityDataOtherAlarmRaisedBuilder
	CreateSecurityDataOtherAlarmRaisedBuilder() SecurityDataOtherAlarmRaisedBuilder
}

// _SecurityDataOtherAlarmRaised is the data-structure of this message
type _SecurityDataOtherAlarmRaised struct {
	SecurityDataContract
}

var _ SecurityDataOtherAlarmRaised = (*_SecurityDataOtherAlarmRaised)(nil)
var _ SecurityDataRequirements = (*_SecurityDataOtherAlarmRaised)(nil)

// NewSecurityDataOtherAlarmRaised factory function for _SecurityDataOtherAlarmRaised
func NewSecurityDataOtherAlarmRaised(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataOtherAlarmRaised {
	_result := &_SecurityDataOtherAlarmRaised{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataOtherAlarmRaisedBuilder is a builder for SecurityDataOtherAlarmRaised
type SecurityDataOtherAlarmRaisedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SecurityDataOtherAlarmRaisedBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() SecurityDataBuilder
	// Build builds the SecurityDataOtherAlarmRaised or returns an error if something is wrong
	Build() (SecurityDataOtherAlarmRaised, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataOtherAlarmRaised
}

// NewSecurityDataOtherAlarmRaisedBuilder() creates a SecurityDataOtherAlarmRaisedBuilder
func NewSecurityDataOtherAlarmRaisedBuilder() SecurityDataOtherAlarmRaisedBuilder {
	return &_SecurityDataOtherAlarmRaisedBuilder{_SecurityDataOtherAlarmRaised: new(_SecurityDataOtherAlarmRaised)}
}

type _SecurityDataOtherAlarmRaisedBuilder struct {
	*_SecurityDataOtherAlarmRaised

	parentBuilder *_SecurityDataBuilder

	err *utils.MultiError
}

var _ (SecurityDataOtherAlarmRaisedBuilder) = (*_SecurityDataOtherAlarmRaisedBuilder)(nil)

func (b *_SecurityDataOtherAlarmRaisedBuilder) setParent(contract SecurityDataContract) {
	b.SecurityDataContract = contract
	contract.(*_SecurityData)._SubType = b._SecurityDataOtherAlarmRaised
}

func (b *_SecurityDataOtherAlarmRaisedBuilder) WithMandatoryFields() SecurityDataOtherAlarmRaisedBuilder {
	return b
}

func (b *_SecurityDataOtherAlarmRaisedBuilder) Build() (SecurityDataOtherAlarmRaised, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SecurityDataOtherAlarmRaised.deepCopy(), nil
}

func (b *_SecurityDataOtherAlarmRaisedBuilder) MustBuild() SecurityDataOtherAlarmRaised {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SecurityDataOtherAlarmRaisedBuilder) Done() SecurityDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewSecurityDataBuilder().(*_SecurityDataBuilder)
	}
	return b.parentBuilder
}

func (b *_SecurityDataOtherAlarmRaisedBuilder) buildForSecurityData() (SecurityData, error) {
	return b.Build()
}

func (b *_SecurityDataOtherAlarmRaisedBuilder) DeepCopy() any {
	_copy := b.CreateSecurityDataOtherAlarmRaisedBuilder().(*_SecurityDataOtherAlarmRaisedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSecurityDataOtherAlarmRaisedBuilder creates a SecurityDataOtherAlarmRaisedBuilder
func (b *_SecurityDataOtherAlarmRaised) CreateSecurityDataOtherAlarmRaisedBuilder() SecurityDataOtherAlarmRaisedBuilder {
	if b == nil {
		return NewSecurityDataOtherAlarmRaisedBuilder()
	}
	return &_SecurityDataOtherAlarmRaisedBuilder{_SecurityDataOtherAlarmRaised: b.deepCopy()}
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

func (m *_SecurityDataOtherAlarmRaised) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// Deprecated: use the interface for direct cast
func CastSecurityDataOtherAlarmRaised(structType any) SecurityDataOtherAlarmRaised {
	if casted, ok := structType.(SecurityDataOtherAlarmRaised); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataOtherAlarmRaised); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataOtherAlarmRaised) GetTypeName() string {
	return "SecurityDataOtherAlarmRaised"
}

func (m *_SecurityDataOtherAlarmRaised) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataOtherAlarmRaised) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataOtherAlarmRaised) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataOtherAlarmRaised SecurityDataOtherAlarmRaised, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataOtherAlarmRaised"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataOtherAlarmRaised")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataOtherAlarmRaised"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataOtherAlarmRaised")
	}

	return m, nil
}

func (m *_SecurityDataOtherAlarmRaised) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataOtherAlarmRaised) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataOtherAlarmRaised"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataOtherAlarmRaised")
		}

		if popErr := writeBuffer.PopContext("SecurityDataOtherAlarmRaised"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataOtherAlarmRaised")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataOtherAlarmRaised) IsSecurityDataOtherAlarmRaised() {}

func (m *_SecurityDataOtherAlarmRaised) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataOtherAlarmRaised) deepCopy() *_SecurityDataOtherAlarmRaised {
	if m == nil {
		return nil
	}
	_SecurityDataOtherAlarmRaisedCopy := &_SecurityDataOtherAlarmRaised{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
	}
	_SecurityDataOtherAlarmRaisedCopy.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataOtherAlarmRaisedCopy
}

func (m *_SecurityDataOtherAlarmRaised) String() string {
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