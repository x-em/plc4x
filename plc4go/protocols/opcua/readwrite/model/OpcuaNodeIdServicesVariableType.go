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

// OpcuaNodeIdServicesVariableType is an enum
type OpcuaNodeIdServicesVariableType int32

type IOpcuaNodeIdServicesVariableType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableType_BaseVariableType                          OpcuaNodeIdServicesVariableType = 62
	OpcuaNodeIdServicesVariableType_BaseDataVariableType                      OpcuaNodeIdServicesVariableType = 63
	OpcuaNodeIdServicesVariableType_PropertyType                              OpcuaNodeIdServicesVariableType = 68
	OpcuaNodeIdServicesVariableType_DataTypeDescriptionType                   OpcuaNodeIdServicesVariableType = 69
	OpcuaNodeIdServicesVariableType_DataTypeDictionaryType                    OpcuaNodeIdServicesVariableType = 72
	OpcuaNodeIdServicesVariableType_ServerVendorCapabilityType                OpcuaNodeIdServicesVariableType = 2137
	OpcuaNodeIdServicesVariableType_ServerStatusType                          OpcuaNodeIdServicesVariableType = 2138
	OpcuaNodeIdServicesVariableType_ServerDiagnosticsSummaryType              OpcuaNodeIdServicesVariableType = 2150
	OpcuaNodeIdServicesVariableType_SamplingIntervalDiagnosticsArrayType      OpcuaNodeIdServicesVariableType = 2164
	OpcuaNodeIdServicesVariableType_SamplingIntervalDiagnosticsType           OpcuaNodeIdServicesVariableType = 2165
	OpcuaNodeIdServicesVariableType_SubscriptionDiagnosticsArrayType          OpcuaNodeIdServicesVariableType = 2171
	OpcuaNodeIdServicesVariableType_SubscriptionDiagnosticsType               OpcuaNodeIdServicesVariableType = 2172
	OpcuaNodeIdServicesVariableType_SessionDiagnosticsArrayType               OpcuaNodeIdServicesVariableType = 2196
	OpcuaNodeIdServicesVariableType_SessionDiagnosticsVariableType            OpcuaNodeIdServicesVariableType = 2197
	OpcuaNodeIdServicesVariableType_SessionSecurityDiagnosticsArrayType       OpcuaNodeIdServicesVariableType = 2243
	OpcuaNodeIdServicesVariableType_SessionSecurityDiagnosticsType            OpcuaNodeIdServicesVariableType = 2244
	OpcuaNodeIdServicesVariableType_DataItemType                              OpcuaNodeIdServicesVariableType = 2365
	OpcuaNodeIdServicesVariableType_AnalogItemType                            OpcuaNodeIdServicesVariableType = 2368
	OpcuaNodeIdServicesVariableType_DiscreteItemType                          OpcuaNodeIdServicesVariableType = 2372
	OpcuaNodeIdServicesVariableType_TwoStateDiscreteType                      OpcuaNodeIdServicesVariableType = 2373
	OpcuaNodeIdServicesVariableType_MultiStateDiscreteType                    OpcuaNodeIdServicesVariableType = 2376
	OpcuaNodeIdServicesVariableType_ProgramDiagnosticType                     OpcuaNodeIdServicesVariableType = 2380
	OpcuaNodeIdServicesVariableType_StateVariableType                         OpcuaNodeIdServicesVariableType = 2755
	OpcuaNodeIdServicesVariableType_FiniteStateVariableType                   OpcuaNodeIdServicesVariableType = 2760
	OpcuaNodeIdServicesVariableType_TransitionVariableType                    OpcuaNodeIdServicesVariableType = 2762
	OpcuaNodeIdServicesVariableType_FiniteTransitionVariableType              OpcuaNodeIdServicesVariableType = 2767
	OpcuaNodeIdServicesVariableType_BuildInfoType                             OpcuaNodeIdServicesVariableType = 3051
	OpcuaNodeIdServicesVariableType_TwoStateVariableType                      OpcuaNodeIdServicesVariableType = 8995
	OpcuaNodeIdServicesVariableType_ConditionVariableType                     OpcuaNodeIdServicesVariableType = 9002
	OpcuaNodeIdServicesVariableType_MultiStateValueDiscreteType               OpcuaNodeIdServicesVariableType = 11238
	OpcuaNodeIdServicesVariableType_OptionSetType                             OpcuaNodeIdServicesVariableType = 11487
	OpcuaNodeIdServicesVariableType_ArrayItemType                             OpcuaNodeIdServicesVariableType = 12021
	OpcuaNodeIdServicesVariableType_YArrayItemType                            OpcuaNodeIdServicesVariableType = 12029
	OpcuaNodeIdServicesVariableType_XYArrayItemType                           OpcuaNodeIdServicesVariableType = 12038
	OpcuaNodeIdServicesVariableType_ImageItemType                             OpcuaNodeIdServicesVariableType = 12047
	OpcuaNodeIdServicesVariableType_CubeItemType                              OpcuaNodeIdServicesVariableType = 12057
	OpcuaNodeIdServicesVariableType_NDimensionArrayItemType                   OpcuaNodeIdServicesVariableType = 12068
	OpcuaNodeIdServicesVariableType_GuardVariableType                         OpcuaNodeIdServicesVariableType = 15113
	OpcuaNodeIdServicesVariableType_ExpressionGuardVariableType               OpcuaNodeIdServicesVariableType = 15128
	OpcuaNodeIdServicesVariableType_ElseGuardVariableType                     OpcuaNodeIdServicesVariableType = 15317
	OpcuaNodeIdServicesVariableType_BaseAnalogType                            OpcuaNodeIdServicesVariableType = 15318
	OpcuaNodeIdServicesVariableType_ProgramDiagnostic2Type                    OpcuaNodeIdServicesVariableType = 15383
	OpcuaNodeIdServicesVariableType_SelectionListType                         OpcuaNodeIdServicesVariableType = 16309
	OpcuaNodeIdServicesVariableType_AlarmRateVariableType                     OpcuaNodeIdServicesVariableType = 17277
	OpcuaNodeIdServicesVariableType_AnalogUnitType                            OpcuaNodeIdServicesVariableType = 17497
	OpcuaNodeIdServicesVariableType_AnalogUnitRangeType                       OpcuaNodeIdServicesVariableType = 17570
	OpcuaNodeIdServicesVariableType_RationalNumberType                        OpcuaNodeIdServicesVariableType = 17709
	OpcuaNodeIdServicesVariableType_VectorType                                OpcuaNodeIdServicesVariableType = 17714
	OpcuaNodeIdServicesVariableType_ThreeDVectorType                          OpcuaNodeIdServicesVariableType = 17716
	OpcuaNodeIdServicesVariableType_AudioVariableType                         OpcuaNodeIdServicesVariableType = 17986
	OpcuaNodeIdServicesVariableType_CartesianCoordinatesType                  OpcuaNodeIdServicesVariableType = 18772
	OpcuaNodeIdServicesVariableType_ThreeDCartesianCoordinatesType            OpcuaNodeIdServicesVariableType = 18774
	OpcuaNodeIdServicesVariableType_OrientationType                           OpcuaNodeIdServicesVariableType = 18779
	OpcuaNodeIdServicesVariableType_ThreeDOrientationType                     OpcuaNodeIdServicesVariableType = 18781
	OpcuaNodeIdServicesVariableType_FrameType                                 OpcuaNodeIdServicesVariableType = 18786
	OpcuaNodeIdServicesVariableType_ThreeDFrameType                           OpcuaNodeIdServicesVariableType = 18791
	OpcuaNodeIdServicesVariableType_MultiStateDictionaryEntryDiscreteBaseType OpcuaNodeIdServicesVariableType = 19077
	OpcuaNodeIdServicesVariableType_MultiStateDictionaryEntryDiscreteType     OpcuaNodeIdServicesVariableType = 19084
	OpcuaNodeIdServicesVariableType_PubSubDiagnosticsCounterType              OpcuaNodeIdServicesVariableType = 19725
	OpcuaNodeIdServicesVariableType_AlarmStateVariableType                    OpcuaNodeIdServicesVariableType = 32244
	OpcuaNodeIdServicesVariableType_BitFieldType                              OpcuaNodeIdServicesVariableType = 32431
	OpcuaNodeIdServicesVariableType_ReferenceDescriptionVariableType          OpcuaNodeIdServicesVariableType = 32657
)

