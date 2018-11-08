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
    define(['ApiClient', 'model/Tmsv1instrumentidentifiersLinks', 'model/Tmsv1instrumentidentifiersMetadata', 'model/Tmsv1paymentinstrumentsBankAccount', 'model/Tmsv1paymentinstrumentsBillTo', 'model/Tmsv1paymentinstrumentsBuyerInformation', 'model/Tmsv1paymentinstrumentsCard', 'model/Tmsv1paymentinstrumentsInstrumentIdentifier', 'model/Tmsv1paymentinstrumentsMerchantInformation', 'model/Tmsv1paymentinstrumentsProcessingInformation'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./Tmsv1instrumentidentifiersLinks'), require('./Tmsv1instrumentidentifiersMetadata'), require('./Tmsv1paymentinstrumentsBankAccount'), require('./Tmsv1paymentinstrumentsBillTo'), require('./Tmsv1paymentinstrumentsBuyerInformation'), require('./Tmsv1paymentinstrumentsCard'), require('./Tmsv1paymentinstrumentsInstrumentIdentifier'), require('./Tmsv1paymentinstrumentsMerchantInformation'), require('./Tmsv1paymentinstrumentsProcessingInformation'));
  } else {
    // Browser globals (root is window)
    if (!root.CyberSource) {
      root.CyberSource = {};
    }
    root.CyberSource.TmsV1PaymentinstrumentsPost201Response = factory(root.CyberSource.ApiClient, root.CyberSource.Tmsv1instrumentidentifiersLinks, root.CyberSource.Tmsv1instrumentidentifiersMetadata, root.CyberSource.Tmsv1paymentinstrumentsBankAccount, root.CyberSource.Tmsv1paymentinstrumentsBillTo, root.CyberSource.Tmsv1paymentinstrumentsBuyerInformation, root.CyberSource.Tmsv1paymentinstrumentsCard, root.CyberSource.Tmsv1paymentinstrumentsInstrumentIdentifier, root.CyberSource.Tmsv1paymentinstrumentsMerchantInformation, root.CyberSource.Tmsv1paymentinstrumentsProcessingInformation);
  }
}(this, function(ApiClient, Tmsv1instrumentidentifiersLinks, Tmsv1instrumentidentifiersMetadata, Tmsv1paymentinstrumentsBankAccount, Tmsv1paymentinstrumentsBillTo, Tmsv1paymentinstrumentsBuyerInformation, Tmsv1paymentinstrumentsCard, Tmsv1paymentinstrumentsInstrumentIdentifier, Tmsv1paymentinstrumentsMerchantInformation, Tmsv1paymentinstrumentsProcessingInformation) {
  'use strict';




  /**
   * The TmsV1PaymentinstrumentsPost201Response model module.
   * @module model/TmsV1PaymentinstrumentsPost201Response
   * @version 0.0.1
   */

  /**
   * Constructs a new <code>TmsV1PaymentinstrumentsPost201Response</code>.
   * @alias module:model/TmsV1PaymentinstrumentsPost201Response
   * @class
   */
  var exports = function() {
    var _this = this;













  };

  /**
   * Constructs a <code>TmsV1PaymentinstrumentsPost201Response</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/TmsV1PaymentinstrumentsPost201Response} obj Optional instance to populate.
   * @return {module:model/TmsV1PaymentinstrumentsPost201Response} The populated <code>TmsV1PaymentinstrumentsPost201Response</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('_links')) {
        obj['_links'] = Tmsv1instrumentidentifiersLinks.constructFromObject(data['_links']);
      }
      if (data.hasOwnProperty('id')) {
        obj['id'] = ApiClient.convertToType(data['id'], 'String');
      }
      if (data.hasOwnProperty('object')) {
        obj['object'] = ApiClient.convertToType(data['object'], 'String');
      }
      if (data.hasOwnProperty('state')) {
        obj['state'] = ApiClient.convertToType(data['state'], 'String');
      }
      if (data.hasOwnProperty('bankAccount')) {
        obj['bankAccount'] = Tmsv1paymentinstrumentsBankAccount.constructFromObject(data['bankAccount']);
      }
      if (data.hasOwnProperty('card')) {
        obj['card'] = Tmsv1paymentinstrumentsCard.constructFromObject(data['card']);
      }
      if (data.hasOwnProperty('buyerInformation')) {
        obj['buyerInformation'] = Tmsv1paymentinstrumentsBuyerInformation.constructFromObject(data['buyerInformation']);
      }
      if (data.hasOwnProperty('billTo')) {
        obj['billTo'] = Tmsv1paymentinstrumentsBillTo.constructFromObject(data['billTo']);
      }
      if (data.hasOwnProperty('processingInformation')) {
        obj['processingInformation'] = Tmsv1paymentinstrumentsProcessingInformation.constructFromObject(data['processingInformation']);
      }
      if (data.hasOwnProperty('merchantInformation')) {
        obj['merchantInformation'] = Tmsv1paymentinstrumentsMerchantInformation.constructFromObject(data['merchantInformation']);
      }
      if (data.hasOwnProperty('metaData')) {
        obj['metaData'] = Tmsv1instrumentidentifiersMetadata.constructFromObject(data['metaData']);
      }
      if (data.hasOwnProperty('instrumentIdentifier')) {
        obj['instrumentIdentifier'] = Tmsv1paymentinstrumentsInstrumentIdentifier.constructFromObject(data['instrumentIdentifier']);
      }
    }
    return obj;
  }

  /**
   * @member {module:model/Tmsv1instrumentidentifiersLinks} _links
   */
  exports.prototype['_links'] = undefined;
  /**
   * Unique identification number assigned by CyberSource to the submitted request.
   * @member {String} id
   */
  exports.prototype['id'] = undefined;
  /**
   * Describes type of token. For example: customer, paymentInstrument or instrumentIdentifier.
   * @member {module:model/TmsV1PaymentinstrumentsPost201Response.ObjectEnum} object
   */
  exports.prototype['object'] = undefined;
  /**
   * Current state of the token.
   * @member {module:model/TmsV1PaymentinstrumentsPost201Response.StateEnum} state
   */
  exports.prototype['state'] = undefined;
  /**
   * @member {module:model/Tmsv1paymentinstrumentsBankAccount} bankAccount
   */
  exports.prototype['bankAccount'] = undefined;
  /**
   * @member {module:model/Tmsv1paymentinstrumentsCard} card
   */
  exports.prototype['card'] = undefined;
  /**
   * @member {module:model/Tmsv1paymentinstrumentsBuyerInformation} buyerInformation
   */
  exports.prototype['buyerInformation'] = undefined;
  /**
   * @member {module:model/Tmsv1paymentinstrumentsBillTo} billTo
   */
  exports.prototype['billTo'] = undefined;
  /**
   * @member {module:model/Tmsv1paymentinstrumentsProcessingInformation} processingInformation
   */
  exports.prototype['processingInformation'] = undefined;
  /**
   * @member {module:model/Tmsv1paymentinstrumentsMerchantInformation} merchantInformation
   */
  exports.prototype['merchantInformation'] = undefined;
  /**
   * @member {module:model/Tmsv1instrumentidentifiersMetadata} metaData
   */
  exports.prototype['metaData'] = undefined;
  /**
   * @member {module:model/Tmsv1paymentinstrumentsInstrumentIdentifier} instrumentIdentifier
   */
  exports.prototype['instrumentIdentifier'] = undefined;


  /**
   * Allowed values for the <code>object</code> property.
   * @enum {String}
   * @readonly
   */
  exports.ObjectEnum = {
    /**
     * value: "paymentInstrument"
     * @const
     */
    "paymentInstrument": "paymentInstrument"  };

  /**
   * Allowed values for the <code>state</code> property.
   * @enum {String}
   * @readonly
   */
  exports.StateEnum = {
    /**
     * value: "ACTIVE"
     * @const
     */
    "ACTIVE": "ACTIVE",
    /**
     * value: "CLOSED"
     * @const
     */
    "CLOSED": "CLOSED"  };


  return exports;
}));

