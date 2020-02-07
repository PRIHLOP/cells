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
import ServiceQuery from './ServiceQuery';





/**
* The JobsActionOutputFilter model module.
* @module model/JobsActionOutputFilter
* @version 1.0
*/
export default class JobsActionOutputFilter {
    /**
    * Constructs a new <code>JobsActionOutputFilter</code>.
    * @alias module:model/JobsActionOutputFilter
    * @class
    */

    constructor() {
        

        
        

        

        
    }

    /**
    * Constructs a <code>JobsActionOutputFilter</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/JobsActionOutputFilter} obj Optional instance to populate.
    * @return {module:model/JobsActionOutputFilter} The populated <code>JobsActionOutputFilter</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new JobsActionOutputFilter();

            
            
            

            if (data.hasOwnProperty('Query')) {
                obj['Query'] = ServiceQuery.constructFromObject(data['Query']);
            }
        }
        return obj;
    }

    /**
    * @member {module:model/ServiceQuery} Query
    */
    Query = undefined;








}

