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

// BACnetConstructedDataPulseConverterAll is the corresponding interface of BACnetConstructedDataPulseConverterAll
type BACnetConstructedDataPulseConverterAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataPulseConverterAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataPulseConverterAll()
	// CreateBuilder creates a BACnetConstructedDataPulseConverterAllBuilder
	CreateBACnetConstructedDataPulseConverterAllBuilder() BACnetConstructedDataPulseConverterAllBuilder
}

// _BACnetConstructedDataPulseConverterAll is the data-structure of this message
type _BACnetConstructedDataPulseConverterAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataPulseConverterAll = (*_BACnetConstructedDataPulseConverterAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataPulseConverterAll)(nil)

// NewBACnetConstructedDataPulseConverterAll factory function for _BACnetConstructedDataPulseConverterAll
func NewBACnetConstructedDataPulseConverterAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPulseConverterAll {
	_result := &_BACnetConstructedDataPulseConverterAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataPulseConverterAllBuilder is a builder for BACnetConstructedDataPulseConverterAll
type BACnetConstructedDataPulseConverterAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataPulseConverterAllBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataPulseConverterAll or returns an error if something is wrong
	Build() (BACnetConstructedDataPulseConverterAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataPulseConverterAll
}

// NewBACnetConstructedDataPulseConverterAllBuilder() creates a BACnetConstructedDataPulseConverterAllBuilder
func NewBACnetConstructedDataPulseConverterAllBuilder() BACnetConstructedDataPulseConverterAllBuilder {
	return &_BACnetConstructedDataPulseConverterAllBuilder{_BACnetConstructedDataPulseConverterAll: new(_BACnetConstructedDataPulseConverterAll)}
}

type _BACnetConstructedDataPulseConverterAllBuilder struct {
	*_BACnetConstructedDataPulseConverterAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataPulseConverterAllBuilder) = (*_BACnetConstructedDataPulseConverterAllBuilder)(nil)

func (b *_BACnetConstructedDataPulseConverterAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataPulseConverterAll
}

func (b *_BACnetConstructedDataPulseConverterAllBuilder) WithMandatoryFields() BACnetConstructedDataPulseConverterAllBuilder {
	return b
}

func (b *_BACnetConstructedDataPulseConverterAllBuilder) Build() (BACnetConstructedDataPulseConverterAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataPulseConverterAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataPulseConverterAllBuilder) MustBuild() BACnetConstructedDataPulseConverterAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataPulseConverterAllBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataPulseConverterAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataPulseConverterAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataPulseConverterAllBuilder().(*_BACnetConstructedDataPulseConverterAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataPulseConverterAllBuilder creates a BACnetConstructedDataPulseConverterAllBuilder
func (b *_BACnetConstructedDataPulseConverterAll) CreateBACnetConstructedDataPulseConverterAllBuilder() BACnetConstructedDataPulseConverterAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataPulseConverterAllBuilder()
	}
	return &_BACnetConstructedDataPulseConverterAllBuilder{_BACnetConstructedDataPulseConverterAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPulseConverterAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_PULSE_CONVERTER
}

func (m *_BACnetConstructedDataPulseConverterAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPulseConverterAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPulseConverterAll(structType any) BACnetConstructedDataPulseConverterAll {
	if casted, ok := structType.(BACnetConstructedDataPulseConverterAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPulseConverterAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPulseConverterAll) GetTypeName() string {
	return "BACnetConstructedDataPulseConverterAll"
}

func (m *_BACnetConstructedDataPulseConverterAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataPulseConverterAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataPulseConverterAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataPulseConverterAll BACnetConstructedDataPulseConverterAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPulseConverterAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPulseConverterAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPulseConverterAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPulseConverterAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataPulseConverterAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPulseConverterAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPulseConverterAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPulseConverterAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPulseConverterAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPulseConverterAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPulseConverterAll) IsBACnetConstructedDataPulseConverterAll() {}

func (m *_BACnetConstructedDataPulseConverterAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataPulseConverterAll) deepCopy() *_BACnetConstructedDataPulseConverterAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataPulseConverterAllCopy := &_BACnetConstructedDataPulseConverterAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	_BACnetConstructedDataPulseConverterAllCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataPulseConverterAllCopy
}

func (m *_BACnetConstructedDataPulseConverterAll) String() string {
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