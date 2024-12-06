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

// BACnetConstructedDataLifeSafetyPointAll is the corresponding interface of BACnetConstructedDataLifeSafetyPointAll
type BACnetConstructedDataLifeSafetyPointAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataLifeSafetyPointAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLifeSafetyPointAll()
	// CreateBuilder creates a BACnetConstructedDataLifeSafetyPointAllBuilder
	CreateBACnetConstructedDataLifeSafetyPointAllBuilder() BACnetConstructedDataLifeSafetyPointAllBuilder
}

// _BACnetConstructedDataLifeSafetyPointAll is the data-structure of this message
type _BACnetConstructedDataLifeSafetyPointAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataLifeSafetyPointAll = (*_BACnetConstructedDataLifeSafetyPointAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLifeSafetyPointAll)(nil)

// NewBACnetConstructedDataLifeSafetyPointAll factory function for _BACnetConstructedDataLifeSafetyPointAll
func NewBACnetConstructedDataLifeSafetyPointAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLifeSafetyPointAll {
	_result := &_BACnetConstructedDataLifeSafetyPointAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLifeSafetyPointAllBuilder is a builder for BACnetConstructedDataLifeSafetyPointAll
type BACnetConstructedDataLifeSafetyPointAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataLifeSafetyPointAllBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLifeSafetyPointAll or returns an error if something is wrong
	Build() (BACnetConstructedDataLifeSafetyPointAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLifeSafetyPointAll
}

// NewBACnetConstructedDataLifeSafetyPointAllBuilder() creates a BACnetConstructedDataLifeSafetyPointAllBuilder
func NewBACnetConstructedDataLifeSafetyPointAllBuilder() BACnetConstructedDataLifeSafetyPointAllBuilder {
	return &_BACnetConstructedDataLifeSafetyPointAllBuilder{_BACnetConstructedDataLifeSafetyPointAll: new(_BACnetConstructedDataLifeSafetyPointAll)}
}

type _BACnetConstructedDataLifeSafetyPointAllBuilder struct {
	*_BACnetConstructedDataLifeSafetyPointAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLifeSafetyPointAllBuilder) = (*_BACnetConstructedDataLifeSafetyPointAllBuilder)(nil)

func (b *_BACnetConstructedDataLifeSafetyPointAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataLifeSafetyPointAll
}

func (b *_BACnetConstructedDataLifeSafetyPointAllBuilder) WithMandatoryFields() BACnetConstructedDataLifeSafetyPointAllBuilder {
	return b
}

func (b *_BACnetConstructedDataLifeSafetyPointAllBuilder) Build() (BACnetConstructedDataLifeSafetyPointAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLifeSafetyPointAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataLifeSafetyPointAllBuilder) MustBuild() BACnetConstructedDataLifeSafetyPointAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLifeSafetyPointAllBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLifeSafetyPointAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLifeSafetyPointAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLifeSafetyPointAllBuilder().(*_BACnetConstructedDataLifeSafetyPointAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLifeSafetyPointAllBuilder creates a BACnetConstructedDataLifeSafetyPointAllBuilder
func (b *_BACnetConstructedDataLifeSafetyPointAll) CreateBACnetConstructedDataLifeSafetyPointAllBuilder() BACnetConstructedDataLifeSafetyPointAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataLifeSafetyPointAllBuilder()
	}
	return &_BACnetConstructedDataLifeSafetyPointAllBuilder{_BACnetConstructedDataLifeSafetyPointAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyPointAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LIFE_SAFETY_POINT
}

func (m *_BACnetConstructedDataLifeSafetyPointAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLifeSafetyPointAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLifeSafetyPointAll(structType any) BACnetConstructedDataLifeSafetyPointAll {
	if casted, ok := structType.(BACnetConstructedDataLifeSafetyPointAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLifeSafetyPointAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLifeSafetyPointAll) GetTypeName() string {
	return "BACnetConstructedDataLifeSafetyPointAll"
}

func (m *_BACnetConstructedDataLifeSafetyPointAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataLifeSafetyPointAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLifeSafetyPointAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLifeSafetyPointAll BACnetConstructedDataLifeSafetyPointAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLifeSafetyPointAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLifeSafetyPointAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLifeSafetyPointAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLifeSafetyPointAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLifeSafetyPointAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLifeSafetyPointAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLifeSafetyPointAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLifeSafetyPointAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLifeSafetyPointAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLifeSafetyPointAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLifeSafetyPointAll) IsBACnetConstructedDataLifeSafetyPointAll() {}

func (m *_BACnetConstructedDataLifeSafetyPointAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLifeSafetyPointAll) deepCopy() *_BACnetConstructedDataLifeSafetyPointAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLifeSafetyPointAllCopy := &_BACnetConstructedDataLifeSafetyPointAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	_BACnetConstructedDataLifeSafetyPointAllCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLifeSafetyPointAllCopy
}

func (m *_BACnetConstructedDataLifeSafetyPointAll) String() string {
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
