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

// BACnetShedLevelLevel is the corresponding interface of BACnetShedLevelLevel
type BACnetShedLevelLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetShedLevel
	// GetLevel returns Level (property field)
	GetLevel() BACnetContextTagUnsignedInteger
	// IsBACnetShedLevelLevel is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetShedLevelLevel()
	// CreateBuilder creates a BACnetShedLevelLevelBuilder
	CreateBACnetShedLevelLevelBuilder() BACnetShedLevelLevelBuilder
}

// _BACnetShedLevelLevel is the data-structure of this message
type _BACnetShedLevelLevel struct {
	BACnetShedLevelContract
	Level BACnetContextTagUnsignedInteger
}

var _ BACnetShedLevelLevel = (*_BACnetShedLevelLevel)(nil)
var _ BACnetShedLevelRequirements = (*_BACnetShedLevelLevel)(nil)

// NewBACnetShedLevelLevel factory function for _BACnetShedLevelLevel
func NewBACnetShedLevelLevel(peekedTagHeader BACnetTagHeader, level BACnetContextTagUnsignedInteger) *_BACnetShedLevelLevel {
	if level == nil {
		panic("level of type BACnetContextTagUnsignedInteger for BACnetShedLevelLevel must not be nil")
	}
	_result := &_BACnetShedLevelLevel{
		BACnetShedLevelContract: NewBACnetShedLevel(peekedTagHeader),
		Level:                   level,
	}
	_result.BACnetShedLevelContract.(*_BACnetShedLevel)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetShedLevelLevelBuilder is a builder for BACnetShedLevelLevel
type BACnetShedLevelLevelBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(level BACnetContextTagUnsignedInteger) BACnetShedLevelLevelBuilder
	// WithLevel adds Level (property field)
	WithLevel(BACnetContextTagUnsignedInteger) BACnetShedLevelLevelBuilder
	// WithLevelBuilder adds Level (property field) which is build by the builder
	WithLevelBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetShedLevelLevelBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetShedLevelBuilder
	// Build builds the BACnetShedLevelLevel or returns an error if something is wrong
	Build() (BACnetShedLevelLevel, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetShedLevelLevel
}

// NewBACnetShedLevelLevelBuilder() creates a BACnetShedLevelLevelBuilder
func NewBACnetShedLevelLevelBuilder() BACnetShedLevelLevelBuilder {
	return &_BACnetShedLevelLevelBuilder{_BACnetShedLevelLevel: new(_BACnetShedLevelLevel)}
}

type _BACnetShedLevelLevelBuilder struct {
	*_BACnetShedLevelLevel

	parentBuilder *_BACnetShedLevelBuilder

	err *utils.MultiError
}

var _ (BACnetShedLevelLevelBuilder) = (*_BACnetShedLevelLevelBuilder)(nil)

func (b *_BACnetShedLevelLevelBuilder) setParent(contract BACnetShedLevelContract) {
	b.BACnetShedLevelContract = contract
	contract.(*_BACnetShedLevel)._SubType = b._BACnetShedLevelLevel
}

func (b *_BACnetShedLevelLevelBuilder) WithMandatoryFields(level BACnetContextTagUnsignedInteger) BACnetShedLevelLevelBuilder {
	return b.WithLevel(level)
}

func (b *_BACnetShedLevelLevelBuilder) WithLevel(level BACnetContextTagUnsignedInteger) BACnetShedLevelLevelBuilder {
	b.Level = level
	return b
}

func (b *_BACnetShedLevelLevelBuilder) WithLevelBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetShedLevelLevelBuilder {
	builder := builderSupplier(b.Level.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.Level, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetShedLevelLevelBuilder) Build() (BACnetShedLevelLevel, error) {
	if b.Level == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'level' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetShedLevelLevel.deepCopy(), nil
}

func (b *_BACnetShedLevelLevelBuilder) MustBuild() BACnetShedLevelLevel {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetShedLevelLevelBuilder) Done() BACnetShedLevelBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetShedLevelBuilder().(*_BACnetShedLevelBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetShedLevelLevelBuilder) buildForBACnetShedLevel() (BACnetShedLevel, error) {
	return b.Build()
}

func (b *_BACnetShedLevelLevelBuilder) DeepCopy() any {
	_copy := b.CreateBACnetShedLevelLevelBuilder().(*_BACnetShedLevelLevelBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetShedLevelLevelBuilder creates a BACnetShedLevelLevelBuilder
func (b *_BACnetShedLevelLevel) CreateBACnetShedLevelLevelBuilder() BACnetShedLevelLevelBuilder {
	if b == nil {
		return NewBACnetShedLevelLevelBuilder()
	}
	return &_BACnetShedLevelLevelBuilder{_BACnetShedLevelLevel: b.deepCopy()}
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

func (m *_BACnetShedLevelLevel) GetParent() BACnetShedLevelContract {
	return m.BACnetShedLevelContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetShedLevelLevel) GetLevel() BACnetContextTagUnsignedInteger {
	return m.Level
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetShedLevelLevel(structType any) BACnetShedLevelLevel {
	if casted, ok := structType.(BACnetShedLevelLevel); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetShedLevelLevel); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetShedLevelLevel) GetTypeName() string {
	return "BACnetShedLevelLevel"
}

func (m *_BACnetShedLevelLevel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetShedLevelContract.(*_BACnetShedLevel).getLengthInBits(ctx))

	// Simple field (level)
	lengthInBits += m.Level.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetShedLevelLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetShedLevelLevel) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetShedLevel) (__bACnetShedLevelLevel BACnetShedLevelLevel, err error) {
	m.BACnetShedLevelContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetShedLevelLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetShedLevelLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	level, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "level", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'level' field"))
	}
	m.Level = level

	if closeErr := readBuffer.CloseContext("BACnetShedLevelLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetShedLevelLevel")
	}

	return m, nil
}

func (m *_BACnetShedLevelLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetShedLevelLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetShedLevelLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetShedLevelLevel")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "level", m.GetLevel(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'level' field")
		}

		if popErr := writeBuffer.PopContext("BACnetShedLevelLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetShedLevelLevel")
		}
		return nil
	}
	return m.BACnetShedLevelContract.(*_BACnetShedLevel).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetShedLevelLevel) IsBACnetShedLevelLevel() {}

func (m *_BACnetShedLevelLevel) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetShedLevelLevel) deepCopy() *_BACnetShedLevelLevel {
	if m == nil {
		return nil
	}
	_BACnetShedLevelLevelCopy := &_BACnetShedLevelLevel{
		m.BACnetShedLevelContract.(*_BACnetShedLevel).deepCopy(),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.Level),
	}
	_BACnetShedLevelLevelCopy.BACnetShedLevelContract.(*_BACnetShedLevel)._SubType = m
	return _BACnetShedLevelLevelCopy
}

func (m *_BACnetShedLevelLevel) String() string {
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