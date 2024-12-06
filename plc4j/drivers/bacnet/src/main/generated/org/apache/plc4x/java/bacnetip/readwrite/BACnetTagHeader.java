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

public class BACnetTagHeader implements Message {

  // Properties.
  protected final byte tagNumber;
  protected final TagClass tagClass;
  protected final byte lengthValueType;
  protected final Short extTagNumber;
  protected final Short extLength;
  protected final Integer extExtLength;
  protected final Long extExtExtLength;

  public BACnetTagHeader(
      byte tagNumber,
      TagClass tagClass,
      byte lengthValueType,
      Short extTagNumber,
      Short extLength,
      Integer extExtLength,
      Long extExtExtLength) {
    super();
    this.tagNumber = tagNumber;
    this.tagClass = tagClass;
    this.lengthValueType = lengthValueType;
    this.extTagNumber = extTagNumber;
    this.extLength = extLength;
    this.extExtLength = extExtLength;
    this.extExtExtLength = extExtExtLength;
  }

  public byte getTagNumber() {
    return tagNumber;
  }

  public TagClass getTagClass() {
    return tagClass;
  }

  public byte getLengthValueType() {
    return lengthValueType;
  }

  public Short getExtTagNumber() {
    return extTagNumber;
  }

  public Short getExtLength() {
    return extLength;
  }

  public Integer getExtExtLength() {
    return extExtLength;
  }

  public Long getExtExtExtLength() {
    return extExtExtLength;
  }

  public short getActualTagNumber() {
    return (short) ((((getTagNumber()) < (15)) ? getTagNumber() : getExtTagNumber()));
  }

  public boolean getIsBoolean() {
    return (boolean)
        (((getTagNumber()) == (1)) && ((getTagClass()) == (TagClass.APPLICATION_TAGS)));
  }

  public boolean getIsConstructed() {
    return (boolean)
        (((getTagClass()) == (TagClass.CONTEXT_SPECIFIC_TAGS)) && ((getLengthValueType()) == (6)));
  }

  public boolean getIsPrimitiveAndNotBoolean() {
    return (boolean) ((!(getIsConstructed())) && (!(getIsBoolean())));
  }

