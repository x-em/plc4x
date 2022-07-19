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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// MeasurementUnits is an enum
type MeasurementUnits uint8

type IMeasurementUnits interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	MeasurementUnits_CELSIUS              MeasurementUnits = 0x00
	MeasurementUnits_AMPS                 MeasurementUnits = 0x01
	MeasurementUnits_ANGLE_DEGREES        MeasurementUnits = 0x02
	MeasurementUnits_COULOMB              MeasurementUnits = 0x03
	MeasurementUnits_BOOLEANLOGIC         MeasurementUnits = 0x04
	MeasurementUnits_FARADS               MeasurementUnits = 0x05
	MeasurementUnits_HENRYS               MeasurementUnits = 0x06
	MeasurementUnits_HERTZ                MeasurementUnits = 0x07
	MeasurementUnits_JOULES               MeasurementUnits = 0x08
	MeasurementUnits_KATAL                MeasurementUnits = 0x09
	MeasurementUnits_KG_PER_M3            MeasurementUnits = 0x0A
	MeasurementUnits_KILOGRAMS            MeasurementUnits = 0x0B
	MeasurementUnits_LITRES               MeasurementUnits = 0x0C
	MeasurementUnits_LITRES_PER_HOUR      MeasurementUnits = 0x0D
	MeasurementUnits_LITRES_PER_MINUTE    MeasurementUnits = 0x0E
	MeasurementUnits_LITRES_PER_SECOND    MeasurementUnits = 0x0F
	MeasurementUnits_LUX                  MeasurementUnits = 0x10
	MeasurementUnits_METRES               MeasurementUnits = 0x11
	MeasurementUnits_METRES_PER_MINUTE    MeasurementUnits = 0x12
	MeasurementUnits_METRES_PER_SECOND    MeasurementUnits = 0x13
	MeasurementUnits_METRES_PER_S_SQUARED MeasurementUnits = 0x14
	MeasurementUnits_MOLE                 MeasurementUnits = 0x15
	MeasurementUnits_NEWTON_METRE         MeasurementUnits = 0x16
	MeasurementUnits_NEWTONS              MeasurementUnits = 0x17
	MeasurementUnits_OHMS                 MeasurementUnits = 0x18
	MeasurementUnits_PASCAL               MeasurementUnits = 0x19
	MeasurementUnits_PERCENT              MeasurementUnits = 0x1A
	MeasurementUnits_DECIBELS             MeasurementUnits = 0x1B
	MeasurementUnits_PPM                  MeasurementUnits = 0x1C
	MeasurementUnits_RPM                  MeasurementUnits = 0x1D
	MeasurementUnits_SECOND               MeasurementUnits = 0x1E
	MeasurementUnits_MINUTES              MeasurementUnits = 0x1F
	MeasurementUnits_HOURS                MeasurementUnits = 0x20
	MeasurementUnits_SIEVERTS             MeasurementUnits = 0x21
	MeasurementUnits_STERADIAN            MeasurementUnits = 0x22
	MeasurementUnits_TESLA                MeasurementUnits = 0x23
	MeasurementUnits_VOLTS                MeasurementUnits = 0x24
	MeasurementUnits_WATT_HOURS           MeasurementUnits = 0x25
	MeasurementUnits_WATTS                MeasurementUnits = 0x26
	MeasurementUnits_WEBERS               MeasurementUnits = 0x27
	MeasurementUnits_NO_UNITS             MeasurementUnits = 0xFE
	MeasurementUnits_CUSTOM               MeasurementUnits = 0xFF
)

var MeasurementUnitsValues []MeasurementUnits

