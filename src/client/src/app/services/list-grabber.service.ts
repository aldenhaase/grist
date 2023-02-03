import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { collection } from '../types/listTypes';

@Injectable({
  providedIn: 'root'
})
export class ListGrabberService {

  constructor(private http:HttpClient) { }
  grab(){
    return this.http.get<collection>(environment.API_URL + "/listGrabber", {withCredentials: true})
  }
}
