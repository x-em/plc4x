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

// PermissionType is an enum
type PermissionType uint32

type IPermissionType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	PermissionType_permissionTypeNone                 PermissionType = 0
	PermissionType_permissionTypeBrowse               PermissionType = 1
	PermissionType_permissionTypeReadRolePermissions  PermissionType = 2
	PermissionType_permissionTypeWriteAttribute       PermissionType = 4
	PermissionType_permissionTypeWriteRolePermissions PermissionType = 8
	PermissionType_permissionTypeWriteHistorizing     PermissionType = 16
	PermissionType_permissionTypeRead                 PermissionType = 32
	PermissionType_permissionTypeWrite                PermissionType = 64
	PermissionType_permissionTypeReadHistory          PermissionType = 128
	PermissionType_permissionTypeInsertHistory        PermissionType = 256
	PermissionType_permissionTypeModifyHistory        PermissionType = 512
	PermissionType_permissionTypeDeleteHistory        PermissionType = 1024
	PermissionType_permissionTypeReceiveEvents        PermissionType = 2048
	PermissionType_permissionTypeCall                 PermissionType = 4096
	PermissionType_permissionTypeAddReference         PermissionType = 8192
	PermissionType_permissionTypeRemoveReference      PermissionType = 16384
	PermissionType_permissionTypeDeleteNode           PermissionType = 32768
	PermissionType_permissionTypeAddNode              PermissionType = 65536
)

var PermissionTypeValues []PermissionType

func init() {
	_ = errors.New
	PermissionTypeValues = []PermissionType{
		PermissionType_permissionTypeNone,
		PermissionType_permissionTypeBrowse,
		PermissionType_permissionTypeReadRolePermissions,
		PermissionType_permissionTypeWriteAttribute,
		PermissionType_permissionTypeWriteRolePermissions,
		PermissionType_permissionTypeWriteHistorizing,
		PermissionType_permissionTypeRead,
		PermissionType_permissionTypeWrite,
		PermissionType_permissionTypeReadHistory,
		PermissionType_permissionTypeInsertHistory,
		PermissionType_permissionTypeModifyHistory,
		PermissionType_permissionTypeDeleteHistory,
		PermissionType_permissionTypeReceiveEvents,
		PermissionType_permissionTypeCall,
		PermissionType_permissionTypeAddReference,
		PermissionType_permissionTypeRemoveReference,
		PermissionType_permissionTypeDeleteNode,
		PermissionType_permissionTypeAddNode,
	}
}

func PermissionTypeByValue(value uint32) (enum PermissionType, ok bool) {
	switch value {
	case 0:
		return PermissionType_permissionTypeNone, true
	case 1:
		return PermissionType_permissionTypeBrowse, true
	case 1024:
		return PermissionType_permissionTypeDeleteHistory, true
	case 128:
		return PermissionType_permissionTypeReadHistory, true
	case 16:
		return PermissionType_permissionTypeWriteHistorizing, true
	case 16384:
		return PermissionType_permissionTypeRemoveReference, true
	case 2:
		return PermissionType_permissionTypeReadRolePermissions, true
	case 2048:
		return PermissionType_permissionTypeReceiveEvents, true
	case 256:
		return PermissionType_permissionTypeInsertHistory, true
	case 32:
		return PermissionType_permissionTypeRead, true
	case 32768:
		return PermissionType_permissionTypeDeleteNode, true
	case 4:
		return PermissionType_permissionTypeWriteAttribute, true
	case 4096:
		return PermissionType_permissionTypeCall, true
	case 512:
		return PermissionType_permissionTypeModifyHistory, true
	case 64:
		return PermissionType_permissionTypeWrite, true
	case 65536:
		return PermissionType_permissionTypeAddNode, true
	case 8:
		return PermissionType_permissionTypeWriteRolePermissions, true
	case 8192:
		return PermissionType_permissionTypeAddReference, true
	}
	return 0, false
}

