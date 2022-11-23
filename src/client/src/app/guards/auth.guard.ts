import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, RouterStateSnapshot, UrlTree } from '@angular/router';
import { Router } from '@angular/router';
import { Observable, of } from 'rxjs';
import { catchError, map } from 'rxjs/operators';
import { AuthStatusService } from '../services/auth-status.service';
import { sessionAuthenticationResponse } from '../types/sessionAuthentication';
@Injectable({
  providedIn: 'root'
})



export class AuthGuard implements CanActivate {
  constructor(private authenticator: AuthStatusService, private router: Router){};
  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): Observable<boolean> {
    return this.authenticator.checkForSessionCookie().pipe(map((response: sessionAuthenticationResponse)=>{
            if(response.Authenticated){
              return true;
            }
            this.router.navigate(['login']); 
            return false;
            }), 
            catchError((error)=>{
              this.router.navigate(['login']);
              return of(false)
            }));
  }
}
