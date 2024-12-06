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

public class BACnetConstructedDataIPv6ZoneIndex extends BACnetConstructedData implements Message {

  // Accessors for discriminator values.
  public BACnetObjectType getObjectTypeArgument() {
    return null;
  }

  public BACnetPropertyIdentifier getPropertyIdentifierArgument() {
    return BACnetPropertyIdentifier.IPV6_ZONE_INDEX;
  }

  // Properties.
  protected final BACnetApplicationTagCharacterString ipv6ZoneIndex;

  // Arguments.
  protected final Short tagNumber;
  protected final BACnetTagPayloadUnsignedInteger arrayIndexArgument;

  public BACnetConstructedDataIPv6ZoneIndex(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetApplicationTagCharacterString ipv6ZoneIndex,
      Short tagNumber,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument);
    this.ipv6ZoneIndex = ipv6ZoneIndex;
    this.tagNumber = tagNumber;
    this.arrayIndexArgument = arrayIndexArgument;
  }

  public BACnetApplicationTagCharacterString getIpv6ZoneIndex() {
    return ipv6ZoneIndex;
  }

  public BACnetApplicationTagCharacterString getActualValue() {
    return (BACnetApplicationTagCharacterString) (getIpv6ZoneIndex());
  }

  @Override
  protected void serializeBACnetConstructedDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetConstructedDataIPv6ZoneIndex");

    // Simple Field (ipv6ZoneIndex)
    writeSimpleField("ipv6ZoneIndex", ipv6ZoneIndex, writeComplex(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetApplicationTagCharacterString actualValue = getActualValue();
    writeBuffer.writeVirtual("actualValue", actualValue);

    writeBuffer.popContext("BACnetConstructedDataIPv6ZoneIndex");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConstructedDataIPv6ZoneIndex _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (ipv6ZoneIndex)
    lengthInBits += ipv6ZoneIndex.getLengthInBits();

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static BACnetConstructedDataBuilder staticParseBACnetConstructedDataBuilder(
      ReadBuffer readBuffer,
      Short tagNumber,
      BACnetObjectType objectTypeArgument,
      BACnetPropertyIdentifier propertyIdentifierArgument,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument)
      throws ParseException {
    readBuffer.pullContext("BACnetConstructedDataIPv6ZoneIndex");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagCharacterString ipv6ZoneIndex =
        readSimpleField(
            "ipv6ZoneIndex",
            readComplex(
                () ->
                    (BACnetApplicationTagCharacterString)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));
    BACnetApplicationTagCharacterString actualValue =
        readVirtualField("actualValue", BACnetApplicationTagCharacterString.class, ipv6ZoneIndex);

    readBuffer.closeContext("BACnetConstructedDataIPv6ZoneIndex");
    // Create the instance
    return new BACnetConstructedDataIPv6ZoneIndexBuilderImpl(
        ipv6ZoneIndex, tagNumber, arrayIndexArgument);
  }

  public static class BACnetConstructedDataIPv6ZoneIndexBuilderImpl
      implements BACnetConstructedData.BACnetConstructedDataBuilder {
    private final BACnetApplicationTagCharacterString ipv6ZoneIndex;
    private final Short tagNumber;
    private final BACnetTagPayloadUnsignedInteger arrayIndexArgument;

    public BACnetConstructedDataIPv6ZoneIndexBuilderImpl(
        BACnetApplicationTagCharacterString ipv6ZoneIndex,
        Short tagNumber,
        BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
      this.ipv6ZoneIndex = ipv6ZoneIndex;
      this.tagNumber = tagNumber;
      this.arrayIndexArgument = arrayIndexArgument;
    }

    public BACnetConstructedDataIPv6ZoneIndex build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber,
        BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
      BACnetConstructedDataIPv6ZoneIndex bACnetConstructedDataIPv6ZoneIndex =
          new BACnetConstructedDataIPv6ZoneIndex(
              openingTag,
              peekedTagHeader,
              closingTag,
              ipv6ZoneIndex,
              tagNumber,
              arrayIndexArgument);
      return bACnetConstructedDataIPv6ZoneIndex;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConstructedDataIPv6ZoneIndex)) {
      return false;
    }
    BACnetConstructedDataIPv6ZoneIndex that = (BACnetConstructedDataIPv6ZoneIndex) o;
    return (getIpv6ZoneIndex() == that.getIpv6ZoneIndex()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getIpv6ZoneIndex());
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