var OpcuaNodeIdServicesVariableTypeValues []OpcuaNodeIdServicesVariableType

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableTypeValues = []OpcuaNodeIdServicesVariableType{
		OpcuaNodeIdServicesVariableType_BaseVariableType,
		OpcuaNodeIdServicesVariableType_BaseDataVariableType,
		OpcuaNodeIdServicesVariableType_PropertyType,
		OpcuaNodeIdServicesVariableType_DataTypeDescriptionType,
		OpcuaNodeIdServicesVariableType_DataTypeDictionaryType,
		OpcuaNodeIdServicesVariableType_ServerVendorCapabilityType,
		OpcuaNodeIdServicesVariableType_ServerStatusType,
		OpcuaNodeIdServicesVariableType_ServerDiagnosticsSummaryType,
		OpcuaNodeIdServicesVariableType_SamplingIntervalDiagnosticsArrayType,
		OpcuaNodeIdServicesVariableType_SamplingIntervalDiagnosticsType,
		OpcuaNodeIdServicesVariableType_SubscriptionDiagnosticsArrayType,
		OpcuaNodeIdServicesVariableType_SubscriptionDiagnosticsType,
		OpcuaNodeIdServicesVariableType_SessionDiagnosticsArrayType,
		OpcuaNodeIdServicesVariableType_SessionDiagnosticsVariableType,
		OpcuaNodeIdServicesVariableType_SessionSecurityDiagnosticsArrayType,
		OpcuaNodeIdServicesVariableType_SessionSecurityDiagnosticsType,
		OpcuaNodeIdServicesVariableType_DataItemType,
		OpcuaNodeIdServicesVariableType_AnalogItemType,
		OpcuaNodeIdServicesVariableType_DiscreteItemType,
		OpcuaNodeIdServicesVariableType_TwoStateDiscreteType,
		OpcuaNodeIdServicesVariableType_MultiStateDiscreteType,
		OpcuaNodeIdServicesVariableType_ProgramDiagnosticType,
		OpcuaNodeIdServicesVariableType_StateVariableType,
		OpcuaNodeIdServicesVariableType_FiniteStateVariableType,
		OpcuaNodeIdServicesVariableType_TransitionVariableType,
		OpcuaNodeIdServicesVariableType_FiniteTransitionVariableType,
		OpcuaNodeIdServicesVariableType_BuildInfoType,
		OpcuaNodeIdServicesVariableType_TwoStateVariableType,
		OpcuaNodeIdServicesVariableType_ConditionVariableType,
		OpcuaNodeIdServicesVariableType_MultiStateValueDiscreteType,
		OpcuaNodeIdServicesVariableType_OptionSetType,
		OpcuaNodeIdServicesVariableType_ArrayItemType,
		OpcuaNodeIdServicesVariableType_YArrayItemType,
		OpcuaNodeIdServicesVariableType_XYArrayItemType,
		OpcuaNodeIdServicesVariableType_ImageItemType,
		OpcuaNodeIdServicesVariableType_CubeItemType,
		OpcuaNodeIdServicesVariableType_NDimensionArrayItemType,
		OpcuaNodeIdServicesVariableType_GuardVariableType,
		OpcuaNodeIdServicesVariableType_ExpressionGuardVariableType,
		OpcuaNodeIdServicesVariableType_ElseGuardVariableType,
		OpcuaNodeIdServicesVariableType_BaseAnalogType,
		OpcuaNodeIdServicesVariableType_ProgramDiagnostic2Type,
		OpcuaNodeIdServicesVariableType_SelectionListType,
		OpcuaNodeIdServicesVariableType_AlarmRateVariableType,
		OpcuaNodeIdServicesVariableType_AnalogUnitType,
		OpcuaNodeIdServicesVariableType_AnalogUnitRangeType,
		OpcuaNodeIdServicesVariableType_RationalNumberType,
		OpcuaNodeIdServicesVariableType_VectorType,
		OpcuaNodeIdServicesVariableType_ThreeDVectorType,
		OpcuaNodeIdServicesVariableType_AudioVariableType,
		OpcuaNodeIdServicesVariableType_CartesianCoordinatesType,
		OpcuaNodeIdServicesVariableType_ThreeDCartesianCoordinatesType,
		OpcuaNodeIdServicesVariableType_OrientationType,
		OpcuaNodeIdServicesVariableType_ThreeDOrientationType,
		OpcuaNodeIdServicesVariableType_FrameType,
		OpcuaNodeIdServicesVariableType_ThreeDFrameType,
		OpcuaNodeIdServicesVariableType_MultiStateDictionaryEntryDiscreteBaseType,
		OpcuaNodeIdServicesVariableType_MultiStateDictionaryEntryDiscreteType,
		OpcuaNodeIdServicesVariableType_PubSubDiagnosticsCounterType,
		OpcuaNodeIdServicesVariableType_AlarmStateVariableType,
		OpcuaNodeIdServicesVariableType_BitFieldType,
		OpcuaNodeIdServicesVariableType_ReferenceDescriptionVariableType,
	}
}

