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

package readwrite

import (
	"context"

	"github.com/pkg/errors"

	. "github.com/apache/plc4x/plc4go/protocols/modbus/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

type ModbusParserHelper struct {
}

func (m ModbusParserHelper) Parse(typeName string, arguments []string, io utils.ReadBuffer) (any, error) {
	switch typeName {
	case "ModbusPDUWriteFileRecordRequestItem":
		return ModbusPDUWriteFileRecordRequestItemParseWithBuffer(context.Background(), io)
	case "DataItem":
		dataType, _ := ModbusDataTypeByName(arguments[0])
		numberOfValues, err := utils.StrToUint16(arguments[1])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		bigEndian, err := utils.StrToBool(arguments[2])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return DataItemParseWithBuffer(context.Background(), io, dataType, numberOfValues, bigEndian)
	case "ModbusPDUReadFileRecordResponseItem":
		return ModbusPDUReadFileRecordResponseItemParseWithBuffer(context.Background(), io)
	case "ModbusDeviceInformationObject":
		return ModbusDeviceInformationObjectParseWithBuffer(context.Background(), io)
	case "ModbusConstants":
		return ModbusConstantsParseWithBuffer(context.Background(), io)
	case "ModbusPDUWriteFileRecordResponseItem":
		return ModbusPDUWriteFileRecordResponseItemParseWithBuffer(context.Background(), io)
	case "ModbusPDU":
		response, err := utils.StrToBool(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return ModbusPDUParseWithBuffer[ModbusPDU](context.Background(), io, response)
	case "ModbusPDUReadFileRecordRequestItem":
		return ModbusPDUReadFileRecordRequestItemParseWithBuffer(context.Background(), io)
	case "ModbusADU":
		driverType, _ := DriverTypeByName(arguments[0])
		response, err := utils.StrToBool(arguments[1])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return ModbusADUParseWithBuffer[ModbusADU](context.Background(), io, driverType, response)
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
