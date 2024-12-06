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
package org.apache.plc4x.java.opcua.readwrite;

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

public class BrowseNextRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 533;
  }

  // Properties.
  protected final RequestHeader requestHeader;
  protected final boolean releaseContinuationPoints;
  protected final List<PascalByteString> continuationPoints;

  public BrowseNextRequest(
      RequestHeader requestHeader,
      boolean releaseContinuationPoints,
      List<PascalByteString> continuationPoints) {
    super();
    this.requestHeader = requestHeader;
    this.releaseContinuationPoints = releaseContinuationPoints;
    this.continuationPoints = continuationPoints;
  }

  public RequestHeader getRequestHeader() {
    return requestHeader;
  }

  public boolean getReleaseContinuationPoints() {
    return releaseContinuationPoints;
  }

  public List<PascalByteString> getContinuationPoints() {
    return continuationPoints;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BrowseNextRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, writeComplex(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0x00, writeUnsignedByte(writeBuffer, 7));

    // Simple Field (releaseContinuationPoints)
    writeSimpleField(
        "releaseContinuationPoints", releaseContinuationPoints, writeBoolean(writeBuffer));

    // Implicit Field (noOfContinuationPoints) (Used for parsing, but its value is not stored as
    // it's implicitly given by the objects content)
    int noOfContinuationPoints =
        (int) ((((getContinuationPoints()) == (null)) ? -(1) : COUNT(getContinuationPoints())));
    writeImplicitField(
        "noOfContinuationPoints", noOfContinuationPoints, writeSignedInt(writeBuffer, 32));

    // Array Field (continuationPoints)
    writeComplexTypeArrayField("continuationPoints", continuationPoints, writeBuffer);

    writeBuffer.popContext("BrowseNextRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BrowseNextRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Reserved Field (reserved)
    lengthInBits += 7;

    // Simple field (releaseContinuationPoints)
    lengthInBits += 1;

    // Implicit Field (noOfContinuationPoints)
    lengthInBits += 32;

    // Array field
    if (continuationPoints != null) {
      int i = 0;
      for (PascalByteString element : continuationPoints) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= continuationPoints.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("BrowseNextRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    RequestHeader requestHeader =
        readSimpleField(
            "requestHeader",
            readComplex(
                () ->
                    (RequestHeader) ExtensionObjectDefinition.staticParse(readBuffer, (int) (391)),
                readBuffer));

    Byte reservedField0 =
        readReservedField("reserved", readUnsignedByte(readBuffer, 7), (byte) 0x00);

    boolean releaseContinuationPoints =
        readSimpleField("releaseContinuationPoints", readBoolean(readBuffer));

    int noOfContinuationPoints =
        readImplicitField("noOfContinuationPoints", readSignedInt(readBuffer, 32));

    List<PascalByteString> continuationPoints =
        readCountArrayField(
            "continuationPoints",
            readComplex(() -> PascalByteString.staticParse(readBuffer), readBuffer),
            noOfContinuationPoints);

    readBuffer.closeContext("BrowseNextRequest");
    // Create the instance
    return new BrowseNextRequestBuilderImpl(
        requestHeader, releaseContinuationPoints, continuationPoints);
  }

  public static class BrowseNextRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final RequestHeader requestHeader;
    private final boolean releaseContinuationPoints;
    private final List<PascalByteString> continuationPoints;

    public BrowseNextRequestBuilderImpl(
        RequestHeader requestHeader,
        boolean releaseContinuationPoints,
        List<PascalByteString> continuationPoints) {
      this.requestHeader = requestHeader;
      this.releaseContinuationPoints = releaseContinuationPoints;
      this.continuationPoints = continuationPoints;
    }

    public BrowseNextRequest build() {
      BrowseNextRequest browseNextRequest =
          new BrowseNextRequest(requestHeader, releaseContinuationPoints, continuationPoints);
      return browseNextRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BrowseNextRequest)) {
      return false;
    }
    BrowseNextRequest that = (BrowseNextRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getReleaseContinuationPoints() == that.getReleaseContinuationPoints())
        && (getContinuationPoints() == that.getContinuationPoints())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getRequestHeader(),
        getReleaseContinuationPoints(),
        getContinuationPoints());
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