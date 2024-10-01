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

public class MultistateInput implements Message {

  // Properties.
  protected final ReadableProperty objectIdentifier;
  protected final ReadableProperty objectName;
  protected final ReadableProperty objectType;
  protected final ReadableProperty presentValue;
  protected final OptionalProperty description;
  protected final OptionalProperty deviceType;
  protected final ReadableProperty statusFlags;
  protected final ReadableProperty eventState;
  protected final OptionalProperty reliability;
  protected final ReadableProperty outOfService;
  protected final ReadableProperty numberOfStates;
  protected final OptionalProperty stateText;
  protected final OptionalProperty timeDelay;
  protected final OptionalProperty notificationClass;
  protected final OptionalProperty alarmValues;
  protected final OptionalProperty faultValues;
  protected final OptionalProperty eventEnable;
  protected final OptionalProperty ackedTransitions;
  protected final OptionalProperty notifyType;
  protected final OptionalProperty eventTimeStamps;
  protected final OptionalProperty eventMessageTexts;
  protected final OptionalProperty eventMessageTextsConfig;
  protected final OptionalProperty eventDetectionEnable;
  protected final OptionalProperty eventAlgorithmInhibitRef;
  protected final OptionalProperty eventAlgorithmInhibit;
  protected final OptionalProperty timeDelayNormal;
  protected final OptionalProperty reliabilityEvaluationInhibit;
  protected final ReadableProperty propertyList;
  protected final OptionalProperty interfaceValue;
  protected final OptionalProperty tags;
  protected final OptionalProperty profileLocation;
  protected final OptionalProperty profileName;

  public MultistateInput(
      ReadableProperty objectIdentifier,
      ReadableProperty objectName,
      ReadableProperty objectType,
      ReadableProperty presentValue,
      OptionalProperty description,
      OptionalProperty deviceType,
      ReadableProperty statusFlags,
      ReadableProperty eventState,
      OptionalProperty reliability,
      ReadableProperty outOfService,
      ReadableProperty numberOfStates,
      OptionalProperty stateText,
      OptionalProperty timeDelay,
      OptionalProperty notificationClass,
      OptionalProperty alarmValues,
      OptionalProperty faultValues,
      OptionalProperty eventEnable,
      OptionalProperty ackedTransitions,
      OptionalProperty notifyType,
      OptionalProperty eventTimeStamps,
      OptionalProperty eventMessageTexts,
      OptionalProperty eventMessageTextsConfig,
      OptionalProperty eventDetectionEnable,
      OptionalProperty eventAlgorithmInhibitRef,
      OptionalProperty eventAlgorithmInhibit,
      OptionalProperty timeDelayNormal,
      OptionalProperty reliabilityEvaluationInhibit,
      ReadableProperty propertyList,
      OptionalProperty interfaceValue,
      OptionalProperty tags,
      OptionalProperty profileLocation,
      OptionalProperty profileName) {
    super();
    this.objectIdentifier = objectIdentifier;
    this.objectName = objectName;
    this.objectType = objectType;
    this.presentValue = presentValue;
    this.description = description;
    this.deviceType = deviceType;
    this.statusFlags = statusFlags;
    this.eventState = eventState;
    this.reliability = reliability;
    this.outOfService = outOfService;
    this.numberOfStates = numberOfStates;
    this.stateText = stateText;
    this.timeDelay = timeDelay;
    this.notificationClass = notificationClass;
    this.alarmValues = alarmValues;
    this.faultValues = faultValues;
    this.eventEnable = eventEnable;
    this.ackedTransitions = ackedTransitions;
    this.notifyType = notifyType;
    this.eventTimeStamps = eventTimeStamps;
    this.eventMessageTexts = eventMessageTexts;
    this.eventMessageTextsConfig = eventMessageTextsConfig;
    this.eventDetectionEnable = eventDetectionEnable;
    this.eventAlgorithmInhibitRef = eventAlgorithmInhibitRef;
    this.eventAlgorithmInhibit = eventAlgorithmInhibit;
    this.timeDelayNormal = timeDelayNormal;
    this.reliabilityEvaluationInhibit = reliabilityEvaluationInhibit;
    this.propertyList = propertyList;
    this.interfaceValue = interfaceValue;
    this.tags = tags;
    this.profileLocation = profileLocation;
    this.profileName = profileName;
  }

  public ReadableProperty getObjectIdentifier() {
    return objectIdentifier;
  }

