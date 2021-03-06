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
    EntBill,
    EntBillFromJSON,
    EntBillFromJSONTyped,
    EntBillToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntProductEdges
 */
export interface EntProductEdges {
    /**
     * Products holds the value of the products edge.
     * @type {Array<EntBill>}
     * @memberof EntProductEdges
     */
    products?: Array<EntBill>;
}

export function EntProductEdgesFromJSON(json: any): EntProductEdges {
    return EntProductEdgesFromJSONTyped(json, false);
}

export function EntProductEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntProductEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'products': !exists(json, 'products') ? undefined : ((json['products'] as Array<any>).map(EntBillFromJSON)),
    };
}

export function EntProductEdgesToJSON(value?: EntProductEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'products': value.products === undefined ? undefined : ((value.products as Array<any>).map(EntBillToJSON)),
    };
}


