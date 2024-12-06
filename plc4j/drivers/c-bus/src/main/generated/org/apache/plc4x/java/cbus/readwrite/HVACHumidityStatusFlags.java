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

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class HVACHumidityStatusFlags implements Message {

  // Properties.
  protected final boolean expansion;
  protected final boolean error;
  protected final boolean busy;
  protected final boolean damperState;
  protected final boolean fanActive;
  protected final boolean dehumidifyingPlant;
  protected final boolean humidifyingPlant;

  // Reserved Fields
  private Boolean reservedField0;

  public HVACHumidityStatusFlags(
      boolean expansion,
      boolean error,
      boolean busy,
      boolean damperState,
      boolean fanActive,
      boolean dehumidifyingPlant,
      boolean humidifyingPlant) {
    super();
    this.expansion = expansion;
    this.error = error;
    this.busy = busy;
    this.damperState = damperState;
    this.fanActive = fanActive;
    this.dehumidifyingPlant = dehumidifyingPlant;
    this.humidifyingPlant = humidifyingPlant;
  }

  public boolean getExpansion() {
    return expansion;
  }

  public boolean getError() {
    return error;
  }

  public boolean getBusy() {
    return busy;
  }

  public boolean getDamperState() {
    return damperState;
  }

  public boolean getFanActive() {
    return fanActive;
  }

  public boolean getDehumidifyingPlant() {
    return dehumidifyingPlant;
  }

  public boolean getHumidifyingPlant() {
    return humidifyingPlant;
  }

  public boolean getIsDamperStateClosed() {
    return (boolean) (!(getDamperState()));
  }

  public boolean getIsDamperStateOpen() {
    return (boolean) (getDamperState());
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("HVACHumidityStatusFlags");

    // Simple Field (expansion)
    writeSimpleField("expansion", expansion, writeBoolean(writeBuffer));

    // Simple Field (error)
    writeSimpleField("error", error, writeBoolean(writeBuffer));

    // Simple Field (busy)
    writeSimpleField("busy", busy, writeBoolean(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (boolean) false,
        writeBoolean(writeBuffer));

    // Simple Field (damperState)
    writeSimpleField("damperState", damperState, writeBoolean(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isDamperStateClosed = getIsDamperStateClosed();
    writeBuffer.writeVirtual("isDamperStateClosed", isDamperStateClosed);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isDamperStateOpen = getIsDamperStateOpen();
    writeBuffer.writeVirtual("isDamperStateOpen", isDamperStateOpen);

    // Simple Field (fanActive)
    writeSimpleField("fanActive", fanActive, writeBoolean(writeBuffer));

    // Simple Field (dehumidifyingPlant)
    writeSimpleField("dehumidifyingPlant", dehumidifyingPlant, writeBoolean(writeBuffer));

    // Simple Field (humidifyingPlant)
    writeSimpleField("humidifyingPlant", humidifyingPlant, writeBoolean(writeBuffer));

    writeBuffer.popContext("HVACHumidityStatusFlags");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    HVACHumidityStatusFlags _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (expansion)
    lengthInBits += 1;

    // Simple field (error)
    lengthInBits += 1;

    // Simple field (busy)
    lengthInBits += 1;

    // Reserved Field (reserved)
    lengthInBits += 1;

    // Simple field (damperState)
    lengthInBits += 1;

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // Simple field (fanActive)
    lengthInBits += 1;

    // Simple field (dehumidifyingPlant)
    lengthInBits += 1;

    // Simple field (humidifyingPlant)
    lengthInBits += 1;

    return lengthInBits;
  }

  public static HVACHumidityStatusFlags staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("HVACHumidityStatusFlags");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    boolean expansion = readSimpleField("expansion", readBoolean(readBuffer));

    boolean error = readSimpleField("error", readBoolean(readBuffer));

    boolean busy = readSimpleField("busy", readBoolean(readBuffer));

    Boolean reservedField0 =
        readReservedField("reserved", readBoolean(readBuffer), (boolean) false);

    boolean damperState = readSimpleField("damperState", readBoolean(readBuffer));
    boolean isDamperStateClosed =
        readVirtualField("isDamperStateClosed", boolean.class, !(damperState));
    boolean isDamperStateOpen = readVirtualField("isDamperStateOpen", boolean.class, damperState);

    boolean fanActive = readSimpleField("fanActive", readBoolean(readBuffer));

    boolean dehumidifyingPlant = readSimpleField("dehumidifyingPlant", readBoolean(readBuffer));

    boolean humidifyingPlant = readSimpleField("humidifyingPlant", readBoolean(readBuffer));

    readBuffer.closeContext("HVACHumidityStatusFlags");
    // Create the instance
    HVACHumidityStatusFlags _hVACHumidityStatusFlags;
    _hVACHumidityStatusFlags =
        new HVACHumidityStatusFlags(
            expansion, error, busy, damperState, fanActive, dehumidifyingPlant, humidifyingPlant);
    _hVACHumidityStatusFlags.reservedField0 = reservedField0;
    return _hVACHumidityStatusFlags;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof HVACHumidityStatusFlags)) {
      return false;
    }
    HVACHumidityStatusFlags that = (HVACHumidityStatusFlags) o;
    return (getExpansion() == that.getExpansion())
        && (getError() == that.getError())
        && (getBusy() == that.getBusy())
        && (getDamperState() == that.getDamperState())
        && (getFanActive() == that.getFanActive())
        && (getDehumidifyingPlant() == that.getDehumidifyingPlant())
        && (getHumidifyingPlant() == that.getHumidifyingPlant())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getExpansion(),
        getError(),
        getBusy(),
        getDamperState(),
        getFanActive(),
        getDehumidifyingPlant(),
        getHumidifyingPlant());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
