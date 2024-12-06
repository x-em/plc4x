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

// BACnetPropertyStatesAction is the corresponding interface of BACnetPropertyStatesAction
type BACnetPropertyStatesAction interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetAction returns Action (property field)
	GetAction() BACnetActionTagged
	// IsBACnetPropertyStatesAction is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesAction()
	// CreateBuilder creates a BACnetPropertyStatesActionBuilder
	CreateBACnetPropertyStatesActionBuilder() BACnetPropertyStatesActionBuilder
}

// _BACnetPropertyStatesAction is the data-structure of this message
type _BACnetPropertyStatesAction struct {
	BACnetPropertyStatesContract
	Action BACnetActionTagged
}

var _ BACnetPropertyStatesAction = (*_BACnetPropertyStatesAction)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesAction)(nil)

// NewBACnetPropertyStatesAction factory function for _BACnetPropertyStatesAction
func NewBACnetPropertyStatesAction(peekedTagHeader BACnetTagHeader, action BACnetActionTagged) *_BACnetPropertyStatesAction {
	if action == nil {
		panic("action of type BACnetActionTagged for BACnetPropertyStatesAction must not be nil")
	}
	_result := &_BACnetPropertyStatesAction{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		Action:                       action,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesActionBuilder is a builder for BACnetPropertyStatesAction
type BACnetPropertyStatesActionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(action BACnetActionTagged) BACnetPropertyStatesActionBuilder
	// WithAction adds Action (property field)
	WithAction(BACnetActionTagged) BACnetPropertyStatesActionBuilder
	// WithActionBuilder adds Action (property field) which is build by the builder
	WithActionBuilder(func(BACnetActionTaggedBuilder) BACnetActionTaggedBuilder) BACnetPropertyStatesActionBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetPropertyStatesBuilder
	// Build builds the BACnetPropertyStatesAction or returns an error if something is wrong
	Build() (BACnetPropertyStatesAction, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesAction
}

// NewBACnetPropertyStatesActionBuilder() creates a BACnetPropertyStatesActionBuilder
func NewBACnetPropertyStatesActionBuilder() BACnetPropertyStatesActionBuilder {
	return &_BACnetPropertyStatesActionBuilder{_BACnetPropertyStatesAction: new(_BACnetPropertyStatesAction)}
}

type _BACnetPropertyStatesActionBuilder struct {
	*_BACnetPropertyStatesAction

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesActionBuilder) = (*_BACnetPropertyStatesActionBuilder)(nil)

func (b *_BACnetPropertyStatesActionBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
	contract.(*_BACnetPropertyStates)._SubType = b._BACnetPropertyStatesAction
}

func (b *_BACnetPropertyStatesActionBuilder) WithMandatoryFields(action BACnetActionTagged) BACnetPropertyStatesActionBuilder {
	return b.WithAction(action)
}

func (b *_BACnetPropertyStatesActionBuilder) WithAction(action BACnetActionTagged) BACnetPropertyStatesActionBuilder {
	b.Action = action
	return b
}

func (b *_BACnetPropertyStatesActionBuilder) WithActionBuilder(builderSupplier func(BACnetActionTaggedBuilder) BACnetActionTaggedBuilder) BACnetPropertyStatesActionBuilder {
	builder := builderSupplier(b.Action.CreateBACnetActionTaggedBuilder())
	var err error
	b.Action, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetActionTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesActionBuilder) Build() (BACnetPropertyStatesAction, error) {
	if b.Action == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'action' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesAction.deepCopy(), nil
}

func (b *_BACnetPropertyStatesActionBuilder) MustBuild() BACnetPropertyStatesAction {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPropertyStatesActionBuilder) Done() BACnetPropertyStatesBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetPropertyStatesBuilder().(*_BACnetPropertyStatesBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesActionBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesActionBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesActionBuilder().(*_BACnetPropertyStatesActionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesActionBuilder creates a BACnetPropertyStatesActionBuilder
func (b *_BACnetPropertyStatesAction) CreateBACnetPropertyStatesActionBuilder() BACnetPropertyStatesActionBuilder {
	if b == nil {
		return NewBACnetPropertyStatesActionBuilder()
	}
	return &_BACnetPropertyStatesActionBuilder{_BACnetPropertyStatesAction: b.deepCopy()}
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

func (m *_BACnetPropertyStatesAction) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesAction) GetAction() BACnetActionTagged {
	return m.Action
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesAction(structType any) BACnetPropertyStatesAction {
	if casted, ok := structType.(BACnetPropertyStatesAction); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesAction); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesAction) GetTypeName() string {
	return "BACnetPropertyStatesAction"
}

func (m *_BACnetPropertyStatesAction) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (action)
	lengthInBits += m.Action.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesAction) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesAction) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesAction BACnetPropertyStatesAction, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesAction"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesAction")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	action, err := ReadSimpleField[BACnetActionTagged](ctx, "action", ReadComplex[BACnetActionTagged](BACnetActionTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'action' field"))
	}
	m.Action = action

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesAction"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesAction")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesAction) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesAction) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesAction"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesAction")
		}

		if err := WriteSimpleField[BACnetActionTagged](ctx, "action", m.GetAction(), WriteComplex[BACnetActionTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'action' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesAction"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesAction")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesAction) IsBACnetPropertyStatesAction() {}

func (m *_BACnetPropertyStatesAction) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesAction) deepCopy() *_BACnetPropertyStatesAction {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesActionCopy := &_BACnetPropertyStatesAction{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		utils.DeepCopy[BACnetActionTagged](m.Action),
	}
	_BACnetPropertyStatesActionCopy.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesActionCopy
}

func (m *_BACnetPropertyStatesAction) String() string {
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