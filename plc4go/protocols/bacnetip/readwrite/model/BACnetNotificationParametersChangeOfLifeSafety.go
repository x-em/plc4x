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

// BACnetNotificationParametersChangeOfLifeSafety is the corresponding interface of BACnetNotificationParametersChangeOfLifeSafety
type BACnetNotificationParametersChangeOfLifeSafety interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetNewState returns NewState (property field)
	GetNewState() BACnetLifeSafetyStateTagged
	// GetNewMode returns NewMode (property field)
	GetNewMode() BACnetLifeSafetyModeTagged
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetOperationExpected returns OperationExpected (property field)
	GetOperationExpected() BACnetLifeSafetyOperationTagged
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
	// IsBACnetNotificationParametersChangeOfLifeSafety is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersChangeOfLifeSafety()
	// CreateBuilder creates a BACnetNotificationParametersChangeOfLifeSafetyBuilder
	CreateBACnetNotificationParametersChangeOfLifeSafetyBuilder() BACnetNotificationParametersChangeOfLifeSafetyBuilder
}

// _BACnetNotificationParametersChangeOfLifeSafety is the data-structure of this message
type _BACnetNotificationParametersChangeOfLifeSafety struct {
	BACnetNotificationParametersContract
	InnerOpeningTag   BACnetOpeningTag
	NewState          BACnetLifeSafetyStateTagged
	NewMode           BACnetLifeSafetyModeTagged
	StatusFlags       BACnetStatusFlagsTagged
	OperationExpected BACnetLifeSafetyOperationTagged
	InnerClosingTag   BACnetClosingTag
}

var _ BACnetNotificationParametersChangeOfLifeSafety = (*_BACnetNotificationParametersChangeOfLifeSafety)(nil)
var _ BACnetNotificationParametersRequirements = (*_BACnetNotificationParametersChangeOfLifeSafety)(nil)

