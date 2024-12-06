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

// BACnetConstructedDataMinimumOutput is the corresponding interface of BACnetConstructedDataMinimumOutput
type BACnetConstructedDataMinimumOutput interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMinimumOutput returns MinimumOutput (property field)
	GetMinimumOutput() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataMinimumOutput is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMinimumOutput()
	// CreateBuilder creates a BACnetConstructedDataMinimumOutputBuilder
	CreateBACnetConstructedDataMinimumOutputBuilder() BACnetConstructedDataMinimumOutputBuilder
}

// _BACnetConstructedDataMinimumOutput is the data-structure of this message
type _BACnetConstructedDataMinimumOutput struct {
	BACnetConstructedDataContract
	MinimumOutput BACnetApplicationTagReal
}

var _ BACnetConstructedDataMinimumOutput = (*_BACnetConstructedDataMinimumOutput)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMinimumOutput)(nil)

// NewBACnetConstructedDataMinimumOutput factory function for _BACnetConstructedDataMinimumOutput
func NewBACnetConstructedDataMinimumOutput(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, minimumOutput BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMinimumOutput {
	if minimumOutput == nil {
		panic("minimumOutput of type BACnetApplicationTagReal for BACnetConstructedDataMinimumOutput must not be nil")
	}
	_result := &_BACnetConstructedDataMinimumOutput{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MinimumOutput:                 minimumOutput,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataMinimumOutputBuilder is a builder for BACnetConstructedDataMinimumOutput
type BACnetConstructedDataMinimumOutputBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(minimumOutput BACnetApplicationTagReal) BACnetConstructedDataMinimumOutputBuilder
	// WithMinimumOutput adds MinimumOutput (property field)
	WithMinimumOutput(BACnetApplicationTagReal) BACnetConstructedDataMinimumOutputBuilder
	// WithMinimumOutputBuilder adds MinimumOutput (property field) which is build by the builder
	WithMinimumOutputBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataMinimumOutputBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataMinimumOutput or returns an error if something is wrong
	Build() (BACnetConstructedDataMinimumOutput, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataMinimumOutput
}

// NewBACnetConstructedDataMinimumOutputBuilder() creates a BACnetConstructedDataMinimumOutputBuilder
func NewBACnetConstructedDataMinimumOutputBuilder() BACnetConstructedDataMinimumOutputBuilder {
	return &_BACnetConstructedDataMinimumOutputBuilder{_BACnetConstructedDataMinimumOutput: new(_BACnetConstructedDataMinimumOutput)}
}

type _BACnetConstructedDataMinimumOutputBuilder struct {
	*_BACnetConstructedDataMinimumOutput

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataMinimumOutputBuilder) = (*_BACnetConstructedDataMinimumOutputBuilder)(nil)

func (b *_BACnetConstructedDataMinimumOutputBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataMinimumOutput
}

func (b *_BACnetConstructedDataMinimumOutputBuilder) WithMandatoryFields(minimumOutput BACnetApplicationTagReal) BACnetConstructedDataMinimumOutputBuilder {
	return b.WithMinimumOutput(minimumOutput)
}

func (b *_BACnetConstructedDataMinimumOutputBuilder) WithMinimumOutput(minimumOutput BACnetApplicationTagReal) BACnetConstructedDataMinimumOutputBuilder {
	b.MinimumOutput = minimumOutput
	return b
}

func (b *_BACnetConstructedDataMinimumOutputBuilder) WithMinimumOutputBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataMinimumOutputBuilder {
	builder := builderSupplier(b.MinimumOutput.CreateBACnetApplicationTagRealBuilder())
	var err error
	b.MinimumOutput, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataMinimumOutputBuilder) Build() (BACnetConstructedDataMinimumOutput, error) {
	if b.MinimumOutput == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'minimumOutput' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataMinimumOutput.deepCopy(), nil
}

func (b *_BACnetConstructedDataMinimumOutputBuilder) MustBuild() BACnetConstructedDataMinimumOutput {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataMinimumOutputBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataMinimumOutputBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataMinimumOutputBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataMinimumOutputBuilder().(*_BACnetConstructedDataMinimumOutputBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataMinimumOutputBuilder creates a BACnetConstructedDataMinimumOutputBuilder
func (b *_BACnetConstructedDataMinimumOutput) CreateBACnetConstructedDataMinimumOutputBuilder() BACnetConstructedDataMinimumOutputBuilder {
	if b == nil {
		return NewBACnetConstructedDataMinimumOutputBuilder()
	}
	return &_BACnetConstructedDataMinimumOutputBuilder{_BACnetConstructedDataMinimumOutput: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMinimumOutput) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMinimumOutput) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MINIMUM_OUTPUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMinimumOutput) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMinimumOutput) GetMinimumOutput() BACnetApplicationTagReal {
	return m.MinimumOutput
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMinimumOutput) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetMinimumOutput())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMinimumOutput(structType any) BACnetConstructedDataMinimumOutput {
	if casted, ok := structType.(BACnetConstructedDataMinimumOutput); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMinimumOutput); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMinimumOutput) GetTypeName() string {
	return "BACnetConstructedDataMinimumOutput"
}

func (m *_BACnetConstructedDataMinimumOutput) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (minimumOutput)
	lengthInBits += m.MinimumOutput.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMinimumOutput) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMinimumOutput) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMinimumOutput BACnetConstructedDataMinimumOutput, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMinimumOutput"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMinimumOutput")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	minimumOutput, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "minimumOutput", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minimumOutput' field"))
	}
	m.MinimumOutput = minimumOutput

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), minimumOutput)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMinimumOutput"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMinimumOutput")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMinimumOutput) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMinimumOutput) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMinimumOutput"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMinimumOutput")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "minimumOutput", m.GetMinimumOutput(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'minimumOutput' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMinimumOutput"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMinimumOutput")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMinimumOutput) IsBACnetConstructedDataMinimumOutput() {}

func (m *_BACnetConstructedDataMinimumOutput) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataMinimumOutput) deepCopy() *_BACnetConstructedDataMinimumOutput {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataMinimumOutputCopy := &_BACnetConstructedDataMinimumOutput{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagReal](m.MinimumOutput),
	}
	_BACnetConstructedDataMinimumOutputCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataMinimumOutputCopy
}

func (m *_BACnetConstructedDataMinimumOutput) String() string {
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