# CdnSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CachingTime** | Pointer to **int32** |  | [optional] 
**FollowOriginRedirect** | Pointer to [**CdnSettingsFollowOriginRedirect**](CdnSettingsFollowOriginRedirect.md) |  | [optional] 
**CachingTimeBrowser** | Pointer to [**CdnSettingsCachingTimeBrowser**](CdnSettingsCachingTimeBrowser.md) |  | [optional] 
**IgnoreCookie** | Pointer to **bool** |  | [optional] 
**IgnoreQueryString** | Pointer to **bool** |  | [optional] 
**Stale** | Pointer to [**CdnSettingsStale**](CdnSettingsStale.md) |  | [optional] 
**GeoAcl** | Pointer to [**CdnSettingsGeoAcl**](CdnSettingsGeoAcl.md) |  | [optional] 
**RefererAcl** | Pointer to [**CdnSettingsRefererAcl**](CdnSettingsRefererAcl.md) |  | [optional] 
**IpAddressAcl** | Pointer to [**CdnSettingsIpAddressAcl**](CdnSettingsIpAddressAcl.md) |  | [optional] 
**RedirectHttpToHttps** | Pointer to **bool** |  | [optional] 
**UserAgentAcl** | Pointer to [**CdnSettingsUserAgentAcl**](CdnSettingsUserAgentAcl.md) |  | [optional] 
**TokenizedUrlSecureKey** | Pointer to [**CdnSettingsTokenizedUrlSecureKey**](CdnSettingsTokenizedUrlSecureKey.md) |  | [optional] 
**AllowedHttpMethods** | Pointer to [**CdnSettingsAllowedHttpMethods**](CdnSettingsAllowedHttpMethods.md) |  | [optional] 
**Http3Enabled** | Pointer to **bool** |  | [optional] 
**GzipCompression** | Pointer to [**CdnSettingsGzipCompression**](CdnSettingsGzipCompression.md) |  | [optional] 
**ContentSlice** | Pointer to **bool** |  | [optional] 
**StaticRequestHeaders** | Pointer to [**CdnSettingsStaticRequestHeaders**](CdnSettingsStaticRequestHeaders.md) |  | [optional] 
**Cors** | Pointer to [**CdnSettingsCors**](CdnSettingsCors.md) |  | [optional] 
**StaticResponseHeaders** | Pointer to [**CdnSettingsStaticResponseHeader**](CdnSettingsStaticResponseHeader.md) |  | [optional] 
**ResponseHeadersHidingPolicy** | Pointer to [**CdnSettingsResponseHeadersHidingPolicy**](CdnSettingsResponseHeadersHidingPolicy.md) |  | [optional] 

## Methods

### NewCdnSettings

`func NewCdnSettings() *CdnSettings`

NewCdnSettings instantiates a new CdnSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnSettingsWithDefaults

`func NewCdnSettingsWithDefaults() *CdnSettings`

NewCdnSettingsWithDefaults instantiates a new CdnSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCachingTime

`func (o *CdnSettings) GetCachingTime() int32`

GetCachingTime returns the CachingTime field if non-nil, zero value otherwise.

### GetCachingTimeOk

`func (o *CdnSettings) GetCachingTimeOk() (*int32, bool)`

GetCachingTimeOk returns a tuple with the CachingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachingTime

`func (o *CdnSettings) SetCachingTime(v int32)`

SetCachingTime sets CachingTime field to given value.

### HasCachingTime

`func (o *CdnSettings) HasCachingTime() bool`

HasCachingTime returns a boolean if a field has been set.

### GetFollowOriginRedirect

`func (o *CdnSettings) GetFollowOriginRedirect() CdnSettingsFollowOriginRedirect`

GetFollowOriginRedirect returns the FollowOriginRedirect field if non-nil, zero value otherwise.

### GetFollowOriginRedirectOk

`func (o *CdnSettings) GetFollowOriginRedirectOk() (*CdnSettingsFollowOriginRedirect, bool)`

GetFollowOriginRedirectOk returns a tuple with the FollowOriginRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowOriginRedirect

`func (o *CdnSettings) SetFollowOriginRedirect(v CdnSettingsFollowOriginRedirect)`

SetFollowOriginRedirect sets FollowOriginRedirect field to given value.

### HasFollowOriginRedirect

`func (o *CdnSettings) HasFollowOriginRedirect() bool`

HasFollowOriginRedirect returns a boolean if a field has been set.

### GetCachingTimeBrowser

`func (o *CdnSettings) GetCachingTimeBrowser() CdnSettingsCachingTimeBrowser`