  public long getActualLength() {
    return (long)
        (((((getLengthValueType()) == (5)) && ((getExtLength()) == (255)))
            ? getExtExtExtLength()
            : (((((getLengthValueType()) == (5)) && ((getExtLength()) == (254)))
                ? getExtExtLength()
                : ((((getLengthValueType()) == (5)) ? getExtLength() : getLengthValueType()))))));
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetTagHeader");

    // Simple Field (tagNumber)
    writeSimpleField("tagNumber", tagNumber, writeUnsignedByte(writeBuffer, 4));

    // Simple Field (tagClass)
    writeSimpleEnumField(
        "tagClass",
        "TagClass",
        tagClass,
        writeEnum(TagClass::getValue, TagClass::name, writeUnsignedByte(writeBuffer, 1)));

    // Simple Field (lengthValueType)
    writeSimpleField("lengthValueType", lengthValueType, writeUnsignedByte(writeBuffer, 3));

    // Optional Field (extTagNumber) (Can be skipped, if the value is null)
    writeOptionalField(
        "extTagNumber", extTagNumber, writeUnsignedShort(writeBuffer, 8), (getTagNumber()) == (15));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    short actualTagNumber = getActualTagNumber();
    writeBuffer.writeVirtual("actualTagNumber", actualTagNumber);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isBoolean = getIsBoolean();
    writeBuffer.writeVirtual("isBoolean", isBoolean);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isConstructed = getIsConstructed();
    writeBuffer.writeVirtual("isConstructed", isConstructed);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isPrimitiveAndNotBoolean = getIsPrimitiveAndNotBoolean();
    writeBuffer.writeVirtual("isPrimitiveAndNotBoolean", isPrimitiveAndNotBoolean);

    // Optional Field (extLength) (Can be skipped, if the value is null)
    writeOptionalField(
        "extLength",
        extLength,
        writeUnsignedShort(writeBuffer, 8),
        (getIsPrimitiveAndNotBoolean()) && ((getLengthValueType()) == (5)));

    // Optional Field (extExtLength) (Can be skipped, if the value is null)
    writeOptionalField(
        "extExtLength",
        extExtLength,
        writeUnsignedInt(writeBuffer, 16),
        ((getIsPrimitiveAndNotBoolean()) && ((getLengthValueType()) == (5)))
            && ((getExtLength()) == (254)));

    // Optional Field (extExtExtLength) (Can be skipped, if the value is null)
    writeOptionalField(
        "extExtExtLength",
        extExtExtLength,
        writeUnsignedLong(writeBuffer, 32),
        ((getIsPrimitiveAndNotBoolean()) && ((getLengthValueType()) == (5)))
            && ((getExtLength()) == (255)));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    long actualLength = getActualLength();
    writeBuffer.writeVirtual("actualLength", actualLength);

    writeBuffer.popContext("BACnetTagHeader");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetTagHeader _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (tagNumber)
    lengthInBits += 4;

    // Simple field (tagClass)
    lengthInBits += 1;

    // Simple field (lengthValueType)
    lengthInBits += 3;

    // Optional Field (extTagNumber)
    if (extTagNumber != null) {
      lengthInBits += 8;
    }

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // Optional Field (extLength)
    if (extLength != null) {
      lengthInBits += 8;
    }

    // Optional Field (extExtLength)
    if (extExtLength != null) {
      lengthInBits += 16;
    }

    // Optional Field (extExtExtLength)
    if (extExtExtLength != null) {
      lengthInBits += 32;
    }

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static BACnetTagHeader staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetTagHeader");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte tagNumber = readSimpleField("tagNumber", readUnsignedByte(readBuffer, 4));

    TagClass tagClass =
        readEnumField(
            "tagClass",
            "TagClass",
            readEnum(TagClass::enumForValue, readUnsignedByte(readBuffer, 1)));

    byte lengthValueType = readSimpleField("lengthValueType", readUnsignedByte(readBuffer, 3));

    Short extTagNumber =
        readOptionalField("extTagNumber", readUnsignedShort(readBuffer, 8), (tagNumber) == (15));
    short actualTagNumber =
        readVirtualField(
            "actualTagNumber", short.class, (((tagNumber) < (15)) ? tagNumber : extTagNumber));
    boolean isBoolean =
        readVirtualField(
            "isBoolean",
            boolean.class,
            ((tagNumber) == (1)) && ((tagClass) == (TagClass.APPLICATION_TAGS)));
    boolean isConstructed =
        readVirtualField(
            "isConstructed",
            boolean.class,
            ((tagClass) == (TagClass.CONTEXT_SPECIFIC_TAGS)) && ((lengthValueType) == (6)));
    boolean isPrimitiveAndNotBoolean =
        readVirtualField(
            "isPrimitiveAndNotBoolean", boolean.class, (!(isConstructed)) && (!(isBoolean)));

    Short extLength =
        readOptionalField(
            "extLength",
            readUnsignedShort(readBuffer, 8),
            (isPrimitiveAndNotBoolean) && ((lengthValueType) == (5)));

    Integer extExtLength =
        readOptionalField(
            "extExtLength",
            readUnsignedInt(readBuffer, 16),
            ((isPrimitiveAndNotBoolean) && ((lengthValueType) == (5))) && ((extLength) == (254)));

    Long extExtExtLength =
        readOptionalField(
            "extExtExtLength",
            readUnsignedLong(readBuffer, 32),
            ((isPrimitiveAndNotBoolean) && ((lengthValueType) == (5))) && ((extLength) == (255)));
    long actualLength =
        readVirtualField(
            "actualLength",
            long.class,
            ((((lengthValueType) == (5)) && ((extLength) == (255)))
                ? extExtExtLength
                : (((((lengthValueType) == (5)) && ((extLength) == (254)))
                    ? extExtLength
                    : ((((lengthValueType) == (5)) ? extLength : lengthValueType))))));

    readBuffer.closeContext("BACnetTagHeader");
    // Create the instance
    BACnetTagHeader _bACnetTagHeader;
    _bACnetTagHeader =
        new BACnetTagHeader(
            tagNumber,
            tagClass,
            lengthValueType,
            extTagNumber,
            extLength,
            extExtLength,
            extExtExtLength);
    return _bACnetTagHeader;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetTagHeader)) {
      return false;
    }
    BACnetTagHeader that = (BACnetTagHeader) o;
    return (getTagNumber() == that.getTagNumber())
        && (getTagClass() == that.getTagClass())
        && (getLengthValueType() == that.getLengthValueType())
        && (getExtTagNumber() == that.getExtTagNumber())
        && (getExtLength() == that.getExtLength())
        && (getExtExtLength() == that.getExtExtLength())
        && (getExtExtExtLength() == that.getExtExtExtLength())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getTagNumber(),
        getTagClass(),
        getLengthValueType(),
        getExtTagNumber(),
        getExtLength(),
        getExtExtLength(),
        getExtExtExtLength());
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
