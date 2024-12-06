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
package org.apache.plc4x.java.knxnetip.readwrite;

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

public class KnxGroupAddress2Level extends KnxGroupAddress implements Message {

  // Accessors for discriminator values.
  public Byte getNumLevels() {
    return (byte) 2;
  }

  // Properties.
  protected final byte mainGroup;
  protected final short subGroup;

  public KnxGroupAddress2Level(byte mainGroup, short subGroup) {
    super();
    this.mainGroup = mainGroup;
    this.subGroup = subGroup;
  }

  public byte getMainGroup() {
    return mainGroup;
  }

  public short getSubGroup() {
    return subGroup;
  }

  @Override
  protected void serializeKnxGroupAddressChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("KnxGroupAddress2Level");

    // Simple Field (mainGroup)
    writeSimpleField("mainGroup", mainGroup, writeUnsignedByte(writeBuffer, 5));

    // Simple Field (subGroup)
    writeSimpleField("subGroup", subGroup, writeUnsignedShort(writeBuffer, 11));

    writeBuffer.popContext("KnxGroupAddress2Level");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    KnxGroupAddress2Level _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (mainGroup)
    lengthInBits += 5;

    // Simple field (subGroup)
    lengthInBits += 11;

    return lengthInBits;
  }

  public static KnxGroupAddressBuilder staticParseKnxGroupAddressBuilder(
      ReadBuffer readBuffer, Byte numLevels) throws ParseException {
    readBuffer.pullContext("KnxGroupAddress2Level");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte mainGroup = readSimpleField("mainGroup", readUnsignedByte(readBuffer, 5));

    short subGroup = readSimpleField("subGroup", readUnsignedShort(readBuffer, 11));

    readBuffer.closeContext("KnxGroupAddress2Level");
    // Create the instance
    return new KnxGroupAddress2LevelBuilderImpl(mainGroup, subGroup);
  }

  public static class KnxGroupAddress2LevelBuilderImpl
      implements KnxGroupAddress.KnxGroupAddressBuilder {
    private final byte mainGroup;
    private final short subGroup;

    public KnxGroupAddress2LevelBuilderImpl(byte mainGroup, short subGroup) {
      this.mainGroup = mainGroup;
      this.subGroup = subGroup;
    }

    public KnxGroupAddress2Level build() {
      KnxGroupAddress2Level knxGroupAddress2Level = new KnxGroupAddress2Level(mainGroup, subGroup);
      return knxGroupAddress2Level;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof KnxGroupAddress2Level)) {
      return false;
    }
    KnxGroupAddress2Level that = (KnxGroupAddress2Level) o;
    return (getMainGroup() == that.getMainGroup())
        && (getSubGroup() == that.getSubGroup())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getMainGroup(), getSubGroup());
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