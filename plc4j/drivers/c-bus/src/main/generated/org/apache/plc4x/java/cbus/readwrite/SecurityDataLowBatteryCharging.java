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

public class SecurityDataLowBatteryCharging extends SecurityData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final byte startStop;

  public SecurityDataLowBatteryCharging(
      SecurityCommandTypeContainer commandTypeContainer, byte argument, byte startStop) {
    super(commandTypeContainer, argument);
    this.startStop = startStop;
  }

  public byte getStartStop() {
    return startStop;
  }

  public boolean getChargeStopped() {
    return (boolean) ((getStartStop()) == (0x00));
  }

  public boolean getChargeStarted() {
    return (boolean) ((getStartStop()) > (0xFE));
  }

  @Override
  protected void serializeSecurityDataChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SecurityDataLowBatteryCharging");

    // Simple Field (startStop)
    writeSimpleField("startStop", startStop, writeByte(writeBuffer, 8));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean chargeStopped = getChargeStopped();
    writeBuffer.writeVirtual("chargeStopped", chargeStopped);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean chargeStarted = getChargeStarted();
    writeBuffer.writeVirtual("chargeStarted", chargeStarted);

    writeBuffer.popContext("SecurityDataLowBatteryCharging");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SecurityDataLowBatteryCharging _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (startStop)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static SecurityDataBuilder staticParseSecurityDataBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("SecurityDataLowBatteryCharging");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte startStop = readSimpleField("startStop", readByte(readBuffer, 8));
    boolean chargeStopped = readVirtualField("chargeStopped", boolean.class, (startStop) == (0x00));
    boolean chargeStarted = readVirtualField("chargeStarted", boolean.class, (startStop) > (0xFE));

    readBuffer.closeContext("SecurityDataLowBatteryCharging");
    // Create the instance
    return new SecurityDataLowBatteryChargingBuilderImpl(startStop);
  }

  public static class SecurityDataLowBatteryChargingBuilderImpl
      implements SecurityData.SecurityDataBuilder {
    private final byte startStop;

    public SecurityDataLowBatteryChargingBuilderImpl(byte startStop) {
      this.startStop = startStop;
    }

    public SecurityDataLowBatteryCharging build(
        SecurityCommandTypeContainer commandTypeContainer, byte argument) {
      SecurityDataLowBatteryCharging securityDataLowBatteryCharging =
          new SecurityDataLowBatteryCharging(commandTypeContainer, argument, startStop);
      return securityDataLowBatteryCharging;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SecurityDataLowBatteryCharging)) {
      return false;
    }
    SecurityDataLowBatteryCharging that = (SecurityDataLowBatteryCharging) o;
    return (getStartStop() == that.getStartStop()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getStartStop());
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