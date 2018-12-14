/**
 * CyberSource Flex API
 * Simple PAN tokenization service
 *
 * OpenAPI spec version: 0.0.1
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.2.3
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/Body1', 'model/InlineResponse400', 'model/TmsV1InstrumentidentifiersDelete409Response', 'model/TmsV1InstrumentidentifiersPost200Response'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/Body1'), require('../model/InlineResponse400'), require('../model/TmsV1InstrumentidentifiersDelete409Response'), require('../model/TmsV1InstrumentidentifiersPost200Response'));
  } else {
    // Browser globals (root is window)
    if (!root.CyberSource) {
      root.CyberSource = {};
    }
    root.CyberSource.InstrumentIdentifierApi = factory(root.CyberSource.ApiClient, root.CyberSource.Body1, root.CyberSource.InlineResponse400, root.CyberSource.TmsV1InstrumentidentifiersDelete409Response, root.CyberSource.TmsV1InstrumentidentifiersPost200Response);
  }
}(this, function(ApiClient, Body1, InlineResponse400, TmsV1InstrumentidentifiersDelete409Response, TmsV1InstrumentidentifiersPost200Response) {
  'use strict';

  /**
   * InstrumentIdentifier service.
   * @module api/InstrumentIdentifierApi
   * @version 0.0.1
   */

  /**
   * Constructs a new InstrumentIdentifierApi. 
   * @alias module:api/InstrumentIdentifierApi
   * @class
   * @param {module:ApiClient} apiClient Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(configObject, apiClient) {
    this.apiClient = apiClient || ApiClient.instance;

    this.apiClient.setConfiguration(configObject);


    /**
     * Callback function to receive the result of the tmsV1InstrumentidentifiersTokenIdDelete operation.
     * @callback module:api/InstrumentIdentifierApi~tmsV1InstrumentidentifiersTokenIdDeleteCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete an Instrument Identifier
     * @param {String} profileId The id of a profile containing user specific TMS configuration.
     * @param {String} tokenId The TokenId of an Instrument Identifier.
     * @param {module:api/InstrumentIdentifierApi~tmsV1InstrumentidentifiersTokenIdDeleteCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.tmsV1InstrumentidentifiersTokenIdDelete = function(profileId, tokenId, callback) {
      var postBody = null;

      // verify the required parameter 'profileId' is set
      if (profileId === undefined || profileId === null) {
        throw new Error("Missing the required parameter 'profileId' when calling tmsV1InstrumentidentifiersTokenIdDelete");
      }

      // verify the required parameter 'tokenId' is set
      if (tokenId === undefined || tokenId === null) {
        throw new Error("Missing the required parameter 'tokenId' when calling tmsV1InstrumentidentifiersTokenIdDelete");
      }


      var pathParams = {
        'tokenId': tokenId
      };
      var queryParams = {
      };
      var headerParams = {
        'profile-id': profileId
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json;charset=utf-8'];
      var accepts = ['application/json;charset=utf-8'];
      var returnType = null;

      return this.apiClient.callApi(
        '/tms/v1/instrumentidentifiers/{tokenId}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the tmsV1InstrumentidentifiersTokenIdGet operation.
     * @callback module:api/InstrumentIdentifierApi~tmsV1InstrumentidentifiersTokenIdGetCallback
     * @param {String} error Error message, if any.
     * @param {module:model/TmsV1InstrumentidentifiersPost200Response} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Retrieve an Instrument Identifier
     * @param {String} profileId The id of a profile containing user specific TMS configuration.
     * @param {String} tokenId The TokenId of an Instrument Identifier.
     * @param {module:api/InstrumentIdentifierApi~tmsV1InstrumentidentifiersTokenIdGetCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/TmsV1InstrumentidentifiersPost200Response}
     */
    this.tmsV1InstrumentidentifiersTokenIdGet = function(profileId, tokenId, callback) {
      var postBody = null;

      // verify the required parameter 'profileId' is set
      if (profileId === undefined || profileId === null) {
        throw new Error("Missing the required parameter 'profileId' when calling tmsV1InstrumentidentifiersTokenIdGet");
      }

      // verify the required parameter 'tokenId' is set
      if (tokenId === undefined || tokenId === null) {
        throw new Error("Missing the required parameter 'tokenId' when calling tmsV1InstrumentidentifiersTokenIdGet");
      }


      var pathParams = {
        'tokenId': tokenId
      };
      var queryParams = {
      };
      var headerParams = {
        'profile-id': profileId
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json;charset=utf-8'];
      var accepts = ['application/json;charset=utf-8'];
      var returnType = TmsV1InstrumentidentifiersPost200Response;

      return this.apiClient.callApi(
        '/tms/v1/instrumentidentifiers/{tokenId}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the tmsV1InstrumentidentifiersTokenIdPatch operation.
     * @callback module:api/InstrumentIdentifierApi~tmsV1InstrumentidentifiersTokenIdPatchCallback
     * @param {String} error Error message, if any.
     * @param {module:model/TmsV1InstrumentidentifiersPost200Response} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update a Instrument Identifier
     * @param {String} profileId The id of a profile containing user specific TMS configuration.
     * @param {String} tokenId The TokenId of an Instrument Identifier.
     * @param {module:model/Body1} body Please specify the previous transaction Id to update.
     * @param {module:api/InstrumentIdentifierApi~tmsV1InstrumentidentifiersTokenIdPatchCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/TmsV1InstrumentidentifiersPost200Response}
     */
    this.tmsV1InstrumentidentifiersTokenIdPatch = function(profileId, tokenId, body, callback) {
      var postBody = body;

      // verify the required parameter 'profileId' is set
      if (profileId === undefined || profileId === null) {
        throw new Error("Missing the required parameter 'profileId' when calling tmsV1InstrumentidentifiersTokenIdPatch");
      }

      // verify the required parameter 'tokenId' is set
      if (tokenId === undefined || tokenId === null) {
        throw new Error("Missing the required parameter 'tokenId' when calling tmsV1InstrumentidentifiersTokenIdPatch");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling tmsV1InstrumentidentifiersTokenIdPatch");
      }


      var pathParams = {
        'tokenId': tokenId
      };
      var queryParams = {
      };
      var headerParams = {
        'profile-id': profileId
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json;charset=utf-8'];
      var accepts = ['application/json;charset=utf-8'];
      var returnType = TmsV1InstrumentidentifiersPost200Response;

      return this.apiClient.callApi(
        '/tms/v1/instrumentidentifiers/{tokenId}', 'PATCH',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));
