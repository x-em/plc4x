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

// OpcuaNodeIdServicesVariableTrust is an enum
type OpcuaNodeIdServicesVariableTrust int32

type IOpcuaNodeIdServicesVariableTrust interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableTrust_TrustListType_LastUpdateTime                   OpcuaNodeIdServicesVariableTrust = 12542
	OpcuaNodeIdServicesVariableTrust_TrustListType_OpenWithMasks_InputArguments     OpcuaNodeIdServicesVariableTrust = 12544
	OpcuaNodeIdServicesVariableTrust_TrustListType_OpenWithMasks_OutputArguments    OpcuaNodeIdServicesVariableTrust = 12545
	OpcuaNodeIdServicesVariableTrust_TrustListType_CloseAndUpdate_OutputArguments   OpcuaNodeIdServicesVariableTrust = 12547
	OpcuaNodeIdServicesVariableTrust_TrustListType_AddCertificate_InputArguments    OpcuaNodeIdServicesVariableTrust = 12549
	OpcuaNodeIdServicesVariableTrust_TrustListType_RemoveCertificate_InputArguments OpcuaNodeIdServicesVariableTrust = 12551
	OpcuaNodeIdServicesVariableTrust_TrustListMasks_EnumValues                      OpcuaNodeIdServicesVariableTrust = 12553
	OpcuaNodeIdServicesVariableTrust_TrustListType_CloseAndUpdate_InputArguments    OpcuaNodeIdServicesVariableTrust = 12705
	OpcuaNodeIdServicesVariableTrust_TrustListType_UpdateFrequency                  OpcuaNodeIdServicesVariableTrust = 19296
	OpcuaNodeIdServicesVariableTrust_TrustListOutOfDateAlarmType_TrustListId        OpcuaNodeIdServicesVariableTrust = 19446
	OpcuaNodeIdServicesVariableTrust_TrustListOutOfDateAlarmType_LastUpdateTime     OpcuaNodeIdServicesVariableTrust = 19447
	OpcuaNodeIdServicesVariableTrust_TrustListOutOfDateAlarmType_UpdateFrequency    OpcuaNodeIdServicesVariableTrust = 19448
	OpcuaNodeIdServicesVariableTrust_TrustListType_DefaultValidationOptions         OpcuaNodeIdServicesVariableTrust = 23563
	OpcuaNodeIdServicesVariableTrust_TrustListValidationOptions_OptionSetValues     OpcuaNodeIdServicesVariableTrust = 23565
	OpcuaNodeIdServicesVariableTrust_TrustListType_ActivityTimeout                  OpcuaNodeIdServicesVariableTrust = 32254
	OpcuaNodeIdServicesVariableTrust_TrustListUpdatedAuditEventType_TrustListId     OpcuaNodeIdServicesVariableTrust = 32281
)

var OpcuaNodeIdServicesVariableTrustValues []OpcuaNodeIdServicesVariableTrust

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableTrustValues = []OpcuaNodeIdServicesVariableTrust{
		OpcuaNodeIdServicesVariableTrust_TrustListType_LastUpdateTime,
		OpcuaNodeIdServicesVariableTrust_TrustListType_OpenWithMasks_InputArguments,
		OpcuaNodeIdServicesVariableTrust_TrustListType_OpenWithMasks_OutputArguments,
		OpcuaNodeIdServicesVariableTrust_TrustListType_CloseAndUpdate_OutputArguments,
		OpcuaNodeIdServicesVariableTrust_TrustListType_AddCertificate_InputArguments,
		OpcuaNodeIdServicesVariableTrust_TrustListType_RemoveCertificate_InputArguments,
		OpcuaNodeIdServicesVariableTrust_TrustListMasks_EnumValues,
		OpcuaNodeIdServicesVariableTrust_TrustListType_CloseAndUpdate_InputArguments,
		OpcuaNodeIdServicesVariableTrust_TrustListType_UpdateFrequency,
		OpcuaNodeIdServicesVariableTrust_TrustListOutOfDateAlarmType_TrustListId,
		OpcuaNodeIdServicesVariableTrust_TrustListOutOfDateAlarmType_LastUpdateTime,
		OpcuaNodeIdServicesVariableTrust_TrustListOutOfDateAlarmType_UpdateFrequency,
		OpcuaNodeIdServicesVariableTrust_TrustListType_DefaultValidationOptions,
		OpcuaNodeIdServicesVariableTrust_TrustListValidationOptions_OptionSetValues,
		OpcuaNodeIdServicesVariableTrust_TrustListType_ActivityTimeout,
		OpcuaNodeIdServicesVariableTrust_TrustListUpdatedAuditEventType_TrustListId,
	}
}

