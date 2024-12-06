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

// BACnetConstructedDataAnalogInputAll is the corresponding interface of BACnetConstructedDataAnalogInputAll
type BACnetConstructedDataAnalogInputAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataAnalogInputAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAnalogInputAll()
	// CreateBuilder creates a BACnetConstructedDataAnalogInputAllBuilder
	CreateBACnetConstructedDataAnalogInputAllBuilder() BACnetConstructedDataAnalogInputAllBuilder
}

// _BACnetConstructedDataAnalogInputAll is the data-structure of this message
type _BACnetConstructedDataAnalogInputAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataAnalogInputAll = (*_BACnetConstructedDataAnalogInputAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAnalogInputAll)(nil)

// NewBACnetConstructedDataAnalogInputAll factory function for _BACnetConstructedDataAnalogInputAll
func NewBACnetConstructedDataAnalogInputAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAnalogInputAll {
	_result := &_BACnetConstructedDataAnalogInputAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAnalogInputAllBuilder is a builder for BACnetConstructedDataAnalogInputAll
type BACnetConstructedDataAnalogInputAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataAnalogInputAllBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataAnalogInputAll or returns an error if something is wrong
	Build() (BACnetConstructedDataAnalogInputAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAnalogInputAll
}

// NewBACnetConstructedDataAnalogInputAllBuilder() creates a BACnetConstructedDataAnalogInputAllBuilder
func NewBACnetConstructedDataAnalogInputAllBuilder() BACnetConstructedDataAnalogInputAllBuilder {
	return &_BACnetConstructedDataAnalogInputAllBuilder{_BACnetConstructedDataAnalogInputAll: new(_BACnetConstructedDataAnalogInputAll)}
}

type _BACnetConstructedDataAnalogInputAllBuilder struct {
	*_BACnetConstructedDataAnalogInputAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAnalogInputAllBuilder) = (*_BACnetConstructedDataAnalogInputAllBuilder)(nil)

func (b *_BACnetConstructedDataAnalogInputAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataAnalogInputAll
}

func (b *_BACnetConstructedDataAnalogInputAllBuilder) WithMandatoryFields() BACnetConstructedDataAnalogInputAllBuilder {
	return b
}

func (b *_BACnetConstructedDataAnalogInputAllBuilder) Build() (BACnetConstructedDataAnalogInputAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAnalogInputAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataAnalogInputAllBuilder) MustBuild() BACnetConstructedDataAnalogInputAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataAnalogInputAllBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAnalogInputAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAnalogInputAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAnalogInputAllBuilder().(*_BACnetConstructedDataAnalogInputAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAnalogInputAllBuilder creates a BACnetConstructedDataAnalogInputAllBuilder
func (b *_BACnetConstructedDataAnalogInputAll) CreateBACnetConstructedDataAnalogInputAllBuilder() BACnetConstructedDataAnalogInputAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataAnalogInputAllBuilder()
	}
	return &_BACnetConstructedDataAnalogInputAllBuilder{_BACnetConstructedDataAnalogInputAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAnalogInputAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ANALOG_INPUT
}

func (m *_BACnetConstructedDataAnalogInputAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAnalogInputAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAnalogInputAll(structType any) BACnetConstructedDataAnalogInputAll {
	if casted, ok := structType.(BACnetConstructedDataAnalogInputAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAnalogInputAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAnalogInputAll) GetTypeName() string {
	return "BACnetConstructedDataAnalogInputAll"
}

func (m *_BACnetConstructedDataAnalogInputAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataAnalogInputAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAnalogInputAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAnalogInputAll BACnetConstructedDataAnalogInputAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAnalogInputAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAnalogInputAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAnalogInputAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAnalogInputAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAnalogInputAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAnalogInputAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAnalogInputAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAnalogInputAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAnalogInputAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAnalogInputAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAnalogInputAll) IsBACnetConstructedDataAnalogInputAll() {}

func (m *_BACnetConstructedDataAnalogInputAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAnalogInputAll) deepCopy() *_BACnetConstructedDataAnalogInputAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAnalogInputAllCopy := &_BACnetConstructedDataAnalogInputAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	_BACnetConstructedDataAnalogInputAllCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAnalogInputAllCopy
}

func (m *_BACnetConstructedDataAnalogInputAll) String() string {
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