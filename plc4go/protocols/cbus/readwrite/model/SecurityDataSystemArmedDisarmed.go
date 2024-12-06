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

// SecurityDataSystemArmedDisarmed is the corresponding interface of SecurityDataSystemArmedDisarmed
type SecurityDataSystemArmedDisarmed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// GetArmCodeType returns ArmCodeType (property field)
	GetArmCodeType() SecurityArmCode
	// IsSecurityDataSystemArmedDisarmed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataSystemArmedDisarmed()
	// CreateBuilder creates a SecurityDataSystemArmedDisarmedBuilder
	CreateSecurityDataSystemArmedDisarmedBuilder() SecurityDataSystemArmedDisarmedBuilder
}

// _SecurityDataSystemArmedDisarmed is the data-structure of this message
type _SecurityDataSystemArmedDisarmed struct {
	SecurityDataContract
	ArmCodeType SecurityArmCode
}

var _ SecurityDataSystemArmedDisarmed = (*_SecurityDataSystemArmedDisarmed)(nil)
var _ SecurityDataRequirements = (*_SecurityDataSystemArmedDisarmed)(nil)

// NewSecurityDataSystemArmedDisarmed factory function for _SecurityDataSystemArmedDisarmed
func NewSecurityDataSystemArmedDisarmed(commandTypeContainer SecurityCommandTypeContainer, argument byte, armCodeType SecurityArmCode) *_SecurityDataSystemArmedDisarmed {
	if armCodeType == nil {
		panic("armCodeType of type SecurityArmCode for SecurityDataSystemArmedDisarmed must not be nil")
	}
	_result := &_SecurityDataSystemArmedDisarmed{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
		ArmCodeType:          armCodeType,
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataSystemArmedDisarmedBuilder is a builder for SecurityDataSystemArmedDisarmed
type SecurityDataSystemArmedDisarmedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(armCodeType SecurityArmCode) SecurityDataSystemArmedDisarmedBuilder
	// WithArmCodeType adds ArmCodeType (property field)
	WithArmCodeType(SecurityArmCode) SecurityDataSystemArmedDisarmedBuilder
	// WithArmCodeTypeBuilder adds ArmCodeType (property field) which is build by the builder
	WithArmCodeTypeBuilder(func(SecurityArmCodeBuilder) SecurityArmCodeBuilder) SecurityDataSystemArmedDisarmedBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() SecurityDataBuilder
	// Build builds the SecurityDataSystemArmedDisarmed or returns an error if something is wrong
	Build() (SecurityDataSystemArmedDisarmed, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataSystemArmedDisarmed
}

// NewSecurityDataSystemArmedDisarmedBuilder() creates a SecurityDataSystemArmedDisarmedBuilder
func NewSecurityDataSystemArmedDisarmedBuilder() SecurityDataSystemArmedDisarmedBuilder {
	return &_SecurityDataSystemArmedDisarmedBuilder{_SecurityDataSystemArmedDisarmed: new(_SecurityDataSystemArmedDisarmed)}
}

type _SecurityDataSystemArmedDisarmedBuilder struct {
	*_SecurityDataSystemArmedDisarmed

	parentBuilder *_SecurityDataBuilder

	err *utils.MultiError
}

var _ (SecurityDataSystemArmedDisarmedBuilder) = (*_SecurityDataSystemArmedDisarmedBuilder)(nil)

func (b *_SecurityDataSystemArmedDisarmedBuilder) setParent(contract SecurityDataContract) {
	b.SecurityDataContract = contract
	contract.(*_SecurityData)._SubType = b._SecurityDataSystemArmedDisarmed
}

func (b *_SecurityDataSystemArmedDisarmedBuilder) WithMandatoryFields(armCodeType SecurityArmCode) SecurityDataSystemArmedDisarmedBuilder {
	return b.WithArmCodeType(armCodeType)
}

func (b *_SecurityDataSystemArmedDisarmedBuilder) WithArmCodeType(armCodeType SecurityArmCode) SecurityDataSystemArmedDisarmedBuilder {
	b.ArmCodeType = armCodeType
	return b
}

func (b *_SecurityDataSystemArmedDisarmedBuilder) WithArmCodeTypeBuilder(builderSupplier func(SecurityArmCodeBuilder) SecurityArmCodeBuilder) SecurityDataSystemArmedDisarmedBuilder {
	builder := builderSupplier(b.ArmCodeType.CreateSecurityArmCodeBuilder())
	var err error
	b.ArmCodeType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "SecurityArmCodeBuilder failed"))
	}
	return b
}

