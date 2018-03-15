/* tslint:disable */

export interface ApiKubernetesVolume {
  gceDisk: string;
  mountPath: string;
  name: string;
  pvcName: string;
  pvcSizeGB: number;
}