func init() {
	_ = errors.New
	MeasurementUnitsValues = []MeasurementUnits{
		MeasurementUnits_CELSIUS,
		MeasurementUnits_AMPS,
		MeasurementUnits_ANGLE_DEGREES,
		MeasurementUnits_COULOMB,
		MeasurementUnits_BOOLEANLOGIC,
		MeasurementUnits_FARADS,
		MeasurementUnits_HENRYS,
		MeasurementUnits_HERTZ,
		MeasurementUnits_JOULES,
		MeasurementUnits_KATAL,
		MeasurementUnits_KG_PER_M3,
		MeasurementUnits_KILOGRAMS,
		MeasurementUnits_LITRES,
		MeasurementUnits_LITRES_PER_HOUR,
		MeasurementUnits_LITRES_PER_MINUTE,
		MeasurementUnits_LITRES_PER_SECOND,
		MeasurementUnits_LUX,
		MeasurementUnits_METRES,
		MeasurementUnits_METRES_PER_MINUTE,
		MeasurementUnits_METRES_PER_SECOND,
		MeasurementUnits_METRES_PER_S_SQUARED,
		MeasurementUnits_MOLE,
		MeasurementUnits_NEWTON_METRE,
		MeasurementUnits_NEWTONS,
		MeasurementUnits_OHMS,
		MeasurementUnits_PASCAL,
		MeasurementUnits_PERCENT,
		MeasurementUnits_DECIBELS,
		MeasurementUnits_PPM,
		MeasurementUnits_RPM,
		MeasurementUnits_SECOND,
		MeasurementUnits_MINUTES,
		MeasurementUnits_HOURS,
		MeasurementUnits_SIEVERTS,
		MeasurementUnits_STERADIAN,
		MeasurementUnits_TESLA,
		MeasurementUnits_VOLTS,
		MeasurementUnits_WATT_HOURS,
		MeasurementUnits_WATTS,
		MeasurementUnits_WEBERS,
		MeasurementUnits_NO_UNITS,
		MeasurementUnits_CUSTOM,
	}
}

func MeasurementUnitsByValue(value uint8) (enum MeasurementUnits, ok bool) {
	switch value {
	case 0x00:
		return MeasurementUnits_CELSIUS, true
	case 0x01:
		return MeasurementUnits_AMPS, true
	case 0x02:
		return MeasurementUnits_ANGLE_DEGREES, true
	case 0x03:
		return MeasurementUnits_COULOMB, true
	case 0x04:
		return MeasurementUnits_BOOLEANLOGIC, true
	case 0x05:
		return MeasurementUnits_FARADS, true
	case 0x06:
		return MeasurementUnits_HENRYS, true
	case 0x07:
		return MeasurementUnits_HERTZ, true
	case 0x08:
		return MeasurementUnits_JOULES, true
	case 0x09:
		return MeasurementUnits_KATAL, true
	case 0x0A:
		return MeasurementUnits_KG_PER_M3, true
	case 0x0B:
		return MeasurementUnits_KILOGRAMS, true
	case 0x0C:
		return MeasurementUnits_LITRES, true
	case 0x0D:
		return MeasurementUnits_LITRES_PER_HOUR, true
	case 0x0E:
		return MeasurementUnits_LITRES_PER_MINUTE, true
	case 0x0F:
		return MeasurementUnits_LITRES_PER_SECOND, true
	case 0x10:
		return MeasurementUnits_LUX, true
	case 0x11:
		return MeasurementUnits_METRES, true
	case 0x12:
		return MeasurementUnits_METRES_PER_MINUTE, true
	case 0x13:
		return MeasurementUnits_METRES_PER_SECOND, true
	case 0x14:
		return MeasurementUnits_METRES_PER_S_SQUARED, true
	case 0x15:
		return MeasurementUnits_MOLE, true
	case 0x16:
		return MeasurementUnits_NEWTON_METRE, true
	case 0x17:
		return MeasurementUnits_NEWTONS, true
	case 0x18:
		return MeasurementUnits_OHMS, true
	case 0x19:
		return MeasurementUnits_PASCAL, true
	case 0x1A:
		return MeasurementUnits_PERCENT, true
	case 0x1B:
		return MeasurementUnits_DECIBELS, true
	case 0x1C:
		return MeasurementUnits_PPM, true
	case 0x1D:
		return MeasurementUnits_RPM, true
	case 0x1E:
		return MeasurementUnits_SECOND, true
	case 0x1F:
		return MeasurementUnits_MINUTES, true
	case 0x20:
		return MeasurementUnits_HOURS, true
	case 0x21:
		return MeasurementUnits_SIEVERTS, true
	case 0x22:
		return MeasurementUnits_STERADIAN, true
	case 0x23:
		return MeasurementUnits_TESLA, true
	case 0x24:
		return MeasurementUnits_VOLTS, true
	case 0x25:
		return MeasurementUnits_WATT_HOURS, true
	case 0x26:
		return MeasurementUnits_WATTS, true
	case 0x27:
		return MeasurementUnits_WEBERS, true
	case 0xFE:
		return MeasurementUnits_NO_UNITS, true
	case 0xFF:
		return MeasurementUnits_CUSTOM, true
	}
	return 0, false
}

