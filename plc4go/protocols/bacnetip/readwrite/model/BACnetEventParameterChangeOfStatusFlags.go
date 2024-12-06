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

// BACnetEventParameterChangeOfStatusFlags is the corresponding interface of BACnetEventParameterChangeOfStatusFlags
type BACnetEventParameterChangeOfStatusFlags interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetSelectedFlags returns SelectedFlags (property field)
	GetSelectedFlags() BACnetStatusFlagsTagged
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventParameterChangeOfStatusFlags is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterChangeOfStatusFlags()
	// CreateBuilder creates a BACnetEventParameterChangeOfStatusFlagsBuilder
	CreateBACnetEventParameterChangeOfStatusFlagsBuilder() BACnetEventParameterChangeOfStatusFlagsBuilder
}

// _BACnetEventParameterChangeOfStatusFlags is the data-structure of this message
type _BACnetEventParameterChangeOfStatusFlags struct {
	BACnetEventParameterContract
	OpeningTag    BACnetOpeningTag
	TimeDelay     BACnetContextTagUnsignedInteger
	SelectedFlags BACnetStatusFlagsTagged
	ClosingTag    BACnetClosingTag
}

var _ BACnetEventParameterChangeOfStatusFlags = (*_BACnetEventParameterChangeOfStatusFlags)(nil)
var _ BACnetEventParameterRequirements = (*_BACnetEventParameterChangeOfStatusFlags)(nil)

// NewBACnetEventParameterChangeOfStatusFlags factory function for _BACnetEventParameterChangeOfStatusFlags
func NewBACnetEventParameterChangeOfStatusFlags(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, selectedFlags BACnetStatusFlagsTagged, closingTag BACnetClosingTag) *_BACnetEventParameterChangeOfStatusFlags {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventParameterChangeOfStatusFlags must not be nil")
	}
	if timeDelay == nil {
		panic("timeDelay of type BACnetContextTagUnsignedInteger for BACnetEventParameterChangeOfStatusFlags must not be nil")
	}
	if selectedFlags == nil {
		panic("selectedFlags of type BACnetStatusFlagsTagged for BACnetEventParameterChangeOfStatusFlags must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventParameterChangeOfStatusFlags must not be nil")
	}
	_result := &_BACnetEventParameterChangeOfStatusFlags{
		BACnetEventParameterContract: NewBACnetEventParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		TimeDelay:                    timeDelay,
		SelectedFlags:                selectedFlags,
		ClosingTag:                   closingTag,
	}
	_result.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventParameterChangeOfStatusFlagsBuilder is a builder for BACnetEventParameterChangeOfStatusFlags
type BACnetEventParameterChangeOfStatusFlagsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, selectedFlags BACnetStatusFlagsTagged, closingTag BACnetClosingTag) BACnetEventParameterChangeOfStatusFlagsBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetEventParameterChangeOfStatusFlagsBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterChangeOfStatusFlagsBuilder
	// WithTimeDelay adds TimeDelay (property field)
	WithTimeDelay(BACnetContextTagUnsignedInteger) BACnetEventParameterChangeOfStatusFlagsBuilder
	// WithTimeDelayBuilder adds TimeDelay (property field) which is build by the builder
	WithTimeDelayBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterChangeOfStatusFlagsBuilder
	// WithSelectedFlags adds SelectedFlags (property field)
	WithSelectedFlags(BACnetStatusFlagsTagged) BACnetEventParameterChangeOfStatusFlagsBuilder
	// WithSelectedFlagsBuilder adds SelectedFlags (property field) which is build by the builder
	WithSelectedFlagsBuilder(func(BACnetStatusFlagsTaggedBuilder) BACnetStatusFlagsTaggedBuilder) BACnetEventParameterChangeOfStatusFlagsBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetEventParameterChangeOfStatusFlagsBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterChangeOfStatusFlagsBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetEventParameterBuilder
	// Build builds the BACnetEventParameterChangeOfStatusFlags or returns an error if something is wrong
	Build() (BACnetEventParameterChangeOfStatusFlags, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventParameterChangeOfStatusFlags
}

