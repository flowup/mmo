/* tslint:disable */
import {
  ApiKubernetesEnvVar,
  ApiKubernetesPort,
  ApiKubernetesVolume,
} from './..';

export interface ApiKubernetesServiceForm {
  ports: ApiKubernetesPort[];
  serviceName: string;
  variables: ApiKubernetesEnvVar[];
  volumes: ApiKubernetesVolume[];
}
