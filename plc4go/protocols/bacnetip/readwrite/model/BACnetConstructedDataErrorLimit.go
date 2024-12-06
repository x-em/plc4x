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

// BACnetConstructedDataErrorLimit is the corresponding interface of BACnetConstructedDataErrorLimit
type BACnetConstructedDataErrorLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetErrorLimit returns ErrorLimit (property field)
	GetErrorLimit() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataErrorLimit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataErrorLimit()
	// CreateBuilder creates a BACnetConstructedDataErrorLimitBuilder
	CreateBACnetConstructedDataErrorLimitBuilder() BACnetConstructedDataErrorLimitBuilder
}

// _BACnetConstructedDataErrorLimit is the data-structure of this message
type _BACnetConstructedDataErrorLimit struct {
	BACnetConstructedDataContract
	ErrorLimit BACnetApplicationTagReal
}

var _ BACnetConstructedDataErrorLimit = (*_BACnetConstructedDataErrorLimit)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataErrorLimit)(nil)

// NewBACnetConstructedDataErrorLimit factory function for _BACnetConstructedDataErrorLimit
func NewBACnetConstructedDataErrorLimit(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, errorLimit BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataErrorLimit {
	if errorLimit == nil {
		panic("errorLimit of type BACnetApplicationTagReal for BACnetConstructedDataErrorLimit must not be nil")
	}
	_result := &_BACnetConstructedDataErrorLimit{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ErrorLimit:                    errorLimit,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataErrorLimitBuilder is a builder for BACnetConstructedDataErrorLimit
type BACnetConstructedDataErrorLimitBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(errorLimit BACnetApplicationTagReal) BACnetConstructedDataErrorLimitBuilder
	// WithErrorLimit adds ErrorLimit (property field)
	WithErrorLimit(BACnetApplicationTagReal) BACnetConstructedDataErrorLimitBuilder
	// WithErrorLimitBuilder adds ErrorLimit (property field) which is build by the builder
	WithErrorLimitBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataErrorLimitBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataErrorLimit or returns an error if something is wrong
	Build() (BACnetConstructedDataErrorLimit, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataErrorLimit
}

// NewBACnetConstructedDataErrorLimitBuilder() creates a BACnetConstructedDataErrorLimitBuilder
func NewBACnetConstructedDataErrorLimitBuilder() BACnetConstructedDataErrorLimitBuilder {
	return &_BACnetConstructedDataErrorLimitBuilder{_BACnetConstructedDataErrorLimit: new(_BACnetConstructedDataErrorLimit)}
}

type _BACnetConstructedDataErrorLimitBuilder struct {
	*_BACnetConstructedDataErrorLimit

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataErrorLimitBuilder) = (*_BACnetConstructedDataErrorLimitBuilder)(nil)

func (b *_BACnetConstructedDataErrorLimitBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataErrorLimit
}

func (b *_BACnetConstructedDataErrorLimitBuilder) WithMandatoryFields(errorLimit BACnetApplicationTagReal) BACnetConstructedDataErrorLimitBuilder {
	return b.WithErrorLimit(errorLimit)
}

func (b *_BACnetConstructedDataErrorLimitBuilder) WithErrorLimit(errorLimit BACnetApplicationTagReal) BACnetConstructedDataErrorLimitBuilder {
	b.ErrorLimit = errorLimit
	return b
}

func (b *_BACnetConstructedDataErrorLimitBuilder) WithErrorLimitBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataErrorLimitBuilder {
	builder := builderSupplier(b.ErrorLimit.CreateBACnetApplicationTagRealBuilder())
	var err error
	b.ErrorLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataErrorLimitBuilder) Build() (BACnetConstructedDataErrorLimit, error) {
	if b.ErrorLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'errorLimit' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataErrorLimit.deepCopy(), nil
}

func (b *_BACnetConstructedDataErrorLimitBuilder) MustBuild() BACnetConstructedDataErrorLimit {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataErrorLimitBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataErrorLimitBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataErrorLimitBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataErrorLimitBuilder().(*_BACnetConstructedDataErrorLimitBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataErrorLimitBuilder creates a BACnetConstructedDataErrorLimitBuilder
func (b *_BACnetConstructedDataErrorLimit) CreateBACnetConstructedDataErrorLimitBuilder() BACnetConstructedDataErrorLimitBuilder {
	if b == nil {
		return NewBACnetConstructedDataErrorLimitBuilder()
	}
	return &_BACnetConstructedDataErrorLimitBuilder{_BACnetConstructedDataErrorLimit: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataErrorLimit) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataErrorLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ERROR_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataErrorLimit) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataErrorLimit) GetErrorLimit() BACnetApplicationTagReal {
	return m.ErrorLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataErrorLimit) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetErrorLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataErrorLimit(structType any) BACnetConstructedDataErrorLimit {
	if casted, ok := structType.(BACnetConstructedDataErrorLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataErrorLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataErrorLimit) GetTypeName() string {
	return "BACnetConstructedDataErrorLimit"
}

func (m *_BACnetConstructedDataErrorLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (errorLimit)
	lengthInBits += m.ErrorLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataErrorLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataErrorLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataErrorLimit BACnetConstructedDataErrorLimit, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataErrorLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataErrorLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	errorLimit, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "errorLimit", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorLimit' field"))
	}
	m.ErrorLimit = errorLimit

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), errorLimit)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataErrorLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataErrorLimit")
	}

	return m, nil
}

func (m *_BACnetConstructedDataErrorLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataErrorLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataErrorLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataErrorLimit")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "errorLimit", m.GetErrorLimit(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'errorLimit' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataErrorLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataErrorLimit")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataErrorLimit) IsBACnetConstructedDataErrorLimit() {}

func (m *_BACnetConstructedDataErrorLimit) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataErrorLimit) deepCopy() *_BACnetConstructedDataErrorLimit {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataErrorLimitCopy := &_BACnetConstructedDataErrorLimit{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagReal](m.ErrorLimit),
	}
	_BACnetConstructedDataErrorLimitCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataErrorLimitCopy
}

func (m *_BACnetConstructedDataErrorLimit) String() string {
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