  public ReadableProperty getObjectName() {
    return objectName;
  }

  public ReadableProperty getObjectType() {
    return objectType;
  }

  public ReadableProperty getPresentValue() {
    return presentValue;
  }

  public OptionalProperty getDescription() {
    return description;
  }

  public OptionalProperty getDeviceType() {
    return deviceType;
  }

  public ReadableProperty getStatusFlags() {
    return statusFlags;
  }

  public ReadableProperty getEventState() {
    return eventState;
  }

  public OptionalProperty getReliability() {
    return reliability;
  }

  public ReadableProperty getOutOfService() {
    return outOfService;
  }

  public ReadableProperty getNumberOfStates() {
    return numberOfStates;
  }

  public OptionalProperty getStateText() {
    return stateText;
  }

  public OptionalProperty getTimeDelay() {
    return timeDelay;
  }

  public OptionalProperty getNotificationClass() {
    return notificationClass;
  }

  public OptionalProperty getAlarmValues() {
    return alarmValues;
  }

  public OptionalProperty getFaultValues() {
    return faultValues;
  }

  public OptionalProperty getEventEnable() {
    return eventEnable;
  }

  public OptionalProperty getAckedTransitions() {
    return ackedTransitions;
  }

  public OptionalProperty getNotifyType() {
    return notifyType;
  }

  public OptionalProperty getEventTimeStamps() {
    return eventTimeStamps;
  }

  public OptionalProperty getEventMessageTexts() {
    return eventMessageTexts;
  }

  public OptionalProperty getEventMessageTextsConfig() {
    return eventMessageTextsConfig;
  }

  public OptionalProperty getEventDetectionEnable() {
    return eventDetectionEnable;
  }

  public OptionalProperty getEventAlgorithmInhibitRef() {
    return eventAlgorithmInhibitRef;
  }

  public OptionalProperty getEventAlgorithmInhibit() {
    return eventAlgorithmInhibit;
  }

  public OptionalProperty getTimeDelayNormal() {
    return timeDelayNormal;
  }

  public OptionalProperty getReliabilityEvaluationInhibit() {
    return reliabilityEvaluationInhibit;
  }

  public ReadableProperty getPropertyList() {
    return propertyList;
  }

  public OptionalProperty getInterfaceValue() {
    return interfaceValue;
  }

  public OptionalProperty getTags() {
    return tags;
  }

  public OptionalProperty getProfileLocation() {
    return profileLocation;
  }

