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

// BACnetLandingCallStatusCommandDirection is the corresponding interface of BACnetLandingCallStatusCommandDirection
type BACnetLandingCallStatusCommandDirection interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetLandingCallStatusCommand
	// GetDirection returns Direction (property field)
	GetDirection() BACnetLiftCarDirectionTagged
	// IsBACnetLandingCallStatusCommandDirection is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLandingCallStatusCommandDirection()
	// CreateBuilder creates a BACnetLandingCallStatusCommandDirectionBuilder
	CreateBACnetLandingCallStatusCommandDirectionBuilder() BACnetLandingCallStatusCommandDirectionBuilder
}

// _BACnetLandingCallStatusCommandDirection is the data-structure of this message
type _BACnetLandingCallStatusCommandDirection struct {
	BACnetLandingCallStatusCommandContract
	Direction BACnetLiftCarDirectionTagged
}

var _ BACnetLandingCallStatusCommandDirection = (*_BACnetLandingCallStatusCommandDirection)(nil)
var _ BACnetLandingCallStatusCommandRequirements = (*_BACnetLandingCallStatusCommandDirection)(nil)

// NewBACnetLandingCallStatusCommandDirection factory function for _BACnetLandingCallStatusCommandDirection
func NewBACnetLandingCallStatusCommandDirection(peekedTagHeader BACnetTagHeader, direction BACnetLiftCarDirectionTagged) *_BACnetLandingCallStatusCommandDirection {
	if direction == nil {
		panic("direction of type BACnetLiftCarDirectionTagged for BACnetLandingCallStatusCommandDirection must not be nil")
	}
	_result := &_BACnetLandingCallStatusCommandDirection{
		BACnetLandingCallStatusCommandContract: NewBACnetLandingCallStatusCommand(peekedTagHeader),
		Direction:                              direction,
	}
	_result.BACnetLandingCallStatusCommandContract.(*_BACnetLandingCallStatusCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLandingCallStatusCommandDirectionBuilder is a builder for BACnetLandingCallStatusCommandDirection
type BACnetLandingCallStatusCommandDirectionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(direction BACnetLiftCarDirectionTagged) BACnetLandingCallStatusCommandDirectionBuilder
	// WithDirection adds Direction (property field)
	WithDirection(BACnetLiftCarDirectionTagged) BACnetLandingCallStatusCommandDirectionBuilder
	// WithDirectionBuilder adds Direction (property field) which is build by the builder
	WithDirectionBuilder(func(BACnetLiftCarDirectionTaggedBuilder) BACnetLiftCarDirectionTaggedBuilder) BACnetLandingCallStatusCommandDirectionBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetLandingCallStatusCommandBuilder
	// Build builds the BACnetLandingCallStatusCommandDirection or returns an error if something is wrong
	Build() (BACnetLandingCallStatusCommandDirection, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLandingCallStatusCommandDirection
}

// NewBACnetLandingCallStatusCommandDirectionBuilder() creates a BACnetLandingCallStatusCommandDirectionBuilder
func NewBACnetLandingCallStatusCommandDirectionBuilder() BACnetLandingCallStatusCommandDirectionBuilder {
	return &_BACnetLandingCallStatusCommandDirectionBuilder{_BACnetLandingCallStatusCommandDirection: new(_BACnetLandingCallStatusCommandDirection)}
}

type _BACnetLandingCallStatusCommandDirectionBuilder struct {
	*_BACnetLandingCallStatusCommandDirection

	parentBuilder *_BACnetLandingCallStatusCommandBuilder

	err *utils.MultiError
}

var _ (BACnetLandingCallStatusCommandDirectionBuilder) = (*_BACnetLandingCallStatusCommandDirectionBuilder)(nil)

func (b *_BACnetLandingCallStatusCommandDirectionBuilder) setParent(contract BACnetLandingCallStatusCommandContract) {
	b.BACnetLandingCallStatusCommandContract = contract
	contract.(*_BACnetLandingCallStatusCommand)._SubType = b._BACnetLandingCallStatusCommandDirection
}

func (b *_BACnetLandingCallStatusCommandDirectionBuilder) WithMandatoryFields(direction BACnetLiftCarDirectionTagged) BACnetLandingCallStatusCommandDirectionBuilder {
	return b.WithDirection(direction)
}

func (b *_BACnetLandingCallStatusCommandDirectionBuilder) WithDirection(direction BACnetLiftCarDirectionTagged) BACnetLandingCallStatusCommandDirectionBuilder {
	b.Direction = direction
	return b
}

func (b *_BACnetLandingCallStatusCommandDirectionBuilder) WithDirectionBuilder(builderSupplier func(BACnetLiftCarDirectionTaggedBuilder) BACnetLiftCarDirectionTaggedBuilder) BACnetLandingCallStatusCommandDirectionBuilder {
	builder := builderSupplier(b.Direction.CreateBACnetLiftCarDirectionTaggedBuilder())
	var err error
	b.Direction, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLiftCarDirectionTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetLandingCallStatusCommandDirectionBuilder) Build() (BACnetLandingCallStatusCommandDirection, error) {
	if b.Direction == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'direction' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetLandingCallStatusCommandDirection.deepCopy(), nil
}

func (b *_BACnetLandingCallStatusCommandDirectionBuilder) MustBuild() BACnetLandingCallStatusCommandDirection {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetLandingCallStatusCommandDirectionBuilder) Done() BACnetLandingCallStatusCommandBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetLandingCallStatusCommandBuilder().(*_BACnetLandingCallStatusCommandBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetLandingCallStatusCommandDirectionBuilder) buildForBACnetLandingCallStatusCommand() (BACnetLandingCallStatusCommand, error) {
	return b.Build()
}

func (b *_BACnetLandingCallStatusCommandDirectionBuilder) DeepCopy() any {
	_copy := b.CreateBACnetLandingCallStatusCommandDirectionBuilder().(*_BACnetLandingCallStatusCommandDirectionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetLandingCallStatusCommandDirectionBuilder creates a BACnetLandingCallStatusCommandDirectionBuilder
func (b *_BACnetLandingCallStatusCommandDirection) CreateBACnetLandingCallStatusCommandDirectionBuilder() BACnetLandingCallStatusCommandDirectionBuilder {
	if b == nil {
		return NewBACnetLandingCallStatusCommandDirectionBuilder()
	}
	return &_BACnetLandingCallStatusCommandDirectionBuilder{_BACnetLandingCallStatusCommandDirection: b.deepCopy()}
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

func (m *_BACnetLandingCallStatusCommandDirection) GetParent() BACnetLandingCallStatusCommandContract {
	return m.BACnetLandingCallStatusCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLandingCallStatusCommandDirection) GetDirection() BACnetLiftCarDirectionTagged {
	return m.Direction
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLandingCallStatusCommandDirection(structType any) BACnetLandingCallStatusCommandDirection {
	if casted, ok := structType.(BACnetLandingCallStatusCommandDirection); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLandingCallStatusCommandDirection); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLandingCallStatusCommandDirection) GetTypeName() string {
	return "BACnetLandingCallStatusCommandDirection"
}

func (m *_BACnetLandingCallStatusCommandDirection) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLandingCallStatusCommandContract.(*_BACnetLandingCallStatusCommand).getLengthInBits(ctx))

	// Simple field (direction)
	lengthInBits += m.Direction.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLandingCallStatusCommandDirection) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLandingCallStatusCommandDirection) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLandingCallStatusCommand) (__bACnetLandingCallStatusCommandDirection BACnetLandingCallStatusCommandDirection, err error) {
	m.BACnetLandingCallStatusCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLandingCallStatusCommandDirection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLandingCallStatusCommandDirection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	direction, err := ReadSimpleField[BACnetLiftCarDirectionTagged](ctx, "direction", ReadComplex[BACnetLiftCarDirectionTagged](BACnetLiftCarDirectionTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'direction' field"))
	}
	m.Direction = direction

	if closeErr := readBuffer.CloseContext("BACnetLandingCallStatusCommandDirection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLandingCallStatusCommandDirection")
	}

	return m, nil
}

func (m *_BACnetLandingCallStatusCommandDirection) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLandingCallStatusCommandDirection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLandingCallStatusCommandDirection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLandingCallStatusCommandDirection")
		}

		if err := WriteSimpleField[BACnetLiftCarDirectionTagged](ctx, "direction", m.GetDirection(), WriteComplex[BACnetLiftCarDirectionTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'direction' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLandingCallStatusCommandDirection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLandingCallStatusCommandDirection")
		}
		return nil
	}
	return m.BACnetLandingCallStatusCommandContract.(*_BACnetLandingCallStatusCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLandingCallStatusCommandDirection) IsBACnetLandingCallStatusCommandDirection() {}

func (m *_BACnetLandingCallStatusCommandDirection) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLandingCallStatusCommandDirection) deepCopy() *_BACnetLandingCallStatusCommandDirection {
	if m == nil {
		return nil
	}
	_BACnetLandingCallStatusCommandDirectionCopy := &_BACnetLandingCallStatusCommandDirection{
		m.BACnetLandingCallStatusCommandContract.(*_BACnetLandingCallStatusCommand).deepCopy(),
		utils.DeepCopy[BACnetLiftCarDirectionTagged](m.Direction),
	}
	_BACnetLandingCallStatusCommandDirectionCopy.BACnetLandingCallStatusCommandContract.(*_BACnetLandingCallStatusCommand)._SubType = m
	return _BACnetLandingCallStatusCommandDirectionCopy
}

func (m *_BACnetLandingCallStatusCommandDirection) String() string {
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
