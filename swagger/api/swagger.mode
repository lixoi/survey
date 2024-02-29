


# api/api.proto
  

## Informations

### Version

version not set

## Tags

  ### <span id="tag-i-c-h-survey"></span>ICHSurvey

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json

## All endpoints

###  i_c_h_survey

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v1/set_answer/{userId} | [i c h survey set answer](#i-c-h-survey-set-answer) |  |
| GET | /v1/start_survey/{userId} | [i c h survey start survey](#i-c-h-survey-start-survey) |  |
  


## Paths

### <span id="i-c-h-survey-set-answer"></span> i c h survey set answer (*ICHSurvey_SetAnswer*)

```
POST /v1/set_answer/{userId}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| userId | `path` | uint64 (formatted string) | `string` |  | ✓ |  |  |
| body | `body` | [ICHSurveySetAnswerBody](#i-c-h-survey-set-answer-body) | `models.ICHSurveySetAnswerBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#i-c-h-survey-set-answer-200) | OK | A successful response. |  | [schema](#i-c-h-survey-set-answer-200-schema) |
| [default](#i-c-h-survey-set-answer-default) | | An unexpected error response. |  | [schema](#i-c-h-survey-set-answer-default-schema) |

#### Responses


##### <span id="i-c-h-survey-set-answer-200"></span> 200 - A successful response.
Status: OK

###### <span id="i-c-h-survey-set-answer-200-schema"></span> Schema
   
  

[APIQuestionResponse](#api-question-response)

##### <span id="i-c-h-survey-set-answer-default"></span> Default Response
An unexpected error response.

###### <span id="i-c-h-survey-set-answer-default-schema"></span> Schema

  

[RPCStatus](#rpc-status)

### <span id="i-c-h-survey-start-survey"></span> i c h survey start survey (*ICHSurvey_StartSurvey*)

```
GET /v1/start_survey/{userId}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| userId | `path` | uint64 (formatted string) | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#i-c-h-survey-start-survey-200) | OK | A successful response. |  | [schema](#i-c-h-survey-start-survey-200-schema) |
| [default](#i-c-h-survey-start-survey-default) | | An unexpected error response. |  | [schema](#i-c-h-survey-start-survey-default-schema) |

#### Responses


##### <span id="i-c-h-survey-start-survey-200"></span> 200 - A successful response.
Status: OK

###### <span id="i-c-h-survey-start-survey-200-schema"></span> Schema
   
  

[APIQuestionResponse](#api-question-response)

##### <span id="i-c-h-survey-start-survey-default"></span> Default Response
An unexpected error response.

###### <span id="i-c-h-survey-start-survey-default-schema"></span> Schema

  

[RPCStatus](#rpc-status)

## Models

### <span id="i-c-h-survey-set-answer-body"></span> ICHSurveySetAnswerBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| answer | string| `string` |  | |  |  |
| number | int64 (formatted integer)| `int64` |  | |  |  |



### <span id="api-class-questions"></span> apiClassQuestions


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| apiClassQuestions | string| string | `"UNKNOWN_QUESTIONS_CLASS"`|  |  |



### <span id="api-question-response"></span> apiQuestionResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| message | string| `string` |  | |  |  |
| number | int64 (formatted integer)| `int64` |  | |  |  |
| question | string| `string` |  | |  |  |
| userId | uint64 (formatted string)| `string` |  | |  |  |



### <span id="api-status-response"></span> apiStatusResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| message | string| `string` |  | |  |  |



### <span id="api-survey"></span> apiSurvey


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| answer | string| `string` |  | |  |  |
| latency | string| `string` |  | |  |  |
| number | int64 (formatted integer)| `int64` |  | |  |  |
| question | string| `string` |  | |  |  |
| title | string| `string` |  | |  |  |
| userId | uint64 (formatted string)| `string` |  | |  |  |



### <span id="api-survey-response"></span> apiSurveyResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| mesage | string| `string` |  | |  |  |
| qs | [][APISurvey](#api-survey)| `[]*APISurvey` |  | |  |  |
| startSurvey | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |



### <span id="protobuf-any"></span> protobufAny


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| @type | string| `string` |  | |  |  |



### <span id="rpc-status"></span> rpcStatus


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| code | int32 (formatted integer)| `int32` |  | |  |  |
| details | [][ProtobufAny](#protobuf-any)| `[]*ProtobufAny` |  | |  |  |
| message | string| `string` |  | |  |  |


