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

package object

import (
	"github.com/pkg/errors"

	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/basetypes"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/comp"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/primitivedata"
)

type PulseConverterObject struct {
	Object
}

func NewPulseConverterObject(options ...Option) (*PulseConverterObject, error) {
	p := new(PulseConverterObject)
	objectType := "pulseConverter"
	_object_supports_cov := true
	properties := []Property{
		NewReadableProperty("presentValue", V2P(NewReal)),
		NewOptionalProperty("inputReference", V2P(NewObjectPropertyReference)),
		NewReadableProperty("statusFlags", V2P(NewStatusFlags)),
		NewReadableProperty("eventState", V2P(NewEventState)),
		NewOptionalProperty("reliability", V2P(NewReliability)),
		NewReadableProperty("outOfService", V2P(NewBoolean)),
		NewReadableProperty("units", V2P(NewEngineeringUnits)),
		NewReadableProperty("scaleFactor", V2P(NewReal)),
		NewWritableProperty("adjustValue", V2P(NewReal)),
		NewReadableProperty("count", V2P(NewUnsigned)),
		NewReadableProperty("updateTime", V2P(NewDateTime)),
		NewReadableProperty("countChangeTime", V2P(NewDateTime)),
		NewReadableProperty("countBeforeChange", V2P(NewUnsigned)),
		NewOptionalProperty("covIncrement", V2P(NewReal)),
		NewOptionalProperty("covPeriod", V2P(NewUnsigned)),
		NewOptionalProperty("notificationClass", V2P(NewUnsigned)),
		NewOptionalProperty("timeDelay", V2P(NewUnsigned)),
		NewOptionalProperty("highLimit", V2P(NewReal)),
		NewOptionalProperty("lowLimit", V2P(NewReal)),
		NewOptionalProperty("deadband", V2P(NewReal)),
		NewOptionalProperty("limitEnable", V2P(NewLimitEnable)),
		NewOptionalProperty("eventEnable", V2P(NewEventTransitionBits)),
		NewOptionalProperty("ackedTransitions", V2P(NewEventTransitionBits)),
		NewOptionalProperty("notifyType", V2P(NewNotifyType)),
		NewOptionalProperty("eventTimeStamps", ArrayOfP(NewTimeStamp, 3, 0)),
		NewOptionalProperty("eventMessageTexts", ArrayOfP(NewCharacterString, 3, 0)),
		NewOptionalProperty("eventMessageTextsConfig", ArrayOfP(NewCharacterString, 3, 0)),
		NewOptionalProperty("eventDetectionEnable", V2P(NewBoolean)),
		NewOptionalProperty("eventAlgorithmInhibitRef", V2P(NewObjectPropertyReference)),
		NewOptionalProperty("eventAlgorithmInhibit", V2P(NewBoolean)),
		NewOptionalProperty("timeDelayNormal", V2P(NewUnsigned)),
		NewOptionalProperty("reliabilityEvaluationInhibit", V2P(NewBoolean)),
	}
	var err error
	p.Object, err = NewObject(Combine(options, WithObjectType(objectType), WithObjectExtraProperties(properties), WithObjectSupportsCov(_object_supports_cov))...)
	if err != nil {
		return nil, errors.Wrap(err, "error creating object")
	}
	if _, err := RegisterObjectType(NKW(KWCls, p)); err != nil {
		return nil, errors.Wrap(err, "error registering object type")
	}
	return p, nil
}