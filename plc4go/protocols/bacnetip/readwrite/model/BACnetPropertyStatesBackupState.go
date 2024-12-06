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

// BACnetPropertyStatesBackupState is the corresponding interface of BACnetPropertyStatesBackupState
type BACnetPropertyStatesBackupState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetBackupState returns BackupState (property field)
	GetBackupState() BACnetBackupStateTagged
	// IsBACnetPropertyStatesBackupState is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesBackupState()
	// CreateBuilder creates a BACnetPropertyStatesBackupStateBuilder
	CreateBACnetPropertyStatesBackupStateBuilder() BACnetPropertyStatesBackupStateBuilder
}

// _BACnetPropertyStatesBackupState is the data-structure of this message
type _BACnetPropertyStatesBackupState struct {
	BACnetPropertyStatesContract
	BackupState BACnetBackupStateTagged
}

var _ BACnetPropertyStatesBackupState = (*_BACnetPropertyStatesBackupState)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesBackupState)(nil)

// NewBACnetPropertyStatesBackupState factory function for _BACnetPropertyStatesBackupState
func NewBACnetPropertyStatesBackupState(peekedTagHeader BACnetTagHeader, backupState BACnetBackupStateTagged) *_BACnetPropertyStatesBackupState {
	if backupState == nil {
		panic("backupState of type BACnetBackupStateTagged for BACnetPropertyStatesBackupState must not be nil")
	}
	_result := &_BACnetPropertyStatesBackupState{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		BackupState:                  backupState,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesBackupStateBuilder is a builder for BACnetPropertyStatesBackupState
type BACnetPropertyStatesBackupStateBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(backupState BACnetBackupStateTagged) BACnetPropertyStatesBackupStateBuilder
	// WithBackupState adds BackupState (property field)
	WithBackupState(BACnetBackupStateTagged) BACnetPropertyStatesBackupStateBuilder
	// WithBackupStateBuilder adds BackupState (property field) which is build by the builder
	WithBackupStateBuilder(func(BACnetBackupStateTaggedBuilder) BACnetBackupStateTaggedBuilder) BACnetPropertyStatesBackupStateBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetPropertyStatesBuilder
	// Build builds the BACnetPropertyStatesBackupState or returns an error if something is wrong
	Build() (BACnetPropertyStatesBackupState, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesBackupState
}

// NewBACnetPropertyStatesBackupStateBuilder() creates a BACnetPropertyStatesBackupStateBuilder
func NewBACnetPropertyStatesBackupStateBuilder() BACnetPropertyStatesBackupStateBuilder {
	return &_BACnetPropertyStatesBackupStateBuilder{_BACnetPropertyStatesBackupState: new(_BACnetPropertyStatesBackupState)}
}

type _BACnetPropertyStatesBackupStateBuilder struct {
	*_BACnetPropertyStatesBackupState

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesBackupStateBuilder) = (*_BACnetPropertyStatesBackupStateBuilder)(nil)

func (b *_BACnetPropertyStatesBackupStateBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
	contract.(*_BACnetPropertyStates)._SubType = b._BACnetPropertyStatesBackupState
}

func (b *_BACnetPropertyStatesBackupStateBuilder) WithMandatoryFields(backupState BACnetBackupStateTagged) BACnetPropertyStatesBackupStateBuilder {
	return b.WithBackupState(backupState)
}

func (b *_BACnetPropertyStatesBackupStateBuilder) WithBackupState(backupState BACnetBackupStateTagged) BACnetPropertyStatesBackupStateBuilder {
	b.BackupState = backupState
	return b
}

func (b *_BACnetPropertyStatesBackupStateBuilder) WithBackupStateBuilder(builderSupplier func(BACnetBackupStateTaggedBuilder) BACnetBackupStateTaggedBuilder) BACnetPropertyStatesBackupStateBuilder {
	builder := builderSupplier(b.BackupState.CreateBACnetBackupStateTaggedBuilder())
	var err error
	b.BackupState, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetBackupStateTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesBackupStateBuilder) Build() (BACnetPropertyStatesBackupState, error) {
	if b.BackupState == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'backupState' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesBackupState.deepCopy(), nil
}

func (b *_BACnetPropertyStatesBackupStateBuilder) MustBuild() BACnetPropertyStatesBackupState {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPropertyStatesBackupStateBuilder) Done() BACnetPropertyStatesBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetPropertyStatesBuilder().(*_BACnetPropertyStatesBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesBackupStateBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesBackupStateBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesBackupStateBuilder().(*_BACnetPropertyStatesBackupStateBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesBackupStateBuilder creates a BACnetPropertyStatesBackupStateBuilder
func (b *_BACnetPropertyStatesBackupState) CreateBACnetPropertyStatesBackupStateBuilder() BACnetPropertyStatesBackupStateBuilder {
	if b == nil {
		return NewBACnetPropertyStatesBackupStateBuilder()
	}
	return &_BACnetPropertyStatesBackupStateBuilder{_BACnetPropertyStatesBackupState: b.deepCopy()}
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

func (m *_BACnetPropertyStatesBackupState) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesBackupState) GetBackupState() BACnetBackupStateTagged {
	return m.BackupState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesBackupState(structType any) BACnetPropertyStatesBackupState {
	if casted, ok := structType.(BACnetPropertyStatesBackupState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesBackupState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesBackupState) GetTypeName() string {
	return "BACnetPropertyStatesBackupState"
}

func (m *_BACnetPropertyStatesBackupState) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (backupState)
	lengthInBits += m.BackupState.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesBackupState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesBackupState) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesBackupState BACnetPropertyStatesBackupState, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesBackupState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesBackupState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	backupState, err := ReadSimpleField[BACnetBackupStateTagged](ctx, "backupState", ReadComplex[BACnetBackupStateTagged](BACnetBackupStateTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'backupState' field"))
	}
	m.BackupState = backupState

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesBackupState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesBackupState")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesBackupState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesBackupState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesBackupState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesBackupState")
		}

		if err := WriteSimpleField[BACnetBackupStateTagged](ctx, "backupState", m.GetBackupState(), WriteComplex[BACnetBackupStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'backupState' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesBackupState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesBackupState")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesBackupState) IsBACnetPropertyStatesBackupState() {}

func (m *_BACnetPropertyStatesBackupState) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesBackupState) deepCopy() *_BACnetPropertyStatesBackupState {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesBackupStateCopy := &_BACnetPropertyStatesBackupState{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		utils.DeepCopy[BACnetBackupStateTagged](m.BackupState),
	}
	_BACnetPropertyStatesBackupStateCopy.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesBackupStateCopy
}

func (m *_BACnetPropertyStatesBackupState) String() string {
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