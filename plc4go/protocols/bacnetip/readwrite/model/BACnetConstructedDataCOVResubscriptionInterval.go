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

// BACnetConstructedDataCOVResubscriptionInterval is the corresponding interface of BACnetConstructedDataCOVResubscriptionInterval
type BACnetConstructedDataCOVResubscriptionInterval interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetCovResubscriptionInterval returns CovResubscriptionInterval (property field)
	GetCovResubscriptionInterval() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataCOVResubscriptionInterval is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataCOVResubscriptionInterval()
	// CreateBuilder creates a BACnetConstructedDataCOVResubscriptionIntervalBuilder
	CreateBACnetConstructedDataCOVResubscriptionIntervalBuilder() BACnetConstructedDataCOVResubscriptionIntervalBuilder
}

// _BACnetConstructedDataCOVResubscriptionInterval is the data-structure of this message
type _BACnetConstructedDataCOVResubscriptionInterval struct {
	BACnetConstructedDataContract
	CovResubscriptionInterval BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataCOVResubscriptionInterval = (*_BACnetConstructedDataCOVResubscriptionInterval)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCOVResubscriptionInterval)(nil)

// NewBACnetConstructedDataCOVResubscriptionInterval factory function for _BACnetConstructedDataCOVResubscriptionInterval
func NewBACnetConstructedDataCOVResubscriptionInterval(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, covResubscriptionInterval BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCOVResubscriptionInterval {
	if covResubscriptionInterval == nil {
		panic("covResubscriptionInterval of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataCOVResubscriptionInterval must not be nil")
	}
	_result := &_BACnetConstructedDataCOVResubscriptionInterval{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		CovResubscriptionInterval:     covResubscriptionInterval,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataCOVResubscriptionIntervalBuilder is a builder for BACnetConstructedDataCOVResubscriptionInterval
type BACnetConstructedDataCOVResubscriptionIntervalBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(covResubscriptionInterval BACnetApplicationTagUnsignedInteger) BACnetConstructedDataCOVResubscriptionIntervalBuilder
	// WithCovResubscriptionInterval adds CovResubscriptionInterval (property field)
	WithCovResubscriptionInterval(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataCOVResubscriptionIntervalBuilder
	// WithCovResubscriptionIntervalBuilder adds CovResubscriptionInterval (property field) which is build by the builder
	WithCovResubscriptionIntervalBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataCOVResubscriptionIntervalBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataCOVResubscriptionInterval or returns an error if something is wrong
	Build() (BACnetConstructedDataCOVResubscriptionInterval, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataCOVResubscriptionInterval
}

// NewBACnetConstructedDataCOVResubscriptionIntervalBuilder() creates a BACnetConstructedDataCOVResubscriptionIntervalBuilder
func NewBACnetConstructedDataCOVResubscriptionIntervalBuilder() BACnetConstructedDataCOVResubscriptionIntervalBuilder {
	return &_BACnetConstructedDataCOVResubscriptionIntervalBuilder{_BACnetConstructedDataCOVResubscriptionInterval: new(_BACnetConstructedDataCOVResubscriptionInterval)}
}

type _BACnetConstructedDataCOVResubscriptionIntervalBuilder struct {
	*_BACnetConstructedDataCOVResubscriptionInterval

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataCOVResubscriptionIntervalBuilder) = (*_BACnetConstructedDataCOVResubscriptionIntervalBuilder)(nil)

func (b *_BACnetConstructedDataCOVResubscriptionIntervalBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataCOVResubscriptionInterval
}

func (b *_BACnetConstructedDataCOVResubscriptionIntervalBuilder) WithMandatoryFields(covResubscriptionInterval BACnetApplicationTagUnsignedInteger) BACnetConstructedDataCOVResubscriptionIntervalBuilder {
	return b.WithCovResubscriptionInterval(covResubscriptionInterval)
}

func (b *_BACnetConstructedDataCOVResubscriptionIntervalBuilder) WithCovResubscriptionInterval(covResubscriptionInterval BACnetApplicationTagUnsignedInteger) BACnetConstructedDataCOVResubscriptionIntervalBuilder {
	b.CovResubscriptionInterval = covResubscriptionInterval
	return b
}

func (b *_BACnetConstructedDataCOVResubscriptionIntervalBuilder) WithCovResubscriptionIntervalBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataCOVResubscriptionIntervalBuilder {
	builder := builderSupplier(b.CovResubscriptionInterval.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.CovResubscriptionInterval, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataCOVResubscriptionIntervalBuilder) Build() (BACnetConstructedDataCOVResubscriptionInterval, error) {
	if b.CovResubscriptionInterval == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'covResubscriptionInterval' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataCOVResubscriptionInterval.deepCopy(), nil
}

func (b *_BACnetConstructedDataCOVResubscriptionIntervalBuilder) MustBuild() BACnetConstructedDataCOVResubscriptionInterval {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataCOVResubscriptionIntervalBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataCOVResubscriptionIntervalBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataCOVResubscriptionIntervalBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataCOVResubscriptionIntervalBuilder().(*_BACnetConstructedDataCOVResubscriptionIntervalBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataCOVResubscriptionIntervalBuilder creates a BACnetConstructedDataCOVResubscriptionIntervalBuilder
func (b *_BACnetConstructedDataCOVResubscriptionInterval) CreateBACnetConstructedDataCOVResubscriptionIntervalBuilder() BACnetConstructedDataCOVResubscriptionIntervalBuilder {
	if b == nil {
		return NewBACnetConstructedDataCOVResubscriptionIntervalBuilder()
	}
	return &_BACnetConstructedDataCOVResubscriptionIntervalBuilder{_BACnetConstructedDataCOVResubscriptionInterval: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCOVResubscriptionInterval) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCOVResubscriptionInterval) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_COV_RESUBSCRIPTION_INTERVAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCOVResubscriptionInterval) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCOVResubscriptionInterval) GetCovResubscriptionInterval() BACnetApplicationTagUnsignedInteger {
	return m.CovResubscriptionInterval
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCOVResubscriptionInterval) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetCovResubscriptionInterval())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCOVResubscriptionInterval(structType any) BACnetConstructedDataCOVResubscriptionInterval {
	if casted, ok := structType.(BACnetConstructedDataCOVResubscriptionInterval); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCOVResubscriptionInterval); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCOVResubscriptionInterval) GetTypeName() string {
	return "BACnetConstructedDataCOVResubscriptionInterval"
}

func (m *_BACnetConstructedDataCOVResubscriptionInterval) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (covResubscriptionInterval)
	lengthInBits += m.CovResubscriptionInterval.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCOVResubscriptionInterval) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCOVResubscriptionInterval) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCOVResubscriptionInterval BACnetConstructedDataCOVResubscriptionInterval, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCOVResubscriptionInterval"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCOVResubscriptionInterval")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	covResubscriptionInterval, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "covResubscriptionInterval", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'covResubscriptionInterval' field"))
	}
	m.CovResubscriptionInterval = covResubscriptionInterval

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), covResubscriptionInterval)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCOVResubscriptionInterval"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCOVResubscriptionInterval")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCOVResubscriptionInterval) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCOVResubscriptionInterval) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCOVResubscriptionInterval"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCOVResubscriptionInterval")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "covResubscriptionInterval", m.GetCovResubscriptionInterval(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'covResubscriptionInterval' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCOVResubscriptionInterval"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCOVResubscriptionInterval")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCOVResubscriptionInterval) IsBACnetConstructedDataCOVResubscriptionInterval() {
}

func (m *_BACnetConstructedDataCOVResubscriptionInterval) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataCOVResubscriptionInterval) deepCopy() *_BACnetConstructedDataCOVResubscriptionInterval {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataCOVResubscriptionIntervalCopy := &_BACnetConstructedDataCOVResubscriptionInterval{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.CovResubscriptionInterval),
	}
	_BACnetConstructedDataCOVResubscriptionIntervalCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataCOVResubscriptionIntervalCopy
}

func (m *_BACnetConstructedDataCOVResubscriptionInterval) String() string {
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