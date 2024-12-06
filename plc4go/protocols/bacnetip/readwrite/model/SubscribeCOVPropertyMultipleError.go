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

// SubscribeCOVPropertyMultipleError is the corresponding interface of SubscribeCOVPropertyMultipleError
type SubscribeCOVPropertyMultipleError interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetError
	// GetErrorType returns ErrorType (property field)
	GetErrorType() ErrorEnclosed
	// GetFirstFailedSubscription returns FirstFailedSubscription (property field)
	GetFirstFailedSubscription() SubscribeCOVPropertyMultipleErrorFirstFailedSubscription
	// IsSubscribeCOVPropertyMultipleError is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSubscribeCOVPropertyMultipleError()
	// CreateBuilder creates a SubscribeCOVPropertyMultipleErrorBuilder
	CreateSubscribeCOVPropertyMultipleErrorBuilder() SubscribeCOVPropertyMultipleErrorBuilder
}

// _SubscribeCOVPropertyMultipleError is the data-structure of this message
type _SubscribeCOVPropertyMultipleError struct {
	BACnetErrorContract
	ErrorType               ErrorEnclosed
	FirstFailedSubscription SubscribeCOVPropertyMultipleErrorFirstFailedSubscription
}

var _ SubscribeCOVPropertyMultipleError = (*_SubscribeCOVPropertyMultipleError)(nil)
var _ BACnetErrorRequirements = (*_SubscribeCOVPropertyMultipleError)(nil)

// NewSubscribeCOVPropertyMultipleError factory function for _SubscribeCOVPropertyMultipleError
func NewSubscribeCOVPropertyMultipleError(errorType ErrorEnclosed, firstFailedSubscription SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) *_SubscribeCOVPropertyMultipleError {
	if errorType == nil {
		panic("errorType of type ErrorEnclosed for SubscribeCOVPropertyMultipleError must not be nil")
	}
	if firstFailedSubscription == nil {
		panic("firstFailedSubscription of type SubscribeCOVPropertyMultipleErrorFirstFailedSubscription for SubscribeCOVPropertyMultipleError must not be nil")
	}
	_result := &_SubscribeCOVPropertyMultipleError{
		BACnetErrorContract:     NewBACnetError(),
		ErrorType:               errorType,
		FirstFailedSubscription: firstFailedSubscription,
	}
	_result.BACnetErrorContract.(*_BACnetError)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SubscribeCOVPropertyMultipleErrorBuilder is a builder for SubscribeCOVPropertyMultipleError
type SubscribeCOVPropertyMultipleErrorBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(errorType ErrorEnclosed, firstFailedSubscription SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) SubscribeCOVPropertyMultipleErrorBuilder
	// WithErrorType adds ErrorType (property field)
	WithErrorType(ErrorEnclosed) SubscribeCOVPropertyMultipleErrorBuilder
	// WithErrorTypeBuilder adds ErrorType (property field) which is build by the builder
	WithErrorTypeBuilder(func(ErrorEnclosedBuilder) ErrorEnclosedBuilder) SubscribeCOVPropertyMultipleErrorBuilder
	// WithFirstFailedSubscription adds FirstFailedSubscription (property field)
	WithFirstFailedSubscription(SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) SubscribeCOVPropertyMultipleErrorBuilder
	// WithFirstFailedSubscriptionBuilder adds FirstFailedSubscription (property field) which is build by the builder
	WithFirstFailedSubscriptionBuilder(func(SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder) SubscribeCOVPropertyMultipleErrorBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetErrorBuilder
	// Build builds the SubscribeCOVPropertyMultipleError or returns an error if something is wrong
	Build() (SubscribeCOVPropertyMultipleError, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SubscribeCOVPropertyMultipleError
}

// NewSubscribeCOVPropertyMultipleErrorBuilder() creates a SubscribeCOVPropertyMultipleErrorBuilder
func NewSubscribeCOVPropertyMultipleErrorBuilder() SubscribeCOVPropertyMultipleErrorBuilder {
	return &_SubscribeCOVPropertyMultipleErrorBuilder{_SubscribeCOVPropertyMultipleError: new(_SubscribeCOVPropertyMultipleError)}
}

type _SubscribeCOVPropertyMultipleErrorBuilder struct {
	*_SubscribeCOVPropertyMultipleError

	parentBuilder *_BACnetErrorBuilder

	err *utils.MultiError
}

var _ (SubscribeCOVPropertyMultipleErrorBuilder) = (*_SubscribeCOVPropertyMultipleErrorBuilder)(nil)

func (b *_SubscribeCOVPropertyMultipleErrorBuilder) setParent(contract BACnetErrorContract) {
	b.BACnetErrorContract = contract
	contract.(*_BACnetError)._SubType = b._SubscribeCOVPropertyMultipleError
}

func (b *_SubscribeCOVPropertyMultipleErrorBuilder) WithMandatoryFields(errorType ErrorEnclosed, firstFailedSubscription SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) SubscribeCOVPropertyMultipleErrorBuilder {
	return b.WithErrorType(errorType).WithFirstFailedSubscription(firstFailedSubscription)
}

func (b *_SubscribeCOVPropertyMultipleErrorBuilder) WithErrorType(errorType ErrorEnclosed) SubscribeCOVPropertyMultipleErrorBuilder {
	b.ErrorType = errorType
	return b
}

func (b *_SubscribeCOVPropertyMultipleErrorBuilder) WithErrorTypeBuilder(builderSupplier func(ErrorEnclosedBuilder) ErrorEnclosedBuilder) SubscribeCOVPropertyMultipleErrorBuilder {
	builder := builderSupplier(b.ErrorType.CreateErrorEnclosedBuilder())
	var err error
	b.ErrorType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ErrorEnclosedBuilder failed"))
	}
	return b
}

func (b *_SubscribeCOVPropertyMultipleErrorBuilder) WithFirstFailedSubscription(firstFailedSubscription SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) SubscribeCOVPropertyMultipleErrorBuilder {
	b.FirstFailedSubscription = firstFailedSubscription
	return b
}

func (b *_SubscribeCOVPropertyMultipleErrorBuilder) WithFirstFailedSubscriptionBuilder(builderSupplier func(SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder) SubscribeCOVPropertyMultipleErrorBuilder {
	builder := builderSupplier(b.FirstFailedSubscription.CreateSubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder())
	var err error
	b.FirstFailedSubscription, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder failed"))
	}
	return b
}

func (b *_SubscribeCOVPropertyMultipleErrorBuilder) Build() (SubscribeCOVPropertyMultipleError, error) {
	if b.ErrorType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'errorType' not set"))
	}
	if b.FirstFailedSubscription == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'firstFailedSubscription' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SubscribeCOVPropertyMultipleError.deepCopy(), nil
}

