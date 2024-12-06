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

public class BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString
    extends BACnetNotificationParametersChangeOfDiscreteValueNewValue implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetApplicationTagOctetString octetStringValue;

  // Arguments.
  protected final Short tagNumber;

  public BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetApplicationTagOctetString octetStringValue,
      Short tagNumber) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber);
    this.octetStringValue = octetStringValue;
    this.tagNumber = tagNumber;
  }

  public BACnetApplicationTagOctetString getOctetStringValue() {
    return octetStringValue;
  }

  @Override
  protected void serializeBACnetNotificationParametersChangeOfDiscreteValueNewValueChild(
      WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString");

    // Simple Field (octetStringValue)
    writeSimpleField("octetStringValue", octetStringValue, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (octetStringValue)
    lengthInBits += octetStringValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder
      staticParseBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder(
          ReadBuffer readBuffer, Short tagNumber) throws ParseException {
    readBuffer.pullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagOctetString octetStringValue =
        readSimpleField(
            "octetStringValue",
            readComplex(
                () ->
                    (BACnetApplicationTagOctetString) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString");
    // Create the instance
    return new BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetStringBuilderImpl(
        octetStringValue, tagNumber);
  }

  public static
  class BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetStringBuilderImpl
      implements BACnetNotificationParametersChangeOfDiscreteValueNewValue
          .BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder {
    private final BACnetApplicationTagOctetString octetStringValue;
    private final Short tagNumber;

    public BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetStringBuilderImpl(
        BACnetApplicationTagOctetString octetStringValue, Short tagNumber) {
      this.octetStringValue = octetStringValue;
      this.tagNumber = tagNumber;
    }

    public BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber) {
      BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString
          bACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString =
              new BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString(
                  openingTag, peekedTagHeader, closingTag, octetStringValue, tagNumber);
      return bACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString)) {
      return false;
    }
    BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString that =
        (BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString) o;
    return (getOctetStringValue() == that.getOctetStringValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getOctetStringValue());
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