func PermissionTypeByName(value string) (enum PermissionType, ok bool) {
	switch value {
	case "permissionTypeNone":
		return PermissionType_permissionTypeNone, true
	case "permissionTypeBrowse":
		return PermissionType_permissionTypeBrowse, true
	case "permissionTypeDeleteHistory":
		return PermissionType_permissionTypeDeleteHistory, true
	case "permissionTypeReadHistory":
		return PermissionType_permissionTypeReadHistory, true
	case "permissionTypeWriteHistorizing":
		return PermissionType_permissionTypeWriteHistorizing, true
	case "permissionTypeRemoveReference":
		return PermissionType_permissionTypeRemoveReference, true
	case "permissionTypeReadRolePermissions":
		return PermissionType_permissionTypeReadRolePermissions, true
	case "permissionTypeReceiveEvents":
		return PermissionType_permissionTypeReceiveEvents, true
	case "permissionTypeInsertHistory":
		return PermissionType_permissionTypeInsertHistory, true
	case "permissionTypeRead":
		return PermissionType_permissionTypeRead, true
	case "permissionTypeDeleteNode":
		return PermissionType_permissionTypeDeleteNode, true
	case "permissionTypeWriteAttribute":
		return PermissionType_permissionTypeWriteAttribute, true
	case "permissionTypeCall":
		return PermissionType_permissionTypeCall, true
	case "permissionTypeModifyHistory":
		return PermissionType_permissionTypeModifyHistory, true
	case "permissionTypeWrite":
		return PermissionType_permissionTypeWrite, true
	case "permissionTypeAddNode":
		return PermissionType_permissionTypeAddNode, true
	case "permissionTypeWriteRolePermissions":
		return PermissionType_permissionTypeWriteRolePermissions, true
	case "permissionTypeAddReference":
		return PermissionType_permissionTypeAddReference, true
	}
	return 0, false
}

func PermissionTypeKnows(value uint32) bool {
	for _, typeValue := range PermissionTypeValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastPermissionType(structType any) PermissionType {
	castFunc := func(typ any) PermissionType {
		if sPermissionType, ok := typ.(PermissionType); ok {
			return sPermissionType
		}
		return 0
	}
	return castFunc(structType)
}

func (m PermissionType) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m PermissionType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PermissionTypeParse(ctx context.Context, theBytes []byte) (PermissionType, error) {
	return PermissionTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func PermissionTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (PermissionType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("PermissionType", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading PermissionType")
	}
	if enum, ok := PermissionTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for PermissionType")
		return PermissionType(val), nil
	} else {
		return enum, nil
	}
}

func (e PermissionType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e PermissionType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint32("PermissionType", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e PermissionType) GetValue() uint32 {
	return uint32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e PermissionType) PLC4XEnumName() string {
	switch e {
	case PermissionType_permissionTypeNone:
		return "permissionTypeNone"
	case PermissionType_permissionTypeBrowse:
		return "permissionTypeBrowse"
	case PermissionType_permissionTypeDeleteHistory:
		return "permissionTypeDeleteHistory"
	case PermissionType_permissionTypeReadHistory:
		return "permissionTypeReadHistory"
	case PermissionType_permissionTypeWriteHistorizing:
		return "permissionTypeWriteHistorizing"
	case PermissionType_permissionTypeRemoveReference:
		return "permissionTypeRemoveReference"
	case PermissionType_permissionTypeReadRolePermissions:
		return "permissionTypeReadRolePermissions"
	case PermissionType_permissionTypeReceiveEvents:
		return "permissionTypeReceiveEvents"
	case PermissionType_permissionTypeInsertHistory:
		return "permissionTypeInsertHistory"
	case PermissionType_permissionTypeRead:
		return "permissionTypeRead"
	case PermissionType_permissionTypeDeleteNode:
		return "permissionTypeDeleteNode"
	case PermissionType_permissionTypeWriteAttribute:
		return "permissionTypeWriteAttribute"
	case PermissionType_permissionTypeCall:
		return "permissionTypeCall"
	case PermissionType_permissionTypeModifyHistory:
		return "permissionTypeModifyHistory"
	case PermissionType_permissionTypeWrite:
		return "permissionTypeWrite"
	case PermissionType_permissionTypeAddNode:
		return "permissionTypeAddNode"
	case PermissionType_permissionTypeWriteRolePermissions:
		return "permissionTypeWriteRolePermissions"
	case PermissionType_permissionTypeAddReference:
		return "permissionTypeAddReference"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e PermissionType) String() string {
	return e.PLC4XEnumName()
}