  public OptionalProperty getProfileName() {
    return profileName;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("MultistateInput");

    // Simple Field (objectIdentifier)
    writeSimpleField("objectIdentifier", objectIdentifier, writeComplex(writeBuffer));

    // Simple Field (objectName)
    writeSimpleField("objectName", objectName, writeComplex(writeBuffer));

    // Simple Field (objectType)
    writeSimpleField("objectType", objectType, writeComplex(writeBuffer));

    // Simple Field (presentValue)
    writeSimpleField("presentValue", presentValue, writeComplex(writeBuffer));

    // Simple Field (description)
    writeSimpleField("description", description, writeComplex(writeBuffer));

    // Simple Field (deviceType)
    writeSimpleField("deviceType", deviceType, writeComplex(writeBuffer));

    // Simple Field (statusFlags)
    writeSimpleField("statusFlags", statusFlags, writeComplex(writeBuffer));

    // Simple Field (eventState)
    writeSimpleField("eventState", eventState, writeComplex(writeBuffer));

    // Simple Field (reliability)
    writeSimpleField("reliability", reliability, writeComplex(writeBuffer));

    // Simple Field (outOfService)
    writeSimpleField("outOfService", outOfService, writeComplex(writeBuffer));

    // Simple Field (numberOfStates)
    writeSimpleField("numberOfStates", numberOfStates, writeComplex(writeBuffer));

    // Simple Field (stateText)
    writeSimpleField("stateText", stateText, writeComplex(writeBuffer));

    // Simple Field (timeDelay)
    writeSimpleField("timeDelay", timeDelay, writeComplex(writeBuffer));

    // Simple Field (notificationClass)
    writeSimpleField("notificationClass", notificationClass, writeComplex(writeBuffer));

    // Simple Field (alarmValues)
    writeSimpleField("alarmValues", alarmValues, writeComplex(writeBuffer));

    // Simple Field (faultValues)
    writeSimpleField("faultValues", faultValues, writeComplex(writeBuffer));

    // Simple Field (eventEnable)
    writeSimpleField("eventEnable", eventEnable, writeComplex(writeBuffer));

    // Simple Field (ackedTransitions)
    writeSimpleField("ackedTransitions", ackedTransitions, writeComplex(writeBuffer));

    // Simple Field (notifyType)
    writeSimpleField("notifyType", notifyType, writeComplex(writeBuffer));

    // Simple Field (eventTimeStamps)
    writeSimpleField("eventTimeStamps", eventTimeStamps, writeComplex(writeBuffer));

    // Simple Field (eventMessageTexts)
    writeSimpleField("eventMessageTexts", eventMessageTexts, writeComplex(writeBuffer));

    // Simple Field (eventMessageTextsConfig)
    writeSimpleField("eventMessageTextsConfig", eventMessageTextsConfig, writeComplex(writeBuffer));

    // Simple Field (eventDetectionEnable)
    writeSimpleField("eventDetectionEnable", eventDetectionEnable, writeComplex(writeBuffer));

    // Simple Field (eventAlgorithmInhibitRef)
    writeSimpleField(
        "eventAlgorithmInhibitRef", eventAlgorithmInhibitRef, writeComplex(writeBuffer));

    // Simple Field (eventAlgorithmInhibit)
    writeSimpleField("eventAlgorithmInhibit", eventAlgorithmInhibit, writeComplex(writeBuffer));

    // Simple Field (timeDelayNormal)
    writeSimpleField("timeDelayNormal", timeDelayNormal, writeComplex(writeBuffer));

    // Simple Field (reliabilityEvaluationInhibit)
    writeSimpleField(
        "reliabilityEvaluationInhibit", reliabilityEvaluationInhibit, writeComplex(writeBuffer));

    // Simple Field (propertyList)
    writeSimpleField("propertyList", propertyList, writeComplex(writeBuffer));

    // Simple Field (interfaceValue)
    writeSimpleField("interfaceValue", interfaceValue, writeComplex(writeBuffer));

    // Simple Field (tags)
    writeSimpleField("tags", tags, writeComplex(writeBuffer));

    // Simple Field (profileLocation)
    writeSimpleField("profileLocation", profileLocation, writeComplex(writeBuffer));

    // Simple Field (profileName)
    writeSimpleField("profileName", profileName, writeComplex(writeBuffer));

    writeBuffer.popContext("MultistateInput");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    MultistateInput _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (objectIdentifier)
    lengthInBits += objectIdentifier.getLengthInBits();

    // Simple field (objectName)
    lengthInBits += objectName.getLengthInBits();

    // Simple field (objectType)
    lengthInBits += objectType.getLengthInBits();

    // Simple field (presentValue)
    lengthInBits += presentValue.getLengthInBits();

    // Simple field (description)
    lengthInBits += description.getLengthInBits();

    // Simple field (deviceType)
    lengthInBits += deviceType.getLengthInBits();

    // Simple field (statusFlags)
    lengthInBits += statusFlags.getLengthInBits();

    // Simple field (eventState)
    lengthInBits += eventState.getLengthInBits();

    // Simple field (reliability)
    lengthInBits += reliability.getLengthInBits();

    // Simple field (outOfService)
    lengthInBits += outOfService.getLengthInBits();

    // Simple field (numberOfStates)
    lengthInBits += numberOfStates.getLengthInBits();

    // Simple field (stateText)
    lengthInBits += stateText.getLengthInBits();

    // Simple field (timeDelay)
    lengthInBits += timeDelay.getLengthInBits();

    // Simple field (notificationClass)
    lengthInBits += notificationClass.getLengthInBits();

    // Simple field (alarmValues)
    lengthInBits += alarmValues.getLengthInBits();

    // Simple field (faultValues)
    lengthInBits += faultValues.getLengthInBits();

    // Simple field (eventEnable)
    lengthInBits += eventEnable.getLengthInBits();

    // Simple field (ackedTransitions)
    lengthInBits += ackedTransitions.getLengthInBits();

    // Simple field (notifyType)
    lengthInBits += notifyType.getLengthInBits();

    // Simple field (eventTimeStamps)
    lengthInBits += eventTimeStamps.getLengthInBits();

    // Simple field (eventMessageTexts)
    lengthInBits += eventMessageTexts.getLengthInBits();

    // Simple field (eventMessageTextsConfig)
    lengthInBits += eventMessageTextsConfig.getLengthInBits();

    // Simple field (eventDetectionEnable)
    lengthInBits += eventDetectionEnable.getLengthInBits();

    // Simple field (eventAlgorithmInhibitRef)
    lengthInBits += eventAlgorithmInhibitRef.getLengthInBits();

    // Simple field (eventAlgorithmInhibit)
    lengthInBits += eventAlgorithmInhibit.getLengthInBits();

    // Simple field (timeDelayNormal)
    lengthInBits += timeDelayNormal.getLengthInBits();

    // Simple field (reliabilityEvaluationInhibit)
    lengthInBits += reliabilityEvaluationInhibit.getLengthInBits();

    // Simple field (propertyList)
    lengthInBits += propertyList.getLengthInBits();

    // Simple field (interfaceValue)
    lengthInBits += interfaceValue.getLengthInBits();

    // Simple field (tags)
    lengthInBits += tags.getLengthInBits();

    // Simple field (profileLocation)
    lengthInBits += profileLocation.getLengthInBits();

    // Simple field (profileName)
    lengthInBits += profileName.getLengthInBits();

    return lengthInBits;
  }

