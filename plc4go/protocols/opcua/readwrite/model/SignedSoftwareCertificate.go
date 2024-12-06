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

// SignedSoftwareCertificate is the corresponding interface of SignedSoftwareCertificate
type SignedSoftwareCertificate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetCertificateData returns CertificateData (property field)
	GetCertificateData() PascalByteString
	// GetSignature returns Signature (property field)
	GetSignature() PascalByteString
	// IsSignedSoftwareCertificate is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSignedSoftwareCertificate()
	// CreateBuilder creates a SignedSoftwareCertificateBuilder
	CreateSignedSoftwareCertificateBuilder() SignedSoftwareCertificateBuilder
}

// _SignedSoftwareCertificate is the data-structure of this message
type _SignedSoftwareCertificate struct {
	ExtensionObjectDefinitionContract
	CertificateData PascalByteString
	Signature       PascalByteString
}

var _ SignedSoftwareCertificate = (*_SignedSoftwareCertificate)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_SignedSoftwareCertificate)(nil)

// NewSignedSoftwareCertificate factory function for _SignedSoftwareCertificate
func NewSignedSoftwareCertificate(certificateData PascalByteString, signature PascalByteString) *_SignedSoftwareCertificate {
	if certificateData == nil {
		panic("certificateData of type PascalByteString for SignedSoftwareCertificate must not be nil")
	}
	if signature == nil {
		panic("signature of type PascalByteString for SignedSoftwareCertificate must not be nil")
	}
	_result := &_SignedSoftwareCertificate{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		CertificateData:                   certificateData,
		Signature:                         signature,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SignedSoftwareCertificateBuilder is a builder for SignedSoftwareCertificate
type SignedSoftwareCertificateBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(certificateData PascalByteString, signature PascalByteString) SignedSoftwareCertificateBuilder
	// WithCertificateData adds CertificateData (property field)
	WithCertificateData(PascalByteString) SignedSoftwareCertificateBuilder
	// WithCertificateDataBuilder adds CertificateData (property field) which is build by the builder
	WithCertificateDataBuilder(func(PascalByteStringBuilder) PascalByteStringBuilder) SignedSoftwareCertificateBuilder
	// WithSignature adds Signature (property field)
	WithSignature(PascalByteString) SignedSoftwareCertificateBuilder
	// WithSignatureBuilder adds Signature (property field) which is build by the builder
	WithSignatureBuilder(func(PascalByteStringBuilder) PascalByteStringBuilder) SignedSoftwareCertificateBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the SignedSoftwareCertificate or returns an error if something is wrong
	Build() (SignedSoftwareCertificate, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SignedSoftwareCertificate
}

// NewSignedSoftwareCertificateBuilder() creates a SignedSoftwareCertificateBuilder
func NewSignedSoftwareCertificateBuilder() SignedSoftwareCertificateBuilder {
	return &_SignedSoftwareCertificateBuilder{_SignedSoftwareCertificate: new(_SignedSoftwareCertificate)}
}

type _SignedSoftwareCertificateBuilder struct {
	*_SignedSoftwareCertificate

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (SignedSoftwareCertificateBuilder) = (*_SignedSoftwareCertificateBuilder)(nil)

func (b *_SignedSoftwareCertificateBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._SignedSoftwareCertificate
}

func (b *_SignedSoftwareCertificateBuilder) WithMandatoryFields(certificateData PascalByteString, signature PascalByteString) SignedSoftwareCertificateBuilder {
	return b.WithCertificateData(certificateData).WithSignature(signature)
}

func (b *_SignedSoftwareCertificateBuilder) WithCertificateData(certificateData PascalByteString) SignedSoftwareCertificateBuilder {
	b.CertificateData = certificateData
	return b
}

func (b *_SignedSoftwareCertificateBuilder) WithCertificateDataBuilder(builderSupplier func(PascalByteStringBuilder) PascalByteStringBuilder) SignedSoftwareCertificateBuilder {
	builder := builderSupplier(b.CertificateData.CreatePascalByteStringBuilder())
	var err error
	b.CertificateData, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalByteStringBuilder failed"))
	}
	return b
}

func (b *_SignedSoftwareCertificateBuilder) WithSignature(signature PascalByteString) SignedSoftwareCertificateBuilder {
	b.Signature = signature
	return b
}

func (b *_SignedSoftwareCertificateBuilder) WithSignatureBuilder(builderSupplier func(PascalByteStringBuilder) PascalByteStringBuilder) SignedSoftwareCertificateBuilder {
	builder := builderSupplier(b.Signature.CreatePascalByteStringBuilder())
	var err error
	b.Signature, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalByteStringBuilder failed"))
	}
	return b
}

func (b *_SignedSoftwareCertificateBuilder) Build() (SignedSoftwareCertificate, error) {
	if b.CertificateData == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'certificateData' not set"))
	}
	if b.Signature == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'signature' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SignedSoftwareCertificate.deepCopy(), nil
}

