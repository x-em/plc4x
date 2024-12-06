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

public class BrowsePathResult extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 551;
  }

  // Properties.
  protected final StatusCode statusCode;
  protected final List<BrowsePathTarget> targets;

  public BrowsePathResult(StatusCode statusCode, List<BrowsePathTarget> targets) {
    super();
    this.statusCode = statusCode;
    this.targets = targets;
  }

  public StatusCode getStatusCode() {
    return statusCode;
  }

  public List<BrowsePathTarget> getTargets() {
    return targets;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BrowsePathResult");

    // Simple Field (statusCode)
    writeSimpleField("statusCode", statusCode, writeComplex(writeBuffer));

    // Implicit Field (noOfTargets) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfTargets = (int) ((((getTargets()) == (null)) ? -(1) : COUNT(getTargets())));
    writeImplicitField("noOfTargets", noOfTargets, writeSignedInt(writeBuffer, 32));

    // Array Field (targets)
    writeComplexTypeArrayField("targets", targets, writeBuffer);

    writeBuffer.popContext("BrowsePathResult");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BrowsePathResult _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (statusCode)
    lengthInBits += statusCode.getLengthInBits();

    // Implicit Field (noOfTargets)
    lengthInBits += 32;

    // Array field
    if (targets != null) {
      int i = 0;
      for (BrowsePathTarget element : targets) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= targets.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("BrowsePathResult");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    StatusCode statusCode =
        readSimpleField(
            "statusCode", readComplex(() -> StatusCode.staticParse(readBuffer), readBuffer));

    int noOfTargets = readImplicitField("noOfTargets", readSignedInt(readBuffer, 32));

    List<BrowsePathTarget> targets =
        readCountArrayField(
            "targets",
            readComplex(
                () ->
                    (BrowsePathTarget)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (548)),
                readBuffer),
            noOfTargets);

    readBuffer.closeContext("BrowsePathResult");
    // Create the instance
    return new BrowsePathResultBuilderImpl(statusCode, targets);
  }

  public static class BrowsePathResultBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final StatusCode statusCode;
    private final List<BrowsePathTarget> targets;

    public BrowsePathResultBuilderImpl(StatusCode statusCode, List<BrowsePathTarget> targets) {
      this.statusCode = statusCode;
      this.targets = targets;
    }

    public BrowsePathResult build() {
      BrowsePathResult browsePathResult = new BrowsePathResult(statusCode, targets);
      return browsePathResult;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BrowsePathResult)) {
      return false;
    }
    BrowsePathResult that = (BrowsePathResult) o;
    return (getStatusCode() == that.getStatusCode())
        && (getTargets() == that.getTargets())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getStatusCode(), getTargets());
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
