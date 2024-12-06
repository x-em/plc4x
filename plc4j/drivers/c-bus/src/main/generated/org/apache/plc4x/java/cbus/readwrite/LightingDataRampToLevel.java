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

public class LightingDataRampToLevel extends LightingData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final byte group;
  protected final byte level;

  public LightingDataRampToLevel(
      LightingCommandTypeContainer commandTypeContainer, byte group, byte level) {
    super(commandTypeContainer);
    this.group = group;
    this.level = level;
  }

  public byte getGroup() {
    return group;
  }

  public byte getLevel() {
    return level;
  }

  @Override
  protected void serializeLightingDataChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("LightingDataRampToLevel");

    // Simple Field (group)
    writeSimpleField("group", group, writeByte(writeBuffer, 8));

    // Simple Field (level)
    writeSimpleField("level", level, writeByte(writeBuffer, 8));

    writeBuffer.popContext("LightingDataRampToLevel");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    LightingDataRampToLevel _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (group)
    lengthInBits += 8;

    // Simple field (level)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static LightingDataBuilder staticParseLightingDataBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("LightingDataRampToLevel");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte group = readSimpleField("group", readByte(readBuffer, 8));

    byte level = readSimpleField("level", readByte(readBuffer, 8));

    readBuffer.closeContext("LightingDataRampToLevel");
    // Create the instance
    return new LightingDataRampToLevelBuilderImpl(group, level);
  }

  public static class LightingDataRampToLevelBuilderImpl
      implements LightingData.LightingDataBuilder {
    private final byte group;
    private final byte level;

    public LightingDataRampToLevelBuilderImpl(byte group, byte level) {
      this.group = group;
      this.level = level;
    }

    public LightingDataRampToLevel build(LightingCommandTypeContainer commandTypeContainer) {
      LightingDataRampToLevel lightingDataRampToLevel =
          new LightingDataRampToLevel(commandTypeContainer, group, level);
      return lightingDataRampToLevel;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof LightingDataRampToLevel)) {
      return false;
    }
    LightingDataRampToLevel that = (LightingDataRampToLevel) o;
    return (getGroup() == that.getGroup())
        && (getLevel() == that.getLevel())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getGroup(), getLevel());
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