// NewBACnetNotificationParametersChangeOfLifeSafety factory function for _BACnetNotificationParametersChangeOfLifeSafety
func NewBACnetNotificationParametersChangeOfLifeSafety(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, innerOpeningTag BACnetOpeningTag, newState BACnetLifeSafetyStateTagged, newMode BACnetLifeSafetyModeTagged, statusFlags BACnetStatusFlagsTagged, operationExpected BACnetLifeSafetyOperationTagged, innerClosingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersChangeOfLifeSafety {
	if innerOpeningTag == nil {
		panic("innerOpeningTag of type BACnetOpeningTag for BACnetNotificationParametersChangeOfLifeSafety must not be nil")
	}
	if newState == nil {
		panic("newState of type BACnetLifeSafetyStateTagged for BACnetNotificationParametersChangeOfLifeSafety must not be nil")
	}
	if newMode == nil {
		panic("newMode of type BACnetLifeSafetyModeTagged for BACnetNotificationParametersChangeOfLifeSafety must not be nil")
	}
	if statusFlags == nil {
		panic("statusFlags of type BACnetStatusFlagsTagged for BACnetNotificationParametersChangeOfLifeSafety must not be nil")
	}
	if operationExpected == nil {
		panic("operationExpected of type BACnetLifeSafetyOperationTagged for BACnetNotificationParametersChangeOfLifeSafety must not be nil")
	}
	if innerClosingTag == nil {
		panic("innerClosingTag of type BACnetClosingTag for BACnetNotificationParametersChangeOfLifeSafety must not be nil")
	}
	_result := &_BACnetNotificationParametersChangeOfLifeSafety{
		BACnetNotificationParametersContract: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
		InnerOpeningTag:                      innerOpeningTag,
		NewState:                             newState,
		NewMode:                              newMode,
		StatusFlags:                          statusFlags,
		OperationExpected:                    operationExpected,
		InnerClosingTag:                      innerClosingTag,
	}
	_result.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNotificationParametersChangeOfLifeSafetyBuilder is a builder for BACnetNotificationParametersChangeOfLifeSafety
type BACnetNotificationParametersChangeOfLifeSafetyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(innerOpeningTag BACnetOpeningTag, newState BACnetLifeSafetyStateTagged, newMode BACnetLifeSafetyModeTagged, statusFlags BACnetStatusFlagsTagged, operationExpected BACnetLifeSafetyOperationTagged, innerClosingTag BACnetClosingTag) BACnetNotificationParametersChangeOfLifeSafetyBuilder
	// WithInnerOpeningTag adds InnerOpeningTag (property field)
	WithInnerOpeningTag(BACnetOpeningTag) BACnetNotificationParametersChangeOfLifeSafetyBuilder
	// WithInnerOpeningTagBuilder adds InnerOpeningTag (property field) which is build by the builder
	WithInnerOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersChangeOfLifeSafetyBuilder
	// WithNewState adds NewState (property field)
	WithNewState(BACnetLifeSafetyStateTagged) BACnetNotificationParametersChangeOfLifeSafetyBuilder
	// WithNewStateBuilder adds NewState (property field) which is build by the builder
	WithNewStateBuilder(func(BACnetLifeSafetyStateTaggedBuilder) BACnetLifeSafetyStateTaggedBuilder) BACnetNotificationParametersChangeOfLifeSafetyBuilder
	// WithNewMode adds NewMode (property field)
	WithNewMode(BACnetLifeSafetyModeTagged) BACnetNotificationParametersChangeOfLifeSafetyBuilder
	// WithNewModeBuilder adds NewMode (property field) which is build by the builder
	WithNewModeBuilder(func(BACnetLifeSafetyModeTaggedBuilder) BACnetLifeSafetyModeTaggedBuilder) BACnetNotificationParametersChangeOfLifeSafetyBuilder
	// WithStatusFlags adds StatusFlags (property field)
	WithStatusFlags(BACnetStatusFlagsTagged) BACnetNotificationParametersChangeOfLifeSafetyBuilder
	// WithStatusFlagsBuilder adds StatusFlags (property field) which is build by the builder
	WithStatusFlagsBuilder(func(BACnetStatusFlagsTaggedBuilder) BACnetStatusFlagsTaggedBuilder) BACnetNotificationParametersChangeOfLifeSafetyBuilder
	// WithOperationExpected adds OperationExpected (property field)
	WithOperationExpected(BACnetLifeSafetyOperationTagged) BACnetNotificationParametersChangeOfLifeSafetyBuilder
	// WithOperationExpectedBuilder adds OperationExpected (property field) which is build by the builder
	WithOperationExpectedBuilder(func(BACnetLifeSafetyOperationTaggedBuilder) BACnetLifeSafetyOperationTaggedBuilder) BACnetNotificationParametersChangeOfLifeSafetyBuilder
	// WithInnerClosingTag adds InnerClosingTag (property field)
	WithInnerClosingTag(BACnetClosingTag) BACnetNotificationParametersChangeOfLifeSafetyBuilder
	// WithInnerClosingTagBuilder adds InnerClosingTag (property field) which is build by the builder
	WithInnerClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersChangeOfLifeSafetyBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetNotificationParametersBuilder
	// Build builds the BACnetNotificationParametersChangeOfLifeSafety or returns an error if something is wrong
	Build() (BACnetNotificationParametersChangeOfLifeSafety, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNotificationParametersChangeOfLifeSafety
}

// NewBACnetNotificationParametersChangeOfLifeSafetyBuilder() creates a BACnetNotificationParametersChangeOfLifeSafetyBuilder
func NewBACnetNotificationParametersChangeOfLifeSafetyBuilder() BACnetNotificationParametersChangeOfLifeSafetyBuilder {
	return &_BACnetNotificationParametersChangeOfLifeSafetyBuilder{_BACnetNotificationParametersChangeOfLifeSafety: new(_BACnetNotificationParametersChangeOfLifeSafety)}
}

type _BACnetNotificationParametersChangeOfLifeSafetyBuilder struct {
	*_BACnetNotificationParametersChangeOfLifeSafety

	parentBuilder *_BACnetNotificationParametersBuilder

	err *utils.MultiError
}

var _ (BACnetNotificationParametersChangeOfLifeSafetyBuilder) = (*_BACnetNotificationParametersChangeOfLifeSafetyBuilder)(nil)

func (b *_BACnetNotificationParametersChangeOfLifeSafetyBuilder) setParent(contract BACnetNotificationParametersContract) {
	b.BACnetNotificationParametersContract = contract
	contract.(*_BACnetNotificationParameters)._SubType = b._BACnetNotificationParametersChangeOfLifeSafety
}

func (b *_BACnetNotificationParametersChangeOfLifeSafetyBuilder) WithMandatoryFields(innerOpeningTag BACnetOpeningTag, newState BACnetLifeSafetyStateTagged, newMode BACnetLifeSafetyModeTagged, statusFlags BACnetStatusFlagsTagged, operationExpected BACnetLifeSafetyOperationTagged, innerClosingTag BACnetClosingTag) BACnetNotificationParametersChangeOfLifeSafetyBuilder {
	return b.WithInnerOpeningTag(innerOpeningTag).WithNewState(newState).WithNewMode(newMode).WithStatusFlags(statusFlags).WithOperationExpected(operationExpected).WithInnerClosingTag(innerClosingTag)
}

func (b *_BACnetNotificationParametersChangeOfLifeSafetyBuilder) WithInnerOpeningTag(innerOpeningTag BACnetOpeningTag) BACnetNotificationParametersChangeOfLifeSafetyBuilder {
	b.InnerOpeningTag = innerOpeningTag
	return b
}

func (b *_BACnetNotificationParametersChangeOfLifeSafetyBuilder) WithInnerOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersChangeOfLifeSafetyBuilder {
	builder := builderSupplier(b.InnerOpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.InnerOpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersChangeOfLifeSafetyBuilder) WithNewState(newState BACnetLifeSafetyStateTagged) BACnetNotificationParametersChangeOfLifeSafetyBuilder {
	b.NewState = newState
	return b
}

func (b *_BACnetNotificationParametersChangeOfLifeSafetyBuilder) WithNewStateBuilder(builderSupplier func(BACnetLifeSafetyStateTaggedBuilder) BACnetLifeSafetyStateTaggedBuilder) BACnetNotificationParametersChangeOfLifeSafetyBuilder {
	builder := builderSupplier(b.NewState.CreateBACnetLifeSafetyStateTaggedBuilder())
	var err error
	b.NewState, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLifeSafetyStateTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersChangeOfLifeSafetyBuilder) WithNewMode(newMode BACnetLifeSafetyModeTagged) BACnetNotificationParametersChangeOfLifeSafetyBuilder {
	b.NewMode = newMode
	return b
}

func (b *_BACnetNotificationParametersChangeOfLifeSafetyBuilder) WithNewModeBuilder(builderSupplier func(BACnetLifeSafetyModeTaggedBuilder) BACnetLifeSafetyModeTaggedBuilder) BACnetNotificationParametersChangeOfLifeSafetyBuilder {
	builder := builderSupplier(b.NewMode.CreateBACnetLifeSafetyModeTaggedBuilder())
	var err error
	b.NewMode, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLifeSafetyModeTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersChangeOfLifeSafetyBuilder) WithStatusFlags(statusFlags BACnetStatusFlagsTagged) BACnetNotificationParametersChangeOfLifeSafetyBuilder {
	b.StatusFlags = statusFlags
	return b
}

func (b *_BACnetNotificationParametersChangeOfLifeSafetyBuilder) WithStatusFlagsBuilder(builderSupplier func(BACnetStatusFlagsTaggedBuilder) BACnetStatusFlagsTaggedBuilder) BACnetNotificationParametersChangeOfLifeSafetyBuilder {
	builder := builderSupplier(b.StatusFlags.CreateBACnetStatusFlagsTaggedBuilder())
	var err error
	b.StatusFlags, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetStatusFlagsTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersChangeOfLifeSafetyBuilder) WithOperationExpected(operationExpected BACnetLifeSafetyOperationTagged) BACnetNotificationParametersChangeOfLifeSafetyBuilder {
	b.OperationExpected = operationExpected
	return b
}

func (b *_BACnetNotificationParametersChangeOfLifeSafetyBuilder) WithOperationExpectedBuilder(builderSupplier func(BACnetLifeSafetyOperationTaggedBuilder) BACnetLifeSafetyOperationTaggedBuilder) BACnetNotificationParametersChangeOfLifeSafetyBuilder {
	builder := builderSupplier(b.OperationExpected.CreateBACnetLifeSafetyOperationTaggedBuilder())
	var err error
	b.OperationExpected, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLifeSafetyOperationTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersChangeOfLifeSafetyBuilder) WithInnerClosingTag(innerClosingTag BACnetClosingTag) BACnetNotificationParametersChangeOfLifeSafetyBuilder {
	b.InnerClosingTag = innerClosingTag
	return b
}

func (b *_BACnetNotificationParametersChangeOfLifeSafetyBuilder) WithInnerClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersChangeOfLifeSafetyBuilder {
	builder := builderSupplier(b.InnerClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.InnerClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersChangeOfLifeSafetyBuilder) Build() (BACnetNotificationParametersChangeOfLifeSafety, error) {
	if b.InnerOpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'innerOpeningTag' not set"))
	}
	if b.NewState == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'newState' not set"))
	}
	if b.NewMode == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'newMode' not set"))
	}
	if b.StatusFlags == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'statusFlags' not set"))
	}
	if b.OperationExpected == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'operationExpected' not set"))
	}
	if b.InnerClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'innerClosingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetNotificationParametersChangeOfLifeSafety.deepCopy(), nil
}