func (b *_SubscribeCOVPropertyMultipleErrorBuilder) MustBuild() SubscribeCOVPropertyMultipleError {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SubscribeCOVPropertyMultipleErrorBuilder) Done() BACnetErrorBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetErrorBuilder().(*_BACnetErrorBuilder)
	}
	return b.parentBuilder
}

func (b *_SubscribeCOVPropertyMultipleErrorBuilder) buildForBACnetError() (BACnetError, error) {
	return b.Build()
}

func (b *_SubscribeCOVPropertyMultipleErrorBuilder) DeepCopy() any {
	_copy := b.CreateSubscribeCOVPropertyMultipleErrorBuilder().(*_SubscribeCOVPropertyMultipleErrorBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSubscribeCOVPropertyMultipleErrorBuilder creates a SubscribeCOVPropertyMultipleErrorBuilder
func (b *_SubscribeCOVPropertyMultipleError) CreateSubscribeCOVPropertyMultipleErrorBuilder() SubscribeCOVPropertyMultipleErrorBuilder {
	if b == nil {
		return NewSubscribeCOVPropertyMultipleErrorBuilder()
	}
	return &_SubscribeCOVPropertyMultipleErrorBuilder{_SubscribeCOVPropertyMultipleError: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SubscribeCOVPropertyMultipleError) GetErrorChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SubscribeCOVPropertyMultipleError) GetParent() BACnetErrorContract {
	return m.BACnetErrorContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SubscribeCOVPropertyMultipleError) GetErrorType() ErrorEnclosed {
	return m.ErrorType
}

func (m *_SubscribeCOVPropertyMultipleError) GetFirstFailedSubscription() SubscribeCOVPropertyMultipleErrorFirstFailedSubscription {
	return m.FirstFailedSubscription
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSubscribeCOVPropertyMultipleError(structType any) SubscribeCOVPropertyMultipleError {
	if casted, ok := structType.(SubscribeCOVPropertyMultipleError); ok {
		return casted
	}
	if casted, ok := structType.(*SubscribeCOVPropertyMultipleError); ok {
		return *casted
	}
	return nil
}

func (m *_SubscribeCOVPropertyMultipleError) GetTypeName() string {
	return "SubscribeCOVPropertyMultipleError"
}

func (m *_SubscribeCOVPropertyMultipleError) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetErrorContract.(*_BACnetError).getLengthInBits(ctx))

	// Simple field (errorType)
	lengthInBits += m.ErrorType.GetLengthInBits(ctx)

	// Simple field (firstFailedSubscription)
	lengthInBits += m.FirstFailedSubscription.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SubscribeCOVPropertyMultipleError) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SubscribeCOVPropertyMultipleError) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetError, errorChoice BACnetConfirmedServiceChoice) (__subscribeCOVPropertyMultipleError SubscribeCOVPropertyMultipleError, err error) {
	m.BACnetErrorContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SubscribeCOVPropertyMultipleError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SubscribeCOVPropertyMultipleError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	errorType, err := ReadSimpleField[ErrorEnclosed](ctx, "errorType", ReadComplex[ErrorEnclosed](ErrorEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorType' field"))
	}
	m.ErrorType = errorType

	firstFailedSubscription, err := ReadSimpleField[SubscribeCOVPropertyMultipleErrorFirstFailedSubscription](ctx, "firstFailedSubscription", ReadComplex[SubscribeCOVPropertyMultipleErrorFirstFailedSubscription](SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'firstFailedSubscription' field"))
	}
	m.FirstFailedSubscription = firstFailedSubscription

	if closeErr := readBuffer.CloseContext("SubscribeCOVPropertyMultipleError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SubscribeCOVPropertyMultipleError")
	}

	return m, nil
}

func (m *_SubscribeCOVPropertyMultipleError) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SubscribeCOVPropertyMultipleError) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SubscribeCOVPropertyMultipleError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SubscribeCOVPropertyMultipleError")
		}

		if err := WriteSimpleField[ErrorEnclosed](ctx, "errorType", m.GetErrorType(), WriteComplex[ErrorEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'errorType' field")
		}

		if err := WriteSimpleField[SubscribeCOVPropertyMultipleErrorFirstFailedSubscription](ctx, "firstFailedSubscription", m.GetFirstFailedSubscription(), WriteComplex[SubscribeCOVPropertyMultipleErrorFirstFailedSubscription](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'firstFailedSubscription' field")
		}

		if popErr := writeBuffer.PopContext("SubscribeCOVPropertyMultipleError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SubscribeCOVPropertyMultipleError")
		}
		return nil
	}
	return m.BACnetErrorContract.(*_BACnetError).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SubscribeCOVPropertyMultipleError) IsSubscribeCOVPropertyMultipleError() {}

func (m *_SubscribeCOVPropertyMultipleError) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SubscribeCOVPropertyMultipleError) deepCopy() *_SubscribeCOVPropertyMultipleError {
	if m == nil {
		return nil
	}
	_SubscribeCOVPropertyMultipleErrorCopy := &_SubscribeCOVPropertyMultipleError{
		m.BACnetErrorContract.(*_BACnetError).deepCopy(),
		utils.DeepCopy[ErrorEnclosed](m.ErrorType),
		utils.DeepCopy[SubscribeCOVPropertyMultipleErrorFirstFailedSubscription](m.FirstFailedSubscription),
	}
	_SubscribeCOVPropertyMultipleErrorCopy.BACnetErrorContract.(*_BACnetError)._SubType = m
	return _SubscribeCOVPropertyMultipleErrorCopy
}

func (m *_SubscribeCOVPropertyMultipleError) String() string {
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