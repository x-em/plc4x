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

public class KnxGroupAddressFreeLevel extends KnxGroupAddress implements Message {

  // Accessors for discriminator values.
  public Byte getNumLevels() {
    return (byte) 1;
  }

  // Properties.
  protected final int subGroup;

  public KnxGroupAddressFreeLevel(int subGroup) {
    super();
    this.subGroup = subGroup;
  }

  public int getSubGroup() {
    return subGroup;
  }

  @Override
  protected void serializeKnxGroupAddressChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("KnxGroupAddressFreeLevel");

    // Simple Field (subGroup)
    writeSimpleField("subGroup", subGroup, writeUnsignedInt(writeBuffer, 16));

    writeBuffer.popContext("KnxGroupAddressFreeLevel");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    KnxGroupAddressFreeLevel _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (subGroup)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static KnxGroupAddressBuilder staticParseKnxGroupAddressBuilder(
      ReadBuffer readBuffer, Byte numLevels) throws ParseException {
    readBuffer.pullContext("KnxGroupAddressFreeLevel");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int subGroup = readSimpleField("subGroup", readUnsignedInt(readBuffer, 16));

    readBuffer.closeContext("KnxGroupAddressFreeLevel");
    // Create the instance
    return new KnxGroupAddressFreeLevelBuilderImpl(subGroup);
  }

  public static class KnxGroupAddressFreeLevelBuilderImpl
      implements KnxGroupAddress.KnxGroupAddressBuilder {
    private final int subGroup;

    public KnxGroupAddressFreeLevelBuilderImpl(int subGroup) {
      this.subGroup = subGroup;
    }

    public KnxGroupAddressFreeLevel build() {
      KnxGroupAddressFreeLevel knxGroupAddressFreeLevel = new KnxGroupAddressFreeLevel(subGroup);
      return knxGroupAddressFreeLevel;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof KnxGroupAddressFreeLevel)) {
      return false;
    }
    KnxGroupAddressFreeLevel that = (KnxGroupAddressFreeLevel) o;
    return (getSubGroup() == that.getSubGroup()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getSubGroup());
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