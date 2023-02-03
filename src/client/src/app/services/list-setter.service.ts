import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { collection} from '../types/listTypes';
import { Observable, firstValueFrom} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ListSetterService {

  constructor(private http:HttpClient) { }
  async set(collection: collection){
    const collection$ = await firstValueFrom(this.http.post<collection>(environment.API_URL + "/listSetter", JSON.stringify(collection), {withCredentials: true}))
  }
}
