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

public class BACnetCOVSubscription implements Message {

  // Properties.
  protected final BACnetRecipientProcessEnclosed recipient;
  protected final BACnetObjectPropertyReferenceEnclosed monitoredPropertyReference;
  protected final BACnetContextTagBoolean issueConfirmedNotifications;
  protected final BACnetContextTagUnsignedInteger timeRemaining;
  protected final BACnetContextTagReal covIncrement;

  public BACnetCOVSubscription(
      BACnetRecipientProcessEnclosed recipient,
      BACnetObjectPropertyReferenceEnclosed monitoredPropertyReference,
      BACnetContextTagBoolean issueConfirmedNotifications,
      BACnetContextTagUnsignedInteger timeRemaining,
      BACnetContextTagReal covIncrement) {
    super();
    this.recipient = recipient;
    this.monitoredPropertyReference = monitoredPropertyReference;
    this.issueConfirmedNotifications = issueConfirmedNotifications;
    this.timeRemaining = timeRemaining;
    this.covIncrement = covIncrement;
  }

  public BACnetRecipientProcessEnclosed getRecipient() {
    return recipient;
  }

  public BACnetObjectPropertyReferenceEnclosed getMonitoredPropertyReference() {
    return monitoredPropertyReference;
  }

  public BACnetContextTagBoolean getIssueConfirmedNotifications() {
    return issueConfirmedNotifications;
  }

  public BACnetContextTagUnsignedInteger getTimeRemaining() {
    return timeRemaining;
  }

  public BACnetContextTagReal getCovIncrement() {
    return covIncrement;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetCOVSubscription");

    // Simple Field (recipient)
    writeSimpleField("recipient", recipient, writeComplex(writeBuffer));

    // Simple Field (monitoredPropertyReference)
    writeSimpleField(
        "monitoredPropertyReference", monitoredPropertyReference, writeComplex(writeBuffer));

    // Simple Field (issueConfirmedNotifications)
    writeSimpleField(
        "issueConfirmedNotifications", issueConfirmedNotifications, writeComplex(writeBuffer));

    // Simple Field (timeRemaining)
    writeSimpleField("timeRemaining", timeRemaining, writeComplex(writeBuffer));

    // Optional Field (covIncrement) (Can be skipped, if the value is null)
    writeOptionalField("covIncrement", covIncrement, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetCOVSubscription");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetCOVSubscription _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (recipient)
    lengthInBits += recipient.getLengthInBits();

    // Simple field (monitoredPropertyReference)
    lengthInBits += monitoredPropertyReference.getLengthInBits();

    // Simple field (issueConfirmedNotifications)
    lengthInBits += issueConfirmedNotifications.getLengthInBits();

    // Simple field (timeRemaining)
    lengthInBits += timeRemaining.getLengthInBits();

    // Optional Field (covIncrement)
    if (covIncrement != null) {
      lengthInBits += covIncrement.getLengthInBits();
    }

    return lengthInBits;
  }

  public static BACnetCOVSubscription staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetCOVSubscription");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetRecipientProcessEnclosed recipient =
        readSimpleField(
            "recipient",
            readComplex(
                () -> BACnetRecipientProcessEnclosed.staticParse(readBuffer, (short) (0)),
                readBuffer));

    BACnetObjectPropertyReferenceEnclosed monitoredPropertyReference =
        readSimpleField(
            "monitoredPropertyReference",
            readComplex(
                () -> BACnetObjectPropertyReferenceEnclosed.staticParse(readBuffer, (short) (1)),
                readBuffer));

    BACnetContextTagBoolean issueConfirmedNotifications =
        readSimpleField(
            "issueConfirmedNotifications",
            readComplex(
                () ->
                    (BACnetContextTagBoolean)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (2), (BACnetDataType) (BACnetDataType.BOOLEAN)),
                readBuffer));

    BACnetContextTagUnsignedInteger timeRemaining =
        readSimpleField(
            "timeRemaining",
            readComplex(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (3),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetContextTagReal covIncrement =
        readOptionalField(
            "covIncrement",
            readComplex(
                () ->
                    (BACnetContextTagReal)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (4), (BACnetDataType) (BACnetDataType.REAL)),
                readBuffer));

    readBuffer.closeContext("BACnetCOVSubscription");
    // Create the instance
    BACnetCOVSubscription _bACnetCOVSubscription;
    _bACnetCOVSubscription =
        new BACnetCOVSubscription(
            recipient,
            monitoredPropertyReference,
            issueConfirmedNotifications,
            timeRemaining,
            covIncrement);
    return _bACnetCOVSubscription;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetCOVSubscription)) {
      return false;
    }
    BACnetCOVSubscription that = (BACnetCOVSubscription) o;
    return (getRecipient() == that.getRecipient())
        && (getMonitoredPropertyReference() == that.getMonitoredPropertyReference())
        && (getIssueConfirmedNotifications() == that.getIssueConfirmedNotifications())
        && (getTimeRemaining() == that.getTimeRemaining())
        && (getCovIncrement() == that.getCovIncrement())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getRecipient(),
        getMonitoredPropertyReference(),
        getIssueConfirmedNotifications(),
        getTimeRemaining(),
        getCovIncrement());
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