// NewBACnetEventParameterChangeOfStatusFlagsBuilder() creates a BACnetEventParameterChangeOfStatusFlagsBuilder
func NewBACnetEventParameterChangeOfStatusFlagsBuilder() BACnetEventParameterChangeOfStatusFlagsBuilder {
	return &_BACnetEventParameterChangeOfStatusFlagsBuilder{_BACnetEventParameterChangeOfStatusFlags: new(_BACnetEventParameterChangeOfStatusFlags)}
}

type _BACnetEventParameterChangeOfStatusFlagsBuilder struct {
	*_BACnetEventParameterChangeOfStatusFlags

	parentBuilder *_BACnetEventParameterBuilder

	err *utils.MultiError
}

var _ (BACnetEventParameterChangeOfStatusFlagsBuilder) = (*_BACnetEventParameterChangeOfStatusFlagsBuilder)(nil)

func (b *_BACnetEventParameterChangeOfStatusFlagsBuilder) setParent(contract BACnetEventParameterContract) {
	b.BACnetEventParameterContract = contract
	contract.(*_BACnetEventParameter)._SubType = b._BACnetEventParameterChangeOfStatusFlags
}

func (b *_BACnetEventParameterChangeOfStatusFlagsBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, selectedFlags BACnetStatusFlagsTagged, closingTag BACnetClosingTag) BACnetEventParameterChangeOfStatusFlagsBuilder {
	return b.WithOpeningTag(openingTag).WithTimeDelay(timeDelay).WithSelectedFlags(selectedFlags).WithClosingTag(closingTag)
}

