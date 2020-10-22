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
    EntBillEdges,
    EntBillEdgesFromJSON,
    EntBillEdgesFromJSONTyped,
    EntBillEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntBill
 */
export interface EntBill {
    /**
     * AddedTime holds the value of the "added_time" field.
     * @type {string}
     * @memberof EntBill
     */
    addedTime?: string;
    /**
     * 
     * @type {EntBillEdges}
     * @memberof EntBill
     */
    edges?: EntBillEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntBill
     */
    id?: number;
    /**
     * Quantity holds the value of the "quantity" field.
     * @type {number}
     * @memberof EntBill
     */
    quantity?: number;
}

export function EntBillFromJSON(json: any): EntBill {
    return EntBillFromJSONTyped(json, false);
}

export function EntBillFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntBill {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'addedTime': !exists(json, 'added_time') ? undefined : json['added_time'],
        'edges': !exists(json, 'edges') ? undefined : EntBillEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'quantity': !exists(json, 'quantity') ? undefined : json['quantity'],
    };
}

export function EntBillToJSON(value?: EntBill | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'added_time': value.addedTime,
        'edges': EntBillEdgesToJSON(value.edges),
        'id': value.id,
        'quantity': value.quantity,
    };
}


