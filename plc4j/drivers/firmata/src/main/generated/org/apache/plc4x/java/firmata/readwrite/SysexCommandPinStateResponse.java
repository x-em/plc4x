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
package org.apache.plc4x.java.firmata.readwrite;

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

public class SysexCommandPinStateResponse extends SysexCommand implements Message {

  // Accessors for discriminator values.
  public Short getCommandType() {
    return (short) 0x6E;
  }

  public Boolean getResponse() {
    return false;
  }

  // Properties.
  protected final short pin;
  protected final short pinMode;
  protected final short pinState;

  public SysexCommandPinStateResponse(short pin, short pinMode, short pinState) {
    super();
    this.pin = pin;
    this.pinMode = pinMode;
    this.pinState = pinState;
  }

  public short getPin() {
    return pin;
  }

  public short getPinMode() {
    return pinMode;
  }

  public short getPinState() {
    return pinState;
  }

  @Override
  protected void serializeSysexCommandChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("SysexCommandPinStateResponse");

    // Simple Field (pin)
    writeSimpleField("pin", pin, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (pinMode)
    writeSimpleField("pinMode", pinMode, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (pinState)
    writeSimpleField("pinState", pinState, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("SysexCommandPinStateResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SysexCommandPinStateResponse _value = this;

    // Simple field (pin)
    lengthInBits += 8;

    // Simple field (pinMode)
    lengthInBits += 8;

    // Simple field (pinState)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static SysexCommandBuilder staticParseSysexCommandBuilder(
      ReadBuffer readBuffer, Boolean response) throws ParseException {
    readBuffer.pullContext("SysexCommandPinStateResponse");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    short pin = readSimpleField("pin", readUnsignedShort(readBuffer, 8));

    short pinMode = readSimpleField("pinMode", readUnsignedShort(readBuffer, 8));

    short pinState = readSimpleField("pinState", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("SysexCommandPinStateResponse");
    // Create the instance
    return new SysexCommandPinStateResponseBuilderImpl(pin, pinMode, pinState);
  }

  public static class SysexCommandPinStateResponseBuilderImpl
      implements SysexCommand.SysexCommandBuilder {
    private final short pin;
    private final short pinMode;
    private final short pinState;

    public SysexCommandPinStateResponseBuilderImpl(short pin, short pinMode, short pinState) {
      this.pin = pin;
      this.pinMode = pinMode;
      this.pinState = pinState;
    }

    public SysexCommandPinStateResponse build() {
      SysexCommandPinStateResponse sysexCommandPinStateResponse =
          new SysexCommandPinStateResponse(pin, pinMode, pinState);
      return sysexCommandPinStateResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SysexCommandPinStateResponse)) {
      return false;
    }
    SysexCommandPinStateResponse that = (SysexCommandPinStateResponse) o;
    return (getPin() == that.getPin())
        && (getPinMode() == that.getPinMode())
        && (getPinState() == that.getPinState())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getPin(), getPinMode(), getPinState());
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
