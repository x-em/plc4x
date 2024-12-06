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

public class BACnetConfirmedServiceRequestAtomicReadFile extends BACnetConfirmedServiceRequest
    implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getServiceChoice() {
    return BACnetConfirmedServiceChoice.ATOMIC_READ_FILE;
  }

  // Properties.
  protected final BACnetApplicationTagObjectIdentifier fileIdentifier;
  protected final BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord accessMethod;

  // Arguments.
  protected final Long serviceRequestLength;

  public BACnetConfirmedServiceRequestAtomicReadFile(
      BACnetApplicationTagObjectIdentifier fileIdentifier,
      BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord accessMethod,
      Long serviceRequestLength) {
    super(serviceRequestLength);
    this.fileIdentifier = fileIdentifier;
    this.accessMethod = accessMethod;
    this.serviceRequestLength = serviceRequestLength;
  }

  public BACnetApplicationTagObjectIdentifier getFileIdentifier() {
    return fileIdentifier;
  }

  public BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord getAccessMethod() {
    return accessMethod;
  }

  @Override
  protected void serializeBACnetConfirmedServiceRequestChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetConfirmedServiceRequestAtomicReadFile");

    // Simple Field (fileIdentifier)
    writeSimpleField("fileIdentifier", fileIdentifier, writeComplex(writeBuffer));

    // Simple Field (accessMethod)
    writeSimpleField("accessMethod", accessMethod, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetConfirmedServiceRequestAtomicReadFile");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConfirmedServiceRequestAtomicReadFile _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (fileIdentifier)
    lengthInBits += fileIdentifier.getLengthInBits();

    // Simple field (accessMethod)
    lengthInBits += accessMethod.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetConfirmedServiceRequestBuilder
      staticParseBACnetConfirmedServiceRequestBuilder(
          ReadBuffer readBuffer, Long serviceRequestLength) throws ParseException {
    readBuffer.pullContext("BACnetConfirmedServiceRequestAtomicReadFile");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagObjectIdentifier fileIdentifier =
        readSimpleField(
            "fileIdentifier",
            readComplex(
                () ->
                    (BACnetApplicationTagObjectIdentifier)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord accessMethod =
        readSimpleField(
            "accessMethod",
            readComplex(
                () ->
                    BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord.staticParse(
                        readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetConfirmedServiceRequestAtomicReadFile");
    // Create the instance
    return new BACnetConfirmedServiceRequestAtomicReadFileBuilderImpl(
        fileIdentifier, accessMethod, serviceRequestLength);
  }

  public static class BACnetConfirmedServiceRequestAtomicReadFileBuilderImpl
      implements BACnetConfirmedServiceRequest.BACnetConfirmedServiceRequestBuilder {
    private final BACnetApplicationTagObjectIdentifier fileIdentifier;
    private final BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord accessMethod;
    private final Long serviceRequestLength;

    public BACnetConfirmedServiceRequestAtomicReadFileBuilderImpl(
        BACnetApplicationTagObjectIdentifier fileIdentifier,
        BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord accessMethod,
        Long serviceRequestLength) {
      this.fileIdentifier = fileIdentifier;
      this.accessMethod = accessMethod;
      this.serviceRequestLength = serviceRequestLength;
    }

    public BACnetConfirmedServiceRequestAtomicReadFile build(Long serviceRequestLength) {

      BACnetConfirmedServiceRequestAtomicReadFile bACnetConfirmedServiceRequestAtomicReadFile =
          new BACnetConfirmedServiceRequestAtomicReadFile(
              fileIdentifier, accessMethod, serviceRequestLength);
      return bACnetConfirmedServiceRequestAtomicReadFile;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConfirmedServiceRequestAtomicReadFile)) {
      return false;
    }
    BACnetConfirmedServiceRequestAtomicReadFile that =
        (BACnetConfirmedServiceRequestAtomicReadFile) o;
    return (getFileIdentifier() == that.getFileIdentifier())
        && (getAccessMethod() == that.getAccessMethod())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getFileIdentifier(), getAccessMethod());
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