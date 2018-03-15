/* tslint:disable */
import {
  ApiKubernetesEnvVar,
  ApiKubernetesPort,
  ApiKubernetesVolume,
} from './..';

export interface ApiKubernetesServiceForm {
  ports: ApiKubernetesPort[];
  projectName: string;
  serviceName: string;
  variables: ApiKubernetesEnvVar[];
  volumes: ApiKubernetesVolume[];
}