func OpcuaNodeIdServicesVariableTrustByValue(value int32) (enum OpcuaNodeIdServicesVariableTrust, ok bool) {
	switch value {
	case 12542:
		return OpcuaNodeIdServicesVariableTrust_TrustListType_LastUpdateTime, true
	case 12544:
		return OpcuaNodeIdServicesVariableTrust_TrustListType_OpenWithMasks_InputArguments, true
	case 12545:
		return OpcuaNodeIdServicesVariableTrust_TrustListType_OpenWithMasks_OutputArguments, true
	case 12547:
		return OpcuaNodeIdServicesVariableTrust_TrustListType_CloseAndUpdate_OutputArguments, true
	case 12549:
		return OpcuaNodeIdServicesVariableTrust_TrustListType_AddCertificate_InputArguments, true
	case 12551:
		return OpcuaNodeIdServicesVariableTrust_TrustListType_RemoveCertificate_InputArguments, true
	case 12553:
		return OpcuaNodeIdServicesVariableTrust_TrustListMasks_EnumValues, true
	case 12705:
		return OpcuaNodeIdServicesVariableTrust_TrustListType_CloseAndUpdate_InputArguments, true
	case 19296:
		return OpcuaNodeIdServicesVariableTrust_TrustListType_UpdateFrequency, true
	case 19446:
		return OpcuaNodeIdServicesVariableTrust_TrustListOutOfDateAlarmType_TrustListId, true
	case 19447:
		return OpcuaNodeIdServicesVariableTrust_TrustListOutOfDateAlarmType_LastUpdateTime, true
	case 19448:
		return OpcuaNodeIdServicesVariableTrust_TrustListOutOfDateAlarmType_UpdateFrequency, true
	case 23563:
		return OpcuaNodeIdServicesVariableTrust_TrustListType_DefaultValidationOptions, true
	case 23565:
		return OpcuaNodeIdServicesVariableTrust_TrustListValidationOptions_OptionSetValues, true
	case 32254:
		return OpcuaNodeIdServicesVariableTrust_TrustListType_ActivityTimeout, true
	case 32281:
		return OpcuaNodeIdServicesVariableTrust_TrustListUpdatedAuditEventType_TrustListId, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableTrustByName(value string) (enum OpcuaNodeIdServicesVariableTrust, ok bool) {
	switch value {
	case "TrustListType_LastUpdateTime":
		return OpcuaNodeIdServicesVariableTrust_TrustListType_LastUpdateTime, true
	case "TrustListType_OpenWithMasks_InputArguments":
		return OpcuaNodeIdServicesVariableTrust_TrustListType_OpenWithMasks_InputArguments, true
	case "TrustListType_OpenWithMasks_OutputArguments":
		return OpcuaNodeIdServicesVariableTrust_TrustListType_OpenWithMasks_OutputArguments, true
	case "TrustListType_CloseAndUpdate_OutputArguments":
		return OpcuaNodeIdServicesVariableTrust_TrustListType_CloseAndUpdate_OutputArguments, true
	case "TrustListType_AddCertificate_InputArguments":
		return OpcuaNodeIdServicesVariableTrust_TrustListType_AddCertificate_InputArguments, true
	case "TrustListType_RemoveCertificate_InputArguments":
		return OpcuaNodeIdServicesVariableTrust_TrustListType_RemoveCertificate_InputArguments, true
	case "TrustListMasks_EnumValues":
		return OpcuaNodeIdServicesVariableTrust_TrustListMasks_EnumValues, true
	case "TrustListType_CloseAndUpdate_InputArguments":
		return OpcuaNodeIdServicesVariableTrust_TrustListType_CloseAndUpdate_InputArguments, true
	case "TrustListType_UpdateFrequency":
		return OpcuaNodeIdServicesVariableTrust_TrustListType_UpdateFrequency, true
	case "TrustListOutOfDateAlarmType_TrustListId":
		return OpcuaNodeIdServicesVariableTrust_TrustListOutOfDateAlarmType_TrustListId, true
	case "TrustListOutOfDateAlarmType_LastUpdateTime":
		return OpcuaNodeIdServicesVariableTrust_TrustListOutOfDateAlarmType_LastUpdateTime, true
	case "TrustListOutOfDateAlarmType_UpdateFrequency":
		return OpcuaNodeIdServicesVariableTrust_TrustListOutOfDateAlarmType_UpdateFrequency, true
	case "TrustListType_DefaultValidationOptions":
		return OpcuaNodeIdServicesVariableTrust_TrustListType_DefaultValidationOptions, true
	case "TrustListValidationOptions_OptionSetValues":
		return OpcuaNodeIdServicesVariableTrust_TrustListValidationOptions_OptionSetValues, true
	case "TrustListType_ActivityTimeout":
		return OpcuaNodeIdServicesVariableTrust_TrustListType_ActivityTimeout, true
	case "TrustListUpdatedAuditEventType_TrustListId":
		return OpcuaNodeIdServicesVariableTrust_TrustListUpdatedAuditEventType_TrustListId, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableTrustKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableTrustValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableTrust(structType any) OpcuaNodeIdServicesVariableTrust {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableTrust {
		if sOpcuaNodeIdServicesVariableTrust, ok := typ.(OpcuaNodeIdServicesVariableTrust); ok {
			return sOpcuaNodeIdServicesVariableTrust
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableTrust) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableTrust) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableTrustParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableTrust, error) {
	return OpcuaNodeIdServicesVariableTrustParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableTrustParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableTrust, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableTrust", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableTrust")
	}
	if enum, ok := OpcuaNodeIdServicesVariableTrustByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableTrust")
		return OpcuaNodeIdServicesVariableTrust(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableTrust) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableTrust) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableTrust", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableTrust) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableTrust) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableTrust_TrustListType_LastUpdateTime:
		return "TrustListType_LastUpdateTime"
	case OpcuaNodeIdServicesVariableTrust_TrustListType_OpenWithMasks_InputArguments:
		return "TrustListType_OpenWithMasks_InputArguments"
	case OpcuaNodeIdServicesVariableTrust_TrustListType_OpenWithMasks_OutputArguments:
		return "TrustListType_OpenWithMasks_OutputArguments"
	case OpcuaNodeIdServicesVariableTrust_TrustListType_CloseAndUpdate_OutputArguments:
		return "TrustListType_CloseAndUpdate_OutputArguments"
	case OpcuaNodeIdServicesVariableTrust_TrustListType_AddCertificate_InputArguments:
		return "TrustListType_AddCertificate_InputArguments"
	case OpcuaNodeIdServicesVariableTrust_TrustListType_RemoveCertificate_InputArguments:
		return "TrustListType_RemoveCertificate_InputArguments"
	case OpcuaNodeIdServicesVariableTrust_TrustListMasks_EnumValues:
		return "TrustListMasks_EnumValues"
	case OpcuaNodeIdServicesVariableTrust_TrustListType_CloseAndUpdate_InputArguments:
		return "TrustListType_CloseAndUpdate_InputArguments"
	case OpcuaNodeIdServicesVariableTrust_TrustListType_UpdateFrequency:
		return "TrustListType_UpdateFrequency"
	case OpcuaNodeIdServicesVariableTrust_TrustListOutOfDateAlarmType_TrustListId:
		return "TrustListOutOfDateAlarmType_TrustListId"
	case OpcuaNodeIdServicesVariableTrust_TrustListOutOfDateAlarmType_LastUpdateTime:
		return "TrustListOutOfDateAlarmType_LastUpdateTime"
	case OpcuaNodeIdServicesVariableTrust_TrustListOutOfDateAlarmType_UpdateFrequency:
		return "TrustListOutOfDateAlarmType_UpdateFrequency"
	case OpcuaNodeIdServicesVariableTrust_TrustListType_DefaultValidationOptions:
		return "TrustListType_DefaultValidationOptions"
	case OpcuaNodeIdServicesVariableTrust_TrustListValidationOptions_OptionSetValues:
		return "TrustListValidationOptions_OptionSetValues"
	case OpcuaNodeIdServicesVariableTrust_TrustListType_ActivityTimeout:
		return "TrustListType_ActivityTimeout"
	case OpcuaNodeIdServicesVariableTrust_TrustListUpdatedAuditEventType_TrustListId:
		return "TrustListUpdatedAuditEventType_TrustListId"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableTrust) String() string {
	return e.PLC4XEnumName()
}
