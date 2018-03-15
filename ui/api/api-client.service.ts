/* tslint:disable */

import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Inject, Injectable, InjectionToken, Optional } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { HttpOptions } from './';
import * as models from './models';

export const USE_DOMAIN = new InjectionToken<string>('USE_DOMAIN');
export const USE_HTTP_OPTIONS = new InjectionToken<HttpOptions>('USE_HTTP_OPTIONS');

/**
 * Created with https://github.com/flowup/api-client-generator
 */
@Injectable()
export class APIClient {

  readonly options: HttpOptions;
  private domain: string = `http://host`;

  constructor(private http: HttpClient,
              @Optional() @Inject(USE_DOMAIN) domain: string,
              @Optional() @Inject(USE_HTTP_OPTIONS) options: HttpOptions) {

    if (domain) {
      this.domain = domain;
    }

    this.options = {
      headers: options && options.headers ? options.headers : new HttpHeaders(),
      params: options && options.params ? options.params : new HttpParams()
    };
  }

  getGlobalPlugins(options?: HttpOptions): Observable<models.ApiPlugins> {
    const path = `/plugins`;
    options = {...this.options, ...options};

    return this.sendRequest<models.ApiPlugins>('GET', path, options);
  }

  getServices(options?: HttpOptions): Observable<models.ApiServices> {
    const path = `/services`;
    options = {...this.options, ...options};

    return this.sendRequest<models.ApiServices>('GET', path, options);
  }

  getKubernetesConfigs(name: string, description: string, options?: HttpOptions): Observable<models.ApiKubernetesConfigs> {
    const path = `/services/${name}/kubernetes`;
    options = {...this.options, ...options};

    if (description) {
      options.params = options.params.set('description', String(description));
    }
    return this.sendRequest<models.ApiKubernetesConfigs>('GET', path, options);
  }

  kubernetesFormFromPlugins(name: string, description: string, options?: HttpOptions): Observable<models.ApiKubernetesServiceForm> {
    const path = `/services/${name}/kubernetes/form`;
    options = {...this.options, ...options};

    if (description) {
      options.params = options.params.set('description', String(description));
    }
    return this.sendRequest<models.ApiKubernetesServiceForm>('GET', path, options);
  }

  getPlugins(name: string, description: string, options?: HttpOptions): Observable<models.ApiPlugins> {
    const path = `/services/${name}/plugins`;
    options = {...this.options, ...options};

    if (description) {
      options.params = options.params.set('description', String(description));
    }
    return this.sendRequest<models.ApiPlugins>('GET', path, options);
  }

  kubernetesConfigFromForm(serviceName: string, options?: HttpOptions): Observable<models.ApiKubernetesConfigs> {
    const path = `/services/${serviceName}/kubernetes/create`;
    options = {...this.options, ...options};

    return this.sendRequest<models.ApiKubernetesConfigs>('POST', path, options);
  }

  getVersion(options?: HttpOptions): Observable<models.ApiVersion> {
    const path = `/version`;
    options = {...this.options, ...options};

    return this.sendRequest<models.ApiVersion>('GET', path, options);
  }

  private sendRequest<T>(method: string, path: string, options: HttpOptions, body?: any): Observable<T> {
    switch (method) {
      case 'GET':
        return this.http.get<T>(`${this.domain}${path}`, options);
      case 'PUT':
        return this.http.put<T>(`${this.domain}${path}`, body, options);
      case 'POST':
        return this.http.post<T>(`${this.domain}${path}`, body, options);
      case 'DELETE':
        return this.http.delete<T>(`${this.domain}${path}`, options);
      default:
        console.error(`Unsupported request: ${method}`);
        return Observable.throw(`Unsupported request: ${method}`);
    }
  }
}
