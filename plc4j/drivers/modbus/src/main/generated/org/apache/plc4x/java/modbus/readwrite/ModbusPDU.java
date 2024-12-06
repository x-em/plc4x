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
package org.apache.plc4x.java.modbus.readwrite;

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

public abstract class ModbusPDU implements Message {

  // Abstract accessors for discriminator values.
  public abstract Boolean getErrorFlag();

  public abstract Byte getFunctionFlag();

  public abstract Boolean getResponse();

  public ModbusPDU() {
    super();
  }

  protected abstract void serializeModbusPDUChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ModbusPDU");

    // Discriminator Field (errorFlag) (Used as input to a switch field)
    writeDiscriminatorField("errorFlag", getErrorFlag(), writeBoolean(writeBuffer));

    // Discriminator Field (functionFlag) (Used as input to a switch field)
    writeDiscriminatorField("functionFlag", getFunctionFlag(), writeUnsignedByte(writeBuffer, 7));

    // Switch field (Serialize the sub-type)
    serializeModbusPDUChild(writeBuffer);

    writeBuffer.popContext("ModbusPDU");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    ModbusPDU _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Discriminator Field (errorFlag)
    lengthInBits += 1;

    // Discriminator Field (functionFlag)
    lengthInBits += 7;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static ModbusPDU staticParse(ReadBuffer readBuffer, Boolean response)
      throws ParseException {
    readBuffer.pullContext("ModbusPDU");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    boolean errorFlag = readDiscriminatorField("errorFlag", readBoolean(readBuffer));

    byte functionFlag = readDiscriminatorField("functionFlag", readUnsignedByte(readBuffer, 7));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    ModbusPDUBuilder builder = null;
    if (EvaluationHelper.equals(errorFlag, (boolean) true)) {
      builder = ModbusPDUError.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x02)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder =
          ModbusPDUReadDiscreteInputsRequest.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x02)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder =
          ModbusPDUReadDiscreteInputsResponse.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x01)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = ModbusPDUReadCoilsRequest.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x01)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = ModbusPDUReadCoilsResponse.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x05)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = ModbusPDUWriteSingleCoilRequest.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x05)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = ModbusPDUWriteSingleCoilResponse.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x0F)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder =
          ModbusPDUWriteMultipleCoilsRequest.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x0F)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder =
          ModbusPDUWriteMultipleCoilsResponse.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x04)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder =
          ModbusPDUReadInputRegistersRequest.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x04)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder =
          ModbusPDUReadInputRegistersResponse.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x03)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder =
          ModbusPDUReadHoldingRegistersRequest.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x03)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder =
          ModbusPDUReadHoldingRegistersResponse.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x06)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder =
          ModbusPDUWriteSingleRegisterRequest.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x06)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder =
          ModbusPDUWriteSingleRegisterResponse.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x10)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder =
          ModbusPDUWriteMultipleHoldingRegistersRequest.staticParseModbusPDUBuilder(
              readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x10)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder =
          ModbusPDUWriteMultipleHoldingRegistersResponse.staticParseModbusPDUBuilder(
              readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x17)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder =
          ModbusPDUReadWriteMultipleHoldingRegistersRequest.staticParseModbusPDUBuilder(
              readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x17)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder =
          ModbusPDUReadWriteMultipleHoldingRegistersResponse.staticParseModbusPDUBuilder(
              readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x16)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder =
          ModbusPDUMaskWriteHoldingRegisterRequest.staticParseModbusPDUBuilder(
              readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x16)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder =
          ModbusPDUMaskWriteHoldingRegisterResponse.staticParseModbusPDUBuilder(
              readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x18)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = ModbusPDUReadFifoQueueRequest.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x18)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = ModbusPDUReadFifoQueueResponse.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x14)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = ModbusPDUReadFileRecordRequest.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x14)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = ModbusPDUReadFileRecordResponse.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x15)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = ModbusPDUWriteFileRecordRequest.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x15)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = ModbusPDUWriteFileRecordResponse.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x07)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder =
          ModbusPDUReadExceptionStatusRequest.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x07)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder =
          ModbusPDUReadExceptionStatusResponse.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x08)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = ModbusPDUDiagnosticRequest.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x08)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = ModbusPDUDiagnosticResponse.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x0B)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder =
          ModbusPDUGetComEventCounterRequest.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x0B)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder =
          ModbusPDUGetComEventCounterResponse.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x0C)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = ModbusPDUGetComEventLogRequest.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x0C)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = ModbusPDUGetComEventLogResponse.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x11)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = ModbusPDUReportServerIdRequest.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x11)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = ModbusPDUReportServerIdResponse.staticParseModbusPDUBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x2B)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder =
          ModbusPDUReadDeviceIdentificationRequest.staticParseModbusPDUBuilder(
              readBuffer, response);
    } else if (EvaluationHelper.equals(errorFlag, (boolean) false)
        && EvaluationHelper.equals(functionFlag, (byte) 0x2B)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder =
          ModbusPDUReadDeviceIdentificationResponse.staticParseModbusPDUBuilder(
              readBuffer, response);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "errorFlag="
              + errorFlag
              + " "
              + "functionFlag="
              + functionFlag
              + " "
              + "response="
              + response
              + "]");
    }

    readBuffer.closeContext("ModbusPDU");
    // Create the instance
    ModbusPDU _modbusPDU = builder.build();
    return _modbusPDU;
  }

  public interface ModbusPDUBuilder {
    ModbusPDU build();
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ModbusPDU)) {
      return false;
    }
    ModbusPDU that = (ModbusPDU) o;
    return true;
  }

  @Override
  public int hashCode() {
    return Objects.hash();
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