func (b *_SignedSoftwareCertificateBuilder) MustBuild() SignedSoftwareCertificate {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SignedSoftwareCertificateBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_SignedSoftwareCertificateBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_SignedSoftwareCertificateBuilder) DeepCopy() any {
	_copy := b.CreateSignedSoftwareCertificateBuilder().(*_SignedSoftwareCertificateBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSignedSoftwareCertificateBuilder creates a SignedSoftwareCertificateBuilder
func (b *_SignedSoftwareCertificate) CreateSignedSoftwareCertificateBuilder() SignedSoftwareCertificateBuilder {
	if b == nil {
		return NewSignedSoftwareCertificateBuilder()
	}
	return &_SignedSoftwareCertificateBuilder{_SignedSoftwareCertificate: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SignedSoftwareCertificate) GetExtensionId() int32 {
	return int32(346)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SignedSoftwareCertificate) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SignedSoftwareCertificate) GetCertificateData() PascalByteString {
	return m.CertificateData
}

func (m *_SignedSoftwareCertificate) GetSignature() PascalByteString {
	return m.Signature
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSignedSoftwareCertificate(structType any) SignedSoftwareCertificate {
	if casted, ok := structType.(SignedSoftwareCertificate); ok {
		return casted
	}
	if casted, ok := structType.(*SignedSoftwareCertificate); ok {
		return *casted
	}
	return nil
}

func (m *_SignedSoftwareCertificate) GetTypeName() string {
	return "SignedSoftwareCertificate"
}

func (m *_SignedSoftwareCertificate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (certificateData)
	lengthInBits += m.CertificateData.GetLengthInBits(ctx)

	// Simple field (signature)
	lengthInBits += m.Signature.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SignedSoftwareCertificate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SignedSoftwareCertificate) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__signedSoftwareCertificate SignedSoftwareCertificate, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SignedSoftwareCertificate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SignedSoftwareCertificate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	certificateData, err := ReadSimpleField[PascalByteString](ctx, "certificateData", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'certificateData' field"))
	}
	m.CertificateData = certificateData

	signature, err := ReadSimpleField[PascalByteString](ctx, "signature", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'signature' field"))
	}
	m.Signature = signature

	if closeErr := readBuffer.CloseContext("SignedSoftwareCertificate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SignedSoftwareCertificate")
	}

	return m, nil
}

func (m *_SignedSoftwareCertificate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SignedSoftwareCertificate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SignedSoftwareCertificate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SignedSoftwareCertificate")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "certificateData", m.GetCertificateData(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'certificateData' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "signature", m.GetSignature(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'signature' field")
		}

		if popErr := writeBuffer.PopContext("SignedSoftwareCertificate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SignedSoftwareCertificate")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SignedSoftwareCertificate) IsSignedSoftwareCertificate() {}

func (m *_SignedSoftwareCertificate) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SignedSoftwareCertificate) deepCopy() *_SignedSoftwareCertificate {
	if m == nil {
		return nil
	}
	_SignedSoftwareCertificateCopy := &_SignedSoftwareCertificate{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[PascalByteString](m.CertificateData),
		utils.DeepCopy[PascalByteString](m.Signature),
	}
	_SignedSoftwareCertificateCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _SignedSoftwareCertificateCopy
}

func (m *_SignedSoftwareCertificate) String() string {
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