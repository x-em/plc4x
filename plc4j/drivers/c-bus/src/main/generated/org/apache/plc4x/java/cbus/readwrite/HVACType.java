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
package org.apache.plc4x.java.cbus.readwrite;

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum HVACType {
  NONE((short) 0x00),
  FURNACE_GAS_OIL_ELECTRIC((short) 0x01),
  EVAPORATIVE((short) 0x02),
  HEAT_PUMP_REVERSE_CYCLE((short) 0x03),
  HEAT_PUMP_HEATING_ONLY((short) 0x04),
  HEAT_PUMP_COOLING_ONLY((short) 0x05),
  FURNANCE_EVAP_COOLING((short) 0x06),
  FURNANCE_HEAT_PUMP_COOLING_ONLY((short) 0x07),
  HYDRONIC((short) 0x08),
  HYDRONIC_HEAT_PUMP_COOLING_ONLY((short) 0x09),
  HYDRONIC_EVAPORATIVE((short) 0x0A),
  ANY((short) 0xFF);
  private static final Map<Short, HVACType> map;

  static {
    map = new HashMap<>();
    for (HVACType value : HVACType.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private short value;

  HVACType(short value) {
    this.value = value;
  }

  public short getValue() {
    return value;
  }

  public static HVACType enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}
