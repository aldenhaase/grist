import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { collection, list } from '../types/listTypes';
import { Observable, firstValueFrom} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ListDeleterService {

  constructor(private http:HttpClient) { }
  async delete(list: list){
    await firstValueFrom(this.http.post<list>(environment.API_URL + "/listDeleter", JSON.stringify(list), {withCredentials: true}))
  }
}
