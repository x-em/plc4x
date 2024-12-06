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

// BACnetConstructedDataTransactionNotificationClass is the corresponding interface of BACnetConstructedDataTransactionNotificationClass
type BACnetConstructedDataTransactionNotificationClass interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetTransactionNotificationClass returns TransactionNotificationClass (property field)
	GetTransactionNotificationClass() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataTransactionNotificationClass is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTransactionNotificationClass()
	// CreateBuilder creates a BACnetConstructedDataTransactionNotificationClassBuilder
	CreateBACnetConstructedDataTransactionNotificationClassBuilder() BACnetConstructedDataTransactionNotificationClassBuilder
}

// _BACnetConstructedDataTransactionNotificationClass is the data-structure of this message
type _BACnetConstructedDataTransactionNotificationClass struct {
	BACnetConstructedDataContract
	TransactionNotificationClass BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataTransactionNotificationClass = (*_BACnetConstructedDataTransactionNotificationClass)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTransactionNotificationClass)(nil)

// NewBACnetConstructedDataTransactionNotificationClass factory function for _BACnetConstructedDataTransactionNotificationClass
func NewBACnetConstructedDataTransactionNotificationClass(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, transactionNotificationClass BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTransactionNotificationClass {
	if transactionNotificationClass == nil {
		panic("transactionNotificationClass of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataTransactionNotificationClass must not be nil")
	}
	_result := &_BACnetConstructedDataTransactionNotificationClass{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		TransactionNotificationClass:  transactionNotificationClass,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataTransactionNotificationClassBuilder is a builder for BACnetConstructedDataTransactionNotificationClass
type BACnetConstructedDataTransactionNotificationClassBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(transactionNotificationClass BACnetApplicationTagUnsignedInteger) BACnetConstructedDataTransactionNotificationClassBuilder
	// WithTransactionNotificationClass adds TransactionNotificationClass (property field)
	WithTransactionNotificationClass(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataTransactionNotificationClassBuilder
	// WithTransactionNotificationClassBuilder adds TransactionNotificationClass (property field) which is build by the builder
	WithTransactionNotificationClassBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataTransactionNotificationClassBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataTransactionNotificationClass or returns an error if something is wrong
	Build() (BACnetConstructedDataTransactionNotificationClass, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataTransactionNotificationClass
}

// NewBACnetConstructedDataTransactionNotificationClassBuilder() creates a BACnetConstructedDataTransactionNotificationClassBuilder
func NewBACnetConstructedDataTransactionNotificationClassBuilder() BACnetConstructedDataTransactionNotificationClassBuilder {
	return &_BACnetConstructedDataTransactionNotificationClassBuilder{_BACnetConstructedDataTransactionNotificationClass: new(_BACnetConstructedDataTransactionNotificationClass)}
}

type _BACnetConstructedDataTransactionNotificationClassBuilder struct {
	*_BACnetConstructedDataTransactionNotificationClass

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataTransactionNotificationClassBuilder) = (*_BACnetConstructedDataTransactionNotificationClassBuilder)(nil)

func (b *_BACnetConstructedDataTransactionNotificationClassBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataTransactionNotificationClass
}

func (b *_BACnetConstructedDataTransactionNotificationClassBuilder) WithMandatoryFields(transactionNotificationClass BACnetApplicationTagUnsignedInteger) BACnetConstructedDataTransactionNotificationClassBuilder {
	return b.WithTransactionNotificationClass(transactionNotificationClass)
}

func (b *_BACnetConstructedDataTransactionNotificationClassBuilder) WithTransactionNotificationClass(transactionNotificationClass BACnetApplicationTagUnsignedInteger) BACnetConstructedDataTransactionNotificationClassBuilder {
	b.TransactionNotificationClass = transactionNotificationClass
	return b
}

func (b *_BACnetConstructedDataTransactionNotificationClassBuilder) WithTransactionNotificationClassBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataTransactionNotificationClassBuilder {
	builder := builderSupplier(b.TransactionNotificationClass.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.TransactionNotificationClass, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataTransactionNotificationClassBuilder) Build() (BACnetConstructedDataTransactionNotificationClass, error) {
	if b.TransactionNotificationClass == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'transactionNotificationClass' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataTransactionNotificationClass.deepCopy(), nil
}

func (b *_BACnetConstructedDataTransactionNotificationClassBuilder) MustBuild() BACnetConstructedDataTransactionNotificationClass {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataTransactionNotificationClassBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataTransactionNotificationClassBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataTransactionNotificationClassBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataTransactionNotificationClassBuilder().(*_BACnetConstructedDataTransactionNotificationClassBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataTransactionNotificationClassBuilder creates a BACnetConstructedDataTransactionNotificationClassBuilder
func (b *_BACnetConstructedDataTransactionNotificationClass) CreateBACnetConstructedDataTransactionNotificationClassBuilder() BACnetConstructedDataTransactionNotificationClassBuilder {
	if b == nil {
		return NewBACnetConstructedDataTransactionNotificationClassBuilder()
	}
	return &_BACnetConstructedDataTransactionNotificationClassBuilder{_BACnetConstructedDataTransactionNotificationClass: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTransactionNotificationClass) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTransactionNotificationClass) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TRANSACTION_NOTIFICATION_CLASS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTransactionNotificationClass) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTransactionNotificationClass) GetTransactionNotificationClass() BACnetApplicationTagUnsignedInteger {
	return m.TransactionNotificationClass
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTransactionNotificationClass) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetTransactionNotificationClass())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTransactionNotificationClass(structType any) BACnetConstructedDataTransactionNotificationClass {
	if casted, ok := structType.(BACnetConstructedDataTransactionNotificationClass); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTransactionNotificationClass); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTransactionNotificationClass) GetTypeName() string {
	return "BACnetConstructedDataTransactionNotificationClass"
}

func (m *_BACnetConstructedDataTransactionNotificationClass) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (transactionNotificationClass)
	lengthInBits += m.TransactionNotificationClass.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTransactionNotificationClass) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTransactionNotificationClass) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTransactionNotificationClass BACnetConstructedDataTransactionNotificationClass, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTransactionNotificationClass"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTransactionNotificationClass")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	transactionNotificationClass, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "transactionNotificationClass", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transactionNotificationClass' field"))
	}
	m.TransactionNotificationClass = transactionNotificationClass

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), transactionNotificationClass)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTransactionNotificationClass"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTransactionNotificationClass")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTransactionNotificationClass) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTransactionNotificationClass) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTransactionNotificationClass"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTransactionNotificationClass")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "transactionNotificationClass", m.GetTransactionNotificationClass(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'transactionNotificationClass' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTransactionNotificationClass"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTransactionNotificationClass")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTransactionNotificationClass) IsBACnetConstructedDataTransactionNotificationClass() {
}

func (m *_BACnetConstructedDataTransactionNotificationClass) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataTransactionNotificationClass) deepCopy() *_BACnetConstructedDataTransactionNotificationClass {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataTransactionNotificationClassCopy := &_BACnetConstructedDataTransactionNotificationClass{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.TransactionNotificationClass),
	}
	_BACnetConstructedDataTransactionNotificationClassCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataTransactionNotificationClassCopy
}

func (m *_BACnetConstructedDataTransactionNotificationClass) String() string {
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