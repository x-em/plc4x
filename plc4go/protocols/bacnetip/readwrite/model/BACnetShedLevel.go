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

// BACnetShedLevel is the corresponding interface of BACnetShedLevel
type BACnetShedLevel interface {
	BACnetShedLevelContract
	BACnetShedLevelRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBACnetShedLevel is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetShedLevel()
	// CreateBuilder creates a BACnetShedLevelBuilder
	CreateBACnetShedLevelBuilder() BACnetShedLevelBuilder
}

// BACnetShedLevelContract provides a set of functions which can be overwritten by a sub struct
type BACnetShedLevelContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// IsBACnetShedLevel is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetShedLevel()
	// CreateBuilder creates a BACnetShedLevelBuilder
	CreateBACnetShedLevelBuilder() BACnetShedLevelBuilder
}

// BACnetShedLevelRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetShedLevelRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetShedLevel is the data-structure of this message
type _BACnetShedLevel struct {
	_SubType interface {
		BACnetShedLevelContract
		BACnetShedLevelRequirements
	}
	PeekedTagHeader BACnetTagHeader
}

var _ BACnetShedLevelContract = (*_BACnetShedLevel)(nil)

// NewBACnetShedLevel factory function for _BACnetShedLevel
func NewBACnetShedLevel(peekedTagHeader BACnetTagHeader) *_BACnetShedLevel {
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetShedLevel must not be nil")
	}
	return &_BACnetShedLevel{PeekedTagHeader: peekedTagHeader}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetShedLevelBuilder is a builder for BACnetShedLevel
type BACnetShedLevelBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(peekedTagHeader BACnetTagHeader) BACnetShedLevelBuilder
	// WithPeekedTagHeader adds PeekedTagHeader (property field)
	WithPeekedTagHeader(BACnetTagHeader) BACnetShedLevelBuilder
	// WithPeekedTagHeaderBuilder adds PeekedTagHeader (property field) which is build by the builder
	WithPeekedTagHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetShedLevelBuilder
	// AsBACnetShedLevelPercent converts this build to a subType of BACnetShedLevel. It is always possible to return to current builder using Done()
	AsBACnetShedLevelPercent() BACnetShedLevelPercentBuilder
	// AsBACnetShedLevelLevel converts this build to a subType of BACnetShedLevel. It is always possible to return to current builder using Done()
	AsBACnetShedLevelLevel() BACnetShedLevelLevelBuilder
	// AsBACnetShedLevelAmount converts this build to a subType of BACnetShedLevel. It is always possible to return to current builder using Done()
	AsBACnetShedLevelAmount() BACnetShedLevelAmountBuilder
	// Build builds the BACnetShedLevel or returns an error if something is wrong
	PartialBuild() (BACnetShedLevelContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() BACnetShedLevelContract
	// Build builds the BACnetShedLevel or returns an error if something is wrong
	Build() (BACnetShedLevel, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetShedLevel
}

// NewBACnetShedLevelBuilder() creates a BACnetShedLevelBuilder
func NewBACnetShedLevelBuilder() BACnetShedLevelBuilder {
	return &_BACnetShedLevelBuilder{_BACnetShedLevel: new(_BACnetShedLevel)}
}

type _BACnetShedLevelChildBuilder interface {
	utils.Copyable
	setParent(BACnetShedLevelContract)
	buildForBACnetShedLevel() (BACnetShedLevel, error)
}

type _BACnetShedLevelBuilder struct {
	*_BACnetShedLevel

	childBuilder _BACnetShedLevelChildBuilder

	err *utils.MultiError
}

var _ (BACnetShedLevelBuilder) = (*_BACnetShedLevelBuilder)(nil)

func (b *_BACnetShedLevelBuilder) WithMandatoryFields(peekedTagHeader BACnetTagHeader) BACnetShedLevelBuilder {
	return b.WithPeekedTagHeader(peekedTagHeader)
}

func (b *_BACnetShedLevelBuilder) WithPeekedTagHeader(peekedTagHeader BACnetTagHeader) BACnetShedLevelBuilder {
	b.PeekedTagHeader = peekedTagHeader
	return b
}

func (b *_BACnetShedLevelBuilder) WithPeekedTagHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetShedLevelBuilder {
	builder := builderSupplier(b.PeekedTagHeader.CreateBACnetTagHeaderBuilder())
	var err error
	b.PeekedTagHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return b
}

func (b *_BACnetShedLevelBuilder) PartialBuild() (BACnetShedLevelContract, error) {
	if b.PeekedTagHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'peekedTagHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetShedLevel.deepCopy(), nil
}

func (b *_BACnetShedLevelBuilder) PartialMustBuild() BACnetShedLevelContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetShedLevelBuilder) AsBACnetShedLevelPercent() BACnetShedLevelPercentBuilder {
	if cb, ok := b.childBuilder.(BACnetShedLevelPercentBuilder); ok {
		return cb
	}
	cb := NewBACnetShedLevelPercentBuilder().(*_BACnetShedLevelPercentBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetShedLevelBuilder) AsBACnetShedLevelLevel() BACnetShedLevelLevelBuilder {
	if cb, ok := b.childBuilder.(BACnetShedLevelLevelBuilder); ok {
		return cb
	}
	cb := NewBACnetShedLevelLevelBuilder().(*_BACnetShedLevelLevelBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetShedLevelBuilder) AsBACnetShedLevelAmount() BACnetShedLevelAmountBuilder {
	if cb, ok := b.childBuilder.(BACnetShedLevelAmountBuilder); ok {
		return cb
	}
	cb := NewBACnetShedLevelAmountBuilder().(*_BACnetShedLevelAmountBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetShedLevelBuilder) Build() (BACnetShedLevel, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForBACnetShedLevel()
}

func (b *_BACnetShedLevelBuilder) MustBuild() BACnetShedLevel {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetShedLevelBuilder) DeepCopy() any {
	_copy := b.CreateBACnetShedLevelBuilder().(*_BACnetShedLevelBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_BACnetShedLevelChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetShedLevelBuilder creates a BACnetShedLevelBuilder
func (b *_BACnetShedLevel) CreateBACnetShedLevelBuilder() BACnetShedLevelBuilder {
	if b == nil {
		return NewBACnetShedLevelBuilder()
	}
	return &_BACnetShedLevelBuilder{_BACnetShedLevel: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetShedLevel) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetShedLevel) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetShedLevel(structType any) BACnetShedLevel {
	if casted, ok := structType.(BACnetShedLevel); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetShedLevel); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetShedLevel) GetTypeName() string {
	return "BACnetShedLevel"
}

func (m *_BACnetShedLevel) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetShedLevel) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_BACnetShedLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetShedLevelParse[T BACnetShedLevel](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetShedLevelParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetShedLevelParseWithBufferProducer[T BACnetShedLevel]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetShedLevelParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetShedLevelParseWithBuffer[T BACnetShedLevel](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetShedLevel{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_BACnetShedLevel) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetShedLevel BACnetShedLevel, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetShedLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetShedLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetShedLevel
	switch {
	case peekedTagNumber == uint8(0): // BACnetShedLevelPercent
		if _child, err = new(_BACnetShedLevelPercent).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetShedLevelPercent for type-switch of BACnetShedLevel")
		}
	case peekedTagNumber == uint8(1): // BACnetShedLevelLevel
		if _child, err = new(_BACnetShedLevelLevel).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetShedLevelLevel for type-switch of BACnetShedLevel")
		}
	case peekedTagNumber == uint8(2): // BACnetShedLevelAmount
		if _child, err = new(_BACnetShedLevelAmount).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetShedLevelAmount for type-switch of BACnetShedLevel")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	if closeErr := readBuffer.CloseContext("BACnetShedLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetShedLevel")
	}

	return _child, nil
}

func (pm *_BACnetShedLevel) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetShedLevel, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetShedLevel"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetShedLevel")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetShedLevel"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetShedLevel")
	}
	return nil
}

func (m *_BACnetShedLevel) IsBACnetShedLevel() {}

func (m *_BACnetShedLevel) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetShedLevel) deepCopy() *_BACnetShedLevel {
	if m == nil {
		return nil
	}
	_BACnetShedLevelCopy := &_BACnetShedLevel{
		nil, // will be set by child
		utils.DeepCopy[BACnetTagHeader](m.PeekedTagHeader),
	}
	return _BACnetShedLevelCopy
}