  public static MultistateInput staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("MultistateInput");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ReadableProperty objectIdentifier =
        readSimpleField(
            "objectIdentifier",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetObjectIdentifier")),
                readBuffer));

    ReadableProperty objectName =
        readSimpleField(
            "objectName",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("CharacterString")),
                readBuffer));

    ReadableProperty objectType =
        readSimpleField(
            "objectType",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetObjectType")),
                readBuffer));

    ReadableProperty presentValue =
        readSimpleField(
            "presentValue",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty description =
        readSimpleField(
            "description",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("CharacterString")),
                readBuffer));

    OptionalProperty deviceType =
        readSimpleField(
            "deviceType",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("CharacterString")),
                readBuffer));

    ReadableProperty statusFlags =
        readSimpleField(
            "statusFlags",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetStatusFlags")),
                readBuffer));

    ReadableProperty eventState =
        readSimpleField(
            "eventState",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetEventState")),
                readBuffer));

    OptionalProperty reliability =
        readSimpleField(
            "reliability",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BACnetReliability")),
                readBuffer));

    ReadableProperty outOfService =
        readSimpleField(
            "outOfService",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    ReadableProperty numberOfStates =
        readSimpleField(
            "numberOfStates",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty stateText =
        readSimpleField(
            "stateText",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[N] of CharacterString")),
                readBuffer));

    OptionalProperty timeDelay =
        readSimpleField(
            "timeDelay",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty notificationClass =
        readSimpleField(
            "notificationClass",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty alarmValues =
        readSimpleField(
            "alarmValues",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BACnetLIST of Unsigned")),
                readBuffer));

    OptionalProperty faultValues =
        readSimpleField(
            "faultValues",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BACnetLIST of Unsigned")),
                readBuffer));

    OptionalProperty eventEnable =
        readSimpleField(
            "eventEnable",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetEventTransitionBits")),
                readBuffer));

    OptionalProperty ackedTransitions =
        readSimpleField(
            "ackedTransitions",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetEventTransitionBits")),
                readBuffer));

    OptionalProperty notifyType =
        readSimpleField(
            "notifyType",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BACnetNotifyType")),
                readBuffer));

    OptionalProperty eventTimeStamps =
        readSimpleField(
            "eventTimeStamps",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[3] of BACnetTimeStamp")),
                readBuffer));

    OptionalProperty eventMessageTexts =
        readSimpleField(
            "eventMessageTexts",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[3] of CharacterString")),
                readBuffer));

    OptionalProperty eventMessageTextsConfig =
        readSimpleField(
            "eventMessageTextsConfig",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[3] of CharacterString")),
                readBuffer));

    OptionalProperty eventDetectionEnable =
        readSimpleField(
            "eventDetectionEnable",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    OptionalProperty eventAlgorithmInhibitRef =
        readSimpleField(
            "eventAlgorithmInhibitRef",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetObjectPropertyReference")),
                readBuffer));

    OptionalProperty eventAlgorithmInhibit =
        readSimpleField(
            "eventAlgorithmInhibit",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    OptionalProperty timeDelayNormal =
        readSimpleField(
            "timeDelayNormal",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty reliabilityEvaluationInhibit =
        readSimpleField(
            "reliabilityEvaluationInhibit",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    ReadableProperty propertyList =
        readSimpleField(
            "propertyList",
            readComplex(
                () ->
                    ReadableProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[N] of BACnetPropertyIdentifier")),
                readBuffer));

    OptionalProperty interfaceValue =
        readSimpleField(
            "interfaceValue",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BACnetOptionalUnsigned")),
                readBuffer));

    OptionalProperty tags =
        readSimpleField(
            "tags",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[N] of BACnetNameValue")),
                readBuffer));

    OptionalProperty profileLocation =
        readSimpleField(
            "profileLocation",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("CharacterString")),
                readBuffer));

    OptionalProperty profileName =
        readSimpleField(
            "profileName",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("CharacterString")),
                readBuffer));

    readBuffer.closeContext("MultistateInput");
    // Create the instance
    MultistateInput _multistateInput;
    _multistateInput =
        new MultistateInput(
            objectIdentifier,
            objectName,
            objectType,
            presentValue,
            description,
            deviceType,
            statusFlags,
            eventState,
            reliability,
            outOfService,
            numberOfStates,
            stateText,
            timeDelay,
            notificationClass,
            alarmValues,
            faultValues,
            eventEnable,
            ackedTransitions,
            notifyType,
            eventTimeStamps,
            eventMessageTexts,
            eventMessageTextsConfig,
            eventDetectionEnable,
            eventAlgorithmInhibitRef,
            eventAlgorithmInhibit,
            timeDelayNormal,
            reliabilityEvaluationInhibit,
            propertyList,
            interfaceValue,
            tags,
            profileLocation,
            profileName);
    return _multistateInput;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof MultistateInput)) {
      return false;
    }
    MultistateInput that = (MultistateInput) o;
    return (getObjectIdentifier() == that.getObjectIdentifier())
        && (getObjectName() == that.getObjectName())
        && (getObjectType() == that.getObjectType())
        && (getPresentValue() == that.getPresentValue())
        && (getDescription() == that.getDescription())
        && (getDeviceType() == that.getDeviceType())
        && (getStatusFlags() == that.getStatusFlags())
        && (getEventState() == that.getEventState())
        && (getReliability() == that.getReliability())
        && (getOutOfService() == that.getOutOfService())
        && (getNumberOfStates() == that.getNumberOfStates())
        && (getStateText() == that.getStateText())
        && (getTimeDelay() == that.getTimeDelay())
        && (getNotificationClass() == that.getNotificationClass())
        && (getAlarmValues() == that.getAlarmValues())
        && (getFaultValues() == that.getFaultValues())
        && (getEventEnable() == that.getEventEnable())
        && (getAckedTransitions() == that.getAckedTransitions())
        && (getNotifyType() == that.getNotifyType())
        && (getEventTimeStamps() == that.getEventTimeStamps())
        && (getEventMessageTexts() == that.getEventMessageTexts())
        && (getEventMessageTextsConfig() == that.getEventMessageTextsConfig())
        && (getEventDetectionEnable() == that.getEventDetectionEnable())
        && (getEventAlgorithmInhibitRef() == that.getEventAlgorithmInhibitRef())
        && (getEventAlgorithmInhibit() == that.getEventAlgorithmInhibit())
        && (getTimeDelayNormal() == that.getTimeDelayNormal())
        && (getReliabilityEvaluationInhibit() == that.getReliabilityEvaluationInhibit())
        && (getPropertyList() == that.getPropertyList())
        && (getInterfaceValue() == that.getInterfaceValue())
        && (getTags() == that.getTags())
        && (getProfileLocation() == that.getProfileLocation())
        && (getProfileName() == that.getProfileName())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getObjectIdentifier(),
        getObjectName(),
        getObjectType(),
        getPresentValue(),
        getDescription(),
        getDeviceType(),
        getStatusFlags(),
        getEventState(),
        getReliability(),
        getOutOfService(),
        getNumberOfStates(),
        getStateText(),
        getTimeDelay(),
        getNotificationClass(),
        getAlarmValues(),
        getFaultValues(),
        getEventEnable(),
        getAckedTransitions(),
        getNotifyType(),
        getEventTimeStamps(),
        getEventMessageTexts(),
        getEventMessageTextsConfig(),
        getEventDetectionEnable(),
        getEventAlgorithmInhibitRef(),
        getEventAlgorithmInhibit(),
        getTimeDelayNormal(),
        getReliabilityEvaluationInhibit(),
        getPropertyList(),
        getInterfaceValue(),
        getTags(),
        getProfileLocation(),
        getProfileName());
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