func OpcuaNodeIdServicesVariableTypeByValue(value int32) (enum OpcuaNodeIdServicesVariableType, ok bool) {
	switch value {
	case 11238:
		return OpcuaNodeIdServicesVariableType_MultiStateValueDiscreteType, true
	case 11487:
		return OpcuaNodeIdServicesVariableType_OptionSetType, true
	case 12021:
		return OpcuaNodeIdServicesVariableType_ArrayItemType, true
	case 12029:
		return OpcuaNodeIdServicesVariableType_YArrayItemType, true
	case 12038:
		return OpcuaNodeIdServicesVariableType_XYArrayItemType, true
	case 12047:
		return OpcuaNodeIdServicesVariableType_ImageItemType, true
	case 12057:
		return OpcuaNodeIdServicesVariableType_CubeItemType, true
	case 12068:
		return OpcuaNodeIdServicesVariableType_NDimensionArrayItemType, true
	case 15113:
		return OpcuaNodeIdServicesVariableType_GuardVariableType, true
	case 15128:
		return OpcuaNodeIdServicesVariableType_ExpressionGuardVariableType, true
	case 15317:
		return OpcuaNodeIdServicesVariableType_ElseGuardVariableType, true
	case 15318:
		return OpcuaNodeIdServicesVariableType_BaseAnalogType, true
	case 15383:
		return OpcuaNodeIdServicesVariableType_ProgramDiagnostic2Type, true
	case 16309:
		return OpcuaNodeIdServicesVariableType_SelectionListType, true
	case 17277:
		return OpcuaNodeIdServicesVariableType_AlarmRateVariableType, true
	case 17497:
		return OpcuaNodeIdServicesVariableType_AnalogUnitType, true
	case 17570:
		return OpcuaNodeIdServicesVariableType_AnalogUnitRangeType, true
	case 17709:
		return OpcuaNodeIdServicesVariableType_RationalNumberType, true
	case 17714:
		return OpcuaNodeIdServicesVariableType_VectorType, true
	case 17716:
		return OpcuaNodeIdServicesVariableType_ThreeDVectorType, true
	case 17986:
		return OpcuaNodeIdServicesVariableType_AudioVariableType, true
	case 18772:
		return OpcuaNodeIdServicesVariableType_CartesianCoordinatesType, true
	case 18774:
		return OpcuaNodeIdServicesVariableType_ThreeDCartesianCoordinatesType, true
	case 18779:
		return OpcuaNodeIdServicesVariableType_OrientationType, true
	case 18781:
		return OpcuaNodeIdServicesVariableType_ThreeDOrientationType, true
	case 18786:
		return OpcuaNodeIdServicesVariableType_FrameType, true
	case 18791:
		return OpcuaNodeIdServicesVariableType_ThreeDFrameType, true
	case 19077:
		return OpcuaNodeIdServicesVariableType_MultiStateDictionaryEntryDiscreteBaseType, true
	case 19084:
		return OpcuaNodeIdServicesVariableType_MultiStateDictionaryEntryDiscreteType, true
	case 19725:
		return OpcuaNodeIdServicesVariableType_PubSubDiagnosticsCounterType, true
	case 2137:
		return OpcuaNodeIdServicesVariableType_ServerVendorCapabilityType, true
	case 2138:
		return OpcuaNodeIdServicesVariableType_ServerStatusType, true
	case 2150:
		return OpcuaNodeIdServicesVariableType_ServerDiagnosticsSummaryType, true
	case 2164:
		return OpcuaNodeIdServicesVariableType_SamplingIntervalDiagnosticsArrayType, true
	case 2165:
		return OpcuaNodeIdServicesVariableType_SamplingIntervalDiagnosticsType, true
	case 2171:
		return OpcuaNodeIdServicesVariableType_SubscriptionDiagnosticsArrayType, true
	case 2172:
		return OpcuaNodeIdServicesVariableType_SubscriptionDiagnosticsType, true
	case 2196:
		return OpcuaNodeIdServicesVariableType_SessionDiagnosticsArrayType, true
	case 2197:
		return OpcuaNodeIdServicesVariableType_SessionDiagnosticsVariableType, true
	case 2243:
		return OpcuaNodeIdServicesVariableType_SessionSecurityDiagnosticsArrayType, true
	case 2244:
		return OpcuaNodeIdServicesVariableType_SessionSecurityDiagnosticsType, true
	case 2365:
		return OpcuaNodeIdServicesVariableType_DataItemType, true
	case 2368:
		return OpcuaNodeIdServicesVariableType_AnalogItemType, true
	case 2372:
		return OpcuaNodeIdServicesVariableType_DiscreteItemType, true
	case 2373:
		return OpcuaNodeIdServicesVariableType_TwoStateDiscreteType, true
	case 2376:
		return OpcuaNodeIdServicesVariableType_MultiStateDiscreteType, true
	case 2380:
		return OpcuaNodeIdServicesVariableType_ProgramDiagnosticType, true
	case 2755:
		return OpcuaNodeIdServicesVariableType_StateVariableType, true
	case 2760:
		return OpcuaNodeIdServicesVariableType_FiniteStateVariableType, true
	case 2762:
		return OpcuaNodeIdServicesVariableType_TransitionVariableType, true
	case 2767:
		return OpcuaNodeIdServicesVariableType_FiniteTransitionVariableType, true
	case 3051:
		return OpcuaNodeIdServicesVariableType_BuildInfoType, true
	case 32244:
		return OpcuaNodeIdServicesVariableType_AlarmStateVariableType, true
	case 32431:
		return OpcuaNodeIdServicesVariableType_BitFieldType, true
	case 32657:
		return OpcuaNodeIdServicesVariableType_ReferenceDescriptionVariableType, true
	case 62:
		return OpcuaNodeIdServicesVariableType_BaseVariableType, true
	case 63:
		return OpcuaNodeIdServicesVariableType_BaseDataVariableType, true
	case 68:
		return OpcuaNodeIdServicesVariableType_PropertyType, true
	case 69:
		return OpcuaNodeIdServicesVariableType_DataTypeDescriptionType, true
	case 72:
		return OpcuaNodeIdServicesVariableType_DataTypeDictionaryType, true
	case 8995:
		return OpcuaNodeIdServicesVariableType_TwoStateVariableType, true
	case 9002:
		return OpcuaNodeIdServicesVariableType_ConditionVariableType, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableTypeByName(value string) (enum OpcuaNodeIdServicesVariableType, ok bool) {
	switch value {
	case "MultiStateValueDiscreteType":
		return OpcuaNodeIdServicesVariableType_MultiStateValueDiscreteType, true
	case "OptionSetType":
		return OpcuaNodeIdServicesVariableType_OptionSetType, true
	case "ArrayItemType":
		return OpcuaNodeIdServicesVariableType_ArrayItemType, true
	case "YArrayItemType":
		return OpcuaNodeIdServicesVariableType_YArrayItemType, true
	case "XYArrayItemType":
		return OpcuaNodeIdServicesVariableType_XYArrayItemType, true
	case "ImageItemType":
		return OpcuaNodeIdServicesVariableType_ImageItemType, true
	case "CubeItemType":
		return OpcuaNodeIdServicesVariableType_CubeItemType, true
	case "NDimensionArrayItemType":
		return OpcuaNodeIdServicesVariableType_NDimensionArrayItemType, true
	case "GuardVariableType":
		return OpcuaNodeIdServicesVariableType_GuardVariableType, true
	case "ExpressionGuardVariableType":
		return OpcuaNodeIdServicesVariableType_ExpressionGuardVariableType, true
	case "ElseGuardVariableType":
		return OpcuaNodeIdServicesVariableType_ElseGuardVariableType, true
	case "BaseAnalogType":
		return OpcuaNodeIdServicesVariableType_BaseAnalogType, true
	case "ProgramDiagnostic2Type":
		return OpcuaNodeIdServicesVariableType_ProgramDiagnostic2Type, true
	case "SelectionListType":
		return OpcuaNodeIdServicesVariableType_SelectionListType, true
	case "AlarmRateVariableType":
		return OpcuaNodeIdServicesVariableType_AlarmRateVariableType, true
	case "AnalogUnitType":
		return OpcuaNodeIdServicesVariableType_AnalogUnitType, true
	case "AnalogUnitRangeType":
		return OpcuaNodeIdServicesVariableType_AnalogUnitRangeType, true
	case "RationalNumberType":
		return OpcuaNodeIdServicesVariableType_RationalNumberType, true
	case "VectorType":
		return OpcuaNodeIdServicesVariableType_VectorType, true
	case "ThreeDVectorType":
		return OpcuaNodeIdServicesVariableType_ThreeDVectorType, true
	case "AudioVariableType":
		return OpcuaNodeIdServicesVariableType_AudioVariableType, true
	case "CartesianCoordinatesType":
		return OpcuaNodeIdServicesVariableType_CartesianCoordinatesType, true
	case "ThreeDCartesianCoordinatesType":
		return OpcuaNodeIdServicesVariableType_ThreeDCartesianCoordinatesType, true
	case "OrientationType":
		return OpcuaNodeIdServicesVariableType_OrientationType, true
	case "ThreeDOrientationType":
		return OpcuaNodeIdServicesVariableType_ThreeDOrientationType, true
	case "FrameType":
		return OpcuaNodeIdServicesVariableType_FrameType, true
	case "ThreeDFrameType":
		return OpcuaNodeIdServicesVariableType_ThreeDFrameType, true
	case "MultiStateDictionaryEntryDiscreteBaseType":
		return OpcuaNodeIdServicesVariableType_MultiStateDictionaryEntryDiscreteBaseType, true
	case "MultiStateDictionaryEntryDiscreteType":
		return OpcuaNodeIdServicesVariableType_MultiStateDictionaryEntryDiscreteType, true
	case "PubSubDiagnosticsCounterType":
		return OpcuaNodeIdServicesVariableType_PubSubDiagnosticsCounterType, true
	case "ServerVendorCapabilityType":
		return OpcuaNodeIdServicesVariableType_ServerVendorCapabilityType, true
	case "ServerStatusType":
		return OpcuaNodeIdServicesVariableType_ServerStatusType, true
	case "ServerDiagnosticsSummaryType":
		return OpcuaNodeIdServicesVariableType_ServerDiagnosticsSummaryType, true
	case "SamplingIntervalDiagnosticsArrayType":
		return OpcuaNodeIdServicesVariableType_SamplingIntervalDiagnosticsArrayType, true
	case "SamplingIntervalDiagnosticsType":
		return OpcuaNodeIdServicesVariableType_SamplingIntervalDiagnosticsType, true
	case "SubscriptionDiagnosticsArrayType":
		return OpcuaNodeIdServicesVariableType_SubscriptionDiagnosticsArrayType, true
	case "SubscriptionDiagnosticsType":
		return OpcuaNodeIdServicesVariableType_SubscriptionDiagnosticsType, true
	case "SessionDiagnosticsArrayType":
		return OpcuaNodeIdServicesVariableType_SessionDiagnosticsArrayType, true
	case "SessionDiagnosticsVariableType":
		return OpcuaNodeIdServicesVariableType_SessionDiagnosticsVariableType, true
	case "SessionSecurityDiagnosticsArrayType":
		return OpcuaNodeIdServicesVariableType_SessionSecurityDiagnosticsArrayType, true
	case "SessionSecurityDiagnosticsType":
		return OpcuaNodeIdServicesVariableType_SessionSecurityDiagnosticsType, true
	case "DataItemType":
		return OpcuaNodeIdServicesVariableType_DataItemType, true
	case "AnalogItemType":
		return OpcuaNodeIdServicesVariableType_AnalogItemType, true
	case "DiscreteItemType":
		return OpcuaNodeIdServicesVariableType_DiscreteItemType, true
	case "TwoStateDiscreteType":
		return OpcuaNodeIdServicesVariableType_TwoStateDiscreteType, true
	case "MultiStateDiscreteType":
		return OpcuaNodeIdServicesVariableType_MultiStateDiscreteType, true
	case "ProgramDiagnosticType":
		return OpcuaNodeIdServicesVariableType_ProgramDiagnosticType, true
	case "StateVariableType":
		return OpcuaNodeIdServicesVariableType_StateVariableType, true
	case "FiniteStateVariableType":
		return OpcuaNodeIdServicesVariableType_FiniteStateVariableType, true
	case "TransitionVariableType":
		return OpcuaNodeIdServicesVariableType_TransitionVariableType, true
	case "FiniteTransitionVariableType":
		return OpcuaNodeIdServicesVariableType_FiniteTransitionVariableType, true
	case "BuildInfoType":
		return OpcuaNodeIdServicesVariableType_BuildInfoType, true
	case "AlarmStateVariableType":
		return OpcuaNodeIdServicesVariableType_AlarmStateVariableType, true
	case "BitFieldType":
		return OpcuaNodeIdServicesVariableType_BitFieldType, true
	case "ReferenceDescriptionVariableType":
		return OpcuaNodeIdServicesVariableType_ReferenceDescriptionVariableType, true
	case "BaseVariableType":
		return OpcuaNodeIdServicesVariableType_BaseVariableType, true
	case "BaseDataVariableType":
		return OpcuaNodeIdServicesVariableType_BaseDataVariableType, true
	case "PropertyType":
		return OpcuaNodeIdServicesVariableType_PropertyType, true
	case "DataTypeDescriptionType":
		return OpcuaNodeIdServicesVariableType_DataTypeDescriptionType, true
	case "DataTypeDictionaryType":
		return OpcuaNodeIdServicesVariableType_DataTypeDictionaryType, true
	case "TwoStateVariableType":
		return OpcuaNodeIdServicesVariableType_TwoStateVariableType, true
	case "ConditionVariableType":
		return OpcuaNodeIdServicesVariableType_ConditionVariableType, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableTypeKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableTypeValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableType(structType any) OpcuaNodeIdServicesVariableType {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableType {
		if sOpcuaNodeIdServicesVariableType, ok := typ.(OpcuaNodeIdServicesVariableType); ok {
			return sOpcuaNodeIdServicesVariableType
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableType) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableTypeParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableType, error) {
	return OpcuaNodeIdServicesVariableTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableType", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableType")
	}
	if enum, ok := OpcuaNodeIdServicesVariableTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableType")
		return OpcuaNodeIdServicesVariableType(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableType", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableType) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableType) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableType_MultiStateValueDiscreteType:
		return "MultiStateValueDiscreteType"
	case OpcuaNodeIdServicesVariableType_OptionSetType:
		return "OptionSetType"
	case OpcuaNodeIdServicesVariableType_ArrayItemType:
		return "ArrayItemType"
	case OpcuaNodeIdServicesVariableType_YArrayItemType:
		return "YArrayItemType"
	case OpcuaNodeIdServicesVariableType_XYArrayItemType:
		return "XYArrayItemType"
	case OpcuaNodeIdServicesVariableType_ImageItemType:
		return "ImageItemType"
	case OpcuaNodeIdServicesVariableType_CubeItemType:
		return "CubeItemType"
	case OpcuaNodeIdServicesVariableType_NDimensionArrayItemType:
		return "NDimensionArrayItemType"
	case OpcuaNodeIdServicesVariableType_GuardVariableType:
		return "GuardVariableType"
	case OpcuaNodeIdServicesVariableType_ExpressionGuardVariableType:
		return "ExpressionGuardVariableType"
	case OpcuaNodeIdServicesVariableType_ElseGuardVariableType:
		return "ElseGuardVariableType"
	case OpcuaNodeIdServicesVariableType_BaseAnalogType:
		return "BaseAnalogType"
	case OpcuaNodeIdServicesVariableType_ProgramDiagnostic2Type:
		return "ProgramDiagnostic2Type"
	case OpcuaNodeIdServicesVariableType_SelectionListType:
		return "SelectionListType"
	case OpcuaNodeIdServicesVariableType_AlarmRateVariableType:
		return "AlarmRateVariableType"
	case OpcuaNodeIdServicesVariableType_AnalogUnitType:
		return "AnalogUnitType"
	case OpcuaNodeIdServicesVariableType_AnalogUnitRangeType:
		return "AnalogUnitRangeType"
	case OpcuaNodeIdServicesVariableType_RationalNumberType:
		return "RationalNumberType"
	case OpcuaNodeIdServicesVariableType_VectorType:
		return "VectorType"
	case OpcuaNodeIdServicesVariableType_ThreeDVectorType:
		return "ThreeDVectorType"
	case OpcuaNodeIdServicesVariableType_AudioVariableType:
		return "AudioVariableType"
	case OpcuaNodeIdServicesVariableType_CartesianCoordinatesType:
		return "CartesianCoordinatesType"
	case OpcuaNodeIdServicesVariableType_ThreeDCartesianCoordinatesType:
		return "ThreeDCartesianCoordinatesType"
	case OpcuaNodeIdServicesVariableType_OrientationType:
		return "OrientationType"
	case OpcuaNodeIdServicesVariableType_ThreeDOrientationType:
		return "ThreeDOrientationType"
	case OpcuaNodeIdServicesVariableType_FrameType:
		return "FrameType"
	case OpcuaNodeIdServicesVariableType_ThreeDFrameType:
		return "ThreeDFrameType"
	case OpcuaNodeIdServicesVariableType_MultiStateDictionaryEntryDiscreteBaseType:
		return "MultiStateDictionaryEntryDiscreteBaseType"
	case OpcuaNodeIdServicesVariableType_MultiStateDictionaryEntryDiscreteType:
		return "MultiStateDictionaryEntryDiscreteType"
	case OpcuaNodeIdServicesVariableType_PubSubDiagnosticsCounterType:
		return "PubSubDiagnosticsCounterType"
	case OpcuaNodeIdServicesVariableType_ServerVendorCapabilityType:
		return "ServerVendorCapabilityType"
	case OpcuaNodeIdServicesVariableType_ServerStatusType:
		return "ServerStatusType"
	case OpcuaNodeIdServicesVariableType_ServerDiagnosticsSummaryType:
		return "ServerDiagnosticsSummaryType"
	case OpcuaNodeIdServicesVariableType_SamplingIntervalDiagnosticsArrayType:
		return "SamplingIntervalDiagnosticsArrayType"
	case OpcuaNodeIdServicesVariableType_SamplingIntervalDiagnosticsType:
		return "SamplingIntervalDiagnosticsType"
	case OpcuaNodeIdServicesVariableType_SubscriptionDiagnosticsArrayType:
		return "SubscriptionDiagnosticsArrayType"
	case OpcuaNodeIdServicesVariableType_SubscriptionDiagnosticsType:
		return "SubscriptionDiagnosticsType"
	case OpcuaNodeIdServicesVariableType_SessionDiagnosticsArrayType:
		return "SessionDiagnosticsArrayType"
	case OpcuaNodeIdServicesVariableType_SessionDiagnosticsVariableType:
		return "SessionDiagnosticsVariableType"
	case OpcuaNodeIdServicesVariableType_SessionSecurityDiagnosticsArrayType:
		return "SessionSecurityDiagnosticsArrayType"
	case OpcuaNodeIdServicesVariableType_SessionSecurityDiagnosticsType:
		return "SessionSecurityDiagnosticsType"
	case OpcuaNodeIdServicesVariableType_DataItemType:
		return "DataItemType"
	case OpcuaNodeIdServicesVariableType_AnalogItemType:
		return "AnalogItemType"
	case OpcuaNodeIdServicesVariableType_DiscreteItemType:
		return "DiscreteItemType"
	case OpcuaNodeIdServicesVariableType_TwoStateDiscreteType:
		return "TwoStateDiscreteType"
	case OpcuaNodeIdServicesVariableType_MultiStateDiscreteType:
		return "MultiStateDiscreteType"
	case OpcuaNodeIdServicesVariableType_ProgramDiagnosticType:
		return "ProgramDiagnosticType"
	case OpcuaNodeIdServicesVariableType_StateVariableType:
		return "StateVariableType"
	case OpcuaNodeIdServicesVariableType_FiniteStateVariableType:
		return "FiniteStateVariableType"
	case OpcuaNodeIdServicesVariableType_TransitionVariableType:
		return "TransitionVariableType"
	case OpcuaNodeIdServicesVariableType_FiniteTransitionVariableType:
		return "FiniteTransitionVariableType"
	case OpcuaNodeIdServicesVariableType_BuildInfoType:
		return "BuildInfoType"
	case OpcuaNodeIdServicesVariableType_AlarmStateVariableType:
		return "AlarmStateVariableType"
	case OpcuaNodeIdServicesVariableType_BitFieldType:
		return "BitFieldType"
	case OpcuaNodeIdServicesVariableType_ReferenceDescriptionVariableType:
		return "ReferenceDescriptionVariableType"
	case OpcuaNodeIdServicesVariableType_BaseVariableType:
		return "BaseVariableType"
	case OpcuaNodeIdServicesVariableType_BaseDataVariableType:
		return "BaseDataVariableType"
	case OpcuaNodeIdServicesVariableType_PropertyType:
		return "PropertyType"
	case OpcuaNodeIdServicesVariableType_DataTypeDescriptionType:
		return "DataTypeDescriptionType"
	case OpcuaNodeIdServicesVariableType_DataTypeDictionaryType:
		return "DataTypeDictionaryType"
	case OpcuaNodeIdServicesVariableType_TwoStateVariableType:
		return "TwoStateVariableType"
	case OpcuaNodeIdServicesVariableType_ConditionVariableType:
		return "ConditionVariableType"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableType) String() string {
	return e.PLC4XEnumName()
}