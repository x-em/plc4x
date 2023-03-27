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
package org.apache.plc4x.java.s7.readwrite;

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

public class DateAndTime implements Message {

  // Properties.
  protected final short year;
  protected final short month;
  protected final short day;
  protected final short hour;
  protected final short minutes;
  protected final short seconds;
  protected final int msec;
  protected final byte dow;

  public DateAndTime(
      short year,
      short month,
      short day,
      short hour,
      short minutes,
      short seconds,
      int msec,
      byte dow) {
    super();
    this.year = year;
    this.month = month;
    this.day = day;
    this.hour = hour;
    this.minutes = minutes;
    this.seconds = seconds;
    this.msec = msec;
    this.dow = dow;
  }

  public short getYear() {
    return year;
  }

  public short getMonth() {
    return month;
  }

  public short getDay() {
    return day;
  }

  public short getHour() {
    return hour;
  }

  public short getMinutes() {
    return minutes;
  }

  public short getSeconds() {
    return seconds;
  }

  public int getMsec() {
    return msec;
  }

  public byte getDow() {
    return dow;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("DateAndTime");

    // Manual Field (year)
    writeManualField(
        "year",
        () -> org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.ByteToBcd(writeBuffer, year),
        writeBuffer);

    // Manual Field (month)
    writeManualField(
        "month",
        () -> org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.ByteToBcd(writeBuffer, month),
        writeBuffer);

    // Manual Field (day)
    writeManualField(
        "day",
        () -> org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.ByteToBcd(writeBuffer, day),
        writeBuffer);

    // Manual Field (hour)
    writeManualField(
        "hour",
        () -> org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.ByteToBcd(writeBuffer, hour),
        writeBuffer);

    // Manual Field (minutes)
    writeManualField(
        "minutes",
        () -> org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.ByteToBcd(writeBuffer, minutes),
        writeBuffer);

    // Manual Field (seconds)
    writeManualField(
        "seconds",
        () -> org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.ByteToBcd(writeBuffer, seconds),
        writeBuffer);

    // Manual Field (msec)
    writeManualField(
        "msec",
        () -> org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.IntToS7msec(writeBuffer, msec),
        writeBuffer);

    // Simple Field (dow)
    writeSimpleField("dow", dow, writeUnsignedByte(writeBuffer, 4));

    writeBuffer.popContext("DateAndTime");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    DateAndTime _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Manual Field (year)
    lengthInBits += 8;

    // Manual Field (month)
    lengthInBits += 8;

    // Manual Field (day)
    lengthInBits += 8;

    // Manual Field (hour)
    lengthInBits += 8;

    // Manual Field (minutes)
    lengthInBits += 8;

    // Manual Field (seconds)
    lengthInBits += 8;

    // Manual Field (msec)
    lengthInBits += 12;

    // Simple field (dow)
    lengthInBits += 4;

    return lengthInBits;
  }

  public static DateAndTime staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static DateAndTime staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("DateAndTime");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short year =
        readManualField(
            "year",
            readBuffer,
            () ->
                (short)
                    (org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.BcdToInt(readBuffer)));

    short month =
        readManualField(
            "month",
            readBuffer,
            () ->
                (short)
                    (org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.BcdToInt(readBuffer)));

    short day =
        readManualField(
            "day",
            readBuffer,
            () ->
                (short)
                    (org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.BcdToInt(readBuffer)));

    short hour =
        readManualField(
            "hour",
            readBuffer,
            () ->
                (short)
                    (org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.BcdToInt(readBuffer)));

    short minutes =
        readManualField(
            "minutes",
            readBuffer,
            () ->
                (short)
                    (org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.BcdToInt(readBuffer)));

    short seconds =
        readManualField(
            "seconds",
            readBuffer,
            () ->
                (short)
                    (org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.BcdToInt(readBuffer)));

    int msec =
        readManualField(
            "msec",
            readBuffer,
            () ->
                (int)
                    (org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.S7msecToInt(
                        readBuffer)));

    byte dow = readSimpleField("dow", readUnsignedByte(readBuffer, 4));

    readBuffer.closeContext("DateAndTime");
    // Create the instance
    DateAndTime _dateAndTime;
    _dateAndTime = new DateAndTime(year, month, day, hour, minutes, seconds, msec, dow);
    return _dateAndTime;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DateAndTime)) {
      return false;
    }
    DateAndTime that = (DateAndTime) o;
    return (getYear() == that.getYear())
        && (getMonth() == that.getMonth())
        && (getDay() == that.getDay())
        && (getHour() == that.getHour())
        && (getMinutes() == that.getMinutes())
        && (getSeconds() == that.getSeconds())
        && (getMsec() == that.getMsec())
        && (getDow() == that.getDow())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getYear(),
        getMonth(),
        getDay(),
        getHour(),
        getMinutes(),
        getSeconds(),
        getMsec(),
        getDow());
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
