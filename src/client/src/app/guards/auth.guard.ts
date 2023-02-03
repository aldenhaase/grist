import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, RouterStateSnapshot, UrlTree } from '@angular/router';
import { Router } from '@angular/router';
import { Observable, of } from 'rxjs';
import { catchError, map } from 'rxjs/operators';
import { AuthStatusService } from '../services/auth-status.service';
@Injectable({
  providedIn: 'root'
})



export class AuthGuard implements CanActivate {
  constructor(private authenticator: AuthStatusService, private router: Router){};
  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): Observable<boolean> {
    return this.authenticator.checkForSessionCookie().pipe(map((authenticated: boolean)=>{
            if(!authenticated){
              this.router.navigate(['login']); 
              return false
            }
            return true;
            }), 
            catchError((error)=>{
              this.router.navigate(['login']);
              return of(false)
            }));
  }
}
