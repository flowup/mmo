/* tslint:disable */

import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Inject, Injectable, Optional } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import {
  ApiKubernetesConfigs,
  ApiKubernetesServiceForm,
  ApiPlugins,
  ApiServices,
  ApiVersion,
} from '.';

interface HttpOptions {
  headers?: HttpHeaders,
  params?: HttpParams,
  reportProgress?: boolean,
  withCredentials?: boolean,
}

/**
 * Created with ngx-swagger-client-generator (https://github.com/flowup/ngx-swagger-client-generator)
 */
@Injectable()
export class ApiClientService {

  readonly options: HttpOptions;
  private domain = 'http://localhost:8080';

  constructor(private http: HttpClient,
              @Optional() @Inject('domain') domain: string,
              @Optional() @Inject('httpOptions') options: HttpOptions) {
    if (domain) {
      this.domain = domain;
    }

    this.options = {
      headers: options && options.headers ? options.headers : new HttpHeaders(),
      params: options && options.params ? options.params : new HttpParams()
    };
  }

  getGlobalPlugins(options?: HttpOptions): Observable<ApiPlugins> {
    const path = `/plugins`;
    options = {...this.options, ...options};

    return this.sendRequest<ApiPlugins>('GET', path, options);
  }

  getServices(options?: HttpOptions): Observable<ApiServices> {
    const path = `/services`;
    options = {...this.options, ...options};

    return this.sendRequest<ApiServices>('GET', path, options);
  }

  getKubernetesConfigs(name: string, description: string, options?: HttpOptions): Observable<ApiKubernetesConfigs> {
    const path = `/services/${name}/kubernetes`;
    options = {...this.options, ...options};

    if (description) {
      options.params = options.params.set('description', String(description));
    }
    return this.sendRequest<ApiKubernetesConfigs>('GET', path, options);
  }

  kubernetesFormFromPlugins(name: string, description: string, options?: HttpOptions): Observable<ApiKubernetesServiceForm> {
    const path = `/services/${name}/kubernetes/form`;
    options = {...this.options, ...options};

    if (description) {
      options.params = options.params.set('description', String(description));
    }
    return this.sendRequest<ApiKubernetesServiceForm>('GET', path, options);
  }

  getPlugins(name: string, description: string, options?: HttpOptions): Observable<ApiPlugins> {
    const path = `/services/${name}/plugins`;
    options = {...this.options, ...options};

    if (description) {
      options.params = options.params.set('description', String(description));
    }
    return this.sendRequest<ApiPlugins>('GET', path, options);
  }

  kubernetesConfigFromForm(serviceName: string, options?: HttpOptions): Observable<ApiKubernetesConfigs> {
    const path = `/services/${serviceName}/kubernetes/create`;
    options = {...this.options, ...options};

    return this.sendRequest<ApiKubernetesConfigs>('POST', path, options);
  }

  getVersion(options?: HttpOptions): Observable<ApiVersion> {
    const path = `/version`;
    options = {...this.options, ...options};

    return this.sendRequest<ApiVersion>('GET', path, options);
  }

  private sendRequest<T>(method: string, path: string, options: HttpOptions, body?: any): Observable<T> {
    if (method === 'GET') {
      return this.http.get<T>(`${this.domain}${path}`, options);
    } else if (method === 'PUT') {
      return this.http.put<T>(`${this.domain}${path}`, body, options);
    } else if (method === 'POST') {
      return this.http.post<T>(`${this.domain}${path}`, body, options);
    } else if (method === 'DELETE') {
      return this.http.delete<T>(`${this.domain}${path}`, options);
    } else {
      console.error('Unsupported request: ' + method);
      return Observable.throw('Unsupported request: ' + method);
    }
  }
}
