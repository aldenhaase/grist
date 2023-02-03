import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})

export class AuthStatusService {
  constructor(private http:HttpClient) { }

  checkForSessionCookie(): Observable<boolean>{
    return this.http.get<boolean>(environment.API_URL + '/cookieAuthenticator',{withCredentials: true});
  }
}