GetCachingTimeBrowser returns the CachingTimeBrowser field if non-nil, zero value otherwise.

### GetCachingTimeBrowserOk

`func (o *CdnSettings) GetCachingTimeBrowserOk() (*CdnSettingsCachingTimeBrowser, bool)`

GetCachingTimeBrowserOk returns a tuple with the CachingTimeBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachingTimeBrowser

`func (o *CdnSettings) SetCachingTimeBrowser(v CdnSettingsCachingTimeBrowser)`

SetCachingTimeBrowser sets CachingTimeBrowser field to given value.

### HasCachingTimeBrowser

`func (o *CdnSettings) HasCachingTimeBrowser() bool`

HasCachingTimeBrowser returns a boolean if a field has been set.

### GetIgnoreCookie

`func (o *CdnSettings) GetIgnoreCookie() bool`

GetIgnoreCookie returns the IgnoreCookie field if non-nil, zero value otherwise.

### GetIgnoreCookieOk

`func (o *CdnSettings) GetIgnoreCookieOk() (*bool, bool)`

GetIgnoreCookieOk returns a tuple with the IgnoreCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCookie

`func (o *CdnSettings) SetIgnoreCookie(v bool)`

SetIgnoreCookie sets IgnoreCookie field to given value.

### HasIgnoreCookie

`func (o *CdnSettings) HasIgnoreCookie() bool`

HasIgnoreCookie returns a boolean if a field has been set.

### GetIgnoreQueryString

`func (o *CdnSettings) GetIgnoreQueryString() bool`

GetIgnoreQueryString returns the IgnoreQueryString field if non-nil, zero value otherwise.

### GetIgnoreQueryStringOk

`func (o *CdnSettings) GetIgnoreQueryStringOk() (*bool, bool)`

GetIgnoreQueryStringOk returns a tuple with the IgnoreQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreQueryString

`func (o *CdnSettings) SetIgnoreQueryString(v bool)`

SetIgnoreQueryString sets IgnoreQueryString field to given value.

### HasIgnoreQueryString

`func (o *CdnSettings) HasIgnoreQueryString() bool`

HasIgnoreQueryString returns a boolean if a field has been set.

### GetStale

