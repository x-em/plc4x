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

// OpcuaNodeIdServicesVariableDefault is an enum
type OpcuaNodeIdServicesVariableDefault int32

type IOpcuaNodeIdServicesVariableDefault interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableDefault_DefaultInstanceBrowseName                                            OpcuaNodeIdServicesVariableDefault = 17605
	OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_AggregateConfiguration_TreatUncertainAsBad    OpcuaNodeIdServicesVariableDefault = 32639
	OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_AggregateConfiguration_PercentDataBad         OpcuaNodeIdServicesVariableDefault = 32640
	OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_AggregateConfiguration_PercentDataGood        OpcuaNodeIdServicesVariableDefault = 32641
	OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_AggregateConfiguration_UseSlopedExtrapolation OpcuaNodeIdServicesVariableDefault = 32642
	OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_Stepped                                       OpcuaNodeIdServicesVariableDefault = 32644
	OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_Definition                                    OpcuaNodeIdServicesVariableDefault = 32645
	OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_MaxTimeInterval                               OpcuaNodeIdServicesVariableDefault = 32646
	OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_MinTimeInterval                               OpcuaNodeIdServicesVariableDefault = 32647
	OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_ExceptionDeviation                            OpcuaNodeIdServicesVariableDefault = 32648
	OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_ExceptionDeviationFormat                      OpcuaNodeIdServicesVariableDefault = 32649
	OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_StartOfArchive                                OpcuaNodeIdServicesVariableDefault = 32650
	OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_StartOfOnlineArchive                          OpcuaNodeIdServicesVariableDefault = 32656
	OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_ServerTimestampSupported                      OpcuaNodeIdServicesVariableDefault = 32682
	OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_MaxTimeStoredValues                           OpcuaNodeIdServicesVariableDefault = 32752
	OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_MaxCountStoredValues                          OpcuaNodeIdServicesVariableDefault = 32753
	OpcuaNodeIdServicesVariableDefault_DefaultHEConfiguration_StartOfArchive                                OpcuaNodeIdServicesVariableDefault = 32756
	OpcuaNodeIdServicesVariableDefault_DefaultHEConfiguration_StartOfOnlineArchive                          OpcuaNodeIdServicesVariableDefault = 32757
)

var OpcuaNodeIdServicesVariableDefaultValues []OpcuaNodeIdServicesVariableDefault

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableDefaultValues = []OpcuaNodeIdServicesVariableDefault{
		OpcuaNodeIdServicesVariableDefault_DefaultInstanceBrowseName,
		OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_AggregateConfiguration_TreatUncertainAsBad,
		OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_AggregateConfiguration_PercentDataBad,
		OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_AggregateConfiguration_PercentDataGood,
		OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_AggregateConfiguration_UseSlopedExtrapolation,
		OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_Stepped,
		OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_Definition,
		OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_MaxTimeInterval,
		OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_MinTimeInterval,
		OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_ExceptionDeviation,
		OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_ExceptionDeviationFormat,
		OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_StartOfArchive,
		OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_StartOfOnlineArchive,
		OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_ServerTimestampSupported,
		OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_MaxTimeStoredValues,
		OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_MaxCountStoredValues,
		OpcuaNodeIdServicesVariableDefault_DefaultHEConfiguration_StartOfArchive,
		OpcuaNodeIdServicesVariableDefault_DefaultHEConfiguration_StartOfOnlineArchive,
	}
}

