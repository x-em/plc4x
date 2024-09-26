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

// CreateMonitoredItemsResponse is the corresponding interface of CreateMonitoredItemsResponse
type CreateMonitoredItemsResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetNoOfResults returns NoOfResults (property field)
	GetNoOfResults() int32
	// GetResults returns Results (property field)
	GetResults() []ExtensionObjectDefinition
	// GetNoOfDiagnosticInfos returns NoOfDiagnosticInfos (property field)
	GetNoOfDiagnosticInfos() int32
	// GetDiagnosticInfos returns DiagnosticInfos (property field)
	GetDiagnosticInfos() []DiagnosticInfo
	// IsCreateMonitoredItemsResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCreateMonitoredItemsResponse()
}

// _CreateMonitoredItemsResponse is the data-structure of this message
type _CreateMonitoredItemsResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader      ExtensionObjectDefinition
	NoOfResults         int32
	Results             []ExtensionObjectDefinition
	NoOfDiagnosticInfos int32
	DiagnosticInfos     []DiagnosticInfo
}

var _ CreateMonitoredItemsResponse = (*_CreateMonitoredItemsResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CreateMonitoredItemsResponse)(nil)

// NewCreateMonitoredItemsResponse factory function for _CreateMonitoredItemsResponse
func NewCreateMonitoredItemsResponse(responseHeader ExtensionObjectDefinition, noOfResults int32, results []ExtensionObjectDefinition, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo) *_CreateMonitoredItemsResponse {
	if responseHeader == nil {
		panic("responseHeader of type ExtensionObjectDefinition for CreateMonitoredItemsResponse must not be nil")
	}
	_result := &_CreateMonitoredItemsResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		NoOfResults:                       noOfResults,
		Results:                           results,
		NoOfDiagnosticInfos:               noOfDiagnosticInfos,
		DiagnosticInfos:                   diagnosticInfos,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CreateMonitoredItemsResponse) GetIdentifier() string {
	return "754"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CreateMonitoredItemsResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CreateMonitoredItemsResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_CreateMonitoredItemsResponse) GetNoOfResults() int32 {
	return m.NoOfResults
}

func (m *_CreateMonitoredItemsResponse) GetResults() []ExtensionObjectDefinition {
	return m.Results
}

func (m *_CreateMonitoredItemsResponse) GetNoOfDiagnosticInfos() int32 {
	return m.NoOfDiagnosticInfos
}

func (m *_CreateMonitoredItemsResponse) GetDiagnosticInfos() []DiagnosticInfo {
	return m.DiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCreateMonitoredItemsResponse(structType any) CreateMonitoredItemsResponse {
	if casted, ok := structType.(CreateMonitoredItemsResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CreateMonitoredItemsResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CreateMonitoredItemsResponse) GetTypeName() string {
	return "CreateMonitoredItemsResponse"
}

func (m *_CreateMonitoredItemsResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (noOfResults)
	lengthInBits += 32

	// Array field
	if len(m.Results) > 0 {
		for _curItem, element := range m.Results {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Results), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfDiagnosticInfos)
	lengthInBits += 32

	// Array field
	if len(m.DiagnosticInfos) > 0 {
		for _curItem, element := range m.DiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_CreateMonitoredItemsResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CreateMonitoredItemsResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__createMonitoredItemsResponse CreateMonitoredItemsResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CreateMonitoredItemsResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CreateMonitoredItemsResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("394")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	noOfResults, err := ReadSimpleField(ctx, "noOfResults", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfResults' field"))
	}
	m.NoOfResults = noOfResults

	results, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "results", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("748")), readBuffer), uint64(noOfResults))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'results' field"))
	}
	m.Results = results

	noOfDiagnosticInfos, err := ReadSimpleField(ctx, "noOfDiagnosticInfos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDiagnosticInfos' field"))
	}
	m.NoOfDiagnosticInfos = noOfDiagnosticInfos

	diagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "diagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'diagnosticInfos' field"))
	}
	m.DiagnosticInfos = diagnosticInfos

	if closeErr := readBuffer.CloseContext("CreateMonitoredItemsResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CreateMonitoredItemsResponse")
	}

	return m, nil
}

func (m *_CreateMonitoredItemsResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CreateMonitoredItemsResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CreateMonitoredItemsResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CreateMonitoredItemsResponse")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfResults", m.GetNoOfResults(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfResults' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "results", m.GetResults(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'results' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfDiagnosticInfos", m.GetNoOfDiagnosticInfos(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "diagnosticInfos", m.GetDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'diagnosticInfos' field")
		}

		if popErr := writeBuffer.PopContext("CreateMonitoredItemsResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CreateMonitoredItemsResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CreateMonitoredItemsResponse) IsCreateMonitoredItemsResponse() {}

func (m *_CreateMonitoredItemsResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
