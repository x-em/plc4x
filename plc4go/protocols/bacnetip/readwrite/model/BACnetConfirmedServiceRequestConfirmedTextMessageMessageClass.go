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

// BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass is the corresponding interface of BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass interface {
	BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract
	BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass()
	// CreateBuilder creates a BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder
	CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder
}

// BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract provides a set of functions which can be overwritten by a sub struct
type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetTagNumber() returns a parser argument
	GetTagNumber() uint8
	// IsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass()
	// CreateBuilder creates a BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder
	CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder
}

// BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass is the data-structure of this message
type _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass struct {
	_SubType interface {
		BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract
		BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassRequirements
	}
	OpeningTag      BACnetOpeningTag
	PeekedTagHeader BACnetTagHeader
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract = (*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass)(nil)

// NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass factory function for _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
func NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass must not be nil")
	}
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass must not be nil")
	}
	return &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder is a builder for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder
	// WithPeekedTagHeader adds PeekedTagHeader (property field)
	WithPeekedTagHeader(BACnetTagHeader) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder
	// WithPeekedTagHeaderBuilder adds PeekedTagHeader (property field) which is build by the builder
	WithPeekedTagHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder
	// AsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric converts this build to a subType of BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass. It is always possible to return to current builder using Done()
	AsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder
	// AsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter converts this build to a subType of BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass. It is always possible to return to current builder using Done()
	AsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder
	// Build builds the BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass or returns an error if something is wrong
	PartialBuild() (BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract
	// Build builds the BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
}

// NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder() creates a BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder
func NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder {
	return &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder{_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass: new(_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass)}
}

type _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassChildBuilder interface {
	utils.Copyable
	setParent(BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract)
	buildForBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass() (BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass, error)
}

type _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder struct {
	*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass

	childBuilder _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassChildBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder) = (*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder {
	return b.WithOpeningTag(openingTag).WithPeekedTagHeader(peekedTagHeader).WithClosingTag(closingTag)
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder {
	builder := builderSupplier(b.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.OpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder) WithPeekedTagHeader(peekedTagHeader BACnetTagHeader) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder {
	b.PeekedTagHeader = peekedTagHeader
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder) WithPeekedTagHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder {
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

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder {
	builder := builderSupplier(b.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.ClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder) WithArgTagNumber(tagNumber uint8) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder {
	b.TagNumber = tagNumber
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder) PartialBuild() (BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.PeekedTagHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'peekedTagHeader' not set"))
	}
	if b.ClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder) PartialMustBuild() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder) AsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder {
	if cb, ok := b.childBuilder.(BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder); ok {
		return cb
	}
	cb := NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder().(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder) AsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder {
	if cb, ok := b.childBuilder.(BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder); ok {
		return cb
	}
	cb := NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder().(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder) Build() (BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass()
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder) MustBuild() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder().(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder creates a BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder
func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder()
	}
	return &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder{_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) GetPeekedTagNumber() uint8 {
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
func CastBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass(structType any) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass"
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassParse[T BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass](ctx context.Context, theBytes []byte, tagNumber uint8) (T, error) {
	return BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassParseWithBufferProducer[T BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass](tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassParseWithBuffer[T](ctx, readBuffer, tagNumber)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassParseWithBuffer[T BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass](ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (T, error) {
	v, err := (&_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
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

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetConfirmedServiceRequestConfirmedTextMessageMessageClass BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

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
	var _child BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	switch {
	case peekedTagNumber == uint8(0): // BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric
		if _child, err = new(_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric for type-switch of BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass")
		}
	case peekedTagNumber == uint8(1): // BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter
		if _child, err = new(_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter for type-switch of BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass")
	}

	return _child, nil
}

func (pm *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
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

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) IsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass() {
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) deepCopy() *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCopy := &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass{
		nil, // will be set by child
		utils.DeepCopy[BACnetOpeningTag](m.OpeningTag),
		utils.DeepCopy[BACnetTagHeader](m.PeekedTagHeader),
		utils.DeepCopy[BACnetClosingTag](m.ClosingTag),
		m.TagNumber,
	}
	return _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCopy
}