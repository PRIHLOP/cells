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
* The InstallTLSLetsEncrypt model module.
* @module model/InstallTLSLetsEncrypt
* @version 1.0
*/
export default class InstallTLSLetsEncrypt {
    /**
    * Constructs a new <code>InstallTLSLetsEncrypt</code>.
    * @alias module:model/InstallTLSLetsEncrypt
    * @class
    */

    constructor() {
        

        
        

        

        
    }

    /**
    * Constructs a <code>InstallTLSLetsEncrypt</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/InstallTLSLetsEncrypt} obj Optional instance to populate.
    * @return {module:model/InstallTLSLetsEncrypt} The populated <code>InstallTLSLetsEncrypt</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstallTLSLetsEncrypt();

            
            
            

            if (data.hasOwnProperty('Email')) {
                obj['Email'] = ApiClient.convertToType(data['Email'], 'String');
            }
            if (data.hasOwnProperty('AcceptEULA')) {
                obj['AcceptEULA'] = ApiClient.convertToType(data['AcceptEULA'], 'Boolean');
            }
            if (data.hasOwnProperty('StagingCA')) {
                obj['StagingCA'] = ApiClient.convertToType(data['StagingCA'], 'Boolean');
            }
        }
        return obj;
    }

    /**
    * @member {String} Email
    */
    Email = undefined;
    /**
    * @member {Boolean} AcceptEULA
    */
    AcceptEULA = undefined;
    /**
    * @member {Boolean} StagingCA
    */
    StagingCA = undefined;








}