func (b *_SecurityDataSystemArmedDisarmedBuilder) Build() (SecurityDataSystemArmedDisarmed, error) {
	if b.ArmCodeType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'armCodeType' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SecurityDataSystemArmedDisarmed.deepCopy(), nil
}

func (b *_SecurityDataSystemArmedDisarmedBuilder) MustBuild() SecurityDataSystemArmedDisarmed {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SecurityDataSystemArmedDisarmedBuilder) Done() SecurityDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewSecurityDataBuilder().(*_SecurityDataBuilder)
	}
	return b.parentBuilder
}

func (b *_SecurityDataSystemArmedDisarmedBuilder) buildForSecurityData() (SecurityData, error) {
	return b.Build()
}

func (b *_SecurityDataSystemArmedDisarmedBuilder) DeepCopy() any {
	_copy := b.CreateSecurityDataSystemArmedDisarmedBuilder().(*_SecurityDataSystemArmedDisarmedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSecurityDataSystemArmedDisarmedBuilder creates a SecurityDataSystemArmedDisarmedBuilder
func (b *_SecurityDataSystemArmedDisarmed) CreateSecurityDataSystemArmedDisarmedBuilder() SecurityDataSystemArmedDisarmedBuilder {
	if b == nil {
		return NewSecurityDataSystemArmedDisarmedBuilder()
	}
	return &_SecurityDataSystemArmedDisarmedBuilder{_SecurityDataSystemArmedDisarmed: b.deepCopy()}
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

func (m *_SecurityDataSystemArmedDisarmed) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataSystemArmedDisarmed) GetArmCodeType() SecurityArmCode {
	return m.ArmCodeType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSecurityDataSystemArmedDisarmed(structType any) SecurityDataSystemArmedDisarmed {
	if casted, ok := structType.(SecurityDataSystemArmedDisarmed); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataSystemArmedDisarmed); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataSystemArmedDisarmed) GetTypeName() string {
	return "SecurityDataSystemArmedDisarmed"
}

func (m *_SecurityDataSystemArmedDisarmed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	// Simple field (armCodeType)
	lengthInBits += m.ArmCodeType.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SecurityDataSystemArmedDisarmed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataSystemArmedDisarmed) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataSystemArmedDisarmed SecurityDataSystemArmedDisarmed, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataSystemArmedDisarmed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataSystemArmedDisarmed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	armCodeType, err := ReadSimpleField[SecurityArmCode](ctx, "armCodeType", ReadComplex[SecurityArmCode](SecurityArmCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'armCodeType' field"))
	}
	m.ArmCodeType = armCodeType

	if closeErr := readBuffer.CloseContext("SecurityDataSystemArmedDisarmed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataSystemArmedDisarmed")
	}

	return m, nil
}

func (m *_SecurityDataSystemArmedDisarmed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataSystemArmedDisarmed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataSystemArmedDisarmed"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataSystemArmedDisarmed")
		}

		if err := WriteSimpleField[SecurityArmCode](ctx, "armCodeType", m.GetArmCodeType(), WriteComplex[SecurityArmCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'armCodeType' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataSystemArmedDisarmed"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataSystemArmedDisarmed")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataSystemArmedDisarmed) IsSecurityDataSystemArmedDisarmed() {}

func (m *_SecurityDataSystemArmedDisarmed) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataSystemArmedDisarmed) deepCopy() *_SecurityDataSystemArmedDisarmed {
	if m == nil {
		return nil
	}
	_SecurityDataSystemArmedDisarmedCopy := &_SecurityDataSystemArmedDisarmed{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
		utils.DeepCopy[SecurityArmCode](m.ArmCodeType),
	}
	_SecurityDataSystemArmedDisarmedCopy.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataSystemArmedDisarmedCopy
}

func (m *_SecurityDataSystemArmedDisarmed) String() string {
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
