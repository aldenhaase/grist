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
      local.lists = remote.lists
  }
}
