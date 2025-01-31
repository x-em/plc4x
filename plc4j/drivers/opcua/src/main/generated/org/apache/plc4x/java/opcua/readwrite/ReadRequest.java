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

public class ReadRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 631;
  }

  // Properties.
  protected final RequestHeader requestHeader;
  protected final double maxAge;
  protected final TimestampsToReturn timestampsToReturn;
  protected final List<ReadValueId> nodesToRead;

  public ReadRequest(
      RequestHeader requestHeader,
      double maxAge,
      TimestampsToReturn timestampsToReturn,
      List<ReadValueId> nodesToRead) {
    super();
    this.requestHeader = requestHeader;
    this.maxAge = maxAge;
    this.timestampsToReturn = timestampsToReturn;
    this.nodesToRead = nodesToRead;
  }

  public RequestHeader getRequestHeader() {
    return requestHeader;
  }

  public double getMaxAge() {
    return maxAge;
  }

  public TimestampsToReturn getTimestampsToReturn() {
    return timestampsToReturn;
  }

  public List<ReadValueId> getNodesToRead() {
    return nodesToRead;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ReadRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, writeComplex(writeBuffer));

    // Simple Field (maxAge)
    writeSimpleField("maxAge", maxAge, writeDouble(writeBuffer, 64));

    // Simple Field (timestampsToReturn)
    writeSimpleEnumField(
        "timestampsToReturn",
        "TimestampsToReturn",
        timestampsToReturn,
        writeEnum(
            TimestampsToReturn::getValue,
            TimestampsToReturn::name,
            writeUnsignedLong(writeBuffer, 32)));

    // Implicit Field (noOfNodesToRead) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfNodesToRead = (int) ((((getNodesToRead()) == (null)) ? -(1) : COUNT(getNodesToRead())));
    writeImplicitField("noOfNodesToRead", noOfNodesToRead, writeSignedInt(writeBuffer, 32));

    // Array Field (nodesToRead)
    writeComplexTypeArrayField("nodesToRead", nodesToRead, writeBuffer);

    writeBuffer.popContext("ReadRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ReadRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Simple field (maxAge)
    lengthInBits += 64;

    // Simple field (timestampsToReturn)
    lengthInBits += 32;

    // Implicit Field (noOfNodesToRead)
    lengthInBits += 32;

    // Array field
    if (nodesToRead != null) {
      int i = 0;
      for (ReadValueId element : nodesToRead) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= nodesToRead.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("ReadRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    RequestHeader requestHeader =
        readSimpleField(
            "requestHeader",
            readComplex(
                () ->
                    (RequestHeader) ExtensionObjectDefinition.staticParse(readBuffer, (int) (391)),
                readBuffer));

    double maxAge = readSimpleField("maxAge", readDouble(readBuffer, 64));

    TimestampsToReturn timestampsToReturn =
        readEnumField(
            "timestampsToReturn",
            "TimestampsToReturn",
            readEnum(TimestampsToReturn::enumForValue, readUnsignedLong(readBuffer, 32)));

    int noOfNodesToRead = readImplicitField("noOfNodesToRead", readSignedInt(readBuffer, 32));

    List<ReadValueId> nodesToRead =
        readCountArrayField(
            "nodesToRead",
            readComplex(
                () -> (ReadValueId) ExtensionObjectDefinition.staticParse(readBuffer, (int) (628)),
                readBuffer),
            noOfNodesToRead);

    readBuffer.closeContext("ReadRequest");
    // Create the instance
    return new ReadRequestBuilderImpl(requestHeader, maxAge, timestampsToReturn, nodesToRead);
  }

  public static class ReadRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final RequestHeader requestHeader;
    private final double maxAge;
    private final TimestampsToReturn timestampsToReturn;
    private final List<ReadValueId> nodesToRead;

    public ReadRequestBuilderImpl(
        RequestHeader requestHeader,
        double maxAge,
        TimestampsToReturn timestampsToReturn,
        List<ReadValueId> nodesToRead) {
      this.requestHeader = requestHeader;
      this.maxAge = maxAge;
      this.timestampsToReturn = timestampsToReturn;
      this.nodesToRead = nodesToRead;
    }

    public ReadRequest build() {
      ReadRequest readRequest =
          new ReadRequest(requestHeader, maxAge, timestampsToReturn, nodesToRead);
      return readRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ReadRequest)) {
      return false;
    }
    ReadRequest that = (ReadRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getMaxAge() == that.getMaxAge())
        && (getTimestampsToReturn() == that.getTimestampsToReturn())
        && (getNodesToRead() == that.getNodesToRead())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getRequestHeader(),
        getMaxAge(),
        getTimestampsToReturn(),
        getNodesToRead());
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
