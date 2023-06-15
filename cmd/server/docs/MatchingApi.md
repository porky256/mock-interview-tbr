# {{classname}}

All URIs are relative to */*

| Method                                                        | HTTP request                              | Description               |
|---------------------------------------------------------------|-------------------------------------------|---------------------------|
| [**GetMatch**](MatchingApi.md#getmatch)                       | **Get** /matching/match/{matchID}         | Get a match               |
| [**GetMatchByUserID**](MatchingApi.md#getmatchbyuserid)       | **Get** /matching/match/findByUserID      | Get a match               |
| [**GetInterviewByID**](MatchingApi.md#getinterviewbyid)       | **Get** /matching/interview/{interviewID} | Get a interview           |
| [**GetMatchByID**](MatchingApi.md#getmatchbyid)               | **Post** /matching/match/{matchID}        | Approve a match           |
| [**PostMatchingRequest**](MatchingApi.md#postmatchingrequest) | **Post** /matching/request                | Create a matching request |

# **GetMatch**
> Match GetMatch(ctx, matchID)
Get a match

This can only be done by the logged-in user.

### Required Parameters

| Name        | Type                | Description                                                                 | Notes |
|-------------|---------------------|-----------------------------------------------------------------------------|-------|
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **matchID** | **int64**           | ID of match to return                                                       |       |

### Return type

[**Match**](Match.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMatchByUserID**
> []Match GetMatchByUserID(ctx, userID)
Get a match

This can only be done by the logged-in user.

### Required Parameters

| Name       | Type                | Description                                                                 | Notes |
|------------|---------------------|-----------------------------------------------------------------------------|-------|
| **ctx**    | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **userID** | **int64**           | ID of user                                                                  | 

### Return type

[**[]Match**](Match.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInterviewByID**
> Interview GetInterviewByID(ctx, interviewID)
Get interview

This can only be done by the logged-in user.

### Required Parameters

| Name            | Type                | Description                                                                 | Notes |
|-----------------|---------------------|-----------------------------------------------------------------------------|-------|
| **ctx**         | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **interviewID** | **int64**           | ID of interview to return                                                   | 

### Return type

[**Interview**](Interview.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMatchByID**
> GetMatchByID(ctx, matchID, date)
Approve a match

This can only be done by the logged-in user.

### Required Parameters

| Name        | Type                | Description                                                                 | Notes |
|-------------|---------------------|-----------------------------------------------------------------------------|-------|
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **matchID** | **int64**           | ID of match to approve                                                      |
| **date**    | **string**          | date of interview                                                           | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMatchingRequest**
> Match PostMatchingRequest(ctx, skillID, userID, requestedSkills)
Create a matching request

This can only be done by the logged-in user.

### Required Parameters

| Name                | Type                    | Description                                                                 | Notes |
|---------------------|-------------------------|-----------------------------------------------------------------------------|-------|
| **ctx**             | **context.Context**     | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **skillID**         | **int64**               | ID of skill                                                                 |
| **userID**          | **int64**               | ID of user                                                                  |
| **requestedSkills** | [**[]Skill**](Skill.md) | list of requested skills                                                    | 

### Return type

[**Match**](Match.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

