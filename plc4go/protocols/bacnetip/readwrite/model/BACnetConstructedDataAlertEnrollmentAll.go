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

// BACnetConstructedDataAlertEnrollmentAll is the corresponding interface of BACnetConstructedDataAlertEnrollmentAll
type BACnetConstructedDataAlertEnrollmentAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataAlertEnrollmentAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAlertEnrollmentAll()
	// CreateBuilder creates a BACnetConstructedDataAlertEnrollmentAllBuilder
	CreateBACnetConstructedDataAlertEnrollmentAllBuilder() BACnetConstructedDataAlertEnrollmentAllBuilder
}

// _BACnetConstructedDataAlertEnrollmentAll is the data-structure of this message
type _BACnetConstructedDataAlertEnrollmentAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataAlertEnrollmentAll = (*_BACnetConstructedDataAlertEnrollmentAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAlertEnrollmentAll)(nil)

// NewBACnetConstructedDataAlertEnrollmentAll factory function for _BACnetConstructedDataAlertEnrollmentAll
func NewBACnetConstructedDataAlertEnrollmentAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAlertEnrollmentAll {
	_result := &_BACnetConstructedDataAlertEnrollmentAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAlertEnrollmentAllBuilder is a builder for BACnetConstructedDataAlertEnrollmentAll
type BACnetConstructedDataAlertEnrollmentAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataAlertEnrollmentAllBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataAlertEnrollmentAll or returns an error if something is wrong
	Build() (BACnetConstructedDataAlertEnrollmentAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAlertEnrollmentAll
}

// NewBACnetConstructedDataAlertEnrollmentAllBuilder() creates a BACnetConstructedDataAlertEnrollmentAllBuilder
func NewBACnetConstructedDataAlertEnrollmentAllBuilder() BACnetConstructedDataAlertEnrollmentAllBuilder {
	return &_BACnetConstructedDataAlertEnrollmentAllBuilder{_BACnetConstructedDataAlertEnrollmentAll: new(_BACnetConstructedDataAlertEnrollmentAll)}
}

type _BACnetConstructedDataAlertEnrollmentAllBuilder struct {
	*_BACnetConstructedDataAlertEnrollmentAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAlertEnrollmentAllBuilder) = (*_BACnetConstructedDataAlertEnrollmentAllBuilder)(nil)

func (b *_BACnetConstructedDataAlertEnrollmentAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataAlertEnrollmentAll
}

func (b *_BACnetConstructedDataAlertEnrollmentAllBuilder) WithMandatoryFields() BACnetConstructedDataAlertEnrollmentAllBuilder {
	return b
}

func (b *_BACnetConstructedDataAlertEnrollmentAllBuilder) Build() (BACnetConstructedDataAlertEnrollmentAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAlertEnrollmentAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataAlertEnrollmentAllBuilder) MustBuild() BACnetConstructedDataAlertEnrollmentAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataAlertEnrollmentAllBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAlertEnrollmentAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAlertEnrollmentAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAlertEnrollmentAllBuilder().(*_BACnetConstructedDataAlertEnrollmentAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAlertEnrollmentAllBuilder creates a BACnetConstructedDataAlertEnrollmentAllBuilder
func (b *_BACnetConstructedDataAlertEnrollmentAll) CreateBACnetConstructedDataAlertEnrollmentAllBuilder() BACnetConstructedDataAlertEnrollmentAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataAlertEnrollmentAllBuilder()
	}
	return &_BACnetConstructedDataAlertEnrollmentAllBuilder{_BACnetConstructedDataAlertEnrollmentAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAlertEnrollmentAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ALERT_ENROLLMENT
}

func (m *_BACnetConstructedDataAlertEnrollmentAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAlertEnrollmentAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAlertEnrollmentAll(structType any) BACnetConstructedDataAlertEnrollmentAll {
	if casted, ok := structType.(BACnetConstructedDataAlertEnrollmentAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAlertEnrollmentAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAlertEnrollmentAll) GetTypeName() string {
	return "BACnetConstructedDataAlertEnrollmentAll"
}

func (m *_BACnetConstructedDataAlertEnrollmentAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataAlertEnrollmentAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAlertEnrollmentAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAlertEnrollmentAll BACnetConstructedDataAlertEnrollmentAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAlertEnrollmentAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAlertEnrollmentAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAlertEnrollmentAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAlertEnrollmentAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAlertEnrollmentAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAlertEnrollmentAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAlertEnrollmentAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAlertEnrollmentAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAlertEnrollmentAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAlertEnrollmentAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAlertEnrollmentAll) IsBACnetConstructedDataAlertEnrollmentAll() {}

func (m *_BACnetConstructedDataAlertEnrollmentAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAlertEnrollmentAll) deepCopy() *_BACnetConstructedDataAlertEnrollmentAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAlertEnrollmentAllCopy := &_BACnetConstructedDataAlertEnrollmentAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	_BACnetConstructedDataAlertEnrollmentAllCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAlertEnrollmentAllCopy
}

func (m *_BACnetConstructedDataAlertEnrollmentAll) String() string {
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