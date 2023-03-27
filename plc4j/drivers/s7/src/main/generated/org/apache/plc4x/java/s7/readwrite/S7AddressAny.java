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
package org.apache.plc4x.java.s7.readwrite;

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

public class S7AddressAny extends S7Address implements Message {

  // Accessors for discriminator values.
  public Short getAddressType() {
    return (short) 0x10;
  }

  // Properties.
  protected final TransportSize transportSize;
  protected final int numberOfElements;
  protected final int dbNumber;
  protected final MemoryArea area;
  protected final int byteAddress;
  protected final byte bitAddress;

  public S7AddressAny(
      TransportSize transportSize,
      int numberOfElements,
      int dbNumber,
      MemoryArea area,
      int byteAddress,
      byte bitAddress) {
    super();
    this.transportSize = transportSize;
    this.numberOfElements = numberOfElements;
    this.dbNumber = dbNumber;
    this.area = area;
    this.byteAddress = byteAddress;
    this.bitAddress = bitAddress;
  }

  public TransportSize getTransportSize() {
    return transportSize;
  }

  public int getNumberOfElements() {
    return numberOfElements;
  }

  public int getDbNumber() {
    return dbNumber;
  }

  public MemoryArea getArea() {
    return area;
  }

  public int getByteAddress() {
    return byteAddress;
  }

  public byte getBitAddress() {
    return bitAddress;
  }

  @Override
  protected void serializeS7AddressChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("S7AddressAny");

    // Enum field (transportSize)
    writeEnumField(
        "transportSize",
        "TransportSize",
        transportSize,
        new DataWriterEnumDefault<>(
            TransportSize::getCode, TransportSize::name, writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (numberOfElements)
    writeSimpleField("numberOfElements", numberOfElements, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (dbNumber)
    writeSimpleField("dbNumber", dbNumber, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (area)
    writeSimpleEnumField(
        "area",
        "MemoryArea",
        area,
        new DataWriterEnumDefault<>(
            MemoryArea::getValue, MemoryArea::name, writeUnsignedShort(writeBuffer, 8)));

    // Reserved Field (reserved)
    writeReservedField("reserved", (short) 0x00, writeUnsignedShort(writeBuffer, 5));

    // Simple Field (byteAddress)
    writeSimpleField("byteAddress", byteAddress, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (bitAddress)
    writeSimpleField("bitAddress", bitAddress, writeUnsignedByte(writeBuffer, 3));

    writeBuffer.popContext("S7AddressAny");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    S7AddressAny _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Enum Field (transportSize)
    lengthInBits += 8;

    // Simple field (numberOfElements)
    lengthInBits += 16;

    // Simple field (dbNumber)
    lengthInBits += 16;

    // Simple field (area)
    lengthInBits += 8;

    // Reserved Field (reserved)
    lengthInBits += 5;

    // Simple field (byteAddress)
    lengthInBits += 16;

    // Simple field (bitAddress)
    lengthInBits += 3;

    return lengthInBits;
  }

  public static S7AddressBuilder staticParseS7AddressBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("S7AddressAny");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    TransportSize transportSize =
        readEnumField(
            "transportSize",
            "TransportSize",
            readEnum(TransportSize::firstEnumForFieldCode, readUnsignedShort(readBuffer, 8)));

    int numberOfElements = readSimpleField("numberOfElements", readUnsignedInt(readBuffer, 16));

    int dbNumber = readSimpleField("dbNumber", readUnsignedInt(readBuffer, 16));

    MemoryArea area =
        readEnumField(
            "area",
            "MemoryArea",
            new DataReaderEnumDefault<>(
                MemoryArea::enumForValue, readUnsignedShort(readBuffer, 8)));

    Short reservedField0 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 5), (short) 0x00);

    int byteAddress = readSimpleField("byteAddress", readUnsignedInt(readBuffer, 16));

    byte bitAddress = readSimpleField("bitAddress", readUnsignedByte(readBuffer, 3));

    readBuffer.closeContext("S7AddressAny");
    // Create the instance
    return new S7AddressAnyBuilderImpl(
        transportSize, numberOfElements, dbNumber, area, byteAddress, bitAddress);
  }

  public static class S7AddressAnyBuilderImpl implements S7Address.S7AddressBuilder {
    private final TransportSize transportSize;
    private final int numberOfElements;
    private final int dbNumber;
    private final MemoryArea area;
    private final int byteAddress;
    private final byte bitAddress;

    public S7AddressAnyBuilderImpl(
        TransportSize transportSize,
        int numberOfElements,
        int dbNumber,
        MemoryArea area,
        int byteAddress,
        byte bitAddress) {
      this.transportSize = transportSize;
      this.numberOfElements = numberOfElements;
      this.dbNumber = dbNumber;
      this.area = area;
      this.byteAddress = byteAddress;
      this.bitAddress = bitAddress;
    }

    public S7AddressAny build() {
      S7AddressAny s7AddressAny =
          new S7AddressAny(
              transportSize, numberOfElements, dbNumber, area, byteAddress, bitAddress);
      return s7AddressAny;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof S7AddressAny)) {
      return false;
    }
    S7AddressAny that = (S7AddressAny) o;
    return (getTransportSize() == that.getTransportSize())
        && (getNumberOfElements() == that.getNumberOfElements())
        && (getDbNumber() == that.getDbNumber())
        && (getArea() == that.getArea())
        && (getByteAddress() == that.getByteAddress())
        && (getBitAddress() == that.getBitAddress())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getTransportSize(),
        getNumberOfElements(),
        getDbNumber(),
        getArea(),
        getByteAddress(),
        getBitAddress());
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