`func (o *CdnSettings) GetStale() CdnSettingsStale`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *CdnSettings) GetStaleOk() (*CdnSettingsStale, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *CdnSettings) SetStale(v CdnSettingsStale)`

SetStale sets Stale field to given value.

### HasStale

`func (o *CdnSettings) HasStale() bool`

HasStale returns a boolean if a field has been set.

### GetGeoAcl

`func (o *CdnSettings) GetGeoAcl() CdnSettingsGeoAcl`

GetGeoAcl returns the GeoAcl field if non-nil, zero value otherwise.

### GetGeoAclOk

`func (o *CdnSettings) GetGeoAclOk() (*CdnSettingsGeoAcl, bool)`

GetGeoAclOk returns a tuple with the GeoAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoAcl

`func (o *CdnSettings) SetGeoAcl(v CdnSettingsGeoAcl)`

SetGeoAcl sets GeoAcl field to given value.

### HasGeoAcl

`func (o *CdnSettings) HasGeoAcl() bool`

HasGeoAcl returns a boolean if a field has been set.

### GetRefererAcl

`func (o *CdnSettings) GetRefererAcl() CdnSettingsRefererAcl`

GetRefererAcl returns the RefererAcl field if non-nil, zero value otherwise.

### GetRefererAclOk

`func (o *CdnSettings) GetRefererAclOk() (*CdnSettingsRefererAcl, bool)`

GetRefererAclOk returns a tuple with the RefererAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefererAcl

`func (o *CdnSettings) SetRefererAcl(v CdnSettingsRefererAcl)`

SetRefererAcl sets RefererAcl field to given value.

### HasRefererAcl

`func (o *CdnSettings) HasRefererAcl() bool`

HasRefererAcl returns a boolean if a field has been set.

### GetIpAddressAcl

`func (o *CdnSettings) GetIpAddressAcl() CdnSettingsIpAddressAcl`

GetIpAddressAcl returns the IpAddressAcl field if non-nil, zero value otherwise.

### GetIpAddressAclOk

`func (o *CdnSettings) GetIpAddressAclOk() (*CdnSettingsIpAddressAcl, bool)`

GetIpAddressAclOk returns a tuple with the IpAddressAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressAcl

`func (o *CdnSettings) SetIpAddressAcl(v CdnSettingsIpAddressAcl)`

SetIpAddressAcl sets IpAddressAcl field to given value.

### HasIpAddressAcl

`func (o *CdnSettings) HasIpAddressAcl() bool`

HasIpAddressAcl returns a boolean if a field has been set.

### GetRedirectHttpToHttps

`func (o *CdnSettings) GetRedirectHttpToHttps() bool`

GetRedirectHttpToHttps returns the RedirectHttpToHttps field if non-nil, zero value otherwise.

### GetRedirectHttpToHttpsOk

`func (o *CdnSettings) GetRedirectHttpToHttpsOk() (*bool, bool)`

GetRedirectHttpToHttpsOk returns a tuple with the RedirectHttpToHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectHttpToHttps

`func (o *CdnSettings) SetRedirectHttpToHttps(v bool)`

SetRedirectHttpToHttps sets RedirectHttpToHttps field to given value.

### HasRedirectHttpToHttps

`func (o *CdnSettings) HasRedirectHttpToHttps() bool`

HasRedirectHttpToHttps returns a boolean if a field has been set.

### GetUserAgentAcl

`func (o *CdnSettings) GetUserAgentAcl() CdnSettingsUserAgentAcl`

GetUserAgentAcl returns the UserAgentAcl field if non-nil, zero value otherwise.

### GetUserAgentAclOk

`func (o *CdnSettings) GetUserAgentAclOk() (*CdnSettingsUserAgentAcl, bool)`

GetUserAgentAclOk returns a tuple with the UserAgentAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgentAcl

`func (o *CdnSettings) SetUserAgentAcl(v CdnSettingsUserAgentAcl)`

SetUserAgentAcl sets UserAgentAcl field to given value.

### HasUserAgentAcl

`func (o *CdnSettings) HasUserAgentAcl() bool`

HasUserAgentAcl returns a boolean if a field has been set.

### GetTokenizedUrlSecureKey

`func (o *CdnSettings) GetTokenizedUrlSecureKey() CdnSettingsTokenizedUrlSecureKey`

GetTokenizedUrlSecureKey returns the TokenizedUrlSecureKey field if non-nil, zero value otherwise.

### GetTokenizedUrlSecureKeyOk

`func (o *CdnSettings) GetTokenizedUrlSecureKeyOk() (*CdnSettingsTokenizedUrlSecureKey, bool)`

GetTokenizedUrlSecureKeyOk returns a tuple with the TokenizedUrlSecureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizedUrlSecureKey

`func (o *CdnSettings) SetTokenizedUrlSecureKey(v CdnSettingsTokenizedUrlSecureKey)`

SetTokenizedUrlSecureKey sets TokenizedUrlSecureKey field to given value.

### HasTokenizedUrlSecureKey

`func (o *CdnSettings) HasTokenizedUrlSecureKey() bool`

HasTokenizedUrlSecureKey returns a boolean if a field has been set.

### GetAllowedHttpMethods

`func (o *CdnSettings) GetAllowedHttpMethods() CdnSettingsAllowedHttpMethods`

GetAllowedHttpMethods returns the AllowedHttpMethods field if non-nil, zero value otherwise.

### GetAllowedHttpMethodsOk

`func (o *CdnSettings) GetAllowedHttpMethodsOk() (*CdnSettingsAllowedHttpMethods, bool)`

GetAllowedHttpMethodsOk returns a tuple with the AllowedHttpMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedHttpMethods

`func (o *CdnSettings) SetAllowedHttpMethods(v CdnSettingsAllowedHttpMethods)`

SetAllowedHttpMethods sets AllowedHttpMethods field to given value.

### HasAllowedHttpMethods

`func (o *CdnSettings) HasAllowedHttpMethods() bool`

HasAllowedHttpMethods returns a boolean if a field has been set.

### GetHttp3Enabled

`func (o *CdnSettings) GetHttp3Enabled() bool`

GetHttp3Enabled returns the Http3Enabled field if non-nil, zero value otherwise.

### GetHttp3EnabledOk

`func (o *CdnSettings) GetHttp3EnabledOk() (*bool, bool)`

GetHttp3EnabledOk returns a tuple with the Http3Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp3Enabled

`func (o *CdnSettings) SetHttp3Enabled(v bool)`

SetHttp3Enabled sets Http3Enabled field to given value.

### HasHttp3Enabled

`func (o *CdnSettings) HasHttp3Enabled() bool`

HasHttp3Enabled returns a boolean if a field has been set.

### GetGzipCompression

`func (o *CdnSettings) GetGzipCompression() CdnSettingsGzipCompression`

GetGzipCompression returns the GzipCompression field if non-nil, zero value otherwise.

### GetGzipCompressionOk

`func (o *CdnSettings) GetGzipCompressionOk() (*CdnSettingsGzipCompression, bool)`

GetGzipCompressionOk returns a tuple with the GzipCompression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGzipCompression

`func (o *CdnSettings) SetGzipCompression(v CdnSettingsGzipCompression)`

SetGzipCompression sets GzipCompression field to given value.

### HasGzipCompression

`func (o *CdnSettings) HasGzipCompression() bool`

HasGzipCompression returns a boolean if a field has been set.

### GetContentSlice

`func (o *CdnSettings) GetContentSlice() bool`

GetContentSlice returns the ContentSlice field if non-nil, zero value otherwise.

### GetContentSliceOk

`func (o *CdnSettings) GetContentSliceOk() (*bool, bool)`

GetContentSliceOk returns a tuple with the ContentSlice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSlice

`func (o *CdnSettings) SetContentSlice(v bool)`

SetContentSlice sets ContentSlice field to given value.

### HasContentSlice

`func (o *CdnSettings) HasContentSlice() bool`

HasContentSlice returns a boolean if a field has been set.

### GetStaticRequestHeaders

`func (o *CdnSettings) GetStaticRequestHeaders() CdnSettingsStaticRequestHeaders`

GetStaticRequestHeaders returns the StaticRequestHeaders field if non-nil, zero value otherwise.

### GetStaticRequestHeadersOk

`func (o *CdnSettings) GetStaticRequestHeadersOk() (*CdnSettingsStaticRequestHeaders, bool)`

GetStaticRequestHeadersOk returns a tuple with the StaticRequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRequestHeaders

`func (o *CdnSettings) SetStaticRequestHeaders(v CdnSettingsStaticRequestHeaders)`

SetStaticRequestHeaders sets StaticRequestHeaders field to given value.

### HasStaticRequestHeaders

`func (o *CdnSettings) HasStaticRequestHeaders() bool`

HasStaticRequestHeaders returns a boolean if a field has been set.

### GetCors

`func (o *CdnSettings) GetCors() CdnSettingsCors`

GetCors returns the Cors field if non-nil, zero value otherwise.

### GetCorsOk

`func (o *CdnSettings) GetCorsOk() (*CdnSettingsCors, bool)`

GetCorsOk returns a tuple with the Cors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCors

`func (o *CdnSettings) SetCors(v CdnSettingsCors)`

SetCors sets Cors field to given value.

### HasCors

`func (o *CdnSettings) HasCors() bool`

HasCors returns a boolean if a field has been set.

### GetStaticResponseHeaders

`func (o *CdnSettings) GetStaticResponseHeaders() CdnSettingsStaticResponseHeader`

GetStaticResponseHeaders returns the StaticResponseHeaders field if non-nil, zero value otherwise.

### GetStaticResponseHeadersOk

`func (o *CdnSettings) GetStaticResponseHeadersOk() (*CdnSettingsStaticResponseHeader, bool)`

GetStaticResponseHeadersOk returns a tuple with the StaticResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticResponseHeaders

`func (o *CdnSettings) SetStaticResponseHeaders(v CdnSettingsStaticResponseHeader)`

SetStaticResponseHeaders sets StaticResponseHeaders field to given value.

### HasStaticResponseHeaders

`func (o *CdnSettings) HasStaticResponseHeaders() bool`

HasStaticResponseHeaders returns a boolean if a field has been set.

### GetResponseHeadersHidingPolicy

`func (o *CdnSettings) GetResponseHeadersHidingPolicy() CdnSettingsResponseHeadersHidingPolicy`

GetResponseHeadersHidingPolicy returns the ResponseHeadersHidingPolicy field if non-nil, zero value otherwise.

### GetResponseHeadersHidingPolicyOk

`func (o *CdnSettings) GetResponseHeadersHidingPolicyOk() (*CdnSettingsResponseHeadersHidingPolicy, bool)`

GetResponseHeadersHidingPolicyOk returns a tuple with the ResponseHeadersHidingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeadersHidingPolicy

`func (o *CdnSettings) SetResponseHeadersHidingPolicy(v CdnSettingsResponseHeadersHidingPolicy)`

SetResponseHeadersHidingPolicy sets ResponseHeadersHidingPolicy field to given value.

### HasResponseHeadersHidingPolicy

`func (o *CdnSettings) HasResponseHeadersHidingPolicy() bool`

HasResponseHeadersHidingPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


