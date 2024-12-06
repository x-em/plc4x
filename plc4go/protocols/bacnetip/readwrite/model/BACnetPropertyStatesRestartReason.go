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

// BACnetPropertyStatesRestartReason is the corresponding interface of BACnetPropertyStatesRestartReason
type BACnetPropertyStatesRestartReason interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetRestartReason returns RestartReason (property field)
	GetRestartReason() BACnetRestartReasonTagged
	// IsBACnetPropertyStatesRestartReason is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesRestartReason()
	// CreateBuilder creates a BACnetPropertyStatesRestartReasonBuilder
	CreateBACnetPropertyStatesRestartReasonBuilder() BACnetPropertyStatesRestartReasonBuilder
}

// _BACnetPropertyStatesRestartReason is the data-structure of this message
type _BACnetPropertyStatesRestartReason struct {
	BACnetPropertyStatesContract
	RestartReason BACnetRestartReasonTagged
}

var _ BACnetPropertyStatesRestartReason = (*_BACnetPropertyStatesRestartReason)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesRestartReason)(nil)

// NewBACnetPropertyStatesRestartReason factory function for _BACnetPropertyStatesRestartReason
func NewBACnetPropertyStatesRestartReason(peekedTagHeader BACnetTagHeader, restartReason BACnetRestartReasonTagged) *_BACnetPropertyStatesRestartReason {
	if restartReason == nil {
		panic("restartReason of type BACnetRestartReasonTagged for BACnetPropertyStatesRestartReason must not be nil")
	}
	_result := &_BACnetPropertyStatesRestartReason{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		RestartReason:                restartReason,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesRestartReasonBuilder is a builder for BACnetPropertyStatesRestartReason
type BACnetPropertyStatesRestartReasonBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(restartReason BACnetRestartReasonTagged) BACnetPropertyStatesRestartReasonBuilder
	// WithRestartReason adds RestartReason (property field)
	WithRestartReason(BACnetRestartReasonTagged) BACnetPropertyStatesRestartReasonBuilder
	// WithRestartReasonBuilder adds RestartReason (property field) which is build by the builder
	WithRestartReasonBuilder(func(BACnetRestartReasonTaggedBuilder) BACnetRestartReasonTaggedBuilder) BACnetPropertyStatesRestartReasonBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetPropertyStatesBuilder
	// Build builds the BACnetPropertyStatesRestartReason or returns an error if something is wrong
	Build() (BACnetPropertyStatesRestartReason, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesRestartReason
}

// NewBACnetPropertyStatesRestartReasonBuilder() creates a BACnetPropertyStatesRestartReasonBuilder
func NewBACnetPropertyStatesRestartReasonBuilder() BACnetPropertyStatesRestartReasonBuilder {
	return &_BACnetPropertyStatesRestartReasonBuilder{_BACnetPropertyStatesRestartReason: new(_BACnetPropertyStatesRestartReason)}
}

type _BACnetPropertyStatesRestartReasonBuilder struct {
	*_BACnetPropertyStatesRestartReason

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesRestartReasonBuilder) = (*_BACnetPropertyStatesRestartReasonBuilder)(nil)

func (b *_BACnetPropertyStatesRestartReasonBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
	contract.(*_BACnetPropertyStates)._SubType = b._BACnetPropertyStatesRestartReason
}

func (b *_BACnetPropertyStatesRestartReasonBuilder) WithMandatoryFields(restartReason BACnetRestartReasonTagged) BACnetPropertyStatesRestartReasonBuilder {
	return b.WithRestartReason(restartReason)
}

func (b *_BACnetPropertyStatesRestartReasonBuilder) WithRestartReason(restartReason BACnetRestartReasonTagged) BACnetPropertyStatesRestartReasonBuilder {
	b.RestartReason = restartReason
	return b
}

func (b *_BACnetPropertyStatesRestartReasonBuilder) WithRestartReasonBuilder(builderSupplier func(BACnetRestartReasonTaggedBuilder) BACnetRestartReasonTaggedBuilder) BACnetPropertyStatesRestartReasonBuilder {
	builder := builderSupplier(b.RestartReason.CreateBACnetRestartReasonTaggedBuilder())
	var err error
	b.RestartReason, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetRestartReasonTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesRestartReasonBuilder) Build() (BACnetPropertyStatesRestartReason, error) {
	if b.RestartReason == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'restartReason' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesRestartReason.deepCopy(), nil
}

func (b *_BACnetPropertyStatesRestartReasonBuilder) MustBuild() BACnetPropertyStatesRestartReason {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPropertyStatesRestartReasonBuilder) Done() BACnetPropertyStatesBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetPropertyStatesBuilder().(*_BACnetPropertyStatesBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesRestartReasonBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesRestartReasonBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesRestartReasonBuilder().(*_BACnetPropertyStatesRestartReasonBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesRestartReasonBuilder creates a BACnetPropertyStatesRestartReasonBuilder
func (b *_BACnetPropertyStatesRestartReason) CreateBACnetPropertyStatesRestartReasonBuilder() BACnetPropertyStatesRestartReasonBuilder {
	if b == nil {
		return NewBACnetPropertyStatesRestartReasonBuilder()
	}
	return &_BACnetPropertyStatesRestartReasonBuilder{_BACnetPropertyStatesRestartReason: b.deepCopy()}
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

func (m *_BACnetPropertyStatesRestartReason) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesRestartReason) GetRestartReason() BACnetRestartReasonTagged {
	return m.RestartReason
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesRestartReason(structType any) BACnetPropertyStatesRestartReason {
	if casted, ok := structType.(BACnetPropertyStatesRestartReason); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesRestartReason); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesRestartReason) GetTypeName() string {
	return "BACnetPropertyStatesRestartReason"
}

func (m *_BACnetPropertyStatesRestartReason) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (restartReason)
	lengthInBits += m.RestartReason.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesRestartReason) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesRestartReason) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesRestartReason BACnetPropertyStatesRestartReason, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesRestartReason"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesRestartReason")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	restartReason, err := ReadSimpleField[BACnetRestartReasonTagged](ctx, "restartReason", ReadComplex[BACnetRestartReasonTagged](BACnetRestartReasonTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'restartReason' field"))
	}
	m.RestartReason = restartReason

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesRestartReason"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesRestartReason")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesRestartReason) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesRestartReason) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesRestartReason"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesRestartReason")
		}

		if err := WriteSimpleField[BACnetRestartReasonTagged](ctx, "restartReason", m.GetRestartReason(), WriteComplex[BACnetRestartReasonTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'restartReason' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesRestartReason"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesRestartReason")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesRestartReason) IsBACnetPropertyStatesRestartReason() {}

func (m *_BACnetPropertyStatesRestartReason) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesRestartReason) deepCopy() *_BACnetPropertyStatesRestartReason {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesRestartReasonCopy := &_BACnetPropertyStatesRestartReason{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		utils.DeepCopy[BACnetRestartReasonTagged](m.RestartReason),
	}
	_BACnetPropertyStatesRestartReasonCopy.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesRestartReasonCopy
}

func (m *_BACnetPropertyStatesRestartReason) String() string {
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