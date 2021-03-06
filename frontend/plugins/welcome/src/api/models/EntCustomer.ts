/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntCustomerEdges,
    EntCustomerEdgesFromJSON,
    EntCustomerEdgesFromJSONTyped,
    EntCustomerEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCustomer
 */
export interface EntCustomer {
    /**
     * CustomerName holds the value of the "customer_name" field.
     * @type {string}
     * @memberof EntCustomer
     */
    customerName?: string;
    /**
     * 
     * @type {EntCustomerEdges}
     * @memberof EntCustomer
     */
    edges?: EntCustomerEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntCustomer
     */
    id?: number;
}

export function EntCustomerFromJSON(json: any): EntCustomer {
    return EntCustomerFromJSONTyped(json, false);
}

export function EntCustomerFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCustomer {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'customerName': !exists(json, 'customer_name') ? undefined : json['customer_name'],
        'edges': !exists(json, 'edges') ? undefined : EntCustomerEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntCustomerToJSON(value?: EntCustomer | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'customer_name': value.customerName,
        'edges': EntCustomerEdgesToJSON(value.edges),
        'id': value.id,
    };
}