func MeasurementUnitsByName(value string) (enum MeasurementUnits, ok bool) {
	switch value {
	case "CELSIUS":
		return MeasurementUnits_CELSIUS, true
	case "AMPS":
		return MeasurementUnits_AMPS, true
	case "ANGLE_DEGREES":
		return MeasurementUnits_ANGLE_DEGREES, true
	case "COULOMB":
		return MeasurementUnits_COULOMB, true
	case "BOOLEANLOGIC":
		return MeasurementUnits_BOOLEANLOGIC, true
	case "FARADS":
		return MeasurementUnits_FARADS, true
	case "HENRYS":
		return MeasurementUnits_HENRYS, true
	case "HERTZ":
		return MeasurementUnits_HERTZ, true
	case "JOULES":
		return MeasurementUnits_JOULES, true
	case "KATAL":
		return MeasurementUnits_KATAL, true
	case "KG_PER_M3":
		return MeasurementUnits_KG_PER_M3, true
	case "KILOGRAMS":
		return MeasurementUnits_KILOGRAMS, true
	case "LITRES":
		return MeasurementUnits_LITRES, true
	case "LITRES_PER_HOUR":
		return MeasurementUnits_LITRES_PER_HOUR, true
	case "LITRES_PER_MINUTE":
		return MeasurementUnits_LITRES_PER_MINUTE, true
	case "LITRES_PER_SECOND":
		return MeasurementUnits_LITRES_PER_SECOND, true
	case "LUX":
		return MeasurementUnits_LUX, true
	case "METRES":
		return MeasurementUnits_METRES, true
	case "METRES_PER_MINUTE":
		return MeasurementUnits_METRES_PER_MINUTE, true
	case "METRES_PER_SECOND":
		return MeasurementUnits_METRES_PER_SECOND, true
	case "METRES_PER_S_SQUARED":
		return MeasurementUnits_METRES_PER_S_SQUARED, true
	case "MOLE":
		return MeasurementUnits_MOLE, true
	case "NEWTON_METRE":
		return MeasurementUnits_NEWTON_METRE, true
	case "NEWTONS":
		return MeasurementUnits_NEWTONS, true
	case "OHMS":
		return MeasurementUnits_OHMS, true
	case "PASCAL":
		return MeasurementUnits_PASCAL, true
	case "PERCENT":
		return MeasurementUnits_PERCENT, true
	case "DECIBELS":
		return MeasurementUnits_DECIBELS, true
	case "PPM":
		return MeasurementUnits_PPM, true
	case "RPM":
		return MeasurementUnits_RPM, true
	case "SECOND":
		return MeasurementUnits_SECOND, true
	case "MINUTES":
		return MeasurementUnits_MINUTES, true
	case "HOURS":
		return MeasurementUnits_HOURS, true
	case "SIEVERTS":
		return MeasurementUnits_SIEVERTS, true
	case "STERADIAN":
		return MeasurementUnits_STERADIAN, true
	case "TESLA":
		return MeasurementUnits_TESLA, true
	case "VOLTS":
		return MeasurementUnits_VOLTS, true
	case "WATT_HOURS":
		return MeasurementUnits_WATT_HOURS, true
	case "WATTS":
		return MeasurementUnits_WATTS, true
	case "WEBERS":
		return MeasurementUnits_WEBERS, true
	case "NO_UNITS":
		return MeasurementUnits_NO_UNITS, true
	case "CUSTOM":
		return MeasurementUnits_CUSTOM, true
	}
	return 0, false
}

