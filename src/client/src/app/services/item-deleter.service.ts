import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { collection, item, list } from '../types/listTypes';
import { Observable, firstValueFrom} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ItemDeleterService {

  constructor(private http:HttpClient) { }
  async delete(item: item[]){
    await  firstValueFrom(this.http.post<item>(environment.API_URL + "/itemDeleter", JSON.stringify(item), {withCredentials: true}))
  }
}
