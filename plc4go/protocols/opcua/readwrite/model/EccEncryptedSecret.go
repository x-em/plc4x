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

// EccEncryptedSecret is the corresponding interface of EccEncryptedSecret
type EccEncryptedSecret interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsEccEncryptedSecret is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsEccEncryptedSecret()
	// CreateBuilder creates a EccEncryptedSecretBuilder
	CreateEccEncryptedSecretBuilder() EccEncryptedSecretBuilder
}

// _EccEncryptedSecret is the data-structure of this message
type _EccEncryptedSecret struct {
}

var _ EccEncryptedSecret = (*_EccEncryptedSecret)(nil)

// NewEccEncryptedSecret factory function for _EccEncryptedSecret
func NewEccEncryptedSecret() *_EccEncryptedSecret {
	return &_EccEncryptedSecret{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// EccEncryptedSecretBuilder is a builder for EccEncryptedSecret
type EccEncryptedSecretBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() EccEncryptedSecretBuilder
	// Build builds the EccEncryptedSecret or returns an error if something is wrong
	Build() (EccEncryptedSecret, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() EccEncryptedSecret
}

// NewEccEncryptedSecretBuilder() creates a EccEncryptedSecretBuilder
func NewEccEncryptedSecretBuilder() EccEncryptedSecretBuilder {
	return &_EccEncryptedSecretBuilder{_EccEncryptedSecret: new(_EccEncryptedSecret)}
}

type _EccEncryptedSecretBuilder struct {
	*_EccEncryptedSecret

	err *utils.MultiError
}

var _ (EccEncryptedSecretBuilder) = (*_EccEncryptedSecretBuilder)(nil)

func (b *_EccEncryptedSecretBuilder) WithMandatoryFields() EccEncryptedSecretBuilder {
	return b
}

func (b *_EccEncryptedSecretBuilder) Build() (EccEncryptedSecret, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._EccEncryptedSecret.deepCopy(), nil
}

func (b *_EccEncryptedSecretBuilder) MustBuild() EccEncryptedSecret {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_EccEncryptedSecretBuilder) DeepCopy() any {
	_copy := b.CreateEccEncryptedSecretBuilder().(*_EccEncryptedSecretBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateEccEncryptedSecretBuilder creates a EccEncryptedSecretBuilder
func (b *_EccEncryptedSecret) CreateEccEncryptedSecretBuilder() EccEncryptedSecretBuilder {
	if b == nil {
		return NewEccEncryptedSecretBuilder()
	}
	return &_EccEncryptedSecretBuilder{_EccEncryptedSecret: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastEccEncryptedSecret(structType any) EccEncryptedSecret {
	if casted, ok := structType.(EccEncryptedSecret); ok {
		return casted
	}
	if casted, ok := structType.(*EccEncryptedSecret); ok {
		return *casted
	}
	return nil
}

func (m *_EccEncryptedSecret) GetTypeName() string {
	return "EccEncryptedSecret"
}

func (m *_EccEncryptedSecret) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_EccEncryptedSecret) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func EccEncryptedSecretParse(ctx context.Context, theBytes []byte) (EccEncryptedSecret, error) {
	return EccEncryptedSecretParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func EccEncryptedSecretParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (EccEncryptedSecret, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (EccEncryptedSecret, error) {
		return EccEncryptedSecretParseWithBuffer(ctx, readBuffer)
	}
}

func EccEncryptedSecretParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (EccEncryptedSecret, error) {
	v, err := (&_EccEncryptedSecret{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_EccEncryptedSecret) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__eccEncryptedSecret EccEncryptedSecret, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EccEncryptedSecret"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EccEncryptedSecret")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("EccEncryptedSecret"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EccEncryptedSecret")
	}

	return m, nil
}

func (m *_EccEncryptedSecret) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EccEncryptedSecret) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("EccEncryptedSecret"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for EccEncryptedSecret")
	}

	if popErr := writeBuffer.PopContext("EccEncryptedSecret"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for EccEncryptedSecret")
	}
	return nil
}

func (m *_EccEncryptedSecret) IsEccEncryptedSecret() {}

func (m *_EccEncryptedSecret) DeepCopy() any {
	return m.deepCopy()
}

func (m *_EccEncryptedSecret) deepCopy() *_EccEncryptedSecret {
	if m == nil {
		return nil
	}
	_EccEncryptedSecretCopy := &_EccEncryptedSecret{}
	return _EccEncryptedSecretCopy
}

func (m *_EccEncryptedSecret) String() string {
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