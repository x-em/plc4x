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

public class BACnetAccessAuthenticationFactorDisableTagged implements Message {

  // Properties.
  protected final BACnetTagHeader header;
  protected final BACnetAccessAuthenticationFactorDisable value;
  protected final long proprietaryValue;

  // Arguments.
  protected final Short tagNumber;
  protected final TagClass tagClass;

  public BACnetAccessAuthenticationFactorDisableTagged(
      BACnetTagHeader header,
      BACnetAccessAuthenticationFactorDisable value,
      long proprietaryValue,
      Short tagNumber,
      TagClass tagClass) {
    super();
    this.header = header;
    this.value = value;
    this.proprietaryValue = proprietaryValue;
    this.tagNumber = tagNumber;
    this.tagClass = tagClass;
  }

  public BACnetTagHeader getHeader() {
    return header;
  }

  public BACnetAccessAuthenticationFactorDisable getValue() {
    return value;
  }

  public long getProprietaryValue() {
    return proprietaryValue;
  }

  public boolean getIsProprietary() {
    return (boolean)
        ((getValue()) == (BACnetAccessAuthenticationFactorDisable.VENDOR_PROPRIETARY_VALUE));
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetAccessAuthenticationFactorDisableTagged");

    // Simple Field (header)
    writeSimpleField("header", header, writeComplex(writeBuffer));

    // Manual Field (value)
    writeManualField(
        "value",
        () ->
            org.apache.plc4x.java.bacnetip.readwrite.utils.StaticHelper.writeEnumGeneric(
                writeBuffer, value),
        writeBuffer);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isProprietary = getIsProprietary();
    writeBuffer.writeVirtual("isProprietary", isProprietary);

    // Manual Field (proprietaryValue)
    writeManualField(
        "proprietaryValue",
        () ->
            org.apache.plc4x.java.bacnetip.readwrite.utils.StaticHelper.writeProprietaryEnumGeneric(
                writeBuffer, proprietaryValue, isProprietary),
        writeBuffer);

    writeBuffer.popContext("BACnetAccessAuthenticationFactorDisableTagged");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetAccessAuthenticationFactorDisableTagged _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (header)
    lengthInBits += header.getLengthInBits();

    // Manual Field (value)
    lengthInBits += ((_value.getIsProprietary()) ? 0 : ((header.getActualLength()) * (8)));

    // A virtual field doesn't have any in- or output.

    // Manual Field (proprietaryValue)
    lengthInBits += ((_value.getIsProprietary()) ? ((header.getActualLength()) * (8)) : 0);

    return lengthInBits;
  }

  public static BACnetAccessAuthenticationFactorDisableTagged staticParse(
      ReadBuffer readBuffer, Short tagNumber, TagClass tagClass) throws ParseException {
    readBuffer.pullContext("BACnetAccessAuthenticationFactorDisableTagged");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetTagHeader header =
        readSimpleField(
            "header", readComplex(() -> BACnetTagHeader.staticParse(readBuffer), readBuffer));
    // Validation
    if (!((header.getTagClass()) == (tagClass))) {
      throw new ParseValidationException("tag class doesn't match");
    }
    // Validation
    if (!((((header.getTagClass()) == (TagClass.APPLICATION_TAGS)))
        || (((header.getActualTagNumber()) == (tagNumber))))) {
      throw new ParseAssertException("tagnumber doesn't match");
    }

    BACnetAccessAuthenticationFactorDisable value =
        readManualField(
            "value",
            readBuffer,
            () ->
                (BACnetAccessAuthenticationFactorDisable)
                    (org.apache.plc4x.java.bacnetip.readwrite.utils.StaticHelper.readEnumGeneric(
                        readBuffer,
                        header.getActualLength(),
                        BACnetAccessAuthenticationFactorDisable.VENDOR_PROPRIETARY_VALUE)));
    boolean isProprietary =
        readVirtualField(
            "isProprietary",
            boolean.class,
            (value) == (BACnetAccessAuthenticationFactorDisable.VENDOR_PROPRIETARY_VALUE));

    long proprietaryValue =
        readManualField(
            "proprietaryValue",
            readBuffer,
            () ->
                (long)
                    (org.apache.plc4x.java.bacnetip.readwrite.utils.StaticHelper
                        .readProprietaryEnumGeneric(
                            readBuffer, header.getActualLength(), isProprietary)));

    readBuffer.closeContext("BACnetAccessAuthenticationFactorDisableTagged");
    // Create the instance
    BACnetAccessAuthenticationFactorDisableTagged _bACnetAccessAuthenticationFactorDisableTagged;
    _bACnetAccessAuthenticationFactorDisableTagged =
        new BACnetAccessAuthenticationFactorDisableTagged(
            header, value, proprietaryValue, tagNumber, tagClass);
    return _bACnetAccessAuthenticationFactorDisableTagged;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetAccessAuthenticationFactorDisableTagged)) {
      return false;
    }
    BACnetAccessAuthenticationFactorDisableTagged that =
        (BACnetAccessAuthenticationFactorDisableTagged) o;
    return (getHeader() == that.getHeader())
        && (getValue() == that.getValue())
        && (getProprietaryValue() == that.getProprietaryValue())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getHeader(), getValue(), getProprietaryValue());
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
