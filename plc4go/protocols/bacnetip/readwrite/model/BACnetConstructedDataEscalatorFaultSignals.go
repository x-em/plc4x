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

// BACnetConstructedDataEscalatorFaultSignals is the corresponding interface of BACnetConstructedDataEscalatorFaultSignals
type BACnetConstructedDataEscalatorFaultSignals interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetFaultSignals returns FaultSignals (property field)
	GetFaultSignals() []BACnetEscalatorFaultTagged
	// IsBACnetConstructedDataEscalatorFaultSignals is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataEscalatorFaultSignals()
	// CreateBuilder creates a BACnetConstructedDataEscalatorFaultSignalsBuilder
	CreateBACnetConstructedDataEscalatorFaultSignalsBuilder() BACnetConstructedDataEscalatorFaultSignalsBuilder
}

// _BACnetConstructedDataEscalatorFaultSignals is the data-structure of this message
type _BACnetConstructedDataEscalatorFaultSignals struct {
	BACnetConstructedDataContract
	FaultSignals []BACnetEscalatorFaultTagged
}

var _ BACnetConstructedDataEscalatorFaultSignals = (*_BACnetConstructedDataEscalatorFaultSignals)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataEscalatorFaultSignals)(nil)

// NewBACnetConstructedDataEscalatorFaultSignals factory function for _BACnetConstructedDataEscalatorFaultSignals
func NewBACnetConstructedDataEscalatorFaultSignals(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, faultSignals []BACnetEscalatorFaultTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEscalatorFaultSignals {
	_result := &_BACnetConstructedDataEscalatorFaultSignals{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FaultSignals:                  faultSignals,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataEscalatorFaultSignalsBuilder is a builder for BACnetConstructedDataEscalatorFaultSignals
type BACnetConstructedDataEscalatorFaultSignalsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(faultSignals []BACnetEscalatorFaultTagged) BACnetConstructedDataEscalatorFaultSignalsBuilder
	// WithFaultSignals adds FaultSignals (property field)
	WithFaultSignals(...BACnetEscalatorFaultTagged) BACnetConstructedDataEscalatorFaultSignalsBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataEscalatorFaultSignals or returns an error if something is wrong
	Build() (BACnetConstructedDataEscalatorFaultSignals, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataEscalatorFaultSignals
}

// NewBACnetConstructedDataEscalatorFaultSignalsBuilder() creates a BACnetConstructedDataEscalatorFaultSignalsBuilder
func NewBACnetConstructedDataEscalatorFaultSignalsBuilder() BACnetConstructedDataEscalatorFaultSignalsBuilder {
	return &_BACnetConstructedDataEscalatorFaultSignalsBuilder{_BACnetConstructedDataEscalatorFaultSignals: new(_BACnetConstructedDataEscalatorFaultSignals)}
}

type _BACnetConstructedDataEscalatorFaultSignalsBuilder struct {
	*_BACnetConstructedDataEscalatorFaultSignals

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataEscalatorFaultSignalsBuilder) = (*_BACnetConstructedDataEscalatorFaultSignalsBuilder)(nil)

func (b *_BACnetConstructedDataEscalatorFaultSignalsBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataEscalatorFaultSignals
}

func (b *_BACnetConstructedDataEscalatorFaultSignalsBuilder) WithMandatoryFields(faultSignals []BACnetEscalatorFaultTagged) BACnetConstructedDataEscalatorFaultSignalsBuilder {
	return b.WithFaultSignals(faultSignals...)
}

func (b *_BACnetConstructedDataEscalatorFaultSignalsBuilder) WithFaultSignals(faultSignals ...BACnetEscalatorFaultTagged) BACnetConstructedDataEscalatorFaultSignalsBuilder {
	b.FaultSignals = faultSignals
	return b
}

func (b *_BACnetConstructedDataEscalatorFaultSignalsBuilder) Build() (BACnetConstructedDataEscalatorFaultSignals, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataEscalatorFaultSignals.deepCopy(), nil
}

func (b *_BACnetConstructedDataEscalatorFaultSignalsBuilder) MustBuild() BACnetConstructedDataEscalatorFaultSignals {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataEscalatorFaultSignalsBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataEscalatorFaultSignalsBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataEscalatorFaultSignalsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataEscalatorFaultSignalsBuilder().(*_BACnetConstructedDataEscalatorFaultSignalsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataEscalatorFaultSignalsBuilder creates a BACnetConstructedDataEscalatorFaultSignalsBuilder
func (b *_BACnetConstructedDataEscalatorFaultSignals) CreateBACnetConstructedDataEscalatorFaultSignalsBuilder() BACnetConstructedDataEscalatorFaultSignalsBuilder {
	if b == nil {
		return NewBACnetConstructedDataEscalatorFaultSignalsBuilder()
	}
	return &_BACnetConstructedDataEscalatorFaultSignalsBuilder{_BACnetConstructedDataEscalatorFaultSignals: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEscalatorFaultSignals) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ESCALATOR
}

func (m *_BACnetConstructedDataEscalatorFaultSignals) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_SIGNALS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEscalatorFaultSignals) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEscalatorFaultSignals) GetFaultSignals() []BACnetEscalatorFaultTagged {
	return m.FaultSignals
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEscalatorFaultSignals(structType any) BACnetConstructedDataEscalatorFaultSignals {
	if casted, ok := structType.(BACnetConstructedDataEscalatorFaultSignals); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEscalatorFaultSignals); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEscalatorFaultSignals) GetTypeName() string {
	return "BACnetConstructedDataEscalatorFaultSignals"
}

func (m *_BACnetConstructedDataEscalatorFaultSignals) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.FaultSignals) > 0 {
		for _, element := range m.FaultSignals {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataEscalatorFaultSignals) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEscalatorFaultSignals) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEscalatorFaultSignals BACnetConstructedDataEscalatorFaultSignals, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEscalatorFaultSignals"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEscalatorFaultSignals")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	faultSignals, err := ReadTerminatedArrayField[BACnetEscalatorFaultTagged](ctx, "faultSignals", ReadComplex[BACnetEscalatorFaultTagged](BACnetEscalatorFaultTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'faultSignals' field"))
	}
	m.FaultSignals = faultSignals

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEscalatorFaultSignals"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEscalatorFaultSignals")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEscalatorFaultSignals) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEscalatorFaultSignals) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEscalatorFaultSignals"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEscalatorFaultSignals")
		}

		if err := WriteComplexTypeArrayField(ctx, "faultSignals", m.GetFaultSignals(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'faultSignals' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEscalatorFaultSignals"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEscalatorFaultSignals")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEscalatorFaultSignals) IsBACnetConstructedDataEscalatorFaultSignals() {
}

func (m *_BACnetConstructedDataEscalatorFaultSignals) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataEscalatorFaultSignals) deepCopy() *_BACnetConstructedDataEscalatorFaultSignals {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataEscalatorFaultSignalsCopy := &_BACnetConstructedDataEscalatorFaultSignals{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetEscalatorFaultTagged, BACnetEscalatorFaultTagged](m.FaultSignals),
	}
	_BACnetConstructedDataEscalatorFaultSignalsCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataEscalatorFaultSignalsCopy
}

func (m *_BACnetConstructedDataEscalatorFaultSignals) String() string {
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