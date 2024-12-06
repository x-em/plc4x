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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataTimerAll is the corresponding interface of BACnetConstructedDataTimerAll
type BACnetConstructedDataTimerAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataTimerAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTimerAll()
	// CreateBuilder creates a BACnetConstructedDataTimerAllBuilder
	CreateBACnetConstructedDataTimerAllBuilder() BACnetConstructedDataTimerAllBuilder
}

// _BACnetConstructedDataTimerAll is the data-structure of this message
type _BACnetConstructedDataTimerAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataTimerAll = (*_BACnetConstructedDataTimerAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTimerAll)(nil)

// NewBACnetConstructedDataTimerAll factory function for _BACnetConstructedDataTimerAll
func NewBACnetConstructedDataTimerAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimerAll {
	_result := &_BACnetConstructedDataTimerAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataTimerAllBuilder is a builder for BACnetConstructedDataTimerAll
type BACnetConstructedDataTimerAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataTimerAllBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataTimerAll or returns an error if something is wrong
	Build() (BACnetConstructedDataTimerAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataTimerAll
}

// NewBACnetConstructedDataTimerAllBuilder() creates a BACnetConstructedDataTimerAllBuilder
func NewBACnetConstructedDataTimerAllBuilder() BACnetConstructedDataTimerAllBuilder {
	return &_BACnetConstructedDataTimerAllBuilder{_BACnetConstructedDataTimerAll: new(_BACnetConstructedDataTimerAll)}
}

type _BACnetConstructedDataTimerAllBuilder struct {
	*_BACnetConstructedDataTimerAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataTimerAllBuilder) = (*_BACnetConstructedDataTimerAllBuilder)(nil)

func (b *_BACnetConstructedDataTimerAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataTimerAll
}

func (b *_BACnetConstructedDataTimerAllBuilder) WithMandatoryFields() BACnetConstructedDataTimerAllBuilder {
	return b
}

func (b *_BACnetConstructedDataTimerAllBuilder) Build() (BACnetConstructedDataTimerAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataTimerAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataTimerAllBuilder) MustBuild() BACnetConstructedDataTimerAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataTimerAllBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataTimerAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataTimerAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataTimerAllBuilder().(*_BACnetConstructedDataTimerAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataTimerAllBuilder creates a BACnetConstructedDataTimerAllBuilder
func (b *_BACnetConstructedDataTimerAll) CreateBACnetConstructedDataTimerAllBuilder() BACnetConstructedDataTimerAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataTimerAllBuilder()
	}
	return &_BACnetConstructedDataTimerAllBuilder{_BACnetConstructedDataTimerAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimerAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TIMER
}

func (m *_BACnetConstructedDataTimerAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimerAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimerAll(structType any) BACnetConstructedDataTimerAll {
	if casted, ok := structType.(BACnetConstructedDataTimerAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimerAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimerAll) GetTypeName() string {
	return "BACnetConstructedDataTimerAll"
}

func (m *_BACnetConstructedDataTimerAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataTimerAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTimerAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTimerAll BACnetConstructedDataTimerAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimerAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimerAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimerAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimerAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTimerAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimerAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimerAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimerAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimerAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimerAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimerAll) IsBACnetConstructedDataTimerAll() {}

func (m *_BACnetConstructedDataTimerAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataTimerAll) deepCopy() *_BACnetConstructedDataTimerAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataTimerAllCopy := &_BACnetConstructedDataTimerAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	_BACnetConstructedDataTimerAllCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataTimerAllCopy
}

func (m *_BACnetConstructedDataTimerAll) String() string {
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