func OpcuaNodeIdServicesVariableDefaultByValue(value int32) (enum OpcuaNodeIdServicesVariableDefault, ok bool) {
	switch value {
	case 17605:
		return OpcuaNodeIdServicesVariableDefault_DefaultInstanceBrowseName, true
	case 32639:
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_AggregateConfiguration_TreatUncertainAsBad, true
	case 32640:
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_AggregateConfiguration_PercentDataBad, true
	case 32641:
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_AggregateConfiguration_PercentDataGood, true
	case 32642:
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_AggregateConfiguration_UseSlopedExtrapolation, true
	case 32644:
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_Stepped, true
	case 32645:
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_Definition, true
	case 32646:
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_MaxTimeInterval, true
	case 32647:
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_MinTimeInterval, true
	case 32648:
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_ExceptionDeviation, true
	case 32649:
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_ExceptionDeviationFormat, true
	case 32650:
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_StartOfArchive, true
	case 32656:
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_StartOfOnlineArchive, true
	case 32682:
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_ServerTimestampSupported, true
	case 32752:
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_MaxTimeStoredValues, true
	case 32753:
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_MaxCountStoredValues, true
	case 32756:
		return OpcuaNodeIdServicesVariableDefault_DefaultHEConfiguration_StartOfArchive, true
	case 32757:
		return OpcuaNodeIdServicesVariableDefault_DefaultHEConfiguration_StartOfOnlineArchive, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableDefaultByName(value string) (enum OpcuaNodeIdServicesVariableDefault, ok bool) {
	switch value {
	case "DefaultInstanceBrowseName":
		return OpcuaNodeIdServicesVariableDefault_DefaultInstanceBrowseName, true
	case "DefaultHAConfiguration_AggregateConfiguration_TreatUncertainAsBad":
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_AggregateConfiguration_TreatUncertainAsBad, true
	case "DefaultHAConfiguration_AggregateConfiguration_PercentDataBad":
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_AggregateConfiguration_PercentDataBad, true
	case "DefaultHAConfiguration_AggregateConfiguration_PercentDataGood":
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_AggregateConfiguration_PercentDataGood, true
	case "DefaultHAConfiguration_AggregateConfiguration_UseSlopedExtrapolation":
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_AggregateConfiguration_UseSlopedExtrapolation, true
	case "DefaultHAConfiguration_Stepped":
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_Stepped, true
	case "DefaultHAConfiguration_Definition":
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_Definition, true
	case "DefaultHAConfiguration_MaxTimeInterval":
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_MaxTimeInterval, true
	case "DefaultHAConfiguration_MinTimeInterval":
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_MinTimeInterval, true
	case "DefaultHAConfiguration_ExceptionDeviation":
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_ExceptionDeviation, true
	case "DefaultHAConfiguration_ExceptionDeviationFormat":
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_ExceptionDeviationFormat, true
	case "DefaultHAConfiguration_StartOfArchive":
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_StartOfArchive, true
	case "DefaultHAConfiguration_StartOfOnlineArchive":
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_StartOfOnlineArchive, true
	case "DefaultHAConfiguration_ServerTimestampSupported":
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_ServerTimestampSupported, true
	case "DefaultHAConfiguration_MaxTimeStoredValues":
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_MaxTimeStoredValues, true
	case "DefaultHAConfiguration_MaxCountStoredValues":
		return OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_MaxCountStoredValues, true
	case "DefaultHEConfiguration_StartOfArchive":
		return OpcuaNodeIdServicesVariableDefault_DefaultHEConfiguration_StartOfArchive, true
	case "DefaultHEConfiguration_StartOfOnlineArchive":
		return OpcuaNodeIdServicesVariableDefault_DefaultHEConfiguration_StartOfOnlineArchive, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableDefaultKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableDefaultValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableDefault(structType any) OpcuaNodeIdServicesVariableDefault {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableDefault {
		if sOpcuaNodeIdServicesVariableDefault, ok := typ.(OpcuaNodeIdServicesVariableDefault); ok {
			return sOpcuaNodeIdServicesVariableDefault
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableDefault) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableDefault) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableDefaultParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableDefault, error) {
	return OpcuaNodeIdServicesVariableDefaultParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableDefaultParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableDefault, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableDefault", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableDefault")
	}
	if enum, ok := OpcuaNodeIdServicesVariableDefaultByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableDefault")
		return OpcuaNodeIdServicesVariableDefault(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableDefault) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableDefault) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableDefault", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableDefault) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableDefault) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableDefault_DefaultInstanceBrowseName:
		return "DefaultInstanceBrowseName"
	case OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_AggregateConfiguration_TreatUncertainAsBad:
		return "DefaultHAConfiguration_AggregateConfiguration_TreatUncertainAsBad"
	case OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_AggregateConfiguration_PercentDataBad:
		return "DefaultHAConfiguration_AggregateConfiguration_PercentDataBad"
	case OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_AggregateConfiguration_PercentDataGood:
		return "DefaultHAConfiguration_AggregateConfiguration_PercentDataGood"
	case OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_AggregateConfiguration_UseSlopedExtrapolation:
		return "DefaultHAConfiguration_AggregateConfiguration_UseSlopedExtrapolation"
	case OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_Stepped:
		return "DefaultHAConfiguration_Stepped"
	case OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_Definition:
		return "DefaultHAConfiguration_Definition"
	case OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_MaxTimeInterval:
		return "DefaultHAConfiguration_MaxTimeInterval"
	case OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_MinTimeInterval:
		return "DefaultHAConfiguration_MinTimeInterval"
	case OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_ExceptionDeviation:
		return "DefaultHAConfiguration_ExceptionDeviation"
	case OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_ExceptionDeviationFormat:
		return "DefaultHAConfiguration_ExceptionDeviationFormat"
	case OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_StartOfArchive:
		return "DefaultHAConfiguration_StartOfArchive"
	case OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_StartOfOnlineArchive:
		return "DefaultHAConfiguration_StartOfOnlineArchive"
	case OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_ServerTimestampSupported:
		return "DefaultHAConfiguration_ServerTimestampSupported"
	case OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_MaxTimeStoredValues:
		return "DefaultHAConfiguration_MaxTimeStoredValues"
	case OpcuaNodeIdServicesVariableDefault_DefaultHAConfiguration_MaxCountStoredValues:
		return "DefaultHAConfiguration_MaxCountStoredValues"
	case OpcuaNodeIdServicesVariableDefault_DefaultHEConfiguration_StartOfArchive:
		return "DefaultHEConfiguration_StartOfArchive"
	case OpcuaNodeIdServicesVariableDefault_DefaultHEConfiguration_StartOfOnlineArchive:
		return "DefaultHEConfiguration_StartOfOnlineArchive"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableDefault) String() string {
	return e.PLC4XEnumName()
}