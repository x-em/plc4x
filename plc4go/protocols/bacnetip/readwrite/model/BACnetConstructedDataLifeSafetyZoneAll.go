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

// BACnetConstructedDataLifeSafetyZoneAll is the corresponding interface of BACnetConstructedDataLifeSafetyZoneAll
type BACnetConstructedDataLifeSafetyZoneAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataLifeSafetyZoneAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLifeSafetyZoneAll()
	// CreateBuilder creates a BACnetConstructedDataLifeSafetyZoneAllBuilder
	CreateBACnetConstructedDataLifeSafetyZoneAllBuilder() BACnetConstructedDataLifeSafetyZoneAllBuilder
}

// _BACnetConstructedDataLifeSafetyZoneAll is the data-structure of this message
type _BACnetConstructedDataLifeSafetyZoneAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataLifeSafetyZoneAll = (*_BACnetConstructedDataLifeSafetyZoneAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLifeSafetyZoneAll)(nil)

// NewBACnetConstructedDataLifeSafetyZoneAll factory function for _BACnetConstructedDataLifeSafetyZoneAll
func NewBACnetConstructedDataLifeSafetyZoneAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLifeSafetyZoneAll {
	_result := &_BACnetConstructedDataLifeSafetyZoneAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLifeSafetyZoneAllBuilder is a builder for BACnetConstructedDataLifeSafetyZoneAll
type BACnetConstructedDataLifeSafetyZoneAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataLifeSafetyZoneAllBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLifeSafetyZoneAll or returns an error if something is wrong
	Build() (BACnetConstructedDataLifeSafetyZoneAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLifeSafetyZoneAll
}

// NewBACnetConstructedDataLifeSafetyZoneAllBuilder() creates a BACnetConstructedDataLifeSafetyZoneAllBuilder
func NewBACnetConstructedDataLifeSafetyZoneAllBuilder() BACnetConstructedDataLifeSafetyZoneAllBuilder {
	return &_BACnetConstructedDataLifeSafetyZoneAllBuilder{_BACnetConstructedDataLifeSafetyZoneAll: new(_BACnetConstructedDataLifeSafetyZoneAll)}
}

type _BACnetConstructedDataLifeSafetyZoneAllBuilder struct {
	*_BACnetConstructedDataLifeSafetyZoneAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLifeSafetyZoneAllBuilder) = (*_BACnetConstructedDataLifeSafetyZoneAllBuilder)(nil)

func (b *_BACnetConstructedDataLifeSafetyZoneAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataLifeSafetyZoneAll
}

func (b *_BACnetConstructedDataLifeSafetyZoneAllBuilder) WithMandatoryFields() BACnetConstructedDataLifeSafetyZoneAllBuilder {
	return b
}

func (b *_BACnetConstructedDataLifeSafetyZoneAllBuilder) Build() (BACnetConstructedDataLifeSafetyZoneAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLifeSafetyZoneAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataLifeSafetyZoneAllBuilder) MustBuild() BACnetConstructedDataLifeSafetyZoneAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLifeSafetyZoneAllBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLifeSafetyZoneAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLifeSafetyZoneAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLifeSafetyZoneAllBuilder().(*_BACnetConstructedDataLifeSafetyZoneAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLifeSafetyZoneAllBuilder creates a BACnetConstructedDataLifeSafetyZoneAllBuilder
func (b *_BACnetConstructedDataLifeSafetyZoneAll) CreateBACnetConstructedDataLifeSafetyZoneAllBuilder() BACnetConstructedDataLifeSafetyZoneAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataLifeSafetyZoneAllBuilder()
	}
	return &_BACnetConstructedDataLifeSafetyZoneAllBuilder{_BACnetConstructedDataLifeSafetyZoneAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyZoneAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LIFE_SAFETY_ZONE
}

func (m *_BACnetConstructedDataLifeSafetyZoneAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLifeSafetyZoneAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLifeSafetyZoneAll(structType any) BACnetConstructedDataLifeSafetyZoneAll {
	if casted, ok := structType.(BACnetConstructedDataLifeSafetyZoneAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLifeSafetyZoneAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLifeSafetyZoneAll) GetTypeName() string {
	return "BACnetConstructedDataLifeSafetyZoneAll"
}

func (m *_BACnetConstructedDataLifeSafetyZoneAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataLifeSafetyZoneAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLifeSafetyZoneAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLifeSafetyZoneAll BACnetConstructedDataLifeSafetyZoneAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLifeSafetyZoneAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLifeSafetyZoneAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLifeSafetyZoneAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLifeSafetyZoneAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLifeSafetyZoneAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLifeSafetyZoneAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLifeSafetyZoneAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLifeSafetyZoneAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLifeSafetyZoneAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLifeSafetyZoneAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLifeSafetyZoneAll) IsBACnetConstructedDataLifeSafetyZoneAll() {}

func (m *_BACnetConstructedDataLifeSafetyZoneAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLifeSafetyZoneAll) deepCopy() *_BACnetConstructedDataLifeSafetyZoneAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLifeSafetyZoneAllCopy := &_BACnetConstructedDataLifeSafetyZoneAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	_BACnetConstructedDataLifeSafetyZoneAllCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLifeSafetyZoneAllCopy
}

func (m *_BACnetConstructedDataLifeSafetyZoneAll) String() string {
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
