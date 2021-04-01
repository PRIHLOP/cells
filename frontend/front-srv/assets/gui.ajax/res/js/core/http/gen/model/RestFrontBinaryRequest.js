/**
 * Pydio Cells Rest API
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 */


import ApiClient from '../ApiClient';





/**
* The RestFrontBinaryRequest model module.
* @module model/RestFrontBinaryRequest
* @version 1.0
*/
export default class RestFrontBinaryRequest {
    /**
    * Constructs a new <code>RestFrontBinaryRequest</code>.
    * @alias module:model/RestFrontBinaryRequest
    * @class
    */

    constructor() {
        

        
        

        

        
    }

    /**
    * Constructs a <code>RestFrontBinaryRequest</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/RestFrontBinaryRequest} obj Optional instance to populate.
    * @return {module:model/RestFrontBinaryRequest} The populated <code>RestFrontBinaryRequest</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RestFrontBinaryRequest();

            
            
            

            if (data.hasOwnProperty('BinaryType')) {
                obj['BinaryType'] = ApiClient.convertToType(data['BinaryType'], 'String');
            }
            if (data.hasOwnProperty('Uuid')) {
                obj['Uuid'] = ApiClient.convertToType(data['Uuid'], 'String');
            }
        }
        return obj;
    }

    /**
    * @member {String} BinaryType
    */
    BinaryType = undefined;
    /**
    * @member {String} Uuid
    */
    Uuid = undefined;








}

