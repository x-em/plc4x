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

// BACnetConstructedDataManualSlaveAddressBinding is the corresponding interface of BACnetConstructedDataManualSlaveAddressBinding
type BACnetConstructedDataManualSlaveAddressBinding interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetManualSlaveAddressBinding returns ManualSlaveAddressBinding (property field)
	GetManualSlaveAddressBinding() []BACnetAddressBinding
	// IsBACnetConstructedDataManualSlaveAddressBinding is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataManualSlaveAddressBinding()
	// CreateBuilder creates a BACnetConstructedDataManualSlaveAddressBindingBuilder
	CreateBACnetConstructedDataManualSlaveAddressBindingBuilder() BACnetConstructedDataManualSlaveAddressBindingBuilder
}

// _BACnetConstructedDataManualSlaveAddressBinding is the data-structure of this message
type _BACnetConstructedDataManualSlaveAddressBinding struct {
	BACnetConstructedDataContract
	ManualSlaveAddressBinding []BACnetAddressBinding
}

var _ BACnetConstructedDataManualSlaveAddressBinding = (*_BACnetConstructedDataManualSlaveAddressBinding)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataManualSlaveAddressBinding)(nil)

// NewBACnetConstructedDataManualSlaveAddressBinding factory function for _BACnetConstructedDataManualSlaveAddressBinding
func NewBACnetConstructedDataManualSlaveAddressBinding(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, manualSlaveAddressBinding []BACnetAddressBinding, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataManualSlaveAddressBinding {
	_result := &_BACnetConstructedDataManualSlaveAddressBinding{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ManualSlaveAddressBinding:     manualSlaveAddressBinding,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataManualSlaveAddressBindingBuilder is a builder for BACnetConstructedDataManualSlaveAddressBinding
type BACnetConstructedDataManualSlaveAddressBindingBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(manualSlaveAddressBinding []BACnetAddressBinding) BACnetConstructedDataManualSlaveAddressBindingBuilder
	// WithManualSlaveAddressBinding adds ManualSlaveAddressBinding (property field)
	WithManualSlaveAddressBinding(...BACnetAddressBinding) BACnetConstructedDataManualSlaveAddressBindingBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataManualSlaveAddressBinding or returns an error if something is wrong
	Build() (BACnetConstructedDataManualSlaveAddressBinding, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataManualSlaveAddressBinding
}

// NewBACnetConstructedDataManualSlaveAddressBindingBuilder() creates a BACnetConstructedDataManualSlaveAddressBindingBuilder
func NewBACnetConstructedDataManualSlaveAddressBindingBuilder() BACnetConstructedDataManualSlaveAddressBindingBuilder {
	return &_BACnetConstructedDataManualSlaveAddressBindingBuilder{_BACnetConstructedDataManualSlaveAddressBinding: new(_BACnetConstructedDataManualSlaveAddressBinding)}
}

type _BACnetConstructedDataManualSlaveAddressBindingBuilder struct {
	*_BACnetConstructedDataManualSlaveAddressBinding

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataManualSlaveAddressBindingBuilder) = (*_BACnetConstructedDataManualSlaveAddressBindingBuilder)(nil)

func (b *_BACnetConstructedDataManualSlaveAddressBindingBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataManualSlaveAddressBinding
}

func (b *_BACnetConstructedDataManualSlaveAddressBindingBuilder) WithMandatoryFields(manualSlaveAddressBinding []BACnetAddressBinding) BACnetConstructedDataManualSlaveAddressBindingBuilder {
	return b.WithManualSlaveAddressBinding(manualSlaveAddressBinding...)
}

func (b *_BACnetConstructedDataManualSlaveAddressBindingBuilder) WithManualSlaveAddressBinding(manualSlaveAddressBinding ...BACnetAddressBinding) BACnetConstructedDataManualSlaveAddressBindingBuilder {
	b.ManualSlaveAddressBinding = manualSlaveAddressBinding
	return b
}

func (b *_BACnetConstructedDataManualSlaveAddressBindingBuilder) Build() (BACnetConstructedDataManualSlaveAddressBinding, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataManualSlaveAddressBinding.deepCopy(), nil
}

func (b *_BACnetConstructedDataManualSlaveAddressBindingBuilder) MustBuild() BACnetConstructedDataManualSlaveAddressBinding {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataManualSlaveAddressBindingBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataManualSlaveAddressBindingBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataManualSlaveAddressBindingBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataManualSlaveAddressBindingBuilder().(*_BACnetConstructedDataManualSlaveAddressBindingBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataManualSlaveAddressBindingBuilder creates a BACnetConstructedDataManualSlaveAddressBindingBuilder
func (b *_BACnetConstructedDataManualSlaveAddressBinding) CreateBACnetConstructedDataManualSlaveAddressBindingBuilder() BACnetConstructedDataManualSlaveAddressBindingBuilder {
	if b == nil {
		return NewBACnetConstructedDataManualSlaveAddressBindingBuilder()
	}
	return &_BACnetConstructedDataManualSlaveAddressBindingBuilder{_BACnetConstructedDataManualSlaveAddressBinding: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MANUAL_SLAVE_ADDRESS_BINDING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetManualSlaveAddressBinding() []BACnetAddressBinding {
	return m.ManualSlaveAddressBinding
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataManualSlaveAddressBinding(structType any) BACnetConstructedDataManualSlaveAddressBinding {
	if casted, ok := structType.(BACnetConstructedDataManualSlaveAddressBinding); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataManualSlaveAddressBinding); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetTypeName() string {
	return "BACnetConstructedDataManualSlaveAddressBinding"
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.ManualSlaveAddressBinding) > 0 {
		for _, element := range m.ManualSlaveAddressBinding {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataManualSlaveAddressBinding BACnetConstructedDataManualSlaveAddressBinding, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataManualSlaveAddressBinding"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataManualSlaveAddressBinding")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	manualSlaveAddressBinding, err := ReadTerminatedArrayField[BACnetAddressBinding](ctx, "manualSlaveAddressBinding", ReadComplex[BACnetAddressBinding](BACnetAddressBindingParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'manualSlaveAddressBinding' field"))
	}
	m.ManualSlaveAddressBinding = manualSlaveAddressBinding

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataManualSlaveAddressBinding"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataManualSlaveAddressBinding")
	}

	return m, nil
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataManualSlaveAddressBinding"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataManualSlaveAddressBinding")
		}

		if err := WriteComplexTypeArrayField(ctx, "manualSlaveAddressBinding", m.GetManualSlaveAddressBinding(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'manualSlaveAddressBinding' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataManualSlaveAddressBinding"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataManualSlaveAddressBinding")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) IsBACnetConstructedDataManualSlaveAddressBinding() {
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) deepCopy() *_BACnetConstructedDataManualSlaveAddressBinding {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataManualSlaveAddressBindingCopy := &_BACnetConstructedDataManualSlaveAddressBinding{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetAddressBinding, BACnetAddressBinding](m.ManualSlaveAddressBinding),
	}
	_BACnetConstructedDataManualSlaveAddressBindingCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataManualSlaveAddressBindingCopy
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) String() string {
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
