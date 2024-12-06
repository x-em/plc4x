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
package org.apache.plc4x.java.cbus.readwrite;

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

public class ReplyNetwork implements Message {

  // Properties.
  protected final NetworkRoute networkRoute;
  protected final UnitAddress unitAddress;

  public ReplyNetwork(NetworkRoute networkRoute, UnitAddress unitAddress) {
    super();
    this.networkRoute = networkRoute;
    this.unitAddress = unitAddress;
  }

  public NetworkRoute getNetworkRoute() {
    return networkRoute;
  }

  public UnitAddress getUnitAddress() {
    return unitAddress;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ReplyNetwork");

    // Simple Field (networkRoute)
    writeSimpleField("networkRoute", networkRoute, writeComplex(writeBuffer));

    // Simple Field (unitAddress)
    writeSimpleField("unitAddress", unitAddress, writeComplex(writeBuffer));

    writeBuffer.popContext("ReplyNetwork");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    ReplyNetwork _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (networkRoute)
    lengthInBits += networkRoute.getLengthInBits();

    // Simple field (unitAddress)
    lengthInBits += unitAddress.getLengthInBits();

    return lengthInBits;
  }

  public static ReplyNetwork staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("ReplyNetwork");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NetworkRoute networkRoute =
        readSimpleField(
            "networkRoute", readComplex(() -> NetworkRoute.staticParse(readBuffer), readBuffer));

    UnitAddress unitAddress =
        readSimpleField(
            "unitAddress", readComplex(() -> UnitAddress.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("ReplyNetwork");
    // Create the instance
    ReplyNetwork _replyNetwork;
    _replyNetwork = new ReplyNetwork(networkRoute, unitAddress);
    return _replyNetwork;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ReplyNetwork)) {
      return false;
    }
    ReplyNetwork that = (ReplyNetwork) o;
    return (getNetworkRoute() == that.getNetworkRoute())
        && (getUnitAddress() == that.getUnitAddress())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getNetworkRoute(), getUnitAddress());
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