func (b *_BACnetEventParameterChangeOfStatusFlagsBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetEventParameterChangeOfStatusFlagsBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetEventParameterChangeOfStatusFlagsBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterChangeOfStatusFlagsBuilder {
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

func (b *_BACnetEventParameterChangeOfStatusFlagsBuilder) WithTimeDelay(timeDelay BACnetContextTagUnsignedInteger) BACnetEventParameterChangeOfStatusFlagsBuilder {
	b.TimeDelay = timeDelay
	return b
}

func (b *_BACnetEventParameterChangeOfStatusFlagsBuilder) WithTimeDelayBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterChangeOfStatusFlagsBuilder {
	builder := builderSupplier(b.TimeDelay.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.TimeDelay, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterChangeOfStatusFlagsBuilder) WithSelectedFlags(selectedFlags BACnetStatusFlagsTagged) BACnetEventParameterChangeOfStatusFlagsBuilder {
	b.SelectedFlags = selectedFlags
	return b
}

func (b *_BACnetEventParameterChangeOfStatusFlagsBuilder) WithSelectedFlagsBuilder(builderSupplier func(BACnetStatusFlagsTaggedBuilder) BACnetStatusFlagsTaggedBuilder) BACnetEventParameterChangeOfStatusFlagsBuilder {
	builder := builderSupplier(b.SelectedFlags.CreateBACnetStatusFlagsTaggedBuilder())
	var err error
	b.SelectedFlags, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetStatusFlagsTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterChangeOfStatusFlagsBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetEventParameterChangeOfStatusFlagsBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetEventParameterChangeOfStatusFlagsBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterChangeOfStatusFlagsBuilder {
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

func (b *_BACnetEventParameterChangeOfStatusFlagsBuilder) Build() (BACnetEventParameterChangeOfStatusFlags, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.TimeDelay == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeDelay' not set"))
	}
	if b.SelectedFlags == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'selectedFlags' not set"))
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
	return b._BACnetEventParameterChangeOfStatusFlags.deepCopy(), nil
}

func (b *_BACnetEventParameterChangeOfStatusFlagsBuilder) MustBuild() BACnetEventParameterChangeOfStatusFlags {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetEventParameterChangeOfStatusFlagsBuilder) Done() BACnetEventParameterBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetEventParameterBuilder().(*_BACnetEventParameterBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetEventParameterChangeOfStatusFlagsBuilder) buildForBACnetEventParameter() (BACnetEventParameter, error) {
	return b.Build()
}

func (b *_BACnetEventParameterChangeOfStatusFlagsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetEventParameterChangeOfStatusFlagsBuilder().(*_BACnetEventParameterChangeOfStatusFlagsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetEventParameterChangeOfStatusFlagsBuilder creates a BACnetEventParameterChangeOfStatusFlagsBuilder
func (b *_BACnetEventParameterChangeOfStatusFlags) CreateBACnetEventParameterChangeOfStatusFlagsBuilder() BACnetEventParameterChangeOfStatusFlagsBuilder {
	if b == nil {
		return NewBACnetEventParameterChangeOfStatusFlagsBuilder()
	}
	return &_BACnetEventParameterChangeOfStatusFlagsBuilder{_BACnetEventParameterChangeOfStatusFlags: b.deepCopy()}
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

func (m *_BACnetEventParameterChangeOfStatusFlags) GetParent() BACnetEventParameterContract {
	return m.BACnetEventParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfStatusFlags) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfStatusFlags) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterChangeOfStatusFlags) GetSelectedFlags() BACnetStatusFlagsTagged {
	return m.SelectedFlags
}

func (m *_BACnetEventParameterChangeOfStatusFlags) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfStatusFlags(structType any) BACnetEventParameterChangeOfStatusFlags {
	if casted, ok := structType.(BACnetEventParameterChangeOfStatusFlags); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfStatusFlags); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfStatusFlags) GetTypeName() string {
	return "BACnetEventParameterChangeOfStatusFlags"
}

func (m *_BACnetEventParameterChangeOfStatusFlags) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventParameterContract.(*_BACnetEventParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits(ctx)

	// Simple field (selectedFlags)
	lengthInBits += m.SelectedFlags.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfStatusFlags) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventParameterChangeOfStatusFlags) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventParameter) (__bACnetEventParameterChangeOfStatusFlags BACnetEventParameterChangeOfStatusFlags, err error) {
	m.BACnetEventParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfStatusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfStatusFlags")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(18))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	timeDelay, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeDelay' field"))
	}
	m.TimeDelay = timeDelay

	selectedFlags, err := ReadSimpleField[BACnetStatusFlagsTagged](ctx, "selectedFlags", ReadComplex[BACnetStatusFlagsTagged](BACnetStatusFlagsTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'selectedFlags' field"))
	}
	m.SelectedFlags = selectedFlags

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(18))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfStatusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfStatusFlags")
	}

	return m, nil
}

func (m *_BACnetEventParameterChangeOfStatusFlags) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfStatusFlags) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfStatusFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfStatusFlags")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", m.GetTimeDelay(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeDelay' field")
		}

		if err := WriteSimpleField[BACnetStatusFlagsTagged](ctx, "selectedFlags", m.GetSelectedFlags(), WriteComplex[BACnetStatusFlagsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'selectedFlags' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfStatusFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfStatusFlags")
		}
		return nil
	}
	return m.BACnetEventParameterContract.(*_BACnetEventParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterChangeOfStatusFlags) IsBACnetEventParameterChangeOfStatusFlags() {}

func (m *_BACnetEventParameterChangeOfStatusFlags) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventParameterChangeOfStatusFlags) deepCopy() *_BACnetEventParameterChangeOfStatusFlags {
	if m == nil {
		return nil
	}
	_BACnetEventParameterChangeOfStatusFlagsCopy := &_BACnetEventParameterChangeOfStatusFlags{
		m.BACnetEventParameterContract.(*_BACnetEventParameter).deepCopy(),
		utils.DeepCopy[BACnetOpeningTag](m.OpeningTag),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.TimeDelay),
		utils.DeepCopy[BACnetStatusFlagsTagged](m.SelectedFlags),
		utils.DeepCopy[BACnetClosingTag](m.ClosingTag),
	}
	_BACnetEventParameterChangeOfStatusFlagsCopy.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = m
	return _BACnetEventParameterChangeOfStatusFlagsCopy
}

func (m *_BACnetEventParameterChangeOfStatusFlags) String() string {
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