func MeasurementUnitsKnows(value uint8) bool {
	for _, typeValue := range MeasurementUnitsValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastMeasurementUnits(structType interface{}) MeasurementUnits {
	castFunc := func(typ interface{}) MeasurementUnits {
		if sMeasurementUnits, ok := typ.(MeasurementUnits); ok {
			return sMeasurementUnits
		}
		return 0
	}
	return castFunc(structType)
}

func (m MeasurementUnits) GetLengthInBits() uint16 {
	return 8
}

func (m MeasurementUnits) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MeasurementUnitsParse(readBuffer utils.ReadBuffer) (MeasurementUnits, error) {
	val, err := readBuffer.ReadUint8("MeasurementUnits", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading MeasurementUnits")
	}
	if enum, ok := MeasurementUnitsByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return MeasurementUnits(val), nil
	} else {
		return enum, nil
	}
}

func (e MeasurementUnits) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("MeasurementUnits", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e MeasurementUnits) PLC4XEnumName() string {
	switch e {
	case MeasurementUnits_CELSIUS:
		return "CELSIUS"
	case MeasurementUnits_AMPS:
		return "AMPS"
	case MeasurementUnits_ANGLE_DEGREES:
		return "ANGLE_DEGREES"
	case MeasurementUnits_COULOMB:
		return "COULOMB"
	case MeasurementUnits_BOOLEANLOGIC:
		return "BOOLEANLOGIC"
	case MeasurementUnits_FARADS:
		return "FARADS"
	case MeasurementUnits_HENRYS:
		return "HENRYS"
	case MeasurementUnits_HERTZ:
		return "HERTZ"
	case MeasurementUnits_JOULES:
		return "JOULES"
	case MeasurementUnits_KATAL:
		return "KATAL"
	case MeasurementUnits_KG_PER_M3:
		return "KG_PER_M3"
	case MeasurementUnits_KILOGRAMS:
		return "KILOGRAMS"
	case MeasurementUnits_LITRES:
		return "LITRES"
	case MeasurementUnits_LITRES_PER_HOUR:
		return "LITRES_PER_HOUR"
	case MeasurementUnits_LITRES_PER_MINUTE:
		return "LITRES_PER_MINUTE"
	case MeasurementUnits_LITRES_PER_SECOND:
		return "LITRES_PER_SECOND"
	case MeasurementUnits_LUX:
		return "LUX"
	case MeasurementUnits_METRES:
		return "METRES"
	case MeasurementUnits_METRES_PER_MINUTE:
		return "METRES_PER_MINUTE"
	case MeasurementUnits_METRES_PER_SECOND:
		return "METRES_PER_SECOND"
	case MeasurementUnits_METRES_PER_S_SQUARED:
		return "METRES_PER_S_SQUARED"
	case MeasurementUnits_MOLE:
		return "MOLE"
	case MeasurementUnits_NEWTON_METRE:
		return "NEWTON_METRE"
	case MeasurementUnits_NEWTONS:
		return "NEWTONS"
	case MeasurementUnits_OHMS:
		return "OHMS"
	case MeasurementUnits_PASCAL:
		return "PASCAL"
	case MeasurementUnits_PERCENT:
		return "PERCENT"
	case MeasurementUnits_DECIBELS:
		return "DECIBELS"
	case MeasurementUnits_PPM:
		return "PPM"
	case MeasurementUnits_RPM:
		return "RPM"
	case MeasurementUnits_SECOND:
		return "SECOND"
	case MeasurementUnits_MINUTES:
		return "MINUTES"
	case MeasurementUnits_HOURS:
		return "HOURS"
	case MeasurementUnits_SIEVERTS:
		return "SIEVERTS"
	case MeasurementUnits_STERADIAN:
		return "STERADIAN"
	case MeasurementUnits_TESLA:
		return "TESLA"
	case MeasurementUnits_VOLTS:
		return "VOLTS"
	case MeasurementUnits_WATT_HOURS:
		return "WATT_HOURS"
	case MeasurementUnits_WATTS:
		return "WATTS"
	case MeasurementUnits_WEBERS:
		return "WEBERS"
	case MeasurementUnits_NO_UNITS:
		return "NO_UNITS"
	case MeasurementUnits_CUSTOM:
		return "CUSTOM"
	}
	return ""
}

func (e MeasurementUnits) String() string {
	return e.PLC4XEnumName()
}
