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

// ApduDataExtFileStreamInfoReport is the corresponding interface of ApduDataExtFileStreamInfoReport
type ApduDataExtFileStreamInfoReport interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// IsApduDataExtFileStreamInfoReport is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtFileStreamInfoReport()
	// CreateBuilder creates a ApduDataExtFileStreamInfoReportBuilder
	CreateApduDataExtFileStreamInfoReportBuilder() ApduDataExtFileStreamInfoReportBuilder
}

// _ApduDataExtFileStreamInfoReport is the data-structure of this message
type _ApduDataExtFileStreamInfoReport struct {
	ApduDataExtContract
}

var _ ApduDataExtFileStreamInfoReport = (*_ApduDataExtFileStreamInfoReport)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtFileStreamInfoReport)(nil)

// NewApduDataExtFileStreamInfoReport factory function for _ApduDataExtFileStreamInfoReport
func NewApduDataExtFileStreamInfoReport(length uint8) *_ApduDataExtFileStreamInfoReport {
	_result := &_ApduDataExtFileStreamInfoReport{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataExtFileStreamInfoReportBuilder is a builder for ApduDataExtFileStreamInfoReport
type ApduDataExtFileStreamInfoReportBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ApduDataExtFileStreamInfoReportBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ApduDataExtBuilder
	// Build builds the ApduDataExtFileStreamInfoReport or returns an error if something is wrong
	Build() (ApduDataExtFileStreamInfoReport, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataExtFileStreamInfoReport
}

// NewApduDataExtFileStreamInfoReportBuilder() creates a ApduDataExtFileStreamInfoReportBuilder
func NewApduDataExtFileStreamInfoReportBuilder() ApduDataExtFileStreamInfoReportBuilder {
	return &_ApduDataExtFileStreamInfoReportBuilder{_ApduDataExtFileStreamInfoReport: new(_ApduDataExtFileStreamInfoReport)}
}

type _ApduDataExtFileStreamInfoReportBuilder struct {
	*_ApduDataExtFileStreamInfoReport

	parentBuilder *_ApduDataExtBuilder

	err *utils.MultiError
}

var _ (ApduDataExtFileStreamInfoReportBuilder) = (*_ApduDataExtFileStreamInfoReportBuilder)(nil)

func (b *_ApduDataExtFileStreamInfoReportBuilder) setParent(contract ApduDataExtContract) {
	b.ApduDataExtContract = contract
	contract.(*_ApduDataExt)._SubType = b._ApduDataExtFileStreamInfoReport
}

func (b *_ApduDataExtFileStreamInfoReportBuilder) WithMandatoryFields() ApduDataExtFileStreamInfoReportBuilder {
	return b
}

func (b *_ApduDataExtFileStreamInfoReportBuilder) Build() (ApduDataExtFileStreamInfoReport, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ApduDataExtFileStreamInfoReport.deepCopy(), nil
}

func (b *_ApduDataExtFileStreamInfoReportBuilder) MustBuild() ApduDataExtFileStreamInfoReport {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ApduDataExtFileStreamInfoReportBuilder) Done() ApduDataExtBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewApduDataExtBuilder().(*_ApduDataExtBuilder)
	}
	return b.parentBuilder
}

func (b *_ApduDataExtFileStreamInfoReportBuilder) buildForApduDataExt() (ApduDataExt, error) {
	return b.Build()
}

func (b *_ApduDataExtFileStreamInfoReportBuilder) DeepCopy() any {
	_copy := b.CreateApduDataExtFileStreamInfoReportBuilder().(*_ApduDataExtFileStreamInfoReportBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateApduDataExtFileStreamInfoReportBuilder creates a ApduDataExtFileStreamInfoReportBuilder
func (b *_ApduDataExtFileStreamInfoReport) CreateApduDataExtFileStreamInfoReportBuilder() ApduDataExtFileStreamInfoReportBuilder {
	if b == nil {
		return NewApduDataExtFileStreamInfoReportBuilder()
	}
	return &_ApduDataExtFileStreamInfoReportBuilder{_ApduDataExtFileStreamInfoReport: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtFileStreamInfoReport) GetExtApciType() uint8 {
	return 0x30
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtFileStreamInfoReport) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// Deprecated: use the interface for direct cast
func CastApduDataExtFileStreamInfoReport(structType any) ApduDataExtFileStreamInfoReport {
	if casted, ok := structType.(ApduDataExtFileStreamInfoReport); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtFileStreamInfoReport); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtFileStreamInfoReport) GetTypeName() string {
	return "ApduDataExtFileStreamInfoReport"
}

func (m *_ApduDataExtFileStreamInfoReport) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtFileStreamInfoReport) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtFileStreamInfoReport) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtFileStreamInfoReport ApduDataExtFileStreamInfoReport, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtFileStreamInfoReport"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtFileStreamInfoReport")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtFileStreamInfoReport"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtFileStreamInfoReport")
	}

	return m, nil
}

func (m *_ApduDataExtFileStreamInfoReport) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtFileStreamInfoReport) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtFileStreamInfoReport"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtFileStreamInfoReport")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtFileStreamInfoReport"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtFileStreamInfoReport")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtFileStreamInfoReport) IsApduDataExtFileStreamInfoReport() {}

func (m *_ApduDataExtFileStreamInfoReport) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtFileStreamInfoReport) deepCopy() *_ApduDataExtFileStreamInfoReport {
	if m == nil {
		return nil
	}
	_ApduDataExtFileStreamInfoReportCopy := &_ApduDataExtFileStreamInfoReport{
		m.ApduDataExtContract.(*_ApduDataExt).deepCopy(),
	}
	_ApduDataExtFileStreamInfoReportCopy.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtFileStreamInfoReportCopy
}

func (m *_ApduDataExtFileStreamInfoReport) String() string {
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