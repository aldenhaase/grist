import { Injectable } from '@angular/core';
import { collection } from 'src/app/types/listTypes';
import { Observable, firstValueFrom} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class MergeService {

  constructor() { }
  async merge(remote$: Observable<collection>, local: collection){
    const remote = await firstValueFrom(remote$)
    const remoteString = JSON.stringify(remote)
    const localString = JSON.stringify(local)
    if(remoteString != localString){
      local.lists = remote.lists
    }
  }
}
