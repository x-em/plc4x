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
package org.apache.plc4x.java.bacnetip.readwrite;

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum BACnetEventType {
  CHANGE_OF_BITSTRING((int) 0),
  CHANGE_OF_STATE((int) 1),
  CHANGE_OF_VALUE((int) 2),
  COMMAND_FAILURE((int) 3),
  FLOATING_LIMIT((int) 4),
  OUT_OF_RANGE((int) 5),
  CHANGE_OF_LIFE_SAFETY((int) 8),
  EXTENDED((int) 9),
  BUFFER_READY((int) 10),
  UNSIGNED_RANGE((int) 11),
  ACCESS_EVENT((int) 13),
  DOUBLE_OUT_OF_RANGE((int) 14),
  SIGNED_OUT_OF_RANGE((int) 15),
  UNSIGNED_OUT_OF_RANGE((int) 16),
  CHANGE_OF_CHARACTERSTRING((int) 17),
  CHANGE_OF_STATUS_FLAGS((int) 18),
  CHANGE_OF_RELIABILITY((int) 19),
  NONE((int) 20),
  CHANGE_OF_DISCRETE_VALUE((int) 21),
  CHANGE_OF_TIMER((int) 22),
  VENDOR_PROPRIETARY_VALUE((int) 0xFFFF);
  private static final Map<Integer, BACnetEventType> map;

  static {
    map = new HashMap<>();
    for (BACnetEventType value : BACnetEventType.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private final int value;

  BACnetEventType(int value) {
    this.value = value;
  }

  public int getValue() {
    return value;
  }

  public static BACnetEventType enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}