/* tslint:disable */
/* eslint-disable */
/**
 * Deepfence ThreatMapper
 * Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.
 *
 * The version of the OpenAPI document: 2.0.0
 * Contact: community@deepfence.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ModelUpdateUserIdRequest
 */
export interface ModelUpdateUserIdRequest {
    /**
     * 
     * @type {string}
     * @memberof ModelUpdateUserIdRequest
     */
    first_name?: string;
    /**
     * 
     * @type {boolean}
     * @memberof ModelUpdateUserIdRequest
     */
    is_active?: boolean;
    /**
     * 
     * @type {string}
     * @memberof ModelUpdateUserIdRequest
     */
    last_name?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelUpdateUserIdRequest
     */
    role?: string;
}

/**
 * Check if a given object implements the ModelUpdateUserIdRequest interface.
 */
export function instanceOfModelUpdateUserIdRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ModelUpdateUserIdRequestFromJSON(json: any): ModelUpdateUserIdRequest {
    return ModelUpdateUserIdRequestFromJSONTyped(json, false);
}

export function ModelUpdateUserIdRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelUpdateUserIdRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'first_name': !exists(json, 'first_name') ? undefined : json['first_name'],
        'is_active': !exists(json, 'is_active') ? undefined : json['is_active'],
        'last_name': !exists(json, 'last_name') ? undefined : json['last_name'],
        'role': !exists(json, 'role') ? undefined : json['role'],
    };
}

export function ModelUpdateUserIdRequestToJSON(value?: ModelUpdateUserIdRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'first_name': value.first_name,
        'is_active': value.is_active,
        'last_name': value.last_name,
        'role': value.role,
    };
}