func (b *_BACnetNotificationParametersChangeOfLifeSafetyBuilder) MustBuild() BACnetNotificationParametersChangeOfLifeSafety {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetNotificationParametersChangeOfLifeSafetyBuilder) Done() BACnetNotificationParametersBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetNotificationParametersBuilder().(*_BACnetNotificationParametersBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetNotificationParametersChangeOfLifeSafetyBuilder) buildForBACnetNotificationParameters() (BACnetNotificationParameters, error) {
	return b.Build()
}

func (b *_BACnetNotificationParametersChangeOfLifeSafetyBuilder) DeepCopy() any {
	_copy := b.CreateBACnetNotificationParametersChangeOfLifeSafetyBuilder().(*_BACnetNotificationParametersChangeOfLifeSafetyBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetNotificationParametersChangeOfLifeSafetyBuilder creates a BACnetNotificationParametersChangeOfLifeSafetyBuilder
func (b *_BACnetNotificationParametersChangeOfLifeSafety) CreateBACnetNotificationParametersChangeOfLifeSafetyBuilder() BACnetNotificationParametersChangeOfLifeSafetyBuilder {
	if b == nil {
		return NewBACnetNotificationParametersChangeOfLifeSafetyBuilder()
	}
	return &_BACnetNotificationParametersChangeOfLifeSafetyBuilder{_BACnetNotificationParametersChangeOfLifeSafety: b.deepCopy()}
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

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetParent() BACnetNotificationParametersContract {
	return m.BACnetNotificationParametersContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetNewState() BACnetLifeSafetyStateTagged {
	return m.NewState
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetNewMode() BACnetLifeSafetyModeTagged {
	return m.NewMode
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetOperationExpected() BACnetLifeSafetyOperationTagged {
	return m.OperationExpected
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfLifeSafety(structType any) BACnetNotificationParametersChangeOfLifeSafety {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfLifeSafety); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfLifeSafety); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfLifeSafety"
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).getLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Simple field (newState)
	lengthInBits += m.NewState.GetLengthInBits(ctx)

	// Simple field (newMode)
	lengthInBits += m.NewMode.GetLengthInBits(ctx)

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits(ctx)

	// Simple field (operationExpected)
	lengthInBits += m.OperationExpected.GetLengthInBits(ctx)

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParameters, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (__bACnetNotificationParametersChangeOfLifeSafety BACnetNotificationParametersChangeOfLifeSafety, err error) {
	m.BACnetNotificationParametersContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfLifeSafety"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfLifeSafety")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	innerOpeningTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerOpeningTag' field"))
	}
	m.InnerOpeningTag = innerOpeningTag

	newState, err := ReadSimpleField[BACnetLifeSafetyStateTagged](ctx, "newState", ReadComplex[BACnetLifeSafetyStateTagged](BACnetLifeSafetyStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'newState' field"))
	}
	m.NewState = newState

	newMode, err := ReadSimpleField[BACnetLifeSafetyModeTagged](ctx, "newMode", ReadComplex[BACnetLifeSafetyModeTagged](BACnetLifeSafetyModeTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'newMode' field"))
	}
	m.NewMode = newMode

	statusFlags, err := ReadSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", ReadComplex[BACnetStatusFlagsTagged](BACnetStatusFlagsTaggedParseWithBufferProducer((uint8)(uint8(2)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusFlags' field"))
	}
	m.StatusFlags = statusFlags

	operationExpected, err := ReadSimpleField[BACnetLifeSafetyOperationTagged](ctx, "operationExpected", ReadComplex[BACnetLifeSafetyOperationTagged](BACnetLifeSafetyOperationTaggedParseWithBufferProducer((uint8)(uint8(3)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'operationExpected' field"))
	}
	m.OperationExpected = operationExpected

	innerClosingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "innerClosingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerClosingTag' field"))
	}
	m.InnerClosingTag = innerClosingTag

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfLifeSafety"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfLifeSafety")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfLifeSafety"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfLifeSafety")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", m.GetInnerOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerOpeningTag' field")
		}

		if err := WriteSimpleField[BACnetLifeSafetyStateTagged](ctx, "newState", m.GetNewState(), WriteComplex[BACnetLifeSafetyStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'newState' field")
		}

		if err := WriteSimpleField[BACnetLifeSafetyModeTagged](ctx, "newMode", m.GetNewMode(), WriteComplex[BACnetLifeSafetyModeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'newMode' field")
		}

		if err := WriteSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", m.GetStatusFlags(), WriteComplex[BACnetStatusFlagsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusFlags' field")
		}

		if err := WriteSimpleField[BACnetLifeSafetyOperationTagged](ctx, "operationExpected", m.GetOperationExpected(), WriteComplex[BACnetLifeSafetyOperationTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'operationExpected' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "innerClosingTag", m.GetInnerClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfLifeSafety"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfLifeSafety")
		}
		return nil
	}
	return m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) IsBACnetNotificationParametersChangeOfLifeSafety() {
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) deepCopy() *_BACnetNotificationParametersChangeOfLifeSafety {
	if m == nil {
		return nil
	}
	_BACnetNotificationParametersChangeOfLifeSafetyCopy := &_BACnetNotificationParametersChangeOfLifeSafety{
		m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).deepCopy(),
		utils.DeepCopy[BACnetOpeningTag](m.InnerOpeningTag),
		utils.DeepCopy[BACnetLifeSafetyStateTagged](m.NewState),
		utils.DeepCopy[BACnetLifeSafetyModeTagged](m.NewMode),
		utils.DeepCopy[BACnetStatusFlagsTagged](m.StatusFlags),
		utils.DeepCopy[BACnetLifeSafetyOperationTagged](m.OperationExpected),
		utils.DeepCopy[BACnetClosingTag](m.InnerClosingTag),
	}
	_BACnetNotificationParametersChangeOfLifeSafetyCopy.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = m
	return _BACnetNotificationParametersChangeOfLifeSafetyCopy
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) String() string {
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