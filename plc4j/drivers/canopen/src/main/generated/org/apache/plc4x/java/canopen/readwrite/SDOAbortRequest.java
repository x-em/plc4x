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
package org.apache.plc4x.java.canopen.readwrite;

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

public class SDOAbortRequest extends SDORequest implements Message {

  // Accessors for discriminator values.
  public SDORequestCommand getCommand() {
    return SDORequestCommand.ABORT;
  }

  // Properties.
  protected final SDOAbort abort;

  public SDOAbortRequest(SDOAbort abort) {
    super();
    this.abort = abort;
  }

  public SDOAbort getAbort() {
    return abort;
  }

  @Override
  protected void serializeSDORequestChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SDOAbortRequest");

    // Simple Field (abort)
    writeSimpleField("abort", abort, writeComplex(writeBuffer));

    writeBuffer.popContext("SDOAbortRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SDOAbortRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (abort)
    lengthInBits += abort.getLengthInBits();

    return lengthInBits;
  }

  public static SDORequestBuilder staticParseSDORequestBuilder(
      ReadBuffer readBuffer, SDORequestCommand command) throws ParseException {
    readBuffer.pullContext("SDOAbortRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    SDOAbort abort =
        readSimpleField("abort", readComplex(() -> SDOAbort.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("SDOAbortRequest");
    // Create the instance
    return new SDOAbortRequestBuilderImpl(abort);
  }

  public static class SDOAbortRequestBuilderImpl implements SDORequest.SDORequestBuilder {
    private final SDOAbort abort;

    public SDOAbortRequestBuilderImpl(SDOAbort abort) {
      this.abort = abort;
    }

    public SDOAbortRequest build() {
      SDOAbortRequest sDOAbortRequest = new SDOAbortRequest(abort);
      return sDOAbortRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SDOAbortRequest)) {
      return false;
    }
    SDOAbortRequest that = (SDOAbortRequest) o;
    return (getAbort() == that.getAbort()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getAbort());
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