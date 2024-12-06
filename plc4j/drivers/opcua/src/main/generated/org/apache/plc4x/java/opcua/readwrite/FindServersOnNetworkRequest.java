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

public class FindServersOnNetworkRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 12192;
  }

  // Properties.
  protected final RequestHeader requestHeader;
  protected final long startingRecordId;
  protected final long maxRecordsToReturn;
  protected final List<PascalString> serverCapabilityFilter;

  public FindServersOnNetworkRequest(
      RequestHeader requestHeader,
      long startingRecordId,
      long maxRecordsToReturn,
      List<PascalString> serverCapabilityFilter) {
    super();
    this.requestHeader = requestHeader;
    this.startingRecordId = startingRecordId;
    this.maxRecordsToReturn = maxRecordsToReturn;
    this.serverCapabilityFilter = serverCapabilityFilter;
  }

  public RequestHeader getRequestHeader() {
    return requestHeader;
  }

  public long getStartingRecordId() {
    return startingRecordId;
  }

  public long getMaxRecordsToReturn() {
    return maxRecordsToReturn;
  }

  public List<PascalString> getServerCapabilityFilter() {
    return serverCapabilityFilter;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("FindServersOnNetworkRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, writeComplex(writeBuffer));

    // Simple Field (startingRecordId)
    writeSimpleField("startingRecordId", startingRecordId, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (maxRecordsToReturn)
    writeSimpleField("maxRecordsToReturn", maxRecordsToReturn, writeUnsignedLong(writeBuffer, 32));

    // Implicit Field (noOfServerCapabilityFilter) (Used for parsing, but its value is not stored as
    // it's implicitly given by the objects content)
    int noOfServerCapabilityFilter =
        (int)
            ((((getServerCapabilityFilter()) == (null))
                ? -(1)
                : COUNT(getServerCapabilityFilter())));
    writeImplicitField(
        "noOfServerCapabilityFilter", noOfServerCapabilityFilter, writeSignedInt(writeBuffer, 32));

    // Array Field (serverCapabilityFilter)
    writeComplexTypeArrayField("serverCapabilityFilter", serverCapabilityFilter, writeBuffer);

    writeBuffer.popContext("FindServersOnNetworkRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    FindServersOnNetworkRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Simple field (startingRecordId)
    lengthInBits += 32;

    // Simple field (maxRecordsToReturn)
    lengthInBits += 32;

    // Implicit Field (noOfServerCapabilityFilter)
    lengthInBits += 32;

    // Array field
    if (serverCapabilityFilter != null) {
      int i = 0;
      for (PascalString element : serverCapabilityFilter) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= serverCapabilityFilter.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("FindServersOnNetworkRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    RequestHeader requestHeader =
        readSimpleField(
            "requestHeader",
            readComplex(
                () ->
                    (RequestHeader) ExtensionObjectDefinition.staticParse(readBuffer, (int) (391)),
                readBuffer));

    long startingRecordId = readSimpleField("startingRecordId", readUnsignedLong(readBuffer, 32));

    long maxRecordsToReturn =
        readSimpleField("maxRecordsToReturn", readUnsignedLong(readBuffer, 32));

    int noOfServerCapabilityFilter =
        readImplicitField("noOfServerCapabilityFilter", readSignedInt(readBuffer, 32));

    List<PascalString> serverCapabilityFilter =
        readCountArrayField(
            "serverCapabilityFilter",
            readComplex(() -> PascalString.staticParse(readBuffer), readBuffer),
            noOfServerCapabilityFilter);

    readBuffer.closeContext("FindServersOnNetworkRequest");
    // Create the instance
    return new FindServersOnNetworkRequestBuilderImpl(
        requestHeader, startingRecordId, maxRecordsToReturn, serverCapabilityFilter);
  }

  public static class FindServersOnNetworkRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final RequestHeader requestHeader;
    private final long startingRecordId;
    private final long maxRecordsToReturn;
    private final List<PascalString> serverCapabilityFilter;

    public FindServersOnNetworkRequestBuilderImpl(
        RequestHeader requestHeader,
        long startingRecordId,
        long maxRecordsToReturn,
        List<PascalString> serverCapabilityFilter) {
      this.requestHeader = requestHeader;
      this.startingRecordId = startingRecordId;
      this.maxRecordsToReturn = maxRecordsToReturn;
      this.serverCapabilityFilter = serverCapabilityFilter;
    }

    public FindServersOnNetworkRequest build() {
      FindServersOnNetworkRequest findServersOnNetworkRequest =
          new FindServersOnNetworkRequest(
              requestHeader, startingRecordId, maxRecordsToReturn, serverCapabilityFilter);
      return findServersOnNetworkRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof FindServersOnNetworkRequest)) {
      return false;
    }
    FindServersOnNetworkRequest that = (FindServersOnNetworkRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getStartingRecordId() == that.getStartingRecordId())
        && (getMaxRecordsToReturn() == that.getMaxRecordsToReturn())
        && (getServerCapabilityFilter() == that.getServerCapabilityFilter())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getRequestHeader(),
        getStartingRecordId(),
        getMaxRecordsToReturn(),
        getServerCapabilityFilter());
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