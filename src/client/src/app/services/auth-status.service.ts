import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { sessionAuthenticationResponse } from '../types/sessionAuthentication';

@Injectable({
  providedIn: 'root'
})

export class AuthStatusService {
  constructor(private http:HttpClient) { }

  checkForSessionCookie(): Observable<sessionAuthenticationResponse>{
    return this.http.get<sessionAuthenticationResponse>(environment.API_URL + '/cookieAuthenticator',{withCredentials: true});
  }
}
