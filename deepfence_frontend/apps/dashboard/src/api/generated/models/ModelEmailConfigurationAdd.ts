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
 * @interface ModelEmailConfigurationAdd
 */
export interface ModelEmailConfigurationAdd {
    /**
     * 
     * @type {string}
     * @memberof ModelEmailConfigurationAdd
     */
    amazon_access_key?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelEmailConfigurationAdd
     */
    amazon_secret_key?: string;
    /**
     * 
     * @type {number}
     * @memberof ModelEmailConfigurationAdd
     */
    created_by_user_id?: number;
    /**
     * 
     * @type {string}
     * @memberof ModelEmailConfigurationAdd
     */
    email_id?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelEmailConfigurationAdd
     */
    email_provider?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelEmailConfigurationAdd
     */
    password?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelEmailConfigurationAdd
     */
    port?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelEmailConfigurationAdd
     */
    ses_region?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelEmailConfigurationAdd
     */
    smtp?: string;
}

/**
 * Check if a given object implements the ModelEmailConfigurationAdd interface.
 */
export function instanceOfModelEmailConfigurationAdd(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ModelEmailConfigurationAddFromJSON(json: any): ModelEmailConfigurationAdd {
    return ModelEmailConfigurationAddFromJSONTyped(json, false);
}

export function ModelEmailConfigurationAddFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelEmailConfigurationAdd {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'amazon_access_key': !exists(json, 'amazon_access_key') ? undefined : json['amazon_access_key'],
        'amazon_secret_key': !exists(json, 'amazon_secret_key') ? undefined : json['amazon_secret_key'],
        'created_by_user_id': !exists(json, 'created_by_user_id') ? undefined : json['created_by_user_id'],
        'email_id': !exists(json, 'email_id') ? undefined : json['email_id'],
        'email_provider': !exists(json, 'email_provider') ? undefined : json['email_provider'],
        'password': !exists(json, 'password') ? undefined : json['password'],
        'port': !exists(json, 'port') ? undefined : json['port'],
        'ses_region': !exists(json, 'ses_region') ? undefined : json['ses_region'],
        'smtp': !exists(json, 'smtp') ? undefined : json['smtp'],
    };
}

export function ModelEmailConfigurationAddToJSON(value?: ModelEmailConfigurationAdd | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'amazon_access_key': value.amazon_access_key,
        'amazon_secret_key': value.amazon_secret_key,
        'created_by_user_id': value.created_by_user_id,
        'email_id': value.email_id,
        'email_provider': value.email_provider,
        'password': value.password,
        'port': value.port,
        'ses_region': value.ses_region,
        'smtp': value.smtp,
    };
}
