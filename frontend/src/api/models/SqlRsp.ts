/* tslint:disable */
/* eslint-disable */
/**
 * GinSqlBlog
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface SqlRsp
 */
export interface SqlRsp {
    /**
     * 
     * @type {string}
     * @memberof SqlRsp
     */
    context: string;
}

/**
 * Check if a given object implements the SqlRsp interface.
 */
export function instanceOfSqlRsp(value: object): value is SqlRsp {
    if (!('context' in value) || value['context'] === undefined) return false;
    return true;
}

export function SqlRspFromJSON(json: any): SqlRsp {
    return SqlRspFromJSONTyped(json, false);
}

export function SqlRspFromJSONTyped(json: any, ignoreDiscriminator: boolean): SqlRsp {
    if (json == null) {
        return json;
    }
    return {
        
        'context': json['context'],
    };
}

export function SqlRspToJSON(json: any): SqlRsp {
    return SqlRspToJSONTyped(json, false);
}

export function SqlRspToJSONTyped(value?: SqlRsp | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'context': value['context'],
    };
}

